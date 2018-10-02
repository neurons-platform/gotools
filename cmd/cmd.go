package cmd

type OsType int
type CMDType int

const (
	Linux OsType = iota
	Windows
	Mac
	Other
)

const (
	HTTP CMDType = iota
	GRPC
	SSH
	RPC
)

type CMD struct {
	CMDType CMDType
	HostName string
	IP string
	Port int
	User string
	Password string
	Path string
	OsType OsType
	Command string
	Out CMDOut
	TimeOut int
}

type CMDOut interface {
	GetStdOut() string
	GetStdErr() string
}
