// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "course_name", Type: field.TypeString},
		{Name: "course_location", Type: field.TypeString},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
	}
	// HolesColumns holds the columns for the "holes" table.
	HolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hole_number", Type: field.TypeInt},
		{Name: "hole_name", Type: field.TypeString},
	}
	// HolesTable holds the schema information for the "holes" table.
	HolesTable = &schema.Table{
		Name:       "holes",
		Columns:    HolesColumns,
		PrimaryKey: []*schema.Column{HolesColumns[0]},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
	}
	// RoundsColumns holds the columns for the "rounds" table.
	RoundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_of_round", Type: field.TypeTime},
	}
	// RoundsTable holds the schema information for the "rounds" table.
	RoundsTable = &schema.Table{
		Name:       "rounds",
		Columns:    RoundsColumns,
		PrimaryKey: []*schema.Column{RoundsColumns[0]},
	}
	// TeesColumns holds the columns for the "tees" table.
	TeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "color", Type: field.TypeString},
		{Name: "hole_name", Type: field.TypeInt},
	}
	// TeesTable holds the schema information for the "tees" table.
	TeesTable = &schema.Table{
		Name:       "tees",
		Columns:    TeesColumns,
		PrimaryKey: []*schema.Column{TeesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoursesTable,
		HolesTable,
		PlayersTable,
		RoundsTable,
		TeesTable,
		UsersTable,
	}
)

func init() {
}