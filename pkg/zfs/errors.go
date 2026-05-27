package zfs

// Event reason constants for Kubernetes Warning/Normal events on ZFS CRs.
// These are operation-level identifiers used as the reason field in
// recorder.Event calls — they name what failed, not why.
const (
	ReasonProvisionFailed   = "ProvisionFailed"
	ReasonCloneFailed       = "CloneFailed"
	ReasonDestroyFailed     = "DestroyFailed"
	ReasonSetPropertyFailed = "SetPropertyFailed"
	ReasonResizeFailed      = "ResizeFailed"
	ReasonSnapshotFailed    = "SnapshotFailed"
	ReasonSnapDestroyFailed = "SnapshotDestroyFailed"
	ReasonBackupFailed      = "BackupFailed"
	ReasonRestoreFailed     = "RestoreFailed"

	ReasonProvisioned      = "Provisioned"
	ReasonCloned           = "Cloned"
	ReasonDestroyed        = "Destroyed"
	ReasonPropertySet      = "PropertySet"
	ReasonResized          = "Resized"
	ReasonSnapshotted      = "Snapshotted"
	ReasonSnapDestroyed    = "SnapshotDestroyed"
	ReasonBackupCompleted  = "BackupCompleted"
	ReasonRestoreCompleted = "RestoreCompleted"
)

// Error is the typed error returned by every zfs/zpool shell-out.
// It preserves the raw stderr from the binary verbatim — no pattern
// matching or state probing — so the actual ZFS message always reaches
// logs and Kubernetes events unchanged.
type Error struct {
	Op      string // e.g. "zfs create", "zfs destroy"
	Dataset string // target dataset (may be empty)
	Stderr  string // trimmed combined output from the binary
	Err     error  // original error (exec.ExitError, etc.)
}

// Error implements the error interface and is safe to emit verbatim as
// a Kubernetes Event message.
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap exposes the underlying exec error so errors.As/errors.Is
// work for callers that need to inspect the original exec.ExitError.
func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

// NewZFSError wraps a failed zfs/zpool exec result. Returns nil when
// err is nil so callers can use it unconditionally as a return value.
func NewZFSError(op, dataset string, err error, stderr []byte) error {
	_ = "STUB: not implemented"
	return nil
}
