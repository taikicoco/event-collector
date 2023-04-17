package config

import (

)

const (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
)

type ServerConfig struct {
	Environment       string `json:"ENV,omitempty" env:"ENV,required"`
	Port              string `json:"PORT,omitempty" env:"PORT,required"`
	DriverName        string `json:"DRIVER,omitempty" env:"DRIVER,required"`
	DataSource        string `json:"DATASOURCE,omitempty" env:"DATASOURCE,required"`
	CognitoUserPoolID string `json:"CognitoUserPoolID,omitempty" env:"CognitoUserPoolID,required"`
	CognitoClientID   string `json:"CognitoClientID,omitempty" env:"CognitoClientID,required"`
	BucketName        string `json:"BucketName,omitempty" env:"BucketName,required"`
}
