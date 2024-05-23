package main

import (
	"github.com/hovhannisyangevorg/Data-Procesor/internal/client/http/AWSClient"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DataProcesorConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/service"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
)

func main() {
	cfg := DataProcesorConfig.LoadConfig()
	awsClient, err := AWSClient.NewAWSService(cfg)
	if err != nil {
		utils.WrapError("main", err)
	}
	service.DataProcesor(cfg, awsClient)
}
