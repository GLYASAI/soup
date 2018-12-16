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

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GLYASAI/soup/cmd/chestnut/option"
	"github.com/GLYASAI/soup/cmd/chestnut/server"
	"github.com/mattn/go-oci8"
	"github.com/spf13/pflag"
	"log"
	"os"
	"time"
)

func main() {
	c := option.NewConfig()
	c.AddFlags(pflag.CommandLine)
	pflag.Parse()
	c.SetLog()
	if err := server.Run(c); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main2() {
	oci8.OCI8Driver.Logger = log.New(os.Stderr, "oci8 ", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)

	// [username/[password]@]host[:port][/instance_name][?param1=value1&...&paramN=valueN]
	// A normal simple Open to localhost would look like:
	// db, err := sql.Open("oci8", "127.0.0.1")
	// For testing, need to use additional variables
	db, err := sql.Open("oci8", "huangrh/12345678@127.0.0.1:1521/helowin")
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
		}
	}()

	var rows *sql.Rows
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()
	rows, err = db.QueryContext(ctx, "select 1 from dual")
	if err != nil {
		fmt.Println("QueryContext error is not nil:", err)
		return
	}
	if !rows.Next() {
		fmt.Println("no Next rows")
		return
	}

	dest := make([]interface{}, 1)
	destPointer := make([]interface{}, 1)
	destPointer[0] = &dest[0]
	err = rows.Scan(destPointer...)
	if err != nil {
		fmt.Println("Scan error is not nil:", err)
		return
	}

	if len(dest) != 1 {
		fmt.Println("len dest != 1")
		return
	}
	data, ok := dest[0].(float64)
	if !ok {
		fmt.Println("dest type not float64")
		return
	}
	if data != 1 {
		fmt.Println("data not equal to 1")
		return
	}

	if rows.Next() {
		fmt.Println("has Next rows")
		return
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Err error is not nil:", err)
		return
	}
	err = rows.Close()
	if err != nil {
		fmt.Println("Close error is not nil:", err)
		return
	}

	fmt.Println(data)
}
