// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

import "time"

// MeiHuaShiJianQiGua 梅花易数时间起卦
func MeiHuaShiJianQiGua(ctx DivinationContext) *PaipanResult {
	yearZhiRunes := []rune(ctx.GanZhi[0])
	yearZhi := string(yearZhiRunes[1])
	yearZhiNum := GetZhiIndex(yearZhi)

	hourZhiRunes := []rune(ctx.GanZhi[3])
	hourZhi := string(hourZhiRunes[1])
	hourZhiNum := GetZhiIndex(hourZhi)

	sumYMD := yearZhiNum + ctx.LunarMonth + ctx.LunarDay
	sumYMDH := sumYMD + hourZhiNum

	shangNum := sumYMD % 8
	xiaNum := sumYMDH % 8
	if shangNum == 0 {
		shangNum = 8
	}
	if xiaNum == 0 {
		xiaNum = 8
	}

	bianPos := 6 - (sumYMDH % 6)
	if bianPos == 6 {
		bianPos = 0
	}

	return BuildLiuyaoResult(ctx, shangNum, xiaNum, bianPos, "梅花时间起卦")
}

// MeiHuaShuZiQiGua 梅花易数数字起卦
func MeiHuaShuZiQiGua(ctx DivinationContext, numbers [3]int) *PaipanResult {
	shangNum := numbers[0] % 8
	xiaNum := numbers[1] % 8
	if shangNum == 0 {
		shangNum = 8
	}
	if xiaNum == 0 {
		xiaNum = 8
	}
	bianPos := 6 - ((numbers[0] + numbers[1] + numbers[2]) % 6)
	if bianPos == 6 {
		bianPos = 0
	}
	return BuildLiuyaoResult(ctx, shangNum, xiaNum, bianPos, "梅花数字起卦")
}

// MeiHuaSuiJiQiGua 梅花易数随机起卦
func MeiHuaSuiJiQiGua(ctx DivinationContext) *PaipanResult {
	now := time.Now().UnixMilli()
	n0 := int((now/1000)%9) + 1
	n1 := int((now/10000)%9) + 1
	n2 := int((now/100000)%9) + 1
	return MeiHuaShuZiQiGua(ctx, [3]int{n0, n1, n2})
}

// MeiHuaShouYaoQiGua 梅花易数手摇起卦
func MeiHuaShouYaoQiGua(ctx DivinationContext, yaos [6]int) *PaipanResult {
	shangYinYang := [3]int{YaoToYinYang(yaos[2]), YaoToYinYang(yaos[1]), YaoToYinYang(yaos[0])}
	xiaYinYang := [3]int{YaoToYinYang(yaos[5]), YaoToYinYang(yaos[4]), YaoToYinYang(yaos[3])}

	shangGua := Yaos2BaGuaNum(shangYinYang)
	xiaGua := Yaos2BaGuaNum(xiaYinYang)

	benGuaName := Guas64[shangGua-1][xiaGua-1]
	benGua := GuaInfo{Name: benGuaName, Yaos: yaos}

	huYaos := GetHuGuaYaos(yaos)
	huGua := GuaInfo{Name: Get64GuaNameByYaos(huYaos), Yaos: huYaos}

	bianPos := -1
	for i := 0; i < 6; i++ {
		if yaos[i] == YaoLaoYang || yaos[i] == YaoLaoYin {
			bianPos = i
			break
		}
	}
	if bianPos == -1 {
		bianPos = 0
	}

	bianYaos := GetBianGuaYaos(yaos, bianPos)
	bianGua := &GuaInfo{Name: Get64GuaNameByYaos(bianYaos), Yaos: bianYaos}

	return &PaipanResult{
		Method:         "梅花手摇起卦",
		GanZhi:         ctx.GanZhi,
		XunKong:        ctx.XunKong,
		BenGua:         benGua,
		HuGua:          huGua,
		BianGua:        bianGua,
		BianYao:        bianPos,
		ShangGua:       GuaNum2Name(shangGua),
		XiaGua:         GuaNum2Name(xiaGua),
		ShangGuaWuXing: GuaNum2WuXing(shangGua),
		XiaGuaWuXing:   GuaNum2WuXing(xiaGua),
	}
}
