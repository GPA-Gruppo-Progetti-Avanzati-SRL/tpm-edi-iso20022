// Package pain_008_001_02
// Do not Edit. This stuff it's been automatically generated.
package pain_008_001_02

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_CstmrDrctDbtInitn                                                                                                                = "CstmrDrctDbtInitn"
	Path_CstmrDrctDbtInitn_GrpHdr                                                                                                         = "CstmrDrctDbtInitn.GrpHdr"
	Path_CstmrDrctDbtInitn_GrpHdr_MsgId                                                                                                   = "CstmrDrctDbtInitn.GrpHdr.MsgId"
	Path_CstmrDrctDbtInitn_GrpHdr_CreDtTm                                                                                                 = "CstmrDrctDbtInitn.GrpHdr.CreDtTm"
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Item                                                                                            = "CstmrDrctDbtInitn.GrpHdr.Authstn[]"
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn                                                                                                 = "CstmrDrctDbtInitn.GrpHdr.Authstn" // ARRAY
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Cd                                                                                              = "CstmrDrctDbtInitn.GrpHdr.Authstn[].Cd"
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Prtry                                                                                           = "CstmrDrctDbtInitn.GrpHdr.Authstn[].Prtry"
	Path_CstmrDrctDbtInitn_GrpHdr_NbOfTxs                                                                                                 = "CstmrDrctDbtInitn.GrpHdr.NbOfTxs"
	Path_CstmrDrctDbtInitn_GrpHdr_CtrlSum                                                                                                 = "CstmrDrctDbtInitn.GrpHdr.CtrlSum"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty                                                                                                = "CstmrDrctDbtInitn.GrpHdr.InitgPty"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Nm                                                                                             = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Nm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr                                                                                        = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrTp                                                                                  = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_Dept                                                                                   = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_SubDept                                                                                = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_StrtNm                                                                                 = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_BldgNb                                                                                 = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_PstCd                                                                                  = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_TwnNm                                                                                  = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn                                                                            = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_Ctry                                                                                   = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item                                                                           = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrLine                                                                                = "CstmrDrctDbtInitn.GrpHdr.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id                                                                                             = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId                                                                                       = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_BICOrBEI                                                                              = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item                                                                             = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr                                                                                  = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id                                                                               = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm                                                                          = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd                                                                       = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry                                                                    = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr                                                                             = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId                                                                                      = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth                                                                      = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                              = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                          = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                          = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                          = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item                                                                            = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr                                                                                 = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id                                                                              = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm                                                                         = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd                                                                      = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry                                                                   = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr                                                                            = "CstmrDrctDbtInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtryOfRes                                                                                      = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtryOfRes"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls                                                                                       = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx                                                                                = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_Nm                                                                                    = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_PhneNb                                                                                = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_MobNb                                                                                 = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_FaxNb                                                                                 = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr                                                                              = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_Othr                                                                                  = "CstmrDrctDbtInitn.GrpHdr.InitgPty.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt                                                                                                 = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId                                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_BIC                                                                                  = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId                                                                          = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                                 = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                              = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                           = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId                                                                    = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Nm                                                                                   = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr                                                                              = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp                                                                        = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm                                                                       = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb                                                                       = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd                                                                        = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm                                                                        = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                                  = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item                                                                 = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr                                                                                 = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id                                                                              = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry                                                                   = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr                                                                            = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId                                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_Id                                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_Nm                                                                                      = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr                                                                                 = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp                                                                           = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept                                                                            = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm                                                                          = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb                                                                          = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd                                                                           = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm                                                                           = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                     = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry                                                                            = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item                                                                    = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine                                                                         = "CstmrDrctDbtInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_Item                                                                                                    = "CstmrDrctDbtInitn.PmtInf[]"
	Path_CstmrDrctDbtInitn_PmtInf                                                                                                         = "CstmrDrctDbtInitn.PmtInf" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_PmtInfId                                                                                                = "CstmrDrctDbtInitn.PmtInf[].PmtInfId"
	Path_CstmrDrctDbtInitn_PmtInf_PmtMtd                                                                                                  = "CstmrDrctDbtInitn.PmtInf[].PmtMtd"
	Path_CstmrDrctDbtInitn_PmtInf_BtchBookg                                                                                               = "CstmrDrctDbtInitn.PmtInf[].BtchBookg"
	Path_CstmrDrctDbtInitn_PmtInf_NbOfTxs                                                                                                 = "CstmrDrctDbtInitn.PmtInf[].NbOfTxs"
	Path_CstmrDrctDbtInitn_PmtInf_CtrlSum                                                                                                 = "CstmrDrctDbtInitn.PmtInf[].CtrlSum"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf                                                                                                = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_InstrPrty                                                                                      = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.InstrPrty"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl                                                                                         = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.SvcLvl"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl_Cd                                                                                      = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.SvcLvl.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl_Prtry                                                                                   = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.SvcLvl.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm                                                                                      = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.LclInstrm"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm_Cd                                                                                   = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm_Prtry                                                                                = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SeqTp                                                                                          = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.SeqTp"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp                                                                                       = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.CtgyPurp"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp_Cd                                                                                    = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry                                                                                 = "CstmrDrctDbtInitn.PmtInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_ReqdColltnDt                                                                                            = "CstmrDrctDbtInitn.PmtInf[].ReqdColltnDt"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr                                                                                                    = "CstmrDrctDbtInitn.PmtInf[].Cdtr"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Nm                                                                                                 = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr                                                                                            = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrTp                                                                                      = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_Dept                                                                                       = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_SubDept                                                                                    = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_StrtNm                                                                                     = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_BldgNb                                                                                     = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_PstCd                                                                                      = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_TwnNm                                                                                      = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_CtrySubDvsn                                                                                = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_Ctry                                                                                       = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrLine_Item                                                                               = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrLine                                                                                    = "CstmrDrctDbtInitn.PmtInf[].Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id                                                                                                 = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId                                                                                           = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_BICOrBEI                                                                                  = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Item                                                                                 = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr                                                                                      = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Id                                                                                   = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm                                                                              = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                                                                           = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                                                                        = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Issr                                                                                 = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId                                                                                          = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth                                                                          = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                                  = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                              = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                              = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                              = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Item                                                                                = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr                                                                                     = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Id                                                                                  = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm                                                                             = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                                                                          = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                                       = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Issr                                                                                = "CstmrDrctDbtInitn.PmtInf[].Cdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtryOfRes                                                                                          = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls                                                                                           = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_NmPrfx                                                                                    = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_Nm                                                                                        = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_PhneNb                                                                                    = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_MobNb                                                                                     = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_FaxNb                                                                                     = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_EmailAdr                                                                                  = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_Othr                                                                                      = "CstmrDrctDbtInitn.PmtInf[].Cdtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct                                                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id                                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_IBAN                                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr                                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_Id                                                                                     = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm                                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm_Cd                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm_Prtry                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_Issr                                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp                                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp_Cd                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp_Prtry                                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Ccy                                                                                            = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Nm                                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt                                                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId                                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_BIC                                                                                  = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Nm                                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_Dept                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_SubDept                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_PstCd                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                                  = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_Ctry                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr                                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_Id                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_Issr                                                                            = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId                                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_Id                                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_Nm                                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr                                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrTp                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_Dept                                                                            = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_SubDept                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_StrtNm                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_BldgNb                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_PstCd                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_TwnNm                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                     = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_Ctry                                                                            = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrLine                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct                                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_IBAN                                                                                     = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr                                                                                     = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_Id                                                                                  = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_Issr                                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp_Cd                                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp_Prtry                                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Ccy                                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Nm                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrAgtAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr                                                                                               = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Nm                                                                                            = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr                                                                                       = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrTp                                                                                 = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_Dept                                                                                  = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_SubDept                                                                               = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_StrtNm                                                                                = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_BldgNb                                                                                = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_PstCd                                                                                 = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_TwnNm                                                                                 = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_CtrySubDvsn                                                                           = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_Ctry                                                                                  = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrLine_Item                                                                          = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrLine                                                                               = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id                                                                                            = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId                                                                                      = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_BICOrBEI                                                                             = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Item                                                                            = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr                                                                                 = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Id                                                                              = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm                                                                         = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                                                                      = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                                                                   = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Issr                                                                            = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId                                                                                     = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                                                                     = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                             = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                         = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                         = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                         = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Item                                                                           = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr                                                                                = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Id                                                                             = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                                                                        = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                                                                     = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                                  = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Issr                                                                           = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtryOfRes                                                                                     = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls                                                                                      = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_NmPrfx                                                                               = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_Nm                                                                                   = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_PhneNb                                                                               = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_MobNb                                                                                = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_FaxNb                                                                                = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_EmailAdr                                                                             = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_Othr                                                                                 = "CstmrDrctDbtInitn.PmtInf[].UltmtCdtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgBr                                                                                                  = "CstmrDrctDbtInitn.PmtInf[].ChrgBr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct                                                                                               = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id                                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_IBAN                                                                                       = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr                                                                                       = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_Id                                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm                                                                               = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry                                                                         = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_Issr                                                                                  = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp                                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp_Cd                                                                                         = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp_Prtry                                                                                      = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Ccy                                                                                           = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Nm                                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt                                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId                                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BIC                                                                             = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId                                                                     = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                         = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                      = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId                                                               = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm                                                                              = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr                                                                         = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp                                                                   = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm                                                                  = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb                                                                  = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd                                                                   = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm                                                                   = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                             = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id                                                                         = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry                                                              = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr                                                                       = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId                                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id                                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm                                                                                 = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr                                                                            = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp                                                                      = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept                                                                       = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm                                                                     = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb                                                                     = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd                                                                      = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm                                                                      = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry                                                                       = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item                                                               = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine                                                                    = "CstmrDrctDbtInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId                                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Nm                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr                                                                                     = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrTp                                                                               = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_Dept                                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_SubDept                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_StrtNm                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_BldgNb                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_PstCd                                                                               = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_TwnNm                                                                               = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_CtrySubDvsn                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_Ctry                                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrLine_Item                                                                        = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrLine                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id                                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId                                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_BICOrBEI                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Item                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr                                                                               = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Id                                                                            = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm                                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Issr                                                                          = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId                                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                       = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Item                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Id                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm                                                                      = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry                                                                = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Issr                                                                         = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtryOfRes                                                                                   = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls                                                                                    = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_NmPrfx                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_Nm                                                                                 = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_PhneNb                                                                             = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_MobNb                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_FaxNb                                                                              = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_EmailAdr                                                                           = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_Othr                                                                               = "CstmrDrctDbtInitn.PmtInf[].CdtrSchmeId.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Item                                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf                                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId                                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId_InstrId                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtId.InstrId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId_EndToEndId                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtId.EndToEndId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_InstrPrty                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.InstrPrty"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.SvcLvl"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl_Cd                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.SvcLvl.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl_Prtry                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.SvcLvl.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.LclInstrm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm_Cd                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm_Prtry                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SeqTp                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.SeqTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.CtgyPurp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp_Cd                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp_Prtry                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].InstdAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt_Ccy                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].InstdAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt_Value                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].InstdAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_ChrgBr                                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].ChrgBr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx                                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_MndtId                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.MndtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_DtOfSgntr                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.DtOfSgntr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInd                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlMndtId                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlMndtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Nm                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Dept                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_SubDept                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_StrtNm                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNb                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstCd                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnNm                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_CtrySubDvsn                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Ctry                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine_Item                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_BICOrBEI                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Item                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Id                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Issr                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Item                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Id                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Issr                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtryOfRes                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_NmPrfx                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Nm                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PhneNb                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_MobNb                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_FaxNb                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailAdr                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_BIC                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_MmbId              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Nm                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Dept                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_SubDept                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_StrtNm                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNb                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstCd                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnNm                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Ctry                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine_Item           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Id                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Cd                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Prtry             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Issr                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Id                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Nm                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Dept                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_SubDept                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_StrtNm                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNb                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstCd                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnNm                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_CtrySubDvsn               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Ctry                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine_Item              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_IBAN                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Id                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Cd                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Prtry                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Issr                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Cd                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Prtry                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Ccy                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Nm                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Nm                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Dept                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_SubDept                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_StrtNm                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNb                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstCd                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnNm                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_CtrySubDvsn                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Ctry                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine_Item                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_BICOrBEI                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Item                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Id                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Cd                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Prtry                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Issr                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Item                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Id                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Cd                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Issr                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtryOfRes                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_NmPrfx                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Nm                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PhneNb                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_MobNb                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_FaxNb                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailAdr                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_IBAN                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Id                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Cd                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Prtry                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Issr                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Cd                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Prtry                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Ccy                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Nm                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_BIC                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_MmbId              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Nm                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Dept                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_SubDept                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_StrtNm                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNb                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstCd                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnNm                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Ctry                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine_Item           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Id                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Cd                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Prtry             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Issr                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Id                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Nm                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Dept                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_SubDept                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_StrtNm                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNb                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstCd                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnNm                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_CtrySubDvsn               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Ctry                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine_Item              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_IBAN                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Id                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Cd                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Prtry                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Issr                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Cd                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Prtry                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Ccy                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Nm                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlFnlColltnDt                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlFnlColltnDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.AmdmntInfDtls.OrgnlFrqcy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_ElctrncSgntr                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.ElctrncSgntr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_FrstColltnDt                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.FrstColltnDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_FnlColltnDt                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.FnlColltnDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_Frqcy                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.MndtRltdInf.Frqcy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Nm                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrTp                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_Dept                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_SubDept                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_StrtNm                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_BldgNb                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_PstCd                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_TwnNm                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_CtrySubDvsn                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_Ctry                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrLine_Item                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrLine                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_BICOrBEI                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Item                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Id                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Issr                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Item                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Id                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Issr                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtryOfRes                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_NmPrfx                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_Nm                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_PhneNb                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_MobNb                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_FaxNb                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_EmailAdr                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_Othr                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.CdtrSchmeId.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_PreNtfctnId                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.PreNtfctnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_PreNtfctnDt                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DrctDbtTx.PreNtfctnDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr                                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Nm                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrTp                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_Dept                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_SubDept                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_StrtNm                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_BldgNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_PstCd                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_TwnNm                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_Ctry                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrLine_Item                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrLine                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_BICOrBEI                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Item                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Id                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Issr                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Item                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Id                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtryOfRes                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_NmPrfx                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_Nm                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_PhneNb                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_MobNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_FaxNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_EmailAdr                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_Othr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtCdtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt                                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_BIC                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.BIC"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Nm                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_Dept                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_SubDept                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_PstCd                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_Ctry                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_Id                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_Issr                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_Id                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_Nm                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrTp                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_Dept                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_SubDept                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_StrtNm                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_BldgNb                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_PstCd                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_TwnNm                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_Ctry                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrLine                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_IBAN                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_Id                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_Issr                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp_Cd                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp_Prtry                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Ccy                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Nm                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAgtAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr                                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Nm                                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrTp                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_Dept                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_SubDept                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_StrtNm                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_BldgNb                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_PstCd                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_TwnNm                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_CtrySubDvsn                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_Ctry                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrLine_Item                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrLine                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id                                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_BICOrBEI                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Item                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Id                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Issr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Item                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Id                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Issr                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtryOfRes                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_NmPrfx                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_Nm                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_PhneNb                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_MobNb                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_FaxNb                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_EmailAdr                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_Othr                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Dbtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_IBAN                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.IBAN"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_Id                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm_Cd                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm_Prtry                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_Issr                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Id.Othr.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp_Cd                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Tp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp_Prtry                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Tp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Ccy                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Nm                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].DbtrAcct.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr                                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Nm                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrTp                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_Dept                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_SubDept                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_StrtNm                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_BldgNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_PstCd                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_TwnNm                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_Ctry                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrLine_Item                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrLine                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_BICOrBEI                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Item                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Id                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Issr                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Item                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Id                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtryOfRes                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_NmPrfx                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_Nm                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_PhneNb                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_MobNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_FaxNb                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_EmailAdr                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_Othr                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].UltmtDbtr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstrForCdtrAgt                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].InstrForCdtrAgt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp                                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Purp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp_Cd                                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Purp.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp_Prtry                                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Purp.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Item                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg                                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_DbtCdtRptgInd                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].DbtCdtRptgInd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Authrty"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty_Nm                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Authrty.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty_Ctry                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Authrty.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Item                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Tp                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Dt                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Dt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Ctry                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Cd                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Amt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt_Ccy                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Amt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt_Value                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Amt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Inf_Item                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Inf[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Inf                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RgltryRptg[].Dtls[].Inf" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax                                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Cdtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_TaxId                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Cdtr.TaxId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_RegnId                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Cdtr.RegnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_TaxTp                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Cdtr.TaxTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_TaxId                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.TaxId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_RegnId                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.RegnId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_TaxTp                                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.TaxTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.Authstn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn_Titl                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.Authstn.Titl"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn_Nm                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dbtr.Authstn.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_AdmstnZn                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.AdmstnZn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_RefNb                                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.RefNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Mtd                                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Mtd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxblBaseAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt_Ccy                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxblBaseAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt_Value                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxblBaseAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt_Ccy                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt_Value                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.TtlTaxAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dt                                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Dt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_SeqNb                                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.SeqNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Item                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd                                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Tp                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Ctgy                                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Ctgy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_CtgyDtls                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].CtgyDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_DbtrSts                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].DbtrSts"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_CertId                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].CertId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_FrmsCd                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].FrmsCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_Yr                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd.Yr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_Tp                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt                                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd.FrToDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt_FrDt                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd.FrToDt.FrDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt_ToDt                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].Prd.FrToDt.ToDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Rate                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Rate"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Item                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt                                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_AddtlInf                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].Tax.Rcrd[].AddtlInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_Item                                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf                                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtId                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnMtd                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnMtd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnElctrncAdr                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnElctrncAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Nm                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf                                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Ustrd_Item                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Ustrd[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Ustrd                                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Ustrd" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Item                                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd                                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Item                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Nb                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_RltdDt                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf                                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Tp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Ref                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].CdtrRefInf.Ref"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr                                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Nm                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr                                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id                                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtryOfRes                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm                                                              = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee                                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Nm                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr                                                                 = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry                                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id                                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.BICOrBEI"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                                                        = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                                                      = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                   = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                                                  = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry                                            = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                                                     = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtryOfRes                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtryOfRes"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls                                                                = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm                                                             = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb                                                         = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb                                                          = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                                                       = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr                                                           = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_AddtlRmtInf_Item                                                               = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].AddtlRmtInf[]"
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_AddtlRmtInf                                                                    = "CstmrDrctDbtInitn.PmtInf[].DrctDbtTxInf[].RmtInf.Strd[].AddtlRmtInf" // ARRAY
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_CstmrDrctDbtInitn:                                                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr:                                                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_MsgId:                                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_CreDtTm:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Item:                                                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn:                                                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Cd:                                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_Authstn_Prtry:                                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_NbOfTxs:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_CtrlSum:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty:                                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Nm:                                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr:                                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrTp:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_Dept:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_SubDept:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_StrtNm:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_BldgNb:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_PstCd:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_TwnNm:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_Ctry:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item:                                                                           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_PstlAdr_AdrLine:                                                                                PathTypeArray,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id:                                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_BICOrBEI:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item:                                                                             PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr:                                                                                  PathTypeArray,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth:                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item:                                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr:                                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtryOfRes:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_Nm:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_PhneNb:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_MobNb:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_FaxNb:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_InitgPty_CtctDtls_Othr:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt:                                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_BIC:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Nm:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                                 PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine:                                                                      PathTypeArray,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr:                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId:                                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_Id:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_Nm:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr:                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item:                                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine:                                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_Item:                                                                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf:                                                                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_PmtInfId:                                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtMtd:                                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_BtchBookg:                                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_NbOfTxs:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CtrlSum:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf:                                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_InstrPrty:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl:                                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl_Cd:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SvcLvl_Prtry:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm_Cd:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_LclInstrm_Prtry:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_SeqTp:                                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp_Cd:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ReqdColltnDt:                                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr:                                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Nm:                                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr:                                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrTp:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_Dept:                                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_SubDept:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_StrtNm:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_BldgNb:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_PstCd:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_TwnNm:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_CtrySubDvsn:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_Ctry:                                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrLine_Item:                                                                               PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_PstlAdr_AdrLine:                                                                                    PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id:                                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId:                                                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_BICOrBEI:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Item:                                                                                 PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr:                                                                                      PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Id:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_OrgId_Othr_Issr:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId:                                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Item:                                                                                PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr:                                                                                     PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Id:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm:                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_Id_PrvtId_Othr_Issr:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtryOfRes:                                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls:                                                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_NmPrfx:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_Nm:                                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_PhneNb:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_MobNb:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_FaxNb:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_EmailAdr:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_Cdtr_CtctDtls_Othr:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct:                                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id:                                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_IBAN:                                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr:                                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_Id:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm:                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm_Cd:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_SchmeNm_Prtry:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Id_Othr_Issr:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp:                                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp_Cd:                                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Tp_Prtry:                                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Ccy:                                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAcct_Nm:                                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt:                                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_BIC:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Nm:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_Dept:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_SubDept:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_PstCd:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_Ctry:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                                 PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                                                                      PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr:                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_Id:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_FinInstnId_Othr_Issr:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId:                                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_Id:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_Nm:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr:                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrTp:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_Dept:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_SubDept:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_StrtNm:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_BldgNb:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_PstCd:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_TwnNm:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_Ctry:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgt_BrnchId_PstlAdr_AdrLine:                                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct:                                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id:                                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_IBAN:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr:                                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_Id:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm:                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Id_Othr_Issr:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp:                                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp_Cd:                                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Tp_Prtry:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Ccy:                                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrAgtAcct_Nm:                                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr:                                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Nm:                                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrTp:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_Dept:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_SubDept:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_StrtNm:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_BldgNb:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_PstCd:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_TwnNm:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_CtrySubDvsn:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_Ctry:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrLine_Item:                                                                          PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_PstlAdr_AdrLine:                                                                               PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id:                                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_BICOrBEI:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Item:                                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr:                                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Id:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_OrgId_Othr_Issr:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId:                                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Item:                                                                           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr:                                                                                PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Id:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_Id_PrvtId_Othr_Issr:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtryOfRes:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_NmPrfx:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_Nm:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_PhneNb:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_MobNb:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_FaxNb:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_EmailAdr:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_UltmtCdtr_CtctDtls_Othr:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgBr:                                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct:                                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id:                                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_IBAN:                                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_Id:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm:                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Id_Othr_Issr:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp:                                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp_Cd:                                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Tp_Prtry:                                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Ccy:                                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcct_Nm:                                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt:                                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId:                                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BIC:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId:                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine:                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr:                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm:                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId:                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr:                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item:                                                               PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine:                                                                    PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId:                                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Nm:                                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr:                                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrTp:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_Dept:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_SubDept:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_StrtNm:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_BldgNb:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_PstCd:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_TwnNm:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_CtrySubDvsn:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_Ctry:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrLine_Item:                                                                        PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_PstlAdr_AdrLine:                                                                             PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id:                                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId:                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_BICOrBEI:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Item:                                                                          PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr:                                                                               PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Id:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm:                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_OrgId_Othr_Issr:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth:                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Item:                                                                         PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr:                                                                              PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Id:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm:                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_Id_PrvtId_Othr_Issr:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtryOfRes:                                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls:                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_NmPrfx:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_Nm:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_PhneNb:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_MobNb:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_FaxNb:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_EmailAdr:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_CdtrSchmeId_CtctDtls_Othr:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Item:                                                                                       PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf:                                                                                            PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId:                                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId_InstrId:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtId_EndToEndId:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_InstrPrty:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl:                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl_Cd:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SvcLvl_Prtry:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm_Cd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_LclInstrm_Prtry:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_SeqTp:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp_Cd:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_PmtTpInf_CtgyPurp_Prtry:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt_Ccy:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstdAmt_Value:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_ChrgBr:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx:                                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf:                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_MndtId:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_DtOfSgntr:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInd:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlMndtId:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId:                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Nm:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr:                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId:                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr:                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId:                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr:                        PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtryOfRes:                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls:                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Nm:                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_MobNb:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr:                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt:                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId:                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_BIC:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId:                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_MmbId:              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Nm:                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr:                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Dept:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_SubDept:                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_StrtNm:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNb:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstCd:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnNm:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Ctry:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine:                PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr:                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Id:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm:                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Issr:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId:                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Id:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Nm:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr:                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Dept:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_SubDept:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_StrtNm:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNb:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstCd:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnNm:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Ctry:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine_Item:              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine:                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct:                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_IBAN:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr:                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Id:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Issr:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Cd:                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Prtry:                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Ccy:                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Nm:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr:                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Nm:                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr:                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Dept:                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_SubDept:                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_StrtNm:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNb:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstCd:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnNm:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_CtrySubDvsn:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Ctry:                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine_Item:                         PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine:                              PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id:                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId:                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_BICOrBEI:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Item:                           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr:                                PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Id:                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm:                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Cd:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Issr:                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth:                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Item:                          PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr:                               PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Id:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtryOfRes:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls:                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_NmPrfx:                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Nm:                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PhneNb:                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_MobNb:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_FaxNb:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailAdr:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct:                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id:                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_IBAN:                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr:                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Id:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm:                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Cd:                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Prtry:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Issr:                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp:                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Cd:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Prtry:                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Ccy:                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Nm:                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt:                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId:                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_BIC:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId:                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_MmbId:              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Nm:                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr:                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Dept:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_SubDept:                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_StrtNm:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNb:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstCd:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnNm:                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Ctry:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine:                PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr:                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Id:                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm:                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Issr:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId:                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Id:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Nm:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr:                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Dept:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_SubDept:                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_StrtNm:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNb:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstCd:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnNm:                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Ctry:                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine_Item:              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine:                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct:                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_IBAN:                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr:                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Id:                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Issr:                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp:                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Cd:                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Prtry:                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Ccy:                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Nm:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlFnlColltnDt:                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy:                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_ElctrncSgntr:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_FrstColltnDt:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_FnlColltnDt:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_MndtRltdInf_Frqcy:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId:                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Nm:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr:                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrTp:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_Dept:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_SubDept:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_StrtNm:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_BldgNb:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_PstCd:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_TwnNm:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_CtrySubDvsn:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_Ctry:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrLine_Item:                                                 PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_PstlAdr_AdrLine:                                                      PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id:                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId:                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_BICOrBEI:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Item:                                                   PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr:                                                        PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Id:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_OrgId_Othr_Issr:                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth:                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Item:                                                  PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr:                                                       PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Id:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm:                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_Id_PrvtId_Othr_Issr:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtryOfRes:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls:                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_NmPrfx:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_Nm:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_PhneNb:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_MobNb:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_FaxNb:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_EmailAdr:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_CdtrSchmeId_CtctDtls_Othr:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_PreNtfctnId:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DrctDbtTx_PreNtfctnDt:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr:                                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Nm:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrTp:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_Dept:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_SubDept:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_StrtNm:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_BldgNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_PstCd:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_TwnNm:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_Ctry:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrLine_Item:                                                             PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_PstlAdr_AdrLine:                                                                  PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id:                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_BICOrBEI:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Item:                                                               PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr:                                                                    PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Id:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_OrgId_Othr_Issr:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Item:                                                              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr:                                                                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Id:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtryOfRes:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_NmPrfx:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_Nm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_PhneNb:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_MobNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_FaxNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_EmailAdr:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtCdtr_CtctDtls_Othr:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt:                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_BIC:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId:                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Nm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_Dept:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr:                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_Id:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_FinInstnId_Othr_Issr:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId:                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_Id:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_Nm:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr:                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_Dept:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_SubDept:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_PstCd:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_Ctry:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                       PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                                            PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct:                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id:                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_IBAN:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_Id:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm:                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Id_Othr_Issr:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp:                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp_Cd:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Tp_Prtry:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Ccy:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAgtAcct_Nm:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Nm:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr:                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrTp:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_Dept:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_SubDept:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_StrtNm:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_BldgNb:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_PstCd:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_TwnNm:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_CtrySubDvsn:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_Ctry:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrLine_Item:                                                                  PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_PstlAdr_AdrLine:                                                                       PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id:                                                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_BICOrBEI:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Item:                                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr:                                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Id:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_OrgId_Othr_Issr:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId:                                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Item:                                                                   PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr:                                                                        PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Id:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm:                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_Id_PrvtId_Othr_Issr:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtryOfRes:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_NmPrfx:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_Nm:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_PhneNb:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_MobNb:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_FaxNb:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_EmailAdr:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Dbtr_CtctDtls_Othr:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id:                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_IBAN:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr:                                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_Id:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm:                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm_Cd:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_SchmeNm_Prtry:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Id_Othr_Issr:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp:                                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp_Cd:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Tp_Prtry:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Ccy:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_DbtrAcct_Nm:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr:                                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Nm:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrTp:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_Dept:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_SubDept:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_StrtNm:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_BldgNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_PstCd:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_TwnNm:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_Ctry:                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrLine_Item:                                                             PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_PstlAdr_AdrLine:                                                                  PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id:                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_BICOrBEI:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Item:                                                               PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr:                                                                    PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Id:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_OrgId_Othr_Issr:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Item:                                                              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr:                                                                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Id:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtryOfRes:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_NmPrfx:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_Nm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_PhneNb:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_MobNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_FaxNb:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_EmailAdr:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_UltmtDbtr_CtctDtls_Othr:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_InstrForCdtrAgt:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp:                                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp_Cd:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Purp_Prtry:                                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Item:                                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg:                                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_DbtCdtRptgInd:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty_Nm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Authrty_Ctry:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Item:                                                                       PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls:                                                                            PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Tp:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Dt:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Ctry:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Cd:                                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt_Ccy:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Amt_Value:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Inf_Item:                                                                   PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RgltryRptg_Dtls_Inf:                                                                        PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax:                                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_TaxId:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_RegnId:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Cdtr_TaxTp:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr:                                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_TaxId:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_RegnId:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_TaxTp:                                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn:                                                                           PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn_Titl:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dbtr_Authstn_Nm:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_AdmstnZn:                                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_RefNb:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Mtd:                                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt_Ccy:                                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxblBaseAmt_Value:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt:                                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt_Ccy:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_TtlTaxAmt_Value:                                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Dt:                                                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_SeqNb:                                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Item:                                                                              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd:                                                                                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Tp:                                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Ctgy:                                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_CtgyDtls:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_DbtrSts:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_CertId:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_FrmsCd:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd:                                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_Yr:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_Tp:                                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt:                                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt_FrDt:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_Prd_FrToDt_ToDt:                                                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt:                                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Rate:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt:                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt:                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Item:                                                                  PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls:                                                                       PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd:                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                                                            PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt:                                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_Tax_Rcrd_AddtlInf:                                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_Item:                                                                            PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf:                                                                                 PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtId:                                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnMtd:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnElctrncAdr:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr:                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Nm:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr:                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item:                                                 PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine:                                                      PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf:                                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Ustrd_Item:                                                                          PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Ustrd:                                                                               PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Item:                                                                           PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd:                                                                                PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Item:                                                                PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf:                                                                     PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp:                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_Nb:                                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocInf_RltdDt:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt:                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt:                                                              PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:                                              PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                                                   PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                                                             PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf:                                                                     PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp:                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                                                        PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                                                  PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_CdtrRefInf_Ref:                                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr:                                                                          PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Nm:                                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr:                                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                                                     PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                                                          PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id:                                                                       PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                                                       PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr:                                                            PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                                                    PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:                                                 PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId:                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                    PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                                                      PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                                                           PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtryOfRes:                                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm:                                                              PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee:                                                                         PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Nm:                                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr:                                                                 PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry:                                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                                                    PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                                                         PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id:                                                                      PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId:                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                                                      PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr:                                                           PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                                                        PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                                                   PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:                                                PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                                                      PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId:                                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:                                               PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                   PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                                                     PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                                                          PathTypeArray,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                                                  PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:                                            PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                                                     PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtryOfRes:                                                               PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls:                                                                PathTypeStruct,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm:                                                             PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                                                         PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                                                          PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                                                       PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr:                                                           PathTypeProperty,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_AddtlRmtInf_Item:                                                               PathTypeArrayItem,
	Path_CstmrDrctDbtInitn_PmtInf_DrctDbtTxInf_RmtInf_Strd_AddtlRmtInf:                                                                    PathTypeArray,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in pain_008_001_02", p)
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
