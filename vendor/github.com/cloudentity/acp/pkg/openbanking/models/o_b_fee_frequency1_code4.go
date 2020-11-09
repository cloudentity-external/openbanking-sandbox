// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OBFeeFrequency1Code4 Period e.g. day, week, month etc. for which the fee/charge is capped
//
// swagger:model OB_FeeFrequency1Code_4
type OBFeeFrequency1Code4 string

const (

	// OBFeeFrequency1Code4FEAC captures enum value "FEAC"
	OBFeeFrequency1Code4FEAC OBFeeFrequency1Code4 = "FEAC"

	// OBFeeFrequency1Code4FEAO captures enum value "FEAO"
	OBFeeFrequency1Code4FEAO OBFeeFrequency1Code4 = "FEAO"

	// OBFeeFrequency1Code4FECP captures enum value "FECP"
	OBFeeFrequency1Code4FECP OBFeeFrequency1Code4 = "FECP"

	// OBFeeFrequency1Code4FEDA captures enum value "FEDA"
	OBFeeFrequency1Code4FEDA OBFeeFrequency1Code4 = "FEDA"

	// OBFeeFrequency1Code4FEHO captures enum value "FEHO"
	OBFeeFrequency1Code4FEHO OBFeeFrequency1Code4 = "FEHO"

	// OBFeeFrequency1Code4FEI captures enum value "FEI"
	OBFeeFrequency1Code4FEI OBFeeFrequency1Code4 = "FEI"

	// OBFeeFrequency1Code4FEMO captures enum value "FEMO"
	OBFeeFrequency1Code4FEMO OBFeeFrequency1Code4 = "FEMO"

	// OBFeeFrequency1Code4FEOA captures enum value "FEOA"
	OBFeeFrequency1Code4FEOA OBFeeFrequency1Code4 = "FEOA"

	// OBFeeFrequency1Code4FEOT captures enum value "FEOT"
	OBFeeFrequency1Code4FEOT OBFeeFrequency1Code4 = "FEOT"

	// OBFeeFrequency1Code4FEPC captures enum value "FEPC"
	OBFeeFrequency1Code4FEPC OBFeeFrequency1Code4 = "FEPC"

	// OBFeeFrequency1Code4FEPH captures enum value "FEPH"
	OBFeeFrequency1Code4FEPH OBFeeFrequency1Code4 = "FEPH"

	// OBFeeFrequency1Code4FEPO captures enum value "FEPO"
	OBFeeFrequency1Code4FEPO OBFeeFrequency1Code4 = "FEPO"

	// OBFeeFrequency1Code4FEPS captures enum value "FEPS"
	OBFeeFrequency1Code4FEPS OBFeeFrequency1Code4 = "FEPS"

	// OBFeeFrequency1Code4FEPT captures enum value "FEPT"
	OBFeeFrequency1Code4FEPT OBFeeFrequency1Code4 = "FEPT"

	// OBFeeFrequency1Code4FEPTA captures enum value "FEPTA"
	OBFeeFrequency1Code4FEPTA OBFeeFrequency1Code4 = "FEPTA"

	// OBFeeFrequency1Code4FEPTP captures enum value "FEPTP"
	OBFeeFrequency1Code4FEPTP OBFeeFrequency1Code4 = "FEPTP"

	// OBFeeFrequency1Code4FEQU captures enum value "FEQU"
	OBFeeFrequency1Code4FEQU OBFeeFrequency1Code4 = "FEQU"

	// OBFeeFrequency1Code4FESM captures enum value "FESM"
	OBFeeFrequency1Code4FESM OBFeeFrequency1Code4 = "FESM"

	// OBFeeFrequency1Code4FEST captures enum value "FEST"
	OBFeeFrequency1Code4FEST OBFeeFrequency1Code4 = "FEST"

	// OBFeeFrequency1Code4FEWE captures enum value "FEWE"
	OBFeeFrequency1Code4FEWE OBFeeFrequency1Code4 = "FEWE"

	// OBFeeFrequency1Code4FEYE captures enum value "FEYE"
	OBFeeFrequency1Code4FEYE OBFeeFrequency1Code4 = "FEYE"
)

// for schema
var oBFeeFrequency1Code4Enum []interface{}

func init() {
	var res []OBFeeFrequency1Code4
	if err := json.Unmarshal([]byte(`["FEAC","FEAO","FECP","FEDA","FEHO","FEI","FEMO","FEOA","FEOT","FEPC","FEPH","FEPO","FEPS","FEPT","FEPTA","FEPTP","FEQU","FESM","FEST","FEWE","FEYE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBFeeFrequency1Code4Enum = append(oBFeeFrequency1Code4Enum, v)
	}
}

func (m OBFeeFrequency1Code4) validateOBFeeFrequency1Code4Enum(path, location string, value OBFeeFrequency1Code4) error {
	if err := validate.Enum(path, location, value, oBFeeFrequency1Code4Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b fee frequency1 code 4
func (m OBFeeFrequency1Code4) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBFeeFrequency1Code4Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
