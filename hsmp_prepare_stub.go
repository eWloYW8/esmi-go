//go:build !linux || !cgo

package esmi

func ensureHSMPReady() error {
	return nil
}
