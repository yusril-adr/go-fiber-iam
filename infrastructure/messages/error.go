package messages

import "fmt"

func ErrSortByColInvalid(column string, allowedColumns []string) string {
	return fmt.Sprintf("sort_by column %s is invalid, allowed columns are %v", column, allowedColumns)
}
