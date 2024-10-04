package libs

import (
	"fmt"
	"os"
	"strconv"
)

const envVarPreFix = "LUNGFISH_"

type config struct {
	PublicPort        int
	CommunicationPort int
	BroadcastPort     int
	LogLevel          int
}

func initConfig() *config {
	return &config{
		PublicPort:        getIntEnvVar("PUBLIC_PORT", 8000),
		CommunicationPort: getIntEnvVar("COMMS_PORT", 8010),
		BroadcastPort:     getIntEnvVar("BROADCAST_PORT", 8020),
		LogLevel:          getIntEnvVar("LOG_LEVEL", 4),
	}
}

func generateEnvKey(envName string) string {
	return fmt.Sprintf("%s%s", envVarPreFix, envName)
}

func logEnvVarErr(envKey string, varType string, err error) {
	GetServices().GetLogger().Fatal(
		fmt.Sprintf("Failed to converting environment variable '%s' to %s. Error: %s", envKey, varType, err),
	)
}

func getStringEnvVar(envName string, defaultValue interface{}) string {
	var osEnvVal string

	if osEnvVal = os.Getenv(generateEnvKey(envName)); osEnvVal == "" {
		return fmt.Sprintf("%v", defaultValue)
	}

	return osEnvVal
}

func getIntEnvVar(envName string, defaultValue int) int {
	envVal, err := strconv.Atoi(
		getStringEnvVar(envName, defaultValue),
	)
	if err != nil {
		logEnvVarErr(generateEnvKey(envName), "integer", err)
	}

	return envVal
}

func getFloatEnvVar(envName string, defaultValue float64) float64 {
	envVal, err := strconv.ParseFloat(
		getStringEnvVar(envName, defaultValue), 64,
	)
	if err != nil {
		logEnvVarErr(generateEnvKey(envName), "float", err)
	}

	return envVal
}

func getBoolEnvVar(envName string, defaultValue bool) bool {
	envVal, err := strconv.ParseBool(
		getStringEnvVar(envName, defaultValue),
	)
	if err != nil {
		logEnvVarErr(generateEnvKey(envName), "bool", err)
	}

	return envVal
}

// ShowSettings will log the settings to stdout
func (c *config) ShowSettings() {
	logger := GetServices().GetLogger()

	logger.Info(fmt.Sprintf("Public Port: %d", c.PublicPort))
	logger.Info(fmt.Sprintf("Communication Port: %d", c.CommunicationPort))
	logger.Info(fmt.Sprintf("Broadcast Port: %d", c.BroadcastPort))
	logger.Info(fmt.Sprintf("Log Level: %d", c.LogLevel))
}
