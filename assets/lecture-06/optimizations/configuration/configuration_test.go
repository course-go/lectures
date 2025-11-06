package configuration_test

// Benchmark for config module

import (
	"os"
	"testing"

	"github.com/course-go/lectures/assets/lecture-06/optimizations/configuration"
)

const (
	configFileEnvVariableName = "CCX_NOTIFICATION_SERVICE_CONFIG_FILE"
	defaultConfigFileName     = "../config"
)

// BENCHMARK START OMIT

func BenchmarkGetStorageConfigurationFunctionByValue(b *testing.B) {
	config := mustLoadConfiguration(b)
	for b.Loop() {
		configuration.StorageConfigurationByValue(config)
	}
}

func BenchmarkGetStorageConfigurationFunctionByReference(b *testing.B) {
	config := mustLoadConfiguration(b)
	for b.Loop() {
		configuration.StorageConfigurationByReference(&config)
	}
}

// BENCHMARK MIDDLE OMIT

func BenchmarkGetStorageConfigurationMethodByValue(b *testing.B) {
	config := mustLoadConfiguration(b)
	for b.Loop() {
		config.StorageConfigurationByValue()
	}
}

func BenchmarkGetStorageConfigurationMethodByReference(b *testing.B) {
	config := mustLoadConfiguration(b)
	for b.Loop() {
		config.StorageConfigurationByReference()
	}
}

// BENCHMARK END OMIT
// HELPER START OMIT

func mustLoadConfiguration(b *testing.B) configuration.Config {
	b.Helper()
	configuration, err := loadConfiguration()
	if err != nil {
		b.Fatal(err)
	}
	return configuration
}

func loadConfiguration() (configuration.Config, error) {
	os.Clearenv()
	err := os.Setenv(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return configuration.Config{}, err
	}
	config, err := configuration.LoadConfiguration(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return configuration.Config{}, err
	}
	return config, nil
}

// HELPER END OMIT
