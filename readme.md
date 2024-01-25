Ideal use for in a canvas with linked notes!

# Features
## Gitlab
1. Fetches your starred projects open merge requests (assigned to you) into a file called `Merge Requests.md`
2. Fetches your starred projects open merge requests (where you are the reviewer) into a file called `Review Requests.md`

### setup gitlab
Create a new personal access token
1. Go to: https://gitlab.com/-/user_settings/personal_access_tokens
2. Create new token with scopes: `read_api` + `read_repository`

## Outlook (mac only)
Parses a local json file used by the Outlook widget on your mac into 2 files `🌞 Today.md` `🌞 Tomorrow.md`.

# Setup cron schedule
replace {TOKEN} by your own
```cronexp
*/5 8-20 * * * export GITLAB_TOKEN='{token}' && /{path}/main -p={notes_folder} >/dev/null 2>&1
```
