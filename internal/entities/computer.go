package entities

type Computer struct {
	Id              int    `json:"id" db:"id,omitempty"`
	CpuName         string `json:"cpu_name" db:"cpu_name,omitempty"`
	CpuModel        string `json:"cpu_model" db:"cpu_model,omitempty"`
	CpuCoreCount    int    `json:"cpu_core_count" db:"cpu_core_count,omitempty"`
	CpuThreadCount  int    `json:"cpu_thread_count" db:"cpu_thread_count,omitempty"`
	GpuName         string `json:"gpu_name" db:"gpu_name,omitempty"`
	GpuModel        string `json:"gpu_model" db:"gpu_model,omitempty"`
	GpuMem          int    `json:"gpu_mem" db:"gpu_mem,omitempty"`
	GpuPowerInput   int    `json:"gpu_power_input" db:"gpu_power_input,omitempty"`
	MotherBoardName string `json:"mother_board_name" db:"mother_board_name,omitempty"`
	Ram             int    `json:"ram" db:"ram,omitempty"`
	PowerUnit       int    `json:"power_unit" db:"power_unit,omitempty"`
}
