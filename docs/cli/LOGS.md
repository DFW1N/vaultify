<!-- // ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ######################################################################################## -->

# Vaultify CLI Command - Logs

## Overview
The `logs` command in the Vaultify CLI allows users to view and analyze user-specific action logs. It provides a flexible way to retrieve, filter, and format the history of operations performed by specific users, such as pushing secrets or state files, pulling data, or retrieving secrets.

## Functionality
- **Retrieve User Logs:** The function retrieves and displays a list of actions performed by a specific user.
- **Filter Options:** Users can filter the logs by date range, action type, and resource pattern.
- **Flexible Output:** The command supports both human-readable and JSON output formats.
- **Limit Results:** Users can specify the number of entries to display.
- **Search Capability:** Allows searching for specific patterns in resources.

## Usage
To use the `logs` command in the Vaultify CLI, run the following command:

```bash
vaultify logs [options]
```

### Options
- `--user <username>`: Username to view logs for (required)
- `--from <date>`: Start date for filtering (format: YYYY-MM-DD)
- `--to <date>`: End date for filtering (format: YYYY-MM-DD)
- `--action <type>`: Filter by action type (e.g., push, pull, get)
- `--limit <n>`: Limit the number of entries displayed
- `--json`: Output in JSON format
- `--search <pattern>`: Search for a specific pattern in resources

## Example Use Cases

1. View all logs for a specific user:
```bash
vaultify logs --user sacha
```

2. View logs for a specific date range:
```bash
vaultify logs --user sacha --from 2023-05-01 --to 2023-05-31
```

3. View only 'push' actions for a user:
```bash
vaultify logs --user sacha --action push
```

4. Search for actions related to a specific project:
```bash
vaultify logs --user sacha --search "myproject"
```

5. Limit the output to 10 entries in JSON format:
```bash
vaultify logs --user sacha --limit 10 --json
```

## Output Examples

### Human-readable format:

```bash
User Activity Log for sacha:
2024-09-01T00:01:34+10:00 - pull: vault:kv/data/vaultify/vaultify/default_terraform.tfstate
2024-09-01T00:03:47+10:00 - get: vault:kv/data/secrets/my/test/mysecret

Total actions: 2
```

### JSON format:

```json
{
  "username": "sacha",
  "actions": [
    {
      "timestamp": "2024-09-01T00:01:34+10:00",
      "action": "pull",
      "resource": "vault:kv/data/vaultify/vaultify/default_terraform.tfstate"
    },
    {
      "timestamp": "2024-09-01T00:03:47+10:00",
      "action": "get",
      "resource": "vault:kv/data/secrets/my/test/mysecret"
    }
  ],
  "total_actions": 2
}
```

## Log File Location
The user log entries are stored in JSON files located at:

`~/.vaultify/logs/users/<username>.json`

These files are automatically created and updated when users perform actions with Vaultify.

## Error Handling
- If the specified user log file doesn't exist, it will inform the user that no logs were found.
- Any errors in reading or parsing the log file will be reported to the user.
- Invalid date formats or ranges will result in an error message.

## Auditing and Compliance
The logs command provides several features useful for auditing and compliance purposes:
- User-specific logging allows for tracking individual user actions.
- Date range filtering allows for reviewing actions within specific time periods.
- Action type filtering helps in tracking specific operations (e.g., all 'push' actions).
- JSON output facilitates integration with other tools and scripts for further analysis.
- Search functionality helps in quickly locating actions related to specific projects or resources.

## Best Practices
1. Regularly review user logs to ensure all actions are authorized and expected.
2. Use the date range filtering for periodic (e.g., weekly, monthly) audits of user activities.
3. Utilize the JSON output for integrating log data with other analysis or reporting tools.
4. Use the search functionality to quickly identify actions related to sensitive projects or environments.
5. Combine filters to create targeted reports for specific audit requirements or user activity patterns.

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |