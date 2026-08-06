// Copyright (c) 2026 不惑之心 (buhuo.xin)
// 作者：净志 | 微信：haitaojingzhi | 官网：https://www.buhuo.xin
// 仅开源排盘算法，不含卦爻辞、解卦等内容。

package bushi

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed data/gua_texts.json
var guaTextsJSON []byte

// GuaTexts 卦爻辞与解卦文本
type GuaTexts struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	YaoCi  string `json:"yaoCi"`  // 卦爻辞（原文、译文、注释）
	JieXi  string `json:"jieXi"`  // 解卦（白话文、邵雍、傅佩荣、张铭仁等解析）
}

var guaTextsCache map[string]*GuaTexts

func initGuaTexts() {
	if guaTextsCache != nil {
		return
	}
	guaTextsCache = make(map[string]*GuaTexts)
	json.Unmarshal(guaTextsJSON, &guaTextsCache)
}

// GetGuaTexts 根据卦名获取卦爻辞与解卦文本
func GetGuaTexts(guaName string) (*GuaTexts, error) {
	initGuaTexts()
	gt, ok := guaTextsCache[guaName]
	if !ok {
		return nil, fmt.Errorf("未找到卦文本: %s", guaName)
	}
	return gt, nil
}

// GetAllGuaTexts 获取全部64卦文本数据
func GetAllGuaTexts() map[string]*GuaTexts {
	initGuaTexts()
	return guaTextsCache
}
