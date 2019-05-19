# Harvest CLI

Integrates with Jira to automate time recording against issues.

## Installation

0. Install Go
1. Compile binary:

    ```
    go get github.com/stuart-xyz/harvest-cli
    cd $GOPATH/src/github.com/stuart-xyz/harvest-cli && go build -o $BIN/harvest
    ```

2. Generate a [Jira API token](https://id.atlassian.com)
3. Generate a [personal access token in Harvest](https://id.getharvest.com/developers)
4. Create `$HOME/.harvest.d/config.yml` and add the following:

    ```
    jira-endpoint: ... // e.g. https://<your_company>.atlassian.net
    jira-email: ...
    jira-api-token: ...
    harvest-personal-access-token: ...
    harvest-account-id: ...
    ```

5. Add Harvest task list to `$HOME/.harvest.d/task-list.csv` with the following CSV format:

    ```
    <project_id>,<task_id>,<task_description>
    ```
    
    You may need to ask your Harvest administrator to provide this list. Alternatively, [install HCL](https://github.com/zenhob/hcl) and run `hcl tasks` (however, this relies on the [now-deprecated V1 Harvest API](https://help.getharvest.com/api-v1/)). 

## Usage

```
Usage:
  harvest view [--yesterday]
  harvest log <hours> [--yesterday]
  harvest log <issue_ref> <hours> [--yesterday]
  harvest -h | --help
  harvest --version

Options:
  -h --help     Show this screen.
  --version     Show version.
```
