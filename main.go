package main

import (
	"fmt"
	"sort"
)

type DiscordUser struct {
	DiscordUserID string            `json:"id"`
	VotingPhases  map[string]uint32 `json:"voting_phases"`
}

func main() {

	discordUsers := []DiscordUser{}
	discordUsers = addUser("3", uint32(20), discordUsers)
	discordUsers = addUser("1", uint32(60), discordUsers)
	discordUsers = addUser("4", uint32(10), discordUsers)
	discordUsers = addUser("3", uint32(40), discordUsers)

	sort.Slice(discordUsers[:], func(i, j int) bool {
		return discordUsers[i].VotingPhases["phase_1"] > discordUsers[j].VotingPhases["phase_1"]
	})

	max := 2
	lastDiscordUsers := []DiscordUser{}
	for i := 0; i < max; i++ {
		lastDiscordUsers = append(lastDiscordUsers, discordUsers[i])
	}
	fmt.Println(lastDiscordUsers)
	fmt.Println(discordUsers)
}

func addUser(userId string, points uint32, discordUsers []DiscordUser) []DiscordUser {
	discordUser := DiscordUser{
		DiscordUserID: userId,
		VotingPhases: map[string]uint32{
			"phase_1": points,
		},
	}

	discordUsers = append(discordUsers, discordUser)
	return discordUsers
}
