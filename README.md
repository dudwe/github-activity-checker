# Github Activity Checker

A command-line tool to check the activity of a GitHub user by retrieving their activity data.

## Prerequisites

- Go 1.23.7+ installed on your machine.
- An internet connection to access GitHub API.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/dudwe/github-activity-checker.git
    cd github-activity-checker
    ```

2. Build the application:
    ```bash
    go build -o github-activity-checker
    ```

3. Now, you can run the `github-activity-checker` binary.

## Usage

Run the command with the `--user` flag to specify the GitHub username of the person whose activity you want to check:

```bash
./github-activity-checker --user <github-username>
