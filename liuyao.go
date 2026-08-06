// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

import "time"

// BuildLiuyaoResult 构建六爻排盘结果（包含完整装卦信息）
func BuildLiuyaoResult(ctx DivinationContext, shangGua, xiaGua, bianPos int, method string) *PaipanResult {
	shang := BaGuaNum2Yaos(shangGua)
	xia := BaGuaNum2Yaos(xiaGua)

	benYaos := [6]int{
		shang[2], // 六爻
		shang[1], // 五爻
		shang[0], // 四爻
		xia[2],   // 三爻
		xia[1],   // 二爻
		xia[0],   // 初爻
	}

	if bianPos >= 0 && bianPos < 6 {
		benYaos[bianPos] += 2
	}

	benGuaName := Guas64[shangGua-1][xiaGua-1]
	benGua := GuaInfo{Name: benGuaName, Yaos: benYaos}

	huYaos := GetHuGuaYaos(benYaos)
	huGua := GuaInfo{Name: Get64GuaNameByYaos(huYaos), Yaos: huYaos}

	bianYaos := GetBianGuaYaos(benYaos, bianPos)
	bianGuaName := Get64GuaNameByYaos(bianYaos)
	bianGua := &GuaInfo{Name: bianGuaName, Yaos: bianYaos}

	benGuaEx := GetGuaEx(benGuaName)
	bianGuaExRaw := GetGuaEx(bianGuaName)

	var finalBianGuaEx *GuaExData
	if benGuaEx != nil && bianGuaExRaw != nil {
		bianLiuQin := ReloadLiuqin(benGuaEx.WuXing, bianGuaExRaw.GanZhi)
		lqStr := ""
		for _, lq := range bianLiuQin {
			lqStr += lq
		}
		finalBianGuaEx = &GuaExData{
			Name:    bianGuaExRaw.Name,
			LiuQin:  lqStr,
			GanZhi:  bianGuaExRaw.GanZhi,
			Shi:     bianGuaExRaw.Shi,
			Ying:    bianGuaExRaw.Ying,
			FuCang:  bianGuaExRaw.FuCang,
			GuaShen: bianGuaExRaw.GuaShen,
			BaGong:  bianGuaExRaw.BaGong,
			WuXing:  bianGuaExRaw.WuXing,
			Kind:    bianGuaExRaw.Kind,
		}
	}

	var liuShen []string
	dayGanRunes := []rune(ctx.GanZhi[2])
	if len(dayGanRunes) >= 1 {
		liuShen = GetLiuShen(string(dayGanRunes[0]))
	}

	return &PaipanResult{
		Method:         method,
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
		BenGuaEx:       benGuaEx,
		BianGuaEx:      finalBianGuaEx,
		LiuShen:        liuShen,
	}
}

// LiuYaoShiJianQiGua 六爻时间起卦
func LiuYaoShiJianQiGua(ctx DivinationContext) *PaipanResult {
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

	return BuildLiuyaoResult(ctx, shangNum, xiaNum, bianPos, "六爻时间起卦")
}

// LiuYaoShuZiQiGua 六爻数字起卦。numbers=[上卦数, 下卦数, 变爻数]
func LiuYaoShuZiQiGua(ctx DivinationContext, numbers [3]int) *PaipanResult {
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
	return BuildLiuyaoResult(ctx, shangNum, xiaNum, bianPos, "六爻数字起卦")
}

// LiuYaoSuiJiQiGua 六爻随机起卦
func LiuYaoSuiJiQiGua(ctx DivinationContext) *PaipanResult {
	now := time.Now().UnixMilli()
	n0 := int((now/1000)%9) + 1
	n1 := int((now/10000)%9) + 1
	n2 := int((now/100000)%9) + 1
	return LiuYaoShuZiQiGua(ctx, [3]int{n0, n1, n2})
}

// LiuYaoShouYaoQiGua 六爻手动起卦。yaos索引0=六爻(最上), 5=初爻(最下)
func LiuYaoShouYaoQiGua(ctx DivinationContext, yaos [6]int) *PaipanResult {
	shangYinYang := [3]int{YaoToYinYang(yaos[2]), YaoToYinYang(yaos[1]), YaoToYinYang(yaos[0])}
	xiaYinYang := [3]int{YaoToYinYang(yaos[5]), YaoToYinYang(yaos[4]), YaoToYinYang(yaos[3])}
	shangGua := Yaos2BaGuaNum(shangYinYang)
	xiaGua := Yaos2BaGuaNum(xiaYinYang)

	benGuaName := Guas64[shangGua-1][xiaGua-1]
	benGua := GuaInfo{Name: benGuaName, Yaos: yaos}

	huYaos := GetHuGuaYaos(yaos)
	huGua := GuaInfo{Name: Get64GuaNameByYaos(huYaos), Yaos: huYaos}

	bianYaos := GetBianGuaYaos(yaos, -1)
	bianGua := &GuaInfo{Name: Get64GuaNameByYaos(bianYaos), Yaos: bianYaos}

	benGuaEx := GetGuaEx(benGuaName)
	bianGuaExRaw := GetGuaEx(bianGua.Name)

	var finalBianGuaEx *GuaExData
	if benGuaEx != nil && bianGuaExRaw != nil {
		bianLiuQin := ReloadLiuqin(benGuaEx.WuXing, bianGuaExRaw.GanZhi)
		lqStr := ""
		for _, lq := range bianLiuQin {
			lqStr += lq
		}
		finalBianGuaEx = &GuaExData{
			Name:    bianGuaExRaw.Name,
			LiuQin:  lqStr,
			GanZhi:  bianGuaExRaw.GanZhi,
			Shi:     bianGuaExRaw.Shi,
			Ying:    bianGuaExRaw.Ying,
			FuCang:  bianGuaExRaw.FuCang,
			GuaShen: bianGuaExRaw.GuaShen,
			BaGong:  bianGuaExRaw.BaGong,
			WuXing:  bianGuaExRaw.WuXing,
			Kind:    bianGuaExRaw.Kind,
		}
	}

	var liuShen []string
	dayGanRunes := []rune(ctx.GanZhi[2])
	if len(dayGanRunes) >= 1 {
		liuShen = GetLiuShen(string(dayGanRunes[0]))
	}

	bianPos := -1
	for i, yao := range yaos {
		if yao == YaoLaoYang || yao == YaoLaoYin {
			bianPos = i
			break
		}
	}

	return &PaipanResult{
		Method:         "六爻手动起卦",
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
		BenGuaEx:       benGuaEx,
		BianGuaEx:      finalBianGuaEx,
		LiuShen:        liuShen,
	}
}
