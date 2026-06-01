package utils

import (
	"errors"
	"net"
	"net/url"
)

// IsInternalHost 检查 host 是否解析为内部地址（loopback、private、link-local）
func IsInternalHost(host string) bool {
	ip := net.ParseIP(host)
	if ip != nil {
		return ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
	}
	ips, err := net.LookupIP(host)
	if err != nil {
		return false
	}
	for _, ip := range ips {
		if ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
			return true
		}
	}
	return false
}

// CheckURLForSSRF 解析 URL 并校验其 Host 是否指向内部地址
func CheckURLForSSRF(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	if IsInternalHost(u.Hostname()) {
		return errors.New("内部地址被禁止")
	}
	return nil
}
