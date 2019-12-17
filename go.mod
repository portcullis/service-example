module github.com/portcullis/service-example

go 1.13

replace github.com/portcullis/application => ../application

replace github.com/portcullis/module => ../module

replace github.com/portcullis/logging => ../logging

require (
	github.com/portcullis/application v0.0.0-20191217055053-b424562131ad
	github.com/portcullis/logging v0.0.0-20191217055113-b398c1c4c962
	github.com/portcullis/module v0.0.0-20191217055137-df89d896d642
)
