// Copyright 2019 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model and datastore CRUD methods for an award.

package model

import "sort"

type Award struct {
	Id         int64 `db:"id"`
	Type       AwardType
	AwardName  string
	TeamId     int
	PersonName string
}

type AwardType int

const (
	JudgedAward AwardType = iota
	FinalistAward
	WinnerAward
)

func (database *Database) CreateAward(award *Award) error {
	return database.tables[Award{}].create(award)
}

func (database *Database) GetAwardById(id int64) (*Award, error) {
	var award *Award
	err := database.tables[Award{}].getById(id, &award)
	return award, err
}

func (database *Database) UpdateAward(award *Award) error {
	return database.tables[Award{}].update(award)
}

func (database *Database) DeleteAward(id int64) error {
	return database.tables[Award{}].delete(id)
}

func (database *Database) TruncateAwards() error {
	return database.tables[Award{}].truncate()
}

func (database *Database) GetAllAwards() ([]Award, error) {
	var awards []Award
	if err := database.tables[Award{}].getAll(&awards); err != nil {
		return nil, err
	}
	sort.Slice(awards, func(i, j int) bool {
		return awards[i].Id < awards[j].Id
	})
	return awards, nil
}

func (database *Database) GetAwardsByType(awardType AwardType) ([]Award, error) {
	awards, err := database.GetAllAwards()
	if err != nil {
		return nil, err
	}

	var matchingAwards []Award
	for _, award := range awards {
		if award.Type == awardType {
			matchingAwards = append(matchingAwards, award)
		}
	}
	return matchingAwards, nil
}
