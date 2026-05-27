package driver

import (
	"sync"
)

type volumeLock struct {
	cond   sync.Cond
	locked map[string]struct{}
}

func newVolumeLock() *volumeLock { _ = "STUB: not implemented"; return nil }

func (l *volumeLock) LockVolume(volume string) func() { _ = "STUB: not implemented"; return nil }

func (l *volumeLock) LockVolumeWithSnapshot(volume string, snapshot string) func() {
	_ = "STUB: not implemented"
	return nil
}
