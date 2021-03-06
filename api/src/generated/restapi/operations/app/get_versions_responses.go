// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// GetVersionsOKCode is the HTTP code returned for type GetVersionsOK
const GetVersionsOKCode int = 200

/*GetVersionsOK OK

swagger:response getVersionsOK
*/
type GetVersionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Versions `json:"body,omitempty"`
}

// NewGetVersionsOK creates GetVersionsOK with default headers values
func NewGetVersionsOK() *GetVersionsOK {

	return &GetVersionsOK{}
}

// WithPayload adds the payload to the get versions o k response
func (o *GetVersionsOK) WithPayload(payload *models.Versions) *GetVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get versions o k response
func (o *GetVersionsOK) SetPayload(payload *models.Versions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetVersionsDefault Unexpected error

swagger:response getVersionsDefault
*/
type GetVersionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVersionsDefault creates GetVersionsDefault with default headers values
func NewGetVersionsDefault(code int) *GetVersionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetVersionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get versions default response
func (o *GetVersionsDefault) WithStatusCode(code int) *GetVersionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get versions default response
func (o *GetVersionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get versions default response
func (o *GetVersionsDefault) WithPayload(payload *models.Error) *GetVersionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get versions default response
func (o *GetVersionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
