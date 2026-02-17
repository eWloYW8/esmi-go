//go:build linux && cgo

package esmi

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var hsmpModules = [...]string{"hsmp_acpi", "amd_hsmp"}

func ensureHSMPReady() error {
	if err := loadHSMPModule(); err != nil {
		return err
	}

	info, err := os.Stat(HSMPCharDevfileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("esmi: hsmp device %s not found after module load: %w", HSMPCharDevfileName, err)
		}
		return fmt.Errorf("esmi: failed to check hsmp device %s: %w", HSMPCharDevfileName, err)
	}
	if info.Mode()&os.ModeCharDevice == 0 {
		return fmt.Errorf("esmi: hsmp path %s is not a character device", HSMPCharDevfileName)
	}

	dev, err := os.OpenFile(HSMPCharDevfileName, os.O_RDWR, 0)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return fmt.Errorf("esmi: no permission to access hsmp device %s: %w", HSMPCharDevfileName, err)
		}
		return fmt.Errorf("esmi: cannot access hsmp device %s: %w", HSMPCharDevfileName, err)
	}
	if err := dev.Close(); err != nil {
		return fmt.Errorf("esmi: failed to close hsmp device %s: %w", HSMPCharDevfileName, err)
	}
	return nil
}

func loadHSMPModule() error {
	modprobe, err := findModprobe()
	if err != nil {
		return fmt.Errorf("esmi: %w", err)
	}

	var errs []string
	for _, module := range hsmpModules {
		cmd := exec.Command(modprobe, module)
		out, cmdErr := cmd.CombinedOutput()
		if cmdErr == nil {
			return nil
		}

		detail := strings.TrimSpace(string(out))
		if detail == "" {
			errs = append(errs, fmt.Sprintf("%s: %v", module, cmdErr))
			continue
		}
		errs = append(errs, fmt.Sprintf("%s: %v: %s", module, cmdErr, detail))
	}

	return fmt.Errorf("esmi: failed to load hsmp kernel module (%s)", strings.Join(errs, "; "))
}

func findModprobe() (string, error) {
	if p, err := exec.LookPath("modprobe"); err == nil {
		return p, nil
	}
	for _, p := range []string{"/sbin/modprobe", "/usr/sbin/modprobe"} {
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p, nil
		}
	}
	return "", errors.New("modprobe not found in PATH, /sbin, or /usr/sbin")
}
