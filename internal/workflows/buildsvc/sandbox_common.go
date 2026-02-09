// File: internal/workflows/buildsvc/sandbox_common.go
// Brief: Internal buildsvc package implementation for 'sandbox common'.

// Package buildsvc provides buildsvc helpers.

package buildsvc

import (
	"context"
	"os"
)

const (
	sandboxActiveEnvKey        = "VERIFIER_SANDBOX_ACTIVE"
	legacySandboxActiveEnvKey  = "VERIFIER_NSJAIL_ACTIVE"
	sandboxLogPathEnvKey       = "VERIFIER_SANDBOX_LOG_PATH"
	legacySandboxLogPathEnv    = "VERIFIER_NSJAIL_LOG_PATH"
	sandboxDisableEnvKey       = "VERIFIER_SANDBOX_DISABLE"
	legacySandboxDisableEnv    = "VERIFIER_NSJAIL_DISABLE"
	sandboxContextEnvKey       = "VERIFIER_SANDBOX_CONTEXT"
	legacySandboxContextEnvKey = "VERIFIER_NSJAIL_CONTEXT"
	sandboxCacheEnvKey         = "VERIFIER_SANDBOX_CACHE"
	legacySandboxCacheEnvKey   = "VERIFIER_NSJAIL_CACHE"
	sandboxBuilderEnvKey       = "VERIFIER_SANDBOX_BUILDER"
	legacySandboxBuilderEnvKey = "VERIFIER_NSJAIL_BUILDER"
)

type sandboxInjector func(ctx context.Context, opts *Options, streams Streams, contextAbs string) (bool, error)

func sandboxActive() bool {
	return os.Getenv(sandboxActiveEnvKey) == "1" || os.Getenv(legacySandboxActiveEnvKey) == "1"
}

func sandboxLogPathFromEnv() string {
	if v := os.Getenv(sandboxLogPathEnvKey); v != "" {
		return v
	}
	return os.Getenv(legacySandboxLogPathEnv)
}

func sandboxContextFromEnv() string {
	if v := os.Getenv(sandboxContextEnvKey); v != "" {
		return v
	}
	return os.Getenv(legacySandboxContextEnvKey)
}

func sandboxCacheFromEnv() string {
	if v := os.Getenv(sandboxCacheEnvKey); v != "" {
		return v
	}
	return os.Getenv(legacySandboxCacheEnvKey)
}
