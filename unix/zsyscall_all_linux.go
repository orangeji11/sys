//go:build linux && sw64
// +build linux,sw64

package unix

import "unsafe"

// copied from zsyscall_linux.go for sw64 compiling
func ioctlPtr(fd int, req uint, arg unsafe.Pointer) (err error) {
	_, _, e1 := Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func shmctl(id int, cmd int, buf *SysvShmDesc) (result int, err error) {
	r0, _, e1 := Syscall(SYS_SHMCTL, uintptr(id), uintptr(cmd), uintptr(unsafe.Pointer(buf)))
	result = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func shmat(id int, addr uintptr, flag int) (ret uintptr, err error) {
	r0, _, e1 := Syscall(SYS_SHMAT, uintptr(id), uintptr(addr), uintptr(flag))
	ret = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func shmdt(addr uintptr) (err error) {
	_, _, e1 := Syscall(SYS_SHMDT, uintptr(addr), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func shmget(key int, size int, flag int) (id int, err error) {
	r0, _, e1 := Syscall(SYS_SHMGET, uintptr(key), uintptr(size), uintptr(flag))
	id = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func getitimer(which int, currValue *Itimerval) (err error) {
	_, _, e1 := Syscall(SYS_GETITIMER, uintptr(which), uintptr(unsafe.Pointer(currValue)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func setitimer(which int, newValue *Itimerval, oldValue *Itimerval) (err error) {
	_, _, e1 := Syscall(SYS_SETITIMER, uintptr(which), uintptr(unsafe.Pointer(newValue)), uintptr(unsafe.Pointer(oldValue)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// copied from zsyscall_linux.go for sw64 compiling
func mountSetattr(dirfd int, pathname string, flags uint, attr *MountAttr, size uintptr) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(pathname)
	if err != nil {
		return
	}
	_, _, e1 := Syscall6(SYS_MOUNT_SETATTR, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(unsafe.Pointer(attr)), uintptr(size), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
