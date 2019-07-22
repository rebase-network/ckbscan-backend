package address

type Address struct {
	// balance decimal "balance"
	// address_hash binary "address_hash"
	// cell_consumed decimal "cell_consumed"
	// ckb_transactions_count decimal "ckb_transactions_count"
	// created_at *time.Time "created_at"
	// updated_at *time.Time "updated_at"
	// lock_hash binary "lock_hash"
	pending_reward_blocks_count int `gorm:"pending_reward_blocks_count"`
	// index ["address_hash"]
	// index ["lock_hash"]
}