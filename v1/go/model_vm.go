/*
 * CharVstack-API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Vm -
type Vm struct {
	Name string `json:"name,omitempty"`

	Metadata VmMetadata `json:"metadata,omitempty"`

	Memory int32 `json:"memory,omitempty"`

	Vcpu int32 `json:"vcpu,omitempty"`

	Devices VmDevices `json:"devices,omitempty"`
}
