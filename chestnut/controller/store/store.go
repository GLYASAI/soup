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

package store

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

// NotExistsError is returned when an object does not exist in a local store.
type NotExistsError string

// Error implements the error interface.
func (e NotExistsError) Error() string {
	return fmt.Sprintf("no object matching key %q in local store", string(e))
}

type store struct {
	informers *Informer
	listers   *Lister
}

// Lister contains object listers (stores).
type Lister struct {
	Endpoint EndpointLister
}

func New(client kubernetes.Interface, ns string) {
	s := store{
		informers: &Informer{},
		listers:   &Lister{},
	}

	// create informers factory, enable and assign required informers
	infFactory := informers.NewFilteredSharedInformerFactory(client, time.Second, ns,
		func(options *metav1.ListOptions) {
			//options.LabelSelector = "creater=Rainbond"
		})

	s.informers.Endpoint = infFactory.Core().V1().Endpoints().Informer()
	s.listers.Endpoint.Store = s.informers.Endpoint.GetStore()

	// Endpoint Event Handler
	epEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// TODO
		},
		DeleteFunc: func(obj interface{}) {
			// TODO
		},
		UpdateFunc: func(old, cur interface{}) {
			oep := old.(*corev1.Endpoints)
			cep := cur.(*corev1.Endpoints)
			_ = oep
			_ = cep
			// TODO
		},
	}

	s.informers.Endpoint.AddEventHandlerWithResyncPeriod(epEventHandler, 10*time.Second)
}
