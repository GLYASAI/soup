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
	"github.com/GLYASAI/soup/chestnut/dao"
	"github.com/GLYASAI/soup/chestnut/dao/model"
	"github.com/Sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"reflect"
	"time"
)

// NotExistsError is returned when an object does not exist in a local Store.
type NotExistsError string

// Error implements the error interface.
func (e NotExistsError) Error() string {
	return fmt.Sprintf("no object matching key %q in local Store", string(e))
}

type Store struct {
	informers *Informer
	listers   *Lister
}

// Lister contains object listers (stores).
type Lister struct {
	Pod PodLister
}

func New(client kubernetes.Interface, ns string, ts *dao.TServerImpl, tss *dao.TServerSegImpl) Store {
	s := Store{
		informers: &Informer{},
		listers:   &Lister{},
	}

	// create informers factory, enable and assign required informers
	infFactory := informers.NewFilteredSharedInformerFactory(client, time.Second, ns,
		func(options *metav1.ListOptions) {})

	s.informers.Pod = infFactory.Core().V1().Pods().Informer()
	s.listers.Pod.Store = s.informers.Pod.GetStore()

	// Endpoint Event Handler
	podEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				logrus.Warningf("can not convert %s to *corev1.Endpoints", reflect.TypeOf(obj), pod)
				return
			}
			if pod.Status.Phase != corev1.PodRunning {
				return
			}

			for _, c := range pod.Spec.Containers {
				var segpref string
				var ver string
				for _, e := range c.Env {
					if e.Name == "SEGPREF" {
						segpref = e.Value
					} else if e.Name == "VER" {
						ver = e.Value
					}
				}
				if segpref == "" || ver == "" {
					continue
				}

				serverID, err := ts.GetServerIDByIP(pod.Status.HostIP)
				if err != nil {
					logrus.Warningf("can not get server_id by ip(%s): %v", pod.Status.HostIP, err)
					continue
				}
				seg := model.TServerSeg{serverID, segpref, ver}
				if err := tss.AddOrUpdate(seg); err != nil {
					logrus.Warningf("can not add or update t_server_seg: %v", err)
					continue
				}
				break
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				logrus.Warningf("can not convert %s to *corev1.Endpoints", reflect.TypeOf(obj), pod)
				return
			}

			serverID, err := ts.GetServerIDByIP(pod.Status.HostIP)
			if err != nil {
				logrus.Warningf("cant not get t_server by ip(%s): %v", pod.Status.HostIP, err)
			}

			if err := tss.Delete(serverID); err != nil {
				logrus.Warningf("can not deleting t_server_reg: %v", err)
			}
		},
		UpdateFunc: func(old, cur interface{}) {
			opod, ok := old.(*corev1.Pod)
			if !ok {
				logrus.Warningf("can not convert %s to *corev1.Endpoints", reflect.TypeOf(old), opod)
				return
			}
			cpod, ok := cur.(*corev1.Pod)
			if !ok {
				logrus.Warningf("can not convert %s to *corev1.Endpoints", reflect.TypeOf(cur), cpod)
				return
			}
			// ignore the same secret as the old one
			if opod.ResourceVersion == cpod.ResourceVersion || reflect.DeepEqual(opod, cpod) {
				return
			}
			for _, c := range cpod.Spec.Containers {
				var segpref string
				var ver string
				for _, e := range c.Env {
					if e.Name == "SEGPREF" {
						segpref = e.Value
					} else if e.Name == "VER" {
						ver = e.Value
					}
				}
				if segpref == "" || ver == "" {
					continue
				}

				serverID, err := ts.GetServerIDByIP(opod.Status.HostIP)
				if err != nil {
					logrus.Warningf("can not get server_id by ip(%s): %v", opod.Status.HostIP, err)
					continue
				}
				seg := model.TServerSeg{serverID, segpref, ver}
				if err := tss.AddOrUpdate(seg); err != nil {
					logrus.Warningf("can not add or update t_server_seg: %v", err)
					continue
				}
				break
			}
		},
	}

	s.informers.Pod.AddEventHandlerWithResyncPeriod(podEventHandler, 10*time.Second)

	return s
}

func (s *Store) Run(stopCh chan struct{}) {
	s.informers.Run(stopCh)
}
