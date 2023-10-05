// Package pain_002_001_10
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_10

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.10/common"
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

	const semLogContext = "pain_002_001_10::clear-path"

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

	const semLogContext = "pain_002_001_10::get-property"

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

	const semLogContext = "pain_002_001_10::copy-to-dest"

	var err error
	switch typedDest := dest.(type) {
	case *common.ActiveCurrencyCode:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ActiveCurrencyCode data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ActiveCurrencyCodeZero
		} else {
			*typedDest, err = common.ToActiveCurrencyCode(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
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
	case *common.AnyBICDec2014Identifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.AnyBICDec2014IdentifierZero
		} else {
			*typedDest, err = common.ToAnyBICDec2014Identifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.BICFIDec2014Identifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.BICFIDec2014IdentifierZero
		} else {
			*typedDest, err = common.ToBICFIDec2014Identifier(src)
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
	case *common.ClearingChannel2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ClearingChannel2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ClearingChannel2CodeZero
		} else {
			*typedDest, err = common.ToClearingChannel2Code(src)
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
	case *common.DocumentType6Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.DocumentType6Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.DocumentType6CodeZero
		} else {
			*typedDest, err = common.ToDocumentType6Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Exact2NumericText:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Exact2NumericTextZero
		} else {
			*typedDest, err = common.ToExact2NumericText(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Exact4AlphaNumericText:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Exact4AlphaNumericTextZero
		} else {
			*typedDest, err = common.ToExact4AlphaNumericText(src)
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
	case *common.ExternalCashAccountType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalCashAccountType1CodeZero
		} else {
			*typedDest, err = common.ToExternalCashAccountType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalCashClearingSystem1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalCashClearingSystem1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalCashClearingSystem1CodeZero
		} else {
			*typedDest, err = common.ToExternalCashClearingSystem1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalCategoryPurpose1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalCategoryPurpose1CodeZero
		} else {
			*typedDest, err = common.ToExternalCategoryPurpose1Code(src)
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
	case *common.ExternalDiscountAmountType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalDiscountAmountType1CodeZero
		} else {
			*typedDest, err = common.ToExternalDiscountAmountType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalDocumentLineType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalDocumentLineType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalDocumentLineType1CodeZero
		} else {
			*typedDest, err = common.ToExternalDocumentLineType1Code(src)
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
	case *common.ExternalGarnishmentType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalGarnishmentType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalGarnishmentType1CodeZero
		} else {
			*typedDest, err = common.ToExternalGarnishmentType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalLocalInstrument1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalLocalInstrument1CodeZero
		} else {
			*typedDest, err = common.ToExternalLocalInstrument1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalMandateSetupReason1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalMandateSetupReason1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalMandateSetupReason1CodeZero
		} else {
			*typedDest, err = common.ToExternalMandateSetupReason1Code(src)
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
	case *common.ExternalPaymentGroupStatus1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalPaymentGroupStatus1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalPaymentGroupStatus1CodeZero
		} else {
			*typedDest, err = common.ToExternalPaymentGroupStatus1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalPaymentTransactionStatus1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalPaymentTransactionStatus1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalPaymentTransactionStatus1CodeZero
		} else {
			*typedDest, err = common.ToExternalPaymentTransactionStatus1Code(src)
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
	case *common.ExternalProxyAccountType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalProxyAccountType1CodeZero
		} else {
			*typedDest, err = common.ToExternalProxyAccountType1Code(src)
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
	case *common.ExternalServiceLevel1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalServiceLevel1CodeZero
		} else {
			*typedDest, err = common.ToExternalServiceLevel1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalStatusReason1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalStatusReason1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalStatusReason1CodeZero
		} else {
			*typedDest, err = common.ToExternalStatusReason1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExternalTaxAmountType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.ExternalTaxAmountType1CodeZero
		} else {
			*typedDest, err = common.ToExternalTaxAmountType1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Frequency6Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Frequency6CodeZero
		} else {
			*typedDest, err = common.ToFrequency6Code(src)
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
	case *common.LEIIdentifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.LEIIdentifierZero
		} else {
			*typedDest, err = common.ToLEIIdentifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max1025Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max1025Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max1025TextZero
		} else {
			*typedDest, err = common.ToMax1025Text(src)
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
	case *common.Max128Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max128Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max128TextZero
		} else {
			*typedDest, err = common.ToMax128Text(src)
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
	case *common.Max350Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max350Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Max350TextZero
		} else {
			*typedDest, err = common.ToMax350Text(src)
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
	case *common.NamePrefix2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.NamePrefix2CodeZero
		} else {
			*typedDest, err = common.ToNamePrefix2Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.PaymentMethod4Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.PaymentMethod4Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.PaymentMethod4CodeZero
		} else {
			*typedDest, err = common.ToPaymentMethod4Code(src)
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
	case *common.PreferredContactMethod1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.PreferredContactMethod1CodeZero
		} else {
			*typedDest, err = common.ToPreferredContactMethod1Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Priority2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.Priority2CodeZero
		} else {
			*typedDest, err = common.ToPriority2Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.SequenceType3Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.SequenceType3Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.SequenceType3CodeZero
		} else {
			*typedDest, err = common.ToSequenceType3Code(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.SettlementMethod1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.SettlementMethod1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.SettlementMethod1CodeZero
		} else {
			*typedDest, err = common.ToSettlementMethod1Code(src)
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
	case *common.UUIDv4Identifier:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.UUIDv4Identifier data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = common.UUIDv4IdentifierZero
		} else {
			*typedDest, err = common.ToUUIDv4Identifier(src)
		}

		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *xsdt.AnyType:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		if options.isClearOperation {
			*typedDest = xsdt.AnyTypeZero
		} else {
			*typedDest, err = xsdt.ToAnyType(src)
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
	case *common.ActiveCurrencyCode:
		return *tv, nil
	case *common.ActiveOrHistoricCurrencyCode:
		return *tv, nil
	case *common.AddressType2Code:
		return *tv, nil
	case *common.AnyBICDec2014Identifier:
		return *tv, nil
	case *common.BICFIDec2014Identifier:
		return *tv, nil
	case *common.ChargeBearerType1Code:
		return *tv, nil
	case *common.ClearingChannel2Code:
		return *tv, nil
	case *common.CountryCode:
		return *tv, nil
	case *common.CreditDebitCode:
		return *tv, nil
	case *common.DocumentType3Code:
		return *tv, nil
	case *common.DocumentType6Code:
		return *tv, nil
	case *common.Exact2NumericText:
		return *tv, nil
	case *common.Exact4AlphaNumericText:
		return *tv, nil
	case *common.ExternalAccountIdentification1Code:
		return *tv, nil
	case *common.ExternalCashAccountType1Code:
		return *tv, nil
	case *common.ExternalCashClearingSystem1Code:
		return *tv, nil
	case *common.ExternalCategoryPurpose1Code:
		return *tv, nil
	case *common.ExternalClearingSystemIdentification1Code:
		return *tv, nil
	case *common.ExternalDiscountAmountType1Code:
		return *tv, nil
	case *common.ExternalDocumentLineType1Code:
		return *tv, nil
	case *common.ExternalFinancialInstitutionIdentification1Code:
		return *tv, nil
	case *common.ExternalGarnishmentType1Code:
		return *tv, nil
	case *common.ExternalLocalInstrument1Code:
		return *tv, nil
	case *common.ExternalMandateSetupReason1Code:
		return *tv, nil
	case *common.ExternalOrganisationIdentification1Code:
		return *tv, nil
	case *common.ExternalPaymentGroupStatus1Code:
		return *tv, nil
	case *common.ExternalPaymentTransactionStatus1Code:
		return *tv, nil
	case *common.ExternalPersonIdentification1Code:
		return *tv, nil
	case *common.ExternalProxyAccountType1Code:
		return *tv, nil
	case *common.ExternalPurpose1Code:
		return *tv, nil
	case *common.ExternalServiceLevel1Code:
		return *tv, nil
	case *common.ExternalStatusReason1Code:
		return *tv, nil
	case *common.ExternalTaxAmountType1Code:
		return *tv, nil
	case *common.Frequency6Code:
		return *tv, nil
	case *common.IBAN2007Identifier:
		return *tv, nil
	case *common.ISODate:
		return *tv, nil
	case *common.ISODateTime:
		return *tv, nil
	case *common.LEIIdentifier:
		return *tv, nil
	case *common.Max1025Text:
		return *tv, nil
	case *common.Max105Text:
		return *tv, nil
	case *common.Max128Text:
		return *tv, nil
	case *common.Max140Text:
		return *tv, nil
	case *common.Max15NumericText:
		return *tv, nil
	case *common.Max16Text:
		return *tv, nil
	case *common.Max2048Text:
		return *tv, nil
	case *common.Max34Text:
		return *tv, nil
	case *common.Max350Text:
		return *tv, nil
	case *common.Max35Text:
		return *tv, nil
	case *common.Max4Text:
		return *tv, nil
	case *common.Max70Text:
		return *tv, nil
	case *common.NamePrefix2Code:
		return *tv, nil
	case *common.PaymentMethod4Code:
		return *tv, nil
	case *common.PhoneNumber:
		return *tv, nil
	case *common.PreferredContactMethod1Code:
		return *tv, nil
	case *common.Priority2Code:
		return *tv, nil
	case *common.SequenceType3Code:
		return *tv, nil
	case *common.SettlementMethod1Code:
		return *tv, nil
	case *common.TaxRecordPeriod1Code:
		return *tv, nil
	case *common.UUIDv4Identifier:
		return *tv, nil
	case *xsdt.AnyType:
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
