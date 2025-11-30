package sglogger

// LoggerConfig defines base configuration for all loggers and providers.
// Contains common settings that apply to all logging components.
type LoggerConfig struct {
}

// ProviderConfig extends LoggerConfig with provider-specific settings.
// Embeds common configuration and adds provider-specific parameters.
type ProviderConfig struct {
	LoggerConfig        // Embedded base logger configuration
	Level       Level   // Provider-specific log level
}