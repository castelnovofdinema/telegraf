//go:build !windows
// +build !windows

package portname

func servicesPath() string {
	return "/etc/services"
}
