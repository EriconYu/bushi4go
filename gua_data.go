package bushi

// Guas64 64卦表 [上卦-1][下卦-1] = 卦名
var Guas64 = [8][8]string{
	{"乾为天", "天泽履", "天火同人", "天雷无妄", "天风姤", "天水讼", "天山遁", "天地否"},
	{"泽天夬", "兑为泽", "泽火革", "泽雷随", "泽风大过", "泽水困", "泽山咸", "泽地萃"},
	{"火天大有", "火泽睽", "离为火", "火雷噬嗑", "火风鼎", "火水未济", "火山旅", "火地晋"},
	{"雷天大壮", "雷泽归妹", "雷火丰", "震为雷", "雷风恒", "雷水解", "雷山小过", "雷地豫"},
	{"风天小畜", "风泽中孚", "风火家人", "风雷益", "巽为风", "风水涣", "风山渐", "风地观"},
	{"水天需", "水泽节", "水火既济", "水雷屯", "水风井", "坎为水", "水山蹇", "水地比"},
	{"山天大畜", "山泽损", "山火贲", "山雷颐", "山风蛊", "山水蒙", "艮为山", "山地剥"},
	{"地天泰", "地泽临", "地火明夷", "地雷复", "地风升", "地水师", "地山谦", "坤为地"},
}

// BaGuaNames 八卦名称
var BaGuaNames = [8]string{"乾", "兑", "离", "震", "巽", "坎", "艮", "坤"}

// BaGuaWuXing 八卦五行
var BaGuaWuXing = [8]string{"金", "金", "火", "木", "木", "水", "土", "土"}

// DiZhi12 12地支
var DiZhi12 = [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

// BaGuaNum2Yaos 八卦序号转爻数组（返回 [初爻, 二爻, 三爻] bottom→top）
func BaGuaNum2Yaos(num int) [3]int {
	switch num {
	case 1:
		return [3]int{1, 1, 1} // 乾
	case 2:
		return [3]int{1, 1, 2} // 兑
	case 3:
		return [3]int{1, 2, 1} // 离
	case 4:
		return [3]int{1, 2, 2} // 震
	case 5:
		return [3]int{2, 1, 1} // 巽
	case 6:
		return [3]int{2, 1, 2} // 坎
	case 7:
		return [3]int{2, 2, 1} // 艮
	case 8:
		return [3]int{2, 2, 2} // 坤
	default:
		return [3]int{0, 0, 0}
	}
}

// Yaos2BaGuaNum 爻数组转八卦序号
func Yaos2BaGuaNum(yaos [3]int) int {
	sum := yaos[0]*100 + yaos[1]*10 + yaos[2]
	switch sum {
	case 111:
		return 1 // 乾
	case 112:
		return 2 // 兑
	case 121:
		return 3 // 离
	case 122:
		return 4 // 震
	case 211:
		return 5 // 巽
	case 212:
		return 6 // 坎
	case 221:
		return 7 // 艮
	case 222:
		return 8 // 坤
	default:
		return 0
	}
}

// YaoToYinYang 爻值转阴阳值（3→1, 4→2, 其他不变）
func YaoToYinYang(yao int) int {
	if yao > 2 {
		return yao - 2
	}
	return yao
}

// GuaNum2Name 八卦序号转卦名
func GuaNum2Name(num int) string {
	if num < 1 || num > 8 {
		return ""
	}
	return BaGuaNames[num-1]
}

// GuaNum2WuXing 八卦序号转五行
func GuaNum2WuXing(num int) string {
	if num < 1 || num > 8 {
		return ""
	}
	return BaGuaWuXing[num-1]
}

// GetZhiIndex 获取地支序号(1-12)
func GetZhiIndex(zhi string) int {
	for i := 0; i < 12; i++ {
		if DiZhi12[i] == zhi {
			return i + 1
		}
	}
	return 0
}
