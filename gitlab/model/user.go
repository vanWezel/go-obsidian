package model

import "time"

type User struct {
	Id              int         `json:"id"`
	Username        string      `json:"username"`
	Name            string      `json:"name"`
	State           string      `json:"state"`
	Locked          bool        `json:"locked"`
	AvatarUrl       string      `json:"avatar_url"`
	WebUrl          string      `json:"web_url"`
	CreatedAt       time.Time   `json:"created_at"`
	Bio             string      `json:"bio"`
	Location        string      `json:"location"`
	PublicEmail     string      `json:"public_email"`
	Skype           string      `json:"skype"`
	Linkedin        string      `json:"linkedin"`
	Twitter         string      `json:"twitter"`
	Discord         string      `json:"discord"`
	WebsiteUrl      string      `json:"website_url"`
	Organization    string      `json:"organization"`
	JobTitle        string      `json:"job_title"`
	Pronouns        string      `json:"pronouns"`
	Bot             bool        `json:"bot"`
	WorkInformation interface{} `json:"work_information"`
	LocalTime       interface{} `json:"local_time"`
	LastSignInAt    time.Time   `json:"last_sign_in_at"`
	ConfirmedAt     time.Time   `json:"confirmed_at"`
	LastActivityOn  string      `json:"last_activity_on"`
	Email           string      `json:"email"`
	ThemeId         int         `json:"theme_id"`
	ColorSchemeId   int         `json:"color_scheme_id"`
	ProjectsLimit   int         `json:"projects_limit"`
	CurrentSignInAt time.Time   `json:"current_sign_in_at"`
	Identities      []struct {
		Provider       string `json:"provider"`
		ExternUid      string `json:"extern_uid"`
		SamlProviderId int    `json:"saml_provider_id"`
	} `json:"identities"`
	CanCreateGroup                 bool          `json:"can_create_group"`
	CanCreateProject               bool          `json:"can_create_project"`
	TwoFactorEnabled               bool          `json:"two_factor_enabled"`
	External                       bool          `json:"external"`
	PrivateProfile                 bool          `json:"private_profile"`
	CommitEmail                    string        `json:"commit_email"`
	SharedRunnersMinutesLimit      interface{}   `json:"shared_runners_minutes_limit"`
	ExtraSharedRunnersMinutesLimit interface{}   `json:"extra_shared_runners_minutes_limit"`
	ScimIdentities                 []interface{} `json:"scim_identities"`
}
