//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaThingsCreateOKCode is the HTTP code returned for type SchemaThingsCreateOK
const SchemaThingsCreateOKCode int = 200

/*SchemaThingsCreateOK Added the new Thing class to the schema.

swagger:response schemaThingsCreateOK
*/
type SchemaThingsCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Class `json:"body,omitempty"`
}

// NewSchemaThingsCreateOK creates SchemaThingsCreateOK with default headers values
func NewSchemaThingsCreateOK() *SchemaThingsCreateOK {

	return &SchemaThingsCreateOK{}
}

// WithPayload adds the payload to the schema things create o k response
func (o *SchemaThingsCreateOK) WithPayload(payload *models.Class) *SchemaThingsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things create o k response
func (o *SchemaThingsCreateOK) SetPayload(payload *models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsCreateUnauthorizedCode is the HTTP code returned for type SchemaThingsCreateUnauthorized
const SchemaThingsCreateUnauthorizedCode int = 401

/*SchemaThingsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaThingsCreateUnauthorized
*/
type SchemaThingsCreateUnauthorized struct {
}

// NewSchemaThingsCreateUnauthorized creates SchemaThingsCreateUnauthorized with default headers values
func NewSchemaThingsCreateUnauthorized() *SchemaThingsCreateUnauthorized {

	return &SchemaThingsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaThingsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaThingsCreateForbiddenCode is the HTTP code returned for type SchemaThingsCreateForbidden
const SchemaThingsCreateForbiddenCode int = 403

/*SchemaThingsCreateForbidden Forbidden

swagger:response schemaThingsCreateForbidden
*/
type SchemaThingsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsCreateForbidden creates SchemaThingsCreateForbidden with default headers values
func NewSchemaThingsCreateForbidden() *SchemaThingsCreateForbidden {

	return &SchemaThingsCreateForbidden{}
}

// WithPayload adds the payload to the schema things create forbidden response
func (o *SchemaThingsCreateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaThingsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things create forbidden response
func (o *SchemaThingsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsCreateUnprocessableEntityCode is the HTTP code returned for type SchemaThingsCreateUnprocessableEntity
const SchemaThingsCreateUnprocessableEntityCode int = 422

/*SchemaThingsCreateUnprocessableEntity Invalid Thing class.

swagger:response schemaThingsCreateUnprocessableEntity
*/
type SchemaThingsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsCreateUnprocessableEntity creates SchemaThingsCreateUnprocessableEntity with default headers values
func NewSchemaThingsCreateUnprocessableEntity() *SchemaThingsCreateUnprocessableEntity {

	return &SchemaThingsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema things create unprocessable entity response
func (o *SchemaThingsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaThingsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things create unprocessable entity response
func (o *SchemaThingsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsCreateInternalServerErrorCode is the HTTP code returned for type SchemaThingsCreateInternalServerError
const SchemaThingsCreateInternalServerErrorCode int = 500

/*SchemaThingsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaThingsCreateInternalServerError
*/
type SchemaThingsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsCreateInternalServerError creates SchemaThingsCreateInternalServerError with default headers values
func NewSchemaThingsCreateInternalServerError() *SchemaThingsCreateInternalServerError {

	return &SchemaThingsCreateInternalServerError{}
}

// WithPayload adds the payload to the schema things create internal server error response
func (o *SchemaThingsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaThingsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things create internal server error response
func (o *SchemaThingsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
