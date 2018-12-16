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

package controller

import (
	"github.com/GLYASAI/soup/chestnut/controller/store"
	"github.com/GLYASAI/soup/chestnut/util"
	"github.com/GLYASAI/soup/cmd/chestnut/option"
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	cfg       *option.Config
	K8sClient kubernetes.Interface
}

func New(cfg *option.Config) (*Controller, error) {
	clientset, err := util.NewCK8sClient(cfg.K8SConfPath)
	if err != nil {
		return nil, err
	}

	c := &Controller{
		cfg:       cfg,
		K8sClient: clientset,
	}

	return c, nil
}

func (c *Controller) Start() {
	store.New(c.K8sClient, c.cfg.Namespace)
}
