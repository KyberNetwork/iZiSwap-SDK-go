//go:generate go run github.com/tinylib/msgp -unexported -tests=false -v
//msgp:tuple PoolInfo
//msgp:shim *big.Int as:[]byte using:msgpencode.EncodeInt/msgpencode.DecodeInt
//msgp:ignore SwapResult

package swap

import (
	"math/big"
)

type SwapResult struct {
	AmountX      *big.Int
	AmountY      *big.Int
	CurrentPoint int
	Liquidity    *big.Int
	LiquidityX   *big.Int
}

type PoolInfo struct {
	CurrentPoint int
	PointDelta   int
	LeftMostPt   int
	RightMostPt  int
	Fee          int
	Liquidity    *big.Int
	LiquidityX   *big.Int
	Liquidities  []LiquidityPoint
	LimitOrders  []LimitOrderPoint
}
