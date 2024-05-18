package main

import "github.com/hovhannisyangevorg/Data-Procesor/internal/config"

func loadConfig() *config.Config {
	cfg := config.LoadConfig()
	return cfg
}

func main() {
	config := loadConfig()
	//name := "cheetah"
	//apiKey := "tvMSHFjLOks7Wi6tpKR90w==piwwP2aqnczND54z"
	//apiUrl := fmt.Sprintf("https://api.api-ninjas.com/v1/animals?name=%s", name)
	//
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", apiUrl, nil)
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//
	//req.Header.Set("X-Api-Key", apiKey)
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode == http.StatusOK {
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		fmt.Println("Error reading response body:", err)
	//		return
	//	}
	//	fmt.Println(string(body))
	//} else {
	//	fmt.Println("Error:", resp.Status)
	//}
}
