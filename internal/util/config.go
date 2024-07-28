package util

import (
    "io/ioutil"

    "github.com/dawitel/addispay-project/internal/util"
    
    "gopkg.in/yaml.v2"
)

type Config struct {
    Pulsar struct {
        serviceURL string `yaml:"serviceURL"`
        Functions struct {
                orderProcessing struct {
                    inputTopic string `yaml:"orders-topic"`
                    outputTopic string `yaml:"processed-orders-topic"` 
                }`yaml:"orderProcessing"`
                paymentProcessing struct {
                    inputTopic string `yaml:"processed-orders-topic"`
                    outputTopic string `yaml:"payment-results-topic"` 
                }`yaml:"paymentProcessing"`
                orderFinalization struct {
                    inputTopic string `yaml:"payment-results-topic"`
                }`yaml:"orderFinalization"`
        }`yaml:"functions"`
    }`yaml:"pulsar"`
   

    Grpc struct {
        server struct {
            host string `yaml:"host"`
            port int `yaml:"port"`
            
            TLS struct {
                enabled bool `yaml:"enabled"`
                cert_file string `yaml:"cert_file"`
                key_file string `yaml:"key_file"`
            }`yaml:"tls"`
        }`yaml:"server"`
    } `yaml:"grpc"`
   
    Logging struct {
        Level string `yaml:"level"`
        Format string `yaml:"format"`
        Output string `yaml:"output"`
    } `yaml:"logging"`
}

// LoadConfig gets the path to theconfig.yaml file and loads the configuration settings
func LoadConfig(path string) *Config {
    var config Config
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        logger.Error("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        logger.Error("Unmarshal: %v", err)
    }
    return &config
}
