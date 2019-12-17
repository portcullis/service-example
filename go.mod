module github.com/portcullis/service-example

go 1.13

replace github.com/portcullis/application => ../application

replace github.com/portcullis/module => ../module

replace github.com/portcullis/logging => ../logging

require (
	github.com/portcullis/application v0.0.0-20191216211626-a70f0d307f3f
	github.com/portcullis/logging v0.0.0-20191216195124-dbd89d13ff83
	github.com/portcullis/module v0.0.0-20191216195407-4a3750be7c0f
)
