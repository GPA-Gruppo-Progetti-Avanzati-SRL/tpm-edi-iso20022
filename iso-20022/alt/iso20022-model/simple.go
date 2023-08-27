package iso20022_model

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"regexp"
	"time"
)

// ActiveOrHistoricCurrencyCode: Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// isValid evaluates simple ActiveOrHistoricCurrencyCode type validity accordingly with xsd schema rules
func (t *ActiveOrHistoricCurrencyCode) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z]{3,3}
	valid, _ := regexp.MatchString(`[A-Z]{3,3}`, fmt.Sprintf("%v", *t))
	return valid
}

// AddressType2Code: May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// isValid evaluates simple AddressType2Code type validity accordingly with xsd schema rules
func (t *AddressType2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// AnyBICDec2014Identifier: Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// isValid evaluates simple AnyBICDec2014Identifier type validity accordingly with xsd schema rules
func (t *AnyBICDec2014Identifier) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
	valid, _ := regexp.MatchString(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`, fmt.Sprintf("%v", *t))
	return valid
}

// BICFIDec2014Identifier: Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

// isValid evaluates simple BICFIDec2014Identifier type validity accordingly with xsd schema rules
func (t *BICFIDec2014Identifier) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
	valid, _ := regexp.MatchString(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`, fmt.Sprintf("%v", *t))
	return valid
}

// CancellationIndividualStatus1Code: May be one of RJCR, ACCR, PDCR
type CancellationIndividualStatus1Code string

// isValid evaluates simple CancellationIndividualStatus1Code type validity accordingly with xsd schema rules
func (t *CancellationIndividualStatus1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"RJCR", "ACCR", "PDCR"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// ChargeBearerType1Code: May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

// isValid evaluates simple ChargeBearerType1Code type validity accordingly with xsd schema rules
func (t *ChargeBearerType1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"DEBT", "CRED", "SHAR", "SLEV"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// ChequeDelivery1Code: May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

// isValid evaluates simple ChequeDelivery1Code type validity accordingly with xsd schema rules
func (t *ChequeDelivery1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"MLDB", "MLCD", "MLFA", "CRDB", "CRCD", "CRFA", "PUDB", "PUCD", "PUFA", "RGDB", "RGCD", "RGFA"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// ChequeType2Code: May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

// isValid evaluates simple ChequeType2Code type validity accordingly with xsd schema rules
func (t *ChequeType2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CCHQ", "CCCH", "BCHQ", "DRFT", "ELDR"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// ClearingChannel2Code: May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

// isValid evaluates simple ClearingChannel2Code type validity accordingly with xsd schema rules
func (t *ClearingChannel2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"RTGS", "RTNS", "MPNS", "BOOK"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// CountryCode: Must match the pattern [A-Z]{2,2}
type CountryCode string

// isValid evaluates simple CountryCode type validity accordingly with xsd schema rules
func (t *CountryCode) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z]{2,2}
	valid, _ := regexp.MatchString(`[A-Z]{2,2}`, fmt.Sprintf("%v", *t))
	return valid
}

// CreditDebitCode: May be one of CRDT, DBIT
type CreditDebitCode string

// isValid evaluates simple CreditDebitCode type validity accordingly with xsd schema rules
func (t *CreditDebitCode) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CRDT", "DBIT"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// DocumentType3Code: May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// isValid evaluates simple DocumentType3Code type validity accordingly with xsd schema rules
func (t *DocumentType3Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"RADM", "RPIN", "FXDR", "DISP", "PUOR", "SCOR"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// DocumentType6Code: May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// isValid evaluates simple DocumentType6Code type validity accordingly with xsd schema rules
func (t *DocumentType6Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"MSIN", "CNFA", "DNFA", "CINV", "CREN", "DEBN", "HIRI", "SBIN", "CMCN", "SOAC", "DISP", "BOLD", "VCHR", "AROI", "TSUT", "PUOR"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// Exact2NumericText: Must match the pattern [0-9]{2}
type Exact2NumericText string

// isValid evaluates simple Exact2NumericText type validity accordingly with xsd schema rules
func (t *Exact2NumericText) isValid() bool {
	//checks if the value matches the specified pattern
	// [0-9]{2}
	valid, _ := regexp.MatchString(`[0-9]{2}`, fmt.Sprintf("%v", *t))
	return valid
}

// Exact4AlphaNumericText: Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// isValid evaluates simple Exact4AlphaNumericText type validity accordingly with xsd schema rules
func (t *Exact4AlphaNumericText) isValid() bool {
	//checks if the value matches the specified pattern
	// [a-zA-Z0-9]{4}
	valid, _ := regexp.MatchString(`[a-zA-Z0-9]{4}`, fmt.Sprintf("%v", *t))
	return valid
}

// ExternalAccountIdentification1Code: May be no more than 4 items long
type ExternalAccountIdentification1Code string

// isValid evaluates simple ExternalAccountIdentification1Code type validity accordingly with xsd schema rules
func (t *ExternalAccountIdentification1Code) isValid() bool {
	return true
}

// ExternalCancellationReason1Code: May be no more than 4 items long
type ExternalCancellationReason1Code string

// isValid evaluates simple ExternalCancellationReason1Code type validity accordingly with xsd schema rules
func (t *ExternalCancellationReason1Code) isValid() bool {
	return true
}

// ExternalCashAccountType1Code: May be no more than 4 items long
type ExternalCashAccountType1Code string

// isValid evaluates simple ExternalCashAccountType1Code type validity accordingly with xsd schema rules
func (t *ExternalCashAccountType1Code) isValid() bool {
	return true
}

// ExternalCashClearingSystem1Code: May be no more than 3 items long
type ExternalCashClearingSystem1Code string

// isValid evaluates simple ExternalCashClearingSystem1Code type validity accordingly with xsd schema rules
func (t *ExternalCashClearingSystem1Code) isValid() bool {
	return true
}

// ExternalCategoryPurpose1Code: May be no more than 4 items long
type ExternalCategoryPurpose1Code string

// isValid evaluates simple ExternalCategoryPurpose1Code type validity accordingly with xsd schema rules
func (t *ExternalCategoryPurpose1Code) isValid() bool {
	return true
}

// ExternalChargeType1Code: May be no more than 4 items long
type ExternalChargeType1Code string

// isValid evaluates simple ExternalChargeType1Code type validity accordingly with xsd schema rules
func (t *ExternalChargeType1Code) isValid() bool {
	return true
}

// ExternalClaimNonReceiptRejection1Code: May be no more than 4 items long
type ExternalClaimNonReceiptRejection1Code string

// isValid evaluates simple ExternalClaimNonReceiptRejection1Code type validity accordingly with xsd schema rules
func (t *ExternalClaimNonReceiptRejection1Code) isValid() bool {
	return true
}

// ExternalClearingSystemIdentification1Code: May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// isValid evaluates simple ExternalClearingSystemIdentification1Code type validity accordingly with xsd schema rules
func (t *ExternalClearingSystemIdentification1Code) isValid() bool {
	return true
}

// ExternalDiscountAmountType1Code: May be no more than 4 items long
type ExternalDiscountAmountType1Code string

// isValid evaluates simple ExternalDiscountAmountType1Code type validity accordingly with xsd schema rules
func (t *ExternalDiscountAmountType1Code) isValid() bool {
	return true
}

// ExternalDocumentFormat1Code: May be no more than 4 items long
type ExternalDocumentFormat1Code string

// isValid evaluates simple ExternalDocumentFormat1Code type validity accordingly with xsd schema rules
func (t *ExternalDocumentFormat1Code) isValid() bool {
	return true
}

// ExternalDocumentLineType1Code: May be no more than 4 items long
type ExternalDocumentLineType1Code string

// isValid evaluates simple ExternalDocumentLineType1Code type validity accordingly with xsd schema rules
func (t *ExternalDocumentLineType1Code) isValid() bool {
	return true
}

// ExternalDocumentType1Code: May be no more than 4 items long
type ExternalDocumentType1Code string

// isValid evaluates simple ExternalDocumentType1Code type validity accordingly with xsd schema rules
func (t *ExternalDocumentType1Code) isValid() bool {
	return true
}

// ExternalFinancialInstitutionIdentification1Code: May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// isValid evaluates simple ExternalFinancialInstitutionIdentification1Code type validity accordingly with xsd schema rules
func (t *ExternalFinancialInstitutionIdentification1Code) isValid() bool {
	return true
}

// ExternalGarnishmentType1Code: May be no more than 4 items long
type ExternalGarnishmentType1Code string

// isValid evaluates simple ExternalGarnishmentType1Code type validity accordingly with xsd schema rules
func (t *ExternalGarnishmentType1Code) isValid() bool {
	return true
}

// ExternalInvestigationExecutionConfirmation1Code: May be no more than 4 items long
type ExternalInvestigationExecutionConfirmation1Code string

// isValid evaluates simple ExternalInvestigationExecutionConfirmation1Code type validity accordingly with xsd schema rules
func (t *ExternalInvestigationExecutionConfirmation1Code) isValid() bool {
	return true
}

// ExternalLocalInstrument1Code: May be no more than 35 items long
type ExternalLocalInstrument1Code string

// isValid evaluates simple ExternalLocalInstrument1Code type validity accordingly with xsd schema rules
func (t *ExternalLocalInstrument1Code) isValid() bool {
	return true
}

// ExternalMandateSetupReason1Code: May be no more than 4 items long
type ExternalMandateSetupReason1Code string

// isValid evaluates simple ExternalMandateSetupReason1Code type validity accordingly with xsd schema rules
func (t *ExternalMandateSetupReason1Code) isValid() bool {
	return true
}

// ExternalOrganisationIdentification1Code: May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// isValid evaluates simple ExternalOrganisationIdentification1Code type validity accordingly with xsd schema rules
func (t *ExternalOrganisationIdentification1Code) isValid() bool {
	return true
}

// ExternalPaymentCancellationRejection1Code: May be no more than 4 items long
type ExternalPaymentCancellationRejection1Code string

// isValid evaluates simple ExternalPaymentCancellationRejection1Code type validity accordingly with xsd schema rules
func (t *ExternalPaymentCancellationRejection1Code) isValid() bool {
	return true
}

// ExternalPaymentCompensationReason1Code: May be no more than 4 items long
type ExternalPaymentCompensationReason1Code string

// isValid evaluates simple ExternalPaymentCompensationReason1Code type validity accordingly with xsd schema rules
func (t *ExternalPaymentCompensationReason1Code) isValid() bool {
	return true
}

// ExternalPaymentGroupStatus1Code: May be no more than 4 items long
type ExternalPaymentGroupStatus1Code string

// isValid evaluates simple ExternalPaymentGroupStatus1Code type validity accordingly with xsd schema rules
func (t *ExternalPaymentGroupStatus1Code) isValid() bool {
	return true
}

// ExternalPaymentModificationRejection1Code: May be no more than 4 items long
type ExternalPaymentModificationRejection1Code string

// isValid evaluates simple ExternalPaymentModificationRejection1Code type validity accordingly with xsd schema rules
func (t *ExternalPaymentModificationRejection1Code) isValid() bool {
	return true
}

// ExternalPaymentTransactionStatus1Code: May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code string

// isValid evaluates simple ExternalPaymentTransactionStatus1Code type validity accordingly with xsd schema rules
func (t *ExternalPaymentTransactionStatus1Code) isValid() bool {
	return true
}

// ExternalPersonIdentification1Code: May be no more than 4 items long
type ExternalPersonIdentification1Code string

// isValid evaluates simple ExternalPersonIdentification1Code type validity accordingly with xsd schema rules
func (t *ExternalPersonIdentification1Code) isValid() bool {
	return true
}

// ExternalProxyAccountType1Code: May be no more than 4 items long
type ExternalProxyAccountType1Code string

// isValid evaluates simple ExternalProxyAccountType1Code type validity accordingly with xsd schema rules
func (t *ExternalProxyAccountType1Code) isValid() bool {
	return true
}

// ExternalPurpose1Code: May be no more than 4 items long
type ExternalPurpose1Code string

// isValid evaluates simple ExternalPurpose1Code type validity accordingly with xsd schema rules
func (t *ExternalPurpose1Code) isValid() bool {
	return true
}

// ExternalServiceLevel1Code: May be no more than 4 items long
type ExternalServiceLevel1Code string

// isValid evaluates simple ExternalServiceLevel1Code type validity accordingly with xsd schema rules
func (t *ExternalServiceLevel1Code) isValid() bool {
	return true
}

// ExternalStatusReason1Code: May be no more than 4 items long
type ExternalStatusReason1Code string

// isValid evaluates simple ExternalStatusReason1Code type validity accordingly with xsd schema rules
func (t *ExternalStatusReason1Code) isValid() bool {
	return true
}

// ExternalTaxAmountType1Code: May be no more than 4 items long
type ExternalTaxAmountType1Code string

// isValid evaluates simple ExternalTaxAmountType1Code type validity accordingly with xsd schema rules
func (t *ExternalTaxAmountType1Code) isValid() bool {
	return true
}

// Frequency6Code: May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

// isValid evaluates simple Frequency6Code type validity accordingly with xsd schema rules
func (t *Frequency6Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "FRTN"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// GroupCancellationStatus1Code: May be one of PACR, RJCR, ACCR, PDCR
type GroupCancellationStatus1Code string

// isValid evaluates simple GroupCancellationStatus1Code type validity accordingly with xsd schema rules
func (t *GroupCancellationStatus1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"PACR", "RJCR", "ACCR", "PDCR"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// IBAN2007Identifier: Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// isValid evaluates simple IBAN2007Identifier type validity accordingly with xsd schema rules
func (t *IBAN2007Identifier) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
	valid, _ := regexp.MatchString(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`, fmt.Sprintf("%v", *t))
	return valid
}

// Instruction3Code: May be one of CHQB, HOLD, PHOB, TELB
type Instruction3Code string

// isValid evaluates simple Instruction3Code type validity accordingly with xsd schema rules
func (t *Instruction3Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CHQB", "HOLD", "PHOB", "TELB"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// LEIIdentifier: Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// isValid evaluates simple LEIIdentifier type validity accordingly with xsd schema rules
func (t *LEIIdentifier) isValid() bool {
	//checks if the value matches the specified pattern
	// [A-Z0-9]{18,18}[0-9]{2,2}
	valid, _ := regexp.MatchString(`[A-Z0-9]{18,18}[0-9]{2,2}`, fmt.Sprintf("%v", *t))
	return valid
}

// Max1025Text: May be no more than 1025 items long
type Max1025Text string

// isValid evaluates simple Max1025Text type validity accordingly with xsd schema rules
func (t *Max1025Text) isValid() bool {
	return true
}

// Max105Text: May be no more than 105 items long
type Max105Text string

// isValid evaluates simple Max105Text type validity accordingly with xsd schema rules
func (t *Max105Text) isValid() bool {
	return true
}

// Max10Text: May be no more than 10 items long
type Max10Text string

// isValid evaluates simple Max10Text type validity accordingly with xsd schema rules
func (t *Max10Text) isValid() bool {
	return true
}

// Max128Text: May be no more than 128 items long
type Max128Text string

// isValid evaluates simple Max128Text type validity accordingly with xsd schema rules
func (t *Max128Text) isValid() bool {
	return true
}

// Max140Text: May be no more than 140 items long
type Max140Text string

// isValid evaluates simple Max140Text type validity accordingly with xsd schema rules
func (t *Max140Text) isValid() bool {
	return true
}

// Max15NumericText: Must match the pattern [0-9]{1,15}
type Max15NumericText string

// isValid evaluates simple Max15NumericText type validity accordingly with xsd schema rules
func (t *Max15NumericText) isValid() bool {
	//checks if the value matches the specified pattern
	// [0-9]{1,15}
	valid, _ := regexp.MatchString(`[0-9]{1,15}`, fmt.Sprintf("%v", *t))
	return valid
}

// Max16Text: May be no more than 16 items long
type Max16Text string

// isValid evaluates simple Max16Text type validity accordingly with xsd schema rules
func (t *Max16Text) isValid() bool {
	return true
}

// Max2048Text: May be no more than 2048 items long
type Max2048Text string

// isValid evaluates simple Max2048Text type validity accordingly with xsd schema rules
func (t *Max2048Text) isValid() bool {
	return true
}

// Max34Text: May be no more than 34 items long
type Max34Text string

// isValid evaluates simple Max34Text type validity accordingly with xsd schema rules
func (t *Max34Text) isValid() bool {
	return true
}

// Max350Text: May be no more than 350 items long
type Max350Text string

// isValid evaluates simple Max350Text type validity accordingly with xsd schema rules
func (t *Max350Text) isValid() bool {
	return true
}

// Max35Text: May be no more than 35 items long
type Max35Text string

// isValid evaluates simple Max35Text type validity accordingly with xsd schema rules
func (t *Max35Text) isValid() bool {
	return true
}

// Max4Text: May be no more than 4 items long
type Max4Text string

// isValid evaluates simple Max4Text type validity accordingly with xsd schema rules
func (t *Max4Text) isValid() bool {
	return true
}

// Max70Text: May be no more than 70 items long
type Max70Text string

// isValid evaluates simple Max70Text type validity accordingly with xsd schema rules
func (t *Max70Text) isValid() bool {
	return true
}

// NamePrefix2Code: May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

// isValid evaluates simple NamePrefix2Code type validity accordingly with xsd schema rules
func (t *NamePrefix2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"DOCT", "MADM", "MISS", "MIST", "MIKS"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// PaymentMethod4Code: May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

// isValid evaluates simple PaymentMethod4Code type validity accordingly with xsd schema rules
func (t *PaymentMethod4Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CHK", "TRF", "DD", "TRA"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// PaymentMethod7Code: May be one of CHK, TRF
type PaymentMethod7Code string

// isValid evaluates simple PaymentMethod7Code type validity accordingly with xsd schema rules
func (t *PaymentMethod7Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CHK", "TRF"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// PhoneNumber: Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

// isValid evaluates simple PhoneNumber type validity accordingly with xsd schema rules
func (t *PhoneNumber) isValid() bool {
	//checks if the value matches the specified pattern
	// \+[0-9]{1,3}-[0-9()+\-]{1,30}
	valid, _ := regexp.MatchString(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`, fmt.Sprintf("%v", *t))
	return valid
}

// PreferredContactMethod1Code: May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// isValid evaluates simple PreferredContactMethod1Code type validity accordingly with xsd schema rules
func (t *PreferredContactMethod1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"LETT", "MAIL", "PHON", "FAXX", "CELL"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// Priority2Code: May be one of HIGH, NORM
type Priority2Code string

// isValid evaluates simple Priority2Code type validity accordingly with xsd schema rules
func (t *Priority2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"HIGH", "NORM"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// RegulatoryReportingType1Code: May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

// isValid evaluates simple RegulatoryReportingType1Code type validity accordingly with xsd schema rules
func (t *RegulatoryReportingType1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"CRED", "DEBT", "BOTH"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// RemittanceLocationMethod2Code: May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

// isValid evaluates simple RemittanceLocationMethod2Code type validity accordingly with xsd schema rules
func (t *RemittanceLocationMethod2Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"FAXI", "EDIC", "URID", "EMAL", "POST", "SMSM"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// SequenceType3Code: May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code string

// isValid evaluates simple SequenceType3Code type validity accordingly with xsd schema rules
func (t *SequenceType3Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"FRST", "RCUR", "FNAL", "OOFF", "RPRE"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// SettlementMethod1Code: May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

// isValid evaluates simple SettlementMethod1Code type validity accordingly with xsd schema rules
func (t *SettlementMethod1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"INDA", "INGA", "COVE", "CLRG"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// TaxRecordPeriod1Code: May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// isValid evaluates simple TaxRecordPeriod1Code type validity accordingly with xsd schema rules
func (t *TaxRecordPeriod1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"MM01", "MM02", "MM03", "MM04", "MM05", "MM06", "MM07", "MM08", "MM09", "MM10", "MM11", "MM12", "QTR1", "QTR2", "QTR3", "QTR4", "HLF1", "HLF2"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// TransactionIndividualStatus1Code: May be one of ACTC, RJCT, PDNG, ACCP, ACSP, ACSC, ACCR, ACWC
type TransactionIndividualStatus1Code string

// isValid evaluates simple TransactionIndividualStatus1Code type validity accordingly with xsd schema rules
func (t *TransactionIndividualStatus1Code) isValid() bool {
	//checks if the value is in the enumeration list
	values := []string{"ACTC", "RJCT", "PDNG", "ACCP", "ACSP", "ACSC", "ACCR", "ACWC"}
	for _, v := range values {
		if v == fmt.Sprintf("%v", *t) {
			return true
		}
	}
	return false
}

// UUIDv4Identifier: Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

// isValid evaluates simple UUIDv4Identifier type validity accordingly with xsd schema rules
func (t *UUIDv4Identifier) isValid() bool {
	//checks if the value matches the specified pattern
	// [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
	valid, _ := regexp.MatchString(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`, fmt.Sprintf("%v", *t))
	return valid
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}
func (b xsdBase64Binary) isValid() bool {
	return true
}

type Max10MbBinary []byte

func (t *Max10MbBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10MbBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}
func (t Max10MbBinary) isValid() bool {
	return true
}

// xsdDate
type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}

func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}

func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}

func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

// isValid returning true by default
func (t xsdDate) isValid() bool {
	return true
}

// xsdDateTime
type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

// isValid returning true by default
func (t xsdDateTime) isValid() bool {
	return true
}

// ISODate
type ISODate time.Time

// UnmarshalText
func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}

// MarshalText
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// isValid returning true by default
func (t ISODate) isValid() bool {
	return true
}

// ISODateTime
type ISODateTime time.Time

// UnmarshalText
func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}

// MarshalText
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// isValid returning true by default
func (t ISODateTime) isValid() bool {
	return true
}
