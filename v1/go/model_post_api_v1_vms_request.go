/*
 * CharVstack-API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostApiV1VmsRequest struct {
	Name string `json:"name"`

	Memory int32 `json:"memory"`

	Vcpu int32 `json:"vcpu"`
}
