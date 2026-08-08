// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

// 爻值常量（Yaos数组索引：0=六爻(上)，5=初爻(下)）
const (
	YaoShaoYang = 1 // 少阳（静爻）
	YaoShaoYin  = 2 // 少阴（静爻）
	YaoLaoYang  = 3 // 老阳（动爻→变少阴）
	YaoLaoYin   = 4 // 老阴（动爻→变少阳）
)

// FuCang 伏藏信息
type FuCang struct {
	Pos   int    `json:"pos"`   // 伏藏位置, 1=六爻, 6=初爻
	Value string `json:"value"` // 伏藏爻的值
}

// GuaExData 64卦扩展数据
type GuaExData struct {
	Name    string   `json:"name"`    // 卦名
	LiuQin  string   `json:"liuQin"`  // 六亲排列(6字)
	GanZhi  []string `json:"ganZhi"`  // 每爻干支(6个)
	Shi     int      `json:"shi"`     // 世爻位置(1=六爻,6=初爻)
	Ying    int      `json:"ying"`    // 应爻位置
	FuCang  []FuCang `json:"fuCang"`  // 伏藏列表
	GuaShen int      `json:"guaShen"` // 卦身位置(0=不存在)
	BaGong  string   `json:"baGong"`  // 八宫归属
	WuXing  string   `json:"wuXing"`  // 卦五行
	Kind    string   `json:"kind"`    // 类型(六冲/六合/游魂/归魂)
}

// GuaInfo 卦信息
type GuaInfo struct {
	Name string `json:"name"`
	Yaos [6]int `json:"yaos"` // 爻值: 1少阳 2少阴 3老阳 4老阴
	// 索引: [0]=六爻(上) [5]=初爻(下)
}

// DivinationContext 起卦上下文
type DivinationContext struct {
	GanZhi     [4]string `json:"ganZhi"`     // 四柱干支 [年柱, 月柱, 日柱, 时柱]
	XunKong    string    `json:"xunKong"`    // 旬空（如"戌亥"）
	LunarMonth int       `json:"lunarMonth"` // 农历月
	LunarDay   int       `json:"lunarDay"`   // 农历日
}

// PaipanResult 排盘结果
type PaipanResult struct {
	Method         string     `json:"method"`
	GanZhi         [4]string  `json:"ganZhi"`
	XunKong        string     `json:"xunKong"`
	BenGua         GuaInfo    `json:"benGua"`
	HuGua          GuaInfo    `json:"huGua"`
	BianGua        *GuaInfo   `json:"bianGua"`
	BianYao        int        `json:"bianYao"` // 首个动爻索引: 0=上爻,5=初爻,-1=无; 多动爻检查BenGua.Yaos中的3/4
	ShangGua       string     `json:"shangGua"`
	XiaGua         string     `json:"xiaGua"`
	ShangGuaWuXing string     `json:"shangGuaWuXing"`
	XiaGuaWuXing   string     `json:"xiaGuaWuXing"`
	BenGuaEx       *GuaExData `json:"benGuaEx"`
	BianGuaEx      *GuaExData `json:"bianGuaEx"`
	LiuShen        []string   `json:"liuShen"` // 六神6项，从上爻到初爻
}
