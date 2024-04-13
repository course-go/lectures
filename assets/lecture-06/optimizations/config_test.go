package optimizations_test

// Benchmark for config module

import (
	"os"
	"testing"

	"github.com/course-go/lectures/assets/lecture-06/optimizations"
)

const (
	configFileEnvVariableName = "CCX_NOTIFICATION_SERVICE_CONFIG_FILE"
	defaultConfigFileName     = "./config"
)

// HELPER START OMIT

func loadConfiguration() (optimizations.ConfigStruct, error) {
	os.Clearenv()
	err := os.Setenv(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return optimizations.ConfigStruct{}, err
	}
	config, err := optimizations.LoadConfiguration(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return optimizations.ConfigStruct{}, err
	}
	return config, nil
}

func mustLoadBenchmarkConfiguration(b *testing.B) optimizations.ConfigStruct {
	configuration, err := loadConfiguration()
	if err != nil {
		b.Fatal(err)
	}
	return configuration
}

// HELPER END OMIT
// BENCHMARK START OMIT

func BenchmarkGetStorageConfigurationFunctionByValue(b *testing.B) {
	configuration := mustLoadBenchmarkConfiguration(b)
	b.ResetTimer()
	for range b.N {
		optimizations.GetStorageConfigurationByValue(configuration)
	}
}

func BenchmarkGetStorageConfigurationFunctionByReference(b *testing.B) {
	configuration := mustLoadBenchmarkConfiguration(b)
	b.ResetTimer()
	for range b.N {
		optimizations.GetStorageConfigurationByReference(&configuration)
	}
}

// BENCHMARK MIDDLE OMIT

func BenchmarkGetStorageConfigurationMethodByValue(b *testing.B) {
	configuration := mustLoadBenchmarkConfiguration(b)
	b.ResetTimer()
	for range b.N {
		configuration.GetStorageConfigurationByValue()
	}
}

func BenchmarkGetStorageConfigurationMethodByReference(b *testing.B) {
	configuration := mustLoadBenchmarkConfiguration(b)
	b.ResetTimer()
	for range b.N {
		configuration.GetStorageConfigurationByReference()
	}
}

// BENCHMARK END OMIT
