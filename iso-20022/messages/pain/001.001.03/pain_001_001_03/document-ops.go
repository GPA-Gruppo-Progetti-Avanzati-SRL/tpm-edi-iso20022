// Package pain_001_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_001_001_03

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/rs/zerolog/log"
	"reflect"
	"strings"
)

type SetOpOptions struct {
	logVerbose  bool
	skipIfEmpty bool
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

// Get
// Deprecated: the simply calls the GetProperty method
func (d *Document) Get(path string, opts ...GetOpOption) (interface{}, error) {
	return d.GetProperty(path, opts...)
}

func (d *Document) GetProperty(path string, opts ...GetOpOption) (interface{}, error) {

	const semLogContext = "pain_001_001_03::get-property"

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

	const semLogContext = "pain_001_001_03::copy-to-dest"

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

		*typedDest, err = common.ToActiveOrHistoricCurrencyCode(src)
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

		*typedDest, err = common.ToAddressType2Code(src)
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

		*typedDest, err = common.ToAnyBICIdentifier(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Authorisation1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Authorisation1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToAuthorisation1Code(src)
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

		*typedDest, err = common.ToBICIdentifier(src)
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

		*typedDest, err = common.ToCashAccountType4Code(src)
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

		*typedDest, err = common.ToChargeBearerType1Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ChequeDelivery1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ChequeDelivery1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToChequeDelivery1Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ChequeType2Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ChequeType2Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToChequeType2Code(src)
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

		*typedDest, err = common.ToCountryCode(src)
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

		*typedDest, err = common.ToCreditDebitCode(src)
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

		*typedDest, err = common.ToDocumentType3Code(src)
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

		*typedDest, err = common.ToDocumentType5Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.ExchangeRateType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.ExchangeRateType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToExchangeRateType1Code(src)
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

		*typedDest, err = common.ToExternalAccountIdentification1Code(src)
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

		*typedDest, err = common.ToExternalCategoryPurpose1Code(src)
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

		*typedDest, err = common.ToExternalClearingSystemIdentification1Code(src)
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

		*typedDest, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
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

		*typedDest, err = common.ToExternalLocalInstrument1Code(src)
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

		*typedDest, err = common.ToExternalOrganisationIdentification1Code(src)
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

		*typedDest, err = common.ToExternalPersonIdentification1Code(src)
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

		*typedDest, err = common.ToExternalPurpose1Code(src)
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

		*typedDest, err = common.ToExternalServiceLevel1Code(src)
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

		*typedDest, err = common.ToIBAN2007Identifier(src)
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

		*typedDest, err = common.ToISODate(src)
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

		*typedDest, err = common.ToISODateTime(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Instruction3Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Instruction3Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToInstruction3Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.Max10Text:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.Max10Text data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToMax10Text(src)
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

		*typedDest, err = common.ToMax128Text(src)
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

		*typedDest, err = common.ToMax140Text(src)
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

		*typedDest, err = common.ToMax15NumericText(src)
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

		*typedDest, err = common.ToMax16Text(src)
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

		*typedDest, err = common.ToMax2048Text(src)
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

		*typedDest, err = common.ToMax34Text(src)
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

		*typedDest, err = common.ToMax35Text(src)
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

		*typedDest, err = common.ToMax4Text(src)
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

		*typedDest, err = common.ToMax70Text(src)
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

		*typedDest, err = common.ToNamePrefix1Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.PaymentMethod3Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.PaymentMethod3Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToPaymentMethod3Code(src)
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

		*typedDest, err = common.ToPhoneNumber(src)
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

		*typedDest, err = common.ToPriority2Code(src)
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *common.RegulatoryReportingType1Code:
		if typedDest == nil {
			err = errors.New("nil pointer... in unmarshalling common.RegulatoryReportingType1Code data for" + string(docPath))
			if err != nil && options.logVerbose {
				log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
			}
			return err
		}

		*typedDest, err = common.ToRegulatoryReportingType1Code(src)
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

		*typedDest, err = common.ToRemittanceLocationMethod2Code(src)
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

		*typedDest, err = common.ToTaxRecordPeriod1Code(src)
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

		*typedDest, err = xsdt.ToBoolean(src)
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

		*typedDest, err = xsdt.ToDecimal(src)
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
	case *common.Authorisation1Code:
		return *tv, nil
	case *common.BICIdentifier:
		return *tv, nil
	case *common.CashAccountType4Code:
		return *tv, nil
	case *common.ChargeBearerType1Code:
		return *tv, nil
	case *common.ChequeDelivery1Code:
		return *tv, nil
	case *common.ChequeType2Code:
		return *tv, nil
	case *common.CountryCode:
		return *tv, nil
	case *common.CreditDebitCode:
		return *tv, nil
	case *common.DocumentType3Code:
		return *tv, nil
	case *common.DocumentType5Code:
		return *tv, nil
	case *common.ExchangeRateType1Code:
		return *tv, nil
	case *common.ExternalAccountIdentification1Code:
		return *tv, nil
	case *common.ExternalCategoryPurpose1Code:
		return *tv, nil
	case *common.ExternalClearingSystemIdentification1Code:
		return *tv, nil
	case *common.ExternalFinancialInstitutionIdentification1Code:
		return *tv, nil
	case *common.ExternalLocalInstrument1Code:
		return *tv, nil
	case *common.ExternalOrganisationIdentification1Code:
		return *tv, nil
	case *common.ExternalPersonIdentification1Code:
		return *tv, nil
	case *common.ExternalPurpose1Code:
		return *tv, nil
	case *common.ExternalServiceLevel1Code:
		return *tv, nil
	case *common.IBAN2007Identifier:
		return *tv, nil
	case *common.ISODate:
		return *tv, nil
	case *common.ISODateTime:
		return *tv, nil
	case *common.Instruction3Code:
		return *tv, nil
	case *common.Max10Text:
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
	case *common.Max35Text:
		return *tv, nil
	case *common.Max4Text:
		return *tv, nil
	case *common.Max70Text:
		return *tv, nil
	case *common.NamePrefix1Code:
		return *tv, nil
	case *common.PaymentMethod3Code:
		return *tv, nil
	case *common.PhoneNumber:
		return *tv, nil
	case *common.Priority2Code:
		return *tv, nil
	case *common.RegulatoryReportingType1Code:
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
