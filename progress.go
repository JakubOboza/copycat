package copycat

import "io"

//ProgressManager is holding info about listeners
type ProgressManager struct {
	reader    io.Reader
	writer    io.Writer
	listeners []ProgressListener
}

type ProgressListener interface {
	ProgressUpdate(progress int) //You must know the size to conver it into % sadly
}

//NewProgressReader creaters progress manager from Reader
func NewProgressReader(readerToWrap io.Reader) *ProgressManager {
	return &ProgressManager{reader: readerToWrap}
}

//NewProgressWriter creates progress manager from Writer
func NewProgressWriter(writerToWrap io.Writer) *ProgressManager {
	return &ProgressManager{writer: writerToWrap}
}

//Reader implementes the io.Reader interface
func (pm *ProgressManager) Read(p []byte) (n int, err error) {
	n, err = pm.reader.Read(p)
	pm.update(n)
	return n, err
}

//Writer implementes the io.Writer interface
func (pm *ProgressManager) Write(p []byte) (n int, err error) {
	n, err = pm.writer.Write(p)
	pm.update(n)
	return n, err
}

//AddListener simply adds listeners to the list of observing for events on progress
func (pm *ProgressManager) AddListener(newListener ProgressListener) {
	pm.listeners = append(pm.listeners, newListener)
}

type wrapProgressFunc struct {
	callback func(int)
}

func (wp *wrapProgressFunc) ProgressUpdate(progress int) {
	wp.callback(progress)
}

//AddListenerFunc gives a shorthand to specify only callback function rather entire type
func (pm *ProgressManager) AddListenerFunc(fn func(int)) {
	wrapper := &wrapProgressFunc{callback: fn}
	pm.AddListener(wrapper)
}

func (pm *ProgressManager) update(progress int) {
	for _, l := range pm.listeners {
		l.ProgressUpdate(progress)
	}
}
