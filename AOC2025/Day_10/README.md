Installation instructions for `github.com/draffensperger/golp`

```bash
go get -d github.com/draffensperger/golp
sudo apt-get update
sudo apt-get install liblpsolve55-dev
sudo apt-get install libsuitesparse-dev
export CGO_CFLAGS="-I/usr/include/lpsolve"
export CGO_LDFLAGS="-llpsolve55 -lcolamd -ldl -lm"
```