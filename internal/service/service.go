package service

import (
	"encoding/json"
	"fmt"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/client/http/DPClient"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DPConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/structs"
	"io/ioutil"
	"log"
)

func DataProcesor(cfg *DPConfig.Config) {
	global := structs.GlobalInfrastructura{}
	client := DPClient.NewClient()
	animalOperation := structs.NewAnimalOperation()

	for _, name := range animalOperation.AnimalName {
		err := client.SetRequest("GET", fmt.Sprintf(cfg.CoreUrl, name), cfg.ApiKey, nil)
		if err != nil {
			log.Fatalf("DataProcesor: %v", err)
		}
		resp, err := client.Do()
		if err != nil {
			log.Fatalf("DataProcesor: %v", err)
		}
		defer resp.Body.Close()
		JSONBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("DataProcesor: %v", err)
		}
		err = json.Unmarshal(JSONBody, &global.Animal)
		if err != nil {
			log.Fatalf("DataProcesor: %v", err)
		}

		fmt.Println(global.Animal[0])
		//fmt.Println("UREEEE:", fmt.Sprintf(cfg.CoreUrl, name))
	}
}
