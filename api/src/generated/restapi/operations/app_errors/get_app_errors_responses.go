// Code generated by go-swagger; DO NOT EDIT.

package app_errors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// GetAppErrorsOKCode is the HTTP code returned for type GetAppErrorsOK
const GetAppErrorsOKCode int = 200

/*GetAppErrorsOK OK

swagger:response getAppErrorsOK
*/
type GetAppErrorsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AppError `json:"body,omitempty"`
}

// NewGetAppErrorsOK creates GetAppErrorsOK with default headers values
func NewGetAppErrorsOK() *GetAppErrorsOK {

	return &GetAppErrorsOK{}
}

// WithPayload adds the payload to the get app errors o k response
func (o *GetAppErrorsOK) WithPayload(payload []*models.AppError) *GetAppErrorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app errors o k response
func (o *GetAppErrorsOK) SetPayload(payload []*models.AppError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppErrorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.AppError, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetAppErrorsDefault Unexpected error

swagger:response getAppErrorsDefault
*/
type GetAppErrorsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAppErrorsDefault creates GetAppErrorsDefault with default headers values
func NewGetAppErrorsDefault(code int) *GetAppErrorsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppErrorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get app errors default response
func (o *GetAppErrorsDefault) WithStatusCode(code int) *GetAppErrorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get app errors default response
func (o *GetAppErrorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get app errors default response
func (o *GetAppErrorsDefault) WithPayload(payload *models.Error) *GetAppErrorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app errors default response
func (o *GetAppErrorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppErrorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
