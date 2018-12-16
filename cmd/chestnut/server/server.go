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

package server

import (
	"database/sql"
	"fmt"
	"github.com/GLYASAI/soup/chestnut/controller"
	"github.com/GLYASAI/soup/cmd/chestnut/option"
	"github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-oci8"
	"os"
	"os/signal"
	"syscall"
)

//Run starts the entire program
func Run(cfg *option.Config) error {
	db, err := sql.Open("oci8", "huangrh/12345678@127.0.0.1:1521/helowin")
	if err != nil {
		logrus.Error("Open error is not nil: %v", err)
		return fmt.Errorf("Open error is not nil: %v", err)
	}
	if db == nil {
		return fmt.Errorf("db is nil")
	}
	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			logrus.Errorf("Close error is not nil: %v", err)
		}
	}()

	errCh := make(chan error)
	stopCh := make(chan struct{})
	c, err := controller.New(cfg, db, stopCh)
	if err != nil {
		return err
	}
	c.Start()
	logrus.Info("Successfully start chestnut.")

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-stopCh:
		logrus.Warn("Received stop chan, exiting gracefully...")
	case <-term:
		logrus.Warn("Received SIGTERM, exiting gracefully...")
	case err := <-errCh:
		logrus.Errorf("Received a error %s, exiting gracefully...", err.Error())
	}
	logrus.Info("See you next time!")
	return nil
}
