// Package pain_002_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_03

// IsValid checks if GroupHeader36 is valid
func (s GroupHeader36) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && (s.InitgPty == nil || (s.InitgPty != nil && s.InitgPty.IsValid(true)))

	valid = valid && (s.FwdgAgt == nil || (s.FwdgAgt != nil && s.FwdgAgt.IsValid(true)))

	valid = valid && (s.DbtrAgt == nil || (s.DbtrAgt != nil && s.DbtrAgt.IsValid(true)))

	valid = valid && (s.CdtrAgt == nil || (s.CdtrAgt != nil && s.CdtrAgt.IsValid(true)))

	return valid
}

// IsValid checks if OrganisationIdentificationSchemeName1Choice is valid
func (s OrganisationIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if NumberOfTransactionsPerStatus3 is valid
func (s NumberOfTransactionsPerStatus3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DtldNbOfTxs.IsValid(false)
	valid = valid && s.DtldSts.IsValid(false)
	valid = valid && s.DtldCtrlSum.IsValid(true)

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

// IsValid checks if OrganisationIdentification4 is valid
func (s OrganisationIdentification4) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.BICOrBEI.IsValid(true)
	for j := 0; j < len(s.Othr); j++ {
		valid = valid && s.Othr[j].IsValid(true)
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

// IsValid checks if ClearingSystemIdentification2Choice is valid
func (s ClearingSystemIdentification2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if ReferredDocumentType1Choice is valid
func (s ReferredDocumentType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if StatusReason6Choice is valid
func (s StatusReason6Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if EquivalentAmount2 is valid
func (s EquivalentAmount2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.CcyOfTrf.IsValid(false)

	return valid
}

// IsValid checks if OriginalPaymentInformation1 is valid
func (s OriginalPaymentInformation1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlPmtInfId.IsValid(false)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.PmtInfSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerSts); j++ {
		valid = valid && s.NbOfTxsPerSts[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInfAndSts); j++ {
		valid = valid && s.TxInfAndSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CustomerPaymentStatusReportV03 is valid
func (s CustomerPaymentStatusReportV03) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr != nil && s.GrpHdr.IsValid(false)

	valid = valid && s.OrgnlGrpInfAndSts != nil && s.OrgnlGrpInfAndSts.IsValid(false)

	for j := 0; j < len(s.OrgnlPmtInfAndSts); j++ {
		valid = valid && s.OrgnlPmtInfAndSts[j].IsValid(true)
	}

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

// IsValid checks if Party6Choice is valid
func (s Party6Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgId == nil || (s.OrgId != nil && s.OrgId.IsValid(true)))

	valid = valid && (s.PrvtId == nil || (s.PrvtId != nil && s.PrvtId.IsValid(true)))

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

// IsValid checks if ClearingSystemMemberIdentification2 is valid
func (s ClearingSystemMemberIdentification2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.ClrSysId == nil || (s.ClrSysId != nil && s.ClrSysId.IsValid(true)))

	valid = valid && s.MmbId.IsValid(false)

	return valid
}

// IsValid checks if CategoryPurpose1Choice is valid
func (s CategoryPurpose1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if PersonIdentificationSchemeName1Choice is valid
func (s PersonIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if PaymentTypeInformation22 is valid
func (s PaymentTypeInformation22) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.InstrPrty.IsValid(true)
	valid = valid && s.ClrChanl.IsValid(true)
	valid = valid && (s.SvcLvl == nil || (s.SvcLvl != nil && s.SvcLvl.IsValid(true)))

	valid = valid && (s.LclInstrm == nil || (s.LclInstrm != nil && s.LclInstrm.IsValid(true)))

	valid = valid && s.SeqTp.IsValid(true)
	valid = valid && (s.CtgyPurp == nil || (s.CtgyPurp != nil && s.CtgyPurp.IsValid(true)))

	return valid
}

// IsValid checks if ServiceLevel8Choice is valid
func (s ServiceLevel8Choice) IsValid(optional bool) bool {

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

// IsValid checks if ReferredDocumentInformation3 is valid
func (s ReferredDocumentInformation3) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Nb.IsValid(true)
	valid = valid && s.RltdDt.IsValid(true)

	return valid
}

// IsValid checks if CreditorReferenceType2 is valid
func (s CreditorReferenceType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ActiveOrHistoricCurrencyAndAmount is valid
func (s ActiveOrHistoricCurrencyAndAmount) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Ccy.IsValid(false)
	valid = valid && s.Value.IsValid(false)

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

// IsValid checks if GenericFinancialIdentification1 is valid
func (s GenericFinancialIdentification1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && (s.SchmeNm == nil || (s.SchmeNm != nil && s.SchmeNm.IsValid(true)))

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if OriginalGroupInformation20 is valid
func (s OriginalGroupInformation20) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlMsgId.IsValid(false)
	valid = valid && s.OrgnlMsgNmId.IsValid(false)
	valid = valid && s.OrgnlCreDtTm.IsValid(true)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.GrpSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerSts); j++ {
		valid = valid && s.NbOfTxsPerSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if SettlementInformation13 is valid
func (s SettlementInformation13) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.SttlmMtd.IsValid(false)
	valid = valid && (s.SttlmAcct == nil || (s.SttlmAcct != nil && s.SttlmAcct.IsValid(true)))

	valid = valid && (s.ClrSys == nil || (s.ClrSys != nil && s.ClrSys.IsValid(true)))

	valid = valid && (s.InstgRmbrsmntAgt == nil || (s.InstgRmbrsmntAgt != nil && s.InstgRmbrsmntAgt.IsValid(true)))

	valid = valid && (s.InstgRmbrsmntAgtAcct == nil || (s.InstgRmbrsmntAgtAcct != nil && s.InstgRmbrsmntAgtAcct.IsValid(true)))

	valid = valid && (s.InstdRmbrsmntAgt == nil || (s.InstdRmbrsmntAgt != nil && s.InstdRmbrsmntAgt.IsValid(true)))

	valid = valid && (s.InstdRmbrsmntAgtAcct == nil || (s.InstdRmbrsmntAgtAcct != nil && s.InstdRmbrsmntAgtAcct.IsValid(true)))

	valid = valid && (s.ThrdRmbrsmntAgt == nil || (s.ThrdRmbrsmntAgt != nil && s.ThrdRmbrsmntAgt.IsValid(true)))

	valid = valid && (s.ThrdRmbrsmntAgtAcct == nil || (s.ThrdRmbrsmntAgtAcct != nil && s.ThrdRmbrsmntAgtAcct.IsValid(true)))

	return valid
}

// IsValid checks if AccountIdentification4Choice is valid
func (s AccountIdentification4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.IBAN.IsValid(true)
	valid = valid && (s.Othr == nil || (s.Othr != nil && s.Othr.IsValid(true)))

	return valid
}

// IsValid checks if AccountSchemeName1Choice is valid
func (s AccountSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

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

// IsValid checks if OriginalTransactionReference13 is valid
func (s OriginalTransactionReference13) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.IntrBkSttlmAmt == nil || (s.IntrBkSttlmAmt != nil && s.IntrBkSttlmAmt.IsValid(true)))

	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	valid = valid && s.IntrBkSttlmDt.IsValid(true)
	valid = valid && s.ReqdColltnDt.IsValid(true)
	valid = valid && s.ReqdExctnDt.IsValid(true)
	valid = valid && (s.CdtrSchmeId == nil || (s.CdtrSchmeId != nil && s.CdtrSchmeId.IsValid(true)))

	valid = valid && (s.SttlmInf == nil || (s.SttlmInf != nil && s.SttlmInf.IsValid(true)))

	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && s.PmtMtd.IsValid(true)
	valid = valid && (s.MndtRltdInf == nil || (s.MndtRltdInf != nil && s.MndtRltdInf.IsValid(true)))

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.Dbtr == nil || (s.Dbtr != nil && s.Dbtr.IsValid(true)))

	valid = valid && (s.DbtrAcct == nil || (s.DbtrAcct != nil && s.DbtrAcct.IsValid(true)))

	valid = valid && (s.DbtrAgt == nil || (s.DbtrAgt != nil && s.DbtrAgt.IsValid(true)))

	valid = valid && (s.DbtrAgtAcct == nil || (s.DbtrAgtAcct != nil && s.DbtrAgtAcct.IsValid(true)))

	valid = valid && (s.CdtrAgt == nil || (s.CdtrAgt != nil && s.CdtrAgt.IsValid(true)))

	valid = valid && (s.CdtrAgtAcct == nil || (s.CdtrAgtAcct != nil && s.CdtrAgtAcct.IsValid(true)))

	valid = valid && (s.Cdtr == nil || (s.Cdtr != nil && s.Cdtr.IsValid(true)))

	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

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

// IsValid checks if FinancialIdentificationSchemeName1Choice is valid
func (s FinancialIdentificationSchemeName1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if StatusReasonInformation8 is valid
func (s StatusReasonInformation8) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

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

// IsValid checks if PaymentTransactionInformation25 is valid
func (s PaymentTransactionInformation25) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.StsId.IsValid(true)
	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.TxSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.ChrgsInf); j++ {
		valid = valid && s.ChrgsInf[j].IsValid(true)
	}

	valid = valid && s.AccptncDtTm.IsValid(true)
	valid = valid && s.AcctSvcrRef.IsValid(true)
	valid = valid && s.ClrSysRef.IsValid(true)
	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	return valid
}

// IsValid checks if ChargesInformation5 is valid
func (s ChargesInformation5) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt != nil && s.Amt.IsValid(false)

	valid = valid && s.Pty != nil && s.Pty.IsValid(false)

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

// IsValid checks if BranchData2 is valid
func (s BranchData2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(true)
	valid = valid && s.Nm.IsValid(true)
	valid = valid && (s.PstlAdr == nil || (s.PstlAdr != nil && s.PstlAdr.IsValid(true)))

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

// IsValid checks if AmountType3Choice is valid
func (s AmountType3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.InstdAmt == nil || (s.InstdAmt != nil && s.InstdAmt.IsValid(true)))

	valid = valid && (s.EqvtAmt == nil || (s.EqvtAmt != nil && s.EqvtAmt.IsValid(true)))

	return valid
}

// IsValid checks if CashAccountType2 is valid
func (s CashAccountType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if LocalInstrument2Choice is valid
func (s LocalInstrument2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ReferredDocumentType2 is valid
func (s ReferredDocumentType2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CdOrPrtry != nil && s.CdOrPrtry.IsValid(false)

	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if ClearingSystemIdentification3Choice is valid
func (s ClearingSystemIdentification3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CreditorReferenceType1Choice is valid
func (s CreditorReferenceType1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}
