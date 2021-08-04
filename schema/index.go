package schema

type Index struct {
	Table *Table
	TableName string
	Column *Column
	ColumnIndex int
}
