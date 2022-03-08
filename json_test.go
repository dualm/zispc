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
				ProcessFlowName:      "flow",
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
			} else {
				b, err := got.Encode()
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("%s\n", string(b))
			}
		})
	}
}

func Test_makeSiteList(t *testing.T) {
	t.Parallel()
	type args struct {
		sampleName string
		sites      map[string]string
	}
	tests := []struct {
		name string
		args args
		want []site
	}{
		{
			name: "1",
			args: args{
				sampleName: "sample",
				sites:      map[string]string{"001": "1", "002": "2", "003": "3"},
			},
			want: []site{
				{
					SampleMaterialName: "sample",
					SiteName:           "001",
					SiteValue:          "1",
				},
				{
					SampleMaterialName: "sample",
					SiteName:           "002",
					SiteValue:          "2",
				},
				{
					SampleMaterialName: "sample",
					SiteName:           "003",
					SiteValue:          "3",
				},
			},
		},
		{
			name: "2",
			args: args{
				sampleName: "",
				sites:      map[string]string{"001": "1", "002": "2", "003": "3"},
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
			if got := makeSiteList(tt.args.sampleName, tt.args.sites); !reflect.DeepEqual(got, tt.want) {
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
		sites     []Sites
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
				sites: func() []Sites {
					re := make([]Sites, 0)

					re = AddSite(re, "1", "sample1", map[string]string{"001": "1", "002": "2", "003": "3"})
					re = AddSite(re, "2", "sample1", map[string]string{"001": "1", "002": "2", "003": "3"})

					return re
				}(),
			},
			want: &jsonProcessDataMulti{
				FactoryName:          "factory",
				ProductSpecName:      "spec",
				ProcessFlowName:      "flow",
				ProcessOperationName: "operation",
				MachineName:          "machine",
				MachineRecipeName:    "recipe",
				UnitName:             "unit",
				LotName:              "lot",
				ProductName:          "product",
				SiteList: []Sites{
					{
						ItemName: "1",
						Sites: []site{
							{
								SampleMaterialName: "sample1",
								SiteName:           "001",
								SiteValue:          "1",
							},
							{
								SampleMaterialName: "sample1",
								SiteName:           "002",
								SiteValue:          "2",
							},
							{
								SampleMaterialName: "sample1",
								SiteName:           "003",
								SiteValue:          "3",
							},
						},
					},
					{
						ItemName: "2",
						Sites: []site{
							{
								SampleMaterialName: "sample1",
								SiteName:           "001",
								SiteValue:          "1",
							},
							{
								SampleMaterialName: "sample1",
								SiteName:           "002",
								SiteValue:          "2",
							},
							{
								SampleMaterialName: "sample1",
								SiteName:           "003",
								SiteValue:          "3",
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
			} else {
				b, err := got.Encode()
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("%s\n", string(b))
			}
		})
	}
}

func TestAddSite(t *testing.T) {
	t.Parallel()

	type args struct {
		s          []Sites
		item       string
		sampleName string
		sites      map[string]string
	}
	tests := []struct {
		name string
		args args
		want []Sites
	}{
		{
			name: "1",
			args: args{
				s:          nil,
				item:       "item1",
				sampleName: "sample1",
				sites: map[string]string{
					"k1": "v1",
					"k2": "v2",
				},
			},
			want: []Sites{
				{
					ItemName: "item1",
					Sites: []site{
						{
							SampleMaterialName: "sample1",
							SiteName:           "k1",
							SiteValue:          "v1",
						},
						{
							SampleMaterialName: "sample1",
							SiteName:           "k2",
							SiteValue:          "v2",
						},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				s: []Sites{
					{
						ItemName: "item1",
						Sites: []site{
							{
								SampleMaterialName: "sample1",
								SiteName:           "k1",
								SiteValue:          "v1",
							},
							{
								SampleMaterialName: "sample1",
								SiteName:           "k2",
								SiteValue:          "v2",
							},
						},
					},
				},
				item:       "item2",
				sampleName: "sample1",
				sites: map[string]string{
					"k3": "v3",
					"k4": "v4",
				},
			},
			want: []Sites{
				{
					ItemName: "item1",
					Sites: []site{
						{
							SampleMaterialName: "sample1",
							SiteName:           "k1",
							SiteValue:          "v1",
						},
						{
							SampleMaterialName: "sample1",
							SiteName:           "k2",
							SiteValue:          "v2",
						},
					},
				},
				{
					ItemName: "item2",
					Sites: []site{
						{
							SampleMaterialName: "sample1",
							SiteName:           "k3",
							SiteValue:          "v3",
						},
						{
							SampleMaterialName: "sample1",
							SiteName:           "k4",
							SiteValue:          "v4",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSite(tt.args.s, tt.args.item, tt.args.sampleName, tt.args.sites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
