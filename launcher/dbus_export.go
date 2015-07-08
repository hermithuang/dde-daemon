package launcher

import (
	. "pkg.deepin.io/dde/daemon/launcher/interfaces"
)

type ItemInfoExport struct {
	Path          string
	Name          string
	Id            ItemId
	Icon          string
	CategoryId    CategoryId
	TimeInstalled int64
}

func NewItemInfoExport(item ItemInfoInterface) ItemInfoExport {
	if item == nil {
		return ItemInfoExport{}
	}
	return ItemInfoExport{
		Path:          item.Path(),
		Name:          item.Name(),
		Id:            item.Id(),
		Icon:          item.Icon(),
		CategoryId:    item.GetCategoryId(),
		TimeInstalled: item.GetTimeInstalled(),
	}
}

type CategoryInfoExport struct {
	Name  string
	Id    CategoryId
	Items []ItemId
}

func NewCategoryInfoExport(c CategoryInfoInterface) CategoryInfoExport {
	if c == nil {
		return CategoryInfoExport{}
	}
	return CategoryInfoExport{
		Name:  c.Name(),
		Id:    c.Id(),
		Items: c.Items(),
	}
}

type FrequencyExport struct {
	Id        ItemId
	Frequency uint64
}

type TimeInstalledExport struct {
	Id   ItemId
	Time int64
}
