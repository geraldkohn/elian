package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/staff"
	"google.golang.org/grpc/grpclog"
)

func staffRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req staffRegisterReq
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

		createStaffReq := &pb.CreateStaffRequset{
			Name:         req.Name,
			Password:     req.Password,
			IdCardNumber: req.IdCardNumber,
		}

		createStaffRes, err := c.CreateStaff(ctx, createStaffReq)

		if err != nil {
			res.Msg = createStaffRes.Msg
			res.Token = createStaffRes.Token
			res.Status = 406
			return
		}

		res.Msg = createStaffRes.Msg
		res.Token = createStaffRes.Token
		res.Status = 200
		return
	}
}
