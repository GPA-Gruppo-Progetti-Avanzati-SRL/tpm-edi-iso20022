// Package pain_002_001_10
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_10

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.10/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

// GroupHeader86 type definition
type GroupHeader86 struct {
	MsgId    common.Max35Text                              `xml:"MsgId,omitempty"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm,omitempty"`
	InitgPty *PartyIdentification135                       `xml:"InitgPty,omitempty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"FwdgAgt,omitempty"`
	DbtrAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	CdtrAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
}

// DateAndPlaceOfBirth1 type definition
type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt,omitempty"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth,omitempty"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth,omitempty"`
}

// TaxRecordDetails2 type definition
type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                        `xml:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// DiscountAmountType1Choice type definition
type DiscountAmountType1Choice struct {
	Cd    common.ExternalDiscountAmountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                       `xml:"Prtry,omitempty"`
}

// GenericIdentification30 type definition
type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id,omitempty"`
	Issr    common.Max35Text              `xml:"Issr,omitempty"`
	SchmeNm common.Max35Text              `xml:"SchmeNm,omitempty"`
}

// BranchData3 type definition
type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress24     `xml:"PstlAdr,omitempty"`
}

// OriginalTransactionReference28 type definition
type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty"`
	IntrBkSttlmDt  common.ISODate                                `xml:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   common.ISODate                                `xml:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty"`
	PmtMtd         common.PaymentMethod4Code                     `xml:"PmtMtd,omitempty"`
	MndtRltdInf    *MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    common.ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// DocumentLineType1 type definition
type DocumentLineType1 struct {
	CdOrPrtry *DocumentLineType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text         `xml:"Issr,omitempty"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    common.ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// BranchAndFinancialInstitutionIdentification6 type definition
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId,omitempty"`
	BrnchId    *BranchData3                          `xml:"BrnchId,omitempty"`
}

// SupplementaryDataEnvelope1 type definition
type SupplementaryDataEnvelope1 struct {
	Item xsdt.AnyType `xml:",omitempty,any"`
}

// MandateSetupReason1Choice type definition
type MandateSetupReason1Choice struct {
	Cd    common.ExternalMandateSetupReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max70Text                       `xml:"Prtry,omitempty"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   common.Max35Text   `xml:"TaxId,omitempty"`
	RegnId  common.Max35Text   `xml:"RegnId,omitempty"`
	TaxTp   common.Max35Text   `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
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

// OriginalGroupHeader17 type definition
type OriginalGroupHeader17 struct {
	OrgnlMsgId    common.Max35Text                       `xml:"OrgnlMsgId,omitempty"`
	OrgnlMsgNmId  common.Max35Text                       `xml:"OrgnlMsgNmId,omitempty"`
	OrgnlCreDtTm  common.ISODateTime                     `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText                `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                           `xml:"OrgnlCtrlSum,omitempty"`
	GrpSts        common.ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12            `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5       `xml:"NbOfTxsPerSts,omitempty"`
}

// ProxyAccountType1Choice type definition
type ProxyAccountType1Choice struct {
	Cd    common.ExternalProxyAccountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                     `xml:"Prtry,omitempty"`
}

// FrequencyAndMoment1 type definition
type FrequencyAndMoment1 struct {
	Tp     common.Frequency6Code    `xml:"Tp,omitempty"`
	PtInTm common.Exact2NumericText `xml:"PtInTm,omitempty"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id,omitempty"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// DatePeriod2 type definition
type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt,omitempty"`
	ToDt common.ISODate `xml:"ToDt,omitempty"`
}

// SupplementaryData1 type definition
type SupplementaryData1 struct {
	PlcAndNm common.Max350Text           `xml:"PlcAndNm,omitempty"`
	Envlp    *SupplementaryDataEnvelope1 `xml:"Envlp,omitempty"`
}

// DocumentLineType1Choice type definition
type DocumentLineType1Choice struct {
	Cd    common.ExternalDocumentLineType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                     `xml:"Prtry,omitempty"`
}

// TaxAmountAndType1 type definition
type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice              `xml:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// GarnishmentType1 type definition
type GarnishmentType1 struct {
	CdOrPrtry *GarnishmentType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text        `xml:"Issr,omitempty"`
}

// SettlementInstruction7 type definition
type SettlementInstruction7 struct {
	SttlmMtd             common.SettlementMethod1Code                  `xml:"SttlmMtd,omitempty"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// ReferredDocumentInformation7 type definition
type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty"`
	Nb       common.Max35Text           `xml:"Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty"`
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

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty"`
}

// CustomerPaymentStatusReportV10 type definition
type CustomerPaymentStatusReportV10 struct {
	GrpHdr            *GroupHeader86                 `xml:"GrpHdr,omitempty"`
	OrgnlGrpInfAndSts *OriginalGroupHeader17         `xml:"OrgnlGrpInfAndSts,omitempty"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction32 `xml:"OrgnlPmtInfAndSts,omitempty"`
	SplmtryData       []SupplementaryData1           `xml:"SplmtryData,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id,omitempty"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
}

// CurrencyExchange13 type definition
type CurrencyExchange13 struct {
	SrcCcy   common.ActiveCurrencyCode `xml:"SrcCcy,omitempty"`
	TrgtCcy  common.ActiveCurrencyCode `xml:"TrgtCcy,omitempty"`
	XchgRate xsdt.Decimal              `xml:"XchgRate,omitempty"`
	UnitCcy  common.ActiveCurrencyCode `xml:"UnitCcy,omitempty"`
}

// ReferredDocumentType3Choice type definition
type ReferredDocumentType3Choice struct {
	Cd    common.DocumentType6Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty"`
}

// TaxAmount2 type definition
type TaxAmount2 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text              `xml:"Issr,omitempty"`
}

// TrackerRecord1 type definition
type TrackerRecord1 struct {
	Agt          *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
	ChrgBr       common.ChargeBearerType1Code                  `xml:"ChrgBr,omitempty"`
	ChrgsAmt     *ActiveCurrencyAndAmount                      `xml:"ChrgsAmt,omitempty"`
	XchgRateData *CurrencyExchange13                           `xml:"XchgRateData,omitempty"`
}

// AmendmentInformationDetails13 type definition
type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                              `xml:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                                `xml:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   common.Exact2NumericText                      `xml:"OrgnlTrckgDays,omitempty"`
}

// FrequencyPeriod1 type definition
type FrequencyPeriod1 struct {
	Tp        common.Frequency6Code `xml:"Tp,omitempty"`
	CntPerPrd xsdt.Decimal          `xml:"CntPerPrd,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
}

// ClearingSystemIdentification3Choice type definition
type ClearingSystemIdentification3Choice struct {
	Cd    common.ExternalCashClearingSystem1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                       `xml:"Prtry,omitempty"`
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

// TaxAmountType1Choice type definition
type TaxAmountType1Choice struct {
	Cd    common.ExternalTaxAmountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                  `xml:"Prtry,omitempty"`
}

// PersonIdentification13 type definition
type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
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

// FinancialInstitutionIdentification18 type definition
type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier        `xml:"BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                 `xml:"LEI,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// DateAndDateTime2Choice type definition
type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt,omitempty"`
	DtTm common.ISODateTime `xml:"DtTm,omitempty"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CdtDbtInd common.CreditDebitCode             `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                    `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                  `xml:"AddtlInf,omitempty"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id,omitempty"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// CashAccountType2Choice type definition
type CashAccountType2Choice struct {
	Cd    common.ExternalCashAccountType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// StatusReasonInformation12 type definition
type StatusReasonInformation12 struct {
	Orgtr    *PartyIdentification135 `xml:"Orgtr,omitempty"`
	Rsn      *StatusReason6Choice    `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text     `xml:"AddtlInf,omitempty"`
}

// AmountType4Choice type definition
type AmountType4Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// DocumentLineInformation1 type definition
type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id,omitempty"`
	Desc common.Max2048Text            `xml:"Desc,omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id,omitempty"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// TrackerData1 type definition
type TrackerData1 struct {
	ConfdDt   *DateAndDateTime2Choice  `xml:"ConfdDt,omitempty"`
	ConfdAmt  *ActiveCurrencyAndAmount `xml:"ConfdAmt,omitempty"`
	TrckrRcrd []TrackerRecord1         `xml:"TrckrRcrd,omitempty"`
}

// DocumentLineIdentification1 type definition
type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty"`
	Nb     common.Max35Text   `xml:"Nb,omitempty"`
	RltdDt common.ISODate     `xml:"RltdDt,omitempty"`
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

// OtherContact1 type definition
type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp,omitempty"`
	Id      common.Max128Text `xml:"Id,omitempty"`
}

// Charges7 type definition
type Charges7 struct {
	Amt *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    common.ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
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

// AddressType3Choice type definition
type AddressType3Choice struct {
	Cd    common.AddressType2Code  `xml:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"Prtry,omitempty"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// Party40Choice type definition
type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
}

// StatusReason6Choice type definition
type StatusReason6Choice struct {
	Cd    common.ExternalStatusReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// ActiveCurrencyAndAmount type definition
type ActiveCurrencyAndAmount struct {
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr,omitempty"`
	Value xsdt.Decimal              `xml:",chardata"`
}

// CashAccount38 type definition
type CashAccount38 struct {
	Id   *AccountIdentification4Choice       `xml:"Id,omitempty"`
	Tp   *CashAccountType2Choice             `xml:"Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1        `xml:"Prxy,omitempty"`
}

// OriginalPaymentInstruction32 type definition
type OriginalPaymentInstruction32 struct {
	OrgnlPmtInfId common.Max35Text                       `xml:"OrgnlPmtInfId,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText                `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                           `xml:"OrgnlCtrlSum,omitempty"`
	PmtInfSts     common.ExternalPaymentGroupStatus1Code `xml:"PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12            `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5       `xml:"NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransaction105                `xml:"TxInfAndSts,omitempty"`
}

// MandateRelatedInformation14 type definition
type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text               `xml:"MndtId,omitempty"`
	DtOfSgntr     common.ISODate                 `xml:"DtOfSgntr,omitempty"`
	AmdmntInd     xsdt.Boolean                   `xml:"AmdmntInd,omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text             `xml:"ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                 `xml:"FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                 `xml:"FnlColltnDt,omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty"`
	TrckgDays     common.Exact2NumericText       `xml:"TrckgDays,omitempty"`
}

// Frequency36Choice type definition
type Frequency36Choice struct {
	Tp     common.Frequency6Code `xml:"Tp,omitempty"`
	Prd    *FrequencyPeriod1     `xml:"Prd,omitempty"`
	PtInTm *FrequencyAndMoment1  `xml:"PtInTm,omitempty"`
}

// PaymentTransaction105 type definition
type PaymentTransaction105 struct {
	StsId           common.Max35Text                             `xml:"StsId,omitempty"`
	OrgnlInstrId    common.Max35Text                             `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                             `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier                      `xml:"OrgnlUETR,omitempty"`
	TxSts           common.ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation12                  `xml:"StsRsnInf,omitempty"`
	ChrgsInf        []Charges7                                   `xml:"ChrgsInf,omitempty"`
	TrckrData       *TrackerData1                                `xml:"TrckrData,omitempty"`
	AccptncDtTm     common.ISODateTime                           `xml:"AccptncDtTm,omitempty"`
	AcctSvcrRef     common.Max35Text                             `xml:"AcctSvcrRef,omitempty"`
	ClrSysRef       common.Max35Text                             `xml:"ClrSysRef,omitempty"`
	OrgnlTxRef      *OriginalTransactionReference28              `xml:"OrgnlTxRef,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"SplmtryData,omitempty"`
}

// GarnishmentType1Choice type definition
type GarnishmentType1Choice struct {
	Cd    common.ExternalGarnishmentType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// Party38Choice type definition
type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"PrvtId,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId,omitempty"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
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

// EquivalentAmount2 type definition
type EquivalentAmount2 struct {
	Amt      *ActiveOrHistoricCurrencyAndAmount  `xml:"Amt,omitempty"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf,omitempty"`
}

// ProxyAccountIdentification1 type definition
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty"`
	Id common.Max2048Text       `xml:"Id,omitempty"`
}

// PaymentTypeInformation27 type definition
type PaymentTypeInformation27 struct {
	InstrPrty common.Priority2Code        `xml:"InstrPrty,omitempty"`
	ClrChanl  common.ClearingChannel2Code `xml:"ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice       `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice     `xml:"LclInstrm,omitempty"`
	SeqTp     common.SequenceType3Code    `xml:"SeqTp,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice     `xml:"CtgyPurp,omitempty"`
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

// DiscountAmountAndType1 type definition
type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice         `xml:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// OrganisationIdentification29 type definition
type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// NumberOfTransactionsPerStatus5 type definition
type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText                      `xml:"DtldNbOfTxs,omitempty"`
	DtldSts     common.ExternalPaymentTransactionStatus1Code `xml:"DtldSts,omitempty"`
	DtldCtrlSum xsdt.Decimal                                 `xml:"DtldCtrlSum,omitempty"`
}

// CategoryPurpose1Choice type definition
type CategoryPurpose1Choice struct {
	Cd    common.ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// RemittanceInformation16 type definition
type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty"`
}

// ReferredDocumentType4 type definition
type ReferredDocumentType4 struct {
	CdOrPrtry *ReferredDocumentType3Choice `xml:"CdOrPrtry,omitempty"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
}

// PartyIdentification135 type definition
type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress24   `xml:"PstlAdr,omitempty"`
	Id        *Party38Choice     `xml:"Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty"`
	CtctDtls  *Contact4          `xml:"CtctDtls,omitempty"`
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

// TaxPeriod2 type definition
type TaxPeriod2 struct {
	Yr     common.ISODate              `xml:"Yr,omitempty"`
	Tp     common.TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriod2                `xml:"FrToDt,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr,omitempty"`
	Value xsdt.Decimal                        `xml:",chardata"`
}
