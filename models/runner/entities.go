package runner

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

// Job represents a job within Runner
type Job struct {
	Id                  string
	AccountAlias        string
	Name                string
	Description         string
	ProductId           string
	Repository          *Repository `json:"repository,omitempty"`
	Hosts               []Host
	UseDynamicInventory bool
	Properties          *map[string]interface{}
	Status              string
	Callbacks           *[]Callback `json:"callbacks,omitempty"`
	CreatedTime         base.JSTimeStamp
	LastUpdatedTime     base.JSTimeStamp
	//bootstrapKeyPairAlias
	//playbookTags
	//requirements
	ExecutionTtl int64
	//options
	//vaultPassword
	Links             models.Links `json:"links"`
	JobRunImmediately bool
}

// Repository represents a Git repository where the playbook is stored
type Repository struct {
	Credentials     Credential `json:"credentials,omitempty"`
	Url             string     `json:"url"`
	Branch          string     `json:"branch"`
	DefaultPlaybook string     `json:"defaultPlaybook"`
}

// Credential contains the username/password to use when connecting to a private Git repository
type Credential struct {
	Username string
	Password string
}

// Host represents a host that is available to the playbook when the play or task is executed
type Host struct {
	Id       string
	HostVars map[string]string
}

// Callback defines webhook urls that will be called at various stages of eexecution
type Callback struct {
	Url     string
	Level   string
	Type    string
	To      []string
	Subject string
}
