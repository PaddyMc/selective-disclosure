package main

import (
	"encoding/hex"
	"fmt"
	merkletree "github.com/wealdtech/go-merkletree"
)

func MerkleSimple() {
	secret := "mysecret"

	// Data for the tree
	data := [][]byte{
		[]byte("john"),
		[]byte("1/2/1970"),
		[]byte("frankfurt/germany"),
		[]byte("1234569"),
		[]byte("3531234567"),
	}

	// Create the tree
	tree, err := merkletree.NewUsing(data, New(secret), false)

	if err != nil {
		panic(err)
	}

	// Fetch the root hash of the tree
	root := tree.Root()

	// formatter
	lf := &lFormatter{}
	bf := &bFormatter{}
	_ = tree.DOT(lf, bf)
	//fmt.Println(dot)

	// verifing using the merkle proof
	proof, err := tree.GenerateProof(data[2], 0)
	if err != nil {
		panic(err)
	}
	verified, err := merkletree.VerifyProofUsing(
		data[2],
		false,
		proof,
		[][]byte{root},
		New(secret),
	)
	if !verified {
		fmt.Println("failed to verify proof for data in tree")
	}

	// verifing using a pollard
	pollard := tree.Pollard(2)
	for _, data := range pollard {
		fmt.Println(hex.EncodeToString(data))
	}
	verified = merkletree.VerifyPollardUsing(pollard, New(secret))
	if err != nil {
		panic(err)
	}
	if !verified {
		fmt.Println("failed to verify proof for data in tree")
	}

	fmt.Println("verified proof for data in tree")
}
