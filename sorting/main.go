package main

import (
	"fmt"
	"sort"
)

type Influencer struct {
	Selfies  int
	BioLinks int
}

// vanity calculates an influencer's vanity score.
// Time complexity: O(1), Space complexity: O(1)
func vanity(influencer Influencer) int {
	return influencer.BioLinks*5 + influencer.Selfies
}

// vanitySort sorts influencers from least to greatest using key(vanity score).
// Time complexity: O(n log n), Space complexity: O(n)
func vanitySort(influencers []Influencer, key func(Influencer) int) []Influencer {
	sort.Slice(influencers, func(i, j int) bool {
		return key(influencers[i]) < key(influencers[j])
	})
	return influencers
}

func main() {
	influencers := []Influencer{
		{Selfies: 10, BioLinks: 2},
		{Selfies: 5, BioLinks: 1},
		{Selfies: 20, BioLinks: 0},
		{Selfies: 3, BioLinks: 3},
	}

	sorted := vanitySort(influencers, vanity)

	for _, influencer := range sorted {
		fmt.Printf(
			"Selfies: %d, Bio links: %d, Vanity: %d\n",
			influencer.Selfies,
			influencer.BioLinks,
			vanity(influencer),
		)
	}
}
