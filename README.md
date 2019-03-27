# Harvest CLI

Integrates with JIRA to automate time recording against tickets.

## Installation

1. Install Go
2. Download the [pre-built JIRA CLI binary for your system](https://github.com/Netflix-Skunkworks/go-jira/releases)
3. Rename the binary to `jira` and copy it to `/usr/local/bin/`
4. Generate a [JIRA API token](https://id.atlassian.com)
5. Edit `$HOME/.jira.d/config.yml` and add the following:

    ```
    endpoint: <your_jira_endpoint> // e.g. https://<your_company>.atlassian.net
    login: <your_email_address>
    password-source: keyring // add this to store your API token in your system keychain
    ```

6. `go get github.com/stuart-xyz/harvest-cli && GOBIN=/usr/local/bin/ go install $GOPATH/src/github.com/stuart-xyz/harvest-cli/`
7. Generate a [personal access token in Harvest](https://id.getharvest.com/developers)
8. Edit `$HOME/.jira.d/config.yml` and add the following:

    ```
    harvest-personal-access-token: <token>
    harvest-account-id: <account_id>
    ```

9. Add Harvest task list to `$HOME/.jira.d/harvest-task-list.csv` with the following CSV format:

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
