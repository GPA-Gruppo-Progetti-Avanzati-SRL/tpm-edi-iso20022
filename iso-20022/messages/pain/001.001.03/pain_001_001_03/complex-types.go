// Package pain_001_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_001_001_03

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03/xsdt"
)

// RegulatoryAuthority2 type definition
type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"Nm,omitempty"`
	Ctry common.CountryCode `xml:"Ctry,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
}

// ContactDetails2 type definition
type ContactDetails2 struct {
	NmPrfx   common.NamePrefix1Code `xml:"NmPrfx,omitempty"`
	Nm       common.Max140Text      `xml:"Nm,omitempty"`
	PhneNb   common.PhoneNumber     `xml:"PhneNb,omitempty"`
	MobNb    common.PhoneNumber     `xml:"MobNb,omitempty"`
	FaxNb    common.PhoneNumber     `xml:"FaxNb,omitempty"`
	EmailAdr common.Max2048Text     `xml:"EmailAdr,omitempty"`
	Othr     common.Max35Text       `xml:"Othr,omitempty"`
}

// Cheque6 type definition
type Cheque6 struct {
	ChqTp       common.ChequeType2Code       `xml:"ChqTp,omitempty"`
	ChqNb       common.Max35Text             `xml:"ChqNb,omitempty"`
	ChqFr       *NameAndAddress10            `xml:"ChqFr,omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`
	DlvrTo      *NameAndAddress10            `xml:"DlvrTo,omitempty"`
	InstrPrty   common.Priority2Code         `xml:"InstrPrty,omitempty"`
	ChqMtrtyDt  common.ISODate               `xml:"ChqMtrtyDt,omitempty"`
	FrmsCd      common.Max35Text             `xml:"FrmsCd,omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty"`
	RgnlClrZone common.Max35Text             `xml:"RgnlClrZone,omitempty"`
	PrtLctn     common.Max35Text             `xml:"PrtLctn,omitempty"`
}

// LocalInstrument2Choice type definition
type LocalInstrument2Choice struct {
	Cd    common.ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// TaxInformation3 type definition
type TaxInformation3 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty"`
	AdmstnZn        common.Max35Text                   `xml:"AdmstnZn,omitempty"`
	RefNb           common.Max140Text                  `xml:"RefNb,omitempty"`
	Mtd             common.Max35Text                   `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              common.ISODate                     `xml:"Dt,omitempty"`
	SeqNb           xsdt.Decimal                       `xml:"SeqNb,omitempty"`
	Rcrd            []TaxRecord1                       `xml:"Rcrd,omitempty"`
}

// TaxPeriod1 type definition
type TaxPeriod1 struct {
	Yr     common.ISODate              `xml:"Yr,omitempty"`
	Tp     common.TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriodDetails          `xml:"FrToDt,omitempty"`
}

// ReferredDocumentType2 type definition
type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text            `xml:"Issr,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
	Value xsdt.Decimal                        `xml:",chardata"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// ChequeDeliveryMethod1Choice type definition
type ChequeDeliveryMethod1Choice struct {
	Cd    common.ChequeDelivery1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text           `xml:"Prtry,omitempty"`
}

// PartyIdentification32 type definition
type PartyIdentification32 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress6    `xml:"PstlAdr,omitempty"`
	Id        *Party6Choice      `xml:"Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2   `xml:"CtctDtls,omitempty"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// CashAccount16 type definition
type CashAccount16 struct {
	Id  AccountIdentification4Choice        `xml:"Id"`
	Tp  *CashAccountType2                   `xml:"Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"Nm,omitempty"`
}

// TaxAmount1 type definition
type TaxAmount1 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty"`
}

// ReferredDocumentInformation3 type definition
type ReferredDocumentInformation3 struct {
	Tp     *ReferredDocumentType2 `xml:"Tp,omitempty"`
	Nb     common.Max35Text       `xml:"Nb,omitempty"`
	RltdDt common.ISODate         `xml:"RltdDt,omitempty"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// GroupHeader32 type definition
type GroupHeader32 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	Authstn  []Authorisation1Choice                        `xml:"Authstn,omitempty"`
	NbOfTxs  common.Max15NumericText                       `xml:"NbOfTxs"`
	CtrlSum  xsdt.Decimal                                  `xml:"CtrlSum,omitempty"`
	InitgPty PartyIdentification32                         `xml:"InitgPty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
}

// EquivalentAmount2 type definition
type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    common.ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    common.ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
}

// TaxRecordDetails1 type definition
type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                       `xml:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// CustomerCreditTransferInitiationV03 type definition
type CustomerCreditTransferInitiationV03 struct {
	GrpHdr GroupHeader32                    `xml:"GrpHdr"`
	PmtInf []PaymentInstructionInformation3 `xml:"PmtInf"`
}

// CashAccountType2 type definition
type CashAccountType2 struct {
	Cd    common.CashAccountType4Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// BranchAndFinancialInstitutionIdentification4 type definition
type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"FinInstnId"`
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty"`
}

// PaymentIdentification1 type definition
type PaymentIdentification1 struct {
	InstrId    common.Max35Text `xml:"InstrId,omitempty"`
	EndToEndId common.Max35Text `xml:"EndToEndId"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   common.Max35Text   `xml:"TaxId,omitempty"`
	RegnId  common.Max35Text   `xml:"RegnId,omitempty"`
	TaxTp   common.Max35Text   `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// StructuredRemittanceInformation7 type definition
type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount1             `xml:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification32         `xml:"Invcr,omitempty"`
	Invcee      *PartyIdentification32         `xml:"Invcee,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
}

// PaymentTypeInformation19 type definition
type PaymentTypeInformation19 struct {
	InstrPrty common.Priority2Code    `xml:"InstrPrty,omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// PersonIdentification5 type definition
type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// ExchangeRateInformation1 type definition
type ExchangeRateInformation1 struct {
	XchgRate xsdt.Decimal                 `xml:"XchgRate,omitempty"`
	RateTp   common.ExchangeRateType1Code `xml:"RateTp,omitempty"`
	CtrctId  common.Max35Text             `xml:"CtrctId,omitempty"`
}

// RemittanceInformation5 type definition
type RemittanceInformation5 struct {
	Ustrd []common.Max140Text                `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

// BranchData2 type definition
type BranchData2 struct {
	Id      common.Max35Text  `xml:"Id,omitempty"`
	Nm      common.Max140Text `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress6   `xml:"PstlAdr,omitempty"`
}

// AmountType3Choice type definition
type AmountType3Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// CreditTransferTransactionInformation10 type definition
type CreditTransferTransactionInformation10 struct {
	PmtId           PaymentIdentification1                        `xml:"PmtId"`
	PmtTpInf        *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty"`
	Amt             AmountType3Choice                             `xml:"Amt"`
	XchgRateInf     *ExchangeRateInformation1                     `xml:"XchgRateInf,omitempty"`
	ChrgBr          common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	ChqInstr        *Cheque6                                      `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct  *CashAccount16                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct  *CashAccount16                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct  *CashAccount16                                `xml:"IntrmyAgt3Acct,omitempty"`
	CdtrAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct     *CashAccount16                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr            *PartyIdentification32                        `xml:"Cdtr,omitempty"`
	CdtrAcct        *CashAccount16                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty"`
	InstrForDbtrAgt common.Max140Text                             `xml:"InstrForDbtrAgt,omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation3                              `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation2                         `xml:"RltdRmtInf,omitempty"`
	RmtInf          *RemittanceInformation5                       `xml:"RmtInf,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// Party6Choice type definition
type Party6Choice struct {
	OrgId  *OrganisationIdentification4 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

// FinancialInstitutionIdentification7 type definition
type FinancialInstitutionIdentification7 struct {
	BIC         common.BICIdentifier                 `xml:"BIC,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// NameAndAddress10 type definition
type NameAndAddress10 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress6    `xml:"Adr"`
}

// PostalAddress6 type definition
type PostalAddress6 struct {
	AdrTp       common.AddressType2Code `xml:"AdrTp,omitempty"`
	Dept        common.Max70Text        `xml:"Dept,omitempty"`
	SubDept     common.Max70Text        `xml:"SubDept,omitempty"`
	StrtNm      common.Max70Text        `xml:"StrtNm,omitempty"`
	BldgNb      common.Max16Text        `xml:"BldgNb,omitempty"`
	PstCd       common.Max16Text        `xml:"PstCd,omitempty"`
	TwnNm       common.Max35Text        `xml:"TwnNm,omitempty"`
	CtrySubDvsn common.Max35Text        `xml:"CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode      `xml:"Ctry,omitempty"`
	AdrLine     []common.Max70Text      `xml:"AdrLine,omitempty"`
}

// TaxRecord1 type definition
type TaxRecord1 struct {
	Tp       common.Max35Text  `xml:"Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"FrmsCd,omitempty"`
	Prd      *TaxPeriod1       `xml:"Prd,omitempty"`
	TaxAmt   *TaxAmount1       `xml:"TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"AddtlInf,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// RemittanceAmount1 type definition
type RemittanceAmount1 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      *ActiveOrHistoricCurrencyAndAmount `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            *ActiveOrHistoricCurrencyAndAmount `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// PaymentInstructionInformation3 type definition
type PaymentInstructionInformation3 struct {
	PmtInfId        common.Max35Text                              `xml:"PmtInfId"`
	PmtMtd          common.PaymentMethod3Code                     `xml:"PmtMtd"`
	BtchBookg       xsdt.Boolean                                  `xml:"BtchBookg,omitempty"`
	NbOfTxs         common.Max15NumericText                       `xml:"NbOfTxs,omitempty"`
	CtrlSum         xsdt.Decimal                                  `xml:"CtrlSum,omitempty"`
	PmtTpInf        *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty"`
	ReqdExctnDt     common.ISODate                                `xml:"ReqdExctnDt"`
	PoolgAdjstmntDt common.ISODate                                `xml:"PoolgAdjstmntDt,omitempty"`
	Dbtr            PartyIdentification32                         `xml:"Dbtr"`
	DbtrAcct        CashAccount16                                 `xml:"DbtrAcct"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification4  `xml:"DbtrAgt"`
	DbtrAgtAcct     *CashAccount16                                `xml:"DbtrAgtAcct,omitempty"`
	UltmtDbtr       *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	ChrgBr          common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	ChrgsAcct       *CashAccount16                                `xml:"ChrgsAcct,omitempty"`
	ChrgsAcctAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"ChrgsAcctAgt,omitempty"`
	CdtTrfTxInf     []CreditTransferTransactionInformation10      `xml:"CdtTrfTxInf"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    common.ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// OrganisationIdentification4 type definition
type OrganisationIdentification4 struct {
	BICOrBEI common.AnyBICIdentifier              `xml:"BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// RegulatoryReporting3 type definition
type RegulatoryReporting3 struct {
	DbtCdtRptgInd common.RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2               `xml:"Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3    `xml:"Dtls,omitempty"`
}

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty"`
}

// DateAndPlaceOfBirth type definition
type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

// CategoryPurpose1Choice type definition
type CategoryPurpose1Choice struct {
	Cd    common.ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty"`
}

// ReferredDocumentType1Choice type definition
type ReferredDocumentType1Choice struct {
	Cd    common.DocumentType5Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty"`
}

// DatePeriodDetails type definition
type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

// RemittanceLocation2 type definition
type RemittanceLocation2 struct {
	RmtId             common.Max35Text                     `xml:"RmtId,omitempty"`
	RmtLctnMtd        common.RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr common.Max2048Text                   `xml:"RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    *NameAndAddress10                    `xml:"RmtLctnPstlAdr,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    common.ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                               `xml:"Prtry,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
}

// StructuredRegulatoryReporting3 type definition
type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                   `xml:"Tp,omitempty"`
	Dt   common.ISODate                     `xml:"Dt,omitempty"`
	Ctry common.CountryCode                 `xml:"Ctry,omitempty"`
	Cd   common.Max10Text                   `xml:"Cd,omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	Inf  []common.Max35Text                 `xml:"Inf,omitempty"`
}

// InstructionForCreditorAgent1 type definition
type InstructionForCreditorAgent1 struct {
	Cd       common.Instruction3Code `xml:"Cd,omitempty"`
	InstrInf common.Max140Text       `xml:"InstrInf,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

// Authorisation1Choice type definition
type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"Cd,omitempty"`
	Prtry common.Max128Text         `xml:"Prtry,omitempty"`
}
