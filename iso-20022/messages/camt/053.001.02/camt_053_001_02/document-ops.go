// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/rs/zerolog/log"
	"reflect"
	"strings"
)

type SetOpOptions struct {
	logVerbose       bool
	skipIfEmpty      bool
	isClearOperation bool
}

type SetOpOption func(opts *SetOpOptions)

func SetOpWithLog(b bool) SetOpOption {
	return func(opts *SetOpOptions) {
		opts.logVerbose = b
	}
}

func SetOpWithSkipIfEmpty(b bool) SetOpOption {
	return func(opts *SetOpOptions) {
		opts.skipIfEmpty = b
	}
}

func SetOpWithIsClear(b bool) SetOpOption {
	return func(opts *SetOpOptions) {
		opts.isClearOperation = b
	}
}

type GetOpOptions struct {
	logVerbose bool
}

type GetOpOption func(opts *GetOpOptions)

func GetOpWithLog(b bool) GetOpOption {
	return func(opts *GetOpOptions) {
		opts.logVerbose = b
	}
}

// Set
// Deprecated: simply calls the SetProperty method
func (d *Document) Set(path string, src interface{}, opts ...SetOpOption) error {
	return d.SetProperty(path, src, opts...)
}

func (d *Document) SetProperty(path string, src interface{}, opts ...SetOpOption) error {

	const semLogContext = "camt_053_001_02::set-property"

	options := SetOpOptions{}
	for _, o := range opts {
		o(&options)
	}

	if options.skipIfEmpty && isEmpty(src) {
		return nil
	}

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(string(path))
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Interface("path", path).Interface("value", src).Msg(semLogContext)
		}

		return err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, false)
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Interface("path", path).Interface("value", src).Msg(semLogContext)
		}
		return err
	}

	return copy2Dest(path, values[0], src, &options)
}

func isEmpty(i interface{}) bool {
	b := false

	var s string
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
		b = strings.TrimSpace(s) == ""
	case string:
		s = ti
		b = strings.TrimSpace(s) == ""

	}

	return b
}

func (d *Document) ClearPath(p string, opts ...SetOpOption) error {

	const semLogContext = "camt_053_001_02::clear-path"

	options := SetOpOptions{}
	for _, o := range opts {
		o(&options)
	}

	pt, err := PathType(p)
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(p)).Msg(semLogContext)
		}

		return err
	}

	switch pt {
	case PathTypeProperty:
		i, err := d.GetProperty(p)
		if err != nil {
			return err
		}

		if i != nil {
			err = d.SetProperty(p, "", SetOpWithIsClear(true), SetOpWithLog(true))
		}
	case PathTypeArrayItem:
		err = fmt.Errorf("unsupported operation on array items")
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(p)).Msg(semLogContext)
		}
	case PathTypeArray:
		fallthrough
	case PathTypeStruct:
		err = d.ClearNode(p, SetOpWithLog(true))
	}

	return err
}

// Get
// Deprecated: the simply calls the GetProperty method
func (d *Document) Get(path string, opts ...GetOpOption) (interface{}, error) {
	return d.GetProperty(path, opts...)
}

func (d *Document) GetProperty(path string, opts ...GetOpOption) (interface{}, error) {

	const semLogContext = "camt_053_001_02::get-property"

	options := GetOpOptions{}
	for _, o := range opts {
		o(&options)
	}

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(string(path))
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(path)).Msg(semLogContext)
		}
		return nil, err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, true)
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(path)).Msg(semLogContext)
		}
		return nil, err
	}

	/*
		rv := reflect.ValueOf(values[0])
		fmt.Println("Indirect type is:", reflect.Indirect(rv), reflect.Indirect(rv).Type(), rv.Kind(), rv.Elem(), rv.Elem().Type()) // prints main.CustomStruct

		if tv, ok := values[0].(*common.Max35Text); ok {
			return *tv, nil
		}
	*/

	return derefProperty(path, values[0])
}

func copy2Dest(docPath string, dest, src interface{}, options *SetOpOptions) error {

	const semLogContext = "camt_053_001_02::copy-to-dest"

	var err error
	switch typedDest := dest.(type) {
	case *common.ActiveOrHistoricCurrencyCode:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ActiveOrHistoricCurrencyCodeZero
		} else {
			*typedDest, err = common.ToActiveOrHistoricCurrencyCode(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.AddressType2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.AddressType2CodeZero
		} else {
			*typedDest, err = common.ToAddressType2Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.AnyBICIdentifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.AnyBICIdentifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.AnyBICIdentifierZero
		} else {
			*typedDest, err = common.ToAnyBICIdentifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.BICIdentifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.BICIdentifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.BICIdentifierZero
		} else {
			*typedDest, err = common.ToBICIdentifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.BalanceType12Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.BalanceType12Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.BalanceType12CodeZero
		} else {
			*typedDest, err = common.ToBalanceType12Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.CashAccountType4Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.CashAccountType4Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.CashAccountType4CodeZero
		} else {
			*typedDest, err = common.ToCashAccountType4Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ChargeBearerType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ChargeBearerType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ChargeBearerType1CodeZero
		} else {
			*typedDest, err = common.ToChargeBearerType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ChargeType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ChargeType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ChargeType1CodeZero
		} else {
			*typedDest, err = common.ToChargeType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.CopyDuplicate1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.CopyDuplicate1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.CopyDuplicate1CodeZero
		} else {
			*typedDest, err = common.ToCopyDuplicate1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.CountryCode:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.CountryCode data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.CountryCodeZero
		} else {
			*typedDest, err = common.ToCountryCode(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.CreditDebitCode:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.CreditDebitCodeZero
		} else {
			*typedDest, err = common.ToCreditDebitCode(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.DocumentType3Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.DocumentType3CodeZero
		} else {
			*typedDest, err = common.ToDocumentType3Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.DocumentType5Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.DocumentType5Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.DocumentType5CodeZero
		} else {
			*typedDest, err = common.ToDocumentType5Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.EntryStatus2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.EntryStatus2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.EntryStatus2CodeZero
		} else {
			*typedDest, err = common.ToEntryStatus2Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalAccountIdentification1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalAccountIdentification1CodeZero
		} else {
			*typedDest, err = common.ToExternalAccountIdentification1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalBalanceSubType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalBalanceSubType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalBalanceSubType1CodeZero
		} else {
			*typedDest, err = common.ToExternalBalanceSubType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalBankTransactionDomain1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionDomain1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalBankTransactionDomain1CodeZero
		} else {
			*typedDest, err = common.ToExternalBankTransactionDomain1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalBankTransactionFamily1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionFamily1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalBankTransactionFamily1CodeZero
		} else {
			*typedDest, err = common.ToExternalBankTransactionFamily1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalBankTransactionSubFamily1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalBankTransactionSubFamily1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalBankTransactionSubFamily1CodeZero
		} else {
			*typedDest, err = common.ToExternalBankTransactionSubFamily1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalClearingSystemIdentification1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalClearingSystemIdentification1CodeZero
		} else {
			*typedDest, err = common.ToExternalClearingSystemIdentification1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalFinancialInstitutionIdentification1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalFinancialInstitutionIdentification1CodeZero
		} else {
			*typedDest, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalOrganisationIdentification1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalOrganisationIdentification1CodeZero
		} else {
			*typedDest, err = common.ToExternalOrganisationIdentification1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalPersonIdentification1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalPersonIdentification1CodeZero
		} else {
			*typedDest, err = common.ToExternalPersonIdentification1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalPurpose1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalPurpose1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalPurpose1CodeZero
		} else {
			*typedDest, err = common.ToExternalPurpose1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalReportingSource1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalReportingSource1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalReportingSource1CodeZero
		} else {
			*typedDest, err = common.ToExternalReportingSource1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalReturnReason1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalReturnReason1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalReturnReason1CodeZero
		} else {
			*typedDest, err = common.ToExternalReturnReason1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalTechnicalInputChannel1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalTechnicalInputChannel1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalTechnicalInputChannel1CodeZero
		} else {
			*typedDest, err = common.ToExternalTechnicalInputChannel1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.IBAN2007Identifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.IBAN2007IdentifierZero
		} else {
			*typedDest, err = common.ToIBAN2007Identifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ISINIdentifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ISINIdentifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ISINIdentifierZero
		} else {
			*typedDest, err = common.ToISINIdentifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ISODate:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ISODate data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ISODateZero
		} else {
			*typedDest, err = common.ToISODate(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ISODateTime:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ISODateTimeZero
		} else {
			*typedDest, err = common.ToISODateTime(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.InterestType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.InterestType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.InterestType1CodeZero
		} else {
			*typedDest, err = common.ToInterestType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max105Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max105Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max105TextZero
		} else {
			*typedDest, err = common.ToMax105Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max140Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max140Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max140TextZero
		} else {
			*typedDest, err = common.ToMax140Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max15NumericText:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max15NumericTextZero
		} else {
			*typedDest, err = common.ToMax15NumericText(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max15PlusSignedNumericText:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max15PlusSignedNumericText data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max15PlusSignedNumericTextZero
		} else {
			*typedDest, err = common.ToMax15PlusSignedNumericText(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max16Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max16Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max16TextZero
		} else {
			*typedDest, err = common.ToMax16Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max2048Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max2048TextZero
		} else {
			*typedDest, err = common.ToMax2048Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max34Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max34Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max34TextZero
		} else {
			*typedDest, err = common.ToMax34Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max35Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max35Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max35TextZero
		} else {
			*typedDest, err = common.ToMax35Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max4Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max4Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max4TextZero
		} else {
			*typedDest, err = common.ToMax4Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max500Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max500Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max500TextZero
		} else {
			*typedDest, err = common.ToMax500Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max5NumericText:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max5NumericText data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max5NumericTextZero
		} else {
			*typedDest, err = common.ToMax5NumericText(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max70Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max70Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max70TextZero
		} else {
			*typedDest, err = common.ToMax70Text(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.NamePrefix1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.NamePrefix1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.NamePrefix1CodeZero
		} else {
			*typedDest, err = common.ToNamePrefix1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.PhoneNumber:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.PhoneNumberZero
		} else {
			*typedDest, err = common.ToPhoneNumber(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.RemittanceLocationMethod2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.RemittanceLocationMethod2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.RemittanceLocationMethod2CodeZero
		} else {
			*typedDest, err = common.ToRemittanceLocationMethod2Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.TaxRecordPeriod1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.TaxRecordPeriod1CodeZero
		} else {
			*typedDest, err = common.ToTaxRecordPeriod1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *xsdt.Boolean:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = xsdt.BooleanZero
		} else {
			*typedDest, err = xsdt.ToBoolean(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *xsdt.Decimal:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = xsdt.DecimalZero
		} else {
			*typedDest, err = xsdt.ToDecimal(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", docPath, dest)
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	}
}

func derefProperty(docPath string, val interface{}) (interface{}, error) {

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
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), val)
	}

	return val, err
}
