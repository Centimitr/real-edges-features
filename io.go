package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (e *Edges) sortedTrainSetKeys() []int {
	keys := make([]int, len(e.train))
	i := 0
	for k := range e.train {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func (e *Edges) Output(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Write csv", err)
		return
	}
	defer f.Close()

	//w := csv.NewWriter(f)
	//defer w.Flush()
	//for i, lnk := range e.links() {
	//	values := IntsToStrings([]int{i, lnk.A, lnk.B,})
	//	values = append(values, lnk.features(e).StringSlices()...)
	//	w.Write(values)
	//}
	total := 0
	for _, v := range e.train {
		total += len(v)
	}

	//w := bufio.NewWriter(f)
	wg := &sync.WaitGroup{}
	wg.Add(len(e.train))
	ch := make(chan string)
	i := 0
	output := ""
	for k, followings := range e.train {
		go func(a int, fs []int, wg *sync.WaitGroup) {
			for _, f := range fs {
				lnk := EdgeLink{a, f}
				//ch <- fmt.Sprintf("%d,%d,", a, f) + lnk.features(e).CSVString()
				//fmt.Fprintln(w, fmt.Sprintf("%d,%d,", a, f)+lnk.features(e).CSVString())
				output += fmt.Sprintf("%d,%d,", a, f) + lnk.features(e).CSVString() + "\n"
				i++
				//if i > 1000 {
				//	wg.Done()
				//	return
				//}
			}
			wg.Done()

		}(k, followings, wg)
	}
	go func() {
		start := time.Now()
		//lastTime := start
		//lastI := i
		for {
			//speed := strconv.Itoa((i-lastI)/(int(time.Since(lastTime).Seconds() + 1))) + "/ps"
			progress := strconv.Itoa(i) + "/" + strconv.Itoa(total)
			percentage := fmt.Sprintf(" %.2f", (float64(i)*float64(100))/float64(total)) + "%"
			fmt.Print("\r", time.Since(start), " ", progress, " ", percentage)
			time.Sleep(1 * time.Second)
			//lastTime = time.Now()
			//lastI = i
		}
	}()
	//go func() {
	wg.Wait()
	close(ch)
	//}()
	//w.Flush()
	io.WriteString(f, output)
}

func (e *Edges) ReadTestSetAndGenerateNegatives(inPath string, outPath string) {
	csvf, err := os.Open(inPath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvf.Close()

	out, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	i := 0
	for vs := range ProcessCSV(csvf) {
		//if i < 2 {
		//vs := strings.Split(v[0], "	")
		a, _ := strconv.Atoi(vs[1])
		b, _ := strconv.Atoi(vs[2])
		//fmt.Println(a, b)

		lnk := EdgeLink{a, b}
		//ch <- fmt.Sprintf("%d,%d,", a, f) + lnk.features(e).CSVString()
		//fmt.Fprintln(w, fmt.Sprintf("%d,%d,", a, f)+lnk.features(e).CSVString())
		io.WriteString(out, fmt.Sprintf("%d,%d,", a, b)+lnk.features(e).CSVString()+"\n")
		i++
		//}
	}

}

func (e *Edges) SaveAllIds(filename string) {

	allIdsM := make(map[int]struct{})
	for k, followings := range e.train {
		allIdsM[k] = struct{}{}
		for _, f := range followings {
			allIdsM[f] = struct{}{}
		}
	}
	allIds := make([]int, len(allIdsM))
	i := 0
	for k := range allIdsM {
		allIds[i] = k
		i++
	}
	//sort.Ints(allIds)
	e.allIds = allIds
	err := ioutil.WriteFile(filename, []byte(strings.Join(IntsToStrings(allIds), " ")), 0755)
	if err != nil {
		fmt.Println("Write AllIds", err)
	}
}

func (e *Edges) LoadAllIds(filename string) {
	byts, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Read AllIds", err)
		return
	}
	e.allIds = StringsToInts(strings.Split(string(byts), " "))
}
