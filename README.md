# Harvest CLI

Integrates with JIRA to automate time recording against tickets.

## Installation

1. Install Go
2. Download the [pre-built JIRA CLI binary for your system](https://github.com/Netflix-Skunkworks/go-jira/releases)
3. Rename the binary to `jira` and copy it to `/usr/local/bin/`
4. `go get github.com/stuart-xyz/harvest-cli && GOBIN=/usr/local/bin/ go install $GOPATH/src/github.com/stuart-xyz/harvest-cli/`
5. Generate a [personal access token in Harvest](https://id.getharvest.com/developers)
6. Edit `$HOME/.jira.d/config.yml` and add the entries `harvest-personal-access-token: $TOKEN` and `harvest-account-id: $ACCOUNT_ID`
7. Add Harvest task list to `$HOME/.jira.d/harvest-task-list.csv` with the following CSV format:

    ```
    <project_id>,<task_id>,<task_description>
    ```

## Usage

```
Usage:
  harvest log <ticket_ref> <hours>
  harvest -h | --help
  harvest --version

Options:
  -h --help     Show this screen.
  --version     Show version.
```
