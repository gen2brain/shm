package shm

// Perm is used to pass permission information to IPC operations.
type Perm struct {
	// Key.
	Key int32
	// Owner's user ID.
	Uid uint32
	// Owner's group ID.
	Gid uint32
	// Creator's user ID.
	Cuid uint32
	// Creator's group ID.
	Cgid uint32
	// Read/write permission.
	Mode uint16
	// Padding.
	X__pad1 uint16
	// Sequence number.
	X__seq uint16
	// Padding.
	X__pad2 uint16
	// Reserved.
	X__glibc_reserved1 uint32
	// Reserved.
	X__glibc_reserved2 uint32
}

// IdDs describes shared memory segment.
type IdDs struct {
	// Operation permission struct.
	Perm Perm
	// Size of segment in bytes.
	SegSz uint32
	// Last attach time.
	Atime int32
	// Reserved.
	X__glibc_reserved1 uint32
	// Last detach time.
	Dtime int32
	// Reserved.
	X__glibc_reserved2 uint32
	// Last change time.
	Ctime int32
	// Reserved.
	X__glibc_reserved3 uint32
	// Pid of creator.
	Cpid int32
	// Pid of last shmat/shmdt.
	Lpid int32
	// Number of current attaches.
	Nattch uint32
	// Reserved.
	X__glibc_reserved4 uint32
	// Reserved.
	X__glibc_reserved5 uint32
}
