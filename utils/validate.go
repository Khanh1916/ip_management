package utils
/* Chua cac ham kiem tra tinh hop le cua dia chi IP, 
   su dung trong cac APi de xac thuc dau vao truoc khi luu database hoa xu ly. */

import (
	"net"	// Chua cac phuong thuc xu ly va xac thuc dia chi IP
)

// IsValidIP kiem tra tinh hop le cua IP (IPv4 hoac IPv6)
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil	// Neu invalid ip thi bang nil, tra ve false va nguoc lai.
} 

// IsIPv4 kiem tra dia chi IPv4 hop le
func IsIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil	// Neu valid ip va co dinh dang V4 thi tra ve true
}

// IsIPv6 kiem tra dia chi IPv6 hop le
func IsIPv6(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() == nil	// Neu valid ip va khong co dinh dang V4 tra ve true
}