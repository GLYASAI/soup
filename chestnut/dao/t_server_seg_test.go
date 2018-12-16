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
	"github.com/GLYASAI/soup/chestnut/dao/model"
	"testing"
)

func TestTServerSeg_AddOrUpdate(t *testing.T) {
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

	_, err = db.Exec(`
	create table t_server_seg
	(
		server_seg_id varchar(255),
		server_id varchar(255),
		seg_pref varchar(255),
		ver varchar(255)
	)`)
	if err != nil {
		t.Fatalf("error create table %s: %v", "t_server_seg", err)
	}

	tss := &TServerSegImpl{
		db: db,
	}
	tServerSeg := model.TServerSeg{
		ServerID: "999",
		SegPref: "prefix",
		Ver: "111111",
	}
	err = tss.AddOrUpdate(tServerSeg)
	if err != nil {
		t.Fatalf("error adding or updating t_server_seg: %v", err)
	}

	_, err = db.Exec("drop table %s", "t_server_seg")
	if err != nil {
		t.Fatalf("error droping table %s: %v", "t_server_seg", err)
	}
}
