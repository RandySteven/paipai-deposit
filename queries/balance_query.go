package queries

const (
	SelectAllBalances         GoQuery = `SELECT id, account_id, balance_amount, created_at, updated_at, deleted_at FROM balances WHERE deleted_at IS NULL`
	SelectBalancesByAccountID GoQuery = `SELECT id, account_id, balance_amount, created_at, updated_at, deleted_at FROM balances WHERE account_id = ? AND deleted_at IS NULL`
	SelectBalanceByID         GoQuery = `SELECT id, account_id, balance_amount, created_at, updated_at, deleted_at FROM balances WHERE id = ? AND deleted_at IS NULL`
	InsertBalance             GoQuery = `INSERT INTO balances (account_id, balance_amount) VALUES (?, ?)`
	UpdateBalanceByID         GoQuery = `UPDATE balances SET account_id = ?, balance_amount = ?, updated_at = ? WHERE id = ?`
	DeleteBalanceByID         GoQuery = `UPDATE balances SET deleted_at = ? WHERE id = ?`
)
