package api

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/geraldkohn/elian/proto/record"
	"google.golang.org/grpc/grpclog"
)

func staffSetRWPermissionHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "POST" {
		var req staffSetReadAndWritePermissionReq
		var res staffSetReadAndWritePermissionRes
		defer func() {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				grpclog.Error(err)
			}
			grpclog.Info("staff set read and write permission error")
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.StatusAndMsgRes.Msg = "参数错误"
			res.StatusAndMsgRes.Status = 400
		}

		c := pb.NewRecordServiceClient(connStaff)
		ctx := context.Background()
		var permissionReqs []*pb.PermissionRequest
		permissionReq := &pb.PermissionRequest{
			Hospital:   req.Hospital,
			Department: req.Department,
			Name:       req.Name,
		}
		permissionReqs = append(permissionReqs, permissionReq)
		rpcReq := &pb.StaffSetReadAndWritePermissionRequest{
			StaffToken:         token,
			RecordUid:          req.RecordUid,
			PermissionRequests: permissionReqs,
		}
		rpcRes, err := c.StaffSetReadAndWritePermission(ctx, rpcReq)
		if err != nil {
			res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
			res.StatusAndMsgRes.Status = 406
			return
		}

		res.StatusAndMsgRes.Msg = rpcRes.ErrorCodeAndInfo.Msg
		res.StatusAndMsgRes.Status = 200
		return
	}
}
