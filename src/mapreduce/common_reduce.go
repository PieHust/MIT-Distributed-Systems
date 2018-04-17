package mapreduce

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"sort"
)

func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	//
	// doReduce manages one reduce task: it should read the intermediate
	// files for the task, sort the intermediate key/value pairs by key,
	// call the user-defined reduce function (reduceF) for each key, and
	// write reduceF's output to disk.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTask) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Your code here (Part I).
	//
	var keyValues []KeyValue
	for i := 0; i < nMap; i++ {
		interFile := reduceName(jobName, i, reduceTask)
		file, ferr := os.Open(interFile)
		if ferr != nil {
			panic(ferr)
		}
		defer file.Close()
		br := bufio.NewReader(file)
		for {
			line, _, next := br.ReadLine()
			var tmp KeyValue
			json.Unmarshal(line, &tmp)
			if next == io.EOF {
				break
			}
			keyValues = append(keyValues, tmp)
		}
	}
	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Key < keyValues[j].Key
	})

	oFile, _ := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0666)
	defer oFile.Close()
	enc := json.NewEncoder(oFile)

	tmp := keyValues[0].Key
	var values []string
	for _, keyValue := range keyValues {
		if keyValue.Key == tmp {
			values = append(values, keyValue.Value)
		} else {
			enc.Encode(KeyValue{tmp, reduceF(tmp, values)})
			values = []string{keyValue.Value}
			tmp = keyValue.Key
		}
	}
	enc.Encode(KeyValue{tmp, reduceF(tmp, values)})
}
