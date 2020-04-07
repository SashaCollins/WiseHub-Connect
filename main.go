// credit - go-graphql hello world example
package main

import (
	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"strings"

	//"go/types"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"os"
)

//TODO: ask for token in gui for admin or maybe get token from user credentials
//TODO: save token in file and read from file on startup
//TODO: error msg if no token
const GithubToken = "43779c73fba2eff18728728abb10b0561d90ef81"

type commit struct {
	Author githubV4Actor
	Committer githubV4Actor
}

type comment struct {
	Body   githubv4.String
	Author struct {
		Login githubv4.String
	}
	Editor struct {
		Login githubv4.String
	}
	ViewerCanReact bool
}

type githubV4Actor struct {
	Login     githubv4.String
	URL       githubv4.URI
}

var allOrganizations []organization
type organization struct {
		Login githubv4.String
		ViewerCanAdminister githubv4.Boolean
		DatabaseID	githubv4.Int
		URL			githubv4.URI
		Name 	githubv4.String
		MembersWithRole struct {
			TotalCount githubv4.Int
			Nodes []githubV4Actor
		}`graphql:"membersWithRole(first:$orgaMembersFirst)"`
		Teams struct {
			TotalCount githubv4.Int
			Nodes []team
			PageInfo pageInfo
		}`graphql:"teams(first:$teamFirst,after:$teamAfter)"`
}

//var currentPageInfo pageInfo
type pageInfo struct {
	StartCursor githubv4.String
	HasPreviousPage githubv4.Boolean
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}

//var currentRateLimit rateLimit
type rateLimit struct {
	Cost      githubv4.Int
	Limit     githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}

type repository struct {
	Repository struct {
		ID githubv4.String
		NameWithOwner githubv4.String
		CreatedAt githubv4.DateTime
		Description githubv4.String
		URL        githubv4.URI
		IsArchived githubv4.Boolean
		IsPrivate githubv4.Boolean
		DatabaseID githubv4.Int

		Issue struct {
			Title			githubv4.String
			Author         githubV4Actor
			PublishedAt    githubv4.DateTime
			LastEditedAt   *githubv4.DateTime
			Editor         *githubV4Actor
			Body           githubv4.String
			ViewerCanUpdate githubv4.Boolean
			Comments struct {
				TotalCount githubv4.Int
				Nodes []comment
				PageInfo pageInfo
			} `graphql:"comments(first:$commentsFirst,after:$commentsAfter)"`
		} `graphql:"issue(number:$issueNumber)"`
		//Commits struct {
		//	TotalCount githubv4.Int
		//	Nodes []commit
		//}
		//DefaultBranchRef ref
	} `graphql:"repository(owner:$viewer,name:$repositoryName)"`
}

//var currentRepository repository
var allRepos []teamRepository
type teamRepository struct {
	ID githubv4.String
	NameWithOwner githubv4.String
	CreatedAt githubv4.DateTime
}

var allTeams []team
type team struct {
	Name githubv4.String
	CombinedSlug githubv4.String
	CreatedAt githubv4.DateTime
	Description githubv4.String
	Privacy githubv4.TeamPrivacy
	ViewerCanAdminister githubv4.Boolean
	Members struct {
		TotalCount githubv4.Int
		Nodes []githubV4Actor
	}`graphql:"members(first:$teamMembersFirst)"`
	Repositories struct {
		TotalCount githubv4.Int
		Nodes []teamRepository
		PageInfo pageInfo
	}`graphql:"repositories(first:$repositoryFirst,after:$repositoryAfter)"`
	RepositoriesResourcePath githubv4.URI
	RepositoriesUrl githubv4.URI
}

var currentUser user
type user struct {
	User struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		Organizations struct {
			Nodes []organization
			TotalCount githubv4.Int
			PageInfo pageInfo
		} `graphql:"organizations(first:$organizationFirst,after:$organizationAfter)"`
	} `graphql:"user(login:$viewer)"`
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
var currentViewer viewer
type viewer struct {
	Viewer struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		ID         githubv4.ID
		DatabaseID githubv4.Int
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
//					owner{
//						login
//						id
//						__typename
//						url
//					}
//					assignableUsers{
//						totalCount
//					}
//					licenseInfo{
//						key
//					}
//					defaultBranchRef{
//						target{
//							... on Commit{
//								history(first:10){
//									totalCount
//									edges{
//										node{
//											... on Commit{
//												committedDate

var currentRepo repo
type repo struct {
	Repository struct {
		ID githubv4.String
		NameWithOwner githubv4.String
		CreatedAt githubv4.DateTime
		Description githubv4.String
		IsArchived githubv4.Boolean
		IsPrivate githubv4.Boolean
		DatabaseID githubv4.Int
		URL        githubv4.URI

		Issue struct {
			Title			githubv4.String
			Author         githubV4Actor
			PublishedAt    githubv4.DateTime
			LastEditedAt   *githubv4.DateTime
			Editor         *githubV4Actor
			Body           githubv4.String
			ViewerCanUpdate githubv4.Boolean
			Comments struct {
				TotalCount githubv4.Int
				Nodes []comment
				PageInfo pageInfo
			} `graphql:"comments(first:$commentsFirst,after:$commentsAfter)"`
		} `graphql:"issue(number:$issueNumber)"`
		//Commits struct {
		//	TotalCount githubv4.Int
		//	Nodes []commit
		//}
		//DefaultBranchRef ref
	} `graphql:"repository(owner:$viewer,name:$repositoryName)"`
	RateLimit rateLimit
}

var repoVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"repositoryName":  githubv4.String("firstRepository"),
	"issueNumber":     githubv4.Int(1),
	"commentsFirst":   githubv4.NewInt(1),
	"commentsAfter":   (*githubv4.String)(nil),
}
var userVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"organizationFirst": githubv4.NewInt(1),
	"organizationAfter": (*githubv4.String)(nil),
	"orgaMembersFirst": githubv4.NewInt(1),
	"teamFirst": githubv4.NewInt(1),
	"teamAfter": (*githubv4.String)(nil),
	"teamMembersFirst": githubv4.NewInt(1),
	"repositoryFirst": githubv4.NewInt(1),
	"repositoryAfter": (*githubv4.String)(nil),
	//"issueNumber":     githubv4.Int(1),
	//"organizationName": githubv4.String("WiseHub-Connector"),
}
var countVariables = map[string]interface{}{
	"viewer": (*githubv4.String)(nil), //githubv4.String("SashaCollins"),
	"organizationFirst": githubv4.NewInt(1),
	"organizationAfter": (*githubv4.String)(nil),
	"teamFirst": githubv4.NewInt(1),
	"teamAfter": (*githubv4.String)(nil),
}

func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func run(client *githubv4.Client, query interface{}, localVar *map[string]interface{}) error {
	fmt.Println("in run")
	switch currentQuery := query.(type) {
	case *repo:
		fmt.Println("\tin Repo")
		var allComments []comment
		for {
			err := client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery repository failed with:")
				return err
			}
			allComments = append(allComments, currentQuery.Repository.Issue.Comments.Nodes...)

			if !currentQuery.Repository.Issue.Comments.PageInfo.HasNextPage {
				break
			}
			repoVariables["commentsAfter"] = githubv4.NewString(currentQuery.Repository.Issue.Comments.PageInfo.EndCursor)
		}
		//printJSON(allComments)
		return nil

	case *user:
		fmt.Println("\tin User")
		var countQuery count
		for {
			err := client.Query(context.Background(), &countQuery, countVariables)
			if err != nil {
				fmt.Println("\tQuery user in line 1 failed with")
				return err
			}
			userVariables["orgaMembersFirst"] = githubv4.NewInt(countQuery.User.Organizations.Nodes[0].MembersWithRole.TotalCount)
			err = client.Query(context.Background(), &currentQuery, *localVar)
			if err != nil {
				fmt.Println("\tQuery user in line 2 failed with")
				return err
			}
			allOrganizations = append(allOrganizations, currentQuery.User.Organizations.Nodes...)
			latestOrga := len(allOrganizations)-1
			for {
				err := client.Query(context.Background(), &countQuery, countVariables)
				if err != nil {
					fmt.Println("\tQuery user in line 3 failed with")
					return err
				}
				userVariables["teamMembersFirst"] = githubv4.NewInt(countQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Members.TotalCount)

				err = client.Query(context.Background(), &currentQuery, *localVar)
				if err != nil {
					fmt.Println("\tQuery user in line 4 failed with")
					return err
				}
				//safe current team in slice at correct position
				if !(strings.EqualFold(string(allOrganizations[latestOrga].Teams.Nodes[0].CombinedSlug), string(currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].CombinedSlug))) {
					allOrganizations[latestOrga].Teams.Nodes = append(allOrganizations[latestOrga].Teams.Nodes, currentQuery.User.Organizations.Nodes[0].Teams.Nodes...)
				} else {
					allOrganizations[latestOrga].Teams.Nodes[0] = currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0]
				}
				latestTeam := len(allOrganizations[len(allOrganizations)-1].Teams.Nodes)-1
				for {
					err := client.Query(context.Background(), &currentQuery, *localVar)
					if err != nil {
						fmt.Println("\tQuery user in line 5 failed with")
						return err
					}
					//safe current Repo in slice at correct position
					if !(strings.EqualFold(string(allOrganizations[latestOrga].Teams.Nodes[latestTeam].Repositories.Nodes[0].NameWithOwner), string(currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Repositories.Nodes[0].NameWithOwner))) {
						allOrganizations[latestOrga].Teams.Nodes[latestTeam].Repositories.Nodes = append(allOrganizations[latestOrga].Teams.Nodes[latestTeam].Repositories.Nodes, currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Repositories.Nodes...)
					}else {
						allOrganizations[latestOrga].Teams.Nodes[latestTeam].Repositories.Nodes[0] = currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Repositories.Nodes[0]
					}
					if !currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Repositories.PageInfo.HasNextPage {
						userVariables["repositoryAfter"] = (*githubv4.String)(nil)
						break
					}
					userVariables["repositoryAfter"] = githubv4.NewString(currentQuery.User.Organizations.Nodes[0].Teams.Nodes[0].Repositories.PageInfo.EndCursor)
				}
				if !currentQuery.User.Organizations.Nodes[0].Teams.PageInfo.HasNextPage {
					userVariables["teamAfter"] = (*githubv4.String)(nil)
					countVariables["teamAfter"] = (*githubv4.String)(nil)
					break
				}
				userVariables["teamAfter"] = githubv4.NewString(currentQuery.User.Organizations.Nodes[0].Teams.PageInfo.EndCursor)
				countVariables["teamAfter"] = githubv4.NewString(currentQuery.User.Organizations.Nodes[0].Teams.PageInfo.EndCursor)
			}
			if !currentQuery.User.Organizations.PageInfo.HasNextPage {
				break
			}
			userVariables["organizationAfter"] = githubv4.NewString(currentQuery.User.Organizations.PageInfo.EndCursor)
			countVariables["organizationAfter"] = githubv4.NewString(currentQuery.User.Organizations.PageInfo.EndCursor)
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
		fmt.Println(currentQuery)
		err := client.Query(context.Background(), &currentQuery, *localVar)
		if err != nil {
			fmt.Println("\tQuery default failed with:")
			return err
		}
		printJSON(currentQuery)
		return nil
	}
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	// Use client...
	fmt.Println("before run")
	fuckingErr := run(client, &currentViewer, nil)
	if fuckingErr != nil {
		fmt.Println(fuckingErr)
	}
	repoVariables["viewer"] = currentViewer.Viewer.Login
	userVariables["viewer"] = currentViewer.Viewer.Login
	countVariables["viewer"] = currentViewer.Viewer.Login
	//printJSON(currentViewer)

	fuckingErr = run(client, &currentRepo, &repoVariables)
	if fuckingErr != nil {
		fmt.Println(fuckingErr)
	}
	//printJSON(currentRepo)

	fuckingErr = run(client, &currentUser, &userVariables)
	if fuckingErr != nil {
		fmt.Println(fuckingErr)
	}
	fmt.Println("#############################################")
	//printJSON(currentUser)
	//printJSON(allOrganizations)
	fmt.Println("after run")
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