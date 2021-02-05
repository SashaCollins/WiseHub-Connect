/*
Plugin for GitHub
fetches data from GitHub via graphQL API
@author SashaCollins
@version 1.0
*/
package main

import (
	"encoding/json"
	"github.com/shurcooL/githubv4"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var (
	PluginName    string
	GithubClient  *githubv4.Client
	CurrentViewer Viewer
)

func init() {
	PluginName = "Github"
}

type Github struct {}

func NewPlugin() plugins.PluginI {
	return &Github{}
}

func getPluginName() string {
	return PluginName
}

/*
These structs are used to fetch data from githubv4 API
They don't contain every information available,
for further information please read
https://docs.github.com/en/graphql
and github.com/shurcooL/githubv4
 */

//type Commit struct {
//	Author    ShortUser
//	Committer ShortUser
//}

type Credentials struct {
	UserName string
	Token string
}

type Issue struct {
	Number			githubv4.Int
	Title			githubv4.String
	//Body           githubv4.String
	//State			githubv4.IssueState
}

type Organization struct {
	Login githubv4.String
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

type Repository struct {
	Name  githubv4.String
	Owner ShortUser
	URL   githubv4.URI
}

//Branch
//type Ref struct {
//	Name githubv4.String
//	Prefix githubv4.String
//
//	Target struct {
//		AbbreviatedOid githubv4.String
//		ID githubv4.GitObjectID
//		//History struct {
//		//	TotalCount githubv4.Int
//		//}`graphql:"history(first:0)"`
//	}`graphql:"... on Commit"`
//}
	//`graphql:"target(first:$targetFirst)"`

// so sollte man an die commits eines Repos rankommen können
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
		//Owner       ShortUser
		//CreatedAt   githubv4.DateTime
		//Description githubv4.String
		//IsPrivate   githubv4.Boolean
		Issues      struct {
			TotalCount githubv4.Int
			Nodes      []Issue
			PageInfo   PageInfo
		}`graphql:"issues(first:$issueFirst,after:$issueAfter,filterBy:{assignee:$assignee},states:[$issueState])"` //,states:$issueState
		//Refs struct {
		//	TotalCount githubv4.Int //number of branches
		//	Nodes      []Ref
		//	PageInfo   PageInfo
		//}`graphql:"refs(refPrefix:$prefix,first:$refFirst,after:$refAfter,orderBy:$orderBy)"`
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
			//ViewerCanAdminister githubv4.Boolean
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




/*
This is what the response from FetchData looks for GitHub
*/
type Response struct {
	Organization struct{
		Login githubv4.String `json:"orgaName"`
		Teams []RTeams `json:"teams"`
	} `json:"organization"`
}
type RTeams struct{
	Name githubv4.String `json:"teamName"`
	Members []RMembers `json:"members"`
	Repositories []RRepositories `json:"repositories"`
}
type RMembers struct {
	Login     githubv4.String `json:"memberName"`
}
type RRepositories struct {
	Name  githubv4.String `json:"repoName"`
	URL   githubv4.URI `json:"repoUrl"`
	Issues []RIssues `json:"issues"`
}
type RIssues struct {
	Number			githubv4.Int `json:"issueNumber"`
	Title			githubv4.String `json:"issueTitle"`
}

func (g *Github) FetchData() (string, error) {
	var response []Response

	viewer, err := g.getViewer()
	if err != nil {
		log.Println(err)
		return "", err
	}
	allOrgas, err := g.getOrganizations(viewer.Viewer.Login)
	if err != nil {
		log.Println(err)
		return "", err
	}
	for _, orga := range *allOrgas {
		var resp Response
		resp.Organization.Login = orga.Login
		allTeams, err := g.getTeamsPerOrganization((githubv4.String)(orga.Login))
		if err != nil {
			log.Println(err)
			return "", err
		}
		for _, team := range *allTeams {
			var rt RTeams
			rt.Name = team.Slug
			//resp.Organization.Teams = append(resp.Organization.Teams, )
			//resp.Organization.Teams[j].Name = team.Slug
			allTeamMembersAndRepos, err := g.getTeamMembersAndRepositories((githubv4.String)(orga.Login), (githubv4.String)(team.Slug))
			if err != nil {
				log.Println(err)
				return "", err
			}
			for _, member := range allTeamMembersAndRepos.Organization.Team.Members.Nodes {
				var rm RMembers
				rm.Login = member.Login
				rt.Members = append(rt.Members, rm)
			}
			for _, repo := range allTeamMembersAndRepos.Organization.Team.Repositories.Nodes {
				var rr RRepositories
				rr.Name = repo.Name
				rr.URL = repo.URL

				allIssuesAssigned, err := g.getRepositoryInfo((githubv4.String)(repo.Name), (githubv4.String)(repo.Owner.Login), viewer.Viewer.Login)
				if err != nil {
					log.Println(err)
					return "", err
				}
				for _, issue := range *allIssuesAssigned {
					var ri RIssues
					ri.Number = issue.Number
					ri.Title = issue.Title
					rr.Issues = append(rr.Issues, ri)
				}
				rt.Repositories = append(rt.Repositories, rr)
			}
			resp.Organization.Teams = append(resp.Organization.Teams, rt)
		}
		response = append(response, resp)
	}
	r, _ := json.Marshal(response)
	return string(r), nil
}

func (g *Github) SubmitCredentials(_, token string) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	GithubClient = githubv4.NewClient(httpClient)
}

func (g *Github) FetchPluginName() string {
	return getPluginName()
}

/*
For debugging purposes,
to display readable data on console
 */
func (g *Github) printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
/*
Fetches the current User, the User with given credentials
 */
func (g *Github) getViewer() (*Viewer, error) {
	err := GithubClient.Query(context.Background(), &CurrentViewer, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &CurrentViewer, nil
}
/*
Fetches all organizations which the user with the name 'ownerLogin' is is a part of.
Input Parameters can be fetched from getViewer()
 */
func (g *Github) getOrganizations(ownerLogin githubv4.String) (*[]Organization, error) {
	var allOrganizations []Organization
	var user User
	variables := map[string]interface{} {
		"login": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
		"organizationFirst": githubv4.NewInt(1),
		"organizationAfter": (*githubv4.String)(nil),
	}
	variables["login"] = ownerLogin
	for {
		err := GithubClient.Query(context.Background(), &user, variables)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		allOrganizations = append(allOrganizations, user.User.Organizations.Nodes...)
		if !user.User.Organizations.PageInfo.HasNextPage {
			break
		}
		variables["organizationAfter"] = githubv4.NewString(user.User.Organizations.PageInfo.EndCursor)
	}
	return &allOrganizations, nil
}
/*
Fetches all teams from the organization with the name 'organizationLogin'
Input Parameters can be fetched from getOrganizations(ownerLogin githubv4.String) (*[]Organization, error)
 */
func (g *Github) getTeamsPerOrganization(organizationLogin githubv4.String) (*[]ShortTeam, error) {
	var organizationTeams OrganizationTeams
	var allTeams []ShortTeam
	variables := map[string]interface{} {
		"login": (*githubv4.String)(nil),
		"teamFirst": githubv4.NewInt(100),
		"teamAfter": (*githubv4.String)(nil),
	}
	variables["login"] = organizationLogin
	for {
		err := GithubClient.Query(context.Background(), &organizationTeams, variables)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		allTeams = append(allTeams, organizationTeams.Organization.Teams.Nodes...)

		if !organizationTeams.Organization.Teams.PageInfo.HasNextPage {
			break
		}
		variables["teamAfter"] = githubv4.NewString(organizationTeams.Organization.Teams.PageInfo.EndCursor)
	}
	return &allTeams, nil
}
/*
Fetches all team members and repositories from a given team 'teamName' in an organization 'organizationName'
Input Parameters can be fetched from getTeamsPerOrganization(organizationLogin githubv4.String) (*[]ShortTeam, error)
 */
func (g *Github) getTeamMembersAndRepositories(organizationLogin githubv4.String, teamName githubv4.String) (*Team, error){
	var team Team
	allTeamMembersAndRepos := Team{}
	variables := map[string]interface{} {
		"login": (*githubv4.String)(nil),
		"teamName": (*githubv4.String)(nil),
		"teamMembersFirst": githubv4.NewInt(1),
		"teamMembersAfter": (*githubv4.String)(nil),
		"repositoryFirst": githubv4.NewInt(1),
		"repositoryAfter": (*githubv4.String)(nil),
	}
	variables["login"] = organizationLogin
	variables["teamName"] = teamName
	firstLoop := true
	for {
		err := GithubClient.Query(context.Background(), &team, variables)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if firstLoop {
			allTeamMembersAndRepos = team
			firstLoop = false
		} else {
			//to avoid appending redundant node
			if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Members.Nodes)) != team.Organization.Team.Members.TotalCount {
				allTeamMembersAndRepos.Organization.Team.Members.Nodes = append(allTeamMembersAndRepos.Organization.Team.Members.Nodes, team.Organization.Team.Members.Nodes...)

			}
			if (githubv4.Int)(len(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes)) != team.Organization.Team.Repositories.TotalCount {
				allTeamMembersAndRepos.Organization.Team.Repositories.Nodes = append(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes, team.Organization.Team.Repositories.Nodes...)
			}
		}
		if !team.Organization.Team.Repositories.PageInfo.HasNextPage && !team.Organization.Team.Members.PageInfo.HasNextPage {
			break
		}
		if team.Organization.Team.Members.PageInfo.HasNextPage {
			variables["teamMembersAfter"] = githubv4.NewString(team.Organization.Team.Members.PageInfo.EndCursor)
		}
		if team.Organization.Team.Repositories.PageInfo.HasNextPage {
			variables["repositoryAfter"] = githubv4.NewString(team.Organization.Team.Repositories.PageInfo.EndCursor)
		}
	}
	return &allTeamMembersAndRepos, nil
}
/*
Fetches Information about a given Repository 'repositoryName' and the owners name 'ownerLogin'
Input Parameters can be fetched from getTeamMembersAndRepositories(organizationLogin githubv4.String, teamName githubv4.String) (*Team, error)
 */
func (g *Github) getRepositoryInfo(repositoryName githubv4.String, ownerLogin githubv4.String, assigneeLogin githubv4.String) (*[]Issue, error) {
	var repositoryInfo RepositoryInfo
	var allIssuesAssigned []Issue

	variables := map[string]interface{} {
		"login": (*githubv4.String)(nil),
		"repositoryName": (*githubv4.String)(nil),
		"assignee": (*githubv4.String)(nil),
		"issueState": githubv4.IssueStateOpen,
		"issueFirst": githubv4.NewInt(1),
		"issueAfter": (*githubv4.String)(nil),
	}
	variables["repositoryName"] = repositoryName
	variables["login"] = ownerLogin
	variables["assignee"] = assigneeLogin
	for {
		err := GithubClient.Query(context.Background(), &repositoryInfo, variables)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		allIssuesAssigned = append(allIssuesAssigned, repositoryInfo.Repository.Issues.Nodes...)
		if !repositoryInfo.Repository.Issues.PageInfo.HasNextPage {
			break
		}
		variables["issueAfter"] = githubv4.NewString(repositoryInfo.Repository.Issues.PageInfo.EndCursor)
	}
	return &allIssuesAssigned, nil
}

