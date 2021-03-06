// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBTransaction6Basic Provides further details on an entry in the report.
//
// swagger:model OBTransaction6Basic
type OBTransaction6Basic struct {

	// account Id
	// Required: true
	AccountID AccountID `json:"AccountId"`

	// address line
	AddressLine AddressLine `json:"AddressLine,omitempty"`

	// amount
	// Required: true
	Amount *OBActiveOrHistoricCurrencyAndAmount9 `json:"Amount"`

	// bank transaction code
	BankTransactionCode *OBBankTransactionCodeStructure1 `json:"BankTransactionCode,omitempty"`

	// booking date time
	// Required: true
	// Format: date-time
	BookingDateTime BookingDateTime `json:"BookingDateTime"`

	// card instrument
	CardInstrument *OBTransactionCardInstrument1 `json:"CardInstrument,omitempty"`

	// charge amount
	ChargeAmount *OBActiveOrHistoricCurrencyAndAmount10 `json:"ChargeAmount,omitempty"`

	// credit debit indicator
	// Required: true
	CreditDebitIndicator OBCreditDebitCode1 `json:"CreditDebitIndicator"`

	// currency exchange
	CurrencyExchange *OBCurrencyExchange5 `json:"CurrencyExchange,omitempty"`

	// proprietary bank transaction code
	ProprietaryBankTransactionCode *ProprietaryBankTransactionCodeStructure1 `json:"ProprietaryBankTransactionCode,omitempty"`

	// statement reference
	StatementReference []StatementReference `json:"StatementReference"`

	// status
	// Required: true
	Status OBEntryStatus1Code `json:"Status"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`

	// transaction Id
	TransactionID TransactionID `json:"TransactionId,omitempty"`

	// transaction mutability
	TransactionMutability OBTransactionMutability1Code `json:"TransactionMutability,omitempty"`

	// transaction reference
	TransactionReference TransactionReference `json:"TransactionReference,omitempty"`

	// value date time
	// Format: date-time
	ValueDateTime ValueDateTime `json:"ValueDateTime,omitempty"`
}

// Validate validates this o b transaction6 basic
func (m *OBTransaction6Basic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankTransactionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBookingDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditDebitIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyExchange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProprietaryBankTransactionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionMutability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBTransaction6Basic) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateAddressLine(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine) { // not required
		return nil
	}

	if err := m.AddressLine.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressLine")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateBankTransactionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankTransactionCode) { // not required
		return nil
	}

	if m.BankTransactionCode != nil {
		if err := m.BankTransactionCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BankTransactionCode")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateBookingDateTime(formats strfmt.Registry) error {

	if err := m.BookingDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BookingDateTime")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateCardInstrument(formats strfmt.Registry) error {

	if swag.IsZero(m.CardInstrument) { // not required
		return nil
	}

	if m.CardInstrument != nil {
		if err := m.CardInstrument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CardInstrument")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateChargeAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargeAmount) { // not required
		return nil
	}

	if m.ChargeAmount != nil {
		if err := m.ChargeAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChargeAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateCreditDebitIndicator(formats strfmt.Registry) error {

	if err := m.CreditDebitIndicator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CreditDebitIndicator")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateCurrencyExchange(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrencyExchange) { // not required
		return nil
	}

	if m.CurrencyExchange != nil {
		if err := m.CurrencyExchange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrencyExchange")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateProprietaryBankTransactionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ProprietaryBankTransactionCode) { // not required
		return nil
	}

	if m.ProprietaryBankTransactionCode != nil {
		if err := m.ProprietaryBankTransactionCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProprietaryBankTransactionCode")
			}
			return err
		}
	}

	return nil
}

func (m *OBTransaction6Basic) validateStatementReference(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementReference) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementReference); i++ {

		if err := m.StatementReference[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StatementReference" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *OBTransaction6Basic) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateTransactionID(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionID) { // not required
		return nil
	}

	if err := m.TransactionID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransactionId")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateTransactionMutability(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionMutability) { // not required
		return nil
	}

	if err := m.TransactionMutability.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransactionMutability")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateTransactionReference(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionReference) { // not required
		return nil
	}

	if err := m.TransactionReference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransactionReference")
		}
		return err
	}

	return nil
}

func (m *OBTransaction6Basic) validateValueDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueDateTime) { // not required
		return nil
	}

	if err := m.ValueDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ValueDateTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBTransaction6Basic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBTransaction6Basic) UnmarshalBinary(b []byte) error {
	var res OBTransaction6Basic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
