// Package pain_002_001_03_test
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_03_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/pain_002_001_03"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_pain_002_001_03 = "example-document-pain_002_001_03.xml"

func TestDocumentpain_002_001_03(t *testing.T) {

	d := pain_002_001_03.Document{
		CstmrPmtStsRpt: pain_002_001_03.CustomerPaymentStatusReportV03{
			GrpHdr: pain_002_001_03.GroupHeader36{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				InitgPty: &pain_002_001_03.PartyIdentification32{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_002_001_03.PostalAddress6{
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
					Id: &pain_002_001_03.Party6Choice{
						OrgId: &pain_002_001_03.OrganisationIdentification4{
							BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
							Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_002_001_03.PersonIdentification5{
							DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_002_001_03.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_002_001_03.ContactDetails2{
						NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
						Nm:       common.MustToMax140Text(common.Max140TextSample),
						PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
						Othr:     common.MustToMax35Text(common.Max35TextSample),
					},
				},
				FwdgAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
						Othr: &pain_002_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
				DbtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
						Othr: &pain_002_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
				CdtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
					FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
						BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
						ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
						Othr: &pain_002_001_03.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_03.BranchData2{
						Id: common.MustToMax35Text(common.Max35TextSample),
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
			OrgnlGrpInfAndSts: pain_002_001_03.OriginalGroupInformation20{
				OrgnlMsgId:   common.MustToMax35Text(common.Max35TextSample),
				OrgnlMsgNmId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlCreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				OrgnlNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
				OrgnlCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
				GrpSts:       common.MustToTransactionGroupStatus3Code(common.TransactionGroupStatus3CodeSample),
				StsRsnInf: []pain_002_001_03.StatusReasonInformation8{{
					Orgtr: &pain_002_001_03.PartyIdentification32{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
						Id: &pain_002_001_03.Party6Choice{
							OrgId: &pain_002_001_03.OrganisationIdentification4{
								BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
								Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_002_001_03.PersonIdentification5{
								DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_002_001_03.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_002_001_03.ContactDetails2{
							NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
							Nm:       common.MustToMax140Text(common.Max140TextSample),
							PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
							Othr:     common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Rsn: &pain_002_001_03.StatusReason6Choice{
						Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					AddtlInf: []common.Max105Text{
						common.MustToMax105Text(common.Max105TextSample),
					}},
				},
				NbOfTxsPerSts: []pain_002_001_03.NumberOfTransactionsPerStatus3{{
					DtldNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
					DtldSts:     common.MustToTransactionIndividualStatus3Code(common.TransactionIndividualStatus3CodeSample),
					DtldCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample)},
				},
			},
			OrgnlPmtInfAndSts: []pain_002_001_03.OriginalPaymentInformation1{{
				OrgnlPmtInfId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlNbOfTxs:  common.MustToMax15NumericText(common.Max15NumericTextSample),
				OrgnlCtrlSum:  xsdt.MustToDecimal(xsdt.DecimalSample),
				PmtInfSts:     common.MustToTransactionGroupStatus3Code(common.TransactionGroupStatus3CodeSample),
				StsRsnInf: []pain_002_001_03.StatusReasonInformation8{{
					Orgtr: &pain_002_001_03.PartyIdentification32{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_03.PostalAddress6{
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
						Id: &pain_002_001_03.Party6Choice{
							OrgId: &pain_002_001_03.OrganisationIdentification4{
								BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
								Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_002_001_03.PersonIdentification5{
								DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_002_001_03.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_002_001_03.ContactDetails2{
							NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
							Nm:       common.MustToMax140Text(common.Max140TextSample),
							PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
							Othr:     common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Rsn: &pain_002_001_03.StatusReason6Choice{
						Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					AddtlInf: []common.Max105Text{
						common.MustToMax105Text(common.Max105TextSample),
					}},
				},
				NbOfTxsPerSts: []pain_002_001_03.NumberOfTransactionsPerStatus3{{
					DtldNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
					DtldSts:     common.MustToTransactionIndividualStatus3Code(common.TransactionIndividualStatus3CodeSample),
					DtldCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample)},
				},
				TxInfAndSts: []pain_002_001_03.PaymentTransactionInformation25{{
					StsId:           common.MustToMax35Text(common.Max35TextSample),
					OrgnlInstrId:    common.MustToMax35Text(common.Max35TextSample),
					OrgnlEndToEndId: common.MustToMax35Text(common.Max35TextSample),
					TxSts:           common.MustToTransactionIndividualStatus3Code(common.TransactionIndividualStatus3CodeSample),
					StsRsnInf: []pain_002_001_03.StatusReasonInformation8{{
						Orgtr: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Rsn: &pain_002_001_03.StatusReason6Choice{
							Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						AddtlInf: []common.Max105Text{
							common.MustToMax105Text(common.Max105TextSample),
						}},
					},
					ChrgsInf: []pain_002_001_03.ChargesInformation5{{
						Amt: pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Pty: pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
							FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
								BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
								ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
								Othr: &pain_002_001_03.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_03.BranchData2{
								Id: common.MustToMax35Text(common.Max35TextSample),
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
					AccptncDtTm: common.MustToISODateTime(common.ISODateTimeSample),
					AcctSvcrRef: common.MustToMax35Text(common.Max35TextSample),
					ClrSysRef:   common.MustToMax35Text(common.Max35TextSample),
					OrgnlTxRef: &pain_002_001_03.OriginalTransactionReference13{
						IntrBkSttlmAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Amt: &pain_002_001_03.AmountType3Choice{
							InstdAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							EqvtAmt: &pain_002_001_03.EquivalentAmount2{
								Amt: pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							},
						},
						IntrBkSttlmDt: common.MustToISODate(common.ISODateSample),
						ReqdColltnDt:  common.MustToISODate(common.ISODateSample),
						ReqdExctnDt:   common.MustToISODate(common.ISODateSample),
						CdtrSchmeId: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						SttlmInf: &pain_002_001_03.SettlementInformation13{
							SttlmMtd: common.MustToSettlementMethod1Code(common.SettlementMethod1CodeSample),
							SttlmAcct: &pain_002_001_03.CashAccount16{
								Id: pain_002_001_03.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_03.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_03.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							ClrSys: &pain_002_001_03.ClearingSystemIdentification3Choice{
								Cd:    common.MustToExternalCashClearingSystem1Code(common.ExternalCashClearingSystem1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							InstgRmbrsmntAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
								FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
									BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
									ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Othr: &pain_002_001_03.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_03.BranchData2{
									Id: common.MustToMax35Text(common.Max35TextSample),
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
							InstgRmbrsmntAgtAcct: &pain_002_001_03.CashAccount16{
								Id: pain_002_001_03.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_03.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_03.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							InstdRmbrsmntAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
								FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
									BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
									ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Othr: &pain_002_001_03.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_03.BranchData2{
									Id: common.MustToMax35Text(common.Max35TextSample),
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
							InstdRmbrsmntAgtAcct: &pain_002_001_03.CashAccount16{
								Id: pain_002_001_03.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_03.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_03.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
							ThrdRmbrsmntAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
								FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
									BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
									ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Othr: &pain_002_001_03.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_03.BranchData2{
									Id: common.MustToMax35Text(common.Max35TextSample),
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
							ThrdRmbrsmntAgtAcct: &pain_002_001_03.CashAccount16{
								Id: pain_002_001_03.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_03.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_03.CashAccountType2{
									Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
							},
						},
						PmtTpInf: &pain_002_001_03.PaymentTypeInformation22{
							InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
							ClrChanl:  common.MustToClearingChannel2Code(common.ClearingChannel2CodeSample),
							SvcLvl: &pain_002_001_03.ServiceLevel8Choice{
								Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							LclInstrm: &pain_002_001_03.LocalInstrument2Choice{
								Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							SeqTp: common.MustToSequenceType1Code(common.SequenceType1CodeSample),
							CtgyPurp: &pain_002_001_03.CategoryPurpose1Choice{
								Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						PmtMtd: common.MustToPaymentMethod4Code(common.PaymentMethod4CodeSample),
						MndtRltdInf: &pain_002_001_03.MandateRelatedInformation6{
							MndtId:    common.MustToMax35Text(common.Max35TextSample),
							DtOfSgntr: common.MustToISODate(common.ISODateSample),
							AmdmntInd: xsdt.MustToBoolean(xsdt.BooleanSample),
							AmdmntInfDtls: &pain_002_001_03.AmendmentInformationDetails6{
								OrgnlMndtId: common.MustToMax35Text(common.Max35TextSample),
								OrgnlCdtrSchmeId: &pain_002_001_03.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Id: &pain_002_001_03.Party6Choice{
										OrgId: &pain_002_001_03.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_03.PersonIdentification5{
											DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_03.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_03.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								OrgnlCdtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
											ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_03.PostalAddress6{
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
										Othr: &pain_002_001_03.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &pain_002_001_03.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_03.PostalAddress6{
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
								OrgnlCdtrAgtAcct: &pain_002_001_03.CashAccount16{
									Id: pain_002_001_03.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_03.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_03.CashAccountType2{
										Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
								},
								OrgnlDbtr: &pain_002_001_03.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Id: &pain_002_001_03.Party6Choice{
										OrgId: &pain_002_001_03.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_03.PersonIdentification5{
											DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_03.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_03.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								OrgnlDbtrAcct: &pain_002_001_03.CashAccount16{
									Id: pain_002_001_03.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_03.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_03.CashAccountType2{
										Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
								},
								OrgnlDbtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
									FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
										BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
										ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
											ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_03.PostalAddress6{
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
										Othr: &pain_002_001_03.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &pain_002_001_03.BranchData2{
										Id: common.MustToMax35Text(common.Max35TextSample),
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_03.PostalAddress6{
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
								OrgnlDbtrAgtAcct: &pain_002_001_03.CashAccount16{
									Id: pain_002_001_03.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_03.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_03.CashAccountType2{
										Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
								},
								OrgnlFnlColltnDt: common.MustToISODate(common.ISODateSample),
								OrgnlFrqcy:       common.MustToFrequency1Code(common.Frequency1CodeSample),
							},
							ElctrncSgntr: common.MustToMax1025Text(common.Max1025TextSample),
							FrstColltnDt: common.MustToISODate(common.ISODateSample),
							FnlColltnDt:  common.MustToISODate(common.ISODateSample),
							Frqcy:        common.MustToFrequency1Code(common.Frequency1CodeSample),
						},
						RmtInf: &pain_002_001_03.RemittanceInformation5{
							Ustrd: []common.Max140Text{
								common.MustToMax140Text(common.Max140TextSample),
							},
							Strd: []pain_002_001_03.StructuredRemittanceInformation7{{
								RfrdDocInf: []pain_002_001_03.ReferredDocumentInformation3{{
									Tp: &pain_002_001_03.ReferredDocumentType2{
										CdOrPrtry: pain_002_001_03.ReferredDocumentType1Choice{
											Cd:    common.MustToDocumentType5Code(common.DocumentType5CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Nb:     common.MustToMax35Text(common.Max35TextSample),
									RltdDt: common.MustToISODate(common.ISODateSample)},
								},
								RfrdDocAmt: &pain_002_001_03.RemittanceAmount1{
									DuePyblAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									DscntApldAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CdtNoteAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									TaxAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									AdjstmntAmtAndRsn: []pain_002_001_03.DocumentAdjustment1{{
										Amt: pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
										Rsn:       common.MustToMax4Text(common.Max4TextSample),
										AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
									},
									RmtdAmt: &pain_002_001_03.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
								},
								CdtrRefInf: &pain_002_001_03.CreditorReferenceInformation2{
									Tp: &pain_002_001_03.CreditorReferenceType2{
										CdOrPrtry: pain_002_001_03.CreditorReferenceType1Choice{
											Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Ref: common.MustToMax35Text(common.Max35TextSample),
								},
								Invcr: &pain_002_001_03.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Id: &pain_002_001_03.Party6Choice{
										OrgId: &pain_002_001_03.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_03.PersonIdentification5{
											DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_03.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_03.ContactDetails2{
										NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
										Nm:       common.MustToMax140Text(common.Max140TextSample),
										PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
										Othr:     common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Invcee: &pain_002_001_03.PartyIdentification32{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_03.PostalAddress6{
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
									Id: &pain_002_001_03.Party6Choice{
										OrgId: &pain_002_001_03.OrganisationIdentification4{
											BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
											Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_03.PersonIdentification5{
											DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_03.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_03.ContactDetails2{
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
						UltmtDbtr: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dbtr: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						DbtrAcct: &pain_002_001_03.CashAccount16{
							Id: pain_002_001_03.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_03.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_03.CashAccountType2{
								Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
						},
						DbtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
							FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
								BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
								ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
								Othr: &pain_002_001_03.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_03.BranchData2{
								Id: common.MustToMax35Text(common.Max35TextSample),
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
						DbtrAgtAcct: &pain_002_001_03.CashAccount16{
							Id: pain_002_001_03.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_03.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_03.CashAccountType2{
								Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
						},
						CdtrAgt: &pain_002_001_03.BranchAndFinancialInstitutionIdentification4{
							FinInstnId: pain_002_001_03.FinancialInstitutionIdentification7{
								BIC: common.MustToBICIdentifier(common.BICIdentifierSample),
								ClrSysMmbId: &pain_002_001_03.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_03.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
								Othr: &pain_002_001_03.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_03.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_03.BranchData2{
								Id: common.MustToMax35Text(common.Max35TextSample),
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_03.PostalAddress6{
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
						CdtrAgtAcct: &pain_002_001_03.CashAccount16{
							Id: pain_002_001_03.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_03.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_03.CashAccountType2{
								Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
						},
						Cdtr: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
						CdtrAcct: &pain_002_001_03.CashAccount16{
							Id: pain_002_001_03.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_03.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_03.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_03.CashAccountType2{
								Cd:    common.MustToCashAccountType4Code(common.CashAccountType4CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
						},
						UltmtCdtr: &pain_002_001_03.PartyIdentification32{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_03.PostalAddress6{
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
							Id: &pain_002_001_03.Party6Choice{
								OrgId: &pain_002_001_03.OrganisationIdentification4{
									BICOrBEI: common.MustToAnyBICIdentifier(common.AnyBICIdentifierSample),
									Othr: []pain_002_001_03.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_03.PersonIdentification5{
									DtAndPlcOfBirth: &pain_002_001_03.DateAndPlaceOfBirth{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_03.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_03.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_03.ContactDetails2{
								NmPrfx:   common.MustToNamePrefix1Code(common.NamePrefix1CodeSample),
								Nm:       common.MustToMax140Text(common.Max140TextSample),
								PhneNb:   common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr: common.MustToMax2048Text(common.Max2048TextSample),
								Othr:     common.MustToMax35Text(common.Max35TextSample),
							},
						},
					}},
				}},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_pain_002_001_03, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_pain_002_001_03)

}
