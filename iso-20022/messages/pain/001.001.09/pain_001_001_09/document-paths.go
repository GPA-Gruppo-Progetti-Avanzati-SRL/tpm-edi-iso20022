// Package pain_001_001_09
// Do not Edit. This stuff it's been automatically generated.
package pain_001_001_09

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_CstmrCdtTrfInitn                                                                                               = "CstmrCdtTrfInitn"
	Path_CstmrCdtTrfInitn_GrpHdr                                                                                        = "CstmrCdtTrfInitn.GrpHdr"
	Path_CstmrCdtTrfInitn_GrpHdr_MsgId                                                                                  = "CstmrCdtTrfInitn.GrpHdr.MsgId"
	Path_CstmrCdtTrfInitn_GrpHdr_CreDtTm                                                                                = "CstmrCdtTrfInitn.GrpHdr.CreDtTm"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Item                                                                           = "CstmrCdtTrfInitn.GrpHdr.Authstn[]"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn                                                                                = "CstmrCdtTrfInitn.GrpHdr.Authstn" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Cd                                                                             = "CstmrCdtTrfInitn.GrpHdr.Authstn[].Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Prtry                                                                          = "CstmrCdtTrfInitn.GrpHdr.Authstn[].Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_NbOfTxs                                                                                = "CstmrCdtTrfInitn.GrpHdr.NbOfTxs"
	Path_CstmrCdtTrfInitn_GrpHdr_CtrlSum                                                                                = "CstmrCdtTrfInitn.GrpHdr.CtrlSum"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty                                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Nm                                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr                                                                       = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd                                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry                                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id                                                        = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr                                                      = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Dept                                                                  = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_SubDept                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_StrtNm                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNb                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNm                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Flr                                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstBx                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Room                                                                  = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstCd                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnNm                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnLctnNm                                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_DstrctNm                                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn                                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Ctry                                                                  = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item                                                          = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id                                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId                                                                      = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_AnyBIC                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_LEI                                                                  = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id                                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm                                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd                                                      = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId                                                                     = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth                                                     = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item                                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id                                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm                                                        = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd                                                     = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry                                                  = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr                                                           = "CstmrCdtTrfInitn.GrpHdr.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtryOfRes                                                                     = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtryOfRes"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls                                                                      = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Nm                                                                   = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PhneNb                                                               = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_MobNb                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_FaxNb                                                                = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr                                                             = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailPurp                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_JobTitl                                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Rspnsblty                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Dept                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_Item                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr                                                                 = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp                                                         = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_Id                                                              = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PrefrdMtd                                                            = "CstmrCdtTrfInitn.GrpHdr.InitgPty.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt                                                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId                                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_BICFI                                                               = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId                                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                             = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                          = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_LEI                                                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Nm                                                                  = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr                                                             = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Cd                                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                              = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                            = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm                                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb                                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNm                                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Flr                                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstBx                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Room                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnLctnNm                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_DstrctNm                                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr                                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id                                                             = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry                                                  = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr                                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId                                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Id                                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_LEI                                                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Nm                                                                     = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr                                                                = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp                                                          = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Cd                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry                                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                                 = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                               = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                            = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept                                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm                                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb                                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNm                                                         = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Flr                                                            = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstBx                                                          = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Room                                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd                                                          = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm                                                          = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnLctnNm                                                      = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_DstrctNm                                                       = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn                                                    = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry                                                           = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item                                                   = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine                                                        = "CstmrCdtTrfInitn.GrpHdr.FwdgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Item                                                                                   = "CstmrCdtTrfInitn.PmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf                                                                                        = "CstmrCdtTrfInitn.PmtInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_PmtInfId                                                                               = "CstmrCdtTrfInitn.PmtInf[].PmtInfId"
	Path_CstmrCdtTrfInitn_PmtInf_PmtMtd                                                                                 = "CstmrCdtTrfInitn.PmtInf[].PmtMtd"
	Path_CstmrCdtTrfInitn_PmtInf_BtchBookg                                                                              = "CstmrCdtTrfInitn.PmtInf[].BtchBookg"
	Path_CstmrCdtTrfInitn_PmtInf_NbOfTxs                                                                                = "CstmrCdtTrfInitn.PmtInf[].NbOfTxs"
	Path_CstmrCdtTrfInitn_PmtInf_CtrlSum                                                                                = "CstmrCdtTrfInitn.PmtInf[].CtrlSum"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf                                                                               = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_InstrPrty                                                                     = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Item                                                                   = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl[]"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl                                                                        = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Cd                                                                     = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Prtry                                                                  = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.SvcLvl[].Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm                                                                     = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Cd                                                                  = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Prtry                                                               = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp                                                                      = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Cd                                                                   = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry                                                                = "CstmrCdtTrfInitn.PmtInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt                                                                            = "CstmrCdtTrfInitn.PmtInf[].ReqdExctnDt"
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt_Dt                                                                         = "CstmrCdtTrfInitn.PmtInf[].ReqdExctnDt.Dt"
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt_DtTm                                                                       = "CstmrCdtTrfInitn.PmtInf[].ReqdExctnDt.DtTm"
	Path_CstmrCdtTrfInitn_PmtInf_PoolgAdjstmntDt                                                                        = "CstmrCdtTrfInitn.PmtInf[].PoolgAdjstmntDt"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr                                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Nm                                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr                                                                           = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Cd                                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry                                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Id                                                            = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Issr                                                          = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Dept                                                                      = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_SubDept                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_StrtNm                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNb                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNm                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Flr                                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstBx                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Room                                                                      = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstCd                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnNm                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnLctnNm                                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_DstrctNm                                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_CtrySubDvsn                                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Ctry                                                                      = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine_Item                                                              = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id                                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId                                                                          = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_AnyBIC                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_LEI                                                                      = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Item                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Id                                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm                                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                                                          = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Issr                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId                                                                         = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth                                                         = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Item                                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Id                                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm                                                            = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                         = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                      = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Issr                                                               = "CstmrCdtTrfInitn.PmtInf[].Dbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtryOfRes                                                                         = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls                                                                          = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_NmPrfx                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Nm                                                                       = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PhneNb                                                                   = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_MobNb                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_FaxNb                                                                    = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailAdr                                                                 = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailPurp                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_JobTitl                                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Rspnsblty                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Dept                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_Item                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr                                                                     = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_ChanlTp                                                             = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_Id                                                                  = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PrefrdMtd                                                                = "CstmrCdtTrfInitn.PmtInf[].Dbtr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct                                                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id                                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_IBAN                                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr                                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Id                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm                                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Issr                                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp                                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Cd                                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Prtry                                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Ccy                                                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Nm                                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy                                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp                                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp_Cd                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp_Prtry                                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Id                                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAcct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt                                                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId                                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_BICFI                                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_LEI                                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Nm                                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr                                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNm                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Flr                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstBx                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Room                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_DstrctNm                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr                                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Id                                                             = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                                  = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Issr                                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId                                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Id                                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_LEI                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Nm                                                                     = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr                                                                = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept                                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNm                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Flr                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstBx                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Room                                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm                                                          = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_DstrctNm                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry                                                           = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct                                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id                                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_IBAN                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Id                                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm                                                            = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Issr                                                               = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp                                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Cd                                                                      = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Prtry                                                                   = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Ccy                                                                        = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Nm                                                                         = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy                                                                       = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp_Cd                                                                 = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp_Prtry                                                              = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Id                                                                    = "CstmrCdtTrfInitn.PmtInf[].DbtrAgtAcct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_InstrForDbtrAgt                                                                        = "CstmrCdtTrfInitn.PmtInf[].InstrForDbtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr                                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Nm                                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr                                                                      = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Cd                                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry                                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id                                                       = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr                                                     = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Dept                                                                 = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_SubDept                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_StrtNm                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNb                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNm                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Flr                                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstBx                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Room                                                                 = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstCd                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnNm                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnLctnNm                                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_DstrctNm                                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn                                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Ctry                                                                 = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item                                                         = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id                                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId                                                                     = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_AnyBIC                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_LEI                                                                 = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id                                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                                     = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId                                                                    = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                                    = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item                                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id                                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                                       = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                    = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                 = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr                                                          = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtryOfRes                                                                    = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls                                                                     = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_NmPrfx                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Nm                                                                  = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PhneNb                                                              = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_MobNb                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_FaxNb                                                               = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailAdr                                                            = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailPurp                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_JobTitl                                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Rspnsblty                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Dept                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_Item                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr                                                                = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_ChanlTp                                                        = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_Id                                                             = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PrefrdMtd                                                           = "CstmrCdtTrfInitn.PmtInf[].UltmtDbtr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgBr                                                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgBr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct                                                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id                                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_IBAN                                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr                                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Id                                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm                                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry                                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Issr                                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp                                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Cd                                                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Prtry                                                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Ccy                                                                          = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Nm                                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy                                                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp                                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp_Cd                                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp_Prtry                                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Id                                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt                                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId                                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BICFI                                                          = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_LEI                                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm                                                             = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr                                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Cd                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                         = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                       = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNm                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Flr                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstBx                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Room                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnLctnNm                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_DstrctNm                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id                                                        = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry                                             = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId                                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id                                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_LEI                                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm                                                                = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr                                                           = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp                                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Cd                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                            = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                          = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                       = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNm                                                    = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Flr                                                       = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstBx                                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Room                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd                                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm                                                     = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnLctnNm                                                 = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_DstrctNm                                                  = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn                                               = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry                                                      = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item                                              = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine                                                   = "CstmrCdtTrfInitn.PmtInf[].ChrgsAcctAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Item                                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf                                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId                                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_InstrId                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId.InstrId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_EndToEndId                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId.EndToEndId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_UETR                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtId.UETR"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_InstrPrty                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Item                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Cd                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Prtry                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.SvcLvl[].Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Cd                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt                                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Ccy                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Value                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.InstdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Ccy                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Value                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_CcyOfTrf                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Amt.EqvtAmt.CcyOfTrf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_UnitCcy                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.UnitCcy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_XchgRate                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.XchgRate"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_RateTp                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.RateTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_CtrctId                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].XchgRateInf.CtrctId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChrgBr                                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChrgBr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqTp                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqNb                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Nm                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Cd                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Id                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Issr                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_SchmeNm                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Dept                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_SubDept                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_StrtNm                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNb                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNm                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Flr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstBx                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Room                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstCd                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnLctnNm                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_DstrctNm                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_CtrySubDvsn                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Ctry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine_Item                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqFr.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvryMtd.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Nm                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Cd                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Id                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Issr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_SchmeNm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Dept                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_SubDept                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_StrtNm                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNb                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNm                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Flr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstBx                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Room                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstCd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnNm                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnLctnNm                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_DstrctNm                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_CtrySubDvsn                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Ctry                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine_Item                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.DlvrTo.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_InstrPrty                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.InstrPrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqMtrtyDt                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.ChqMtrtyDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_FrmsCd                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.FrmsCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld_Item                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.MemoFld[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.MemoFld" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_RgnlClrZone                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.RgnlClrZone"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_PrtLctn                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.PrtLctn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_Sgntr_Item                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.Sgntr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_Sgntr                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].ChqInstr.Sgntr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr                                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Nm                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Cd                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Dept                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_SubDept                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_StrtNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Flr                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstBx                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Room                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstCd                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnNm                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnLctnNm                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_DstrctNm                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Ctry                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine_Item                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_AnyBIC                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Issr                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Item                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Id                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtryOfRes                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_NmPrfx                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PhneNb                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_MobNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_FaxNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailAdr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailPurp                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_JobTitl                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Rspnsblty                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Dept                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_ChanlTp                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PrefrdMtd                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtDbtr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_BICFI                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_LEI                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Nm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Dept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_SubDept                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Flr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstBx                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Room                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstCd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnLctnNm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_DstrctNm                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Ctry                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Id                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Issr                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_SubDept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_StrtNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Flr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstBx                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Room                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstCd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnLctnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_DstrctNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Ctry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_IBAN                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Id                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Cd                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Prtry                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Issr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Ccy                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Nm                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp_Cd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp_Prtry                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Id                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt1Acct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_BICFI                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_LEI                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Nm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Dept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_SubDept                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Flr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstBx                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Room                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstCd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnLctnNm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_DstrctNm                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Ctry                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Id                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Issr                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_SubDept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_StrtNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Flr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstBx                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Room                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstCd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnLctnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_DstrctNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Ctry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_IBAN                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Id                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Cd                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Prtry                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Issr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Ccy                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Nm                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp_Cd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp_Prtry                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Id                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt2Acct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_BICFI                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_LEI                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Nm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Dept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_SubDept                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Flr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstBx                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Room                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstCd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnLctnNm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_DstrctNm                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Ctry                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Id                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Issr                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_SubDept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_StrtNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Flr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstBx                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Room                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstCd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnLctnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_DstrctNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Ctry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_IBAN                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Id                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Cd                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Prtry                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Issr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Cd                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Prtry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Ccy                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Nm                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp_Cd                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp_Prtry                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Id                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].IntrmyAgt3Acct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt                                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_BICFI                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.BICFI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_SubDept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Flr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstBx                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Room                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstCd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_DstrctNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Ctry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Issr                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.FinInstnId.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Id                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_LEI                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Nm                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Dept                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_SubDept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_StrtNm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNb                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Flr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstBx                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Room                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstCd                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnNm                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_DstrctNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Ctry                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_IBAN                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Id                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Issr                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Cd                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Prtry                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Ccy                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Nm                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp_Cd                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp_Prtry                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Id                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAgtAcct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr                                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Nm                                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Cd                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_Id                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_Issr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Dept                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_SubDept                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_StrtNm                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNb                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNm                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Flr                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstBx                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Room                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstCd                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnNm                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnLctnNm                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_DstrctNm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_CtrySubDvsn                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Ctry                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine_Item                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id                                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_AnyBIC                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_LEI                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Item                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Issr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Item                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Id                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Issr                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtryOfRes                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_NmPrfx                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Nm                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PhneNb                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_MobNb                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_FaxNb                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailAdr                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailPurp                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_JobTitl                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Rspnsblty                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Dept                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_Item                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_ChanlTp                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PrefrdMtd                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Cdtr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_IBAN                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.IBAN"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Id                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Cd                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Prtry                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Issr                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Id.Othr.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Cd                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Prtry                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Ccy                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Nm                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Prxy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Prxy.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp_Cd                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Prxy.Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp_Prtry                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Prxy.Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Id                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].CdtrAcct.Prxy.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr                                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Nm                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Cd                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Dept                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_SubDept                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_StrtNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNm                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Flr                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstBx                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Room                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstCd                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnNm                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnLctnNm                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_DstrctNm                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Ctry                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine_Item                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_AnyBIC                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_LEI                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Issr                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Item                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Id                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtryOfRes                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_NmPrfx                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PhneNb                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_MobNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_FaxNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailAdr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailPurp                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_JobTitl                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Rspnsblty                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Dept                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_ChanlTp                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_Id                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PrefrdMtd                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].UltmtCdtr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Item                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Cd                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_InstrInf                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForCdtrAgt[].InstrInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForDbtrAgt                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].InstrForDbtrAgt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp                                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Cd                                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Prtry                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Purp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Item                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_DbtCdtRptgInd                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].DbtCdtRptgInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Ctry                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Authrty.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Item                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Tp                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Dt                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Ctry                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Cd                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Ccy                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Value                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf_Item                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Inf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RgltryRptg[].Dtls[].Inf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax                                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxId                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_RegnId                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxTp                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Cdtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxId                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_RegnId                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxTp                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Titl                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn.Titl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Nm                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dbtr.Authstn.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_AdmstnZone                                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.AdmstnZone"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_RefNb                                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.RefNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Mtd                                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Mtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Ccy                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Value                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Ccy                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Value                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.TtlTaxAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dt                                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_SeqNb                                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.SeqNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Item                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd                                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Tp                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Ctgy                                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Ctgy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CtgyDtls                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].CtgyDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_DbtrSts                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].DbtrSts"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CertId                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].CertId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_FrmsCd                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].FrmsCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Yr                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Tp                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt                                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_FrDt                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_ToDt                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Rate                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Rate"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Item                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_AddtlInf                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].Tax.Rcrd[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_Item                                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf                                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtId                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_Item                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_Mtd                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].Mtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_ElctrncAdr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].ElctrncAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Nm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Cd                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Id                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Issr                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_SchmeNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Dept                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_SubDept                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_StrtNm                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNb                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNm                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Flr                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstBx                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Room                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstCd                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnNm                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnLctnNm                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_DstrctNm                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_CtrySubDvsn                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Ctry                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine_Item                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf                                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd_Item                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Ustrd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd                                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Ustrd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Item                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Item                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Nb                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_RltdDt                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Item                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Nb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].RltdDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Desc                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Desc"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Item                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Item                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Item                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Rsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Item                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Item                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Ref                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].CdtrRefInf.Ref"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Nm                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Flr                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstBx                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Room                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_DstrctNm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_LEI                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtryOfRes                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailPurp                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_JobTitl                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_Item                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_Id                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Nm                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Flr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstBx                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Room                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_DstrctNm                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_LEI                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtryOfRes                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailPurp                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_JobTitl                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Dept                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_Item                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_Id                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].Invcee.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt                                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Cdtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_TaxId                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Cdtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_RegnId                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Cdtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_TaxTp                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Cdtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_TaxId                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_RegnId                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_TaxTp                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn.Titl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.TaxId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.RegnId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.TaxTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn.Titl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_AdmstnZone                                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.AdmstnZone"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_RefNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.RefNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Mtd                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Mtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.TtlTaxAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dt                                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_SeqNb                                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.SeqNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Tp                                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Ctgy                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Ctgy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].CtgyDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts                                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].DbtrSts"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_CertId                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].CertId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].FrmsCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Rate"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt                                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Item                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].TaxRmt.Rcrd[].AddtlInf"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Tp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry                                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_Issr                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Tp.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine_Item                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Item                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Item                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm                                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Item                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id                            = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.SubDept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.StrtNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.BldgNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.BldgNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Flr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.PstBx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Room"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.PstCd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.TwnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.TwnLctnNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.DstrctNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.CtrySubDvsn"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Ctry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine_Item                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrLine[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrLine" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.AnyBIC"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.LEI"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Item                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Item                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].Issr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtryOfRes"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls                              = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.NmPrfx"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Nm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.PhneNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.MobNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb                        = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.FaxNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.EmailAdr"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.EmailPurp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.JobTitl"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Rspnsblty"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Dept"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Item                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp                 = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[].ChanlTp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id                      = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[].Id"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.PrefrdMtd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RefNb                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.RefNb"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Dt                                                  = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.Dt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt                                             = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy                                         = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt.Ccy"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt.Value"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd                                   = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.FmlyMdclInsrncInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].GrnshmtRmt.MplyeeTermntnInd"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf_Item                                               = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].AddtlRmtInf[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf                                                    = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].RmtInf.Strd[].AddtlRmtInf" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Item                                                           = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].SplmtryData[]"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData                                                                = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].SplmtryData" // ARRAY
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_PlcAndNm                                                       = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].SplmtryData[].PlcAndNm"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Envlp                                                          = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].SplmtryData[].Envlp"
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Envlp_Item                                                     = "CstmrCdtTrfInitn.PmtInf[].CdtTrfTxInf[].SplmtryData[].Envlp.Item"
	Path_CstmrCdtTrfInitn_SplmtryData_Item                                                                              = "CstmrCdtTrfInitn.SplmtryData[]"
	Path_CstmrCdtTrfInitn_SplmtryData                                                                                   = "CstmrCdtTrfInitn.SplmtryData" // ARRAY
	Path_CstmrCdtTrfInitn_SplmtryData_PlcAndNm                                                                          = "CstmrCdtTrfInitn.SplmtryData[].PlcAndNm"
	Path_CstmrCdtTrfInitn_SplmtryData_Envlp                                                                             = "CstmrCdtTrfInitn.SplmtryData[].Envlp"
	Path_CstmrCdtTrfInitn_SplmtryData_Envlp_Item                                                                        = "CstmrCdtTrfInitn.SplmtryData[].Envlp.Item"
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_CstmrCdtTrfInitn:                                                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr:                                                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_MsgId:                                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_CreDtTm:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Item:                                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn:                                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Cd:                                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_Authstn_Prtry:                                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_NbOfTxs:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_CtrlSum:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty:                                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Nm:                                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp:                                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Dept:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_SubDept:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_StrtNm:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNb:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_BldgNm:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Flr:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstBx:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Room:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_PstCd:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnNm:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_TwnLctnNm:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_DstrctNm:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_Ctry:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine_Item:                                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_PstlAdr_AdrLine:                                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id:                                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_AnyBIC:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_LEI:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Item:                                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr:                                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Id:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Item:                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr:                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtryOfRes:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_NmPrfx:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Nm:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PhneNb:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_MobNb:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_FaxNb:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailAdr:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_EmailPurp:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_JobTitl:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Rspnsblty:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Dept:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_Item:                                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr:                                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_Othr_Id:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_InitgPty_CtctDtls_PrefrdMtd:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt:                                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_BICFI:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_ClrSysMmbId_MmbId:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_LEI:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Nm:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Dept:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_SubDept:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_StrtNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNb:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_BldgNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Flr:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstBx:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Room:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_PstCd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_TwnLctnNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_DstrctNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_Ctry:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_PstlAdr_AdrLine:                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Cd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_FinInstnId_Othr_Issr:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId:                                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Id:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_LEI:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_Nm:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Dept:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_SubDept:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_StrtNm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNb:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_BldgNm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Flr:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstBx:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Room:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_PstCd:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnNm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_TwnLctnNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_DstrctNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_Ctry:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine_Item:                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_GrpHdr_FwdgAgt_BrnchId_PstlAdr_AdrLine:                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Item:                                                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf:                                                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_PmtInfId:                                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtMtd:                                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_BtchBookg:                                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_NbOfTxs:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CtrlSum:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf:                                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_InstrPrty:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Item:                                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl:                                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Cd:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_SvcLvl_Prtry:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Cd:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_LclInstrm_Prtry:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Cd:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PmtTpInf_CtgyPurp_Prtry:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt:                                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt_Dt:                                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ReqdExctnDt_DtTm:                                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_PoolgAdjstmntDt:                                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr:                                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Nm:                                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr:                                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Cd:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Id:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Issr:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Dept:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_SubDept:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_StrtNm:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNb:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_BldgNm:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Flr:                                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstBx:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Room:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_PstCd:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnNm:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_TwnLctnNm:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_DstrctNm:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_CtrySubDvsn:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_Ctry:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine_Item:                                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_PstlAdr_AdrLine:                                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id:                                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId:                                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_AnyBIC:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_LEI:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Item:                                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr:                                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Id:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_OrgId_Othr_Issr:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId:                                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Item:                                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr:                                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Id:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_Id_PrvtId_Othr_Issr:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtryOfRes:                                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls:                                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_NmPrfx:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Nm:                                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PhneNb:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_MobNb:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_FaxNb:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailAdr:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_EmailPurp:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_JobTitl:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Rspnsblty:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Dept:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_Item:                                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr:                                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_ChanlTp:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_Othr_Id:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_Dbtr_CtctDtls_PrefrdMtd:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct:                                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id:                                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_IBAN:                                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Id:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Id_Othr_Issr:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp:                                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Cd:                                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Tp_Prtry:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Ccy:                                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Nm:                                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy:                                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp_Cd:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Tp_Prtry:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAcct_Prxy_Id:                                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt:                                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_BICFI:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_LEI:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Nm:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Flr:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstBx:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Room:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_FinInstnId_Othr_Issr:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId:                                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Id:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_LEI:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_Nm:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Flr:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstBx:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Room:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_DstrctNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct:                                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id:                                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_IBAN:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr:                                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Id:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Id_Othr_Issr:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp:                                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Cd:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Tp_Prtry:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Ccy:                                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Nm:                                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp:                                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp_Cd:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Tp_Prtry:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_DbtrAgtAcct_Prxy_Id:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_InstrForDbtrAgt:                                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr:                                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Nm:                                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Cd:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Dept:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_SubDept:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_StrtNm:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNb:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_BldgNm:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Flr:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstBx:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Room:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_PstCd:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnNm:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_TwnLctnNm:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_DstrctNm:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_Ctry:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item:                                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_PstlAdr_AdrLine:                                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id:                                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_AnyBIC:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_LEI:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item:                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr:                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId:                                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item:                                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr:                                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtryOfRes:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_NmPrfx:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Nm:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PhneNb:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_MobNb:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_FaxNb:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailAdr:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_EmailPurp:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_JobTitl:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Rspnsblty:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Dept:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_Item:                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr:                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_ChanlTp:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_Othr_Id:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_UltmtDbtr_CtctDtls_PrefrdMtd:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgBr:                                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct:                                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id:                                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_IBAN:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Id:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Cd:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_SchmeNm_Prtry:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Id_Othr_Issr:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp:                                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Cd:                                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Tp_Prtry:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Ccy:                                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Nm:                                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy:                                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp_Cd:                                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Tp_Prtry:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcct_Prxy_Id:                                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt:                                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_BICFI:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_ClrSysMmbId_MmbId:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_LEI:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Nm:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Dept:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_SubDept:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_StrtNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNb:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_BldgNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Flr:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstBx:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Room:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_PstCd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_TwnLctnNm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_DstrctNm:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_Ctry:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine_Item:                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_PstlAdr_AdrLine:                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Id:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Cd:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_SchmeNm_Prtry:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_FinInstnId_Othr_Issr:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Id:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_LEI:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_Nm:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Cd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Dept:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_SubDept:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_StrtNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNb:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_BldgNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Flr:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstBx:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Room:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_PstCd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnNm:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_TwnLctnNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_DstrctNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_CtrySubDvsn:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_Ctry:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine_Item:                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_ChrgsAcctAgt_BrnchId_PstlAdr_AdrLine:                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Item:                                                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf:                                                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId:                                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_InstrId:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_EndToEndId:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtId_UETR:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_InstrPrty:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Item:                                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl:                                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Cd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_SvcLvl_Prtry:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Cd:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_PmtTpInf_CtgyPurp_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt:                                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Ccy:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_InstdAmt_Value:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Ccy:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_Amt_Value:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Amt_EqvtAmt_CcyOfTrf:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_UnitCcy:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_XchgRate:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_RateTp:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_XchgRateInf_CtrctId:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChrgBr:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqTp:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqNb:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Cd:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Id:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Issr:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrTp_Prtry_SchmeNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Dept:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_SubDept:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_StrtNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_BldgNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Flr:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstBx:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Room:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_PstCd:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_TwnLctnNm:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_DstrctNm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_CtrySubDvsn:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_Ctry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine_Item:                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqFr_Adr_AdrLine:                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvryMtd_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Nm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Cd:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Id:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Issr:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_SchmeNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Dept:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_SubDept:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_StrtNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNb:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_BldgNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Flr:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstBx:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Room:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_PstCd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnNm:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_TwnLctnNm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_DstrctNm:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_CtrySubDvsn:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_Ctry:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine_Item:                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_DlvrTo_Adr_AdrLine:                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_InstrPrty:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_ChqMtrtyDt:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_FrmsCd:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld_Item:                                                      PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_MemoFld:                                                           PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_RgnlClrZone:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_PrtLctn:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_Sgntr_Item:                                                        PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_ChqInstr_Sgntr:                                                             PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr:                                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Nm:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Cd:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_BldgNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Flr:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstBx:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Room:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_TwnLctnNm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_DstrctNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_AnyBIC:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_OrgId_Othr_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Item:                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr:                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Id:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtryOfRes:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_NmPrfx:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PhneNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_MobNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_FaxNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailAdr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_EmailPurp:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_JobTitl:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Rspnsblty:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Dept:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_ChanlTp:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtDbtr_CtctDtls_PrefrdMtd:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1:                                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Id:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_IBAN:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Id:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Cd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_SchmeNm_Prtry:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Id_Othr_Issr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Tp_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Ccy:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp_Cd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Tp_Prtry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt1Acct_Prxy_Id:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2:                                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Id:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_IBAN:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Id:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Cd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_SchmeNm_Prtry:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Id_Othr_Issr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Tp_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Ccy:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp_Cd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Tp_Prtry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt2Acct_Prxy_Id:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3:                                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Id:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_IBAN:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Id:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Cd:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_SchmeNm_Prtry:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Id_Othr_Issr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Cd:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Tp_Prtry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Ccy:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Nm:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp_Cd:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Tp_Prtry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_IntrmyAgt3Acct_Prxy_Id:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt:                                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_BICFI:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_FinInstnId_Othr_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Id:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_LEI:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_Nm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Dept:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_SubDept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_StrtNm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNb:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_BldgNm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Flr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstBx:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Room:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_PstCd:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnNm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_DstrctNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_Ctry:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgt_BrnchId_PstlAdr_AdrLine:                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_IBAN:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Id:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Id_Othr_Issr:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Cd:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Tp_Prtry:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Ccy:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Nm:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp_Cd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Tp_Prtry:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAgtAcct_Prxy_Id:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Nm:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Cd:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_Id:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_Issr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Dept:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_SubDept:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_StrtNm:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNb:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_BldgNm:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Flr:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstBx:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Room:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_PstCd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnNm:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_TwnLctnNm:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_DstrctNm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_CtrySubDvsn:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_Ctry:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine_Item:                                                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_PstlAdr_AdrLine:                                                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id:                                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_AnyBIC:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_LEI:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Item:                                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr:                                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Id:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_OrgId_Othr_Issr:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId:                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Item:                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr:                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Id:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_Id_PrvtId_Othr_Issr:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtryOfRes:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_NmPrfx:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Nm:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PhneNb:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_MobNb:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_FaxNb:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailAdr:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_EmailPurp:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_JobTitl:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Rspnsblty:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Dept:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_Item:                                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr:                                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_ChanlTp:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_Othr_Id:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Cdtr_CtctDtls_PrefrdMtd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_IBAN:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Id:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Cd:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_SchmeNm_Prtry:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Id_Othr_Issr:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp:                                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Cd:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Tp_Prtry:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Ccy:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Nm:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp_Cd:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Tp_Prtry:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_CdtrAcct_Prxy_Id:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr:                                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Nm:                                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Cd:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry:                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_BldgNm:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Flr:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstBx:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Room:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_TwnLctnNm:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_DstrctNm:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_AnyBIC:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_LEI:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_OrgId_Othr_Issr:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Item:                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr:                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Id:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_Id_PrvtId_Othr_Issr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtryOfRes:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_NmPrfx:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PhneNb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_MobNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_FaxNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailAdr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_EmailPurp:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_JobTitl:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Rspnsblty:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Dept:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_ChanlTp:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_Othr_Id:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_UltmtCdtr_CtctDtls_PrefrdMtd:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Item:                                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt:                                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_Cd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForCdtrAgt_InstrInf:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_InstrForDbtrAgt:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp:                                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Cd:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Purp_Prtry:                                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Item:                                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg:                                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_DbtCdtRptgInd:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Authrty_Ctry:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Item:                                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls:                                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Tp:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Dt:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Ctry:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Cd:                                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Ccy:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Amt_Value:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf_Item:                                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RgltryRptg_Dtls_Inf:                                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax:                                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxId:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_RegnId:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Cdtr_TaxTp:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr:                                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxId:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_RegnId:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_TaxTp:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn:                                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Titl:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dbtr_Authstn_Nm:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_AdmstnZone:                                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_RefNb:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Mtd:                                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Ccy:                                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxblBaseAmt_Value:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt:                                                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Ccy:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_TtlTaxAmt_Value:                                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Dt:                                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_SeqNb:                                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Item:                                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd:                                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Tp:                                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Ctgy:                                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CtgyDtls:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_DbtrSts:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_CertId:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_FrmsCd:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd:                                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Yr:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_Tp:                                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt:                                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_FrDt:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_Prd_FrToDt_ToDt:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt:                                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Rate:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_TtlAmt_Value:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Item:                                                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls:                                                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt:                                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_Tax_Rcrd_AddtlInf:                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_Item:                                                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf:                                                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtId:                                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_Item:                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls:                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_Mtd:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_ElctrncAdr:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Nm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Cd:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Id:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Issr:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_SchmeNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Dept:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_SubDept:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_StrtNm:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNb:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNm:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Flr:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstBx:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Room:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstCd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnNm:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnLctnNm:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_DstrctNm:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_CtrySubDvsn:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Ctry:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine_Item:                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine:                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf:                                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd_Item:                                                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Ustrd:                                                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Item:                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd:                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Item:                                                PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf:                                                     PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Tp_Issr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_Nb:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_RltdDt:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Item:                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls:                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry:                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Desc:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Item:                      PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt:                           PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp:                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt:                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Item:                            PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt:                                 PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Item:                 PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn:                      PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt:                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd:            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn:                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf:             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Item:                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Item:                                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt:                                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Tp_Issr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_CdtrRefInf_Ref:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Nm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Dept:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_SubDept:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_BldgNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Flr:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstBx:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Room:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_PstCd:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_DstrctNm:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_Ctry:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id:                                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_LEI:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr:                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                                      PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                                           PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtryOfRes:                                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Nm:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_MobNb:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_EmailPurp:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_JobTitl:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_Item:                                       PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr:                                            PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_Othr_Id:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Nm:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr:                                                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Room:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id:                                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_LEI:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                                      PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr:                                           PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                                     PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                                          PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtryOfRes:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Nm:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_MobNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_EmailPurp:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_JobTitl:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Dept:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_Item:                                      PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr:                                           PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_Othr_Id:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd:                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt:                                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_TaxId:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_RegnId:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Cdtr_TaxTp:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr:                                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_TaxId:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_RegnId:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_TaxTp:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn:                                            PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_AdmstnZone:                                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_RefNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Mtd:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt:                                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Dt:                                                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_SeqNb:                                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Tp:                                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Ctgy:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts:                                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_CertId:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd:                                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt:                                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate:                                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt:                                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt:                                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Item:                                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls:                                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf:                                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt:                                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp:                                                  PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Tp_Issr:                                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm:                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr:                                     PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry:                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine_Item:                        PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine:                             PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id:                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI:                                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Item:                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr:                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId:                                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth:                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Item:                         PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr:                              PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm:                      PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm:                                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb:                              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept:                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Item:                          PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr:                               PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id:                            PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr:                                       PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr:                               PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp:                         PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry:                   PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id:                PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm:           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id:                                    PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI:                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr:                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId:                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr:                        PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes:                             PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls:                              PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm:                           PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept:                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Item:                    PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr:                         PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp:                 PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id:                      PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd:                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RefNb:                                               PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_Dt:                                                  PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt:                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy:                                         PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value:                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd:                                   PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd:                                    PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf_Item:                                               PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_RmtInf_Strd_AddtlRmtInf:                                                    PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Item:                                                           PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData:                                                                PathTypeArray,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_PlcAndNm:                                                       PathTypeProperty,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Envlp:                                                          PathTypeStruct,
	Path_CstmrCdtTrfInitn_PmtInf_CdtTrfTxInf_SplmtryData_Envlp_Item:                                                     PathTypeProperty,
	Path_CstmrCdtTrfInitn_SplmtryData_Item:                                                                              PathTypeArrayItem,
	Path_CstmrCdtTrfInitn_SplmtryData:                                                                                   PathTypeArray,
	Path_CstmrCdtTrfInitn_SplmtryData_PlcAndNm:                                                                          PathTypeProperty,
	Path_CstmrCdtTrfInitn_SplmtryData_Envlp:                                                                             PathTypeStruct,
	Path_CstmrCdtTrfInitn_SplmtryData_Envlp_Item:                                                                        PathTypeProperty,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in pain_001_001_09", p)
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
