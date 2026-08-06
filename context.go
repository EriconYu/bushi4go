package bushi

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// ZiHourCalculationTime 晚子时换日：23:00~00:00 归入次日子时。
func ZiHourCalculationTime(t time.Time) time.Time {
	if t.Hour() == 23 {
		return t.Add(1 * time.Hour)
	}
	return t
}

// BuildContext 从公历时间构建排盘上下文。
// 23:00~00:00（晚子时）自动按次日子时排盘。
func BuildContext(t time.Time) DivinationContext {
	calcTime := ZiHourCalculationTime(t)

	lunar := calendar.NewLunarFromDate(calcTime)

	ganZhi := [4]string{
		lunar.GetYearInGanZhi(),
		lunar.GetMonthInGanZhi(),
		lunar.GetDayInGanZhi(),
		lunar.GetTimeInGanZhi(),
	}

	// 旬空
	xunKong := getXunKong(ganZhi[2])

	// 农历月日（取绝对值，闰月标记不影响起卦计算）
	lunarMonth := lunar.GetMonth()
	if lunarMonth < 0 {
		lunarMonth = -lunarMonth
	}
	lunarDay := lunar.GetDay()

	return DivinationContext{
		GanZhi:     ganZhi,
		XunKong:    xunKong,
		LunarMonth: lunarMonth,
		LunarDay:   lunarDay,
	}
}

// 六十甲子表
var jiaZiTable = [60]string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
}

var xunKongTable = [6]string{"戌亥", "申酉", "午未", "辰巳", "寅卯", "子丑"}

func getXunKong(dayGanZhi string) string {
	for i, v := range jiaZiTable {
		if v == dayGanZhi {
			return xunKongTable[i/10]
		}
	}
	return ""
}
