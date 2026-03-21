package queries

const (
	SelectAllAccounts   GoQuery = `SELECT id, account_number, cif_number, status, created_at, updated_at, deleted_at FROM accounts WHERE deleted_at IS NULL`
	SelectAccountsByCIF GoQuery = `SELECT id, account_number, cif_number, status, created_at, updated_at, deleted_at FROM accounts WHERE cif_number = ? AND deleted_at IS NULL`
	InsertAccount       GoQuery = `INSERT INTO accounts (account_number, cif_number, status) VALUES (?, ?, ?)`
	UpdateAccountByID   GoQuery = `UPDATE accounts SET account_number = ?, cif_number = ?, status = ?, updated_at = ? WHERE id = ?`
	DeleteAccountByID   GoQuery = `UPDATE accounts SET deleted_at = ? WHERE id = ?`
)
