package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/agency"
	"google.golang.org/grpc/grpclog"
)

func agencyLoginHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	var req agencyLoginReq
	var res registerAndLoginRes
	defer func() {
		res.Attribute = "agency"
		if err := json.NewEncoder(w).Encode(res); err != nil {
			grpclog.Error(err)
		}
		grpclog.Info("agency encode error")
	}()

	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.Status = 400
			res.Token = ""
			res.Msg = "参数错误"
			return
		}

		c := pb.NewAgencyServiceClient(connPatient)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		agencyLoginReq := &pb.AgencyLoginRequest{
			License: req.License,
			Token:   token,
		}

		agencyLoginRes, err := c.AgencyLogin(ctx, agencyLoginReq)
		if err != nil {
			res.Msg = agencyLoginRes.Msg
			res.Token = ""
			res.Status = 406
		}
		res.Msg = agencyLoginRes.Msg
		res.Token = agencyLoginRes.Token
		res.Status = 200
	}
}
