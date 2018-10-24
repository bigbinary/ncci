package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	models "github.com/semaphoreci/cli/api/models"
)

type WorkflowApiV1AlphaApi struct {
	BaseClient           BaseClient
	ResourceNameSingular string
	ResourceNamePlural   string
}

func NewWorkflowV1AlphaApi() WorkflowApiV1AlphaApi {
	baseClient := NewBaseClientFromConfig()
	baseClient.SetApiVersion("v1alpha")

	return WorkflowApiV1AlphaApi{
		BaseClient:           baseClient,
		ResourceNamePlural:   "plumber-workflows",
		ResourceNameSingular: "plumber-workflow",
	}
}

func (c *WorkflowApiV1AlphaApi) ListWorkflows(project_id string) (*models.WorkflowListV1Alpha, error) {
	urlEncode := fmt.Sprintf("%s?project_id=%s", c.ResourceNamePlural, project_id)
	body, status, err := c.BaseClient.List(urlEncode)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("connecting to Semaphore failed '%s'", err))
	}

	if status != 200 {
		return nil, errors.New(fmt.Sprintf("http status %d with message \"%s\" received from upstream", status, body))
	}

	return models.NewWorkflowListV1AlphaFromJson(body)
}

func (c *WorkflowApiV1AlphaApi) CreateSnapshotWf(project_id string) ([]byte, error) {
	requestToken, err := uuid.NewUUID()

	req, err := models.NewWorkflowV1AlphaSnapshotRequest(project_id, "archive content", requestToken.String())

	if err != nil {
		return nil, fmt.Errorf("creating rebuild request failed '%s'", err)
	}

	requestBody, _ := json.Marshal(req)

	if err != nil {
		return nil, fmt.Errorf("request body creation failed '%s'", err)
	}

	body, status, err := c.BaseClient.Post(c.ResourceNamePlural, requestBody)

	switch {
	case err != nil:
		return nil, fmt.Errorf("connecting to Semaphore failed '%s'", err)
	case status != 200:
		return nil, fmt.Errorf("http status %d with message \"%s\" received from upstream", status, body)
	}

	return body, nil
}
