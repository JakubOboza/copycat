package copycat

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestProgressManager struct {
	Updates []int
}

func (tp *TestProgressManager) ProgressUpdate(progress int) {
	tp.Updates = append(tp.Updates, progress)
}

type TestReaderSource struct {
	max   int
	count int
	chunk []byte
}

func (trs *TestReaderSource) Read(p []byte) (n int, err error) {
	if trs.count < trs.max {
		copy(p, trs.chunk)
		trs.count++
		return len(trs.chunk), nil
	} else {
		return 0, io.EOF
	}
}

func TestProgressManagerFromReader(t *testing.T) {

	source := bytes.NewBufferString("This is pretty long text that we will use as simple test for our progress manager foo")
	var dest bytes.Buffer

	pm := NewProgressReader(source)

	tp := &TestProgressManager{}
	pm.AddListener(tp)

	if _, err := io.Copy(&dest, pm); err != nil {
		t.Fatal("Copy didnt work at all")
	}

	assert.Equal(t, []int{85, 0}, tp.Updates)
}

func TestProgressManagerFromWriter(t *testing.T) {
	source := bytes.NewBufferString("This is pretty long text that we will use as simple test for our progress manager foo")
	dest := &bytes.Buffer{}

	pm := NewProgressWriter(dest)

	tp := &TestProgressManager{}
	pm.AddListener(tp)

	if _, err := io.Copy(pm, source); err != nil {
		t.Fatal("Copy didnt work at all")
	}

	assert.Equal(t, []int{85}, tp.Updates)
}

func TestProgressManagerFromCustomReader(t *testing.T) {

	source := &TestReaderSource{max: 50, chunk: []byte{'h', 'e', 'l', 'l', 'o'}}
	var dest bytes.Buffer

	pm := NewProgressReader(source)

	tp := &TestProgressManager{}
	pm.AddListener(tp)

	if _, err := io.Copy(&dest, pm); err != nil {
		t.Fatal("Copy didnt work at all")
	}

	assert.Equal(t, 51, len(tp.Updates)) // 51 because it includes starting 0
	destStr, _ := dest.ReadString('\n')
	assert.Equal(t, 250, len(destStr))

	resultShouldBe := "hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello"

	assert.Equal(t, resultShouldBe, destStr)
}
