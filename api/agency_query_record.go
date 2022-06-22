package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc/grpclog"
)

func agencyQueryRecordHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "GET" {
		var req agencyQueryRecordReq
		var res agencyQueryRecordRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("agency_query_record encode error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Status = 400
			res.StatusAndMsgRes.Msg = "参数错误"
			res.RecordResArray = nil
			return
		}

		c := pb.NewRecordServiceClient(connAgency)
		ctx := context.Background()
		rpcReq := &pb.AgencyQueryRecordsRequest{
			AgencyToken:          token,
		}
		rpcRes, err := c.AgencyQueryRecords(ctx, rpcReq)
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