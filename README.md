# go-sizedwriter

Size limited writer

## Usage

```go
sw := sizedwriter.NewWriter(filename, 500, 0644, func(sw *sizedwriter.Writer) error {
	println("limited!")
	return nil
})
```

## Installation

```
go get github.com/mattn/go-sizedwriter
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
