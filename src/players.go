package main

type Players struct {
	AI string
	Human string
}

func NewPlayers(ai string, human string) *Players {
	players := &Players{
		AI: ai,
		Human: human,
	}
	return players
}

func NewPlayersWithHuman(human string) *Players {
	var players *Players = nil
	if human == "X" {
		players = &Players{
			AI: "O",
			Human: "X",
		}
	} else if human == "O"{
		players = &Players{
			AI: "O",
			Human: "X",
		}
	}
	return players
}