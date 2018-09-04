package main

const ID_COUNT = 4867136

//const ID_COUNT = 1000

func main() {
	e := Edges{}
	e.ReadTrainSet("data/train.txt")
	//e.LoadAllIds("data/all_ids.txt")
	e.OutputRandomPositivePairs("tp_pairs_20000.txt", 20000)
	e.OutputRandomNegativePairs("pn_pairs_20000.txt", 20000)
	//e.Output("features_5.txt")
	//e.ReadTestSetAndGenerateNegatives("data/test-public.txt", "testcases_3.txt")

	//CombineMultipleCSV("csv", "X1.csv")
	//CombineMultipleCSV("csv2", "X2.csv")
	//CombineMultipleCSV("csv3", "X3.csv")
}
