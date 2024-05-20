package config

import (
	"database/sql"
	"fmt"
)

type DatabaseConfig interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetName() string
	GetSSLMode() string
	GetConnectString() string
}

type DatabaseConfigEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (c *DatabaseConfigEnv) GetHost() string {
	return c.Host
}

func (c *DatabaseConfigEnv) GetPort() string {
	return c.Port
}

func (c *DatabaseConfigEnv) GetUser() string {
	return c.User
}

func (c *DatabaseConfigEnv) GetPassword() string {
	return c.Password
}

func (c *DatabaseConfigEnv) GetName() string {
	return c.Name
}

func (c *DatabaseConfigEnv) GetSSLMode() string {
	return c.SSLMode
}

func (c *DatabaseConfigEnv) GetConnectString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), c.GetName(), c.GetSSLMode())
}

func ConnectToDB(dbConfig *DatabaseConfigEnv) (*sql.DB, error) {
	psqlInfo := dbConfig.GetConnectString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("ConnectToDB: %s", err)
	}
	return db, nil
}
