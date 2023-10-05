// Package pain_002_001_10_test
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_10_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.10/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.10/pain_002_001_10"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_pain_002_001_10 = "example-document-pain_002_001_10.xml"

func TestDocumentpain_002_001_10(t *testing.T) {

	d := pain_002_001_10.Document{
		CstmrPmtStsRpt: &pain_002_001_10.CustomerPaymentStatusReportV10{
			GrpHdr: &pain_002_001_10.GroupHeader86{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				InitgPty: &pain_002_001_10.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_002_001_10.PostalAddress24{
						AdrTp: &pain_002_001_10.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_002_001_10.GenericIdentification30{
								Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Dept:        common.MustToMax70Text(common.Max70TextSample),
						SubDept:     common.MustToMax70Text(common.Max70TextSample),
						StrtNm:      common.MustToMax70Text(common.Max70TextSample),
						BldgNb:      common.MustToMax16Text(common.Max16TextSample),
						BldgNm:      common.MustToMax35Text(common.Max35TextSample),
						Flr:         common.MustToMax70Text(common.Max70TextSample),
						PstBx:       common.MustToMax16Text(common.Max16TextSample),
						Room:        common.MustToMax70Text(common.Max70TextSample),
						PstCd:       common.MustToMax16Text(common.Max16TextSample),
						TwnNm:       common.MustToMax35Text(common.Max35TextSample),
						TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
						DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
						CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
						Ctry:        common.MustToCountryCode(common.CountryCodeSample),
						AdrLine: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					Id: &pain_002_001_10.Party38Choice{
						OrgId: &pain_002_001_10.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_002_001_10.PersonIdentification13{
							DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_002_001_10.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_002_001_10.Contact4{
						NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
						Nm:        common.MustToMax140Text(common.Max140TextSample),
						PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
						MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
						EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
						EmailPurp: common.MustToMax35Text(common.Max35TextSample),
						JobTitl:   common.MustToMax35Text(common.Max35TextSample),
						Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
						Dept:      common.MustToMax70Text(common.Max70TextSample),
						Othr: []pain_002_001_10.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				FwdgAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &pain_002_001_10.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_10.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				DbtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &pain_002_001_10.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_10.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				CdtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &pain_002_001_10.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_002_001_10.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
			},
			OrgnlGrpInfAndSts: &pain_002_001_10.OriginalGroupHeader17{
				OrgnlMsgId:   common.MustToMax35Text(common.Max35TextSample),
				OrgnlMsgNmId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlCreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				OrgnlNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
				OrgnlCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
				GrpSts:       common.MustToExternalPaymentGroupStatus1Code(common.ExternalPaymentGroupStatus1CodeSample),
				StsRsnInf: []pain_002_001_10.StatusReasonInformation12{{
					Orgtr: &pain_002_001_10.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &pain_002_001_10.Party38Choice{
							OrgId: &pain_002_001_10.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_002_001_10.PersonIdentification13{
								DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_002_001_10.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_002_001_10.Contact4{
							NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
							Nm:        common.MustToMax140Text(common.Max140TextSample),
							PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
							EmailPurp: common.MustToMax35Text(common.Max35TextSample),
							JobTitl:   common.MustToMax35Text(common.Max35TextSample),
							Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
							Dept:      common.MustToMax70Text(common.Max70TextSample),
							Othr: []pain_002_001_10.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					Rsn: &pain_002_001_10.StatusReason6Choice{
						Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					AddtlInf: []common.Max105Text{
						common.MustToMax105Text(common.Max105TextSample),
					}},
				},
				NbOfTxsPerSts: []pain_002_001_10.NumberOfTransactionsPerStatus5{{
					DtldNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
					DtldSts:     common.MustToExternalPaymentTransactionStatus1Code(common.ExternalPaymentTransactionStatus1CodeSample),
					DtldCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample)},
				},
			},
			OrgnlPmtInfAndSts: []pain_002_001_10.OriginalPaymentInstruction32{{
				OrgnlPmtInfId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlNbOfTxs:  common.MustToMax15NumericText(common.Max15NumericTextSample),
				OrgnlCtrlSum:  xsdt.MustToDecimal(xsdt.DecimalSample),
				PmtInfSts:     common.MustToExternalPaymentGroupStatus1Code(common.ExternalPaymentGroupStatus1CodeSample),
				StsRsnInf: []pain_002_001_10.StatusReasonInformation12{{
					Orgtr: &pain_002_001_10.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_002_001_10.PostalAddress24{
							AdrTp: &pain_002_001_10.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_002_001_10.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &pain_002_001_10.Party38Choice{
							OrgId: &pain_002_001_10.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_002_001_10.PersonIdentification13{
								DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_002_001_10.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_002_001_10.Contact4{
							NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
							Nm:        common.MustToMax140Text(common.Max140TextSample),
							PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
							EmailPurp: common.MustToMax35Text(common.Max35TextSample),
							JobTitl:   common.MustToMax35Text(common.Max35TextSample),
							Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
							Dept:      common.MustToMax70Text(common.Max70TextSample),
							Othr: []pain_002_001_10.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					Rsn: &pain_002_001_10.StatusReason6Choice{
						Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					AddtlInf: []common.Max105Text{
						common.MustToMax105Text(common.Max105TextSample),
					}},
				},
				NbOfTxsPerSts: []pain_002_001_10.NumberOfTransactionsPerStatus5{{
					DtldNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
					DtldSts:     common.MustToExternalPaymentTransactionStatus1Code(common.ExternalPaymentTransactionStatus1CodeSample),
					DtldCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample)},
				},
				TxInfAndSts: []pain_002_001_10.PaymentTransaction105{{
					StsId:           common.MustToMax35Text(common.Max35TextSample),
					OrgnlInstrId:    common.MustToMax35Text(common.Max35TextSample),
					OrgnlEndToEndId: common.MustToMax35Text(common.Max35TextSample),
					OrgnlUETR:       common.MustToUUIDv4Identifier(common.UUIDv4IdentifierSample),
					TxSts:           common.MustToExternalPaymentTransactionStatus1Code(common.ExternalPaymentTransactionStatus1CodeSample),
					StsRsnInf: []pain_002_001_10.StatusReasonInformation12{{
						Orgtr: &pain_002_001_10.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_10.PostalAddress24{
								AdrTp: &pain_002_001_10.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_002_001_10.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_002_001_10.Party38Choice{
								OrgId: &pain_002_001_10.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_10.PersonIdentification13{
									DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_10.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_10.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []pain_002_001_10.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Rsn: &pain_002_001_10.StatusReason6Choice{
							Cd:    common.MustToExternalStatusReason1Code(common.ExternalStatusReason1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						AddtlInf: []common.Max105Text{
							common.MustToMax105Text(common.Max105TextSample),
						}},
					},
					ChrgsInf: []pain_002_001_10.Charges7{{
						Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &pain_002_001_10.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_10.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						}},
					},
					TrckrData: &pain_002_001_10.TrackerData1{
						ConfdDt: &pain_002_001_10.DateAndDateTime2Choice{
							Dt:   common.MustToISODate(common.ISODateSample),
							DtTm: common.MustToISODateTime(common.ISODateTimeSample),
						},
						ConfdAmt: &pain_002_001_10.ActiveCurrencyAndAmount{
							Ccy:   common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						TrckrRcrd: []pain_002_001_10.TrackerRecord1{{
							Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
							ChrgsAmt: &pain_002_001_10.ActiveCurrencyAndAmount{
								Ccy:   common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							XchgRateData: &pain_002_001_10.CurrencyExchange13{
								SrcCcy:   common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
								TrgtCcy:  common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
								XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
								UnitCcy:  common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
							}},
						},
					},
					AccptncDtTm: common.MustToISODateTime(common.ISODateTimeSample),
					AcctSvcrRef: common.MustToMax35Text(common.Max35TextSample),
					ClrSysRef:   common.MustToMax35Text(common.Max35TextSample),
					OrgnlTxRef: &pain_002_001_10.OriginalTransactionReference28{
						IntrBkSttlmAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Amt: &pain_002_001_10.AmountType4Choice{
							InstdAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							EqvtAmt: &pain_002_001_10.EquivalentAmount2{
								Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							},
						},
						IntrBkSttlmDt: common.MustToISODate(common.ISODateSample),
						ReqdColltnDt:  common.MustToISODate(common.ISODateSample),
						ReqdExctnDt: &pain_002_001_10.DateAndDateTime2Choice{
							Dt:   common.MustToISODate(common.ISODateSample),
							DtTm: common.MustToISODateTime(common.ISODateTimeSample),
						},
						CdtrSchmeId: &pain_002_001_10.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_002_001_10.PostalAddress24{
								AdrTp: &pain_002_001_10.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_002_001_10.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &pain_002_001_10.Party38Choice{
								OrgId: &pain_002_001_10.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &pain_002_001_10.PersonIdentification13{
									DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []pain_002_001_10.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &pain_002_001_10.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []pain_002_001_10.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						SttlmInf: &pain_002_001_10.SettlementInstruction7{
							SttlmMtd: common.MustToSettlementMethod1Code(common.SettlementMethod1CodeSample),
							SttlmAcct: &pain_002_001_10.CashAccount38{
								Id: &pain_002_001_10.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_10.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_10.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_002_001_10.ProxyAccountIdentification1{
									Tp: &pain_002_001_10.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							ClrSys: &pain_002_001_10.ClearingSystemIdentification3Choice{
								Cd:    common.MustToExternalCashClearingSystem1Code(common.ExternalCashClearingSystem1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							InstgRmbrsmntAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							InstgRmbrsmntAgtAcct: &pain_002_001_10.CashAccount38{
								Id: &pain_002_001_10.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_10.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_10.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_002_001_10.ProxyAccountIdentification1{
									Tp: &pain_002_001_10.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							InstdRmbrsmntAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							InstdRmbrsmntAgtAcct: &pain_002_001_10.CashAccount38{
								Id: &pain_002_001_10.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_10.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_10.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_002_001_10.ProxyAccountIdentification1{
									Tp: &pain_002_001_10.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							ThrdRmbrsmntAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							ThrdRmbrsmntAgtAcct: &pain_002_001_10.CashAccount38{
								Id: &pain_002_001_10.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &pain_002_001_10.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &pain_002_001_10.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &pain_002_001_10.ProxyAccountIdentification1{
									Tp: &pain_002_001_10.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
						},
						PmtTpInf: &pain_002_001_10.PaymentTypeInformation27{
							InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
							ClrChanl:  common.MustToClearingChannel2Code(common.ClearingChannel2CodeSample),
							SvcLvl: []pain_002_001_10.ServiceLevel8Choice{{
								Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample)},
							},
							LclInstrm: &pain_002_001_10.LocalInstrument2Choice{
								Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							SeqTp: common.MustToSequenceType3Code(common.SequenceType3CodeSample),
							CtgyPurp: &pain_002_001_10.CategoryPurpose1Choice{
								Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						PmtMtd: common.MustToPaymentMethod4Code(common.PaymentMethod4CodeSample),
						MndtRltdInf: &pain_002_001_10.MandateRelatedInformation14{
							MndtId:    common.MustToMax35Text(common.Max35TextSample),
							DtOfSgntr: common.MustToISODate(common.ISODateSample),
							AmdmntInd: xsdt.MustToBoolean(xsdt.BooleanSample),
							AmdmntInfDtls: &pain_002_001_10.AmendmentInformationDetails13{
								OrgnlMndtId: common.MustToMax35Text(common.Max35TextSample),
								OrgnlCdtrSchmeId: &pain_002_001_10.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &pain_002_001_10.Party38Choice{
										OrgId: &pain_002_001_10.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_10.PersonIdentification13{
											DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_10.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_10.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []pain_002_001_10.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								OrgnlCdtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
									FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
										BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
										ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
											ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Nm:  common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &pain_002_001_10.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &pain_002_001_10.BranchData3{
										Id:  common.MustToMax35Text(common.Max35TextSample),
										LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Nm:  common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								OrgnlCdtrAgtAcct: &pain_002_001_10.CashAccount38{
									Id: &pain_002_001_10.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_10.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_10.CashAccountType2Choice{
										Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
									Prxy: &pain_002_001_10.ProxyAccountIdentification1{
										Tp: &pain_002_001_10.ProxyAccountType1Choice{
											Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Id: common.MustToMax2048Text(common.Max2048TextSample),
									},
								},
								OrgnlDbtr: &pain_002_001_10.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &pain_002_001_10.Party38Choice{
										OrgId: &pain_002_001_10.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_10.PersonIdentification13{
											DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_10.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_10.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []pain_002_001_10.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								OrgnlDbtrAcct: &pain_002_001_10.CashAccount38{
									Id: &pain_002_001_10.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_10.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_10.CashAccountType2Choice{
										Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
									Prxy: &pain_002_001_10.ProxyAccountIdentification1{
										Tp: &pain_002_001_10.ProxyAccountType1Choice{
											Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Id: common.MustToMax2048Text(common.Max2048TextSample),
									},
								},
								OrgnlDbtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
									FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
										BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
										ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
											ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
												Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											MmbId: common.MustToMax35Text(common.Max35TextSample),
										},
										LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Nm:  common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Othr: &pain_002_001_10.GenericFinancialIdentification1{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									BrnchId: &pain_002_001_10.BranchData3{
										Id:  common.MustToMax35Text(common.Max35TextSample),
										LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Nm:  common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
									},
								},
								OrgnlDbtrAgtAcct: &pain_002_001_10.CashAccount38{
									Id: &pain_002_001_10.AccountIdentification4Choice{
										IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
										Othr: &pain_002_001_10.GenericAccountIdentification1{
											Id: common.MustToMax34Text(common.Max34TextSample),
											SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
												Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Tp: &pain_002_001_10.CashAccountType2Choice{
										Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Nm:  common.MustToMax70Text(common.Max70TextSample),
									Prxy: &pain_002_001_10.ProxyAccountIdentification1{
										Tp: &pain_002_001_10.ProxyAccountType1Choice{
											Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Id: common.MustToMax2048Text(common.Max2048TextSample),
									},
								},
								OrgnlFnlColltnDt: common.MustToISODate(common.ISODateSample),
								OrgnlFrqcy: &pain_002_001_10.Frequency36Choice{
									Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
									Prd: &pain_002_001_10.FrequencyPeriod1{
										Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
										CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									PtInTm: &pain_002_001_10.FrequencyAndMoment1{
										Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
										PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
									},
								},
								OrgnlRsn: &pain_002_001_10.MandateSetupReason1Choice{
									Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
									Prtry: common.MustToMax70Text(common.Max70TextSample),
								},
								OrgnlTrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
							},
							ElctrncSgntr: common.MustToMax1025Text(common.Max1025TextSample),
							FrstColltnDt: common.MustToISODate(common.ISODateSample),
							FnlColltnDt:  common.MustToISODate(common.ISODateSample),
							Frqcy: &pain_002_001_10.Frequency36Choice{
								Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
								Prd: &pain_002_001_10.FrequencyPeriod1{
									Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
									CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								PtInTm: &pain_002_001_10.FrequencyAndMoment1{
									Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
									PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
								},
							},
							Rsn: &pain_002_001_10.MandateSetupReason1Choice{
								Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
								Prtry: common.MustToMax70Text(common.Max70TextSample),
							},
							TrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
						},
						RmtInf: &pain_002_001_10.RemittanceInformation16{
							Ustrd: []common.Max140Text{
								common.MustToMax140Text(common.Max140TextSample),
							},
							Strd: []pain_002_001_10.StructuredRemittanceInformation16{{
								RfrdDocInf: []pain_002_001_10.ReferredDocumentInformation7{{
									Tp: &pain_002_001_10.ReferredDocumentType4{
										CdOrPrtry: &pain_002_001_10.ReferredDocumentType3Choice{
											Cd:    common.MustToDocumentType6Code(common.DocumentType6CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Nb:     common.MustToMax35Text(common.Max35TextSample),
									RltdDt: common.MustToISODate(common.ISODateSample),
									LineDtls: []pain_002_001_10.DocumentLineInformation1{{
										Id: []pain_002_001_10.DocumentLineIdentification1{{
											Tp: &pain_002_001_10.DocumentLineType1{
												CdOrPrtry: &pain_002_001_10.DocumentLineType1Choice{
													Cd:    common.MustToExternalDocumentLineType1Code(common.ExternalDocumentLineType1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample),
											},
											Nb:     common.MustToMax35Text(common.Max35TextSample),
											RltdDt: common.MustToISODate(common.ISODateSample)},
										},
										Desc: common.MustToMax2048Text(common.Max2048TextSample),
										Amt: &pain_002_001_10.RemittanceAmount3{
											DuePyblAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											DscntApldAmt: []pain_002_001_10.DiscountAmountAndType1{{
												Tp: &pain_002_001_10.DiscountAmountType1Choice{
													Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
													Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
													Value: xsdt.MustToDecimal(xsdt.DecimalSample),
												}},
											},
											CdtNoteAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											TaxAmt: []pain_002_001_10.TaxAmountAndType1{{
												Tp: &pain_002_001_10.TaxAmountType1Choice{
													Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
													Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
													Value: xsdt.MustToDecimal(xsdt.DecimalSample),
												}},
											},
											AdjstmntAmtAndRsn: []pain_002_001_10.DocumentAdjustment1{{
												Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
													Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
													Value: xsdt.MustToDecimal(xsdt.DecimalSample),
												},
												CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
												Rsn:       common.MustToMax4Text(common.Max4TextSample),
												AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
											},
											RmtdAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
										}},
									}},
								},
								RfrdDocAmt: &pain_002_001_10.RemittanceAmount2{
									DuePyblAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									DscntApldAmt: []pain_002_001_10.DiscountAmountAndType1{{
										Tp: &pain_002_001_10.DiscountAmountType1Choice{
											Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										}},
									},
									CdtNoteAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									TaxAmt: []pain_002_001_10.TaxAmountAndType1{{
										Tp: &pain_002_001_10.TaxAmountType1Choice{
											Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										}},
									},
									AdjstmntAmtAndRsn: []pain_002_001_10.DocumentAdjustment1{{
										Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
										Rsn:       common.MustToMax4Text(common.Max4TextSample),
										AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
									},
									RmtdAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
								},
								CdtrRefInf: &pain_002_001_10.CreditorReferenceInformation2{
									Tp: &pain_002_001_10.CreditorReferenceType2{
										CdOrPrtry: &pain_002_001_10.CreditorReferenceType1Choice{
											Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Ref: common.MustToMax35Text(common.Max35TextSample),
								},
								Invcr: &pain_002_001_10.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &pain_002_001_10.Party38Choice{
										OrgId: &pain_002_001_10.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_10.PersonIdentification13{
											DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_10.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_10.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []pain_002_001_10.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								Invcee: &pain_002_001_10.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &pain_002_001_10.Party38Choice{
										OrgId: &pain_002_001_10.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_002_001_10.PersonIdentification13{
											DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_002_001_10.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_002_001_10.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []pain_002_001_10.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								TaxRmt: &pain_002_001_10.TaxInformation7{
									Cdtr: &pain_002_001_10.TaxParty1{
										TaxId:  common.MustToMax35Text(common.Max35TextSample),
										RegnId: common.MustToMax35Text(common.Max35TextSample),
										TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									},
									Dbtr: &pain_002_001_10.TaxParty2{
										TaxId:  common.MustToMax35Text(common.Max35TextSample),
										RegnId: common.MustToMax35Text(common.Max35TextSample),
										TaxTp:  common.MustToMax35Text(common.Max35TextSample),
										Authstn: &pain_002_001_10.TaxAuthorisation1{
											Titl: common.MustToMax35Text(common.Max35TextSample),
											Nm:   common.MustToMax140Text(common.Max140TextSample),
										},
									},
									UltmtDbtr: &pain_002_001_10.TaxParty2{
										TaxId:  common.MustToMax35Text(common.Max35TextSample),
										RegnId: common.MustToMax35Text(common.Max35TextSample),
										TaxTp:  common.MustToMax35Text(common.Max35TextSample),
										Authstn: &pain_002_001_10.TaxAuthorisation1{
											Titl: common.MustToMax35Text(common.Max35TextSample),
											Nm:   common.MustToMax140Text(common.Max140TextSample),
										},
									},
									AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
									RefNb:      common.MustToMax140Text(common.Max140TextSample),
									Mtd:        common.MustToMax35Text(common.Max35TextSample),
									TtlTaxblBaseAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									TtlTaxAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									Dt:    common.MustToISODate(common.ISODateSample),
									SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
									Rcrd: []pain_002_001_10.TaxRecord2{{
										Tp:       common.MustToMax35Text(common.Max35TextSample),
										Ctgy:     common.MustToMax35Text(common.Max35TextSample),
										CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
										DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
										CertId:   common.MustToMax35Text(common.Max35TextSample),
										FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
										Prd: &pain_002_001_10.TaxPeriod2{
											Yr: common.MustToISODate(common.ISODateSample),
											Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
											FrToDt: &pain_002_001_10.DatePeriod2{
												FrDt: common.MustToISODate(common.ISODateSample),
												ToDt: common.MustToISODate(common.ISODateSample),
											},
										},
										TaxAmt: &pain_002_001_10.TaxAmount2{
											Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
											TaxblBaseAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											TtlAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											Dtls: []pain_002_001_10.TaxRecordDetails2{{
												Prd: &pain_002_001_10.TaxPeriod2{
													Yr: common.MustToISODate(common.ISODateSample),
													Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
													FrToDt: &pain_002_001_10.DatePeriod2{
														FrDt: common.MustToISODate(common.ISODateSample),
														ToDt: common.MustToISODate(common.ISODateSample),
													},
												},
												Amt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
													Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
													Value: xsdt.MustToDecimal(xsdt.DecimalSample),
												}},
											},
										},
										AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
									},
								},
								GrnshmtRmt: &pain_002_001_10.Garnishment3{
									Tp: &pain_002_001_10.GarnishmentType1{
										CdOrPrtry: &pain_002_001_10.GarnishmentType1Choice{
											Cd:    common.MustToExternalGarnishmentType1Code(common.ExternalGarnishmentType1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
									Grnshee: &pain_002_001_10.PartyIdentification135{
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Id: &pain_002_001_10.Party38Choice{
											OrgId: &pain_002_001_10.OrganisationIdentification29{
												AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
												LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
												Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
											PrvtId: &pain_002_001_10.PersonIdentification13{
												DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
													BirthDt:     common.MustToISODate(common.ISODateSample),
													PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
												},
												Othr: []pain_002_001_10.GenericPersonIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
										},
										CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
										CtctDtls: &pain_002_001_10.Contact4{
											NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
											Nm:        common.MustToMax140Text(common.Max140TextSample),
											PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
											FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
											EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
											EmailPurp: common.MustToMax35Text(common.Max35TextSample),
											JobTitl:   common.MustToMax35Text(common.Max35TextSample),
											Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
											Dept:      common.MustToMax70Text(common.Max70TextSample),
											Othr: []pain_002_001_10.OtherContact1{{
												ChanlTp: common.MustToMax4Text(common.Max4TextSample),
												Id:      common.MustToMax128Text(common.Max128TextSample)},
											},
											PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
										},
									},
									GrnshmtAdmstr: &pain_002_001_10.PartyIdentification135{
										Nm: common.MustToMax140Text(common.Max140TextSample),
										PstlAdr: &pain_002_001_10.PostalAddress24{
											AdrTp: &pain_002_001_10.AddressType3Choice{
												Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
												Prtry: &pain_002_001_10.GenericIdentification30{
													Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
													Issr:    common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: common.MustToMax35Text(common.Max35TextSample),
												},
											},
											Dept:        common.MustToMax70Text(common.Max70TextSample),
											SubDept:     common.MustToMax70Text(common.Max70TextSample),
											StrtNm:      common.MustToMax70Text(common.Max70TextSample),
											BldgNb:      common.MustToMax16Text(common.Max16TextSample),
											BldgNm:      common.MustToMax35Text(common.Max35TextSample),
											Flr:         common.MustToMax70Text(common.Max70TextSample),
											PstBx:       common.MustToMax16Text(common.Max16TextSample),
											Room:        common.MustToMax70Text(common.Max70TextSample),
											PstCd:       common.MustToMax16Text(common.Max16TextSample),
											TwnNm:       common.MustToMax35Text(common.Max35TextSample),
											TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
											DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
											CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
											Ctry:        common.MustToCountryCode(common.CountryCodeSample),
											AdrLine: []common.Max70Text{
												common.MustToMax70Text(common.Max70TextSample),
											},
										},
										Id: &pain_002_001_10.Party38Choice{
											OrgId: &pain_002_001_10.OrganisationIdentification29{
												AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
												LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
												Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
											PrvtId: &pain_002_001_10.PersonIdentification13{
												DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
													BirthDt:     common.MustToISODate(common.ISODateSample),
													PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
													CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
												},
												Othr: []pain_002_001_10.GenericPersonIdentification1{{
													Id: common.MustToMax35Text(common.Max35TextSample),
													SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
														Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
														Prtry: common.MustToMax35Text(common.Max35TextSample),
													},
													Issr: common.MustToMax35Text(common.Max35TextSample)},
												},
											},
										},
										CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
										CtctDtls: &pain_002_001_10.Contact4{
											NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
											Nm:        common.MustToMax140Text(common.Max140TextSample),
											PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
											MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
											FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
											EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
											EmailPurp: common.MustToMax35Text(common.Max35TextSample),
											JobTitl:   common.MustToMax35Text(common.Max35TextSample),
											Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
											Dept:      common.MustToMax70Text(common.Max70TextSample),
											Othr: []pain_002_001_10.OtherContact1{{
												ChanlTp: common.MustToMax4Text(common.Max4TextSample),
												Id:      common.MustToMax128Text(common.Max128TextSample)},
											},
											PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
										},
									},
									RefNb: common.MustToMax140Text(common.Max140TextSample),
									Dt:    common.MustToISODate(common.ISODateSample),
									RmtdAmt: &pain_002_001_10.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									FmlyMdclInsrncInd: xsdt.MustToBoolean(xsdt.BooleanSample),
									MplyeeTermntnInd:  xsdt.MustToBoolean(xsdt.BooleanSample),
								},
								AddtlRmtInf: []common.Max140Text{
									common.MustToMax140Text(common.Max140TextSample),
								}},
							},
						},
						UltmtDbtr: &pain_002_001_10.Party40Choice{
							Pty: &pain_002_001_10.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_002_001_10.Party38Choice{
									OrgId: &pain_002_001_10.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_002_001_10.PersonIdentification13{
										DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_002_001_10.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_002_001_10.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_002_001_10.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
						},
						Dbtr: &pain_002_001_10.Party40Choice{
							Pty: &pain_002_001_10.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_002_001_10.Party38Choice{
									OrgId: &pain_002_001_10.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_002_001_10.PersonIdentification13{
										DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_002_001_10.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_002_001_10.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_002_001_10.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
						},
						DbtrAcct: &pain_002_001_10.CashAccount38{
							Id: &pain_002_001_10.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_10.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_10.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &pain_002_001_10.ProxyAccountIdentification1{
								Tp: &pain_002_001_10.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						DbtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &pain_002_001_10.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_10.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						DbtrAgtAcct: &pain_002_001_10.CashAccount38{
							Id: &pain_002_001_10.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_10.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_10.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &pain_002_001_10.ProxyAccountIdentification1{
								Tp: &pain_002_001_10.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						CdtrAgt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
									ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &pain_002_001_10.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &pain_002_001_10.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						CdtrAgtAcct: &pain_002_001_10.CashAccount38{
							Id: &pain_002_001_10.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_10.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_10.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &pain_002_001_10.ProxyAccountIdentification1{
								Tp: &pain_002_001_10.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						Cdtr: &pain_002_001_10.Party40Choice{
							Pty: &pain_002_001_10.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_002_001_10.Party38Choice{
									OrgId: &pain_002_001_10.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_002_001_10.PersonIdentification13{
										DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_002_001_10.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_002_001_10.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_002_001_10.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
						},
						CdtrAcct: &pain_002_001_10.CashAccount38{
							Id: &pain_002_001_10.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &pain_002_001_10.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &pain_002_001_10.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &pain_002_001_10.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &pain_002_001_10.ProxyAccountIdentification1{
								Tp: &pain_002_001_10.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						UltmtCdtr: &pain_002_001_10.Party40Choice{
							Pty: &pain_002_001_10.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_002_001_10.PostalAddress24{
									AdrTp: &pain_002_001_10.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_002_001_10.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &pain_002_001_10.Party38Choice{
									OrgId: &pain_002_001_10.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_002_001_10.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_002_001_10.PersonIdentification13{
										DtAndPlcOfBirth: &pain_002_001_10.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_002_001_10.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_002_001_10.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_002_001_10.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []pain_002_001_10.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Agt: &pain_002_001_10.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: &pain_002_001_10.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &pain_002_001_10.ClearingSystemMemberIdentification2{
										ClrSysId: &pain_002_001_10.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &pain_002_001_10.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &pain_002_001_10.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &pain_002_001_10.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_002_001_10.PostalAddress24{
										AdrTp: &pain_002_001_10.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_002_001_10.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
						},
						Purp: &pain_002_001_10.Purpose2Choice{
							Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					SplmtryData: []pain_002_001_10.SupplementaryData1{{
						PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
						Envlp: &pain_002_001_10.SupplementaryDataEnvelope1{
							Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
						}},
					}},
				}},
			},
			SplmtryData: []pain_002_001_10.SupplementaryData1{{
				PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
				Envlp: &pain_002_001_10.SupplementaryDataEnvelope1{
					Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
				}},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_pain_002_001_10, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_pain_002_001_10)

}
