package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
	"github.com/mtnmunuklu/davudpasha-api-client-go/api/davudpasha"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	body := davudpasha.SourcesSaveRequest{
		Lgs: &davudpasha.SourcesItem{
			Enabled:          common.PtrBool(true),
			Name:             common.PtrString("test-file-source"),
			LogSourceDefCode: common.PtrString("test-file-lgs-defcode"),
			LogReaderType:    davudpasha.LOGREADERTYPE_FILE,
			AlertTimeout:     common.PtrInt64(120),
			LogReaderData: &davudpasha.SourcesLogReaderData{
				FileInfo: &davudpasha.SourcesFileInfo{
					Path:           common.PtrString("/home/temp/"),
					FileName:       common.PtrString("test-20240828.log"),
					FileNameFormat: *common.NewNullableString(common.PtrString("test-{yyyy}{MM}{dd}.log")),
				},
			},
			DiscardedLogsConfig: davudpasha.DISCARDEDLOGSCONFIG_TODETAILS,
			WriteRawLogs:        common.PtrBool(true),
			UseSecondaryWriter:  common.PtrBool(true),
		},
	}

	ctx := common.NewDefaultContext(context.Background())
	//ctx := common.NewContextWithEnvParams(context.Background(), "DP_SITE", "DP_API_KEY")
	configuration := common.NewConfiguration()
	configuration.SetHTTPClientWithInsecureSkipVerify()
	apiClient := common.NewAPIClient(configuration)
	api := davudpasha.NewSourcesApi(apiClient)
	resp, r, err := api.SaveSources(ctx, *davudpasha.NewSaveSourcesOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.SaveSources`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", " ")
	fmt.Fprintf(os.Stdout, "Response from `SourcesApi.SaveSources`:\n%s\n", responseContent)

}
