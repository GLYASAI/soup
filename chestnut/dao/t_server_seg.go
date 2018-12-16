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
	"fmt"
	"github.com/GLYASAI/soup/chestnut/dao/model"
	"github.com/twinj/uuid"
)

type TServerSegImpl struct {
	db *sql.DB
}

func (t *TServerSegImpl) AddOrUpdate(seg model.TServerSeg) error {
	rows, err := t.db.Query("select * from T_SERVER_SEG where SERVER_ID = :1", seg.ServerID)
	if err != nil {
		return fmt.Errorf("error selecting data from t_server_seg: %v", err)
	}
	defer rows.Close()

	// add
	if !rows.Next() {
		_, err := t.db.Exec("insert into T_SERVER_SEG(SERVER_SEG_ID, SERVER_ID, SEG_PREF, VER) VALUES (:1, :2, :3, :4)",
			uuid.NewV4(), seg.ServerID, seg.SegPref, seg.Ver)
		if err != nil {
			return fmt.Errorf("error inserting t_server_seg: %v", err)
		}
		return nil
	}
	// update
	var server_seg_id string
	err = rows.Scan(&server_seg_id)
	if err != nil {
		return fmt.Errorf("error scaning rows: %v", err)
	}
	_, err = t.db.Exec("update T_SERVER_SEG set SERVER_ID=:1, SEG_PREF=:2, VER=:3 WHERE SERVER_SEG_ID=:4",
		seg.ServerID, seg.SegPref, seg.Ver, server_seg_id)
	if err != nil {
		return fmt.Errorf("error updating t_server_seg: %v", err)
	}

	return nil
}

// TServerSegImpl deletes record in t_server_seg table by server_id
func (t *TServerSegImpl) Delete(serverID string) error {
	_, err := t.db.Exec("delete from T_SERVER_REG where SERVER_ID=:1", serverID)
	if err != nil {
		return fmt.Errorf("error deleting t_server_reg by server_id(%s): %v", serverID, err)
	}
	return nil
}
