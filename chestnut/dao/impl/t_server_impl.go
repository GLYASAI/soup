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

package impl

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
)

// TServerImpl is the implementation of TServer
type TServerImpl struct {
	db *sql.DB
}

// GetServerIDByIP gets server_id from t_server by ipaddr
func (t *TServerImpl) GetServerIDByIP(ip string) (string, error) {
	var server_id string
	err := t.db.QueryRow("select IP_ADDR from t_server where IP_ADDR=:1", ip).Scan(&server_id)
	if err != nil {
		logrus.Errorf("error getting server_id by ip: %v", err)
		return "", fmt.Errorf("error getting server_id by ip: %v", err)
	}
	return server_id, nil
}