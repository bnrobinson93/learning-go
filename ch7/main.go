package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type (
	TeamName string
)

// Team is a part of Ex 1
type Team struct {
	Name    TeamName
	Players []string
}

func (t *Team) AddPlayer(name string) {
	t.Players = append(t.Players, name)
}

// League is a part of Ex 1
type League struct {
	Teams []Team
	Wins  map[TeamName]int
}

func (l *League) PrintTeams() {
	fmt.Println("All teams:")
	for _, t := range l.Teams {
		fmt.Println(t.Name, "Roster:")
		for _, p := range t.Players {
			fmt.Println(p)
		}
	}
}

// MatchResult is a part of Ex 2 -- assume only one league
func (l *League) MatchResult(name1 TeamName, score1 int, name2 TeamName, score2 int) {
	if score1 == score2 {
		fmt.Println("It is a draw")
		return
	}
	if _, ok := l.Wins[name1]; !ok {
		return
	}
	if _, ok := l.Wins[name2]; !ok {
		return
	}

	if score1 > score2 {
		fmt.Println(name1, "won")
		l.Wins[name1]++
	} else {
		fmt.Println(name2, "won")
		l.Wins[name2]++
	}
}

// Ranking is a part of Ex 2
func (l *League) Ranking() []TeamName {
	names := make([]TeamName, 0, len(l.Wins))
	for k := range l.Wins {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

func (l *League) WinCount(name string) int {
	if wins, ok := l.Wins[TeamName(name)]; ok {
		return wins
	}
	return 0
}

// Ranker is used in Ex 3
// I'm purposefully making it hard by making this a string, not a TeamName (=string)
// so that I can try my hand at an adapter
type Ranker interface {
	Ranking() []string
}

/********* Struct embed pattern **********/
type RankerStruct struct {
	// Note that this would work with VALUE/copy
	// The league only has a map that gets updated so it kinda doesn't matter here
	// but If we want to mutate the league, we should really use a pointer
	// Also, using a pointer means method access sugar sticks, where it auto translates if addressable
	*League // now I have all the fields and methods from league in RankerStruct
}

func (r RankerStruct) Ranking() []string {
	source := r.League.Ranking()
	output := make([]string, 0, len(source))
	for _, v := range source {
		output = append(output, string(v))
	}
	return output
}

/********* End struct embed pattern **********/

/********* Adapater pattern **********/
// Since it's just one function and no state, this is better than struct embed
type RankerAdapter func() []string

// Make RankerAdapter match the Ranker interface requirement
func (f RankerAdapter) Ranking() []string { return f() }

// The actual translation function
func (l *League) RankingStrings() []string {
	output := make([]string, 0, len(l.Wins))
	for _, r := range l.Ranking() {
		output = append(output, string(r))
	}
	return output
}

func RankPrinter(r Ranker, w io.Writer) error {
	counter, hasCounts := r.(interface{ WinCount(string) int })
	for _, rank := range r.Ranking() {
		output := rank
		if hasCounts {
			output = fmt.Sprintf("%s (%d)", rank, counter.WinCount(rank))
		}
		_, err := io.WriteString(w, output+"\n")
		if err != nil {
			return fmt.Errorf("unable to write %v: %w", rank, err)
		}
	}
	return nil
}

/********* End adapater pattern **********/

func main() {
	team1 := Team{
		Name:    "Cowboys",
		Players: []string{"Brad", "Christy"},
	}
	team2 := Team{
		Name:    "Broncos",
		Players: []string{"Madhu", "Harrison"},
	}
	league := League{
		Teams: []Team{team1, team2},
		Wins:  map[TeamName]int{"Cowboys": 0, "Broncos": 0},
	}

	league.MatchResult("Cowboys", 1, "Broncos", 2)
	league.MatchResult("Cowboys", 3, "Broncos", 2)
	league.MatchResult("Cowboys", 4, "Broncos", 3)
	league.MatchResult("Cowboys", 2, "BadTeam", 5)
	league.MatchResult("BadTeam", 2, "Broncos", 5)

	// Ex 3
	// struct version
	if err := RankPrinter(RankerStruct{&league}, os.Stdout); err != nil {
		fmt.Println(err)
	}

	// type conversion, the league method to a RankerAdapter
	// allowed since RankingStrings == func() []string == RankerAdapter
	// same as what http.HandlerFunc encourages
	adapter := RankerAdapter(league.RankingStrings)
	// adapter carries .Ranking, which meets the interface req
	if err := RankPrinter(adapter, os.Stdout); err != nil {
		fmt.Println(err)
	}

	// Added this to play with pointers in functions
	// league.PrintTeams()
	// league.Teams[0].AddPlayer("Dr. Franklin")
	// league.PrintTeams() // missing Dr. Franklin
}
