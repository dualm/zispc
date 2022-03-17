package zispc

import (
	"reflect"
	"testing"
)

func TestNewProcessData(t *testing.T) {
	type args struct {
		headerCount int
		machine     string
		recipe      string
		unit        string
		spec        string
		flow        string
		lot         string
		product     string
		factory     string
		operation   string
		dv          map[string]string
		dvItems     []XMLItem
	}

	tests := []struct {
		name string
		args args
		want *xmlProcessData
	}{
		{
			name: "1",
			args: args{
				headerCount: 10,
				machine:     "machine",
				recipe:      "recipe",
				unit:        "unit",
				spec:        "spec",
				flow:        "flow",
				lot:         "lot",
				product:     "product",
				factory:     "factory",
				operation:   "operation",
				dv: map[string]string{
					"k1": "v1",
					"k2": "v2",
				},
				dvItems: nil,
			},
			want: &xmlProcessData{
				Header:               getHeader(10),
				FactoryName:          "factory",
				ProductSpecName:      "spec",
				ProcessFlowName:      "flow",
				ProcessOperationName: "operation",
				MachineName:          "machine",
				MachineRecipeName:    "recipe",
				UnitName:             "unit",
				LotName:              "lot",
				ProductName:          "product",
				ItemList: []XMLItem{
					{
						ItemName: "k1",
						SiteList: []dvSite{
							{
								SiteName:           "S01",
								SiteValue:          "v1",
								SampleMaterialName: "product",
							},
						},
					},

					{
						ItemName: "k2",
						SiteList: []dvSite{
							{
								SiteName:           "S01",
								SiteValue:          "v2",
								SampleMaterialName: "product",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewXMLProcessData(
				tt.args.headerCount,
				tt.args.machine,
				tt.args.recipe,
				tt.args.unit,
				tt.args.spec,
				tt.args.flow,
				tt.args.lot,
				tt.args.product,
				tt.args.factory,
				tt.args.operation,
				tt.args.dv,
				tt.args.dvItems,
			).(*xmlProcessData)

			if got.FactoryName != tt.args.factory || !reflect.DeepEqual(got.ItemList, tt.want.ItemList) || got.LotName != tt.want.LotName {
				t.Errorf("NewProcessData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	t.Parallel()

	type args struct {
		list         []XMLItem
		name         string
		value        string
		materialName string
		sites        map[string]string
	}

	tests := []struct {
		name string
		args args
		want []XMLItem
	}{
		{
			name: "1",
			args: args{
				list:         make([]XMLItem, 0),
				name:         "k1",
				value:        "v1",
				materialName: "prod",
				sites:        nil,
			},
			want: []XMLItem{
				{
					ItemName: "k1",
					SiteList: []dvSite{
						{
							SiteName:           "S01",
							SiteValue:          "v1",
							SampleMaterialName: "prod",
						},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				list:         make([]XMLItem, 0),
				name:         "k1",
				value:        "v1",
				materialName: "prod",
				sites: map[string]string{
					"S01":  "1",
					"S02":  "2",
					"S100": "100",
				},
			},
			want: []XMLItem{
				{
					ItemName: "k1",
					SiteList: []dvSite{
						{
							SiteName:           "S01",
							SiteValue:          "1",
							SampleMaterialName: "prod",
						},
						{
							SiteName:           "S02",
							SiteValue:          "2",
							SampleMaterialName: "prod",
						},
						{
							SiteName:           "S100",
							SiteValue:          "100",
							SampleMaterialName: "prod",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddItemToXML(tt.args.list, tt.args.name, tt.args.value, tt.args.materialName, tt.args.sites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSiteName(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				k: "S01",
			},
			want: "S01",
		},
		{
			name: "2",
			args: args{
				k: "S1",
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				k: "001",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSiteName(tt.args.k); got != tt.want {
				t.Errorf("checkKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
