package esmi

import (
	"errors"
	"fmt"
	"sync"
)

var (
	libMu      sync.Mutex
	clientRefs int
)

var (
	errNilClient       = errors.New("esmi: nil client")
	errClientClosed    = errors.New("esmi: client is closed")
	errLibraryNotReady = errors.New("esmi: library is not initialized")
)

type Error struct {
	Op     string
	Status Status
}

func (e *Error) Error() string {
	msg := ErrMsg(e.Status)
	if msg == "" {
		msg = e.Status.String()
	}
	if e.Op == "" {
		return fmt.Sprintf("esmi: %s (%d)", msg, e.Status)
	}
	return fmt.Sprintf("esmi %s: %s (%d)", e.Op, msg, e.Status)
}

func (e *Error) Unwrap() error {
	return e.Status
}

func IsStatus(err error, st Status) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Status == st
	}
	return false
}

type Client struct {
	closed bool
}

func NewClient() (*Client, error) {
	libMu.Lock()
	defer libMu.Unlock()

	if clientRefs == 0 {
		st := Init()
		if !st.OK() {
			return nil, &Error{Op: "Init", Status: st}
		}
	}

	clientRefs++
	return &Client{}, nil
}

func (c *Client) Close() error {
	libMu.Lock()
	defer libMu.Unlock()

	if c == nil {
		return errNilClient
	}
	if c.closed {
		return nil
	}

	c.closed = true
	if clientRefs > 0 {
		clientRefs--
	}
	if clientRefs == 0 {
		Exit()
	}
	return nil
}

func (c *Client) isReadyLocked() error {
	if c == nil {
		return errNilClient
	}
	if c.closed {
		return errClientClosed
	}
	if clientRefs <= 0 {
		return errLibraryNotReady
	}
	return nil
}

func (c *Client) call(op string, fn func() Status) error {
	libMu.Lock()
	defer libMu.Unlock()

	if err := c.isReadyLocked(); err != nil {
		return err
	}

	st := fn()
	if st.OK() {
		return nil
	}
	return &Error{Op: op, Status: st}
}

func (c *Client) CoreEnergy(coreInd uint32) (uint64, error) {
	var out uint64
	if err := c.call("CoreEnergy", func() Status { return CoreEnergyGet(coreInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SocketEnergy(socketIdx uint32) (uint64, error) {
	var out uint64
	if err := c.call("SocketEnergy", func() Status { return SocketEnergyGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) AllEnergies() ([]uint64, error) {
	var cpus uint32
	if err := c.call("NumberOfCPUs", func() Status { return NumberOfCPUsGet(&cpus) }); err != nil {
		return nil, err
	}
	if cpus == 0 {
		return nil, nil
	}
	out := make([]uint64, cpus)
	if err := c.call("AllEnergies", func() Status { return AllEnergiesGet(out) }); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) RaplUnitsHSMPMailbox(sockInd uint32) (uint8, uint8, error) {
	var tu, esu uint8
	if err := c.call("RaplUnitsHSMPMailbox", func() Status {
		return RaplUnitsHSMPMailboxGet(sockInd, &tu, &esu)
	}); err != nil {
		return 0, 0, err
	}
	return tu, esu, nil
}

func (c *Client) RaplPackageCounterHSMPMailbox(sockInd uint32) (uint32, uint32, error) {
	var c1, c0 uint32
	if err := c.call("RaplPackageCounterHSMPMailbox", func() Status {
		return RaplPackageCounterHSMPMailboxGet(sockInd, &c1, &c0)
	}); err != nil {
		return 0, 0, err
	}
	return c1, c0, nil
}

func (c *Client) RaplCoreCounterHSMPMailbox(coreInd uint32) (uint32, uint32, error) {
	var c1, c0 uint32
	if err := c.call("RaplCoreCounterHSMPMailbox", func() Status {
		return RaplCoreCounterHSMPMailboxGet(coreInd, &c1, &c0)
	}); err != nil {
		return 0, 0, err
	}
	return c1, c0, nil
}

func (c *Client) CoreEnergyHSMPMailbox(coreInd uint32) (uint64, error) {
	var out uint64
	if err := c.call("CoreEnergyHSMPMailbox", func() Status {
		return CoreEnergyHSMPMailboxGet(coreInd, &out)
	}); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) PackageEnergyHSMPMailbox(sockInd uint32) (uint64, error) {
	var out uint64
	if err := c.call("PackageEnergyHSMPMailbox", func() Status {
		return PackageEnergyHSMPMailboxGet(sockInd, &out)
	}); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) HSMPDriverVersion() (HSMPDriverVersion, error) {
	var out HSMPDriverVersion
	if err := c.call("HSMPDriverVersion", func() Status { return HSMPDriverVersionGet(&out) }); err != nil {
		return HSMPDriverVersion{}, err
	}
	return out, nil
}

func (c *Client) SMUFWVersion() (SMUFWVersion, error) {
	var out SMUFWVersion
	if err := c.call("SMUFWVersion", func() Status { return SMUFWVersionGet(&out) }); err != nil {
		return SMUFWVersion{}, err
	}
	return out, nil
}

func (c *Client) ProchotStatus(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("ProchotStatus", func() Status { return ProchotStatusGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) FclkMclk(socketIdx uint32) (uint32, uint32, error) {
	var fclk, mclk uint32
	if err := c.call("FclkMclk", func() Status { return FclkMclkGet(socketIdx, &fclk, &mclk) }); err != nil {
		return 0, 0, err
	}
	return fclk, mclk, nil
}

func (c *Client) CclkLimit(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("CclkLimit", func() Status { return CclkLimitGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) HSMPProtoVersion() (uint32, error) {
	var out uint32
	if err := c.call("HSMPProtoVersion", func() Status { return HSMPProtoVerGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SocketCurrentActiveFreqLimit(sockInd uint32) (uint16, []string, error) {
	var limit uint16
	var src []string
	if err := c.call("SocketCurrentActiveFreqLimit", func() Status {
		return SocketCurrentActiveFreqLimitGet(sockInd, &limit, &src)
	}); err != nil {
		return 0, nil, err
	}
	return limit, src, nil
}

func (c *Client) SocketFreqRange(sockInd uint8) (uint16, uint16, error) {
	var fmax, fmin uint16
	if err := c.call("SocketFreqRange", func() Status { return SocketFreqRangeGet(sockInd, &fmax, &fmin) }); err != nil {
		return 0, 0, err
	}
	return fmax, fmin, nil
}

func (c *Client) CurrentFreqLimitCore(coreID uint32) (uint32, error) {
	var out uint32
	if err := c.call("CurrentFreqLimitCore", func() Status { return CurrentFreqLimitCoreGet(coreID, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) CpurailIsofreqPolicy(sockInd uint8) (bool, error) {
	var out bool
	if err := c.call("CpurailIsofreqPolicy", func() Status { return CpurailIsofreqPolicyGet(sockInd, &out) }); err != nil {
		return false, err
	}
	return out, nil
}

func (c *Client) DfcCtrlSetting(sockInd uint8) (bool, error) {
	var out bool
	if err := c.call("DfcCtrlSetting", func() Status { return DfcCtrlSettingGet(sockInd, &out) }); err != nil {
		return false, err
	}
	return out, nil
}

func (c *Client) SocketPower(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("SocketPower", func() Status { return SocketPowerGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SocketPowerCap(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("SocketPowerCap", func() Status { return SocketPowerCapGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SocketPowerCapMax(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("SocketPowerCapMax", func() Status { return SocketPowerCapMaxGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) PwrSVITelemetryAllRails(sockInd uint32) (uint32, error) {
	var out uint32
	if err := c.call("PwrSVITelemetryAllRails", func() Status { return PwrSVITelemetryAllRailsGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) PwrEfficiencyMode(sockInd uint8) (uint8, uint32, uint32, error) {
	var mode uint8
	var util, ppt uint32
	if err := c.call("PwrEfficiencyMode", func() Status {
		return PwrEfficiencyModeGet(sockInd, &mode, &util, &ppt)
	}); err != nil {
		return 0, 0, 0, err
	}
	return mode, util, ppt, nil
}

func (c *Client) ReadCCDPower(coreID uint32) (uint32, error) {
	var out uint32
	if err := c.call("ReadCCDPower", func() Status { return ReadCCDPower(coreID, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SetSocketPowerCap(socketIdx uint32, cap uint32) error {
	return c.call("SetSocketPowerCap", func() Status { return SocketPowerCapSet(socketIdx, cap) })
}

func (c *Client) SetPwrEfficiencyMode(sockInd uint8, mode uint8, util uint32, pptLimit uint32) (uint32, uint32, error) {
	u := util
	p := pptLimit
	if err := c.call("SetPwrEfficiencyMode", func() Status {
		return PwrEfficiencyModeSet(sockInd, mode, &u, &p)
	}); err != nil {
		return 0, 0, err
	}
	return u, p, nil
}

func (c *Client) SetCpurailIsofreqPolicy(sockInd uint8, val bool) (bool, error) {
	out := val
	if err := c.call("SetCpurailIsofreqPolicy", func() Status { return CpurailIsofreqPolicySet(sockInd, &out) }); err != nil {
		return false, err
	}
	return out, nil
}

func (c *Client) SetDfcEnable(sockInd uint8, val bool) (bool, error) {
	out := val
	if err := c.call("SetDfcEnable", func() Status { return DfcEnableSet(sockInd, &out) }); err != nil {
		return false, err
	}
	return out, nil
}

func (c *Client) CoreBoostlimit(cpuInd uint32) (uint32, error) {
	var out uint32
	if err := c.call("CoreBoostlimit", func() Status { return CoreBoostlimitGet(cpuInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) Floorlimit(coreOrSockInd uint32, typ SetGetFloorlimit) (uint32, error) {
	var out uint32
	if err := c.call("Floorlimit", func() Status { return FloorlimitSetGet(coreOrSockInd, &out, typ) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SetFloorlimit(coreOrSockInd uint32, floorlimit uint32, typ SetGetFloorlimit) (uint32, error) {
	out := floorlimit
	if err := c.call("SetFloorlimit", func() Status { return FloorlimitSetGet(coreOrSockInd, &out, typ) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SetMSRFloorlimit(coreOrSockInd uint32, msrFloorlimit uint32, typ SetGetFloorlimit, fmax uint32) error {
	return c.call("SetMSRFloorlimit", func() Status {
		return MSRFloorlimitSet(coreOrSockInd, msrFloorlimit, typ, fmax)
	})
}

func (c *Client) SetSDPSLimit(sockInd uint8, sdpsLimit uint32) (uint32, error) {
	out := sdpsLimit
	if err := c.call("SetSDPSLimit", func() Status { return SDPSLimitSet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SDPSLimit(sockInd uint8) (uint32, error) {
	var out uint32
	if err := c.call("SDPSLimit", func() Status { return SDPSLimitGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SocketC0Residency(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("SocketC0Residency", func() Status { return SocketC0ResidencyGet(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SetCoreBoostlimit(cpuInd uint32, boostlimit uint32) error {
	return c.call("SetCoreBoostlimit", func() Status { return CoreBoostlimitSet(cpuInd, boostlimit) })
}

func (c *Client) SetSocketBoostlimit(socketIdx uint32, boostlimit uint32) error {
	return c.call("SetSocketBoostlimit", func() Status { return SocketBoostlimitSet(socketIdx, boostlimit) })
}

func (c *Client) DDRBW(sockInd uint8) (DDRBWMetrics, error) {
	var out DDRBWMetrics
	if err := c.call("DDRBW", func() Status { return DDRBWGet(sockInd, &out) }); err != nil {
		return DDRBWMetrics{}, err
	}
	return out, nil
}

func (c *Client) SocketTemperature(sockInd uint32) (uint32, error) {
	var out uint32
	if err := c.call("SocketTemperature", func() Status { return SocketTemperatureGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) ReadTdelta(sockInd uint8) (uint8, error) {
	var out uint8
	if err := c.call("ReadTdelta", func() Status { return ReadTdelta(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) GetSVI3VRControllerTemp(sockInd uint8, railSelection uint8, railIndex uint8) (SVI3Info, error) {
	inout := SVI3Info{RailSelection: railSelection, RailIndex: railIndex}
	if err := c.call("GetSVI3VRControllerTemp", func() Status { return GetSVI3VRControllerTemp(sockInd, &inout) }); err != nil {
		return SVI3Info{}, err
	}
	return inout, nil
}

func (c *Client) DimmTempRangeAndRefreshRate(sockInd uint8, dimmAddr uint8) (TempRangeRefreshRate, error) {
	var out TempRangeRefreshRate
	if err := c.call("DimmTempRangeAndRefreshRate", func() Status {
		return DimmTempRangeAndRefreshRateGet(sockInd, dimmAddr, &out)
	}); err != nil {
		return TempRangeRefreshRate{}, err
	}
	return out, nil
}

func (c *Client) DimmPowerConsumption(sockInd uint8, dimmAddr uint8) (DimmPower, error) {
	var out DimmPower
	if err := c.call("DimmPowerConsumption", func() Status {
		return DimmPowerConsumptionGet(sockInd, dimmAddr, &out)
	}); err != nil {
		return DimmPower{}, err
	}
	return out, nil
}

func (c *Client) DimmThermalSensor(sockInd uint8, dimmAddr uint8) (DimmThermal, error) {
	var out DimmThermal
	if err := c.call("DimmThermalSensor", func() Status {
		return DimmThermalSensorGet(sockInd, dimmAddr, &out)
	}); err != nil {
		return DimmThermal{}, err
	}
	return out, nil
}

func (c *Client) DimmSBRegRead(sockInd uint8, req DimmSBInfo) (DimmSBInfo, error) {
	out := req
	if err := c.call("DimmSBRegRead", func() Status { return DimmSBRegRead(sockInd, &out) }); err != nil {
		return DimmSBInfo{}, err
	}
	return out, nil
}

func (c *Client) DimmSBRegWrite(sockInd uint8, req DimmSBInfo) error {
	inout := req
	return c.call("DimmSBRegWrite", func() Status { return DimmSBRegWrite(sockInd, &inout) })
}

func (c *Client) SetXGMIWidth(sockInd uint8, min uint8, max uint8) error {
	return c.call("SetXGMIWidth", func() Status { return XGMIWidthSet(sockInd, min, max) })
}

func (c *Client) XGMIWidth(sockInd uint32) (uint8, uint8, error) {
	var min, max uint8
	if err := c.call("XGMIWidth", func() Status { return XGMIWidthGet(sockInd, &min, &max) }); err != nil {
		return 0, 0, err
	}
	return min, max, nil
}

func (c *Client) SetGMI3LinkWidthRange(sockInd uint8, minLinkWidth uint8, maxLinkWidth uint8) error {
	return c.call("SetGMI3LinkWidthRange", func() Status {
		return GMI3LinkWidthRangeSet(sockInd, minLinkWidth, maxLinkWidth)
	})
}

func (c *Client) APBEnable(sockInd uint32) error {
	return c.call("APBEnable", func() Status { return APBEnable(sockInd) })
}

func (c *Client) APBDisable(sockInd uint32, pstate uint8) error {
	return c.call("APBDisable", func() Status { return APBDisable(sockInd, pstate) })
}

func (c *Client) APBStatus(sockInd uint32) (uint8, uint8, error) {
	var apbDisabled, pstate uint8
	if err := c.call("APBStatus", func() Status { return APBStatusGet(sockInd, &apbDisabled, &pstate) }); err != nil {
		return 0, 0, err
	}
	return apbDisabled, pstate, nil
}

func (c *Client) SetSocketLclkDpmLevel(sockInd uint32, nbioID uint8, min uint8, max uint8) error {
	return c.call("SetSocketLclkDpmLevel", func() Status {
		return SocketLclkDpmLevelSet(sockInd, nbioID, min, max)
	})
}

func (c *Client) SocketLclkDpmLevel(sockInd uint8, nbioID uint8) (DPMLevel, error) {
	var out DPMLevel
	if err := c.call("SocketLclkDpmLevel", func() Status {
		return SocketLclkDpmLevelGet(sockInd, nbioID, &out)
	}); err != nil {
		return DPMLevel{}, err
	}
	return out, nil
}

func (c *Client) SetPCIeLinkRate(sockInd uint8, rateCtrl uint8) (uint8, error) {
	var prevMode uint8
	if err := c.call("SetPCIeLinkRate", func() Status { return PCIeLinkRateSet(sockInd, rateCtrl, &prevMode) }); err != nil {
		return 0, err
	}
	return prevMode, nil
}

func (c *Client) SetDFPstateRange(sockInd uint8, minPstate uint8, maxPstate uint8) error {
	return c.call("SetDFPstateRange", func() Status { return DFPstateRangeSet(sockInd, minPstate, maxPstate) })
}

func (c *Client) DFPstateRange(sockInd uint8) (uint8, uint8, error) {
	var minPstate, maxPstate uint8
	if err := c.call("DFPstateRange", func() Status { return DFPstateRangeGet(sockInd, &minPstate, &maxPstate) }); err != nil {
		return 0, 0, err
	}
	return minPstate, maxPstate, nil
}

func (c *Client) SetXGMIPstateRange(sockInd uint8, minState uint8, maxState uint8) error {
	return c.call("SetXGMIPstateRange", func() Status { return XGMIPstateRangeSet(sockInd, minState, maxState) })
}

func (c *Client) XGMIPstateRange(sockInd uint8) (uint8, uint8, error) {
	var minState, maxState uint8
	if err := c.call("XGMIPstateRange", func() Status { return XGMIPstateRangeGet(sockInd, &minState, &maxState) }); err != nil {
		return 0, 0, err
	}
	return minState, maxState, nil
}

func (c *Client) SetPC6Enable(sockInd uint8, pc6Enable uint8) error {
	return c.call("SetPC6Enable", func() Status { return PC6EnableSet(sockInd, pc6Enable) })
}

func (c *Client) PC6Enable(sockInd uint8) (uint8, error) {
	var out uint8
	if err := c.call("PC6Enable", func() Status { return PC6EnableGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) SetCC6Enable(sockInd uint8, cc6Enable uint8) error {
	return c.call("SetCC6Enable", func() Status { return CC6EnableSet(sockInd, cc6Enable) })
}

func (c *Client) CC6Enable(sockInd uint8) (uint8, error) {
	var out uint8
	if err := c.call("CC6Enable", func() Status { return CC6EnableGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) CurrentIOBandwidth(sockInd uint8, link LinkIDBWType) (uint32, error) {
	var out uint32
	if err := c.call("CurrentIOBandwidth", func() Status { return CurrentIOBandwidthGet(sockInd, link, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) CurrentXGMIBandwidth(sockInd uint8, link LinkIDBWType) (uint32, error) {
	var out uint32
	if err := c.call("CurrentXGMIBandwidth", func() Status { return CurrentXGMIBWGet(sockInd, link, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) MetricsTableVersion() (uint32, error) {
	var out uint32
	if err := c.call("MetricsTableVersion", func() Status { return MetricsTableVersionGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) MetricsTable(sockInd uint8) (MetricTable, error) {
	var out MetricTable
	if err := c.call("MetricsTable", func() Status { return MetricsTableGet(sockInd, &out) }); err != nil {
		return MetricTable{}, err
	}
	return out, nil
}

func (c *Client) DRAMAddressMetricsTable(sockInd uint8) (uint64, error) {
	var out uint64
	if err := c.call("DRAMAddressMetricsTable", func() Status { return DRAMAddressMetricsTableGet(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) TestHSMPMailbox(sockInd uint8, data uint32) (uint32, error) {
	out := data
	if err := c.call("TestHSMPMailbox", func() Status { return TestHSMPMailbox(sockInd, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) CPUFamily() (uint32, error) {
	var out uint32
	if err := c.call("CPUFamily", func() Status { return CPUFamilyGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) CPUModel() (uint32, error) {
	var out uint32
	if err := c.call("CPUModel", func() Status { return CPUModelGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) ThreadsPerCore() (uint32, error) {
	var out uint32
	if err := c.call("ThreadsPerCore", func() Status { return ThreadsPerCoreGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) NumberOfCPUs() (uint32, error) {
	var out uint32
	if err := c.call("NumberOfCPUs", func() Status { return NumberOfCPUsGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) NumberOfSockets() (uint32, error) {
	var out uint32
	if err := c.call("NumberOfSockets", func() Status { return NumberOfSocketsGet(&out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) FirstOnlineCoreOnSocket(socketIdx uint32) (uint32, error) {
	var out uint32
	if err := c.call("FirstOnlineCoreOnSocket", func() Status { return FirstOnlineCoreOnSocket(socketIdx, &out) }); err != nil {
		return 0, err
	}
	return out, nil
}

func (c *Client) EnabledCommands(sockInd uint8, readMask bool) (EnabledCommandsInfo, error) {
	info := EnabledCommandsInfo{ReadMask: readMask}
	if err := c.call("EnabledCommands", func() Status { return GetEnabledCommands(sockInd, &info) }); err != nil {
		return EnabledCommandsInfo{}, err
	}
	return info, nil
}
