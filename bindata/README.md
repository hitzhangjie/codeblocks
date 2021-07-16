# bindata
An easy to use utility that can embed any files into go binary

Run `bindata -h` to see the usage:

```bash
Usage of bindata:
  -gopkg string
    	write transformed data to *.go, whose package is $package (default "gobin")
  -input string
    	read data from input, which could be a regular file or directory
  -output string
    	write transformed data to named *.go, which could be linked with binary
```

If there's an folder named `DIR` which may contains files and subfolders, then we want to convert it into a go file named `dir.go`, we also want the package named as `package myasset`.

Run  `bindata -gopkg mydir -input DIR -output dir.go` to finish the conversion.

