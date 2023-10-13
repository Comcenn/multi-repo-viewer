package git

import (
	"automated_retro/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DEPENDABOT_QUERY = `{"query": "query { repository(owner: \"Rentecarlo\", name:\"%s\") { vulnerabilityAlerts(first:20, states: OPEN) { nodes { number state createdAt securityVulnerability{ severity }}}}}"}`

type Git struct {
	config       *config.Config
	client       *http.Client
	errorHandler func(error)
}

func (g Git) createGraphQlRequest(query string) *http.Request {
	req, err := http.NewRequest("POST", g.config.Github.Host+g.config.Github.Graphql, bytes.NewBufferString(query))
	if err != nil {
		g.errorHandler(err)
	}
	req.Header.Set("Authorization", "bearer "+g.config.Github.Token)
	return req
}

func (g Git) makeRequest(req *http.Request) (*http.Response, error) {
	resp, err := g.client.Do(req)
	return resp, err
}

func (g Git) getDependabotQuery(repo string) string {
	return fmt.Sprintf(DEPENDABOT_QUERY, repo)
}

func (g Git) parseJson(data []byte) map[string]interface{} {
	var jsonMap map[string]interface{}
	json.Unmarshal(data, &jsonMap)
	return jsonMap
}

func (g Git) GetDependabotAlerts(repo string) map[string]interface{} {
	query := g.getDependabotQuery(repo)
	req := g.createGraphQlRequest(query)
	resp, err := g.makeRequest(req)
	if err != nil {
		g.errorHandler(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.errorHandler(err)
	}

	return g.parseJson(body)

}

func CreateGit(cfg *config.Config, errHandler func(error)) Git {
	client := &http.Client{}
	return Git{cfg, client, errHandler}
}
