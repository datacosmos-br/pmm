// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeNodeExporterReader is a Reader for the ChangeNodeExporter structure.
type ChangeNodeExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeNodeExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeNodeExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeNodeExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeNodeExporterOK creates a ChangeNodeExporterOK with default headers values
func NewChangeNodeExporterOK() *ChangeNodeExporterOK {
	return &ChangeNodeExporterOK{}
}

/* ChangeNodeExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeNodeExporterOK struct {
	Payload *ChangeNodeExporterOKBody
}

func (o *ChangeNodeExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeNodeExporter][%d] changeNodeExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeNodeExporterOK) GetPayload() *ChangeNodeExporterOKBody {
	return o.Payload
}

func (o *ChangeNodeExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeNodeExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeNodeExporterDefault creates a ChangeNodeExporterDefault with default headers values
func NewChangeNodeExporterDefault(code int) *ChangeNodeExporterDefault {
	return &ChangeNodeExporterDefault{
		_statusCode: code,
	}
}

/* ChangeNodeExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeNodeExporterDefault struct {
	_statusCode int

	Payload *ChangeNodeExporterDefaultBody
}

// Code gets the status code for the change node exporter default response
func (o *ChangeNodeExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeNodeExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeNodeExporter][%d] ChangeNodeExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeNodeExporterDefault) GetPayload() *ChangeNodeExporterDefaultBody {
	return o.Payload
}

func (o *ChangeNodeExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeNodeExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeNodeExporterBody change node exporter body
swagger:model ChangeNodeExporterBody
*/
type ChangeNodeExporterBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeNodeExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change node exporter body
func (o *ChangeNodeExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterBody) validateCommon(formats strfmt.Registry) error {
	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change node exporter body based on the context it is used
func (o *ChangeNodeExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterBody) contextValidateCommon(ctx context.Context, formats strfmt.Registry) error {
	if o.Common != nil {
		if err := o.Common.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterDefaultBody change node exporter default body
swagger:model ChangeNodeExporterDefaultBody
*/
type ChangeNodeExporterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeNodeExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change node exporter default body
func (o *ChangeNodeExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeNodeExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeNodeExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change node exporter default body based on the context it is used
func (o *ChangeNodeExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeNodeExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeNodeExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterDefaultBodyDetailsItems0 change node exporter default body details items0
swagger:model ChangeNodeExporterDefaultBodyDetailsItems0
*/
type ChangeNodeExporterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change node exporter default body details items0
func (o *ChangeNodeExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change node exporter default body details items0 based on context it is used
func (o *ChangeNodeExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterOKBody change node exporter OK body
swagger:model ChangeNodeExporterOKBody
*/
type ChangeNodeExporterOKBody struct {
	// node exporter
	NodeExporter *ChangeNodeExporterOKBodyNodeExporter `json:"node_exporter,omitempty"`
}

// Validate validates this change node exporter OK body
func (o *ChangeNodeExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterOKBody) validateNodeExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	if o.NodeExporter != nil {
		if err := o.NodeExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeNodeExporterOk" + "." + "node_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeNodeExporterOk" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change node exporter OK body based on the context it is used
func (o *ChangeNodeExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNodeExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterOKBody) contextValidateNodeExporter(ctx context.Context, formats strfmt.Registry) error {
	if o.NodeExporter != nil {
		if err := o.NodeExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeNodeExporterOk" + "." + "node_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeNodeExporterOk" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterOKBodyNodeExporter NodeExporter runs on Generic or Container Node and exposes its metrics.
swagger:model ChangeNodeExporterOKBodyNodeExporter
*/
type ChangeNodeExporterOKBodyNodeExporter struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this change node exporter OK body node exporter
func (o *ChangeNodeExporterOKBodyNodeExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum = append(changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeNodeExporterOKBodyNodeExporterStatusSTARTING captures enum value "STARTING"
	ChangeNodeExporterOKBodyNodeExporterStatusSTARTING string = "STARTING"

	// ChangeNodeExporterOKBodyNodeExporterStatusRUNNING captures enum value "RUNNING"
	ChangeNodeExporterOKBodyNodeExporterStatusRUNNING string = "RUNNING"

	// ChangeNodeExporterOKBodyNodeExporterStatusWAITING captures enum value "WAITING"
	ChangeNodeExporterOKBodyNodeExporterStatusWAITING string = "WAITING"

	// ChangeNodeExporterOKBodyNodeExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeNodeExporterOKBodyNodeExporterStatusSTOPPING string = "STOPPING"

	// ChangeNodeExporterOKBodyNodeExporterStatusDONE captures enum value "DONE"
	ChangeNodeExporterOKBodyNodeExporterStatusDONE string = "DONE"

	// ChangeNodeExporterOKBodyNodeExporterStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeNodeExporterOKBodyNodeExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeNodeExporterOKBodyNodeExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeNodeExporterOKBodyNodeExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeNodeExporterOk"+"."+"node_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change node exporter OK body node exporter based on context it is used
func (o *ChangeNodeExporterOKBodyNodeExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterOKBodyNodeExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterOKBodyNodeExporter) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterOKBodyNodeExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeNodeExporterParamsBodyCommon
*/
type ChangeNodeExporterParamsBodyCommon struct {
	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change node exporter params body common
func (o *ChangeNodeExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change node exporter params body common based on context it is used
func (o *ChangeNodeExporterParamsBodyCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
