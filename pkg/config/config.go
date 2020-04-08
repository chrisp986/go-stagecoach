package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Config is used to initialize the connection string for the database
type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
}

type MailConfig struct {
	SmtpServer  string `yaml:"smtpServer"`
	Loginname   string `yaml:"loginname"`
	Password    string `yaml:"password"`
	FromAddress string `yaml:"fromAddress"`
}

//GetConfig is used to read the config file that stores the database information
func (d *DBConfig) ReadDBConfig() (dsn string) {

	yamlFile, err := ioutil.ReadFile(filepath.Join("configs", "db_config.yaml"))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if d.User == "" || d.IP == "" || d.DBName == "" || d.Password == "" || d.Port == "" {
		log.Fatalf("ERROR -> Field in db config file is empty")
	}

	return d.User + ":" + d.Password + "@/" + d.DBName + "?parseTime=true"
}

func (m *MailConfig) ReadMailConfig() *MailConfig {

	yamlFile, err := ioutil.ReadFile(filepath.Join("configs", "mail_config.yaml"))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if m.SmtpServer == "" || m.Loginname == "" || m.Password == "" || m.FromAddress == "" {
		log.Fatalf("ERROR -> Field in mail config file is empty")
	}
	return m
}
