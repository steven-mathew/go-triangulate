# `go-triangulate`
[![License][License-Image]][License-URL]
[![Go Report Card][GoReportCard-Image]][GoReportCard-URL]

`go-triangulate` uses the technique covered in [Triangulation by Ear Clipping](https://www.geometrictools.com/Documentation/TriangulationByEarClipping.pdf).

A port of [`triangulate`](https://github.com/ekzhang/triangulate) to golang and wasm.

## Running `go-triangulate`
```bash
go run server.go
GOARCH=wasm GOOS=js go build -o lib.wasm main.go
```

## Docker
```
TODO
```

## License

Licensed under either of

 * Apache License, Version 2.0
   ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)

at your option.


[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg
[License-URL]: http://opensource.org/licenses/Apache
[GoReportCard-Image]: https://goreportcard.com/badge/github.com/steven-mathew/go-triangulate
[GoReportCard-URL]: https://goreportcard.com/report/github.com/steven-mathew/go-triangulate
