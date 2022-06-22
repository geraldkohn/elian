package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/patient"
	"google.golang.org/grpc/grpclog"
)

func patientLoginHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	var req patientLoginReq
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

		patientLoginReq := &pb.PatientLoginRequest{
			Token:        token,
			IdCardNumber: req.IdCardNumber,
			Password:     req.Password,
		}

		patientLoginRes, err := c.PatientLogin(ctx, patientLoginReq)
		if err != nil {
			res.Msg = patientLoginRes.Msg
			res.Token = ""
			res.Status = 406
		}
		res.Msg = patientLoginRes.Msg
		res.Token = patientLoginRes.Token
		res.Status = 200
	}
}
