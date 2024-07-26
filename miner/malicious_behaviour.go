package miner

type MBConfig struct {
	DoubleSign           bool   // Generate two consecutive blocks for the same parent block
	SkipOffsetInturn     uint64 // Skip block production for in-turn validators at a specified offset
	BroadcastDelayBlocks uint64 // Delay broadcasting mined blocks by a specified number of blocks
	LastBlockMiningTime  uint64 // Mining time (milliseconds) for the last block in every turn
}

var DefaultMBConfig = MBConfig{
	DoubleSign:           false,
	SkipOffsetInturn:     100,
	BroadcastDelayBlocks: 0,
	LastBlockMiningTime:  2500,
}
