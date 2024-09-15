package cmd

// GlobalConfig holds the global flags like swaggerPath.
type GlobalConfig struct {
	SwaggerPath string
}

// Initialize the global config instance
var config GlobalConfig
