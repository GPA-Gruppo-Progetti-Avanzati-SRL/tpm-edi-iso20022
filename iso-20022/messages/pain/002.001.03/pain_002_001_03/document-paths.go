// Package pain_002_001_03
// Do not Edit. This stuff it's been automatically generated.
package pain_002_001_03

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_CstmrPmtStsRpt                                                                                                                           = "CstmrPmtStsRpt"
	Path_CstmrPmtStsRpt_GrpHdr                                                                                                                    = "CstmrPmtStsRpt.GrpHdr"
	Path_CstmrPmtStsRpt_GrpHdr_MsgId                                                                                                              = "CstmrPmtStsRpt.GrpHdr.MsgId"
	Path_CstmrPmtStsRpt_GrpHdr_CreDtTm                                                                                                            = "CstmrPmtStsRpt.GrpHdr.CreDtTm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty                                                                                                           = "CstmrPmtStsRpt.GrpHdr.InitgPty"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Nm                                                                                                        = "CstmrPmtStsRpt.GrpHdr.InitgPty.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr                                                                                                   = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp                                                                                             = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_Dept                                                                                              = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_SubDept                                                                                           = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_StrtNm                                                                                            = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_BldgNb                                                                                            = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_PstCd                                                                                             = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_TwnNm                                                                                             = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn                                                                                       = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_Ctry                                                                                              = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrLine_Item                                                                                      = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrLine                                                                                           = "CstmrPmtStsRpt.GrpHdr.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id                                                                                                        = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId                                                                                                  = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_BICOrBEI                                                                                         = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Item                                                                                        = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr                                                                                             = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Id                                                                                          = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm                                                                                     = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd                                                                                  = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry                                                                               = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Issr                                                                                        = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId                                                                                                 = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth                                                                                 = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                                         = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                                     = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                                     = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                                     = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Item                                                                                       = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr                                                                                            = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Id                                                                                         = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm                                                                                    = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd                                                                                 = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry                                                                              = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr                                                                                       = "CstmrPmtStsRpt.GrpHdr.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtryOfRes                                                                                                 = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtryOfRes"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls                                                                                                  = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_NmPrfx                                                                                           = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_Nm                                                                                               = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_PhneNb                                                                                           = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_MobNb                                                                                            = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_FaxNb                                                                                            = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_EmailAdr                                                                                         = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_Othr                                                                                             = "CstmrPmtStsRpt.GrpHdr.InitgPty.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt                                                                                                            = "CstmrPmtStsRpt.GrpHdr.FwdgAgt"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId                                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_BIC                                                                                             = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId                                                                                     = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                                            = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                                         = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                                      = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId                                                                               = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Nm                                                                                              = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr                                                                                         = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp                                                                                   = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm                                                                                  = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb                                                                                  = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd                                                                                   = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm                                                                                   = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                                             = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item                                                                            = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr                                                                                            = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_Id                                                                                         = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry                                                                              = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr                                                                                       = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId                                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_Id                                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_Nm                                                                                                 = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr                                                                                            = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp                                                                                      = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept                                                                                       = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm                                                                                     = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb                                                                                     = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd                                                                                      = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm                                                                                      = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                                = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry                                                                                       = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item                                                                               = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine                                                                                    = "CstmrPmtStsRpt.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt                                                                                                            = "CstmrPmtStsRpt.GrpHdr.DbtrAgt"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId                                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_BIC                                                                                             = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId                                                                                     = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                                            = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                                         = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                                      = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                                               = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Nm                                                                                              = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr                                                                                         = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                                                                   = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Dept                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_SubDept                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                                                                  = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                                                                  = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_PstCd                                                                                   = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                                                                   = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                                             = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Ctry                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                                                            = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr                                                                                            = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Id                                                                                         = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                                                              = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Issr                                                                                       = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId                                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_Id                                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_Nm                                                                                                 = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr                                                                                            = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp                                                                                      = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Dept                                                                                       = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_SubDept                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_StrtNm                                                                                     = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_BldgNb                                                                                     = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_PstCd                                                                                      = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_TwnNm                                                                                      = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                                = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Ctry                                                                                       = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                                               = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrLine                                                                                    = "CstmrPmtStsRpt.GrpHdr.DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt                                                                                                            = "CstmrPmtStsRpt.GrpHdr.CdtrAgt"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId                                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_BIC                                                                                             = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId                                                                                     = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                                            = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                                                         = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                                                      = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                                               = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Nm                                                                                              = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr                                                                                         = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp                                                                                   = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Dept                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_SubDept                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_StrtNm                                                                                  = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_BldgNb                                                                                  = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_PstCd                                                                                   = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_TwnNm                                                                                   = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                                             = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Ctry                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                                                            = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrLine                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr                                                                                            = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Id                                                                                         = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                                                              = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Issr                                                                                       = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId                                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_Id                                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_Nm                                                                                                 = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr                                                                                            = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp                                                                                      = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Dept                                                                                       = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_SubDept                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_StrtNm                                                                                     = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_BldgNb                                                                                     = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_PstCd                                                                                      = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_TwnNm                                                                                      = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                                                = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Ctry                                                                                       = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                                               = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrLine                                                                                    = "CstmrPmtStsRpt.GrpHdr.CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts                                                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlMsgId                                                                                              = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlMsgNmId                                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlCreDtTm                                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlNbOfTxs                                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.OrgnlNbOfTxs"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlCtrlSum                                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.OrgnlCtrlSum"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_GrpSts                                                                                                  = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.GrpSts"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Item                                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf                                                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr                                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Nm                                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr                                                                                 = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn                                                                     = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry                                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item                                                                    = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id                                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId                                                                                = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI                                                                       = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id                                                                        = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm                                                                   = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd                                                                = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry                                                             = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId                                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                       = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                   = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                   = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                   = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item                                                                     = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id                                                                       = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm                                                                  = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                            = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr                                                                     = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtryOfRes                                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls                                                                                = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm                                                                             = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb                                                                         = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb                                                                          = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr                                                                       = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Orgtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn                                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Rsn"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Cd                                                                                        = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Rsn.Cd"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Prtry                                                                                     = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].Rsn.Prtry"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_AddtlInf_Item                                                                                 = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].AddtlInf[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_AddtlInf                                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.StsRsnInf[].AddtlInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_Item                                                                                      = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.NbOfTxsPerSts[]"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts                                                                                           = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.NbOfTxsPerSts" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldNbOfTxs                                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.NbOfTxsPerSts[].DtldNbOfTxs"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldSts                                                                                   = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.NbOfTxsPerSts[].DtldSts"
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldCtrlSum                                                                               = "CstmrPmtStsRpt.OrgnlGrpInfAndSts.NbOfTxsPerSts[].DtldCtrlSum"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_Item                                                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts                                                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlPmtInfId                                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].OrgnlPmtInfId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlNbOfTxs                                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].OrgnlNbOfTxs"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlCtrlSum                                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].OrgnlCtrlSum"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_PmtInfSts                                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].PmtInfSts"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Item                                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf                                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr                                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Nm                                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr                                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id                                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId                                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI                                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id                                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtryOfRes                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls                                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr                                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn                                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Rsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Cd                                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Rsn.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Prtry                                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].Rsn.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_AddtlInf_Item                                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].AddtlInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_AddtlInf                                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].StsRsnInf[].AddtlInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_Item                                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].NbOfTxsPerSts[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts                                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].NbOfTxsPerSts" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldNbOfTxs                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].NbOfTxsPerSts[].DtldNbOfTxs"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldSts                                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].NbOfTxsPerSts[].DtldSts"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldCtrlSum                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].NbOfTxsPerSts[].DtldCtrlSum"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_Item                                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts                                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsId                                                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlInstrId                                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlInstrId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlEndToEndId                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlEndToEndId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_TxSts                                                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].TxSts"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Item                                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf                                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Nm                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtryOfRes                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Orgtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Rsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Cd                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Rsn.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Prtry                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].Rsn.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_AddtlInf_Item                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].AddtlInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_AddtlInf                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].StsRsnInf[].AddtlInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Item                                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf                                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt                                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Amt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Ccy                                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Amt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Value                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Amt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty                                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_BIC                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_MmbId                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Nm                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrTp                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_Dept                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_SubDept                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_StrtNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_BldgNb                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_PstCd                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_TwnNm                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_CtrySubDvsn                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_Ctry                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrLine_Item                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrLine                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_Id                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm_Cd                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm_Prtry                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_Issr                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_Id                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_Nm                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrTp                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_Dept                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_SubDept                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_StrtNm                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_BldgNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_PstCd                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_TwnNm                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_CtrySubDvsn                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_Ctry                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrLine_Item                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrLine                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ChrgsInf[].Pty.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AccptncDtTm                                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].AccptncDtTm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AcctSvcrRef                                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].AcctSvcrRef"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ClrSysRef                                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].ClrSysRef"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef                                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.IntrBkSttlmAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt_Ccy                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.IntrBkSttlmAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt_Value                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.IntrBkSttlmAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt                                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.InstdAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Ccy                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.InstdAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Value                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.InstdAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.EqvtAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.EqvtAmt.Amt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Ccy                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.EqvtAmt.Amt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Value                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.EqvtAmt.Amt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_CcyOfTrf                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Amt.EqvtAmt.CcyOfTrf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmDt                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.IntrBkSttlmDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdColltnDt                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.ReqdColltnDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdExctnDt                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.ReqdExctnDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Nm                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_Dept                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_SubDept                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_StrtNm                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_BldgNb                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_PstCd                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_TwnNm                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_CtrySubDvsn                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_Ctry                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrLine_Item                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrLine                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_BICOrBEI                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Item                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Id                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Issr                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Item                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Id                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Issr                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtryOfRes                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_NmPrfx                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_Nm                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_PhneNb                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_MobNb                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_FaxNb                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_EmailAdr                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_Othr                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrSchmeId.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmMtd                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmMtd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_IBAN                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Id                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Cd                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Prtry                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Issr                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Cd                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Prtry                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Ccy                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.SttlmAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ClrSys"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys_Cd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ClrSys.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys_Prtry                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ClrSys.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_BIC                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Nm                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Dept                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_SubDept                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_PstCd                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Ctry                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Id                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Issr                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Id                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Nm                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Dept                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_SubDept                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_StrtNm                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_BldgNb                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_PstCd                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_TwnNm                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Ctry                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrLine                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_IBAN                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Id                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Issr                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Cd                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Prtry                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Ccy                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Nm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstgRmbrsmntAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_BIC                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Nm                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Dept                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Id                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Issr                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Id                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Nm                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Dept                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_SubDept                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_PstCd                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Ctry                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_IBAN                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Id                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Issr                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Cd                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Prtry                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Ccy                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Nm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.InstdRmbrsmntAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_BIC                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Nm                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Dept                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Id                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Issr                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Id                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Nm                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Dept                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_SubDept                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_PstCd                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Ctry                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_IBAN                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Id                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Issr                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Cd                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Prtry                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Ccy                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Nm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.SttlmInf.ThrdRmbrsmntAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_InstrPrty                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.InstrPrty"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_ClrChanl                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.ClrChanl"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.SvcLvl"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Cd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.SvcLvl.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Prtry                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.SvcLvl.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.LclInstrm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Cd                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.LclInstrm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SeqTp                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.SeqTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.CtgyPurp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.CtgyPurp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Prtry                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtMtd                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.PmtMtd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_MndtId                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.MndtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_DtOfSgntr                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.DtOfSgntr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInd                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlMndtId                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlMndtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Nm                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Dept                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_SubDept                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_StrtNm                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNb                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstCd                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnNm                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_CtrySubDvsn                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Ctry                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine_Item                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_BICOrBEI                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Item                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Id                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Issr                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Item                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Id                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Issr                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtryOfRes                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_NmPrfx                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Nm                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PhneNb                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_MobNb                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_FaxNb                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailAdr                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrSchmeId.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_BIC                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_MmbId              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Nm                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Dept                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_SubDept                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_StrtNm                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNb                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstCd                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnNm                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Ctry                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine_Item           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Id                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Cd                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Prtry             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Issr                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Id                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Nm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Dept                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_SubDept                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_StrtNm                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNb                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstCd                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnNm                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_CtrySubDvsn               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Ctry                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine_Item              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_IBAN                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Id                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Cd                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Prtry                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Issr                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Cd                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Prtry                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Ccy                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Nm                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlCdtrAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Nm                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Dept                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_SubDept                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_StrtNm                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNb                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstCd                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnNm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_CtrySubDvsn                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Ctry                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine_Item                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_BICOrBEI                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Item                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Id                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Cd                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Prtry                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Issr                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Item                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Id                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Cd                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Issr                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtryOfRes                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_NmPrfx                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Nm                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PhneNb                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_MobNb                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_FaxNb                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailAdr                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_IBAN                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Id                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Cd                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Prtry                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Issr                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Cd                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Prtry                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Ccy                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Nm                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_BIC                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_MmbId              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Nm                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Dept                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_SubDept                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_StrtNm                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNb                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstCd                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnNm                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Ctry                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine_Item           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Id                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Cd                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Prtry             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Issr                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Id                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Nm                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Dept                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_SubDept                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_StrtNm                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNb                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstCd                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnNm                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_CtrySubDvsn               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Ctry                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine_Item              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_IBAN                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Id                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Cd                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Prtry                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Issr                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Cd                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Prtry                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Ccy                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Nm                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlDbtrAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFnlColltnDt                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlFnlColltnDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.AmdmntInfDtls.OrgnlFrqcy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_ElctrncSgntr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.ElctrncSgntr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_FrstColltnDt                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.FrstColltnDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_FnlColltnDt                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.FnlColltnDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_Frqcy                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.MndtRltdInf.Frqcy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf                                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Ustrd_Item                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Ustrd[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Ustrd                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Ustrd" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Item                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Item                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_Issr                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Nb                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_RltdDt                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DscntApldAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.TaxAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.TaxAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.TaxAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_Issr                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Ref                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].CdtrRefInf.Ref"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Nm                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Dept                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_SubDept                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_StrtNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNb                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstCd                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnNm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Ctry                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtryOfRes                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Nm                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PhneNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_MobNb                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_FaxNb                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Dept                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_SubDept                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_StrtNm                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstCd                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Ctry                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtryOfRes                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Nm                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PhneNb                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_MobNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_FaxNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].Invcee.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf_Item                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].AddtlRmtInf[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.RmtInf.Strd[].AddtlRmtInf" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Nm                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Dept                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_SubDept                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_StrtNm                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_BldgNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_PstCd                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_TwnNm                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_CtrySubDvsn                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Ctry                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrLine_Item                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrLine                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_BICOrBEI                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Item                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Id                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Issr                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Item                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Id                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Issr                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtryOfRes                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_NmPrfx                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_PhneNb                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_MobNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_FaxNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_EmailAdr                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtDbtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Nm                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Dept                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_SubDept                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_StrtNm                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_BldgNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_PstCd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_TwnNm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_CtrySubDvsn                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Ctry                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrLine_Item                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrLine                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_BICOrBEI                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Item                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Id                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Issr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Item                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Id                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Issr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtryOfRes                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_NmPrfx                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Nm                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_PhneNb                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_MobNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_FaxNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_EmailAdr                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Dbtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_IBAN                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Id                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Cd                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Prtry                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Issr                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Cd                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Prtry                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Ccy                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Nm                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_BIC                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Dept                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_SubDept                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstCd                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Ctry                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Id                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Issr                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Id                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Nm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Dept                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_SubDept                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_StrtNm                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNb                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstCd                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnNm                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Ctry                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_IBAN                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Id                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Cd                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Issr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp_Cd                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp_Prtry                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Ccy                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Nm                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.DbtrAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_BIC                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.BIC"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Dept                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_SubDept                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_StrtNm                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNb                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstCd                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Ctry                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Id                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Issr                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Id                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Nm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Dept                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_SubDept                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_StrtNm                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNb                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstCd                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnNm                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Ctry                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_IBAN                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Id                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Cd                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Issr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp_Cd                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp_Prtry                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Ccy                                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Nm                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAgtAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr                                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Nm                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Dept                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_SubDept                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_StrtNm                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_BldgNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_PstCd                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_TwnNm                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_CtrySubDvsn                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Ctry                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrLine_Item                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrLine                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id                                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_BICOrBEI                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Item                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Id                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Issr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Item                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Id                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Issr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtryOfRes                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls                                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_NmPrfx                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Nm                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_PhneNb                                                             = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_MobNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_FaxNb                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_EmailAdr                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Othr                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.Cdtr.CtctDtls.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct                                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_IBAN                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.IBAN"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr                                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Id                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr.SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Cd                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Prtry                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Issr                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Id.Othr.Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Tp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Cd                                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Tp.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Prtry                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Tp.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Ccy                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Ccy"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Nm                                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.CdtrAcct.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr                                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Nm                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr                                                                = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.AdrTp"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Dept                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.Dept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_SubDept                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.SubDept"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_StrtNm                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.StrtNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_BldgNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.BldgNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_PstCd                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.PstCd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_TwnNm                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.TwnNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_CtrySubDvsn                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Ctry                                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.Ctry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrLine_Item                                                   = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrLine                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id                                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_BICOrBEI                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.BICOrBEI"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Item                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Id                                                       = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm                                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Issr                                                     = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                  = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Item                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Id                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                                                 = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                           = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Issr                                                    = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtryOfRes                                                              = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtryOfRes"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls                                                               = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_NmPrfx                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.NmPrfx"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Nm                                                            = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.Nm"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_PhneNb                                                        = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.PhneNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_MobNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.MobNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_FaxNb                                                         = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.FaxNb"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_EmailAdr                                                      = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.EmailAdr"
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Othr                                                          = "CstmrPmtStsRpt.OrgnlPmtInfAndSts[].TxInfAndSts[].OrgnlTxRef.UltmtCdtr.CtctDtls.Othr"
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_CstmrPmtStsRpt:                                                                                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr:                                                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_MsgId:                                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CreDtTm:                                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty:                                                                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Nm:                                                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr:                                                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrTp:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_Dept:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_SubDept:                                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_StrtNm:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_BldgNb:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_PstCd:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_TwnNm:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_Ctry:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrLine_Item:                                                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_PstlAdr_AdrLine:                                                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id:                                                                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId:                                                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_BICOrBEI:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Item:                                                                                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr:                                                                                             PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Id:                                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm:                                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:                                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId:                                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth:                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Item:                                                                                       PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr:                                                                                            PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm:                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtryOfRes:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls:                                                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_NmPrfx:                                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_Nm:                                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_PhneNb:                                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_MobNb:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_FaxNb:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_EmailAdr:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_InitgPty_CtctDtls_Othr:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt:                                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId:                                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_BIC:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId:                                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Nm:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr:                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine:                                                                                 PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_Id:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm:                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId:                                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_Id:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_Nm:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item:                                                                               PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine:                                                                                    PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt:                                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId:                                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_BIC:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId:                                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Nm:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr:                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Dept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                                                                 PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Id:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm:                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_FinInstnId_Othr_Issr:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId:                                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_Id:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_Nm:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Dept:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_SubDept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_PstCd:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_Ctry:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                                               PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                                                                    PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt:                                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId:                                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_BIC:                                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId:                                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Nm:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr:                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Dept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_SubDept:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_PstCd:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_Ctry:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                                                                                 PathTypeArray,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Id:                                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm:                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_FinInstnId_Othr_Issr:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId:                                                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_Id:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_Nm:                                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr:                                                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrTp:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Dept:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_SubDept:                                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_StrtNm:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_BldgNb:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_PstCd:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_TwnNm:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_Ctry:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                                               PathTypeArrayItem,
	Path_CstmrPmtStsRpt_GrpHdr_CdtrAgt_BrnchId_PstlAdr_AdrLine:                                                                                    PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts:                                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlMsgId:                                                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlMsgNmId:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlCreDtTm:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlNbOfTxs:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_OrgnlCtrlSum:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_GrpSts:                                                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Item:                                                                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf:                                                                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr:                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Nm:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr:                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item:                                                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:                                                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id:                                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item:                                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr:                                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId:                                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item:                                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr:                                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtryOfRes:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn:                                                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Cd:                                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_Rsn_Prtry:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_AddtlInf_Item:                                                                                 PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_StsRsnInf_AddtlInf:                                                                                      PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_Item:                                                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts:                                                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldNbOfTxs:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldSts:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlGrpInfAndSts_NbOfTxsPerSts_DtldCtrlSum:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_Item:                                                                                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts:                                                                                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlPmtInfId:                                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlNbOfTxs:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_OrgnlCtrlSum:                                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_PmtInfSts:                                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Item:                                                                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf:                                                                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr:                                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Nm:                                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr:                                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item:                                                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:                                                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id:                                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item:                                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr:                                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId:                                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item:                                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr:                                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtryOfRes:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn:                                                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Cd:                                                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_Rsn_Prtry:                                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_AddtlInf_Item:                                                                                 PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_StsRsnInf_AddtlInf:                                                                                      PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_Item:                                                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts:                                                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldNbOfTxs:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldSts:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_NbOfTxsPerSts_DtldCtrlSum:                                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_Item:                                                                                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts:                                                                                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsId:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlInstrId:                                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlEndToEndId:                                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_TxSts:                                                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Item:                                                                              PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf:                                                                                   PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr:                                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Nm:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrTp:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Dept:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_SubDept:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_StrtNm:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_BldgNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_PstCd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_TwnNm:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_CtrySubDvsn:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_Ctry:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine_Item:                                                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_PstlAdr_AdrLine:                                                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id:                                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_BICOrBEI:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Item:                                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr:                                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Id:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_OrgId_Othr_Issr:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Item:                                                         PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr:                                                              PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Id:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_Id_PrvtId_Othr_Issr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtryOfRes:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_NmPrfx:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Nm:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_PhneNb:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_MobNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_FaxNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_EmailAdr:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Orgtr_CtctDtls_Othr:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn:                                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Cd:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_Rsn_Prtry:                                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_AddtlInf_Item:                                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_StsRsnInf_AddtlInf:                                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Item:                                                                               PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf:                                                                                    PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Ccy:                                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Amt_Value:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty:                                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_BIC:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId:                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_ClrSysMmbId_MmbId:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Nm:                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr:                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrTp:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_Dept:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_SubDept:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_StrtNm:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_BldgNb:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_PstCd:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_TwnNm:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_CtrySubDvsn:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_Ctry:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrLine_Item:                                                PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_PstlAdr_AdrLine:                                                     PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm_Cd:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_SchmeNm_Prtry:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_FinInstnId_Othr_Issr:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId:                                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_Id:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_Nm:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrTp:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_Dept:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_SubDept:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_StrtNm:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_BldgNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_PstCd:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_TwnNm:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_CtrySubDvsn:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_Ctry:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrLine_Item:                                                   PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ChrgsInf_Pty_BrnchId_PstlAdr_AdrLine:                                                        PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AccptncDtTm:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_AcctSvcrRef:                                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_ClrSysRef:                                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef:                                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt_Ccy:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmAmt_Value:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt:                                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Ccy:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_InstdAmt_Value:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Ccy:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_Amt_Value:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Amt_EqvtAmt_CcyOfTrf:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_IntrBkSttlmDt:                                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdColltnDt:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_ReqdExctnDt:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Nm:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr:                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrTp:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_Dept:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_SubDept:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_StrtNm:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_BldgNb:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_PstCd:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_TwnNm:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_CtrySubDvsn:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_Ctry:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrLine_Item:                                                 PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_PstlAdr_AdrLine:                                                      PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId:                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_BICOrBEI:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Item:                                                   PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr:                                                        PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Id:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_OrgId_Othr_Issr:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId:                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth:                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Item:                                                  PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr:                                                       PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Id:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm:                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_Id_PrvtId_Othr_Issr:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtryOfRes:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls:                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_NmPrfx:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_Nm:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_PhneNb:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_MobNb:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_FaxNb:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_EmailAdr:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrSchmeId_CtctDtls_Othr:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf:                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmMtd:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id:                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_IBAN:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Id:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm:                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Cd:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_SchmeNm_Prtry:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Id_Othr_Issr:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp:                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Cd:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Tp_Prtry:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Ccy:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_SttlmAcct_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys_Cd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ClrSys_Prtry:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId:                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_BIC:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId:                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId:                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Nm:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Dept:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item:                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr:                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Id:                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm:                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_FinInstnId_Othr_Issr:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Id:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_Nm:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr:                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrTp:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Dept:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_SubDept:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_PstCd:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_Ctry:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item:                           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:                                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct:                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_IBAN:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr:                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Id:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Id_Othr_Issr:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Cd:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Tp_Prtry:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Ccy:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstgRmbrsmntAgtAcct_Nm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId:                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_BIC:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId:                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId:                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Nm:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Dept:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item:                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr:                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Id:                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm:                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_FinInstnId_Othr_Issr:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Id:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_Nm:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr:                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Dept:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_SubDept:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_PstCd:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_Ctry:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item:                           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:                                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct:                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_IBAN:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr:                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Id:                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Id_Othr_Issr:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Cd:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Tp_Prtry:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Ccy:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_InstdRmbrsmntAgtAcct_Nm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt:                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_BIC:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId:                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId:                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_ClrSysMmbId_MmbId:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Nm:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr:                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrTp:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Dept:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_SubDept:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_StrtNm:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_BldgNb:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_PstCd:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_TwnNm:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_CtrySubDvsn:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_Ctry:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine_Item:                         PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_PstlAdr_AdrLine:                              PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr:                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Id:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm:                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Cd:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_SchmeNm_Prtry:                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_FinInstnId_Othr_Issr:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Id:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_Nm:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr:                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrTp:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Dept:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_SubDept:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_StrtNm:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_BldgNb:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_PstCd:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_TwnNm:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_CtrySubDvsn:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_Ctry:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine_Item:                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgt_BrnchId_PstlAdr_AdrLine:                                 PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct:                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_IBAN:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr:                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Id:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Cd:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_SchmeNm_Prtry:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Id_Othr_Issr:                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Cd:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Tp_Prtry:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Ccy:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_SttlmInf_ThrdRmbrsmntAgtAcct_Nm:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf:                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_InstrPrty:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_ClrChanl:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Cd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SvcLvl_Prtry:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Cd:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_SeqTp:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Cd:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtTpInf_CtgyPurp_Prtry:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_PmtMtd:                                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_MndtId:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_DtOfSgntr:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInd:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlMndtId:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId:                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Nm:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr:                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId:                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr:                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId:                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr:                        PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtryOfRes:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls:                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Nm:                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_MobNb:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrSchmeId_CtctDtls_Othr:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt:                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId:                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_BIC:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId:                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_ClrSysMmbId_MmbId:              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Nm:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr:                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrTp:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Dept:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_SubDept:                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_StrtNm:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_BldgNb:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_PstCd:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_TwnNm:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_Ctry:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_PstlAdr_AdrLine:                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr:                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Id:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm:                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_FinInstnId_Othr_Issr:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId:                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Id:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_Nm:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr:                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrTp:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Dept:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_SubDept:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_StrtNm:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_BldgNb:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_PstCd:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_TwnNm:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_Ctry:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine_Item:              PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgt_BrnchId_PstlAdr_AdrLine:                   PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct:                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_IBAN:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr:                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Id:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Id_Othr_Issr:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Cd:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Tp_Prtry:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Ccy:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlCdtrAgtAcct_Nm:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Nm:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr:                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrTp:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Dept:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_SubDept:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_StrtNm:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_BldgNb:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_PstCd:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_TwnNm:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_CtrySubDvsn:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_Ctry:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine_Item:                         PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_PstlAdr_AdrLine:                              PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id:                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_BICOrBEI:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Item:                           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr:                                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Id:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm:                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Cd:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_OrgId_Othr_Issr:                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth:                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Item:                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr:                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Id:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_Id_PrvtId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtryOfRes:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_NmPrfx:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Nm:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_PhneNb:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_MobNb:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_FaxNb:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_EmailAdr:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtr_CtctDtls_Othr:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct:                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id:                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_IBAN:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr:                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Id:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm:                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Cd:                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_SchmeNm_Prtry:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Id_Othr_Issr:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp:                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Cd:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Tp_Prtry:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Ccy:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAcct_Nm:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt:                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId:                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_BIC:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId:                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_ClrSysMmbId_MmbId:              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Nm:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr:                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrTp:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Dept:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_SubDept:                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_StrtNm:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_BldgNb:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_PstCd:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_TwnNm:                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_Ctry:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_PstlAdr_AdrLine:                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr:                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Id:                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm:                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_FinInstnId_Othr_Issr:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId:                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Id:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_Nm:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr:                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrTp:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Dept:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_SubDept:                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_StrtNm:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_BldgNb:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_PstCd:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_TwnNm:                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_Ctry:                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine_Item:              PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgt_BrnchId_PstlAdr_AdrLine:                   PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct:                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_IBAN:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr:                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Id:                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Id_Othr_Issr:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp:                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Cd:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Tp_Prtry:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Ccy:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlDbtrAgtAcct_Nm:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFnlColltnDt:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_AmdmntInfDtls_OrgnlFrqcy:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_ElctrncSgntr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_FrstColltnDt:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_FnlColltnDt:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_MndtRltdInf_Frqcy:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf:                                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Ustrd_Item:                                                                PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Ustrd:                                                                     PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Item:                                                                 PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd:                                                                      PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Item:                                                      PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf:                                                           PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Tp_Issr:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_Nb:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocInf_RltdDt:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt:                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value:                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt:                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf:                                                           PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Tp_Issr:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_CdtrRefInf_Ref:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Nm:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr:                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Dept:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_SubDept:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_PstCd:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_Ctry:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                                           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                                                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id:                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                                             PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr:                                                  PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                                                 PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtryOfRes:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Nm:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_MobNb:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcr_CtctDtls_Othr:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Dept:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_SubDept:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_PstCd:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_Ctry:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id:                                                            PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                                            PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr:                                                 PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId:                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                                           PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                                                PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtryOfRes:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Nm:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_MobNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_Invcee_CtctDtls_Othr:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf_Item:                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_RmtInf_Strd_AddtlRmtInf:                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr:                                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Nm:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrTp:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Dept:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_SubDept:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_StrtNm:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_BldgNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_PstCd:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_TwnNm:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_CtrySubDvsn:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_Ctry:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrLine_Item:                                                   PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_PstlAdr_AdrLine:                                                        PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_BICOrBEI:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Item:                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr:                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Id:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_OrgId_Othr_Issr:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId:                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Item:                                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr:                                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Id:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_Id_PrvtId_Othr_Issr:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtryOfRes:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_NmPrfx:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_PhneNb:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_MobNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_FaxNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_EmailAdr:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtDbtr_CtctDtls_Othr:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr:                                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Nm:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrTp:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Dept:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_SubDept:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_StrtNm:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_BldgNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_PstCd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_TwnNm:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_CtrySubDvsn:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_Ctry:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrLine_Item:                                                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_PstlAdr_AdrLine:                                                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id:                                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_BICOrBEI:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Item:                                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr:                                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Id:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_OrgId_Othr_Issr:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Item:                                                         PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr:                                                              PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Id:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_Id_PrvtId_Othr_Issr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtryOfRes:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_NmPrfx:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Nm:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_PhneNb:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_MobNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_FaxNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_EmailAdr:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Dbtr_CtctDtls_Othr:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct:                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_IBAN:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr:                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Id:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm:                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Cd:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_SchmeNm_Prtry:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Id_Othr_Issr:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Cd:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Tp_Prtry:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Ccy:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAcct_Nm:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt:                                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_BIC:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Dept:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr:                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Id:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_FinInstnId_Othr_Issr:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Id:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_Nm:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr:                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_IBAN:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr:                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Id:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Id_Othr_Issr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp_Cd:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Tp_Prtry:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Ccy:                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_DbtrAgtAcct_Nm:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt:                                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_BIC:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Dept:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_SubDept:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_PstCd:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_Ctry:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr:                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Id:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_FinInstnId_Othr_Issr:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId:                                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Id:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_Nm:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr:                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrTp:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgt_BrnchId_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_IBAN:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr:                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Id:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Id_Othr_Issr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp_Cd:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Tp_Prtry:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Ccy:                                                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAgtAcct_Nm:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr:                                                                             PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Nm:                                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrTp:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Dept:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_SubDept:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_StrtNm:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_BldgNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_PstCd:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_TwnNm:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_CtrySubDvsn:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_Ctry:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrLine_Item:                                                        PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_PstlAdr_AdrLine:                                                             PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id:                                                                          PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_BICOrBEI:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Item:                                                          PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr:                                                               PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Id:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm:                                                       PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_OrgId_Othr_Issr:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId:                                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                                                   PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Item:                                                         PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr:                                                              PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Id:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_Id_PrvtId_Othr_Issr:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtryOfRes:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls:                                                                    PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_NmPrfx:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Nm:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_PhneNb:                                                             PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_MobNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_FaxNb:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_EmailAdr:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_Cdtr_CtctDtls_Othr:                                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct:                                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_IBAN:                                                                 PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr:                                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Id:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm:                                                         PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Cd:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_SchmeNm_Prtry:                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Id_Othr_Issr:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp:                                                                      PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Cd:                                                                   PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Tp_Prtry:                                                                PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Ccy:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_CdtrAcct_Nm:                                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr:                                                                        PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Nm:                                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr:                                                                PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrTp:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Dept:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_SubDept:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_StrtNm:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_BldgNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_PstCd:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_TwnNm:                                                          PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_CtrySubDvsn:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_Ctry:                                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrLine_Item:                                                   PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_PstlAdr_AdrLine:                                                        PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id:                                                                     PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_BICOrBEI:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Item:                                                     PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr:                                                          PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Id:                                                       PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                                                  PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                                               PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_OrgId_Othr_Issr:                                                     PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId:                                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                                              PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                  PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Item:                                                    PathTypeArrayItem,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr:                                                         PathTypeArray,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Id:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                                                 PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                           PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_Id_PrvtId_Othr_Issr:                                                    PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtryOfRes:                                                              PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls:                                                               PathTypeStruct,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_NmPrfx:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Nm:                                                            PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_PhneNb:                                                        PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_MobNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_FaxNb:                                                         PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_EmailAdr:                                                      PathTypeProperty,
	Path_CstmrPmtStsRpt_OrgnlPmtInfAndSts_TxInfAndSts_OrgnlTxRef_UltmtCdtr_CtctDtls_Othr:                                                          PathTypeProperty,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in pain_002_001_03", p)
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
