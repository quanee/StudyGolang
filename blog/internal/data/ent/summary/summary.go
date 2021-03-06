// Code generated by entc, DO NOT EDIT.

package summary

const (
	// Label holds the string label denoting the summary type in the database.
	Label = "summary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// Table holds the table name of the summary in the database.
	Table = "summaries"
)

// Columns holds all SQL columns for summary fields.
var Columns = []string{
	FieldID,
	FieldTitle,
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
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
