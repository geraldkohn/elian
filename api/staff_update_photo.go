package api

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc/grpclog"
)

func staffUpdatePhotoHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "POST" {
		var req staffUpdatePhotoReq
		var res staffUpdatePhotoRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("staff update photo error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Msg = "参数错误"
			res.StatusAndMsgRes.Status = 400
		}
		c := pb.NewRecordServiceClient(connStaff)
		ctx := context.Background()
		rpcReq := &pb.StaffUpdatePhotoHashRequest{
			StaffToken: token,
			PhotoHash: req.PhotoHash,
			RecordUid: req.RecordUid,
		}
		rpcRes, err := c.StaffUpdatePhotoHash(ctx, rpcReq)
		if err != nil {
			res.StatusAndMsgRes.Status = 406
			res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
			return
		}
		res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
		res.StatusAndMsgRes.Status = 200
		return
	}
}