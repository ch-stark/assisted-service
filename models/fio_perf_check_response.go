// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FioPerfCheckResponse fio perf check response
//
// swagger:model fio_perf_check_response
type FioPerfCheckResponse struct {

	// The 99th percentile of fdatasync durations in milliseconds.
	IoSyncDuration int64 `json:"io_sync_duration,omitempty"`

	// The device path.
	Path string `json:"path,omitempty"`
}

// Validate validates this fio perf check response
func (m *FioPerfCheckResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FioPerfCheckResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FioPerfCheckResponse) UnmarshalBinary(b []byte) error {
	var res FioPerfCheckResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
