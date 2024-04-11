package computer_repository

import (
	"database/sql"
	"fmt"
	"github.com/lallison/h_skills_project/internal/entities"
	"log"
	"log/slog"
)

type ComputerRepository struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *ComputerRepository {
	repo := &ComputerRepository{db: db}
	return repo
}

func (r *ComputerRepository) Computers(logger *slog.Logger) ([]*entities.Computer, error) {
	var res []*entities.Computer

	stmt, err := r.db.Prepare("SELECT * FROM computers")
	if err != nil {
		return res, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return res, fmt.Errorf("failed to query rows: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	for rows.Next() {
		var c entities.Computer

		if err := rows.Scan(&c.Id, &c.CpuName, &c.CpuModel, &c.CpuCoreCount, &c.CpuThreadCount, &c.GpuName, &c.GpuModel, &c.GpuMem, &c.GpuPowerInput, &c.MotherBoardName, &c.Ram, &c.PowerUnit); err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return res, fmt.Errorf("failed to scan rows: %w", err)
	}

	return res, nil
}

func (r *ComputerRepository) ComputerByID(logger *slog.Logger, id string) (*entities.Computer, error) {
	var c entities.Computer

	stmt, err := r.db.Prepare("SELECT * FROM computers WHERE id = $1")
	if err != nil {
		return &c, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close query: %v", err)
		}
	}()

	row := stmt.QueryRow(id)
	if err := row.Scan(&c.Id, &c.CpuName, &c.CpuModel, &c.CpuCoreCount, &c.CpuThreadCount, &c.GpuName, &c.GpuModel, &c.GpuMem, &c.GpuPowerInput, &c.MotherBoardName, &c.Ram, &c.PowerUnit); err != nil {
		return &c, fmt.Errorf("failed to scan row: %w", err)
	}

	return &c, nil
}

func (r *ComputerRepository) CreateComputer(logger *slog.Logger, computer *entities.Computer) error {
	stmt, err := r.db.Prepare("INSERT INTO computers (cpu_name, cpu_model, cpu_core_count, cpu_thread_count, gpu_name, gpu_model, gpu_mem, gpu_power_input, mother_board_name, ram, power_unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close query: %v", err)
		}
	}()

	_, err = stmt.Exec(computer.CpuName, computer.CpuModel, computer.CpuCoreCount, computer.CpuThreadCount, computer.GpuName, computer.GpuModel, computer.GpuMem, computer.GpuPowerInput, computer.MotherBoardName, computer.Ram, computer.PowerUnit)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
