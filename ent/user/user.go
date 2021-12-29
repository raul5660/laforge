// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// EdgeUserToTag holds the string denoting the usertotag edge name in mutations.
	EdgeUserToTag = "UserToTag"
	// EdgeUserToEnvironment holds the string denoting the usertoenvironment edge name in mutations.
	EdgeUserToEnvironment = "UserToEnvironment"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserToTagTable is the table that holds the UserToTag relation/edge.
	UserToTagTable = "tags"
	// UserToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	UserToTagInverseTable = "tags"
	// UserToTagColumn is the table column denoting the UserToTag relation/edge.
	UserToTagColumn = "user_user_to_tag"
	// UserToEnvironmentTable is the table that holds the UserToEnvironment relation/edge. The primary key declared below.
	UserToEnvironmentTable = "environment_EnvironmentToUser"
	// UserToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	UserToEnvironmentInverseTable = "environments"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUUID,
	FieldEmail,
	FieldHclID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"command_command_to_user",
	"finding_finding_to_user",
	"host_host_to_user",
	"script_script_to_user",
}

var (
	// UserToEnvironmentPrimaryKey and UserToEnvironmentColumn2 are the table columns denoting the
	// primary key for the UserToEnvironment relation (M2M).
	UserToEnvironmentPrimaryKey = []string{"environment_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)