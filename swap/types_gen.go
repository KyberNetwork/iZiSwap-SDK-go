package swap

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/KyberNetwork/iZiSwap-SDK-go/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PoolInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.CurrentPoint, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "CurrentPoint")
		return
	}
	z.PointDelta, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "PointDelta")
		return
	}
	z.LeftMostPt, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "LeftMostPt")
		return
	}
	z.RightMostPt, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "RightMostPt")
		return
	}
	z.Fee, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Fee")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Liquidity")
			return
		}
		z.Liquidity = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Liquidity))
			if err != nil {
				err = msgp.WrapError(err, "Liquidity")
				return
			}
			z.Liquidity = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "LiquidityX")
			return
		}
		z.LiquidityX = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.LiquidityX))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityX")
				return
			}
			z.LiquidityX = msgpencode.DecodeInt(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Liquidities")
		return
	}
	if cap(z.Liquidities) >= int(zb0004) {
		z.Liquidities = (z.Liquidities)[:zb0004]
	} else {
		z.Liquidities = make([]LiquidityPoint, zb0004)
	}
	for za0001 := range z.Liquidities {
		err = z.Liquidities[za0001].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Liquidities", za0001)
			return
		}
	}
	var zb0005 uint32
	zb0005, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "LimitOrders")
		return
	}
	if cap(z.LimitOrders) >= int(zb0005) {
		z.LimitOrders = (z.LimitOrders)[:zb0005]
	} else {
		z.LimitOrders = make([]LimitOrderPoint, zb0005)
	}
	for za0002 := range z.LimitOrders {
		err = z.LimitOrders[za0002].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "LimitOrders", za0002)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 9
	err = en.Append(0x99)
	if err != nil {
		return
	}
	err = en.WriteInt(z.CurrentPoint)
	if err != nil {
		err = msgp.WrapError(err, "CurrentPoint")
		return
	}
	err = en.WriteInt(z.PointDelta)
	if err != nil {
		err = msgp.WrapError(err, "PointDelta")
		return
	}
	err = en.WriteInt(z.LeftMostPt)
	if err != nil {
		err = msgp.WrapError(err, "LeftMostPt")
		return
	}
	err = en.WriteInt(z.RightMostPt)
	if err != nil {
		err = msgp.WrapError(err, "RightMostPt")
		return
	}
	err = en.WriteInt(z.Fee)
	if err != nil {
		err = msgp.WrapError(err, "Fee")
		return
	}
	if z.Liquidity == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Liquidity))
		if err != nil {
			err = msgp.WrapError(err, "Liquidity")
			return
		}
	}
	if z.LiquidityX == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.LiquidityX))
		if err != nil {
			err = msgp.WrapError(err, "LiquidityX")
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.Liquidities)))
	if err != nil {
		err = msgp.WrapError(err, "Liquidities")
		return
	}
	for za0001 := range z.Liquidities {
		err = z.Liquidities[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Liquidities", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.LimitOrders)))
	if err != nil {
		err = msgp.WrapError(err, "LimitOrders")
		return
	}
	for za0002 := range z.LimitOrders {
		err = z.LimitOrders[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "LimitOrders", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 9
	o = append(o, 0x99)
	o = msgp.AppendInt(o, z.CurrentPoint)
	o = msgp.AppendInt(o, z.PointDelta)
	o = msgp.AppendInt(o, z.LeftMostPt)
	o = msgp.AppendInt(o, z.RightMostPt)
	o = msgp.AppendInt(o, z.Fee)
	if z.Liquidity == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Liquidity))
	}
	if z.LiquidityX == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.LiquidityX))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.Liquidities)))
	for za0001 := range z.Liquidities {
		o, err = z.Liquidities[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Liquidities", za0001)
			return
		}
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.LimitOrders)))
	for za0002 := range z.LimitOrders {
		o, err = z.LimitOrders[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "LimitOrders", za0002)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.CurrentPoint, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CurrentPoint")
		return
	}
	z.PointDelta, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "PointDelta")
		return
	}
	z.LeftMostPt, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LeftMostPt")
		return
	}
	z.RightMostPt, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "RightMostPt")
		return
	}
	z.Fee, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Fee")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Liquidity = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Liquidity))
			if err != nil {
				err = msgp.WrapError(err, "Liquidity")
				return
			}
			z.Liquidity = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.LiquidityX = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.LiquidityX))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityX")
				return
			}
			z.LiquidityX = msgpencode.DecodeInt(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Liquidities")
		return
	}
	if cap(z.Liquidities) >= int(zb0004) {
		z.Liquidities = (z.Liquidities)[:zb0004]
	} else {
		z.Liquidities = make([]LiquidityPoint, zb0004)
	}
	for za0001 := range z.Liquidities {
		bts, err = z.Liquidities[za0001].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Liquidities", za0001)
			return
		}
	}
	var zb0005 uint32
	zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LimitOrders")
		return
	}
	if cap(z.LimitOrders) >= int(zb0005) {
		z.LimitOrders = (z.LimitOrders)[:zb0005]
	} else {
		z.LimitOrders = make([]LimitOrderPoint, zb0005)
	}
	for za0002 := range z.LimitOrders {
		bts, err = z.LimitOrders[za0002].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "LimitOrders", za0002)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolInfo) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.IntSize + msgp.IntSize + msgp.IntSize + msgp.IntSize
	if z.Liquidity == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Liquidity))
	}
	if z.LiquidityX == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.LiquidityX))
	}
	s += msgp.ArrayHeaderSize
	for za0001 := range z.Liquidities {
		s += z.Liquidities[za0001].Msgsize()
	}
	s += msgp.ArrayHeaderSize
	for za0002 := range z.LimitOrders {
		s += z.LimitOrders[za0002].Msgsize()
	}
	return
}