package zispc

import (
	"reflect"
	"testing"
)

func TestNewJSONProcessData(t *testing.T) {
	t.Parallel()
	type args struct {
		machine   string
		lot       string
		recipe    string
		factory   string
		unit      string
		product   string
		spec      string
		flow      string
		operation string
		item      string
		sites     map[string]string
	}
	tests := []struct {
		name string
		args args
		want ProcessData
	}{
		{
			name: "1",
			args: args{
				machine:   "machine",
				lot:       "lot",
				recipe:    "recipe",
				factory:   "factory",
				unit:      "unit",
				product:   "product",
				spec:      "spec",
				flow:      "flow",
				operation: "operation",
				item:      "item",
				sites: map[string]string{
					"001": "1",
					"002": "2",
					"003": "3",
				},
			},
			want: &jsonProcessData{
				FactoryName:          "factory",
				ProductSpecName:      "spec",
				ProductFlowName:      "flow",
				ProcessOperationName: "operation",
				MachineName:          "machine",
				MachineRecipeName:    "recipe",
				UnitName:             "unit",
				LotName:              "lot",
				ProductName:          "product",
				ItemName:             "item",
				SiteList: []site{
					{
						SiteName:  "001",
						SiteValue: "1",
					},
					{
						SiteName:  "002",
						SiteValue: "2",
					},
					{
						SiteName:  "003",
						SiteValue: "3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONProcessData(tt.args.machine, tt.args.lot, tt.args.recipe, tt.args.factory, tt.args.unit, tt.args.product, tt.args.spec, tt.args.flow, tt.args.operation, tt.args.item, tt.args.sites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONProcessData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSiteList(t *testing.T) {
	t.Parallel()
	type args struct {
		sites map[string]string
	}
	tests := []struct {
		name string
		args args
		want []site
	}{
		{
			name: "1",
			args: args{
				sites: map[string]string{
					"001": "1",
					"002": "2",
					"003": "3",
				},
			},
			want: []site{
				{
					SiteName:  "001",
					SiteValue: "1",
				},
				{
					SiteName:  "002",
					SiteValue: "2",
				},
				{
					SiteName:  "003",
					SiteValue: "3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSiteList(tt.args.sites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSiteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJSONProcessDataMulti(t *testing.T) {
	t.Parallel()
	type args struct {
		machine   string
		lot       string
		recipe    string
		factory   string
		unit      string
		product   string
		spec      string
		flow      string
		operation string
		sites     []JSONSites
	}
	tests := []struct {
		name string
		args args
		want ProcessData
	}{
		{
			name: "1",
			args: args{
				machine:   "machine",
				lot:       "lot",
				recipe:    "recipe",
				factory:   "factory",
				unit:      "unit",
				product:   "product",
				spec:      "spec",
				flow:      "flow",
				operation: "operation",
				sites: []JSONSites{
					{
						ItemName: "1",
						Sites: []site{
							{
								SiteName:  "001",
								SiteValue: "1",
							},
							{
								SiteName:  "002",
								SiteValue: "2",
							},
							{
								SiteName:  "003",
								SiteValue: "3",
							},
						},
					},
					{
						ItemName: "2",
						Sites: []site{
							{
								SiteName:  "001",
								SiteValue: "1",
							},
							{
								SiteName:  "002",
								SiteValue: "2",
							},
							{
								SiteName:  "003",
								SiteValue: "3",
							},
						},
					},
				},
			},
			want: &jsonProcessDataMulti{
				FactoryName:          "factory",
				ProductSpecName:      "spec",
				ProductFlowName:      "flow",
				ProcessOperationName: "operation",
				MachineName:          "machine",
				MachineRecipeName:    "recipe",
				UnitName:             "unit",
				LotName:              "lot",
				ProductName:          "product",
				SiteList: []JSONSites{
					{
						ItemName: "1",
						Sites: []site{
							{
								SiteName:  "001",
								SiteValue: "1",
							},
							{
								SiteName:  "002",
								SiteValue: "2",
							},
							{
								SiteName:  "003",
								SiteValue: "3",
							},
						},
					},
					{
						ItemName: "2",
						Sites: []site{
							{
								SiteName:  "001",
								SiteValue: "1",
							},
							{
								SiteName:  "002",
								SiteValue: "2",
							},
							{
								SiteName:  "003",
								SiteValue: "3",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONProcessDataMulti(tt.args.machine, tt.args.lot, tt.args.recipe, tt.args.factory, tt.args.unit, tt.args.product, tt.args.spec, tt.args.flow, tt.args.operation, tt.args.sites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONProcessDataMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
