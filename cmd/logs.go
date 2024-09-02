// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ########################################################################################

package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type UserAction struct {
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
	Resource  string `json:"resource"`
}

type UserLog struct {
	Username string       `json:"username"`
	Actions  []UserAction `json:"actions"`
}

func LogUserAction(action, resource string) error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home directory: %w", err)
	}

	logDir := filepath.Join(homeDir, ".vaultify", "logs", "users")
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return fmt.Errorf("error creating user logs directory: %w", err)
	}

	logFile := filepath.Join(logDir, currentUser.Username+".json")

	var userLog UserLog

	if _, err := os.Stat(logFile); err == nil {
		fileContent, err := os.ReadFile(logFile)
		if err != nil {
			return fmt.Errorf("error reading user log file: %w", err)
		}
		if err := json.Unmarshal(fileContent, &userLog); err != nil {
			return fmt.Errorf("error parsing user log file: %w", err)
		}
	} else {
		userLog = UserLog{Username: currentUser.Username}
	}

	newAction := UserAction{
		Timestamp: time.Now().Format(time.RFC3339),
		Action:    action,
		Resource:  resource,
	}
	userLog.Actions = append(userLog.Actions, newAction)

	fileContent, err := json.MarshalIndent(userLog, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling user log data: %w", err)
	}

	if err := os.WriteFile(logFile, fileContent, 0600); err != nil {
		return fmt.Errorf("error writing user log file: %w", err)
	}

	return nil
}

func Logs(args []string) {
	logsCmd := flag.NewFlagSet("logs", flag.ExitOnError)
	userFlag := logsCmd.String("user", "", "Username to view logs for")
	fromDate := logsCmd.String("from", "", "Start date for filtering (YYYY-MM-DD)")
	toDate := logsCmd.String("to", "", "End date for filtering (YYYY-MM-DD)")
	actionType := logsCmd.String("action", "", "Filter by action type")
	limit := logsCmd.Int("limit", 0, "Limit the number of entries displayed")
	jsonOutput := logsCmd.Bool("json", false, "Output in JSON format")
	searchPattern := logsCmd.String("search", "", "Search for a specific pattern in resources")

	if err := logsCmd.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	if *userFlag == "" {
		fmt.Println("Please specify a username with --user flag")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}

	logFile := filepath.Join(homeDir, ".vaultify", "logs", "users", *userFlag+".json")

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		fmt.Printf("No logs found for user: %s\n", *userFlag)
		return
	}

	fileContent, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("Error reading user log file:", err)
		return
	}

	var userLog UserLog
	if err := json.Unmarshal(fileContent, &userLog); err != nil {
		fmt.Println("Error parsing user log file:", err)
		return
	}

	filteredActions, err := filterUserActions(userLog.Actions, *fromDate, *toDate, *actionType, *searchPattern)
	if err != nil {
		fmt.Println("Error filtering user actions:", err)
		return
	}

	if *limit > 0 && *limit < len(filteredActions) {
		filteredActions = filteredActions[:*limit]
	}
	if *jsonOutput {
		outputJSONData(userLog.Username, filteredActions)
	} else {
		outputHumanReadable(userLog.Username, filteredActions)
	}
}

func filterUserActions(actions []UserAction, fromDate, toDate, actionType, searchPattern string) ([]UserAction, error) {
	var filteredActions []UserAction

	from, err := parseDate(fromDate)
	if err != nil {
		return nil, fmt.Errorf("invalid from date: %w", err)
	}

	to, err := parseDate(toDate)
	if err != nil {
		return nil, fmt.Errorf("invalid to date: %w", err)
	}

	for _, action := range actions {
		actionTime, err := time.Parse(time.RFC3339, action.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp in log entry: %w", err)
		}

		if (from.IsZero() || actionTime.After(from)) &&
			(to.IsZero() || actionTime.Before(to)) &&
			(actionType == "" || action.Action == actionType) &&
			(searchPattern == "" || strings.Contains(action.Resource, searchPattern)) {
			filteredActions = append(filteredActions, action)
		}
	}

	return filteredActions, nil
}

func outputJSONData(username string, actions []UserAction) {
	output := struct {
		Username string       `json:"username"`
		Actions  []UserAction `json:"actions"`
		Total    int          `json:"total_actions"`
	}{
		Username: username,
		Actions:  actions,
		Total:    len(actions),
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func outputHumanReadable(username string, actions []UserAction) {
	fmt.Printf("User Activity Log for %s:\n", username)
	for _, action := range actions {
		fmt.Printf("%s - %s: %s\n", action.Timestamp, action.Action, action.Resource)
	}
	fmt.Printf("\nTotal actions: %d\n", len(actions))
}
