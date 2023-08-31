// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

// SecurityIdentification4Choice type definition
type SecurityIdentification4Choice struct {
	ISIN  common.ISINIdentifier             `xml:"ISIN,omitempty"`
	Prtry *AlternateSecurityIdentification2 `xml:"Prtry,omitempty"`
}

// CorporateAction1 type definition
type CorporateAction1 struct {
	Cd    common.Max35Text `xml:"Cd,omitempty"`
	Nb    common.Max35Text `xml:"Nb,omitempty"`
	Prtry common.Max35Text `xml:"Prtry,omitempty"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// MessageIdentification2 type definition
type MessageIdentification2 struct {
	MsgNmId common.Max35Text `xml:"MsgNmId,omitempty"`
	MsgId   common.Max35Text `xml:"MsgId,omitempty"`
}

// BranchAndFinancialInstitutionIdentification4 type definition
type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"FinInstnId"`
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty"`
}

// ProprietaryDate2 type definition
type ProprietaryDate2 struct {
	Tp common.Max35Text      `xml:"Tp"`
	Dt DateAndDateTimeChoice `xml:"Dt"`
}

// ProprietaryPrice2 type definition
type ProprietaryPrice2 struct {
	Tp   common.Max35Text                  `xml:"Tp"`
	Pric ActiveOrHistoricCurrencyAndAmount `xml:"Pric"`
}

// CurrencyExchange5 type definition
type CurrencyExchange5 struct {
	SrcCcy   common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy  common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy,omitempty"`
	UnitCcy  common.ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`
	XchgRate xsdt.Decimal                        `xml:"XchgRate"`
	CtrctId  common.Max35Text                    `xml:"CtrctId,omitempty"`
	QtnDt    common.ISODateTime                  `xml:"QtnDt,omitempty"`
}

// AmountAndCurrencyExchangeDetails4 type definition
type AmountAndCurrencyExchangeDetails4 struct {
	Tp      common.Max35Text                  `xml:"Tp"`
	Amt     ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyXchg *CurrencyExchange5                `xml:"CcyXchg,omitempty"`
}

// TransactionAgents2 type definition
type TransactionAgents2 struct {
	DbtrAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	CdtrAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`
	RcvgAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"RcvgAgt,omitempty"`
	DlvrgAgt   *BranchAndFinancialInstitutionIdentification4 `xml:"DlvrgAgt,omitempty"`
	IssgAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"IssgAgt,omitempty"`
	SttlmPlc   *BranchAndFinancialInstitutionIdentification4 `xml:"SttlmPlc,omitempty"`
	Prtry      []ProprietaryAgent2                           `xml:"Prtry,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
}

// TotalsPerBankTransactionCode2 type definition
type TotalsPerBankTransactionCode2 struct {
	NbOfNtries    common.Max15NumericText       `xml:"NbOfNtries,omitempty"`
	Sum           xsdt.Decimal                  `xml:"Sum,omitempty"`
	TtlNetNtryAmt xsdt.Decimal                  `xml:"TtlNetNtryAmt,omitempty"`
	CdtDbtInd     common.CreditDebitCode        `xml:"CdtDbtInd,omitempty"`
	FcstInd       xsdt.Boolean                  `xml:"FcstInd,omitempty"`
	BkTxCd        BankTransactionCodeStructure4 `xml:"BkTxCd"`
	Avlbty        []CashBalanceAvailability2    `xml:"Avlbty,omitempty"`
}

// CashBalanceAvailability2 type definition
type CashBalanceAvailability2 struct {
	Dt        CashBalanceAvailabilityDate1      `xml:"Dt"`
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
}

// BankTransactionCodeStructure4 type definition
type BankTransactionCodeStructure4 struct {
	Domn  *BankTransactionCodeStructure5            `xml:"Domn,omitempty"`
	Prtry *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// BranchData2 type definition
type BranchData2 struct {
	Id      common.Max35Text  `xml:"Id,omitempty"`
	Nm      common.Max140Text `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress6   `xml:"PstlAdr,omitempty"`
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

// FinancialInstitutionIdentification7 type definition
type FinancialInstitutionIdentification7 struct {
	BIC         common.BICIdentifier                 `xml:"BIC,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// ReferredDocumentType1Choice type definition
type ReferredDocumentType1Choice struct {
	Cd    common.DocumentType5Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// TaxPeriod1 type definition
type TaxPeriod1 struct {
	Yr     common.ISODate              `xml:"Yr,omitempty"`
	Tp     common.TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriodDetails          `xml:"FrToDt,omitempty"`
}

// BankToCustomerStatementV02 type definition
type BankToCustomerStatementV02 struct {
	GrpHdr GroupHeader42       `xml:"GrpHdr"`
	Stmt   []AccountStatement2 `xml:"Stmt"`
}

// CashAccount20 type definition
type CashAccount20 struct {
	Id   AccountIdentification4Choice                  `xml:"Id"`
	Tp   *CashAccountType2                             `xml:"Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode           `xml:"Ccy,omitempty"`
	Nm   common.Max70Text                              `xml:"Nm,omitempty"`
	Ownr *PartyIdentification32                        `xml:"Ownr,omitempty"`
	Svcr *BranchAndFinancialInstitutionIdentification4 `xml:"Svcr,omitempty"`
}

// ImpliedCurrencyAmountRangeChoice type definition
type ImpliedCurrencyAmountRangeChoice struct {
	FrAmt   *AmountRangeBoundary1 `xml:"FrAmt,omitempty"`
	ToAmt   *AmountRangeBoundary1 `xml:"ToAmt,omitempty"`
	FrToAmt *FromToAmountRange    `xml:"FrToAmt,omitempty"`
	EQAmt   xsdt.Decimal          `xml:"EQAmt,omitempty"`
	NEQAmt  xsdt.Decimal          `xml:"NEQAmt,omitempty"`
}

// AmountRangeBoundary1 type definition
type AmountRangeBoundary1 struct {
	BdryAmt xsdt.Decimal `xml:"BdryAmt"`
	Incl    xsdt.Boolean `xml:"Incl"`
}

// EntryDetails1 type definition
type EntryDetails1 struct {
	Btch   *BatchInformation2  `xml:"Btch,omitempty"`
	TxDtls []EntryTransaction2 `xml:"TxDtls,omitempty"`
}

// PersonIdentification5 type definition
type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    common.ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
}

// BalanceType5Choice type definition
type BalanceType5Choice struct {
	Cd    common.BalanceType12Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// ReportEntry2 type definition
type ReportEntry2 struct {
	NtryRef       common.Max35Text                  `xml:"NtryRef,omitempty"`
	Amt           ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd     common.CreditDebitCode            `xml:"CdtDbtInd"`
	RvslInd       xsdt.Boolean                      `xml:"RvslInd,omitempty"`
	Sts           common.EntryStatus2Code           `xml:"Sts"`
	BookgDt       *DateAndDateTimeChoice            `xml:"BookgDt,omitempty"`
	ValDt         *DateAndDateTimeChoice            `xml:"ValDt,omitempty"`
	AcctSvcrRef   common.Max35Text                  `xml:"AcctSvcrRef,omitempty"`
	Avlbty        []CashBalanceAvailability2        `xml:"Avlbty,omitempty"`
	BkTxCd        BankTransactionCodeStructure4     `xml:"BkTxCd"`
	ComssnWvrInd  xsdt.Boolean                      `xml:"ComssnWvrInd,omitempty"`
	AddtlInfInd   *MessageIdentification2           `xml:"AddtlInfInd,omitempty"`
	AmtDtls       *AmountAndCurrencyExchange3       `xml:"AmtDtls,omitempty"`
	Chrgs         []ChargesInformation6             `xml:"Chrgs,omitempty"`
	TechInptChanl *TechnicalInputChannel1Choice     `xml:"TechInptChanl,omitempty"`
	Intrst        []TransactionInterest2            `xml:"Intrst,omitempty"`
	NtryDtls      []EntryDetails1                   `xml:"NtryDtls,omitempty"`
	AddtlNtryInf  common.Max500Text                 `xml:"AddtlNtryInf,omitempty"`
}

// FinancialInstrumentQuantityChoice type definition
type FinancialInstrumentQuantityChoice struct {
	Unit     xsdt.Decimal `xml:"Unit,omitempty"`
	FaceAmt  xsdt.Decimal `xml:"FaceAmt,omitempty"`
	AmtsdVal xsdt.Decimal `xml:"AmtsdVal,omitempty"`
}

// CashAccount16 type definition
type CashAccount16 struct {
	Id  AccountIdentification4Choice        `xml:"Id"`
	Tp  *CashAccountType2                   `xml:"Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"Nm,omitempty"`
}

// OrganisationIdentification4 type definition
type OrganisationIdentification4 struct {
	BICOrBEI common.AnyBICIdentifier              `xml:"BICOrBEI,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// BalanceSubType1Choice type definition
type BalanceSubType1Choice struct {
	Cd    common.ExternalBalanceSubType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                   `xml:"Prtry,omitempty"`
}

// ReportingSource1Choice type definition
type ReportingSource1Choice struct {
	Cd    common.ExternalReportingSource1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// DateAndDateTimeChoice type definition
type DateAndDateTimeChoice struct {
	Dt   common.ISODate     `xml:"Dt,omitempty"`
	DtTm common.ISODateTime `xml:"DtTm,omitempty"`
}

// RemittanceInformation5 type definition
type RemittanceInformation5 struct {
	Ustrd []common.Max140Text                `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty"`
}

// ReturnReason5Choice type definition
type ReturnReason5Choice struct {
	Cd    common.ExternalReturnReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

// FromToAmountRange type definition
type FromToAmountRange struct {
	FrAmt AmountRangeBoundary1 `xml:"FrAmt"`
	ToAmt AmountRangeBoundary1 `xml:"ToAmt"`
}

// BankTransactionCodeStructure5 type definition
type BankTransactionCodeStructure5 struct {
	Cd   common.ExternalBankTransactionDomain1Code `xml:"Cd"`
	Fmly BankTransactionCodeStructure6             `xml:"Fmly"`
}

// ChargesInformation6 type definition
type ChargesInformation6 struct {
	TtlChrgsAndTaxAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"TtlChrgsAndTaxAmt,omitempty"`
	Amt               ActiveOrHistoricCurrencyAndAmount             `xml:"Amt"`
	CdtDbtInd         common.CreditDebitCode                        `xml:"CdtDbtInd,omitempty"`
	Tp                *ChargeType2Choice                            `xml:"Tp,omitempty"`
	Rate              xsdt.Decimal                                  `xml:"Rate,omitempty"`
	Br                common.ChargeBearerType1Code                  `xml:"Br,omitempty"`
	Pty               *BranchAndFinancialInstitutionIdentification4 `xml:"Pty,omitempty"`
	Tax               *TaxCharges2                                  `xml:"Tax,omitempty"`
}

// TaxCharges2 type definition
type TaxCharges2 struct {
	Id   common.Max35Text                   `xml:"Id,omitempty"`
	Rate xsdt.Decimal                       `xml:"Rate,omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    common.ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                               `xml:"Prtry,omitempty"`
}

// AccountStatement2 type definition
type AccountStatement2 struct {
	Id           common.Max35Text          `xml:"Id"`
	ElctrncSeqNb xsdt.Decimal              `xml:"ElctrncSeqNb,omitempty"`
	LglSeqNb     xsdt.Decimal              `xml:"LglSeqNb,omitempty"`
	CreDtTm      common.ISODateTime        `xml:"CreDtTm"`
	FrToDt       *DateTimePeriodDetails    `xml:"FrToDt,omitempty"`
	CpyDplctInd  common.CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`
	RptgSrc      *ReportingSource1Choice   `xml:"RptgSrc,omitempty"`
	Acct         CashAccount20             `xml:"Acct"`
	RltdAcct     *CashAccount16            `xml:"RltdAcct,omitempty"`
	Intrst       []AccountInterest2        `xml:"Intrst,omitempty"`
	Bal          []CashBalance3            `xml:"Bal"`
	TxsSummry    *TotalTransactions2       `xml:"TxsSummry,omitempty"`
	Ntry         []ReportEntry2            `xml:"Ntry,omitempty"`
	AddtlStmtInf common.Max500Text         `xml:"AddtlStmtInf,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// ChargeType2Choice type definition
type ChargeType2Choice struct {
	Cd    common.ChargeType1Code  `xml:"Cd,omitempty"`
	Prtry *GenericIdentification3 `xml:"Prtry,omitempty"`
}

// RemittanceLocation2 type definition
type RemittanceLocation2 struct {
	RmtId             common.Max35Text                     `xml:"RmtId,omitempty"`
	RmtLctnMtd        common.RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr common.Max2048Text                   `xml:"RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    *NameAndAddress10                    `xml:"RmtLctnPstlAdr,omitempty"`
}

// PartyIdentification32 type definition
type PartyIdentification32 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress6    `xml:"PstlAdr,omitempty"`
	Id        *Party6Choice      `xml:"Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2   `xml:"CtctDtls,omitempty"`
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

// NumberAndSumOfTransactions2 type definition
type NumberAndSumOfTransactions2 struct {
	NbOfNtries    common.Max15NumericText `xml:"NbOfNtries,omitempty"`
	Sum           xsdt.Decimal            `xml:"Sum,omitempty"`
	TtlNetNtryAmt xsdt.Decimal            `xml:"TtlNetNtryAmt,omitempty"`
	CdtDbtInd     common.CreditDebitCode  `xml:"CdtDbtInd,omitempty"`
}

// TransactionReferences2 type definition
type TransactionReferences2 struct {
	MsgId       common.Max35Text       `xml:"MsgId,omitempty"`
	AcctSvcrRef common.Max35Text       `xml:"AcctSvcrRef,omitempty"`
	PmtInfId    common.Max35Text       `xml:"PmtInfId,omitempty"`
	InstrId     common.Max35Text       `xml:"InstrId,omitempty"`
	EndToEndId  common.Max35Text       `xml:"EndToEndId,omitempty"`
	TxId        common.Max35Text       `xml:"TxId,omitempty"`
	MndtId      common.Max35Text       `xml:"MndtId,omitempty"`
	ChqNb       common.Max35Text       `xml:"ChqNb,omitempty"`
	ClrSysRef   common.Max35Text       `xml:"ClrSysRef,omitempty"`
	Prtry       *ProprietaryReference1 `xml:"Prtry,omitempty"`
}

// CreditLine2 type definition
type CreditLine2 struct {
	Incl xsdt.Boolean                       `xml:"Incl"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
	Value xsdt.Decimal                        `xml:",chardata"`
}

// CashBalanceAvailabilityDate1 type definition
type CashBalanceAvailabilityDate1 struct {
	NbOfDays common.Max15PlusSignedNumericText `xml:"NbOfDays,omitempty"`
	ActlDt   common.ISODate                    `xml:"ActlDt,omitempty"`
}

// NumberAndSumOfTransactions1 type definition
type NumberAndSumOfTransactions1 struct {
	NbOfNtries common.Max15NumericText `xml:"NbOfNtries,omitempty"`
	Sum        xsdt.Decimal            `xml:"Sum,omitempty"`
}

// AmountAndCurrencyExchangeDetails3 type definition
type AmountAndCurrencyExchangeDetails3 struct {
	Amt     ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyXchg *CurrencyExchange5                `xml:"CcyXchg,omitempty"`
}

// Party6Choice type definition
type Party6Choice struct {
	OrgId  *OrganisationIdentification4 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

// TransactionDates2 type definition
type TransactionDates2 struct {
	AccptncDtTm             common.ISODateTime `xml:"AccptncDtTm,omitempty"`
	TradActvtyCtrctlSttlmDt common.ISODate     `xml:"TradActvtyCtrctlSttlmDt,omitempty"`
	TradDt                  common.ISODate     `xml:"TradDt,omitempty"`
	IntrBkSttlmDt           common.ISODate     `xml:"IntrBkSttlmDt,omitempty"`
	StartDt                 common.ISODate     `xml:"StartDt,omitempty"`
	EndDt                   common.ISODate     `xml:"EndDt,omitempty"`
	TxDtTm                  common.ISODateTime `xml:"TxDtTm,omitempty"`
	Prtry                   []ProprietaryDate2 `xml:"Prtry,omitempty"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   common.Max35Text   `xml:"TaxId,omitempty"`
	RegnId  common.Max35Text   `xml:"RegnId,omitempty"`
	TaxTp   common.Max35Text   `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    common.ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// TaxRecordDetails1 type definition
type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                       `xml:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// CashAccountType2 type definition
type CashAccountType2 struct {
	Cd    common.CashAccountType4Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text            `xml:"Prtry,omitempty"`
}

// AmountAndCurrencyExchange3 type definition
type AmountAndCurrencyExchange3 struct {
	InstdAmt      *AmountAndCurrencyExchangeDetails3  `xml:"InstdAmt,omitempty"`
	TxAmt         *AmountAndCurrencyExchangeDetails3  `xml:"TxAmt,omitempty"`
	CntrValAmt    *AmountAndCurrencyExchangeDetails3  `xml:"CntrValAmt,omitempty"`
	AnncdPstngAmt *AmountAndCurrencyExchangeDetails3  `xml:"AnncdPstngAmt,omitempty"`
	PrtryAmt      []AmountAndCurrencyExchangeDetails4 `xml:"PrtryAmt,omitempty"`
}

// TechnicalInputChannel1Choice type definition
type TechnicalInputChannel1Choice struct {
	Cd    common.ExternalTechnicalInputChannel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
}

// TransactionInterest2 type definition
type TransactionInterest2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
	Tp        *InterestType1Choice              `xml:"Tp,omitempty"`
	Rate      []Rate3                           `xml:"Rate,omitempty"`
	FrToDt    *DateTimePeriodDetails            `xml:"FrToDt,omitempty"`
	Rsn       common.Max35Text                  `xml:"Rsn,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// ProprietaryBankTransactionCodeStructure1 type definition
type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   common.Max35Text `xml:"Cd"`
	Issr common.Max35Text `xml:"Issr,omitempty"`
}

// GroupHeader42 type definition
type GroupHeader42 struct {
	MsgId    common.Max35Text       `xml:"MsgId"`
	CreDtTm  common.ISODateTime     `xml:"CreDtTm"`
	MsgRcpt  *PartyIdentification32 `xml:"MsgRcpt,omitempty"`
	MsgPgntn *Pagination            `xml:"MsgPgntn,omitempty"`
	AddtlInf common.Max500Text      `xml:"AddtlInf,omitempty"`
}

// TransactionParty2 type definition
type TransactionParty2 struct {
	InitgPty  *PartyIdentification32 `xml:"InitgPty,omitempty"`
	Dbtr      *PartyIdentification32 `xml:"Dbtr,omitempty"`
	DbtrAcct  *CashAccount16         `xml:"DbtrAcct,omitempty"`
	UltmtDbtr *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`
	Cdtr      *PartyIdentification32 `xml:"Cdtr,omitempty"`
	CdtrAcct  *CashAccount16         `xml:"CdtrAcct,omitempty"`
	UltmtCdtr *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`
	TradgPty  *PartyIdentification32 `xml:"TradgPty,omitempty"`
	Prtry     []ProprietaryParty2    `xml:"Prtry,omitempty"`
}

// DateAndPlaceOfBirth type definition
type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
}

// GenericIdentification3 type definition
type GenericIdentification3 struct {
	Id   common.Max35Text `xml:"Id"`
	Issr common.Max35Text `xml:"Issr,omitempty"`
}

// ReturnReasonInformation10 type definition
type ReturnReasonInformation10 struct {
	OrgnlBkTxCd *BankTransactionCodeStructure4 `xml:"OrgnlBkTxCd,omitempty"`
	Orgtr       *PartyIdentification32         `xml:"Orgtr,omitempty"`
	Rsn         *ReturnReason5Choice           `xml:"Rsn,omitempty"`
	AddtlInf    []common.Max105Text            `xml:"AddtlInf,omitempty"`
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

// AccountInterest2 type definition
type AccountInterest2 struct {
	Tp     *InterestType1Choice   `xml:"Tp,omitempty"`
	Rate   []Rate3                `xml:"Rate,omitempty"`
	FrToDt *DateTimePeriodDetails `xml:"FrToDt,omitempty"`
	Rsn    common.Max35Text       `xml:"Rsn,omitempty"`
}

// BalanceType12 type definition
type BalanceType12 struct {
	CdOrPrtry BalanceType5Choice     `xml:"CdOrPrtry"`
	SubTp     *BalanceSubType1Choice `xml:"SubTp,omitempty"`
}

// Pagination type definition
type Pagination struct {
	PgNb      common.Max5NumericText `xml:"PgNb"`
	LastPgInd xsdt.Boolean           `xml:"LastPgInd"`
}

// TaxAmount1 type definition
type TaxAmount1 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty"`
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

// CashBalance3 type definition
type CashBalance3 struct {
	Tp        BalanceType12                     `xml:"Tp"`
	CdtLine   *CreditLine2                      `xml:"CdtLine,omitempty"`
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
	Dt        DateAndDateTimeChoice             `xml:"Dt"`
	Avlbty    []CashBalanceAvailability2        `xml:"Avlbty,omitempty"`
}

// AlternateSecurityIdentification2 type definition
type AlternateSecurityIdentification2 struct {
	Tp common.Max35Text `xml:"Tp"`
	Id common.Max35Text `xml:"Id"`
}

// InterestType1Choice type definition
type InterestType1Choice struct {
	Cd    common.InterestType1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// ProprietaryReference1 type definition
type ProprietaryReference1 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Ref common.Max35Text `xml:"Ref"`
}

// ProprietaryAgent2 type definition
type ProprietaryAgent2 struct {
	Tp  common.Max35Text                             `xml:"Tp"`
	Agt BranchAndFinancialInstitutionIdentification4 `xml:"Agt"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// DatePeriodDetails type definition
type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

// DateTimePeriodDetails type definition
type DateTimePeriodDetails struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

// TotalTransactions2 type definition
type TotalTransactions2 struct {
	TtlNtries          *NumberAndSumOfTransactions2    `xml:"TtlNtries,omitempty"`
	TtlCdtNtries       *NumberAndSumOfTransactions1    `xml:"TtlCdtNtries,omitempty"`
	TtlDbtNtries       *NumberAndSumOfTransactions1    `xml:"TtlDbtNtries,omitempty"`
	TtlNtriesPerBkTxCd []TotalsPerBankTransactionCode2 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// Rate3 type definition
type Rate3 struct {
	Tp      RateType4Choice          `xml:"Tp"`
	VldtyRg *CurrencyAndAmountRange2 `xml:"VldtyRg,omitempty"`
}

// CurrencyAndAmountRange2 type definition
type CurrencyAndAmountRange2 struct {
	Amt       ImpliedCurrencyAmountRangeChoice    `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode              `xml:"CdtDbtInd,omitempty"`
	Ccy       common.ActiveOrHistoricCurrencyCode `xml:"Ccy"`
}

// EntryTransaction2 type definition
type EntryTransaction2 struct {
	Refs        *TransactionReferences2        `xml:"Refs,omitempty"`
	AmtDtls     *AmountAndCurrencyExchange3    `xml:"AmtDtls,omitempty"`
	Avlbty      []CashBalanceAvailability2     `xml:"Avlbty,omitempty"`
	BkTxCd      *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`
	Chrgs       []ChargesInformation6          `xml:"Chrgs,omitempty"`
	Intrst      []TransactionInterest2         `xml:"Intrst,omitempty"`
	RltdPties   *TransactionParty2             `xml:"RltdPties,omitempty"`
	RltdAgts    *TransactionAgents2            `xml:"RltdAgts,omitempty"`
	Purp        *Purpose2Choice                `xml:"Purp,omitempty"`
	RltdRmtInf  []RemittanceLocation2          `xml:"RltdRmtInf,omitempty"`
	RmtInf      *RemittanceInformation5        `xml:"RmtInf,omitempty"`
	RltdDts     *TransactionDates2             `xml:"RltdDts,omitempty"`
	RltdPric    *TransactionPrice2Choice       `xml:"RltdPric,omitempty"`
	RltdQties   []TransactionQuantities1Choice `xml:"RltdQties,omitempty"`
	FinInstrmId *SecurityIdentification4Choice `xml:"FinInstrmId,omitempty"`
	Tax         *TaxInformation3               `xml:"Tax,omitempty"`
	RtrInf      *ReturnReasonInformation10     `xml:"RtrInf,omitempty"`
	CorpActn    *CorporateAction1              `xml:"CorpActn,omitempty"`
	SfkpgAcct   *CashAccount16                 `xml:"SfkpgAcct,omitempty"`
	AddtlTxInf  common.Max500Text              `xml:"AddtlTxInf,omitempty"`
}

// NameAndAddress10 type definition
type NameAndAddress10 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress6    `xml:"Adr"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty"`
}

// ProprietaryQuantity1 type definition
type ProprietaryQuantity1 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Qty common.Max35Text `xml:"Qty"`
}

// BankTransactionCodeStructure6 type definition
type BankTransactionCodeStructure6 struct {
	Cd        common.ExternalBankTransactionFamily1Code    `xml:"Cd"`
	SubFmlyCd common.ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

// ProprietaryParty2 type definition
type ProprietaryParty2 struct {
	Tp  common.Max35Text      `xml:"Tp"`
	Pty PartyIdentification32 `xml:"Pty"`
}

// TransactionQuantities1Choice type definition
type TransactionQuantities1Choice struct {
	Qty   *FinancialInstrumentQuantityChoice `xml:"Qty,omitempty"`
	Prtry *ProprietaryQuantity1              `xml:"Prtry,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// TransactionPrice2Choice type definition
type TransactionPrice2Choice struct {
	DealPric *ActiveOrHistoricCurrencyAndAmount `xml:"DealPric,omitempty"`
	Prtry    []ProprietaryPrice2                `xml:"Prtry,omitempty"`
}

// ReferredDocumentInformation3 type definition
type ReferredDocumentInformation3 struct {
	Tp     *ReferredDocumentType2 `xml:"Tp,omitempty"`
	Nb     common.Max35Text       `xml:"Nb,omitempty"`
	RltdDt common.ISODate         `xml:"RltdDt,omitempty"`
}

// ReferredDocumentType2 type definition
type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text            `xml:"Issr,omitempty"`
}

// RateType4Choice type definition
type RateType4Choice struct {
	Pctg xsdt.Decimal     `xml:"Pctg,omitempty"`
	Othr common.Max35Text `xml:"Othr,omitempty"`
}

// BatchInformation2 type definition
type BatchInformation2 struct {
	MsgId     common.Max35Text                   `xml:"MsgId,omitempty"`
	PmtInfId  common.Max35Text                   `xml:"PmtInfId,omitempty"`
	NbOfTxs   common.Max15NumericText            `xml:"NbOfTxs,omitempty"`
	TtlAmt    *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	CdtDbtInd common.CreditDebitCode             `xml:"CdtDbtInd,omitempty"`
}

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty"`
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