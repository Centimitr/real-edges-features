package main

const ID_COUNT = 4867136

//const ID_COUNT = 1000

func main() {
	//e := Edges{}
	//e.ReadTrainSet("data/train.txt")
	//e.LoadAllIds("data/all_ids.txt")
	//e.OutputRandomNegativePairs("pn_pairs_6000.txt", 6000)
	//e.Output("features_5.txt")
	//e.ReadTestSetAndGenerateNegatives("data/test-public.txt", "testcases_3.txt")
	CombineMultipleCSV("csv", "X1.csv")
	CombineMultipleCSV("csv2", "X2.csv")
}
