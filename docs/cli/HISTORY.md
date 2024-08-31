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

# Vaultify CLI Command - History

## Overview
The `history` command in the Vaultify CLI allows users to view and analyze a log of their actions performed with Vaultify. It provides a flexible way to retrieve, filter, and format the history of operations such as pushing secrets or state files to Vault or other storage backends, and injecting secrets into Vault.

## Functionality
- **Retrieve Action History:** The function retrieves and displays a list of actions performed by the user, filtered by the current passphrase.
- **Filter Options:** Users can filter the history by date range, action type, and storage location.
- **Flexible Output:** The command supports both human-readable and JSON output formats.
- **Limit Results:** Users can specify the number of entries to display.
- **Search Capability:** Allows searching for specific storage locations or patterns.

## Usage
To use the `history` command in the Vaultify CLI, run the following command:

```bash
vaultify history [options]
```

### Options
- `--from <date>`: Start date for filtering (format: YYYY-MM-DD)
- `--to <date>`: End date for filtering (format: YYYY-MM-DD)
- `--action <type>`: Filter by action type (e.g., push, inject)
- `--limit <n>`: Limit the number of entries displayed
- `--json`: Output in JSON format
- `--search <pattern>`: Search for a specific pattern in storage locations
- `--output <file>`: Write the output to a file instead of stdout

## Example Use Cases

1. View all history:

```bash
vaultify history
```
> This command displays all history entries in human-readable format.

---

2. View history in JSON format:
```bash
vaultify history --json
```
> This command outputs the entire history in JSON format, useful for parsing with other tools.

---

3. View history for a specific date range:
```bash
vaultify history --from 2023-04-01 --to 2023-04-30
```
> This command shows all actions performed between April 1st and April 30th, 2023.

---

4. View last 10 'push' actions:
```bash
vaultify history --action push --limit 10
```
> This command displays the 10 most recent 'push' actions.

---

5. Search for actions related to a specific project:
```bash
vaultify history --search "vault:kv/data/vaultify/vaultify/default_terraform.tfstate"
```
> This command shows all actions where the storage location contains the string "vault:...".

---

6. Export all 'inject' actions to a JSON file:
```bash
vaultify history --action inject --json --output inject_history.json
```
> This command filters for 'inject' actions, formats the output as JSON, and saves it to a file named 'inject_history.json'.

---

7. View actions from the last week:
```bash
vaultify history --from $(date -d "7 days ago" +%Y-%m-%d)
```
> This command shows all actions from the last 7 days.

---

8. Combine multiple filters:
```bash
vaultify history --from 2023-05-01 --to 2023-05-31 --action push --search "production" --json
```
> This command filters for 'push' actions in May 2023 related to "production", and outputs in JSON format.


## Output
The command will display output similar to the following:

```bash
Action History:
2023-04-20T10:30:00Z - push: vault:/secret/myproject/dev_terraform.tfstate
2023-04-20T11:15:00Z - inject: vault:/secret/myapp/api_key
2023-04-20T12:00:00Z - push: azure_storage:mystorageaccount
Total actions for this passphrase: 3
```

### JSON File Entry Structure
`~/.vaultify/logs/history.json`
```json
{
  "history": [
    {
      "timestamp": "2024-08-31T22:39:54+10:00",
      "passphrase_hash": "de17504111a7a61ccc82c0ed46fb61b8bcd29b24a106a969a9e42aad7eb29488",
      "action": "inject",
      "storage": "kv/data/secrets/foo/bar/tfc_token"
    },
    {
      "timestamp": "2024-08-31T23:17:38+10:00",
      "passphrase_hash": "de17504111a7a61ccc82c0ed46fb61b8bcd29b24a106a969a9e42aad7eb29488",
      "action": "push",
      "storage": "vault:kv/data/vaultify/vaultify/default_terraform.tfstate"
    }
  ],
  "total_actions": 2
}
```

## History File Location
The history entries are stored in a JSON file located at:

`~/.vaultify/logs/history.json`

This file is automatically created and updated when you perform actions with Vaultify.

## Security Note
The history file uses a hash of your passphrase to associate actions with a particular user. The actual passphrase is never stored. Each user with a different passphrase will only see their own actions in the history.

## Environment Variables
- `VAULTIFY_PASSPHRASE`: Must be set to view the history. This environment variable is used to generate the passphrase hash for filtering the history entries.

## Error Handling
- If the `VAULTIFY_PASSPHRASE` is not set, the command will display an error message.
- If the history file doesn't exist or is empty, it will inform the user that no history was found.
- Any errors in reading or parsing the history file will be reported to the user.
- Invalid date formats or ranges will result in an error message.

## Auditing and Compliance
The enhanced history command provides several features useful for auditing and compliance purposes:
- Date range filtering allows for reviewing actions within specific time periods.
- Action type filtering helps in tracking specific operations (e.g., all 'push' actions).
- JSON output facilitates integration with other tools and scripts for further analysis.
- The ability to export to a file allows for maintaining separate audit logs.
- Search functionality helps in quickly locating actions related to specific projects or resources.

## Best Practices
1. Regularly review the history to ensure all actions are authorized and expected.
2. Use the date range filtering for periodic (e.g., weekly, monthly) audits.
3. Export history to JSON files for long-term record keeping and analysis.
4. Use the search functionality to quickly identify actions related to sensitive projects or environments.
5. Combine filters to create targeted reports for specific audit requirements.

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |