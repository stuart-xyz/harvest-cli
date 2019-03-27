# Harvest CLI

Integrates with JIRA to automate time recording against tickets.

## Installation

1. Install JIRA CLI
2. Download + build source
3. Generate a personal access token in Harvest
4. Edit `$HOME/.jira.d/config.yml` and add the entries `harvest-personal-access-token: $TOKEN` and `harvest-account-id: $ACCOUNT_ID`
5. Add Harvest task list to `$HOME/.jira.d/harvest-task-list.csv` with the following CSV format:

  ```
  <project_id>,<task_id>,<task_description>
  ```

## Usage

```
Usage:
  harvest log <ticket_ref> <hours>...
  harvest -h | --help
  harvest --version

Options:
  -h --help     Show this screen.
  --version     Show version.
```
