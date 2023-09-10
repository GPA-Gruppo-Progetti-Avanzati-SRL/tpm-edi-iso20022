// Package pain_013_001_07_test
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/013.001.07/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/013.001.07/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_pain_013_001_07 = "example-document-pain_013_001_07.xml"

func TestDocumentpain_013_001_07(t *testing.T) {

	d := pain_013_001_07.Document{
		CdtrPmtActvtnReq: &pain_013_001_07.CreditorPaymentActivationRequestV07{
			GrpHdr: &pain_013_001_07.GroupHeader78{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				NbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
				CtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample),
				InitgPty: &pain_013_001_07.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_013_001_07.PostalAddress24{
						AdrTp: &pain_013_001_07.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_013_001_07.GenericIdentification30{
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
					Id: &pain_013_001_07.Party38Choice{
						OrgId: &pain_013_001_07.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_013_001_07.PersonIdentification13{
							DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_013_001_07.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_013_001_07.Contact4{
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
						Othr: []pain_013_001_07.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
			},
			PmtInf: []pain_013_001_07.PaymentInstruction31{{
				PmtInfId: common.MustToMax35Text(common.Max35TextSample),
				PmtMtd:   common.MustToPaymentMethod7Code(common.PaymentMethod7CodeSample),
				PmtTpInf: &pain_013_001_07.PaymentTypeInformation26{
					InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
					SvcLvl: []pain_013_001_07.ServiceLevel8Choice{{
						Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample)},
					},
					LclInstrm: &pain_013_001_07.LocalInstrument2Choice{
						Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					CtgyPurp: &pain_013_001_07.CategoryPurpose1Choice{
						Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				ReqdExctnDt: &pain_013_001_07.DateAndDateTime2Choice{
					Dt:   common.MustToISODate(common.ISODateSample),
					DtTm: common.MustToISODateTime(common.ISODateTimeSample),
				},
				XpryDt: &pain_013_001_07.DateAndDateTime2Choice{
					Dt:   common.MustToISODate(common.ISODateSample),
					DtTm: common.MustToISODateTime(common.ISODateTimeSample),
				},
				PmtCond: &pain_013_001_07.PaymentCondition1{
					AmtModAllwd:   xsdt.MustToBoolean(xsdt.BooleanSample),
					EarlyPmtAllwd: xsdt.MustToBoolean(xsdt.BooleanSample),
					DelyPnlty:     common.MustToMax140Text(common.Max140TextSample),
					ImdtPmtRbt: &pain_013_001_07.AmountOrRate1Choice{
						Amt: &pain_013_001_07.ActiveCurrencyAndAmount{
							Ccy:   common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					GrntedPmtReqd: xsdt.MustToBoolean(xsdt.BooleanSample),
				},
				Dbtr: &pain_013_001_07.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_013_001_07.PostalAddress24{
						AdrTp: &pain_013_001_07.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_013_001_07.GenericIdentification30{
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
					Id: &pain_013_001_07.Party38Choice{
						OrgId: &pain_013_001_07.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_013_001_07.PersonIdentification13{
							DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_013_001_07.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_013_001_07.Contact4{
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
						Othr: []pain_013_001_07.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				DbtrAcct: &pain_013_001_07.CashAccount38{
					Id: &pain_013_001_07.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
						Othr: &pain_013_001_07.GenericAccountIdentification1{
							Id: common.MustToMax34Text(common.Max34TextSample),
							SchmeNm: &pain_013_001_07.AccountSchemeName1Choice{
								Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					Tp: &pain_013_001_07.CashAccountType2Choice{
						Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
					Nm:  common.MustToMax70Text(common.Max70TextSample),
					Prxy: &pain_013_001_07.ProxyAccountIdentification1{
						Tp: &pain_013_001_07.ProxyAccountType1Choice{
							Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Id: common.MustToMax2048Text(common.Max2048TextSample),
					},
				},
				DbtrAgt: &pain_013_001_07.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: &pain_013_001_07.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &pain_013_001_07.ClearingSystemMemberIdentification2{
							ClrSysId: &pain_013_001_07.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_013_001_07.PostalAddress24{
							AdrTp: &pain_013_001_07.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_013_001_07.GenericIdentification30{
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
						Othr: &pain_013_001_07.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &pain_013_001_07.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &pain_013_001_07.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_013_001_07.PostalAddress24{
							AdrTp: &pain_013_001_07.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_013_001_07.GenericIdentification30{
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
				UltmtDbtr: &pain_013_001_07.PartyIdentification135{
					Nm: common.MustToMax140Text(common.Max140TextSample),
					PstlAdr: &pain_013_001_07.PostalAddress24{
						AdrTp: &pain_013_001_07.AddressType3Choice{
							Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
							Prtry: &pain_013_001_07.GenericIdentification30{
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
					Id: &pain_013_001_07.Party38Choice{
						OrgId: &pain_013_001_07.OrganisationIdentification29{
							AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
							LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
						PrvtId: &pain_013_001_07.PersonIdentification13{
							DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
								BirthDt:     common.MustToISODate(common.ISODateSample),
								PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
								CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
							},
							Othr: []pain_013_001_07.GenericPersonIdentification1{{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample)},
							},
						},
					},
					CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
					CtctDtls: &pain_013_001_07.Contact4{
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
						Othr: []pain_013_001_07.OtherContact1{{
							ChanlTp: common.MustToMax4Text(common.Max4TextSample),
							Id:      common.MustToMax128Text(common.Max128TextSample)},
						},
						PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
					},
				},
				ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
				CdtTrfTx: []pain_013_001_07.CreditTransferTransaction35{{
					PmtId: &pain_013_001_07.PaymentIdentification6{
						InstrId:    common.MustToMax35Text(common.Max35TextSample),
						EndToEndId: common.MustToMax35Text(common.Max35TextSample),
						UETR:       common.MustToUUIDv4Identifier(common.UUIDv4IdentifierSample),
					},
					PmtTpInf: &pain_013_001_07.PaymentTypeInformation26{
						InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
						SvcLvl: []pain_013_001_07.ServiceLevel8Choice{{
							Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample)},
						},
						LclInstrm: &pain_013_001_07.LocalInstrument2Choice{
							Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						CtgyPurp: &pain_013_001_07.CategoryPurpose1Choice{
							Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					PmtCond: &pain_013_001_07.PaymentCondition1{
						AmtModAllwd:   xsdt.MustToBoolean(xsdt.BooleanSample),
						EarlyPmtAllwd: xsdt.MustToBoolean(xsdt.BooleanSample),
						DelyPnlty:     common.MustToMax140Text(common.Max140TextSample),
						ImdtPmtRbt: &pain_013_001_07.AmountOrRate1Choice{
							Amt: &pain_013_001_07.ActiveCurrencyAndAmount{
								Ccy:   common.MustToActiveCurrencyCode(common.ActiveCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						GrntedPmtReqd: xsdt.MustToBoolean(xsdt.BooleanSample),
					},
					Amt: &pain_013_001_07.AmountType4Choice{
						InstdAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						EqvtAmt: &pain_013_001_07.EquivalentAmount2{
							Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						},
					},
					ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSample),
					ChqInstr: &pain_013_001_07.Cheque11{
						ChqTp: common.MustToChequeType2Code(common.ChequeType2CodeSample),
						ChqNb: common.MustToMax35Text(common.Max35TextSample),
						ChqFr: &pain_013_001_07.NameAndAddress16{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							Adr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
						DlvryMtd: &pain_013_001_07.ChequeDeliveryMethod1Choice{
							Cd:    common.MustToChequeDelivery1Code(common.ChequeDelivery1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						DlvrTo: &pain_013_001_07.NameAndAddress16{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							Adr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
					UltmtDbtr: &pain_013_001_07.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_013_001_07.PostalAddress24{
							AdrTp: &pain_013_001_07.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_013_001_07.GenericIdentification30{
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
						Id: &pain_013_001_07.Party38Choice{
							OrgId: &pain_013_001_07.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_013_001_07.PersonIdentification13{
								DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_013_001_07.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_013_001_07.Contact4{
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
							Othr: []pain_013_001_07.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					IntrmyAgt1: &pain_013_001_07.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_013_001_07.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_013_001_07.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_013_001_07.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
							Othr: &pain_013_001_07.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_013_001_07.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
					IntrmyAgt2: &pain_013_001_07.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_013_001_07.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_013_001_07.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_013_001_07.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
							Othr: &pain_013_001_07.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_013_001_07.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
					IntrmyAgt3: &pain_013_001_07.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_013_001_07.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_013_001_07.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_013_001_07.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
							Othr: &pain_013_001_07.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_013_001_07.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
					CdtrAgt: &pain_013_001_07.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: &pain_013_001_07.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &pain_013_001_07.ClearingSystemMemberIdentification2{
								ClrSysId: &pain_013_001_07.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
							Othr: &pain_013_001_07.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &pain_013_001_07.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &pain_013_001_07.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &pain_013_001_07.PostalAddress24{
								AdrTp: &pain_013_001_07.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &pain_013_001_07.GenericIdentification30{
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
					Cdtr: &pain_013_001_07.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_013_001_07.PostalAddress24{
							AdrTp: &pain_013_001_07.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_013_001_07.GenericIdentification30{
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
						Id: &pain_013_001_07.Party38Choice{
							OrgId: &pain_013_001_07.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_013_001_07.PersonIdentification13{
								DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_013_001_07.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_013_001_07.Contact4{
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
							Othr: []pain_013_001_07.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					CdtrAcct: &pain_013_001_07.CashAccount38{
						Id: &pain_013_001_07.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &pain_013_001_07.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &pain_013_001_07.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &pain_013_001_07.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &pain_013_001_07.ProxyAccountIdentification1{
							Tp: &pain_013_001_07.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					UltmtCdtr: &pain_013_001_07.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &pain_013_001_07.PostalAddress24{
							AdrTp: &pain_013_001_07.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &pain_013_001_07.GenericIdentification30{
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
						Id: &pain_013_001_07.Party38Choice{
							OrgId: &pain_013_001_07.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &pain_013_001_07.PersonIdentification13{
								DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []pain_013_001_07.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &pain_013_001_07.Contact4{
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
							Othr: []pain_013_001_07.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					InstrForCdtrAgt: []pain_013_001_07.InstructionForCreditorAgent1{{
						Cd:       common.MustToInstruction3Code(common.Instruction3CodeSample),
						InstrInf: common.MustToMax140Text(common.Max140TextSample)},
					},
					Purp: &pain_013_001_07.Purpose2Choice{
						Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
					RgltryRptg: []pain_013_001_07.RegulatoryReporting3{{
						DbtCdtRptgInd: common.MustToRegulatoryReportingType1Code(common.RegulatoryReportingType1CodeSample),
						Authrty: &pain_013_001_07.RegulatoryAuthority2{
							Nm:   common.MustToMax140Text(common.Max140TextSample),
							Ctry: common.MustToCountryCode(common.CountryCodeSample),
						},
						Dtls: []pain_013_001_07.StructuredRegulatoryReporting3{{
							Tp:   common.MustToMax35Text(common.Max35TextSample),
							Dt:   common.MustToISODate(common.ISODateSample),
							Ctry: common.MustToCountryCode(common.CountryCodeSample),
							Cd:   common.MustToMax10Text(common.Max10TextSample),
							Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							Inf: []common.Max35Text{
								common.MustToMax35Text(common.Max35TextSample),
							}},
						}},
					},
					Tax: &pain_013_001_07.TaxInformation8{
						Cdtr: &pain_013_001_07.TaxParty1{
							TaxId:  common.MustToMax35Text(common.Max35TextSample),
							RegnId: common.MustToMax35Text(common.Max35TextSample),
							TaxTp:  common.MustToMax35Text(common.Max35TextSample),
						},
						Dbtr: &pain_013_001_07.TaxParty2{
							TaxId:  common.MustToMax35Text(common.Max35TextSample),
							RegnId: common.MustToMax35Text(common.Max35TextSample),
							TaxTp:  common.MustToMax35Text(common.Max35TextSample),
							Authstn: &pain_013_001_07.TaxAuthorisation1{
								Titl: common.MustToMax35Text(common.Max35TextSample),
								Nm:   common.MustToMax140Text(common.Max140TextSample),
							},
						},
						AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
						RefNb:      common.MustToMax140Text(common.Max140TextSample),
						Mtd:        common.MustToMax35Text(common.Max35TextSample),
						TtlTaxblBaseAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						TtlTaxAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						Dt:    common.MustToISODate(common.ISODateSample),
						SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
						Rcrd: []pain_013_001_07.TaxRecord2{{
							Tp:       common.MustToMax35Text(common.Max35TextSample),
							Ctgy:     common.MustToMax35Text(common.Max35TextSample),
							CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
							DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
							CertId:   common.MustToMax35Text(common.Max35TextSample),
							FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
							Prd: &pain_013_001_07.TaxPeriod2{
								Yr: common.MustToISODate(common.ISODateSample),
								Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
								FrToDt: &pain_013_001_07.DatePeriod2{
									FrDt: common.MustToISODate(common.ISODateSample),
									ToDt: common.MustToISODate(common.ISODateSample),
								},
							},
							TaxAmt: &pain_013_001_07.TaxAmount2{
								Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
								TaxblBaseAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dtls: []pain_013_001_07.TaxRecordDetails2{{
									Prd: &pain_013_001_07.TaxPeriod2{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &pain_013_001_07.DatePeriod2{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
							},
							AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
						},
					},
					RltdRmtInf: []pain_013_001_07.RemittanceLocation7{{
						RmtId: common.MustToMax35Text(common.Max35TextSample),
						RmtLctnDtls: []pain_013_001_07.RemittanceLocationData1{{
							Mtd:        common.MustToRemittanceLocationMethod2Code(common.RemittanceLocationMethod2CodeSample),
							ElctrncAdr: common.MustToMax2048Text(common.Max2048TextSample),
							PstlAdr: &pain_013_001_07.NameAndAddress16{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								Adr: &pain_013_001_07.PostalAddress24{
									AdrTp: &pain_013_001_07.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_013_001_07.GenericIdentification30{
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
					RmtInf: &pain_013_001_07.RemittanceInformation16{
						Ustrd: []common.Max140Text{
							common.MustToMax140Text(common.Max140TextSample),
						},
						Strd: []pain_013_001_07.StructuredRemittanceInformation16{{
							RfrdDocInf: []pain_013_001_07.ReferredDocumentInformation7{{
								Tp: &pain_013_001_07.ReferredDocumentType4{
									CdOrPrtry: &pain_013_001_07.ReferredDocumentType3Choice{
										Cd:    common.MustToDocumentType6Code(common.DocumentType6CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Nb:     common.MustToMax35Text(common.Max35TextSample),
								RltdDt: common.MustToISODate(common.ISODateSample),
								LineDtls: []pain_013_001_07.DocumentLineInformation1{{
									Id: []pain_013_001_07.DocumentLineIdentification1{{
										Tp: &pain_013_001_07.DocumentLineType1{
											CdOrPrtry: &pain_013_001_07.DocumentLineType1Choice{
												Cd:    common.MustToExternalDocumentLineType1Code(common.ExternalDocumentLineType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
										Nb:     common.MustToMax35Text(common.Max35TextSample),
										RltdDt: common.MustToISODate(common.ISODateSample)},
									},
									Desc: common.MustToMax2048Text(common.Max2048TextSample),
									Amt: &pain_013_001_07.RemittanceAmount3{
										DuePyblAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										DscntApldAmt: []pain_013_001_07.DiscountAmountAndType1{{
											Tp: &pain_013_001_07.DiscountAmountType1Choice{
												Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										CdtNoteAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TaxAmt: []pain_013_001_07.TaxAmountAndType1{{
											Tp: &pain_013_001_07.TaxAmountType1Choice{
												Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										AdjstmntAmtAndRsn: []pain_013_001_07.DocumentAdjustment1{{
											Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
											Rsn:       common.MustToMax4Text(common.Max4TextSample),
											AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
										},
										RmtdAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
									}},
								}},
							},
							RfrdDocAmt: &pain_013_001_07.RemittanceAmount2{
								DuePyblAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								DscntApldAmt: []pain_013_001_07.DiscountAmountAndType1{{
									Tp: &pain_013_001_07.DiscountAmountType1Choice{
										Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								CdtNoteAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TaxAmt: []pain_013_001_07.TaxAmountAndType1{{
									Tp: &pain_013_001_07.TaxAmountType1Choice{
										Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								AdjstmntAmtAndRsn: []pain_013_001_07.DocumentAdjustment1{{
									Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
									Rsn:       common.MustToMax4Text(common.Max4TextSample),
									AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
								},
								RmtdAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
							},
							CdtrRefInf: &pain_013_001_07.CreditorReferenceInformation2{
								Tp: &pain_013_001_07.CreditorReferenceType2{
									CdOrPrtry: &pain_013_001_07.CreditorReferenceType1Choice{
										Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Ref: common.MustToMax35Text(common.Max35TextSample),
							},
							Invcr: &pain_013_001_07.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_013_001_07.PostalAddress24{
									AdrTp: &pain_013_001_07.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_013_001_07.GenericIdentification30{
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
								Id: &pain_013_001_07.Party38Choice{
									OrgId: &pain_013_001_07.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_013_001_07.PersonIdentification13{
										DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_013_001_07.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_013_001_07.Contact4{
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
									Othr: []pain_013_001_07.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Invcee: &pain_013_001_07.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_013_001_07.PostalAddress24{
									AdrTp: &pain_013_001_07.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_013_001_07.GenericIdentification30{
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
								Id: &pain_013_001_07.Party38Choice{
									OrgId: &pain_013_001_07.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_013_001_07.PersonIdentification13{
										DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_013_001_07.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_013_001_07.Contact4{
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
									Othr: []pain_013_001_07.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							TaxRmt: &pain_013_001_07.TaxInformation7{
								Cdtr: &pain_013_001_07.TaxParty1{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								},
								Dbtr: &pain_013_001_07.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &pain_013_001_07.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								UltmtDbtr: &pain_013_001_07.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &pain_013_001_07.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
								RefNb:      common.MustToMax140Text(common.Max140TextSample),
								Mtd:        common.MustToMax35Text(common.Max35TextSample),
								TtlTaxblBaseAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlTaxAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dt:    common.MustToISODate(common.ISODateSample),
								SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
								Rcrd: []pain_013_001_07.TaxRecord2{{
									Tp:       common.MustToMax35Text(common.Max35TextSample),
									Ctgy:     common.MustToMax35Text(common.Max35TextSample),
									CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
									DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
									CertId:   common.MustToMax35Text(common.Max35TextSample),
									FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
									Prd: &pain_013_001_07.TaxPeriod2{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &pain_013_001_07.DatePeriod2{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									TaxAmt: &pain_013_001_07.TaxAmount2{
										Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
										TaxblBaseAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TtlAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										Dtls: []pain_013_001_07.TaxRecordDetails2{{
											Prd: &pain_013_001_07.TaxPeriod2{
												Yr: common.MustToISODate(common.ISODateSample),
												Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
												FrToDt: &pain_013_001_07.DatePeriod2{
													FrDt: common.MustToISODate(common.ISODateSample),
													ToDt: common.MustToISODate(common.ISODateSample),
												},
											},
											Amt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
									},
									AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
								},
							},
							GrnshmtRmt: &pain_013_001_07.Garnishment3{
								Tp: &pain_013_001_07.GarnishmentType1{
									CdOrPrtry: &pain_013_001_07.GarnishmentType1Choice{
										Cd:    common.MustToExternalGarnishmentType1Code(common.ExternalGarnishmentType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Grnshee: &pain_013_001_07.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_013_001_07.PostalAddress24{
										AdrTp: &pain_013_001_07.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_013_001_07.GenericIdentification30{
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
									Id: &pain_013_001_07.Party38Choice{
										OrgId: &pain_013_001_07.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_013_001_07.PersonIdentification13{
											DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_013_001_07.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_013_001_07.Contact4{
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
										Othr: []pain_013_001_07.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								GrnshmtAdmstr: &pain_013_001_07.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &pain_013_001_07.PostalAddress24{
										AdrTp: &pain_013_001_07.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &pain_013_001_07.GenericIdentification30{
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
									Id: &pain_013_001_07.Party38Choice{
										OrgId: &pain_013_001_07.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &pain_013_001_07.PersonIdentification13{
											DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []pain_013_001_07.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &pain_013_001_07.Contact4{
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
										Othr: []pain_013_001_07.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								RefNb: common.MustToMax140Text(common.Max140TextSample),
								Dt:    common.MustToISODate(common.ISODateSample),
								RmtdAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
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
					NclsdFile: []pain_013_001_07.Document12{{
						Tp: &pain_013_001_07.DocumentType1Choice{
							Cd: common.MustToExternalDocumentType1Code(common.ExternalDocumentType1CodeSample),
							Prtry: &pain_013_001_07.GenericIdentification1{
								Id:      common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Id: common.MustToMax35Text(common.Max35TextSample),
						IsseDt: &pain_013_001_07.DateAndDateTime2Choice{
							Dt:   common.MustToISODate(common.ISODateSample),
							DtTm: common.MustToISODateTime(common.ISODateTimeSample),
						},
						Nm:     common.MustToMax140Text(common.Max140TextSample),
						LangCd: xsdt.MustToString(xsdt.StringSample),
						Frmt: &pain_013_001_07.DocumentFormat1Choice{
							Cd: common.MustToExternalDocumentFormat1Code(common.ExternalDocumentFormat1CodeSample),
							Prtry: &pain_013_001_07.GenericIdentification1{
								Id:      common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								Issr:    common.MustToMax35Text(common.Max35TextSample),
							},
						},
						FileNm: common.MustToMax140Text(common.Max140TextSample),
						DgtlSgntr: &pain_013_001_07.PartyAndSignature3{
							Pty: &pain_013_001_07.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &pain_013_001_07.PostalAddress24{
									AdrTp: &pain_013_001_07.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &pain_013_001_07.GenericIdentification30{
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
								Id: &pain_013_001_07.Party38Choice{
									OrgId: &pain_013_001_07.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []pain_013_001_07.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &pain_013_001_07.PersonIdentification13{
										DtAndPlcOfBirth: &pain_013_001_07.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []pain_013_001_07.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &pain_013_001_07.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &pain_013_001_07.Contact4{
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
									Othr: []pain_013_001_07.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Sgntr: &pain_013_001_07.SkipPayload{
								Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
							},
						},
						Nclsr: common.MustToMax10MbBinary(common.Max10MbBinarySample)},
					},
					SplmtryData: []pain_013_001_07.SupplementaryData1{{
						PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
						Envlp: &pain_013_001_07.SupplementaryDataEnvelope1{
							Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
						}},
					}},
				}},
			},
			SplmtryData: []pain_013_001_07.SupplementaryData1{{
				PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
				Envlp: &pain_013_001_07.SupplementaryDataEnvelope1{
					Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
				}},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_pain_013_001_07, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_pain_013_001_07)

}
