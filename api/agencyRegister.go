package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/geraldkohn/elian/proto/agency"
	"google.golang.org/grpc/grpclog"
)

func agencyRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req agencyRegisterReq
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
			log.Println(err.Error())
			return
		}

		c := pb.NewAgencyServiceClient(connAgency)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rpcReq := &pb.CreateAgencyRequest{
			License: req.License,
		}

		rpcRes, err := c.CreateAgency(ctx, rpcReq)

		if err != nil {
			res.Msg = rpcRes.GetMsg()
			res.Token = rpcRes.GetToken()
			res.Status = 406
			return
		}

		res.Msg = rpcRes.GetMsg()
		res.Token = rpcRes.GetToken()
		res.Status = 200
		return
	}
}
