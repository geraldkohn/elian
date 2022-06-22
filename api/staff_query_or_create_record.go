package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc/grpclog"
)

func staffQueryOrCreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "POST" {
		var req staffCreateRecordReq
		var res staffCreateRecordRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("staff_create_record encode error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Msg = "参数错误"
			res.StatusAndMsgRes.Status = 400
			return
		}

		c := pb.NewRecordServiceClient(connRecord)
		ctx := context.Background()

		rpcReq := &pb.StaffCreateRecordRequest{
			StaffToken: token,
			RecordRequest: &pb.RecordRequest{
				PatientIdCardNumber: req.PatientIdCard,
				PhotoHash:           req.PhotoHash,
				DocumentHash:        req.DocumentHash,
				Description:         req.Description,
			},
		}
		rpcRes, err := c.StaffCreateRecord(ctx, rpcReq)
		if err != nil {
			res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
			res.StatusAndMsgRes.Status = 406
		}
		res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
		res.StatusAndMsgRes.Status = 200
		return
	}

	if r.Method == "GET" {
		var req staffQueryRecordReq
		var res staffQueryRecordRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("staff_query_record encode error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Status = 400
			res.StatusAndMsgRes.Msg = "参数错误"
			res.RecordResArray = nil
			return
		}

		c := pb.NewRecordServiceClient(connStaff)
		ctx := context.Background()
		rpcReq := &pb.StaffQueryRecordsRequest{
			StaffToken:          token,
			PatientIdCardNumber: req.PatientIdCardNumber,
		}
		rpcRes, err := c.StaffQueryRecords(ctx, rpcReq)
		if err != nil {
			res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
			res.StatusAndMsgRes.Status = 406
			res.RecordResArray = nil
			return
		}

		res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
		res.StatusAndMsgRes.Status = 200
		for _, rpcRecord := range rpcRes.RecordResponses {
			httpRecord := recordRes{
				PatientIdCardNumber: rpcRecord.PatientIdCardNumber,
				PhotoHash:           rpcRecord.PhotoHash,
				DocumentHash:        rpcRecord.DocumentHash,
				LastChangeTime:      rpcRecord.LastChangeTime,
				Description:         rpcRecord.Description,
				RecordUid:           rpcRecord.RecordUid,
			}
			res.RecordResArray = append(res.RecordResArray, httpRecord)
		}
		return
	}
}
