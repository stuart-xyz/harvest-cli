package services

type Config struct {
	HomeDir          *string
	JiraEndpoint     string `yaml:"jira-endpoint"`
	JiraEmail        string `yaml:"jira-email"`
	JiraApiToken     string `yaml:"jira-api-token"`
	HarvestApiToken  string `yaml:"harvest-personal-access-token"`
	HarvestAccountId int    `yaml:"harvest-account-id"`
}
