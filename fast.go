package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func (r *reader) addWord(word []byte) { // adding words without channels
	state, index := r.contains(word)
	if state {
		(*r.words)[index].counter++
	} else {
		record := record{word, 1, false}
		*r.words = append(*r.words, record)
	}
}

func strip(s []byte) []byte {
	n := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {


			/*if 'A' <= b && b <= 'Z' {
				b += 'a' - 'A'
				s[n] = b
			} else {
				s[n] = b
			}*/

			s[n] = b
			n++
		} else if (b > 122 || b < 65) ||
					(b > 90 && b < 97) {
			s[n] = ' '
			n++
		}
	}
	return s[:n]
}

func Fast(out io.Writer) {
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sc := bufio.NewScanner(file)

	words := make([]record, 0)
	reader := &reader{words: &words}

	var row []byte
	var lineWords [][]byte

	//var byteVal byte
	//var word []byte

	//ch := make(chan []byte)

	//go reader.read_from_chan(ch)

	for ;sc.Scan(); {

		row = sc.Bytes()

		row = bytes.ToLower(row)
		row = strip(row)

		lineWords = bytes.Split(row, []byte(" "))


		for _, res := range lineWords {
			if len(res) != 0 {
				reader.addWord(res)
			}
		}

		/*for i:=0; i<len(row); i++ {
			byteVal = row[i]

			if byteVal >= 65 && byteVal <= 90 {

				byteVal = byteVal + 32
				word = append(word, byteVal)


			} else if byteVal >= 97 && byteVal <= 122 {

				word = append(word, byteVal)

			} else if byteVal == 32 && len(word) != 0 {

				//ch <- word
				reader.addWord(word)
				word = word[:0]

			} else if ((byteVal > 122 || byteVal < 65) || (byteVal > 90 && byteVal < 97)) && len(word) != 0 {

				//ch <- word
				reader.addWord(word)
				word = word[:0]


			} else {
				continue
			}

			if i == len(row)-1 && byteVal != 32 && len(word)!=0 {
				reader.addWord(word)
				//ch <- word
				word = word[:0]

			}

		}*/
	}
	//close(ch)

	reader.get20mostfrequentwords() //getting 20 most frequent words, and write it to rating slice
	reader.print(out)

	/*readingBuf := make([]byte, 1) //read file by one letter only

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
	reader.print(out)*/                  //print elements from words according to the rating list
}
