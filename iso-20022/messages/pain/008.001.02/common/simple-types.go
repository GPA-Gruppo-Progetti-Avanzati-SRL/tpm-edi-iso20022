// Package common
// Do not Edit. This stuff it's been automatically generated.
package common

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02/xsdt"
)

// AddressType2Code May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code xsdt.String

// Max70Text May be no more than 70 items long
type Max70Text xsdt.String

// CountryCode Must match the pattern [A-Z]{2,2}
type CountryCode xsdt.String

// ExternalAccountIdentification1Code May be no more than 4 items long
type ExternalAccountIdentification1Code xsdt.String

// RegulatoryReportingType1Code May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code xsdt.String

// ExternalServiceLevel1Code May be no more than 4 items long
type ExternalServiceLevel1Code xsdt.String

// PhoneNumber Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber xsdt.String

// ISODateTime type definition
type ISODateTime xsdt.DateTime

// ISODate type definition
type ISODate xsdt.Date

// ExternalCategoryPurpose1Code May be no more than 4 items long
type ExternalCategoryPurpose1Code xsdt.String

// ExternalPersonIdentification1Code May be no more than 4 items long
type ExternalPersonIdentification1Code xsdt.String

// DocumentType3Code May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code xsdt.String

// ExternalPurpose1Code May be no more than 4 items long
type ExternalPurpose1Code xsdt.String

// Max128Text May be no more than 128 items long
type Max128Text xsdt.String

// DocumentType5Code May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT
type DocumentType5Code xsdt.String

// BICIdentifier Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier xsdt.String

// CreditDebitCode May be one of CRDT, DBIT
type CreditDebitCode xsdt.String

// ExternalOrganisationIdentification1Code May be no more than 4 items long
type ExternalOrganisationIdentification1Code xsdt.String

// Priority2Code May be one of HIGH, NORM
type Priority2Code xsdt.String

// Max34Text May be no more than 34 items long
type Max34Text xsdt.String

// TaxRecordPeriod1Code May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code xsdt.String

// Max35Text May be no more than 35 items long
type Max35Text xsdt.String

// NamePrefix1Code May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code xsdt.String

// ActiveOrHistoricCurrencyCode Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode xsdt.String

// PaymentMethod2Code May be one of DD
type PaymentMethod2Code xsdt.String

// RemittanceLocationMethod2Code May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code xsdt.String

// ExternalFinancialInstitutionIdentification1Code May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code xsdt.String

// Max15NumericText Must match the pattern [0-9]{1,15}
type Max15NumericText xsdt.String

// AnyBICIdentifier Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier xsdt.String

// Authorisation1Code May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code xsdt.String

// ChargeBearerType1Code May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code xsdt.String

// Max2048Text May be no more than 2048 items long
type Max2048Text xsdt.String

// Frequency1Code May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code xsdt.String

// ExternalClearingSystemIdentification1Code May be no more than 5 items long
type ExternalClearingSystemIdentification1Code xsdt.String

// IBAN2007Identifier Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier xsdt.String

// Max1025Text May be no more than 1025 items long
type Max1025Text xsdt.String

// Max140Text May be no more than 140 items long
type Max140Text xsdt.String

// Max16Text May be no more than 16 items long
type Max16Text xsdt.String

// CashAccountType4Code May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code xsdt.String

// Max4Text May be no more than 4 items long
type Max4Text xsdt.String

// ExternalLocalInstrument1Code May be no more than 35 items long
type ExternalLocalInstrument1Code xsdt.String

// SequenceType1Code May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code xsdt.String

// Max10Text May be no more than 10 items long
type Max10Text xsdt.String
