package aop

import "database/sql"

func HandleError(err error) {
	switch err {
	case sql.ErrNoRows:
	}
}
