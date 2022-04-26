package search

import "strings"

//HasStickerByName returns if the weapon has a sticker if any of the stickers contains the provided name
func (s *SearchItem) HasStickerByName(name string) bool {
	for _, v := range s.AssetInfo.Info.Stickers {
		if strings.Contains(v.Name, name) {
			return true
		}
	}
	return false
}

//HasStickerById returns if the weapon has a sticker if any of the stickers ids match the given id
func (s *SearchItem) HasStickerById(id int) bool {
	for _, v := range s.AssetInfo.Info.Stickers {
		if v.StickerId == id {
			return true
		}
	}
	return false
}