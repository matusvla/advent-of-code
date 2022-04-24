package main

import (
	"fmt"
	"math"
	"strconv"
)

const input = "620D49005AD2245800D0C9E72BD279CAFB0016B1FA2B1802DC00D0CC611A47FCE2A4ACE1DD144BFABBFACA002FB2C6F33DFF4A0C0119B169B013005F003720004263644384800087C3B8B51C26B449130802D1A0068A5BD7D49DE793A48B5400D8293B1F95C5A3005257B880F5802A00084C788AD0440010F8490F608CACE034401AB4D0F5802726B3392EE2199628CEA007001884005C92015CC8051800130EC0468A01042803B8300D8E200788018C027890088CE0049006028012AB00342A0060801B2EBE400424933980453EFB2ABB36032274C026E4976001237D964FF736AFB56F254CB84CDF136C1007E7EB42298FE713749F973F7283005656F902A004067CD27CC1C00D9CB5FDD4D0014348010C8331C21710021304638C513006E234308B060094BEB76CE3966AA007C6588A5670DC3754395485007A718A7F149CA2DD3B6E7B777800118E7B59C0ECF5AE5D3B6CB1496BAE53B7ADD78C013C00CD2629BF5371D1D4C537EA6E3A3E95A3E180592AC7246B34032CF92804001A1CCF9BA521782ECBD69A98648BC18025800F8C9C37C827CA7BEFB31EADF0AE801BA42B87935B8EF976194EEC426AAF640168CECAF84BC004AE7D1673A6A600B4AB65802D230D35CF81B803D3775683F3A3860087802132FB32F322C92A4C402524F2DE006E8000854378F710C0010D8F30FE224AE428C015E00D40401987F06E3600021D0CE3EC228DA000574E4C3080182931E936E953B200BF656E15400D3496E4A725B92998027C00A84EEEE6B347D30BE60094E537AA73A1D600B880371AA36C3200043235C4C866C018E4963B7E7AA2B379918C639F1550086064BB148BA499EC731004E1AC966BDBC7646600C080370822AC4C1007E38C428BE0008741689D0ECC01197CF216EA16802D3748FE91B25CAF6D5F11C463004E4FD08FAF381F6004D3232CC93E7715B463F780"

func main() {
	var binaryInput string
	for i := range input {
		bi, err := strconv.ParseInt(input[i:i+1], 16, 64)
		if err != nil {
			panic(err)
		}
		b := fmt.Sprintf("%04s", strconv.FormatInt(bi, 2))
		binaryInput += b
	}
	//fmt.Println(binaryInput)
	p, _ := NewPacket(binaryInput)
	fmt.Println(p.versionSum())
	fmt.Println(p.value())

}

const (
	literalType = 4
)

func NewPacket(s string) (Packet, int) {
	version, err := strconv.ParseInt(s[0:3], 2, 64)
	if err != nil {
		panic(err)
	}

	typeID, err := strconv.ParseInt(s[3:6], 2, 64)
	if err != nil {
		panic(err)
	}

	if typeID == literalType {
		p, i := NewLiteralPacket(version, s[6:])
		return p, i + 6
	}
	i, p := NewOperationPacket(version, typeID, s[6:])
	return i, p + 6
}

type Packet interface {
	versionSum() int64
	value() int64
}

type LiteralPacket struct {
	version int64
	val     int64
}

func NewLiteralPacket(version int64, body string) (*LiteralPacket, int) {
	var valueStr string

	var usedCount int
	var outStrs []string
	for i := 0; ; i += 5 {
		usedCount += 5
		outStrs = append(outStrs, body[i:i+5])
		valueStr += body[i+1 : i+5]
		if body[i] == '0' {
			break
		}
	}
	value, err := strconv.ParseInt(valueStr, 2, 64)
	if err != nil {
		panic(err)
	}
	return &LiteralPacket{
		version: version,
		val:     value,
	}, usedCount
}

func (l *LiteralPacket) versionSum() int64 {
	return l.version
}

func (l *LiteralPacket) value() int64 {
	return l.val
}

type OperationPacket struct {
	version      int64
	typeID       int64
	lengthTypeID uint8
	subpackets   []Packet
}

func (o *OperationPacket) value() int64 {
	var result int64
	switch o.typeID {
	case 0:
		result = int64(0)
		for _, sp := range o.subpackets {
			result += sp.value()
		}
	case 1:
		result = int64(1)
		for _, sp := range o.subpackets {
			result *= sp.value()
		}
	case 2:
		result = math.MaxInt64
		for _, sp := range o.subpackets {
			if sp.value() < result {
				result = sp.value()
			}
		}
	case 3:
		result = -1
		for _, sp := range o.subpackets {
			if sp.value() > result {
				result = sp.value()
			}
		}
	case 5:
		result = 0
		if o.subpackets[0].value() > o.subpackets[1].value() {
			result = 1
		}
	case 6:
		result = 0
		if o.subpackets[0].value() < o.subpackets[1].value() {
			result = 1
		}
	case 7:
		result = 0
		if o.subpackets[0].value() == o.subpackets[1].value() {
			result = 1
		}
	}
	return result
}

func NewOperationPacket(version, typeID int64, body string) (*OperationPacket, int) {
	lengthTypeID := body[0] - '0'
	switch lengthTypeID {
	case 0:
		p, i := NewOperationLengthPacket(version, typeID, body[1:])
		return p, i + 1
	case 1:
		p, i := NewOperationSizePacket(version, typeID, body[1:])
		return p, i + 1
	default:
		panic("ltid")
	}
}

func NewOperationLengthPacket(version, typeID int64, body string) (*OperationPacket, int) {
	var subpackets []Packet
	key := body[0:15]
	totalLengthInBits, err := strconv.ParseInt(key, 2, 64)
	if err != nil {
		panic(err)
	}
	j := 15
	for j < 15+int(totalLengthInBits) {
		p, index := NewPacket(body[j:])
		j += index
		subpackets = append(subpackets, p)
	}
	return &OperationPacket{
		version:      version,
		typeID:       typeID,
		lengthTypeID: 0,
		subpackets:   subpackets,
	}, j
}

func NewOperationSizePacket(version, typeID int64, body string) (*OperationPacket, int) {
	var subpackets []Packet
	numberOfSubPacketsImmediatelyContained, err := strconv.ParseInt(body[0:11], 2, 64)
	if err != nil {
		panic(err)
	}
	i := 11
	for j := 0; j < int(numberOfSubPacketsImmediatelyContained); j++ {
		p, index := NewPacket(body[i:])
		i += index
		subpackets = append(subpackets, p)
	}
	return &OperationPacket{
		version:      version,
		typeID:       typeID,
		lengthTypeID: 1,
		subpackets:   subpackets,
	}, i
}

func (o *OperationPacket) versionSum() int64 {
	result := o.version
	for _, sp := range o.subpackets {
		result += sp.versionSum()
	}
	return result
}
