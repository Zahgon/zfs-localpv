package zfs

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

// maxEventMessageLen is the Kubernetes API server hard limit for Event.Message.
// Messages longer than this are rejected with a validation error.
const maxEventMessageLen = 1024

func truncateEventMessage(msg string) string { _ = "STUB: not implemented"; return "" }

// EmitFailureEvent records a Warning event on obj. The event message is
// err.Error(), which for *ZFSError includes the verbatim ZFS stderr so
// the real failure cause is visible in kubectl describe output.
func EmitFailureEvent(recorder record.EventRecorder, obj runtime.Object, reason string, err error) {
	_ = "STUB: not implemented"
	return
}

// EmitSuccessEvent records a Normal event on obj for terminal
// user-visible transitions (provisioned, destroyed, etc.).
func EmitSuccessEvent(recorder record.EventRecorder, obj runtime.Object, reason, message string) {
	_ = "STUB: not implemented"
	return
}
