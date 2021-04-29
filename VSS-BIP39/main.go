package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	r := "94b8c423610d2ce4adb08ca74bae438e"
	m, err := readLine("english.txt")
	if err != nil {
		log.Panic(err)
	}

	rhex, err := hex.DecodeString(r)
	if err != nil {
		log.Panic(err)
	}

	//add checksum to the end of the string
	h := sha256.Sum256(rhex)
	fmt.Printf("hash: [%s]\n", hex.EncodeToString(h[:]))
	r = r + string(hex.EncodeToString(h[0:1])[0])

	var a []string
	b := new(big.Int)
	b.SetString(r, 16)

	n := big.NewInt(0)
	for i := 0; i < 12; i++ {
		n.And(b, big.NewInt(**yournumberhere**))

		//prepend to array
		a = append([]string{m[uint(n.Uint64())]}, a...)

		b = b.**operation**(b, **yournumberhere**)
	}
	words := strings.Join(a, " ")

	fmt.Printf("words are: [%s]\n", words)
	bip39seed := pbkdf2.Key([]byte(words), []byte("mnemonic"), 2048, 64, sha512.New)
	fmt.Printf("bip39seed: [%s]\n", hex.EncodeToString(bip39seed))
}

func readLine(filename string) (map[uint]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := map[uint]string{}
	i := uint(0);
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
