// Package pain_002_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_03

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
)

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    common.ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// OriginalGroupInformation20 type definition
type OriginalGroupInformation20 struct {
	OrgnlMsgId    common.Max35Text                   `xml:"OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                   `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm  common.ISODateTime                 `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText            `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                       `xml:"OrgnlCtrlSum,omitempty"`
	GrpSts        common.TransactionGroupStatus3Code `xml:"GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation8         `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3   `xml:"NbOfTxsPerSts,omitempty"`
}

// LocalInstrument2Choice type definition
type LocalInstrument2Choice struct {
	Cd    common.ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// PersonIdentification5 type definition
type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// FinancialInstitutionIdentification7 type definition
type FinancialInstitutionIdentification7 struct {
	BIC         common.BICIdentifier                 `xml:"BIC,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                    `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// StatusReasonInformation8 type definition
type StatusReasonInformation8 struct {
	Orgtr    *PartyIdentification32 `xml:"Orgtr,omitempty"`
	Rsn      *StatusReason6Choice   `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text    `xml:"AddtlInf,omitempty"`
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

// GroupHeader36 type definition
type GroupHeader36 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	InitgPty *PartyIdentification32                        `xml:"InitgPty,omitempty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
	DbtrAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	CdtrAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
}

// EquivalentAmount2 type definition
type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text          `xml:"Issr,omitempty"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                             `xml:"Issr,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    common.ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                               `xml:"Prtry,omitempty"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    common.ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                          `xml:"Prtry,omitempty"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref common.Max35Text        `xml:"Ref,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                          `xml:"Issr,omitempty"`
}

// AmountType3Choice type definition
type AmountType3Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    common.Max35Text                       `xml:"Issr,omitempty"`
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

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text             `xml:"Issr,omitempty"`
}

// ReferredDocumentType1Choice type definition
type ReferredDocumentType1Choice struct {
	Cd    common.DocumentType5Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
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

// BranchData2 type definition
type BranchData2 struct {
	Id      common.Max35Text  `xml:"Id,omitempty"`
	Nm      common.Max140Text `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress6   `xml:"PstlAdr,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    common.ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                         `xml:"Prtry,omitempty"`
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

// Party6Choice type definition
type Party6Choice struct {
	OrgId  *OrganisationIdentification4 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

// DateAndPlaceOfBirth type definition
type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
	Value xsdt.Decimal                        `xml:",chardata"`
}

// PaymentTypeInformation22 type definition
type PaymentTypeInformation22 struct {
	InstrPrty common.Priority2Code        `xml:"InstrPrty,omitempty"`
	ClrChanl  common.ClearingChannel2Code `xml:"ClrChanl,omitempty"`
	SvcLvl    *ServiceLevel8Choice        `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice     `xml:"LclInstrm,omitempty"`
	SeqTp     common.SequenceType1Code    `xml:"SeqTp,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice     `xml:"CtgyPurp,omitempty"`
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

// CategoryPurpose1Choice type definition
type CategoryPurpose1Choice struct {
	Cd    common.ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                    `xml:"Prtry,omitempty"`
}

// NumberOfTransactionsPerStatus3 type definition
type NumberOfTransactionsPerStatus3 struct {
	DtldNbOfTxs common.Max15NumericText                 `xml:"DtldNbOfTxs"`
	DtldSts     common.TransactionIndividualStatus3Code `xml:"DtldSts"`
	DtldCtrlSum xsdt.Decimal                            `xml:"DtldCtrlSum,omitempty"`
}

// RemittanceInformation5 type definition
type RemittanceInformation5 struct {
	Ustrd []common.Max140Text                `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

// ReferredDocumentType2 type definition
type ReferredDocumentType2 struct {
	CdOrPrtry ReferredDocumentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text            `xml:"Issr,omitempty"`
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

// SettlementInformation13 type definition
type SettlementInformation13 struct {
	SttlmMtd             common.SettlementMethod1Code                  `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount16                                `xml:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount16                                `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount16                                `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification4 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount16                                `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// ReferredDocumentInformation3 type definition
type ReferredDocumentInformation3 struct {
	Tp     *ReferredDocumentType2 `xml:"Tp,omitempty"`
	Nb     common.Max35Text       `xml:"Nb,omitempty"`
	RltdDt common.ISODate         `xml:"RltdDt,omitempty"`
}

// OriginalPaymentInformation1 type definition
type OriginalPaymentInformation1 struct {
	OrgnlPmtInfId common.Max35Text                   `xml:"OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText            `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                       `xml:"OrgnlCtrlSum,omitempty"`
	PmtInfSts     common.TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation8         `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3   `xml:"NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransactionInformation25  `xml:"TxInfAndSts,omitempty"`
}

// OriginalTransactionReference13 type definition
type OriginalTransactionReference13 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountType3Choice                            `xml:"Amt,omitempty"`
	IntrBkSttlmDt  common.ISODate                                `xml:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   common.ISODate                                `xml:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    common.ISODate                                `xml:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *PartyIdentification32                        `xml:"CdtrSchmeId,omitempty"`
	SttlmInf       *SettlementInformation13                      `xml:"SttlmInf,omitempty"`
	PmtTpInf       *PaymentTypeInformation22                     `xml:"PmtTpInf,omitempty"`
	PmtMtd         common.PaymentMethod4Code                     `xml:"PmtMtd,omitempty"`
	MndtRltdInf    *MandateRelatedInformation6                   `xml:"MndtRltdInf,omitempty"`
	RmtInf         *RemittanceInformation5                       `xml:"RmtInf,omitempty"`
	UltmtDbtr      *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	Dbtr           *PartyIdentification32                        `xml:"Dbtr,omitempty"`
	DbtrAcct       *CashAccount16                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *CashAccount16                                `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *CashAccount16                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr           *PartyIdentification32                        `xml:"Cdtr,omitempty"`
	CdtrAcct       *CashAccount16                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr      *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
}

// ChargesInformation5 type definition
type ChargesInformation5 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Pty BranchAndFinancialInstitutionIdentification4 `xml:"Pty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier      `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
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

// CustomerPaymentStatusReportV03 type definition
type CustomerPaymentStatusReportV03 struct {
	GrpHdr            GroupHeader36                 `xml:"GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation20    `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInformation1 `xml:"OrgnlPmtInfAndSts,omitempty"`
}

// ClearingSystemIdentification3Choice type definition
type ClearingSystemIdentification3Choice struct {
	Cd    common.ExternalCashClearingSystem1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                       `xml:"Prtry,omitempty"`
}

// CashAccount16 type definition
type CashAccount16 struct {
	Id  AccountIdentification4Choice        `xml:"Id"`
	Tp  *CashAccountType2                   `xml:"Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"Nm,omitempty"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    common.ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    common.DocumentType3Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text         `xml:"Prtry,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    common.ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                       `xml:"Prtry,omitempty"`
}

// StatusReason6Choice type definition
type StatusReason6Choice struct {
	Cd    common.ExternalStatusReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// PaymentTransactionInformation25 type definition
type PaymentTransactionInformation25 struct {
	StsId           common.Max35Text                        `xml:"StsId,omitempty"`
	OrgnlInstrId    common.Max35Text                        `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                        `xml:"OrgnlEndToEndId,omitempty"`
	TxSts           common.TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation8              `xml:"StsRsnInf,omitempty"`
	ChrgsInf        []ChargesInformation5                   `xml:"ChrgsInf,omitempty"`
	AccptncDtTm     common.ISODateTime                      `xml:"AccptncDtTm,omitempty"`
	AcctSvcrRef     common.Max35Text                        `xml:"AcctSvcrRef,omitempty"`
	ClrSysRef       common.Max35Text                        `xml:"ClrSysRef,omitempty"`
	OrgnlTxRef      *OriginalTransactionReference13         `xml:"OrgnlTxRef,omitempty"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}
