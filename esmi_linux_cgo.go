//go:build linux && cgo

package esmi

/*
#cgo CFLAGS: -I${SRCDIR}/cinclude
#cgo LDFLAGS: -lpthread -lrt -lm
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <e_smi/e_smi.h>

struct go_hsmp_metric_table {
	uint32_t accumulation_counter;

	uint32_t max_socket_temperature;
	uint32_t max_vr_temperature;
	uint32_t max_hbm_temperature;
	uint64_t max_socket_temperature_acc;
	uint64_t max_vr_temperature_acc;
	uint64_t max_hbm_temperature_acc;

	uint32_t socket_power_limit;
	uint32_t max_socket_power_limit;
	uint32_t socket_power;

	uint64_t timestamp;
	uint64_t socket_energy_acc;
	uint64_t ccd_energy_acc;
	uint64_t xcd_energy_acc;
	uint64_t aid_energy_acc;
	uint64_t hbm_energy_acc;

	uint32_t cclk_frequency_limit;
	uint32_t gfxclk_frequency_limit;
	uint32_t fclk_frequency;
	uint32_t uclk_frequency;
	uint32_t socclk_frequency[4];
	uint32_t vclk_frequency[4];
	uint32_t dclk_frequency[4];
	uint32_t lclk_frequency[4];
	uint64_t gfxclk_frequency_acc[8];
	uint64_t cclk_frequency_acc[96];

	uint32_t max_cclk_frequency;
	uint32_t min_cclk_frequency;
	uint32_t max_gfxclk_frequency;
	uint32_t min_gfxclk_frequency;
	uint32_t fclk_frequency_table[4];
	uint32_t uclk_frequency_table[4];
	uint32_t socclk_frequency_table[4];
	uint32_t vclk_frequency_table[4];
	uint32_t dclk_frequency_table[4];
	uint32_t lclk_frequency_table[4];
	uint32_t max_lclk_dpm_range;
	uint32_t min_lclk_dpm_range;

	uint32_t xgmi_width;
	uint32_t xgmi_bitrate;
	uint64_t xgmi_read_bandwidth_acc[8];
	uint64_t xgmi_write_bandwidth_acc[8];

	uint32_t socket_c0_residency;
	uint32_t socket_gfx_busy;
	uint32_t dram_bandwidth_utilization;
	uint64_t socket_c0_residency_acc;
	uint64_t socket_gfx_busy_acc;
	uint64_t dram_bandwidth_acc;
	uint32_t max_dram_bandwidth;
	uint64_t dram_bandwidth_utilization_acc;
	uint64_t pcie_bandwidth_acc[4];

	uint32_t prochot_residency_acc;
	uint32_t ppt_residency_acc;
	uint32_t socket_thm_residency_acc;
	uint32_t vr_thm_residency_acc;
	uint32_t hbm_thm_residency_acc;
	uint32_t spare;

	uint32_t gfxclk_frequency[8];
};

static void go_copy_metric_table(struct hsmp_metric_table *in, struct go_hsmp_metric_table *out)
{
	int i = 0;

	out->accumulation_counter = in->accumulation_counter;

	out->max_socket_temperature = in->max_socket_temperature;
	out->max_vr_temperature = in->max_vr_temperature;
	out->max_hbm_temperature = in->max_hbm_temperature;
	out->max_socket_temperature_acc = in->max_socket_temperature_acc;
	out->max_vr_temperature_acc = in->max_vr_temperature_acc;
	out->max_hbm_temperature_acc = in->max_hbm_temperature_acc;

	out->socket_power_limit = in->socket_power_limit;
	out->max_socket_power_limit = in->max_socket_power_limit;
	out->socket_power = in->socket_power;

	out->timestamp = in->timestamp;
	out->socket_energy_acc = in->socket_energy_acc;
	out->ccd_energy_acc = in->ccd_energy_acc;
	out->xcd_energy_acc = in->xcd_energy_acc;
	out->aid_energy_acc = in->aid_energy_acc;
	out->hbm_energy_acc = in->hbm_energy_acc;

	out->cclk_frequency_limit = in->cclk_frequency_limit;
	out->gfxclk_frequency_limit = in->gfxclk_frequency_limit;
	out->fclk_frequency = in->fclk_frequency;
	out->uclk_frequency = in->uclk_frequency;
	for (i = 0; i < 4; i++) {
		out->socclk_frequency[i] = in->socclk_frequency[i];
		out->vclk_frequency[i] = in->vclk_frequency[i];
		out->dclk_frequency[i] = in->dclk_frequency[i];
		out->lclk_frequency[i] = in->lclk_frequency[i];

		out->fclk_frequency_table[i] = in->fclk_frequency_table[i];
		out->uclk_frequency_table[i] = in->uclk_frequency_table[i];
		out->socclk_frequency_table[i] = in->socclk_frequency_table[i];
		out->vclk_frequency_table[i] = in->vclk_frequency_table[i];
		out->dclk_frequency_table[i] = in->dclk_frequency_table[i];
		out->lclk_frequency_table[i] = in->lclk_frequency_table[i];

		out->pcie_bandwidth_acc[i] = in->pcie_bandwidth_acc[i];
	}
	for (i = 0; i < 8; i++) {
		out->gfxclk_frequency_acc[i] = in->gfxclk_frequency_acc[i];
		out->xgmi_read_bandwidth_acc[i] = in->xgmi_read_bandwidth_acc[i];
		out->xgmi_write_bandwidth_acc[i] = in->xgmi_write_bandwidth_acc[i];
		out->gfxclk_frequency[i] = in->gfxclk_frequency[i];
	}
	for (i = 0; i < 96; i++) {
		out->cclk_frequency_acc[i] = in->cclk_frequency_acc[i];
	}

	out->max_cclk_frequency = in->max_cclk_frequency;
	out->min_cclk_frequency = in->min_cclk_frequency;
	out->max_gfxclk_frequency = in->max_gfxclk_frequency;
	out->min_gfxclk_frequency = in->min_gfxclk_frequency;
	out->max_lclk_dpm_range = in->max_lclk_dpm_range;
	out->min_lclk_dpm_range = in->min_lclk_dpm_range;

	out->xgmi_width = in->xgmi_width;
	out->xgmi_bitrate = in->xgmi_bitrate;

	out->socket_c0_residency = in->socket_c0_residency;
	out->socket_gfx_busy = in->socket_gfx_busy;
	out->dram_bandwidth_utilization = in->dram_bandwidth_utilization;
	out->socket_c0_residency_acc = in->socket_c0_residency_acc;
	out->socket_gfx_busy_acc = in->socket_gfx_busy_acc;
	out->dram_bandwidth_acc = in->dram_bandwidth_acc;
	out->max_dram_bandwidth = in->max_dram_bandwidth;
	out->dram_bandwidth_utilization_acc = in->dram_bandwidth_utilization_acc;

	out->prochot_residency_acc = in->prochot_residency_acc;
	out->ppt_residency_acc = in->ppt_residency_acc;
	out->socket_thm_residency_acc = in->socket_thm_residency_acc;
	out->vr_thm_residency_acc = in->vr_thm_residency_acc;
	out->hbm_thm_residency_acc = in->hbm_thm_residency_acc;
	out->spare = in->spare;
}

static esmi_status_t go_esmi_metrics_table_get(uint8_t sock_ind, struct go_hsmp_metric_table *out)
{
	struct hsmp_metric_table in = {0};
	esmi_status_t st;

	st = esmi_metrics_table_get(sock_ind, &in);
	if (st == ESMI_SUCCESS && out != NULL) {
		go_copy_metric_table(&in, out);
	}
	return st;
}

static esmi_status_t go_esmi_dimm_temp_range_and_refresh_rate_get(
	uint8_t sock_ind,
	uint8_t dimm_addr,
	uint8_t *range,
	uint8_t *ref_rate)
{
	struct temp_range_refresh_rate rate = {0};
	esmi_status_t st = esmi_dimm_temp_range_and_refresh_rate_get(sock_ind, dimm_addr, &rate);
	if (st == ESMI_SUCCESS) {
		if (range) {
			*range = rate.range;
		}
		if (ref_rate) {
			*ref_rate = rate.ref_rate;
		}
	}
	return st;
}

static esmi_status_t go_esmi_dimm_power_consumption_get(
	uint8_t sock_ind,
	uint8_t dimm_addr,
	uint16_t *power,
	uint16_t *update_rate,
	uint8_t *out_dimm_addr)
{
	struct dimm_power dimm_pow = {0};
	esmi_status_t st = esmi_dimm_power_consumption_get(sock_ind, dimm_addr, &dimm_pow);
	if (st == ESMI_SUCCESS) {
		if (power) {
			*power = dimm_pow.power;
		}
		if (update_rate) {
			*update_rate = dimm_pow.update_rate;
		}
		if (out_dimm_addr) {
			*out_dimm_addr = dimm_pow.dimm_addr;
		}
	}
	return st;
}

static esmi_status_t go_esmi_dimm_thermal_sensor_get(
	uint8_t sock_ind,
	uint8_t dimm_addr,
	uint16_t *sensor,
	uint16_t *update_rate,
	uint8_t *out_dimm_addr,
	float *temp)
{
	struct dimm_thermal dimm_temp = {0};
	esmi_status_t st = esmi_dimm_thermal_sensor_get(sock_ind, dimm_addr, &dimm_temp);
	if (st == ESMI_SUCCESS) {
		if (sensor) {
			*sensor = dimm_temp.sensor;
		}
		if (update_rate) {
			*update_rate = dimm_temp.update_rate;
		}
		if (out_dimm_addr) {
			*out_dimm_addr = dimm_temp.dimm_addr;
		}
		if (temp) {
			*temp = dimm_temp.temp;
		}
	}
	return st;
}

static esmi_status_t go_esmi_get_svi3_vr_controller_temp(
	uint8_t sock_ind,
	uint8_t rail_selection,
	uint8_t rail_index,
	uint8_t *out_rail_index,
	uint32_t *temperature)
{
	struct svi3_info info = {0};
	info.m_svi3_info_inarg.info.svi3_rail_selection = rail_selection;
	info.m_svi3_info_inarg.info.svi3_rail_index = rail_index;
	esmi_status_t st = esmi_get_svi3_vr_controller_temp(sock_ind, &info);
	if (st == ESMI_SUCCESS) {
		if (out_rail_index) {
			*out_rail_index = info.m_svi3_info_inarg.info.svi3_rail_index;
		}
		if (temperature) {
			*temperature = info.m_svi3_info_inarg.info.svi3_temperature;
		}
	}
	return st;
}

static esmi_status_t go_esmi_dimm_sb_reg_read(
	uint8_t sock_ind,
	uint8_t dimm_addr,
	uint8_t lid,
	uint16_t reg_offset,
	uint8_t reg_space,
	uint8_t write_data,
	uint32_t *read_data)
{
	struct dimm_sb_info info = {0};
	info.m_dimm_sb_info_inarg.info.dimm_addr = dimm_addr;
	info.m_dimm_sb_info_inarg.info.lid = lid;
	info.m_dimm_sb_info_inarg.info.reg_offset = reg_offset;
	info.m_dimm_sb_info_inarg.info.reg_space = reg_space;
	info.m_dimm_sb_info_inarg.info.write_data = write_data;
	esmi_status_t st = esmi_dimm_sb_reg_read(sock_ind, &info);
	if (st == ESMI_SUCCESS && read_data) {
		*read_data = info.read_data;
	}
	return st;
}

static esmi_status_t go_esmi_dimm_sb_reg_write(
	uint8_t sock_ind,
	uint8_t dimm_addr,
	uint8_t lid,
	uint16_t reg_offset,
	uint8_t reg_space,
	uint8_t write_data)
{
	struct dimm_sb_info info = {0};
	info.m_dimm_sb_info_inarg.info.dimm_addr = dimm_addr;
	info.m_dimm_sb_info_inarg.info.lid = lid;
	info.m_dimm_sb_info_inarg.info.reg_offset = reg_offset;
	info.m_dimm_sb_info_inarg.info.reg_space = reg_space;
	info.m_dimm_sb_info_inarg.info.write_data = write_data;
	return esmi_dimm_sb_reg_write(sock_ind, &info);
}

static esmi_status_t go_esmi_socket_lclk_dpm_level_get(
	uint8_t sock_ind,
	uint8_t nbio_id,
	uint8_t *min,
	uint8_t *max)
{
	struct dpm_level dpm = {0};
	esmi_status_t st = esmi_socket_lclk_dpm_level_get(sock_ind, nbio_id, &dpm);
	if (st == ESMI_SUCCESS) {
		if (min) {
			*min = dpm.min_dpm_level;
		}
		if (max) {
			*max = dpm.max_dpm_level;
		}
	}
	return st;
}

static esmi_status_t go_esmi_current_io_bandwidth_get(
	uint8_t sock_ind,
	uint32_t bw_type,
	char *link_name,
	uint32_t *io_bw)
{
	struct link_id_bw_type link = {0};
	link.bw_type = (io_bw_encoding)bw_type;
	link.link_name = link_name;
	return esmi_current_io_bandwidth_get(sock_ind, link, io_bw);
}

static esmi_status_t go_esmi_current_xgmi_bw_get(
	uint8_t sock_ind,
	uint32_t bw_type,
	char *link_name,
	uint32_t *xgmi_bw)
{
	struct link_id_bw_type link = {0};
	link.bw_type = (io_bw_encoding)bw_type;
	link.link_name = link_name;
	return esmi_current_xgmi_bw_get(sock_ind, link, xgmi_bw);
}

static esmi_status_t go_esmi_get_enabled_commands(
	uint8_t sock_ind,
	bool read_mask,
	uint32_t *arg0,
	uint32_t *arg1,
	uint32_t *arg2)
{
	struct hsmp_enabled_commands_info info = {0};
	info.read_mask = read_mask;
	esmi_status_t st = esmi_get_enabled_commands(sock_ind, &info);
	if (st == ESMI_SUCCESS) {
		if (arg0) {
			*arg0 = info.arg0;
		}
		if (arg1) {
			*arg1 = info.arg1;
		}
		if (arg2) {
			*arg2 = info.arg2;
		}
	}
	return st;
}
*/
import "C"

import (
	"unsafe"
)

const freqLimitSourceCount = 10

func cU8Ptr(p *uint8) *C.uint8_t {
	if p == nil {
		return nil
	}
	return (*C.uint8_t)(unsafe.Pointer(p))
}

func cU16Ptr(p *uint16) *C.uint16_t {
	if p == nil {
		return nil
	}
	return (*C.uint16_t)(unsafe.Pointer(p))
}

func cU32Ptr(p *uint32) *C.uint32_t {
	if p == nil {
		return nil
	}
	return (*C.uint32_t)(unsafe.Pointer(p))
}

func cU64Ptr(p *uint64) *C.uint64_t {
	if p == nil {
		return nil
	}
	return (*C.uint64_t)(unsafe.Pointer(p))
}

func cBoolPtr(p *bool, local *C.bool) *C.bool {
	if p == nil {
		return nil
	}
	if *p {
		*local = C.bool(true)
	} else {
		*local = C.bool(false)
	}
	return local
}

func updateBool(p *bool, v C.bool) {
	if p != nil {
		*p = v != C.bool(false)
	}
}

func status(v C.esmi_status_t) Status {
	return Status(v)
}

func Init() Status {
	return status(C.esmi_init())
}

func Exit() {
	C.esmi_exit()
}

func CoreEnergyGet(coreInd uint32, energy *uint64) Status {
	return status(C.esmi_core_energy_get(C.uint32_t(coreInd), cU64Ptr(energy)))
}

func SocketEnergyGet(socketIdx uint32, energy *uint64) Status {
	return status(C.esmi_socket_energy_get(C.uint32_t(socketIdx), cU64Ptr(energy)))
}

func AllEnergiesGet(energy []uint64) Status {
	var ptr *C.uint64_t
	if len(energy) > 0 {
		ptr = (*C.uint64_t)(unsafe.Pointer(&energy[0]))
	}
	return status(C.esmi_all_energies_get(ptr))
}

func RaplUnitsHSMPMailboxGet(sockInd uint32, tu *uint8, esu *uint8) Status {
	return status(C.esmi_rapl_units_hsmp_mailbox_get(C.uint32_t(sockInd), cU8Ptr(tu), cU8Ptr(esu)))
}

func RaplPackageCounterHSMPMailboxGet(sockInd uint32, counter1 *uint32, counter0 *uint32) Status {
	return status(C.esmi_rapl_package_counter_hsmp_mailbox_get(C.uint32_t(sockInd), cU32Ptr(counter1), cU32Ptr(counter0)))
}

func RaplCoreCounterHSMPMailboxGet(coreInd uint32, counter1 *uint32, counter0 *uint32) Status {
	return status(C.esmi_rapl_core_counter_hsmp_mailbox_get(C.uint32_t(coreInd), cU32Ptr(counter1), cU32Ptr(counter0)))
}

func CoreEnergyHSMPMailboxGet(coreInd uint32, energy *uint64) Status {
	return status(C.esmi_core_energy_hsmp_mailbox_get(C.uint32_t(coreInd), cU64Ptr(energy)))
}

func PackageEnergyHSMPMailboxGet(sockInd uint32, energy *uint64) Status {
	return status(C.esmi_package_energy_hsmp_mailbox_get(C.uint32_t(sockInd), cU64Ptr(energy)))
}

func HSMPDriverVersionGet(hsmpDriverVer *HSMPDriverVersion) Status {
	if hsmpDriverVer == nil {
		return status(C.esmi_hsmp_driver_version_get(nil))
	}
	var cVer C.struct_hsmp_driver_version
	st := status(C.esmi_hsmp_driver_version_get(&cVer))
	if st == StatusSuccess {
		hsmpDriverVer.Major = uint32(cVer.major)
		hsmpDriverVer.Minor = uint32(cVer.minor)
	}
	return st
}

func SMUFWVersionGet(smuFW *SMUFWVersion) Status {
	if smuFW == nil {
		return status(C.esmi_smu_fw_version_get(nil))
	}
	var cVer C.struct_smu_fw_version
	st := status(C.esmi_smu_fw_version_get(&cVer))
	if st == StatusSuccess {
		smuFW.Debug = uint8(cVer.debug)
		smuFW.Minor = uint8(cVer.minor)
		smuFW.Major = uint8(cVer.major)
		smuFW.Unused = uint8(cVer.unused)
	}
	return st
}

func ProchotStatusGet(socketIdx uint32, prochot *uint32) Status {
	return status(C.esmi_prochot_status_get(C.uint32_t(socketIdx), cU32Ptr(prochot)))
}

func FclkMclkGet(socketIdx uint32, fclk *uint32, mclk *uint32) Status {
	return status(C.esmi_fclk_mclk_get(C.uint32_t(socketIdx), cU32Ptr(fclk), cU32Ptr(mclk)))
}

func CclkLimitGet(socketIdx uint32, cclk *uint32) Status {
	return status(C.esmi_cclk_limit_get(C.uint32_t(socketIdx), cU32Ptr(cclk)))
}

func HSMPProtoVerGet(protoVer *uint32) Status {
	return status(C.esmi_hsmp_proto_ver_get(cU32Ptr(protoVer)))
}

func SocketCurrentActiveFreqLimitGet(sockInd uint32, freq *uint16, srcType *[]string) Status {
	var cFreq C.uint16_t
	var freqPtr *C.uint16_t
	if freq != nil {
		freqPtr = &cFreq
	}

	var srcPtrs [freqLimitSourceCount]*C.char
	var srcArg **C.char
	if srcType != nil {
		srcArg = (**C.char)(unsafe.Pointer(&srcPtrs[0]))
	}

	st := status(C.esmi_socket_current_active_freq_limit_get(C.uint32_t(sockInd), freqPtr, srcArg))
	if freq != nil {
		*freq = uint16(cFreq)
	}
	if srcType != nil {
		names := make([]string, 0, len(srcPtrs))
		for i := range srcPtrs {
			if srcPtrs[i] == nil {
				break
			}
			names = append(names, C.GoString(srcPtrs[i]))
		}
		*srcType = names
	}
	return st
}

func SocketFreqRangeGet(sockInd uint8, fmax *uint16, fmin *uint16) Status {
	return status(C.esmi_socket_freq_range_get(C.uint8_t(sockInd), cU16Ptr(fmax), cU16Ptr(fmin)))
}

func CurrentFreqLimitCoreGet(coreID uint32, freq *uint32) Status {
	return status(C.esmi_current_freq_limit_core_get(C.uint32_t(coreID), cU32Ptr(freq)))
}

func CpurailIsofreqPolicyGet(sockInd uint8, val *bool) Status {
	var cVal C.bool
	ptr := cBoolPtr(val, &cVal)
	st := status(C.esmi_cpurail_isofreq_policy_get(C.uint8_t(sockInd), ptr))
	if ptr != nil {
		updateBool(val, cVal)
	}
	return st
}

func DfcCtrlSettingGet(sockInd uint8, val *bool) Status {
	var cVal C.bool
	ptr := cBoolPtr(val, &cVal)
	st := status(C.esmi_dfc_ctrl_setting_get(C.uint8_t(sockInd), ptr))
	if ptr != nil {
		updateBool(val, cVal)
	}
	return st
}

func SocketPowerGet(socketIdx uint32, power *uint32) Status {
	return status(C.esmi_socket_power_get(C.uint32_t(socketIdx), cU32Ptr(power)))
}

func SocketPowerCapGet(socketIdx uint32, cap *uint32) Status {
	return status(C.esmi_socket_power_cap_get(C.uint32_t(socketIdx), cU32Ptr(cap)))
}

func SocketPowerCapMaxGet(socketIdx uint32, max *uint32) Status {
	return status(C.esmi_socket_power_cap_max_get(C.uint32_t(socketIdx), cU32Ptr(max)))
}

func PwrSVITelemetryAllRailsGet(sockInd uint32, power *uint32) Status {
	return status(C.esmi_pwr_svi_telemetry_all_rails_get(C.uint32_t(sockInd), cU32Ptr(power)))
}

func PwrEfficiencyModeGet(sockInd uint8, mode *uint8, util *uint32, pptLimit *uint32) Status {
	return status(C.esmi_pwr_efficiency_mode_get(C.uint8_t(sockInd), cU8Ptr(mode), cU32Ptr(util), cU32Ptr(pptLimit)))
}

func ReadCCDPower(coreID uint32, power *uint32) Status {
	return status(C.esmi_read_ccd_power(C.uint32_t(coreID), cU32Ptr(power)))
}

func SocketPowerCapSet(socketIdx uint32, cap uint32) Status {
	return status(C.esmi_socket_power_cap_set(C.uint32_t(socketIdx), C.uint32_t(cap)))
}

func PwrEfficiencyModeSet(sockInd uint8, mode uint8, util *uint32, pptLimit *uint32) Status {
	return status(C.esmi_pwr_efficiency_mode_set(C.uint8_t(sockInd), C.uint8_t(mode), cU32Ptr(util), cU32Ptr(pptLimit)))
}

func CpurailIsofreqPolicySet(sockInd uint8, val *bool) Status {
	var cVal C.bool
	ptr := cBoolPtr(val, &cVal)
	st := status(C.esmi_cpurail_isofreq_policy_set(C.uint8_t(sockInd), ptr))
	if ptr != nil {
		updateBool(val, cVal)
	}
	return st
}

func DfcEnableSet(sockInd uint8, val *bool) Status {
	var cVal C.bool
	ptr := cBoolPtr(val, &cVal)
	st := status(C.esmi_dfc_enable_set(C.uint8_t(sockInd), ptr))
	if ptr != nil {
		updateBool(val, cVal)
	}
	return st
}

func CoreBoostlimitGet(cpuInd uint32, boostlimit *uint32) Status {
	return status(C.esmi_core_boostlimit_get(C.uint32_t(cpuInd), cU32Ptr(boostlimit)))
}

func FloorlimitSetGet(coreOrSockInd uint32, floorlimit *uint32, typ SetGetFloorlimit) Status {
	return status(C.esmi_floorlimit_set_get(C.uint32_t(coreOrSockInd), cU32Ptr(floorlimit), C.enum_set_get_floorlimit(typ)))
}

func MSRFloorlimitSet(coreOrSockInd uint32, msrFloorlimit uint32, typ SetGetFloorlimit, fmax uint32) Status {
	return status(C.esmi_msr_floorlimit_set(C.uint32_t(coreOrSockInd), C.uint32_t(msrFloorlimit), C.enum_set_get_floorlimit(typ), C.uint32_t(fmax)))
}

func SDPSLimitSet(sockInd uint8, sdpsLimit *uint32) Status {
	return status(C.esmi_sdps_limit_set(C.uint8_t(sockInd), cU32Ptr(sdpsLimit)))
}

func SDPSLimitGet(sockInd uint8, currentSDPSLimit *uint32) Status {
	return status(C.esmi_sdps_limit_get(C.uint8_t(sockInd), cU32Ptr(currentSDPSLimit)))
}

func SocketC0ResidencyGet(socketIdx uint32, c0Residency *uint32) Status {
	return status(C.esmi_socket_c0_residency_get(C.uint32_t(socketIdx), cU32Ptr(c0Residency)))
}

func CoreBoostlimitSet(cpuInd uint32, boostlimit uint32) Status {
	return status(C.esmi_core_boostlimit_set(C.uint32_t(cpuInd), C.uint32_t(boostlimit)))
}

func SocketBoostlimitSet(socketIdx uint32, boostlimit uint32) Status {
	return status(C.esmi_socket_boostlimit_set(C.uint32_t(socketIdx), C.uint32_t(boostlimit)))
}

func DDRBWGet(sockInd uint8, ddrBW *DDRBWMetrics) Status {
	if ddrBW == nil {
		return status(C.esmi_ddr_bw_get(C.uint8_t(sockInd), nil))
	}
	var cDDR C.struct_ddr_bw_metrics
	st := status(C.esmi_ddr_bw_get(C.uint8_t(sockInd), &cDDR))
	if st == StatusSuccess {
		ddrBW.MaxBW = uint32(cDDR.max_bw)
		ddrBW.UtilizedBW = uint32(cDDR.utilized_bw)
		ddrBW.UtilizedPct = uint32(cDDR.utilized_pct)
	}
	return st
}

func SocketTemperatureGet(sockInd uint32, tmon *uint32) Status {
	return status(C.esmi_socket_temperature_get(C.uint32_t(sockInd), cU32Ptr(tmon)))
}

func ReadTdelta(sockInd uint8, statusVal *uint8) Status {
	return status(C.esmi_read_tdelta(C.uint8_t(sockInd), cU8Ptr(statusVal)))
}

func GetSVI3VRControllerTemp(sockInd uint8, inout *SVI3Info) Status {
	if inout == nil {
		return StatusArgPtrNull
	}
	var outRailIndex C.uint8_t
	var temp C.uint32_t
	st := status(C.go_esmi_get_svi3_vr_controller_temp(
		C.uint8_t(sockInd),
		C.uint8_t(inout.RailSelection),
		C.uint8_t(inout.RailIndex),
		&outRailIndex,
		&temp,
	))
	if st == StatusSuccess {
		inout.RailIndex = uint8(outRailIndex)
		inout.Temperature = uint32(temp)
	}
	return st
}

func DimmTempRangeAndRefreshRateGet(sockInd uint8, dimmAddr uint8, rate *TempRangeRefreshRate) Status {
	if rate == nil {
		return StatusArgPtrNull
	}
	var r C.uint8_t
	var ref C.uint8_t
	st := status(C.go_esmi_dimm_temp_range_and_refresh_rate_get(C.uint8_t(sockInd), C.uint8_t(dimmAddr), &r, &ref))
	if st == StatusSuccess {
		rate.Range = uint8(r)
		rate.RefRate = uint8(ref)
	}
	return st
}

func DimmPowerConsumptionGet(sockInd uint8, dimmAddr uint8, dimmPower *DimmPower) Status {
	if dimmPower == nil {
		return StatusArgPtrNull
	}
	var power C.uint16_t
	var updateRate C.uint16_t
	var outAddr C.uint8_t
	st := status(C.go_esmi_dimm_power_consumption_get(C.uint8_t(sockInd), C.uint8_t(dimmAddr), &power, &updateRate, &outAddr))
	if st == StatusSuccess {
		dimmPower.Power = uint16(power)
		dimmPower.UpdateRate = uint16(updateRate)
		dimmPower.DimmAddr = uint8(outAddr)
	}
	return st
}

func DimmThermalSensorGet(sockInd uint8, dimmAddr uint8, dimmThermal *DimmThermal) Status {
	if dimmThermal == nil {
		return StatusArgPtrNull
	}
	var sensor C.uint16_t
	var updateRate C.uint16_t
	var outAddr C.uint8_t
	var temp C.float
	st := status(C.go_esmi_dimm_thermal_sensor_get(C.uint8_t(sockInd), C.uint8_t(dimmAddr), &sensor, &updateRate, &outAddr, &temp))
	if st == StatusSuccess {
		dimmThermal.Sensor = uint16(sensor)
		dimmThermal.UpdateRate = uint16(updateRate)
		dimmThermal.DimmAddr = uint8(outAddr)
		dimmThermal.Temp = float32(temp)
	}
	return st
}

func DimmSBRegRead(sockInd uint8, inout *DimmSBInfo) Status {
	if inout == nil {
		return StatusArgPtrNull
	}
	var readData C.uint32_t
	st := status(C.go_esmi_dimm_sb_reg_read(
		C.uint8_t(sockInd),
		C.uint8_t(inout.DimmAddr),
		C.uint8_t(inout.LID),
		C.uint16_t(inout.RegOffset),
		C.uint8_t(inout.RegSpace),
		C.uint8_t(inout.WriteData),
		&readData,
	))
	if st == StatusSuccess {
		inout.ReadData = uint32(readData)
	}
	return st
}

func DimmSBRegWrite(sockInd uint8, inout *DimmSBInfo) Status {
	if inout == nil {
		return StatusArgPtrNull
	}
	return status(C.go_esmi_dimm_sb_reg_write(
		C.uint8_t(sockInd),
		C.uint8_t(inout.DimmAddr),
		C.uint8_t(inout.LID),
		C.uint16_t(inout.RegOffset),
		C.uint8_t(inout.RegSpace),
		C.uint8_t(inout.WriteData),
	))
}

func XGMIWidthSet(sockInd uint8, min uint8, max uint8) Status {
	return status(C.esmi_xgmi_width_set(C.uint8_t(sockInd), C.uint8_t(min), C.uint8_t(max)))
}

func XGMIWidthGet(sockInd uint32, min *uint8, max *uint8) Status {
	return status(C.esmi_xgmi_width_get(C.uint32_t(sockInd), cU8Ptr(min), cU8Ptr(max)))
}

func GMI3LinkWidthRangeSet(sockInd uint8, minLinkWidth uint8, maxLinkWidth uint8) Status {
	return status(C.esmi_gmi3_link_width_range_set(C.uint8_t(sockInd), C.uint8_t(minLinkWidth), C.uint8_t(maxLinkWidth)))
}

func APBEnable(sockInd uint32) Status {
	return status(C.esmi_apb_enable(C.uint32_t(sockInd)))
}

func APBDisable(sockInd uint32, pstate uint8) Status {
	return status(C.esmi_apb_disable(C.uint32_t(sockInd), C.uint8_t(pstate)))
}

func APBStatusGet(sockInd uint32, apbDisabled *uint8, pstate *uint8) Status {
	return status(C.esmi_apb_status_get(C.uint32_t(sockInd), cU8Ptr(apbDisabled), cU8Ptr(pstate)))
}

func SocketLclkDpmLevelSet(sockInd uint32, nbioID uint8, min uint8, max uint8) Status {
	return status(C.esmi_socket_lclk_dpm_level_set(C.uint32_t(sockInd), C.uint8_t(nbioID), C.uint8_t(min), C.uint8_t(max)))
}

func SocketLclkDpmLevelGet(sockInd uint8, nbioID uint8, nbio *DPMLevel) Status {
	if nbio == nil {
		return StatusArgPtrNull
	}
	var min C.uint8_t
	var max C.uint8_t
	st := status(C.go_esmi_socket_lclk_dpm_level_get(C.uint8_t(sockInd), C.uint8_t(nbioID), &min, &max))
	if st == StatusSuccess {
		nbio.MinDPMLevel = uint8(min)
		nbio.MaxDPMLevel = uint8(max)
	}
	return st
}

func PCIeLinkRateSet(sockInd uint8, rateCtrl uint8, prevMode *uint8) Status {
	return status(C.esmi_pcie_link_rate_set(C.uint8_t(sockInd), C.uint8_t(rateCtrl), cU8Ptr(prevMode)))
}

func DFPstateRangeSet(sockInd uint8, minPstate uint8, maxPstate uint8) Status {
	return status(C.esmi_df_pstate_range_set(C.uint8_t(sockInd), C.uint8_t(minPstate), C.uint8_t(maxPstate)))
}

func DFPstateRangeGet(sockInd uint8, minPstate *uint8, maxPstate *uint8) Status {
	return status(C.esmi_df_pstate_range_get(C.uint8_t(sockInd), cU8Ptr(minPstate), cU8Ptr(maxPstate)))
}

func XGMIPstateRangeSet(sockInd uint8, minState uint8, maxState uint8) Status {
	return status(C.esmi_xgmi_pstate_range_set(C.uint8_t(sockInd), C.uint8_t(minState), C.uint8_t(maxState)))
}

func XGMIPstateRangeGet(sockInd uint8, minState *uint8, maxState *uint8) Status {
	return status(C.esmi_xgmi_pstate_range_get(C.uint8_t(sockInd), cU8Ptr(minState), cU8Ptr(maxState)))
}

func PC6EnableSet(sockInd uint8, pc6Enable uint8) Status {
	return status(C.esmi_pc6_enable_set(C.uint8_t(sockInd), C.uint8_t(pc6Enable)))
}

func PC6EnableGet(sockInd uint8, currentPC6Enable *uint8) Status {
	return status(C.esmi_pc6_enable_get(C.uint8_t(sockInd), cU8Ptr(currentPC6Enable)))
}

func CC6EnableSet(sockInd uint8, cc6Enable uint8) Status {
	return status(C.esmi_cc6_enable_set(C.uint8_t(sockInd), C.uint8_t(cc6Enable)))
}

func CC6EnableGet(sockInd uint8, currentCC6Enable *uint8) Status {
	return status(C.esmi_cc6_enable_get(C.uint8_t(sockInd), cU8Ptr(currentCC6Enable)))
}

func CurrentIOBandwidthGet(sockInd uint8, link LinkIDBWType, ioBW *uint32) Status {
	cName := C.CString(link.LinkName)
	defer C.free(unsafe.Pointer(cName))
	return status(C.go_esmi_current_io_bandwidth_get(C.uint8_t(sockInd), C.uint32_t(link.BWType), cName, cU32Ptr(ioBW)))
}

func CurrentXGMIBWGet(sockInd uint8, link LinkIDBWType, xgmiBW *uint32) Status {
	cName := C.CString(link.LinkName)
	defer C.free(unsafe.Pointer(cName))
	return status(C.go_esmi_current_xgmi_bw_get(C.uint8_t(sockInd), C.uint32_t(link.BWType), cName, cU32Ptr(xgmiBW)))
}

func MetricsTableVersionGet(metricsVersion *uint32) Status {
	return status(C.esmi_metrics_table_version_get(cU32Ptr(metricsVersion)))
}

func MetricsTableGet(sockInd uint8, metricsTable *MetricTable) Status {
	if metricsTable == nil {
		return StatusArgPtrNull
	}
	var cTable C.struct_go_hsmp_metric_table
	st := status(C.go_esmi_metrics_table_get(C.uint8_t(sockInd), &cTable))
	if st == StatusSuccess {
		*metricsTable = metricTableFromC(cTable)
	}
	return st
}

func DRAMAddressMetricsTableGet(sockInd uint8, dramAddr *uint64) Status {
	return status(C.esmi_dram_address_metrics_table_get(C.uint8_t(sockInd), cU64Ptr(dramAddr)))
}

func TestHSMPMailbox(sockInd uint8, data *uint32) Status {
	return status(C.esmi_test_hsmp_mailbox(C.uint8_t(sockInd), cU32Ptr(data)))
}

func CPUFamilyGet(family *uint32) Status {
	return status(C.esmi_cpu_family_get(cU32Ptr(family)))
}

func CPUModelGet(model *uint32) Status {
	return status(C.esmi_cpu_model_get(cU32Ptr(model)))
}

func ThreadsPerCoreGet(threads *uint32) Status {
	return status(C.esmi_threads_per_core_get(cU32Ptr(threads)))
}

func NumberOfCPUsGet(cpus *uint32) Status {
	return status(C.esmi_number_of_cpus_get(cU32Ptr(cpus)))
}

func NumberOfSocketsGet(sockets *uint32) Status {
	return status(C.esmi_number_of_sockets_get(cU32Ptr(sockets)))
}

func FirstOnlineCoreOnSocket(socketIdx uint32, coreInd *uint32) Status {
	return status(C.esmi_first_online_core_on_socket(C.uint32_t(socketIdx), cU32Ptr(coreInd)))
}

func GetEnabledCommands(sockInd uint8, info *EnabledCommandsInfo) Status {
	if info == nil {
		return StatusArgPtrNull
	}
	var arg0 C.uint32_t
	var arg1 C.uint32_t
	var arg2 C.uint32_t
	st := status(C.go_esmi_get_enabled_commands(
		C.uint8_t(sockInd),
		C.bool(info.ReadMask),
		&arg0,
		&arg1,
		&arg2,
	))
	if st == StatusSuccess {
		info.Arg0 = uint32(arg0)
		info.Arg1 = uint32(arg1)
		info.Arg2 = uint32(arg2)
	}
	return st
}

func ErrMsg(esmiErr Status) string {
	msg := C.esmi_get_err_msg(C.esmi_status_t(esmiErr))
	if msg == nil {
		return esmiErr.String()
	}
	return C.GoString(msg)
}

func metricTableFromC(in C.struct_go_hsmp_metric_table) MetricTable {
	out := MetricTable{
		AccumulationCounter: uint32(in.accumulation_counter),

		MaxSocketTemperature:    uint32(in.max_socket_temperature),
		MaxVRTemperature:        uint32(in.max_vr_temperature),
		MaxHBMTemperature:       uint32(in.max_hbm_temperature),
		MaxSocketTemperatureAcc: uint64(in.max_socket_temperature_acc),
		MaxVRTemperatureAcc:     uint64(in.max_vr_temperature_acc),
		MaxHBMTemperatureAcc:    uint64(in.max_hbm_temperature_acc),

		SocketPowerLimit:    uint32(in.socket_power_limit),
		MaxSocketPowerLimit: uint32(in.max_socket_power_limit),
		SocketPower:         uint32(in.socket_power),

		Timestamp:       uint64(in.timestamp),
		SocketEnergyAcc: uint64(in.socket_energy_acc),
		CCDEnergyAcc:    uint64(in.ccd_energy_acc),
		XCDEnergyAcc:    uint64(in.xcd_energy_acc),
		AIDEnergyAcc:    uint64(in.aid_energy_acc),
		HBMEnergyAcc:    uint64(in.hbm_energy_acc),

		CCLKFrequencyLimit:   uint32(in.cclk_frequency_limit),
		GFXCLKFrequencyLimit: uint32(in.gfxclk_frequency_limit),
		FCLKFrequency:        uint32(in.fclk_frequency),
		UCLKFrequency:        uint32(in.uclk_frequency),

		MaxCCLKFrequency:   uint32(in.max_cclk_frequency),
		MinCCLKFrequency:   uint32(in.min_cclk_frequency),
		MaxGFXCLKFrequency: uint32(in.max_gfxclk_frequency),
		MinGFXCLKFrequency: uint32(in.min_gfxclk_frequency),
		MaxLclkDPMRange:    uint32(in.max_lclk_dpm_range),
		MinLclkDPMRange:    uint32(in.min_lclk_dpm_range),

		XGMIWidth:   uint32(in.xgmi_width),
		XGMIBitrate: uint32(in.xgmi_bitrate),

		SocketC0Residency:           uint32(in.socket_c0_residency),
		SocketGFXBusy:               uint32(in.socket_gfx_busy),
		DRAMBandwidthUtilization:    uint32(in.dram_bandwidth_utilization),
		SocketC0ResidencyAcc:        uint64(in.socket_c0_residency_acc),
		SocketGFXBusyAcc:            uint64(in.socket_gfx_busy_acc),
		DRAMBandwidthAcc:            uint64(in.dram_bandwidth_acc),
		MaxDRAMBandwidth:            uint32(in.max_dram_bandwidth),
		DRAMBandwidthUtilizationAcc: uint64(in.dram_bandwidth_utilization_acc),

		PROCHOTResidencyAcc:   uint32(in.prochot_residency_acc),
		PPTResidencyAcc:       uint32(in.ppt_residency_acc),
		SocketTHMResidencyAcc: uint32(in.socket_thm_residency_acc),
		VRTHMResidencyAcc:     uint32(in.vr_thm_residency_acc),
		HBMTHMResidencyAcc:    uint32(in.hbm_thm_residency_acc),
		Spare:                 uint32(in.spare),
	}

	for i := 0; i < len(out.SocclkFrequency); i++ {
		out.SocclkFrequency[i] = uint32(in.socclk_frequency[i])
		out.VclkFrequency[i] = uint32(in.vclk_frequency[i])
		out.DclkFrequency[i] = uint32(in.dclk_frequency[i])
		out.LclkFrequency[i] = uint32(in.lclk_frequency[i])

		out.FCLKFrequencyTable[i] = uint32(in.fclk_frequency_table[i])
		out.UCLKFrequencyTable[i] = uint32(in.uclk_frequency_table[i])
		out.SocclkFrequencyTable[i] = uint32(in.socclk_frequency_table[i])
		out.VclkFrequencyTable[i] = uint32(in.vclk_frequency_table[i])
		out.DclkFrequencyTable[i] = uint32(in.dclk_frequency_table[i])
		out.LclkFrequencyTable[i] = uint32(in.lclk_frequency_table[i])

		out.PCIEBandwidthAcc[i] = uint64(in.pcie_bandwidth_acc[i])
	}

	for i := 0; i < len(out.GFXCLKFrequencyAcc); i++ {
		out.GFXCLKFrequencyAcc[i] = uint64(in.gfxclk_frequency_acc[i])
		out.XGMIReadBandwidthAcc[i] = uint64(in.xgmi_read_bandwidth_acc[i])
		out.XGMIWriteBandwidthAcc[i] = uint64(in.xgmi_write_bandwidth_acc[i])
		out.GFXCLKFrequency[i] = uint32(in.gfxclk_frequency[i])
	}

	for i := 0; i < len(out.CCLKFrequencyAcc); i++ {
		out.CCLKFrequencyAcc[i] = uint64(in.cclk_frequency_acc[i])
	}

	return out
}
