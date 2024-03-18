package entities

type Computer struct {
	Id              int    `json:"id"`
	CpuName         string `json:"chip_name"`
	CpuModel        string `json:"cpu_model"`
	CpuCoreCount    int    `json:"cpu_core_count"`
	CpuThreadCount  int    `json:"cpu_thread_count"`
	GpuName         string `json:"gpu_name"`
	GpuModel        string `json:"gpu_model"`
	GpuMem          int    `json:"gpu_mem"`
	GpuPowerInput   int    `json:"gpu_power_input"`
	MotherBoardName string `json:"mother_board_name"`
	Ram             int    `json:"ram"`
	PowerUnit       int    `json:"power_unit"`
}
