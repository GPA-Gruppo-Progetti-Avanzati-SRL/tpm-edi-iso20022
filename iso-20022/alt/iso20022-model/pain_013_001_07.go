package iso20022_model

import (
	"bytes"
	"encoding/xml"
)

// Document of Pain_013_001_07 ISO20022 message
type Pain_013_001_07_Document struct {
	XMLName          xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07 Document"`
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV07 `xml:"CdtrPmtActvtnReq"`
}

func (d *Pain_013_001_07_Document) isValid() bool {
	return true
}

/* Snippet mnanually added */
func (d *Pain_013_001_07_Document) ToXML() ([]byte, error) {
	w := &bytes.Buffer{}
	w.Write([]byte(xml.Header))

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	err := enc.Encode(d)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// Cheque11
type Cheque11 struct {
	ChqTp       ChequeType2Code             `xml:"ChqTp,omitempty"`
	ChqNb       Max35Text                   `xml:"ChqNb,omitempty"`
	ChqFr       NameAndAddress16            `xml:"ChqFr,omitempty"`
	DlvryMtd    ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`
	DlvrTo      NameAndAddress16            `xml:"DlvrTo,omitempty"`
	InstrPrty   Priority2Code               `xml:"InstrPrty,omitempty"`
	ChqMtrtyDt  ISODate                     `xml:"ChqMtrtyDt,omitempty"`
	FrmsCd      Max35Text                   `xml:"FrmsCd,omitempty"`
	MemoFld     Max35Text                   `xml:"MemoFld,omitempty"`
	RgnlClrZone Max35Text                   `xml:"RgnlClrZone,omitempty"`
	PrtLctn     Max35Text                   `xml:"PrtLctn,omitempty"`
	Sgntr       Max70Text                   `xml:"Sgntr,omitempty"`
}

// isValid checks if Cheque11 is valid
func (s *Cheque11) isValid() bool {
	valid :=
		s.ChqTp.isValid() &&
			s.ChqNb.isValid() &&
			s.ChqFr.isValid() &&
			s.DlvryMtd.isValid() &&
			s.DlvrTo.isValid() &&
			s.InstrPrty.isValid() &&
			s.ChqMtrtyDt.isValid() &&
			s.FrmsCd.isValid() &&
			s.MemoFld.isValid() &&
			s.RgnlClrZone.isValid() &&
			s.PrtLctn.isValid() &&
			s.Sgntr.isValid()

	return valid

}

// ChequeDeliveryMethod1Choice
type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"Cd,omitempty"`
	Prtry Max35Text           `xml:"Prtry,omitempty"`
}

// isValid checks if ChequeDeliveryMethod1Choice is valid
func (s *ChequeDeliveryMethod1Choice) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.Prtry.isValid()

	return valid

}

// CreditTransferTransaction35
type CreditTransferTransaction35 struct {
	PmtId           PaymentIdentification6                       `xml:"PmtId"`
	PmtTpInf        PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtCond         PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Amt             AmountType4Choice                            `xml:"Amt"`
	ChrgBr          ChargeBearerType1Code                        `xml:"ChrgBr"`
	ChqInstr        Cheque11                                     `xml:"ChqInstr,omitempty"`
	UltmtDbtr       PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2      BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3      BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Cdtr            PartyIdentification135                       `xml:"Cdtr"`
	CdtrAcct        CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt InstructionForCreditorAgent1                 `xml:"InstrForCdtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      RegulatoryReporting3                         `xml:"RgltryRptg,omitempty"`
	Tax             TaxInformation8                              `xml:"Tax,omitempty"`
	RltdRmtInf      RemittanceLocation7                          `xml:"RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile       Document12                                   `xml:"NclsdFile,omitempty"`
	SplmtryData     SupplementaryData1                           `xml:"SplmtryData,omitempty"`
}

// isValid checks if CreditTransferTransaction35 is valid
func (s *CreditTransferTransaction35) isValid() bool {
	valid :=
		s.PmtId.isValid() &&
			s.PmtTpInf.isValid() &&
			s.PmtCond.isValid() &&
			s.Amt.isValid() &&
			s.ChrgBr.isValid() &&
			s.ChqInstr.isValid() &&
			s.UltmtDbtr.isValid() &&
			s.IntrmyAgt1.isValid() &&
			s.IntrmyAgt2.isValid() &&
			s.IntrmyAgt3.isValid() &&
			s.CdtrAgt.isValid() &&
			s.Cdtr.isValid() &&
			s.CdtrAcct.isValid() &&
			s.UltmtCdtr.isValid() &&
			s.InstrForCdtrAgt.isValid() &&
			s.Purp.isValid() &&
			s.RgltryRptg.isValid() &&
			s.Tax.isValid() &&
			s.RltdRmtInf.isValid() &&
			s.RmtInf.isValid() &&
			s.NclsdFile.isValid() &&
			s.SplmtryData.isValid()

	return valid

}

// CreditorPaymentActivationRequestV07
type CreditorPaymentActivationRequestV07 struct {
	GrpHdr      GroupHeader78        `xml:"GrpHdr"`
	PmtInf      PaymentInstruction31 `xml:"PmtInf"`
	SplmtryData SupplementaryData1   `xml:"SplmtryData,omitempty"`
}

// isValid checks if CreditorPaymentActivationRequestV07 is valid
func (s *CreditorPaymentActivationRequestV07) isValid() bool {
	valid :=
		s.GrpHdr.isValid() &&
			s.PmtInf.isValid() &&
			s.SplmtryData.isValid()

	return valid

}

// GroupHeader78
type GroupHeader78 struct {
	MsgId    Max35Text              `xml:"MsgId"`
	CreDtTm  ISODateTime            `xml:"CreDtTm"`
	NbOfTxs  Max15NumericText       `xml:"NbOfTxs"`
	CtrlSum  float64                `xml:"CtrlSum,omitempty"`
	InitgPty PartyIdentification135 `xml:"InitgPty"`
}

// isValid checks if GroupHeader78 is valid
func (s *GroupHeader78) isValid() bool {
	valid :=
		s.MsgId.isValid() &&
			s.CreDtTm.isValid() &&
			s.NbOfTxs.isValid() &&
			true &&
			s.InitgPty.isValid()

	return valid

}

// InstructionForCreditorAgent1
type InstructionForCreditorAgent1 struct {
	Cd       Instruction3Code `xml:"Cd,omitempty"`
	InstrInf Max140Text       `xml:"InstrInf,omitempty"`
}

// isValid checks if InstructionForCreditorAgent1 is valid
func (s *InstructionForCreditorAgent1) isValid() bool {
	valid :=
		s.Cd.isValid() &&
			s.InstrInf.isValid()

	return valid

}

// PaymentIdentification6
type PaymentIdentification6 struct {
	InstrId    Max35Text        `xml:"InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"EndToEndId"`
	UETR       UUIDv4Identifier `xml:"UETR,omitempty"`
}

// isValid checks if PaymentIdentification6 is valid
func (s *PaymentIdentification6) isValid() bool {
	valid :=
		s.InstrId.isValid() &&
			s.EndToEndId.isValid() &&
			s.UETR.isValid()

	return valid

}

// PaymentInstruction31
type PaymentInstruction31 struct {
	PmtInfId    Max35Text                                    `xml:"PmtInfId,omitempty"`
	PmtMtd      PaymentMethod7Code                           `xml:"PmtMtd"`
	PmtTpInf    PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `xml:"ReqdExctnDt"`
	XpryDt      DateAndDateTime2Choice                       `xml:"XpryDt,omitempty"`
	PmtCond     PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Dbtr        PartyIdentification135                       `xml:"Dbtr"`
	DbtrAcct    CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	UltmtDbtr   PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	ChrgBr      ChargeBearerType1Code                        `xml:"ChrgBr,omitempty"`
	CdtTrfTx    CreditTransferTransaction35                  `xml:"CdtTrfTx"`
}

// isValid checks if PaymentInstruction31 is valid
func (s *PaymentInstruction31) isValid() bool {
	valid :=
		s.PmtInfId.isValid() &&
			s.PmtMtd.isValid() &&
			s.PmtTpInf.isValid() &&
			s.ReqdExctnDt.isValid() &&
			s.XpryDt.isValid() &&
			s.PmtCond.isValid() &&
			s.Dbtr.isValid() &&
			s.DbtrAcct.isValid() &&
			s.DbtrAgt.isValid() &&
			s.UltmtDbtr.isValid() &&
			s.ChrgBr.isValid() &&
			s.CdtTrfTx.isValid()

	return valid

}

// RegulatoryAuthority2
type RegulatoryAuthority2 struct {
	Nm   Max140Text  `xml:"Nm,omitempty"`
	Ctry CountryCode `xml:"Ctry,omitempty"`
}

// isValid checks if RegulatoryAuthority2 is valid
func (s *RegulatoryAuthority2) isValid() bool {
	valid :=
		s.Nm.isValid() &&
			s.Ctry.isValid()

	return valid

}

// RegulatoryReporting3
type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code   `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2           `xml:"Authrty,omitempty"`
	Dtls          StructuredRegulatoryReporting3 `xml:"Dtls,omitempty"`
}

// isValid checks if RegulatoryReporting3 is valid
func (s *RegulatoryReporting3) isValid() bool {
	valid :=
		s.DbtCdtRptgInd.isValid() &&
			s.Authrty.isValid() &&
			s.Dtls.isValid()

	return valid

}

// RemittanceLocation7
type RemittanceLocation7 struct {
	RmtId       Max35Text               `xml:"RmtId,omitempty"`
	RmtLctnDtls RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty"`
}

// isValid checks if RemittanceLocation7 is valid
func (s *RemittanceLocation7) isValid() bool {
	valid :=
		s.RmtId.isValid() &&
			s.RmtLctnDtls.isValid()

	return valid

}

// RemittanceLocationData1
type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr Max2048Text                   `xml:"ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty"`
}

// isValid checks if RemittanceLocationData1 is valid
func (s *RemittanceLocationData1) isValid() bool {
	valid :=
		s.Mtd.isValid() &&
			s.ElctrncAdr.isValid() &&
			s.PstlAdr.isValid()

	return valid

}

// StructuredRegulatoryReporting3
type StructuredRegulatoryReporting3 struct {
	Tp   Max35Text                         `xml:"Tp,omitempty"`
	Dt   ISODate                           `xml:"Dt,omitempty"`
	Ctry CountryCode                       `xml:"Ctry,omitempty"`
	Cd   Max10Text                         `xml:"Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	Inf  Max35Text                         `xml:"Inf,omitempty"`
}

// isValid checks if StructuredRegulatoryReporting3 is valid
func (s *StructuredRegulatoryReporting3) isValid() bool {
	valid :=
		s.Tp.isValid() &&
			s.Dt.isValid() &&
			s.Ctry.isValid() &&
			s.Cd.isValid() &&
			s.Amt.isValid() &&
			s.Inf.isValid()

	return valid

}

// TaxInformation8
type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty"`
	Rcrd            TaxRecord2                        `xml:"Rcrd,omitempty"`
}

// isValid checks if TaxInformation8 is valid
func (s *TaxInformation8) isValid() bool {
	valid :=
		s.Cdtr.isValid() &&
			s.Dbtr.isValid() &&
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
