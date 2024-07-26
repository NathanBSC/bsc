package miner

//go:generate go run github.com/fjl/gencodec -type MBConfig -formats toml -out gen_mb_config.go
type MBConfig struct {
	// Generate two consecutive blocks for the same parent block
	DoubleSign bool
	// Skip block production for in-turn validators at a specified offset
	SkipOffsetInturn *uint64 `toml:",omitempty"`
	// Delay broadcasting mined blocks by a specified number of blocks
	BroadcastDelayBlocks uint64
	// Mining time (milliseconds) for the last block in every turn
	LastBlockMiningTime uint64
}

var DefaultMBConfig = MBConfig{
	DoubleSign:           false,
	BroadcastDelayBlocks: 0,
	LastBlockMiningTime:  2500,
}
