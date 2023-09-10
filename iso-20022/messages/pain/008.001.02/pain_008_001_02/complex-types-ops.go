// Package pain_008_001_02
// Do not Edit. This stuff it's been automatically generated.
package pain_008_001_02

// IsValid checks if DatePeriodDetails is valid
func (s DatePeriodDetails) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FrDt.IsValid(false)
	valid = valid && s.ToDt.IsValid(false)

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

// IsValid checks if AmendmentInformationDetails6 is valid
func (s AmendmentInformationDetails6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlMndtId.IsValid(true)
	valid = valid && (s.OrgnlCdtrSchmeId == nil || (s.OrgnlCdtrSchmeId != nil && s.OrgnlCdtrSchmeId.IsValid(true)))

	valid = valid && (s.OrgnlCdtrAgt == nil || (s.OrgnlCdtrAgt != nil && s.OrgnlCdtrAgt.IsValid(true)))

	valid = valid && (s.OrgnlCdtrAgtAcct == nil || (s.OrgnlCdtrAgtAcct != nil && s.OrgnlCdtrAgtAcct.IsValid(true)))

	valid = valid && (s.OrgnlDbtr == nil || (s.OrgnlDbtr != nil && s.OrgnlDbtr.IsValid(true)))

	valid = valid && (s.OrgnlDbtrAcct == nil || (s.OrgnlDbtrAcct != nil && s.OrgnlDbtrAcct.IsValid(true)))

	valid = valid && (s.OrgnlDbtrAgt == nil || (s.OrgnlDbtrAgt != nil && s.OrgnlDbtrAgt.IsValid(true)))

	valid = valid && (s.OrgnlDbtrAgtAcct == nil || (s.OrgnlDbtrAgtAcct != nil && s.OrgnlDbtrAgtAcct.IsValid(true)))

	valid = valid && s.OrgnlFnlColltnDt.IsValid(true)
	valid = valid && s.OrgnlFrqcy.IsValid(true)

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

// IsValid checks if GroupHeader39 is valid
func (s GroupHeader39) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	for j := 0; j < len(s.Authstn); j++ {
		valid = valid && s.Authstn[j].IsValid(true)
	}

	valid = valid && s.NbOfTxs.IsValid(false)
	valid = valid && s.CtrlSum.IsValid(true)
	valid = valid && s.InitgPty != nil && s.InitgPty.IsValid(false)

	valid = valid && (s.FwdgAgt == nil || (s.FwdgAgt != nil && s.FwdgAgt.IsValid(true)))

	return valid
}

// IsValid checks if PersonIdentificationSchemeName1Choice is valid
func (s PersonIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if DirectDebitTransaction6 is valid
func (s DirectDebitTransaction6) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.MndtRltdInf == nil || (s.MndtRltdInf != nil && s.MndtRltdInf.IsValid(true)))

	valid = valid && (s.CdtrSchmeId == nil || (s.CdtrSchmeId != nil && s.CdtrSchmeId.IsValid(true)))

	valid = valid && s.PreNtfctnId.IsValid(true)
	valid = valid && s.PreNtfctnDt.IsValid(true)

	return valid
}

// IsValid checks if MandateRelatedInformation6 is valid
func (s MandateRelatedInformation6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MndtId.IsValid(true)
	valid = valid && s.DtOfSgntr.IsValid(true)
	valid = valid && s.AmdmntInd.IsValid(true)
	valid = valid && (s.AmdmntInfDtls == nil || (s.AmdmntInfDtls != nil && s.AmdmntInfDtls.IsValid(true)))

	valid = valid && s.ElctrncSgntr.IsValid(true)
	valid = valid && s.FrstColltnDt.IsValid(true)
	valid = valid && s.FnlColltnDt.IsValid(true)
	valid = valid && s.Frqcy.IsValid(true)

	return valid
}

// IsValid checks if TaxRecordDetails1 is valid
func (s TaxRecordDetails1) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Prd == nil || (s.Prd != nil && s.Prd.IsValid(true)))

	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	return valid
}

// IsValid checks if PaymentTypeInformation20 is valid
func (s PaymentTypeInformation20) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.InstrPrty.IsValid(true)
	valid = valid && (s.SvcLvl == nil || (s.SvcLvl != nil && s.SvcLvl.IsValid(true)))

	valid = valid && (s.LclInstrm == nil || (s.LclInstrm != nil && s.LclInstrm.IsValid(true)))

	valid = valid && s.SeqTp.IsValid(true)
	valid = valid && (s.CtgyPurp == nil || (s.CtgyPurp != nil && s.CtgyPurp.IsValid(true)))

	return valid
}

// IsValid checks if NameAndAddress10 is valid
func (s NameAndAddress10) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(false)
	valid = valid && s.Adr != nil && s.Adr.IsValid(false)

	return valid
}

// IsValid checks if CreditorReferenceInformation2 is valid
func (s CreditorReferenceInformation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Ref.IsValid(true)

	return valid
}

// IsValid checks if LocalInstrument2Choice is valid
func (s LocalInstrument2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ServiceLevel8Choice is valid
func (s ServiceLevel8Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if TaxParty1 is valid
func (s TaxParty1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)

	return valid
}

// IsValid checks if TaxAuthorisation1 is valid
func (s TaxAuthorisation1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Titl.IsValid(true)
	valid = valid && s.Nm.IsValid(true)

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

// IsValid checks if GenericFinancialIdentification1 is valid
func (s GenericFinancialIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

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

// IsValid checks if RegulatoryReporting3 is valid
func (s RegulatoryReporting3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DbtCdtRptgInd.IsValid(true)
	valid = valid && (s.Authrty == nil || (s.Authrty != nil && s.Authrty.IsValid(true)))

	for j := 0; j < len(s.Dtls); j++ {
		valid = valid && s.Dtls[j].IsValid(true)
	}

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

// IsValid checks if ActiveOrHistoricCurrencyAndAmount is valid
func (s ActiveOrHistoricCurrencyAndAmount) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Ccy.IsValid(false)
	valid = valid && s.Value.IsValid(false)

	return valid
}

// IsValid checks if FinancialIdentificationSchemeName1Choice is valid
func (s FinancialIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ClearingSystemIdentification2Choice is valid
func (s ClearingSystemIdentification2Choice) IsValid(optional bool) bool {

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

// IsValid checks if BranchAndFinancialInstitutionIdentification4 is valid
func (s BranchAndFinancialInstitutionIdentification4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.FinInstnId != nil && s.FinInstnId.IsValid(false)

	valid = valid && (s.BrnchId == nil || (s.BrnchId != nil && s.BrnchId.IsValid(true)))

	return valid
}

// IsValid checks if RegulatoryAuthority2 is valid
func (s RegulatoryAuthority2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)

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

// IsValid checks if PaymentIdentification1 is valid
func (s PaymentIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.InstrId.IsValid(true)
	valid = valid && s.EndToEndId.IsValid(false)

	return valid
}

// IsValid checks if CreditorReferenceType1Choice is valid
func (s CreditorReferenceType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if OrganisationIdentificationSchemeName1Choice is valid
func (s OrganisationIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if CashAccount16 is valid
func (s CashAccount16) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id != nil && s.Id.IsValid(false)

	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Ccy.IsValid(true)
	valid = valid && s.Nm.IsValid(true)

	return valid
}

// IsValid checks if DirectDebitTransactionInformation9 is valid
func (s DirectDebitTransactionInformation9) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtId != nil && s.PmtId.IsValid(false)

	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && s.InstdAmt != nil && s.InstdAmt.IsValid(false)

	valid = valid && s.ChrgBr.IsValid(true)
	valid = valid && (s.DrctDbtTx == nil || (s.DrctDbtTx != nil && s.DrctDbtTx.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	valid = valid && s.DbtrAgt != nil && s.DbtrAgt.IsValid(false)

	valid = valid && (s.DbtrAgtAcct == nil || (s.DbtrAgtAcct != nil && s.DbtrAgtAcct.IsValid(true)))

	valid = valid && s.Dbtr != nil && s.Dbtr.IsValid(false)

	valid = valid && s.DbtrAcct != nil && s.DbtrAcct.IsValid(false)

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && s.InstrForCdtrAgt.IsValid(true)
	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	for j := 0; j < len(s.RgltryRptg); j++ {
		valid = valid && s.RgltryRptg[j].IsValid(true)
	}

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	for j := 0; j < len(s.RltdRmtInf); j++ {
		valid = valid && s.RltdRmtInf[j].IsValid(true)
	}

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	return valid
}

// IsValid checks if Authorisation1Choice is valid
func (s Authorisation1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if ClearingSystemMemberIdentification2 is valid
func (s ClearingSystemMemberIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.ClrSysId == nil || (s.ClrSysId != nil && s.ClrSysId.IsValid(true)))

	valid = valid && s.MmbId.IsValid(false)

	return valid
}

// IsValid checks if StructuredRegulatoryReporting3 is valid
func (s StructuredRegulatoryReporting3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(true)
	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)
	valid = valid && s.Cd.IsValid(true)
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	for j := 0; j < len(s.Inf); j++ {
		valid = valid && s.Inf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CategoryPurpose1Choice is valid
func (s CategoryPurpose1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if PaymentInstructionInformation4 is valid
func (s PaymentInstructionInformation4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtInfId.IsValid(false)
	valid = valid && s.PmtMtd.IsValid(false)
	valid = valid && s.BtchBookg.IsValid(true)
	valid = valid && s.NbOfTxs.IsValid(true)
	valid = valid && s.CtrlSum.IsValid(true)
	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && s.ReqdColltnDt.IsValid(false)
	valid = valid && s.Cdtr != nil && s.Cdtr.IsValid(false)

	valid = valid && s.CdtrAcct != nil && s.CdtrAcct.IsValid(false)

	valid = valid && s.CdtrAgt != nil && s.CdtrAgt.IsValid(false)

	valid = valid && (s.CdtrAgtAcct == nil || (s.CdtrAgtAcct != nil && s.CdtrAgtAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	valid = valid && s.ChrgBr.IsValid(true)
	valid = valid && (s.ChrgsAcct == nil || (s.ChrgsAcct != nil && s.ChrgsAcct.IsValid(true)))

	valid = valid && (s.ChrgsAcctAgt == nil || (s.ChrgsAcctAgt != nil && s.ChrgsAcctAgt.IsValid(true)))

	valid = valid && (s.CdtrSchmeId == nil || (s.CdtrSchmeId != nil && s.CdtrSchmeId.IsValid(true)))

	if len(s.DrctDbtTxInf) == 0 {
		valid = false
	}
	for j := 0; j < len(s.DrctDbtTxInf); j++ {
		valid = valid && s.DrctDbtTxInf[j].IsValid(false)
	}

	return valid
}

// IsValid checks if CashAccountType2 is valid
func (s CashAccountType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CustomerDirectDebitInitiationV02 is valid
func (s CustomerDirectDebitInitiationV02) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr != nil && s.GrpHdr.IsValid(false)

	if len(s.PmtInf) == 0 {
		valid = false
	}
	for j := 0; j < len(s.PmtInf); j++ {
		valid = valid && s.PmtInf[j].IsValid(false)
	}

	return valid
}

// IsValid checks if Party6Choice is valid
func (s Party6Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgId == nil || (s.OrgId != nil && s.OrgId.IsValid(true)))

	valid = valid && (s.PrvtId == nil || (s.PrvtId != nil && s.PrvtId.IsValid(true)))

	return valid
}

// IsValid checks if Purpose2Choice is valid
func (s Purpose2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if GenericOrganisationIdentification1 is valid
func (s GenericOrganisationIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

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

// IsValid checks if GenericAccountIdentification1 is valid
func (s GenericAccountIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if AccountIdentification4Choice is valid
func (s AccountIdentification4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.IBAN.IsValid(true)
	valid = valid && (s.Othr == nil || (s.Othr != nil && s.Othr.IsValid(true)))

	return valid
}

// IsValid checks if ReferredDocumentType2 is valid
func (s ReferredDocumentType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && s.Issr.IsValid(true)

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

// IsValid checks if TaxParty2 is valid
func (s TaxParty2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.TaxId.IsValid(true)
	valid = valid && s.RegnId.IsValid(true)
	valid = valid && s.TaxTp.IsValid(true)
	valid = valid && (s.Authstn == nil || (s.Authstn != nil && s.Authstn.IsValid(true)))

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

// IsValid checks if AccountSchemeName1Choice is valid
func (s AccountSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if ReferredDocumentType1Choice is valid
func (s ReferredDocumentType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}
