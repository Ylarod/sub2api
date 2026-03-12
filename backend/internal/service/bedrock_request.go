package service

import (
	"fmt"

	"github.com/tidwall/sjson"
)

// BuildBedrockURL 构建 Bedrock InvokeModel 的 URL
// stream=true 时使用 invoke-with-response-stream 端点
func BuildBedrockURL(region, modelID string, stream bool) string {
	if stream {
		return fmt.Sprintf("https://bedrock-runtime.%s.amazonaws.com/model/%s/invoke-with-response-stream", region, modelID)
	}
	return fmt.Sprintf("https://bedrock-runtime.%s.amazonaws.com/model/%s/invoke", region, modelID)
}

// PrepareBedrockRequestBody 处理请求体以适配 Bedrock API
// 1. 注入 anthropic_version
// 2. 移除 Bedrock 不支持的字段（model 字段由 URL 指定）
func PrepareBedrockRequestBody(body []byte) ([]byte, error) {
	var err error

	// 注入 anthropic_version（Bedrock 要求）
	body, err = sjson.SetBytes(body, "anthropic_version", "bedrock-2023-05-31")
	if err != nil {
		return nil, fmt.Errorf("inject anthropic_version: %w", err)
	}

	// 移除 model 字段（Bedrock 通过 URL 指定模型）
	body, err = sjson.DeleteBytes(body, "model")
	if err != nil {
		return nil, fmt.Errorf("remove model field: %w", err)
	}

	return body, nil
}
