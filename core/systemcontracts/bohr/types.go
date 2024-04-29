package bohr

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/ValidatorContract
	MainnetValidatorContract string
)

// contract codes for Chapel upgrade
var (
	//go:embed chapel/ValidatorContract
	ChapelValidatorContract string
)

// contract codes for Rialto upgrade
var (
	//go:embed rialto/ValidatorContract
	RialtoValidatorContract string
)
