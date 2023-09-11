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
 * ExternalCategoryPurpose1Code Ops
 */

const (
	ExternalCategoryPurpose1CodeZero      = ""
	ExternalCategoryPurpose1CodeSample    = "oJ"
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
 * ExternalAccountIdentification1Code Ops
 */

const (
	ExternalAccountIdentification1CodeZero      = ""
	ExternalAccountIdentification1CodeSample    = "nN"
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
 * ClearingChannel2Code Ops
 */

const (
	ClearingChannel2CodeZero   = ""
	ClearingChannel2CodeSample = "MPNS"
	ClearingChannel2CodeRTGS   = "RTGS"
	ClearingChannel2CodeRTNS   = "RTNS"
	ClearingChannel2CodeMPNS   = "MPNS"
	ClearingChannel2CodeBOOK   = "BOOK"
)

var ClearingChannel2CodeEnumRestriction = []string{ClearingChannel2CodeRTGS, ClearingChannel2CodeRTNS, ClearingChannel2CodeMPNS, ClearingChannel2CodeBOOK}

// IsValid checks if ClearingChannel2Code of type String is valid
func (t ClearingChannel2Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ClearingChannel2CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), ClearingChannel2CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ClearingChannel2Code) String() string {
	return string(t)
}

// ToClearingChannel2Code method for easy conversion with application of restrictions
func ToClearingChannel2Code(i interface{}) (ClearingChannel2Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, ClearingChannel2CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ClearingChannel2Code", s)
	}

	return ClearingChannel2Code(s), nil
}

// MustToClearingChannel2Code method for easy conversion with application of restrictions. Panics on error.
func MustToClearingChannel2Code(s interface{}) ClearingChannel2Code {
	v, err := ToClearingChannel2Code(s)
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
	ExternalOrganisationIdentification1CodeSample    = "PG"
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
 * Max105Text Ops
 */

const (
	Max105TextZero      = ""
	Max105TextSample    = "siuzytMOJPatwtPilfsfykSBGplhxtxVSGpqaJaBRgAvzLXqzRrrU"
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
 * PaymentMethod4Code Ops
 */

const (
	PaymentMethod4CodeZero   = ""
	PaymentMethod4CodeSample = "DD"
	PaymentMethod4CodeCHK    = "CHK"
	PaymentMethod4CodeTRF    = "TRF"
	PaymentMethod4CodeDD     = "DD"
	PaymentMethod4CodeTRA    = "TRA"
)

var PaymentMethod4CodeEnumRestriction = []string{PaymentMethod4CodeCHK, PaymentMethod4CodeTRF, PaymentMethod4CodeDD, PaymentMethod4CodeTRA}

// IsValid checks if PaymentMethod4Code of type String is valid
func (t PaymentMethod4Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == PaymentMethod4CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), PaymentMethod4CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t PaymentMethod4Code) String() string {
	return string(t)
}

// ToPaymentMethod4Code method for easy conversion with application of restrictions
func ToPaymentMethod4Code(i interface{}) (PaymentMethod4Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, PaymentMethod4CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type PaymentMethod4Code", s)
	}

	return PaymentMethod4Code(s), nil
}

// MustToPaymentMethod4Code method for easy conversion with application of restrictions. Panics on error.
func MustToPaymentMethod4Code(s interface{}) PaymentMethod4Code {
	v, err := ToPaymentMethod4Code(s)
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
 * Max35Text Ops
 */

const (
	Max35TextZero      = ""
	Max35TextSample    = "IYvaIujDpHYjxeUBrV"
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
 * ActiveOrHistoricCurrencyCode Ops
 */

const (
	ActiveOrHistoricCurrencyCodeZero   = ""
	ActiveOrHistoricCurrencyCodeSample = "OKQ"
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
 * IBAN2007Identifier Ops
 */

const (
	IBAN2007IdentifierZero   = ""
	IBAN2007IdentifierSample = "HZ09frdBS9f2o"
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
 * PhoneNumber Ops
 */

const (
	PhoneNumberZero   = ""
	PhoneNumberSample = "+331-(119"
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
 * ExternalLocalInstrument1Code Ops
 */

const (
	ExternalLocalInstrument1CodeZero      = ""
	ExternalLocalInstrument1CodeSample    = "fdwUzWHRihFDPRHBMu"
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
 * SequenceType1Code Ops
 */

const (
	SequenceType1CodeZero   = ""
	SequenceType1CodeSample = "FNAL"
	SequenceType1CodeFRST   = "FRST"
	SequenceType1CodeRCUR   = "RCUR"
	SequenceType1CodeFNAL   = "FNAL"
	SequenceType1CodeOOFF   = "OOFF"
)

var SequenceType1CodeEnumRestriction = []string{SequenceType1CodeFRST, SequenceType1CodeRCUR, SequenceType1CodeFNAL, SequenceType1CodeOOFF}

// IsValid checks if SequenceType1Code of type String is valid
func (t SequenceType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == SequenceType1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), SequenceType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SequenceType1Code) String() string {
	return string(t)
}

// ToSequenceType1Code method for easy conversion with application of restrictions
func ToSequenceType1Code(i interface{}) (SequenceType1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, SequenceType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type SequenceType1Code", s)
	}

	return SequenceType1Code(s), nil
}

// MustToSequenceType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToSequenceType1Code(s interface{}) SequenceType1Code {
	v, err := ToSequenceType1Code(s)
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
 * ExternalPersonIdentification1Code Ops
 */

const (
	ExternalPersonIdentification1CodeZero      = ""
	ExternalPersonIdentification1CodeSample    = "EW"
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
 * Max2048Text Ops
 */

const (
	Max2048TextZero      = ""
	Max2048TextSample    = "maNvhHLITWJcJKzJsRPeJbpTqWPUlWAcidRlufmYWddEktBDrWMYZTPdmDAVCuvQutLDKFBGRxnyDNYuRKWsmZnOYgSLPSsuMMmvYmaFEjLgSvndTeranxGMNCcCdlEBYqxqrgAieJuFVZUsgPmweLRmsdTIwtLUDqoBSNtPRViLXLCrSAZCjxTCaZnfHRhEbbvlLkCbvOlRrOWJFNwNQBqyozKVGeninyrTgqubNiIzWcEZsdyUJTDAODnRKcTqZThaTJmRVAClOEaCEmEUmEtNTDVBbvMBtQBeDBVyDFfPWGaZQntPiyJbzjrRNLRzGpzDszPvweGxRQcoxpAQxuFkaowCKoRccTyTJHeKKokJNwUaqKCBRxMSdypnuVQlXTxlLUfRUssUCbZClckQhhyMmVTDwokzVJUlvROmQBJTVsrIgXXlvECpbDIWegrUuUYynIaqJjgVXQtZptcLZIVmypNpLbbaxHYDmKrYnEbqbvWLjZZAHQHWIGRYPwqfeXSWmRncQZpTAzceEsIpdYbrhfkVRdPhRXJzEKTBNCvokhNhSeFRvuKeklvsHwabufMjYxuvCIeoMjxivfYHKXcwedxkRvEookKzpacqgBiXejXjTJPhMQsmGzinFnyUVhVOydnINYczGVyZhBRwWAXjuufUoPJlNOFUvANpjrlwiuIbjJbgkIxMkImqLfMdmksNxvjgyEJAjrjcQDUlVTuFKDLjgBjrpEBlRzIZVSUAgxlZPBVcodWJbYmpBEprbxsRnSFqRIjOBrNaxVzTtXlnbdWOKzddfZvtzFRwXkWBBNMmMfCraUahSxpTkppSxjPLhPnxeSCjulnkeSRvUtXaQHrOqxrwbCXRiWGXOyolORPasKWOdANVLTUjsQCqRuNsAzgVRDgWTxeohgOVmZWEReHPcpuWoyejNDNtfSJWCidgAfEcoSKFvxTVCZoKMHuKFgkhaujRMxfcNGDjkbfeENRdnZuYVnOFfaKgXcxfcXVutzuIkhdUMOniOqWY"
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
 * Max1025Text Ops
 */

const (
	Max1025TextZero      = ""
	Max1025TextSample    = "IJMqLchMkCJZHSenoYIFnHZyTzEfMSdhUvtMBlfUoZYLEdqkoZCuTuRpLKzzuuFjTKIOTjSnLTFioYayJwBotODVVVIrFCqrhOGVehnNZfARfsBVGBRaHjfFcQtXovQyXbjAyWbspgheqXXumwlaMkvqgeVNsLyylTpEsGuymNstQzxPbmXGYPZRoCcAjYwjicELciiQYDuPezahlBGYsXsQyazOWcrSbAuTepBSUGepQDYBzMBKXUbWscOLULiIiQtPiPomlXcvuojfHBHOQiGrmoSuelNipGLXawNpUCRNoVMTKOGgxNQPpRiWmNuFOPctLBlMZxcxWxtuKOMWJPKLdPagMDpbVTfGYQeuDWBbAsCNeGzZKpmgINTkianyMbMixyxcPQjMCSHCDrkbuyVTPEZjtFPuvTZsHZzzaAnhQFNuiUgbXhycXDtAIhVqjQdfuwMFhuNIjQYGTeIzlHVLBSMEFGHRBYTYpySblEWkdaiqYPvySaTkOBngwczeR"
	Max1025TextLength    = 0
	Max1025TextMinLength = 1
	Max1025TextMaxLength = 1025
)

// IsValid checks if Max1025Text of type String is valid
func (t Max1025Text) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Max1025TextZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), Max1025TextLength, Max1025TextMinLength, Max1025TextMaxLength)

	return valid
}

// String method for easy conversion
func (t Max1025Text) String() string {
	return string(t)
}

// ToMax1025Text method for easy conversion with application of restrictions
func ToMax1025Text(i interface{}) (Max1025Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max1025TextLength, Max1025TextMinLength, Max1025TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max1025Text", s)
	}

	return Max1025Text(s), nil
}

// MustToMax1025Text method for easy conversion with application of restrictions. Panics on error.
func MustToMax1025Text(s interface{}) Max1025Text {
	v, err := ToMax1025Text(s)
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
 * Max4Text Ops
 */

const (
	Max4TextZero      = ""
	Max4TextSample    = "vN"
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
 * Max16Text Ops
 */

const (
	Max16TextZero      = ""
	Max16TextSample    = "onCngObQ"
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
 * ExternalClearingSystemIdentification1Code Ops
 */

const (
	ExternalClearingSystemIdentification1CodeZero      = ""
	ExternalClearingSystemIdentification1CodeSample    = "iLP"
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
 * ExternalCashClearingSystem1Code Ops
 */

const (
	ExternalCashClearingSystem1CodeZero      = ""
	ExternalCashClearingSystem1CodeSample    = "JI"
	ExternalCashClearingSystem1CodeLength    = 0
	ExternalCashClearingSystem1CodeMinLength = 1
	ExternalCashClearingSystem1CodeMaxLength = 3
)

// IsValid checks if ExternalCashClearingSystem1Code of type String is valid
func (t ExternalCashClearingSystem1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalCashClearingSystem1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalCashClearingSystem1CodeLength, ExternalCashClearingSystem1CodeMinLength, ExternalCashClearingSystem1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalCashClearingSystem1Code) String() string {
	return string(t)
}

// ToExternalCashClearingSystem1Code method for easy conversion with application of restrictions
func ToExternalCashClearingSystem1Code(i interface{}) (ExternalCashClearingSystem1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalCashClearingSystem1CodeLength, ExternalCashClearingSystem1CodeMinLength, ExternalCashClearingSystem1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalCashClearingSystem1Code", s)
	}

	return ExternalCashClearingSystem1Code(s), nil
}

// MustToExternalCashClearingSystem1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalCashClearingSystem1Code(s interface{}) ExternalCashClearingSystem1Code {
	v, err := ToExternalCashClearingSystem1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * TransactionGroupStatus3Code Ops
 */

const (
	TransactionGroupStatus3CodeZero   = ""
	TransactionGroupStatus3CodeSample = "PDNG"
	TransactionGroupStatus3CodeACTC   = "ACTC"
	TransactionGroupStatus3CodeRCVD   = "RCVD"
	TransactionGroupStatus3CodePART   = "PART"
	TransactionGroupStatus3CodeRJCT   = "RJCT"
	TransactionGroupStatus3CodePDNG   = "PDNG"
	TransactionGroupStatus3CodeACCP   = "ACCP"
	TransactionGroupStatus3CodeACSP   = "ACSP"
	TransactionGroupStatus3CodeACSC   = "ACSC"
	TransactionGroupStatus3CodeACWC   = "ACWC"
)

var TransactionGroupStatus3CodeEnumRestriction = []string{TransactionGroupStatus3CodeACTC, TransactionGroupStatus3CodeRCVD, TransactionGroupStatus3CodePART, TransactionGroupStatus3CodeRJCT, TransactionGroupStatus3CodePDNG, TransactionGroupStatus3CodeACCP, TransactionGroupStatus3CodeACSP, TransactionGroupStatus3CodeACSC, TransactionGroupStatus3CodeACWC}

// IsValid checks if TransactionGroupStatus3Code of type String is valid
func (t TransactionGroupStatus3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == TransactionGroupStatus3CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), TransactionGroupStatus3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TransactionGroupStatus3Code) String() string {
	return string(t)
}

// ToTransactionGroupStatus3Code method for easy conversion with application of restrictions
func ToTransactionGroupStatus3Code(i interface{}) (TransactionGroupStatus3Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, TransactionGroupStatus3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type TransactionGroupStatus3Code", s)
	}

	return TransactionGroupStatus3Code(s), nil
}

// MustToTransactionGroupStatus3Code method for easy conversion with application of restrictions. Panics on error.
func MustToTransactionGroupStatus3Code(s interface{}) TransactionGroupStatus3Code {
	v, err := ToTransactionGroupStatus3Code(s)
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
	Max34TextSample    = "NEKzHKmJBWSEmIYVE"
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
 * SettlementMethod1Code Ops
 */

const (
	SettlementMethod1CodeZero   = ""
	SettlementMethod1CodeSample = "COVE"
	SettlementMethod1CodeINDA   = "INDA"
	SettlementMethod1CodeINGA   = "INGA"
	SettlementMethod1CodeCOVE   = "COVE"
	SettlementMethod1CodeCLRG   = "CLRG"
)

var SettlementMethod1CodeEnumRestriction = []string{SettlementMethod1CodeINDA, SettlementMethod1CodeINGA, SettlementMethod1CodeCOVE, SettlementMethod1CodeCLRG}

// IsValid checks if SettlementMethod1Code of type String is valid
func (t SettlementMethod1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == SettlementMethod1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), SettlementMethod1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t SettlementMethod1Code) String() string {
	return string(t)
}

// ToSettlementMethod1Code method for easy conversion with application of restrictions
func ToSettlementMethod1Code(i interface{}) (SettlementMethod1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, SettlementMethod1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type SettlementMethod1Code", s)
	}

	return SettlementMethod1Code(s), nil
}

// MustToSettlementMethod1Code method for easy conversion with application of restrictions. Panics on error.
func MustToSettlementMethod1Code(s interface{}) SettlementMethod1Code {
	v, err := ToSettlementMethod1Code(s)
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
 * ExternalFinancialInstitutionIdentification1Code Ops
 */

const (
	ExternalFinancialInstitutionIdentification1CodeZero      = ""
	ExternalFinancialInstitutionIdentification1CodeSample    = "Td"
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
 * ExternalStatusReason1Code Ops
 */

const (
	ExternalStatusReason1CodeZero      = ""
	ExternalStatusReason1CodeSample    = "mM"
	ExternalStatusReason1CodeLength    = 0
	ExternalStatusReason1CodeMinLength = 1
	ExternalStatusReason1CodeMaxLength = 4
)

// IsValid checks if ExternalStatusReason1Code of type String is valid
func (t ExternalStatusReason1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == ExternalStatusReason1CodeZero {
		return valid
	}
	valid = valid && isLengthRestrictionValid(t.String(), ExternalStatusReason1CodeLength, ExternalStatusReason1CodeMinLength, ExternalStatusReason1CodeMaxLength)

	return valid
}

// String method for easy conversion
func (t ExternalStatusReason1Code) String() string {
	return string(t)
}

// ToExternalStatusReason1Code method for easy conversion with application of restrictions
func ToExternalStatusReason1Code(i interface{}) (ExternalStatusReason1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, ExternalStatusReason1CodeLength, ExternalStatusReason1CodeMinLength, ExternalStatusReason1CodeMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type ExternalStatusReason1Code", s)
	}

	return ExternalStatusReason1Code(s), nil
}

// MustToExternalStatusReason1Code method for easy conversion with application of restrictions. Panics on error.
func MustToExternalStatusReason1Code(s interface{}) ExternalStatusReason1Code {
	v, err := ToExternalStatusReason1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Frequency1Code Ops
 */

const (
	Frequency1CodeZero   = ""
	Frequency1CodeSample = "WEEK"
	Frequency1CodeYEAR   = "YEAR"
	Frequency1CodeMNTH   = "MNTH"
	Frequency1CodeQURT   = "QURT"
	Frequency1CodeMIAN   = "MIAN"
	Frequency1CodeWEEK   = "WEEK"
	Frequency1CodeDAIL   = "DAIL"
	Frequency1CodeADHO   = "ADHO"
	Frequency1CodeINDA   = "INDA"
)

var Frequency1CodeEnumRestriction = []string{Frequency1CodeYEAR, Frequency1CodeMNTH, Frequency1CodeQURT, Frequency1CodeMIAN, Frequency1CodeWEEK, Frequency1CodeDAIL, Frequency1CodeADHO, Frequency1CodeINDA}

// IsValid checks if Frequency1Code of type String is valid
func (t Frequency1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == Frequency1CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), Frequency1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t Frequency1Code) String() string {
	return string(t)
}

// ToFrequency1Code method for easy conversion with application of restrictions
func ToFrequency1Code(i interface{}) (Frequency1Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, Frequency1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type Frequency1Code", s)
	}

	return Frequency1Code(s), nil
}

// MustToFrequency1Code method for easy conversion with application of restrictions. Panics on error.
func MustToFrequency1Code(s interface{}) Frequency1Code {
	v, err := ToFrequency1Code(s)
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
 * ExternalServiceLevel1Code Ops
 */

const (
	ExternalServiceLevel1CodeZero      = ""
	ExternalServiceLevel1CodeSample    = "ap"
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
 * TransactionIndividualStatus3Code Ops
 */

const (
	TransactionIndividualStatus3CodeZero   = ""
	TransactionIndividualStatus3CodeSample = "ACCP"
	TransactionIndividualStatus3CodeACTC   = "ACTC"
	TransactionIndividualStatus3CodeRJCT   = "RJCT"
	TransactionIndividualStatus3CodePDNG   = "PDNG"
	TransactionIndividualStatus3CodeACCP   = "ACCP"
	TransactionIndividualStatus3CodeACSP   = "ACSP"
	TransactionIndividualStatus3CodeACSC   = "ACSC"
	TransactionIndividualStatus3CodeACWC   = "ACWC"
)

var TransactionIndividualStatus3CodeEnumRestriction = []string{TransactionIndividualStatus3CodeACTC, TransactionIndividualStatus3CodeRJCT, TransactionIndividualStatus3CodePDNG, TransactionIndividualStatus3CodeACCP, TransactionIndividualStatus3CodeACSP, TransactionIndividualStatus3CodeACSC, TransactionIndividualStatus3CodeACWC}

// IsValid checks if TransactionIndividualStatus3Code of type String is valid
func (t TransactionIndividualStatus3Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)

	if optional && t == TransactionIndividualStatus3CodeZero {
		return valid
	}
	valid = valid && isEnumRestrictionValid(t.String(), TransactionIndividualStatus3CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t TransactionIndividualStatus3Code) String() string {
	return string(t)
}

// ToTransactionIndividualStatus3Code method for easy conversion with application of restrictions
func ToTransactionIndividualStatus3Code(i interface{}) (TransactionIndividualStatus3Code, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isEnumRestrictionValid(s, TransactionIndividualStatus3CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type TransactionIndividualStatus3Code", s)
	}

	return TransactionIndividualStatus3Code(s), nil
}

// MustToTransactionIndividualStatus3Code method for easy conversion with application of restrictions. Panics on error.
func MustToTransactionIndividualStatus3Code(s interface{}) TransactionIndividualStatus3Code {
	v, err := ToTransactionIndividualStatus3Code(s)
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
	Max70TextSample    = "yDvfwRbnAOSHNYXsjcDxWEITHgqlDYhLFJM"
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
 * CountryCode Ops
 */

const (
	CountryCodeZero   = ""
	CountryCodeSample = "JG"
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
 * AnyBICIdentifier Ops
 */

const (
	AnyBICIdentifierZero   = ""
	AnyBICIdentifierSample = "FCEGGXLG"
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
 * Max15NumericText Ops
 */

const (
	Max15NumericTextZero   = ""
	Max15NumericTextSample = "29202984"
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
 * Max140Text Ops
 */

const (
	Max140TextZero      = ""
	Max140TextSample    = "OcTZZFWapmAOnXsexQFSvTVoWpQQtMnMOzQQIvZvysnXgPLZTnLDrwaxkScPbAfGypnykB"
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
 * BICIdentifier Ops
 */

const (
	BICIdentifierZero   = ""
	BICIdentifierSample = "UYQQXHMX"
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
