package main

type Players struct {
	AI string
	Human string
}

func NewPlayers(AI string, Human string) *Players {
	players := &Players{
		AI: AI,
		Human: Human,
	}
	return players
}