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

package option

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
)

// Config is used to configure the parameters required by the program
type Config struct {
	LogLevel    string
	K8SConfPath string
	Namespace   string
}

func NewConfig() *Config {
	return &Config{}
}

// AddFlags adds flags
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.LogLevel, "log-level", "debug", "the gateway log level")
	fs.StringVar(&c.K8SConfPath, "kube-conf", "/opt/rainbond/etc/kubernetes/kubecfg/admin.kubeconfig", "absolute path to the kubeconfig file.")
	fs.StringVar(&c.Namespace, "namespace", "chestnut", "the namespace of the resource that needs to watch in kubernetes.")
}

// SetLog sets the standard logger
func (c Config) SetLog() {
	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		fmt.Println("set log level error." + err.Error())
		return
	}
	logrus.SetLevel(level)
}
