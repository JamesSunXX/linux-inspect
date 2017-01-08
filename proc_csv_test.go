package psn

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestProcCSV(t *testing.T) {
	dn, err := GetDevice("/boot")
	if err != nil {
		fmt.Println(err)
		t.Skip()
	}
	nt, err := GetDefaultInterface()
	if err != nil {
		fmt.Println(err)
		t.Skip()
	}

	fpath := filepath.Join(os.TempDir(), fmt.Sprintf("test-%010d.csv", time.Now().UnixNano()))
	defer os.RemoveAll(fpath)

	c := NewCSV(fpath, 1, dn, nt)
	for i := 0; i < 3; i++ {
		fmt.Printf("#%d: collecting data with %s and %s at %s\n", i, dn, nt, fpath)
		if err := c.Add(); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
	}

	// fill-in empty rows
	time.Sleep(2 * time.Second)

	if err := c.Add(); err != nil {
		t.Fatal(err)
	}

	if err := c.Save(); err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err)
		t.Skip()
	}
	fmt.Println("CSV contents:", string(b))
}
