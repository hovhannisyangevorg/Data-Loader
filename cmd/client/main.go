package main

import (
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DPConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/service"
)

func main() {
	//ctx := context.Background()
	cfg := DPConfig.LoadConfig()
	//db, err := DBConfig.ConnectToDB(cfg.Database)
	//if err != nil {
	//	fmt.Errorf("main: %s", err)
	//}
	service.DataProcesor(cfg)

}
