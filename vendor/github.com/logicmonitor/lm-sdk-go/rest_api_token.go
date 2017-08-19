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

type RestApiToken struct {
	AccessId string `json:"accessId,omitempty"`

	AdminName string `json:"adminName,omitempty"`

	Note string `json:"note,omitempty"`

	LastUsedOn int64 `json:"lastUsedOn,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	Roles []string `json:"roles,omitempty"`

	AdminId int32 `json:"adminId,omitempty"`

	Id int32 `json:"id,omitempty"`

	CreatedOn int64 `json:"createdOn,omitempty"`

	Status int32 `json:"status,omitempty"`
}
