package api

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc/grpclog"
)

func patientQueryRecordHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "GET" {
		var req patientQueryRecordReq
		var res patientQueryRecordRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("patient_query_record encode error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Status = 400
			res.StatusAndMsgRes.Msg = "参数错误"
			res.RecordResArray = nil
			return
		}

		c := pb.NewRecordServiceClient(connPatient)
		ctx := context.Background()
		rpcReq := &pb.PatientQueryRecordsRequest{
			PatientToken:          token,
		}
		rpcRes, err := c.PatientQueryRecords(ctx, rpcReq)
		if err != nil {
			res.StatusAndMsgRes.Msg = rpcRes.GetErrorCodeAndInfo().GetMsg()
			res.StatusAndMsgRes.Status = 406
			res.RecordResArray = nil
			return
		}

		res.StatusAndMsgRes.Msg = rpcRes.GetErrorCodeAndInfo().GetMsg()
		res.StatusAndMsgRes.Status = 200
		for _, rpcRecord := range rpcRes.GetRecordResponses() {
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