/*
 * CharVstack-API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oldopenapi

type StoragePool struct {
	Name string `json:"name,omitempty"`

	TotalSize uint64 `json:"total_size,omitempty"`

	UsedSize uint64 `json:"used_size,omitempty"`

	Path string `json:"path,omitempty"`

	Status string `json:"status,omitempty"`
}
