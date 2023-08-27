// Package common
// Do not Edit. This stuff it's been automatically generated.
package common

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/xsdt"
	"regexp"
	"time"
)

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

/*
 * RegulatoryReportingType1Code Ops
 */

const (
	RegulatoryReportingType1CodeZero   = ""
	RegulatoryReportingType1CodeSample = "DEBT"
	RegulatoryReportingType1CodeCRED   = "CRED"
	RegulatoryReportingType1CodeDEBT   = "DEBT"
	RegulatoryReportingType1CodeBOTH   = "BOTH"
)

var RegulatoryReportingType1CodeEnumRestriction = []string{RegulatoryReportingType1CodeCRED, RegulatoryReportingType1CodeDEBT, RegulatoryReportingType1CodeBOTH}

// IsValid checks if RegulatoryReportingType1Code of type String is valid
func (t RegulatoryReportingType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == RegulatoryReportingType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), RegulatoryReportingType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t RegulatoryReportingType1Code) String() string {
	return string(t)
}

// ToRegulatoryReportingType1Code method for easy conversion with application of restrictions
func ToRegulatoryReportingType1Code(i interface{}) (RegulatoryReportingType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, RegulatoryReportingType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type RegulatoryReportingType1Code", s)
	}

	return RegulatoryReportingType1Code(s), nil
}

// MustToRegulatoryReportingType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToRegulatoryReportingType1Code(s interface{}) RegulatoryReportingType1Code {
	v, err := ToRegulatoryReportingType1Code(s)
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
	Max4TextSample    = "oJ"
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
 * PhoneNumber Ops
 */

const (
	PhoneNumberZero   = ""
	PhoneNumberSample = "+07-97689)3(-2"
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
 * ChequeDelivery1Code Ops
 */

const (
	ChequeDelivery1CodeZero   = ""
	ChequeDelivery1CodeSample = "PUDB"
	ChequeDelivery1CodeMLDB   = "MLDB"
	ChequeDelivery1CodeMLCD   = "MLCD"
	ChequeDelivery1CodeMLFA   = "MLFA"
	ChequeDelivery1CodeCRDB   = "CRDB"
	ChequeDelivery1CodeCRCD   = "CRCD"
	ChequeDelivery1CodeCRFA   = "CRFA"
	ChequeDelivery1CodePUDB   = "PUDB"
	ChequeDelivery1CodePUCD   = "PUCD"
	ChequeDelivery1CodePUFA   = "PUFA"
	ChequeDelivery1CodeRGDB   = "RGDB"
	ChequeDelivery1CodeRGCD   = "RGCD"
	ChequeDelivery1CodeRGFA   = "RGFA"
)

var ChequeDelivery1CodeEnumRestriction = []string{ChequeDelivery1CodeMLDB, ChequeDelivery1CodeMLCD, ChequeDelivery1CodeMLFA, ChequeDelivery1CodeCRDB, ChequeDelivery1CodeCRCD, ChequeDelivery1CodeCRFA, ChequeDelivery1CodePUDB, ChequeDelivery1CodePUCD, ChequeDelivery1CodePUFA, ChequeDelivery1CodeRGDB, ChequeDelivery1CodeRGCD, ChequeDelivery1CodeRGFA}

// IsValid checks if ChequeDelivery1Code of type String is valid
func (t ChequeDelivery1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ChequeDelivery1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ChequeDelivery1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeDelivery1Code) String() string {
	return string(t)
}

// ToChequeDelivery1Code method for easy conversion with application of restrictions
func ToChequeDelivery1Code(i interface{}) (ChequeDelivery1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ChequeDelivery1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChequeDelivery1Code", s)
	}

	return ChequeDelivery1Code(s), nil
}

// MustToChequeDelivery1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChequeDelivery1Code(s interface{}) ChequeDelivery1Code {
	v, err := ToChequeDelivery1Code(s)
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
	BICIdentifierSample = "KFMIGT9ATEZ"
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
 * PaymentMethod3Code Ops
 */

const (
	PaymentMethod3CodeZero   = ""
	PaymentMethod3CodeSample = "TRF"
	PaymentMethod3CodeCHK    = "CHK"
	PaymentMethod3CodeTRF    = "TRF"
	PaymentMethod3CodeTRA    = "TRA"
)

var PaymentMethod3CodeEnumRestriction = []string{PaymentMethod3CodeCHK, PaymentMethod3CodeTRF, PaymentMethod3CodeTRA}

// IsValid checks if PaymentMethod3Code of type String is valid
func (t PaymentMethod3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == PaymentMethod3CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), PaymentMethod3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod3Code) String() string {
	return string(t)
}

// ToPaymentMethod3Code method for easy conversion with application of restrictions
func ToPaymentMethod3Code(i interface{}) (PaymentMethod3Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, PaymentMethod3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type PaymentMethod3Code", s)
	}

	return PaymentMethod3Code(s), nil
}

// MustToPaymentMethod3Code method for easy conversion with application of restrictions. Panics on error.
func MustToPaymentMethod3Code(s interface{}) PaymentMethod3Code {
	v, err := ToPaymentMethod3Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExchangeRateType1Code Ops
 */

const (
	ExchangeRateType1CodeZero   = ""
	ExchangeRateType1CodeSample = "SALE"
	ExchangeRateType1CodeSPOT   = "SPOT"
	ExchangeRateType1CodeSALE   = "SALE"
	ExchangeRateType1CodeAGRD   = "AGRD"
)

var ExchangeRateType1CodeEnumRestriction = []string{ExchangeRateType1CodeSPOT, ExchangeRateType1CodeSALE, ExchangeRateType1CodeAGRD}

// IsValid checks if ExchangeRateType1Code of type String is valid
func (t ExchangeRateType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExchangeRateType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ExchangeRateType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ExchangeRateType1Code) String() string {
	return string(t)
}

// ToExchangeRateType1Code method for easy conversion with application of restrictions
func ToExchangeRateType1Code(i interface{}) (ExchangeRateType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ExchangeRateType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ExchangeRateType1Code", s)
	}

	return ExchangeRateType1Code(s), nil
}

// MustToExchangeRateType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExchangeRateType1Code(s interface{}) ExchangeRateType1Code {
	v, err := ToExchangeRateType1Code(s)
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
	Max140TextSample    = "nNPGsiuzytMOJPatwtPilfsfykSBGplhxtxVSGpqaJaBRgAvzLXqzRrrUIYvaIujDpHYjx"
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
 * Max2048Text Ops
 */

const (
	Max2048TextZero      = ""
	Max2048TextSample    = "eUBrVfdwUzWHRihFDPRHBMuEWmaNvhHLITWJcJKzJsRPeJbpTqWPUlWAcidRlufmYWddEktBDrWMYZTPdmDAVCuvQutLDKFBGRxnyDNYuRKWsmZnOYgSLPSsuMMmvYmaFEjLgSvndTeranxGMNCcCdlEBYqxqrgAieJuFVZUsgPmweLRmsdTIwtLUDqoBSNtPRViLXLCrSAZCjxTCaZnfHRhEbbvlLkCbvOlRrOWJFNwNQBqyozKVGeninyrTgqubNiIzWcEZsdyUJTDAODnRKcTqZThaTJmRVAClOEaCEmEUmEtNTDVBbvMBtQBeDBVyDFfPWGaZQntPiyJbzjrRNLRzGpzDszPvweGxRQcoxpAQxuFkaowCKoRccTyTJHeKKokJNwUaqKCBRxMSdypnuVQlXTxlLUfRUssUCbZClckQhhyMmVTDwokzVJUlvROmQBJTVsrIgXXlvECpbDIWegrUuUYynIaqJjgVXQtZptcLZIVmypNpLbbaxHYDmKrYnEbqbvWLjZZAHQHWIGRYPwqfeXSWmRncQZpTAzceEsIpdYbrhfkVRdPhRXJzEKTBNCvokhNhSeFRvuKeklvsHwabufMjYxuvCIeoMjxivfYHKXcwedxkRvEookKzpacqgBiXejXjTJPhMQsmGzinFnyUVhVOydnINYczGVyZhBRwWAXjuufUoPJlNOFUvANpjrlwiuIbjJbgkIxMkImqLfMdmksNxvjgyEJAjrjcQDUlVTuFKDLjgBjrpEBlRzIZVSUAgxlZPBVcodWJbYmpBEprbxsRnSFqRIjOBrNaxVzTtXlnbdWOKzddfZvtzFRwXkWBBNMmMfCraUahSxpTkppSxjPLhPnxeSCjulnkeSRvUtXaQHrOqxrwbCXRiWGXOyolORPasKWOdANVLTUjsQCqRuNsAzgVRDgWTxeohgOVmZWEReHPcpuWoyejNDNtfSJWCidgAfEcoSKFvxTVCZoKMHuKFgkhaujRMxfcNGDjkbfeENRdnZuYVnOFfaK"
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
	ActiveOrHistoricCurrencyCodeSample = "TGW"
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
 * IBAN2007Identifier Ops
 */

const (
	IBAN2007IdentifierZero   = ""
	IBAN2007IdentifierSample = "ND77RgLR5GF"
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
 * ChequeType2Code Ops
 */

const (
	ChequeType2CodeZero   = ""
	ChequeType2CodeSample = "BCHQ"
	ChequeType2CodeCCHQ   = "CCHQ"
	ChequeType2CodeCCCH   = "CCCH"
	ChequeType2CodeBCHQ   = "BCHQ"
	ChequeType2CodeDRFT   = "DRFT"
	ChequeType2CodeELDR   = "ELDR"
)

var ChequeType2CodeEnumRestriction = []string{ChequeType2CodeCCHQ, ChequeType2CodeCCCH, ChequeType2CodeBCHQ, ChequeType2CodeDRFT, ChequeType2CodeELDR}

// IsValid checks if ChequeType2Code of type String is valid
func (t ChequeType2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ChequeType2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ChequeType2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChequeType2Code) String() string {
	return string(t)
}

// ToChequeType2Code method for easy conversion with application of restrictions
func ToChequeType2Code(i interface{}) (ChequeType2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ChequeType2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChequeType2Code", s)
	}

	return ChequeType2Code(s), nil
}

// MustToChequeType2Code method for easy conversion with application of restrictions. Panics on error.
func MustToChequeType2Code(s interface{}) ChequeType2Code {
	v, err := ToChequeType2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max10Text Ops
 */

const (
	Max10TextZero      = ""
	Max10TextSample    = "gXcxf"
	Max10TextLength    = 0
	Max10TextMinLength = 1
	Max10TextMaxLength = 10
)

// IsValid checks if Max10Text of type String is valid
func (t Max10Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max10TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max10TextLength, Max10TextMinLength, Max10TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max10Text) String() string {
	return string(t)
}

// ToMax10Text method for easy conversion with application of restrictions
func ToMax10Text(i interface{}) (Max10Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max10TextLength, Max10TextMinLength, Max10TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max10Text", s)
	}

	return Max10Text(s), nil
}

// MustToMax10Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax10Text(s interface{}) Max10Text {
	v, err := ToMax10Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * ExternalCategoryPurpose1Code Ops
 */

const (
	ExternalCategoryPurpose1CodeZero      = ""
	ExternalCategoryPurpose1CodeSample    = "cX"
	ExternalCategoryPurpose1CodeLength    = 0
	ExternalCategoryPurpose1CodeMinLength = 1
	ExternalCategoryPurpose1CodeMaxLength = 4
)

// IsValid checks if ExternalCategoryPurpose1Code of type String is valid
func (t ExternalCategoryPurpose1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalCategoryPurpose1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCategoryPurpose1CodeLength, ExternalCategoryPurpose1CodeMinLength, ExternalCategoryPurpose1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCategoryPurpose1Code) String() string {
	return string(t)
}

// ToExternalCategoryPurpose1Code method for easy conversion with application of restrictions
func ToExternalCategoryPurpose1Code(i interface{}) (ExternalCategoryPurpose1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalCategoryPurpose1CodeLength, ExternalCategoryPurpose1CodeMinLength, ExternalCategoryPurpose1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCategoryPurpose1Code", s)
	}

	return ExternalCategoryPurpose1Code(s), nil
}

// MustToExternalCategoryPurpose1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCategoryPurpose1Code(s interface{}) ExternalCategoryPurpose1Code {
	v, err := ToExternalCategoryPurpose1Code(s)
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
	Max35TextSample    = "VutzuIkhdUMOniOqWY"
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
 * ExternalServiceLevel1Code Ops
 */

const (
	ExternalServiceLevel1CodeZero      = ""
	ExternalServiceLevel1CodeSample    = "IJ"
	ExternalServiceLevel1CodeLength    = 0
	ExternalServiceLevel1CodeMinLength = 1
	ExternalServiceLevel1CodeMaxLength = 4
)

// IsValid checks if ExternalServiceLevel1Code of type String is valid
func (t ExternalServiceLevel1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalServiceLevel1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalServiceLevel1CodeLength, ExternalServiceLevel1CodeMinLength, ExternalServiceLevel1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalServiceLevel1Code) String() string {
	return string(t)
}

// ToExternalServiceLevel1Code method for easy conversion with application of restrictions
func ToExternalServiceLevel1Code(i interface{}) (ExternalServiceLevel1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalServiceLevel1CodeLength, ExternalServiceLevel1CodeMinLength, ExternalServiceLevel1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalServiceLevel1Code", s)
	}

	return ExternalServiceLevel1Code(s), nil
}

// MustToExternalServiceLevel1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalServiceLevel1Code(s interface{}) ExternalServiceLevel1Code {
	v, err := ToExternalServiceLevel1Code(s)
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
	Max70TextSample    = "MqLchMkCJZHSenoYIFnHZyTzEfMSdhUvtMB"
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
	ExternalAccountIdentification1CodeSample    = "lf"
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
 * Instruction3Code Ops
 */

const (
	Instruction3CodeZero   = ""
	Instruction3CodeSample = "PHOB"
	Instruction3CodeCHQB   = "CHQB"
	Instruction3CodeHOLD   = "HOLD"
	Instruction3CodePHOB   = "PHOB"
	Instruction3CodeTELB   = "TELB"
)

var Instruction3CodeEnumRestriction = []string{Instruction3CodeCHQB, Instruction3CodeHOLD, Instruction3CodePHOB, Instruction3CodeTELB}

// IsValid checks if Instruction3Code of type String is valid
func (t Instruction3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Instruction3CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), Instruction3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Instruction3Code) String() string {
	return string(t)
}

// ToInstruction3Code method for easy conversion with application of restrictions
func ToInstruction3Code(i interface{}) (Instruction3Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, Instruction3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Instruction3Code", s)
	}

	return Instruction3Code(s), nil
}

// MustToInstruction3Code method for easy conversion with application of restrictions. Panics on error.
func MustToInstruction3Code(s interface{}) Instruction3Code {
	v, err := ToInstruction3Code(s)
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
 * Priority2Code Ops
 */

const (
	Priority2CodeZero   = ""
	Priority2CodeSample = "NORM"
	Priority2CodeHIGH   = "HIGH"
	Priority2CodeNORM   = "NORM"
)

var Priority2CodeEnumRestriction = []string{Priority2CodeHIGH, Priority2CodeNORM}

// IsValid checks if Priority2Code of type String is valid
func (t Priority2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Priority2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), Priority2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Priority2Code) String() string {
	return string(t)
}

// ToPriority2Code method for easy conversion with application of restrictions
func ToPriority2Code(i interface{}) (Priority2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, Priority2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Priority2Code", s)
	}

	return Priority2Code(s), nil
}

// MustToPriority2Code method for easy conversion with application of restrictions. Panics on error.
func MustToPriority2Code(s interface{}) Priority2Code {
	v, err := ToPriority2Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Authorisation1Code Ops
 */

const (
	Authorisation1CodeZero   = ""
	Authorisation1CodeSample = "FSUM"
	Authorisation1CodeAUTH   = "AUTH"
	Authorisation1CodeFDET   = "FDET"
	Authorisation1CodeFSUM   = "FSUM"
	Authorisation1CodeILEV   = "ILEV"
)

var Authorisation1CodeEnumRestriction = []string{Authorisation1CodeAUTH, Authorisation1CodeFDET, Authorisation1CodeFSUM, Authorisation1CodeILEV}

// IsValid checks if Authorisation1Code of type String is valid
func (t Authorisation1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Authorisation1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), Authorisation1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Authorisation1Code) String() string {
	return string(t)
}

// ToAuthorisation1Code method for easy conversion with application of restrictions
func ToAuthorisation1Code(i interface{}) (Authorisation1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, Authorisation1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Authorisation1Code", s)
	}

	return Authorisation1Code(s), nil
}

// MustToAuthorisation1Code method for easy conversion with application of restrictions. Panics on error.
func MustToAuthorisation1Code(s interface{}) Authorisation1Code {
	v, err := ToAuthorisation1Code(s)
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
	ExternalOrganisationIdentification1CodeSample    = "Uo"
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
 * ExternalClearingSystemIdentification1Code Ops
 */

const (
	ExternalClearingSystemIdentification1CodeZero      = ""
	ExternalClearingSystemIdentification1CodeSample    = "ZYL"
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
 * ExternalLocalInstrument1Code Ops
 */

const (
	ExternalLocalInstrument1CodeZero      = ""
	ExternalLocalInstrument1CodeSample    = "EdqkoZCuTuRpLKzzuu"
	ExternalLocalInstrument1CodeLength    = 0
	ExternalLocalInstrument1CodeMinLength = 1
	ExternalLocalInstrument1CodeMaxLength = 35
)

// IsValid checks if ExternalLocalInstrument1Code of type String is valid
func (t ExternalLocalInstrument1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalLocalInstrument1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalLocalInstrument1CodeLength, ExternalLocalInstrument1CodeMinLength, ExternalLocalInstrument1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalLocalInstrument1Code) String() string {
	return string(t)
}

// ToExternalLocalInstrument1Code method for easy conversion with application of restrictions
func ToExternalLocalInstrument1Code(i interface{}) (ExternalLocalInstrument1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalLocalInstrument1CodeLength, ExternalLocalInstrument1CodeMinLength, ExternalLocalInstrument1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalLocalInstrument1Code", s)
	}

	return ExternalLocalInstrument1Code(s), nil
}

// MustToExternalLocalInstrument1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalLocalInstrument1Code(s interface{}) ExternalLocalInstrument1Code {
	v, err := ToExternalLocalInstrument1Code(s)
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
	CountryCodeSample = "HW"
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
 * ExternalFinancialInstitutionIdentification1Code Ops
 */

const (
	ExternalFinancialInstitutionIdentification1CodeZero      = ""
	ExternalFinancialInstitutionIdentification1CodeSample    = "Fj"
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
 * Max34Text Ops
 */

const (
	Max34TextZero      = ""
	Max34TextSample    = "TKIOTjSnLTFioYayJ"
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
 * Max15NumericText Ops
 */

const (
	Max15NumericTextZero   = ""
	Max15NumericTextSample = "120136766"
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
 * AnyBICIdentifier Ops
 */

const (
	AnyBICIdentifierZero   = ""
	AnyBICIdentifierSample = "WDDNSX3N"
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
 * Max128Text Ops
 */

const (
	Max128TextZero      = ""
	Max128TextSample    = "wBotODVVVIrFCqrhOGVehnNZfARfsBVGBRaHjfFcQtXovQyXbjAyWbspgheqXXum"
	Max128TextLength    = 0
	Max128TextMinLength = 1
	Max128TextMaxLength = 128
)

// IsValid checks if Max128Text of type String is valid
func (t Max128Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max128TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max128TextLength, Max128TextMinLength, Max128TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max128Text) String() string {
	return string(t)
}

// ToMax128Text method for easy conversion with application of restrictions
func ToMax128Text(i interface{}) (Max128Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max128TextLength, Max128TextMinLength, Max128TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max128Text", s)
	}

	return Max128Text(s), nil
}

// MustToMax128Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax128Text(s interface{}) Max128Text {
	v, err := ToMax128Text(s)
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
	Max16TextSample    = "wlaMkvqg"
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
 * ExternalPersonIdentification1Code Ops
 */

const (
	ExternalPersonIdentification1CodeZero      = ""
	ExternalPersonIdentification1CodeSample    = "eV"
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
 * ExternalPurpose1Code Ops
 */

const (
	ExternalPurpose1CodeZero      = ""
	ExternalPurpose1CodeSample    = "Ns"
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
