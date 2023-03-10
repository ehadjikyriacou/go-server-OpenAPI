/*
 * PetSitter API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Job struct {

	Id int32 `json:"id,omitempty"`

	CreatorUserId int32 `json:"creator_user_id,omitempty"`

	StartTime string `json:"start_time,omitempty"`

	EndTime string `json:"end_time,omitempty"`

	Activity string `json:"activity,omitempty"`

	Dog *Dog `json:"dog,omitempty"`
}
