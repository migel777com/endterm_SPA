package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type writer struct {
	writing_buf *[]byte
}

func (w *writer) write_to_temp_buf(byte byte) {
	*w.writing_buf = append(*w.writing_buf, byte)
	//fmt.Println("WRITER: Written bytes to writing_buf", byte)
}

func (w *writer) write_to_chan(ch chan []byte) {
	ch <- *w.writing_buf
	//fmt.Println("WRITER: Send byte slice to chan", writing_buf)
	*w.writing_buf = nil
}

type reader struct {
	words  *[]record
	rating *[]int
}

type record struct {
	word    []byte
	counter int
	checked bool
}

func (r *reader) contains(element []byte) (bool, int) {
	for index, v := range *r.words {
		if bytes.Equal(v.word, element) {
			return true, index
		}
		index = index + 1
	}
	return false, 0
}

func (r *reader) read_from_chan(ch chan []byte) { //reading word from the channel
	for node := range ch {
		state, index := r.contains(node)
		if state {
			(*r.words)[index].counter++
		} else {
			record := record{node, 1, false}
			*r.words = append(*r.words, record)
		}
	}
}

func (r *reader) get20mostfrequentwords() {
	list := make([]int, 20)
	r.rating = &list
	for index, _ := range *r.rating {
		temp := 0
		inss := 0
		for index, v := range *r.words {
			if (v.checked == false) && (v.counter > temp) {
				temp = v.counter
				inss = index
			}
		}
		(*r.words)[inss].checked = true
		(*r.rating)[index] = inss
	}

}

func (r *reader) print(out io.Writer) {
	for _, v := range *r.rating {
		fmt.Fprintln(out, (*r.words)[v].counter, " ", string((*r.words)[v].word), (*r.words)[v].word)
		//fmt.Println()
	}
}

func Slow(out io.Writer) {
	//start := time.Now()
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	readingBuf := make([]byte, 1) //read file by one letter only

	words := make([]record, 0)
	reader := reader{words: &words} //creating reader

	writingBuf := make([]byte, 0)
	writer := writer{&writingBuf} //writer

	ch := make(chan []byte) //channel that we will use to pass slices of bytes from writer to reader
	//btw reader listens in range of elements that are passed to channel, it will stop working when there are no elements left, so we don't need any wait groups

	go func() {
		for {
			//reading file's letters one by one
			n, err := file.Read(readingBuf)

			if n > 0 {
				byteVal := readingBuf[0]
				if byteVal >= 65 && byteVal <= 90 { //if symbol is uppercase letter

					byteVal = byteVal + 32
					writer.write_to_temp_buf(byteVal) //writing to temporary buffer

				} else if byteVal >= 97 && byteVal <= 122 { //if symbol is lowercase letter

					writer.write_to_temp_buf(byteVal) //writing to temporary buffer

				} else if byteVal == 32 && len(writingBuf) != 0 { //if symbol is [space], and we have letters in our buffer

					writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer

				} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(writingBuf) != 0 { //if symbol is any other than letter or space, and we have letters in our buffer

					writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer

				} else {
					continue
				}
			}

			if err == io.EOF {
				writer.write_to_chan(ch) //send temporary buffer content to channel, empty the temporary buffer
				break
			}
		}
		close(ch) //close channel, so our that our reader will stop working after there are no elements left, in other case reader will cause deadlock
	}()

	reader.read_from_chan(ch) //reading from channel in range of elements in channel

	reader.get20mostfrequentwords() //getting 20 most frequent words, and write it to rating slice
	reader.print(out)                  //print elements from words according to the rating list
	//fmt.Printf("Process took %s\n", time.Since(start))
}
