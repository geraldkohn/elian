package main

import (
	"log"
	"net/http"

	config "github.com/geraldkohn/elian/config"
	"google.golang.org/grpc"
)

var (
	connPatient *grpc.ClientConn
	connAgency  *grpc.ClientConn
	connStaff   *grpc.ClientConn
	connRecord  *grpc.ClientConn
)

func main() {
	var err error
	connPatient, err = grpc.Dial(config.Host + ":" + config.PatientPort)
	if err != nil {
		log.Fatalf("患者模块客户端初始化失败: %v", err)
	}
	connAgency, err = grpc.Dial(config.Host + ":" + config.AgencyPort)
	if err != nil {
		log.Fatalf("机构模块客户端初始化失败: %v", err)
	}
	connStaff, err = grpc.Dial(config.Host + ":" + config.StaffPort)
	if err != nil {
		log.Fatalf("医生模块客户端初始化失败: %v", err)
	}
	connRecord, err = grpc.Dial(config.Host + ":" + config.RecordPort)
	if err != nil {
		log.Fatalf("病历记录模块客户端初始化失败: %v", err)
	}
	defer func() {
		connPatient.Close()
		connAgency.Close()
		connStaff.Close()
		connRecord.Close()
	}()

	log.Println("http server begin")

	http.HandleFunc("/api/patient/register", patientRegisterHandler)
	http.HandleFunc("/api/staff/register", staffRegisterHandler)
	http.HandleFunc("/api/agency/register", agencyRegisterHandler)
	http.HandleFunc("/api/patient/login", patientLoginHandler)
	http.HandleFunc("/api/staff/login", staffLoginHandler)
	http.HandleFunc("/api/agency/login", agencyLoginHandler)
	http.HandleFunc("/api/record/staff", staffQueryOrCreateRecordHandler)
	http.HandleFunc("/api/record/staff-set-rw-permission", staffSetRWPermissionHandler)
	http.HandleFunc("/api/record/staff-set-r-permission", staffSetRPermissionHandler)
	http.HandleFunc("/api/record/staff-update-photo", staffUpdatePhotoHandler)
	http.HandleFunc("/api/record/staff-update-document", staffUpdateDocumentHandler)
	http.HandleFunc("/api/record/staff-update-description", staffUpdateDescriptionHandler)
	http.HandleFunc("/api/record/patient", patientQueryRecordHandler)
	http.HandleFunc("/api/record/agency", agencyQueryRecordHandler)
	http.HandleFunc("/api/record/agency-set-rw-permission", agencySetRWPermissionHandler)
	http.HandleFunc("/api/record/agency-set-r-permission", agencySetRPermissionHandler)

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("http serve error")
	}
}
