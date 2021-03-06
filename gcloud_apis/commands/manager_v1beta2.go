// generated code: api commands
/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	api_client "github.com/googlecloudplatform/gcloud/gcloud_apis/clients/manager/v1beta2"
	"github.com/googlecloudplatform/gcloud/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin

func Manager_v1beta2_DeploymentsDelete(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDeploymentsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "DELETE")
	}

	expectedParams := []string{
		"projectId",
		"region",
		"deploymentName",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_region := paramValues[1]
	param_deploymentName := paramValues[2]

	call := service.Delete(param_projectId, param_region, param_deploymentName)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_DeploymentsGet(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDeploymentsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{
		"projectId",
		"region",
		"deploymentName",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_region := paramValues[1]
	param_deploymentName := paramValues[2]

	call := service.Get(param_projectId, param_region, param_deploymentName)

	var response *api_client.Deployment
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_DeploymentsInsert(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDeploymentsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		commands_util.UsageForMethod(context.InvocationMethod, "POST")
	}

	request := &api_client.Deployment{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	keyValues := flagValues

	err = commands_util.OverwriteRequestWithValues(&request, keyValues)
	if err != nil {
		return err
	}

	expectedParams := []string{
		"projectId",
		"region",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_region := paramValues[1]

	call := service.Insert(param_projectId, param_region,
		request,
	)

	var response *api_client.Deployment
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_DeploymentsList(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDeploymentsService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"pageToken":  false,
	}

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	for k, r := range queryParamNames {
		if _, ok := flagValues[k]; r && !ok {
			return fmt.Errorf("missing required flag %q", "--"+k)
		}
	}

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{
		"projectId",
		"region",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_region := paramValues[1]

	call := service.List(param_projectId, param_region)

	// Set query parameters.
	if value, ok := flagValues["maxResults"]; ok {
		query_maxResults, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.MaxResults(query_maxResults)
	}
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}

	var response *api_client.DeploymentsListResponse
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_TemplatesDelete(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewTemplatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "DELETE")
	}

	expectedParams := []string{
		"projectId",
		"templateName",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_templateName := paramValues[1]

	call := service.Delete(param_projectId, param_templateName)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_TemplatesGet(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewTemplatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{
		"projectId",
		"templateName",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]
	param_templateName := paramValues[1]

	call := service.Get(param_projectId, param_templateName)

	var response *api_client.Template
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_TemplatesInsert(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewTemplatesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		commands_util.UsageForMethod(context.InvocationMethod, "POST")
	}

	request := &api_client.Template{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	keyValues := flagValues

	err = commands_util.OverwriteRequestWithValues(&request, keyValues)
	if err != nil {
		return err
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Insert(param_projectId,
		request,
	)

	var response *api_client.Template
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Manager_v1beta2_TemplatesList(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewTemplatesService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"pageToken":  false,
	}

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	for k, r := range queryParamNames {
		if _, ok := flagValues[k]; r && !ok {
			return fmt.Errorf("missing required flag %q", "--"+k)
		}
	}

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.List(param_projectId)

	// Set query parameters.
	if value, ok := flagValues["maxResults"]; ok {
		query_maxResults, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.MaxResults(query_maxResults)
	}
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}

	var response *api_client.TemplatesListResponse
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}
