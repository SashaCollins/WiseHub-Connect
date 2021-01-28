package continuous_integration

type ContinuousIntegration interface{
	GetRepositories() (interface{}, error)
	GetBuilds() (interface{}, error)
}
