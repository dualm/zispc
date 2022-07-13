package zispc

import (
	"encoding/xml"
	"fmt"
)

const (
	emptyItem = "-"
)

type Header struct {
	XMLName                   xml.Name `xml:"Header"`
	MessageName               string   `xml:"MESSAGENAME"`
	EventComment              string   `xml:"EVENTCOMMENT"`
	EventUser                 string   `xml:"EVENTUSER"`
	OriginalSourceSubjectName string   `xml:"ORIGINALSOURCESUBJECTNAME"`
}

func getHeader(headerLength int, messageName ...string) Header {
	_name := "DataCollectRequest"

	if len(messageName) > 0 {
		_name = messageName[0]
	}

	return Header{
		MessageName:               _name,
		EventComment:              _name,
		EventUser:                 _name,
		OriginalSourceSubjectName: makeSubjectName(headerLength),
	}
}

func makeSubjectName(headerLength int) string {
	return "_INBOX." + getRandStr(headerLength) + getTransactionId()
}

type XmlProcessData struct {
	XMLName              xml.Name  `xml:"Message"`
	Header               Header    `xml:"Header"`
	FactoryName          string    `xml:"Body>FACTORYNAME"`
	ProductSpecName      string    `xml:"Body>PRODUCTSPECNAME"`
	ProcessFlowName      string    `xml:"Body>PROCESSFLOWNAME"`
	ProcessOperationName string    `xml:"Body>PROCESSOPERATIONNAME"`
	MachineName          string    `xml:"Body>MACHINENAME"`
	MachineRecipeName    string    `xml:"Body>MACHINERECIPENAME"`
	UnitName             string    `xml:"Body>UNITNAME"`
	LotName              string    `xml:"Body>LOTNAME"`
	ProductName          string    `xml:"Body>PRODUCTNAME"`
	ItemList             []XMLItem `xml:"Body>ITEMLIST>ITEM"`
}

func (data *XmlProcessData) Encode() ([]byte, error) {
	var _field string

	switch {
	case data.MachineName == emptyItem:
		_field = "MachineName"
	case data.LotName == emptyItem:
		_field = "ProductName"
	default:
		return xml.MarshalIndent(data, "  ", "    ")
	}

	return nil, fmt.Errorf("encode spc message error, invalid %s", _field)
}

type XMLItem struct {
	XMLName  xml.Name `xml:"ITEM"`
	ItemName string   `xml:"ITEMNAME"`
	SiteList []dvSite `xml:"SITELIST>SITE"`
}

type dvSite struct {
	XMLName            xml.Name `xml:"SITE"`
	SiteName           string   `xml:"SITENAME"` // index，从1开始
	SiteValue          string   `xml:"SITEVALUE"`
	SampleMaterialName string   `xml:"SAMPLEMATERIALNAME"` // 产品名称
}

func AddItemToXML(list []XMLItem, name, value, materialName string, sites map[string]string) []XMLItem {
	if sites == nil {
		list = append(list, XMLItem{
			ItemName: name,
			SiteList: []dvSite{
				{
					SiteName:           OnlySiteName(),
					SiteValue:          value,
					SampleMaterialName: materialName,
				},
			},
		})
	} else {
		siteList := make([]dvSite, 0, len(sites))
		for k, v := range sites {
			siteList = append(siteList, dvSite{
				SiteName:           checkSiteName(k),
				SiteValue:          v,
				SampleMaterialName: materialName,
			})
		}
		list = append(list, XMLItem{
			ItemName: name,
			SiteList: siteList,
		})
	}

	return list
}

func NewXMLProcessData(headerCount int, machine, recipe, unit, spec, flow, lot, product, factory, operation string,
	dv map[string]string, dvItems []XMLItem, messageName ...string) ProcessData {
	header := getHeader(headerCount, messageName...)

	items := make([]XMLItem, 0)

	if dvItems == nil {
		for key, val := range dv {
			items = AddItemToXML(items, key, val, makeEmpty(product), nil)
		}
	} else {
		items = dvItems
	}

	processData := XmlProcessData{
		Header:               header,
		FactoryName:          makeEmpty(factory),
		ProductSpecName:      makeEmpty(spec),
		ProcessFlowName:      makeEmpty(flow),
		ProcessOperationName: makeEmpty(operation),
		MachineName:          makeEmpty(machine),
		MachineRecipeName:    makeEmpty(recipe),
		UnitName:             makeEmpty(unit),
		LotName:              makeEmpty(lot),
		ProductName:          makeEmpty(product),
		ItemList:             items,
	}

	return &processData
}

func makeEmpty(s string) string {
	if s == "" {
		return emptyItem
	}

	return s
}
