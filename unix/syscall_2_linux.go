//go:build linux && !sw64
// +build linux,!sw64

package unix

func Pipe(p []int) error {
        return Pipe2(p, 0)
}

//sysnb pipe2(p *[2]_C_int, flags int) (err error)

func Pipe2(p []int, flags int) error {
        if len(p) != 2 {
                return EINVAL
        }
        var pp [2]_C_int
        err := pipe2(&pp, flags)
        if err == nil {
                p[0] = int(pp[0])
                p[1] = int(pp[1])
        }
        return err
}

func Poll(fds []PollFd, timeout int) (n int, err error) {
        var ts *Timespec
        if timeout >= 0 {
                ts = new(Timespec)
                *ts = NsecToTimespec(int64(timeout) * 1e6)
        }
        return Ppoll(fds, ts, nil)
}

func EpollCreate(size int) (fd int, err error) {
        if size <= 0 {
                return -1, EINVAL
        }
        return EpollCreate1(0)
}
