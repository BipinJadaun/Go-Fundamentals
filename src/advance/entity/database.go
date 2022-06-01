package entity

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DBConfig holds db related Configuration
type DBConfig struct {
	HostName string `json:"Hostname" yaml:"Hostname"`
	Port     string `json:"Port" yaml:"Port"`
	UserName string `json:"Username" yaml:"Username"`
	Password string `json:"Paswd" yaml:"Paswd"`
	DBName   string `json:"DBname" yaml:"DBname"`
}

// DBConn holds connection information
type DBConn struct {
	Config     *DBConfig
	Client     *mongo.Client
	Collection *mongo.Collection
}

// DBOutput contain db collection information
type DBOutput struct {
	Collection *mongo.Collection
}

// DBRequest holds request parameters
type DBRequest struct {
	Data       []interface{}
	Collection string
}
