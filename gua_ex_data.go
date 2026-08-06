// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

var guaExMap = map[string]*GuaExData{

	"乾为天": {
		Name:    "乾为天",
		LiuQin:  "父兄官父财子",
		GanZhi:  []string{"戌土", "申金", "午火", "辰土", "寅木", "子水"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 6,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "六冲",
	},

	"天风姤": {
		Name:    "天风姤",
		LiuQin:  "父兄官兄子父",
		GanZhi:  []string{"戌土", "申金", "午火", "酉金", "亥水", "丑土"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 5, Value: "寅木财"}},
		GuaShen: 3,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "",
	},

	"天山遁": {
		Name:    "天山遁",
		LiuQin:  "父兄官兄官父",
		GanZhi:  []string{"戌土", "申金", "午火", "申金", "午火", "辰土"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 5, Value: "寅木财"}, {Pos: 6, Value: "子水子"}},
		GuaShen: 0,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "",
	},

	"天地否": {
		Name:    "天地否",
		LiuQin:  "父兄官财官父",
		GanZhi:  []string{"戌土", "申金", "午火", "卯木", "巳火", "未土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 6, Value: "子水子"}},
		GuaShen: 2,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "六合",
	},

	"风地观": {
		Name:    "风地观",
		LiuQin:  "财官父财官父",
		GanZhi:  []string{"卯木", "巳火", "未土", "卯木", "巳火", "未土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 2, Value: "兄申金"}, {Pos: 6, Value: "子水子"}},
		GuaShen: 0,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "",
	},

	"山地剥": {
		Name:    "山地剥",
		LiuQin:  "财子父财官父",
		GanZhi:  []string{"寅木", "子水", "戌土", "卯木", "巳火", "未土"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 2, Value: "兄申金"}},
		GuaShen: 3,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "",
	},

	"火地晋": {
		Name:    "火地晋",
		LiuQin:  "官父兄财官父",
		GanZhi:  []string{"巳火", "未土", "酉金", "卯木", "巳火", "未土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 6, Value: "子子水"}},
		GuaShen: 4,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "游魂",
	},

	"火天大有": {
		Name:    "火天大有",
		LiuQin:  "官父兄父财子",
		GanZhi:  []string{"巳火", "未土", "酉金", "辰土", "寅木", "子水"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{},
		GuaShen: 5,
		BaGong:  "乾",
		WuXing:  "金",
		Kind:    "归魂",
	},

	"坎为水": {
		Name:    "坎为水",
		LiuQin:  "兄官父财官子",
		GanZhi:  []string{"子水", "戌土", "申金", "午火", "辰土", "寅木"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 0,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "六冲",
	},

	"水泽节": {
		Name:    "水泽节",
		LiuQin:  "兄官父官子财",
		GanZhi:  []string{"子水", "戌土", "申金", "丑土", "卯木", "巳火"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{},
		GuaShen: 1,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "六合",
	},

	"水雷屯": {
		Name:    "水雷屯",
		LiuQin:  "兄官父官子兄",
		GanZhi:  []string{"子水", "戌土", "申金", "辰土", "寅木", "子水"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 4, Value: "财午火"}},
		GuaShen: 0,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "",
	},

	"水火既济": {
		Name:    "水火既济",
		LiuQin:  "兄官父兄官子",
		GanZhi:  []string{"子水", "戌土", "申金", "亥水", "丑土", "卯木"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 4, Value: "财午火"}},
		GuaShen: 6,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "",
	},

	"泽火革": {
		Name:    "泽火革",
		LiuQin:  "官父兄兄官子",
		GanZhi:  []string{"未土", "酉金", "亥水", "亥水", "丑土", "卯木"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 4, Value: "财午火"}},
		GuaShen: 6,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "",
	},

	"雷火丰": {
		Name:    "雷火丰",
		LiuQin:  "官父财兄官子",
		GanZhi:  []string{"戌土", "申金", "午火", "亥水", "丑土", "卯木"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{},
		GuaShen: 6,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "",
	},

	"地火明夷": {
		Name:    "地火明夷",
		LiuQin:  "父兄官兄官子",
		GanZhi:  []string{"酉金", "亥水", "丑土", "亥水", "丑土", "卯木"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 4, Value: "财午火"}},
		GuaShen: 1,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "游魂",
	},

	"地水师": {
		Name:    "地水师",
		LiuQin:  "父兄官财官子",
		GanZhi:  []string{"酉金", "亥水", "丑土", "午火", "辰土", "寅木"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{},
		GuaShen: 3,
		BaGong:  "坎",
		WuXing:  "水",
		Kind:    "归魂",
	},

	"艮为山": {
		Name:    "艮为山",
		LiuQin:  "官财兄子父兄",
		GanZhi:  []string{"寅木", "子水", "戌土", "申金", "午火", "辰土"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 0,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "六冲",
	},

	"山火贲": {
		Name:    "山火贲",
		LiuQin:  "官财兄财兄官",
		GanZhi:  []string{"寅木", "子水", "戌土", "亥水", "丑土", "卯木"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 4, Value: "子申金"}, {Pos: 5, Value: "父午火"}},
		GuaShen: 2,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "六合",
	},

	"山天大畜": {
		Name:    "山天大畜",
		LiuQin:  "官财兄兄官财",
		GanZhi:  []string{"寅木", "子水", "戌土", "辰土", "寅木", "子水"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 4, Value: "子申金"}, {Pos: 5, Value: "父午火"}},
		GuaShen: 6,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "",
	},

	"山泽损": {
		Name:    "山泽损",
		LiuQin:  "官财兄兄官父",
		GanZhi:  []string{"寅木", "子水", "戌土", "丑土", "卯木", "巳火"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 4, Value: "子申金"}},
		GuaShen: 4,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "",
	},

	"火泽睽": {
		Name:    "火泽睽",
		LiuQin:  "父兄子兄官父",
		GanZhi:  []string{"巳火", "未土", "酉金", "丑土", "卯木", "巳火"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 2, Value: "财子水"}},
		GuaShen: 5,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "",
	},

	"天泽履": {
		Name:    "天泽履",
		LiuQin:  "兄子父兄官父",
		GanZhi:  []string{"戌土", "申金", "午火", "丑土", "卯木", "巳火"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 2, Value: "财子水"}},
		GuaShen: 6,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "",
	},

	"风泽中孚": {
		Name:    "风泽中孚",
		LiuQin:  "官父兄兄官父",
		GanZhi:  []string{"卯木", "巳火", "未土", "丑土", "卯木", "巳火"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 2, Value: "财子水"}, {Pos: 4, Value: "子申金"}},
		GuaShen: 0,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "游魂",
	},

	"风山渐": {
		Name:    "风山渐",
		LiuQin:  "官父兄子父兄",
		GanZhi:  []string{"卯木", "巳火", "未土", "申金", "午火", "辰土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 2, Value: "财子水"}},
		GuaShen: 1,
		BaGong:  "艮",
		WuXing:  "土",
		Kind:    "归魂",
	},

	"震为雷": {
		Name:    "震为雷",
		LiuQin:  "财官子财兄父",
		GanZhi:  []string{"戌土", "申金", "午火", "辰土", "寅木", "子水"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 0,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "六冲",
	},

	"雷地豫": {
		Name:    "雷地豫",
		LiuQin:  "财官子兄子财",
		GanZhi:  []string{"戌土", "申金", "午火", "卯木", "巳火", "未土"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 6, Value: "父子水"}},
		GuaShen: 4,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "六合",
	},

	"雷水解": {
		Name:    "雷水解",
		LiuQin:  "财官子子财兄",
		GanZhi:  []string{"戌土", "申金", "午火", "午火", "辰土", "寅木"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 6, Value: "父子水"}},
		GuaShen: 0,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "",
	},

	"雷风恒": {
		Name:    "雷风恒",
		LiuQin:  "财官子官父财",
		GanZhi:  []string{"戌土", "申金", "午火", "酉金", "亥水", "丑土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 5, Value: "兄寅木"}},
		GuaShen: 5,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "",
	},

	"地风升": {
		Name:    "地风升",
		LiuQin:  "官父财官父财",
		GanZhi:  []string{"酉金", "亥水", "丑土", "酉金", "亥水", "丑土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 3, Value: "子午火"}, {Pos: 5, Value: "兄寅木"}},
		GuaShen: 14,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "",
	},

	"水风井": {
		Name:    "水风井",
		LiuQin:  "父财官官父财",
		GanZhi:  []string{"子水", "戌土", "申金", "酉金", "亥水", "丑土"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 3, Value: "子午火"}, {Pos: 5, Value: "兄寅木"}},
		GuaShen: 4,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "",
	},

	"泽风大过": {
		Name:    "泽风大过",
		LiuQin:  "财官父官父财",
		GanZhi:  []string{"未土", "酉金", "亥水", "酉金", "亥水", "丑土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 3, Value: "子午火"}, {Pos: 5, Value: "兄寅木"}},
		GuaShen: 0,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "游魂",
	},

	"泽雷随": {
		Name:    "泽雷随",
		LiuQin:  "财官父财兄父",
		GanZhi:  []string{"未土", "酉金", "亥水", "辰土", "寅木", "子水"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 3, Value: "子午火"}},
		GuaShen: 2,
		BaGong:  "震",
		WuXing:  "木",
		Kind:    "归魂",
	},

	"巽为风": {
		Name:    "巽为风",
		LiuQin:  "兄子财官父财",
		GanZhi:  []string{"卯木", "巳火", "未土", "酉金", "亥水", "丑土"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 2,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "六冲",
	},

	"风天小畜": {
		Name:    "风天小畜",
		LiuQin:  "兄子财财兄父",
		GanZhi:  []string{"卯木", "巳火", "未土", "辰土", "寅木", "子水"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 4, Value: "官酉金"}},
		GuaShen: 6,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "",
	},

	"风火家人": {
		Name:    "风火家人",
		LiuQin:  "兄子财父财兄",
		GanZhi:  []string{"卯木", "巳火", "未土", "亥水", "丑土", "卯木"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 4, Value: "官酉金"}},
		GuaShen: 3,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "",
	},

	"风雷益": {
		Name:    "风雷益",
		LiuQin:  "兄子财财兄父",
		GanZhi:  []string{"卯木", "巳火", "未土", "辰土", "寅木", "子水"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 4, Value: "官酉金"}},
		GuaShen: 5,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "",
	},

	"天雷无妄": {
		Name:    "天雷无妄",
		LiuQin:  "财官子财兄父",
		GanZhi:  []string{"戌土", "申金", "午火", "辰土", "寅木", "子水"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{},
		GuaShen: 1,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "六冲",
	},

	"火雷噬嗑": {
		Name:    "火雷噬嗑",
		LiuQin:  "子财官财兄父",
		GanZhi:  []string{"巳火", "未土", "酉金", "辰土", "寅木", "子水"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{},
		GuaShen: 0,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "",
	},

	"山雷颐": {
		Name:    "山雷颐",
		LiuQin:  "兄父财财兄父",
		GanZhi:  []string{"寅木", "子水", "戌土", "辰土", "寅木", "子水"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 2, Value: "子巳火"}, {Pos: 4, Value: "官酉金"}},
		GuaShen: 4,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "游魂",
	},

	"山风蛊": {
		Name:    "山风蛊",
		LiuQin:  "兄父财官父财",
		GanZhi:  []string{"寅木", "子水", "戌土", "酉金", "亥水", "丑土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 2, Value: "子巳火"}},
		GuaShen: 1,
		BaGong:  "巽",
		WuXing:  "木",
		Kind:    "归魂",
	},

	"离为火": {
		Name:    "离为火",
		LiuQin:  "兄子财官子父",
		GanZhi:  []string{"巳火", "未土", "酉金", "亥水", "丑土", "卯木"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 1,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "六冲",
	},

	"火山旅": {
		Name:    "火山旅",
		LiuQin:  "兄子财财兄子",
		GanZhi:  []string{"巳火", "未土", "酉金", "申金", "午火", "辰土"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 4, Value: "官亥水"}, {Pos: 6, Value: "父卯木"}},
		GuaShen: 5,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "六合",
	},

	"火风鼎": {
		Name:    "火风鼎",
		LiuQin:  "兄子财财官子",
		GanZhi:  []string{"巳火", "未土", "酉金", "酉金", "亥水", "丑土"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{{Pos: 6, Value: "财卯木"}},
		GuaShen: 6,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "",
	},

	"火水未济": {
		Name:    "火水未济",
		LiuQin:  "兄子财兄子父",
		GanZhi:  []string{"巳火", "未土", "酉金", "午火", "辰土", "寅木"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 4, Value: "官亥水"}},
		GuaShen: 0,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "",
	},

	"山水蒙": {
		Name:    "山水蒙",
		LiuQin:  "父官子兄子父",
		GanZhi:  []string{"寅木", "子水", "戌土", "午火", "辰土", "寅木"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 3, Value: "财酉金"}},
		GuaShen: 3,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "",
	},

	"风水涣": {
		Name:    "风水涣",
		LiuQin:  "父兄子兄子父",
		GanZhi:  []string{"卯木", "巳火", "未土", "午火", "辰土", "寅木"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 3, Value: "财酉金"}, {Pos: 4, Value: "官亥水"}},
		GuaShen: 5,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "",
	},

	"天水讼": {
		Name:    "天水讼",
		LiuQin:  "子财兄兄子父",
		GanZhi:  []string{"戌土", "申金", "午火", "午火", "辰土", "寅木"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 4, Value: "官亥水"}},
		GuaShen: 6,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "游魂",
	},

	"天火同人": {
		Name:    "天火同人",
		LiuQin:  "子财兄官子父",
		GanZhi:  []string{"戌土", "申金", "午火", "亥水", "丑土", "卯木"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{},
		GuaShen: 6,
		BaGong:  "离",
		WuXing:  "火",
		Kind:    "归魂",
	},

	"坤为地": {
		Name:    "坤为地",
		LiuQin:  "子财兄官父兄",
		GanZhi:  []string{"酉金", "亥水", "丑土", "卯木", "巳火", "未土"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 2,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "六冲",
	},

	"地雷复": {
		Name:    "地雷复",
		LiuQin:  "子财兄兄官财",
		GanZhi:  []string{"酉金", "亥水", "丑土", "辰土", "寅木", "子水"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{{Pos: 5, Value: "父巳火"}},
		GuaShen: 6,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "六合",
	},

	"地泽临": {
		Name:    "地泽临",
		LiuQin:  "子财兄兄官父",
		GanZhi:  []string{"酉金", "亥水", "丑土", "丑土", "卯木", "巳火"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{},
		GuaShen: 34,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "",
	},

	"地天泰": {
		Name:    "地天泰",
		LiuQin:  "子财兄兄官财",
		GanZhi:  []string{"酉金", "亥水", "丑土", "辰土", "寅木", "子水"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 5, Value: "父巳火"}},
		GuaShen: 5,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "六合",
	},

	"雷天大壮": {
		Name:    "雷天大壮",
		LiuQin:  "兄子父兄官财",
		GanZhi:  []string{"戌土", "申金", "午火", "辰土", "寅木", "子水"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{},
		GuaShen: 4,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "六冲",
	},

	"泽天夬": {
		Name:    "泽天夬",
		LiuQin:  "兄子财兄官财",
		GanZhi:  []string{"未土", "酉金", "亥水", "辰土", "寅木", "子水"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 5, Value: "父巳火"}},
		GuaShen: 4,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "",
	},

	"水天需": {
		Name:    "水天需",
		LiuQin:  "财兄子兄官财",
		GanZhi:  []string{"子水", "戌土", "申金", "辰土", "寅木", "子水"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 5, Value: "父巳火"}},
		GuaShen: 1,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "游魂",
	},

	"水地比": {
		Name:    "水地比",
		LiuQin:  "财兄子官父兄",
		GanZhi:  []string{"子水", "戌土", "申金", "卯木", "巳火", "未土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{},
		GuaShen: 3,
		BaGong:  "坤",
		WuXing:  "土",
		Kind:    "归魂",
	},

	"兑为泽": {
		Name:    "兑为泽",
		LiuQin:  "父兄子父财官",
		GanZhi:  []string{"未土", "酉金", "亥水", "丑土", "卯木", "巳火"},
		Shi:     1,
		Ying:    4,
		FuCang:  []FuCang{},
		GuaShen: 3,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "六冲",
	},

	"泽水困": {
		Name:    "泽水困",
		LiuQin:  "父兄子官父财",
		GanZhi:  []string{"未土", "酉金", "亥水", "午火", "辰土", "寅木"},
		Shi:     6,
		Ying:    3,
		FuCang:  []FuCang{},
		GuaShen: 4,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "六合",
	},

	"泽地萃": {
		Name:    "泽地萃",
		LiuQin:  "父兄子财官父",
		GanZhi:  []string{"未土", "酉金", "亥水", "卯木", "巳火", "未土"},
		Shi:     5,
		Ying:    2,
		FuCang:  []FuCang{},
		GuaShen: 1,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "",
	},

	"泽山咸": {
		Name:    "泽山咸",
		LiuQin:  "父兄子兄官父",
		GanZhi:  []string{"未土", "酉金", "亥水", "申金", "午火", "辰土"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 5, Value: "财卯木"}},
		GuaShen: 6,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "",
	},

	"水山蹇": {
		Name:    "水山蹇",
		LiuQin:  "子父兄兄官父",
		GanZhi:  []string{"子水", "戌土", "申金", "申金", "午火", "辰土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 5, Value: "财卯木"}},
		GuaShen: 2,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "",
	},

	"地山谦": {
		Name:    "地山谦",
		LiuQin:  "兄子父兄官父",
		GanZhi:  []string{"酉金", "亥水", "丑土", "申金", "午火", "辰土"},
		Shi:     2,
		Ying:    5,
		FuCang:  []FuCang{{Pos: 5, Value: "财卯木"}},
		GuaShen: 0,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "",
	},

	"雷山小过": {
		Name:    "雷山小过",
		LiuQin:  "父兄官兄官父",
		GanZhi:  []string{"戌土", "申金", "午火", "申金", "午火", "辰土"},
		Shi:     3,
		Ying:    6,
		FuCang:  []FuCang{{Pos: 3, Value: "子亥水"}, {Pos: 5, Value: "财卯木"}},
		GuaShen: 5,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "游魂",
	},

	"雷泽归妹": {
		Name:    "雷泽归妹",
		LiuQin:  "父兄官父财官",
		GanZhi:  []string{"戌土", "申金", "午火", "丑土", "卯木", "巳火"},
		Shi:     4,
		Ying:    1,
		FuCang:  []FuCang{{Pos: 3, Value: "子亥水"}},
		GuaShen: 0,
		BaGong:  "兑",
		WuXing:  "金",
		Kind:    "归魂",
	},

}

// GetLiuShen 六神映射（根据日天干）。索引0=六爻, 5=初爻。
func GetLiuShen(dayGan string) []string {
	switch dayGan {
	case "甲", "乙":
		return []string{"玄武", "白虎", "螣蛇", "勾陈", "朱雀", "青龙"}
	case "丙", "丁":
		return []string{"青龙", "玄武", "白虎", "螣蛇", "勾陈", "朱雀"}
	case "戊":
		return []string{"朱雀", "青龙", "玄武", "白虎", "螣蛇", "勾陈"}
	case "己":
		return []string{"勾陈", "朱雀", "青龙", "玄武", "白虎", "螣蛇"}
	case "庚", "辛":
		return []string{"螣蛇", "勾陈", "朱雀", "青龙", "玄武", "白虎"}
	case "壬", "癸":
		return []string{"白虎", "螣蛇", "勾陈", "朱雀", "青龙", "玄武"}
	default:
		return nil
	}
}

// 地支五行映射
var diZhiWuXing = map[string]string{
	"子": "水", "丑": "土", "寅": "木", "卯": "木", "辰": "土", "巳": "火",
	"午": "火", "未": "土", "申": "金", "酉": "金", "戌": "土", "亥": "水",
}

// 六亲生成表
var genLiuqin = map[string]map[string]string{
	"木": {"水": "父", "木": "兄", "火": "子", "土": "财", "金": "官"},
	"火": {"水": "官", "木": "父", "火": "兄", "土": "子", "金": "财"},
	"土": {"水": "财", "木": "官", "火": "父", "土": "兄", "金": "子"},
	"金": {"水": "子", "木": "财", "火": "官", "土": "父", "金": "兄"},
	"水": {"水": "兄", "木": "子", "火": "财", "土": "官", "金": "父"},
}

// ReloadLiuqin 根据本卦五行动态计算六亲
func ReloadLiuqin(guaWuXing string, ganZhi []string) []string {
	result := make([]string, 6)
	wxMap, ok := genLiuqin[guaWuXing]
	if !ok {
		return result
	}
	for i := 0; i < 6 && i < len(ganZhi); i++ {
		if len(ganZhi[i]) == 0 {
			continue
		}
		dzRunes := []rune(ganZhi[i])
		dz := string(dzRunes[0])
		dzwx, ok2 := diZhiWuXing[dz]
		if !ok2 {
			continue
		}
		if lq, ok3 := wxMap[dzwx]; ok3 {
			result[i] = lq
		}
	}
	return result
}

// GetGuaEx 根据卦名获取扩展数据
func GetGuaEx(guaName string) *GuaExData {
	return guaExMap[guaName]
}

// GetAllGuaNames 获取所有64卦名
func GetAllGuaNames() []string {
	names := make([]string, 0, len(guaExMap))
	for name := range guaExMap {
		names = append(names, name)
	}
	return names
}
