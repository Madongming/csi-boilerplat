/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hostpath

import (
	"errors"
	"sync"
)

const (
	kib    int64 = 1024
	mib    int64 = kib * 1024
	gib    int64 = mib * 1024
	gib100 int64 = gib * 100
	tib    int64 = gib * 1024
	tib100 int64 = tib * 100
)

type hostPath struct {
	config Config

	// gRPC calls involving any of the fields below must be serialized
	// by locking this mutex before starting. Internal helper
	// functions assume that the mutex has been locked.
	mutex sync.Mutex
}

type Config struct {
	DriverName    string
	Endpoint      string
	NodeID        string
	VendorVersion string

	// TODO: Other config field
	// ...
}

func NewHostPathDriver(cfg Config) (*hostPath, error) {
	if cfg.DriverName == "" {
		return nil, errors.New("no driver name provided")
	}

	if cfg.NodeID == "" {
		return nil, errors.New("no node id provided")
	}

	if cfg.Endpoint == "" {
		return nil, errors.New("no driver endpoint provided")
	}

	// TODO: New your csi
	// ...

	hp := &hostPath{
		config: cfg,
	}
	return hp, nil
}

func (hp *hostPath) Run() error {
	s := NewNonBlockingGRPCServer()
	// hp itself implements ControllerServer, NodeServer, and IdentityServer.
	s.Start(hp.config.Endpoint, hp, hp, hp)
	s.Wait()

	return nil
}
