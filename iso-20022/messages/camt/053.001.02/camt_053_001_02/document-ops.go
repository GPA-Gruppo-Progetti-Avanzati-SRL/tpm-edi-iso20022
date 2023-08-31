// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"reflect"
)

func (d *Document) Set(path string, src interface{}) error {

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(path)
	if err != nil {
		return err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, false)
	if err != nil {
		return err
	}

	return copy2Dest(path, values[0], src)
}

func (d *Document) Get(path string) (interface{}, error) {

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(path)
	if err != nil {
		return nil, err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, true)
	if err != nil {
		return nil, err
	}

	/*
		rv := reflect.ValueOf(values[0])
		fmt.Println("Indirect type is:", reflect.Indirect(rv), reflect.Indirect(rv).Type(), rv.Kind(), rv.Elem(), rv.Elem().Type()) // prints main.CustomStruct

		if tv, ok := values[0].(*common.Max35Text); ok {
			return *tv, nil
		}
	*/

	return deref(path, values[0])
}

func copy2Dest(docPath string, dest, src interface{}) error {

	var err error
	switch typedDest := dest.(type) {
	case *common.ActiveOrHistoricCurrencyCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*typedDest, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err
	case *common.AddressType2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*typedDest, err = common.ToAddressType2Code(src)
		return err
	case *common.AnyBICIdentifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICIdentifier data for" + docPath)
		}

		*typedDest, err = common.ToAnyBICIdentifier(src)
		return err
	case *common.BICIdentifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.BICIdentifier data for" + docPath)
		}

		*typedDest, err = common.ToBICIdentifier(src)
		return err
	case *common.BalanceType12Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.BalanceType12Code data for" + docPath)
		}

		*typedDest, err = common.ToBalanceType12Code(src)
		return err
	case *common.CashAccountType4Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CashAccountType4Code data for" + docPath)
		}

		*typedDest, err = common.ToCashAccountType4Code(src)
		return err
	case *common.ChargeBearerType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ChargeBearerType1Code data for" + docPath)
		}

		*typedDest, err = common.ToChargeBearerType1Code(src)
		return err
	case *common.ChargeType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ChargeType1Code data for" + docPath)
		}

		*typedDest, err = common.ToChargeType1Code(src)
		return err
	case *common.CopyDuplicate1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CopyDuplicate1Code data for" + docPath)
		}

		*typedDest, err = common.ToCopyDuplicate1Code(src)
		return err
	case *common.CountryCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*typedDest, err = common.ToCountryCode(src)
		return err
	case *common.CreditDebitCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*typedDest, err = common.ToCreditDebitCode(src)
		return err
	case *common.DocumentType3Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + docPath)
		}

		*typedDest, err = common.ToDocumentType3Code(src)
		return err
	case *common.DocumentType5Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType5Code data for" + docPath)
		}

		*typedDest, err = common.ToDocumentType5Code(src)
		return err
	case *common.EntryStatus2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.EntryStatus2Code data for" + docPath)
		}

		*typedDest, err = common.ToEntryStatus2Code(src)
		return err
	case *common.ExternalAccountIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalAccountIdentification1Code(src)
		return err
	case *common.ExternalBalanceSubType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalBalanceSubType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalBalanceSubType1Code(src)
		return err
	case *common.ExternalBankTransactionDomain1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionDomain1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalBankTransactionDomain1Code(src)
		return err
	case *common.ExternalBankTransactionFamily1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionFamily1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalBankTransactionFamily1Code(src)
		return err
	case *common.ExternalBankTransactionSubFamily1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionSubFamily1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalBankTransactionSubFamily1Code(src)
		return err
	case *common.ExternalClearingSystemIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err
	case *common.ExternalFinancialInstitutionIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err
	case *common.ExternalOrganisationIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalOrganisationIdentification1Code(src)
		return err
	case *common.ExternalPersonIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalPersonIdentification1Code(src)
		return err
	case *common.ExternalPurpose1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPurpose1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalPurpose1Code(src)
		return err
	case *common.ExternalReportingSource1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalReportingSource1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalReportingSource1Code(src)
		return err
	case *common.ExternalReturnReason1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalReturnReason1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalReturnReason1Code(src)
		return err
	case *common.ExternalTechnicalInputChannel1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTechnicalInputChannel1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalTechnicalInputChannel1Code(src)
		return err
	case *common.IBAN2007Identifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*typedDest, err = common.ToIBAN2007Identifier(src)
		return err
	case *common.ISINIdentifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ISINIdentifier data for" + docPath)
		}

		*typedDest, err = common.ToISINIdentifier(src)
		return err
	case *common.ISODate:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*typedDest, err = common.ToISODate(src)
		return err
	case *common.ISODateTime:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*typedDest, err = common.ToISODateTime(src)
		return err
	case *common.InterestType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.InterestType1Code data for" + docPath)
		}

		*typedDest, err = common.ToInterestType1Code(src)
		return err
	case *common.Max105Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max105Text data for" + docPath)
		}

		*typedDest, err = common.ToMax105Text(src)
		return err
	case *common.Max140Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*typedDest, err = common.ToMax140Text(src)
		return err
	case *common.Max15NumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*typedDest, err = common.ToMax15NumericText(src)
		return err
	case *common.Max15PlusSignedNumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15PlusSignedNumericText data for" + docPath)
		}

		*typedDest, err = common.ToMax15PlusSignedNumericText(src)
		return err
	case *common.Max16Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*typedDest, err = common.ToMax16Text(src)
		return err
	case *common.Max2048Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*typedDest, err = common.ToMax2048Text(src)
		return err
	case *common.Max34Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*typedDest, err = common.ToMax34Text(src)
		return err
	case *common.Max35Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*typedDest, err = common.ToMax35Text(src)
		return err
	case *common.Max4Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*typedDest, err = common.ToMax4Text(src)
		return err
	case *common.Max500Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max500Text data for" + docPath)
		}

		*typedDest, err = common.ToMax500Text(src)
		return err
	case *common.Max5NumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max5NumericText data for" + docPath)
		}

		*typedDest, err = common.ToMax5NumericText(src)
		return err
	case *common.Max70Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*typedDest, err = common.ToMax70Text(src)
		return err
	case *common.NamePrefix1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix1Code data for" + docPath)
		}

		*typedDest, err = common.ToNamePrefix1Code(src)
		return err
	case *common.PhoneNumber:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*typedDest, err = common.ToPhoneNumber(src)
		return err
	case *common.RemittanceLocationMethod2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.RemittanceLocationMethod2Code data for" + docPath)
		}

		*typedDest, err = common.ToRemittanceLocationMethod2Code(src)
		return err
	case *common.TaxRecordPeriod1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*typedDest, err = common.ToTaxRecordPeriod1Code(src)
		return err
	case *xsdt.Boolean:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*typedDest, err = xsdt.ToBoolean(src)
		return err
	case *xsdt.Decimal:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*typedDest, err = xsdt.ToDecimal(src)
		return err
	default:
		return fmt.Errorf("could not find the type to node %s of type %v", docPath, dest)
	}

}

func deref(docPath string, val interface{}) (interface{}, error) {

	if val == nil {
		return nil, nil
	}

	var err error
	switch tv := val.(type) {
	case *common.ActiveOrHistoricCurrencyCode:
		return *tv, nil
	case *common.AddressType2Code:
		return *tv, nil
	case *common.AnyBICIdentifier:
		return *tv, nil
	case *common.BICIdentifier:
		return *tv, nil
	case *common.BalanceType12Code:
		return *tv, nil
	case *common.CashAccountType4Code:
		return *tv, nil
	case *common.ChargeBearerType1Code:
		return *tv, nil
	case *common.ChargeType1Code:
		return *tv, nil
	case *common.CopyDuplicate1Code:
		return *tv, nil
	case *common.CountryCode:
		return *tv, nil
	case *common.CreditDebitCode:
		return *tv, nil
	case *common.DocumentType3Code:
		return *tv, nil
	case *common.DocumentType5Code:
		return *tv, nil
	case *common.EntryStatus2Code:
		return *tv, nil
	case *common.ExternalAccountIdentification1Code:
		return *tv, nil
	case *common.ExternalBalanceSubType1Code:
		return *tv, nil
	case *common.ExternalBankTransactionDomain1Code:
		return *tv, nil
	case *common.ExternalBankTransactionFamily1Code:
		return *tv, nil
	case *common.ExternalBankTransactionSubFamily1Code:
		return *tv, nil
	case *common.ExternalClearingSystemIdentification1Code:
		return *tv, nil
	case *common.ExternalFinancialInstitutionIdentification1Code:
		return *tv, nil
	case *common.ExternalOrganisationIdentification1Code:
		return *tv, nil
	case *common.ExternalPersonIdentification1Code:
		return *tv, nil
	case *common.ExternalPurpose1Code:
		return *tv, nil
	case *common.ExternalReportingSource1Code:
		return *tv, nil
	case *common.ExternalReturnReason1Code:
		return *tv, nil
	case *common.ExternalTechnicalInputChannel1Code:
		return *tv, nil
	case *common.IBAN2007Identifier:
		return *tv, nil
	case *common.ISINIdentifier:
		return *tv, nil
	case *common.ISODate:
		return *tv, nil
	case *common.ISODateTime:
		return *tv, nil
	case *common.InterestType1Code:
		return *tv, nil
	case *common.Max105Text:
		return *tv, nil
	case *common.Max140Text:
		return *tv, nil
	case *common.Max15NumericText:
		return *tv, nil
	case *common.Max15PlusSignedNumericText:
		return *tv, nil
	case *common.Max16Text:
		return *tv, nil
	case *common.Max2048Text:
		return *tv, nil
	case *common.Max34Text:
		return *tv, nil
	case *common.Max35Text:
		return *tv, nil
	case *common.Max4Text:
		return *tv, nil
	case *common.Max500Text:
		return *tv, nil
	case *common.Max5NumericText:
		return *tv, nil
	case *common.Max70Text:
		return *tv, nil
	case *common.NamePrefix1Code:
		return *tv, nil
	case *common.PhoneNumber:
		return *tv, nil
	case *common.RemittanceLocationMethod2Code:
		return *tv, nil
	case *common.TaxRecordPeriod1Code:
		return *tv, nil
	case *xsdt.Boolean:
		return *tv, nil
	case *xsdt.Decimal:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %v", docPath, val)
	}

	return val, err
}
