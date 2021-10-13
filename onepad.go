package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flobii/one-time-pad/pkg/onetimepad"
)

const name string = "onepad"

func main() {
	if len(os.Args) == 1 {
		fmt.Println(name)
		fmt.Println("    encrypts and decrypts files")
		fmt.Println()
		usage(true)
		return
	}
	err := handleArgs()
	if err != nil {
		log(err)
	}
}

func usage(e ...bool) {
	var extended bool
	if len(e) > 0 {
		extended = e[0]
	} else {
		extended = false
	}
	fmt.Printf("usage: %s <filename> <keydir / keypath>\n", name)
	if extended {
		fmt.Println("       if a file with `.key` is provided as the second argument")
		fmt.Println("       the file will be decrypted using this key.")
		fmt.Println("       otherwise the file will be encrypted and the key will be")
		fmt.Println("       stored in the path given as the second argument.")
	}
}

func log(err error) {
	fmt.Printf("%s: %v\n", name, err)
}

func handleArgs() error {
	var err error = nil
	if len(os.Args) > 3 {
		return errors.New("too many arguments")
	}
	filename := os.Args[1]
	var keypath string
	if len(os.Args) < 3 {
		keypath = "./"
	} else {
		keypath = os.Args[2]
	}
	if strings.HasSuffix(keypath, ".key") {
		data, d_err := ioutil.ReadFile(keypath)
		if d_err != nil {
			return d_err
		}

		err = onetimepad.DecryptFile(filename, data)
		if err != nil {
			return err
		}
		k_err := os.Remove(keypath)
		if k_err != nil {
			return k_err
		}
	} else {
		key := []byte{}
		keypath = strings.ReplaceAll(keypath, "\\", "/")
		if !strings.HasSuffix(keypath, "/") {
			keypath += "/"
		}
		key, err = onetimepad.EncryptFile(filename)
		if err != nil {
			return err
		}
		keypath = keypath + filepath.Base(filename) + ".key"
		ioutil.WriteFile(keypath, key, 0644)
	}
	return err
}
