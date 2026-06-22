//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris

// Package shm implements System V shared memory functions (shmctl, shmget, shmat, shmdt).
//
// This platform does not provide System V shared memory; every function returns
// errUnsupported so that programs depending on the package still build.
package shm

import "errors"

// Constants.
const (
	// Mode bits for `shmget`.

	// Create key if key does not exist.
	IPC_CREAT = 01000
	// Fail if key exists.
	IPC_EXCL = 02000
	// Return error on wait.
	IPC_NOWAIT = 04000

	// Special key values.

	// Private key.
	IPC_PRIVATE = 0

	// Flags for `shmat`.

	// Attach read-only access.
	SHM_RDONLY = 010000
	// Round attach address to SHMLBA.
	SHM_RND = 020000
	// Take-over region on attach.
	SHM_REMAP = 040000
	// Execution access.
	SHM_EXEC = 0100000

	// Commands for `shmctl`.

	// Lock segment (root only).
	SHM_LOCK = 1
	// Unlock segment (root only).
	SHM_UNLOCK = 12

	// Control commands for `shmctl`.

	// Remove identifier.
	IPC_RMID = 0
	// Set `ipc_perm` options.
	IPC_SET = 1
	// Get `ipc_perm' options.
	IPC_STAT = 2
)

var errUnsupported = errors.New("shm: unsupported platform")

// Perm is used to pass permission information to IPC operations.
type Perm struct{}

// IdDs describes shared memory segment.
type IdDs struct {
	// Operation permission struct.
	Perm Perm
	// Size of segment in bytes.
	SegSz uint64
}

// Get allocates a shared memory segment.
func Get(key int, size int, shmFlg int) (shmId int, err error) {
	return -1, errUnsupported
}

// At attaches the shared memory segment identified by shmId.
func At(shmId int, shmAddr uintptr, shmFlg int) (data []byte, err error) {
	return nil, errUnsupported
}

// Dt detaches the shared memory segment.
func Dt(data []byte) error {
	return errUnsupported
}

// Ctl performs the control operation specified by cmd on the shared memory segment whose identifier is given in shmId.
func Ctl(shmId int, cmd int, buf *IdDs) (int, error) {
	return -1, errUnsupported
}

// Rm removes the shared memory segment.
func Rm(shmId int) error {
	return errUnsupported
}

// Size returns size of shared memory segment.
func Size(shmId int) (int64, error) {
	return 0, errUnsupported
}
