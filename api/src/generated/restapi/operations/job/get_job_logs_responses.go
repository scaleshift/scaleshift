// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// GetJobLogsOKCode is the HTTP code returned for type GetJobLogsOK
const GetJobLogsOKCode int = 200

/*GetJobLogsOK OK

swagger:response getJobLogsOK
*/
type GetJobLogsOK struct {

	/*
	  In: Body
	*/
	Payload *models.JobLogs `json:"body,omitempty"`
}

// NewGetJobLogsOK creates GetJobLogsOK with default headers values
func NewGetJobLogsOK() *GetJobLogsOK {

	return &GetJobLogsOK{}
}

// WithPayload adds the payload to the get job logs o k response
func (o *GetJobLogsOK) WithPayload(payload *models.JobLogs) *GetJobLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job logs o k response
func (o *GetJobLogsOK) SetPayload(payload *models.JobLogs) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetJobLogsDefault Unexpected error

swagger:response getJobLogsDefault
*/
type GetJobLogsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetJobLogsDefault creates GetJobLogsDefault with default headers values
func NewGetJobLogsDefault(code int) *GetJobLogsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetJobLogsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get job logs default response
func (o *GetJobLogsDefault) WithStatusCode(code int) *GetJobLogsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get job logs default response
func (o *GetJobLogsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get job logs default response
func (o *GetJobLogsDefault) WithPayload(payload *models.Error) *GetJobLogsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job logs default response
func (o *GetJobLogsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobLogsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
