package testing_tools

type TestingTools interface{
	GetRepositories() (interface{}, error)
	GetBuilds() (interface{}, error)
}
