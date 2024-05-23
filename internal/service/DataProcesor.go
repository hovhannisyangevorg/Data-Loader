package service

import (
	"encoding/json"
	"fmt"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/client/http/AWSClient"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/client/http/DataProcesorClient"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DataProcesorConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/structs"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
	"io/ioutil"
)

func DataProcesor(cfg *DataProcesorConfig.Config, awsClient *AWSClient.AWSService) {
	var global structs.GlobalInfrastructura
	client := DataProcesorClient.NewClient()
	animalOperation := structs.NewAnimalOperation()
	var FilePresentation utils.PresentationPath

	for _, name := range animalOperation.AnimalName {
		err := client.SetRequest("GET", fmt.Sprintf(cfg.CoreUrl, name), cfg.ApiKey, nil)
		if err != nil {
			utils.WrapError("DataProcesor", err)
			continue
		}
		resp, err := client.Do()
		if err != nil {
			utils.WrapError("DataProcesor", err)
		}
		defer resp.Body.Close()
		JSONBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utils.WrapError("DataProcesor", err)
		}
		err = json.Unmarshal(JSONBody, &global.Animal)
		if err != nil {
			utils.WrapError("DataProcesor", err)
		}
		FilePresentation.DirName = name
		FilePresentation.FileData = JSONBody
		err = SaveFileToS3(cfg, awsClient, global, FilePresentation)
		if err != nil {
			utils.WrapError("DataProcesor", err)
		}
		//fmt.Println("UREEEE:", fmt.Sprintf(cfg.CoreUrl, name))
	}
}

func SaveFileToS3(cfg *DataProcesorConfig.Config, awsClient *AWSClient.AWSService, GlobalInfo structs.GlobalInfrastructura, FilePresentation utils.PresentationPath) error {

	for _, animal := range GlobalInfo.Animal {
		FilePresentation.FileName = utils.IgnoreSpaces(animal.Name)
		err := awsClient.UploadFile(cfg, fmt.Sprintf("/%s/%s.json", FilePresentation.DirName, FilePresentation.FileName), FilePresentation.FileData)
		if err != nil {
			return utils.WrapError("SaveFileToS3", err)
		}
	}
	return nil
}
