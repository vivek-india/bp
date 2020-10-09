/*
 * Dhanraj Hardware Stores
 *
 * Dhanraj Hardware Stores APIs
 *
 * API version: 1.0.0
 * Contact: viveks.singh@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Person struct {

	// Unique ID of the Business Partner
	Id int32 `json:"id,omitempty"`

	// ID of the Business Partner provided by admin
	Nickname string `json:"nickname"`

	// Name of the Business Partner
	Name string `json:"name"`

	Contact []string `json:"contact,omitempty"`

	// Address of the Business Partner
	Address string `json:"address,omitempty"`

	// City/Village of the Business Partner
	City string `json:"city,omitempty"`
}
