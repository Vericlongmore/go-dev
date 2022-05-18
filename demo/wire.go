//+build wireinject

package main

import "github.com/google/wire"

import "time"
import "errors"
import "fmt"

type Monster struct {
	Name string
}

type Player struct {
	Name string
}

type Mission struct {
	Player  Player
	Monster Monster
}

//func NewPlayer(name string) Player{
//
//	return Player{Name: name}
//}

func NewPlayer(name string) (Player, func(), error) {
	cleanup := func() {
		fmt.Println("cleanup!")
	}
	if time.Now().Unix()%2 == 0 {
		return Player{}, cleanup, errors.New("player dead")
	}
	return Player{Name: name}, cleanup, nil
}

func NewMonster() Monster {

	return Monster{Name: "kitty"}
}

func NewMission() Mission {
	p := Player{Name: "dj"}
	m := Monster{Name: "kitty"}

	return Mission{p, m}
}

func InitMission(name string) (Mission, func(), error) {
	wire.Build(NewPlayer, NewMonster, NewMission)
	return Mission{}, nil, nil
}

func InitPlayer() Player {
	wire.Build(NewMission, wire.FieldsOf(new(Mission), "Player"))
	return Player{}
}
func InitMonster() Monster {
	wire.Build(NewMission, wire.FieldsOf(new(Mission), "Monster"))
	return Monster{}
}
