This repository demonstrates a bug of gocheck.

The bug occurs when two packages are placed into one folder.

### Reproduction:

`go test -v ./...`

### Directory structure:

```.
├── gocheck_one_package
│   └── init_test.go
├── gocheck_two_packages_one_folder
│   ├── init_package2_test.go
│   └── init_test.go
├── one_package
│   └── a_test.go
└── two_packages_one_folder
    ├── two_packages_one_folder_a_test.go
    └── two_packages_one_folder_b_test.go
```

### Output
```
go test -v ./...
=== RUN   Test
OK: 1 passed
--- PASS: Test (0.00s)
PASS
ok  	_____/double-test/gocheck_one_package	0.010s
=== RUN   Test
OK: 2 passed
--- PASS: Test (0.00s)
=== RUN   Test
OK: 2 passed
--- PASS: Test (0.00s)
PASS
ok  	_____/double-test/gocheck_two_packages_one_folder	0.007s
=== RUN   TestOnePackage
--- PASS: TestOnePackage (0.00s)
PASS
ok  	_____/double-test/one_package	0.009s
=== RUN   TestTwoInOneMain
--- PASS: TestTwoInOneMain (0.00s)
=== RUN   TestTwoInOneTestPackage
--- PASS: TestTwoInOneTestPackage (0.00s)
PASS
ok  	_____/double-test/two_packages_one_folder	0.001s
```
