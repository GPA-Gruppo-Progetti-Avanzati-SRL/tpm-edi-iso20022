// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/013.001.07/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// DiscountAmountAndType1 type definition
type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice         `xml:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
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

// GenericIdentification1 type definition
type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"Id,omitempty"`
	SchmeNm common.Max35Text `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text `xml:"Issr,omitempty"`
}

// PartyAndSignature3 type definition
type PartyAndSignature3 struct {
	Pty   *PartyIdentification135 `xml:"Pty,omitempty"`
	Sgntr *SkipPayload            `xml:"Sgntr,omitempty"`
}

// CreditTransferTransaction35 type definition
type CreditTransferTransaction35 struct {
	PmtId           *PaymentIdentification6                       `xml:"PmtId,omitempty"`
	PmtTpInf        *PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtCond         *PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Amt             *AmountType4Choice                            `xml:"Amt,omitempty"`
	ChrgBr          common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	ChqInstr        *Cheque11                                     `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	CdtrAgt         *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	Cdtr            *PartyIdentification135                       `xml:"Cdtr,omitempty"`
	CdtrAcct        *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation8                              `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                         `xml:"RltdRmtInf,omitempty"`
	RmtInf          *RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile       []Document12                                  `xml:"NclsdFile,omitempty"`
	SplmtryData     []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// RemittanceLocation7 type definition
type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty"`
}

// DocumentLineType1Choice type definition
type DocumentLineType1Choice struct {
	Cd    common.ExternalDocumentLineType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                     `xml:"Prtry,omitempty"`
}

// DateAndDateTime2Choice type definition
type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt,omitempty"`
	DtTm common.ISODateTime `xml:"DtTm,omitempty"`
}

// PaymentCondition1 type definition
type PaymentCondition1 struct {
	AmtModAllwd   xsdt.Boolean         `xml:"AmtModAllwd,omitempty"`
	EarlyPmtAllwd xsdt.Boolean         `xml:"EarlyPmtAllwd,omitempty"`
	DelyPnlty     common.Max140Text    `xml:"DelyPnlty,omitempty"`
	ImdtPmtRbt    *AmountOrRate1Choice `xml:"ImdtPmtRbt,omitempty"`
	GrntedPmtReqd xsdt.Boolean         `xml:"GrntedPmtReqd,omitempty"`
}

// CashAccountType2Choice type definition
type CashAccountType2Choice struct {
	Cd    common.ExternalCashAccountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId,omitempty"`
}

// EquivalentAmount2 type definition
type EquivalentAmount2 struct {
	Amt      *ActiveOrHistoricCurrencyAndAmount  `xml:"Amt,omitempty"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf,omitempty"`
}

// ChequeDeliveryMethod1Choice type definition
type ChequeDeliveryMethod1Choice struct {
	Cd    common.ChequeDelivery1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text           `xml:"Prtry,omitempty"`
}

// AddressType3Choice type definition
type AddressType3Choice struct {
	Cd    common.AddressType2Code  `xml:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"Prtry,omitempty"`
}

// BranchAndFinancialInstitutionIdentification6 type definition
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId,omitempty"`
	BrnchId    *BranchData3                          `xml:"BrnchId,omitempty"`
}

// DocumentType1Choice type definition
type DocumentType1Choice struct {
	Cd    common.ExternalDocumentType1Code `xml:"Cd,omitempty"`
	Prtry *GenericIdentification1          `xml:"Prtry,omitempty"`
}

// SupplementaryData1 type definition
type SupplementaryData1 struct {
	PlcAndNm common.Max350Text           `xml:"PlcAndNm,omitempty"`
	Envlp    *SupplementaryDataEnvelope1 `xml:"Envlp,omitempty"`
}

// Contact4 type definition
type Contact4 struct {
	NmPrfx    common.NamePrefix2Code             `xml:"NmPrfx,omitempty"`
	Nm        common.Max140Text                  `xml:"Nm,omitempty"`
	PhneNb    common.PhoneNumber                 `xml:"PhneNb,omitempty"`
	MobNb     common.PhoneNumber                 `xml:"MobNb,omitempty"`
	FaxNb     common.PhoneNumber                 `xml:"FaxNb,omitempty"`
	EmailAdr  common.Max2048Text                 `xml:"EmailAdr,omitempty"`
	EmailPurp common.Max35Text                   `xml:"EmailPurp,omitempty"`
	JobTitl   common.Max35Text                   `xml:"JobTitl,omitempty"`
	Rspnsblty common.Max35Text                   `xml:"Rspnsblty,omitempty"`
	Dept      common.Max70Text                   `xml:"Dept,omitempty"`
	Othr      []OtherContact1                    `xml:"Othr,omitempty"`
	PrefrdMtd common.PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// GarnishmentType1Choice type definition
type GarnishmentType1Choice struct {
	Cd    common.ExternalGarnishmentType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// RemittanceLocationData1 type definition
type RemittanceLocationData1 struct {
	Mtd        common.RemittanceLocationMethod2Code `xml:"Mtd,omitempty"`
	ElctrncAdr common.Max2048Text                   `xml:"ElctrncAdr,omitempty"`
	PstlAdr    *NameAndAddress16                    `xml:"PstlAdr,omitempty"`
}

// RemittanceInformation16 type definition
type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty"`
}

// TaxInformation7 type definition
type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                   `xml:"AdmstnZone,omitempty"`
	RefNb           common.Max140Text                  `xml:"RefNb,omitempty"`
	Mtd             common.Max35Text                   `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              common.ISODate                     `xml:"Dt,omitempty"`
	SeqNb           xsdt.Decimal                       `xml:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty"`
}

// PaymentIdentification6 type definition
type PaymentIdentification6 struct {
	InstrId    common.Max35Text        `xml:"InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"EndToEndId,omitempty"`
	UETR       common.UUIDv4Identifier `xml:"UETR,omitempty"`
}

// RegulatoryReporting3 type definition
type RegulatoryReporting3 struct {
	DbtCdtRptgInd common.RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2               `xml:"Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3    `xml:"Dtls,omitempty"`
}

// DocumentLineInformation1 type definition
type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id,omitempty"`
	Desc common.Max2048Text            `xml:"Desc,omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// OrganisationIdentification29 type definition
type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// CashAccount38 type definition
type CashAccount38 struct {
	Id   *AccountIdentification4Choice       `xml:"Id,omitempty"`
	Tp   *CashAccountType2Choice             `xml:"Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1        `xml:"Prxy,omitempty"`
}

// Document12 type definition
type Document12 struct {
	Tp        *DocumentType1Choice    `xml:"Tp,omitempty"`
	Id        common.Max35Text        `xml:"Id,omitempty"`
	IsseDt    *DateAndDateTime2Choice `xml:"IsseDt,omitempty"`
	Nm        common.Max140Text       `xml:"Nm,omitempty"`
	LangCd    xsdt.String             `xml:"LangCd,omitempty"`
	Frmt      *DocumentFormat1Choice  `xml:"Frmt,omitempty"`
	FileNm    common.Max140Text       `xml:"FileNm,omitempty"`
	DgtlSgntr *PartyAndSignature3     `xml:"DgtlSgntr,omitempty"`
	Nclsr     common.Max10MbBinary    `xml:"Nclsr,omitempty"`
}

// GenericIdentification30 type definition
type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id,omitempty"`
	Issr    common.Max35Text              `xml:"Issr,omitempty"`
	SchmeNm common.Max35Text              `xml:"SchmeNm,omitempty"`
}

// TaxInformation8 type definition
type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty"`
	AdmstnZone      common.Max35Text                   `xml:"AdmstnZone,omitempty"`
	RefNb           common.Max140Text                  `xml:"RefNb,omitempty"`
	Mtd             common.Max35Text                   `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              common.ISODate                     `xml:"Dt,omitempty"`
	SeqNb           xsdt.Decimal                       `xml:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty"`
}

// DiscountAmountType1Choice type definition
type DiscountAmountType1Choice struct {
	Cd    common.ExternalDiscountAmountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                       `xml:"Prtry,omitempty"`
}

// SupplementaryDataEnvelope1 type definition
type SupplementaryDataEnvelope1 struct {
	Item xsdt.AnyType `xml:",omitempty,any"`
}

// AmountOrRate1Choice type definition
type AmountOrRate1Choice struct {
	Amt  *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`
	Rate xsdt.Decimal             `xml:"Rate,omitempty"`
}

// ReferredDocumentType3Choice type definition
type ReferredDocumentType3Choice struct {
	Cd    common.DocumentType6Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// TaxRecord2 type definition
type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"FrmsCd,omitempty"`
	Prd      *TaxPeriod2       `xml:"Prd,omitempty"`
	TaxAmt   *TaxAmount2       `xml:"TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"AddtlInf,omitempty"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    common.ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// ProxyAccountIdentification1 type definition
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty"`
	Id common.Max2048Text       `xml:"Id,omitempty"`
}

// AmountType4Choice type definition
type AmountType4Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// PostalAddress24 type definition
type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty"`
	Dept        common.Max70Text    `xml:"Dept,omitempty"`
	SubDept     common.Max70Text    `xml:"SubDept,omitempty"`
	StrtNm      common.Max70Text    `xml:"StrtNm,omitempty"`
	BldgNb      common.Max16Text    `xml:"BldgNb,omitempty"`
	BldgNm      common.Max35Text    `xml:"BldgNm,omitempty"`
	Flr         common.Max70Text    `xml:"Flr,omitempty"`
	PstBx       common.Max16Text    `xml:"PstBx,omitempty"`
	Room        common.Max70Text    `xml:"Room,omitempty"`
	PstCd       common.Max16Text    `xml:"PstCd,omitempty"`
	TwnNm       common.Max35Text    `xml:"TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text    `xml:"TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text    `xml:"DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text    `xml:"CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode  `xml:"Ctry,omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id,omitempty"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
}

// NameAndAddress16 type definition
type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm,omitempty"`
	Adr *PostalAddress24  `xml:"Adr,omitempty"`
}

// TaxRecordDetails2 type definition
type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                        `xml:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// GroupHeader78 type definition
type GroupHeader78 struct {
	MsgId    common.Max35Text        `xml:"MsgId,omitempty"`
	CreDtTm  common.ISODateTime      `xml:"CreDtTm,omitempty"`
	NbOfTxs  common.Max15NumericText `xml:"NbOfTxs,omitempty"`
	CtrlSum  xsdt.Decimal            `xml:"CtrlSum,omitempty"`
	InitgPty *PartyIdentification135 `xml:"InitgPty,omitempty"`
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

// RemittanceAmount2 type definition
type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// TaxAmountAndType1 type definition
type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice              `xml:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// DocumentFormat1Choice type definition
type DocumentFormat1Choice struct {
	Cd    common.ExternalDocumentFormat1Code `xml:"Cd,omitempty"`
	Prtry *GenericIdentification1            `xml:"Prtry,omitempty"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    common.ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// Cheque11 type definition
type Cheque11 struct {
	ChqTp       common.ChequeType2Code       `xml:"ChqTp,omitempty"`
	ChqNb       common.Max35Text             `xml:"ChqNb,omitempty"`
	ChqFr       *NameAndAddress16            `xml:"ChqFr,omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`
	DlvrTo      *NameAndAddress16            `xml:"DlvrTo,omitempty"`
	InstrPrty   common.Priority2Code         `xml:"InstrPrty,omitempty"`
	ChqMtrtyDt  common.ISODate               `xml:"ChqMtrtyDt,omitempty"`
	FrmsCd      common.Max35Text             `xml:"FrmsCd,omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty"`
	RgnlClrZone common.Max35Text             `xml:"RgnlClrZone,omitempty"`
	PrtLctn     common.Max35Text             `xml:"PrtLctn,omitempty"`
	Sgntr       []common.Max70Text           `xml:"Sgntr,omitempty"`
}

// PaymentInstruction31 type definition
type PaymentInstruction31 struct {
	PmtInfId    common.Max35Text                              `xml:"PmtInfId,omitempty"`
	PmtMtd      common.PaymentMethod7Code                     `xml:"PmtMtd,omitempty"`
	PmtTpInf    *PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	ReqdExctnDt *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty"`
	XpryDt      *DateAndDateTime2Choice                       `xml:"XpryDt,omitempty"`
	PmtCond     *PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Dbtr        *PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DbtrAcct    *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	UltmtDbtr   *PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	ChrgBr      common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	CdtTrfTx    []CreditTransferTransaction35                 `xml:"CdtTrfTx,omitempty"`
}

// PaymentTypeInformation26 type definition
type PaymentTypeInformation26 struct {
	InstrPrty common.Priority2Code    `xml:"InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   common.Max35Text   `xml:"TaxId,omitempty"`
	RegnId  common.Max35Text   `xml:"RegnId,omitempty"`
	TaxTp   common.Max35Text   `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty"`
}

// DatePeriod2 type definition
type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt,omitempty"`
	ToDt common.ISODate `xml:"ToDt,omitempty"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CdtDbtInd common.CreditDebitCode             `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                    `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                  `xml:"AddtlInf,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
}

// Garnishment3 type definition
type Garnishment3 struct {
	Tp                *GarnishmentType1                  `xml:"Tp,omitempty"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                  `xml:"RefNb,omitempty"`
	Dt                common.ISODate                     `xml:"Dt,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd xsdt.Boolean                       `xml:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  xsdt.Boolean                       `xml:"MplyeeTermntnInd,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// TaxAmount2 type definition
type TaxAmount2 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id,omitempty"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// BranchData3 type definition
type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress24     `xml:"PstlAdr,omitempty"`
}

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty"`
}

// CreditorPaymentActivationRequestV07 type definition
type CreditorPaymentActivationRequestV07 struct {
	GrpHdr      *GroupHeader78         `xml:"GrpHdr,omitempty"`
	PmtInf      []PaymentInstruction31 `xml:"PmtInf,omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty"`
}

// PartyIdentification135 type definition
type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress24   `xml:"PstlAdr,omitempty"`
	Id        *Party38Choice     `xml:"Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty"`
	CtctDtls  *Contact4          `xml:"CtctDtls,omitempty"`
}

// DocumentLineIdentification1 type definition
type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty"`
	Nb     common.Max35Text   `xml:"Nb,omitempty"`
	RltdDt common.ISODate     `xml:"RltdDt,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text              `xml:"Issr,omitempty"`
}

// RegulatoryAuthority2 type definition
type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"Nm,omitempty"`
	Ctry common.CountryCode `xml:"Ctry,omitempty"`
}

// DocumentLineType1 type definition
type DocumentLineType1 struct {
	CdOrPrtry *DocumentLineType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text         `xml:"Issr,omitempty"`
}

// ActiveCurrencyAndAmount type definition
type ActiveCurrencyAndAmount struct {
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr,omitempty"`
	Value xsdt.Decimal              `xml:",omitempty,chardata"`
}

// PersonIdentification13 type definition
type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// RemittanceAmount3 type definition
type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// ReferredDocumentType4 type definition
type ReferredDocumentType4 struct {
	CdOrPrtry *ReferredDocumentType3Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
}

// ProxyAccountType1Choice type definition
type ProxyAccountType1Choice struct {
	Cd    common.ExternalProxyAccountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                     `xml:"Prtry,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr,omitempty"`
	Value xsdt.Decimal                        `xml:",omitempty,chardata"`
}

// TaxPeriod2 type definition
type TaxPeriod2 struct {
	Yr     common.ISODate              `xml:"Yr,omitempty"`
	Tp     common.TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriod2                `xml:"FrToDt,omitempty"`
}

// LocalInstrument2Choice type definition
type LocalInstrument2Choice struct {
	Cd    common.ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// FinancialInstitutionIdentification18 type definition
type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier        `xml:"BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                 `xml:"LEI,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// StructuredRemittanceInformation16 type definition
type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty"`
}

// Party38Choice type definition
type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"PrvtId,omitempty"`
}

// InstructionForCreditorAgent1 type definition
type InstructionForCreditorAgent1 struct {
	Cd       common.Instruction3Code `xml:"Cd,omitempty"`
	InstrInf common.Max140Text       `xml:"InstrInf,omitempty"`
}

// OtherContact1 type definition
type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp,omitempty"`
	Id      common.Max128Text `xml:"Id,omitempty"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id,omitempty"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// TaxAmountType1Choice type definition
type TaxAmountType1Choice struct {
	Cd    common.ExternalTaxAmountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                  `xml:"Prtry,omitempty"`
}

// ReferredDocumentInformation7 type definition
type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty"`
	Nb       common.Max35Text           `xml:"Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

// GarnishmentType1 type definition
type GarnishmentType1 struct {
	CdOrPrtry *GarnishmentType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text        `xml:"Issr,omitempty"`
}

// DateAndPlaceOfBirth1 type definition
type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt,omitempty"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth,omitempty"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth,omitempty"`
}

// SkipPayload type definition
type SkipPayload struct {
	Item xsdt.AnyType `xml:",omitempty,any"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id,omitempty"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    common.ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                               `xml:"Prtry,omitempty"`
}
