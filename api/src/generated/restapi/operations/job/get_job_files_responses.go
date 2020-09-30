// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// GetJobFilesOKCode is the HTTP code returned for type GetJobFilesOK
const GetJobFilesOKCode int = 200

/*GetJobFilesOK OK

swagger:response getJobFilesOK
*/
type GetJobFilesOK struct {

	/*
	  In: Body
	*/
	Payload *models.JobFiles `json:"body,omitempty"`
}

// NewGetJobFilesOK creates GetJobFilesOK with default headers values
func NewGetJobFilesOK() *GetJobFilesOK {

	return &GetJobFilesOK{}
}

// WithPayload adds the payload to the get job files o k response
func (o *GetJobFilesOK) WithPayload(payload *models.JobFiles) *GetJobFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job files o k response
func (o *GetJobFilesOK) SetPayload(payload *models.JobFiles) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetJobFilesDefault Unexpected error

swagger:response getJobFilesDefault
*/
type GetJobFilesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetJobFilesDefault creates GetJobFilesDefault with default headers values
func NewGetJobFilesDefault(code int) *GetJobFilesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetJobFilesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get job files default response
func (o *GetJobFilesDefault) WithStatusCode(code int) *GetJobFilesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get job files default response
func (o *GetJobFilesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get job files default response
func (o *GetJobFilesDefault) WithPayload(payload *models.Error) *GetJobFilesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job files default response
func (o *GetJobFilesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobFilesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
