package secrets

import (
	"context"
	"github.com/blend/go-sdk/ex"
)

const errNilConfig ex.Class = "config cannot be nil"

type SecretTraceConfig struct {
	VaultOperation string
	KeyName        string
}

// TraceOption is an option type for secret trace
type TraceOption func(config *SecretTraceConfig) error

// TraceOptConfig allows you to provide the entire secret trace configuration
func TraceOptConfig(providedConfig SecretTraceConfig) TraceOption {
	return func(config *SecretTraceConfig) error {
		*config = providedConfig
		return nil
	}
}

// TraceOptVaultOperation allows you to set the VaultOperation being hit
func TraceOptVaultOperation(path string) TraceOption {
	return func(config *SecretTraceConfig) error {
		if config == nil {
			return errNilConfig
		}
		config.VaultOperation = path
		return nil
	}
}

// TraceOptKeyName allows you to specify the name of the key being interacted with
func TraceOptKeyName(keyName string) TraceOption {
	return func(config *SecretTraceConfig) error {
		if config == nil {
			return errNilConfig
		}
		config.KeyName = keyName
		return nil
	}
}

// Tracer is a tracer for requests.
type Tracer interface {
	Start(ctx context.Context, options ...TraceOption) (TraceFinisher, error)
}

// TraceFinisher is a finisher for traces.
type TraceFinisher interface {
	Finish(ctx context.Context, statusCode int, vaultError error)
}
