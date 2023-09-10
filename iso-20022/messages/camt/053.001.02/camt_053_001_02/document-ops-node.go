// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

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

	const semLogContext = "camt_053_001_02::copy-to-dest"

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
	case **AlternateSecurityIdentification2:
		if typedStruct, ok := src.(AlternateSecurityIdentification2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmountAndCurrencyExchange3:
		if typedStruct, ok := src.(AmountAndCurrencyExchange3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmountAndCurrencyExchangeDetails3:
		if typedStruct, ok := src.(AmountAndCurrencyExchangeDetails3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmountRangeBoundary1:
		if typedStruct, ok := src.(AmountRangeBoundary1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BalanceSubType1Choice:
		if typedStruct, ok := src.(BalanceSubType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BalanceType12:
		if typedStruct, ok := src.(BalanceType12); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BalanceType5Choice:
		if typedStruct, ok := src.(BalanceType5Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BankToCustomerStatementV02:
		if typedStruct, ok := src.(BankToCustomerStatementV02); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BankTransactionCodeStructure4:
		if typedStruct, ok := src.(BankTransactionCodeStructure4); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BankTransactionCodeStructure5:
		if typedStruct, ok := src.(BankTransactionCodeStructure5); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BankTransactionCodeStructure6:
		if typedStruct, ok := src.(BankTransactionCodeStructure6); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **BatchInformation2:
		if typedStruct, ok := src.(BatchInformation2); ok {
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
	case **CashAccount20:
		if typedStruct, ok := src.(CashAccount20); ok {
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
	case **CashBalanceAvailabilityDate1:
		if typedStruct, ok := src.(CashBalanceAvailabilityDate1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ChargeType2Choice:
		if typedStruct, ok := src.(ChargeType2Choice); ok {
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
	case **CorporateAction1:
		if typedStruct, ok := src.(CorporateAction1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CreditLine2:
		if typedStruct, ok := src.(CreditLine2); ok {
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
	case **CurrencyAndAmountRange2:
		if typedStruct, ok := src.(CurrencyAndAmountRange2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CurrencyExchange5:
		if typedStruct, ok := src.(CurrencyExchange5); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **DateAndDateTimeChoice:
		if typedStruct, ok := src.(DateAndDateTimeChoice); ok {
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
	case **DateTimePeriodDetails:
		if typedStruct, ok := src.(DateTimePeriodDetails); ok {
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
	case **FinancialInstrumentQuantityChoice:
		if typedStruct, ok := src.(FinancialInstrumentQuantityChoice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **FromToAmountRange:
		if typedStruct, ok := src.(FromToAmountRange); ok {
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
	case **GenericIdentification3:
		if typedStruct, ok := src.(GenericIdentification3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **GroupHeader42:
		if typedStruct, ok := src.(GroupHeader42); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ImpliedCurrencyAmountRangeChoice:
		if typedStruct, ok := src.(ImpliedCurrencyAmountRangeChoice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **InterestType1Choice:
		if typedStruct, ok := src.(InterestType1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **MessageIdentification2:
		if typedStruct, ok := src.(MessageIdentification2); ok {
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
	case **NumberAndSumOfTransactions1:
		if typedStruct, ok := src.(NumberAndSumOfTransactions1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **NumberAndSumOfTransactions2:
		if typedStruct, ok := src.(NumberAndSumOfTransactions2); ok {
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
	case **Pagination:
		if typedStruct, ok := src.(Pagination); ok {
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
	case **ProprietaryBankTransactionCodeStructure1:
		if typedStruct, ok := src.(ProprietaryBankTransactionCodeStructure1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProprietaryQuantity1:
		if typedStruct, ok := src.(ProprietaryQuantity1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProprietaryReference1:
		if typedStruct, ok := src.(ProprietaryReference1); ok {
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
	case **RateType4Choice:
		if typedStruct, ok := src.(RateType4Choice); ok {
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
	case **ReportingSource1Choice:
		if typedStruct, ok := src.(ReportingSource1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReturnReason5Choice:
		if typedStruct, ok := src.(ReturnReason5Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReturnReasonInformation10:
		if typedStruct, ok := src.(ReturnReasonInformation10); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **SecurityIdentification4Choice:
		if typedStruct, ok := src.(SecurityIdentification4Choice); ok {
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
	case **TaxCharges2:
		if typedStruct, ok := src.(TaxCharges2); ok {
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
	case **TechnicalInputChannel1Choice:
		if typedStruct, ok := src.(TechnicalInputChannel1Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TotalTransactions2:
		if typedStruct, ok := src.(TotalTransactions2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionAgents2:
		if typedStruct, ok := src.(TransactionAgents2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionDates2:
		if typedStruct, ok := src.(TransactionDates2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionParty2:
		if typedStruct, ok := src.(TransactionParty2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionPrice2Choice:
		if typedStruct, ok := src.(TransactionPrice2Choice); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionReferences2:
		if typedStruct, ok := src.(TransactionReferences2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AccountInterest2:
		if typedStruct, ok := src.(AccountInterest2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AccountStatement2:
		if typedStruct, ok := src.(AccountStatement2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **AmountAndCurrencyExchangeDetails4:
		if typedStruct, ok := src.(AmountAndCurrencyExchangeDetails4); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashBalance3:
		if typedStruct, ok := src.(CashBalance3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **CashBalanceAvailability2:
		if typedStruct, ok := src.(CashBalanceAvailability2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ChargesInformation6:
		if typedStruct, ok := src.(ChargesInformation6); ok {
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
	case **EntryDetails1:
		if typedStruct, ok := src.(EntryDetails1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **EntryTransaction2:
		if typedStruct, ok := src.(EntryTransaction2); ok {
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
	case **ProprietaryAgent2:
		if typedStruct, ok := src.(ProprietaryAgent2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProprietaryDate2:
		if typedStruct, ok := src.(ProprietaryDate2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProprietaryParty2:
		if typedStruct, ok := src.(ProprietaryParty2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ProprietaryPrice2:
		if typedStruct, ok := src.(ProprietaryPrice2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **Rate3:
		if typedStruct, ok := src.(Rate3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReferredDocumentInformation3:
		if typedStruct, ok := src.(ReferredDocumentInformation3); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **RemittanceLocation2:
		if typedStruct, ok := src.(RemittanceLocation2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **ReportEntry2:
		if typedStruct, ok := src.(ReportEntry2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **StructuredRemittanceInformation7:
		if typedStruct, ok := src.(StructuredRemittanceInformation7); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxRecord1:
		if typedStruct, ok := src.(TaxRecord1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TaxRecordDetails1:
		if typedStruct, ok := src.(TaxRecordDetails1); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TotalsPerBankTransactionCode2:
		if typedStruct, ok := src.(TotalsPerBankTransactionCode2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionInterest2:
		if typedStruct, ok := src.(TransactionInterest2); ok {
			*typedDest = &typedStruct
		} else {
			err = fmt.Errorf("value of type %T cannot be assigned", src)
		}
		if err != nil && options.logVerbose {
			log.Error().Err(err).Str("path", string(docPath)).Interface("value", src).Msg(semLogContext)
		}
		return err
	case **TransactionQuantities1Choice:
		if typedStruct, ok := src.(TransactionQuantities1Choice); ok {
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
	const semLogContext = "camt_053_001_02::clear-node"

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

	const semLogContext = "camt_053_001_02::clear-node"

	var err error
	switch typedDest := dest.(type) {
	case **AccountIdentification4Choice:
		*typedDest = nil
	case **AccountSchemeName1Choice:
		*typedDest = nil
	case **ActiveOrHistoricCurrencyAndAmount:
		*typedDest = nil
	case **AlternateSecurityIdentification2:
		*typedDest = nil
	case **AmountAndCurrencyExchange3:
		*typedDest = nil
	case **AmountAndCurrencyExchangeDetails3:
		*typedDest = nil
	case **AmountRangeBoundary1:
		*typedDest = nil
	case **BalanceSubType1Choice:
		*typedDest = nil
	case **BalanceType12:
		*typedDest = nil
	case **BalanceType5Choice:
		*typedDest = nil
	case **BankToCustomerStatementV02:
		*typedDest = nil
	case **BankTransactionCodeStructure4:
		*typedDest = nil
	case **BankTransactionCodeStructure5:
		*typedDest = nil
	case **BankTransactionCodeStructure6:
		*typedDest = nil
	case **BatchInformation2:
		*typedDest = nil
	case **BranchAndFinancialInstitutionIdentification4:
		*typedDest = nil
	case **BranchData2:
		*typedDest = nil
	case **CashAccount16:
		*typedDest = nil
	case **CashAccount20:
		*typedDest = nil
	case **CashAccountType2:
		*typedDest = nil
	case **CashBalanceAvailabilityDate1:
		*typedDest = nil
	case **ChargeType2Choice:
		*typedDest = nil
	case **ClearingSystemIdentification2Choice:
		*typedDest = nil
	case **ClearingSystemMemberIdentification2:
		*typedDest = nil
	case **ContactDetails2:
		*typedDest = nil
	case **CorporateAction1:
		*typedDest = nil
	case **CreditLine2:
		*typedDest = nil
	case **CreditorReferenceInformation2:
		*typedDest = nil
	case **CreditorReferenceType1Choice:
		*typedDest = nil
	case **CreditorReferenceType2:
		*typedDest = nil
	case **CurrencyAndAmountRange2:
		*typedDest = nil
	case **CurrencyExchange5:
		*typedDest = nil
	case **DateAndDateTimeChoice:
		*typedDest = nil
	case **DateAndPlaceOfBirth:
		*typedDest = nil
	case **DatePeriodDetails:
		*typedDest = nil
	case **DateTimePeriodDetails:
		*typedDest = nil
	case **FinancialIdentificationSchemeName1Choice:
		*typedDest = nil
	case **FinancialInstitutionIdentification7:
		*typedDest = nil
	case **FinancialInstrumentQuantityChoice:
		*typedDest = nil
	case **FromToAmountRange:
		*typedDest = nil
	case **GenericAccountIdentification1:
		*typedDest = nil
	case **GenericFinancialIdentification1:
		*typedDest = nil
	case **GenericIdentification3:
		*typedDest = nil
	case **GroupHeader42:
		*typedDest = nil
	case **ImpliedCurrencyAmountRangeChoice:
		*typedDest = nil
	case **InterestType1Choice:
		*typedDest = nil
	case **MessageIdentification2:
		*typedDest = nil
	case **NameAndAddress10:
		*typedDest = nil
	case **NumberAndSumOfTransactions1:
		*typedDest = nil
	case **NumberAndSumOfTransactions2:
		*typedDest = nil
	case **OrganisationIdentification4:
		*typedDest = nil
	case **OrganisationIdentificationSchemeName1Choice:
		*typedDest = nil
	case **Pagination:
		*typedDest = nil
	case **Party6Choice:
		*typedDest = nil
	case **PartyIdentification32:
		*typedDest = nil
	case **PersonIdentification5:
		*typedDest = nil
	case **PersonIdentificationSchemeName1Choice:
		*typedDest = nil
	case **PostalAddress6:
		*typedDest = nil
	case **ProprietaryBankTransactionCodeStructure1:
		*typedDest = nil
	case **ProprietaryQuantity1:
		*typedDest = nil
	case **ProprietaryReference1:
		*typedDest = nil
	case **Purpose2Choice:
		*typedDest = nil
	case **RateType4Choice:
		*typedDest = nil
	case **ReferredDocumentType1Choice:
		*typedDest = nil
	case **ReferredDocumentType2:
		*typedDest = nil
	case **RemittanceAmount1:
		*typedDest = nil
	case **RemittanceInformation5:
		*typedDest = nil
	case **ReportingSource1Choice:
		*typedDest = nil
	case **ReturnReason5Choice:
		*typedDest = nil
	case **ReturnReasonInformation10:
		*typedDest = nil
	case **SecurityIdentification4Choice:
		*typedDest = nil
	case **TaxAmount1:
		*typedDest = nil
	case **TaxAuthorisation1:
		*typedDest = nil
	case **TaxCharges2:
		*typedDest = nil
	case **TaxInformation3:
		*typedDest = nil
	case **TaxParty1:
		*typedDest = nil
	case **TaxParty2:
		*typedDest = nil
	case **TaxPeriod1:
		*typedDest = nil
	case **TechnicalInputChannel1Choice:
		*typedDest = nil
	case **TotalTransactions2:
		*typedDest = nil
	case **TransactionAgents2:
		*typedDest = nil
	case **TransactionDates2:
		*typedDest = nil
	case **TransactionParty2:
		*typedDest = nil
	case **TransactionPrice2Choice:
		*typedDest = nil
	case **TransactionReferences2:
		*typedDest = nil
	case **AccountInterest2:
		*typedDest = nil
	case **AccountStatement2:
		*typedDest = nil
	case *AccountStatement2:
		typedDest = nil
	case **AmountAndCurrencyExchangeDetails4:
		*typedDest = nil
	case **CashBalance3:
		*typedDest = nil
	case **CashBalanceAvailability2:
		*typedDest = nil
	case **ChargesInformation6:
		*typedDest = nil
	case **DocumentAdjustment1:
		*typedDest = nil
	case **EntryDetails1:
		*typedDest = nil
	case **EntryTransaction2:
		*typedDest = nil
	case **GenericOrganisationIdentification1:
		*typedDest = nil
	case **GenericPersonIdentification1:
		*typedDest = nil
	case **ProprietaryAgent2:
		*typedDest = nil
	case **ProprietaryDate2:
		*typedDest = nil
	case **ProprietaryParty2:
		*typedDest = nil
	case **ProprietaryPrice2:
		*typedDest = nil
	case **Rate3:
		*typedDest = nil
	case **ReferredDocumentInformation3:
		*typedDest = nil
	case **RemittanceLocation2:
		*typedDest = nil
	case **ReportEntry2:
		*typedDest = nil
	case **StructuredRemittanceInformation7:
		*typedDest = nil
	case **TaxRecord1:
		*typedDest = nil
	case **TaxRecordDetails1:
		*typedDest = nil
	case **TotalsPerBankTransactionCode2:
		*typedDest = nil
	case **TransactionInterest2:
		*typedDest = nil
	case **TransactionQuantities1Choice:
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
	case **AlternateSecurityIdentification2:
		return *tv, nil
	case **AmountAndCurrencyExchange3:
		return *tv, nil
	case **AmountAndCurrencyExchangeDetails3:
		return *tv, nil
	case **AmountRangeBoundary1:
		return *tv, nil
	case **BalanceSubType1Choice:
		return *tv, nil
	case **BalanceType12:
		return *tv, nil
	case **BalanceType5Choice:
		return *tv, nil
	case **BankToCustomerStatementV02:
		return *tv, nil
	case **BankTransactionCodeStructure4:
		return *tv, nil
	case **BankTransactionCodeStructure5:
		return *tv, nil
	case **BankTransactionCodeStructure6:
		return *tv, nil
	case **BatchInformation2:
		return *tv, nil
	case **BranchAndFinancialInstitutionIdentification4:
		return *tv, nil
	case **BranchData2:
		return *tv, nil
	case **CashAccount16:
		return *tv, nil
	case **CashAccount20:
		return *tv, nil
	case **CashAccountType2:
		return *tv, nil
	case **CashBalanceAvailabilityDate1:
		return *tv, nil
	case **ChargeType2Choice:
		return *tv, nil
	case **ClearingSystemIdentification2Choice:
		return *tv, nil
	case **ClearingSystemMemberIdentification2:
		return *tv, nil
	case **ContactDetails2:
		return *tv, nil
	case **CorporateAction1:
		return *tv, nil
	case **CreditLine2:
		return *tv, nil
	case **CreditorReferenceInformation2:
		return *tv, nil
	case **CreditorReferenceType1Choice:
		return *tv, nil
	case **CreditorReferenceType2:
		return *tv, nil
	case **CurrencyAndAmountRange2:
		return *tv, nil
	case **CurrencyExchange5:
		return *tv, nil
	case **DateAndDateTimeChoice:
		return *tv, nil
	case **DateAndPlaceOfBirth:
		return *tv, nil
	case **DatePeriodDetails:
		return *tv, nil
	case **DateTimePeriodDetails:
		return *tv, nil
	case **FinancialIdentificationSchemeName1Choice:
		return *tv, nil
	case **FinancialInstitutionIdentification7:
		return *tv, nil
	case **FinancialInstrumentQuantityChoice:
		return *tv, nil
	case **FromToAmountRange:
		return *tv, nil
	case **GenericAccountIdentification1:
		return *tv, nil
	case **GenericFinancialIdentification1:
		return *tv, nil
	case **GenericIdentification3:
		return *tv, nil
	case **GroupHeader42:
		return *tv, nil
	case **ImpliedCurrencyAmountRangeChoice:
		return *tv, nil
	case **InterestType1Choice:
		return *tv, nil
	case **MessageIdentification2:
		return *tv, nil
	case **NameAndAddress10:
		return *tv, nil
	case **NumberAndSumOfTransactions1:
		return *tv, nil
	case **NumberAndSumOfTransactions2:
		return *tv, nil
	case **OrganisationIdentification4:
		return *tv, nil
	case **OrganisationIdentificationSchemeName1Choice:
		return *tv, nil
	case **Pagination:
		return *tv, nil
	case **Party6Choice:
		return *tv, nil
	case **PartyIdentification32:
		return *tv, nil
	case **PersonIdentification5:
		return *tv, nil
	case **PersonIdentificationSchemeName1Choice:
		return *tv, nil
	case **PostalAddress6:
		return *tv, nil
	case **ProprietaryBankTransactionCodeStructure1:
		return *tv, nil
	case **ProprietaryQuantity1:
		return *tv, nil
	case **ProprietaryReference1:
		return *tv, nil
	case **Purpose2Choice:
		return *tv, nil
	case **RateType4Choice:
		return *tv, nil
	case **ReferredDocumentType1Choice:
		return *tv, nil
	case **ReferredDocumentType2:
		return *tv, nil
	case **RemittanceAmount1:
		return *tv, nil
	case **RemittanceInformation5:
		return *tv, nil
	case **ReportingSource1Choice:
		return *tv, nil
	case **ReturnReason5Choice:
		return *tv, nil
	case **ReturnReasonInformation10:
		return *tv, nil
	case **SecurityIdentification4Choice:
		return *tv, nil
	case **TaxAmount1:
		return *tv, nil
	case **TaxAuthorisation1:
		return *tv, nil
	case **TaxCharges2:
		return *tv, nil
	case **TaxInformation3:
		return *tv, nil
	case **TaxParty1:
		return *tv, nil
	case **TaxParty2:
		return *tv, nil
	case **TaxPeriod1:
		return *tv, nil
	case **TechnicalInputChannel1Choice:
		return *tv, nil
	case **TotalTransactions2:
		return *tv, nil
	case **TransactionAgents2:
		return *tv, nil
	case **TransactionDates2:
		return *tv, nil
	case **TransactionParty2:
		return *tv, nil
	case **TransactionPrice2Choice:
		return *tv, nil
	case **TransactionReferences2:
		return *tv, nil
	case **AccountInterest2:
		return *tv, nil
	case *[]AccountStatement2:
		return *tv, nil
	case *AccountStatement2:
		return *tv, nil
	case **AmountAndCurrencyExchangeDetails4:
		return *tv, nil
	case **CashBalance3:
		return *tv, nil
	case **CashBalanceAvailability2:
		return *tv, nil
	case **ChargesInformation6:
		return *tv, nil
	case **DocumentAdjustment1:
		return *tv, nil
	case **EntryDetails1:
		return *tv, nil
	case **EntryTransaction2:
		return *tv, nil
	case **GenericOrganisationIdentification1:
		return *tv, nil
	case **GenericPersonIdentification1:
		return *tv, nil
	case **ProprietaryAgent2:
		return *tv, nil
	case **ProprietaryDate2:
		return *tv, nil
	case **ProprietaryParty2:
		return *tv, nil
	case **ProprietaryPrice2:
		return *tv, nil
	case **Rate3:
		return *tv, nil
	case **ReferredDocumentInformation3:
		return *tv, nil
	case **RemittanceLocation2:
		return *tv, nil
	case **ReportEntry2:
		return *tv, nil
	case **StructuredRemittanceInformation7:
		return *tv, nil
	case **TaxRecord1:
		return *tv, nil
	case **TaxRecordDetails1:
		return *tv, nil
	case **TotalsPerBankTransactionCode2:
		return *tv, nil
	case **TransactionInterest2:
		return *tv, nil
	case **TransactionQuantities1Choice:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %T", string(docPath), val)
	}

	return val, err
}
