// Code generated by go-swagger; DO NOT EDIT.

package quiz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "frontend/gen/models"
)

// GetQuizzesOKCode is the HTTP code returned for type GetQuizzesOK
const GetQuizzesOKCode int = 200

/*GetQuizzesOK OK

swagger:response getQuizzesOK
*/
type GetQuizzesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.QuizQuestion `json:"body,omitempty"`
}

// NewGetQuizzesOK creates GetQuizzesOK with default headers values
func NewGetQuizzesOK() *GetQuizzesOK {

	return &GetQuizzesOK{}
}

// WithPayload adds the payload to the get quizzes o k response
func (o *GetQuizzesOK) WithPayload(payload []*models.QuizQuestion) *GetQuizzesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get quizzes o k response
func (o *GetQuizzesOK) SetPayload(payload []*models.QuizQuestion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuizzesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.QuizQuestion, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetQuizzesDefault Generic error

swagger:response getQuizzesDefault
*/
type GetQuizzesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuizzesDefault creates GetQuizzesDefault with default headers values
func NewGetQuizzesDefault(code int) *GetQuizzesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetQuizzesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get quizzes default response
func (o *GetQuizzesDefault) WithStatusCode(code int) *GetQuizzesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get quizzes default response
func (o *GetQuizzesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get quizzes default response
func (o *GetQuizzesDefault) WithPayload(payload *models.Error) *GetQuizzesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get quizzes default response
func (o *GetQuizzesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuizzesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
