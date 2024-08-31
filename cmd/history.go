package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type HistoryEntry struct {
	Timestamp      string `json:"timestamp"`
	PassphraseHash string `json:"passphrase_hash"`
	Action         string `json:"action"`
	Storage        string `json:"storage"`
}

type HistoryData struct {
	Entries []HistoryEntry `json:"entries"`
}

func LogHistory(action, storage string) error {
	passphrase := os.Getenv("VAULTIFY_PASSPHRASE")
	if passphrase == "" {
		return fmt.Errorf("VAULTIFY_PASSPHRASE environment variable is not set")
	}

	hasher := sha256.New()
	hasher.Write([]byte(passphrase))
	passphraseHash := hex.EncodeToString(hasher.Sum(nil))

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home directory: %w", err)
	}

	vaultifyDir := filepath.Join(homeDir, ".vaultify")
	logsDir := filepath.Join(vaultifyDir, "logs")
	if err := os.MkdirAll(logsDir, 0700); err != nil {
		return fmt.Errorf("error creating .vaultify/logs directory: %w", err)
	}

	historyFile := filepath.Join(logsDir, "history.json")

	var historyData HistoryData

	if _, err := os.Stat(historyFile); err == nil {
		fileContent, err := os.ReadFile(historyFile)
		if err != nil {
			return fmt.Errorf("error reading history file: %w", err)
		}
		if err := json.Unmarshal(fileContent, &historyData); err != nil {
			return fmt.Errorf("error parsing history file: %w", err)
		}
	}

	newEntry := HistoryEntry{
		Timestamp:      time.Now().Format(time.RFC3339),
		PassphraseHash: passphraseHash,
		Action:         action,
		Storage:        storage,
	}
	historyData.Entries = append(historyData.Entries, newEntry)

	fileContent, err := json.MarshalIndent(historyData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling history data: %w", err)
	}

	if err := os.WriteFile(historyFile, fileContent, 0600); err != nil {
		return fmt.Errorf("error writing history file: %w", err)
	}

	return nil
}

func ViewHistory(args []string) {
	historyCmd := flag.NewFlagSet("history", flag.ExitOnError)
	fromDate := historyCmd.String("from", "", "Start date for filtering (YYYY-MM-DD)")
	toDate := historyCmd.String("to", "", "End date for filtering (YYYY-MM-DD)")
	actionType := historyCmd.String("action", "", "Filter by action type")
	limit := historyCmd.Int("limit", 0, "Limit the number of entries displayed")
	jsonOutput := historyCmd.Bool("json", false, "Output in JSON format")
	searchPattern := historyCmd.String("search", "", "Search for a specific pattern in storage locations")
	outputFile := historyCmd.String("output", "", "Write output to a file")

	if err := historyCmd.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	passphrase := os.Getenv("VAULTIFY_PASSPHRASE")
	if passphrase == "" {
		fmt.Println("VAULTIFY_PASSPHRASE environment variable is not set")
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(passphrase))
	passphraseHash := hex.EncodeToString(hasher.Sum(nil))

	historyData, err := readHistoryFile()
	if err != nil {
		fmt.Println("Error reading history file:", err)
		return
	}

	filteredHistory, err := filterHistory(historyData, passphraseHash, *fromDate, *toDate, *actionType, *searchPattern)
	if err != nil {
		fmt.Println("Error filtering history:", err)
		return
	}

	if *limit > 0 && *limit < len(filteredHistory) {
		filteredHistory = filteredHistory[:*limit]
	}

	var output string
	if *jsonOutput {
		output = formatJSON(filteredHistory)
	} else {
		output = formatHumanReadable(filteredHistory)
	}

	if *outputFile != "" {
		if err := writeToFile(*outputFile, output); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Printf("History written to file: %s\n", *outputFile)
	} else {
		fmt.Println(output)
	}
}

func readHistoryFile() (HistoryData, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return HistoryData{}, fmt.Errorf("error getting user home directory: %w", err)
	}

	historyFile := filepath.Join(homeDir, ".vaultify", "logs", "history.json")
	content, err := os.ReadFile(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return HistoryData{}, nil
		}
		return HistoryData{}, fmt.Errorf("error reading history file: %w", err)
	}

	var historyData HistoryData
	if err := json.Unmarshal(content, &historyData); err != nil {
		return HistoryData{}, fmt.Errorf("error parsing history file: %w", err)
	}

	return historyData, nil
}

func filterHistory(historyData HistoryData, passphraseHash, fromDate, toDate, actionType, searchPattern string) ([]HistoryEntry, error) {
	var filteredEntries []HistoryEntry

	fromTime, err := parseDate(fromDate)
	if err != nil {
		return nil, fmt.Errorf("invalid from date: %w", err)
	}

	toTime, err := parseDate(toDate)
	if err != nil {
		return nil, fmt.Errorf("invalid to date: %w", err)
	}

	for _, entry := range historyData.Entries {
		if entry.PassphraseHash != passphraseHash {
			continue
		}

		entryTime, err := time.Parse(time.RFC3339, entry.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp in history entry: %w", err)
		}

		if (fromTime.IsZero() || entryTime.After(fromTime)) &&
			(toTime.IsZero() || entryTime.Before(toTime)) &&
			(actionType == "" || entry.Action == actionType) &&
			(searchPattern == "" || strings.Contains(entry.Storage, searchPattern)) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries, nil
}

func parseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", date)
}

func formatJSON(entries []HistoryEntry) string {
	type jsonOutput struct {
		History      []HistoryEntry `json:"history"`
		TotalActions int            `json:"total_actions"`
	}

	output := jsonOutput{
		History:      entries,
		TotalActions: len(entries),
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting JSON: %v", err)
	}

	return string(jsonData)
}

func formatHumanReadable(entries []HistoryEntry) string {
	var sb strings.Builder
	sb.WriteString("Action History:\n")
	for _, entry := range entries {
		sb.WriteString(fmt.Sprintf("%s - %s: %s\n", entry.Timestamp, entry.Action, entry.Storage))
	}
	sb.WriteString(fmt.Sprintf("\nTotal actions: %d\n", len(entries)))
	return sb.String()
}

func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0600)
}
