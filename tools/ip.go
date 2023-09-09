package tools

import (
	"github.com/ipipdotnet/ipdb-go"
	"net"
	"path/filepath"
	"strings"
)

// GetIpFromAddr
// 从一个网络地址（net.Addr）中提取出对应的 IPv4 地址。
func GetIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	// 如果 addr 是 *net.IPNet 类型，说明它是一个网络地址（通常是子网），将其中的 IPv4 地址提取出来，赋值给 ip。
	case *net.IPNet:
		ip = v.IP
	// 如果 addr 是 *net.IPAddr 类型，说明它是一个具体的 IP 地址，同样将其中的 IPv4 地址提取出来，赋值给 ip。
	case *net.IPAddr:
		ip = v.IP
	}
	// 检查 ip 是否为空或者是否是回环地址（loopback）。如果 ip 为空（nil）或者是回环地址，就返回 nil，表示无法提取有效的 IPv4 地址。
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	// 将 IPv4 地址转换为 4 字节的形式（IPv4 地址通常是 4 字节的）
	ip = ip.To4()
	if ip == nil {
		return nil
	}

	return ip
}

// GetOutboundIP
// 获取本机计算机用于出站连接的 IP 地址(通常为公网ip地址)
func GetOutboundIP() (net.IP, error) {
	// 使用 net.Dial 函数创建一个 UDP 连接，连接到目标地址 "8.8.8.8:80"。这里使用的是 Google 的 DNS 服务器地址 8.8.8.8 和端口 80。
	// 使用一个可靠的公共服务器地址，如 Google 的 DNS 服务器，是一种常见的方法来测试网络连接是否正常。
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	//  conn.LocalAddr() 方法获取本地端点的地址信息，并将其强制类型转换为 *net.UDPAddr 类型。
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 提取本地出站 IP 地址
	return localAddr.IP, nil
}

// ParseIp 解析ip到城市信息
func ParseIp(ip string) *ipdb.CityInfo {
	realPath, err := filepath.Abs("../config")
	realPath = strings.Replace(realPath+"/city.free.ipdb", "\\", "/", -1)
	db, err := ipdb.NewCity(realPath)
	if err != nil {
		return nil
	}
	db.Reload(realPath)
	c, err := db.FindInfo(ip, "CN")
	if err != nil {
		return nil
	}
	return c
}
