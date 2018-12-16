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
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

// EndpointLister makes a store that lists Endpoints.
type EndpointLister struct {
	cache.Store
}

// ByKey returns the Endpoints of the Service matching key in the local Endpoint store.
func (s *EndpointLister) ByKey(key string) (*apiv1.Endpoints, error) {
	eps, exists, err := s.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotExistsError(key)
	}
	return eps.(*apiv1.Endpoints), nil
}
