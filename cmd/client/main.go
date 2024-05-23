package main

import (
	"github.com/hovhannisyangevorg/Data-Procesor/internal/client/http/AWSClient"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/ProcesConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/service"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
)

func main() {

	cfg := ProcesConfig.LoadConfig()
	awsClient, err := AWSClient.NewAWSService(cfg)
	if err != nil {
		utils.WrapError("main", err)
	}
	service.DataProcesor(cfg, awsClient) //awsClient
}

//fileName := "example1.txt"
//fileContent := []byte("Hello, World!")
//err = service.UploadFile(cfg, fileName, fileContent)
//if err != nil {
//	return nil, utils.WrapError("NewAWSService", err)
//}
