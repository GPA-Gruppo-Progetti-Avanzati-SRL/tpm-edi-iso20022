// Package pain_008_001_02
// Do not Edit. This stuff it's been automatically generated.
package pain_008_001_02

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

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

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
}

// BranchAndFinancialInstitutionIdentification4 type definition
type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId *FinancialInstitutionIdentification7 `xml:"FinInstnId,omitempty"`
	BrnchId    *BranchData2                         `xml:"BrnchId,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// ReferredDocumentInformation3 type definition
type ReferredDocumentInformation3 struct {
	Tp     *ReferredDocumentType2 `xml:"Tp,omitempty"`
	Nb     common.Max35Text       `xml:"Nb,omitempty"`
	RltdDt common.ISODate         `xml:"RltdDt,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id,omitempty"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
}

// RemittanceLocation2 type definition
type RemittanceLocation2 struct {
	RmtId             common.Max35Text                     `xml:"RmtId,omitempty"`
	RmtLctnMtd        common.RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr common.Max2048Text                   `xml:"RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    *NameAndAddress10                    `xml:"RmtLctnPstlAdr,omitempty"`
}

// ReferredDocumentType2 type definition
type ReferredDocumentType2 struct {
	CdOrPrtry *ReferredDocumentType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
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

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// PaymentInstructionInformation4 type definition
type PaymentInstructionInformation4 struct {
	PmtInfId     common.Max35Text                              `xml:"PmtInfId,omitempty"`
	PmtMtd       common.PaymentMethod2Code                     `xml:"PmtMtd,omitempty"`
	BtchBookg    xsdt.Boolean                                  `xml:"BtchBookg,omitempty"`
	NbOfTxs      common.Max15NumericText                       `xml:"NbOfTxs,omitempty"`
	CtrlSum      xsdt.Decimal                                  `xml:"CtrlSum,omitempty"`
	PmtTpInf     *PaymentTypeInformation20                     `xml:"PmtTpInf,omitempty"`
	ReqdColltnDt common.ISODate                                `xml:"ReqdColltnDt,omitempty"`
	Cdtr         *PartyIdentification32                        `xml:"Cdtr,omitempty"`
	CdtrAcct     *CashAccount16                                `xml:"CdtrAcct,omitempty"`
	CdtrAgt      *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct  *CashAccount16                                `xml:"CdtrAgtAcct,omitempty"`
	UltmtCdtr    *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
	ChrgBr       common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	ChrgsAcct    *CashAccount16                                `xml:"ChrgsAcct,omitempty"`
	ChrgsAcctAgt *BranchAndFinancialInstitutionIdentification4 `xml:"ChrgsAcctAgt,omitempty"`
	CdtrSchmeId  *PartyIdentification32                        `xml:"CdtrSchmeId,omitempty"`
	DrctDbtTxInf []DirectDebitTransactionInformation9          `xml:"DrctDbtTxInf,omitempty"`
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

// MandateRelatedInformation6 type definition
type MandateRelatedInformation6 struct {
	MndtId        common.Max35Text              `xml:"MndtId,omitempty"`
	DtOfSgntr     common.ISODate                `xml:"DtOfSgntr,omitempty"`
	AmdmntInd     xsdt.Boolean                  `xml:"AmdmntInd,omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails6 `xml:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                `xml:"FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                `xml:"FnlColltnDt,omitempty"`
	Frqcy         common.Frequency1Code         `xml:"Frqcy,omitempty"`
}

// NameAndAddress10 type definition
type NameAndAddress10 struct {
	Nm  common.Max140Text `xml:"Nm,omitempty"`
	Adr *PostalAddress6   `xml:"Adr,omitempty"`
}

// CustomerDirectDebitInitiationV02 type definition
type CustomerDirectDebitInitiationV02 struct {
	GrpHdr *GroupHeader39                   `xml:"GrpHdr,omitempty"`
	PmtInf []PaymentInstructionInformation4 `xml:"PmtInf,omitempty"`
}

// RegulatoryAuthority2 type definition
type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"Nm,omitempty"`
	Ctry common.CountryCode `xml:"Ctry,omitempty"`
}

// FinancialInstitutionIdentification7 type definition
type FinancialInstitutionIdentification7 struct {
	BIC         common.BICIdentifier                 `xml:"BIC,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr,omitempty"`
	Value xsdt.Decimal                        `xml:",chardata"`
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

// TaxPeriod1 type definition
type TaxPeriod1 struct {
	Yr     common.ISODate              `xml:"Yr,omitempty"`
	Tp     common.TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriodDetails          `xml:"FrToDt,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId,omitempty"`
}

// PersonIdentification5 type definition
type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// DateAndPlaceOfBirth type definition
type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt,omitempty"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth,omitempty"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth,omitempty"`
}

// CategoryPurpose1Choice type definition
type CategoryPurpose1Choice struct {
	Cd    common.ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    common.ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
}

// DirectDebitTransactionInformation9 type definition
type DirectDebitTransactionInformation9 struct {
	PmtId           *PaymentIdentification1                       `xml:"PmtId,omitempty"`
	PmtTpInf        *PaymentTypeInformation20                     `xml:"PmtTpInf,omitempty"`
	InstdAmt        *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty"`
	ChrgBr          common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	DrctDbtTx       *DirectDebitTransaction6                      `xml:"DrctDbtTx,omitempty"`
	UltmtCdtr       *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
	DbtrAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct     *CashAccount16                                `xml:"DbtrAgtAcct,omitempty"`
	Dbtr            *PartyIdentification32                        `xml:"Dbtr,omitempty"`
	DbtrAcct        *CashAccount16                                `xml:"DbtrAcct,omitempty"`
	UltmtDbtr       *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	InstrForCdtrAgt common.Max140Text                             `xml:"InstrForCdtrAgt,omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation3                              `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation2                         `xml:"RltdRmtInf,omitempty"`
	RmtInf          *RemittanceInformation5                       `xml:"RmtInf,omitempty"`
}

// Authorisation1Choice type definition
type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"Cd,omitempty"`
	Prtry common.Max128Text         `xml:"Prtry,omitempty"`
}

// GroupHeader39 type definition
type GroupHeader39 struct {
	MsgId    common.Max35Text                              `xml:"MsgId,omitempty"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm,omitempty"`
	Authstn  []Authorisation1Choice                        `xml:"Authstn,omitempty"`
	NbOfTxs  common.Max15NumericText                       `xml:"NbOfTxs,omitempty"`
	CtrlSum  xsdt.Decimal                                  `xml:"CtrlSum,omitempty"`
	InitgPty *PartyIdentification32                        `xml:"InitgPty,omitempty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id,omitempty"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id,omitempty"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// DirectDebitTransaction6 type definition
type DirectDebitTransaction6 struct {
	MndtRltdInf *MandateRelatedInformation6 `xml:"MndtRltdInf,omitempty"`
	CdtrSchmeId *PartyIdentification32      `xml:"CdtrSchmeId,omitempty"`
	PreNtfctnId common.Max35Text            `xml:"PreNtfctnId,omitempty"`
	PreNtfctnDt common.ISODate              `xml:"PreNtfctnDt,omitempty"`
}

// PartyIdentification32 type definition
type PartyIdentification32 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress6    `xml:"PstlAdr,omitempty"`
	Id        *Party6Choice      `xml:"Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2   `xml:"CtctDtls,omitempty"`
}

// OrganisationIdentification4 type definition
type OrganisationIdentification4 struct {
	BICOrBEI common.AnyBICIdentifier              `xml:"BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// PaymentTypeInformation20 type definition
type PaymentTypeInformation20 struct {
	InstrPrty common.Priority2Code     `xml:"InstrPrty,omitempty"`
	SvcLvl    *ServiceLevel8Choice     `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice  `xml:"LclInstrm,omitempty"`
	SeqTp     common.SequenceType1Code `xml:"SeqTp,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice  `xml:"CtgyPurp,omitempty"`
}

// Party6Choice type definition
type Party6Choice struct {
	OrgId  *OrganisationIdentification4 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    common.ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                               `xml:"Prtry,omitempty"`
}

// LocalInstrument2Choice type definition
type LocalInstrument2Choice struct {
	Cd    common.ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// AmendmentInformationDetails6 type definition
type AmendmentInformationDetails6 struct {
	OrgnlMndtId      common.Max35Text                              `xml:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification32                        `xml:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct *CashAccount16                                `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        *PartyIdentification32                        `xml:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    *CashAccount16                                `xml:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct *CashAccount16                                `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                                `xml:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       common.Frequency1Code                         `xml:"OrgnlFrqcy,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// CashAccount16 type definition
type CashAccount16 struct {
	Id  *AccountIdentification4Choice       `xml:"Id,omitempty"`
	Tp  *CashAccountType2                   `xml:"Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"Nm,omitempty"`
}

// RemittanceInformation5 type definition
type RemittanceInformation5 struct {
	Ustrd []common.Max140Text                `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    common.ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text              `xml:"Issr,omitempty"`
}

// PaymentIdentification1 type definition
type PaymentIdentification1 struct {
	InstrId    common.Max35Text `xml:"InstrId,omitempty"`
	EndToEndId common.Max35Text `xml:"EndToEndId,omitempty"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   common.Max35Text   `xml:"TaxId,omitempty"`
	RegnId  common.Max35Text   `xml:"RegnId,omitempty"`
	TaxTp   common.Max35Text   `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
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

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id,omitempty"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    common.ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// RegulatoryReporting3 type definition
type RegulatoryReporting3 struct {
	DbtCdtRptgInd common.RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2               `xml:"Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3    `xml:"Dtls,omitempty"`
}

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty"`
}

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty"`
}

// DatePeriodDetails type definition
type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt,omitempty"`
	ToDt common.ISODate `xml:"ToDt,omitempty"`
}

// CashAccountType2 type definition
type CashAccountType2 struct {
	Cd    common.CashAccountType4Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// ReferredDocumentType1Choice type definition
type ReferredDocumentType1Choice struct {
	Cd    common.DocumentType5Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// BranchData2 type definition
type BranchData2 struct {
	Id      common.Max35Text  `xml:"Id,omitempty"`
	Nm      common.Max140Text `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress6   `xml:"PstlAdr,omitempty"`
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

// TaxRecordDetails1 type definition
type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                        `xml:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// TaxAmount1 type definition
type TaxAmount1 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CdtDbtInd common.CreditDebitCode             `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                    `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                  `xml:"AddtlInf,omitempty"`
}
