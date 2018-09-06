package main

import (
	"flag"
)

const IdCount = 4867136

//const ID_COUNT = 1000

func main() {
	e := Edges{}
	e.ReadTrainSet("data/train.txt")
	//e.LoadAllIds("ldata/all_ids.txt")
	//e.OutputRandomPositivePairs("tp_pairs_20000.csv", 20000)
	e.OutputRandomNegativePairs("pn_pairs_20000.csv", 20000)
	//e.Output("features_5.txt")
	//e.ReadTestSetAndGenerateNegatives("data/test-public.txt", "testcases_3.txt")
}

func mainCombine() {
	var dir, out, prefix string
	flag.StringVar(&dir, "dir", ".", "dir")
	flag.StringVar(&out, "out", "no_name.csv", "out")
	flag.StringVar(&prefix, "prefix", "", "dir")

	flag.Parse()

	if out == "no_name.csv" && prefix != "" {
		out = prefix + ".csv"
	}
	CombineMultipleCSVByPrefix(dir, out, prefix)
}
