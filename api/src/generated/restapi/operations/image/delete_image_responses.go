// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// DeleteImageNoContentCode is the HTTP code returned for type DeleteImageNoContent
const DeleteImageNoContentCode int = 204

/*DeleteImageNoContent OK

swagger:response deleteImageNoContent
*/
type DeleteImageNoContent struct {
}

// NewDeleteImageNoContent creates DeleteImageNoContent with default headers values
func NewDeleteImageNoContent() *DeleteImageNoContent {

	return &DeleteImageNoContent{}
}

// WriteResponse to the client
func (o *DeleteImageNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteImageDefault Unexpected error

swagger:response deleteImageDefault
*/
type DeleteImageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteImageDefault creates DeleteImageDefault with default headers values
func NewDeleteImageDefault(code int) *DeleteImageDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteImageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete image default response
func (o *DeleteImageDefault) WithStatusCode(code int) *DeleteImageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete image default response
func (o *DeleteImageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete image default response
func (o *DeleteImageDefault) WithPayload(payload *models.Error) *DeleteImageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image default response
func (o *DeleteImageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
