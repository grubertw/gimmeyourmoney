/******************************************************************************
 * Copyright (c) Travis Gruber (2017) All rights reserved.
 *
 * file: order_post_handler.go
 *
 * Contains http handler responsible for POST on the root [/] of the Order
 * Service. ALL order submissions must go through this API endpoint.
 *****************************************************************************/
package http_handlers

import (
	"net/http"

	"github.com/grubertw/gimmeyourmoney/OrderService/core"
)

/**
 * Object for encapsolating components used for processing an Order submission.
 */
type OrderPostHandler struct {
	validator orderValidator
	processor *core.OrderProcessor
}

/** Constructor for  OrderPostHandler */
func NewOrderPostHandler(op *core.OrderProcessor) *OrderPostHandler {

}

/** Handles all POST method calls on [/].
 *
 * This is the main API endpoint for Order submission. Please see documentation
 * in OrderService.md for expected structure of the JSON input.
 */
func (h *OrderPostHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {

}

/**
 * Object responsible for validating the internals of the OrderRequest.
 */
type orderValidator struct {

}

/** Validate product SKU is valid for the OrderItem. */
func (v *orderValidator) validateProduct(or *OrderItem) bool {

	return true
}

/** Validate RuleParameterCard is valid for this lender. */
func (v *orderValidator) validateParameterCard(pcid *string) bool {

	return true
}

/** Data structure deserialzied from the OrderPostHandler.
 *
 * This is the 'top-level' struct, which contains a number of structs
 * underneath.
 */
type OrderRequest struct {
	LenderID		string		`json:"lenderId"`
	OrderID			string		`json:"orderId"`
	Type			string		`json:"type"`
	Items			[]OrderItem	`json:"items"`
	UserID			string		`json:"userId"`
}

type OrderItem struct {
	ProductSKU			string			`json:"productSku"`
	ConsumerItems		[]ConsumerItem	`json:"consumerItems"`
	ParameterCardID 	string			`json:"parameterCardId"`
	CustomRuleParams	[]RuleParameter `json:"customRuleParams"`
	Extras				[]string		`json:"extras"`
}

type ConsumerItem struct {
	ConsumerID			string				`json:"consumerId"`
	ConsumerIdent		ConsumerIdent		`json:"consumerIdent"`
	RequestedAccounts	[]RequestedAccount	`json:"requestedAccounts"`
	ConsumerIsEngaged	bool				`json:"consumerIsEngaged"`
	LenderFlow			string				`json:"lenderFlow"`
}

type ConsumerIdent struct {
	FirstName		string			`json:"firstName"`
	MiddleName		string			`json:"middleName"`
	LastName		string			`json:"lastName"`
	SSN				string 			`json:"ssn"`
	Emails			[]string		`json:"emails"`
	Phones			[]string		`json:"phones"`
	Addresses		[]Address		`json:"addresses"`
}

type Address struct {
	Country			string			`json:"country"`
	State			string			`json:"state"`
	City			string 			`json:"city"`
	Street			string			`json:"street"`
	ZipCode			int				`json:"zipCode"`
}

type RequestedAccount struct {
	AccountNum		int			`json:"accountNum"`
	Required		bool 		`json:"required"`
}

type RuleParameter struct {
	Rule	string		`json:"rule"`
	Args	[]string	`json:"arguments"`
}