package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

// сюда писать код

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	var in chan interface{}
	var out chan interface{}
	for idx, currJob := range jobs {
		wg.Add(1)
		if idx == 0 {
			in = nil
			out = make(chan interface{}, MaxInputDataLen)
		} else {
			in = out
			newOut := make(chan interface{}, MaxInputDataLen)
			out = newOut
		}
		go func(currJob job, wg *sync.WaitGroup, in chan interface{}, out chan interface{}) {
			defer wg.Done()
			currJob(in, out)
			close(out)
		}(currJob, wg, in, out)
	}
	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	md5mutex := &sync.Mutex{}
	singleWg := &sync.WaitGroup{}
	for dataFromChan := range in {
		singleWg.Add(1)
		go countSingleHash(singleWg, out, dataFromChan, md5mutex)
	}
	singleWg.Wait()
}

func MultiHash(in, out chan interface{}) {
	multiWg := &sync.WaitGroup{}
	for dataFromChan := range in {
		multiWg.Add(1)
		go countMultiHashForCurrentNumber(multiWg, out, dataFromChan)
	}
	multiWg.Wait()
}

func CombineResults(in, out chan interface{}) {
	const separator = "_"
	var combine []string
	for result := range in {
		combine = append(combine, result.(string))
	}
	sort.Strings(combine)
	out <- strings.Join(combine, separator)
}

func countSingleHash(singleWg *sync.WaitGroup, out chan interface{}, dataFromChan interface{}, md5mutex *sync.Mutex) {
	defer singleWg.Done()
	const separator = "~"
	inputData := strconv.Itoa(dataFromChan.(int))

	// DataSignerCrc32 counting
	crcDataOut := make(chan string)
	go func(out chan string, inputData string) {
		out <- DataSignerCrc32(inputData)
	}(crcDataOut, inputData)

	// crcMd5Data counting
	crcMd5DataOut := make(chan string)
	go func(out chan string, inputData string, md5mutex *sync.Mutex) {
		md5mutex.Lock()
		md5Data := DataSignerMd5(inputData)
		md5mutex.Unlock()
		crcMd5Data := DataSignerCrc32(md5Data)
		out <- crcMd5Data
	}(crcMd5DataOut, inputData, md5mutex)

	crcData := <-crcDataOut
	crcMd5Data := <-crcMd5DataOut

	close(crcDataOut)
	close(crcMd5DataOut)

	out <- strings.Join([]string{crcData, crcMd5Data}, separator)
}

func countMultiHashForCurrentNumber(multiWg *sync.WaitGroup, out chan interface{}, dataFromChan interface{}) {
	defer multiWg.Done()
	const steps = 6
	inputData := dataFromChan.(string)
	hashesToConcat := make([]string, steps)
	localWg := &sync.WaitGroup{}
	for i := 0; i < steps; i++ {
		localWg.Add(1)
		go func(wg *sync.WaitGroup, hashesToConcat []string, inputData string, i int) {
			defer wg.Done()
			dataToHash := strconv.Itoa(i) + inputData
			hashesToConcat[i] = DataSignerCrc32(dataToHash)
		}(localWg, hashesToConcat, inputData, i)
	}
	localWg.Wait()
	out <- strings.Join(hashesToConcat, "")
}
