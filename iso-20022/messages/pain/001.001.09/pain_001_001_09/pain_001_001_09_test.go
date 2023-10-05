// Package pain_001_001_09_test
// Do not Edit. This stuff it's been automatically generated.
package pain_001_001_09_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.09/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.09/pain_001_001_09"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_pain_001_001_09 = "example-document-pain_001_001_09.xml"

func TestDocumentpain_001_001_09(t *testing.T) {

	d := pain_001_001_09.Document{
		CstmrCdtTrfInitn: &pain_001_001_09.CustomerCreditTransferInitiationV09{
			GrpHdr: &pain_001_001_09.GroupHeader85{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				Authstn: []pain_001_001_09.Authorisation1Choice{{
					Cd:    common.MustToAuthorisation1Code(common.Authorisation1CodeSample),
					Prtry: common.MustToMax128Text(common.Max128TextSample)},
				},
				NbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
				CtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
				InitgPty: &pain_001_001_09.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_09.PostalAddress24{
						AdrTp: &pain_001_001_09.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_001_001_09.GenericIdentification30{
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
					Id: &pain_001_001_09.Party38Choice{
						OrgId: &pain_001_001_09.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_09.PersonIdentification13{
							DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_09.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_09.Contact4{
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
						Othr: []pain_001_001_09.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				FwdgAgt: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Othr: &pain_001_001_09.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_09.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
			PmtInf: []pain_001_001_09.PaymentInstruction30{{
				PmtInfId:  common.MustToMax35Text(common.Max35TextSample),
				PmtMtd:    common.MustToPaymentMethod3Code(common.PaymentMethod3CodeSample),
				BtchBookg: xsdt.MustToBoolean(xsdt.BooleanSample),
				NbOfTxs:   common.MustToMax15NumericText(common.Max15NumericTextSample),
				CtrlSum:   xsdt.MustToDecimal(xsdt.DecimalSample),
				PmtTpInf: &pain_001_001_09.PaymentTypeInformation26{
					InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
					SvcLvl: []pain_001_001_09.ServiceLevel8Choice{{
						Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample)},
					},
					LclInstrm: &pain_001_001_09.LocalInstrument2Choice{
						Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					CtgyPurp: &pain_001_001_09.CategoryPurpose1Choice{
						Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				ReqdExctnDt: &pain_001_001_09.DateAndDateTime2Choice{
					Dt:   common.MustToISODate(common.ISODateSample),
					DtTm: common.MustToISODateTime(common.ISODateTimeSample),
				},
				PoolgAdjstmntDt: common.MustToISODate(common.ISODateSample),
				Dbtr: &pain_001_001_09.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_09.PostalAddress24{
						AdrTp: &pain_001_001_09.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_001_001_09.GenericIdentification30{
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
					Id: &pain_001_001_09.Party38Choice{
						OrgId: &pain_001_001_09.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_09.PersonIdentification13{
							DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_09.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_09.Contact4{
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
						Othr: []pain_001_001_09.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				DbtrAcct: &pain_001_001_09.CashAccount38{
					Id: &pain_001_001_09.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_09.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_09.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_001_001_09.ProxyAccountIdentification1{
						Tp: &pain_001_001_09.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				DbtrAgt: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Othr: &pain_001_001_09.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_09.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
				DbtrAgtAcct: &pain_001_001_09.CashAccount38{
					Id: &pain_001_001_09.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_09.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_09.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_001_001_09.ProxyAccountIdentification1{
						Tp: &pain_001_001_09.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				InstrForDbtrAgt: common.MustToMax140Text(common.Max140TextSample),
				UltmtDbtr: &pain_001_001_09.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_001_001_09.PostalAddress24{
						AdrTp: &pain_001_001_09.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_001_001_09.GenericIdentification30{
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
					Id: &pain_001_001_09.Party38Choice{
						OrgId: &pain_001_001_09.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_001_001_09.PersonIdentification13{
							DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_001_001_09.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_001_001_09.Contact4{
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
						Othr: []pain_001_001_09.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
				ChrgsAcct: &pain_001_001_09.CashAccount38{
					Id: &pain_001_001_09.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_001_001_09.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_001_001_09.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_001_001_09.ProxyAccountIdentification1{
						Tp: &pain_001_001_09.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				ChrgsAcctAgt: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Othr: &pain_001_001_09.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_001_001_09.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
				CdtTrfTxInf: []pain_001_001_09.CreditTransferTransaction34{{
					PmtId: &pain_001_001_09.PaymentIdentification6{
						InstrId:    common.MustToMax35Text(common.Max35TextSample),
						EndToEndId: common.MustToMax35Text(common.Max35TextSample),
						UETR:       common.MustToUUIDv4Identifier(common.UUIDv4IdentifierSample),
					},
					PmtTpInf: &pain_001_001_09.PaymentTypeInformation26{
						InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
						SvcLvl: []pain_001_001_09.ServiceLevel8Choice{{
							Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample)},
						},
						LclInstrm: &pain_001_001_09.LocalInstrument2Choice{
							Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						CtgyPurp: &pain_001_001_09.CategoryPurpose1Choice{
							Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Amt: &pain_001_001_09.AmountType4Choice{
						InstdAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						EqvtAmt: &pain_001_001_09.EquivalentAmount2{
							Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						},
					},
					XchgRateInf: &pain_001_001_09.ExchangeRate1{
						UnitCcy:  common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						XchgRate: xsdt.MustToDecimal(xsdt.DecimalSample),
						RateTp:   common.MustToExchangeRateType1Code(common.ExchangeRateType1CodeSample),
						CtrctId:  common.MustToMax35Text(common.Max35TextSample),
					},
					ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
					ChqInstr: &pain_001_001_09.Cheque11{
						ChqTp: common.MustToChequeType2Code(common.ChequeType2CodeSample),
						ChqNb: common.MustToMax35Text(common.Max35TextSample),
						ChqFr: &pain_001_001_09.NameAndAddress16{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							Adr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
						DlvryMtd: &pain_001_001_09.ChequeDeliveryMethod1Choice{
							Cd:    common.MustToChequeDelivery1Code(common.ChequeDelivery1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						DlvrTo: &pain_001_001_09.NameAndAddress16{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							Adr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
						InstrPrty:  common.MustToPriority2Code(common.Priority2CodeSample),
						ChqMtrtyDt: common.MustToISODate(common.ISODateSample),
						FrmsCd:     common.MustToMax35Text(common.Max35TextSample),
						MemoFld: []common.Max35Text{
							common.MustToMax35Text(common.Max35TextSample),
						},
						RgnlClrZone: common.MustToMax35Text(common.Max35TextSample),
						PrtLctn:     common.MustToMax35Text(common.Max35TextSample),
						Sgntr: []common.Max70Text{
							common.MustToMax70Text(common.Max70TextSample),
						},
					},
					UltmtDbtr: &pain_001_001_09.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Id: &pain_001_001_09.Party38Choice{
							OrgId: &pain_001_001_09.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_001_001_09.PersonIdentification13{
								DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_001_001_09.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_001_001_09.Contact4{
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
							Othr: []pain_001_001_09.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					IntrmyAgt1: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
							Othr: &pain_001_001_09.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_001_001_09.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
					IntrmyAgt1Acct: &pain_001_001_09.CashAccount38{
						Id: &pain_001_001_09.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_001_001_09.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_001_001_09.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_001_001_09.ProxyAccountIdentification1{
							Tp: &pain_001_001_09.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					IntrmyAgt2: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
							Othr: &pain_001_001_09.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_001_001_09.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
					IntrmyAgt2Acct: &pain_001_001_09.CashAccount38{
						Id: &pain_001_001_09.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_001_001_09.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_001_001_09.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_001_001_09.ProxyAccountIdentification1{
							Tp: &pain_001_001_09.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					IntrmyAgt3: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
							Othr: &pain_001_001_09.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_001_001_09.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
					IntrmyAgt3Acct: &pain_001_001_09.CashAccount38{
						Id: &pain_001_001_09.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_001_001_09.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_001_001_09.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_001_001_09.ProxyAccountIdentification1{
							Tp: &pain_001_001_09.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					CdtrAgt: &pain_001_001_09.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_001_001_09.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_001_001_09.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_001_001_09.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
							Othr: &pain_001_001_09.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_001_001_09.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_001_001_09.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_001_001_09.PostalAddress24{
								AdrTp: &pain_001_001_09.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_001_001_09.GenericIdentification30{
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
					CdtrAgtAcct: &pain_001_001_09.CashAccount38{
						Id: &pain_001_001_09.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_001_001_09.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_001_001_09.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_001_001_09.ProxyAccountIdentification1{
							Tp: &pain_001_001_09.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					Cdtr: &pain_001_001_09.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Id: &pain_001_001_09.Party38Choice{
							OrgId: &pain_001_001_09.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_001_001_09.PersonIdentification13{
								DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_001_001_09.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_001_001_09.Contact4{
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
							Othr: []pain_001_001_09.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					CdtrAcct: &pain_001_001_09.CashAccount38{
						Id: &pain_001_001_09.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_001_001_09.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_001_001_09.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_001_001_09.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_001_001_09.ProxyAccountIdentification1{
							Tp: &pain_001_001_09.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					UltmtCdtr: &pain_001_001_09.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_001_001_09.PostalAddress24{
							AdrTp: &pain_001_001_09.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_001_001_09.GenericIdentification30{
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
						Id: &pain_001_001_09.Party38Choice{
							OrgId: &pain_001_001_09.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_001_001_09.PersonIdentification13{
								DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_001_001_09.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_001_001_09.Contact4{
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
							Othr: []pain_001_001_09.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					InstrForCdtrAgt: []pain_001_001_09.InstructionForCreditorAgent1{{
						Cd:       common.MustToInstruction3Code(common.Instruction3CodeSample),
						InstrInf: common.MustToMax140Text(common.Max140TextSample)},
					},
					InstrForDbtrAgt: common.MustToMax140Text(common.Max140TextSample),
					Purp: &pain_001_001_09.Purpose2Choice{
						Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					RgltryRptg: []pain_001_001_09.RegulatoryReporting3{{
						DbtCdtRptgInd: common.MustToRegulatoryReportingType1Code(common.RegulatoryReportingType1CodeSample),
						Authrty: &pain_001_001_09.RegulatoryAuthority2{
							Nm:   common.MustToMax140Text(common.Max140TextSample),
							Ctry: common.MustToCountryCode(common.CountryCodeSample),
						},
						Dtls: []pain_001_001_09.StructuredRegulatoryReporting3{{
							Tp:   common.MustToMax35Text(common.Max35TextSample),
							Dt:   common.MustToISODate(common.ISODateSample),
							Ctry: common.MustToCountryCode(common.CountryCodeSample),
							Cd:   common.MustToMax10Text(common.Max10TextSample),
							Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Inf: []common.Max35Text{
								common.MustToMax35Text(common.Max35TextSample),
							}},
						}},
					},
					Tax: &pain_001_001_09.TaxInformation8{
						Cdtr: &pain_001_001_09.TaxParty1{
							TaxId:  common.MustToMax35Text(common.Max35TextSample),
							RegnId: common.MustToMax35Text(common.Max35TextSample),
							TaxTp:  common.MustToMax35Text(common.Max35TextSample),
						},
						Dbtr: &pain_001_001_09.TaxParty2{
							TaxId:  common.MustToMax35Text(common.Max35TextSample),
							RegnId: common.MustToMax35Text(common.Max35TextSample),
							TaxTp:  common.MustToMax35Text(common.Max35TextSample),
							Authstn: &pain_001_001_09.TaxAuthorisation1{
								Titl: common.MustToMax35Text(common.Max35TextSample),
								Nm:   common.MustToMax140Text(common.Max140TextSample),
							},
						},
						AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
						RefNb:      common.MustToMax140Text(common.Max140TextSample),
						Mtd:        common.MustToMax35Text(common.Max35TextSample),
						TtlTaxblBaseAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						TtlTaxAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Dt:    common.MustToISODate(common.ISODateSample),
						SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
						Rcrd: []pain_001_001_09.TaxRecord2{{
							Tp:       common.MustToMax35Text(common.Max35TextSample),
							Ctgy:     common.MustToMax35Text(common.Max35TextSample),
							CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
							DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
							CertId:   common.MustToMax35Text(common.Max35TextSample),
							FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
							Prd: &pain_001_001_09.TaxPeriod2{
								Yr: common.MustToISODate(common.ISODateSample),
								Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
								FrToDt: &pain_001_001_09.DatePeriod2{
									FrDt: common.MustToISODate(common.ISODateSample),
									ToDt: common.MustToISODate(common.ISODateSample),
								},
							},
							TaxAmt: &pain_001_001_09.TaxAmount2{
								Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
								TaxblBaseAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dtls: []pain_001_001_09.TaxRecordDetails2{{
									Prd: &pain_001_001_09.TaxPeriod2{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &pain_001_001_09.DatePeriod2{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
							},
							AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
						},
					},
					RltdRmtInf: []pain_001_001_09.RemittanceLocation7{{
						RmtId: common.MustToMax35Text(common.Max35TextSample),
						RmtLctnDtls: []pain_001_001_09.RemittanceLocationData1{{
							Mtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
							ElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
							PstlAdr: &pain_001_001_09.NameAndAddress16{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								Adr: &pain_001_001_09.PostalAddress24{
									AdrTp: &pain_001_001_09.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_001_001_09.GenericIdentification30{
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
							}},
						}},
					},
					RmtInf: &pain_001_001_09.RemittanceInformation16{
						Ustrd: []common.Max140Text{
							common.MustToMax140Text(common.Max140TextSample),
						},
						Strd: []pain_001_001_09.StructuredRemittanceInformation16{{
							RfrdDocInf: []pain_001_001_09.ReferredDocumentInformation7{{
								Tp: &pain_001_001_09.ReferredDocumentType4{
									CdOrPrtry: &pain_001_001_09.ReferredDocumentType3Choice{
										Cd:    common.MustToDocumentType6Code(common.DocumentType6CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Nb:     common.MustToMax35Text(common.Max35TextSample),
								RltdDt: common.MustToISODate(common.ISODateSample),
								LineDtls: []pain_001_001_09.DocumentLineInformation1{{
									Id: []pain_001_001_09.DocumentLineIdentification1{{
										Tp: &pain_001_001_09.DocumentLineType1{
											CdOrPrtry: &pain_001_001_09.DocumentLineType1Choice{
												Cd:    common.MustToExternalDocumentLineType1Code(common.ExternalDocumentLineType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
										Nb:     common.MustToMax35Text(common.Max35TextSample),
										RltdDt: common.MustToISODate(common.ISODateSample)},
									},
									Desc: common.MustToMax2048Text(common.Max2048TextSample),
									Amt: &pain_001_001_09.RemittanceAmount3{
										DuePyblAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										DscntApldAmt: []pain_001_001_09.DiscountAmountAndType1{{
											Tp: &pain_001_001_09.DiscountAmountType1Choice{
												Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										CdtNoteAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TaxAmt: []pain_001_001_09.TaxAmountAndType1{{
											Tp: &pain_001_001_09.TaxAmountType1Choice{
												Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										AdjstmntAmtAndRsn: []pain_001_001_09.DocumentAdjustment1{{
											Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
											Rsn:       common.MustToMax4Text(common.Max4TextSample),
											AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
										},
										RmtdAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
									}},
								}},
							},
							RfrdDocAmt: &pain_001_001_09.RemittanceAmount2{
								DuePyblAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								DscntApldAmt: []pain_001_001_09.DiscountAmountAndType1{{
									Tp: &pain_001_001_09.DiscountAmountType1Choice{
										Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								CdtNoteAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TaxAmt: []pain_001_001_09.TaxAmountAndType1{{
									Tp: &pain_001_001_09.TaxAmountType1Choice{
										Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								AdjstmntAmtAndRsn: []pain_001_001_09.DocumentAdjustment1{{
									Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
									Rsn:       common.MustToMax4Text(common.Max4TextSample),
									AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
								},
								RmtdAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
							},
							CdtrRefInf: &pain_001_001_09.CreditorReferenceInformation2{
								Tp: &pain_001_001_09.CreditorReferenceType2{
									CdOrPrtry: &pain_001_001_09.CreditorReferenceType1Choice{
										Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Ref: common.MustToMax35Text(common.Max35TextSample),
							},
							Invcr: &pain_001_001_09.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_001_001_09.PostalAddress24{
									AdrTp: &pain_001_001_09.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_001_001_09.GenericIdentification30{
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
								Id: &pain_001_001_09.Party38Choice{
									OrgId: &pain_001_001_09.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_001_001_09.PersonIdentification13{
										DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_001_001_09.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_001_001_09.Contact4{
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
									Othr: []pain_001_001_09.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Invcee: &pain_001_001_09.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_001_001_09.PostalAddress24{
									AdrTp: &pain_001_001_09.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_001_001_09.GenericIdentification30{
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
								Id: &pain_001_001_09.Party38Choice{
									OrgId: &pain_001_001_09.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_001_001_09.PersonIdentification13{
										DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_001_001_09.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_001_001_09.Contact4{
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
									Othr: []pain_001_001_09.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							TaxRmt: &pain_001_001_09.TaxInformation7{
								Cdtr: &pain_001_001_09.TaxParty1{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								},
								Dbtr: &pain_001_001_09.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &pain_001_001_09.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								UltmtDbtr: &pain_001_001_09.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &pain_001_001_09.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
								RefNb:      common.MustToMax140Text(common.Max140TextSample),
								Mtd:        common.MustToMax35Text(common.Max35TextSample),
								TtlTaxblBaseAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlTaxAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dt:    common.MustToISODate(common.ISODateSample),
								SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
								Rcrd: []pain_001_001_09.TaxRecord2{{
									Tp:       common.MustToMax35Text(common.Max35TextSample),
									Ctgy:     common.MustToMax35Text(common.Max35TextSample),
									CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
									DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
									CertId:   common.MustToMax35Text(common.Max35TextSample),
									FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
									Prd: &pain_001_001_09.TaxPeriod2{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &pain_001_001_09.DatePeriod2{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									TaxAmt: &pain_001_001_09.TaxAmount2{
										Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
										TaxblBaseAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TtlAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										Dtls: []pain_001_001_09.TaxRecordDetails2{{
											Prd: &pain_001_001_09.TaxPeriod2{
												Yr: common.MustToISODate(common.ISODateSample),
												Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
												FrToDt: &pain_001_001_09.DatePeriod2{
													FrDt: common.MustToISODate(common.ISODateSample),
													ToDt: common.MustToISODate(common.ISODateSample),
												},
											},
											Amt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
									},
									AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
								},
							},
							GrnshmtRmt: &pain_001_001_09.Garnishment3{
								Tp: &pain_001_001_09.GarnishmentType1{
									CdOrPrtry: &pain_001_001_09.GarnishmentType1Choice{
										Cd:    common.MustToExternalGarnishmentType1Code(common.ExternalGarnishmentType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Grnshee: &pain_001_001_09.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_001_001_09.PostalAddress24{
										AdrTp: &pain_001_001_09.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_001_001_09.GenericIdentification30{
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
									Id: &pain_001_001_09.Party38Choice{
										OrgId: &pain_001_001_09.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_001_001_09.PersonIdentification13{
											DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_001_001_09.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_001_001_09.Contact4{
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
										Othr: []pain_001_001_09.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								GrnshmtAdmstr: &pain_001_001_09.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_001_001_09.PostalAddress24{
										AdrTp: &pain_001_001_09.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_001_001_09.GenericIdentification30{
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
									Id: &pain_001_001_09.Party38Choice{
										OrgId: &pain_001_001_09.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_001_001_09.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_001_001_09.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_001_001_09.PersonIdentification13{
											DtAndPlcOfBirth: &pain_001_001_09.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_001_001_09.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_001_001_09.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_001_001_09.Contact4{
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
										Othr: []pain_001_001_09.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								RefNb: common.MustToMax140Text(common.Max140TextSample),
								Dt:    common.MustToISODate(common.ISODateSample),
								RmtdAmt: &pain_001_001_09.ActiveOrHistoricCurrencyAndAmount{
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
					SplmtryData: []pain_001_001_09.SupplementaryData1{{
						PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
						Envlp: &pain_001_001_09.SupplementaryDataEnvelope1{
							Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
						}},
					}},
				}},
			},
			SplmtryData: []pain_001_001_09.SupplementaryData1{{
				PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
				Envlp: &pain_001_001_09.SupplementaryDataEnvelope1{
					Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
				}},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_pain_001_001_09, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_pain_001_001_09)

}
