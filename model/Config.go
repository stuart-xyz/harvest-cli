package model

type Config struct {
	JiraEndpoint     string `yaml:"endpoint"`
	HarvestApiToken  string `yaml:"harvest-personal-access-token"`
	HarvestAccountId int    `yaml:"harvest-account-id"`
}
