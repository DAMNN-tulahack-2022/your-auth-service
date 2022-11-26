package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	getUserById = `select id from your_user
						where login = $1`

	createUser = `insert into your_user
							(total_exp, vacancy_id, login, first_name, last_name,
							middle_name, avatar_uri, current_project_id, viewed_ids,
							completed_project_ids, skills_ids, role)
						values
						 	(:total_exp, :vacancy_id, :login, :first_name, :last_name,
							:middle_name, :avatar_uri, :current_project_id, :viewed_ids,
							:completed_project_ids, :skills_ids, 'regular')
								returning *`
)

type Uint32Array []uint32

func (ua Uint32Array) Value() (driver.Value, error) {
	return json.Marshal(ua)
}

func (ua *Uint32Array) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ua)
}

type User struct {
	Id                   uint32      `db:"id"`
	TotalExp             uint32      `db:"total_exp"`
	VacancyId            uint32      `db:"vacancy_id"`
	Login                string      `db:"login"`
	FirstName            string      `db:"first_name"`
	LastName             string      `db:"last_name"`
	MiddleName           string      `db:"middle_name"`
	AvaterURL            string      `db:"avatar_uri"`
	CurrentProjectId     uint32      `db:"current_project_id"`
	ViewedPostIds        Uint32Array `db:"viewed_ids"`
	CompletedProjectsIds Uint32Array `db:"completed_project_ids"`
	SkillsIds            Uint32Array `db:"skills_ids"`
	Role                 string      `db:"role"`
}
