package smash

import (
	"io"
	"sync"
	"bufio"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)
	var wait sync.WaitGroup
	// Parse input scan
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanWords)
	//buf := 
	var result uint32
	for scan.Scan() {
		//chann := make(chan map[uint32]uint, len(strings.Fields(buf.String())))
		get_s := word(scan.Text())
		result = smasher(get_s)
		wait.Add(1)
		go func(key uint32, intmap map[uint32]uint, wait *sync.WaitGroup) {
			//for i=0; i<len(); i++ {}
			intmap[key] = intmap[key] + 1
			wait.Done()
		}(result, m, &wait)
		wait.Wait()
	}
	return m
}
