// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

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

	const semLogContext = "pain_013_001_07::copy-to-dest"

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
	case **AmountOrRate1Choice:
		if typedStruct, ok := src.(AmountOrRate1Choice); ok {
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
	case **Cheque11:
		if typedStruct, ok := src.(Cheque11); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ChequeDeliveryMethod1Choice:
		if typedStruct, ok := src.(ChequeDeliveryMethod1Choice); ok {
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
	case **CreditorPaymentActivationRequestV07:
		if typedStruct, ok := src.(CreditorPaymentActivationRequestV07); ok {
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
	case **DocumentFormat1Choice:
		if typedStruct, ok := src.(DocumentFormat1Choice); ok {
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
	case **DocumentType1Choice:
		if typedStruct, ok := src.(DocumentType1Choice); ok {
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
	case **GenericIdentification1:
		if typedStruct, ok := src.(GenericIdentification1); ok {
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
	case **GroupHeader78:
		if typedStruct, ok := src.(GroupHeader78); ok {
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
	case **NameAndAddress16:
		if typedStruct, ok := src.(NameAndAddress16); ok {
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
	case **PartyAndSignature3:
		if typedStruct, ok := src.(PartyAndSignature3); ok {
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
	case **PaymentCondition1:
		if typedStruct, ok := src.(PaymentCondition1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentIdentification6:
		if typedStruct, ok := src.(PaymentIdentification6); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentTypeInformation26:
		if typedStruct, ok := src.(PaymentTypeInformation26); ok {
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
	case **RegulatoryAuthority2:
		if typedStruct, ok := src.(RegulatoryAuthority2); ok {
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
	case **SkipPayload:
		if typedStruct, ok := src.(SkipPayload); ok {
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
	case **TaxInformation8:
		if typedStruct, ok := src.(TaxInformation8); ok {
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
	case **CreditTransferTransaction35:
		if typedStruct, ok := src.(CreditTransferTransaction35); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DiscountAmountAndType1:
		if typedStruct, ok := src.(DiscountAmountAndType1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Document12:
		if typedStruct, ok := src.(Document12); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DocumentAdjustment1:
		if typedStruct, ok := src.(DocumentAdjustment1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DocumentLineIdentification1:
		if typedStruct, ok := src.(DocumentLineIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DocumentLineInformation1:
		if typedStruct, ok := src.(DocumentLineInformation1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GenericOrganisationIdentification1:
		if typedStruct, ok := src.(GenericOrganisationIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GenericPersonIdentification1:
		if typedStruct, ok := src.(GenericPersonIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **InstructionForCreditorAgent1:
		if typedStruct, ok := src.(InstructionForCreditorAgent1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OtherContact1:
		if typedStruct, ok := src.(OtherContact1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentInstruction31:
		if typedStruct, ok := src.(PaymentInstruction31); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReferredDocumentInformation7:
		if typedStruct, ok := src.(ReferredDocumentInformation7); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RegulatoryReporting3:
		if typedStruct, ok := src.(RegulatoryReporting3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceLocation7:
		if typedStruct, ok := src.(RemittanceLocation7); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceLocationData1:
		if typedStruct, ok := src.(RemittanceLocationData1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ServiceLevel8Choice:
		if typedStruct, ok := src.(ServiceLevel8Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **StructuredRegulatoryReporting3:
		if typedStruct, ok := src.(StructuredRegulatoryReporting3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **StructuredRemittanceInformation16:
		if typedStruct, ok := src.(StructuredRemittanceInformation16); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **SupplementaryData1:
		if typedStruct, ok := src.(SupplementaryData1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxAmountAndType1:
		if typedStruct, ok := src.(TaxAmountAndType1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxRecord2:
		if typedStruct, ok := src.(TaxRecord2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxRecordDetails2:
		if typedStruct, ok := src.(TaxRecordDetails2); ok {
			*typedDest = &typedStruct
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
	const semLogContext = "pain_013_001_07::clear-node"

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

	const semLogContext = "pain_013_001_07::clear-node"

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
	case **AmountOrRate1Choice:
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
	case **Cheque11:
		*typedDest = nil
	case **ChequeDeliveryMethod1Choice:
		*typedDest = nil
	case **ClearingSystemIdentification2Choice:
		*typedDest = nil
	case **ClearingSystemMemberIdentification2:
		*typedDest = nil
	case **Contact4:
		*typedDest = nil
	case **CreditorPaymentActivationRequestV07:
		*typedDest = nil
	case **CreditorReferenceInformation2:
		*typedDest = nil
	case **CreditorReferenceType1Choice:
		*typedDest = nil
	case **CreditorReferenceType2:
		*typedDest = nil
	case **DateAndDateTime2Choice:
		*typedDest = nil
	case **DateAndPlaceOfBirth1:
		*typedDest = nil
	case **DatePeriod2:
		*typedDest = nil
	case **DiscountAmountType1Choice:
		*typedDest = nil
	case **DocumentFormat1Choice:
		*typedDest = nil
	case **DocumentLineType1:
		*typedDest = nil
	case **DocumentLineType1Choice:
		*typedDest = nil
	case **DocumentType1Choice:
		*typedDest = nil
	case **EquivalentAmount2:
		*typedDest = nil
	case **FinancialIdentificationSchemeName1Choice:
		*typedDest = nil
	case **FinancialInstitutionIdentification18:
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
	case **GenericIdentification1:
		*typedDest = nil
	case **GenericIdentification30:
		*typedDest = nil
	case **GroupHeader78:
		*typedDest = nil
	case **LocalInstrument2Choice:
		*typedDest = nil
	case **NameAndAddress16:
		*typedDest = nil
	case **OrganisationIdentification29:
		*typedDest = nil
	case **OrganisationIdentificationSchemeName1Choice:
		*typedDest = nil
	case **Party38Choice:
		*typedDest = nil
	case **PartyAndSignature3:
		*typedDest = nil
	case **PartyIdentification135:
		*typedDest = nil
	case **PaymentCondition1:
		*typedDest = nil
	case **PaymentIdentification6:
		*typedDest = nil
	case **PaymentTypeInformation26:
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
	case **RegulatoryAuthority2:
		*typedDest = nil
	case **RemittanceAmount2:
		*typedDest = nil
	case **RemittanceAmount3:
		*typedDest = nil
	case **RemittanceInformation16:
		*typedDest = nil
	case **SkipPayload:
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
	case **TaxInformation8:
		*typedDest = nil
	case **TaxParty1:
		*typedDest = nil
	case **TaxParty2:
		*typedDest = nil
	case **TaxPeriod2:
		*typedDest = nil
	case **CreditTransferTransaction35:
		*typedDest = nil
	case **DiscountAmountAndType1:
		*typedDest = nil
	case **Document12:
		*typedDest = nil
	case **DocumentAdjustment1:
		*typedDest = nil
	case **DocumentLineIdentification1:
		*typedDest = nil
	case **DocumentLineInformation1:
		*typedDest = nil
	case **GenericOrganisationIdentification1:
		*typedDest = nil
	case **GenericPersonIdentification1:
		*typedDest = nil
	case **InstructionForCreditorAgent1:
		*typedDest = nil
	case **OtherContact1:
		*typedDest = nil
	case **PaymentInstruction31:
		*typedDest = nil
	case **ReferredDocumentInformation7:
		*typedDest = nil
	case **RegulatoryReporting3:
		*typedDest = nil
	case **RemittanceLocation7:
		*typedDest = nil
	case **RemittanceLocationData1:
		*typedDest = nil
	case **ServiceLevel8Choice:
		*typedDest = nil
	case **StructuredRegulatoryReporting3:
		*typedDest = nil
	case **StructuredRemittanceInformation16:
		*typedDest = nil
	case **SupplementaryData1:
		*typedDest = nil
	case **TaxAmountAndType1:
		*typedDest = nil
	case **TaxRecord2:
		*typedDest = nil
	case **TaxRecordDetails2:
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
	case **AmountOrRate1Choice:
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
	case **Cheque11:
		return *tv, nil
	case **ChequeDeliveryMethod1Choice:
		return *tv, nil
	case **ClearingSystemIdentification2Choice:
		return *tv, nil
	case **ClearingSystemMemberIdentification2:
		return *tv, nil
	case **Contact4:
		return *tv, nil
	case **CreditorPaymentActivationRequestV07:
		return *tv, nil
	case **CreditorReferenceInformation2:
		return *tv, nil
	case **CreditorReferenceType1Choice:
		return *tv, nil
	case **CreditorReferenceType2:
		return *tv, nil
	case **DateAndDateTime2Choice:
		return *tv, nil
	case **DateAndPlaceOfBirth1:
		return *tv, nil
	case **DatePeriod2:
		return *tv, nil
	case **DiscountAmountType1Choice:
		return *tv, nil
	case **DocumentFormat1Choice:
		return *tv, nil
	case **DocumentLineType1:
		return *tv, nil
	case **DocumentLineType1Choice:
		return *tv, nil
	case **DocumentType1Choice:
		return *tv, nil
	case **EquivalentAmount2:
		return *tv, nil
	case **FinancialIdentificationSchemeName1Choice:
		return *tv, nil
	case **FinancialInstitutionIdentification18:
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
	case **GenericIdentification1:
		return *tv, nil
	case **GenericIdentification30:
		return *tv, nil
	case **GroupHeader78:
		return *tv, nil
	case **LocalInstrument2Choice:
		return *tv, nil
	case **NameAndAddress16:
		return *tv, nil
	case **OrganisationIdentification29:
		return *tv, nil
	case **OrganisationIdentificationSchemeName1Choice:
		return *tv, nil
	case **Party38Choice:
		return *tv, nil
	case **PartyAndSignature3:
		return *tv, nil
	case **PartyIdentification135:
		return *tv, nil
	case **PaymentCondition1:
		return *tv, nil
	case **PaymentIdentification6:
		return *tv, nil
	case **PaymentTypeInformation26:
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
	case **RegulatoryAuthority2:
		return *tv, nil
	case **RemittanceAmount2:
		return *tv, nil
	case **RemittanceAmount3:
		return *tv, nil
	case **RemittanceInformation16:
		return *tv, nil
	case **SkipPayload:
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
	case **TaxInformation8:
		return *tv, nil
	case **TaxParty1:
		return *tv, nil
	case **TaxParty2:
		return *tv, nil
	case **TaxPeriod2:
		return *tv, nil
	case **CreditTransferTransaction35:
		return *tv, nil
	case **DiscountAmountAndType1:
		return *tv, nil
	case **Document12:
		return *tv, nil
	case **DocumentAdjustment1:
		return *tv, nil
	case **DocumentLineIdentification1:
		return *tv, nil
	case **DocumentLineInformation1:
		return *tv, nil
	case **GenericOrganisationIdentification1:
		return *tv, nil
	case **GenericPersonIdentification1:
		return *tv, nil
	case **InstructionForCreditorAgent1:
		return *tv, nil
	case **OtherContact1:
		return *tv, nil
	case **PaymentInstruction31:
		return *tv, nil
	case **ReferredDocumentInformation7:
		return *tv, nil
	case **RegulatoryReporting3:
		return *tv, nil
	case **RemittanceLocation7:
		return *tv, nil
	case **RemittanceLocationData1:
		return *tv, nil
	case **ServiceLevel8Choice:
		return *tv, nil
	case **StructuredRegulatoryReporting3:
		return *tv, nil
	case **StructuredRemittanceInformation16:
		return *tv, nil
	case **SupplementaryData1:
		return *tv, nil
	case **TaxAmountAndType1:
		return *tv, nil
	case **TaxRecord2:
		return *tv, nil
	case **TaxRecordDetails2:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), val)
	}

	return val, err
}
