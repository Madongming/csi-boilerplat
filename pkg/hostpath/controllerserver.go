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
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func (hp *hostPath) CreateVolume(_ context.Context, _ *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, finalErr error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your CreateVolume implement
	// ...

	return &csi.CreateVolumeResponse{}, nil
}

func (hp *hostPath) DeleteVolume(_ context.Context, _ *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your DeleteVolume implement
	// ...

	return &csi.DeleteVolumeResponse{}, nil
}

func (hp *hostPath) ControllerGetCapabilities(_ context.Context, _ *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	// TODO: your Capabilities
	return &csi.ControllerGetCapabilitiesResponse{}, nil
}

func (hp *hostPath) ValidateVolumeCapabilities(_ context.Context, _ *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement

	return &csi.ValidateVolumeCapabilitiesResponse{}, nil
}

func (hp *hostPath) ControllerPublishVolume(_ context.Context, _ *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	// TODO: Check arguments

	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (hp *hostPath) ControllerUnpublishVolume(_ context.Context, _ *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	// TODO: Check arguments

	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (hp *hostPath) GetCapacity(_ context.Context, _ *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.GetCapacityResponse{}, nil
}

func (hp *hostPath) ListVolumes(_ context.Context, _ *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	volumeRes := &csi.ListVolumesResponse{
		Entries: []*csi.ListVolumesResponse_Entry{},
	}

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...
	// fill in `volumeRes`

	glog.V(5).Infof("Volumes are: %+v", *volumeRes)
	return volumeRes, nil
}

func (hp *hostPath) ControllerGetVolume(_ context.Context, _ *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.ControllerGetVolumeResponse{}, nil
}

// CreateSnapshot uses tar command to create snapshot for hostpath volume. The tar command can quickly create
// archives of entire directories. The host image must have "tar" binaries in /bin, /usr/sbin, or /usr/bin.
func (hp *hostPath) CreateSnapshot(_ context.Context, _ *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.CreateSnapshotResponse{}, nil
}

func (hp *hostPath) DeleteSnapshot(_ context.Context, _ *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.DeleteSnapshotResponse{}, nil
}

func (hp *hostPath) ListSnapshots(_ context.Context, _ *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.ListSnapshotsResponse{}, nil
}

func (hp *hostPath) ControllerExpandVolume(_ context.Context, _ *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	// TODO: Check arguments

	// Lock before acting on global state. A production-quality
	// driver might use more fine-grained locking.
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	// TODO: your implement
	// ...

	return &csi.ControllerExpandVolumeResponse{}, nil
}
