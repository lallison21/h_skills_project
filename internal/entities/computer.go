package entities

type Computer struct {
	Id              int    `json:"id" db:"id"`
	CpuName         string `json:"cpu_name" db:"cpu_name"`
	CpuModel        string `json:"cpu_model" db:"cpu_model"`
	CpuCoreCount    int    `json:"cpu_core_count" db:"cpu_core_count"`
	CpuThreadCount  int    `json:"cpu_thread_count" db:"cpu_thread_count"`
	GpuName         string `json:"gpu_name" db:"gpu_name"`
	GpuModel        string `json:"gpu_model" db:"gpu_model"`
	GpuMem          int    `json:"gpu_mem" db:"gpu_mem"`
	GpuPowerInput   int    `json:"gpu_power_input" db:"gpu_power_input"`
	MotherBoardName string `json:"mother_board_name" db:"mother_board_name"`
	Ram             int    `json:"ram" db:"ram"`
	PowerUnit       int    `json:"power_unit" db:"power_unit"`
}
