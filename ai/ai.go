/*
Copyright © 2026 Dominik Meisner <meisnerd2003@gmail.com>

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
package ai

import (
	"context"

	"google.golang.org/genai"
)

type Service struct {
	client *genai.Client
	model  string
}

func New(apiKey string, model string) (*Service, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		client: client,
		model:  model,
	}, nil
}

func (s *Service) GetCommitMessage(diff string, giveDescription bool) (result string, e error) {
	ctx := context.Background()

	var prompt string
	if giveDescription {
		prompt = "Generate git conventional commit title and description based on the provided diff: " + diff
	} else {
		prompt = "Generate git conventional commit title (!no description!) based on the provided diff: " + diff
	}

	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}
	response, err := s.client.Models.GenerateContent(ctx, s.model, contents, nil)
	if err != nil {
		return "", err
	}

	return response.Text(), nil
}
