package main

import (
	"github.com/geraldkohn/elian/config"
	"github.com/geraldkohn/elian/data/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "data/query",
	})

	g.UseDB(config.DB)

	g.ApplyBasic(model.Agency{}, model.Patient{}, model.Staff{}, model.Record{}, model.PatientuidToRecorduid{})

	g.Execute()
}
