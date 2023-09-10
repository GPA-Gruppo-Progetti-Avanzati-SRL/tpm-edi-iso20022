// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

// IsValid checks if TaxParty1 is valid
func (s TaxParty1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)

	return valid
}

// IsValid checks if ContactDetails2 is valid
func (s ContactDetails2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NmPrfx.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && s.PhneNb.IsValid(true)
	valid = valid && s.MobNb.IsValid(true)
	valid = valid && s.FaxNb.IsValid(true)
	valid = valid && s.EmailAdr.IsValid(true)
	valid = valid && s.Othr.IsValid(true)

	return valid
}

// IsValid checks if ProprietaryAgent2 is valid
func (s ProprietaryAgent2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Agt != nil && s.Agt.IsValid(false)

	return valid
}

// IsValid checks if BalanceType12 is valid
func (s BalanceType12) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && (s.SubTp == nil || (s.SubTp != nil && s.SubTp.IsValid(true)))

	return valid
}

// IsValid checks if BatchInformation2 is valid
func (s BatchInformation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(true)
	valid = valid && s.PmtInfId.IsValid(true)
	valid = valid && s.NbOfTxs.IsValid(true)
	valid = valid && (s.TtlAmt == nil || (s.TtlAmt != nil && s.TtlAmt.IsValid(true)))

	valid = valid && s.CdtDbtInd.IsValid(true)

	return valid
}

// IsValid checks if TransactionDates2 is valid
func (s TransactionDates2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.AccptncDtTm.IsValid(true)
	valid = valid && s.TradActvtyCtrctlSttlmDt.IsValid(true)
	valid = valid && s.TradDt.IsValid(true)
	valid = valid && s.IntrBkSttlmDt.IsValid(true)
	valid = valid && s.StartDt.IsValid(true)
	valid = valid && s.EndDt.IsValid(true)
	valid = valid && s.TxDtTm.IsValid(true)
	for j := 0; j < len(s.Prtry); j++ {
		valid = valid && s.Prtry[j].IsValid(true)
	}

	return valid
}

// IsValid checks if OrganisationIdentificationSchemeName1Choice is valid
func (s OrganisationIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if OrganisationIdentification4 is valid
func (s OrganisationIdentification4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BICOrBEI.IsValid(true)
	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
	}

	return valid
}

// IsValid checks if BalanceSubType1Choice is valid
func (s BalanceSubType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CreditLine2 is valid
func (s CreditLine2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Incl.IsValid(false)
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	return valid
}

// IsValid checks if NameAndAddress10 is valid
func (s NameAndAddress10) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(false)
	valid = valid && s.Adr != nil && s.Adr.IsValid(false)

	return valid
}

// IsValid checks if FinancialInstrumentQuantityChoice is valid
func (s FinancialInstrumentQuantityChoice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Unit.IsValid(true)
	valid = valid && s.FaceAmt.IsValid(true)
	valid = valid && s.AmtsdVal.IsValid(true)

	return valid
}

// IsValid checks if TaxAuthorisation1 is valid
func (s TaxAuthorisation1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Titl.IsValid(true)
	valid = valid && s.Nm.IsValid(true)

	return valid
}

// IsValid checks if ProprietaryParty2 is valid
func (s ProprietaryParty2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Pty != nil && s.Pty.IsValid(false)

	return valid
}

// IsValid checks if GroupHeader42 is valid
func (s GroupHeader42) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && (s.MsgRcpt == nil || (s.MsgRcpt != nil && s.MsgRcpt.IsValid(true)))

	valid = valid && (s.MsgPgntn == nil || (s.MsgPgntn != nil && s.MsgPgntn.IsValid(true)))

	valid = valid && s.AddtlInf.IsValid(true)

	return valid
}

// IsValid checks if TransactionAgents2 is valid
func (s TransactionAgents2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.DbtrAgt == nil || (s.DbtrAgt != nil && s.DbtrAgt.IsValid(true)))

	valid = valid && (s.CdtrAgt == nil || (s.CdtrAgt != nil && s.CdtrAgt.IsValid(true)))

	valid = valid && (s.IntrmyAgt1 == nil || (s.IntrmyAgt1 != nil && s.IntrmyAgt1.IsValid(true)))

	valid = valid && (s.IntrmyAgt2 == nil || (s.IntrmyAgt2 != nil && s.IntrmyAgt2.IsValid(true)))

	valid = valid && (s.IntrmyAgt3 == nil || (s.IntrmyAgt3 != nil && s.IntrmyAgt3.IsValid(true)))

	valid = valid && (s.RcvgAgt == nil || (s.RcvgAgt != nil && s.RcvgAgt.IsValid(true)))

	valid = valid && (s.DlvrgAgt == nil || (s.DlvrgAgt != nil && s.DlvrgAgt.IsValid(true)))

	valid = valid && (s.IssgAgt == nil || (s.IssgAgt != nil && s.IssgAgt.IsValid(true)))

	valid = valid && (s.SttlmPlc == nil || (s.SttlmPlc != nil && s.SttlmPlc.IsValid(true)))

	for j := 0; j < len(s.Prtry); j++ {
		valid = valid && s.Prtry[j].IsValid(true)
	}

	return valid
}

// IsValid checks if TaxInformation3 is valid
func (s TaxInformation3) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Cdtr == nil || (s.Cdtr != nil && s.Cdtr.IsValid(true)))

	valid = valid && (s.Dbtr == nil || (s.Dbtr != nil && s.Dbtr.IsValid(true)))

	valid = valid && s.AdmstnZn.IsValid(true)
	valid = valid && s.RefNb.IsValid(true)
	valid = valid && s.Mtd.IsValid(true)
	valid = valid && (s.TtlTaxblBaseAmt == nil || (s.TtlTaxblBaseAmt != nil && s.TtlTaxblBaseAmt.IsValid(true)))

	valid = valid && (s.TtlTaxAmt == nil || (s.TtlTaxAmt != nil && s.TtlTaxAmt.IsValid(true)))

	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.SeqNb.IsValid(true)
	for j := 0; j < len(s.Rcrd); j++ {
		valid = valid && s.Rcrd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if DateAndDateTimeChoice is valid
func (s DateAndDateTimeChoice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.DtTm.IsValid(true)

	return valid
}

// IsValid checks if TotalTransactions2 is valid
func (s TotalTransactions2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.TtlNtries == nil || (s.TtlNtries != nil && s.TtlNtries.IsValid(true)))

	valid = valid && (s.TtlCdtNtries == nil || (s.TtlCdtNtries != nil && s.TtlCdtNtries.IsValid(true)))

	valid = valid && (s.TtlDbtNtries == nil || (s.TtlDbtNtries != nil && s.TtlDbtNtries.IsValid(true)))

	for j := 0; j < len(s.TtlNtriesPerBkTxCd); j++ {
		valid = valid && s.TtlNtriesPerBkTxCd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if AmountAndCurrencyExchange3 is valid
func (s AmountAndCurrencyExchange3) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.InstdAmt == nil || (s.InstdAmt != nil && s.InstdAmt.IsValid(true)))

	valid = valid && (s.TxAmt == nil || (s.TxAmt != nil && s.TxAmt.IsValid(true)))

	valid = valid && (s.CntrValAmt == nil || (s.CntrValAmt != nil && s.CntrValAmt.IsValid(true)))

	valid = valid && (s.AnncdPstngAmt == nil || (s.AnncdPstngAmt != nil && s.AnncdPstngAmt.IsValid(true)))

	for j := 0; j < len(s.PrtryAmt); j++ {
		valid = valid && s.PrtryAmt[j].IsValid(true)
	}

	return valid
}

// IsValid checks if ReferredDocumentType1Choice is valid
func (s ReferredDocumentType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ProprietaryQuantity1 is valid
func (s ProprietaryQuantity1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Qty.IsValid(false)

	return valid
}

// IsValid checks if CashAccountType2 is valid
func (s CashAccountType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if TaxRecordDetails1 is valid
func (s TaxRecordDetails1) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Prd == nil || (s.Prd != nil && s.Prd.IsValid(true)))

	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	return valid
}

// IsValid checks if ClearingSystemIdentification2Choice is valid
func (s ClearingSystemIdentification2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if BalanceType5Choice is valid
func (s BalanceType5Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CreditorReferenceType2 is valid
func (s CreditorReferenceType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if AmountAndCurrencyExchangeDetails3 is valid
func (s AmountAndCurrencyExchangeDetails3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && (s.CcyXchg == nil || (s.CcyXchg != nil && s.CcyXchg.IsValid(true)))

	return valid
}

// IsValid checks if RemittanceInformation5 is valid
func (s RemittanceInformation5) IsValid(optional bool) bool {

	valid := true
	for j := 0; j < len(s.Ustrd); j++ {
		valid = valid && s.Ustrd[j].IsValid(true)
	}

	for j := 0; j < len(s.Strd); j++ {
		valid = valid && s.Strd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if StructuredRemittanceInformation7 is valid
func (s StructuredRemittanceInformation7) IsValid(optional bool) bool {

	valid := true
	for j := 0; j < len(s.RfrdDocInf); j++ {
		valid = valid && s.RfrdDocInf[j].IsValid(true)
	}

	valid = valid && (s.RfrdDocAmt == nil || (s.RfrdDocAmt != nil && s.RfrdDocAmt.IsValid(true)))

	valid = valid && (s.CdtrRefInf == nil || (s.CdtrRefInf != nil && s.CdtrRefInf.IsValid(true)))

	valid = valid && (s.Invcr == nil || (s.Invcr != nil && s.Invcr.IsValid(true)))

	valid = valid && (s.Invcee == nil || (s.Invcee != nil && s.Invcee.IsValid(true)))

	for j := 0; j < len(s.AddtlRmtInf); j++ {
		valid = valid && s.AddtlRmtInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if BranchAndFinancialInstitutionIdentification4 is valid
func (s BranchAndFinancialInstitutionIdentification4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FinInstnId != nil && s.FinInstnId.IsValid(false)

	valid = valid && (s.BrnchId == nil || (s.BrnchId != nil && s.BrnchId.IsValid(true)))

	return valid
}

// IsValid checks if AccountInterest2 is valid
func (s AccountInterest2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	for j := 0; j < len(s.Rate); j++ {
		valid = valid && s.Rate[j].IsValid(true)
	}

	valid = valid && (s.FrToDt == nil || (s.FrToDt != nil && s.FrToDt.IsValid(true)))

	valid = valid && s.Rsn.IsValid(true)

	return valid
}

// IsValid checks if AmountRangeBoundary1 is valid
func (s AmountRangeBoundary1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BdryAmt.IsValid(false)
	valid = valid && s.Incl.IsValid(false)

	return valid
}

// IsValid checks if ActiveOrHistoricCurrencyAndAmount is valid
func (s ActiveOrHistoricCurrencyAndAmount) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Ccy.IsValid(false)
	valid = valid && s.Value.IsValid(false)

	return valid
}

// IsValid checks if TransactionReferences2 is valid
func (s TransactionReferences2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(true)
	valid = valid && s.AcctSvcrRef.IsValid(true)
	valid = valid && s.PmtInfId.IsValid(true)
	valid = valid && s.InstrId.IsValid(true)
	valid = valid && s.EndToEndId.IsValid(true)
	valid = valid && s.TxId.IsValid(true)
	valid = valid && s.MndtId.IsValid(true)
	valid = valid && s.ChqNb.IsValid(true)
	valid = valid && s.ClrSysRef.IsValid(true)
	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if DatePeriodDetails is valid
func (s DatePeriodDetails) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FrDt.IsValid(false)
	valid = valid && s.ToDt.IsValid(false)

	return valid
}

// IsValid checks if BranchData2 is valid
func (s BranchData2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && (s.PstlAdr == nil || (s.PstlAdr != nil && s.PstlAdr.IsValid(true)))

	return valid
}

// IsValid checks if TransactionQuantities1Choice is valid
func (s TransactionQuantities1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Qty == nil || (s.Qty != nil && s.Qty.IsValid(true)))

	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if Party6Choice is valid
func (s Party6Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgId == nil || (s.OrgId != nil && s.OrgId.IsValid(true)))

	valid = valid && (s.PrvtId == nil || (s.PrvtId != nil && s.PrvtId.IsValid(true)))

	return valid
}

// IsValid checks if TechnicalInputChannel1Choice is valid
func (s TechnicalInputChannel1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if TaxRecord1 is valid
func (s TaxRecord1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(true)
	valid = valid && s.Ctgy.IsValid(true)
	valid = valid && s.CtgyDtls.IsValid(true)
	valid = valid && s.DbtrSts.IsValid(true)
	valid = valid && s.CertId.IsValid(true)
	valid = valid && s.FrmsCd.IsValid(true)
	valid = valid && (s.Prd == nil || (s.Prd != nil && s.Prd.IsValid(true)))

	valid = valid && (s.TaxAmt == nil || (s.TaxAmt != nil && s.TaxAmt.IsValid(true)))

	valid = valid && s.AddtlInf.IsValid(true)

	return valid
}

// IsValid checks if CorporateAction1 is valid
func (s CorporateAction1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Nb.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if FinancialIdentificationSchemeName1Choice is valid
func (s FinancialIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if AmountAndCurrencyExchangeDetails4 is valid
func (s AmountAndCurrencyExchangeDetails4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && (s.CcyXchg == nil || (s.CcyXchg != nil && s.CcyXchg.IsValid(true)))

	return valid
}

// IsValid checks if ChargesInformation6 is valid
func (s ChargesInformation6) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.TtlChrgsAndTaxAmt == nil || (s.TtlChrgsAndTaxAmt != nil && s.TtlChrgsAndTaxAmt.IsValid(true)))

	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Rate.IsValid(true)
	valid = valid && s.Br.IsValid(true)
	valid = valid && (s.Pty == nil || (s.Pty != nil && s.Pty.IsValid(true)))

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	return valid
}

// IsValid checks if EntryTransaction2 is valid
func (s EntryTransaction2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Refs == nil || (s.Refs != nil && s.Refs.IsValid(true)))

	valid = valid && (s.AmtDtls == nil || (s.AmtDtls != nil && s.AmtDtls.IsValid(true)))

	for j := 0; j < len(s.Avlbty); j++ {
		valid = valid && s.Avlbty[j].IsValid(true)
	}

	valid = valid && (s.BkTxCd == nil || (s.BkTxCd != nil && s.BkTxCd.IsValid(true)))

	for j := 0; j < len(s.Chrgs); j++ {
		valid = valid && s.Chrgs[j].IsValid(true)
	}

	for j := 0; j < len(s.Intrst); j++ {
		valid = valid && s.Intrst[j].IsValid(true)
	}

	valid = valid && (s.RltdPties == nil || (s.RltdPties != nil && s.RltdPties.IsValid(true)))

	valid = valid && (s.RltdAgts == nil || (s.RltdAgts != nil && s.RltdAgts.IsValid(true)))

	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	for j := 0; j < len(s.RltdRmtInf); j++ {
		valid = valid && s.RltdRmtInf[j].IsValid(true)
	}

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	valid = valid && (s.RltdDts == nil || (s.RltdDts != nil && s.RltdDts.IsValid(true)))

	valid = valid && (s.RltdPric == nil || (s.RltdPric != nil && s.RltdPric.IsValid(true)))

	for j := 0; j < len(s.RltdQties); j++ {
		valid = valid && s.RltdQties[j].IsValid(true)
	}

	valid = valid && (s.FinInstrmId == nil || (s.FinInstrmId != nil && s.FinInstrmId.IsValid(true)))

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	valid = valid && (s.RtrInf == nil || (s.RtrInf != nil && s.RtrInf.IsValid(true)))

	valid = valid && (s.CorpActn == nil || (s.CorpActn != nil && s.CorpActn.IsValid(true)))

	valid = valid && (s.SfkpgAcct == nil || (s.SfkpgAcct != nil && s.SfkpgAcct.IsValid(true)))

	valid = valid && s.AddtlTxInf.IsValid(true)

	return valid
}

// IsValid checks if RemittanceAmount1 is valid
func (s RemittanceAmount1) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.DuePyblAmt == nil || (s.DuePyblAmt != nil && s.DuePyblAmt.IsValid(true)))

	valid = valid && (s.DscntApldAmt == nil || (s.DscntApldAmt != nil && s.DscntApldAmt.IsValid(true)))

	valid = valid && (s.CdtNoteAmt == nil || (s.CdtNoteAmt != nil && s.CdtNoteAmt.IsValid(true)))

	valid = valid && (s.TaxAmt == nil || (s.TaxAmt != nil && s.TaxAmt.IsValid(true)))

	for j := 0; j < len(s.AdjstmntAmtAndRsn); j++ {
		valid = valid && s.AdjstmntAmtAndRsn[j].IsValid(true)
	}

	valid = valid && (s.RmtdAmt == nil || (s.RmtdAmt != nil && s.RmtdAmt.IsValid(true)))

	return valid
}

// IsValid checks if RemittanceLocation2 is valid
func (s RemittanceLocation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.RmtId.IsValid(true)
	valid = valid && s.RmtLctnMtd.IsValid(true)
	valid = valid && s.RmtLctnElctrncAdr.IsValid(true)
	valid = valid && (s.RmtLctnPstlAdr == nil || (s.RmtLctnPstlAdr != nil && s.RmtLctnPstlAdr.IsValid(true)))

	return valid
}

// IsValid checks if TransactionPrice2Choice is valid
func (s TransactionPrice2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.DealPric == nil || (s.DealPric != nil && s.DealPric.IsValid(true)))

	for j := 0; j < len(s.Prtry); j++ {
		valid = valid && s.Prtry[j].IsValid(true)
	}

	return valid
}

// IsValid checks if TaxParty2 is valid
func (s TaxParty2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)
	valid = valid && (s.Authstn == nil || (s.Authstn != nil && s.Authstn.IsValid(true)))

	return valid
}

// IsValid checks if TaxCharges2 is valid
func (s TaxCharges2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(true)
	valid = valid && s.Rate.IsValid(true)
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	return valid
}

// IsValid checks if PartyIdentification32 is valid
func (s PartyIdentification32) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(true)
	valid = valid && (s.PstlAdr == nil || (s.PstlAdr != nil && s.PstlAdr.IsValid(true)))

	valid = valid && (s.Id == nil || (s.Id != nil && s.Id.IsValid(true)))

	valid = valid && s.CtryOfRes.IsValid(true)
	valid = valid && (s.CtctDtls == nil || (s.CtctDtls != nil && s.CtctDtls.IsValid(true)))

	return valid
}

// IsValid checks if GenericOrganisationIdentification1 is valid
func (s GenericOrganisationIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if AccountSchemeName1Choice is valid
func (s AccountSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if EntryDetails1 is valid
func (s EntryDetails1) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Btch == nil || (s.Btch != nil && s.Btch.IsValid(true)))

	for j := 0; j < len(s.TxDtls); j++ {
		valid = valid && s.TxDtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if Purpose2Choice is valid
func (s Purpose2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CashAccount20 is valid
func (s CashAccount20) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id != nil && s.Id.IsValid(false)

	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Ccy.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && (s.Ownr == nil || (s.Ownr != nil && s.Ownr.IsValid(true)))

	valid = valid && (s.Svcr == nil || (s.Svcr != nil && s.Svcr.IsValid(true)))

	return valid
}

// IsValid checks if GenericAccountIdentification1 is valid
func (s GenericAccountIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if CreditorReferenceInformation2 is valid
func (s CreditorReferenceInformation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Ref.IsValid(true)

	return valid
}

// IsValid checks if FinancialInstitutionIdentification7 is valid
func (s FinancialInstitutionIdentification7) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BIC.IsValid(true)
	valid = valid && (s.ClrSysMmbId == nil || (s.ClrSysMmbId != nil && s.ClrSysMmbId.IsValid(true)))

	valid = valid && s.Nm.IsValid(true)
	valid = valid && (s.PstlAdr == nil || (s.PstlAdr != nil && s.PstlAdr.IsValid(true)))

	valid = valid && (s.Othr == nil || (s.Othr != nil && s.Othr.IsValid(true)))

	return valid
}

// IsValid checks if NumberAndSumOfTransactions2 is valid
func (s NumberAndSumOfTransactions2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NbOfNtries.IsValid(true)
	valid = valid && s.Sum.IsValid(true)
	valid = valid && s.TtlNetNtryAmt.IsValid(true)
	valid = valid && s.CdtDbtInd.IsValid(true)

	return valid
}

// IsValid checks if ProprietaryPrice2 is valid
func (s ProprietaryPrice2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Pric != nil && s.Pric.IsValid(false)

	return valid
}

// IsValid checks if PersonIdentification5 is valid
func (s PersonIdentification5) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.DtAndPlcOfBirth == nil || (s.DtAndPlcOfBirth != nil && s.DtAndPlcOfBirth.IsValid(true)))

	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
	}

	return valid
}

// IsValid checks if Rate3 is valid
func (s Rate3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp != nil && s.Tp.IsValid(false)

	valid = valid && (s.VldtyRg == nil || (s.VldtyRg != nil && s.VldtyRg.IsValid(true)))

	return valid
}

// IsValid checks if SecurityIdentification4Choice is valid
func (s SecurityIdentification4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.ISIN.IsValid(true)
	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if TaxPeriod1 is valid
func (s TaxPeriod1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Yr.IsValid(true)
	valid = valid && s.Tp.IsValid(true)
	valid = valid && (s.FrToDt == nil || (s.FrToDt != nil && s.FrToDt.IsValid(true)))

	return valid
}

// IsValid checks if FromToAmountRange is valid
func (s FromToAmountRange) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FrAmt != nil && s.FrAmt.IsValid(false)

	valid = valid && s.ToAmt != nil && s.ToAmt.IsValid(false)

	return valid
}

// IsValid checks if MessageIdentification2 is valid
func (s MessageIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgNmId.IsValid(true)
	valid = valid && s.MsgId.IsValid(true)

	return valid
}

// IsValid checks if CashAccount16 is valid
func (s CashAccount16) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id != nil && s.Id.IsValid(false)

	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Ccy.IsValid(true)
	valid = valid && s.Nm.IsValid(true)

	return valid
}

// IsValid checks if ReportEntry2 is valid
func (s ReportEntry2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NtryRef.IsValid(true)
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(false)
	valid = valid && s.RvslInd.IsValid(true)
	valid = valid && s.Sts.IsValid(false)
	valid = valid && (s.BookgDt == nil || (s.BookgDt != nil && s.BookgDt.IsValid(true)))

	valid = valid && (s.ValDt == nil || (s.ValDt != nil && s.ValDt.IsValid(true)))

	valid = valid && s.AcctSvcrRef.IsValid(true)
	for j := 0; j < len(s.Avlbty); j++ {
		valid = valid && s.Avlbty[j].IsValid(true)
	}

	valid = valid && s.BkTxCd != nil && s.BkTxCd.IsValid(false)

	valid = valid && s.ComssnWvrInd.IsValid(true)
	valid = valid && (s.AddtlInfInd == nil || (s.AddtlInfInd != nil && s.AddtlInfInd.IsValid(true)))

	valid = valid && (s.AmtDtls == nil || (s.AmtDtls != nil && s.AmtDtls.IsValid(true)))

	for j := 0; j < len(s.Chrgs); j++ {
		valid = valid && s.Chrgs[j].IsValid(true)
	}

	valid = valid && (s.TechInptChanl == nil || (s.TechInptChanl != nil && s.TechInptChanl.IsValid(true)))

	for j := 0; j < len(s.Intrst); j++ {
		valid = valid && s.Intrst[j].IsValid(true)
	}

	for j := 0; j < len(s.NtryDtls); j++ {
		valid = valid && s.NtryDtls[j].IsValid(true)
	}

	valid = valid && s.AddtlNtryInf.IsValid(true)

	return valid
}

// IsValid checks if ProprietaryReference1 is valid
func (s ProprietaryReference1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Ref.IsValid(false)

	return valid
}

// IsValid checks if ReferredDocumentType2 is valid
func (s ReferredDocumentType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if DateTimePeriodDetails is valid
func (s DateTimePeriodDetails) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FrDtTm.IsValid(false)
	valid = valid && s.ToDtTm.IsValid(false)

	return valid
}

// IsValid checks if ClearingSystemMemberIdentification2 is valid
func (s ClearingSystemMemberIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.ClrSysId == nil || (s.ClrSysId != nil && s.ClrSysId.IsValid(true)))

	valid = valid && s.MmbId.IsValid(false)

	return valid
}

// IsValid checks if Pagination is valid
func (s Pagination) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PgNb.IsValid(false)
	valid = valid && s.LastPgInd.IsValid(false)

	return valid
}

// IsValid checks if ReportingSource1Choice is valid
func (s ReportingSource1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CurrencyExchange5 is valid
func (s CurrencyExchange5) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.SrcCcy.IsValid(false)
	valid = valid && s.TrgtCcy.IsValid(true)
	valid = valid && s.UnitCcy.IsValid(true)
	valid = valid && s.XchgRate.IsValid(false)
	valid = valid && s.CtrctId.IsValid(true)
	valid = valid && s.QtnDt.IsValid(true)

	return valid
}

// IsValid checks if DateAndPlaceOfBirth is valid
func (s DateAndPlaceOfBirth) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BirthDt.IsValid(false)
	valid = valid && s.PrvcOfBirth.IsValid(true)
	valid = valid && s.CityOfBirth.IsValid(false)
	valid = valid && s.CtryOfBirth.IsValid(false)

	return valid
}

// IsValid checks if GenericPersonIdentification1 is valid
func (s GenericPersonIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ReferredDocumentInformation3 is valid
func (s ReferredDocumentInformation3) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Nb.IsValid(true)
	valid = valid && s.RltdDt.IsValid(true)

	return valid
}

// IsValid checks if ReturnReason5Choice is valid
func (s ReturnReason5Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if NumberAndSumOfTransactions1 is valid
func (s NumberAndSumOfTransactions1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NbOfNtries.IsValid(true)
	valid = valid && s.Sum.IsValid(true)

	return valid
}

// IsValid checks if BankTransactionCodeStructure4 is valid
func (s BankTransactionCodeStructure4) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Domn == nil || (s.Domn != nil && s.Domn.IsValid(true)))

	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if PersonIdentificationSchemeName1Choice is valid
func (s PersonIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CashBalanceAvailabilityDate1 is valid
func (s CashBalanceAvailabilityDate1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NbOfDays.IsValid(true)
	valid = valid && s.ActlDt.IsValid(true)

	return valid
}

// IsValid checks if CashBalance3 is valid
func (s CashBalance3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp != nil && s.Tp.IsValid(false)

	valid = valid && (s.CdtLine == nil || (s.CdtLine != nil && s.CdtLine.IsValid(true)))

	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(false)
	valid = valid && s.Dt != nil && s.Dt.IsValid(false)

	for j := 0; j < len(s.Avlbty); j++ {
		valid = valid && s.Avlbty[j].IsValid(true)
	}

	return valid
}

// IsValid checks if ChargeType2Choice is valid
func (s ChargeType2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if TransactionInterest2 is valid
func (s TransactionInterest2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(false)
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	for j := 0; j < len(s.Rate); j++ {
		valid = valid && s.Rate[j].IsValid(true)
	}

	valid = valid && (s.FrToDt == nil || (s.FrToDt != nil && s.FrToDt.IsValid(true)))

	valid = valid && s.Rsn.IsValid(true)

	return valid
}

// IsValid checks if AlternateSecurityIdentification2 is valid
func (s AlternateSecurityIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Id.IsValid(false)

	return valid
}

// IsValid checks if GenericFinancialIdentification1 is valid
func (s GenericFinancialIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if CashBalanceAvailability2 is valid
func (s CashBalanceAvailability2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Dt != nil && s.Dt.IsValid(false)

	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(false)

	return valid
}

// IsValid checks if TotalsPerBankTransactionCode2 is valid
func (s TotalsPerBankTransactionCode2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.NbOfNtries.IsValid(true)
	valid = valid && s.Sum.IsValid(true)
	valid = valid && s.TtlNetNtryAmt.IsValid(true)
	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && s.FcstInd.IsValid(true)
	valid = valid && s.BkTxCd != nil && s.BkTxCd.IsValid(false)

	for j := 0; j < len(s.Avlbty); j++ {
		valid = valid && s.Avlbty[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CurrencyAndAmountRange2 is valid
func (s CurrencyAndAmountRange2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && s.Ccy.IsValid(false)

	return valid
}

// IsValid checks if DocumentAdjustment1 is valid
func (s DocumentAdjustment1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && s.Rsn.IsValid(true)
	valid = valid && s.AddtlInf.IsValid(true)

	return valid
}

// IsValid checks if BankTransactionCodeStructure5 is valid
func (s BankTransactionCodeStructure5) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(false)
	valid = valid && s.Fmly != nil && s.Fmly.IsValid(false)

	return valid
}

// IsValid checks if ProprietaryDate2 is valid
func (s ProprietaryDate2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(false)
	valid = valid && s.Dt != nil && s.Dt.IsValid(false)

	return valid
}

// IsValid checks if ReturnReasonInformation10 is valid
func (s ReturnReasonInformation10) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgnlBkTxCd == nil || (s.OrgnlBkTxCd != nil && s.OrgnlBkTxCd.IsValid(true)))

	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if BankToCustomerStatementV02 is valid
func (s BankToCustomerStatementV02) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr != nil && s.GrpHdr.IsValid(false)

	if len(s.Stmt) == 0 {
		valid = false
	}
	for j := 0; j < len(s.Stmt); j++ {
		valid = valid && s.Stmt[j].IsValid(false)
	}

	return valid
}

// IsValid checks if RateType4Choice is valid
func (s RateType4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Pctg.IsValid(true)
	valid = valid && s.Othr.IsValid(true)

	return valid
}

// IsValid checks if GenericIdentification3 is valid
func (s GenericIdentification3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if TransactionParty2 is valid
func (s TransactionParty2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.InitgPty == nil || (s.InitgPty != nil && s.InitgPty.IsValid(true)))

	valid = valid && (s.Dbtr == nil || (s.Dbtr != nil && s.Dbtr.IsValid(true)))

	valid = valid && (s.DbtrAcct == nil || (s.DbtrAcct != nil && s.DbtrAcct.IsValid(true)))

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.Cdtr == nil || (s.Cdtr != nil && s.Cdtr.IsValid(true)))

	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	valid = valid && (s.TradgPty == nil || (s.TradgPty != nil && s.TradgPty.IsValid(true)))

	for j := 0; j < len(s.Prtry); j++ {
		valid = valid && s.Prtry[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CreditorReferenceType1Choice is valid
func (s CreditorReferenceType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if InterestType1Choice is valid
func (s InterestType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if BankTransactionCodeStructure6 is valid
func (s BankTransactionCodeStructure6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(false)
	valid = valid && s.SubFmlyCd.IsValid(false)

	return valid
}

// IsValid checks if PostalAddress6 is valid
func (s PostalAddress6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.AdrTp.IsValid(true)
	valid = valid && s.Dept.IsValid(true)
	valid = valid && s.SubDept.IsValid(true)
	valid = valid && s.StrtNm.IsValid(true)
	valid = valid && s.BldgNb.IsValid(true)
	valid = valid && s.PstCd.IsValid(true)
	valid = valid && s.TwnNm.IsValid(true)
	valid = valid && s.CtrySubDvsn.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)
	for j := 0; j < len(s.AdrLine); j++ {
		valid = valid && s.AdrLine[j].IsValid(true)
	}

	return valid
}

// IsValid checks if ImpliedCurrencyAmountRangeChoice is valid
func (s ImpliedCurrencyAmountRangeChoice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.FrAmt == nil || (s.FrAmt != nil && s.FrAmt.IsValid(true)))

	valid = valid && (s.ToAmt == nil || (s.ToAmt != nil && s.ToAmt.IsValid(true)))

	valid = valid && (s.FrToAmt == nil || (s.FrToAmt != nil && s.FrToAmt.IsValid(true)))

	valid = valid && s.EQAmt.IsValid(true)
	valid = valid && s.NEQAmt.IsValid(true)

	return valid
}

// IsValid checks if AccountIdentification4Choice is valid
func (s AccountIdentification4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.IBAN.IsValid(true)
	valid = valid && (s.Othr == nil || (s.Othr != nil && s.Othr.IsValid(true)))

	return valid
}

// IsValid checks if ProprietaryBankTransactionCodeStructure1 is valid
func (s ProprietaryBankTransactionCodeStructure1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if TaxAmount1 is valid
func (s TaxAmount1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Rate.IsValid(true)
	valid = valid && (s.TaxblBaseAmt == nil || (s.TaxblBaseAmt != nil && s.TaxblBaseAmt.IsValid(true)))

	valid = valid && (s.TtlAmt == nil || (s.TtlAmt != nil && s.TtlAmt.IsValid(true)))

	for j := 0; j < len(s.Dtls); j++ {
		valid = valid && s.Dtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if AccountStatement2 is valid
func (s AccountStatement2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.ElctrncSeqNb.IsValid(true)
	valid = valid && s.LglSeqNb.IsValid(true)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && (s.FrToDt == nil || (s.FrToDt != nil && s.FrToDt.IsValid(true)))

	valid = valid && s.CpyDplctInd.IsValid(true)
	valid = valid && (s.RptgSrc == nil || (s.RptgSrc != nil && s.RptgSrc.IsValid(true)))

	valid = valid && s.Acct != nil && s.Acct.IsValid(false)

	valid = valid && (s.RltdAcct == nil || (s.RltdAcct != nil && s.RltdAcct.IsValid(true)))

	for j := 0; j < len(s.Intrst); j++ {
		valid = valid && s.Intrst[j].IsValid(true)
	}

	if len(s.Bal) == 0 {
		valid = false
	}
	for j := 0; j < len(s.Bal); j++ {
		valid = valid && s.Bal[j].IsValid(false)
	}

	valid = valid && (s.TxsSummry == nil || (s.TxsSummry != nil && s.TxsSummry.IsValid(true)))

	for j := 0; j < len(s.Ntry); j++ {
		valid = valid && s.Ntry[j].IsValid(true)
	}

	valid = valid && s.AddtlStmtInf.IsValid(true)

	return valid
}
