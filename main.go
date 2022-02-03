package main

import (
	"context"
	"fmt"

	"github.com/butlerhq/airbyte-client-go/airbyte"
)

const AIRBYTE_URL = "http://localhost:8001/api"

func createClient() *airbyte.ClientWithResponses {
	client, err := airbyte.NewClientWithResponses(AIRBYTE_URL)
	if err != nil {
		panic(err)
	}
	return client
}

func listSourceDefinitions(ctx context.Context, client *airbyte.ClientWithResponses) {
	resp, err := client.ListSourceDefinitionsWithResponse(ctx)
	if err != nil {
		panic(err)
	}
	for _, source := range resp.JSON200.SourceDefinitions {
		fmt.Printf("\n%+v", source)
	}
}

func createAndListWorkspaces(ctx context.Context, client *airbyte.ClientWithResponses) {
	body := airbyte.CreateWorkspaceJSONRequestBody{
		AnonymousDataCollection: airbyte.PtrBool(true),
		DisplaySetupWizard:      airbyte.PtrBool(false),
		Name:                    "Workspace test",
		News:                    airbyte.PtrBool(false),
		SecurityUpdates:         airbyte.PtrBool(false),
	}

	wsResponse, err := client.CreateWorkspaceWithResponse(ctx, body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nCreated workspace\n%+v", wsResponse.JSON200)

	wsListResponse, err := client.ListWorkspacesWithResponse(ctx)
	if err != nil {
		panic(err)
	}
	for _, workspace := range wsListResponse.JSON200.Workspaces {
		fmt.Printf("\n%s", workspace.WorkspaceId)
	}
}

func main() {
	ctx := context.Background()
	client := createClient()
	listSourceDefinitions(ctx, client)
	// createAndListWorkspaces(ctx, client)
}
