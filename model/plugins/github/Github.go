// credit - go-graphql hello world example
package main

import (
	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"github/SashaCollins/Wisehub-Connect/model/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"os"
)

//TODO: ask for token in gui for admin or maybe get token from user credentials
//TODO: save token in file and read from file on startup
//TODO: error msg if no token

type Github struct {}

var (
	GithubToken string
	githubClient *githubv4.Client
	//lock sync.Mutex

	currentViewer Viewer
	currentUser   User

	currentOrganization OrganizationTeams
	allOrganizations    []Organization

	currentTeam            Team
	allTeams               []ShortTeam
	allTeamMembersAndRepos = Team{}

	currentRepository RepositoryInfo
	allRefs           []Ref
	allIssuesAssigned []Issue

	userVariables = map[string]interface{}{
		"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"organizationFirst": githubv4.NewInt(1),
		"organizationAfter": (*githubv4.String)(nil),
	}
	orgaVariables = map[string]interface{}{
		"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"teamFirst": githubv4.NewInt(100),
		"teamAfter": (*githubv4.String)(nil),
	}
	teamVariables = map[string]interface{}{
		"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"teamName": (*githubv4.String)(nil), //githubv4.String("A-Team"),
		"teamMembersFirst": githubv4.NewInt(1),
		"teamMembersAfter": (*githubv4.String)(nil),
		"repositoryFirst": githubv4.NewInt(1),
		"repositoryAfter": (*githubv4.String)(nil),
	}
	repoVariables = map[string]interface{}{
		"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"repositoryName": (*githubv4.String)(nil), //githubv4.String("project-Tide"),
		"assignee": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"issueState": githubv4.IssueStateOpen,
		"issueFirst": githubv4.NewInt(1),
		"issueAfter": (*githubv4.String)(nil),
		//"refName": githubv4.String("commit"),
		"prefix": githubv4.String("refs/heads/"),
		//"target": githubv4.String("Commit"),
		"refFirst": githubv4.NewInt(10),
		"refAfter": (*githubv4.String)(nil),
		"targetFirst": githubv4.NewInt(10),
		"orderBy": githubv4.RefOrder{githubv4.RefOrderFieldTagCommitDate,githubv4.OrderDirectionDesc },
	}
)

func init() {
	//create an http client with oauth authentication
	conf := config.GetConfig()
	GithubToken = conf.GitHub.APIToken
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	githubClient = githubv4.NewClient(httpClient)
	// Use client...
}

type Commit struct {
	Author    ShortUser
	Committer ShortUser
}

type Issue struct {
	Number			githubv4.Int
	Title			githubv4.String
	Body           githubv4.String
	State			githubv4.IssueState
	ViewerCanUpdate githubv4.Boolean
}

type Organization struct {
	Name 	githubv4.String
	Login githubv4.String
	URL			githubv4.URI
	ViewerCanAdminister githubv4.Boolean
}

type OrganizationTeams struct {
	Organization struct {
		Teams struct {
			TotalCount githubv4.Int
			Nodes      []ShortTeam
			PageInfo   PageInfo
		} `graphql:"teams(first:$teamFirst,after:$teamAfter)"`
	}`graphql:"organization(login:$login)"`
}

//var currentPageInfo pageInfo
type PageInfo struct {
	StartCursor githubv4.String
	HasPreviousPage githubv4.Boolean
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}

/*The GraphQL API v4 rate limit is 5,000 points per hour.
Note that 5,000 points per hour is not the same as 5,000 calls per hour:
the GraphQL API v4 and REST API v3 use different rate limits.
 */
type RateLimit struct {
	Cost      githubv4.Int
	Limit     githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}

//var currentRepository repository
type Repository struct {
	Name  githubv4.String
	Owner ShortUser
	URL   githubv4.URI
}

//Branch
type Ref struct {
	Name githubv4.String
	Prefix githubv4.String

	Target struct {
		AbbreviatedOid githubv4.String
		ID githubv4.GitObjectID
		//History struct {
		//	TotalCount githubv4.Int
		//}`graphql:"history(first:0)"`
	}`graphql:"... on Commit"`
}
	//`graphql:"target(first:$targetFirst)"`


	//`graphql:"... on Commit"`
	//repository(owner: "bertrandmartel", name: "callflow-workshop") {
	//	refs(refPrefix: "refs/heads/", orderBy: {direction: DESC, field: TAG_COMMIT_DATE}, first: 100) {
	//		edges {
	//			node {
	//				... on Ref {
	//					name
	//					target {
	//						... on Commit {
	//							history(first: 2) {
	//								edges {
	//									node {
	//										... on Commit {
	//											committedDate

type RepositoryInfo struct {
	Repository struct {
		Owner       ShortUser
		CreatedAt   githubv4.DateTime
		Description githubv4.String
		IsPrivate   githubv4.Boolean
		Issues      struct {
			TotalCount githubv4.Int
			Nodes      []Issue
			PageInfo   PageInfo
		}`graphql:"issues(first:$issueFirst,after:$issueAfter,filterBy:{assignee:$assignee},states:[$issueState])"` //,states:$issueState
		Refs struct {
			TotalCount githubv4.Int //number of branches
			Nodes      []Ref
			PageInfo   PageInfo
		}`graphql:"refs(refPrefix:$prefix,first:$refFirst,after:$refAfter,orderBy:$orderBy)"`
	} `graphql:"repository(owner:$login,name:$repositoryName)"`
}

//Commits struct {
//	TotalCount githubv4.Int
//	Nodes []commit
//}
//DefaultBranchRef ref

type ShortTeam struct {
	Slug githubv4.String
}

type Team struct {
	Organization struct{
		Team struct {
			Name githubv4.String
			CombinedSlug githubv4.String
			Description         githubv4.String
			Privacy             githubv4.TeamPrivacy
			ViewerCanAdminister githubv4.Boolean
			Members struct {
				TotalCount githubv4.Int
				Nodes      []ShortUser
				PageInfo   PageInfo
			} `graphql:"members(first:$teamMembersFirst,after:$teamMembersAfter)"`
			Repositories struct {
				TotalCount githubv4.Int
				Nodes      []Repository
				PageInfo   PageInfo
			} `graphql:"repositories(first:$repositoryFirst,after:$repositoryAfter)"`
			RepositoriesUrl githubv4.URI
		}`graphql:"team(slug:$teamName)"`
	}`graphql:"organization(login:$login)"`
}

type ShortUser struct {
	Login     githubv4.String
	URL       githubv4.URI
}

type User struct {
	User struct {
		Organizations struct {
			TotalCount githubv4.Int
			Nodes      []Organization
			PageInfo   PageInfo
		} `graphql:"organizations(first:$organizationFirst,after:$organizationAfter)"`
	} `graphql:"user(login:$login)"`
}

type Viewer struct {
	Viewer struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		URL 		githubv4.URI
	}
}

//var listRepos(&queryString: String!) struct {
//	rateLimit{
//		cost
//		remaining
//		resetAt
//	}
//	search(query: &queryString, type:REPOSITORY, first:20){
//		repositoryCount
//		pageInfo{
//			endCursor
//			startCursor
//		}
//		edges{
//			node{
//				... on Repository{
//					id
//					name
//					createdAt
//					description
//					isArchived
//					isPrivate
//					url
//					owner{}
//
//					defaultBranchRef{
//						target{
//							... on Commit{
//								history(first:10){
//									totalCount
//									edges{
//										node{
//											... on Commit{
//												committedDate

func (gr *Github) printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func (gr *Github) fetchData(client *githubv4.Client, query interface{}, localVar *map[string]interface{}) error {
	//fmt.Println("\tin run")
	switch currentQuery := query.(type) {
	case *OrganizationTeams:
		//fmt.Println("\t\tin Organization")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\t\tQuery orga in line 1 failed with")
				return err
			}
			allTeams = append(allTeams, currentQuery.Organization.Teams.Nodes...)

			if !currentQuery.Organization.Teams.PageInfo.HasNextPage {
				break
			}
			orgaVariables["teamAfter"] = githubv4.NewString(currentQuery.Organization.Teams.PageInfo.EndCursor)
		}
		return nil

	case *RepositoryInfo:
		//fmt.Println("\t\tin Repo")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\t\tQuery repository failed with:")
				return err
			}
			allIssuesAssigned = append(allIssuesAssigned, currentQuery.Repository.Issues.Nodes...)
			allRefs = append(allRefs, currentQuery.Repository.Refs.Nodes...)
			if !currentQuery.Repository.Issues.PageInfo.HasNextPage {
				break
			}
			repoVariables["issueAfter"] = githubv4.NewString(currentQuery.Repository.Issues.PageInfo.EndCursor)
		}
		return nil

	case *Team:
		//fmt.Println("\t\tin Team")
		firstLoop := true
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\t\tQuery team in line 1 failed with")
				return err
			}
			if firstLoop {
				allTeamMembersAndRepos = currentTeam
				firstLoop = false
			}else {
				//to avoid appending redundant node
				if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Members.Nodes)) != currentQuery.Organization.Team.Members.TotalCount {
					allTeamMembersAndRepos.Organization.Team.Members.Nodes = append(allTeamMembersAndRepos.Organization.Team.Members.Nodes, currentQuery.Organization.Team.Members.Nodes...)

				}
				if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes)) != currentQuery.Organization.Team.Repositories.TotalCount {
					allTeamMembersAndRepos.Organization.Team.Repositories.Nodes = append(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes, currentQuery.Organization.Team.Repositories.Nodes...)
				}
			}
			if !currentQuery.Organization.Team.Repositories.PageInfo.HasNextPage && !currentQuery.Organization.Team.Members.PageInfo.HasNextPage{
				//fmt.Println("\t\tnobody has next page")
				break
			}
			if currentQuery.Organization.Team.Members.PageInfo.HasNextPage {
				//fmt.Println("\t\tmembers have next")
				teamVariables["teamMembersAfter"] = githubv4.NewString(currentQuery.Organization.Team.Members.PageInfo.EndCursor)
			}
			if currentQuery.Organization.Team.Repositories.PageInfo.HasNextPage{
				//fmt.Println("\t\t repos have next")
				teamVariables["repositoryAfter"] = githubv4.NewString(currentQuery.Organization.Team.Repositories.PageInfo.EndCursor)
			}
		}
		//fmt.Println("\t\texit team")
		return nil

	case *User:
		//fmt.Println("\t\tin User")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery user in line 1 failed with")
				return err
			}
			allOrganizations = append(allOrganizations, currentQuery.User.Organizations.Nodes...)

			if !currentQuery.User.Organizations.PageInfo.HasNextPage {
				break
			}
			userVariables["organizationAfter"] = githubv4.NewString(currentQuery.User.Organizations.PageInfo.EndCursor)
		}
		//fmt.Println("\t\tend User")
		return nil

	case *Viewer:
		//fmt.Println("\t\tin Viewer")
		err := client.Query(context.Background(), &currentQuery, nil)
		if err != nil {
			fmt.Println("\tQuery viewer failed with:")
			return err
		}
		return nil

	default:
		//fmt.Println("\t\tin default")
		return fmt.Errorf("something went wrong with the query %s", currentQuery)
	}
}

func (gr *Github) getViewer() (*Viewer, error) {
	//fmt.Println("in GetViewer")
	err := gr.fetchData(githubClient, &currentViewer, nil)
	if err != nil {
		return nil, err
	}
	return &currentViewer, nil
}

func (gr *Github) getOrganizations(ownerLogin githubv4.String) (*[]Organization, error) {
	//fmt.Println("in GetOrganizations")
	userVariables["login"] = ownerLogin
	err := gr.fetchData(githubClient, &currentUser, &userVariables)
	if err != nil {
		return nil, err
	}
	return &allOrganizations, nil
}

func (gr *Github) getTeamsPerOrganization(organizationLogin githubv4.String) (*[]ShortTeam, error) {
	//fmt.Println("in GetTeamsPerOrganization")
	orgaVariables["login"] = organizationLogin
	err := gr.fetchData(githubClient, &currentOrganization, &orgaVariables)
	if err != nil {
		return nil, err
	}
	return &allTeams, nil
}

func (gr *Github) getTeamMembersAndRepositories(organizationLogin githubv4.String, teamName githubv4.String) (*Team, error){
	//fmt.Println("in GetTeamMembersAndRepositories")
	teamVariables["login"] = organizationLogin
	teamVariables["teamName"] = teamName
	err := gr.fetchData(githubClient, &currentTeam, &teamVariables)
	if err != nil {
		return nil, err
	}
	gr.printJSON(allTeamMembersAndRepos)
	return &allTeamMembersAndRepos, nil
}

func (gr *Github) getRepositoryInfo(repositoryName githubv4.String, ownerLogin githubv4.String, assignee githubv4.String) (*[]Issue, *[]Ref, error) {
	//fmt.Println("in GetRepositoryInfo")
	repoVariables["repositoryName"] = repositoryName
	repoVariables["login"] = ownerLogin
	repoVariables["assignee"] = assignee
	err := gr.fetchData(githubClient, &currentRepository, &repoVariables)
	if err != nil {
		return nil, nil, err
	}
	return &allIssuesAssigned, &allRefs, nil //, &commitCountPerUser, &codeCoverage
}

func (gr *Github) GetOrgaInfo() (string,  *[]Organization) {
	viewer, err := gr.getViewer()
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil
	}
	allOrgas, err := gr.getOrganizations(viewer.Viewer.Login)
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil
	}
	return "Github", allOrgas
}

func (gr *Github) GetTeamInfo(orgaName string) (string, *[]ShortTeam) {
	allTeams, err := gr.getTeamsPerOrganization((githubv4.String)(orgaName))
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil
	}
	return "Github", allTeams
}

func (gr *Github) GetInsightTeamInfo(orgaName, teamName string) (string, *Team) {
	allTeamMembersAndRepos, err := gr.getTeamMembersAndRepositories((githubv4.String)(orgaName), (githubv4.String)(teamName))
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil
	}
	return "Github", allTeamMembersAndRepos
}

func (gr *Github) GetTeamRepoInfo(repoName, repoOwner string) (string, *[]Issue, *[]Ref) {
	viewer, err := gr.getViewer()
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil, nil
	}
	allIssuesAssigned, allCommits, err := gr.getRepositoryInfo((githubv4.String)(repoName), (githubv4.String)(repoOwner), viewer.Viewer.Login)
	if err != nil {
		fmt.Println("error:", err)
		return "Github", nil, nil
	}
	return "Github", allIssuesAssigned, allCommits
}