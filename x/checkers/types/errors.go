package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1103, "black address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
)
