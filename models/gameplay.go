package models

import (
	"fmt"
	"time"
)


type Game struct {
	Id          string    `json:"id"  gorm:"primary_key"`
	TagName     string    `json:"tag_name" gorm:"size:255;"`
	TargetCount int64     `json:"target_count" gorm:"size:255;"`
	WinnersCut  float64   `json:"winners_cut" gorm:"size:255;"`
	Cost        int64     `json:"cost"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}


type GameEntry struct {
	Phone    string    `json:"phone"`
	GuessOne string    `json:"guess_one"`
	GuessTwo string    `json:"guess_two"`
	GameId   string    `json:"game_id"`
	Time     time.Time `json:"time"`
}

func (ge *GameEntry) String() string {
	t:= time.Now().String()
	return fmt.Sprintf("%s %s %s %v -- %v\n", ge.Phone, ge.GuessOne, ge.GuessTwo, ge.GameId, t)
}