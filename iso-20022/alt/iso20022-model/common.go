package iso20022_model

// AccountIdentification4Choice
type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// isValid checks if AccountIdentification4Choice is valid
func (s *AccountIdentification4Choice) isValid() bool {
	valid :=
		s.IBAN.isValid() &&
			s.Othr.isValid()

	return valid

}

// AccountSchemeName1Choice
type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                          `xml:"Prtry,omitempty"`
}

// isValid checks if AccountSchemeName1Choice is valid
func (s *AccountSchemeName1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ActiveCurrencyAndAmount
type ActiveCurrencyAndAmount struct {
}

// isValid checks if ActiveCurrencyAndAmount is valid
func (s *ActiveCurrencyAndAmount) isValid() bool {

	return true

}

// ActiveOrHistoricCurrencyAndAmount
type ActiveOrHistoricCurrencyAndAmount struct {
}

// isValid checks if ActiveOrHistoricCurrencyAndAmount is valid
func (s *ActiveOrHistoricCurrencyAndAmount) isValid() bool {

	return true

}

// AddressType3Choice
type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"Prtry,omitempty"`
}

// isValid checks if AddressType3Choice is valid
func (s *AddressType3Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// AmendmentInformationDetails13
type AmendmentInformationDetails13 struct {
	OrgnlMndtId      Max35Text                                    `xml:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty"`
}

// isValid checks if AmendmentInformationDetails13 is valid
func (s *AmendmentInformationDetails13) isValid() bool {
	valid :=
		s.OrgnlMndtId.isValid() &&
			s.OrgnlCdtrSchmeId.isValid() &&
			s.OrgnlCdtrAgt.isValid() &&
			s.OrgnlCdtrAgtAcct.isValid() &&
			s.OrgnlDbtr.isValid() &&
			s.OrgnlDbtrAcct.isValid() &&
			s.OrgnlDbtrAgt.isValid() &&
			s.OrgnlDbtrAgtAcct.isValid() &&
			s.OrgnlFnlColltnDt.isValid() &&
			s.OrgnlFrqcy.isValid() &&
			s.OrgnlRsn.isValid() &&
			s.OrgnlTrckgDays.isValid()

	return valid

}

// AmountOrRate1Choice
type AmountOrRate1Choice struct {
	Amt  ActiveCurrencyAndAmount `xml:"Amt,omitempty"`
	Rate float64                 `xml:"Rate,omitempty"`
}

// isValid checks if AmountOrRate1Choice is valid
func (s *AmountOrRate1Choice) isValid() bool {
	valid :=
		s.Amt.isValid() &&
			true

	return valid

}

// AmountType4Choice
type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// isValid checks if AmountType4Choice is valid
func (s *AmountType4Choice) isValid() bool {
	valid :=
		s.InstdAmt.isValid() &&
			s.EqvtAmt.isValid()

	return valid

}

// BranchAndFinancialInstitutionIdentification6
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty"`
}

// isValid checks if BranchAndFinancialInstitutionIdentification6 is valid
func (s *BranchAndFinancialInstitutionIdentification6) isValid() bool {
	valid :=
		s.FinInstnId.isValid() &&
			s.BrnchId.isValid()

	return valid

}

// BranchData3
type BranchData3 struct {
	Id      Max35Text       `xml:"Id,omitempty"`
	LEI     LEIIdentifier   `xml:"LEI,omitempty"`
	Nm      Max140Text      `xml:"Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"PstlAdr,omitempty"`
}

// isValid checks if BranchData3 is valid
func (s *BranchData3) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.LEI.isValid() &&
			s.Nm.isValid() &&
			s.PstlAdr.isValid()

	return valid

}

// CancellationReason33Choice
type CancellationReason33Choice struct {
	Cd    ExternalCancellationReason1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty"`
}

// isValid checks if CancellationReason33Choice is valid
func (s *CancellationReason33Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// CancellationStatusReason3Choice
type CancellationStatusReason3Choice struct {
	Cd    ExternalPaymentCancellationRejection1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                 `xml:"Prtry,omitempty"`
}

// isValid checks if CancellationStatusReason3Choice is valid
func (s *CancellationStatusReason3Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// CancellationStatusReason4
type CancellationStatusReason4 struct {
	Orgtr    PartyIdentification135          `xml:"Orgtr,omitempty"`
	Rsn      CancellationStatusReason3Choice `xml:"Rsn,omitempty"`
	AddtlInf Max105Text                      `xml:"AddtlInf,omitempty"`
}

// isValid checks if CancellationStatusReason4 is valid
func (s *CancellationStatusReason4) isValid() bool {
	valid :=
		s.Orgtr.isValid() &&
			s.Rsn.isValid() &&
			s.AddtlInf.isValid()

	return valid

}

// Case5
type Case5 struct {
	Id             Max35Text     `xml:"Id"`
	Cretr          Party40Choice `xml:"Cretr"`
	ReopCaseIndctn bool          `xml:"ReopCaseIndctn,omitempty"`
}

// isValid checks if Case5 is valid
func (s *Case5) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.Cretr.isValid() &&
			true

	return valid

}

// CaseAssignment5
type CaseAssignment5 struct {
	Id      Max35Text     `xml:"Id"`
	Assgnr  Party40Choice `xml:"Assgnr"`
	Assgne  Party40Choice `xml:"Assgne"`
	CreDtTm ISODateTime   `xml:"CreDtTm"`
}

// isValid checks if CaseAssignment5 is valid
func (s *CaseAssignment5) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.Assgnr.isValid() &&
			s.Assgne.isValid() &&
			s.CreDtTm.isValid()

	return valid

}

// CashAccount38
type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"Id"`
	Tp   CashAccountType2Choice       `xml:"Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm   Max70Text                    `xml:"Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"Prxy,omitempty"`
}

// isValid checks if CashAccount38 is valid
func (s *CashAccount38) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.Tp.isValid() &&
			s.Ccy.isValid() &&
			s.Nm.isValid() &&
			s.Prxy.isValid()

	return valid

}

// CashAccountType2Choice
type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// isValid checks if CashAccountType2Choice is valid
func (s *CashAccountType2Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// CategoryPurpose1Choice
type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// isValid checks if CategoryPurpose1Choice is valid
func (s *CategoryPurpose1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// Charges7
type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// isValid checks if Charges7 is valid
func (s *Charges7) isValid() bool {
	valid :=
		s.Amt.isValid() &&
			s.Agt.isValid()

	return valid

}

// ClearingSystemIdentification2Choice
type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                 `xml:"Prtry,omitempty"`
}

// isValid checks if ClearingSystemIdentification2Choice is valid
func (s *ClearingSystemIdentification2Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ClearingSystemIdentification3Choice
type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty"`
}

// isValid checks if ClearingSystemIdentification3Choice is valid
func (s *ClearingSystemIdentification3Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ClearingSystemMemberIdentification2
type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"MmbId"`
}

// isValid checks if ClearingSystemMemberIdentification2 is valid
func (s *ClearingSystemMemberIdentification2) isValid() bool {
	valid :=
		s.ClrSysId.isValid() &&
			s.MmbId.isValid()

	return valid

}

// Compensation2
type Compensation2 struct {
	Amt     ActiveCurrencyAndAmount                      `xml:"Amt"`
	DbtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	CdtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Rsn     CompensationReason1Choice                    `xml:"Rsn"`
}

// isValid checks if Compensation2 is valid
func (s *Compensation2) isValid() bool {
	valid :=
		s.Amt.isValid() &&
			s.DbtrAgt.isValid() &&
			s.CdtrAgt.isValid() &&
			s.Rsn.isValid()

	return valid

}

// CompensationReason1Choice
type CompensationReason1Choice struct {
	Cd    ExternalPaymentCompensationReason1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                              `xml:"Prtry,omitempty"`
}

// isValid checks if CompensationReason1Choice is valid
func (s *CompensationReason1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// Contact4
type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"Dept,omitempty"`
	Othr      OtherContact1               `xml:"Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty"`
}

// isValid checks if Contact4 is valid
func (s *Contact4) isValid() bool {
	valid :=
		s.NmPrfx.isValid() &&
			s.Nm.isValid() &&
			s.PhneNb.isValid() &&
			s.MobNb.isValid() &&
			s.FaxNb.isValid() &&
			s.EmailAdr.isValid() &&
			s.EmailPurp.isValid() &&
			s.JobTitl.isValid() &&
			s.Rspnsblty.isValid() &&
			s.Dept.isValid() &&
			s.Othr.isValid() &&
			s.PrefrdMtd.isValid()

	return valid

}

// CorrectiveGroupInformation1
type CorrectiveGroupInformation1 struct {
	MsgId   Max35Text   `xml:"MsgId"`
	MsgNmId Max35Text   `xml:"MsgNmId"`
	CreDtTm ISODateTime `xml:"CreDtTm,omitempty"`
}

// isValid checks if CorrectiveGroupInformation1 is valid
func (s *CorrectiveGroupInformation1) isValid() bool {
	valid :=
		s.MsgId.isValid() &&
			s.MsgNmId.isValid() &&
			s.CreDtTm.isValid()

	return valid

}

// CreditorReferenceInformation2
type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref Max35Text              `xml:"Ref,omitempty"`
}

// isValid checks if CreditorReferenceInformation2 is valid
func (s *CreditorReferenceInformation2) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Ref.isValid()

	return valid

}

// CreditorReferenceType1Choice
type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty"`
}

// isValid checks if CreditorReferenceType1Choice is valid
func (s *CreditorReferenceType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// CreditorReferenceType2
type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text                    `xml:"Issr,omitempty"`
}

// isValid checks if CreditorReferenceType2 is valid
func (s *CreditorReferenceType2) isValid() bool {
	valid :=
		s.CdOrPrtry.isValid() &&
			s.Issr.isValid()

	return valid

}

// DateAndDateTime2Choice
type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"Dt,omitempty"`
	DtTm ISODateTime `xml:"DtTm,omitempty"`
}

// isValid checks if DateAndDateTime2Choice is valid
func (s *DateAndDateTime2Choice) isValid() bool {
	valid :=
		s.Dt.isValid() &&
			s.DtTm.isValid()

	return valid

}

// DateAndPlaceOfBirth1
type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"BirthDt"`
	PrvcOfBirth Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth CountryCode `xml:"CtryOfBirth"`
}

// isValid checks if DateAndPlaceOfBirth1 is valid
func (s *DateAndPlaceOfBirth1) isValid() bool {
	valid :=
		s.BirthDt.isValid() &&
			s.PrvcOfBirth.isValid() &&
			s.CityOfBirth.isValid() &&
			s.CtryOfBirth.isValid()

	return valid

}

// DatePeriod2
type DatePeriod2 struct {
	FrDt ISODate `xml:"FrDt"`
	ToDt ISODate `xml:"ToDt"`
}

// isValid checks if DatePeriod2 is valid
func (s *DatePeriod2) isValid() bool {
	valid :=
		s.FrDt.isValid() &&
			s.ToDt.isValid()

	return valid

}

// DiscountAmountAndType1
type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// isValid checks if DiscountAmountAndType1 is valid
func (s *DiscountAmountAndType1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Amt.isValid()

	return valid

}

// DiscountAmountType1Choice
type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty"`
}

// isValid checks if DiscountAmountType1Choice is valid
func (s *DiscountAmountType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// Document12
type Document12 struct {
	Tp        DocumentType1Choice    `xml:"Tp"`
	Id        Max35Text              `xml:"Id"`
	IsseDt    DateAndDateTime2Choice `xml:"IsseDt"`
	Nm        Max140Text             `xml:"Nm,omitempty"`
	LangCd    string                 `xml:"LangCd,omitempty"`
	Frmt      DocumentFormat1Choice  `xml:"Frmt"`
	FileNm    Max140Text             `xml:"FileNm,omitempty"`
	DgtlSgntr PartyAndSignature3     `xml:"DgtlSgntr,omitempty"`
	Nclsr     Max10MbBinary          `xml:"Nclsr"`
}

// isValid checks if Document12 is valid
func (s *Document12) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Id.isValid() &&
			s.IsseDt.isValid() &&
			s.Nm.isValid() &&
			true &&
			s.Frmt.isValid() &&
			s.FileNm.isValid() &&
			s.DgtlSgntr.isValid() &&
			s.Nclsr.isValid()

	return valid

}

// DocumentAdjustment1
type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"AddtlInf,omitempty"`
}

// isValid checks if DocumentAdjustment1 is valid
func (s *DocumentAdjustment1) isValid() bool {
	valid :=
		s.Amt.isValid() &&
			s.CdtDbtInd.isValid() &&
			s.Rsn.isValid() &&
			s.AddtlInf.isValid()

	return valid

}

// DocumentFormat1Choice
type DocumentFormat1Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"Cd,omitempty"`
	Prtry GenericIdentification1      `xml:"Prtry,omitempty"`
}

// isValid checks if DocumentFormat1Choice is valid
func (s *DocumentFormat1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// DocumentLineIdentification1
type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"Tp,omitempty"`
	Nb     Max35Text         `xml:"Nb,omitempty"`
	RltdDt ISODate           `xml:"RltdDt,omitempty"`
}

// isValid checks if DocumentLineIdentification1 is valid
func (s *DocumentLineIdentification1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Nb.isValid() &&
			s.RltdDt.isValid()

	return valid

}

// DocumentLineInformation1
type DocumentLineInformation1 struct {
	Id   DocumentLineIdentification1 `xml:"Id"`
	Desc Max2048Text                 `xml:"Desc,omitempty"`
	Amt  RemittanceAmount3           `xml:"Amt,omitempty"`
}

// isValid checks if DocumentLineInformation1 is valid
func (s *DocumentLineInformation1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.Desc.isValid() &&
			s.Amt.isValid()

	return valid

}

// DocumentLineType1
type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text               `xml:"Issr,omitempty"`
}

// isValid checks if DocumentLineType1 is valid
func (s *DocumentLineType1) isValid() bool {
	valid :=
		s.CdOrPrtry.isValid() &&
			s.Issr.isValid()

	return valid

}

// DocumentLineType1Choice
type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty"`
}

// isValid checks if DocumentLineType1Choice is valid
func (s *DocumentLineType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// DocumentType1Choice
type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"Cd,omitempty"`
	Prtry GenericIdentification1    `xml:"Prtry,omitempty"`
}

// isValid checks if DocumentType1Choice is valid
func (s *DocumentType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// EquivalentAmount2
type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"CcyOfTrf"`
}

// isValid checks if EquivalentAmount2 is valid
func (s *EquivalentAmount2) isValid() bool {
	valid :=
		s.Amt.isValid() &&
			s.CcyOfTrf.isValid()

	return valid

}

// FinancialIdentificationSchemeName1Choice
type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                       `xml:"Prtry,omitempty"`
}

// isValid checks if FinancialIdentificationSchemeName1Choice is valid
func (s *FinancialIdentificationSchemeName1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// FinancialInstitutionIdentification18
type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"LEI,omitempty"`
	Nm          Max140Text                          `xml:"Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// isValid checks if FinancialInstitutionIdentification18 is valid
func (s *FinancialInstitutionIdentification18) isValid() bool {
	valid :=
		s.BICFI.isValid() &&
			s.ClrSysMmbId.isValid() &&
			s.LEI.isValid() &&
			s.Nm.isValid() &&
			s.PstlAdr.isValid() &&
			s.Othr.isValid()

	return valid

}

// Frequency36Choice
type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp,omitempty"`
	Prd    FrequencyPeriod1    `xml:"Prd,omitempty"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm,omitempty"`
}

// isValid checks if Frequency36Choice is valid
func (s *Frequency36Choice) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Prd.isValid() &&
			s.PtInTm.isValid()

	return valid

}

// FrequencyAndMoment1
type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm"`
}

// isValid checks if FrequencyAndMoment1 is valid
func (s *FrequencyAndMoment1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.PtInTm.isValid()

	return valid

}

// FrequencyPeriod1
type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

// isValid checks if FrequencyPeriod1 is valid
func (s *FrequencyPeriod1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			true

	return valid

}

// Garnishment3
type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"Tp"`
	Grnshee           PartyIdentification135            `xml:"Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                        `xml:"RefNb,omitempty"`
	Dt                ISODate                           `xml:"Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"MplyeeTermntnInd,omitempty"`
}

// isValid checks if Garnishment3 is valid
func (s *Garnishment3) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Grnshee.isValid() &&
			s.GrnshmtAdmstr.isValid() &&
			s.RefNb.isValid() &&
			s.Dt.isValid() &&
			s.RmtdAmt.isValid()

	return valid

}

// GarnishmentType1
type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text              `xml:"Issr,omitempty"`
}

// isValid checks if GarnishmentType1 is valid
func (s *GarnishmentType1) isValid() bool {
	valid :=
		s.CdOrPrtry.isValid() &&
			s.Issr.isValid()

	return valid

}

// GarnishmentType1Choice
type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// isValid checks if GarnishmentType1Choice is valid
func (s *GarnishmentType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// GenericAccountIdentification1
type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"Issr,omitempty"`
}

// isValid checks if GenericAccountIdentification1 is valid
func (s *GenericAccountIdentification1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.SchmeNm.isValid() &&
			s.Issr.isValid()

	return valid

}

// GenericFinancialIdentification1
type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"Issr,omitempty"`
}

// isValid checks if GenericFinancialIdentification1 is valid
func (s *GenericFinancialIdentification1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.SchmeNm.isValid() &&
			s.Issr.isValid()

	return valid

}

// GenericIdentification1
type GenericIdentification1 struct {
	Id      Max35Text `xml:"Id"`
	SchmeNm Max35Text `xml:"SchmeNm,omitempty"`
	Issr    Max35Text `xml:"Issr,omitempty"`
}

// isValid checks if GenericIdentification1 is valid
func (s *GenericIdentification1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.SchmeNm.isValid() &&
			s.Issr.isValid()

	return valid

}

// GenericIdentification30
type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id"`
	Issr    Max35Text              `xml:"Issr"`
	SchmeNm Max35Text              `xml:"SchmeNm,omitempty"`
}

// isValid checks if GenericIdentification30 is valid
func (s *GenericIdentification30) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.Issr.isValid() &&
			s.SchmeNm.isValid()

	return valid

}

// GenericOrganisationIdentification1
type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"Issr,omitempty"`
}

// isValid checks if GenericOrganisationIdentification1 is valid
func (s *GenericOrganisationIdentification1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.SchmeNm.isValid() &&
			s.Issr.isValid()

	return valid

}

// GenericPersonIdentification1
type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"Issr,omitempty"`
}

// isValid checks if GenericPersonIdentification1 is valid
func (s *GenericPersonIdentification1) isValid() bool {
	valid :=
		s.Id.isValid() &&
			s.SchmeNm.isValid() &&
			s.Issr.isValid()

	return valid

}

// LocalInstrument2Choice
type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// isValid checks if LocalInstrument2Choice is valid
func (s *LocalInstrument2Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// MandateRelatedInformation14
type MandateRelatedInformation14 struct {
	MndtId        Max35Text                     `xml:"MndtId,omitempty"`
	DtOfSgntr     ISODate                       `xml:"DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                   `xml:"ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                       `xml:"FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                       `xml:"FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"Rsn,omitempty"`
	TrckgDays     Exact2NumericText             `xml:"TrckgDays,omitempty"`
}

// isValid checks if MandateRelatedInformation14 is valid
func (s *MandateRelatedInformation14) isValid() bool {
	valid :=
		s.MndtId.isValid() &&
			s.DtOfSgntr.isValid() &&
			true &&
			s.AmdmntInfDtls.isValid() &&
			s.ElctrncSgntr.isValid() &&
			s.FrstColltnDt.isValid() &&
			s.FnlColltnDt.isValid() &&
			s.Frqcy.isValid() &&
			s.Rsn.isValid() &&
			s.TrckgDays.isValid()

	return valid

}

// MandateSetupReason1Choice
type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd,omitempty"`
	Prtry Max70Text                       `xml:"Prtry,omitempty"`
}

// isValid checks if MandateSetupReason1Choice is valid
func (s *MandateSetupReason1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ModificationStatusReason1Choice
type ModificationStatusReason1Choice struct {
	Cd    ExternalPaymentModificationRejection1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                 `xml:"Prtry,omitempty"`
}

// isValid checks if ModificationStatusReason1Choice is valid
func (s *ModificationStatusReason1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// NameAndAddress16
type NameAndAddress16 struct {
	Nm  Max140Text      `xml:"Nm"`
	Adr PostalAddress24 `xml:"Adr"`
}

// isValid checks if NameAndAddress16 is valid
func (s *NameAndAddress16) isValid() bool {
	valid :=
		s.Nm.isValid() &&
			s.Adr.isValid()

	return valid

}

// NumberOfTransactionsPerStatus5
type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs Max15NumericText                      `xml:"DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                               `xml:"DtldCtrlSum,omitempty"`
}

// isValid checks if NumberOfTransactionsPerStatus5 is valid
func (s *NumberOfTransactionsPerStatus5) isValid() bool {
	valid :=
		s.DtldNbOfTxs.isValid() &&
			s.DtldSts.isValid() &&
			true

	return valid

}

// OrganisationIdentification29
type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier            `xml:"AnyBIC,omitempty"`
	LEI    LEIIdentifier                      `xml:"LEI,omitempty"`
	Othr   GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// isValid checks if OrganisationIdentification29 is valid
func (s *OrganisationIdentification29) isValid() bool {
	valid :=
		s.AnyBIC.isValid() &&
			s.LEI.isValid() &&
			s.Othr.isValid()

	return valid

}

// OrganisationIdentificationSchemeName1Choice
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                               `xml:"Prtry,omitempty"`
}

// isValid checks if OrganisationIdentificationSchemeName1Choice is valid
func (s *OrganisationIdentificationSchemeName1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// OriginalGroupInformation29
type OriginalGroupInformation29 struct {
	OrgnlMsgId   Max35Text   `xml:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

// isValid checks if OriginalGroupInformation29 is valid
func (s *OriginalGroupInformation29) isValid() bool {
	valid :=
		s.OrgnlMsgId.isValid() &&
			s.OrgnlMsgNmId.isValid() &&
			s.OrgnlCreDtTm.isValid()

	return valid

}

// OriginalTransactionReference28
type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty"`
}

// isValid checks if OriginalTransactionReference28 is valid
func (s *OriginalTransactionReference28) isValid() bool {
	valid :=
		s.IntrBkSttlmAmt.isValid() &&
			s.Amt.isValid() &&
			s.IntrBkSttlmDt.isValid() &&
			s.ReqdColltnDt.isValid() &&
			s.ReqdExctnDt.isValid() &&
			s.CdtrSchmeId.isValid() &&
			s.SttlmInf.isValid() &&
			s.PmtTpInf.isValid() &&
			s.PmtMtd.isValid() &&
			s.MndtRltdInf.isValid() &&
			s.RmtInf.isValid() &&
			s.UltmtDbtr.isValid() &&
			s.Dbtr.isValid() &&
			s.DbtrAcct.isValid() &&
			s.DbtrAgt.isValid() &&
			s.DbtrAgtAcct.isValid() &&
			s.CdtrAgt.isValid() &&
			s.CdtrAgtAcct.isValid() &&
			s.Cdtr.isValid() &&
			s.CdtrAcct.isValid() &&
			s.UltmtCdtr.isValid() &&
			s.Purp.isValid()

	return valid

}

// OtherContact1
type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"ChanlTp"`
	Id      Max128Text `xml:"Id,omitempty"`
}

// isValid checks if OtherContact1 is valid
func (s *OtherContact1) isValid() bool {
	valid :=
		s.ChanlTp.isValid() &&
			s.Id.isValid()

	return valid

}

// Party38Choice
type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"PrvtId,omitempty"`
}

// isValid checks if Party38Choice is valid
func (s *Party38Choice) isValid() bool {
	valid :=
		s.OrgId.isValid() &&
			s.PrvtId.isValid()

	return valid

}

// Party40Choice
type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty,omitempty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
}

// isValid checks if Party40Choice is valid
func (s *Party40Choice) isValid() bool {
	valid :=
		s.Pty.isValid() &&
			s.Agt.isValid()

	return valid

}

// PartyAndSignature3
type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"Pty"`
	Sgntr SkipPayload            `xml:"Sgntr"`
}

// isValid checks if PartyAndSignature3 is valid
func (s *PartyAndSignature3) isValid() bool {
	valid :=
		s.Pty.isValid() &&
			s.Sgntr.isValid()

	return valid

}

// PartyIdentification135
type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"Id,omitempty"`
	CtryOfRes CountryCode     `xml:"CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"CtctDtls,omitempty"`
}

// isValid checks if PartyIdentification135 is valid
func (s *PartyIdentification135) isValid() bool {
	valid :=
		s.Nm.isValid() &&
			s.PstlAdr.isValid() &&
			s.Id.isValid() &&
			s.CtryOfRes.isValid() &&
			s.CtctDtls.isValid()

	return valid

}

// PaymentCancellationReason5
type PaymentCancellationReason5 struct {
	Orgtr    PartyIdentification135     `xml:"Orgtr,omitempty"`
	Rsn      CancellationReason33Choice `xml:"Rsn,omitempty"`
	AddtlInf Max105Text                 `xml:"AddtlInf,omitempty"`
}

// isValid checks if PaymentCancellationReason5 is valid
func (s *PaymentCancellationReason5) isValid() bool {
	valid :=
		s.Orgtr.isValid() &&
			s.Rsn.isValid() &&
			s.AddtlInf.isValid()

	return valid

}

// PaymentCondition1
type PaymentCondition1 struct {
	AmtModAllwd   bool                `xml:"AmtModAllwd"`
	EarlyPmtAllwd bool                `xml:"EarlyPmtAllwd"`
	DelyPnlty     Max140Text          `xml:"DelyPnlty,omitempty"`
	ImdtPmtRbt    AmountOrRate1Choice `xml:"ImdtPmtRbt,omitempty"`
	GrntedPmtReqd bool                `xml:"GrntedPmtReqd"`
}

// isValid checks if PaymentCondition1 is valid
func (s *PaymentCondition1) isValid() bool {
	valid :=
		s.DelyPnlty.isValid() &&
			s.ImdtPmtRbt.isValid() &&
			true

	return valid

}

// PaymentTypeInformation26
type PaymentTypeInformation26 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// isValid checks if PaymentTypeInformation26 is valid
func (s *PaymentTypeInformation26) isValid() bool {
	valid :=
		s.InstrPrty.isValid() &&
			s.SvcLvl.isValid() &&
			s.LclInstrm.isValid() &&
			s.CtgyPurp.isValid()

	return valid

}

// PaymentTypeInformation27
type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty"`
	SvcLvl    ServiceLevel8Choice    `xml:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// isValid checks if PaymentTypeInformation27 is valid
func (s *PaymentTypeInformation27) isValid() bool {
	valid :=
		s.InstrPrty.isValid() &&
			s.ClrChanl.isValid() &&
			s.SvcLvl.isValid() &&
			s.LclInstrm.isValid() &&
			s.SeqTp.isValid() &&
			s.CtgyPurp.isValid()

	return valid

}

// PersonIdentification13
type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1         `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// isValid checks if PersonIdentification13 is valid
func (s *PersonIdentification13) isValid() bool {
	valid :=
		s.DtAndPlcOfBirth.isValid() &&
			s.Othr.isValid()

	return valid

}

// PersonIdentificationSchemeName1Choice
type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                         `xml:"Prtry,omitempty"`
}

// isValid checks if PersonIdentificationSchemeName1Choice is valid
func (s *PersonIdentificationSchemeName1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// PostalAddress24
type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"AdrTp,omitempty"`
	Dept        Max70Text          `xml:"Dept,omitempty"`
	SubDept     Max70Text          `xml:"SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"BldgNm,omitempty"`
	Flr         Max70Text          `xml:"Flr,omitempty"`
	PstBx       Max16Text          `xml:"PstBx,omitempty"`
	Room        Max70Text          `xml:"Room,omitempty"`
	PstCd       Max16Text          `xml:"PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"Ctry,omitempty"`
	AdrLine     Max70Text          `xml:"AdrLine,omitempty"`
}

// isValid checks if PostalAddress24 is valid
func (s *PostalAddress24) isValid() bool {
	valid :=
		s.AdrTp.isValid() &&
			s.Dept.isValid() &&
			s.SubDept.isValid() &&
			s.StrtNm.isValid() &&
			s.BldgNb.isValid() &&
			s.BldgNm.isValid() &&
			s.Flr.isValid() &&
			s.PstBx.isValid() &&
			s.Room.isValid() &&
			s.PstCd.isValid() &&
			s.TwnNm.isValid() &&
			s.TwnLctnNm.isValid() &&
			s.DstrctNm.isValid() &&
			s.CtrySubDvsn.isValid() &&
			s.Ctry.isValid() &&
			s.AdrLine.isValid()

	return valid

}

// ProxyAccountIdentification1
type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty"`
	Id Max2048Text             `xml:"Id"`
}

// isValid checks if ProxyAccountIdentification1 is valid
func (s *ProxyAccountIdentification1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Id.isValid()

	return valid

}

// ProxyAccountType1Choice
type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty"`
}

// isValid checks if ProxyAccountType1Choice is valid
func (s *ProxyAccountType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// Purpose2Choice
type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry Max35Text            `xml:"Prtry,omitempty"`
}

// isValid checks if Purpose2Choice is valid
func (s *Purpose2Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ReferredDocumentInformation7
type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4    `xml:"Tp,omitempty"`
	Nb       Max35Text                `xml:"Nb,omitempty"`
	RltdDt   ISODate                  `xml:"RltdDt,omitempty"`
	LineDtls DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

// isValid checks if ReferredDocumentInformation7 is valid
func (s *ReferredDocumentInformation7) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Nb.isValid() &&
			s.RltdDt.isValid() &&
			s.LineDtls.isValid()

	return valid

}

// ReferredDocumentType3Choice
type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty"`
}

// isValid checks if ReferredDocumentType3Choice is valid
func (s *ReferredDocumentType3Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// ReferredDocumentType4
type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      Max35Text                   `xml:"Issr,omitempty"`
}

// isValid checks if ReferredDocumentType4 is valid
func (s *ReferredDocumentType4) isValid() bool {
	valid :=
		s.CdOrPrtry.isValid() &&
			s.Issr.isValid()

	return valid

}

// RemittanceAmount2
type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      DiscountAmountAndType1            `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            TaxAmountAndType1                 `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn DocumentAdjustment1               `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// isValid checks if RemittanceAmount2 is valid
func (s *RemittanceAmount2) isValid() bool {
	valid :=
		s.DuePyblAmt.isValid() &&
			s.DscntApldAmt.isValid() &&
			s.CdtNoteAmt.isValid() &&
			s.TaxAmt.isValid() &&
			s.AdjstmntAmtAndRsn.isValid() &&
			s.RmtdAmt.isValid()

	return valid

}

// RemittanceAmount3
type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      DiscountAmountAndType1            `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            TaxAmountAndType1                 `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn DocumentAdjustment1               `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// isValid checks if RemittanceAmount3 is valid
func (s *RemittanceAmount3) isValid() bool {
	valid :=
		s.DuePyblAmt.isValid() &&
			s.DscntApldAmt.isValid() &&
			s.CdtNoteAmt.isValid() &&
			s.TaxAmt.isValid() &&
			s.AdjstmntAmtAndRsn.isValid() &&
			s.RmtdAmt.isValid()

	return valid

}

// RemittanceInformation16
type RemittanceInformation16 struct {
	Ustrd Max140Text                        `xml:"Ustrd,omitempty"`
	Strd  StructuredRemittanceInformation16 `xml:"Strd,omitempty"`
}

// isValid checks if RemittanceInformation16 is valid
func (s *RemittanceInformation16) isValid() bool {
	valid :=
		s.Ustrd.isValid() &&
			s.Strd.isValid()

	return valid

}

// ResolutionData1
type ResolutionData1 struct {
	EndToEndId     Max35Text                         `xml:"EndToEndId,omitempty"`
	TxId           Max35Text                         `xml:"TxId,omitempty"`
	UETR           UUIDv4Identifier                  `xml:"UETR,omitempty"`
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt  ISODate                           `xml:"IntrBkSttlmDt,omitempty"`
	ClrChanl       ClearingChannel2Code              `xml:"ClrChanl,omitempty"`
	Compstn        Compensation2                     `xml:"Compstn,omitempty"`
	Chrgs          Charges7                          `xml:"Chrgs,omitempty"`
}

// isValid checks if ResolutionData1 is valid
func (s *ResolutionData1) isValid() bool {
	valid :=
		s.EndToEndId.isValid() &&
			s.TxId.isValid() &&
			s.UETR.isValid() &&
			s.IntrBkSttlmAmt.isValid() &&
			s.IntrBkSttlmDt.isValid() &&
			s.ClrChanl.isValid() &&
			s.Compstn.isValid() &&
			s.Chrgs.isValid()

	return valid

}

// ServiceLevel8Choice
type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty"`
}

// isValid checks if ServiceLevel8Choice is valid
func (s *ServiceLevel8Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// SettlementInstruction7
type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"SttlmMtd"`
	SttlmAcct            CashAccount38                                `xml:"SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// isValid checks if SettlementInstruction7 is valid
func (s *SettlementInstruction7) isValid() bool {
	valid :=
		s.SttlmMtd.isValid() &&
			s.SttlmAcct.isValid() &&
			s.ClrSys.isValid() &&
			s.InstgRmbrsmntAgt.isValid() &&
			s.InstgRmbrsmntAgtAcct.isValid() &&
			s.InstdRmbrsmntAgt.isValid() &&
			s.InstdRmbrsmntAgtAcct.isValid() &&
			s.ThrdRmbrsmntAgt.isValid() &&
			s.ThrdRmbrsmntAgtAcct.isValid()

	return valid

}

// SkipPayload
type SkipPayload struct {
	Item string `xml:",any"`
}

// isValid checks if SkipPayload is valid
func (s *SkipPayload) isValid() bool {
	valid :=
		true

	return valid

}

// StatusReason6Choice
type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty"`
}

// isValid checks if StatusReason6Choice is valid
func (s *StatusReason6Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// StatusReasonInformation12
type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"Orgtr,omitempty"`
	Rsn      StatusReason6Choice    `xml:"Rsn,omitempty"`
	AddtlInf Max105Text             `xml:"AddtlInf,omitempty"`
}

// isValid checks if StatusReasonInformation12 is valid
func (s *StatusReasonInformation12) isValid() bool {
	valid :=
		s.Orgtr.isValid() &&
			s.Rsn.isValid() &&
			s.AddtlInf.isValid()

	return valid

}

// StructuredRemittanceInformation16
type StructuredRemittanceInformation16 struct {
	RfrdDocInf  ReferredDocumentInformation7  `xml:"RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2             `xml:"RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135        `xml:"Invcr,omitempty"`
	Invcee      PartyIdentification135        `xml:"Invcee,omitempty"`
	TaxRmt      TaxInformation7               `xml:"TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                  `xml:"GrnshmtRmt,omitempty"`
	AddtlRmtInf Max140Text                    `xml:"AddtlRmtInf,omitempty"`
}

// isValid checks if StructuredRemittanceInformation16 is valid
func (s *StructuredRemittanceInformation16) isValid() bool {
	valid :=
		s.RfrdDocInf.isValid() &&
			s.RfrdDocAmt.isValid() &&
			s.CdtrRefInf.isValid() &&
			s.Invcr.isValid() &&
			s.Invcee.isValid() &&
			s.TaxRmt.isValid() &&
			s.GrnshmtRmt.isValid() &&
			s.AddtlRmtInf.isValid()

	return valid

}

// SupplementaryData1
type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

// isValid checks if SupplementaryData1 is valid
func (s *SupplementaryData1) isValid() bool {
	valid :=
		s.PlcAndNm.isValid() &&
			s.Envlp.isValid()

	return valid

}

// SupplementaryDataEnvelope1
type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// isValid checks if SupplementaryDataEnvelope1 is valid
func (s *SupplementaryDataEnvelope1) isValid() bool {
	valid :=
		true

	return valid

}

// TaxAmount2
type TaxAmount2 struct {
	Rate         float64                           `xml:"Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         TaxRecordDetails2                 `xml:"Dtls,omitempty"`
}

// isValid checks if TaxAmount2 is valid
func (s *TaxAmount2) isValid() bool {
	valid :=
		true &&
			s.TaxblBaseAmt.isValid() &&
			s.TtlAmt.isValid() &&
			s.Dtls.isValid()

	return valid

}

// TaxAmountAndType1
type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// isValid checks if TaxAmountAndType1 is valid
func (s *TaxAmountAndType1) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Amt.isValid()

	return valid

}

// TaxAmountType1Choice
type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                  `xml:"Prtry,omitempty"`
}

// isValid checks if TaxAmountType1Choice is valid
func (s *TaxAmountType1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// TaxAuthorisation1
type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"Titl,omitempty"`
	Nm   Max140Text `xml:"Nm,omitempty"`
}

// isValid checks if TaxAuthorisation1 is valid
func (s *TaxAuthorisation1) isValid() bool {
	valid :=
		s.Titl.isValid() &&
			s.Nm.isValid()

	return valid

}

// TaxInformation7
type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty"`
	Rcrd            TaxRecord2                        `xml:"Rcrd,omitempty"`
}

// isValid checks if TaxInformation7 is valid
func (s *TaxInformation7) isValid() bool {
	valid :=
		s.Cdtr.isValid() &&
			s.Dbtr.isValid() &&
			s.UltmtDbtr.isValid() &&
			s.AdmstnZone.isValid() &&
			s.RefNb.isValid() &&
			s.Mtd.isValid() &&
			s.TtlTaxblBaseAmt.isValid() &&
			s.TtlTaxAmt.isValid() &&
			s.Dt.isValid() &&
			true &&
			s.Rcrd.isValid()

	return valid

}

// TaxParty1
type TaxParty1 struct {
	TaxId  Max35Text `xml:"TaxId,omitempty"`
	RegnId Max35Text `xml:"RegnId,omitempty"`
	TaxTp  Max35Text `xml:"TaxTp,omitempty"`
}

// isValid checks if TaxParty1 is valid
func (s *TaxParty1) isValid() bool {
	valid :=
		s.TaxId.isValid() &&
			s.RegnId.isValid() &&
			s.TaxTp.isValid()

	return valid

}

// TaxParty2
type TaxParty2 struct {
	TaxId   Max35Text         `xml:"TaxId,omitempty"`
	RegnId  Max35Text         `xml:"RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// isValid checks if TaxParty2 is valid
func (s *TaxParty2) isValid() bool {
	valid :=
		s.TaxId.isValid() &&
			s.RegnId.isValid() &&
			s.TaxTp.isValid() &&
			s.Authstn.isValid()

	return valid

}

// TaxPeriod2
type TaxPeriod2 struct {
	Yr     ISODate              `xml:"Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty"`
}

// isValid checks if TaxPeriod2 is valid
func (s *TaxPeriod2) isValid() bool {
	valid :=
		s.Yr.isValid() &&
			s.Tp.isValid() &&
			s.FrToDt.isValid()

	return valid

}

// TaxRecord2
type TaxRecord2 struct {
	Tp       Max35Text  `xml:"Tp,omitempty"`
	Ctgy     Max35Text  `xml:"Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"FrmsCd,omitempty"`
	Prd      TaxPeriod2 `xml:"Prd,omitempty"`
	TaxAmt   TaxAmount2 `xml:"TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"AddtlInf,omitempty"`
}

// isValid checks if TaxRecord2 is valid
func (s *TaxRecord2) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Ctgy.isValid() &&
			s.CtgyDtls.isValid() &&
			s.DbtrSts.isValid() &&
			s.CertId.isValid() &&
			s.FrmsCd.isValid() &&
			s.Prd.isValid() &&
			s.TaxAmt.isValid() &&
			s.AddtlInf.isValid()

	return valid

}

// TaxRecordDetails2
type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// isValid checks if TaxRecordDetails2 is valid
func (s *TaxRecordDetails2) isValid() bool {
	valid :=
		s.Prd.isValid() &&
			s.Amt.isValid()

	return valid

}
