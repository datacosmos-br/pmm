// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddHAProxyServiceReader is a Reader for the AddHAProxyService structure.
type AddHAProxyServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHAProxyServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddHAProxyServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddHAProxyServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddHAProxyServiceOK creates a AddHAProxyServiceOK with default headers values
func NewAddHAProxyServiceOK() *AddHAProxyServiceOK {
	return &AddHAProxyServiceOK{}
}

/* AddHAProxyServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddHAProxyServiceOK struct {
	Payload *AddHAProxyServiceOKBody
}

func (o *AddHAProxyServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddHAProxyService][%d] addHaProxyServiceOk  %+v", 200, o.Payload)
}

func (o *AddHAProxyServiceOK) GetPayload() *AddHAProxyServiceOKBody {
	return o.Payload
}

func (o *AddHAProxyServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddHAProxyServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHAProxyServiceDefault creates a AddHAProxyServiceDefault with default headers values
func NewAddHAProxyServiceDefault(code int) *AddHAProxyServiceDefault {
	return &AddHAProxyServiceDefault{
		_statusCode: code,
	}
}

/* AddHAProxyServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddHAProxyServiceDefault struct {
	_statusCode int

	Payload *AddHAProxyServiceDefaultBody
}

// Code gets the status code for the add HA proxy service default response
func (o *AddHAProxyServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddHAProxyServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddHAProxyService][%d] AddHAProxyService default  %+v", o._statusCode, o.Payload)
}

func (o *AddHAProxyServiceDefault) GetPayload() *AddHAProxyServiceDefaultBody {
	return o.Payload
}

func (o *AddHAProxyServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddHAProxyServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddHAProxyServiceBody add HA proxy service body
swagger:model AddHAProxyServiceBody
*/
type AddHAProxyServiceBody struct {
	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add HA proxy service body
func (o *AddHAProxyServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add HA proxy service body based on context it is used
func (o *AddHAProxyServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyServiceBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyServiceDefaultBody add HA proxy service default body
swagger:model AddHAProxyServiceDefaultBody
*/
type AddHAProxyServiceDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddHAProxyServiceDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add HA proxy service default body
func (o *AddHAProxyServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyServiceDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddHAProxyService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddHAProxyService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add HA proxy service default body based on the context it is used
func (o *AddHAProxyServiceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyServiceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddHAProxyService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddHAProxyService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyServiceDefaultBodyDetailsItems0 add HA proxy service default body details items0
swagger:model AddHAProxyServiceDefaultBodyDetailsItems0
*/
type AddHAProxyServiceDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add HA proxy service default body details items0
func (o *AddHAProxyServiceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add HA proxy service default body details items0 based on context it is used
func (o *AddHAProxyServiceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyServiceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyServiceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddHAProxyServiceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyServiceOKBody add HA proxy service OK body
swagger:model AddHAProxyServiceOKBody
*/
type AddHAProxyServiceOKBody struct {
	// haproxy
	Haproxy *AddHAProxyServiceOKBodyHaproxy `json:"haproxy,omitempty"`
}

// Validate validates this add HA proxy service OK body
func (o *AddHAProxyServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyServiceOKBody) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addHaProxyServiceOk" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addHaProxyServiceOk" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add HA proxy service OK body based on the context it is used
func (o *AddHAProxyServiceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyServiceOKBody) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {
	if o.Haproxy != nil {
		if err := o.Haproxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addHaProxyServiceOk" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addHaProxyServiceOk" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyServiceOKBodyHaproxy HAProxyService represents a generic HAProxy service instance.
swagger:model AddHAProxyServiceOKBodyHaproxy
*/
type AddHAProxyServiceOKBodyHaproxy struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this service instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add HA proxy service OK body haproxy
func (o *AddHAProxyServiceOKBodyHaproxy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add HA proxy service OK body haproxy based on context it is used
func (o *AddHAProxyServiceOKBodyHaproxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyServiceOKBodyHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyServiceOKBodyHaproxy) UnmarshalBinary(b []byte) error {
	var res AddHAProxyServiceOKBodyHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
