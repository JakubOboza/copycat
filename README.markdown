# Copycat 

Simple progress tracking/wrapper for io.Reader/io.Writer.

How to use:

import
```
import "github.com/JakubOboza/copycat"
```

wrap with progress manager the io.Reader or io.Writer and set your own callback.
```
	source := bytes.NewBufferString("This is pretty long text that we will use as simple test for our progress manager foo")
	var dest bytes.Buffer

	pm := NewProgressReader(source)

	pm.AddListenerFunc(func(progress int) {
		fmt.Println("Hey something is happening we are not stuck it jus takes long time.")
	})

	if _, err := io.Copy(&dest, pm); err != nil {
		fmt.Println("Something went wrong :(")
	}
```

Much better example can be found in *examples* directory.

# Running exmaples

Simple example shows how this can be used to signal % of done copy when downloading file.

```
cd examples/simple
```

and 

```
go run main.go
```


# Running tests

```
make test
```

or just

```
make
```