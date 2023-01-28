/*
 * PetSitter API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type User struct {
	Id string `json:"id,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	FullName string `json:"full_name,omitempty"`

	Roles []string `json:"roles,omitempty"`
}
