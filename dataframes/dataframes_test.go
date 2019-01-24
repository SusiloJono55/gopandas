package dataframes

import (
	"encoding/json"
	"fmt"
	"gopandas/indices"
	"gopandas/series"
	"gopandas/types"
	"math/rand"
	"reflect"
	"testing"
)

var (
	workHour   = series.New([]int{8, 10, 10, 13, 11, 12, 15})
	person     = series.New([]string{"ali", "manan", "korak", "budi", "tolhal", "udin", "badu"})
	department = series.New([]string{"sales", "operation", "sales", "sales", "marketing", "finance", "marketing"})
	combine    = series.New(map[indices.Index]interface{}{"satu": []int{1, 2}, "dua": []int{3, 4}})
)

type IdxList struct {
	Idx []int
}

func TestNew(t *testing.T) {

	type args struct {
		columns []string
		ss      []*series.Series
	}
	tests := []struct {
		name string
		args args
		want *DataFrame
	}{
		{name: "Test New Dataframe",
			args: args{
				columns: []string{"working", "person", "unit"},
				ss:      []*series.Series{workHour, person, department}},
			want: New([]string{"working", "person", "unit"}, []*series.Series{workHour, person, department}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print(tt.want)
			fmt.Print(workHour)
			fmt.Print(person)
			fmt.Println(department)
			fmt.Println(combine)
			if got := New(tt.args.columns, tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataFrame_Describe(t *testing.T) {
	sr := series.New([]float64{-5, -1, 1.1, 2, 3, 3, 4, 6, 7, 7, 10, 17})
	df := New([]string{"sales"}, []*series.Series{sr})
	fmt.Println(df.Describe())

}

func TestIdxCollection(t *testing.T) {
	m := make(map[string][]int)
	source := map[int]string{0: "sales", 1: "operation", 2: "sales", 3: "sales", 4: "marketing", 5: "finance", 6: "marketing"}

	for k, v := range source {
		m[v] = append(m[v], k)
	}

	fmt.Println(m)

}

func TestDataFrame_GroupBy(t *testing.T) {
	df := New([]string{"working", "person", "unit"}, []*series.Series{workHour, person, department})
	gdf := df.GroupBy("unit")
	fmt.Println(gdf)

}

func TestDataFrame_SelectByIndex(t *testing.T) {

	df := New([]string{"working", "person", "unit"}, []*series.Series{workHour, person, department})
	dfg := df.GroupBy("unit")

	ret := NewEmpty()

	sr := series.NewEmpty()
	for k, v := range dfg.Group {
		indices := reflect.TypeOf(v).String()
		indicesVals := reflect.ValueOf(v)
		idx := fmt.Sprintf("%v%v", indices, indicesVals)
		fmt.Println(indices)
		fmt.Println(idx)
		// dfs := df.SelectByIndex(indices).Df[x].Max()
		sr.Set(k, idx)
	}
	ret.AddSeries("cols", sr)

	fmt.Println(ret)

}

func TestDataFrame_ToJson(t *testing.T) {
	df := New([]string{"working", "person", "unit"}, []*series.Series{workHour, person, department})
	js, err := df.ToJson()
	if err != nil {
		t.Error("failed to Marshall to json")
	}

	fmt.Printf("%s\n", js)
}

func TestDataFrame_Maps(t *testing.T) {
	df := New([]string{"working", "person", "unit"}, []*series.Series{workHour, person, department})
	dfg := df.GroupBy("unit").Max()
	dl := map[string]interface{}{
		"0": 1,
		"2": 4,
	}
	fmt.Println(dl)
	js := dfg.Select("working").Maps()
	// if err != nil {
	// 	t.Error("failed to Marshall to json")
	// }

	fmt.Printf("%s\n", js)

	for _, xx := range js {
		ff := xx.(map[string]interface{})
		js, err := json.Marshal(ff)
		if err != nil {
			fmt.Errorf("error : %v", err)
		}
		fmt.Println(string(js))
		// for _, yy := range ff {
		// 	fmt.Printf("%+v, %T\n", yy, yy)
		// }
	}
}

func Test_MixDataColumn(t *testing.T) {
	s := series.New(map[indices.Index]interface{}{
		0: 1,
		1: 1,
		2: 5,
		3: types.NewNan(),
		4: 7,
		5: 2,
	})

	df := NewEmpty()
	df.AddSeries("mix", s)
	dfg := df.Describe()
	dfs := df.Df["mix"].ValuesCount()
	fmt.Println(df)
	fmt.Println(dfg)
	fmt.Println(dfs)
}

func Test_IterateSeries(t *testing.T) {
	s := series.NewEmpty()

	for idx := 0; idx < 10; idx++ {
		switch idx % 2 {
		case 0:
			s.Set(idx, types.Numeric((rand.Float64()*5)+5))
		default:
			s.Set(idx, types.NewNan())
		}

	}

	fmt.Println(s)

	df := NewEmpty()
	df.AddSeries("numbers", s)
	fmt.Println(df.Df["numbers"].Sum())
}
