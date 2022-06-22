package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/patient"
	"google.golang.org/grpc/grpclog"
)

func patientRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req patientRegisterReq
	var res registerAndLoginRes
	defer func() {
		res.Attribute = "patient"
		if err := json.NewEncoder(w).Encode(res); err != nil {
			grpclog.Error(err)
		}
		grpclog.Info("patient encode error")
	}()

	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			r.Body.Close()
			res.Status = 400
			res.Token = ""
			res.Msg = "参数错误"
			return
		}

		c := pb.NewPatientServiceClient(connPatient)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		createPatientReq := &pb.CreatePatientRequest{
			Name:         req.Name,
			Password:     req.Password,
			IdCardNumber: req.IdCardNumber,
		}

		createPatientRes, err := c.CreatePatient(ctx, createPatientReq)

		if err != nil {
			res.Msg = createPatientRes.Msg
			res.Token = createPatientRes.Token
			res.Status = 406
			return
		}

		res.Msg = createPatientRes.Msg
		res.Token = createPatientRes.Token
		res.Status = 200
		return
	}
}
