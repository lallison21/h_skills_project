CREATE TABLE IF NOT EXISTS computers (
    id SERIAL PRIMARY KEY,
    cpu_name VARCHAR(255) NOT NULL,
    cpu_model VARCHAR(255) NOT NULL,
    cpu_core_count INTEGER NOT NULL,
    cpu_thread_count INTEGER NOT NULL,
    gpu_name VARCHAR(255) NOT NULL,
    gpu_model VARCHAR(255) NOT NULL,
    gpu_mem INTEGER NOT NULL,
    gpu_power_input INTEGER NOT NULL,
    mother_board_name VARCHAR(255) NOT NULL,
    ram INTEGER NOT NULL,
    power_unit INTEGER NOT NULL
);

INSERT INTO computers (cpu_name, cpu_model, cpu_core_count, cpu_thread_count, gpu_name, gpu_model, gpu_mem, gpu_power_input, mother_board_name, ram, power_unit)
VALUES ('Intel Core i7', '8700K', 6, 12, 'Nvidia GeForce RTX 2080', 'Ti', 11, 250, 'Asus ROG Strix Z370-E Gaming', 16, 750);
