# Harvest CLI

Integrates with JIRA to automate time recording against tickets.

## Installation

1. Install Go
2. ```go get gopkg.in/Netflix-Skunkworks/go-jira.v1/cmd/jira && go install -o $GOBIN/jira $GOPATH/src/gopkg.in/Netflix-Skunkworks/go-jira.v1/cmd/jira/```
3. Generate a [JIRA API token](https://id.atlassian.com)
4. Edit `$HOME/.jira.d/config.yml` and add the following:

    ```
    endpoint: <your_jira_endpoint> // e.g. https://<your_company>.atlassian.net
    login: <your_email_address>
    password-source: keyring // add this to store your API token in your system keychain
    ```

5. ```go get github.com/stuart-xyz/harvest-cli && go install -o $GOBIN/harvest $GOPATH/src/github.com/stuart-xyz/harvest-cli/```
6. Generate a [personal access token in Harvest](https://id.getharvest.com/developers)
7. Edit `$HOME/.jira.d/config.yml` and add the following:

    ```
    harvest-personal-access-token: <token>
    harvest-account-id: <account_id>
    ```

8. Add Harvest task list to `$HOME/.jira.d/harvest-task-list.csv` with the following CSV format:

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
