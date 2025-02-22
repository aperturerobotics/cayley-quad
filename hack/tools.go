//go:build deps_only
// +build deps_only

package hack

import (
	// _ imports protowrap
	_ "github.com/aperturerobotics/goprotowrap/cmd/protowrap"
	// _ imports protoc-gen-go
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	// _ imports protoc-gen-go-vtproto
	_ "github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto"
	// _ imports golangci-lint
	_ "github.com/golangci/golangci-lint/pkg/golinters"
	// _ imports golangci-lint commands
	_ "github.com/golangci/golangci-lint/pkg/commands"
	// _ imports go-mod-outdated
	_ "github.com/psampaz/go-mod-outdated"
)
