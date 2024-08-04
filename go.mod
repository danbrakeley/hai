module github.com/danbrakeley/interpreter

go 1.22.5

require (
	github.com/danbrakeley/bsh v0.2.1
	github.com/magefile/mage v1.15.0
)

require github.com/danbrakeley/commandline v1.0.0 // indirect

replace github.com/danbrakeley/bsh => ../../bsh
