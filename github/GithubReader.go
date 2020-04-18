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
type comment struct {
	Body   githubv4.String
	//Author struct {
	//	Login githubv4.String
	//}
	//Editor struct {
	//	Login githubv4.String
	//}
	//ViewerCanReact bool
}
type count struct {
	User struct {
		Organizations struct {
			Nodes []struct {
				MembersWithRole struct {
					TotalCount githubv4.Int
				}
				Teams struct {
					Nodes []struct {
						Members struct {
							TotalCount githubv4.Int
						}
					}
				}`graphql:"teams(first:$teamFirst,after:$teamAfter)"`
			}
		}`graphql:"organizations(first:$organizationFirst,after:$organizationAfter)"`
	}`graphql:"user(login:$viewer)"`
}
type issue struct {
	Number			githubv4.Int
	Title			githubv4.String
	Body           githubv4.String
	State			githubv4.IssueState
	ViewerCanUpdate githubv4.Boolean
	Comments struct {
		TotalCount githubv4.Int
		Nodes []comment
		PageInfo pageInfo
	} `graphql:"comments(first:$commentsFirst,after:$commentsAfter)"`
}
var allOrganizations []shortOrganization
type shortOrganization struct {
	Name 	githubv4.String
	Login githubv4.String
	URL			githubv4.URI
	ViewerCanAdminister githubv4.Boolean
}
var currentOrganization organization
type organization struct {
	Organization struct {
		Name                githubv4.String
		Login               githubv4.String
		URL                 githubv4.URI
		ViewerCanAdminister githubv4.Boolean
		MembersWithRole     struct {
			TotalCount githubv4.Int
			Nodes      []shortUser
			PageInfo   pageInfo
		} `graphql:"membersWithRole(first:$orgaMembersFirst)"`
		Teams struct {
			TotalCount githubv4.Int
			Nodes      []shortTeam
			PageInfo   pageInfo
		} `graphql:"teams(first:$teamFirst,after:$teamAfter)"`
	}`graphql:"organization(owner:$viewer,name:$organizationName)"`
}
var orgaVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"organizationName": (*githubv4.String)(nil), //githubv4.String("WiseHub-Connector"),
	"orgaMembersFirst": githubv4.NewInt(100),
	"teamFirst": githubv4.NewInt(1),
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
var allRepos []shortRepository
type shortRepository struct {
	Name githubv4.String
	NameWithOwner githubv4.String
	CreatedAt githubv4.DateTime
}
//Branch
type ref struct {
	Name githubv4.String
	Target struct {
		Node []commit
		Commit struct {
			TotalCount githubv4.Int
		}
	}'graphql:"history(first: 0)"'
}
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
var allIssues []issue
var currentRepository repository
type repository struct {
	Repository struct {
		ID githubv4.String
		NameWithOwner githubv4.String
		CreatedAt githubv4.DateTime
		Description githubv4.String
		URL        githubv4.URI
		IsPrivate githubv4.Boolean

		Issues struct {
			TotalCount githubv4.Int
			Nodes []issue
			PageInfo pageInfo
		}`graphql:"issue(filterBy:$assignee,states:$issueState)"`
		Refs struct {
			TotalCount githubv4.Int
			Nodes []ref
			PageInfo pageInfo
		}`graphql:"ref(refPrefix:$prefix,first:$refFirst,after:$refAfter,orderBy:$orderBy)"`
		//Commits struct {
		//	TotalCount githubv4.Int
		//	Nodes []commit
		//}
		//DefaultBranchRef ref
	} `graphql:"repository(owner:$viewer,name:$repositoryName)"`
}
var repoVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"repositoryName": (*githubv4.String)(nil), //githubv4.String("project-Tide"),
	"assignee": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"issueState": nil, //githubv4.IssueStateOpen,
	"prefix": githubv4.String("refs/heads/"),
	"refFirst": githubv4.NewInt(1),
	"refAfter": (*githubv4.String)(nil),
	"orderBy": githubv4.RefOrder{githubv4.RefOrderFieldTagCommitDate,githubv4.OrderDirectionDesc },
}

var allTeams []shortTeam
type shortTeam struct {
		Name githubv4.String
		CombinedSlug githubv4.String
		URL githubv4.URI
		RepositoriesUrl githubv4.URI
}
var allFullTeams []team
var currentTeam team
type team struct {
	Team struct {
		Name                githubv4.String
		CombinedSlug        githubv4.String
		CreatedAt           githubv4.DateTime
		Description         githubv4.String
		Privacy             githubv4.TeamPrivacy
		ViewerCanAdminister githubv4.Boolean
		Members             struct {
			TotalCount githubv4.Int
			Nodes      []shortUser
			PageInfo	pageInfo
		} `graphql:"members(first:$teamMembersFirst)"`
		Repositories struct {
			TotalCount githubv4.Int
			Nodes      []shortRepository
			PageInfo   pageInfo
		} `graphql:"repositories(first:$repositoryFirst,after:$repositoryAfter)"`
		RepositoriesUrl githubv4.URI
	}`graphql:"team(name:$teamName)"`
}
var teamVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"teamName": (*githubv4.String)(nil), //githubv4.String("A-Team"),
	"teamMembersFirst": githubv4.NewInt(100),
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
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		Organizations struct {
			Nodes []shortOrganization
			TotalCount githubv4.Int
			PageInfo pageInfo
		} `graphql:"organizations(first:$organizationFirst,after:$organizationAfter)"`
	} `graphql:"user(login:$viewer)"`
}
var userVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
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

var countVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"organizationFirst": githubv4.NewInt(1),
	"organizationAfter": (*githubv4.String)(nil),
	"teamFirst": githubv4.NewInt(1),
	"teamAfter": (*githubv4.String)(nil),
}

func PrintJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
func fetchData(client *githubv4.Client, query interface{}, localVar *map[string]interface{}) error {
	fmt.Println("in run")
	switch currentQuery := query.(type) {
	case *organization:
		fmt.Println("\tin Organization")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery orga in line 1 failed with")
				return err
			}
			//if there are more than 100 Members do a reload
			if currentQuery.Organization.MembersWithRole.PageInfo.HasNextPage {
				orgaVariables["orgaMembersFirst"] = githubv4.NewInt(currentQuery.Organization.MembersWithRole.TotalCount)
				err := client.Query(context.Background(), &currentQuery, *localVar)
				if err != nil {
					fmt.Println("\tQuery orga in line 2 failed with")
					return err
				}
			}
			allTeams = append(allTeams, currentQuery.Organization.Teams.Nodes...)

			if !currentQuery.Organization.Teams.PageInfo.HasNextPage {
				break
			}
			orgaVariables["teamAfter"] = githubv4.NewString(currentQuery.Organization.Teams.PageInfo.EndCursor)
		}
		return nil

	case *repository:
		fmt.Println("\tin Repo")
		//for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery repository failed with:")
				return err
			}
			allIssues = append(allIssues, currentQuery.Repository.Issues.Nodes...)

			//if !currentQuery.Repository.Issues.PageInfo.HasNextPage {
			//	break
			//}
			//repoVariables["issueAfter"] = githubv4.NewString(currentQuery.Repository.Issues.PageInfo.EndCursor)
		//}
		//printJSON(allComments)
		return nil

	case *team:
		fmt.Println("\tin Team")
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery orga in line 1 failed with")
				return err
			}
			//if there are more than 100 Members do a reload
			if currentQuery.Team.Members.PageInfo.HasNextPage {
				teamVariables["teamMembersFirst"] = githubv4.NewInt(currentQuery.Team.Members.TotalCount)
				err := client.Query(context.Background(), &currentQuery, *localVar)
				if err != nil {
					fmt.Println("\tQuery orga in line 2 failed with")
					return err
				}
			}
			allRepos = append(allRepos, currentQuery.Team.Repositories.Nodes...)
			if !currentQuery.Team.Repositories.PageInfo.HasNextPage {
				break
			}
			teamVariables["teamAfter"] = githubv4.NewString(currentQuery.Team.Repositories.PageInfo.EndCursor)
		}
		allFullTeams = append(allFullTeams, *currentQuery)
		return nil

	case *user:
		fmt.Println("\tin User")
		for {
			err := client.Query(context.Background(), &currentQuery, countVariables)
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
		fmt.Println("\tend User")
		return nil

	case *viewer:
		fmt.Println("\tin Viewer")
		err := client.Query(context.Background(), &currentQuery, nil)
		if err != nil {
			fmt.Println("\tQuery viewer failed with:")
			return err
		}
		currentViewer.Viewer.Login = currentQuery.Viewer.Login
		currentViewer.Viewer.CreatedAt = currentQuery.Viewer.CreatedAt
		return nil

	default:
		fmt.Println("\tin default")
		return fmt.Errorf("something went wrong with the query %s", currentQuery)
		//fmt.Println(currentQuery)
		//err := client.Query(context.Background(), &currentQuery, *localVar)
		//if err != nil {
		//	fmt.Println("\tQuery default failed with:")
		//	return err
		//}
		//PrintJSON(currentQuery)
		//return nil
	}
}

func GetViewer() *viewer {
	err := fetchData(client, &currentViewer, nil)
	if err != nil {
		fmt.Println(err)
	}
	userVariables["viewer"] = currentViewer.Viewer.Login
	orgaVariables["viewer"] = currentViewer.Viewer.Login
	teamVariables["viewer"] = currentViewer.Viewer.Login
	repoVariables["viewer"] = currentViewer.Viewer.Login
	return &currentViewer
}
func GetOrganizations(ownerLogin string) * []shortOrganization {
	//userVariables["viewer"] = currentViewer.Viewer.Login
	err := fetchData(client, &currentUser, &userVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allOrganizations
}
func GetTeamsPerOrganization(organizationName string, ownerLogin string) *[]shortTeam {
	orgaVariables["organizationName"] = organizationName
	//orgaVariables["viewer"] = currentViewer.Viewer.Login
	err := fetchData(client, &currentOrganization, &orgaVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allTeams
}
func GetFullTeamInfo(teamName string, ownerLogin string) *[]team{
	teamVariables["teamName"] = teamName
	//teamVariables["viewer"] = currentViewer.Viewer.Login
	err := fetchData(client, &currentTeam, &teamVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allFullTeams
}
func GetRepositoriesPerTeam(teamName string, ownerLogin string) * []shortRepository {
	teamVariables["teamName"] = teamName
	//teamVariables["viewer"] = currentViewer.Viewer.Login
	err := fetchData(client, &currentTeam, &teamVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allRepos
}
func GetRepositoryIssues(repositoryName string, assignee string, ownerLogin string) *[]issue {
	repoVariables["repositoryName"] = repositoryName
	repoVariables["assignee"] = assignee
	//repoVariables["viewer"] = currentViewer.Viewer.Login
	err := fetchData(client, &currentRepository, &repoVariables)
	if err != nil {
		fmt.Println(err)
	}
	return &allIssues
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