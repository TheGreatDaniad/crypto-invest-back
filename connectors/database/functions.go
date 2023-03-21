package database

import "strings"

func IsRecordNotFound(err error) bool {
	return err.Error() == "record not found"
}

func IsForeignKeyNotFound(err error) bool {
	return strings.Contains(err.Error(), "(SQLSTATE 23503)")

}
