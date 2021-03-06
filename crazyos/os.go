/**
 * Auth :   liubo
 * Date :   2019/12/23 11:26
 * Comment: 丰富系统函数
 */

package crazyos

import (
	"github.com/badforlabor/gocrazy/crazyio"
	"github.com/kardianos/osext"
	"net"
	"os"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetExecFolder() string {
	var ret, _ = osext.ExecutableFolder()
	return ret
}

func GetAppName() string {
	if len(os.Args) > 0 {
		return crazyio.GetExeName(os.Args[0])
	}

	return ""
}
