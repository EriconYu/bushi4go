// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

import (
	"testing"
	"time"
)

func TestGuas64(t *testing.T) {
	if Guas64[0][0] != "乾为天" {
		t.Error("期望乾为天")
	}
	if Guas64[7][7] != "坤为地" {
		t.Error("期望坤为地")
	}
	if Guas64[7][5] != "地水师" {
		t.Error("期望地水师")
	}
}

func TestGuaExMap(t *testing.T) {
	if len(guaExMap) != 64 {
		t.Errorf("期望64卦，实际%d", len(guaExMap))
	}
}

func TestLiuYaoShiJianQiGua(t *testing.T) {
	ctx := BuildContext(time.Date(2026, 4, 30, 10, 0, 0, 0, time.Local))
	result := LiuYaoShiJianQiGua(ctx)
	if result.BenGua.Name != "地水师" {
		t.Errorf("期望本卦地水师，实际%s", result.BenGua.Name)
	}
	if result.BianGua.Name != "山水蒙" {
		t.Errorf("期望变卦山水蒙，实际%s", result.BianGua.Name)
	}
}

func TestMeiHuaShiJianQiGua(t *testing.T) {
	ctx := BuildContext(time.Date(2026, 4, 29, 19, 0, 0, 0, time.Local))
	result := MeiHuaShiJianQiGua(ctx)
	if result.BenGua.Name != "山泽损" {
		t.Errorf("期望本卦山泽损，实际%s", result.BenGua.Name)
	}
	if result.BianGua.Name != "火泽睽" {
		t.Errorf("期望变卦火泽睽，实际%s", result.BianGua.Name)
	}
}

func TestZiHour(t *testing.T) {
	ctx23 := BuildContext(time.Date(2026, 4, 30, 23, 0, 0, 0, time.Local))
	ctxNextDay := BuildContext(time.Date(2026, 5, 1, 0, 0, 0, 0, time.Local))
	if ctx23.GanZhi[2] != ctxNextDay.GanZhi[2] {
		t.Errorf("晚子时日柱应一致: %s vs %s", ctx23.GanZhi[2], ctxNextDay.GanZhi[2])
	}
}

func TestBianGua(t *testing.T) {
	yaos := [6]int{YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang}
	yaos[0] += 2 // 标记六爻为老阳
	bian := GetBianGuaYaos(yaos, 0)
	if bian[0] != YaoShaoYin {
		t.Errorf("老阳应变少阴，实际%d", bian[0])
	}
}

func TestShouYaoQiGua(t *testing.T) {
	ctx := BuildContext(time.Date(2026, 4, 30, 10, 0, 0, 0, time.Local))
	yaos := [6]int{YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang, YaoShaoYang}
	result := LiuYaoShouYaoQiGua(ctx, yaos)
	if result.BenGua.Name != "乾为天" {
		t.Errorf("期望乾为天，实际%s", result.BenGua.Name)
	}
}

func TestStandardLiuYaoLayout(t *testing.T) {
	ctx := DivinationContext{
		GanZhi:  [4]string{"丙午", "丙申", "甲寅", "乙亥"},
		XunKong: "子丑", LunarMonth: 6, LunarDay: 26,
	}
	result := LiuYaoShiJianQiGua(ctx)

	if result.BenGua.Name != "山火贲" || result.BenGua.Yaos != [6]int{1, 2, 2, 3, 2, 1} {
		t.Fatalf("本卦装卦错误: %+v", result.BenGua)
	}
	if result.BianGua.Name != "山雷颐" || result.BianGua.Yaos != [6]int{1, 2, 2, 2, 2, 1} {
		t.Fatalf("变卦错误: %+v", result.BianGua)
	}
	expectStrings(t, result.LiuShen, []string{"玄武", "白虎", "螣蛇", "勾陈", "朱雀", "青龙"})

	ben := result.BenGuaEx
	if ben.LiuQin != "官财兄财兄官" || ben.Shi != 6 || ben.Ying != 3 {
		t.Fatalf("本卦六亲世应错误: %+v", ben)
	}
	expectStrings(t, ben.GanZhi, []string{"寅木", "子水", "戌土", "亥水", "丑土", "卯木"})
	if len(ben.FuCang) != 2 || ben.FuCang[0] != (FuCang{Pos: 4, Value: "子申金"}) || ben.FuCang[1] != (FuCang{Pos: 5, Value: "父午火"}) {
		t.Fatalf("本卦伏藏错误: %+v", ben.FuCang)
	}

	bian := result.BianGuaEx
	if bian.LiuQin != "官财兄兄官财" || bian.Shi != 3 || bian.Ying != 6 {
		t.Fatalf("变卦六亲世应错误: %+v", bian)
	}
	if len(bian.FuCang) != 2 || bian.FuCang[0] != (FuCang{Pos: 2, Value: "子巳火"}) || bian.FuCang[1] != (FuCang{Pos: 4, Value: "官酉金"}) {
		t.Fatalf("变卦伏藏错误: %+v", bian.FuCang)
	}
}

func TestManualYaoDirectionAndMultipleChanges(t *testing.T) {
	ctx := DivinationContext{GanZhi: [4]string{"丙午", "丙申", "甲寅", "乙亥"}, XunKong: "子丑"}
	bottomUp := [6]int{YaoShaoYang, YaoShaoYin, YaoLaoYang, YaoShaoYin, YaoShaoYin, YaoShaoYang}
	topDown := [6]int{bottomUp[5], bottomUp[4], bottomUp[3], bottomUp[2], bottomUp[1], bottomUp[0]}
	standard := LiuYaoShouYaoQiGua(ctx, topDown)
	if standard.BenGua.Name != "山火贲" || standard.BianGua.Name != "山雷颐" {
		t.Fatalf("手摇方向错误: %s -> %s", standard.BenGua.Name, standard.BianGua.Name)
	}

	multiple := LiuYaoShouYaoQiGua(ctx, [6]int{YaoLaoYang, 2, 2, YaoLaoYang, 2, 1})
	if multiple.BianYao != 0 || multiple.BianGua.Yaos[0] != YaoShaoYin || multiple.BianGua.Yaos[3] != YaoShaoYin {
		t.Fatalf("多动爻变换错误: %+v", multiple)
	}
}

func TestNumberDivinationUsesExplicitMovingLine(t *testing.T) {
	ctx := DivinationContext{
		GanZhi:  [4]string{"丙午", "丙申", "甲寅", "乙亥"},
		XunKong: "子丑", LunarMonth: 6, LunarDay: 26,
	}
	for yaoNumber := 1; yaoNumber <= 6; yaoNumber++ {
		expectedIndex := 6 - yaoNumber
		if got := MovingLineIndex(yaoNumber); got != expectedIndex {
			t.Fatalf("爻位 %d 应映射到索引 %d，实际 %d", yaoNumber, expectedIndex, got)
		}
		liuYao := LiuYaoShuZiQiGua(ctx, [3]int{1, 1, yaoNumber})
		if liuYao.BianYao != expectedIndex {
			t.Fatalf("六爻数字起卦爻位 %d 错误: %d", yaoNumber, liuYao.BianYao)
		}
		movingCount := 0
		for index, yao := range liuYao.BenGua.Yaos {
			if yao == YaoLaoYang || yao == YaoLaoYin {
				movingCount++
				if index != expectedIndex {
					t.Fatalf("动爻落在索引 %d，预期 %d", index, expectedIndex)
				}
			}
		}
		if movingCount != 1 {
			t.Fatalf("应只有一个动爻，实际 %d", movingCount)
		}
		if got := MeiHuaShuZiQiGua(ctx, [3]int{1, 1, yaoNumber}).BianYao; got != expectedIndex {
			t.Fatalf("梅花数字起卦爻位 %d 错误: %d", yaoNumber, got)
		}
	}
	if got := LiuYaoShuZiQiGua(ctx, [3]int{1, 1, 1}).BianGua.Name; got != "天风姤" {
		t.Fatalf("乾为天初爻动应变天风姤，实际 %s", got)
	}
}

func TestRandomAndManualMovingLineInvariants(t *testing.T) {
	ctx := DivinationContext{
		GanZhi:  [4]string{"丙午", "丙申", "甲寅", "乙亥"},
		XunKong: "子丑", LunarMonth: 6, LunarDay: 26,
	}
	for i := 0; i < 20; i++ {
		result := LiuYaoSuiJiQiGua(ctx)
		if result.BianYao < 0 || result.BianYao > 5 {
			t.Fatalf("随机动爻索引越界: %d", result.BianYao)
		}
		movingCount := 0
		for _, yao := range result.BenGua.Yaos {
			if yao == YaoLaoYang || yao == YaoLaoYin {
				movingCount++
			}
		}
		if movingCount != 1 {
			t.Fatalf("随机起卦应只有一个动爻，实际 %d", movingCount)
		}
	}
	static := MeiHuaShouYaoQiGua(ctx, [6]int{1, 1, 1, 1, 1, 1})
	if static.BianYao != -1 {
		t.Fatalf("无老阴老阳时不应伪造动爻: %d", static.BianYao)
	}
}

func TestGuaListAndTexts(t *testing.T) {
	list := Get64GuaList()
	if len(list) != 64 || list[0].Name != "乾为天" || list[63].Name != "坤为地" {
		t.Fatalf("64卦列表顺序错误: len=%d first=%s last=%s", len(list), list[0].Name, list[63].Name)
	}
	if list[0].Yaos != [6]int{1, 1, 1, 1, 1, 1} || list[0].Texts == nil {
		t.Fatalf("64卦完整数据缺失: %+v", list[0])
	}
	gua, err := GetGuaByName("山火贲")
	if err != nil || gua.Yaos != [6]int{1, 2, 2, 1, 2, 1} {
		t.Fatalf("单卦数据错误: gua=%+v err=%v", gua, err)
	}
	missing, err := GetGuaByName("不存在")
	if err != nil || missing != nil {
		t.Fatalf("不存在的卦名应返回nil: gua=%+v err=%v", missing, err)
	}
}

func expectStrings(t *testing.T, actual, expected []string) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("长度错误: actual=%v expected=%v", actual, expected)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("第%d项错误: actual=%v expected=%v", i, actual, expected)
		}
	}
}
