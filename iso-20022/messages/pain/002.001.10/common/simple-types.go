// Package common
// Do not Edit. This stuff it's been automatically generated.
package common

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

// LEIIdentifier Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier xsdt.String

// ExternalFinancialInstitutionIdentification1Code May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code xsdt.String

// Frequency6Code May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code xsdt.String

// Exact2NumericText Must match the pattern [0-9]{2}
type Exact2NumericText xsdt.String

// ExternalDiscountAmountType1Code May be no more than 4 items long
type ExternalDiscountAmountType1Code xsdt.String

// Max140Text May be no more than 140 items long
type Max140Text xsdt.String

// Exact4AlphaNumericText Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText xsdt.String

// ExternalMandateSetupReason1Code May be no more than 4 items long
type ExternalMandateSetupReason1Code xsdt.String

// Max2048Text May be no more than 2048 items long
type Max2048Text xsdt.String

// Max105Text May be no more than 105 items long
type Max105Text xsdt.String

// ExternalDocumentLineType1Code May be no more than 4 items long
type ExternalDocumentLineType1Code xsdt.String

// Max15NumericText Must match the pattern [0-9]{1,15}
type Max15NumericText xsdt.String

// ExternalPaymentGroupStatus1Code May be no more than 4 items long
type ExternalPaymentGroupStatus1Code xsdt.String

// ExternalAccountIdentification1Code May be no more than 4 items long
type ExternalAccountIdentification1Code xsdt.String

// ExternalCashAccountType1Code May be no more than 4 items long
type ExternalCashAccountType1Code xsdt.String

// Max350Text May be no more than 350 items long
type Max350Text xsdt.String

// ExternalCashClearingSystem1Code May be no more than 3 items long
type ExternalCashClearingSystem1Code xsdt.String

// ClearingChannel2Code May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code xsdt.String

// ExternalServiceLevel1Code May be no more than 4 items long
type ExternalServiceLevel1Code xsdt.String

// DocumentType3Code May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code xsdt.String

// ExternalCategoryPurpose1Code May be no more than 4 items long
type ExternalCategoryPurpose1Code xsdt.String

// TaxRecordPeriod1Code May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code xsdt.String

// ExternalProxyAccountType1Code May be no more than 4 items long
type ExternalProxyAccountType1Code xsdt.String

// IBAN2007Identifier Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier xsdt.String

// Max1025Text May be no more than 1025 items long
type Max1025Text xsdt.String

// AddressType2Code May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code xsdt.String

// Max70Text May be no more than 70 items long
type Max70Text xsdt.String

// ExternalGarnishmentType1Code May be no more than 4 items long
type ExternalGarnishmentType1Code xsdt.String

// PreferredContactMethod1Code May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code xsdt.String

// CountryCode Must match the pattern [A-Z]{2,2}
type CountryCode xsdt.String

// Max4Text May be no more than 4 items long
type Max4Text xsdt.String

// ChargeBearerType1Code May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code xsdt.String

// SequenceType3Code May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code xsdt.String

// Max35Text May be no more than 35 items long
type Max35Text xsdt.String

// AnyBICDec2014Identifier Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier xsdt.String

// ExternalTaxAmountType1Code May be no more than 4 items long
type ExternalTaxAmountType1Code xsdt.String

// NamePrefix2Code May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code xsdt.String

// UUIDv4Identifier Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier xsdt.String

// BICFIDec2014Identifier Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier xsdt.String

// ExternalStatusReason1Code May be no more than 4 items long
type ExternalStatusReason1Code xsdt.String

// ExternalPaymentTransactionStatus1Code May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code xsdt.String

// Priority2Code May be one of HIGH, NORM
type Priority2Code xsdt.String

// ISODate type definition
type ISODate xsdt.Date

// ActiveOrHistoricCurrencyCode Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode xsdt.String

// ISODateTime type definition
type ISODateTime xsdt.DateTime

// DocumentType6Code May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code xsdt.String

// CreditDebitCode May be one of CRDT, DBIT
type CreditDebitCode xsdt.String

// ExternalClearingSystemIdentification1Code May be no more than 5 items long
type ExternalClearingSystemIdentification1Code xsdt.String

// ExternalLocalInstrument1Code May be no more than 35 items long
type ExternalLocalInstrument1Code xsdt.String

// Max16Text May be no more than 16 items long
type Max16Text xsdt.String

// PaymentMethod4Code May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code xsdt.String

// Max34Text May be no more than 34 items long
type Max34Text xsdt.String

// ExternalOrganisationIdentification1Code May be no more than 4 items long
type ExternalOrganisationIdentification1Code xsdt.String

// ExternalPersonIdentification1Code May be no more than 4 items long
type ExternalPersonIdentification1Code xsdt.String

// ExternalPurpose1Code May be no more than 4 items long
type ExternalPurpose1Code xsdt.String

// SettlementMethod1Code May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code xsdt.String

// ActiveCurrencyCode Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode xsdt.String

// PhoneNumber Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber xsdt.String

// Max128Text May be no more than 128 items long
type Max128Text xsdt.String