// Package pain_001_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_001_001_03

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_CstmrCdtTrfInitn                                                                             = "CstmrCdtTrfInitn"
	Path_CstmrCdtTrfInitn_GrpHdr                                                                      = "CstmrCdtTrfInitn.GrpHdr"
	Path_CstmrCdtTrfInitn_GrpHdr_MsgId                                                                = "CstmrCdtTrfInitn.GrpHdr.MsgId"
	Path_CstmrCdtTrfInitn_GrpHdr_CreDtTm                                                              = "CstmrCdtTrfInitn.GrpHdr.CreDtTm"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Item                                                         = "CstmrCdtTrfInitn.GrpHdr.Authstn[]"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn                                                              = "CstmrCdtTrfInitn.GrpHdr.Authstn" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Cd                                                           = "CstmrCdtTrfInitn.GrpHdr.Authstn[].Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Prtry                                                        = "CstmrCdtTrfInitn.GrpHdr.Authstn[].Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_NbOfTxs                                                              = "CstmrCdtTrfInitn.GrpHdr.NbOfTxs"
	Path_CstmrCdtTrfInitn_GrpHdr_CtrlSum                                                              = "CstmrCdtTrfInitn.GrpHdr.CtrlSum"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty                                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Nm                                                          = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr                                                     = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Dept                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_SubDept                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_StrtNm                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNb                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstCd                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnNm                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Ctry                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item                                        = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id                                                          = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId                                                    = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_BICOrBEI                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item                                          = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm                                       = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd                                    = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr                                          = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                       = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                       = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                       = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm                                      = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtryOfRes                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtryOfRes"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls                                                    = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Nm                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PhneNb                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_MobNb                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_FaxNb                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt                                                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_BIC                                               = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Nm                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn                               = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr                                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId                                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Id                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Nm                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr                                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn                                  = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Item                                                                 = "CstmrCdtTrfInitn.PmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf                                                                      = "CstmrCdtTrfInitn.PmtInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_PmtInfId                                                             = "CstmrCdtTrfInitn.PmtInf[].PmtInfId"
	Path_CstmrCdtTrfInitn_PmtInf_PmtMtd                                                               = "CstmrCdtTrfInitn.PmtInf[].PmtMtd"
	Path_CstmrCdtTrfInitn_PmtInf_BtchBookg                                                            = "CstmrCdtTrfInitn.PmtInf[].BtchBookg"
	Path_CstmrCdtTrfInitn_PmtInf_NbOfTxs                                                              = "CstmrCdtTrfInitn.PmtInf[].NbOfTxs"
	Path_CstmrCdtTrfInitn_PmtInf_CtrlSum                                                              = "CstmrCdtTrfInitn.PmtInf[].CtrlSum"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf                                                             = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_InstrPrty                                                   = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl                                                      = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Cd                                                   = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Prtry                                                = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm                                                   = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Cd                                                = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Prtry                                             = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp                                                    = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Cd                                                 = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry                                              = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt                                                          = "CstmrCdtTrfInitn.PmtInf[].ReqdExctnDt"
	Path_CstmrCdtTrfInitn_PmtInf_PoolgAdjstmntDt                                                      = "CstmrCdtTrfInitn.PmtInf[].PoolgAdjstmntDt"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr                                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Nm                                                              = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr                                                         = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Dept                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_SubDept                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_StrtNm                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNb                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstCd                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnNm                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_CtrySubDvsn                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Ctry                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine_Item                                            = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id                                                              = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId                                                        = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_BICOrBEI                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Item                                              = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Id                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm                                           = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Issr                                              = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                           = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                           = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                           = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Item                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Id                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm                                          = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Issr                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtryOfRes                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls                                                        = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_NmPrfx                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Nm                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PhneNb                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_MobNb                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_FaxNb                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailAdr                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct                                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_IBAN                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Id                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Issr                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Ccy                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Nm                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt                                                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_BIC                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Nm                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr                                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Id                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Issr                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Id                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Nm                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr                                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_IBAN                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Id                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Issr                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Cd                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Prtry                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Ccy                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Nm                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr                                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Nm                                                         = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr                                                    = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Dept                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_SubDept                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_StrtNm                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNb                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstCd                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnNm                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Ctry                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item                                       = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id                                                         = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId                                                   = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_BICOrBEI                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item                                         = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                      = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                   = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr                                         = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                      = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                      = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                      = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                     = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtryOfRes                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls                                                   = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_NmPrfx                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Nm                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PhneNb                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_MobNb                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_FaxNb                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailAdr                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgBr                                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgBr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct                                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id                                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_IBAN                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Issr                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp                                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Cd                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Prtry                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Ccy                                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Nm                                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt                                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BIC                                          = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn                          = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn                             = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Item                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_InstrId                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId.InstrId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_EndToEndId                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId.EndToEndId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_InstrPrty                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Cd                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Prtry                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Cd                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Ccy                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Value                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Ccy                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Value                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_CcyOfTrf                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.CcyOfTrf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_XchgRate                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.XchgRate"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_RateTp                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.RateTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_CtrctId                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.CtrctId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChrgBr                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChrgBr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqNb                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Nm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Dept                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_SubDept                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_StrtNm                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNb                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstCd                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnNm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_CtrySubDvsn                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Ctry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine_Item                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Nm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Dept                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_SubDept                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_StrtNm                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNb                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstCd                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnNm                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_CtrySubDvsn                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Ctry                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine_Item                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_InstrPrty                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqMtrtyDt                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqMtrtyDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_FrmsCd                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.FrmsCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.MemoFld[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.MemoFld" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_RgnlClrZone                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.RgnlClrZone"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_PrtLctn                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.PrtLctn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Nm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Dept                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_SubDept                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_StrtNm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstCd                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnNm                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Ctry                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine_Item                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_BICOrBEI                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Item                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Item                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Id                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtryOfRes                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_NmPrfx                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PhneNb                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_MobNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_FaxNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailAdr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_BIC                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Nm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Dept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_SubDept                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstCd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Ctry                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Id                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Issr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_IBAN                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Id                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Cd                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Prtry                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Issr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Ccy                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Nm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_BIC                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Nm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Dept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_SubDept                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstCd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Ctry                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Id                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Issr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_IBAN                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Id                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Cd                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Prtry                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Issr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Ccy                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Nm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_BIC                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Nm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Dept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_SubDept                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstCd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Ctry                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Id                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Issr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_IBAN                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Id                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Cd                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Prtry                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Issr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Ccy                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Nm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_BIC                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.BIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Id                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Nm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Dept                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_SubDept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_StrtNm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNb                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstCd                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnNm                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Ctry                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_IBAN                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Id                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Issr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Ccy                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Nm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Nm                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Dept                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_SubDept                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_StrtNm                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNb                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstCd                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_CtrySubDvsn                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Ctry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine_Item                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_BICOrBEI                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Item                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Issr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Item                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Id                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Issr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtryOfRes                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_NmPrfx                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Nm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PhneNb                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_MobNb                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_FaxNb                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailAdr                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_IBAN                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Id                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Cd                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Prtry                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Issr                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Cd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Prtry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Ccy                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Nm                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Nm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Dept                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_SubDept                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_StrtNm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstCd                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnNm                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Ctry                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine_Item                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_BICOrBEI                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Item                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Item                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Id                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtryOfRes                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_NmPrfx                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PhneNb                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_MobNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_FaxNb                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailAdr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Item                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Cd                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_InstrInf                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[].InstrInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForDbtrAgt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForDbtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Cd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Prtry                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Item                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_DbtCdtRptgInd                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].DbtCdtRptgInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Ctry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Item                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Tp                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Dt                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Ctry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Cd                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Ccy                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Value                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf_Item                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Inf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Inf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxId                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_RegnId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxId                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_RegnId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Titl                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn.Titl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Nm                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_AdmstnZn                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.AdmstnZn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_RefNb                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.RefNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Mtd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Mtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Ccy                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Value                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Ccy                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Value                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dt                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_SeqNb                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.SeqNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Item                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Tp                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Ctgy                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Ctgy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CtgyDtls                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].CtgyDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_DbtrSts                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].DbtrSts"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CertId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].CertId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_FrmsCd                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].FrmsCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Yr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Tp                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_FrDt                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_ToDt                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Rate                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Rate"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Item                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_AddtlInf                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_Item                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtId                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnMtd                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnElctrncAdr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnElctrncAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Nm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd_Item                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Ustrd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Ustrd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Item                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Item                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Nb                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_RltdDt                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Ref                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Ref"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Nm                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtryOfRes                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.BICOrBEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtryOfRes                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf_Item                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].AddtlRmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].AddtlRmtInf" // ARRAY
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_CstmrCdtTrfInitn:                                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_MsgId:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_CreDtTm:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Item:                                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn:                                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Cd:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Prtry:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_NbOfTxs:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_CtrlSum:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Dept:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_SubDept:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_StrtNm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNb:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstCd:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnNm:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Ctry:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item:                                        PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine:                                             PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_BICOrBEI:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item:                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr:                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item:                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr:                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtryOfRes:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Nm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PhneNb:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_MobNb:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_FaxNb:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_BIC:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Nm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item:                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine:                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Id:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Nm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Item:                                                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf:                                                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_PmtInfId:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtMtd:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_BtchBookg:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_NbOfTxs:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CtrlSum:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_InstrPrty:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Cd:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Prtry:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Cd:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Prtry:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Cd:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PoolgAdjstmntDt:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr:                                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Nm:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Dept:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_SubDept:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_StrtNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstCd:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_CtrySubDvsn:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Ctry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine_Item:                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine:                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_BICOrBEI:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Item:                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr:                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Id:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Issr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Item:                                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr:                                                  PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Id:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Issr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtryOfRes:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_NmPrfx:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Nm:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PhneNb:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_MobNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_FaxNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailAdr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_IBAN:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Id:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Issr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Ccy:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_BIC:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Nm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Id:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Nm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_IBAN:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Id:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Issr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Cd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Prtry:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Ccy:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Nm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Nm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Dept:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_SubDept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_StrtNm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNb:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstCd:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnNm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Ctry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item:                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine:                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_BICOrBEI:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item:                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr:                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item:                                        PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr:                                             PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtryOfRes:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_NmPrfx:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Nm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PhneNb:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_MobNb:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_FaxNb:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailAdr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgBr:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_IBAN:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Cd:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Prtry:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Ccy:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Nm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BIC:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId:                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item:                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine:                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item:                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine:                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Item:                                                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf:                                                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_InstrId:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_EndToEndId:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_InstrPrty:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Cd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Prtry:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Cd:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Ccy:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Value:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Ccy:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Value:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_CcyOfTrf:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_XchgRate:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_RateTp:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_CtrctId:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChrgBr:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqTp:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqNb:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Nm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Dept:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_SubDept:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_StrtNm:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNb:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstCd:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnNm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_CtrySubDvsn:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Ctry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine_Item:                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine:                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Nm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Dept:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_SubDept:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_StrtNm:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNb:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstCd:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnNm:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_CtrySubDvsn:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Ctry:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine_Item:                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine:                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_InstrPrty:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqMtrtyDt:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_FrmsCd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_RgnlClrZone:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_PrtLctn:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Nm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Dept:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_SubDept:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_StrtNm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstCd:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnNm:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Ctry:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine_Item:                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine:                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_BICOrBEI:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Item:                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr:                                  PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Item:                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr:                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Id:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtryOfRes:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_NmPrfx:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PhneNb:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_MobNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_FaxNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailAdr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_BIC:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId:               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Nm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Dept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_SubDept:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstCd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Ctry:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item:               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine:                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Id:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Id:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_IBAN:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Id:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Cd:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Prtry:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Issr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Ccy:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Nm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_BIC:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId:               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Nm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Dept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_SubDept:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstCd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Ctry:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item:               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine:                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Id:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Id:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_IBAN:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Id:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Cd:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Prtry:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Issr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Ccy:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Nm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_BIC:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId:               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Nm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Dept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_SubDept:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstCd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Ctry:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item:               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine:                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Id:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Id:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_IBAN:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Id:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Cd:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Prtry:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Issr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Ccy:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Nm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_BIC:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm:                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Id:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Nm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Dept:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_SubDept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_StrtNm:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNb:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstCd:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnNm:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Ctry:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine:                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_IBAN:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Id:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Issr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Ccy:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Nm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Nm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Dept:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_SubDept:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_StrtNm:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNb:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstCd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_CtrySubDvsn:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Ctry:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine_Item:                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine:                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_BICOrBEI:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Item:                                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr:                                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Id:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Issr:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Id:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Issr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtryOfRes:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_NmPrfx:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Nm:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PhneNb:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_MobNb:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_FaxNb:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailAdr:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_IBAN:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Id:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Cd:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Prtry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Issr:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Cd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Prtry:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Ccy:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Nm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Nm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Dept:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_SubDept:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_StrtNm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstCd:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnNm:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Ctry:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine_Item:                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine:                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_BICOrBEI:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Item:                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr:                                  PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Item:                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr:                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Id:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtryOfRes:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_NmPrfx:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PhneNb:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_MobNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_FaxNb:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailAdr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Item:                                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt:                                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Cd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_InstrInf:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForDbtrAgt:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Cd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Prtry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Item:                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg:                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_DbtCdtRptgInd:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Ctry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Item:                                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls:                                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Tp:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Dt:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Ctry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Cd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Ccy:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Value:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxId:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_RegnId:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxTp:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxId:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_RegnId:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxTp:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Titl:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Nm:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_AdmstnZn:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_RefNb:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Mtd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Ccy:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Value:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Ccy:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Value:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dt:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_SeqNb:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Item:                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd:                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Tp:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Ctgy:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CtgyDtls:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_DbtrSts:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CertId:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_FrmsCd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Yr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Tp:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_FrDt:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_ToDt:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Rate:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Item:                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls:                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_AddtlInf:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_Item:                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf:                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtId:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnMtd:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnElctrncAdr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Nm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item:               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine:                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd_Item:                                        PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd:                                             PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Item:                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd:                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Item:                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf:                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Nb:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_RltdDt:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Ref:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Nm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr:                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtryOfRes:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr:                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtryOfRes:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf_Item:                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf:                                  PathTypeArray,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in pain_001_001_03", p)
	}

	return t, nil
}

func MustSetArrayItemPathModifiers(p string, modifiers []string) string {
	var err error
	p, err = SetArrayItemPathModifiers(p, modifiers)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return p
}

func SetArrayItemPathModifiers(p string, modifiers []string) (string, error) {
	numArrSpecifiers := strings.Count(p, "[]")
	if len(modifiers) != numArrSpecifiers {
		err := fmt.Errorf("the number of provided (%d) modifiers doesn't match the path provided", len(modifiers))
		return p, err
	}

	for _, m := range modifiers {
		if m == "" {
			m = "last"
		}
		p = strings.Replace(p, "[]", fmt.Sprintf("[%s]", m), 1)
	}

	return p, nil
}
