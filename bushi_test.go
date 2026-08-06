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
