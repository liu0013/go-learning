// Code generated by entc, DO NOT EDIT.

package feedback

import (
	"time"
)

const (
	// Label holds the string label denoting the feedback type in the database.
	Label = "feedback"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLike holds the string denoting the like edge name in mutations.
	EdgeLike = "like"
	// Table holds the table name of the feedback in the database.
	Table = "feedbacks"
	// LikeTable is the table that holds the like relation/edge.
	LikeTable = "likes"
	// LikeInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikeInverseTable = "likes"
	// LikeColumn is the table column denoting the like relation/edge.
	LikeColumn = "feedback_like"
)

// Columns holds all SQL columns for feedback fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldContent,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)