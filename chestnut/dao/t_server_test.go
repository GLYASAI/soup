// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package dao

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"testing"
)

func TestTServerImpl_GetServerIDByIP(t *testing.T) {
	db, err := sql.Open("oci8", "huangrh/12345678@127.0.0.1:1521/helowin")
	if err != nil {
		t.Fatalf("Open error is not nil: %v", err)
	}
	if db == nil {
		t.Fatalf("db is nil")
	}
	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			t.Fatalf("Close error is not nil: %v", err)
		}
	}()

	// TODO: create table

	_, err = db.Exec("insert into T_SERVER(SERVER_ID, IP_ADDR) values (:1, :2)", "888", "192.168.11.11")
	if err != nil {
		t.Fatalf("ExecContext error is not nil: %v", err)
	}

	tserver := TServerImpl{
		db: db,
	}
	serverID, err := tserver.GetServerIDByIP("192.168.11.11")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(serverID)

	_, err = db.Exec("delete from T_SERVER WHERE SERVER_ID = :1", "888")
	if err != nil {
		t.Fatalf("Exec error is not nil: %v", err)
	}
}
