## create gitlab token
- Go to: https://gitlab.com/-/user_settings/personal_access_tokens
- create new token with scopes: `read_api` + `read_repository`

## create a jira token
- go to: https://id.atlassian.com/manage-profile/security/api-tokens

## setup cron schedule
```cronexp
*/5 8-20 * * * cd /{path} && ./main >/dev/null 2>&1
```