package conf_test

// Benchmark for config module

import (
	"os"
	"testing"

	"config-struct/conf"
)

const (
	configFileEnvVariableName = "CCX_NOTIFICATION_SERVICE_CONFIG_FILE"
	defaultConfigFileName     = "./config"
)

// HELPER START OMIT

func loadConfiguration() (conf.ConfigStruct, error) {
	os.Clearenv()
	err := os.Setenv(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return conf.ConfigStruct{}, err
	}
	config, err := conf.LoadConfiguration(configFileEnvVariableName, defaultConfigFileName)
	if err != nil {
		return conf.ConfigStruct{}, err
	}
	return config, nil
}

func mustLoadBenchmarkConfiguration(b *testing.B) conf.ConfigStruct {
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
		conf.GetStorageConfigurationByValue(configuration)
	}
}

func BenchmarkGetStorageConfigurationFunctionByReference(b *testing.B) {
	configuration := mustLoadBenchmarkConfiguration(b)
	b.ResetTimer()
	for range b.N {
		conf.GetStorageConfigurationByReference(&configuration)
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
