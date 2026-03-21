package queries

const (
	MigrateWorkflow TableMigration = `CREATE TABLE IF NOT EXISTS workflows (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		run_id VARCHAR(255) NOT NULL,
		workflow_id VARCHAR(255) NOT NULL,
		activity_name VARCHAR(255) NOT NULL,
		metadata JSON NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL,

		INDEX idx_run_id (run_id),
		INDEX idx_workflow_id (workflow_id),
		INDEX idx_activity_name (activity_name)
	)`

	MigrateAccount TableMigration = `CREATE TABLE IF NOT EXISTS accounts (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		account_number VARCHAR(255) NOT NULL,
		cif_number VARCHAR(255) NOT NULL,
		status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL,

		INDEX idx_account_number (account_number),
		INDEX idx_cif_number (cif_number),
		INDEX idx_status (status)
	)`

	MigrateBalance TableMigration = `CREATE TABLE IF NOT EXISTS balances (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		account_id BIGINT UNSIGNED NOT NULL,
		balance_amount DOUBLE NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL,

		INDEX idx_account_id (account_id)
	)`

	MigrateTransactionHistory TableMigration = `CREATE TABLE IF NOT EXISTS transaction_histories (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		account_id BIGINT UNSIGNED NOT NULL,
		balance_id BIGINT UNSIGNED NOT NULL,
		amount DOUBLE NOT NULL,
		transaction_type VARCHAR(50) NOT NULL,
		transaction_amount DOUBLE NOT NULL,
		transaction_date TIMESTAMP NOT NULL,
		transaction_status VARCHAR(50) NOT NULL,
		transaction_description TEXT,
		transaction_reference VARCHAR(255),
		transaction_fee DOUBLE NOT NULL DEFAULT 0,
		transaction_tax DOUBLE NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL,

		INDEX idx_account_id (account_id),
		INDEX idx_balance_id (balance_id),
		INDEX idx_transaction_type (transaction_type),
		INDEX idx_transaction_date (transaction_date),
		INDEX idx_transaction_status (transaction_status)
	)`
)
