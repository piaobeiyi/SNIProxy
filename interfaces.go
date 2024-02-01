package main

import (
	"net"

	"golang.org/x/net/proxy"
)

func GetDialer(isSocks5 bool) proxy.Dialer {
	if !isSocks5 {
		return &net.Dialer{}
	}
	auth := &proxy.Auth{User: cfg.SocksUser, Password: cfg.SocksPass}
	if auth.User == "" && auth.Password == "" {
		auth = nil
	}
	proxyDialer, err := proxy.SOCKS5("tcp", cfg.SocksAddr, auth, proxy.Direct)
	if err != nil {
		// FIXME: I am shit
		return &net.Dialer{}
	}
	return proxyDialer
}
