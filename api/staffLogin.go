package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/staff"
	"google.golang.org/grpc/grpclog"
)

func staffLoginHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	var req staffLoginReq
	var res registerAndLoginRes
	defer func() {
		res.Attribute = "staff"
		if err := json.NewEncoder(w).Encode(res); err != nil {
			grpclog.Error(err)
		}
		grpclog.Info("staff encode error")
	}()

	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.Status = 400
			res.Token = ""
			res.Msg = "参数错误"
			return
		}

		c := pb.NewStaffServiceClient(connStaff)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rpcReq := &pb.StaffLoginRequest{
			Token:        token,
			IdCardNumber: req.IdCardNumber,
			Password:     req.Password,
		}

		rpcRes, err := c.StaffLogin(ctx, rpcReq)
		if err != nil {
			res.Msg = rpcRes.GetMsg()
			res.Token = ""
			res.Status = 406
			return 
		}
		res.Msg = rpcRes.GetMsg()
		res.Token = rpcRes.GetToken()
		res.Status = 200
		return 
	}
}
