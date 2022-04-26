package search

import "strings"

func (s *SearchItem) HasStickerByName(name string) bool {
	for _, v := range s.AssetInfo.Info.Stickers {
		if strings.Contains(v.Name, name) {
			return true
		}
	}
	return false
}

func (s *SearchItem) HasStickerById(id int) bool {
	for _, v := range s.AssetInfo.Info.Stickers {
		if v.StickerId == id {
			return true
		}
	}
	return false
}