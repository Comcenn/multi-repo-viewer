package main

import (
	"automated_retro/config"
	"automated_retro/git"
	"fmt"
	"os"
)

func ErrorHandler(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func reduceResults(alertsMap map[string]interface{}) map[string]int {
	reduction := map[string]int{"numberOfAlerts": 0, "aboveThreshold": 0}
	// rawData, ok := alertsMap["data"]
	// if !ok {
	// 	return reduction
	// }
	// data, ok := rawData.(map[string]map[string]map[string][]map[string]interface{})
	// repository, ok := data["repository"]["vulnerabilityAlerts"]

	return reduction
}

func task(repoName string, cfg *config.Config, c chan map[string]int) {
	git_client := git.CreateGit(cfg, ErrorHandler)
	alerts := git_client.GetDependabotAlerts(cfg.Github.Owner, repoName)
	fmt.Println(repoName, ": ", alerts)
	c <- reduceResults(alerts)
}

func main() {
	cfg := config.GetConfig()

	channelSlice := []chan map[string]int{}

	for _, name := range cfg.Github.Repositories {
		c := make(chan map[string]int)

		channelSlice = append(channelSlice, c)
		go task(name, &cfg, c)
	}

	for _, c := range channelSlice {
		<-c
	}

}
