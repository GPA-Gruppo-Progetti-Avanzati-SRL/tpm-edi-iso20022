// Package pain_002_001_10
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_10

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/rs/zerolog/log"
	"reflect"
)

func (d *Document) SetNode(path string, src interface{}, opts ...SetOpOption) error {

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
			log.Error().Err(err).Str("path", string(path)).Interface("value", src).Msg(semLogContext)
		}

		return err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, false)
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(path)).Interface("value", src).Msg(semLogContext)
		}
		return err
	}

	return copy2DestNode(path, values[0], src, &options)
}

func copy2DestNode(docPath string, dest, src interface{}, options *SetOpOptions) error {

	const semLogContext = "pain_002_001_10::copy-to-dest"

	var err error
	switch typedDest := dest.(type) {
	case **AccountIdentification4Choice:
		if typedStruct, ok := src.(AccountIdentification4Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AccountSchemeName1Choice:
		if typedStruct, ok := src.(AccountSchemeName1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ActiveCurrencyAndAmount:
		if typedStruct, ok := src.(ActiveCurrencyAndAmount); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ActiveOrHistoricCurrencyAndAmount:
		if typedStruct, ok := src.(ActiveOrHistoricCurrencyAndAmount); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AddressType3Choice:
		if typedStruct, ok := src.(AddressType3Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmendmentInformationDetails13:
		if typedStruct, ok := src.(AmendmentInformationDetails13); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmountType4Choice:
		if typedStruct, ok := src.(AmountType4Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BranchAndFinancialInstitutionIdentification6:
		if typedStruct, ok := src.(BranchAndFinancialInstitutionIdentification6); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BranchData3:
		if typedStruct, ok := src.(BranchData3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashAccount38:
		if typedStruct, ok := src.(CashAccount38); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashAccountType2Choice:
		if typedStruct, ok := src.(CashAccountType2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CategoryPurpose1Choice:
		if typedStruct, ok := src.(CategoryPurpose1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ClearingSystemIdentification2Choice:
		if typedStruct, ok := src.(ClearingSystemIdentification2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ClearingSystemIdentification3Choice:
		if typedStruct, ok := src.(ClearingSystemIdentification3Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ClearingSystemMemberIdentification2:
		if typedStruct, ok := src.(ClearingSystemMemberIdentification2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Contact4:
		if typedStruct, ok := src.(Contact4); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CreditorReferenceInformation2:
		if typedStruct, ok := src.(CreditorReferenceInformation2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CreditorReferenceType1Choice:
		if typedStruct, ok := src.(CreditorReferenceType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CreditorReferenceType2:
		if typedStruct, ok := src.(CreditorReferenceType2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CurrencyExchange13:
		if typedStruct, ok := src.(CurrencyExchange13); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CustomerPaymentStatusReportV10:
		if typedStruct, ok := src.(CustomerPaymentStatusReportV10); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DateAndDateTime2Choice:
		if typedStruct, ok := src.(DateAndDateTime2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DateAndPlaceOfBirth1:
		if typedStruct, ok := src.(DateAndPlaceOfBirth1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DatePeriod2:
		if typedStruct, ok := src.(DatePeriod2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DiscountAmountType1Choice:
		if typedStruct, ok := src.(DiscountAmountType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DocumentLineType1:
		if typedStruct, ok := src.(DocumentLineType1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DocumentLineType1Choice:
		if typedStruct, ok := src.(DocumentLineType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **EquivalentAmount2:
		if typedStruct, ok := src.(EquivalentAmount2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **FinancialIdentificationSchemeName1Choice:
		if typedStruct, ok := src.(FinancialIdentificationSchemeName1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **FinancialInstitutionIdentification18:
		if typedStruct, ok := src.(FinancialInstitutionIdentification18); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Frequency36Choice:
		if typedStruct, ok := src.(Frequency36Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **FrequencyAndMoment1:
		if typedStruct, ok := src.(FrequencyAndMoment1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **FrequencyPeriod1:
		if typedStruct, ok := src.(FrequencyPeriod1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Garnishment3:
		if typedStruct, ok := src.(Garnishment3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GarnishmentType1:
		if typedStruct, ok := src.(GarnishmentType1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GarnishmentType1Choice:
		if typedStruct, ok := src.(GarnishmentType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GenericAccountIdentification1:
		if typedStruct, ok := src.(GenericAccountIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GenericFinancialIdentification1:
		if typedStruct, ok := src.(GenericFinancialIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GenericIdentification30:
		if typedStruct, ok := src.(GenericIdentification30); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GroupHeader86:
		if typedStruct, ok := src.(GroupHeader86); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **LocalInstrument2Choice:
		if typedStruct, ok := src.(LocalInstrument2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **MandateRelatedInformation14:
		if typedStruct, ok := src.(MandateRelatedInformation14); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **MandateSetupReason1Choice:
		if typedStruct, ok := src.(MandateSetupReason1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OrganisationIdentification29:
		if typedStruct, ok := src.(OrganisationIdentification29); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OrganisationIdentificationSchemeName1Choice:
		if typedStruct, ok := src.(OrganisationIdentificationSchemeName1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OriginalGroupHeader17:
		if typedStruct, ok := src.(OriginalGroupHeader17); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OriginalTransactionReference28:
		if typedStruct, ok := src.(OriginalTransactionReference28); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Party38Choice:
		if typedStruct, ok := src.(Party38Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Party40Choice:
		if typedStruct, ok := src.(Party40Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PartyIdentification135:
		if typedStruct, ok := src.(PartyIdentification135); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentTypeInformation27:
		if typedStruct, ok := src.(PaymentTypeInformation27); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PersonIdentification13:
		if typedStruct, ok := src.(PersonIdentification13); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PersonIdentificationSchemeName1Choice:
		if typedStruct, ok := src.(PersonIdentificationSchemeName1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PostalAddress24:
		if typedStruct, ok := src.(PostalAddress24); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProxyAccountIdentification1:
		if typedStruct, ok := src.(ProxyAccountIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProxyAccountType1Choice:
		if typedStruct, ok := src.(ProxyAccountType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Purpose2Choice:
		if typedStruct, ok := src.(Purpose2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReferredDocumentType3Choice:
		if typedStruct, ok := src.(ReferredDocumentType3Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReferredDocumentType4:
		if typedStruct, ok := src.(ReferredDocumentType4); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceAmount2:
		if typedStruct, ok := src.(RemittanceAmount2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceAmount3:
		if typedStruct, ok := src.(RemittanceAmount3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceInformation16:
		if typedStruct, ok := src.(RemittanceInformation16); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **SettlementInstruction7:
		if typedStruct, ok := src.(SettlementInstruction7); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **StatusReason6Choice:
		if typedStruct, ok := src.(StatusReason6Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **SupplementaryDataEnvelope1:
		if typedStruct, ok := src.(SupplementaryDataEnvelope1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxAmount2:
		if typedStruct, ok := src.(TaxAmount2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxAmountType1Choice:
		if typedStruct, ok := src.(TaxAmountType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxAuthorisation1:
		if typedStruct, ok := src.(TaxAuthorisation1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxInformation7:
		if typedStruct, ok := src.(TaxInformation7); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxParty1:
		if typedStruct, ok := src.(TaxParty1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxParty2:
		if typedStruct, ok := src.(TaxParty2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxPeriod2:
		if typedStruct, ok := src.(TaxPeriod2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TrackerData1:
		if typedStruct, ok := src.(TrackerData1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *Charges7:
		if typedStruct, ok := src.(Charges7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]Charges7:
		if typedStruct, ok := src.([]Charges7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *DiscountAmountAndType1:
		if typedStruct, ok := src.(DiscountAmountAndType1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]DiscountAmountAndType1:
		if typedStruct, ok := src.([]DiscountAmountAndType1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *DocumentAdjustment1:
		if typedStruct, ok := src.(DocumentAdjustment1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]DocumentAdjustment1:
		if typedStruct, ok := src.([]DocumentAdjustment1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *DocumentLineIdentification1:
		if typedStruct, ok := src.(DocumentLineIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]DocumentLineIdentification1:
		if typedStruct, ok := src.([]DocumentLineIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *DocumentLineInformation1:
		if typedStruct, ok := src.(DocumentLineInformation1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]DocumentLineInformation1:
		if typedStruct, ok := src.([]DocumentLineInformation1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *GenericOrganisationIdentification1:
		if typedStruct, ok := src.(GenericOrganisationIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]GenericOrganisationIdentification1:
		if typedStruct, ok := src.([]GenericOrganisationIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *GenericPersonIdentification1:
		if typedStruct, ok := src.(GenericPersonIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]GenericPersonIdentification1:
		if typedStruct, ok := src.([]GenericPersonIdentification1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *NumberOfTransactionsPerStatus5:
		if typedStruct, ok := src.(NumberOfTransactionsPerStatus5); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]NumberOfTransactionsPerStatus5:
		if typedStruct, ok := src.([]NumberOfTransactionsPerStatus5); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *OriginalPaymentInstruction32:
		if typedStruct, ok := src.(OriginalPaymentInstruction32); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]OriginalPaymentInstruction32:
		if typedStruct, ok := src.([]OriginalPaymentInstruction32); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *OtherContact1:
		if typedStruct, ok := src.(OtherContact1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]OtherContact1:
		if typedStruct, ok := src.([]OtherContact1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *PaymentTransaction105:
		if typedStruct, ok := src.(PaymentTransaction105); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]PaymentTransaction105:
		if typedStruct, ok := src.([]PaymentTransaction105); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *ReferredDocumentInformation7:
		if typedStruct, ok := src.(ReferredDocumentInformation7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]ReferredDocumentInformation7:
		if typedStruct, ok := src.([]ReferredDocumentInformation7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *ServiceLevel8Choice:
		if typedStruct, ok := src.(ServiceLevel8Choice); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]ServiceLevel8Choice:
		if typedStruct, ok := src.([]ServiceLevel8Choice); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *StatusReasonInformation12:
		if typedStruct, ok := src.(StatusReasonInformation12); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]StatusReasonInformation12:
		if typedStruct, ok := src.([]StatusReasonInformation12); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *StructuredRemittanceInformation16:
		if typedStruct, ok := src.(StructuredRemittanceInformation16); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]StructuredRemittanceInformation16:
		if typedStruct, ok := src.([]StructuredRemittanceInformation16); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *SupplementaryData1:
		if typedStruct, ok := src.(SupplementaryData1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]SupplementaryData1:
		if typedStruct, ok := src.([]SupplementaryData1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TaxAmountAndType1:
		if typedStruct, ok := src.(TaxAmountAndType1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TaxAmountAndType1:
		if typedStruct, ok := src.([]TaxAmountAndType1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TaxRecord2:
		if typedStruct, ok := src.(TaxRecord2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TaxRecord2:
		if typedStruct, ok := src.([]TaxRecord2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TaxRecordDetails2:
		if typedStruct, ok := src.(TaxRecordDetails2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TaxRecordDetails2:
		if typedStruct, ok := src.([]TaxRecordDetails2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TrackerRecord1:
		if typedStruct, ok := src.(TrackerRecord1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TrackerRecord1:
		if typedStruct, ok := src.([]TrackerRecord1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), dest)
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	}

}

func (d *Document) GetNode(path string, opts ...GetOpOption) (interface{}, error) {

	const semLogContext = "camt_053_001_02::get-node"

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

	return derefNode(path, values[0])
}

func (d *Document) ClearNode(path string, opts ...SetOpOption) error {
	const semLogContext = "pain_002_001_10::clear-node"

	options := SetOpOptions{}
	for _, o := range opts {
		o(&options)
	}

	i, err := d.GetNode(path)
	if err != nil {
		return err
	}

	if i == nil {
		return nil
	}

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(string(path))
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(path)).Msg(semLogContext)
		}

		return err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, false)
	if err != nil {
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(path)).Msg(semLogContext)
		}
		return err
	}

	return clearNode(path, values[0], &options)

}

func clearNode(docPath string, dest interface{}, options *SetOpOptions) error {

	const semLogContext = "pain_002_001_10::clear-node"

	var err error
	switch typedDest := dest.(type) {
	case **AccountIdentification4Choice:
		*typedDest = nil
	case **AccountSchemeName1Choice:
		*typedDest = nil
	case **ActiveCurrencyAndAmount:
		*typedDest = nil
	case **ActiveOrHistoricCurrencyAndAmount:
		*typedDest = nil
	case **AddressType3Choice:
		*typedDest = nil
	case **AmendmentInformationDetails13:
		*typedDest = nil
	case **AmountType4Choice:
		*typedDest = nil
	case **BranchAndFinancialInstitutionIdentification6:
		*typedDest = nil
	case **BranchData3:
		*typedDest = nil
	case **CashAccount38:
		*typedDest = nil
	case **CashAccountType2Choice:
		*typedDest = nil
	case **CategoryPurpose1Choice:
		*typedDest = nil
	case **ClearingSystemIdentification2Choice:
		*typedDest = nil
	case **ClearingSystemIdentification3Choice:
		*typedDest = nil
	case **ClearingSystemMemberIdentification2:
		*typedDest = nil
	case **Contact4:
		*typedDest = nil
	case **CreditorReferenceInformation2:
		*typedDest = nil
	case **CreditorReferenceType1Choice:
		*typedDest = nil
	case **CreditorReferenceType2:
		*typedDest = nil
	case **CurrencyExchange13:
		*typedDest = nil
	case **CustomerPaymentStatusReportV10:
		*typedDest = nil
	case **DateAndDateTime2Choice:
		*typedDest = nil
	case **DateAndPlaceOfBirth1:
		*typedDest = nil
	case **DatePeriod2:
		*typedDest = nil
	case **DiscountAmountType1Choice:
		*typedDest = nil
	case **DocumentLineType1:
		*typedDest = nil
	case **DocumentLineType1Choice:
		*typedDest = nil
	case **EquivalentAmount2:
		*typedDest = nil
	case **FinancialIdentificationSchemeName1Choice:
		*typedDest = nil
	case **FinancialInstitutionIdentification18:
		*typedDest = nil
	case **Frequency36Choice:
		*typedDest = nil
	case **FrequencyAndMoment1:
		*typedDest = nil
	case **FrequencyPeriod1:
		*typedDest = nil
	case **Garnishment3:
		*typedDest = nil
	case **GarnishmentType1:
		*typedDest = nil
	case **GarnishmentType1Choice:
		*typedDest = nil
	case **GenericAccountIdentification1:
		*typedDest = nil
	case **GenericFinancialIdentification1:
		*typedDest = nil
	case **GenericIdentification30:
		*typedDest = nil
	case **GroupHeader86:
		*typedDest = nil
	case **LocalInstrument2Choice:
		*typedDest = nil
	case **MandateRelatedInformation14:
		*typedDest = nil
	case **MandateSetupReason1Choice:
		*typedDest = nil
	case **OrganisationIdentification29:
		*typedDest = nil
	case **OrganisationIdentificationSchemeName1Choice:
		*typedDest = nil
	case **OriginalGroupHeader17:
		*typedDest = nil
	case **OriginalTransactionReference28:
		*typedDest = nil
	case **Party38Choice:
		*typedDest = nil
	case **Party40Choice:
		*typedDest = nil
	case **PartyIdentification135:
		*typedDest = nil
	case **PaymentTypeInformation27:
		*typedDest = nil
	case **PersonIdentification13:
		*typedDest = nil
	case **PersonIdentificationSchemeName1Choice:
		*typedDest = nil
	case **PostalAddress24:
		*typedDest = nil
	case **ProxyAccountIdentification1:
		*typedDest = nil
	case **ProxyAccountType1Choice:
		*typedDest = nil
	case **Purpose2Choice:
		*typedDest = nil
	case **ReferredDocumentType3Choice:
		*typedDest = nil
	case **ReferredDocumentType4:
		*typedDest = nil
	case **RemittanceAmount2:
		*typedDest = nil
	case **RemittanceAmount3:
		*typedDest = nil
	case **RemittanceInformation16:
		*typedDest = nil
	case **SettlementInstruction7:
		*typedDest = nil
	case **StatusReason6Choice:
		*typedDest = nil
	case **SupplementaryDataEnvelope1:
		*typedDest = nil
	case **TaxAmount2:
		*typedDest = nil
	case **TaxAmountType1Choice:
		*typedDest = nil
	case **TaxAuthorisation1:
		*typedDest = nil
	case **TaxInformation7:
		*typedDest = nil
	case **TaxParty1:
		*typedDest = nil
	case **TaxParty2:
		*typedDest = nil
	case **TaxPeriod2:
		*typedDest = nil
	case **TrackerData1:
		*typedDest = nil
	case *Charges7:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]Charges7:
		*typedDest = nil
	case *DiscountAmountAndType1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DiscountAmountAndType1:
		*typedDest = nil
	case *DocumentAdjustment1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DocumentAdjustment1:
		*typedDest = nil
	case *DocumentLineIdentification1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DocumentLineIdentification1:
		*typedDest = nil
	case *DocumentLineInformation1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DocumentLineInformation1:
		*typedDest = nil
	case *GenericOrganisationIdentification1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]GenericOrganisationIdentification1:
		*typedDest = nil
	case *GenericPersonIdentification1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]GenericPersonIdentification1:
		*typedDest = nil
	case *NumberOfTransactionsPerStatus5:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]NumberOfTransactionsPerStatus5:
		*typedDest = nil
	case *OriginalPaymentInstruction32:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]OriginalPaymentInstruction32:
		*typedDest = nil
	case *OtherContact1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]OtherContact1:
		*typedDest = nil
	case *PaymentTransaction105:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]PaymentTransaction105:
		*typedDest = nil
	case *ReferredDocumentInformation7:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]ReferredDocumentInformation7:
		*typedDest = nil
	case *ServiceLevel8Choice:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]ServiceLevel8Choice:
		*typedDest = nil
	case *StatusReasonInformation12:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]StatusReasonInformation12:
		*typedDest = nil
	case *StructuredRemittanceInformation16:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]StructuredRemittanceInformation16:
		*typedDest = nil
	case *SupplementaryData1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]SupplementaryData1:
		*typedDest = nil
	case *TaxAmountAndType1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TaxAmountAndType1:
		*typedDest = nil
	case *TaxRecord2:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TaxRecord2:
		*typedDest = nil
	case *TaxRecordDetails2:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TaxRecordDetails2:
		*typedDest = nil
	case *TrackerRecord1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TrackerRecord1:
		*typedDest = nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), dest)
		if options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Msg(semLogContext)
		}
		return err
	}

	return nil
}

func derefNode(docPath string, val interface{}) (interface{}, error) {

	if val == nil {
		return nil, nil
	}

	var err error
	switch tv := val.(type) {
	case **AccountIdentification4Choice:
		return *tv, nil
	case **AccountSchemeName1Choice:
		return *tv, nil
	case **ActiveCurrencyAndAmount:
		return *tv, nil
	case **ActiveOrHistoricCurrencyAndAmount:
		return *tv, nil
	case **AddressType3Choice:
		return *tv, nil
	case **AmendmentInformationDetails13:
		return *tv, nil
	case **AmountType4Choice:
		return *tv, nil
	case **BranchAndFinancialInstitutionIdentification6:
		return *tv, nil
	case **BranchData3:
		return *tv, nil
	case **CashAccount38:
		return *tv, nil
	case **CashAccountType2Choice:
		return *tv, nil
	case **CategoryPurpose1Choice:
		return *tv, nil
	case **ClearingSystemIdentification2Choice:
		return *tv, nil
	case **ClearingSystemIdentification3Choice:
		return *tv, nil
	case **ClearingSystemMemberIdentification2:
		return *tv, nil
	case **Contact4:
		return *tv, nil
	case **CreditorReferenceInformation2:
		return *tv, nil
	case **CreditorReferenceType1Choice:
		return *tv, nil
	case **CreditorReferenceType2:
		return *tv, nil
	case **CurrencyExchange13:
		return *tv, nil
	case **CustomerPaymentStatusReportV10:
		return *tv, nil
	case **DateAndDateTime2Choice:
		return *tv, nil
	case **DateAndPlaceOfBirth1:
		return *tv, nil
	case **DatePeriod2:
		return *tv, nil
	case **DiscountAmountType1Choice:
		return *tv, nil
	case **DocumentLineType1:
		return *tv, nil
	case **DocumentLineType1Choice:
		return *tv, nil
	case **EquivalentAmount2:
		return *tv, nil
	case **FinancialIdentificationSchemeName1Choice:
		return *tv, nil
	case **FinancialInstitutionIdentification18:
		return *tv, nil
	case **Frequency36Choice:
		return *tv, nil
	case **FrequencyAndMoment1:
		return *tv, nil
	case **FrequencyPeriod1:
		return *tv, nil
	case **Garnishment3:
		return *tv, nil
	case **GarnishmentType1:
		return *tv, nil
	case **GarnishmentType1Choice:
		return *tv, nil
	case **GenericAccountIdentification1:
		return *tv, nil
	case **GenericFinancialIdentification1:
		return *tv, nil
	case **GenericIdentification30:
		return *tv, nil
	case **GroupHeader86:
		return *tv, nil
	case **LocalInstrument2Choice:
		return *tv, nil
	case **MandateRelatedInformation14:
		return *tv, nil
	case **MandateSetupReason1Choice:
		return *tv, nil
	case **OrganisationIdentification29:
		return *tv, nil
	case **OrganisationIdentificationSchemeName1Choice:
		return *tv, nil
	case **OriginalGroupHeader17:
		return *tv, nil
	case **OriginalTransactionReference28:
		return *tv, nil
	case **Party38Choice:
		return *tv, nil
	case **Party40Choice:
		return *tv, nil
	case **PartyIdentification135:
		return *tv, nil
	case **PaymentTypeInformation27:
		return *tv, nil
	case **PersonIdentification13:
		return *tv, nil
	case **PersonIdentificationSchemeName1Choice:
		return *tv, nil
	case **PostalAddress24:
		return *tv, nil
	case **ProxyAccountIdentification1:
		return *tv, nil
	case **ProxyAccountType1Choice:
		return *tv, nil
	case **Purpose2Choice:
		return *tv, nil
	case **ReferredDocumentType3Choice:
		return *tv, nil
	case **ReferredDocumentType4:
		return *tv, nil
	case **RemittanceAmount2:
		return *tv, nil
	case **RemittanceAmount3:
		return *tv, nil
	case **RemittanceInformation16:
		return *tv, nil
	case **SettlementInstruction7:
		return *tv, nil
	case **StatusReason6Choice:
		return *tv, nil
	case **SupplementaryDataEnvelope1:
		return *tv, nil
	case **TaxAmount2:
		return *tv, nil
	case **TaxAmountType1Choice:
		return *tv, nil
	case **TaxAuthorisation1:
		return *tv, nil
	case **TaxInformation7:
		return *tv, nil
	case **TaxParty1:
		return *tv, nil
	case **TaxParty2:
		return *tv, nil
	case **TaxPeriod2:
		return *tv, nil
	case **TrackerData1:
		return *tv, nil
	case *Charges7:
		return tv, nil
	case *[]Charges7:
		return *tv, nil
	case *DiscountAmountAndType1:
		return tv, nil
	case *[]DiscountAmountAndType1:
		return *tv, nil
	case *DocumentAdjustment1:
		return tv, nil
	case *[]DocumentAdjustment1:
		return *tv, nil
	case *DocumentLineIdentification1:
		return tv, nil
	case *[]DocumentLineIdentification1:
		return *tv, nil
	case *DocumentLineInformation1:
		return tv, nil
	case *[]DocumentLineInformation1:
		return *tv, nil
	case *GenericOrganisationIdentification1:
		return tv, nil
	case *[]GenericOrganisationIdentification1:
		return *tv, nil
	case *GenericPersonIdentification1:
		return tv, nil
	case *[]GenericPersonIdentification1:
		return *tv, nil
	case *NumberOfTransactionsPerStatus5:
		return tv, nil
	case *[]NumberOfTransactionsPerStatus5:
		return *tv, nil
	case *OriginalPaymentInstruction32:
		return tv, nil
	case *[]OriginalPaymentInstruction32:
		return *tv, nil
	case *OtherContact1:
		return tv, nil
	case *[]OtherContact1:
		return *tv, nil
	case *PaymentTransaction105:
		return tv, nil
	case *[]PaymentTransaction105:
		return *tv, nil
	case *ReferredDocumentInformation7:
		return tv, nil
	case *[]ReferredDocumentInformation7:
		return *tv, nil
	case *ServiceLevel8Choice:
		return tv, nil
	case *[]ServiceLevel8Choice:
		return *tv, nil
	case *StatusReasonInformation12:
		return tv, nil
	case *[]StatusReasonInformation12:
		return *tv, nil
	case *StructuredRemittanceInformation16:
		return tv, nil
	case *[]StructuredRemittanceInformation16:
		return *tv, nil
	case *SupplementaryData1:
		return tv, nil
	case *[]SupplementaryData1:
		return *tv, nil
	case *TaxAmountAndType1:
		return tv, nil
	case *[]TaxAmountAndType1:
		return *tv, nil
	case *TaxRecord2:
		return tv, nil
	case *[]TaxRecord2:
		return *tv, nil
	case *TaxRecordDetails2:
		return tv, nil
	case *[]TaxRecordDetails2:
		return *tv, nil
	case *TrackerRecord1:
		return tv, nil
	case *[]TrackerRecord1:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), val)
	}

	return val, err
}
