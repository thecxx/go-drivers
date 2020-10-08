package mysql

import (
	"context"
	"testing"
)

func TestTransaction_Exec(t *testing.T) {
	d, err := newDatabase()
	if err != nil {
		t.Errorf("New database failed, err=%s\n", err.Error())
		return
	}

	_, err = d.Exec(context.Background(), "DELETE FROM `test` WHERE name = ?", "unittest")
	if err != nil {
		t.Errorf("Delete unittest failed, err=%s\n", err.Error())
		return
	}

	tran, err := d.BeginTransaction(context.Background())
	if err != nil {
		t.Errorf("Begin transaction failed, err=%s\n", err.Error())
		return
	}
	var txerr error = nil
	defer func(txerr error) {
		if txerr != nil {
			tran.Rollback()
		}
	}(txerr)
	_, txerr = tran.Exec(context.Background(), "INSERT INTO `test` (`name`, `gid`) VALUES (?,?)", "unittest", 1)
	if txerr != nil {
		t.Errorf("Execute transaction failed, err=%s\n", err.Error())
		return
	}
	if tran.Commit() != nil {
		t.Errorf("Commit transaction failed, err=%s\n", err.Error())
		return
	}
	result, err := d.Query(context.Background(), "SELECT * FROM `test` WHERE name = ? LIMIT 1", "unittest")
	if err != nil {
		t.Errorf("Query unittest failed, err=%s\n", err.Error())
		return
	}
	row, err := result.Row()
	if err != nil {
		t.Errorf("Get row failed, err=%s\n", err.Error())
		return
	}
	if row == nil {
		t.Errorf("Invalid row\n")
		return
	}
	if row["name"] != "unittest" {
		t.Errorf("Transaction error\n")
		return
	}
}
