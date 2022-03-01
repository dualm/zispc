package zispc

import "encoding/json"

type jsonProcessData struct {
	FactoryName          string `json:"FACTORYNAME"`
	ProductSpecName      string `json:"PRODUCTSPECNAME"`
	ProductFlowName      string `json:"PRODUCTFLOWNAME"`
	ProcessOperationName string `json:"PROCESSOPERATIONNAME"`
	MachineName          string `json:"MACHINENAME"`
	MachineRecipeName    string `json:"MACHINERECIPENAME"`
	UnitName             string `json:"UNITNAME"`
	LotName              string `json:"LOTNAME"`
	ProductName          string `json:"PRODUCTNAME"`
	ItemName             string `json:"ITEMNAME"`
	SiteList             []site `json:"SITELIST"`
}

type site struct {
	SiteName  string `json:"SITENAME"`
	SiteValue string `json:"SITEVALUE"`
}

func (data *jsonProcessData) Encode() ([]byte, error) {
	return json.MarshalIndent(data, "  ", "    ")
}

func NewJSONProcessData(machine, lot, recipe, factory, unit, product, spec, flow, operation, item string, sites map[string]string) ProcessData {
	return &jsonProcessData{
		FactoryName:          factory,
		ProductSpecName:      spec,
		ProductFlowName:      flow,
		ProcessOperationName: operation,
		MachineName:          machine,
		MachineRecipeName:    recipe,
		UnitName:             unit,
		LotName:              lot,
		ProductName:          product,
		ItemName:             item,
		SiteList:             makeSiteList(sites),
	}
}

type jsonProcessDataMulti struct {
	FactoryName          string      `json:"FACTORYNAME"`
	ProductSpecName      string      `json:"PRODUCTSPECNAME"`
	ProductFlowName      string      `json:"PRODUCTFLOWNAME"`
	ProcessOperationName string      `json:"PROCESSOPERATIONNAME"`
	MachineName          string      `json:"MACHINENAME"`
	MachineRecipeName    string      `json:"MACHINERECIPENAME"`
	UnitName             string      `json:"UNITNAME"`
	LotName              string      `json:"LOTNAME"`
	ProductName          string      `json:"PRODUCTNAME"`
	SiteList             []JSONSites `json:"SITELIST"`
}

type JSONSites struct {
	ItemName string `json:"ITEMNAME"`
	Sites    []site `json:"SITES"`
}

func (data *jsonProcessDataMulti) Encode() ([]byte, error) {
	return json.MarshalIndent(data, "  ", "    ")
}

func NewJSONProcessDataMulti(machine, lot, recipe, factory, unit, product, spec, flow, operation string, sites []JSONSites) ProcessData {
	return &jsonProcessDataMulti{
		FactoryName:          factory,
		ProductSpecName:      spec,
		ProductFlowName:      flow,
		ProcessOperationName: operation,
		MachineName:          machine,
		MachineRecipeName:    recipe,
		UnitName:             unit,
		LotName:              lot,
		ProductName:          product,
		SiteList:             sites,
	}
}

func AddSite(s []JSONSites, item string, sites map[string]string) []JSONSites {
	s = append(s, JSONSites{
		ItemName: item,
		Sites:    makeSiteList(sites),
	})

	return s
}

func makeSiteList(sites map[string]string) []site {
	re := make([]site, 0, len(sites))
	for k, v := range sites {
		re = append(re, site{
			SiteName:  k,
			SiteValue: v,
		})
	}

	return re
}
