//go:build !linux || !cgo

package esmi

const noCGOMsg = "esmi bindings require linux+cgo"

func Init() Status {
	return StatusNotSupported
}

func Exit() {}

func CoreEnergyGet(coreInd uint32, energy *uint64) Status {
	return StatusNotSupported
}

func SocketEnergyGet(socketIdx uint32, energy *uint64) Status {
	return StatusNotSupported
}

func AllEnergiesGet(energy []uint64) Status {
	return StatusNotSupported
}

func RaplUnitsHSMPMailboxGet(sockInd uint32, tu *uint8, esu *uint8) Status {
	return StatusNotSupported
}

func RaplPackageCounterHSMPMailboxGet(sockInd uint32, counter1 *uint32, counter0 *uint32) Status {
	return StatusNotSupported
}

func RaplCoreCounterHSMPMailboxGet(coreInd uint32, counter1 *uint32, counter0 *uint32) Status {
	return StatusNotSupported
}

func CoreEnergyHSMPMailboxGet(coreInd uint32, energy *uint64) Status {
	return StatusNotSupported
}

func PackageEnergyHSMPMailboxGet(sockInd uint32, energy *uint64) Status {
	return StatusNotSupported
}

func HSMPDriverVersionGet(hsmpDriverVer *HSMPDriverVersion) Status {
	return StatusNotSupported
}

func SMUFWVersionGet(smuFW *SMUFWVersion) Status {
	return StatusNotSupported
}

func ProchotStatusGet(socketIdx uint32, prochot *uint32) Status {
	return StatusNotSupported
}

func FclkMclkGet(socketIdx uint32, fclk *uint32, mclk *uint32) Status {
	return StatusNotSupported
}

func CclkLimitGet(socketIdx uint32, cclk *uint32) Status {
	return StatusNotSupported
}

func HSMPProtoVerGet(protoVer *uint32) Status {
	return StatusNotSupported
}

func SocketCurrentActiveFreqLimitGet(sockInd uint32, freq *uint16, srcType *[]string) Status {
	return StatusNotSupported
}

func SocketFreqRangeGet(sockInd uint8, fmax *uint16, fmin *uint16) Status {
	return StatusNotSupported
}

func CurrentFreqLimitCoreGet(coreID uint32, freq *uint32) Status {
	return StatusNotSupported
}

func CpurailIsofreqPolicyGet(sockInd uint8, val *bool) Status {
	return StatusNotSupported
}

func DfcCtrlSettingGet(sockInd uint8, val *bool) Status {
	return StatusNotSupported
}

func SocketPowerGet(socketIdx uint32, power *uint32) Status {
	return StatusNotSupported
}

func SocketPowerCapGet(socketIdx uint32, cap *uint32) Status {
	return StatusNotSupported
}

func SocketPowerCapMaxGet(socketIdx uint32, max *uint32) Status {
	return StatusNotSupported
}

func PwrSVITelemetryAllRailsGet(sockInd uint32, power *uint32) Status {
	return StatusNotSupported
}

func PwrEfficiencyModeGet(sockInd uint8, mode *uint8, util *uint32, pptLimit *uint32) Status {
	return StatusNotSupported
}

func ReadCCDPower(coreID uint32, power *uint32) Status {
	return StatusNotSupported
}

func SocketPowerCapSet(socketIdx uint32, cap uint32) Status {
	return StatusNotSupported
}

func PwrEfficiencyModeSet(sockInd uint8, mode uint8, util *uint32, pptLimit *uint32) Status {
	return StatusNotSupported
}

func CpurailIsofreqPolicySet(sockInd uint8, val *bool) Status {
	return StatusNotSupported
}

func DfcEnableSet(sockInd uint8, val *bool) Status {
	return StatusNotSupported
}

func CoreBoostlimitGet(cpuInd uint32, boostlimit *uint32) Status {
	return StatusNotSupported
}

func FloorlimitSetGet(coreOrSockInd uint32, floorlimit *uint32, typ SetGetFloorlimit) Status {
	return StatusNotSupported
}

func MSRFloorlimitSet(coreOrSockInd uint32, msrFloorlimit uint32, typ SetGetFloorlimit, fmax uint32) Status {
	return StatusNotSupported
}

func SDPSLimitSet(sockInd uint8, sdpsLimit *uint32) Status {
	return StatusNotSupported
}

func SDPSLimitGet(sockInd uint8, currentSDPSLimit *uint32) Status {
	return StatusNotSupported
}

func SocketC0ResidencyGet(socketIdx uint32, c0Residency *uint32) Status {
	return StatusNotSupported
}

func CoreBoostlimitSet(cpuInd uint32, boostlimit uint32) Status {
	return StatusNotSupported
}

func SocketBoostlimitSet(socketIdx uint32, boostlimit uint32) Status {
	return StatusNotSupported
}

func DDRBWGet(sockInd uint8, ddrBW *DDRBWMetrics) Status {
	return StatusNotSupported
}

func SocketTemperatureGet(sockInd uint32, tmon *uint32) Status {
	return StatusNotSupported
}

func ReadTdelta(sockInd uint8, statusVal *uint8) Status {
	return StatusNotSupported
}

func GetSVI3VRControllerTemp(sockInd uint8, inout *SVI3Info) Status {
	return StatusNotSupported
}

func DimmTempRangeAndRefreshRateGet(sockInd uint8, dimmAddr uint8, rate *TempRangeRefreshRate) Status {
	return StatusNotSupported
}

func DimmPowerConsumptionGet(sockInd uint8, dimmAddr uint8, dimmPower *DimmPower) Status {
	return StatusNotSupported
}

func DimmThermalSensorGet(sockInd uint8, dimmAddr uint8, dimmThermal *DimmThermal) Status {
	return StatusNotSupported
}

func DimmSBRegRead(sockInd uint8, inout *DimmSBInfo) Status {
	return StatusNotSupported
}

func DimmSBRegWrite(sockInd uint8, inout *DimmSBInfo) Status {
	return StatusNotSupported
}

func XGMIWidthSet(sockInd uint8, min uint8, max uint8) Status {
	return StatusNotSupported
}

func XGMIWidthGet(sockInd uint32, min *uint8, max *uint8) Status {
	return StatusNotSupported
}

func GMI3LinkWidthRangeSet(sockInd uint8, minLinkWidth uint8, maxLinkWidth uint8) Status {
	return StatusNotSupported
}

func APBEnable(sockInd uint32) Status {
	return StatusNotSupported
}

func APBDisable(sockInd uint32, pstate uint8) Status {
	return StatusNotSupported
}

func APBStatusGet(sockInd uint32, apbDisabled *uint8, pstate *uint8) Status {
	return StatusNotSupported
}

func SocketLclkDpmLevelSet(sockInd uint32, nbioID uint8, min uint8, max uint8) Status {
	return StatusNotSupported
}

func SocketLclkDpmLevelGet(sockInd uint8, nbioID uint8, nbio *DPMLevel) Status {
	return StatusNotSupported
}

func PCIeLinkRateSet(sockInd uint8, rateCtrl uint8, prevMode *uint8) Status {
	return StatusNotSupported
}

func DFPstateRangeSet(sockInd uint8, minPstate uint8, maxPstate uint8) Status {
	return StatusNotSupported
}

func DFPstateRangeGet(sockInd uint8, minPstate *uint8, maxPstate *uint8) Status {
	return StatusNotSupported
}

func XGMIPstateRangeSet(sockInd uint8, minState uint8, maxState uint8) Status {
	return StatusNotSupported
}

func XGMIPstateRangeGet(sockInd uint8, minState *uint8, maxState *uint8) Status {
	return StatusNotSupported
}

func PC6EnableSet(sockInd uint8, pc6Enable uint8) Status {
	return StatusNotSupported
}

func PC6EnableGet(sockInd uint8, currentPC6Enable *uint8) Status {
	return StatusNotSupported
}

func CC6EnableSet(sockInd uint8, cc6Enable uint8) Status {
	return StatusNotSupported
}

func CC6EnableGet(sockInd uint8, currentCC6Enable *uint8) Status {
	return StatusNotSupported
}

func CurrentIOBandwidthGet(sockInd uint8, link LinkIDBWType, ioBW *uint32) Status {
	return StatusNotSupported
}

func CurrentXGMIBWGet(sockInd uint8, link LinkIDBWType, xgmiBW *uint32) Status {
	return StatusNotSupported
}

func MetricsTableVersionGet(metricsVersion *uint32) Status {
	return StatusNotSupported
}

func MetricsTableGet(sockInd uint8, metricsTable *MetricTable) Status {
	return StatusNotSupported
}

func DRAMAddressMetricsTableGet(sockInd uint8, dramAddr *uint64) Status {
	return StatusNotSupported
}

func TestHSMPMailbox(sockInd uint8, data *uint32) Status {
	return StatusNotSupported
}

func CPUFamilyGet(family *uint32) Status {
	return StatusNotSupported
}

func CPUModelGet(model *uint32) Status {
	return StatusNotSupported
}

func ThreadsPerCoreGet(threads *uint32) Status {
	return StatusNotSupported
}

func NumberOfCPUsGet(cpus *uint32) Status {
	return StatusNotSupported
}

func NumberOfSocketsGet(sockets *uint32) Status {
	return StatusNotSupported
}

func FirstOnlineCoreOnSocket(socketIdx uint32, coreInd *uint32) Status {
	return StatusNotSupported
}

func GetEnabledCommands(sockInd uint8, info *EnabledCommandsInfo) Status {
	return StatusNotSupported
}

func ErrMsg(esmiErr Status) string {
	return noCGOMsg
}
