// Package camt_053_001_02
// Do not Edit. This stuff it's been automatically generated.
package camt_053_001_02

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	Path_BkToCstmrStmt                                                                                     = "BkToCstmrStmt"
	Path_BkToCstmrStmt_GrpHdr                                                                              = "BkToCstmrStmt.GrpHdr"
	Path_BkToCstmrStmt_GrpHdr_MsgId                                                                        = "BkToCstmrStmt.GrpHdr.MsgId"
	Path_BkToCstmrStmt_GrpHdr_CreDtTm                                                                      = "BkToCstmrStmt.GrpHdr.CreDtTm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt                                                                      = "BkToCstmrStmt.GrpHdr.MsgRcpt"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Nm                                                                   = "BkToCstmrStmt.GrpHdr.MsgRcpt.Nm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr                                                              = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrTp                                                        = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_Dept                                                         = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.Dept"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_SubDept                                                      = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.SubDept"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_StrtNm                                                       = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_BldgNb                                                       = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_PstCd                                                        = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.PstCd"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_TwnNm                                                        = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_CtrySubDvsn                                                  = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_Ctry                                                         = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.Ctry"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrLine_Item                                                 = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrLine                                                      = "BkToCstmrStmt.GrpHdr.MsgRcpt.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id                                                                   = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId                                                             = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_BICOrBEI                                                    = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Item                                                   = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr                                                        = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Id                                                     = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm                                                = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm_Cd                                             = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm_Prtry                                          = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Issr                                                   = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId                                                            = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth                                            = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                    = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Item                                                  = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr                                                       = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Id                                                    = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm                                               = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm_Cd                                            = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm_Prtry                                         = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Issr                                                  = "BkToCstmrStmt.GrpHdr.MsgRcpt.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtryOfRes                                                            = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtryOfRes"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls                                                             = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_NmPrfx                                                      = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_Nm                                                          = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.Nm"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_PhneNb                                                      = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_MobNb                                                       = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.MobNb"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_FaxNb                                                       = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_EmailAdr                                                    = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_Othr                                                        = "BkToCstmrStmt.GrpHdr.MsgRcpt.CtctDtls.Othr"
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn                                                                     = "BkToCstmrStmt.GrpHdr.MsgPgntn"
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn_PgNb                                                                = "BkToCstmrStmt.GrpHdr.MsgPgntn.PgNb"
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn_LastPgInd                                                           = "BkToCstmrStmt.GrpHdr.MsgPgntn.LastPgInd"
	Path_BkToCstmrStmt_GrpHdr_AddtlInf                                                                     = "BkToCstmrStmt.GrpHdr.AddtlInf"
	Path_BkToCstmrStmt_Stmt_Item                                                                           = "BkToCstmrStmt.Stmt[]"
	Path_BkToCstmrStmt_Stmt                                                                                = "BkToCstmrStmt.Stmt" // ARRAY
	Path_BkToCstmrStmt_Stmt_Id                                                                             = "BkToCstmrStmt.Stmt[].Id"
	Path_BkToCstmrStmt_Stmt_ElctrncSeqNb                                                                   = "BkToCstmrStmt.Stmt[].ElctrncSeqNb"
	Path_BkToCstmrStmt_Stmt_LglSeqNb                                                                       = "BkToCstmrStmt.Stmt[].LglSeqNb"
	Path_BkToCstmrStmt_Stmt_CreDtTm                                                                        = "BkToCstmrStmt.Stmt[].CreDtTm"
	Path_BkToCstmrStmt_Stmt_FrToDt                                                                         = "BkToCstmrStmt.Stmt[].FrToDt"
	Path_BkToCstmrStmt_Stmt_FrToDt_FrDtTm                                                                  = "BkToCstmrStmt.Stmt[].FrToDt.FrDtTm"
	Path_BkToCstmrStmt_Stmt_FrToDt_ToDtTm                                                                  = "BkToCstmrStmt.Stmt[].FrToDt.ToDtTm"
	Path_BkToCstmrStmt_Stmt_CpyDplctInd                                                                    = "BkToCstmrStmt.Stmt[].CpyDplctInd"
	Path_BkToCstmrStmt_Stmt_RptgSrc                                                                        = "BkToCstmrStmt.Stmt[].RptgSrc"
	Path_BkToCstmrStmt_Stmt_RptgSrc_Cd                                                                     = "BkToCstmrStmt.Stmt[].RptgSrc.Cd"
	Path_BkToCstmrStmt_Stmt_RptgSrc_Prtry                                                                  = "BkToCstmrStmt.Stmt[].RptgSrc.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct                                                                           = "BkToCstmrStmt.Stmt[].Acct"
	Path_BkToCstmrStmt_Stmt_Acct_Id                                                                        = "BkToCstmrStmt.Stmt[].Acct.Id"
	Path_BkToCstmrStmt_Stmt_Acct_Id_IBAN                                                                   = "BkToCstmrStmt.Stmt[].Acct.Id.IBAN"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr                                                                   = "BkToCstmrStmt.Stmt[].Acct.Id.Othr"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_Id                                                                = "BkToCstmrStmt.Stmt[].Acct.Id.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm                                                           = "BkToCstmrStmt.Stmt[].Acct.Id.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm_Cd                                                        = "BkToCstmrStmt.Stmt[].Acct.Id.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm_Prtry                                                     = "BkToCstmrStmt.Stmt[].Acct.Id.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_Issr                                                              = "BkToCstmrStmt.Stmt[].Acct.Id.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Acct_Tp                                                                        = "BkToCstmrStmt.Stmt[].Acct.Tp"
	Path_BkToCstmrStmt_Stmt_Acct_Tp_Cd                                                                     = "BkToCstmrStmt.Stmt[].Acct.Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Tp_Prtry                                                                  = "BkToCstmrStmt.Stmt[].Acct.Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Ccy                                                                       = "BkToCstmrStmt.Stmt[].Acct.Ccy"
	Path_BkToCstmrStmt_Stmt_Acct_Nm                                                                        = "BkToCstmrStmt.Stmt[].Acct.Nm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr                                                                      = "BkToCstmrStmt.Stmt[].Acct.Ownr"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Nm                                                                   = "BkToCstmrStmt.Stmt[].Acct.Ownr.Nm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr                                                              = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrTp                                                        = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_Dept                                                         = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_SubDept                                                      = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_StrtNm                                                       = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_BldgNb                                                       = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_PstCd                                                        = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_TwnNm                                                        = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_CtrySubDvsn                                                  = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_Ctry                                                         = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrLine_Item                                                 = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrLine                                                      = "BkToCstmrStmt.Stmt[].Acct.Ownr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id                                                                   = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId                                                             = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_BICOrBEI                                                    = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Item                                                   = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr                                                        = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Id                                                     = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm                                                = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm_Cd                                             = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm_Prtry                                          = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Issr                                                   = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId                                                            = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth                                            = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_BirthDt                                    = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth                                = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth                                = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth                                = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Item                                                  = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr                                                       = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Id                                                    = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm                                               = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm_Cd                                            = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm_Prtry                                         = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Issr                                                  = "BkToCstmrStmt.Stmt[].Acct.Ownr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtryOfRes                                                            = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls                                                             = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_NmPrfx                                                      = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_Nm                                                          = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_PhneNb                                                      = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_MobNb                                                       = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_FaxNb                                                       = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_EmailAdr                                                    = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_Othr                                                        = "BkToCstmrStmt.Stmt[].Acct.Ownr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr                                                                      = "BkToCstmrStmt.Stmt[].Acct.Svcr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId                                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_BIC                                                       = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId                                               = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId                                      = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId_Cd                                   = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                                = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_MmbId                                         = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Nm                                                        = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr                                                   = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrTp                                             = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_Dept                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_SubDept                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_StrtNm                                            = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_BldgNb                                            = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_PstCd                                             = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_TwnNm                                             = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_CtrySubDvsn                                       = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_Ctry                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrLine_Item                                      = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrLine                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr                                                      = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_Id                                                   = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm_Cd                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm_Prtry                                        = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_Issr                                                 = "BkToCstmrStmt.Stmt[].Acct.Svcr.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId                                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_Id                                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_Nm                                                           = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr                                                      = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrTp                                                = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_Dept                                                 = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_SubDept                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_StrtNm                                               = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_BldgNb                                               = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_PstCd                                                = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_TwnNm                                                = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_CtrySubDvsn                                          = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_Ctry                                                 = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrLine_Item                                         = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrLine                                              = "BkToCstmrStmt.Stmt[].Acct.Svcr.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_RltdAcct                                                                       = "BkToCstmrStmt.Stmt[].RltdAcct"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id                                                                    = "BkToCstmrStmt.Stmt[].RltdAcct.Id"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_IBAN                                                               = "BkToCstmrStmt.Stmt[].RltdAcct.Id.IBAN"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr                                                               = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_Id                                                            = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr.Id"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm                                                       = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm_Cd                                                    = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm_Prtry                                                 = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_Issr                                                          = "BkToCstmrStmt.Stmt[].RltdAcct.Id.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp                                                                    = "BkToCstmrStmt.Stmt[].RltdAcct.Tp"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp_Cd                                                                 = "BkToCstmrStmt.Stmt[].RltdAcct.Tp.Cd"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp_Prtry                                                              = "BkToCstmrStmt.Stmt[].RltdAcct.Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Ccy                                                                   = "BkToCstmrStmt.Stmt[].RltdAcct.Ccy"
	Path_BkToCstmrStmt_Stmt_RltdAcct_Nm                                                                    = "BkToCstmrStmt.Stmt[].RltdAcct.Nm"
	Path_BkToCstmrStmt_Stmt_Intrst_Item                                                                    = "BkToCstmrStmt.Stmt[].Intrst[]"
	Path_BkToCstmrStmt_Stmt_Intrst                                                                         = "BkToCstmrStmt.Stmt[].Intrst" // ARRAY
	Path_BkToCstmrStmt_Stmt_Intrst_Tp                                                                      = "BkToCstmrStmt.Stmt[].Intrst[].Tp"
	Path_BkToCstmrStmt_Stmt_Intrst_Tp_Cd                                                                   = "BkToCstmrStmt.Stmt[].Intrst[].Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Intrst_Tp_Prtry                                                                = "BkToCstmrStmt.Stmt[].Intrst[].Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Item                                                               = "BkToCstmrStmt.Stmt[].Intrst[].Rate[]"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate                                                                    = "BkToCstmrStmt.Stmt[].Intrst[].Rate" // ARRAY
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp                                                                 = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].Tp"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp_Pctg                                                            = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].Tp.Pctg"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp_Othr                                                            = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].Tp.Othr"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg                                                            = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt                                                        = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt                                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt                                          = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl                                             = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt                                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt                                          = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl                                             = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt                                                = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt                                          = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl                                     = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt                                          = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl                                     = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_EQAmt                                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.EQAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_NEQAmt                                                 = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Amt.NEQAmt"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_CdtDbtInd                                                  = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Ccy                                                        = "BkToCstmrStmt.Stmt[].Intrst[].Rate[].VldtyRg.Ccy"
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt                                                                  = "BkToCstmrStmt.Stmt[].Intrst[].FrToDt"
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt_FrDtTm                                                           = "BkToCstmrStmt.Stmt[].Intrst[].FrToDt.FrDtTm"
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt_ToDtTm                                                           = "BkToCstmrStmt.Stmt[].Intrst[].FrToDt.ToDtTm"
	Path_BkToCstmrStmt_Stmt_Intrst_Rsn                                                                     = "BkToCstmrStmt.Stmt[].Intrst[].Rsn"
	Path_BkToCstmrStmt_Stmt_Bal_Item                                                                       = "BkToCstmrStmt.Stmt[].Bal[]"
	Path_BkToCstmrStmt_Stmt_Bal                                                                            = "BkToCstmrStmt.Stmt[].Bal" // ARRAY
	Path_BkToCstmrStmt_Stmt_Bal_Tp                                                                         = "BkToCstmrStmt.Stmt[].Bal[].Tp"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry                                                               = "BkToCstmrStmt.Stmt[].Bal[].Tp.CdOrPrtry"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry_Cd                                                            = "BkToCstmrStmt.Stmt[].Bal[].Tp.CdOrPrtry.Cd"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry_Prtry                                                         = "BkToCstmrStmt.Stmt[].Bal[].Tp.CdOrPrtry.Prtry"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp                                                                   = "BkToCstmrStmt.Stmt[].Bal[].Tp.SubTp"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp_Cd                                                                = "BkToCstmrStmt.Stmt[].Bal[].Tp.SubTp.Cd"
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp_Prtry                                                             = "BkToCstmrStmt.Stmt[].Bal[].Tp.SubTp.Prtry"
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine                                                                    = "BkToCstmrStmt.Stmt[].Bal[].CdtLine"
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Incl                                                               = "BkToCstmrStmt.Stmt[].Bal[].CdtLine.Incl"
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt                                                                = "BkToCstmrStmt.Stmt[].Bal[].CdtLine.Amt"
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt_Ccy                                                            = "BkToCstmrStmt.Stmt[].Bal[].CdtLine.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt_Value                                                          = "BkToCstmrStmt.Stmt[].Bal[].CdtLine.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Bal_Amt                                                                        = "BkToCstmrStmt.Stmt[].Bal[].Amt"
	Path_BkToCstmrStmt_Stmt_Bal_Amt_Ccy                                                                    = "BkToCstmrStmt.Stmt[].Bal[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Bal_Amt_Value                                                                  = "BkToCstmrStmt.Stmt[].Bal[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Bal_CdtDbtInd                                                                  = "BkToCstmrStmt.Stmt[].Bal[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Bal_Dt                                                                         = "BkToCstmrStmt.Stmt[].Bal[].Dt"
	Path_BkToCstmrStmt_Stmt_Bal_Dt_Dt                                                                      = "BkToCstmrStmt.Stmt[].Bal[].Dt.Dt"
	Path_BkToCstmrStmt_Stmt_Bal_Dt_DtTm                                                                    = "BkToCstmrStmt.Stmt[].Bal[].Dt.DtTm"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Item                                                                = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[]"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty                                                                     = "BkToCstmrStmt.Stmt[].Bal[].Avlbty" // ARRAY
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt                                                                  = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Dt"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt_NbOfDays                                                         = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Dt.NbOfDays"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt_ActlDt                                                           = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Dt.ActlDt"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt                                                                 = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Amt"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt_Ccy                                                             = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt_Value                                                           = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_CdtDbtInd                                                           = "BkToCstmrStmt.Stmt[].Bal[].Avlbty[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_TxsSummry                                                                      = "BkToCstmrStmt.Stmt[].TxsSummry"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries                                                            = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_NbOfNtries                                                 = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtries.NbOfNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_Sum                                                        = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtries.Sum"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_TtlNetNtryAmt                                              = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtries.TtlNetNtryAmt"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_CdtDbtInd                                                  = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtries.CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries                                                         = "BkToCstmrStmt.Stmt[].TxsSummry.TtlCdtNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries_NbOfNtries                                              = "BkToCstmrStmt.Stmt[].TxsSummry.TtlCdtNtries.NbOfNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries_Sum                                                     = "BkToCstmrStmt.Stmt[].TxsSummry.TtlCdtNtries.Sum"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries                                                         = "BkToCstmrStmt.Stmt[].TxsSummry.TtlDbtNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries_NbOfNtries                                              = "BkToCstmrStmt.Stmt[].TxsSummry.TtlDbtNtries.NbOfNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries_Sum                                                     = "BkToCstmrStmt.Stmt[].TxsSummry.TtlDbtNtries.Sum"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Item                                              = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[]"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd                                                   = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd" // ARRAY
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_NbOfNtries                                        = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].NbOfNtries"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Sum                                               = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Sum"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_TtlNetNtryAmt                                     = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].TtlNetNtryAmt"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_CdtDbtInd                                         = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_FcstInd                                           = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].FcstInd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd                                            = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn                                       = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Domn"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Cd                                    = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Domn.Cd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly                                  = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Domn.Fmly"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly_Cd                               = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Domn.Fmly.Cd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly_SubFmlyCd                        = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Domn.Fmly.SubFmlyCd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry                                      = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Prtry"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry_Cd                                   = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Prtry.Cd"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry_Issr                                 = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].BkTxCd.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Item                                       = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[]"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty                                            = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty" // ARRAY
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt                                         = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Dt"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt_NbOfDays                                = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Dt.NbOfDays"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt_ActlDt                                  = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Dt.ActlDt"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt                                        = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Amt"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt_Ccy                                    = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt_Value                                  = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_CdtDbtInd                                  = "BkToCstmrStmt.Stmt[].TxsSummry.TtlNtriesPerBkTxCd[].Avlbty[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_Item                                                                      = "BkToCstmrStmt.Stmt[].Ntry[]"
	Path_BkToCstmrStmt_Stmt_Ntry                                                                           = "BkToCstmrStmt.Stmt[].Ntry" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryRef                                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryRef"
	Path_BkToCstmrStmt_Stmt_Ntry_Amt                                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Amt_Ccy                                                                   = "BkToCstmrStmt.Stmt[].Ntry[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Amt_Value                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_CdtDbtInd                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_RvslInd                                                                   = "BkToCstmrStmt.Stmt[].Ntry[].RvslInd"
	Path_BkToCstmrStmt_Stmt_Ntry_Sts                                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Sts"
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt                                                                   = "BkToCstmrStmt.Stmt[].Ntry[].BookgDt"
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt_Dt                                                                = "BkToCstmrStmt.Stmt[].Ntry[].BookgDt.Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt_DtTm                                                              = "BkToCstmrStmt.Stmt[].Ntry[].BookgDt.DtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt                                                                     = "BkToCstmrStmt.Stmt[].Ntry[].ValDt"
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt_Dt                                                                  = "BkToCstmrStmt.Stmt[].Ntry[].ValDt.Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt_DtTm                                                                = "BkToCstmrStmt.Stmt[].Ntry[].ValDt.DtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_AcctSvcrRef                                                               = "BkToCstmrStmt.Stmt[].Ntry[].AcctSvcrRef"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Item                                                               = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty                                                                    = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt_NbOfDays                                                        = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Dt.NbOfDays"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt_ActlDt                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Dt.ActlDt"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt                                                                = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt_Ccy                                                            = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt_Value                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_CdtDbtInd                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Avlbty[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd                                                                    = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn                                                               = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Domn"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Cd                                                            = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Domn.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly                                                          = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Domn.Fmly"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly_Cd                                                       = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Domn.Fmly.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly_SubFmlyCd                                                = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Domn.Fmly.SubFmlyCd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry                                                              = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry_Cd                                                           = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Prtry.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry_Issr                                                         = "BkToCstmrStmt.Stmt[].Ntry[].BkTxCd.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_ComssnWvrInd                                                              = "BkToCstmrStmt.Stmt[].Ntry[].ComssnWvrInd"
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd                                                               = "BkToCstmrStmt.Stmt[].Ntry[].AddtlInfInd"
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd_MsgNmId                                                       = "BkToCstmrStmt.Stmt[].Ntry[].AddtlInfInd.MsgNmId"
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd_MsgId                                                         = "BkToCstmrStmt.Stmt[].Ntry[].AddtlInfInd.MsgId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls                                                                   = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt                                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt                                                      = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt_Ccy                                                  = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt_Value                                                = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg                                                  = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_SrcCcy                                           = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_TrgtCcy                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_UnitCcy                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_XchgRate                                         = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_CtrctId                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_QtnDt                                            = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.InstdAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt                                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt                                                         = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt_Ccy                                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt_Value                                                   = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg                                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_SrcCcy                                              = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_TrgtCcy                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_UnitCcy                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_XchgRate                                            = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_CtrctId                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_QtnDt                                               = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.TxAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt                                                        = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt                                                    = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt_Ccy                                                = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt_Value                                              = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg                                                = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_SrcCcy                                         = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_TrgtCcy                                        = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_UnitCcy                                        = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_XchgRate                                       = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_CtrctId                                        = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_QtnDt                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.CntrValAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt                                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt                                                 = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt_Ccy                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt_Value                                           = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg                                             = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_SrcCcy                                      = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_TrgtCcy                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_UnitCcy                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_XchgRate                                    = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_CtrctId                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_QtnDt                                       = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.AnncdPstngAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Item                                                     = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[]"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt                                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Tp                                                       = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt                                                      = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt_Ccy                                                  = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt_Value                                                = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg                                                  = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_SrcCcy                                           = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_TrgtCcy                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_UnitCcy                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_XchgRate                                         = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_CtrctId                                          = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_QtnDt                                            = "BkToCstmrStmt.Stmt[].Ntry[].AmtDtls.PrtryAmt[].CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Item                                                                = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs                                                                     = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt                                                   = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].TtlChrgsAndTaxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt_Ccy                                               = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].TtlChrgsAndTaxAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt_Value                                             = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].TtlChrgsAndTaxAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt_Ccy                                                             = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt_Value                                                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_CdtDbtInd                                                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp                                                                  = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Cd                                                               = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry                                                            = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry_Id                                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tp.Prtry.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry_Issr                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tp.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Rate                                                                = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Rate"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Br                                                                  = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Br"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId                                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_BIC                                                  = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId                                          = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd                              = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_MmbId                                    = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Nm                                                   = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr                                              = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrTp                                        = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_Dept                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_SubDept                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_StrtNm                                       = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_BldgNb                                       = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_PstCd                                        = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_TwnNm                                        = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_CtrySubDvsn                                  = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_Ctry                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine_Item                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_Id                                              = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Cd                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Prtry                                   = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_Issr                                            = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId                                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_Id                                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_Nm                                                      = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrTp                                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_Dept                                            = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_SubDept                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_StrtNm                                          = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_BldgNb                                          = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_PstCd                                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_TwnNm                                           = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_CtrySubDvsn                                     = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_Ctry                                            = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrLine_Item                                    = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrLine                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Pty.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Id                                                              = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Rate                                                            = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax.Rate"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt                                                             = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt_Ccy                                                         = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt_Value                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Chrgs[].Tax.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl                                                             = "BkToCstmrStmt.Stmt[].Ntry[].TechInptChanl"
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl_Cd                                                          = "BkToCstmrStmt.Stmt[].Ntry[].TechInptChanl.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl_Prtry                                                       = "BkToCstmrStmt.Stmt[].Ntry[].TechInptChanl.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Item                                                               = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst                                                                    = "BkToCstmrStmt.Stmt[].Ntry[].Intrst" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt                                                                = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt_Ccy                                                            = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt_Value                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_CdtDbtInd                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp                                                                 = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp_Cd                                                              = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp_Prtry                                                           = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Item                                                          = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[]"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate                                                               = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp                                                            = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp_Pctg                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].Tp.Pctg"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp_Othr                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].Tp.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg                                                       = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt                                                   = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt                                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl                                        = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt                                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl                                        = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt                                           = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl                                = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl                                = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_EQAmt                                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.EQAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_NEQAmt                                            = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Amt.NEQAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_CdtDbtInd                                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Ccy                                                   = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rate[].VldtyRg.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt                                                             = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].FrToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt_FrDtTm                                                      = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].FrToDt.FrDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt_ToDtTm                                                      = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].FrToDt.ToDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rsn                                                                = "BkToCstmrStmt.Stmt[].Ntry[].Intrst[].Rsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Item                                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls                                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch                                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_MsgId                                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.MsgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_PmtInfId                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.PmtInfId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_NbOfTxs                                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.NbOfTxs"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt                                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.TtlAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt_Ccy                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.TtlAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt_Value                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.TtlAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_CdtDbtInd                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].Btch.CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Item                                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls                                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs                                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_MsgId                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.MsgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_AcctSvcrRef                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.AcctSvcrRef"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_PmtInfId                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.PmtInfId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_InstrId                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.InstrId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_EndToEndId                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.EndToEndId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_TxId                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.TxId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_MndtId                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.MndtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_ChqNb                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.ChqNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_ClrSysRef                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.ClrSysRef"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry_Tp                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.Prtry.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry_Ref                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Refs.Prtry.Ref"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt_Ccy                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt_Value                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_SrcCcy                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_TrgtCcy                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_UnitCcy                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_XchgRate                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_CtrctId                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_QtnDt                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.InstdAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt_Ccy                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt_Value                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_SrcCcy                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_TrgtCcy                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_UnitCcy                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_XchgRate                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_CtrctId                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_QtnDt                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.TxAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt_Ccy                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt_Value                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_SrcCcy                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_TrgtCcy                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_UnitCcy                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_XchgRate                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_CtrctId                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_QtnDt                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.CntrValAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt_Ccy                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt_Value                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_SrcCcy                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_TrgtCcy                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_UnitCcy                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_XchgRate                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_CtrctId                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_QtnDt                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.AnncdPstngAmt.CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Item                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Tp                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt_Ccy                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt_Value                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_SrcCcy                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.SrcCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_TrgtCcy                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.TrgtCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_UnitCcy                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.UnitCcy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_XchgRate                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.XchgRate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_CtrctId                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.CtrctId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_QtnDt                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AmtDtls.PrtryAmt[].CcyXchg.QtnDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Item                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt_NbOfDays                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Dt.NbOfDays"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt_ActlDt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Dt.ActlDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt_Ccy                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt_Value                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_CdtDbtInd                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Avlbty[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Domn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Cd                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Domn.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Domn.Fmly"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly_Cd                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Domn.Fmly.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly_SubFmlyCd                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Domn.Fmly.SubFmlyCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry_Cd                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Prtry.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry_Issr                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].BkTxCd.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Item                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs                                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].TtlChrgsAndTaxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt_Ccy                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].TtlChrgsAndTaxAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt_Value                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].TtlChrgsAndTaxAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt_Ccy                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt_Value                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_CdtDbtInd                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Cd                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry_Id                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tp.Prtry.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry_Issr                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tp.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Rate                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Rate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Br                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Br"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_BIC                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_MmbId                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Nm                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrTp                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_Dept                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_SubDept                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_StrtNm                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_BldgNb                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_PstCd                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_TwnNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_CtrySubDvsn                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_Ctry                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine_Item                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_Id                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Cd                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Prtry                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_Issr                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_Id                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_Nm                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrTp                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_Dept                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_SubDept                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_StrtNm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_BldgNb                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_PstCd                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_TwnNm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_CtrySubDvsn                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_Ctry                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrLine_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrLine                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Pty.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Id                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Rate                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax.Rate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt_Ccy                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax.Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt_Value                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Chrgs[].Tax.Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Item                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt_Ccy                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt_Value                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_CdtDbtInd                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp_Cd                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp_Prtry                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Item                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp_Pctg                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].Tp.Pctg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp_Othr                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].Tp.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.FrAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.BdryAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.FrToAmt.ToAmt.Incl"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_EQAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.EQAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_NEQAmt                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Amt.NEQAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_CdtDbtInd                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Ccy                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rate[].VldtyRg.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].FrToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt_FrDtTm                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].FrToDt.FrDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt_ToDtTm                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].FrToDt.ToDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rsn                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Intrst[].Rsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Nm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrTp                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_Dept                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_SubDept                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_StrtNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_BldgNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_PstCd                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_TwnNm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_CtrySubDvsn                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_Ctry                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrLine_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrLine                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_BICOrBEI                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Item                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtryOfRes                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_NmPrfx                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_PhneNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_MobNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_FaxNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_EmailAdr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.InitgPty.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Nm                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrTp                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_Dept                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_SubDept                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_StrtNm                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_BldgNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_PstCd                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_TwnNm                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_CtrySubDvsn                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_Ctry                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrLine_Item                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrLine                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_BICOrBEI                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Item                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Id                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm_Cd                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Issr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Item                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Id                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Issr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtryOfRes                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_NmPrfx                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_Nm                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_PhneNb                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_MobNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_FaxNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_EmailAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_Othr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Dbtr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_IBAN                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.IBAN"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_Id                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm_Cd                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm_Prtry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_Issr                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Id.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp_Cd                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp_Prtry                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Ccy                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Nm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.DbtrAcct.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Nm                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrTp                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_Dept                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_SubDept                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_StrtNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_BldgNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_PstCd                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_TwnNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_CtrySubDvsn                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_Ctry                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrLine_Item                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrLine                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_BICOrBEI                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Id                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Issr                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtryOfRes                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_NmPrfx                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_Nm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_PhneNb                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_MobNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_FaxNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_EmailAdr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtDbtr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Nm                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrTp                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_Dept                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_SubDept                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_StrtNm                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_BldgNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_PstCd                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_TwnNm                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_CtrySubDvsn                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_Ctry                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrLine_Item                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrLine                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_BICOrBEI                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Item                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Id                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm_Cd                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Issr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Item                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Id                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Issr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtryOfRes                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_NmPrfx                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_Nm                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_PhneNb                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_MobNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_FaxNb                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_EmailAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_Othr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Cdtr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_IBAN                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.IBAN"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_Id                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm_Cd                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm_Prtry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_Issr                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Id.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp_Cd                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp_Prtry                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Ccy                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Nm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.CdtrAcct.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Nm                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrTp                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_Dept                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_SubDept                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_StrtNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_BldgNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_PstCd                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_TwnNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_CtrySubDvsn                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_Ctry                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrLine_Item                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrLine                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_BICOrBEI                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Id                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Issr                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtryOfRes                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_NmPrfx                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_Nm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_PhneNb                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_MobNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_FaxNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_EmailAdr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.UltmtCdtr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Nm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrTp                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_Dept                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_SubDept                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_StrtNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_BldgNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_PstCd                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_TwnNm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_CtrySubDvsn                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_Ctry                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrLine_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrLine                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_BICOrBEI                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Item                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtryOfRes                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_NmPrfx                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_PhneNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_MobNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_FaxNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_EmailAdr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.TradgPty.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Item                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Tp                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Nm                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrTp                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_Dept                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_SubDept                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_StrtNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_BldgNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_PstCd                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_TwnNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_CtrySubDvsn                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_Ctry                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrLine_Item                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrLine                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_BICOrBEI                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Id                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm_Cd             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm_Prtry          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Issr                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtryOfRes                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_NmPrfx                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_Nm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_PhneNb                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_MobNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_FaxNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_EmailAdr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPties.Prtry[].Pty.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_BIC                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_Id                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_Nm                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrTp                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_Dept                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_SubDept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_StrtNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_BldgNb                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_PstCd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_TwnNm                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_Ctry                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrLine                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DbtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_BIC                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_Id                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_Nm                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrTp                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_Dept                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_SubDept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_StrtNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_BldgNb                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_PstCd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_TwnNm                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_Ctry                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrLine                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.CdtrAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_BIC                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Nm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_Dept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_SubDept            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_PstCd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_Ctry               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_Id                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_Issr                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_Id                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt1.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_BIC                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Nm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_Dept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_SubDept            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_PstCd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_Ctry               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_Id                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_Issr                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_Id                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt2.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_BIC                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Nm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_Dept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_SubDept            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_PstCd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_Ctry               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_Id                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_Issr                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_Id                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IntrmyAgt3.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_BIC                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_MmbId             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_Id                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_Nm                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrTp                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_Dept                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_SubDept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_StrtNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_BldgNb                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_PstCd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_TwnNm                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_CtrySubDvsn              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_Ctry                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrLine_Item             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrLine                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.RcvgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_BIC                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_MmbId            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Nm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrTp                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_Dept                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_SubDept              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_StrtNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_BldgNb               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_PstCd                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_TwnNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_CtrySubDvsn          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_Ctry                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrLine_Item         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrLine              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_Id                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_Nm                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrTp                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_Dept                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_SubDept                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_StrtNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_BldgNb                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_PstCd                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_TwnNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_CtrySubDvsn             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_Ctry                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrLine_Item            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrLine                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.DlvrgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_BIC                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_MmbId             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrTp                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_Dept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_SubDept               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_StrtNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_BldgNb                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_PstCd                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_TwnNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_CtrySubDvsn           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_Ctry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrLine_Item          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrLine               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_Id                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_Nm                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrTp                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_Dept                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_SubDept                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_StrtNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_BldgNb                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_PstCd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_TwnNm                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_CtrySubDvsn              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_Ctry                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrLine_Item             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrLine                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.IssgAgt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_BIC                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId_Cd      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId_Prtry   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_MmbId            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Nm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrTp                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_Dept                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_SubDept              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_StrtNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_BldgNb               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_PstCd                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_TwnNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_CtrySubDvsn          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_Ctry                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrLine_Item         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrLine              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_Id                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_Nm                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrTp                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_Dept                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_SubDept                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_StrtNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_BldgNb                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_PstCd                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_TwnNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_CtrySubDvsn             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_Ctry                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrLine_Item            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrLine                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.SttlmPlc.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Item                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Tp                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_BIC                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.BIC"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.ClrSysMmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.ClrSysMmbId.ClrSysId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.ClrSysMmbId.ClrSysId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_MmbId           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.ClrSysMmbId.MmbId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Nm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrTp               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_Dept                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_SubDept             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_StrtNm              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_BldgNb              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_PstCd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_TwnNm               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_CtrySubDvsn         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_Ctry                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrLine_Item        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrLine             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_Id                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm_Cd             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm_Prtry          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_Issr                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.FinInstnId.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_Id                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_Nm                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrTp                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_Dept                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_SubDept                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_StrtNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_BldgNb                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_PstCd                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_TwnNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_CtrySubDvsn            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_Ctry                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrLine_Item           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrLine                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdAgts.Prtry[].Agt.BrnchId.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp                                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Purp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp_Cd                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Purp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp_Prtry                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Purp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_Item                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtId                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnMtd                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnMtd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnElctrncAdr                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnElctrncAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Nm                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdRmtInf[].RmtLctnPstlAdr.Adr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Ustrd_Item                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Ustrd[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Ustrd                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Ustrd" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Item                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Item                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Tp.CdOrPrtry.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_Issr                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Tp.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Nb                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].Nb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_RltdDt                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocInf[].RltdDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DuePyblAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.DscntApldAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.CdtNoteAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.TaxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.TaxAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].CdtDbtInd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].Rsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.AdjstmntAmtAndRsn[].AddtlInf"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].RfrdDocAmt.RmtdAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Tp.CdOrPrtry.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_Issr                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Tp.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Ref                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].CdtrRefInf.Ref"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Nm                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrTp                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_Dept                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_SubDept                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_StrtNm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_BldgNb                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_PstCd                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_TwnNm                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_Ctry                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrLine                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtryOfRes                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_NmPrfx                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_Nm                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_PhneNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_MobNb                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_FaxNb                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_EmailAdr                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_Othr                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Nm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrTp                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_Dept                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_SubDept                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_StrtNm                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_BldgNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_PstCd                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_TwnNm                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_Ctry                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrLine                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtryOfRes                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_NmPrfx                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_Nm                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_PhneNb                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_MobNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_FaxNb                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_EmailAdr                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_Othr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].Invcee.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_AddtlRmtInf_Item                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].AddtlRmtInf[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_AddtlRmtInf                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RmtInf.Strd[].AddtlRmtInf" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_AccptncDtTm                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.AccptncDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TradActvtyCtrctlSttlmDt                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.TradActvtyCtrctlSttlmDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TradDt                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.TradDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_IntrBkSttlmDt                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.IntrBkSttlmDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_StartDt                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.StartDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_EndDt                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.EndDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TxDtTm                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.TxDtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Item                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Tp                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry[].Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt_Dt                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry[].Dt.Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt_DtTm                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdDts.Prtry[].Dt.DtTm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.DealPric"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric_Ccy                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.DealPric.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric_Value                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.DealPric.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Item                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Tp                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry[].Pric"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric_Ccy                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry[].Pric.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric_Value                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdPric.Prtry[].Pric.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Item                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Qty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_Unit                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Qty.Unit"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_FaceAmt                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Qty.FaceAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_AmtsdVal                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Qty.AmtsdVal"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry_Tp                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Prtry.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry_Qty                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RltdQties[].Prtry.Qty"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].FinInstrmId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_ISIN                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].FinInstrmId.ISIN"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].FinInstrmId.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry_Tp                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].FinInstrmId.Prtry.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry_Id                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].FinInstrmId.Prtry.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax                                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Cdtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_TaxId                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Cdtr.TaxId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_RegnId                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Cdtr.RegnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_TaxTp                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Cdtr.TaxTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_TaxId                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.TaxId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_RegnId                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.RegnId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_TaxTp                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.TaxTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.Authstn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn_Titl                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.Authstn.Titl"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn_Nm                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dbtr.Authstn.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_AdmstnZn                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.AdmstnZn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_RefNb                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.RefNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Mtd                                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Mtd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxblBaseAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt_Ccy                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxblBaseAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt_Value                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxblBaseAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt_Ccy                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt_Value                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.TtlTaxAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dt                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Dt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_SeqNb                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.SeqNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Item                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Tp                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Ctgy                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Ctgy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_CtgyDtls                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].CtgyDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_DbtrSts                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].DbtrSts"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_CertId                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].CertId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_FrmsCd                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].FrmsCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_Yr                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd.Yr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_Tp                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt                                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd.FrToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt_FrDt                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd.FrToDt.FrDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt_ToDt                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].Prd.FrToDt.ToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Rate                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Rate"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TaxblBaseAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TtlAmt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt_Ccy                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TtlAmt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt_Value                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.TtlAmt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Item                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Yr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.FrDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Prd.FrToDt.ToDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Amt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt_Value                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].TaxAmt.Dtls[].Amt.Value"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_AddtlInf                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].Tax.Rcrd[].AddtlInf"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf                                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn                                   = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Domn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Cd                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Domn.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Domn.Fmly"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly_Cd                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Domn.Fmly.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly_SubFmlyCd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Domn.Fmly.SubFmlyCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry_Cd                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Prtry.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry_Issr                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.OrgnlBkTxCd.Prtry.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Nm                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrTp                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.AdrTp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_Dept                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.Dept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_SubDept                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.SubDept"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_StrtNm                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.StrtNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_BldgNb                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.BldgNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_PstCd                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.PstCd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_TwnNm                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.TwnNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_CtrySubDvsn                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.CtrySubDvsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_Ctry                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.Ctry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrLine_Item                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.AdrLine[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrLine                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.PstlAdr.AdrLine" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_BICOrBEI                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.BICOrBEI"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Item                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Id                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Issr                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.OrgId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.DtAndPlcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Item                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Id                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[].Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm                       = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[].SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[].SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[].SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Issr                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.Id.PrvtId.Othr[].Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtryOfRes                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtryOfRes"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls                                     = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_NmPrfx                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.NmPrfx"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_Nm                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_PhneNb                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.PhneNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_MobNb                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.MobNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_FaxNb                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.FaxNb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_EmailAdr                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.EmailAdr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_Othr                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Orgtr.CtctDtls.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Rsn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn_Cd                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Rsn.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn_Prtry                                          = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.Rsn.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_AddtlInf_Item                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.AddtlInf[]"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_AddtlInf                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].RtrInf.AddtlInf" // ARRAY
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn                                                  = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].CorpActn"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Cd                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].CorpActn.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Nb                                               = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].CorpActn.Nb"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Prtry                                            = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].CorpActn.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct                                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_IBAN                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.IBAN"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr                                         = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_Id                                      = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr.Id"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm                                 = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr.SchmeNm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm_Cd                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr.SchmeNm.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm_Prtry                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr.SchmeNm.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_Issr                                    = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Id.Othr.Issr"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Tp"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp_Cd                                           = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Tp.Cd"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp_Prtry                                        = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Tp.Prtry"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Ccy                                             = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Ccy"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Nm                                              = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].SfkpgAcct.Nm"
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AddtlTxInf                                                = "BkToCstmrStmt.Stmt[].Ntry[].NtryDtls[].TxDtls[].AddtlTxInf"
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlNtryInf                                                              = "BkToCstmrStmt.Stmt[].Ntry[].AddtlNtryInf"
	Path_BkToCstmrStmt_Stmt_AddtlStmtInf                                                                   = "BkToCstmrStmt.Stmt[].AddtlStmtInf"
)

const (
	PathTypeProperty  = "property"
	PathTypeStruct    = "struct"
	PathTypeArray     = "array"
	PathTypeArrayItem = "array-item"
)

var nodeRegistryTypes = map[string]string{
	Path_BkToCstmrStmt:                                                                                     PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr:                                                                              PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgId:                                                                        PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_CreDtTm:                                                                      PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt:                                                                      PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Nm:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr:                                                              PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrTp:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_Dept:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_SubDept:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_StrtNm:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_BldgNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_PstCd:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_TwnNm:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_CtrySubDvsn:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_Ctry:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrLine_Item:                                                 PathTypeArrayItem,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_PstlAdr_AdrLine:                                                      PathTypeArray,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_BICOrBEI:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Item:                                                   PathTypeArrayItem,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr:                                                        PathTypeArray,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Id:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_OrgId_Othr_Issr:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth:                                            PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                    PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Item:                                                  PathTypeArrayItem,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr:                                                       PathTypeArray,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Id:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm:                                               PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm_Cd:                                            PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_SchmeNm_Prtry:                                         PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_Id_PrvtId_Othr_Issr:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtryOfRes:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_NmPrfx:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_Nm:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_PhneNb:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_MobNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_FaxNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_EmailAdr:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgRcpt_CtctDtls_Othr:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn:                                                                     PathTypeStruct,
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn_PgNb:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_MsgPgntn_LastPgInd:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_GrpHdr_AddtlInf:                                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Item:                                                                           PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt:                                                                                PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Id:                                                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_ElctrncSeqNb:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_LglSeqNb:                                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_CreDtTm:                                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_FrToDt:                                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_FrToDt_FrDtTm:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_FrToDt_ToDtTm:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_CpyDplctInd:                                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RptgSrc:                                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RptgSrc_Cd:                                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RptgSrc_Prtry:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct:                                                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Id:                                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Id_IBAN:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_Id:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm:                                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm_Cd:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_SchmeNm_Prtry:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Id_Othr_Issr:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Tp:                                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Tp_Cd:                                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Tp_Prtry:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ccy:                                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Nm:                                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr:                                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Nm:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr:                                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrTp:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_Dept:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_SubDept:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_StrtNm:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_BldgNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_PstCd:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_TwnNm:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_CtrySubDvsn:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_Ctry:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrLine_Item:                                                 PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_PstlAdr_AdrLine:                                                      PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_BICOrBEI:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Item:                                                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr:                                                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Id:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm_Cd:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_SchmeNm_Prtry:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_OrgId_Othr_Issr:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Item:                                                  PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr:                                                       PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Id:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm:                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm_Cd:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_SchmeNm_Prtry:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_Id_PrvtId_Othr_Issr:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtryOfRes:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_NmPrfx:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_Nm:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_PhneNb:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_MobNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_FaxNb:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_EmailAdr:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Ownr_CtctDtls_Othr:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr:                                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId:                                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_BIC:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId:                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_ClrSysMmbId_MmbId:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Nm:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr:                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrTp:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_Dept:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_SubDept:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_StrtNm:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_BldgNb:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_PstCd:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_TwnNm:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_CtrySubDvsn:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_Ctry:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrLine_Item:                                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_PstlAdr_AdrLine:                                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_Id:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm_Cd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_SchmeNm_Prtry:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_FinInstnId_Othr_Issr:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId:                                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_Id:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_Nm:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrTp:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_Dept:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_SubDept:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_StrtNm:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_BldgNb:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_PstCd:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_TwnNm:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_CtrySubDvsn:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_Ctry:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrLine_Item:                                         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Acct_Svcr_BrnchId_PstlAdr_AdrLine:                                              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_RltdAcct:                                                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id:                                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_IBAN:                                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr:                                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_Id:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm:                                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm_Cd:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_SchmeNm_Prtry:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Id_Othr_Issr:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp:                                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp_Cd:                                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Tp_Prtry:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Ccy:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_RltdAcct_Nm:                                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Item:                                                                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Intrst:                                                                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Intrst_Tp:                                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Tp_Cd:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Tp_Prtry:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Item:                                                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate:                                                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp_Pctg:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_Tp_Othr:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt:                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_EQAmt:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Amt_NEQAmt:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_CdtDbtInd:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rate_VldtyRg_Ccy:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt:                                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt_FrDtTm:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_FrToDt_ToDtTm:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Intrst_Rsn:                                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Item:                                                                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Bal:                                                                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Bal_Tp:                                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry:                                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry_Cd:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_CdOrPrtry_Prtry:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp_Cd:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Tp_SubTp_Prtry:                                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine:                                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Incl:                                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt:                                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt_Ccy:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_CdtLine_Amt_Value:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Amt:                                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Amt_Ccy:                                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Amt_Value:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_CdtDbtInd:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Dt:                                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Dt_Dt:                                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Dt_DtTm:                                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Item:                                                                PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty:                                                                     PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt:                                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt_NbOfDays:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Dt_ActlDt:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt_Ccy:                                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_Amt_Value:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Bal_Avlbty_CdtDbtInd:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry:                                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_NbOfNtries:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_Sum:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_TtlNetNtryAmt:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtries_CdtDbtInd:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries:                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries_NbOfNtries:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlCdtNtries_Sum:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries:                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries_NbOfNtries:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlDbtNtries_Sum:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Item:                                              PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd:                                                   PathTypeArray,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_NbOfNtries:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Sum:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_TtlNetNtryAmt:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_CdtDbtInd:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_FcstInd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Cd:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly_Cd:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Domn_Fmly_SubFmlyCd:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry_Cd:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_BkTxCd_Prtry_Issr:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Item:                                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty:                                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt_NbOfDays:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Dt_ActlDt:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt_Ccy:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_Amt_Value:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_TxsSummry_TtlNtriesPerBkTxCd_Avlbty_CdtDbtInd:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Item:                                                                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry:                                                                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryRef:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Amt:                                                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Amt_Ccy:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Amt_Value:                                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_CdtDbtInd:                                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_RvslInd:                                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Sts:                                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt_Dt:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BookgDt_DtTm:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt:                                                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt_Dt:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_ValDt_DtTm:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AcctSvcrRef:                                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Item:                                                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty:                                                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt_NbOfDays:                                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Dt_ActlDt:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt:                                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt_Ccy:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_Amt_Value:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Avlbty_CdtDbtInd:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd:                                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn:                                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Cd:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly:                                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly_Cd:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Domn_Fmly_SubFmlyCd:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry:                                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry_Cd:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_BkTxCd_Prtry_Issr:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_ComssnWvrInd:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd:                                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd_MsgNmId:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlInfInd_MsgId:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls:                                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt:                                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt_Ccy:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_Amt_Value:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_SrcCcy:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_TrgtCcy:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_UnitCcy:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_XchgRate:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_CtrctId:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_InstdAmt_CcyXchg_QtnDt:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt:                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt_Ccy:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_Amt_Value:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg:                                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_SrcCcy:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_TrgtCcy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_UnitCcy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_XchgRate:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_CtrctId:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_TxAmt_CcyXchg_QtnDt:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt:                                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt:                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt_Ccy:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_Amt_Value:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_SrcCcy:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_TrgtCcy:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_UnitCcy:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_XchgRate:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_CtrctId:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_CntrValAmt_CcyXchg_QtnDt:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt:                                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt_Ccy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_Amt_Value:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_SrcCcy:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_TrgtCcy:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_UnitCcy:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_XchgRate:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_CtrctId:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_AnncdPstngAmt_CcyXchg_QtnDt:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Item:                                                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt:                                                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Tp:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt_Ccy:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_Amt_Value:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_SrcCcy:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_TrgtCcy:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_UnitCcy:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_XchgRate:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_CtrctId:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AmtDtls_PrtryAmt_CcyXchg_QtnDt:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Item:                                                                PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs:                                                                     PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt:                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt_Ccy:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_TtlChrgsAndTaxAmt_Value:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt_Ccy:                                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Amt_Value:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_CdtDbtInd:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp:                                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Cd:                                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry_Id:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tp_Prtry_Issr:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Rate:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Br:                                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_BIC:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_ClrSysMmbId_MmbId:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Nm:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrTp:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_Dept:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_SubDept:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_StrtNm:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_BldgNb:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_PstCd:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_TwnNm:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_CtrySubDvsn:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_Ctry:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine_Item:                                 PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine:                                      PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_Id:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Cd:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Prtry:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_FinInstnId_Othr_Issr:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId:                                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_Id:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_Nm:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrTp:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_Dept:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_SubDept:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_StrtNm:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_BldgNb:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_PstCd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_TwnNm:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_CtrySubDvsn:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_Ctry:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrLine_Item:                                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Pty_BrnchId_PstlAdr_AdrLine:                                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Id:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Rate:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt_Ccy:                                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Chrgs_Tax_Amt_Value:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl_Cd:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_TechInptChanl_Prtry:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Item:                                                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst:                                                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt:                                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt_Ccy:                                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Amt_Value:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_CdtDbtInd:                                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp:                                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp_Cd:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Tp_Prtry:                                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Item:                                                          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate:                                                               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp:                                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp_Pctg:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_Tp_Othr:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg:                                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt:                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt:                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_EQAmt:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Amt_NEQAmt:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_CdtDbtInd:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rate_VldtyRg_Ccy:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt_FrDtTm:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_FrToDt_ToDtTm:                                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_Intrst_Rsn:                                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Item:                                                             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls:                                                                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch:                                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_MsgId:                                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_PmtInfId:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_NbOfTxs:                                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt_Ccy:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_TtlAmt_Value:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_Btch_CdtDbtInd:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Item:                                                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls:                                                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_MsgId:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_AcctSvcrRef:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_PmtInfId:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_InstrId:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_EndToEndId:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_TxId:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_MndtId:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_ChqNb:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_ClrSysRef:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry_Tp:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Refs_Prtry_Ref:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls:                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt_Ccy:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_Amt_Value:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_SrcCcy:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_TrgtCcy:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_UnitCcy:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_XchgRate:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_CtrctId:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_InstdAmt_CcyXchg_QtnDt:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt_Ccy:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_Amt_Value:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_SrcCcy:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_TrgtCcy:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_UnitCcy:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_XchgRate:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_CtrctId:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_TxAmt_CcyXchg_QtnDt:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt_Ccy:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_Amt_Value:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_SrcCcy:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_TrgtCcy:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_UnitCcy:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_XchgRate:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_CtrctId:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_CntrValAmt_CcyXchg_QtnDt:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt_Ccy:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_Amt_Value:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_SrcCcy:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_TrgtCcy:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_UnitCcy:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_XchgRate:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_CtrctId:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_AnncdPstngAmt_CcyXchg_QtnDt:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Item:                                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt:                                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Tp:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt_Ccy:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_Amt_Value:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_SrcCcy:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_TrgtCcy:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_UnitCcy:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_XchgRate:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_CtrctId:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AmtDtls_PrtryAmt_CcyXchg_QtnDt:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Item:                                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty:                                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt_NbOfDays:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Dt_ActlDt:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt_Ccy:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_Amt_Value:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Avlbty_CdtDbtInd:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd:                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn:                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Cd:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly_Cd:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Domn_Fmly_SubFmlyCd:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry_Cd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_BkTxCd_Prtry_Issr:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Item:                                                PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs:                                                     PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt_Ccy:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_TtlChrgsAndTaxAmt_Value:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt_Ccy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Amt_Value:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_CdtDbtInd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Cd:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry_Id:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tp_Prtry_Issr:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Rate:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Br:                                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_BIC:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_ClrSysMmbId_MmbId:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Nm:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrTp:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_Dept:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_SubDept:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_StrtNm:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_BldgNb:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_PstCd:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_TwnNm:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_CtrySubDvsn:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_Ctry:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine_Item:                 PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_PstlAdr_AdrLine:                      PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_Id:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Cd:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_SchmeNm_Prtry:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_FinInstnId_Othr_Issr:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_Id:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_Nm:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrTp:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_Dept:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_SubDept:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_StrtNm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_BldgNb:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_PstCd:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_TwnNm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_CtrySubDvsn:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_Ctry:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrLine_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Pty_BrnchId_PstlAdr_AdrLine:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Id:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Rate:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt_Ccy:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Chrgs_Tax_Amt_Value:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Item:                                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst:                                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt_Ccy:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Amt_Value:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_CdtDbtInd:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp_Cd:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Tp_Prtry:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Item:                                          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate:                                               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp_Pctg:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_Tp_Othr:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt_BdryAmt:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrAmt_Incl:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt_BdryAmt:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_ToAmt_Incl:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt:                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt:                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_BdryAmt:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_FrAmt_Incl:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt:                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_BdryAmt:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_FrToAmt_ToAmt_Incl:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_EQAmt:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Amt_NEQAmt:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_CdtDbtInd:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rate_VldtyRg_Ccy:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt_FrDtTm:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_FrToDt_ToDtTm:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Intrst_Rsn:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Nm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrTp:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_Dept:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_SubDept:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_StrtNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_BldgNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_PstCd:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_TwnNm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_CtrySubDvsn:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_Ctry:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrLine_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_PstlAdr_AdrLine:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_BICOrBEI:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Item:                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr:                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_OrgId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth:              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_Id_PrvtId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtryOfRes:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_NmPrfx:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_PhneNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_MobNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_FaxNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_EmailAdr:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_InitgPty_CtctDtls_Othr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Nm:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrTp:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_Dept:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_SubDept:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_StrtNm:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_BldgNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_PstCd:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_TwnNm:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_CtrySubDvsn:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_Ctry:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrLine_Item:                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_PstlAdr_AdrLine:                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_BICOrBEI:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Item:                         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr:                              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Id:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm:                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm_Cd:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_SchmeNm_Prtry:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_OrgId_Othr_Issr:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Item:                        PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr:                             PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Id:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm:                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_SchmeNm_Prtry:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_Id_PrvtId_Othr_Issr:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtryOfRes:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_NmPrfx:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_Nm:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_PhneNb:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_MobNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_FaxNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_EmailAdr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Dbtr_CtctDtls_Othr:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_IBAN:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_Id:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm:                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm_Cd:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_SchmeNm_Prtry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Id_Othr_Issr:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp_Cd:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Tp_Prtry:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Ccy:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_DbtrAcct_Nm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Nm:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_Dept:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_SubDept:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_PstCd:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_Ctry:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_PstlAdr_AdrLine:                       PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtryOfRes:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_Nm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_MobNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtDbtr_CtctDtls_Othr:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr:                                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Nm:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrTp:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_Dept:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_SubDept:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_StrtNm:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_BldgNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_PstCd:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_TwnNm:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_CtrySubDvsn:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_Ctry:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrLine_Item:                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_PstlAdr_AdrLine:                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_BICOrBEI:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Item:                         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr:                              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Id:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm:                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm_Cd:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_SchmeNm_Prtry:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_OrgId_Othr_Issr:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Item:                        PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr:                             PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Id:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm:                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm_Cd:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_SchmeNm_Prtry:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_Id_PrvtId_Othr_Issr:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtryOfRes:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_NmPrfx:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_Nm:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_PhneNb:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_MobNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_FaxNb:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_EmailAdr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Cdtr_CtctDtls_Othr:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_IBAN:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_Id:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm:                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm_Cd:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_SchmeNm_Prtry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Id_Othr_Issr:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp_Cd:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Tp_Prtry:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Ccy:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_CdtrAcct_Nm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Nm:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_Dept:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_SubDept:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_PstCd:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_Ctry:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_PstlAdr_AdrLine:                       PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtryOfRes:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_Nm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_MobNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_UltmtCdtr_CtctDtls_Othr:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Nm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrTp:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_Dept:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_SubDept:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_StrtNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_BldgNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_PstCd:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_TwnNm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_CtrySubDvsn:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_Ctry:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrLine_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_PstlAdr_AdrLine:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_BICOrBEI:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Item:                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr:                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_OrgId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth:              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_Id_PrvtId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtryOfRes:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_NmPrfx:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_PhneNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_MobNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_FaxNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_EmailAdr:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_TradgPty_CtctDtls_Othr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Item:                                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry:                                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Tp:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Nm:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrTp:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_Dept:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_SubDept:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_StrtNm:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_BldgNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_PstCd:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_TwnNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_CtrySubDvsn:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_Ctry:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrLine_Item:                  PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_PstlAdr_AdrLine:                       PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_BICOrBEI:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_OrgId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth:             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_BirthDt:     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Id:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_Id_PrvtId_Othr_Issr:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtryOfRes:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_NmPrfx:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_Nm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_PhneNb:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_MobNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_FaxNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_EmailAdr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPties_Prtry_Pty_CtctDtls_Othr:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_BIC:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId:                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_ClrSysMmbId_MmbId:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_FinInstnId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_Id:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_Nm:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrTp:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_Dept:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_SubDept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_StrtNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_BldgNb:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_PstCd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_TwnNm:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_CtrySubDvsn:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_Ctry:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrLine_Item:             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DbtrAgt_BrnchId_PstlAdr_AdrLine:                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_BIC:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId:                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId:          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_ClrSysMmbId_MmbId:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_FinInstnId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_Id:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_Nm:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrTp:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_Dept:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_SubDept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_StrtNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_BldgNb:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_PstCd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_TwnNm:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_CtrySubDvsn:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_Ctry:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrLine_Item:             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_CdtrAgt_BrnchId_PstlAdr_AdrLine:                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId:                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_BIC:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId:       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Cd:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_ClrSysId_Prtry: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_ClrSysMmbId_MmbId:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Nm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr:                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrTp:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_Dept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_SubDept:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_StrtNm:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_BldgNb:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_PstCd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_TwnNm:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_CtrySubDvsn:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_Ctry:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine_Item:       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_PstlAdr_AdrLine:            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_Id:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm:               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Cd:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_SchmeNm_Prtry:         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_FinInstnId_Othr_Issr:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_Id:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt1_BrnchId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId:                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_BIC:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId:       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Cd:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_ClrSysId_Prtry: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_ClrSysMmbId_MmbId:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Nm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr:                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrTp:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_Dept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_SubDept:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_StrtNm:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_BldgNb:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_PstCd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_TwnNm:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_CtrySubDvsn:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_Ctry:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine_Item:       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_PstlAdr_AdrLine:            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_Id:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm:               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Cd:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_SchmeNm_Prtry:         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_FinInstnId_Othr_Issr:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_Id:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt2_BrnchId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId:                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_BIC:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId:       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Cd:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_ClrSysId_Prtry: PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_ClrSysMmbId_MmbId:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Nm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr:                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrTp:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_Dept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_SubDept:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_StrtNm:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_BldgNb:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_PstCd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_TwnNm:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_CtrySubDvsn:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_Ctry:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine_Item:       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_PstlAdr_AdrLine:            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_Id:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm:               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Cd:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_SchmeNm_Prtry:         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_FinInstnId_Othr_Issr:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_Id:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IntrmyAgt3_BrnchId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_BIC:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId:                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId:          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_ClrSysMmbId_MmbId:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_FinInstnId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_Id:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_Nm:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrTp:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_Dept:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_SubDept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_StrtNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_BldgNb:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_PstCd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_TwnNm:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_CtrySubDvsn:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_Ctry:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrLine_Item:             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_RcvgAgt_BrnchId_PstlAdr_AdrLine:                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_BIC:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId:         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_ClrSysMmbId_MmbId:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Nm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr:                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrTp:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_Dept:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_SubDept:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_StrtNm:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_BldgNb:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_PstCd:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_TwnNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_CtrySubDvsn:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_Ctry:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrLine_Item:         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_PstlAdr_AdrLine:              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_FinInstnId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_Id:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_Nm:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrTp:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_Dept:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_SubDept:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_StrtNm:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_BldgNb:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_PstCd:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_TwnNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_CtrySubDvsn:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_Ctry:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrLine_Item:            PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_DlvrgAgt_BrnchId_PstlAdr_AdrLine:                 PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_BIC:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId:                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId:          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_ClrSysMmbId_MmbId:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrTp:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_Dept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_SubDept:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_StrtNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_BldgNb:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_PstCd:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_TwnNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_CtrySubDvsn:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_Ctry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrLine_Item:          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_PstlAdr_AdrLine:               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_FinInstnId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_Id:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_Nm:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr:                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrTp:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_Dept:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_SubDept:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_StrtNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_BldgNb:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_PstCd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_TwnNm:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_CtrySubDvsn:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_Ctry:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrLine_Item:             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_IssgAgt_BrnchId_PstlAdr_AdrLine:                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_BIC:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId:         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId_Cd:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_ClrSysMmbId_MmbId:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Nm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr:                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrTp:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_Dept:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_SubDept:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_StrtNm:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_BldgNb:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_PstCd:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_TwnNm:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_CtrySubDvsn:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_Ctry:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrLine_Item:         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_PstlAdr_AdrLine:              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_FinInstnId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_Id:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_Nm:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrTp:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_Dept:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_SubDept:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_StrtNm:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_BldgNb:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_PstCd:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_TwnNm:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_CtrySubDvsn:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_Ctry:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrLine_Item:            PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_SttlmPlc_BrnchId_PstlAdr_AdrLine:                 PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Item:                                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry:                                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Tp:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_BIC:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId:        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Cd:     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_ClrSysId_Prtry:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_ClrSysMmbId_MmbId:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Nm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr:                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrTp:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_Dept:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_SubDept:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_StrtNm:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_BldgNb:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_PstCd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_TwnNm:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_CtrySubDvsn:         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_Ctry:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrLine_Item:        PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_PstlAdr_AdrLine:             PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr:                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_Id:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm:                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm_Cd:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_SchmeNm_Prtry:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_FinInstnId_Othr_Issr:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_Id:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_Nm:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr:                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrTp:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_Dept:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_SubDept:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_StrtNm:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_BldgNb:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_PstCd:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_TwnNm:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_CtrySubDvsn:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_Ctry:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrLine_Item:           PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdAgts_Prtry_Agt_BrnchId_PstlAdr_AdrLine:                PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp:                                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp_Cd:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Purp_Prtry:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_Item:                                           PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf:                                                PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtId:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnMtd:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnElctrncAdr:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Nm:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrTp:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_Dept:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_SubDept:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_StrtNm:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_BldgNb:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_PstCd:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_TwnNm:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_CtrySubDvsn:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_Ctry:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine_Item:                PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdRmtInf_RmtLctnPstlAdr_Adr_AdrLine:                     PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf:                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Ustrd_Item:                                         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Ustrd:                                              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Item:                                          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd:                                               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Item:                               PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf:                                    PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Cd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_CdOrPrtry_Prtry:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Tp_Issr:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_Nb:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocInf_RltdDt:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Ccy:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DuePyblAmt_Value:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Ccy:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_DscntApldAmt_Value:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt:                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Ccy:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_CdtNoteAmt_Value:                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt:                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt_Ccy:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_TaxAmt_Value:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Item:             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn:                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt:              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Ccy:          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Amt_Value:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_CdtDbtInd:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_Rsn:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_AdjstmntAmtAndRsn_AddtlInf:         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt:                            PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Ccy:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_RfrdDocAmt_RmtdAmt_Value:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Cd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_CdOrPrtry_Prtry:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Tp_Issr:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_CdtrRefInf_Ref:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Nm:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrTp:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_Dept:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_SubDept:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_StrtNm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_BldgNb:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_PstCd:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_TwnNm:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_CtrySubDvsn:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_Ctry:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrLine_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_PstlAdr_AdrLine:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_BICOrBEI:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Item:                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr:                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Id:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm:                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Cd:                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_SchmeNm_Prtry:             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_OrgId_Othr_Issr:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth:               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Item:                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr:                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_Id_PrvtId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtryOfRes:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_NmPrfx:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_Nm:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_PhneNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_MobNb:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_FaxNb:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_EmailAdr:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcr_CtctDtls_Othr:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Nm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr:                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrTp:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_Dept:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_SubDept:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_StrtNm:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_BldgNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_PstCd:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_TwnNm:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_CtrySubDvsn:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_Ctry:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrLine_Item:                   PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_PstlAdr_AdrLine:                        PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_BICOrBEI:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Item:                     PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr:                          PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Id:                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm:                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Cd:               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_SchmeNm_Prtry:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_OrgId_Othr_Issr:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth:              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_BirthDt:      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Item:                    PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr:                         PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Id:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm:                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Cd:              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_SchmeNm_Prtry:           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_Id_PrvtId_Othr_Issr:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtryOfRes:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls:                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_NmPrfx:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_Nm:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_PhneNb:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_MobNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_FaxNb:                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_EmailAdr:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_Invcee_CtctDtls_Othr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_AddtlRmtInf_Item:                              PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RmtInf_Strd_AddtlRmtInf:                                   PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts:                                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_AccptncDtTm:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TradActvtyCtrctlSttlmDt:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TradDt:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_IntrBkSttlmDt:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_StartDt:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_EndDt:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_TxDtTm:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Item:                                        PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry:                                             PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Tp:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt_Dt:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdDts_Prtry_Dt_DtTm:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric_Ccy:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_DealPric_Value:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Item:                                       PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry:                                            PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Tp:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric_Ccy:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdPric_Prtry_Pric_Value:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Item:                                            PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties:                                                 PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_Unit:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_FaceAmt:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Qty_AmtsdVal:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry:                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry_Tp:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RltdQties_Prtry_Qty:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId:                                               PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_ISIN:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry_Tp:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_FinInstrmId_Prtry_Id:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax:                                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_TaxId:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_RegnId:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Cdtr_TaxTp:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_TaxId:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_RegnId:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_TaxTp:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn:                                          PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn_Titl:                                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dbtr_Authstn_Nm:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_AdmstnZn:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_RefNb:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Mtd:                                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt_Ccy:                                   PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxblBaseAmt_Value:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt:                                             PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt_Ccy:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_TtlTaxAmt_Value:                                       PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Dt:                                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_SeqNb:                                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Item:                                             PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd:                                                  PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Tp:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Ctgy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_CtgyDtls:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_DbtrSts:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_CertId:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_FrmsCd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_Yr:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_Tp:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt:                                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt_FrDt:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_Prd_FrToDt_ToDt:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt:                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Rate:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Ccy:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TaxblBaseAmt_Value:                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt_Ccy:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_TtlAmt_Value:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Item:                                 PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls:                                      PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_Yr:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_Tp:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt:                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_FrDt:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Prd_FrToDt_ToDt:                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt_Ccy:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_TaxAmt_Dtls_Amt_Value:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_Tax_Rcrd_AddtlInf:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf:                                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd:                                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn:                                   PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Cd:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly:                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly_Cd:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Domn_Fmly_SubFmlyCd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry:                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry_Cd:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_OrgnlBkTxCd_Prtry_Issr:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Nm:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr:                                      PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrTp:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_Dept:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_SubDept:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_StrtNm:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_BldgNb:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_PstCd:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_TwnNm:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_CtrySubDvsn:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_Ctry:                                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrLine_Item:                         PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_PstlAdr_AdrLine:                              PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id:                                           PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_BICOrBEI:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Item:                           PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr:                                PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Id:                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm:                        PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd:                     PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_SchmeNm_Prtry:                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_OrgId_Othr_Issr:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId:                                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth:                    PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_BirthDt:            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_PrvcOfBirth:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CityOfBirth:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_DtAndPlcOfBirth_CtryOfBirth:        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Item:                          PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr:                               PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Id:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm:                       PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Cd:                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_SchmeNm_Prtry:                 PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_Id_PrvtId_Othr_Issr:                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtryOfRes:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls:                                     PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_NmPrfx:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_Nm:                                  PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_PhneNb:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_MobNb:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_FaxNb:                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_EmailAdr:                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Orgtr_CtctDtls_Othr:                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn:                                                PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn_Cd:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_Rsn_Prtry:                                          PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_AddtlInf_Item:                                      PathTypeArrayItem,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_RtrInf_AddtlInf:                                           PathTypeArray,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn:                                                  PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Cd:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Nb:                                               PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_CorpActn_Prtry:                                            PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct:                                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_IBAN:                                         PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr:                                         PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_Id:                                      PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm:                                 PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm_Cd:                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_SchmeNm_Prtry:                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Id_Othr_Issr:                                    PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp:                                              PathTypeStruct,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp_Cd:                                           PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Tp_Prtry:                                        PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Ccy:                                             PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_SfkpgAcct_Nm:                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_NtryDtls_TxDtls_AddtlTxInf:                                                PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_Ntry_AddtlNtryInf:                                                              PathTypeProperty,
	Path_BkToCstmrStmt_Stmt_AddtlStmtInf:                                                                   PathTypeProperty,
}

func PathType(p string) (string, error) {
	t, ok := nodeRegistryTypes[p]
	if !ok {
		return "", fmt.Errorf("the path %s cannot be recognized as a valid path in camt_053_001_02", p)
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
