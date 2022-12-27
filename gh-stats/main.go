package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"strings"
)

func main() {
	userPtr := flag.String("user", "", "Required: Github username (case sensitive!)")
	orgPtr := flag.String("org", "", "Required: organisation name")
	repoPtr := flag.String("repo", "", "Required: Repository")
	tokenPtr := flag.String("token", "", "Required: Github auth token")

	flag.Parse()

	if *userPtr == "" || *orgPtr == "" || *repoPtr == "" || *tokenPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *tokenPtr})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// TODO: GH, second argument – activity is not immediately accessible, so perhaps I'd need to refetch it
	activity, _, err := client.Repositories.ListContributorsStats(ctx, *orgPtr, *repoPtr)

	if err != nil {
		log.Fatalln(err)
	}

	var act *github.ContributorStats
	for _, a := range activity {
		if strings.Compare(*a.Author.Login, *userPtr) == 0 {
			act = a
		}
	}

	additions := 0
	deletions := 0
	commits := 0

	for _, week := range act.Weeks {
		additions += *week.Additions
		deletions += *week.Deletions
		commits += *week.Commits
	}

	fmt.Printf("User: " + *userPtr + ", Repo: " + *repoPtr + "\n")
	fmt.Printf("additions: %v\n", additions)
	fmt.Printf("deletions: %v\n", deletions)
	fmt.Printf("commits: %v\n", commits)

}
