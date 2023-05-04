package wasi

import (
	"encoding/binary"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/stealthrocket/wazergo/types"
)

func TestErrors(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(Errno(0)), 2)

	assertEqual(t, int(ESUCCESS), 0)
	assertEqual(t, int(E2BIG), 1)
	assertEqual(t, int(EACCES), 2)
	assertEqual(t, int(EADDRINUSE), 3)
	assertEqual(t, int(EADDRNOTAVAIL), 4)
	assertEqual(t, int(EAFNOSUPPORT), 5)
	assertEqual(t, int(EAGAIN), 6)
	assertEqual(t, int(EALREADY), 7)
	assertEqual(t, int(EBADF), 8)
	assertEqual(t, int(EBADMSG), 9)
	assertEqual(t, int(EBUSY), 10)
	assertEqual(t, int(ECANCELED), 11)
	assertEqual(t, int(ECHILD), 12)
	assertEqual(t, int(ECONNABORTED), 13)
	assertEqual(t, int(ECONNREFUSED), 14)
	assertEqual(t, int(ECONNRESET), 15)
	assertEqual(t, int(EDEADLK), 16)
	assertEqual(t, int(EDESTADDRREQ), 17)
	assertEqual(t, int(EDOM), 18)
	assertEqual(t, int(EDQUOT), 19)
	assertEqual(t, int(EEXIST), 20)
	assertEqual(t, int(EFAULT), 21)
	assertEqual(t, int(EFBIG), 22)
	assertEqual(t, int(EHOSTUNREACH), 23)
	assertEqual(t, int(EIDRM), 24)
	assertEqual(t, int(EILSEQ), 25)
	assertEqual(t, int(EINPROGRESS), 26)
	assertEqual(t, int(EINTR), 27)
	assertEqual(t, int(EINVAL), 28)
	assertEqual(t, int(EIO), 29)
	assertEqual(t, int(EISCONN), 30)
	assertEqual(t, int(EISDIR), 31)
	assertEqual(t, int(ELOOP), 32)
	assertEqual(t, int(EMFILE), 33)
	assertEqual(t, int(EMLINK), 34)
	assertEqual(t, int(EMSGSIZE), 35)
	assertEqual(t, int(EMULTIHOP), 36)
	assertEqual(t, int(ENAMETOOLONG), 37)
	assertEqual(t, int(ENETDOWN), 38)
	assertEqual(t, int(ENETRESET), 39)
	assertEqual(t, int(ENETUNREACH), 40)
	assertEqual(t, int(ENFILE), 41)
	assertEqual(t, int(ENOBUFS), 42)
	assertEqual(t, int(ENODEV), 43)
	assertEqual(t, int(ENOENT), 44)
	assertEqual(t, int(ENOEXEC), 45)
	assertEqual(t, int(ENOLCK), 46)
	assertEqual(t, int(ENOLINK), 47)
	assertEqual(t, int(ENOMEM), 48)
	assertEqual(t, int(ENOMSG), 49)
	assertEqual(t, int(ENOPROTOOPT), 50)
	assertEqual(t, int(ENOSPC), 51)
	assertEqual(t, int(ENOSYS), 52)
	assertEqual(t, int(ENOTCONN), 53)
	assertEqual(t, int(ENOTDIR), 54)
	assertEqual(t, int(ENOTEMPTY), 55)
	assertEqual(t, int(ENOTRECOVERABLE), 56)
	assertEqual(t, int(ENOTSOCK), 57)
	assertEqual(t, int(ENOTSUP), 58)
	assertEqual(t, int(ENOTTY), 59)
	assertEqual(t, int(ENXIO), 60)
	assertEqual(t, int(EOVERFLOW), 61)
	assertEqual(t, int(EOWNERDEAD), 62)
	assertEqual(t, int(EPERM), 63)
	assertEqual(t, int(EPIPE), 64)
	assertEqual(t, int(EPROTO), 65)
	assertEqual(t, int(EPROTONOSUPPORT), 66)
	assertEqual(t, int(EPROTOTYPE), 67)
	assertEqual(t, int(ERANGE), 68)
	assertEqual(t, int(EROFS), 69)
	assertEqual(t, int(ESPIPE), 70)
	assertEqual(t, int(ESRCH), 71)
	assertEqual(t, int(ESTALE), 72)
	assertEqual(t, int(ETIMEDOUT), 73)
	assertEqual(t, int(ETXTBSY), 74)
	assertEqual(t, int(EXDEV), 75)
	assertEqual(t, int(ENOTCAPABLE), 76)

	assertEqual(t, E2BIG.Error(), "Argument list too long")
	assertEqual(t, EACCES.Error(), "Permission denied")
	assertEqual(t, EADDRINUSE.Error(), "Address already in use")
	assertEqual(t, EADDRNOTAVAIL.Error(), "Address not available")
	assertEqual(t, EAFNOSUPPORT.Error(), "Address family not supported by protocol family")
	assertEqual(t, EAGAIN.Error(), "Try again")
	assertEqual(t, EALREADY.Error(), "Socket already connected")
	assertEqual(t, EBADF.Error(), "Bad file number")
	assertEqual(t, EBADMSG.Error(), "Trying to read unreadable message")
	assertEqual(t, EBUSY.Error(), "Device or resource busy")
	assertEqual(t, ECANCELED.Error(), "Operation canceled.")
	assertEqual(t, ECHILD.Error(), "No child processes")
	assertEqual(t, ECONNABORTED.Error(), "Connection aborted")
	assertEqual(t, ECONNREFUSED.Error(), "Connection refused")
	assertEqual(t, ECONNRESET.Error(), "Connection reset by peer")
	assertEqual(t, EDEADLK.Error(), "Deadlock condition")
	assertEqual(t, EDESTADDRREQ.Error(), "Destination address required")
	assertEqual(t, EDOM.Error(), "Math arg out of domain of func")
	assertEqual(t, EDQUOT.Error(), "Quota exceeded")
	assertEqual(t, EEXIST.Error(), "File exists")
	assertEqual(t, EFAULT.Error(), "Bad address")
	assertEqual(t, EFBIG.Error(), "File too large")
	assertEqual(t, EHOSTUNREACH.Error(), "Host is unreachable")
	assertEqual(t, EIDRM.Error(), "Identifier removed")
	assertEqual(t, EILSEQ.Error(), "Illegal byte sequence")
	assertEqual(t, EINPROGRESS.Error(), "Connection already in progress")
	assertEqual(t, EINTR.Error(), "Interrupted system call")
	assertEqual(t, EINVAL.Error(), "Invalid argument")
	assertEqual(t, EIO.Error(), "I/O error")
	assertEqual(t, EISCONN.Error(), "Socket is already connected")
	assertEqual(t, EISDIR.Error(), "Is a directory")
	assertEqual(t, ELOOP.Error(), "Too many symbolic links")
	assertEqual(t, EMFILE.Error(), "Too many open files")
	assertEqual(t, EMLINK.Error(), "Too many links")
	assertEqual(t, EMSGSIZE.Error(), "Message too long")
	assertEqual(t, EMULTIHOP.Error(), "Multihop attempted")
	assertEqual(t, ENAMETOOLONG.Error(), "File name too long")
	assertEqual(t, ENETDOWN.Error(), "Network interface is not configured")
	assertEqual(t, ENETRESET.Error(), "Network dropped connection on reset")
	assertEqual(t, ENETUNREACH.Error(), "Network is unreachable")
	assertEqual(t, ENFILE.Error(), "File table overflow")
	assertEqual(t, ENOBUFS.Error(), "No buffer space available")
	assertEqual(t, ENODEV.Error(), "No such device")
	assertEqual(t, ENOENT.Error(), "No such file or directory")
	assertEqual(t, ENOEXEC.Error(), "Exec format error")
	assertEqual(t, ENOLCK.Error(), "No record locks available")
	assertEqual(t, ENOLINK.Error(), "The link has been severed")
	assertEqual(t, ENOMEM.Error(), "Out of memory")
	assertEqual(t, ENOMSG.Error(), "No message of desired type")
	assertEqual(t, ENOPROTOOPT.Error(), "Protocol not available")
	assertEqual(t, ENOSPC.Error(), "No space left on device")
	assertEqual(t, ENOSYS.Error(), "Not implemented")
	assertEqual(t, ENOTCONN.Error(), "Socket is not connected")
	assertEqual(t, ENOTDIR.Error(), "Not a directory")
	assertEqual(t, ENOTEMPTY.Error(), "Directory not empty")
	assertEqual(t, ENOTRECOVERABLE.Error(), "State not recoverable")
	assertEqual(t, ENOTSOCK.Error(), "Socket operation on non-socket")
	assertEqual(t, ENOTSUP.Error(), "Not supported")
	assertEqual(t, ENOTTY.Error(), "Not a typewriter")
	assertEqual(t, ENXIO.Error(), "No such device or address")
	assertEqual(t, EOVERFLOW.Error(), "Value too large for defined data type")
	assertEqual(t, EOWNERDEAD.Error(), "Owner died")
	assertEqual(t, EPERM.Error(), "Operation not permitted")
	assertEqual(t, EPIPE.Error(), "Broken pipe")
	assertEqual(t, EPROTO.Error(), "Protocol error")
	assertEqual(t, EPROTONOSUPPORT.Error(), "Unknown protocol")
	assertEqual(t, EPROTOTYPE.Error(), "Protocol wrong type for socket")
	assertEqual(t, ERANGE.Error(), "Math result not representable")
	assertEqual(t, EROFS.Error(), "Read-only file system")
	assertEqual(t, ESPIPE.Error(), "Illegal seek")
	assertEqual(t, ESRCH.Error(), "No such process")
	assertEqual(t, ESTALE.Error(), "Stale file handle")
	assertEqual(t, ETIMEDOUT.Error(), "Connection timed out")
	assertEqual(t, ETXTBSY.Error(), "Text file busy")
	assertEqual(t, EXDEV.Error(), "Cross-device link")
	assertEqual(t, ENOTCAPABLE.Error(), "Capabilities insufficient")

	assertEqual(t, E2BIG.Name(), "E2BIG")
	assertEqual(t, EACCES.Name(), "EACCES")
	assertEqual(t, EADDRINUSE.Name(), "EADDRINUSE")
	assertEqual(t, EADDRNOTAVAIL.Name(), "EADDRNOTAVAIL")
	assertEqual(t, EAFNOSUPPORT.Name(), "EAFNOSUPPORT")
	assertEqual(t, EAGAIN.Name(), "EAGAIN")
	assertEqual(t, EALREADY.Name(), "EALREADY")
	assertEqual(t, EBADF.Name(), "EBADF")
	assertEqual(t, EBADMSG.Name(), "EBADMSG")
	assertEqual(t, EBUSY.Name(), "EBUSY")
	assertEqual(t, ECANCELED.Name(), "ECANCELED")
	assertEqual(t, ECHILD.Name(), "ECHILD")
	assertEqual(t, ECONNABORTED.Name(), "ECONNABORTED")
	assertEqual(t, ECONNREFUSED.Name(), "ECONNREFUSED")
	assertEqual(t, ECONNRESET.Name(), "ECONNRESET")
	assertEqual(t, EDEADLK.Name(), "EDEADLK")
	assertEqual(t, EDESTADDRREQ.Name(), "EDESTADDRREQ")
	assertEqual(t, EDOM.Name(), "EDOM")
	assertEqual(t, EDQUOT.Name(), "EDQUOT")
	assertEqual(t, EEXIST.Name(), "EEXIST")
	assertEqual(t, EFAULT.Name(), "EFAULT")
	assertEqual(t, EFBIG.Name(), "EFBIG")
	assertEqual(t, EHOSTUNREACH.Name(), "EHOSTUNREACH")
	assertEqual(t, EIDRM.Name(), "EIDRM")
	assertEqual(t, EILSEQ.Name(), "EILSEQ")
	assertEqual(t, EINPROGRESS.Name(), "EINPROGRESS")
	assertEqual(t, EINTR.Name(), "EINTR")
	assertEqual(t, EINVAL.Name(), "EINVAL")
	assertEqual(t, EIO.Name(), "EIO")
	assertEqual(t, EISCONN.Name(), "EISCONN")
	assertEqual(t, EISDIR.Name(), "EISDIR")
	assertEqual(t, ELOOP.Name(), "ELOOP")
	assertEqual(t, EMFILE.Name(), "EMFILE")
	assertEqual(t, EMLINK.Name(), "EMLINK")
	assertEqual(t, EMSGSIZE.Name(), "EMSGSIZE")
	assertEqual(t, EMULTIHOP.Name(), "EMULTIHOP")
	assertEqual(t, ENAMETOOLONG.Name(), "ENAMETOOLONG")
	assertEqual(t, ENETDOWN.Name(), "ENETDOWN")
	assertEqual(t, ENETRESET.Name(), "ENETRESET")
	assertEqual(t, ENETUNREACH.Name(), "ENETUNREACH")
	assertEqual(t, ENFILE.Name(), "ENFILE")
	assertEqual(t, ENOBUFS.Name(), "ENOBUFS")
	assertEqual(t, ENODEV.Name(), "ENODEV")
	assertEqual(t, ENOENT.Name(), "ENOENT")
	assertEqual(t, ENOEXEC.Name(), "ENOEXEC")
	assertEqual(t, ENOLCK.Name(), "ENOLCK")
	assertEqual(t, ENOLINK.Name(), "ENOLINK")
	assertEqual(t, ENOMEM.Name(), "ENOMEM")
	assertEqual(t, ENOMSG.Name(), "ENOMSG")
	assertEqual(t, ENOPROTOOPT.Name(), "ENOPROTOOPT")
	assertEqual(t, ENOSPC.Name(), "ENOSPC")
	assertEqual(t, ENOSYS.Name(), "ENOSYS")
	assertEqual(t, ENOTCONN.Name(), "ENOTCONN")
	assertEqual(t, ENOTDIR.Name(), "ENOTDIR")
	assertEqual(t, ENOTEMPTY.Name(), "ENOTEMPTY")
	assertEqual(t, ENOTRECOVERABLE.Name(), "ENOTRECOVERABLE")
	assertEqual(t, ENOTSOCK.Name(), "ENOTSOCK")
	assertEqual(t, ENOTSUP.Name(), "ENOTSUP")
	assertEqual(t, ENOTTY.Name(), "ENOTTY")
	assertEqual(t, ENXIO.Name(), "ENXIO")
	assertEqual(t, EOVERFLOW.Name(), "EOVERFLOW")
	assertEqual(t, EOWNERDEAD.Name(), "EOWNERDEAD")
	assertEqual(t, EPERM.Name(), "EPERM")
	assertEqual(t, EPIPE.Name(), "EPIPE")
	assertEqual(t, EPROTO.Name(), "EPROTO")
	assertEqual(t, EPROTONOSUPPORT.Name(), "EPROTONOSUPPORT")
	assertEqual(t, EPROTOTYPE.Name(), "EPROTOTYPE")
	assertEqual(t, ERANGE.Name(), "ERANGE")
	assertEqual(t, EROFS.Name(), "EROFS")
	assertEqual(t, ESPIPE.Name(), "ESPIPE")
	assertEqual(t, ESRCH.Name(), "ESRCH")
	assertEqual(t, ESTALE.Name(), "ESTALE")
	assertEqual(t, ETIMEDOUT.Name(), "ETIMEDOUT")
	assertEqual(t, ETXTBSY.Name(), "ETXTBSY")
	assertEqual(t, EXDEV.Name(), "EXDEV")
	assertEqual(t, ENOTCAPABLE.Name(), "ENOTCAPABLE")
}

func TestExec(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(ExitCode(0)), 4)

	assertEqual(t, unsafe.Sizeof(Signal(0)), 1)

	assertEqual(t, int(SIGNONE), 0)
	assertEqual(t, int(SIGHUP), 1)
	assertEqual(t, int(SIGINT), 2)
	assertEqual(t, int(SIGQUIT), 3)
	assertEqual(t, int(SIGILL), 4)
	assertEqual(t, int(SIGTRAP), 5)
	assertEqual(t, int(SIGABRT), 6)
	assertEqual(t, int(SIGBUS), 7)
	assertEqual(t, int(SIGFPE), 8)
	assertEqual(t, int(SIGKILL), 9)
	assertEqual(t, int(SIGUSR1), 10)
	assertEqual(t, int(SIGSEGV), 11)
	assertEqual(t, int(SIGUSR2), 12)
	assertEqual(t, int(SIGPIPE), 13)
	assertEqual(t, int(SIGALRM), 14)
	assertEqual(t, int(SIGTERM), 15)
	assertEqual(t, int(SIGCHLD), 16)
	assertEqual(t, int(SIGCONT), 17)
	assertEqual(t, int(SIGSTOP), 18)
	assertEqual(t, int(SIGTSTP), 19)
	assertEqual(t, int(SIGTTIN), 20)
	assertEqual(t, int(SIGTTOU), 21)
	assertEqual(t, int(SIGURG), 22)
	assertEqual(t, int(SIGXCPU), 23)
	assertEqual(t, int(SIGXFSZ), 24)
	assertEqual(t, int(SIGVTALRM), 25)
	assertEqual(t, int(SIGPROF), 26)
	assertEqual(t, int(SIGWINCH), 27)
	assertEqual(t, int(SIGPOLL), 28)
	assertEqual(t, int(SIGPWR), 29)
	assertEqual(t, int(SIGSYS), 30)

	assertEqual(t, SIGNONE.String(), "no signal")
	assertEqual(t, SIGHUP.String(), "hangup")
	assertEqual(t, SIGINT.String(), "interrupt")
	assertEqual(t, SIGQUIT.String(), "quit")
	assertEqual(t, SIGILL.String(), "illegal instruction")
	assertEqual(t, SIGTRAP.String(), "trace/breakpoint trap")
	assertEqual(t, SIGABRT.String(), "abort")
	assertEqual(t, SIGBUS.String(), "bus error")
	assertEqual(t, SIGFPE.String(), "floating point exception")
	assertEqual(t, SIGKILL.String(), "killed")
	assertEqual(t, SIGUSR1.String(), "user-defined signal 1")
	assertEqual(t, SIGSEGV.String(), "segmentation fault")
	assertEqual(t, SIGUSR2.String(), "user-defined signal 2")
	assertEqual(t, SIGPIPE.String(), "broken pipe")
	assertEqual(t, SIGALRM.String(), "alarm clock")
	assertEqual(t, SIGTERM.String(), "terminated")
	assertEqual(t, SIGCHLD.String(), "child exited")
	assertEqual(t, SIGCONT.String(), "continued")
	assertEqual(t, SIGSTOP.String(), "stopped (signal)")
	assertEqual(t, SIGTSTP.String(), "stopped")
	assertEqual(t, SIGTTIN.String(), "stopped (tty input)")
	assertEqual(t, SIGTTOU.String(), "stopped (tty output)")
	assertEqual(t, SIGURG.String(), "urgent I/O condition")
	assertEqual(t, SIGXCPU.String(), "CPU time limit exceeded")
	assertEqual(t, SIGXFSZ.String(), "file size limit exceeded")
	assertEqual(t, SIGVTALRM.String(), "virtual timer expired")
	assertEqual(t, SIGPROF.String(), "profiling timer expired")
	assertEqual(t, SIGWINCH.String(), "window changed")
	assertEqual(t, SIGPOLL.String(), "I/O possible")
	assertEqual(t, SIGPWR.String(), "power failure")
	assertEqual(t, SIGSYS.String(), "bad system call")

	assertEqual(t, SIGNONE.Name(), "SIGNONE")
	assertEqual(t, SIGHUP.Name(), "SIGHUP")
	assertEqual(t, SIGINT.Name(), "SIGINT")
	assertEqual(t, SIGQUIT.Name(), "SIGQUIT")
	assertEqual(t, SIGILL.Name(), "SIGILL")
	assertEqual(t, SIGTRAP.Name(), "SIGTRAP")
	assertEqual(t, SIGABRT.Name(), "SIGABRT")
	assertEqual(t, SIGBUS.Name(), "SIGBUS")
	assertEqual(t, SIGFPE.Name(), "SIGFPE")
	assertEqual(t, SIGKILL.Name(), "SIGKILL")
	assertEqual(t, SIGUSR1.Name(), "SIGUSR1")
	assertEqual(t, SIGSEGV.Name(), "SIGSEGV")
	assertEqual(t, SIGUSR2.Name(), "SIGUSR2")
	assertEqual(t, SIGPIPE.Name(), "SIGPIPE")
	assertEqual(t, SIGALRM.Name(), "SIGALRM")
	assertEqual(t, SIGTERM.Name(), "SIGTERM")
	assertEqual(t, SIGCHLD.Name(), "SIGCHLD")
	assertEqual(t, SIGCONT.Name(), "SIGCONT")
	assertEqual(t, SIGSTOP.Name(), "SIGSTOP")
	assertEqual(t, SIGTSTP.Name(), "SIGTSTP")
	assertEqual(t, SIGTTIN.Name(), "SIGTTIN")
	assertEqual(t, SIGTTOU.Name(), "SIGTTOU")
	assertEqual(t, SIGURG.Name(), "SIGURG")
	assertEqual(t, SIGXCPU.Name(), "SIGXCPU")
	assertEqual(t, SIGXFSZ.Name(), "SIGXFSZ")
	assertEqual(t, SIGVTALRM.Name(), "SIGVTALRM")
	assertEqual(t, SIGPROF.Name(), "SIGPROF")
	assertEqual(t, SIGWINCH.Name(), "SIGWINCH")
	assertEqual(t, SIGPOLL.Name(), "SIGPOLL")
	assertEqual(t, SIGPWR.Name(), "SIGPWR")
	assertEqual(t, SIGSYS.Name(), "SIGSYS")
}

func TestFile(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(FD(0)), 4)
	assertEqual(t, unsafe.Sizeof(INode(0)), 8)
	assertEqual(t, unsafe.Sizeof(Device(0)), 8)
	assertEqual(t, unsafe.Sizeof(FileSize(0)), 8)
	assertEqual(t, unsafe.Sizeof(FileDelta(0)), 8)
	assertEqual(t, unsafe.Sizeof(LinkCount(0)), 8)

	assertEqual(t, unsafe.Sizeof(FileStat{}), 64)
	assertEqual(t, unsafe.Sizeof(FileStat{}.Device), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.INode), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.FileType), 1)
	assertEqual(t, unsafe.Sizeof(FileStat{}.NLink), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.Size), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.AccessTime), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.ModifyTime), 8)
	assertEqual(t, unsafe.Sizeof(FileStat{}.ChangeTime), 8)
	assertEqual(t, unsafe.Offsetof(FileStat{}.Device), 0)
	assertEqual(t, unsafe.Offsetof(FileStat{}.INode), 8)
	assertEqual(t, unsafe.Offsetof(FileStat{}.FileType), 16)
	assertEqual(t, unsafe.Offsetof(FileStat{}.NLink), 24)
	assertEqual(t, unsafe.Offsetof(FileStat{}.Size), 32)
	assertEqual(t, unsafe.Offsetof(FileStat{}.AccessTime), 40)
	assertEqual(t, unsafe.Offsetof(FileStat{}.ModifyTime), 48)
	assertEqual(t, unsafe.Offsetof(FileStat{}.ChangeTime), 56)

	assertEqual(t, unsafe.Sizeof(Whence(0)), 1)
	assertEqual(t, SeekStart, 0)
	assertEqual(t, SeekCurrent, 1)
	assertEqual(t, SeekEnd, 2)
	assertEqual(t, SeekStart.String(), "SeekStart")
	assertEqual(t, SeekCurrent.String(), "SeekCurrent")
	assertEqual(t, SeekEnd.String(), "SeekEnd")

	assertEqual(t, unsafe.Sizeof(FileType(0)), 1)
	assertEqual(t, UnknownType, 0)
	assertEqual(t, BlockDeviceType, 1)
	assertEqual(t, CharacterDeviceType, 2)
	assertEqual(t, DirectoryType, 3)
	assertEqual(t, RegularFileType, 4)
	assertEqual(t, SocketDGramType, 5)
	assertEqual(t, SocketStreamType, 6)
	assertEqual(t, SymbolicLinkType, 7)
	assertEqual(t, UnknownType.String(), "UnknownType")
	assertEqual(t, BlockDeviceType.String(), "BlockDeviceType")
	assertEqual(t, CharacterDeviceType.String(), "CharacterDeviceType")
	assertEqual(t, DirectoryType.String(), "DirectoryType")
	assertEqual(t, RegularFileType.String(), "RegularFileType")
	assertEqual(t, SocketDGramType.String(), "SocketDGramType")
	assertEqual(t, SocketStreamType.String(), "SocketStreamType")
	assertEqual(t, SymbolicLinkType.String(), "SymbolicLinkType")

	assertEqual(t, unsafe.Sizeof(FDFlags(0)), 2)
	assertEqual(t, Append, 1<<0)
	assertEqual(t, DSync, 1<<1)
	assertEqual(t, NonBlock, 1<<2)
	assertEqual(t, RSync, 1<<3)
	assertEqual(t, Sync, 1<<4)
	assertEqual(t, Append.String(), "Append")
	assertEqual(t, DSync.String(), "DSync")
	assertEqual(t, NonBlock.String(), "NonBlock")
	assertEqual(t, RSync.String(), "RSync")
	assertEqual(t, Sync.String(), "Sync")
	assertEqual(t, (DSync | NonBlock).String(), "DSync|NonBlock")

	assertEqual(t, unsafe.Sizeof(FDStat{}), 24)
	assertEqual(t, unsafe.Sizeof(FDStat{}.FileType), 1)
	assertEqual(t, unsafe.Offsetof(FDStat{}.FileType), 0)
	assertEqual(t, unsafe.Sizeof(FDStat{}.Flags), 2)
	assertEqual(t, unsafe.Offsetof(FDStat{}.Flags), 2)
	assertEqual(t, unsafe.Sizeof(FDStat{}.RightsBase), 8)
	assertEqual(t, unsafe.Offsetof(FDStat{}.RightsBase), 8)
	assertEqual(t, unsafe.Sizeof(FDStat{}.RightsInheriting), 8)
	assertEqual(t, unsafe.Offsetof(FDStat{}.RightsInheriting), 16)

	assertEqual(t, unsafe.Sizeof(Rights(0)), 8)

	assertEqual(t, FDDataSyncRight, 1<<0)
	assertEqual(t, FDReadRight, 1<<1)
	assertEqual(t, FDSeekRight, 1<<2)
	assertEqual(t, FDStatSetFlagsRight, 1<<3)
	assertEqual(t, FDSyncRight, 1<<4)
	assertEqual(t, FDTellRight, 1<<5)
	assertEqual(t, FDWriteRight, 1<<6)
	assertEqual(t, FDAdviseRight, 1<<7)
	assertEqual(t, FDAllocateRight, 1<<8)
	assertEqual(t, PathCreateDirectoryRight, 1<<9)
	assertEqual(t, PathCreateFileRight, 1<<10)
	assertEqual(t, PathLinkSourceRight, 1<<11)
	assertEqual(t, PathLinkTargetRight, 1<<12)
	assertEqual(t, PathOpenRight, 1<<13)
	assertEqual(t, FDReadDirRight, 1<<14)
	assertEqual(t, PathReadLinkRight, 1<<15)
	assertEqual(t, PathRenameSourceRight, 1<<16)
	assertEqual(t, PathRenameTargetRight, 1<<17)
	assertEqual(t, PathFileStatGetRight, 1<<18)
	assertEqual(t, PathFileStatSetSizeRight, 1<<19)
	assertEqual(t, PathFileStatSetTimesRight, 1<<20)
	assertEqual(t, FDFileStatGetRight, 1<<21)
	assertEqual(t, FDFileStatSetSizeRight, 1<<22)
	assertEqual(t, FDFileStatSetTimesRight, 1<<23)
	assertEqual(t, PathSymlinkRight, 1<<24)
	assertEqual(t, PathRemoveDirectoryRight, 1<<25)
	assertEqual(t, PathUnlinkFileRight, 1<<26)
	assertEqual(t, PollFDReadWriteRight, 1<<27)
	assertEqual(t, SockShutdownRight, 1<<28)
	assertEqual(t, SockAcceptRight, 1<<29)
	for i := 0; i <= 29; i++ {
		assertEqual(t, AllRights.Has(1<<i), true)
	}
	for i := 30; i < 64; i++ {
		assertEqual(t, AllRights.Has(1<<i), false)
	}
	assertEqual(t, FDDataSyncRight.String(), "FDDataSyncRight")
	assertEqual(t, FDReadRight.String(), "FDReadRight")
	assertEqual(t, FDSeekRight.String(), "FDSeekRight")
	assertEqual(t, FDStatSetFlagsRight.String(), "FDStatSetFlagsRight")
	assertEqual(t, FDSyncRight.String(), "FDSyncRight")
	assertEqual(t, FDTellRight.String(), "FDTellRight")
	assertEqual(t, FDWriteRight.String(), "FDWriteRight")
	assertEqual(t, FDAdviseRight.String(), "FDAdviseRight")
	assertEqual(t, FDAllocateRight.String(), "FDAllocateRight")
	assertEqual(t, PathCreateDirectoryRight.String(), "PathCreateDirectoryRight")
	assertEqual(t, PathCreateFileRight.String(), "PathCreateFileRight")
	assertEqual(t, PathLinkSourceRight.String(), "PathLinkSourceRight")
	assertEqual(t, PathLinkTargetRight.String(), "PathLinkTargetRight")
	assertEqual(t, PathOpenRight.String(), "PathOpenRight")
	assertEqual(t, FDReadDirRight.String(), "FDReadDirRight")
	assertEqual(t, PathReadLinkRight.String(), "PathReadLinkRight")
	assertEqual(t, PathRenameSourceRight.String(), "PathRenameSourceRight")
	assertEqual(t, PathRenameTargetRight.String(), "PathRenameTargetRight")
	assertEqual(t, PathFileStatGetRight.String(), "PathFileStatGetRight")
	assertEqual(t, PathFileStatSetSizeRight.String(), "PathFileStatSetSizeRight")
	assertEqual(t, PathFileStatSetTimesRight.String(), "PathFileStatSetTimesRight")
	assertEqual(t, FDFileStatGetRight.String(), "FDFileStatGetRight")
	assertEqual(t, FDFileStatSetSizeRight.String(), "FDFileStatSetSizeRight")
	assertEqual(t, FDFileStatSetTimesRight.String(), "FDFileStatSetTimesRight")
	assertEqual(t, PathSymlinkRight.String(), "PathSymlinkRight")
	assertEqual(t, PathRemoveDirectoryRight.String(), "PathRemoveDirectoryRight")
	assertEqual(t, PathUnlinkFileRight.String(), "PathUnlinkFileRight")
	assertEqual(t, PollFDReadWriteRight.String(), "PollFDReadWriteRight")
	assertEqual(t, SockShutdownRight.String(), "SockShutdownRight")
	assertEqual(t, SockAcceptRight.String(), "SockAcceptRight")
	assertEqual(t, (FDFileStatGetRight | PathSymlinkRight).String(), "FDFileStatGetRight|PathSymlinkRight")
	assertEqual(t, Rights(0).String(), "Rights(0)")
	assertEqual(t, AllRights.String(), "AllRights")
	assertEqual(t, Rights(math.MaxUint32).String(), "AllRights")

	assertEqual(t, unsafe.Sizeof(DirCookie(0)), 8)
	assertEqual(t, unsafe.Sizeof(DirNameLength(0)), 4)

	assertEqual(t, unsafe.Sizeof(Advice(0)), 1)
	assertEqual(t, Normal, 0)
	assertEqual(t, Sequential, 1)
	assertEqual(t, Random, 2)
	assertEqual(t, WillNeed, 3)
	assertEqual(t, DontNeed, 4)
	assertEqual(t, NoReuse, 5)
	assertEqual(t, Normal.String(), "Normal")
	assertEqual(t, Sequential.String(), "Sequential")
	assertEqual(t, Random.String(), "Random")
	assertEqual(t, WillNeed.String(), "WillNeed")
	assertEqual(t, DontNeed.String(), "DontNeed")
	assertEqual(t, NoReuse.String(), "NoReuse")

	assertEqual(t, unsafe.Sizeof(FSTFlags(0)), 2)
	assertEqual(t, AccessTime, 1<<0)
	assertEqual(t, AccessTimeNow, 1<<1)
	assertEqual(t, ModifyTime, 1<<2)
	assertEqual(t, ModifyTimeNow, 1<<3)
	assertEqual(t, AccessTime.String(), "AccessTime")
	assertEqual(t, AccessTimeNow.String(), "AccessTimeNow")
	assertEqual(t, ModifyTime.String(), "ModifyTime")
	assertEqual(t, ModifyTimeNow.String(), "ModifyTimeNow")
	assertEqual(t, (AccessTime | ModifyTimeNow).String(), "AccessTime|ModifyTimeNow")

	assertEqual(t, unsafe.Sizeof(LookupFlags(0)), 4)
	assertEqual(t, SymlinkFollow, 1<<0)
	assertEqual(t, SymlinkFollow.String(), "SymlinkFollow")

	assertEqual(t, unsafe.Sizeof(OpenFlags(0)), 2)
	assertEqual(t, OpenCreate, 1<<0)
	assertEqual(t, OpenDirectory, 1<<1)
	assertEqual(t, OpenExclusive, 1<<2)
	assertEqual(t, OpenTruncate, 1<<3)
	assertEqual(t, OpenCreate.String(), "OpenCreate")
	assertEqual(t, OpenDirectory.String(), "OpenDirectory")
	assertEqual(t, OpenExclusive.String(), "OpenExclusive")
	assertEqual(t, OpenTruncate.String(), "OpenTruncate")
	assertEqual(t, (OpenCreate | OpenExclusive | OpenTruncate).String(), "OpenCreate|OpenExclusive|OpenTruncate")

	assertEqual(t, unsafe.Sizeof(PreOpenType(0)), 1)
	assertEqual(t, PreOpenDir, 0)
	assertEqual(t, PreOpenDir.String(), "PreOpenDir")

	assertEqual(t, unsafe.Sizeof(PreStatDir{}), 4)
	assertEqual(t, unsafe.Sizeof(PreStatDir{}.NameLength), 4)
	assertEqual(t, unsafe.Offsetof(PreStatDir{}.NameLength), 0)

	assertEqual(t, unsafe.Sizeof(PreStat{}), 8)
	assertEqual(t, unsafe.Sizeof(PreStat{}.Type), 1)
	assertEqual(t, unsafe.Offsetof(PreStat{}.Type), 0)
	assertEqual(t, unsafe.Sizeof(PreStat{}.PreStatDir), 4)
	assertEqual(t, unsafe.Offsetof(PreStat{}.PreStatDir), 4)

	assertEqual(t, unsafe.Sizeof(Size(0)), 4)
}

func TestSubscription(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(Subscription{}), 48)
	assertEqual(t, unsafe.Offsetof(Subscription{}.UserData), 0)
	assertEqual(t, unsafe.Sizeof(Subscription{}.UserData), 8)
	assertEqual(t, unsafe.Offsetof(Subscription{}.EventType), 8)
	assertEqual(t, unsafe.Offsetof(Subscription{}.variant), 16)
	assertEqual(t, unsafe.Sizeof(Subscription{}.variant), 32)

	assertEqual(t, unsafe.Sizeof(SubscriptionFDReadWrite{}), 4)
	assertEqual(t, unsafe.Offsetof(SubscriptionFDReadWrite{}.FD), 0)
	assertEqual(t, unsafe.Sizeof(SubscriptionFDReadWrite{}.FD), 4)
	assertEqual(t, format(SubscriptionFDReadWrite{FD: 42}), `{FD:42}`)

	now := Timestamp(time.Date(2023, 5, 2, 12, 2, 0, 0, time.UTC).UnixNano())
	assertEqual(t, unsafe.Sizeof(SubscriptionClock{}), 32)
	assertEqual(t, unsafe.Offsetof(SubscriptionClock{}.ID), 0)
	assertEqual(t, unsafe.Offsetof(SubscriptionClock{}.Timeout), 8)
	assertEqual(t, unsafe.Offsetof(SubscriptionClock{}.Precision), 16)
	assertEqual(t, unsafe.Offsetof(SubscriptionClock{}.Flags), 24)
	assertEqual(t, format(SubscriptionClock{ID: Monotonic, Timeout: 10e3, Precision: 1}), `{ID:Monotonic,Timeout:10µs,Precision:1ns}`)
	assertEqual(t, format(SubscriptionClock{ID: Monotonic, Timeout: now, Precision: 1, Flags: Abstime}), `{ID:Monotonic,Timeout:2023-05-02T12:02:00Z,Precision:1ns}`)

	assertEqual(t, unsafe.Sizeof(SubscriptionClockFlags(0)), 2)
	assertEqual(t, Abstime, 0x1)
	assertEqual(t, Abstime.String(), "Abstime")
}

func format(v types.Formatter) string {
	s := new(strings.Builder)
	v.Format(s)
	return s.String()
}

func TestSubscriptionFDReadWrite(t *testing.T) {
	actual := MakeSubscriptionFDReadWrite(0xFEEDF4CED00DCAFE, FDReadEvent, SubscriptionFDReadWrite{FD(0xABCD)})

	expected := Subscription{
		UserData:  0xFEEDF4CED00DCAFE,
		EventType: FDReadEvent,
		variant:   [32]byte{},
	}
	binary.LittleEndian.PutUint32(expected.variant[0:4], uint32(0xABCD))
	assertEqual(t, actual, expected)

	variant := actual.GetFDReadWrite()
	assertEqual(t, variant.FD, FD(0xABCD))
}

func TestSubscriptionClock(t *testing.T) {
	actual := MakeSubscriptionClock(0xFEEDF4CED00DCAFE, SubscriptionClock{
		ID:        Realtime,
		Timeout:   0xABCD,
		Precision: 0x5678,
		Flags:     Abstime,
	})

	expected := Subscription{
		UserData:  0xFEEDF4CED00DCAFE,
		EventType: ClockEvent,
		variant:   [32]byte{},
	}
	binary.LittleEndian.PutUint32(expected.variant[0:4], uint32(Realtime))
	binary.LittleEndian.PutUint64(expected.variant[8:16], uint64(0xABCD))
	binary.LittleEndian.PutUint64(expected.variant[16:24], uint64(0x5678))
	binary.LittleEndian.PutUint16(expected.variant[24:26], uint16(Abstime))
	assertEqual(t, actual, expected)

	variant := actual.GetClock()
	assertEqual(t, variant.ID, Realtime)
	assertEqual(t, variant.Timeout, 0xABCD)
	assertEqual(t, variant.Precision, 0x5678)
	assertEqual(t, variant.Flags, Abstime)
}

func TestEvent(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(Event{}), 32)
	assertEqual(t, unsafe.Offsetof(Event{}.UserData), 0)
	assertEqual(t, unsafe.Sizeof(Event{}.UserData), 8)
	assertEqual(t, unsafe.Offsetof(Event{}.Errno), 8)
	assertEqual(t, unsafe.Sizeof(Event{}.Errno), 2)
	assertEqual(t, unsafe.Offsetof(Event{}.EventType), 10)
	assertEqual(t, unsafe.Offsetof(Event{}.FDReadWrite), 16)

	assertEqual(t, unsafe.Sizeof(EventFDReadWrite{}), 16)
	assertEqual(t, unsafe.Offsetof(EventFDReadWrite{}.NBytes), 0)
	assertEqual(t, unsafe.Sizeof(EventFDReadWrite{}.NBytes), 8)
	assertEqual(t, unsafe.Offsetof(EventFDReadWrite{}.Flags), 8)

	assertEqual(t, unsafe.Sizeof(EventType(0)), 1)
	assertEqual(t, ClockEvent, 0)
	assertEqual(t, FDReadEvent, 1)
	assertEqual(t, FDWriteEvent, 2)
	assertEqual(t, ClockEvent.String(), "ClockEvent")
	assertEqual(t, FDReadEvent.String(), "FDReadEvent")
	assertEqual(t, FDWriteEvent.String(), "FDWriteEvent")

	assertEqual(t, unsafe.Sizeof(EventFDReadWriteFlags(0)), 2)
	assertEqual(t, Hangup, 0x1)
	assertEqual(t, Hangup.String(), "Hangup")
}

func TestSocket(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(RIFlags(0)), 2)
	assertEqual(t, RecvPeek, 1<<0)
	assertEqual(t, RecvWaitAll, 1<<1)
	assertEqual(t, RecvPeek.String(), "RecvPeek")
	assertEqual(t, RecvWaitAll.String(), "RecvWaitAll")
	assertEqual(t, (RecvPeek | RecvWaitAll).String(), "RecvPeek|RecvWaitAll")

	assertEqual(t, unsafe.Sizeof(ROFlags(0)), 2)
	assertEqual(t, RecvDataTruncated, 1<<0)
	assertEqual(t, RecvDataTruncated.String(), "RecvDataTruncated")

	assertEqual(t, unsafe.Sizeof(SIFlags(0)), 2)

	assertEqual(t, unsafe.Sizeof(SDFlags(0)), 2)
	assertEqual(t, ShutdownRD, 1<<0)
	assertEqual(t, ShutdownWR, 1<<1)
	assertEqual(t, ShutdownRD.String(), "ShutdownRD")
	assertEqual(t, ShutdownWR.String(), "ShutdownWR")
	assertEqual(t, (ShutdownRD | ShutdownWR).String(), "ShutdownRD|ShutdownWR")
}

func TestTime(t *testing.T) {
	assertEqual(t, unsafe.Sizeof(Timestamp(0)), 8)

	assertEqual(t, unsafe.Sizeof(ClockID(0)), 4)
	assertEqual(t, Realtime, 0)
	assertEqual(t, Monotonic, 1)
	assertEqual(t, ProcessCPUTimeID, 2)
	assertEqual(t, ThreadCPUTimeID, 3)
	assertEqual(t, Realtime.String(), "Realtime")
	assertEqual(t, Monotonic.String(), "Monotonic")
	assertEqual(t, ProcessCPUTimeID.String(), "ProcessCPUTimeID")
	assertEqual(t, ThreadCPUTimeID.String(), "ThreadCPUTimeID")
}

func assertEqual[T any](t *testing.T, actual, expected T) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v != %v", actual, expected)
	}
}
