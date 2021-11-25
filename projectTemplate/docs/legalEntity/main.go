package legalEntity

import (
	t "github.com/tvitcom/nla_framework/types"
	"github.com/tvitcom/nla_framework/utils"
)

const (
	name           = "legal_entity"
	name_ru        = "юрЛица"
	name_ru_plural = "юрЛица"
	// файл с иконкой сохранить в директорию projectTemplate/sourceFiles/src/webClient/public/image/legal_entity_icon.svg
	// скачать можно с ресурса https://www.flaticon.com/
	menu_icon       = "image/legal_entity_icon.svg"
	breadcrumb_icon = "far fa-file-alt"
)

func GetDoc(project *t.ProjectType) t.DocType {
	doc := t.DocType{
		Project:    project,
		Name:       name,
		NameRu:     name_ru,
		PathPrefix: "docs",
		Flds: []t.FldType{
			t.GetFldTitle(),
			t.GetFldString("inn", "ИНН", 30, [][]int{{2, 1}}, "col-4"),
			t.GetFldString("kpp", "КПП", 30, [][]int{{2, 2}}, "col-4"),
			t.GetFldString("type", "тип организации", 50, [][]int{{2, 2}}),
			t.GetFldString("address_legal", "юр адрес", 0, [][]int{{3, 1}}, "col-8"),
		},
		Vue: t.DocVue{
			RouteName:      name,
			MenuIcon:       menu_icon,
			BreadcrumbIcon: breadcrumb_icon,
			Roles:          []string{},
		},
		//Templates:   map[string]*t.DocTemplate{"webClient_item.vue": {},},
		IsBaseTemplates: t.DocIsBaseTemplates{Vue: true, Sql: true},
		Sql: t.DocSql{
			IsSearchText:    true,
			IsBeforeTrigger: true,
			IsAfterTrigger:  true,
		},
	}
	// создаем стандартные методы sql "list", "update", "get_by_id" с возможностью ограничения по ролям
	doc.Sql.FillBaseMethods(doc.Name)

	doc.Vue.I18n = map[string]string{
		"listTitle":        utils.UpperCaseFirst(name_ru_plural),
		"listDeletedTitle": "Удаленные " + name_ru_plural,
	}

	doc.Init()

	return doc
}
