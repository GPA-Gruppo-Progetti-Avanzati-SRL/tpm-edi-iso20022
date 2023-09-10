// Package camt_053_001_02_test
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/camt_053_001_02"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_camt_053_001_02 = "example-document-camt_053_001_02.xml"

func TestDocumentcamt_053_001_02_SetOps(t *testing.T) {
	doc := camt_053_001_02.NewDocument()
	_ = doc.Set(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr_AddtlInf, "", camt_053_001_02.SetOpWithLog(true))
	_ = doc.Set(camt_053_001_02.Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_MsgId, "", camt_053_001_02.SetOpWithLog(true))
	getv, err := doc.GetNode(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr_MsgRcpt)
	require.NoError(t, err)
	_ = doc.SetNode(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr_MsgRcpt, camt_053_001_02.PartyIdentification32{Nm: "vabbuo'"}, camt_053_001_02.SetOpWithLog(true))
	getv, err = doc.GetNode(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr_MsgRcpt)
	require.NoError(t, err)
	t.Log(getv)

	err = doc.ClearNode(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr_MsgRcpt)
	require.NoError(t, err)

	err = doc.ClearNode(camt_053_001_02.Path_BkToCstmrStmt_GrpHdr)
	require.NoError(t, err)

	err = doc.ClearNode(camt_053_001_02.Path_BkToCstmrStmt_Stmt)
	require.NoError(t, err)

	err = doc.ClearNode("BkToCstmrStmt.Stmt")
	require.NoError(t, err)

	b, err := doc.ToXML()
	require.NoError(t, err)
	t.Log(string(b))
}

func TestDocumentcamt_053_001_02(t *testing.T) {

	d := camt_053_001_02.Document{
		BkToCstmrStmt: &camt_053_001_02.BankToCustomerStatementV02{
			GrpHdr: &camt_053_001_02.GroupHeader42{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				MsgRcpt: &camt_053_001_02.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &camt_053_001_02.PostalAddress6{
						AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &camt_053_001_02.Party6Choice{
						OrgId: &camt_053_001_02.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &camt_053_001_02.PersonIdentification5{
							DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []camt_053_001_02.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &camt_053_001_02.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				MsgPgntn: &camt_053_001_02.Pagination{
					PgNb:      common.MustToMax5NumericText(common.Max5NumericTextSample),
					LastPgInd: xsdt.MustToBoolean(xsdt.BooleanSample),
				},
				AddtlInf: common.MustToMax500Text(common.Max500TextSample),
			},
			Stmt: []camt_053_001_02.AccountStatement2{{
				Id:           common.MustToMax35Text(common.Max35TextSample),
				ElctrncSeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
				LglSeqNb:     xsdt.MustToDecimal(xsdt.DecimalSample),
				CreDtTm:      common.MustToISODateTime(common.ISODateTimeSample),
				FrToDt: &camt_053_001_02.DateTimePeriodDetails{
					FrDtTm: common.MustToISODateTime(common.ISODateTimeSample),
					ToDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				},
				CpyDplctInd: common.MustToCopyDuplicate1Code(common.CopyDuplicate1CodeSample),
				RptgSrc: &camt_053_001_02.ReportingSource1Choice{
					Cd:    common.MustToExternalReportingSource1Code(common.ExternalReportingSource1CodeSample),
					Prtry: common.MustToMax35Text(common.Max35TextSample),
				},
				Acct: &camt_053_001_02.CashAccount20{
					Id: &camt_053_001_02.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &camt_053_001_02.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &camt_053_001_02.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &camt_053_001_02.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Ownr: &camt_053_001_02.PartyIdentification32{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &camt_053_001_02.PostalAddress6{
							AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &camt_053_001_02.Party6Choice{
							OrgId: &camt_053_001_02.OrganisationIdentification4{
								BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
								Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &camt_053_001_02.PersonIdentification5{
								DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []camt_053_001_02.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &camt_053_001_02.ContactDetails2{
							NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
							Nm:       common.MustToMax140Text(common.Max140TextSample),
							PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
							Othr:     common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Svcr: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
						FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
							BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
							ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
								ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &camt_053_001_02.PostalAddress6{
								AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Othr: &camt_053_001_02.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &camt_053_001_02.BranchData2{
							Id: common.MustToMax35Text(common.Max35TextSample),
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &camt_053_001_02.PostalAddress6{
								AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
						},
					},
				},
				RltdAcct: &camt_053_001_02.CashAccount16{
					Id: &camt_053_001_02.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &camt_053_001_02.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &camt_053_001_02.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &camt_053_001_02.CashAccountType2{
						Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
				},
				Intrst: []camt_053_001_02.AccountInterest2{{
					Tp: &camt_053_001_02.InterestType1Choice{
						Cd:    common.MustToInterestType1Code(common.InterestType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Rate: []camt_053_001_02.Rate3{{
						Tp: &camt_053_001_02.RateType4Choice{
							Pctg: xsdt.MustToDecimal(xsdt.DecimalSample),
							Othr: common.MustToMax35Text(common.Max35TextSample),
						},
						VldtyRg: &camt_053_001_02.CurrencyAndAmountRange2{
							Amt: &camt_053_001_02.ImpliedCurrencyAmountRangeChoice{
								FrAmt: &camt_053_001_02.AmountRangeBoundary1{
									BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
									Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
								},
								ToAmt: &camt_053_001_02.AmountRangeBoundary1{
									BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
									Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
								},
								FrToAmt: &camt_053_001_02.FromToAmountRange{
									FrAmt: &camt_053_001_02.AmountRangeBoundary1{
										BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
										Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
									},
									ToAmt: &camt_053_001_02.AmountRangeBoundary1{
										BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
										Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
									},
								},
								EQAmt:  xsdt.MustToDecimal(xsdt.DecimalSample),
								NEQAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
							Ccy:       common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						}},
					},
					FrToDt: &camt_053_001_02.DateTimePeriodDetails{
						FrDtTm: common.MustToISODateTime(common.ISODateTimeSample),
						ToDtTm: common.MustToISODateTime(common.ISODateTimeSample),
					},
					Rsn: common.MustToMax35Text(common.Max35TextSample)},
				},
				Bal: []camt_053_001_02.CashBalance3{{
					Tp: &camt_053_001_02.BalanceType12{
						CdOrPrtry: &camt_053_001_02.BalanceType5Choice{
							Cd:    common.MustToBalanceType12Code(common.BalanceType12CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						SubTp: &camt_053_001_02.BalanceSubType1Choice{
							Cd:    common.MustToExternalBalanceSubType1Code(common.ExternalBalanceSubType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					CdtLine: &camt_053_001_02.CreditLine2{
						Incl: xsdt.MustToBoolean(xsdt.BooleanSample),
						Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
					},
					Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
					Dt: &camt_053_001_02.DateAndDateTimeChoice{
						Dt:   common.MustToISODate(common.ISODateSample),
						DtTm: common.MustToISODateTime(common.ISODateTimeSample),
					},
					Avlbty: []camt_053_001_02.CashBalanceAvailability2{{
						Dt: &camt_053_001_02.CashBalanceAvailabilityDate1{
							NbOfDays: common.MustToMax15PlusSignedNumericText(common.Max15PlusSignedNumericTextSample),
							ActlDt:   common.MustToISODate(common.ISODateSample),
						},
						Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample)},
					}},
				},
				TxsSummry: &camt_053_001_02.TotalTransactions2{
					TtlNtries: &camt_053_001_02.NumberAndSumOfTransactions2{
						NbOfNtries:    common.MustToMax15NumericText(common.Max15NumericTextSample),
						Sum:           xsdt.MustToDecimal(xsdt.DecimalSample),
						TtlNetNtryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
						CdtDbtInd:     common.MustToCreditDebitCode(common.CreditDebitCodeSample),
					},
					TtlCdtNtries: &camt_053_001_02.NumberAndSumOfTransactions1{
						NbOfNtries: common.MustToMax15NumericText(common.Max15NumericTextSample),
						Sum:        xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					TtlDbtNtries: &camt_053_001_02.NumberAndSumOfTransactions1{
						NbOfNtries: common.MustToMax15NumericText(common.Max15NumericTextSample),
						Sum:        xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					TtlNtriesPerBkTxCd: []camt_053_001_02.TotalsPerBankTransactionCode2{{
						NbOfNtries:    common.MustToMax15NumericText(common.Max15NumericTextSample),
						Sum:           xsdt.MustToDecimal(xsdt.DecimalSample),
						TtlNetNtryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
						CdtDbtInd:     common.MustToCreditDebitCode(common.CreditDebitCodeSample),
						FcstInd:       xsdt.MustToBoolean(xsdt.BooleanSample),
						BkTxCd: &camt_053_001_02.BankTransactionCodeStructure4{
							Domn: &camt_053_001_02.BankTransactionCodeStructure5{
								Cd: common.MustToExternalBankTransactionDomain1Code(common.ExternalBankTransactionDomain1CodeSample),
								Fmly: &camt_053_001_02.BankTransactionCodeStructure6{
									Cd:        common.MustToExternalBankTransactionFamily1Code(common.ExternalBankTransactionFamily1CodeSample),
									SubFmlyCd: common.MustToExternalBankTransactionSubFamily1Code(common.ExternalBankTransactionSubFamily1CodeSample),
								},
							},
							Prtry: &camt_053_001_02.ProprietaryBankTransactionCodeStructure1{
								Cd:   common.MustToMax35Text(common.Max35TextSample),
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Avlbty: []camt_053_001_02.CashBalanceAvailability2{{
							Dt: &camt_053_001_02.CashBalanceAvailabilityDate1{
								NbOfDays: common.MustToMax15PlusSignedNumericText(common.Max15PlusSignedNumericTextSample),
								ActlDt:   common.MustToISODate(common.ISODateSample),
							},
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample)},
						}},
					},
				},
				Ntry: []camt_053_001_02.ReportEntry2{{
					NtryRef: common.MustToMax35Text(common.Max35TextSample),
					Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
					RvslInd:   xsdt.MustToBoolean(xsdt.BooleanSample),
					Sts:       common.MustToEntryStatus2Code(common.EntryStatus2CodeSample),
					BookgDt: &camt_053_001_02.DateAndDateTimeChoice{
						Dt:   common.MustToISODate(common.ISODateSample),
						DtTm: common.MustToISODateTime(common.ISODateTimeSample),
					},
					ValDt: &camt_053_001_02.DateAndDateTimeChoice{
						Dt:   common.MustToISODate(common.ISODateSample),
						DtTm: common.MustToISODateTime(common.ISODateTimeSample),
					},
					AcctSvcrRef: common.MustToMax35Text(common.Max35TextSample),
					Avlbty: []camt_053_001_02.CashBalanceAvailability2{{
						Dt: &camt_053_001_02.CashBalanceAvailabilityDate1{
							NbOfDays: common.MustToMax15PlusSignedNumericText(common.Max15PlusSignedNumericTextSample),
							ActlDt:   common.MustToISODate(common.ISODateSample),
						},
						Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample)},
					},
					BkTxCd: &camt_053_001_02.BankTransactionCodeStructure4{
						Domn: &camt_053_001_02.BankTransactionCodeStructure5{
							Cd: common.MustToExternalBankTransactionDomain1Code(common.ExternalBankTransactionDomain1CodeSample),
							Fmly: &camt_053_001_02.BankTransactionCodeStructure6{
								Cd:        common.MustToExternalBankTransactionFamily1Code(common.ExternalBankTransactionFamily1CodeSample),
								SubFmlyCd: common.MustToExternalBankTransactionSubFamily1Code(common.ExternalBankTransactionSubFamily1CodeSample),
							},
						},
						Prtry: &camt_053_001_02.ProprietaryBankTransactionCodeStructure1{
							Cd:   common.MustToMax35Text(common.Max35TextSample),
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					ComssnWvrInd: xsdt.MustToBoolean(xsdt.BooleanSample),
					AddtlInfInd: &camt_053_001_02.MessageIdentification2{
						MsgNmId: common.MustToMax35Text(common.Max35TextSample),
						MsgId:   common.MustToMax35Text(common.Max35TextSample),
					},
					AmtDtls: &camt_053_001_02.AmountAndCurrencyExchange3{
						InstdAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyXchg: &camt_053_001_02.CurrencyExchange5{
								SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								CtrctId:  common.MustToMax35Text(common.Max35TextSample),
								QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
							},
						},
						TxAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyXchg: &camt_053_001_02.CurrencyExchange5{
								SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								CtrctId:  common.MustToMax35Text(common.Max35TextSample),
								QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
							},
						},
						CntrValAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyXchg: &camt_053_001_02.CurrencyExchange5{
								SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								CtrctId:  common.MustToMax35Text(common.Max35TextSample),
								QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
							},
						},
						AnncdPstngAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyXchg: &camt_053_001_02.CurrencyExchange5{
								SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								CtrctId:  common.MustToMax35Text(common.Max35TextSample),
								QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
							},
						},
						PrtryAmt: []camt_053_001_02.AmountAndCurrencyExchangeDetails4{{
							Tp: common.MustToMax35Text(common.Max35TextSample),
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyXchg: &camt_053_001_02.CurrencyExchange5{
								SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								CtrctId:  common.MustToMax35Text(common.Max35TextSample),
								QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
							}},
						},
					},
					Chrgs: []camt_053_001_02.ChargesInformation6{{
						TtlChrgsAndTaxAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
						Tp: &camt_053_001_02.ChargeType2Choice{
							Cd: common.MustToChargeType1Code(common.ChargeType1CodeSample),
							Prtry: &camt_053_001_02.GenericIdentification3{
								Id:   common.MustToMax35Text(common.Max35TextSample),
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
						Br:   common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
						Pty: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
							FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
								BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
								ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
									ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &camt_053_001_02.PostalAddress6{
									AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &camt_053_001_02.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &camt_053_001_02.BranchData2{
								Id: common.MustToMax35Text(common.Max35TextSample),
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &camt_053_001_02.PostalAddress6{
									AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						Tax: &camt_053_001_02.TaxCharges2{
							Id:   common.MustToMax35Text(common.Max35TextSample),
							Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
							Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
						}},
					},
					TechInptChanl: &camt_053_001_02.TechnicalInputChannel1Choice{
						Cd:    common.MustToExternalTechnicalInputChannel1Code(common.ExternalTechnicalInputChannel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Intrst: []camt_053_001_02.TransactionInterest2{{
						Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
						Tp: &camt_053_001_02.InterestType1Choice{
							Cd:    common.MustToInterestType1Code(common.InterestType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Rate: []camt_053_001_02.Rate3{{
							Tp: &camt_053_001_02.RateType4Choice{
								Pctg: xsdt.MustToDecimal(xsdt.DecimalSample),
								Othr: common.MustToMax35Text(common.Max35TextSample),
							},
							VldtyRg: &camt_053_001_02.CurrencyAndAmountRange2{
								Amt: &camt_053_001_02.ImpliedCurrencyAmountRangeChoice{
									FrAmt: &camt_053_001_02.AmountRangeBoundary1{
										BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
										Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
									},
									ToAmt: &camt_053_001_02.AmountRangeBoundary1{
										BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
										Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
									},
									FrToAmt: &camt_053_001_02.FromToAmountRange{
										FrAmt: &camt_053_001_02.AmountRangeBoundary1{
											BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
											Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
										},
										ToAmt: &camt_053_001_02.AmountRangeBoundary1{
											BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
											Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
										},
									},
									EQAmt:  xsdt.MustToDecimal(xsdt.DecimalSample),
									NEQAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Ccy:       common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							}},
						},
						FrToDt: &camt_053_001_02.DateTimePeriodDetails{
							FrDtTm: common.MustToISODateTime(common.ISODateTimeSample),
							ToDtTm: common.MustToISODateTime(common.ISODateTimeSample),
						},
						Rsn: common.MustToMax35Text(common.Max35TextSample)},
					},
					NtryDtls: []camt_053_001_02.EntryDetails1{{
						Btch: &camt_053_001_02.BatchInformation2{
							MsgId:    common.MustToMax35Text(common.Max35TextSample),
							PmtInfId: common.MustToMax35Text(common.Max35TextSample),
							NbOfTxs:  common.MustToMax15NumericText(common.Max15NumericTextSample),
							TtlAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
						},
						TxDtls: []camt_053_001_02.EntryTransaction2{{
							Refs: &camt_053_001_02.TransactionReferences2{
								MsgId:       common.MustToMax35Text(common.Max35TextSample),
								AcctSvcrRef: common.MustToMax35Text(common.Max35TextSample),
								PmtInfId:    common.MustToMax35Text(common.Max35TextSample),
								InstrId:     common.MustToMax35Text(common.Max35TextSample),
								EndToEndId:  common.MustToMax35Text(common.Max35TextSample),
								TxId:        common.MustToMax35Text(common.Max35TextSample),
								MndtId:      common.MustToMax35Text(common.Max35TextSample),
								ChqNb:       common.MustToMax35Text(common.Max35TextSample),
								ClrSysRef:   common.MustToMax35Text(common.Max35TextSample),
								Prtry: &camt_053_001_02.ProprietaryReference1{
									Tp:  common.MustToMax35Text(common.Max35TextSample),
									Ref: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							AmtDtls: &camt_053_001_02.AmountAndCurrencyExchange3{
								InstdAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CcyXchg: &camt_053_001_02.CurrencyExchange5{
										SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
										CtrctId:  common.MustToMax35Text(common.Max35TextSample),
										QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
									},
								},
								TxAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CcyXchg: &camt_053_001_02.CurrencyExchange5{
										SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
										CtrctId:  common.MustToMax35Text(common.Max35TextSample),
										QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
									},
								},
								CntrValAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CcyXchg: &camt_053_001_02.CurrencyExchange5{
										SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
										CtrctId:  common.MustToMax35Text(common.Max35TextSample),
										QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
									},
								},
								AnncdPstngAmt: &camt_053_001_02.AmountAndCurrencyExchangeDetails3{
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CcyXchg: &camt_053_001_02.CurrencyExchange5{
										SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
										CtrctId:  common.MustToMax35Text(common.Max35TextSample),
										QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
									},
								},
								PrtryAmt: []camt_053_001_02.AmountAndCurrencyExchangeDetails4{{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CcyXchg: &camt_053_001_02.CurrencyExchange5{
										SrcCcy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										TrgtCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
										CtrctId:  common.MustToMax35Text(common.Max35TextSample),
										QtnDt:    common.MustToISODateTime(common.ISODateTimeSample),
									}},
								},
							},
							Avlbty: []camt_053_001_02.CashBalanceAvailability2{{
								Dt: &camt_053_001_02.CashBalanceAvailabilityDate1{
									NbOfDays: common.MustToMax15PlusSignedNumericText(common.Max15PlusSignedNumericTextSample),
									ActlDt:   common.MustToISODate(common.ISODateSample),
								},
								Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample)},
							},
							BkTxCd: &camt_053_001_02.BankTransactionCodeStructure4{
								Domn: &camt_053_001_02.BankTransactionCodeStructure5{
									Cd: common.MustToExternalBankTransactionDomain1Code(common.ExternalBankTransactionDomain1CodeSample),
									Fmly: &camt_053_001_02.BankTransactionCodeStructure6{
										Cd:        common.MustToExternalBankTransactionFamily1Code(common.ExternalBankTransactionFamily1CodeSample),
										SubFmlyCd: common.MustToExternalBankTransactionSubFamily1Code(common.ExternalBankTransactionSubFamily1CodeSample),
									},
								},
								Prtry: &camt_053_001_02.ProprietaryBankTransactionCodeStructure1{
									Cd:   common.MustToMax35Text(common.Max35TextSample),
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Chrgs: []camt_053_001_02.ChargesInformation6{{
								TtlChrgsAndTaxAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Tp: &camt_053_001_02.ChargeType2Choice{
									Cd: common.MustToChargeType1Code(common.ChargeType1CodeSample),
									Prtry: &camt_053_001_02.GenericIdentification3{
										Id:   common.MustToMax35Text(common.Max35TextSample),
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
								Br:   common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
								Pty: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								Tax: &camt_053_001_02.TaxCharges2{
									Id:   common.MustToMax35Text(common.Max35TextSample),
									Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
									Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
								}},
							},
							Intrst: []camt_053_001_02.TransactionInterest2{{
								Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
								Tp: &camt_053_001_02.InterestType1Choice{
									Cd:    common.MustToInterestType1Code(common.InterestType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Rate: []camt_053_001_02.Rate3{{
									Tp: &camt_053_001_02.RateType4Choice{
										Pctg: xsdt.MustToDecimal(xsdt.DecimalSample),
										Othr: common.MustToMax35Text(common.Max35TextSample),
									},
									VldtyRg: &camt_053_001_02.CurrencyAndAmountRange2{
										Amt: &camt_053_001_02.ImpliedCurrencyAmountRangeChoice{
											FrAmt: &camt_053_001_02.AmountRangeBoundary1{
												BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
												Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
											},
											ToAmt: &camt_053_001_02.AmountRangeBoundary1{
												BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
												Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
											},
											FrToAmt: &camt_053_001_02.FromToAmountRange{
												FrAmt: &camt_053_001_02.AmountRangeBoundary1{
													BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
													Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
												},
												ToAmt: &camt_053_001_02.AmountRangeBoundary1{
													BdryAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
													Incl:    xsdt.MustToBoolean(xsdt.BooleanSample),
												},
											},
											EQAmt:  xsdt.MustToDecimal(xsdt.DecimalSample),
											NEQAmt: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
										Ccy:       common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									}},
								},
								FrToDt: &camt_053_001_02.DateTimePeriodDetails{
									FrDtTm: common.MustToISODateTime(common.ISODateTimeSample),
									ToDtTm: common.MustToISODateTime(common.ISODateTimeSample),
								},
								Rsn: common.MustToMax35Text(common.Max35TextSample)},
							},
							RltdPties: &camt_053_001_02.TransactionParty2{
								InitgPty: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dbtr: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								DbtrAcct: &camt_053_001_02.CashAccount16{
									Id: &camt_053_001_02.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &camt_053_001_02.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &camt_053_001_02.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &camt_053_001_02.CashAccountType2{
										Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
								},
								UltmtDbtr: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Cdtr: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								CdtrAcct: &camt_053_001_02.CashAccount16{
									Id: &camt_053_001_02.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &camt_053_001_02.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &camt_053_001_02.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &camt_053_001_02.CashAccountType2{
										Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
								},
								UltmtCdtr: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								TradgPty: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Prtry: []camt_053_001_02.ProprietaryParty2{{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Pty: &camt_053_001_02.PartyIdentification32{
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Id: &camt_053_001_02.Party6Choice{
											OrgId: &camt_053_001_02.OrganisationIdentification4{
												BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
												Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
											PrvtId: &camt_053_001_02.PersonIdentification5{
												DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
													BirthDt:     common.MustToISODate(common.ISODateSample),
													PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
												},
												Othr: []camt_053_001_02.GenericPersonIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
										},
										CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
										CtctDtls: &camt_053_001_02.ContactDetails2{
											NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
											Nm:       common.MustToMax140Text(common.Max140TextSample),
											PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
											MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
											Othr:     common.MustToMax35Text(common.Max35TextSample),
										},
									}},
								},
							},
							RltdAgts: &camt_053_001_02.TransactionAgents2{
								DbtrAgt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								CdtrAgt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								IntrmyAgt1: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								IntrmyAgt2: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								IntrmyAgt3: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								RcvgAgt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								DlvrgAgt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								IssgAgt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								SttlmPlc: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
											ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &camt_053_001_02.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &camt_053_001_02.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								Prtry: []camt_053_001_02.ProprietaryAgent2{{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Agt: &camt_053_001_02.BranchAndFinancialInstitutionIdentification4{
										FinInstnId: &camt_053_001_02.FinancialInstitutionIdentification7{
											BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
											ClrSysMmbId: &camt_053_001_02.ClearingSystemMemberIdentification2{
												ClrSysId: &camt_053_001_02.ClearingSystemIdentification2Choice{
													Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												MmbId: common.MustToMax35Text(common.Max35TextSample),
											},
											Nm: common.MustToMax140Text(common.Max140TextSample),
											PstlAdr: &camt_053_001_02.PostalAddress6{
												AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
												Dept:        common.MustToMax70Text(common.Max70TextSample),
												SubDept:     common.MustToMax70Text(common.Max70TextSample),
												StrtNm:      common.MustToMax70Text(common.Max70TextSample),
												BldgNb:      common.MustToMax16Text(common.Max16TextSample),
												PstCd:       common.MustToMax16Text(common.Max16TextSample),
												TwnNm:       common.MustToMax35Text(common.Max35TextSample),
												CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
												Ctry:        common.MustToCountryCode(common.CountryCodeSample),
												AdrLine: []common.Max70Text{
													common.MustToMax70Text(common.Max70TextSample),
												},
											},
											Othr: &camt_053_001_02.GenericFinancialIdentification1{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.FinancialIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										BrnchId: &camt_053_001_02.BranchData2{
											Id: common.MustToMax35Text(common.Max35TextSample),
											Nm: common.MustToMax140Text(common.Max140TextSample),
											PstlAdr: &camt_053_001_02.PostalAddress6{
												AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
												Dept:        common.MustToMax70Text(common.Max70TextSample),
												SubDept:     common.MustToMax70Text(common.Max70TextSample),
												StrtNm:      common.MustToMax70Text(common.Max70TextSample),
												BldgNb:      common.MustToMax16Text(common.Max16TextSample),
												PstCd:       common.MustToMax16Text(common.Max16TextSample),
												TwnNm:       common.MustToMax35Text(common.Max35TextSample),
												CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
												Ctry:        common.MustToCountryCode(common.CountryCodeSample),
												AdrLine: []common.Max70Text{
													common.MustToMax70Text(common.Max70TextSample),
												},
											},
										},
									}},
								},
							},
							Purp: &camt_053_001_02.Purpose2Choice{
								Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							RltdRmtInf: []camt_053_001_02.RemittanceLocation2{{
								RmtId:             common.MustToMax35Text(common.Max35TextSample),
								RmtLctnMtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
								RmtLctnElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
								RmtLctnPstlAdr: &camt_053_001_02.NameAndAddress10{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									Adr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								}},
							},
							RmtInf: &camt_053_001_02.RemittanceInformation5{
								Ustrd: []common.Max140Text{
									common.MustToMax140Text(common.Max140TextSample),
								},
								Strd: []camt_053_001_02.StructuredRemittanceInformation7{{
									RfrdDocInf: []camt_053_001_02.ReferredDocumentInformation3{{
										Tp: &camt_053_001_02.ReferredDocumentType2{
											CdOrPrtry: &camt_053_001_02.ReferredDocumentType1Choice{
												Cd:    common.MustToDocumentType5Code(common.DocumentType5CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
										Nb:     common.MustToMax35Text(common.Max35TextSample),
										RltdDt: common.MustToISODate(common.ISODateSample)},
									},
									RfrdDocAmt: &camt_053_001_02.RemittanceAmount1{
										DuePyblAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										DscntApldAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										CdtNoteAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TaxAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										AdjstmntAmtAndRsn: []camt_053_001_02.DocumentAdjustment1{{
											Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
											Rsn:       common.MustToMax4Text(common.Max4TextSample),
											AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
										},
										RmtdAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
									},
									CdtrRefInf: &camt_053_001_02.CreditorReferenceInformation2{
										Tp: &camt_053_001_02.CreditorReferenceType2{
											CdOrPrtry: &camt_053_001_02.CreditorReferenceType1Choice{
												Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
										Ref: common.MustToMax35Text(common.Max35TextSample),
									},
									Invcr: &camt_053_001_02.PartyIdentification32{
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Id: &camt_053_001_02.Party6Choice{
											OrgId: &camt_053_001_02.OrganisationIdentification4{
												BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
												Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
											PrvtId: &camt_053_001_02.PersonIdentification5{
												DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
													BirthDt:     common.MustToISODate(common.ISODateSample),
													PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
												},
												Othr: []camt_053_001_02.GenericPersonIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
										},
										CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
										CtctDtls: &camt_053_001_02.ContactDetails2{
											NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
											Nm:       common.MustToMax140Text(common.Max140TextSample),
											PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
											MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
											Othr:     common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Invcee: &camt_053_001_02.PartyIdentification32{
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &camt_053_001_02.PostalAddress6{
											AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Id: &camt_053_001_02.Party6Choice{
											OrgId: &camt_053_001_02.OrganisationIdentification4{
												BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
												Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
											PrvtId: &camt_053_001_02.PersonIdentification5{
												DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
													BirthDt:     common.MustToISODate(common.ISODateSample),
													PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
												},
												Othr: []camt_053_001_02.GenericPersonIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
										},
										CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
										CtctDtls: &camt_053_001_02.ContactDetails2{
											NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
											Nm:       common.MustToMax140Text(common.Max140TextSample),
											PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
											MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
											Othr:     common.MustToMax35Text(common.Max35TextSample),
										},
									},
									AddtlRmtInf: []common.Max140Text{
										common.MustToMax140Text(common.Max140TextSample),
									}},
								},
							},
							RltdDts: &camt_053_001_02.TransactionDates2{
								AccptncDtTm:             common.MustToISODateTime(common.ISODateTimeSample),
								TradActvtyCtrctlSttlmDt: common.MustToISODate(common.ISODateSample),
								TradDt:                  common.MustToISODate(common.ISODateSample),
								IntrBkSttlmDt:           common.MustToISODate(common.ISODateSample),
								StartDt:                 common.MustToISODate(common.ISODateSample),
								EndDt:                   common.MustToISODate(common.ISODateSample),
								TxDtTm:                  common.MustToISODateTime(common.ISODateTimeSample),
								Prtry: []camt_053_001_02.ProprietaryDate2{{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Dt: &camt_053_001_02.DateAndDateTimeChoice{
										Dt:   common.MustToISODate(common.ISODateSample),
										DtTm: common.MustToISODateTime(common.ISODateTimeSample),
									}},
								},
							},
							RltdPric: &camt_053_001_02.TransactionPrice2Choice{
								DealPric: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Prtry: []camt_053_001_02.ProprietaryPrice2{{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Pric: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
							},
							RltdQties: []camt_053_001_02.TransactionQuantities1Choice{{
								Qty: &camt_053_001_02.FinancialInstrumentQuantityChoice{
									Unit:     xsdt.MustToDecimal(xsdt.DecimalSample),
									FaceAmt:  xsdt.MustToDecimal(xsdt.DecimalSample),
									AmtsdVal: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Prtry: &camt_053_001_02.ProprietaryQuantity1{
									Tp:  common.MustToMax35Text(common.Max35TextSample),
									Qty: common.MustToMax35Text(common.Max35TextSample),
								}},
							},
							FinInstrmId: &camt_053_001_02.SecurityIdentification4Choice{
								ISIN: common.MustToISINIdentifier(common.ISINIdentifierSample),
								Prtry: &camt_053_001_02.AlternateSecurityIdentification2{
									Tp: common.MustToMax35Text(common.Max35TextSample),
									Id: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tax: &camt_053_001_02.TaxInformation3{
								Cdtr: &camt_053_001_02.TaxParty1{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								},
								Dbtr: &camt_053_001_02.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &camt_053_001_02.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								AdmstnZn: common.MustToMax35Text(common.Max35TextSample),
								RefNb:    common.MustToMax140Text(common.Max140TextSample),
								Mtd:      common.MustToMax35Text(common.Max35TextSample),
								TtlTaxblBaseAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlTaxAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dt:    common.MustToISODate(common.ISODateSample),
								SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
								Rcrd: []camt_053_001_02.TaxRecord1{{
									Tp:       common.MustToMax35Text(common.Max35TextSample),
									Ctgy:     common.MustToMax35Text(common.Max35TextSample),
									CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
									DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
									CertId:   common.MustToMax35Text(common.Max35TextSample),
									FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
									Prd: &camt_053_001_02.TaxPeriod1{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &camt_053_001_02.DatePeriodDetails{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									TaxAmt: &camt_053_001_02.TaxAmount1{
										Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
										TaxblBaseAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TtlAmt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										Dtls: []camt_053_001_02.TaxRecordDetails1{{
											Prd: &camt_053_001_02.TaxPeriod1{
												Yr: common.MustToISODate(common.ISODateSample),
												Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
												FrToDt: &camt_053_001_02.DatePeriodDetails{
													FrDt: common.MustToISODate(common.ISODateSample),
													ToDt: common.MustToISODate(common.ISODateSample),
												},
											},
											Amt: &camt_053_001_02.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
									},
									AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
								},
							},
							RtrInf: &camt_053_001_02.ReturnReasonInformation10{
								OrgnlBkTxCd: &camt_053_001_02.BankTransactionCodeStructure4{
									Domn: &camt_053_001_02.BankTransactionCodeStructure5{
										Cd: common.MustToExternalBankTransactionDomain1Code(common.ExternalBankTransactionDomain1CodeSample),
										Fmly: &camt_053_001_02.BankTransactionCodeStructure6{
											Cd:        common.MustToExternalBankTransactionFamily1Code(common.ExternalBankTransactionFamily1CodeSample),
											SubFmlyCd: common.MustToExternalBankTransactionSubFamily1Code(common.ExternalBankTransactionSubFamily1CodeSample),
										},
									},
									Prtry: &camt_053_001_02.ProprietaryBankTransactionCodeStructure1{
										Cd:   common.MustToMax35Text(common.Max35TextSample),
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Orgtr: &camt_053_001_02.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &camt_053_001_02.PostalAddress6{
										AdrTp:       common.MustToAddressType2Code(common.AddressType2CodeSample),
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &camt_053_001_02.Party6Choice{
										OrgId: &camt_053_001_02.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []camt_053_001_02.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &camt_053_001_02.PersonIdentification5{
											DtAndPlcOfBirth: &camt_053_001_02.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []camt_053_001_02.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &camt_053_001_02.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &camt_053_001_02.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Rsn: &camt_053_001_02.ReturnReason5Choice{
									Cd:    common.MustToExternalReturnReason1Code(common.ExternalReturnReason1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								AddtlInf: []common.Max105Text{
									common.MustToMax105Text(common.Max105TextSample),
								},
							},
							CorpActn: &camt_053_001_02.CorporateAction1{
								Cd:    common.MustToMax35Text(common.Max35TextSample),
								Nb:    common.MustToMax35Text(common.Max35TextSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							SfkpgAcct: &camt_053_001_02.CashAccount16{
								Id: &camt_053_001_02.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &camt_053_001_02.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &camt_053_001_02.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &camt_053_001_02.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							AddtlTxInf: common.MustToMax500Text(common.Max500TextSample)},
						}},
					},
					AddtlNtryInf: common.MustToMax500Text(common.Max500TextSample)},
				},
				AddtlStmtInf: common.MustToMax500Text(common.Max500TextSample)},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_camt_053_001_02, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_camt_053_001_02)

}
