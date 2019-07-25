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

