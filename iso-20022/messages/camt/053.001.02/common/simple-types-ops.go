// Package common
// Do not Edit. This stuff it's been automatically generated.
package common

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"regexp"
	"time"
)

/*
 * CashAccountType4Code Ops
 */

const (
	CashAccountType4CodeZero   = ""
	CashAccountType4CodeSample = "SVGS"
	CashAccountType4CodeCASH   = "CASH"
	CashAccountType4CodeCHAR   = "CHAR"
	CashAccountType4CodeCOMM   = "COMM"
	CashAccountType4CodeTAXE   = "TAXE"
	CashAccountType4CodeCISH   = "CISH"
	CashAccountType4CodeTRAS   = "TRAS"
	CashAccountType4CodeSACC   = "SACC"
	CashAccountType4CodeCACC   = "CACC"
	CashAccountType4CodeSVGS   = "SVGS"
	CashAccountType4CodeONDP   = "ONDP"
	CashAccountType4CodeMGLD   = "MGLD"
	CashAccountType4CodeNREX   = "NREX"
	CashAccountType4CodeMOMA   = "MOMA"
	CashAccountType4CodeLOAN   = "LOAN"
	CashAccountType4CodeSLRY   = "SLRY"
	CashAccountType4CodeODFT   = "ODFT"
)

var CashAccountType4CodeEnumRestriction = []string{CashAccountType4CodeCASH, CashAccountType4CodeCHAR, CashAccountType4CodeCOMM, CashAccountType4CodeTAXE, CashAccountType4CodeCISH, CashAccountType4CodeTRAS, CashAccountType4CodeSACC, CashAccountType4CodeCACC, CashAccountType4CodeSVGS, CashAccountType4CodeONDP, CashAccountType4CodeMGLD, CashAccountType4CodeNREX, CashAccountType4CodeMOMA, CashAccountType4CodeLOAN, CashAccountType4CodeSLRY, CashAccountType4CodeODFT}

// IsValid checks if CashAccountType4Code of type String is valid
func (t CashAccountType4Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == CashAccountType4CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), CashAccountType4CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CashAccountType4Code) String() string {
	return string(t)
}

// ToCashAccountType4Code method for easy conversion with application of restrictions
func ToCashAccountType4Code(i interface{}) (CashAccountType4Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, CashAccountType4CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type CashAccountType4Code", s)
	}

	return CashAccountType4Code(s), nil
}

// MustToCashAccountType4Code method for easy conversion with application of restrictions. Panics on error.
func MustToCashAccountType4Code(s interface{}) CashAccountType4Code {
	v, err := ToCashAccountType4Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * EntryStatus2Code Ops
 */

const (
	EntryStatus2CodeZero   = ""
	EntryStatus2CodeSample = "PDNG"
	EntryStatus2CodeBOOK   = "BOOK"
	EntryStatus2CodePDNG   = "PDNG"
	EntryStatus2CodeINFO   = "INFO"
)

var EntryStatus2CodeEnumRestriction = []string{EntryStatus2CodeBOOK, EntryStatus2CodePDNG, EntryStatus2CodeINFO}

// IsValid checks if EntryStatus2Code of type String is valid
func (t EntryStatus2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == EntryStatus2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), EntryStatus2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t EntryStatus2Code) String() string {
	return string(t)
}

// ToEntryStatus2Code method for easy conversion with application of restrictions
func ToEntryStatus2Code(i interface{}) (EntryStatus2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, EntryStatus2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type EntryStatus2Code", s)
	}

	return EntryStatus2Code(s), nil
}

// MustToEntryStatus2Code method for easy conversion with application of restrictions. Panics on error.
func MustToEntryStatus2Code(s interface{}) EntryStatus2Code {
	v, err := ToEntryStatus2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalTechnicalInputChannel1Code Ops
 */

const (
	ExternalTechnicalInputChannel1CodeZero      = ""
	ExternalTechnicalInputChannel1CodeSample    = "oJ"
	ExternalTechnicalInputChannel1CodeLength    = 0
	ExternalTechnicalInputChannel1CodeMinLength = 1
	ExternalTechnicalInputChannel1CodeMaxLength = 4
)

// IsValid checks if ExternalTechnicalInputChannel1Code of type String is valid
func (t ExternalTechnicalInputChannel1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalTechnicalInputChannel1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalTechnicalInputChannel1CodeLength, ExternalTechnicalInputChannel1CodeMinLength, ExternalTechnicalInputChannel1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalTechnicalInputChannel1Code) String() string {
	return string(t)
}

// ToExternalTechnicalInputChannel1Code method for easy conversion with application of restrictions
func ToExternalTechnicalInputChannel1Code(i interface{}) (ExternalTechnicalInputChannel1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalTechnicalInputChannel1CodeLength, ExternalTechnicalInputChannel1CodeMinLength, ExternalTechnicalInputChannel1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalTechnicalInputChannel1Code", s)
	}

	return ExternalTechnicalInputChannel1Code(s), nil
}

// MustToExternalTechnicalInputChannel1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalTechnicalInputChannel1Code(s interface{}) ExternalTechnicalInputChannel1Code {
	v, err := ToExternalTechnicalInputChannel1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalFinancialInstitutionIdentification1Code Ops
 */

const (
	ExternalFinancialInstitutionIdentification1CodeZero      = ""
	ExternalFinancialInstitutionIdentification1CodeSample    = "nN"
	ExternalFinancialInstitutionIdentification1CodeLength    = 0
	ExternalFinancialInstitutionIdentification1CodeMinLength = 1
	ExternalFinancialInstitutionIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalFinancialInstitutionIdentification1Code of type String is valid
func (t ExternalFinancialInstitutionIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalFinancialInstitutionIdentification1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalFinancialInstitutionIdentification1CodeLength, ExternalFinancialInstitutionIdentification1CodeMinLength, ExternalFinancialInstitutionIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalFinancialInstitutionIdentification1Code) String() string {
	return string(t)
}

// ToExternalFinancialInstitutionIdentification1Code method for easy conversion with application of restrictions
func ToExternalFinancialInstitutionIdentification1Code(i interface{}) (ExternalFinancialInstitutionIdentification1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalFinancialInstitutionIdentification1CodeLength, ExternalFinancialInstitutionIdentification1CodeMinLength, ExternalFinancialInstitutionIdentification1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalFinancialInstitutionIdentification1Code", s)
	}

	return ExternalFinancialInstitutionIdentification1Code(s), nil
}

// MustToExternalFinancialInstitutionIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalFinancialInstitutionIdentification1Code(s interface{}) ExternalFinancialInstitutionIdentification1Code {
	v, err := ToExternalFinancialInstitutionIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalReturnReason1Code Ops
 */

const (
	ExternalReturnReason1CodeZero      = ""
	ExternalReturnReason1CodeSample    = "PG"
	ExternalReturnReason1CodeLength    = 0
	ExternalReturnReason1CodeMinLength = 1
	ExternalReturnReason1CodeMaxLength = 4
)

// IsValid checks if ExternalReturnReason1Code of type String is valid
func (t ExternalReturnReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalReturnReason1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalReturnReason1CodeLength, ExternalReturnReason1CodeMinLength, ExternalReturnReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalReturnReason1Code) String() string {
	return string(t)
}

// ToExternalReturnReason1Code method for easy conversion with application of restrictions
func ToExternalReturnReason1Code(i interface{}) (ExternalReturnReason1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalReturnReason1CodeLength, ExternalReturnReason1CodeMinLength, ExternalReturnReason1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalReturnReason1Code", s)
	}

	return ExternalReturnReason1Code(s), nil
}

// MustToExternalReturnReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalReturnReason1Code(s interface{}) ExternalReturnReason1Code {
	v, err := ToExternalReturnReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * DocumentType5Code Ops
 */

const (
	DocumentType5CodeZero   = ""
	DocumentType5CodeSample = "SBIN"
	DocumentType5CodeMSIN   = "MSIN"
	DocumentType5CodeCNFA   = "CNFA"
	DocumentType5CodeDNFA   = "DNFA"
	DocumentType5CodeCINV   = "CINV"
	DocumentType5CodeCREN   = "CREN"
	DocumentType5CodeDEBN   = "DEBN"
	DocumentType5CodeHIRI   = "HIRI"
	DocumentType5CodeSBIN   = "SBIN"
	DocumentType5CodeCMCN   = "CMCN"
	DocumentType5CodeSOAC   = "SOAC"
	DocumentType5CodeDISP   = "DISP"
	DocumentType5CodeBOLD   = "BOLD"
	DocumentType5CodeVCHR   = "VCHR"
	DocumentType5CodeAROI   = "AROI"
	DocumentType5CodeTSUT   = "TSUT"
)

var DocumentType5CodeEnumRestriction = []string{DocumentType5CodeMSIN, DocumentType5CodeCNFA, DocumentType5CodeDNFA, DocumentType5CodeCINV, DocumentType5CodeCREN, DocumentType5CodeDEBN, DocumentType5CodeHIRI, DocumentType5CodeSBIN, DocumentType5CodeCMCN, DocumentType5CodeSOAC, DocumentType5CodeDISP, DocumentType5CodeBOLD, DocumentType5CodeVCHR, DocumentType5CodeAROI, DocumentType5CodeTSUT}

// IsValid checks if DocumentType5Code of type String is valid
func (t DocumentType5Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == DocumentType5CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), DocumentType5CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType5Code) String() string {
	return string(t)
}

// ToDocumentType5Code method for easy conversion with application of restrictions
func ToDocumentType5Code(i interface{}) (DocumentType5Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, DocumentType5CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type DocumentType5Code", s)
	}

	return DocumentType5Code(s), nil
}

// MustToDocumentType5Code method for easy conversion with application of restrictions. Panics on error.
func MustToDocumentType5Code(s interface{}) DocumentType5Code {
	v, err := ToDocumentType5Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max2048Text Ops
 */

const (
	Max2048TextZero      = ""
	Max2048TextSample    = "siuzytMOJPatwtPilfsfykSBGplhxtxVSGpqaJaBRgAvzLXqzRrrUIYvaIujDpHYjxeUBrVfdwUzWHRihFDPRHBMuEWmaNvhHLITWJcJKzJsRPeJbpTqWPUlWAcidRlufmYWddEktBDrWMYZTPdmDAVCuvQutLDKFBGRxnyDNYuRKWsmZnOYgSLPSsuMMmvYmaFEjLgSvndTeranxGMNCcCdlEBYqxqrgAieJuFVZUsgPmweLRmsdTIwtLUDqoBSNtPRViLXLCrSAZCjxTCaZnfHRhEbbvlLkCbvOlRrOWJFNwNQBqyozKVGeninyrTgqubNiIzWcEZsdyUJTDAODnRKcTqZThaTJmRVAClOEaCEmEUmEtNTDVBbvMBtQBeDBVyDFfPWGaZQntPiyJbzjrRNLRzGpzDszPvweGxRQcoxpAQxuFkaowCKoRccTyTJHeKKokJNwUaqKCBRxMSdypnuVQlXTxlLUfRUssUCbZClckQhhyMmVTDwokzVJUlvROmQBJTVsrIgXXlvECpbDIWegrUuUYynIaqJjgVXQtZptcLZIVmypNpLbbaxHYDmKrYnEbqbvWLjZZAHQHWIGRYPwqfeXSWmRncQZpTAzceEsIpdYbrhfkVRdPhRXJzEKTBNCvokhNhSeFRvuKeklvsHwabufMjYxuvCIeoMjxivfYHKXcwedxkRvEookKzpacqgBiXejXjTJPhMQsmGzinFnyUVhVOydnINYczGVyZhBRwWAXjuufUoPJlNOFUvANpjrlwiuIbjJbgkIxMkImqLfMdmksNxvjgyEJAjrjcQDUlVTuFKDLjgBjrpEBlRzIZVSUAgxlZPBVcodWJbYmpBEprbxsRnSFqRIjOBrNaxVzTtXlnbdWOKzddfZvtzFRwXkWBBNMmMfCraUahSxpTkppSxjPLhPnxeSCjulnkeSRvUtXaQHrOqxrwbCXRiWGXOyolORPasKWOdANVLTUjsQCqRuNsAzgVRDgWTxeohgOVmZWEReHPcpuWoyejN"
	Max2048TextLength    = 0
	Max2048TextMinLength = 1
	Max2048TextMaxLength = 2048
)

// IsValid checks if Max2048Text of type String is valid
func (t Max2048Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max2048TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max2048TextLength, Max2048TextMinLength, Max2048TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max2048Text) String() string {
	return string(t)
}

// ToMax2048Text method for easy conversion with application of restrictions
func ToMax2048Text(i interface{}) (Max2048Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max2048TextLength, Max2048TextMinLength, Max2048TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max2048Text", s)
	}

	return Max2048Text(s), nil
}

// MustToMax2048Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax2048Text(s interface{}) Max2048Text {
	v, err := ToMax2048Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ActiveOrHistoricCurrencyCode Ops
 */

const (
	ActiveOrHistoricCurrencyCodeZero   = ""
	ActiveOrHistoricCurrencyCodeSample = "NYN"
)

var ActiveOrHistoricCurrencyCodePatternRestriction = regexp.MustCompile(`[A-Z]{3,3}`)

// IsValid checks if ActiveOrHistoricCurrencyCode of type String is valid
func (t ActiveOrHistoricCurrencyCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ActiveOrHistoricCurrencyCodeZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), ActiveOrHistoricCurrencyCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t ActiveOrHistoricCurrencyCode) String() string {
	return string(t)
}

// ToActiveOrHistoricCurrencyCode method for easy conversion with application of restrictions
func ToActiveOrHistoricCurrencyCode(i interface{}) (ActiveOrHistoricCurrencyCode, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, ActiveOrHistoricCurrencyCodePatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type ActiveOrHistoricCurrencyCode", s)
	}

	return ActiveOrHistoricCurrencyCode(s), nil
}

// MustToActiveOrHistoricCurrencyCode method for easy conversion with application of restrictions. Panics on error.
func MustToActiveOrHistoricCurrencyCode(s interface{}) ActiveOrHistoricCurrencyCode {
	v, err := ToActiveOrHistoricCurrencyCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ChargeType1Code Ops
 */

const (
	ChargeType1CodeZero   = ""
	ChargeType1CodeSample = "COMM"
	ChargeType1CodeBRKF   = "BRKF"
	ChargeType1CodeCOMM   = "COMM"
)

var ChargeType1CodeEnumRestriction = []string{ChargeType1CodeBRKF, ChargeType1CodeCOMM}

// IsValid checks if ChargeType1Code of type String is valid
func (t ChargeType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ChargeType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ChargeType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeType1Code) String() string {
	return string(t)
}

// ToChargeType1Code method for easy conversion with application of restrictions
func ToChargeType1Code(i interface{}) (ChargeType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ChargeType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChargeType1Code", s)
	}

	return ChargeType1Code(s), nil
}

// MustToChargeType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChargeType1Code(s interface{}) ChargeType1Code {
	v, err := ToChargeType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max140Text Ops
 */

const (
	Max140TextZero      = ""
	Max140TextSample    = "DNtfSJWCidgAfEcoSKFvxTVCZoKMHuKFgkhaujRMxfcNGDjkbfeENRdnZuYVnOFfaKgXcx"
	Max140TextLength    = 0
	Max140TextMinLength = 1
	Max140TextMaxLength = 140
)

// IsValid checks if Max140Text of type String is valid
func (t Max140Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max140TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max140TextLength, Max140TextMinLength, Max140TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max140Text) String() string {
	return string(t)
}

// ToMax140Text method for easy conversion with application of restrictions
func ToMax140Text(i interface{}) (Max140Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max140TextLength, Max140TextMinLength, Max140TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max140Text", s)
	}

	return Max140Text(s), nil
}

// MustToMax140Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax140Text(s interface{}) Max140Text {
	v, err := ToMax140Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max16Text Ops
 */

const (
	Max16TextZero      = ""
	Max16TextSample    = "fcXVutzu"
	Max16TextLength    = 0
	Max16TextMinLength = 1
	Max16TextMaxLength = 16
)

// IsValid checks if Max16Text of type String is valid
func (t Max16Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max16TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max16TextLength, Max16TextMinLength, Max16TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max16Text) String() string {
	return string(t)
}

// ToMax16Text method for easy conversion with application of restrictions
func ToMax16Text(i interface{}) (Max16Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max16TextLength, Max16TextMinLength, Max16TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max16Text", s)
	}

	return Max16Text(s), nil
}

// MustToMax16Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax16Text(s interface{}) Max16Text {
	v, err := ToMax16Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * BICIdentifier Ops
 */

const (
	BICIdentifierZero   = ""
	BICIdentifierSample = "ABSPAHDF"
)

var BICIdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if BICIdentifier of type String is valid
func (t BICIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == BICIdentifierZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), BICIdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t BICIdentifier) String() string {
	return string(t)
}

// ToBICIdentifier method for easy conversion with application of restrictions
func ToBICIdentifier(i interface{}) (BICIdentifier, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, BICIdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type BICIdentifier", s)
	}

	return BICIdentifier(s), nil
}

// MustToBICIdentifier method for easy conversion with application of restrictions. Panics on error.
func MustToBICIdentifier(s interface{}) BICIdentifier {
	v, err := ToBICIdentifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max5NumericText Ops
 */

const (
	Max5NumericTextZero   = ""
	Max5NumericTextSample = "22"
)

var Max5NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{1,5}`)

// IsValid checks if Max5NumericText of type String is valid
func (t Max5NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max5NumericTextZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), Max5NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Max5NumericText) String() string {
	return string(t)
}

// ToMax5NumericText method for easy conversion with application of restrictions
func ToMax5NumericText(i interface{}) (Max5NumericText, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, Max5NumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Max5NumericText", s)
	}

	return Max5NumericText(s), nil
}

// MustToMax5NumericText method for easy conversion with application of restrictions. Panics on error.
func MustToMax5NumericText(s interface{}) Max5NumericText {
	v, err := ToMax5NumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ChargeBearerType1Code Ops
 */

const (
	ChargeBearerType1CodeZero   = ""
	ChargeBearerType1CodeSample = "SHAR"
	ChargeBearerType1CodeDEBT   = "DEBT"
	ChargeBearerType1CodeCRED   = "CRED"
	ChargeBearerType1CodeSHAR   = "SHAR"
	ChargeBearerType1CodeSLEV   = "SLEV"
)

var ChargeBearerType1CodeEnumRestriction = []string{ChargeBearerType1CodeDEBT, ChargeBearerType1CodeCRED, ChargeBearerType1CodeSHAR, ChargeBearerType1CodeSLEV}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ChargeBearerType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeBearerType1Code) String() string {
	return string(t)
}

// ToChargeBearerType1Code method for easy conversion with application of restrictions
func ToChargeBearerType1Code(i interface{}) (ChargeBearerType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ChargeBearerType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChargeBearerType1Code", s)
	}

	return ChargeBearerType1Code(s), nil
}

// MustToChargeBearerType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChargeBearerType1Code(s interface{}) ChargeBearerType1Code {
	v, err := ToChargeBearerType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max34Text Ops
 */

const (
	Max34TextZero      = ""
	Max34TextSample    = "IkhdUMOniOqWYIJMq"
	Max34TextLength    = 0
	Max34TextMinLength = 1
	Max34TextMaxLength = 34
)

// IsValid checks if Max34Text of type String is valid
func (t Max34Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max34TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max34TextLength, Max34TextMinLength, Max34TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max34Text) String() string {
	return string(t)
}

// ToMax34Text method for easy conversion with application of restrictions
func ToMax34Text(i interface{}) (Max34Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max34TextLength, Max34TextMinLength, Max34TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max34Text", s)
	}

	return Max34Text(s), nil
}

// MustToMax34Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax34Text(s interface{}) Max34Text {
	v, err := ToMax34Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPersonIdentification1Code Ops
 */

const (
	ExternalPersonIdentification1CodeZero      = ""
	ExternalPersonIdentification1CodeSample    = "Lc"
	ExternalPersonIdentification1CodeLength    = 0
	ExternalPersonIdentification1CodeMinLength = 1
	ExternalPersonIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalPersonIdentification1Code of type String is valid
func (t ExternalPersonIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalPersonIdentification1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPersonIdentification1CodeLength, ExternalPersonIdentification1CodeMinLength, ExternalPersonIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPersonIdentification1Code) String() string {
	return string(t)
}

// ToExternalPersonIdentification1Code method for easy conversion with application of restrictions
func ToExternalPersonIdentification1Code(i interface{}) (ExternalPersonIdentification1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalPersonIdentification1CodeLength, ExternalPersonIdentification1CodeMinLength, ExternalPersonIdentification1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPersonIdentification1Code", s)
	}

	return ExternalPersonIdentification1Code(s), nil
}

// MustToExternalPersonIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPersonIdentification1Code(s interface{}) ExternalPersonIdentification1Code {
	v, err := ToExternalPersonIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CopyDuplicate1Code Ops
 */

const (
	CopyDuplicate1CodeZero   = ""
	CopyDuplicate1CodeSample = "COPY"
	CopyDuplicate1CodeCODU   = "CODU"
	CopyDuplicate1CodeCOPY   = "COPY"
	CopyDuplicate1CodeDUPL   = "DUPL"
)

var CopyDuplicate1CodeEnumRestriction = []string{CopyDuplicate1CodeCODU, CopyDuplicate1CodeCOPY, CopyDuplicate1CodeDUPL}

// IsValid checks if CopyDuplicate1Code of type String is valid
func (t CopyDuplicate1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == CopyDuplicate1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), CopyDuplicate1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CopyDuplicate1Code) String() string {
	return string(t)
}

// ToCopyDuplicate1Code method for easy conversion with application of restrictions
func ToCopyDuplicate1Code(i interface{}) (CopyDuplicate1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, CopyDuplicate1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type CopyDuplicate1Code", s)
	}

	return CopyDuplicate1Code(s), nil
}

// MustToCopyDuplicate1Code method for easy conversion with application of restrictions. Panics on error.
func MustToCopyDuplicate1Code(s interface{}) CopyDuplicate1Code {
	v, err := ToCopyDuplicate1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CreditDebitCode Ops
 */

const (
	CreditDebitCodeZero   = ""
	CreditDebitCodeSample = "DBIT"
	CreditDebitCodeCRDT   = "CRDT"
	CreditDebitCodeDBIT   = "DBIT"
)

var CreditDebitCodeEnumRestriction = []string{CreditDebitCodeCRDT, CreditDebitCodeDBIT}

// IsValid checks if CreditDebitCode of type String is valid
func (t CreditDebitCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == CreditDebitCodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), CreditDebitCodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t CreditDebitCode) String() string {
	return string(t)
}

// ToCreditDebitCode method for easy conversion with application of restrictions
func ToCreditDebitCode(i interface{}) (CreditDebitCode, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, CreditDebitCodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type CreditDebitCode", s)
	}

	return CreditDebitCode(s), nil
}

// MustToCreditDebitCode method for easy conversion with application of restrictions. Panics on error.
func MustToCreditDebitCode(s interface{}) CreditDebitCode {
	v, err := ToCreditDebitCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * TaxRecordPeriod1Code Ops
 */

const (
	TaxRecordPeriod1CodeZero   = ""
	TaxRecordPeriod1CodeSample = "MM10"
	TaxRecordPeriod1CodeMM01   = "MM01"
	TaxRecordPeriod1CodeMM02   = "MM02"
	TaxRecordPeriod1CodeMM03   = "MM03"
	TaxRecordPeriod1CodeMM04   = "MM04"
	TaxRecordPeriod1CodeMM05   = "MM05"
	TaxRecordPeriod1CodeMM06   = "MM06"
	TaxRecordPeriod1CodeMM07   = "MM07"
	TaxRecordPeriod1CodeMM08   = "MM08"
	TaxRecordPeriod1CodeMM09   = "MM09"
	TaxRecordPeriod1CodeMM10   = "MM10"
	TaxRecordPeriod1CodeMM11   = "MM11"
	TaxRecordPeriod1CodeMM12   = "MM12"
	TaxRecordPeriod1CodeQTR1   = "QTR1"
	TaxRecordPeriod1CodeQTR2   = "QTR2"
	TaxRecordPeriod1CodeQTR3   = "QTR3"
	TaxRecordPeriod1CodeQTR4   = "QTR4"
	TaxRecordPeriod1CodeHLF1   = "HLF1"
	TaxRecordPeriod1CodeHLF2   = "HLF2"
)

var TaxRecordPeriod1CodeEnumRestriction = []string{TaxRecordPeriod1CodeMM01, TaxRecordPeriod1CodeMM02, TaxRecordPeriod1CodeMM03, TaxRecordPeriod1CodeMM04, TaxRecordPeriod1CodeMM05, TaxRecordPeriod1CodeMM06, TaxRecordPeriod1CodeMM07, TaxRecordPeriod1CodeMM08, TaxRecordPeriod1CodeMM09, TaxRecordPeriod1CodeMM10, TaxRecordPeriod1CodeMM11, TaxRecordPeriod1CodeMM12, TaxRecordPeriod1CodeQTR1, TaxRecordPeriod1CodeQTR2, TaxRecordPeriod1CodeQTR3, TaxRecordPeriod1CodeQTR4, TaxRecordPeriod1CodeHLF1, TaxRecordPeriod1CodeHLF2}

// IsValid checks if TaxRecordPeriod1Code of type String is valid
func (t TaxRecordPeriod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == TaxRecordPeriod1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), TaxRecordPeriod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TaxRecordPeriod1Code) String() string {
	return string(t)
}

// ToTaxRecordPeriod1Code method for easy conversion with application of restrictions
func ToTaxRecordPeriod1Code(i interface{}) (TaxRecordPeriod1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, TaxRecordPeriod1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type TaxRecordPeriod1Code", s)
	}

	return TaxRecordPeriod1Code(s), nil
}

// MustToTaxRecordPeriod1Code method for easy conversion with application of restrictions. Panics on error.
func MustToTaxRecordPeriod1Code(s interface{}) TaxRecordPeriod1Code {
	v, err := ToTaxRecordPeriod1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalBalanceSubType1Code Ops
 */

const (
	ExternalBalanceSubType1CodeZero      = ""
	ExternalBalanceSubType1CodeSample    = "hM"
	ExternalBalanceSubType1CodeLength    = 0
	ExternalBalanceSubType1CodeMinLength = 1
	ExternalBalanceSubType1CodeMaxLength = 4
)

// IsValid checks if ExternalBalanceSubType1Code of type String is valid
func (t ExternalBalanceSubType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalBalanceSubType1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalBalanceSubType1CodeLength, ExternalBalanceSubType1CodeMinLength, ExternalBalanceSubType1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalBalanceSubType1Code) String() string {
	return string(t)
}

// ToExternalBalanceSubType1Code method for easy conversion with application of restrictions
func ToExternalBalanceSubType1Code(i interface{}) (ExternalBalanceSubType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalBalanceSubType1CodeLength, ExternalBalanceSubType1CodeMinLength, ExternalBalanceSubType1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalBalanceSubType1Code", s)
	}

	return ExternalBalanceSubType1Code(s), nil
}

// MustToExternalBalanceSubType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalBalanceSubType1Code(s interface{}) ExternalBalanceSubType1Code {
	v, err := ToExternalBalanceSubType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * CountryCode Ops
 */

const (
	CountryCodeZero   = ""
	CountryCodeSample = "HZ"
)

var CountryCodePatternRestriction = regexp.MustCompile(`[A-Z]{2,2}`)

// IsValid checks if CountryCode of type String is valid
func (t CountryCode) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == CountryCodeZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), CountryCodePatternRestriction)

	return valid
}

// String method for easy conversion
func (t CountryCode) String() string {
	return string(t)
}

// ToCountryCode method for easy conversion with application of restrictions
func ToCountryCode(i interface{}) (CountryCode, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, CountryCodePatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type CountryCode", s)
	}

	return CountryCode(s), nil
}

// MustToCountryCode method for easy conversion with application of restrictions. Panics on error.
func MustToCountryCode(s interface{}) CountryCode {
	v, err := ToCountryCode(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ISODate Ops
 */

const (
	ISODateSample = "Value"
	ISODateZero   = ""
)

// IsValid checks if ISODate of type Date is valid
func (t ISODate) IsValid(optional bool) bool {

	valid := xsdt.Date(t).IsValid(optional)
	if optional && t == ISODateZero {
		return valid
	}
	return valid
}

// String method for easy conversion
func (t ISODate) String() string {
	return string(t)
}

// ToISODate method for easy conversion from time.Time
func ToISODate(tm interface{}) (ISODate, error) {

	switch typedTm := tm.(type) {
	case time.Time:
		return ISODate(typedTm.Format("2006-01-02")), nil
	case string:
		return ISODate(typedTm), nil
	case ISODate:
		return typedTm, nil
	}

	return "", fmt.Errorf("cannot convert %v to ISODate", tm)
}

func MustToISODate(tm interface{}) ISODate {
	d, err := ToISODate(tm)
	if err != nil {
		panic(err)
	}

	return d
}

// ISODateExample method for generation of valid sample data
func ISODateExample() ISODate {
	return ISODate(time.Now().Format("2006-01-02"))
}

/*
 * Max500Text Ops
 */

const (
	Max500TextZero      = ""
	Max500TextSample    = "kCJZHSenoYIFnHZyTzEfMSdhUvtMBlfUoZYLEdqkoZCuTuRpLKzzuuFjTKIOTjSnLTFioYayJwBotODVVVIrFCqrhOGVehnNZfARfsBVGBRaHjfFcQtXovQyXbjAyWbspgheqXXumwlaMkvqgeVNsLyylTpEsGuymNstQzxPbmXGYPZRoCcAjYwjicELciiQYDuPezahlBGYsXsQyazOWcrSbAuTepBSUGepQDYBzMBKXUbWscOLULiIiQ"
	Max500TextLength    = 0
	Max500TextMinLength = 1
	Max500TextMaxLength = 500
)

// IsValid checks if Max500Text of type String is valid
func (t Max500Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max500TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max500TextLength, Max500TextMinLength, Max500TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max500Text) String() string {
	return string(t)
}

// ToMax500Text method for easy conversion with application of restrictions
func ToMax500Text(i interface{}) (Max500Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max500TextLength, Max500TextMinLength, Max500TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max500Text", s)
	}

	return Max500Text(s), nil
}

// MustToMax500Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax500Text(s interface{}) Max500Text {
	v, err := ToMax500Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalClearingSystemIdentification1Code Ops
 */

const (
	ExternalClearingSystemIdentification1CodeZero      = ""
	ExternalClearingSystemIdentification1CodeSample    = "tPi"
	ExternalClearingSystemIdentification1CodeLength    = 0
	ExternalClearingSystemIdentification1CodeMinLength = 1
	ExternalClearingSystemIdentification1CodeMaxLength = 5
)

// IsValid checks if ExternalClearingSystemIdentification1Code of type String is valid
func (t ExternalClearingSystemIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalClearingSystemIdentification1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalClearingSystemIdentification1CodeLength, ExternalClearingSystemIdentification1CodeMinLength, ExternalClearingSystemIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalClearingSystemIdentification1Code) String() string {
	return string(t)
}

// ToExternalClearingSystemIdentification1Code method for easy conversion with application of restrictions
func ToExternalClearingSystemIdentification1Code(i interface{}) (ExternalClearingSystemIdentification1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalClearingSystemIdentification1CodeLength, ExternalClearingSystemIdentification1CodeMinLength, ExternalClearingSystemIdentification1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalClearingSystemIdentification1Code", s)
	}

	return ExternalClearingSystemIdentification1Code(s), nil
}

// MustToExternalClearingSystemIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalClearingSystemIdentification1Code(s interface{}) ExternalClearingSystemIdentification1Code {
	v, err := ToExternalClearingSystemIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max105Text Ops
 */

const (
	Max105TextZero      = ""
	Max105TextSample    = "PomlXcvuojfHBHOQiGrmoSuelNipGLXawNpUCRNoVMTKOGgxNQPpR"
	Max105TextLength    = 0
	Max105TextMinLength = 1
	Max105TextMaxLength = 105
)

// IsValid checks if Max105Text of type String is valid
func (t Max105Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max105TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max105TextLength, Max105TextMinLength, Max105TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max105Text) String() string {
	return string(t)
}

// ToMax105Text method for easy conversion with application of restrictions
func ToMax105Text(i interface{}) (Max105Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max105TextLength, Max105TextMinLength, Max105TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max105Text", s)
	}

	return Max105Text(s), nil
}

// MustToMax105Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax105Text(s interface{}) Max105Text {
	v, err := ToMax105Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalPurpose1Code Ops
 */

const (
	ExternalPurpose1CodeZero      = ""
	ExternalPurpose1CodeSample    = "iW"
	ExternalPurpose1CodeLength    = 0
	ExternalPurpose1CodeMinLength = 1
	ExternalPurpose1CodeMaxLength = 4
)

// IsValid checks if ExternalPurpose1Code of type String is valid
func (t ExternalPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalPurpose1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalPurpose1CodeLength, ExternalPurpose1CodeMinLength, ExternalPurpose1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalPurpose1Code) String() string {
	return string(t)
}

// ToExternalPurpose1Code method for easy conversion with application of restrictions
func ToExternalPurpose1Code(i interface{}) (ExternalPurpose1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalPurpose1CodeLength, ExternalPurpose1CodeMinLength, ExternalPurpose1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalPurpose1Code", s)
	}

	return ExternalPurpose1Code(s), nil
}

// MustToExternalPurpose1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalPurpose1Code(s interface{}) ExternalPurpose1Code {
	v, err := ToExternalPurpose1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * InterestType1Code Ops
 */

const (
	InterestType1CodeZero   = ""
	InterestType1CodeSample = "OVRN"
	InterestType1CodeINDY   = "INDY"
	InterestType1CodeOVRN   = "OVRN"
)

var InterestType1CodeEnumRestriction = []string{InterestType1CodeINDY, InterestType1CodeOVRN}

// IsValid checks if InterestType1Code of type String is valid
func (t InterestType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == InterestType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), InterestType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t InterestType1Code) String() string {
	return string(t)
}

// ToInterestType1Code method for easy conversion with application of restrictions
func ToInterestType1Code(i interface{}) (InterestType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, InterestType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type InterestType1Code", s)
	}

	return InterestType1Code(s), nil
}

// MustToInterestType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToInterestType1Code(s interface{}) InterestType1Code {
	v, err := ToInterestType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * DocumentType3Code Ops
 */

const (
	DocumentType3CodeZero   = ""
	DocumentType3CodeSample = "DISP"
	DocumentType3CodeRADM   = "RADM"
	DocumentType3CodeRPIN   = "RPIN"
	DocumentType3CodeFXDR   = "FXDR"
	DocumentType3CodeDISP   = "DISP"
	DocumentType3CodePUOR   = "PUOR"
	DocumentType3CodeSCOR   = "SCOR"
)

var DocumentType3CodeEnumRestriction = []string{DocumentType3CodeRADM, DocumentType3CodeRPIN, DocumentType3CodeFXDR, DocumentType3CodeDISP, DocumentType3CodePUOR, DocumentType3CodeSCOR}

// IsValid checks if DocumentType3Code of type String is valid
func (t DocumentType3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == DocumentType3CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), DocumentType3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t DocumentType3Code) String() string {
	return string(t)
}

// ToDocumentType3Code method for easy conversion with application of restrictions
func ToDocumentType3Code(i interface{}) (DocumentType3Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, DocumentType3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type DocumentType3Code", s)
	}

	return DocumentType3Code(s), nil
}

// MustToDocumentType3Code method for easy conversion with application of restrictions. Panics on error.
func MustToDocumentType3Code(s interface{}) DocumentType3Code {
	v, err := ToDocumentType3Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * AnyBICIdentifier Ops
 */

const (
	AnyBICIdentifierZero   = ""
	AnyBICIdentifierSample = "MFRVSUPE"
)

var AnyBICIdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)

// IsValid checks if AnyBICIdentifier of type String is valid
func (t AnyBICIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == AnyBICIdentifierZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), AnyBICIdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t AnyBICIdentifier) String() string {
	return string(t)
}

// ToAnyBICIdentifier method for easy conversion with application of restrictions
func ToAnyBICIdentifier(i interface{}) (AnyBICIdentifier, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, AnyBICIdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type AnyBICIdentifier", s)
	}

	return AnyBICIdentifier(s), nil
}

// MustToAnyBICIdentifier method for easy conversion with application of restrictions. Panics on error.
func MustToAnyBICIdentifier(s interface{}) AnyBICIdentifier {
	v, err := ToAnyBICIdentifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NamePrefix1Code Ops
 */

const (
	NamePrefix1CodeZero   = ""
	NamePrefix1CodeSample = "MISS"
	NamePrefix1CodeDOCT   = "DOCT"
	NamePrefix1CodeMIST   = "MIST"
	NamePrefix1CodeMISS   = "MISS"
	NamePrefix1CodeMADM   = "MADM"
)

var NamePrefix1CodeEnumRestriction = []string{NamePrefix1CodeDOCT, NamePrefix1CodeMIST, NamePrefix1CodeMISS, NamePrefix1CodeMADM}

// IsValid checks if NamePrefix1Code of type String is valid
func (t NamePrefix1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == NamePrefix1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), NamePrefix1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t NamePrefix1Code) String() string {
	return string(t)
}

// ToNamePrefix1Code method for easy conversion with application of restrictions
func ToNamePrefix1Code(i interface{}) (NamePrefix1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, NamePrefix1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type NamePrefix1Code", s)
	}

	return NamePrefix1Code(s), nil
}

// MustToNamePrefix1Code method for easy conversion with application of restrictions. Panics on error.
func MustToNamePrefix1Code(s interface{}) NamePrefix1Code {
	v, err := ToNamePrefix1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalReportingSource1Code Ops
 */

const (
	ExternalReportingSource1CodeZero      = ""
	ExternalReportingSource1CodeSample    = "mN"
	ExternalReportingSource1CodeLength    = 0
	ExternalReportingSource1CodeMinLength = 1
	ExternalReportingSource1CodeMaxLength = 4
)

// IsValid checks if ExternalReportingSource1Code of type String is valid
func (t ExternalReportingSource1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalReportingSource1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalReportingSource1CodeLength, ExternalReportingSource1CodeMinLength, ExternalReportingSource1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalReportingSource1Code) String() string {
	return string(t)
}

// ToExternalReportingSource1Code method for easy conversion with application of restrictions
func ToExternalReportingSource1Code(i interface{}) (ExternalReportingSource1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalReportingSource1CodeLength, ExternalReportingSource1CodeMinLength, ExternalReportingSource1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalReportingSource1Code", s)
	}

	return ExternalReportingSource1Code(s), nil
}

// MustToExternalReportingSource1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalReportingSource1Code(s interface{}) ExternalReportingSource1Code {
	v, err := ToExternalReportingSource1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ISINIdentifier Ops
 */

const (
	ISINIdentifierZero   = ""
	ISINIdentifierSample = "3QXFELC5MVFY"
)

var ISINIdentifierPatternRestriction = regexp.MustCompile(`[A-Z0-9]{12,12}`)

// IsValid checks if ISINIdentifier of type String is valid
func (t ISINIdentifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ISINIdentifierZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), ISINIdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t ISINIdentifier) String() string {
	return string(t)
}

// ToISINIdentifier method for easy conversion with application of restrictions
func ToISINIdentifier(i interface{}) (ISINIdentifier, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, ISINIdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type ISINIdentifier", s)
	}

	return ISINIdentifier(s), nil
}

// MustToISINIdentifier method for easy conversion with application of restrictions. Panics on error.
func MustToISINIdentifier(s interface{}) ISINIdentifier {
	v, err := ToISINIdentifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * AddressType2Code Ops
 */

const (
	AddressType2CodeZero   = ""
	AddressType2CodeSample = "BIZZ"
	AddressType2CodeADDR   = "ADDR"
	AddressType2CodePBOX   = "PBOX"
	AddressType2CodeHOME   = "HOME"
	AddressType2CodeBIZZ   = "BIZZ"
	AddressType2CodeMLTO   = "MLTO"
	AddressType2CodeDLVY   = "DLVY"
)

var AddressType2CodeEnumRestriction = []string{AddressType2CodeADDR, AddressType2CodePBOX, AddressType2CodeHOME, AddressType2CodeBIZZ, AddressType2CodeMLTO, AddressType2CodeDLVY}

// IsValid checks if AddressType2Code of type String is valid
func (t AddressType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == AddressType2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), AddressType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t AddressType2Code) String() string {
	return string(t)
}

// ToAddressType2Code method for easy conversion with application of restrictions
func ToAddressType2Code(i interface{}) (AddressType2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, AddressType2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type AddressType2Code", s)
	}

	return AddressType2Code(s), nil
}

// MustToAddressType2Code method for easy conversion with application of restrictions. Panics on error.
func MustToAddressType2Code(s interface{}) AddressType2Code {
	v, err := ToAddressType2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalOrganisationIdentification1Code Ops
 */

const (
	ExternalOrganisationIdentification1CodeZero      = ""
	ExternalOrganisationIdentification1CodeSample    = "uF"
	ExternalOrganisationIdentification1CodeLength    = 0
	ExternalOrganisationIdentification1CodeMinLength = 1
	ExternalOrganisationIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalOrganisationIdentification1Code of type String is valid
func (t ExternalOrganisationIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalOrganisationIdentification1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalOrganisationIdentification1CodeLength, ExternalOrganisationIdentification1CodeMinLength, ExternalOrganisationIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalOrganisationIdentification1Code) String() string {
	return string(t)
}

// ToExternalOrganisationIdentification1Code method for easy conversion with application of restrictions
func ToExternalOrganisationIdentification1Code(i interface{}) (ExternalOrganisationIdentification1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalOrganisationIdentification1CodeLength, ExternalOrganisationIdentification1CodeMinLength, ExternalOrganisationIdentification1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalOrganisationIdentification1Code", s)
	}

	return ExternalOrganisationIdentification1Code(s), nil
}

// MustToExternalOrganisationIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalOrganisationIdentification1Code(s interface{}) ExternalOrganisationIdentification1Code {
	v, err := ToExternalOrganisationIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ISODateTime Ops
 */

const (
	ISODateTimeSample = "Value"
	ISODateTimeZero   = ""
)

// IsValid checks if ISODateTime of type DateTime is valid
func (t ISODateTime) IsValid(optional bool) bool {

	valid := xsdt.DateTime(t).IsValid(optional)
	if optional && t == ISODateTimeZero {
		return valid
	}
	return valid
}

// String method for easy conversion
func (t ISODateTime) String() string {
	return string(t)
}

// ToISODateTime method for easy conversion from time.Time
func ToISODateTime(tm interface{}) (ISODateTime, error) {

	switch typedTm := tm.(type) {
	case time.Time:
		return ISODateTime(typedTm.Format(time.RFC3339)), nil
	case string:
		return ISODateTime(typedTm), nil
	case ISODateTime:
		return typedTm, nil
	}

	return "", fmt.Errorf("cannot convert %v to ISODateTime", tm)
}

func MustToISODateTime(tm interface{}) ISODateTime {
	d, err := ToISODateTime(tm)
	if err != nil {
		panic(err)
	}

	return d
}

// ISODateTimeExample method for generation of valid sample data
func ISODateTimeExample() ISODateTime {
	return ISODateTime(time.Now().Format(time.RFC3339))
}

/*
 * PhoneNumber Ops
 */

const (
	PhoneNumberZero   = ""
	PhoneNumberSample = "+7-476+"
)

var PhoneNumberPatternRestriction = regexp.MustCompile(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`)

// IsValid checks if PhoneNumber of type String is valid
func (t PhoneNumber) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == PhoneNumberZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), PhoneNumberPatternRestriction)

	return valid
}

// String method for easy conversion
func (t PhoneNumber) String() string {
	return string(t)
}

// ToPhoneNumber method for easy conversion with application of restrictions
func ToPhoneNumber(i interface{}) (PhoneNumber, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, PhoneNumberPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type PhoneNumber", s)
	}

	return PhoneNumber(s), nil
}

// MustToPhoneNumber method for easy conversion with application of restrictions. Panics on error.
func MustToPhoneNumber(s interface{}) PhoneNumber {
	v, err := ToPhoneNumber(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * BalanceType12Code Ops
 */

const (
	BalanceType12CodeZero   = ""
	BalanceType12CodeSample = "CLBD"
	BalanceType12CodeXPCD   = "XPCD"
	BalanceType12CodeOPAV   = "OPAV"
	BalanceType12CodeITAV   = "ITAV"
	BalanceType12CodeCLAV   = "CLAV"
	BalanceType12CodeFWAV   = "FWAV"
	BalanceType12CodeCLBD   = "CLBD"
	BalanceType12CodeITBD   = "ITBD"
	BalanceType12CodeOPBD   = "OPBD"
	BalanceType12CodePRCD   = "PRCD"
	BalanceType12CodeINFO   = "INFO"
)

var BalanceType12CodeEnumRestriction = []string{BalanceType12CodeXPCD, BalanceType12CodeOPAV, BalanceType12CodeITAV, BalanceType12CodeCLAV, BalanceType12CodeFWAV, BalanceType12CodeCLBD, BalanceType12CodeITBD, BalanceType12CodeOPBD, BalanceType12CodePRCD, BalanceType12CodeINFO}

// IsValid checks if BalanceType12Code of type String is valid
func (t BalanceType12Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == BalanceType12CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), BalanceType12CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t BalanceType12Code) String() string {
	return string(t)
}

// ToBalanceType12Code method for easy conversion with application of restrictions
func ToBalanceType12Code(i interface{}) (BalanceType12Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, BalanceType12CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type BalanceType12Code", s)
	}

	return BalanceType12Code(s), nil
}

// MustToBalanceType12Code method for easy conversion with application of restrictions. Panics on error.
func MustToBalanceType12Code(s interface{}) BalanceType12Code {
	v, err := ToBalanceType12Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalBankTransactionDomain1Code Ops
 */

const (
	ExternalBankTransactionDomain1CodeZero      = ""
	ExternalBankTransactionDomain1CodeSample    = "OP"
	ExternalBankTransactionDomain1CodeLength    = 0
	ExternalBankTransactionDomain1CodeMinLength = 1
	ExternalBankTransactionDomain1CodeMaxLength = 4
)

// IsValid checks if ExternalBankTransactionDomain1Code of type String is valid
func (t ExternalBankTransactionDomain1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalBankTransactionDomain1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalBankTransactionDomain1CodeLength, ExternalBankTransactionDomain1CodeMinLength, ExternalBankTransactionDomain1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalBankTransactionDomain1Code) String() string {
	return string(t)
}

// ToExternalBankTransactionDomain1Code method for easy conversion with application of restrictions
func ToExternalBankTransactionDomain1Code(i interface{}) (ExternalBankTransactionDomain1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalBankTransactionDomain1CodeLength, ExternalBankTransactionDomain1CodeMinLength, ExternalBankTransactionDomain1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalBankTransactionDomain1Code", s)
	}

	return ExternalBankTransactionDomain1Code(s), nil
}

// MustToExternalBankTransactionDomain1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalBankTransactionDomain1Code(s interface{}) ExternalBankTransactionDomain1Code {
	v, err := ToExternalBankTransactionDomain1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max4Text Ops
 */

const (
	Max4TextZero      = ""
	Max4TextSample    = "ct"
	Max4TextLength    = 0
	Max4TextMinLength = 1
	Max4TextMaxLength = 4
)

// IsValid checks if Max4Text of type String is valid
func (t Max4Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max4TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max4TextLength, Max4TextMinLength, Max4TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max4Text) String() string {
	return string(t)
}

// ToMax4Text method for easy conversion with application of restrictions
func ToMax4Text(i interface{}) (Max4Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max4TextLength, Max4TextMinLength, Max4TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max4Text", s)
	}

	return Max4Text(s), nil
}

// MustToMax4Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax4Text(s interface{}) Max4Text {
	v, err := ToMax4Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * IBAN2007Identifier Ops
 */

const (
	IBAN2007IdentifierZero   = ""
	IBAN2007IdentifierSample = "IS55BY"
)

var IBAN2007IdentifierPatternRestriction = regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)

// IsValid checks if IBAN2007Identifier of type String is valid
func (t IBAN2007Identifier) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == IBAN2007IdentifierZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), IBAN2007IdentifierPatternRestriction)

	return valid
}

// String method for easy conversion
func (t IBAN2007Identifier) String() string {
	return string(t)
}

// ToIBAN2007Identifier method for easy conversion with application of restrictions
func ToIBAN2007Identifier(i interface{}) (IBAN2007Identifier, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, IBAN2007IdentifierPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type IBAN2007Identifier", s)
	}

	return IBAN2007Identifier(s), nil
}

// MustToIBAN2007Identifier method for easy conversion with application of restrictions. Panics on error.
func MustToIBAN2007Identifier(s interface{}) IBAN2007Identifier {
	v, err := ToIBAN2007Identifier(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalBankTransactionFamily1Code Ops
 */

const (
	ExternalBankTransactionFamily1CodeZero      = ""
	ExternalBankTransactionFamily1CodeSample    = "LB"
	ExternalBankTransactionFamily1CodeLength    = 0
	ExternalBankTransactionFamily1CodeMinLength = 1
	ExternalBankTransactionFamily1CodeMaxLength = 4
)

// IsValid checks if ExternalBankTransactionFamily1Code of type String is valid
func (t ExternalBankTransactionFamily1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalBankTransactionFamily1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalBankTransactionFamily1CodeLength, ExternalBankTransactionFamily1CodeMinLength, ExternalBankTransactionFamily1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalBankTransactionFamily1Code) String() string {
	return string(t)
}

// ToExternalBankTransactionFamily1Code method for easy conversion with application of restrictions
func ToExternalBankTransactionFamily1Code(i interface{}) (ExternalBankTransactionFamily1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalBankTransactionFamily1CodeLength, ExternalBankTransactionFamily1CodeMinLength, ExternalBankTransactionFamily1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalBankTransactionFamily1Code", s)
	}

	return ExternalBankTransactionFamily1Code(s), nil
}

// MustToExternalBankTransactionFamily1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalBankTransactionFamily1Code(s interface{}) ExternalBankTransactionFamily1Code {
	v, err := ToExternalBankTransactionFamily1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max35Text Ops
 */

const (
	Max35TextZero      = ""
	Max35TextSample    = "lMZxcxWxtuKOMWJPKL"
	Max35TextLength    = 0
	Max35TextMinLength = 1
	Max35TextMaxLength = 35
)

// IsValid checks if Max35Text of type String is valid
func (t Max35Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max35TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max35TextLength, Max35TextMinLength, Max35TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max35Text) String() string {
	return string(t)
}

// ToMax35Text method for easy conversion with application of restrictions
func ToMax35Text(i interface{}) (Max35Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max35TextLength, Max35TextMinLength, Max35TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max35Text", s)
	}

	return Max35Text(s), nil
}

// MustToMax35Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax35Text(s interface{}) Max35Text {
	v, err := ToMax35Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max15NumericText Ops
 */

const (
	Max15NumericTextZero   = ""
	Max15NumericTextSample = "15885"
)

var Max15NumericTextPatternRestriction = regexp.MustCompile(`[0-9]{1,15}`)

// IsValid checks if Max15NumericText of type String is valid
func (t Max15NumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max15NumericTextZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), Max15NumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Max15NumericText) String() string {
	return string(t)
}

// ToMax15NumericText method for easy conversion with application of restrictions
func ToMax15NumericText(i interface{}) (Max15NumericText, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, Max15NumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Max15NumericText", s)
	}

	return Max15NumericText(s), nil
}

// MustToMax15NumericText method for easy conversion with application of restrictions. Panics on error.
func MustToMax15NumericText(s interface{}) Max15NumericText {
	v, err := ToMax15NumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalBankTransactionSubFamily1Code Ops
 */

const (
	ExternalBankTransactionSubFamily1CodeZero      = ""
	ExternalBankTransactionSubFamily1CodeSample    = "dP"
	ExternalBankTransactionSubFamily1CodeLength    = 0
	ExternalBankTransactionSubFamily1CodeMinLength = 1
	ExternalBankTransactionSubFamily1CodeMaxLength = 4
)

// IsValid checks if ExternalBankTransactionSubFamily1Code of type String is valid
func (t ExternalBankTransactionSubFamily1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalBankTransactionSubFamily1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalBankTransactionSubFamily1CodeLength, ExternalBankTransactionSubFamily1CodeMinLength, ExternalBankTransactionSubFamily1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalBankTransactionSubFamily1Code) String() string {
	return string(t)
}

// ToExternalBankTransactionSubFamily1Code method for easy conversion with application of restrictions
func ToExternalBankTransactionSubFamily1Code(i interface{}) (ExternalBankTransactionSubFamily1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalBankTransactionSubFamily1CodeLength, ExternalBankTransactionSubFamily1CodeMinLength, ExternalBankTransactionSubFamily1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalBankTransactionSubFamily1Code", s)
	}

	return ExternalBankTransactionSubFamily1Code(s), nil
}

// MustToExternalBankTransactionSubFamily1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalBankTransactionSubFamily1Code(s interface{}) ExternalBankTransactionSubFamily1Code {
	v, err := ToExternalBankTransactionSubFamily1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max70Text Ops
 */

const (
	Max70TextZero      = ""
	Max70TextSample    = "agMDpbVTfGYQeuDWBbAsCNeGzZKpmgINTki"
	Max70TextLength    = 0
	Max70TextMinLength = 1
	Max70TextMaxLength = 70
)

// IsValid checks if Max70Text of type String is valid
func (t Max70Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max70TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max70TextLength, Max70TextMinLength, Max70TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max70Text) String() string {
	return string(t)
}

// ToMax70Text method for easy conversion with application of restrictions
func ToMax70Text(i interface{}) (Max70Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max70TextLength, Max70TextMinLength, Max70TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max70Text", s)
	}

	return Max70Text(s), nil
}

// MustToMax70Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax70Text(s interface{}) Max70Text {
	v, err := ToMax70Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalAccountIdentification1Code Ops
 */

const (
	ExternalAccountIdentification1CodeZero      = ""
	ExternalAccountIdentification1CodeSample    = "an"
	ExternalAccountIdentification1CodeLength    = 0
	ExternalAccountIdentification1CodeMinLength = 1
	ExternalAccountIdentification1CodeMaxLength = 4
)

// IsValid checks if ExternalAccountIdentification1Code of type String is valid
func (t ExternalAccountIdentification1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalAccountIdentification1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalAccountIdentification1CodeLength, ExternalAccountIdentification1CodeMinLength, ExternalAccountIdentification1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalAccountIdentification1Code) String() string {
	return string(t)
}

// ToExternalAccountIdentification1Code method for easy conversion with application of restrictions
func ToExternalAccountIdentification1Code(i interface{}) (ExternalAccountIdentification1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalAccountIdentification1CodeLength, ExternalAccountIdentification1CodeMinLength, ExternalAccountIdentification1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalAccountIdentification1Code", s)
	}

	return ExternalAccountIdentification1Code(s), nil
}

// MustToExternalAccountIdentification1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalAccountIdentification1Code(s interface{}) ExternalAccountIdentification1Code {
	v, err := ToExternalAccountIdentification1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max15PlusSignedNumericText Ops
 */

const (
	Max15PlusSignedNumericTextZero   = ""
	Max15PlusSignedNumericTextSample = "+4"
)

var Max15PlusSignedNumericTextPatternRestriction = regexp.MustCompile(`[+]{0,1}[0-9]{1,15}`)

// IsValid checks if Max15PlusSignedNumericText of type String is valid
func (t Max15PlusSignedNumericText) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max15PlusSignedNumericTextZero {
		return valid
	}
	valid = valid && isPatternRestrictionValid(t.String(), Max15PlusSignedNumericTextPatternRestriction)

	return valid
}

// String method for easy conversion
func (t Max15PlusSignedNumericText) String() string {
	return string(t)
}

// ToMax15PlusSignedNumericText method for easy conversion with application of restrictions
func ToMax15PlusSignedNumericText(i interface{}) (Max15PlusSignedNumericText, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isPatternRestrictionValid(s, Max15PlusSignedNumericTextPatternRestriction) {
		return "", fmt.Errorf("cannot satisfy pattern restriction for %s of type Max15PlusSignedNumericText", s)
	}

	return Max15PlusSignedNumericText(s), nil
}

// MustToMax15PlusSignedNumericText method for easy conversion with application of restrictions. Panics on error.
func MustToMax15PlusSignedNumericText(s interface{}) Max15PlusSignedNumericText {
	v, err := ToMax15PlusSignedNumericText(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * RemittanceLocationMethod2Code Ops
 */

const (
	RemittanceLocationMethod2CodeZero   = ""
	RemittanceLocationMethod2CodeSample = "EMAL"
	RemittanceLocationMethod2CodeFAXI   = "FAXI"
	RemittanceLocationMethod2CodeEDIC   = "EDIC"
	RemittanceLocationMethod2CodeURID   = "URID"
	RemittanceLocationMethod2CodeEMAL   = "EMAL"
	RemittanceLocationMethod2CodePOST   = "POST"
	RemittanceLocationMethod2CodeSMSM   = "SMSM"
)

var RemittanceLocationMethod2CodeEnumRestriction = []string{RemittanceLocationMethod2CodeFAXI, RemittanceLocationMethod2CodeEDIC, RemittanceLocationMethod2CodeURID, RemittanceLocationMethod2CodeEMAL, RemittanceLocationMethod2CodePOST, RemittanceLocationMethod2CodeSMSM}

// IsValid checks if RemittanceLocationMethod2Code of type String is valid
func (t RemittanceLocationMethod2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == RemittanceLocationMethod2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), RemittanceLocationMethod2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RemittanceLocationMethod2Code) String() string {
	return string(t)
}

// ToRemittanceLocationMethod2Code method for easy conversion with application of restrictions
func ToRemittanceLocationMethod2Code(i interface{}) (RemittanceLocationMethod2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, RemittanceLocationMethod2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type RemittanceLocationMethod2Code", s)
	}

	return RemittanceLocationMethod2Code(s), nil
}

// MustToRemittanceLocationMethod2Code method for easy conversion with application of restrictions. Panics on error.
func MustToRemittanceLocationMethod2Code(s interface{}) RemittanceLocationMethod2Code {
	v, err := ToRemittanceLocationMethod2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}
