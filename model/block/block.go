package block

type Block struct {
	Version    string `gorm:"version"`
	Number     int    `gorm:"number"`
	Timestamp  string `gorm:"timestamp"`
	TxHash     string `gorm:"tx_hash"`
	ParentHash string `gorm:"parent_hash"`
	WitRoot    string `gorm:"wit_root"`
}
