package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	semver.Sort(releases)
	// set array pointers
	verLoca := 0
	relLoca := len(releases) - 1
	// Find the large release Version and put as first element of versionSlice
	for ; relLoca >= 0; relLoca-- {
		if releases[relLoca].PreRelease != "" {
			continue
		} else if releases[relLoca].Compare(*minVersion) < 0 {
			break
		} else {
			versionSlice = append(versionSlice, releases[relLoca])
			break
		}
	}
	// Compare remaining release in releases to versionSlice[curr]
	for ; relLoca >= 0; relLoca-- {
		if releases[relLoca].PreRelease != "" {
			continue
		} else if releases[relLoca].Compare(*minVersion) < 0 {
			break
		} else {
			// Compare Major, then Minor, no need for patch as first one is largest patch
			if releases[relLoca].Major < versionSlice[verLoca].Major {
				verLoca++
				versionSlice = append(versionSlice, releases[relLoca])
			} else if releases[relLoca].Minor < versionSlice[verLoca].Minor {
				verLoca++
				versionSlice = append(versionSlice, releases[relLoca])
			}
		}
	}
	return versionSlice
}

// Helper function to start a panic and try to recover the panic
func showError(err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n", r)
		}
	}()
	panic(err)
}

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Get input arg which is file path name
	arg := os.Args
	if len(arg) == 2 {
		bytes, err := ioutil.ReadFile(arg[1])
		dat := string(bytes)
		if err != nil {
			showError(err)
		}
		// Split the data
		repoMins := strings.Split(dat, "\n")
		repos := make([]string, len(repoMins))
		pages := make([]string, len(repoMins))
		minVers := make([]string, len(repoMins))
		for i, repoMin := range repoMins {
			firstSplits := strings.Split(repoMin, ",")
			//Split for parsing "0/r" error in last Repo case in file
			minVers[i] = strings.Split(firstSplits[1], "\r")[0]
			repoNotSplited := firstSplits[0]
			secondSplits := strings.Split(repoNotSplited, "/")
			pages[i] = secondSplits[0]
			repos[i] = secondSplits[1]
		}
		// Github
		client := github.NewClient(nil)
		ctx := context.Background()
		opt := &github.ListOptions{PerPage: 50}
		for softNum := 0; softNum < len(repoMins); softNum++ {
			//***Error handling test case*********//
			//releases, _, err := client.Repositories.ListReleases(ctx, "lkubernetes", "kulbernetes", opt)
			//***Original list releases case******//
			//releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
			//***General version list release*****//
			releases, _, err := client.Repositories.ListReleases(ctx, pages[softNum], repos[softNum], opt)

			if err != nil {
				//panic(err) // is this really a good way?
				//No, it is a fatel panic, not good if we still want to
				//retrieve the relaese version of other open source software
				//Try to recover and continue to run instead
				showError(err)
			}
			//Original minVersion is not used when reading from file****************
			//minVersion := semver.New("1.8.0")
			minVersion := semver.New(minVers[softNum])
			allReleases := make([]*semver.Version, len(releases))
			for i, release := range releases {
				versionString := *release.TagName
				if versionString[0] == 'v' {
					versionString = versionString[1:]
				}
				allReleases[i] = semver.New(versionString)
			}
			versionSlice := LatestVersions(allReleases, minVersion)
			//***Original print out*****//
			//fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
			//***General print out******//
			fmt.Printf("latest versions of %s/%s: %s\n", pages[softNum], repos[softNum], versionSlice)
		}
	} else {
		fmt.Printf("Too many/No enough arguments! Required 2 args, you gave %d", len(arg))
	}
}
