// Package pain_008_001_02
// Do not Edit. This stuff it's been automatically generated.
package pain_008_001_02

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

	const semLogContext = "pain_008_001_02::copy-to-dest"

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
	case **AmendmentInformationDetails6:
		if typedStruct, ok := src.(AmendmentInformationDetails6); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BranchAndFinancialInstitutionIdentification4:
		if typedStruct, ok := src.(BranchAndFinancialInstitutionIdentification4); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BranchData2:
		if typedStruct, ok := src.(BranchData2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashAccount16:
		if typedStruct, ok := src.(CashAccount16); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashAccountType2:
		if typedStruct, ok := src.(CashAccountType2); ok {
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
	case **ContactDetails2:
		if typedStruct, ok := src.(ContactDetails2); ok {
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
	case **CustomerDirectDebitInitiationV02:
		if typedStruct, ok := src.(CustomerDirectDebitInitiationV02); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DateAndPlaceOfBirth:
		if typedStruct, ok := src.(DateAndPlaceOfBirth); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DatePeriodDetails:
		if typedStruct, ok := src.(DatePeriodDetails); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DirectDebitTransaction6:
		if typedStruct, ok := src.(DirectDebitTransaction6); ok {
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
	case **FinancialInstitutionIdentification7:
		if typedStruct, ok := src.(FinancialInstitutionIdentification7); ok {
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
	case **GroupHeader39:
		if typedStruct, ok := src.(GroupHeader39); ok {
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
	case **MandateRelatedInformation6:
		if typedStruct, ok := src.(MandateRelatedInformation6); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **NameAndAddress10:
		if typedStruct, ok := src.(NameAndAddress10); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **OrganisationIdentification4:
		if typedStruct, ok := src.(OrganisationIdentification4); ok {
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
	case **Party6Choice:
		if typedStruct, ok := src.(Party6Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PartyIdentification32:
		if typedStruct, ok := src.(PartyIdentification32); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentIdentification1:
		if typedStruct, ok := src.(PaymentIdentification1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PaymentTypeInformation20:
		if typedStruct, ok := src.(PaymentTypeInformation20); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **PersonIdentification5:
		if typedStruct, ok := src.(PersonIdentification5); ok {
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
	case **PostalAddress6:
		if typedStruct, ok := src.(PostalAddress6); ok {
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
	case **ReferredDocumentType1Choice:
		if typedStruct, ok := src.(ReferredDocumentType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReferredDocumentType2:
		if typedStruct, ok := src.(ReferredDocumentType2); ok {
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
	case **RemittanceAmount1:
		if typedStruct, ok := src.(RemittanceAmount1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceInformation5:
		if typedStruct, ok := src.(RemittanceInformation5); ok {
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
	case **TaxAmount1:
		if typedStruct, ok := src.(TaxAmount1); ok {
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
	case **TaxInformation3:
		if typedStruct, ok := src.(TaxInformation3); ok {
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
	case **TaxPeriod1:
		if typedStruct, ok := src.(TaxPeriod1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *Authorisation1Choice:
		if typedStruct, ok := src.(Authorisation1Choice); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]Authorisation1Choice:
		if typedStruct, ok := src.([]Authorisation1Choice); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *DirectDebitTransactionInformation9:
		if typedStruct, ok := src.(DirectDebitTransactionInformation9); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]DirectDebitTransactionInformation9:
		if typedStruct, ok := src.([]DirectDebitTransactionInformation9); ok {
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
	case *PaymentInstructionInformation4:
		if typedStruct, ok := src.(PaymentInstructionInformation4); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]PaymentInstructionInformation4:
		if typedStruct, ok := src.([]PaymentInstructionInformation4); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *ReferredDocumentInformation3:
		if typedStruct, ok := src.(ReferredDocumentInformation3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]ReferredDocumentInformation3:
		if typedStruct, ok := src.([]ReferredDocumentInformation3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *RegulatoryReporting3:
		if typedStruct, ok := src.(RegulatoryReporting3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]RegulatoryReporting3:
		if typedStruct, ok := src.([]RegulatoryReporting3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *RemittanceLocation2:
		if typedStruct, ok := src.(RemittanceLocation2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]RemittanceLocation2:
		if typedStruct, ok := src.([]RemittanceLocation2); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *StructuredRegulatoryReporting3:
		if typedStruct, ok := src.(StructuredRegulatoryReporting3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]StructuredRegulatoryReporting3:
		if typedStruct, ok := src.([]StructuredRegulatoryReporting3); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *StructuredRemittanceInformation7:
		if typedStruct, ok := src.(StructuredRemittanceInformation7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]StructuredRemittanceInformation7:
		if typedStruct, ok := src.([]StructuredRemittanceInformation7); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TaxRecord1:
		if typedStruct, ok := src.(TaxRecord1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TaxRecord1:
		if typedStruct, ok := src.([]TaxRecord1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *TaxRecordDetails1:
		if typedStruct, ok := src.(TaxRecordDetails1); ok {
			*typedDest = typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case *[]TaxRecordDetails1:
		if typedStruct, ok := src.([]TaxRecordDetails1); ok {
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
	const semLogContext = "pain_008_001_02::clear-node"

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

	const semLogContext = "pain_008_001_02::clear-node"

	var err error
	switch typedDest := dest.(type) {
	case **AccountIdentification4Choice:
		*typedDest = nil
	case **AccountSchemeName1Choice:
		*typedDest = nil
	case **ActiveOrHistoricCurrencyAndAmount:
		*typedDest = nil
	case **AmendmentInformationDetails6:
		*typedDest = nil
	case **BranchAndFinancialInstitutionIdentification4:
		*typedDest = nil
	case **BranchData2:
		*typedDest = nil
	case **CashAccount16:
		*typedDest = nil
	case **CashAccountType2:
		*typedDest = nil
	case **CategoryPurpose1Choice:
		*typedDest = nil
	case **ClearingSystemIdentification2Choice:
		*typedDest = nil
	case **ClearingSystemMemberIdentification2:
		*typedDest = nil
	case **ContactDetails2:
		*typedDest = nil
	case **CreditorReferenceInformation2:
		*typedDest = nil
	case **CreditorReferenceType1Choice:
		*typedDest = nil
	case **CreditorReferenceType2:
		*typedDest = nil
	case **CustomerDirectDebitInitiationV02:
		*typedDest = nil
	case **DateAndPlaceOfBirth:
		*typedDest = nil
	case **DatePeriodDetails:
		*typedDest = nil
	case **DirectDebitTransaction6:
		*typedDest = nil
	case **FinancialIdentificationSchemeName1Choice:
		*typedDest = nil
	case **FinancialInstitutionIdentification7:
		*typedDest = nil
	case **GenericAccountIdentification1:
		*typedDest = nil
	case **GenericFinancialIdentification1:
		*typedDest = nil
	case **GroupHeader39:
		*typedDest = nil
	case **LocalInstrument2Choice:
		*typedDest = nil
	case **MandateRelatedInformation6:
		*typedDest = nil
	case **NameAndAddress10:
		*typedDest = nil
	case **OrganisationIdentification4:
		*typedDest = nil
	case **OrganisationIdentificationSchemeName1Choice:
		*typedDest = nil
	case **Party6Choice:
		*typedDest = nil
	case **PartyIdentification32:
		*typedDest = nil
	case **PaymentIdentification1:
		*typedDest = nil
	case **PaymentTypeInformation20:
		*typedDest = nil
	case **PersonIdentification5:
		*typedDest = nil
	case **PersonIdentificationSchemeName1Choice:
		*typedDest = nil
	case **PostalAddress6:
		*typedDest = nil
	case **Purpose2Choice:
		*typedDest = nil
	case **ReferredDocumentType1Choice:
		*typedDest = nil
	case **ReferredDocumentType2:
		*typedDest = nil
	case **RegulatoryAuthority2:
		*typedDest = nil
	case **RemittanceAmount1:
		*typedDest = nil
	case **RemittanceInformation5:
		*typedDest = nil
	case **ServiceLevel8Choice:
		*typedDest = nil
	case **TaxAmount1:
		*typedDest = nil
	case **TaxAuthorisation1:
		*typedDest = nil
	case **TaxInformation3:
		*typedDest = nil
	case **TaxParty1:
		*typedDest = nil
	case **TaxParty2:
		*typedDest = nil
	case **TaxPeriod1:
		*typedDest = nil
	case *Authorisation1Choice:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]Authorisation1Choice:
		*typedDest = nil
	case *DirectDebitTransactionInformation9:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DirectDebitTransactionInformation9:
		*typedDest = nil
	case *DocumentAdjustment1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]DocumentAdjustment1:
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
	case *PaymentInstructionInformation4:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]PaymentInstructionInformation4:
		*typedDest = nil
	case *ReferredDocumentInformation3:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]ReferredDocumentInformation3:
		*typedDest = nil
	case *RegulatoryReporting3:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]RegulatoryReporting3:
		*typedDest = nil
	case *RemittanceLocation2:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]RemittanceLocation2:
		*typedDest = nil
	case *StructuredRegulatoryReporting3:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]StructuredRegulatoryReporting3:
		*typedDest = nil
	case *StructuredRemittanceInformation7:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]StructuredRemittanceInformation7:
		*typedDest = nil
	case *TaxRecord1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TaxRecord1:
		*typedDest = nil
	case *TaxRecordDetails1:
		err = fmt.Errorf("unsupported feature of clearing single array items")
		return err
	case *[]TaxRecordDetails1:
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
	case **ActiveOrHistoricCurrencyAndAmount:
		return *tv, nil
	case **AmendmentInformationDetails6:
		return *tv, nil
	case **BranchAndFinancialInstitutionIdentification4:
		return *tv, nil
	case **BranchData2:
		return *tv, nil
	case **CashAccount16:
		return *tv, nil
	case **CashAccountType2:
		return *tv, nil
	case **CategoryPurpose1Choice:
		return *tv, nil
	case **ClearingSystemIdentification2Choice:
		return *tv, nil
	case **ClearingSystemMemberIdentification2:
		return *tv, nil
	case **ContactDetails2:
		return *tv, nil
	case **CreditorReferenceInformation2:
		return *tv, nil
	case **CreditorReferenceType1Choice:
		return *tv, nil
	case **CreditorReferenceType2:
		return *tv, nil
	case **CustomerDirectDebitInitiationV02:
		return *tv, nil
	case **DateAndPlaceOfBirth:
		return *tv, nil
	case **DatePeriodDetails:
		return *tv, nil
	case **DirectDebitTransaction6:
		return *tv, nil
	case **FinancialIdentificationSchemeName1Choice:
		return *tv, nil
	case **FinancialInstitutionIdentification7:
		return *tv, nil
	case **GenericAccountIdentification1:
		return *tv, nil
	case **GenericFinancialIdentification1:
		return *tv, nil
	case **GroupHeader39:
		return *tv, nil
	case **LocalInstrument2Choice:
		return *tv, nil
	case **MandateRelatedInformation6:
		return *tv, nil
	case **NameAndAddress10:
		return *tv, nil
	case **OrganisationIdentification4:
		return *tv, nil
	case **OrganisationIdentificationSchemeName1Choice:
		return *tv, nil
	case **Party6Choice:
		return *tv, nil
	case **PartyIdentification32:
		return *tv, nil
	case **PaymentIdentification1:
		return *tv, nil
	case **PaymentTypeInformation20:
		return *tv, nil
	case **PersonIdentification5:
		return *tv, nil
	case **PersonIdentificationSchemeName1Choice:
		return *tv, nil
	case **PostalAddress6:
		return *tv, nil
	case **Purpose2Choice:
		return *tv, nil
	case **ReferredDocumentType1Choice:
		return *tv, nil
	case **ReferredDocumentType2:
		return *tv, nil
	case **RegulatoryAuthority2:
		return *tv, nil
	case **RemittanceAmount1:
		return *tv, nil
	case **RemittanceInformation5:
		return *tv, nil
	case **ServiceLevel8Choice:
		return *tv, nil
	case **TaxAmount1:
		return *tv, nil
	case **TaxAuthorisation1:
		return *tv, nil
	case **TaxInformation3:
		return *tv, nil
	case **TaxParty1:
		return *tv, nil
	case **TaxParty2:
		return *tv, nil
	case **TaxPeriod1:
		return *tv, nil
	case *Authorisation1Choice:
		return tv, nil
	case *[]Authorisation1Choice:
		return *tv, nil
	case *DirectDebitTransactionInformation9:
		return tv, nil
	case *[]DirectDebitTransactionInformation9:
		return *tv, nil
	case *DocumentAdjustment1:
		return tv, nil
	case *[]DocumentAdjustment1:
		return *tv, nil
	case *GenericOrganisationIdentification1:
		return tv, nil
	case *[]GenericOrganisationIdentification1:
		return *tv, nil
	case *GenericPersonIdentification1:
		return tv, nil
	case *[]GenericPersonIdentification1:
		return *tv, nil
	case *PaymentInstructionInformation4:
		return tv, nil
	case *[]PaymentInstructionInformation4:
		return *tv, nil
	case *ReferredDocumentInformation3:
		return tv, nil
	case *[]ReferredDocumentInformation3:
		return *tv, nil
	case *RegulatoryReporting3:
		return tv, nil
	case *[]RegulatoryReporting3:
		return *tv, nil
	case *RemittanceLocation2:
		return tv, nil
	case *[]RemittanceLocation2:
		return *tv, nil
	case *StructuredRegulatoryReporting3:
		return tv, nil
	case *[]StructuredRegulatoryReporting3:
		return *tv, nil
	case *StructuredRemittanceInformation7:
		return tv, nil
	case *[]StructuredRemittanceInformation7:
		return *tv, nil
	case *TaxRecord1:
		return tv, nil
	case *[]TaxRecord1:
		return *tv, nil
	case *TaxRecordDetails1:
		return tv, nil
	case *[]TaxRecordDetails1:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), val)
	}

	return val, err
}
