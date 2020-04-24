// credit - go-graphql hello world example
package github

import (
	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	//"go/types"
	"golang.org/x/net/context"
	"os"
)

//TODO: ask for token in gui for admin or maybe get token from user credentials
//TODO: save token in file and read from file on startup
//TODO: error msg if no token
const GithubToken = "43779c73fba2eff18728728abb10b0561d90ef81"

var client *githubv4.Client
func init() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client = githubv4.NewClient(httpClient)
	// Use client...
}

type commit struct {
	Author shortUser
	Committer shortUser
}
type issue struct {
	Number			githubv4.Int
	Title			githubv4.String
	Body           githubv4.String
	State			githubv4.IssueState
	ViewerCanUpdate githubv4.Boolean
}
var allOrganizations []organization
type organization struct {
	Name 	githubv4.String
	Login githubv4.String
	URL			githubv4.URI
	ViewerCanAdminister githubv4.Boolean
}
var currentOrganization organizationTeams
type organizationTeams struct {
	Organization struct {
		Teams struct {
			TotalCount githubv4.Int
			Nodes      []shortTeam
			PageInfo   pageInfo
		} `graphql:"teams(first:$teamFirst,after:$teamAfter)"`
	}`graphql:"organization(login:$login)"`
}
var orgaVariables = map[string]interface{}{
	"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"teamFirst": githubv4.NewInt(100),
	"teamAfter": (*githubv4.String)(nil),
}
//var currentPageInfo pageInfo
type pageInfo struct {
	StartCursor githubv4.String
	HasPreviousPage githubv4.Boolean
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}
/*The GraphQL API v4 rate limit is 5,000 points per hour.
Note that 5,000 points per hour is not the same as 5,000 calls per hour:
the GraphQL API v4 and REST API v3 use different rate limits.
 */
type rateLimit struct {
	Cost      githubv4.Int
	Limit     githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}
//var currentRepository repository
type repository struct {
	Name githubv4.String
	Owner shortUser
	CreatedAt githubv4.DateTime
	Description githubv4.String
	URL        githubv4.URI
	IsPrivate githubv4.Boolean
}
//Branch
//type ref struct {
//	Name githubv4.String
//	Target struct {
//		Node []commit
//		Commit struct {
//			TotalCount githubv4.Int
//		}
//	}'graphql:"history(first: 0)"'
//}
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
var allIssuesAssigned []issue
var currentRepository repositoryInfo
type repositoryInfo struct {
	Repository struct {
		Issues struct {
			TotalCount githubv4.Int
			Nodes []issue
			PageInfo pageInfo
		}`graphql:"issues(first:$issueFirst,after:$issueAfter,filterBy:{assignee:$assignee},states:[$issueState])"` //,states:$issueState
		//Refs struct {
		//	TotalCount githubv4.Int
		//	Nodes []ref
		//	PageInfo pageInfo
		//}`graphql:"ref(refPrefix:$prefix,first:$refFirst,after:$refAfter,orderBy:$orderBy)"`
		//Commits struct {
		//	TotalCount githubv4.Int
		//	Nodes []commit
		//}
		//DefaultBranchRef ref
	} `graphql:"repository(owner:$login,name:$repositoryName)"`
}
var repoVariables = map[string]interface{}{
	"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"repositoryName": (*githubv4.String)(nil), //githubv4.String("project-Tide"),
	"assignee": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"issueState": githubv4.IssueStateOpen,
	"issueFirst": githubv4.Int(1),
	"issueAfter": (*githubv4.String)(nil),
	"commentsFirst": githubv4.Int(1),
	"commentsAfter": (*githubv4.String)(nil),
	//"prefix": githubv4.String("refs/heads/"),
	//"refFirst": githubv4.NewInt(1),
	//"refAfter": (*githubv4.String)(nil),
	//"orderBy": githubv4.RefOrder{githubv4.RefOrderFieldTagCommitDate,githubv4.OrderDirectionDesc },
}

var allTeams []shortTeam
type shortTeam struct {
	Name githubv4.String
	Slug githubv4.String
	CombinedSlug githubv4.String
	Description         githubv4.String
	RepositoriesUrl githubv4.URI
	Privacy             githubv4.TeamPrivacy
	ViewerCanAdminister githubv4.Boolean
}
var allTeamMembersAndRepos = team{}
var currentTeam team
type team struct {
	Organization struct{
		Team struct {
			Members struct {
				TotalCount githubv4.Int
				Nodes      []shortUser
				PageInfo   pageInfo
			} `graphql:"members(first:$teamMembersFirst,after:$teamMembersAfter)"`
			Repositories struct {
				TotalCount githubv4.Int
				Nodes      []repository
				PageInfo   pageInfo
			} `graphql:"repositories(first:$repositoryFirst,after:$repositoryAfter)"`
			RepositoriesUrl githubv4.URI
		}`graphql:"team(slug:$teamName)"`
	}`graphql:"organization(login:$login)"`
}
var teamVariables = map[string]interface{}{
	"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"teamName": (*githubv4.String)(nil), //githubv4.String("A-Team"),
	"teamMembersFirst": githubv4.NewInt(1),
	"teamMembersAfter": (*githubv4.String)(nil),
	"repositoryFirst": githubv4.NewInt(1),
	"repositoryAfter": (*githubv4.String)(nil),
}
type shortUser struct {
	Login     githubv4.String
	URL       githubv4.URI
}
var currentUser user
type user struct {
	User struct {
		Organizations struct {
			TotalCount githubv4.Int
			Nodes []organization
			PageInfo pageInfo
		} `graphql:"organizations(first:$organizationFirst,after:$organizationAfter)"`
	} `graphql:"user(login:$login)"`
}
var userVariables = map[string]interface{}{
	"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"organizationFirst": githubv4.NewInt(1),
	"organizationAfter": (*githubv4.String)(nil),
}

var currentViewer viewer
type viewer struct {
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

func PrintJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func fetchData(client *githubv4.Client, query interface{}, localVar *map[string]interface{}) error {
	fmt.Println("\tin run")
	switch currentQuery := query.(type) {
	case *organizationTeams:
		fmt.Println("\t\tin Organization")
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

	case *repositoryInfo:
		fmt.Println("\t\tin Repo")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\t\tQuery repository failed with:")
				return err
			}
			allIssuesAssigned = append(allIssuesAssigned, currentQuery.Repository.Issues.Nodes...)

			if !currentQuery.Repository.Issues.PageInfo.HasNextPage {
				break
			}
			repoVariables["issueAfter"] = githubv4.NewString(currentQuery.Repository.Issues.PageInfo.EndCursor)
		}
		return nil

	case *team:
		fmt.Println("\t\tin Team")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\t\tQuery team in line 1 failed with")
				return err
			}

			fmt.Println("\t\tin Team2")
			//to avoid appending empty node or redundant node
			fmt.Println((githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Members.Nodes)))
			if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Members.Nodes)) != currentQuery.Organization.Team.Members.TotalCount {
				allTeamMembersAndRepos.Organization.Team.Members.Nodes = append(allTeamMembersAndRepos.Organization.Team.Members.Nodes, currentQuery.Organization.Team.Members.Nodes...)
				fmt.Println("\t\t\tin Team2")
				fmt.Println(allTeamMembersAndRepos)
			}
			if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes)) != currentQuery.Organization.Team.Repositories.TotalCount {
				allTeamMembersAndRepos.Organization.Team.Repositories.Nodes = append(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes, currentQuery.Organization.Team.Repositories.Nodes...)
			}
			fmt.Println("\t\tin Team3")
			if !currentQuery.Organization.Team.Repositories.PageInfo.HasNextPage && !currentQuery.Organization.Team.Members.PageInfo.HasNextPage{
				fmt.Println("\t\tnobody has next page")
				break
			}
			if currentQuery.Organization.Team.Members.PageInfo.HasNextPage {
				fmt.Println("\t\tmembers have next")
				teamVariables["teamMembersAfter"] = githubv4.NewString(currentQuery.Organization.Team.Members.PageInfo.EndCursor)
			}
			if currentQuery.Organization.Team.Repositories.PageInfo.HasNextPage{
				fmt.Println("\t\t repos have next")
				teamVariables["repositoryAfter"] = githubv4.NewString(currentQuery.Organization.Team.Repositories.PageInfo.EndCursor)
			}
		}
		fmt.Println("\t\texit team")
		return nil

	case *user:
		fmt.Println("\t\tin User")
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
		fmt.Println("\t\tend User")
		return nil

	case *viewer:
		fmt.Println("\t\tin Viewer")
		err := client.Query(context.Background(), &currentQuery, nil)
		if err != nil {
			fmt.Println("\tQuery viewer failed with:")
			return err
		}
		return nil

	default:
		fmt.Println("\t\tin default")
		return fmt.Errorf("something went wrong with the query %s", currentQuery)
	}
}

func GetViewer() *viewer {
	fmt.Println("in GetViewer")
	err := fetchData(client, &currentViewer, nil)
	if err != nil {
		fmt.Println(err)
	}
	return &currentViewer
}
func GetOrganizations(ownerLogin githubv4.String) *[]organization {
	fmt.Println("in GetOrganizations")
	userVariables["login"] = ownerLogin
	err := fetchData(client, &currentUser, &userVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allOrganizations
}
func GetTeamsPerOrganization(organizationLogin githubv4.String) *[]shortTeam {
	fmt.Println("in GetTeamsPerOrganization")
	orgaVariables["login"] = organizationLogin
	err := fetchData(client, &currentOrganization, &orgaVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allTeams
}
func GetTeamMembersAndRepositories(organizationLogin githubv4.String, teamName githubv4.String) *team{
	fmt.Println("in GetTeamMembersAndRepositories")
	teamVariables["login"] = organizationLogin
	teamVariables["teamName"] = teamName
	err := fetchData(client, &currentTeam, &teamVariables)
	if err != nil {
		fmt.Println(err)
	}
	PrintJSON(allTeamMembersAndRepos)
	fmt.Println("exit GetTeamMembersAndRepositories")
	return &allTeamMembersAndRepos
}
func GetRepositoryInfo(repositoryName githubv4.String, ownerLogin githubv4.String, assignee githubv4.String) (*[]issue) {
	fmt.Println("in GetRepositoryInfo")
	repoVariables["repositoryName"] = repositoryName
	repoVariables["login"] = ownerLogin
	repoVariables["assignee"] = assignee
	err := fetchData(client, &currentRepository, &repoVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allIssuesAssigned//, &commitCountPerUser, &codeCoverage
}



//There are these ReactionGroups aka emojis, which might be interesting at some point
//
//ReactionGroups []struct {
//	Content githubv4.ReactionContent
//	Users   struct {
//		Nodes []struct {	//User
//			Login githubv4.String
//		}
//		TotalCount githubv4.Int
//	} `graphql:"users(first:10)"`
//	ViewerHasReacted githubv4.Boolean
//}