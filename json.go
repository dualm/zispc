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
	SampleMaterialName string `json:"SAMPLEMATERIALNAME,omitempty"`
	SiteName           string `json:"SITENAME"`
	SiteValue          string `json:"SITEVALUE"`
}

type Sites struct {
	ItemName string `json:"ITEMNAME"`
	Sites    []site `json:"SITES"`
}

func (data *jsonProcessData) Encode() ([]byte, error) {
	return json.MarshalIndent(data, "  ", "    ")
}

func NewJSONProcessData(machine, lot, recipe, factory, unit, product, spec, flow, operation, item string, sites map[string]string) ProcessData {
	return &jsonProcessData{
		FactoryName:          makeEmpty(factory),
		ProductSpecName:      makeEmpty(spec),
		ProductFlowName:      makeEmpty(flow),
		ProcessOperationName: makeEmpty(operation),
		MachineName:          makeEmpty(machine),
		MachineRecipeName:    makeEmpty(recipe),
		UnitName:             makeEmpty(unit),
		LotName:              makeEmpty(lot),
		ProductName:          makeEmpty(product),
		ItemName:             makeEmpty(item),
		SiteList:             makeSiteList("", sites),
	}
}

type jsonProcessDataMulti struct {
	FactoryName          string  `json:"FACTORYNAME"`
	ProductSpecName      string  `json:"PRODUCTSPECNAME"`
	ProductFlowName      string  `json:"PRODUCTFLOWNAME"`
	ProcessOperationName string  `json:"PROCESSOPERATIONNAME"`
	MachineName          string  `json:"MACHINENAME"`
	MachineRecipeName    string  `json:"MACHINERECIPENAME"`
	UnitName             string  `json:"UNITNAME"`
	LotName              string  `json:"LOTNAME"`
	ProductName          string  `json:"PRODUCTNAME"`
	SiteList             []Sites `json:"SITELIST"`
}

func (data *jsonProcessDataMulti) Encode() ([]byte, error) {
	return json.MarshalIndent(data, "  ", "    ")
}

func NewJSONProcessDataMulti(machine, lot, recipe, factory, unit, product, spec, flow, operation string, sites []Sites) ProcessData {
	return &jsonProcessDataMulti{
		FactoryName:          makeEmpty(factory),
		ProductSpecName:      makeEmpty(spec),
		ProductFlowName:      makeEmpty(flow),
		ProcessOperationName: makeEmpty(operation),
		MachineName:          makeEmpty(machine),
		MachineRecipeName:    makeEmpty(recipe),
		UnitName:             makeEmpty(unit),
		LotName:              makeEmpty(lot),
		ProductName:          makeEmpty(product),
		SiteList:             sites,
	}
}

func AddSite(s []Sites, item, sampleName string, sites map[string]string) []Sites {
	if s == nil {
		s = make([]Sites, 0, 1)
	}

	return append(s, makeSites(item, sampleName, sites))
}

func makeSiteList(sampleName string, sites map[string]string) []site {
	re := make([]site, 0, len(sites))
	for k, v := range sites {
		re = append(re, site{
			SampleMaterialName: makeEmpty(sampleName),
			SiteName:           k,
			SiteValue:          v,
		})
	}

	return re
}

func makeSites(item, sampleName string, sites map[string]string) Sites {
	return Sites{
		ItemName: makeEmpty(item),
		Sites:    makeSiteList(sampleName, sites),
	}
}
