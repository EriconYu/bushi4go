// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

// GuaFullInfo 卦的完整信息（含排盘数据 + 卦爻辞）
type GuaFullInfo struct {
	Name  string     `json:"name"`
	Yaos  [6]int     `json:"yaos"` // index 0=上爻（六爻），index 5=初爻
	Ex    *GuaExData `json:"ex"`
	Texts *GuaTexts  `json:"texts"`
}

// Get64GuaList 获取全部64卦完整数据（含卦画、装卦和卦爻辞）。
// 返回顺序稳定：按上卦1...8、下卦1...8排列。
func Get64GuaList() []GuaFullInfo {
	initGuaTexts()
	result := make([]GuaFullInfo, 0, 64)
	for shangGua := 1; shangGua <= 8; shangGua++ {
		for xiaGua := 1; xiaGua <= 8; xiaGua++ {
			name := Guas64[shangGua-1][xiaGua-1]
			info := buildGuaFullInfo(name, shangGua, xiaGua)
			result = append(result, info)
		}
	}
	return result
}

// GetGuaByName 根据卦名获取单卦完整数据（含卦爻辞）
func GetGuaByName(name string) (*GuaFullInfo, error) {
	for shangGua := 1; shangGua <= 8; shangGua++ {
		for xiaGua := 1; xiaGua <= 8; xiaGua++ {
			if Guas64[shangGua-1][xiaGua-1] == name {
				info := buildGuaFullInfo(name, shangGua, xiaGua)
				return &info, nil
			}
		}
	}
	return nil, nil
}

func buildGuaFullInfo(name string, shangGua, xiaGua int) GuaFullInfo {
	shang := BaGuaNum2Yaos(shangGua)
	xia := BaGuaNum2Yaos(xiaGua)
	info := GuaFullInfo{
		Name: name,
		Yaos: [6]int{shang[2], shang[1], shang[0], xia[2], xia[1], xia[0]},
		Ex:   GetGuaEx(name),
	}
	if texts, ok := guaTextsCache[name]; ok {
		info.Texts = texts
	}
	return info
}
