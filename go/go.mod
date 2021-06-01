module github.com/PaddyMc/selective-disclosure

go 1.16

// for some reason you need to clone this locally and link this way
// TODO: please fix me
replace github.com/wealdtech/go-merkletree => /home/ghost/git/wealdtech/go-merkletree

require (
	github.com/cbergoon/merkletree v0.2.0
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/wealdtech/go-merkletree v1.0.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
)
