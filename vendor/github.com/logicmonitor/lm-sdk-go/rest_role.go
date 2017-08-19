/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type RestRole struct {
	EnableRemoteSessionInCompanyLevel bool `json:"enableRemoteSessionInCompanyLevel,omitempty"`

	CustomHelpLabel string `json:"customHelpLabel,omitempty"`

	CustomHelpURL string `json:"customHelpURL,omitempty"`

	Privileges []RestPrivilege `json:"privileges"`

	AssociatedUserCount int32 `json:"associatedUserCount,omitempty"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Id int32 `json:"id,omitempty"`

	TwoFARequired bool `json:"twoFARequired,omitempty"`

	RequireEULA bool `json:"requireEULA,omitempty"`
}
