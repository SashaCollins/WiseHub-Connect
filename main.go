// credit - go-graphql hello world example
package main

import (
	"golang.org/x/net/context"
	"os"
	"time"

	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

//TODO: ask for token in gui for admin or maybe get token from user credentials
//TODO: save token in file and read from file on startup
//TODO: error msg if no token
const GithubToken = "43779c73fba2eff18728728abb10b0561d90ef81"

//var organization struct {
//	Login               githubv4.String
//	CreatedAt           time.Time
//	Name                githubv4.String
//	viewerCanAdminister bool
//	DatabaseID          githubv4.Int
//	URL                 githubv4.URI
//}
var orga struct {
	Organization struct {
		DatabaseID	githubv4.Int
		URL			githubv4.URI
		Name 	githubv4.String
		//Members struct {
		//	totalCount 	githubv4.Int
		//}
	} `graphql:"organization(login:$organizationName)"`
	Viewer struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		ID         githubv4.ID
		DatabaseID githubv4.Int
	}
}
var orgaVariables = map[string]interface{}{
	"organizationName": githubv4.String("WiseHub-Connector"),
	//"organizationFirst": 		githubv4.NewInt(5),
	//"repositoryName":  githubv4.String("firstRepository"),
	//"issueNumber":     githubv4.Int(1),
	//"commentsFirst":   githubv4.NewInt(1),
	//"commentsAfter":   githubv4.NewString("Y3Vyc29yOjE5NTE4NDI1Ng=="),
}
var user struct {
	User struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		Organizations struct {
			Nodes []struct {
				Login githubv4.String
				DatabaseID githubv4.Int
				//viewerCanAdminister githubv4.Boolean
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage githubv4.Boolean
			}
		} `graphql:"organizations(first:$organizationFirst)"`
	} `graphql:"user(login:$repositoryOwner)"`
}
var userVariables = map[string]interface{}{
	"repositoryOwner": githubv4.String("SashaCollins"),
	"organizationFirst": githubv4.NewInt(1),
	//"organizationAfter": githubv4.NewString("Y3Vyc29yOnYyOpHOAiHjZw=="),
	//"organizationName": githubv4.String("WiseHub-Connector"),
	//"repositoryName":  githubv4.String("firstRepository"),
	//"issueNumber":     githubv4.Int(1),
	//"commentsFirst":   githubv4.NewInt(1),
	//"commentsAfter":   githubv4.NewString("Y3Vyc29yOjE5NTE4NDI1Ng=="),
}
var testViewer struct {
	Viewer struct {
		Login     string
		CreatedAt time.Time
	}
}

type githubV4Actor struct {
	Login     githubv4.String
	AvatarURL githubv4.URI `graphql:"avatarUrl(size:72)"`
	URL       githubv4.URI
}

var repo struct {
	Repository struct {
		DatabaseID githubv4.Int
		URL        githubv4.URI

		Issue struct {
			Title			githubv4.String
			Author         githubV4Actor
			PublishedAt    githubv4.DateTime
			LastEditedAt   *githubv4.DateTime
			Editor         *githubV4Actor
			Body           githubv4.String
			ReactionGroups []struct {
				Content githubv4.ReactionContent
				Users   struct {
					Nodes []struct {
						Login githubv4.String
					}

					TotalCount githubv4.Int
				} `graphql:"users(first:10)"`
				ViewerHasReacted githubv4.Boolean
			}
			ViewerCanUpdate githubv4.Boolean

			Comments struct {
				Nodes []struct {
					Body   githubv4.String
					Author struct {
						Login githubv4.String
					}
					Editor struct {
						Login githubv4.String
					}
				}
				PageInfo struct {
					StartCursor githubv4.String
					HasPreviousPage githubv4.Boolean
					EndCursor   githubv4.String
					HasNextPage githubv4.Boolean
				}
			} `graphql:"comments(first:$commentsFirst,after:$commentsAfter)"`
		} `graphql:"issue(number:$issueNumber)"`
	} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
	Viewer struct {
		Login      githubv4.String
		CreatedAt  githubv4.DateTime
		ID         githubv4.ID
		DatabaseID githubv4.Int
	}
	RateLimit struct {
		Cost      githubv4.Int
		Limit     githubv4.Int
		Remaining githubv4.Int
		ResetAt   githubv4.DateTime
	}
}
var variables = map[string]interface{}{
	"repositoryOwner": githubv4.String("SashaCollins"),
	"repositoryName":  githubv4.String("firstRepository"),
	"issueNumber":     githubv4.Int(1),
	"commentsFirst":   githubv4.NewInt(1),
	"commentsAfter":   githubv4.NewString("Y3Vyc29yOnYyOpHOI8NEVw=="),
}

func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func run(client *githubv4.Client, query interface{}, localVar map[string]interface{}) error {
	err := client.Query(context.Background(), &repo, localVar)
	if err != nil {
		return err
	}
	printJSON(repo)
	//goon.Dump(out)
	//fmt.Println(github.Stringify(out))
	return nil
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	// Use client...

	fuckingErr:= run(client, repo, variables)
	if fuckingErr != nil {
		fmt.Println(fuckingErr)
	}

	err := client.Query(context.Background(), &testViewer, nil)
	if err != nil {
		fmt.Printf("\tQuery query failed with: %s\n",err)
	}
	fmt.Println("Login:", testViewer.Viewer.Login)
	fmt.Println("CreatedAt:", testViewer.Viewer.CreatedAt)

	err = client.Query(context.Background(), &orga, orgaVariables)
	if err != nil {
		fmt.Printf("\tQuery Organization failed with: %s\n", err)

	}
	fmt.Printf("orga name: %s\n", orga.Organization.Name)
	fmt.Printf("orga id: %d\n", orga.Organization.DatabaseID)


	err = client.Query(context.Background(), &user, userVariables)
	if err != nil {
		fmt.Printf("\tQuery User failed with: %s\n", err)

	}
	fmt.Printf("user Login: %s\n", user.User.Login)
	fmt.Printf("user createdAt: %s\n", user.User.CreatedAt)
	fmt.Println("user orga Node Login: ", user.User.Organizations.Nodes)


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
//											}
//										}
//									}
//								}
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//}