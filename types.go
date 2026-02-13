package esmi

import "fmt"

const (
	EnergyDeviceName     = "amd_energy"
	HSMPCharDevfileName  = "/dev/hsmp"
	HSMPMetricTablePath  = "/sys/devices/platform/amd_hsmp"
	Core0                = 0
	Socket0              = 0
	CPUMaxCoresPerSocket = 512
)

var FrequencyLimitSourceNames = []string{
	"cHTC-Active",
	"PROCHOT",
	"TDC Limit (CPU rail)",
	"PPT Limit",
	"OPN Max",
	"Reliability Limit",
	"APML Agent",
	"HSMP Agent",
	"VRHOT(Voltage Regulator Hot)",
	"TDC Limit (VDD_MEM_S3 rail)",
}

type Status int32

const (
	StatusSuccess        Status = 0
	StatusInitialized    Status = StatusSuccess
	StatusNoEnergyDrv    Status = 1
	StatusNoMSRDrv       Status = 2
	StatusNoHSMPDrv      Status = 3
	StatusNoHSMPSup      Status = 4
	StatusNoDrv          Status = 5
	StatusFileNotFound   Status = 6
	StatusDevBusy        Status = 7
	StatusPermission     Status = 8
	StatusNotSupported   Status = 9
	StatusFileError      Status = 10
	StatusInterrupted    Status = 11
	StatusIOError        Status = 12
	StatusUnexpectedSize Status = 13
	StatusUnknownError   Status = 14
	StatusArgPtrNull     Status = 15
	StatusNoMemory       Status = 16
	StatusNotInitialized Status = 17
	StatusInvalidInput   Status = 18
	StatusHSMPTimeout    Status = 19
	StatusNoHSMPMsgSup   Status = 20
	StatusPreReqNotSat   Status = 21
	StatusSMUBusy        Status = 22
)

func (s Status) OK() bool {
	return s == StatusSuccess
}

func (s Status) String() string {
	switch s {
	case StatusSuccess:
		return "ESMI_SUCCESS"
	case StatusNoEnergyDrv:
		return "ESMI_NO_ENERGY_DRV"
	case StatusNoMSRDrv:
		return "ESMI_NO_MSR_DRV"
	case StatusNoHSMPDrv:
		return "ESMI_NO_HSMP_DRV"
	case StatusNoHSMPSup:
		return "ESMI_NO_HSMP_SUP"
	case StatusNoDrv:
		return "ESMI_NO_DRV"
	case StatusFileNotFound:
		return "ESMI_FILE_NOT_FOUND"
	case StatusDevBusy:
		return "ESMI_DEV_BUSY"
	case StatusPermission:
		return "ESMI_PERMISSION"
	case StatusNotSupported:
		return "ESMI_NOT_SUPPORTED"
	case StatusFileError:
		return "ESMI_FILE_ERROR"
	case StatusInterrupted:
		return "ESMI_INTERRUPTED"
	case StatusIOError:
		return "ESMI_IO_ERROR"
	case StatusUnexpectedSize:
		return "ESMI_UNEXPECTED_SIZE"
	case StatusUnknownError:
		return "ESMI_UNKNOWN_ERROR"
	case StatusArgPtrNull:
		return "ESMI_ARG_PTR_NULL"
	case StatusNoMemory:
		return "ESMI_NO_MEMORY"
	case StatusNotInitialized:
		return "ESMI_NOT_INITIALIZED"
	case StatusInvalidInput:
		return "ESMI_INVALID_INPUT"
	case StatusHSMPTimeout:
		return "ESMI_HSMP_TIMEOUT"
	case StatusNoHSMPMsgSup:
		return "ESMI_NO_HSMP_MSG_SUP"
	case StatusPreReqNotSat:
		return "ESMI_PRE_REQ_NOT_SAT"
	case StatusSMUBusy:
		return "ESMI_SMU_BUSY"
	default:
		return fmt.Sprintf("ESMI_STATUS_%d", int32(s))
	}
}

func (s Status) Error() string {
	return s.String()
}

type IoBWEncoding uint32

const (
	AggBW IoBWEncoding = 1 << iota
	RdBW
	WrBW
)

type SetGetFloorlimit uint32

const (
	SetFloorFrequencyCore SetGetFloorlimit = iota
	SetFloorFrequencySocket
	GetFloorFrequencyCore
	GetEffFloorFrequencyCore
	GetFloorFrequencySocket
	GetEffFloorFrequencySocket
)

type HSMPDriverVersion struct {
	Major uint32
	Minor uint32
}

type SMUFWVersion struct {
	Debug  uint8
	Minor  uint8
	Major  uint8
	Unused uint8
}

type DDRBWMetrics struct {
	MaxBW       uint32
	UtilizedBW  uint32
	UtilizedPct uint32
}

type TempRangeRefreshRate struct {
	Range   uint8
	RefRate uint8
}

type DimmPower struct {
	Power      uint16
	UpdateRate uint16
	DimmAddr   uint8
}

type DimmThermal struct {
	Sensor     uint16
	UpdateRate uint16
	DimmAddr   uint8
	Temp       float32
}

type DimmSBInfo struct {
	DimmAddr  uint8
	LID       uint8
	RegOffset uint16
	RegSpace  uint8
	WriteData uint8
	ReadData  uint32
}

type SVI3Info struct {
	RailSelection uint8
	RailIndex     uint8
	Temperature   uint32
}

type LinkIDBWType struct {
	BWType   IoBWEncoding
	LinkName string
}

type DPMLevel struct {
	MaxDPMLevel uint8
	MinDPMLevel uint8
}

type EnabledCommandsInfo struct {
	ReadMask bool
	Arg0     uint32
	Arg1     uint32
	Arg2     uint32
}

type MetricTable struct {
	AccumulationCounter uint32

	MaxSocketTemperature    uint32
	MaxVRTemperature        uint32
	MaxHBMTemperature       uint32
	MaxSocketTemperatureAcc uint64
	MaxVRTemperatureAcc     uint64
	MaxHBMTemperatureAcc    uint64

	SocketPowerLimit    uint32
	MaxSocketPowerLimit uint32
	SocketPower         uint32

	Timestamp       uint64
	SocketEnergyAcc uint64
	CCDEnergyAcc    uint64
	XCDEnergyAcc    uint64
	AIDEnergyAcc    uint64
	HBMEnergyAcc    uint64

	CCLKFrequencyLimit   uint32
	GFXCLKFrequencyLimit uint32
	FCLKFrequency        uint32
	UCLKFrequency        uint32
	SocclkFrequency      [4]uint32
	VclkFrequency        [4]uint32
	DclkFrequency        [4]uint32
	LclkFrequency        [4]uint32
	GFXCLKFrequencyAcc   [8]uint64
	CCLKFrequencyAcc     [96]uint64

	MaxCCLKFrequency     uint32
	MinCCLKFrequency     uint32
	MaxGFXCLKFrequency   uint32
	MinGFXCLKFrequency   uint32
	FCLKFrequencyTable   [4]uint32
	UCLKFrequencyTable   [4]uint32
	SocclkFrequencyTable [4]uint32
	VclkFrequencyTable   [4]uint32
	DclkFrequencyTable   [4]uint32
	LclkFrequencyTable   [4]uint32
	MaxLclkDPMRange      uint32
	MinLclkDPMRange      uint32

	XGMIWidth             uint32
	XGMIBitrate           uint32
	XGMIReadBandwidthAcc  [8]uint64
	XGMIWriteBandwidthAcc [8]uint64

	SocketC0Residency           uint32
	SocketGFXBusy               uint32
	DRAMBandwidthUtilization    uint32
	SocketC0ResidencyAcc        uint64
	SocketGFXBusyAcc            uint64
	DRAMBandwidthAcc            uint64
	MaxDRAMBandwidth            uint32
	DRAMBandwidthUtilizationAcc uint64
	PCIEBandwidthAcc            [4]uint64

	PROCHOTResidencyAcc   uint32
	PPTResidencyAcc       uint32
	SocketTHMResidencyAcc uint32
	VRTHMResidencyAcc     uint32
	HBMTHMResidencyAcc    uint32
	Spare                 uint32

	GFXCLKFrequency [8]uint32
}
