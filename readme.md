## create gitlab token
- Go to: https://gitlab.com/-/user_settings/personal_access_tokens
- create new token with scopes: `read_api` + `read_repository`

## setup cron schedule
replace {TOKEN} by your own
```cronexp
*/5 8-20 * * * export GITLAB_TOKEN='{token}' && /{path}/main -p={notes_folder} >/dev/null 2>&1
```