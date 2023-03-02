package pipelines

type PacMetadata struct {
	RegistryUsername string
	RegistryPassword string
	RegistryServer   string

	PersonalAccessToken string
	WebhookSecret       string

	ConfigureLocalResources   bool
	ConfigureClusterResources bool
	ConfigureRemoteResources  bool
}
