package bushi

// GetHuGuaYaos 根据本卦六爻计算互卦
func GetHuGuaYaos(benYaos [6]int) [6]int {
	var hu [6]int
	hu[0] = YaoToYinYang(benYaos[1])
	hu[1] = YaoToYinYang(benYaos[2])
	hu[2] = YaoToYinYang(benYaos[3])
	hu[3] = YaoToYinYang(benYaos[2])
	hu[4] = YaoToYinYang(benYaos[3])
	hu[5] = YaoToYinYang(benYaos[4])
	return hu
}

// GetBianGuaYaos 根据本卦六爻计算变卦
func GetBianGuaYaos(benYaos [6]int, pos int) [6]int {
	var bian [6]int
	for i := 0; i < 6; i++ {
		bian[i] = benYaos[i]
		isMovingYao := bian[i] == YaoLaoYang || bian[i] == YaoLaoYin
		isThisMoving := (pos == -1 && isMovingYao) || (pos >= 0 && i == pos && isMovingYao)
		if isThisMoving {
			if bian[i] == YaoLaoYang {
				bian[i] = YaoShaoYin
			} else {
				bian[i] = YaoShaoYang
			}
		} else {
			bian[i] = YaoToYinYang(bian[i])
		}
	}
	return bian
}

// Get64GuaNameByYaos 根据六爻计算64卦名
func Get64GuaNameByYaos(yaos [6]int) string {
	yaosTemp := [6]int{}
	for i := 0; i < 6; i++ {
		yaosTemp[i] = YaoToYinYang(yaos[i])
	}
	shangYaos := [3]int{yaosTemp[2], yaosTemp[1], yaosTemp[0]}
	xiaYaos := [3]int{yaosTemp[5], yaosTemp[4], yaosTemp[3]}
	shangGua := Yaos2BaGuaNum(shangYaos)
	xiaGua := Yaos2BaGuaNum(xiaYaos)
	if shangGua < 1 || shangGua > 8 || xiaGua < 1 || xiaGua > 8 {
		return ""
	}
	return Guas64[shangGua-1][xiaGua-1]
}
