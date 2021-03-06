package model

import (
	"time"
)

const (
	DeploymentRouteTypeHTTPService string = "http-service"
	DeploymentRouteTypeStatic      string = "static"
)

type DeploymentRoute struct {
	ID         string
	CreatedAt  *time.Time
	Version    string
	Path       string
	Type       string
	TypeConfig RouteTypeConfig
}

type RouteTypeConfig map[string]interface{}

func (r RouteTypeConfig) BackendURL() string {
	if str, ok := r["backend_url"].(string); ok {
		return str
	}
	return ""
}

func (r RouteTypeConfig) TargetPath() string {
	if str, ok := r["target_path"].(string); ok {
		return str
	}
	return ""
}

func (r RouteTypeConfig) AssetPathMapping() map[string]string {
	m := r["asset_path_mapping"].(map[string]interface{})
	mapping := map[string]string{}
	for k, v := range m {
		mapping[k] = v.(string)
	}
	return mapping
}

func (r RouteTypeConfig) AssetFallbackPagePath() string {
	if p, ok := r["asset_fallback_page_path"].(string); ok {
		return p
	}
	return ""
}

func (r RouteTypeConfig) AssetErrorPagePath() string {
	if p, ok := r["asset_error_page_path"].(string); ok {
		return p
	}
	return ""
}

func (r RouteTypeConfig) AssetIndexFile() string {
	if f, ok := r["asset_index_file"].(string); ok {
		return f
	}
	return ""
}
