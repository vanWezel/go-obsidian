Ideal use for in a canvas with linked notes!

# Features

## Gitlab
1. Fetches your starred projects open merge requests (assigned to you) into a file called `Merge Requests.md`
2. Fetches your starred projects open merge requests (where you are the reviewer) into a file called `Review Requests.md`

## Outlook (mac only)
Parses a local json file used by the Outlook widget on your mac into 2 files `🌞 Today.md` `🌞 Tomorrow.md`.

## Jira
Fetches your current tickets / issues that are asigned to you with a few statuses ("Blocked", "In Progress", "Testing & Acceptance")

## copy .env.example to .env
fill in the missing keys

### setup gitlab
Create a new personal access token
1. Go to: https://gitlab.com/-/user_settings/personal_access_tokens
2. Create new token with scopes: `read_api` + `read_repository`
3. Paste token as key `GITLAB_TOKEN` in .env

### setup jira
Create a new API token
1. Go to: https://id.atlassian.com/manage-profile/security/api-tokens
2. Setup the `.env` file
    - Paste your new token in `JIRA_TOKEN`
    - Paste your jira domain in `JIRA_DOMAIN`
    - Paste your jira username in `JIRA_USERNAME`

## setup cron schedule
```cronexp
*/5 8-20 * * * cd /{path} && ./main >/dev/null 2>&1
```
