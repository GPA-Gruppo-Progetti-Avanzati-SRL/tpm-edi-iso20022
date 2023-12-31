// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_CdtrPmtActvtnReq                                                                                            = "CdtrPmtActvtnReq"
	Path_CdtrPmtActvtnReq_GrpHdr                                                                                     = "CdtrPmtActvtnReq.GrpHdr"
	Path_CdtrPmtActvtnReq_GrpHdr_MsgId                                                                               = "CdtrPmtActvtnReq.GrpHdr.MsgId"
	Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm                                                                             = "CdtrPmtActvtnReq.GrpHdr.CreDtTm"
	Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs                                                                             = "CdtrPmtActvtnReq.GrpHdr.NbOfTxs"
	Path_CdtrPmtActvtnReq_GrpHdr_CtrlSum                                                                             = "CdtrPmtActvtnReq.GrpHdr.CtrlSum"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty                                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm                                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr                                                                    = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd                                                           = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry                                                        = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id                                                     = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr                                                   = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm                                                = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Dept                                                               = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_SubDept                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_StrtNm                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNb                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNm                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Flr                                                                = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstBx                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Room                                                               = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstCd                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnNm                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnLctnNm                                                          = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_DstrctNm                                                           = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn                                                        = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Ctry                                                               = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrLine_Item                                                       = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrLine                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id                                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId                                                                   = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_AnyBIC                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_LEI                                                               = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Item                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id                                                           = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm                                                      = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd                                                   = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry                                                = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId                                                                  = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth                                                  = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                          = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                      = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                      = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                      = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Item                                                        = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Id                                                          = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm                                                     = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd                                                  = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry                                               = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr                                                        = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtryOfRes                                                                  = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtryOfRes"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls                                                                   = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_NmPrfx                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Nm                                                                = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PhneNb                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_MobNb                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_FaxNb                                                             = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailAdr                                                          = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailPurp                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_JobTitl                                                           = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Rspnsblty                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Dept                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_Item                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp                                                      = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_Id                                                           = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PrefrdMtd                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_Item                                                                                = "CdtrPmtActvtnReq.PmtInf[]"
	Path_CdtrPmtActvtnReq_PmtInf                                                                                     = "CdtrPmtActvtnReq.PmtInf" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_PmtInfId                                                                            = "CdtrPmtActvtnReq.PmtInf[].PmtInfId"
	Path_CdtrPmtActvtnReq_PmtInf_PmtMtd                                                                              = "CdtrPmtActvtnReq.PmtInf[].PmtMtd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf                                                                            = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_InstrPrty                                                                  = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.InstrPrty"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Item                                                                = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.SvcLvl[]"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl                                                                     = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.SvcLvl" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd                                                                  = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.SvcLvl[].Cd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Prtry                                                               = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.SvcLvl[].Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm                                                                  = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.LclInstrm"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Cd                                                               = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.LclInstrm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry                                                            = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.LclInstrm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp                                                                   = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.CtgyPurp"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd                                                                = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.CtgyPurp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Prtry                                                             = "CdtrPmtActvtnReq.PmtInf[].PmtTpInf.CtgyPurp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt                                                                         = "CdtrPmtActvtnReq.PmtInf[].ReqdExctnDt"
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt                                                                      = "CdtrPmtActvtnReq.PmtInf[].ReqdExctnDt.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_DtTm                                                                    = "CdtrPmtActvtnReq.PmtInf[].ReqdExctnDt.DtTm"
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt                                                                              = "CdtrPmtActvtnReq.PmtInf[].XpryDt"
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt                                                                           = "CdtrPmtActvtnReq.PmtInf[].XpryDt.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt_DtTm                                                                         = "CdtrPmtActvtnReq.PmtInf[].XpryDt.DtTm"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond                                                                             = "CdtrPmtActvtnReq.PmtInf[].PmtCond"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_AmtModAllwd                                                                 = "CdtrPmtActvtnReq.PmtInf[].PmtCond.AmtModAllwd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_EarlyPmtAllwd                                                               = "CdtrPmtActvtnReq.PmtInf[].PmtCond.EarlyPmtAllwd"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_DelyPnlty                                                                   = "CdtrPmtActvtnReq.PmtInf[].PmtCond.DelyPnlty"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt                                                                  = "CdtrPmtActvtnReq.PmtInf[].PmtCond.ImdtPmtRbt"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt                                                              = "CdtrPmtActvtnReq.PmtInf[].PmtCond.ImdtPmtRbt.Amt"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Ccy                                                          = "CdtrPmtActvtnReq.PmtInf[].PmtCond.ImdtPmtRbt.Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Value                                                        = "CdtrPmtActvtnReq.PmtInf[].PmtCond.ImdtPmtRbt.Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Rate                                                             = "CdtrPmtActvtnReq.PmtInf[].PmtCond.ImdtPmtRbt.Rate"
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_GrntedPmtReqd                                                               = "CdtrPmtActvtnReq.PmtInf[].PmtCond.GrntedPmtReqd"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr                                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm                                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr                                                                        = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Cd                                                               = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry                                                            = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Id                                                         = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Issr                                                       = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                                    = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Dept                                                                   = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_SubDept                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_StrtNm                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNb                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNm                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Flr                                                                    = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstBx                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Room                                                                   = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstCd                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnNm                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnLctnNm                                                              = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_DstrctNm                                                               = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_CtrySubDvsn                                                            = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Ctry                                                                   = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine_Item                                                           = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id                                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId                                                                       = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_AnyBIC                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_LEI                                                                   = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Item                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Id                                                               = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm                                                          = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                                                       = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                                                    = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Issr                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId                                                                      = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth                                                      = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                              = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                          = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                          = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                          = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Item                                                            = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id                                                              = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm                                                         = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                      = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                                   = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Issr                                                            = "CdtrPmtActvtnReq.PmtInf[].Dbtr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtryOfRes                                                                      = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls                                                                       = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_NmPrfx                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Nm                                                                    = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PhneNb                                                                = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_MobNb                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_FaxNb                                                                 = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailAdr                                                              = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailPurp                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_JobTitl                                                               = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Rspnsblty                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Dept                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_Item                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr                                                                  = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_ChanlTp                                                          = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_Id                                                               = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PrefrdMtd                                                             = "CdtrPmtActvtnReq.PmtInf[].Dbtr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct                                                                            = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id                                                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN                                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.IBAN"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr                                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Id                                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm                                                            = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd                                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Issr                                                               = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Id.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp                                                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Cd                                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Prtry                                                                   = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Ccy                                                                        = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Nm                                                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy                                                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Prxy"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp                                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Prxy.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Cd                                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Prxy.Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Prtry                                                              = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Prxy.Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Id                                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAcct.Prxy.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt                                                                             = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId                                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI                                                            = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.BICFI"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                             = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                          = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId                                                = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_LEI                                                              = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Nm                                                               = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr                                                          = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                              = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                           = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm                                                   = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb                                                   = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNm                                                   = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Flr                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstBx                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Room                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm                                                = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_DstrctNm                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                              = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                             = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr                                                             = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Id                                                          = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                               = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Issr                                                        = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.FinInstnId.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId                                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Id                                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_LEI                                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Nm                                                                  = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr                                                             = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp                                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                              = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                            = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept                                                        = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNm                                                      = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Flr                                                         = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstBx                                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Room                                                        = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd                                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm                                                       = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm                                                   = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_DstrctNm                                                    = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                                 = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry                                                        = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item                                                = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine                                                     = "CdtrPmtActvtnReq.PmtInf[].DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr                                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Nm                                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr                                                                   = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Cd                                                          = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry                                                       = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id                                                    = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr                                                  = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                               = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Dept                                                              = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_SubDept                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_StrtNm                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNb                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNm                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Flr                                                               = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstBx                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Room                                                              = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstCd                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnNm                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnLctnNm                                                         = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_DstrctNm                                                          = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn                                                       = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Ctry                                                              = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item                                                      = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrLine                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id                                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId                                                                  = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_AnyBIC                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_LEI                                                              = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id                                                          = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                                     = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                                  = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                               = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId                                                                 = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                                 = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                         = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                     = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                     = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                     = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item                                                       = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id                                                         = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                                    = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                                 = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                              = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr                                                       = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtryOfRes                                                                 = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls                                                                  = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_NmPrfx                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Nm                                                               = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PhneNb                                                           = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_MobNb                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_FaxNb                                                            = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailAdr                                                         = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailPurp                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_JobTitl                                                          = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Rspnsblty                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Dept                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_Item                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr                                                             = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_ChanlTp                                                     = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_Id                                                          = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PrefrdMtd                                                        = "CdtrPmtActvtnReq.PmtInf[].UltmtDbtr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_ChrgBr                                                                              = "CdtrPmtActvtnReq.PmtInf[].ChrgBr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Item                                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx                                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId                                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_InstrId                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtId.InstrId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtId.EndToEndId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_UETR                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtId.UETR"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_InstrPrty                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.InstrPrty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Item                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.SvcLvl[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.SvcLvl" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.SvcLvl[].Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Prtry                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.SvcLvl[].Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.LclInstrm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Cd                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.LclInstrm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.LclInstrm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.CtgyPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.CtgyPurp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Prtry                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtTpInf.CtgyPurp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.AmtModAllwd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.EarlyPmtAllwd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_DelyPnlty                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.DelyPnlty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.ImdtPmtRbt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.ImdtPmtRbt.Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Ccy                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.ImdtPmtRbt.Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Value                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.ImdtPmtRbt.Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Rate                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.ImdtPmtRbt.Rate"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].PmtCond.GrntedPmtReqd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt                                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.InstdAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.InstdAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.InstdAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.EqvtAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.EqvtAmt.Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Ccy                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.EqvtAmt.Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Value                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.EqvtAmt.Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_CcyOfTrf                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Amt.EqvtAmt.CcyOfTrf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr                                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChrgBr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqTp                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqNb                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Nm                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Cd                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Id                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Issr                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_SchmeNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Dept                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_SubDept                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_StrtNm                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNb                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNm                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Flr                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstBx                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Room                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstCd                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnLctnNm                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_DstrctNm                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_CtrySubDvsn                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Ctry                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrLine_Item                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrLine                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqFr.Adr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvryMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Cd                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvryMtd.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Prtry                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvryMtd.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Nm                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Cd                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Id                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Issr                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_SchmeNm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Dept                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_SubDept                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_StrtNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNb                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Flr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstBx                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Room                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstCd                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnNm                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnLctnNm                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_DstrctNm                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_CtrySubDvsn                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Ctry                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrLine_Item                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrLine                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.DlvrTo.Adr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_InstrPrty                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.InstrPrty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqMtrtyDt                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.ChqMtrtyDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_FrmsCd                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.FrmsCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_MemoFld_Item                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.MemoFld[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_MemoFld                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.MemoFld" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_RgnlClrZone                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.RgnlClrZone"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_PrtLctn                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.PrtLctn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_Sgntr_Item                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.Sgntr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_Sgntr                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].ChqInstr.Sgntr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr                                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Nm                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Cd                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Dept                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_SubDept                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_StrtNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Flr                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstBx                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Room                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstCd                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnNm                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnLctnNm                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_DstrctNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_CtrySubDvsn                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Ctry                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrLine_Item                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrLine                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_AnyBIC                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Issr                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Item                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Id                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Issr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtryOfRes                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_NmPrfx                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PhneNb                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_MobNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_FaxNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailAdr                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailPurp                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_JobTitl                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Rspnsblty                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Dept                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_ChanlTp                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PrefrdMtd                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtDbtr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_BICFI                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.BICFI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.ClrSysMmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_LEI                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Nm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Dept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_SubDept                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Flr                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstBx                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Room                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstCd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnLctnNm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_DstrctNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Ctry                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Id                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Issr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.FinInstnId.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_SubDept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_StrtNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Flr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstBx                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Room                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstCd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnLctnNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_DstrctNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Ctry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrLine                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt1.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_BICFI                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.BICFI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.ClrSysMmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.ClrSysMmbId.MmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_LEI                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Nm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Dept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_SubDept                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Flr                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstBx                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Room                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstCd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnLctnNm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_DstrctNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Ctry                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Id                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Issr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.FinInstnId.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_SubDept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_StrtNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Flr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstBx                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Room                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstCd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnLctnNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_DstrctNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Ctry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrLine                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt2.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_BICFI                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.BICFI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.ClrSysMmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.ClrSysMmbId.MmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_LEI                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Nm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Id                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Dept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_SubDept                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Flr                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstBx                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Room                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstCd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnLctnNm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_DstrctNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Ctry                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Id                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Issr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.FinInstnId.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_SubDept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_StrtNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Flr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstBx                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Room                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstCd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnLctnNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_DstrctNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Ctry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrLine                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].IntrmyAgt3.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.BICFI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_SubDept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_StrtNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Flr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstBx                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Room                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstCd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_DstrctNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Ctry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrLine                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Issr                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.FinInstnId.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Id                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_LEI                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Nm                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Dept                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_SubDept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_StrtNm                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNb                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNm                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Flr                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstBx                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Room                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstCd                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnNm                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_DstrctNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Ctry                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrLine                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr                                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Cd                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Id                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Issr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Dept                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_SubDept                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_StrtNm                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNb                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNm                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Flr                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstBx                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Room                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstCd                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnNm                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnLctnNm                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_DstrctNm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_CtrySubDvsn                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Ctry                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine_Item                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_AnyBIC                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_LEI                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Item                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Issr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Item                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Id                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Issr                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtryOfRes                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_NmPrfx                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Nm                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PhneNb                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_MobNb                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_FaxNb                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailAdr                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailPurp                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_JobTitl                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Rspnsblty                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Dept                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_Item                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_ChanlTp                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PrefrdMtd                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Cdtr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.IBAN"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Id                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Cd                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Prtry                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Issr                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Id.Othr.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Cd                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Prtry                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Ccy                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Nm                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Prxy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Prxy.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Cd                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Prxy.Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Prtry                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Prxy.Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Id                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].CdtrAcct.Prxy.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr                                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Nm                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Cd                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Dept                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_SubDept                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_StrtNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNm                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Flr                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstBx                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Room                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstCd                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnNm                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnLctnNm                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_DstrctNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_CtrySubDvsn                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Ctry                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrLine_Item                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrLine                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_AnyBIC                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_LEI                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Issr                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Item                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Id                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Issr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtryOfRes                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_NmPrfx                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PhneNb                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_MobNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_FaxNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailAdr                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailPurp                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_JobTitl                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Rspnsblty                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Dept                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_ChanlTp                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PrefrdMtd                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].UltmtCdtr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_Item                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].InstrForCdtrAgt[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].InstrForCdtrAgt" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_Cd                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].InstrForCdtrAgt[].Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_InstrInf                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].InstrForCdtrAgt[].InstrInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp                                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Purp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Cd                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Purp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Prtry                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Purp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Item                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_DbtCdtRptgInd                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].DbtCdtRptgInd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Authrty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Authrty.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Ctry                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Authrty.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Item                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Tp                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Dt                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Dt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Ctry                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Cd                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Ccy                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Value                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Inf_Item                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Inf[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Inf                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RgltryRptg[].Dtls[].Inf" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax                                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Cdtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxId                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Cdtr.TaxId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_RegnId                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Cdtr.RegnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxTp                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Cdtr.TaxTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxId                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.TaxId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_RegnId                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.RegnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxTp                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.TaxTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.Authstn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Titl                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.Authstn.Titl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Nm                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dbtr.Authstn.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_AdmstnZone                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.AdmstnZone"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_RefNb                                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.RefNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Mtd                                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Mtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxblBaseAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Ccy                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxblBaseAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Value                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxblBaseAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Ccy                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Value                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.TtlTaxAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dt                                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_SeqNb                                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.SeqNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Item                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd                                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Tp                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Ctgy                                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Ctgy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CtgyDtls                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].CtgyDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_DbtrSts                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].DbtrSts"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CertId                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].CertId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_FrmsCd                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].FrmsCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Yr                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd.Yr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Tp                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd.FrToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_FrDt                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd.FrToDt.FrDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_ToDt                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].Prd.FrToDt.ToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Rate                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Rate"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TtlAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Ccy                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Value                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Item                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Value                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_AddtlInf                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].Tax.Rcrd[].AddtlInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_Item                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf                                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtId                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_Item                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_Mtd                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].Mtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_ElctrncAdr                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].ElctrncAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Nm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Cd                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Id                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Issr                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_SchmeNm                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Dept                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_SubDept                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_StrtNm                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNb                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNm                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Flr                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstBx                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Room                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstCd                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnNm                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnLctnNm                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_DstrctNm                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_CtrySubDvsn                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Ctry                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine_Item                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RltdRmtInf[].RmtLctnDtls[].PstlAdr.Adr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf                                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd_Item                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Ustrd[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Ustrd" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Item                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Item                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_Issr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Nb                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_RltdDt                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Item                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.CdOrPrtry.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Tp.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].Nb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Id[].RltdDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Desc                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Desc"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DuePyblAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Item                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.DscntApldAmt[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.CdtNoteAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Item                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.TaxAmt[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Item                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].Rsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocInf[].LineDtls[].Amt.RmtdAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Item                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Item                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.TaxAmt[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_Issr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Ref                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].CdtrRefInf.Ref"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Nm                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Dept                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_SubDept                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_StrtNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNb                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Flr                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstBx                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Room                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstCd                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnNm                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_DstrctNm                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Ctry                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrLine                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_LEI                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtryOfRes                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Nm                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PhneNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_MobNb                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_FaxNb                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailPurp                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_JobTitl                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_Item                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_Id                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Nm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Dept                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_SubDept                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_StrtNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Flr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstBx                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Room                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstCd                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnNm                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_DstrctNm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Ctry                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrLine                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_LEI                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtryOfRes                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Nm                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PhneNb                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_MobNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_FaxNb                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailPurp                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_JobTitl                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Dept                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_Item                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_Id                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].Invcee.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Cdtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxId                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Cdtr.TaxId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_RegnId                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Cdtr.RegnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxTp                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Cdtr.TaxTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxId                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.TaxId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_RegnId                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.RegnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxTp                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.TaxTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn.Titl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dbtr.Authstn.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.TaxId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.RegnId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.TaxTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn.Titl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.UltmtDbtr.Authstn.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_AdmstnZone                                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.AdmstnZone"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_RefNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.RefNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Mtd                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Mtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxblBaseAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.TtlTaxAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dt                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_SeqNb                                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.SeqNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Tp                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Ctgy                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Ctgy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].CtgyDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].DbtrSts"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CertId                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].CertId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].FrmsCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.Yr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt.FrDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].Prd.FrToDt.ToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Rate"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Item                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].TaxRmt.Rcrd[].AddtlInf"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Tp.CdOrPrtry.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_Issr                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Tp.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine_Item                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Item                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Item                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Item                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Grnshee.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine_Item                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Item                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Item                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Item                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.GrnshmtAdmstr.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RefNb                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.RefNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Dt                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt.Ccy"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.RmtdAmt.Value"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.FmlyMdclInsrncInd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].GrnshmtRmt.MplyeeTermntnInd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_AddtlRmtInf_Item                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].AddtlRmtInf[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_AddtlRmtInf                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].RmtInf.Strd[].AddtlRmtInf" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Item                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile                                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Cd                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry                                                         = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Id                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_SchmeNm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Issr                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Tp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Id                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].IsseDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_Dt                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].IsseDt.Dt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_DtTm                                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].IsseDt.DtTm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nm                                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_LangCd                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].LangCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt                                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Cd                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Id                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_SchmeNm                                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Issr                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Frmt.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_FileNm                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].FileNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr                                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty                                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Nm                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Cd                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Id                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp.Prtry.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Issr                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp.Prtry.Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrTp.Prtry.SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Dept                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_SubDept                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.SubDept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_StrtNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.StrtNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNb                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.BldgNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNm                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.BldgNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Flr                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.Flr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstBx                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.PstBx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Room                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.Room"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstCd                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.PstCd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnNm                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.TwnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnLctnNm                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.TwnLctnNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_DstrctNm                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.DstrctNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_CtrySubDvsn                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.CtrySubDvsn"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Ctry                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.Ctry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine_Item                               = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrLine[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.PstlAdr.AdrLine" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id                                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_AnyBIC                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.AnyBIC"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_LEI                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.LEI"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Item                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Id                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Cd                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Prtry                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Issr                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.OrgId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Item                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Id                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[].SchmeNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Cd                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Issr                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.Id.PrvtId.Othr[].Issr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtryOfRes                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtryOfRes"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_NmPrfx                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.NmPrfx"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Nm                                        = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Nm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PhneNb                                    = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.PhneNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_MobNb                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.MobNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_FaxNb                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.FaxNb"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailAdr                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.EmailAdr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailPurp                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.EmailPurp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_JobTitl                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.JobTitl"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Rspnsblty                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Rspnsblty"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Dept                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Dept"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Item                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Othr[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr                                      = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Othr" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_ChanlTp                              = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Othr[].ChanlTp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Id                                   = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.Othr[].Id"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PrefrdMtd                                 = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Pty.CtctDtls.PrefrdMtd"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr                                                  = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Sgntr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr_Item                                             = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].DgtlSgntr.Sgntr.Item"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nclsr                                                            = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].NclsdFile[].Nclsr"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Item                                                           = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].SplmtryData[]"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData                                                                = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].SplmtryData" // ARRAY
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_PlcAndNm                                                       = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].SplmtryData[].PlcAndNm"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp                                                          = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].SplmtryData[].Envlp"
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp_Item                                                     = "CdtrPmtActvtnReq.PmtInf[].CdtTrfTx[].SplmtryData[].Envlp.Item"
	Path_CdtrPmtActvtnReq_SplmtryData_Item                                                                           = "CdtrPmtActvtnReq.SplmtryData[]"
	Path_CdtrPmtActvtnReq_SplmtryData                                                                                = "CdtrPmtActvtnReq.SplmtryData" // ARRAY
	Path_CdtrPmtActvtnReq_SplmtryData_PlcAndNm                                                                       = "CdtrPmtActvtnReq.SplmtryData[].PlcAndNm"
	Path_CdtrPmtActvtnReq_SplmtryData_Envlp                                                                          = "CdtrPmtActvtnReq.SplmtryData[].Envlp"
	Path_CdtrPmtActvtnReq_SplmtryData_Envlp_Item                                                                     = "CdtrPmtActvtnReq.SplmtryData[].Envlp.Item"
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_CdtrPmtActvtnReq:                                                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr:                                                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_MsgId:                                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm:                                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs:                                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_CtrlSum:                                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty:                                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm:                                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Cd:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Id:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_Issr:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp_Prtry_SchmeNm:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Dept:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_SubDept:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_StrtNm:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNb:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_BldgNm:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Flr:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstBx:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Room:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_PstCd:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnNm:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_TwnLctnNm:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_DstrctNm:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_CtrySubDvsn:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_Ctry:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrLine_Item:                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrLine:                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id:                                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_AnyBIC:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_LEI:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Item:                                                         PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr:                                                              PathTypeArray,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Item:                                                        PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr:                                                             PathTypeArray,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Id:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_PrvtId_Othr_Issr:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtryOfRes:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_NmPrfx:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Nm:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PhneNb:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_MobNb:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_FaxNb:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailAdr:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_EmailPurp:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_JobTitl:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Rspnsblty:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Dept:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_Item:                                                         PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr:                                                              PathTypeArray,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_ChanlTp:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_Othr_Id:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_CtctDtls_PrefrdMtd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Item:                                                                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf:                                                                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_PmtInfId:                                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtMtd:                                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf:                                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_InstrPrty:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Item:                                                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl:                                                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Prtry:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Cd:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Prtry:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt:                                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt:                                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_DtTm:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt:                                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt:                                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_XpryDt_DtTm:                                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond:                                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_AmtModAllwd:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_EarlyPmtAllwd:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_DelyPnlty:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Ccy:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Amt_Value:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_ImdtPmtRbt_Rate:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_PmtCond_GrntedPmtReqd:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr:                                                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm:                                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr:                                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Cd:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Id:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_Issr:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Dept:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_SubDept:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_StrtNm:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNb:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_BldgNm:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Flr:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstBx:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Room:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_PstCd:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnNm:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_TwnLctnNm:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_DstrctNm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_CtrySubDvsn:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_Ctry:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine_Item:                                                           PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine:                                                                PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id:                                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId:                                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_AnyBIC:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_LEI:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Item:                                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr:                                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Id:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_OrgId_Othr_Issr:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId:                                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Item:                                                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr:                                                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Issr:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtryOfRes:                                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls:                                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_NmPrfx:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Nm:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PhneNb:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_MobNb:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_FaxNb:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailAdr:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_EmailPurp:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_JobTitl:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Rspnsblty:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Dept:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_Item:                                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr:                                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_ChanlTp:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_Othr_Id:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_Dbtr_CtctDtls_PrefrdMtd:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct:                                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id:                                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Id:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Cd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_SchmeNm_Prtry:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_Othr_Issr:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp:                                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Cd:                                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Tp_Prtry:                                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Ccy:                                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Nm:                                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy:                                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Cd:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Tp_Prtry:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Prxy_Id:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt:                                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_LEI:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Nm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_BldgNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Flr:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstBx:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Room:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_TwnLctnNm:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_DstrctNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Id:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_Othr_Issr:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId:                                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Id:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_LEI:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_Nm:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp:                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Cd:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Dept:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_SubDept:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_StrtNm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNb:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_BldgNm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Flr:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstBx:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Room:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_PstCd:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnNm:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_TwnLctnNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_DstrctNm:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_Ctry:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_BrnchId_PstlAdr_AdrLine:                                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr:                                                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Nm:                                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Cd:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry:                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Dept:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_SubDept:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_StrtNm:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNb:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_BldgNm:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Flr:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstBx:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Room:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_PstCd:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnNm:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_TwnLctnNm:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_DstrctNm:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_CtrySubDvsn:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_Ctry:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrLine_Item:                                                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_PstlAdr_AdrLine:                                                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id:                                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_AnyBIC:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_LEI:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Item:                                                        PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr:                                                             PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Id:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_OrgId_Othr_Issr:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId:                                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Item:                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr:                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Id:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_Id_PrvtId_Othr_Issr:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtryOfRes:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_NmPrfx:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Nm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PhneNb:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_MobNb:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_FaxNb:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailAdr:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_EmailPurp:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_JobTitl:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Rspnsblty:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Dept:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_Item:                                                        PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr:                                                             PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_ChanlTp:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_Othr_Id:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_UltmtDbtr_CtctDtls_PrefrdMtd:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_ChrgBr:                                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Item:                                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx:                                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId:                                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_InstrId:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_UETR:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_InstrPrty:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Item:                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl:                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Prtry:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Cd:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Prtry:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_DelyPnlty:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Ccy:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Amt_Value:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_ImdtPmtRbt_Rate:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt:                                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt:                                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Ccy:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_Amt_Value:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_EqvtAmt_CcyOfTrf:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr:                                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqTp:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqNb:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Nm:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp:                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Cd:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Id:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_Issr:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrTp_Prtry_SchmeNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Dept:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_SubDept:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_StrtNm:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNb:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_BldgNm:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Flr:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstBx:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Room:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_PstCd:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_TwnLctnNm:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_DstrctNm:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_CtrySubDvsn:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_Ctry:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrLine_Item:                                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqFr_Adr_AdrLine:                                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Cd:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvryMtd_Prtry:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Nm:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Cd:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Id:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_Issr:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrTp_Prtry_SchmeNm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Dept:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_SubDept:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_StrtNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNb:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_BldgNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Flr:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstBx:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Room:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_PstCd:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnNm:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_TwnLctnNm:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_DstrctNm:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_CtrySubDvsn:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_Ctry:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrLine_Item:                                           PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_DlvrTo_Adr_AdrLine:                                                PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_InstrPrty:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_ChqMtrtyDt:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_FrmsCd:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_MemoFld_Item:                                                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_MemoFld:                                                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_RgnlClrZone:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_PrtLctn:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_Sgntr_Item:                                                        PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChqInstr_Sgntr:                                                             PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Nm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Cd:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Id:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_Issr:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_BldgNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Flr:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstBx:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Room:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_TwnLctnNm:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_DstrctNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_AnyBIC:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Id:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_OrgId_Othr_Issr:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Item:                                              PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr:                                                   PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Id:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_Id_PrvtId_Othr_Issr:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtryOfRes:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_NmPrfx:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PhneNb:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_MobNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_FaxNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailAdr:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_EmailPurp:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_JobTitl:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Rspnsblty:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Dept:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_ChanlTp:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_Othr_Id:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtDbtr_CtctDtls_PrefrdMtd:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1:                                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt1_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2:                                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt2_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3:                                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_BICFI:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_LEI:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry:                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_BldgNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Flr:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstBx:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Room:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_TwnLctnNm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_DstrctNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_IntrmyAgt3_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Room:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Id:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_Othr_Issr:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Id:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_LEI:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_Nm:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Cd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Id:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_Issr:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrTp_Prtry_SchmeNm:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Dept:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_SubDept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_StrtNm:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNb:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_BldgNm:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Flr:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstBx:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Room:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_PstCd:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnNm:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_TwnLctnNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_DstrctNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_Ctry:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_BrnchId_PstlAdr_AdrLine:                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr:                                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Cd:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry:                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Id:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_Issr:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Dept:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_SubDept:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_StrtNm:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNb:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_BldgNm:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Flr:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstBx:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Room:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_PstCd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnNm:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_TwnLctnNm:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_DstrctNm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_CtrySubDvsn:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_Ctry:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine_Item:                                                  PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine:                                                       PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id:                                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_AnyBIC:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_LEI:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Item:                                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr:                                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Issr:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Item:                                                   PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr:                                                        PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Id:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_PrvtId_Othr_Issr:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtryOfRes:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_NmPrfx:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Nm:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PhneNb:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_MobNb:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_FaxNb:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailAdr:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_EmailPurp:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_JobTitl:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Rspnsblty:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Dept:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_Item:                                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr:                                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_ChanlTp:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_Othr_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_CtctDtls_PrefrdMtd:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id:                                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr:                                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Id:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm:                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Cd:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_SchmeNm_Prtry:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_Othr_Issr:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp:                                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Cd:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Tp_Prtry:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Ccy:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Nm:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp:                                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Cd:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Tp_Prtry:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Prxy_Id:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr:                                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Nm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Cd:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry:                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Id:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_Issr:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrTp_Prtry_SchmeNm:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Dept:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_SubDept:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_StrtNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_BldgNm:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Flr:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstBx:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Room:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_PstCd:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnNm:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_TwnLctnNm:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_DstrctNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_CtrySubDvsn:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_Ctry:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrLine_Item:                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_PstlAdr_AdrLine:                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_AnyBIC:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_LEI:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Id:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_OrgId_Othr_Issr:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Item:                                              PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr:                                                   PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Id:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_Id_PrvtId_Othr_Issr:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtryOfRes:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_NmPrfx:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PhneNb:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_MobNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_FaxNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailAdr:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_EmailPurp:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_JobTitl:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Rspnsblty:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Dept:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_ChanlTp:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_Othr_Id:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_UltmtCdtr_CtctDtls_PrefrdMtd:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_Item:                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt:                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_Cd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_InstrForCdtrAgt_InstrInf:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp:                                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Cd:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Purp_Prtry:                                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Item:                                                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg:                                                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_DbtCdtRptgInd:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Authrty_Ctry:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Item:                                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls:                                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Tp:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Dt:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Ctry:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Cd:                                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Ccy:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Amt_Value:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Inf_Item:                                                   PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RgltryRptg_Dtls_Inf:                                                        PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax:                                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxId:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_RegnId:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Cdtr_TaxTp:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr:                                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxId:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_RegnId:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_TaxTp:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn:                                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Titl:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dbtr_Authstn_Nm:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_AdmstnZone:                                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_RefNb:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Mtd:                                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Ccy:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxblBaseAmt_Value:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt:                                                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Ccy:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_TtlTaxAmt_Value:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Dt:                                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_SeqNb:                                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Item:                                                              PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd:                                                                   PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Tp:                                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Ctgy:                                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CtgyDtls:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_DbtrSts:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_CertId:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_FrmsCd:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Yr:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_Tp:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_FrDt:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_Prd_FrToDt_ToDt:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt:                                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Rate:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt:                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_TtlAmt_Value:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Item:                                                  PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls:                                                       PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd:                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt:                                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Tax_Rcrd_AddtlInf:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_Item:                                                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf:                                                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtId:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_Item:                                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls:                                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_Mtd:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_ElctrncAdr:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Nm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp:                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Cd:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Id:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_Issr:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrTp_Prtry_SchmeNm:                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Dept:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_SubDept:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_StrtNm:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNb:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_BldgNm:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Flr:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstBx:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Room:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_PstCd:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnNm:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_TwnLctnNm:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_DstrctNm:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_CtrySubDvsn:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_Ctry:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine_Item:                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RltdRmtInf_RmtLctnDtls_PstlAdr_Adr_AdrLine:                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf:                                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd_Item:                                                          PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd:                                                               PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Item:                                                           PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd:                                                                PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Item:                                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf:                                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Tp_Issr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_Nb:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_RltdDt:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Item:                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls:                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp:                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry:                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Cd:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_CdOrPrtry_Prtry:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Tp_Issr:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_Nb:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Id_RltdDt:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Desc:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Ccy:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DuePyblAmt_Value:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Item:                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt:                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp:                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Cd:                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Tp_Prtry:                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt:                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Ccy:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_DscntApldAmt_Amt_Value:                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Ccy:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_CdtNoteAmt_Value:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Item:                            PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt:                                 PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp:                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Cd:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Tp_Prtry:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Ccy:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_TaxAmt_Amt_Value:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Item:                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn:                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt:                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Ccy:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Amt_Value:            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_CdtDbtInd:            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_Rsn:                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_AdjstmntAmtAndRsn_AddtlInf:             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt:                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Ccy:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocInf_LineDtls_Amt_RmtdAmt_Value:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Item:                                   PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                                        PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Cd:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Tp_Prtry:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Ccy:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Amt_Value:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Item:                                         PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt:                                              PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Tp_Prtry:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Ccy:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_TaxAmt_Amt_Value:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:                              PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                                   PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Tp_Issr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_CdtrRefInf_Ref:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Nm:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Cd:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry:                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Id:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_Issr:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrTp_Prtry_SchmeNm:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Dept:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_SubDept:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_BldgNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Flr:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstBx:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Room:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_PstCd:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_TwnLctnNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_DstrctNm:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_Ctry:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                                     PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                                          PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id:                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_AnyBIC:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_LEI:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr:                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId:                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtryOfRes:                                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Nm:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_MobNb:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_EmailPurp:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_JobTitl:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Rspnsblty:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_Item:                                       PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr:                                            PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_ChanlTp:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_Othr_Id:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcr_CtctDtls_PrefrdMtd:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Nm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Cd:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrTp_Prtry_SchmeNm:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Dept:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_BldgNm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Flr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstBx:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Room:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_TwnLctnNm:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_DstrctNm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id:                                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId:                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_AnyBIC:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_LEI:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr:                                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId:                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                                     PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                                          PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtryOfRes:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls:                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Nm:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_MobNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_EmailPurp:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_JobTitl:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Rspnsblty:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Dept:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_Item:                                      PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr:                                           PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_ChanlTp:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_Othr_Id:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_Invcee_CtctDtls_PrefrdMtd:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxId:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_RegnId:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Cdtr_TaxTp:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxId:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_RegnId:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_TaxTp:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Titl:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dbtr_Authstn_Nm:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr:                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxId:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_RegnId:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_TaxTp:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn:                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Titl:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_UltmtDbtr_Authstn_Nm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_AdmstnZone:                                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_RefNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Mtd:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Ccy:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxblBaseAmt_Value:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt:                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Ccy:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_TtlTaxAmt_Value:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Dt:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_SeqNb:                                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Tp:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Ctgy:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CtgyDtls:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_DbtrSts:                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_CertId:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_FrmsCd:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd:                                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Yr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_Tp:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt:                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_FrDt:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_Prd_FrToDt_ToDt:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Rate:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt:                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TaxblBaseAmt_Value:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt:                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Ccy:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_TtlAmt_Value:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Item:                                   PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls:                                        PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Yr:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_Tp:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Ccy:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_TaxAmt_Dtls_Amt_Value:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_TaxRmt_Rcrd_AddtlInf:                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt:                                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry:                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Cd:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_CdOrPrtry_Prtry:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Tp_Issr:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Nm:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr:                                     PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp:                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Cd:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry:                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Id:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_Issr:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrTp_Prtry_SchmeNm:                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Dept:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_SubDept:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_StrtNm:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNb:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_BldgNm:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Flr:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstBx:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Room:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_PstCd:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnNm:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_TwnLctnNm:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_DstrctNm:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_CtrySubDvsn:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_Ctry:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine_Item:                        PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_PstlAdr_AdrLine:                             PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_AnyBIC:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_LEI:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Item:                          PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr:                               PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Id:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm:                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_OrgId_Othr_Issr:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId:                                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth:                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Item:                         PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr:                              PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Id:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm:                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Cd:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_SchmeNm_Prtry:                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_Id_PrvtId_Othr_Issr:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtryOfRes:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_NmPrfx:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Nm:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PhneNb:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_MobNb:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_FaxNb:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailAdr:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_EmailPurp:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_JobTitl:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Rspnsblty:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Dept:                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Item:                          PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr:                               PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_ChanlTp:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_Othr_Id:                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Grnshee_CtctDtls_PrefrdMtd:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr:                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Nm:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr:                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp:                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Cd:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry:                   PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Id:                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_Issr:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrTp_Prtry_SchmeNm:           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Dept:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_SubDept:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_BldgNm:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Flr:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstBx:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Room:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_PstCd:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_TwnLctnNm:                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_DstrctNm:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_Ctry:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_PstlAdr_AdrLine:                       PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id:                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId:                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_AnyBIC:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_LEI:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr:                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr:                        PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtryOfRes:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls:                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Nm:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_MobNb:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_EmailPurp:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_JobTitl:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Rspnsblty:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Dept:                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Item:                    PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr:                         PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_ChanlTp:                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_Othr_Id:                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_GrnshmtAdmstr_CtctDtls_PrefrdMtd:                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RefNb:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_Dt:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt:                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Ccy:                                         PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_RmtdAmt_Value:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_FmlyMdclInsrncInd:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_GrnshmtRmt_MplyeeTermntnInd:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_AddtlRmtInf_Item:                                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Strd_AddtlRmtInf:                                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Item:                                                             PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile:                                                                  PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp:                                                               PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Cd:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry:                                                         PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Id:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_SchmeNm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Tp_Prtry_Issr:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Id:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt:                                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_Dt:                                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_IsseDt_DtTm:                                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nm:                                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_LangCd:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt:                                                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Cd:                                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry:                                                       PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Id:                                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_SchmeNm:                                               PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Frmt_Prtry_Issr:                                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_FileNm:                                                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr:                                                        PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty:                                                    PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Nm:                                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr:                                            PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp:                                      PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Cd:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry:                                PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Id:                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_Issr:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrTp_Prtry_SchmeNm:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Dept:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_SubDept:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_StrtNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNb:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_BldgNm:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Flr:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstBx:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Room:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_PstCd:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnNm:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_TwnLctnNm:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_DstrctNm:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_CtrySubDvsn:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_Ctry:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine_Item:                               PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_PstlAdr_AdrLine:                                    PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id:                                                 PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_AnyBIC:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_LEI:                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Item:                                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr:                                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Id:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm:                              PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Cd:                           PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_SchmeNm_Prtry:                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_OrgId_Othr_Issr:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId:                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth:                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Item:                                PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr:                                     PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Id:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm:                             PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Cd:                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_Id_PrvtId_Othr_Issr:                                PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtryOfRes:                                          PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls:                                           PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_NmPrfx:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Nm:                                        PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PhneNb:                                    PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_MobNb:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_FaxNb:                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailAdr:                                  PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_EmailPurp:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_JobTitl:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Rspnsblty:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Dept:                                      PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Item:                                 PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr:                                      PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_ChanlTp:                              PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_Othr_Id:                                   PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Pty_CtctDtls_PrefrdMtd:                                 PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr:                                                  PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_DgtlSgntr_Sgntr_Item:                                             PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_NclsdFile_Nclsr:                                                            PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Item:                                                           PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData:                                                                PathTypeArray,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_PlcAndNm:                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp:                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_SplmtryData_Envlp_Item:                                                     PathTypeProperty,
	Path_CdtrPmtActvtnReq_SplmtryData_Item:                                                                           PathTypeArrayItem,
	Path_CdtrPmtActvtnReq_SplmtryData:                                                                                PathTypeArray,
	Path_CdtrPmtActvtnReq_SplmtryData_PlcAndNm:                                                                       PathTypeProperty,
	Path_CdtrPmtActvtnReq_SplmtryData_Envlp:                                                                          PathTypeStruct,
	Path_CdtrPmtActvtnReq_SplmtryData_Envlp_Item:                                                                     PathTypeProperty,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in pain_013_001_07", p)
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
