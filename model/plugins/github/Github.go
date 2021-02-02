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

//TODO: ask for token in gui for admin or maybe get token from user credentials
//TODO: save token in file and read from file on startup
//TODO: error msg if no token


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
	//ViewerCanUpdate githubv4.Boolean
}

type Organization struct {
	Login githubv4.String
	//URL			githubv4.URI
	//ViewerCanAdminister githubv4.Boolean
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



type RTeams struct{
	Name githubv4.String `json:"orga_name"`
	Members []RMembers `json:"member"`
	Repositories []RRepositories `json:"repository"`
}
type RMembers struct {
	Login     githubv4.String `json:"orga_name"`
}
type RRepositories struct {
	Name  githubv4.String `json:"repo_name"`
	URL   githubv4.URI `json:"repo_url"`
	Issues []RIssues `json:"issue"`
}
type RIssues struct {
	Number			githubv4.Int `json:"issue_number"`
	Title			githubv4.String `json:"issue_title"`
}
type Response struct {
	Organization struct{
		Login githubv4.String `json:"orga_name"`
		Teams []RTeams `json:"team"`
	} `json:"organization"`
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

				allIssuesAssigned, err := g.getRepositoryInfo((githubv4.String)(repo.Name), (githubv4.String)(repo.Owner.Login))
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
func (g *Github) SubmitCredentials(username, token string) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	GithubClient = githubv4.NewClient(httpClient)
}
func (g *Github) FetchPluginName() string {
	return getPluginName()
}



func (g *Github) printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func (g *Github) getViewer() (*Viewer, error) {
	err := GithubClient.Query(context.Background(), &CurrentViewer, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &CurrentViewer, nil
}

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

func (g *Github) getRepositoryInfo(repositoryName githubv4.String, ownerLogin githubv4.String) (*[]Issue, error) {
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
	variables["assignee"] = ownerLogin
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

