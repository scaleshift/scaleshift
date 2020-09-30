// Code generated by go-swagger; DO NOT EDIT.

package notebook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/scaleshift/scaleshift/api/src/generated/models"
)

// DeleteNotebookNoContentCode is the HTTP code returned for type DeleteNotebookNoContent
const DeleteNotebookNoContentCode int = 204

/*DeleteNotebookNoContent OK

swagger:response deleteNotebookNoContent
*/
type DeleteNotebookNoContent struct {
}

// NewDeleteNotebookNoContent creates DeleteNotebookNoContent with default headers values
func NewDeleteNotebookNoContent() *DeleteNotebookNoContent {

	return &DeleteNotebookNoContent{}
}

// WriteResponse to the client
func (o *DeleteNotebookNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteNotebookDefault Unexpected error

swagger:response deleteNotebookDefault
*/
type DeleteNotebookDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNotebookDefault creates DeleteNotebookDefault with default headers values
func NewDeleteNotebookDefault(code int) *DeleteNotebookDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteNotebookDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete notebook default response
func (o *DeleteNotebookDefault) WithStatusCode(code int) *DeleteNotebookDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete notebook default response
func (o *DeleteNotebookDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete notebook default response
func (o *DeleteNotebookDefault) WithPayload(payload *models.Error) *DeleteNotebookDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete notebook default response
func (o *DeleteNotebookDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNotebookDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
