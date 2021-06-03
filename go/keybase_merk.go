package main

import (
	"encoding/hex"
	"fmt"
	merkletree "github.com/wealdtech/go-merkletree"
	"strings"
)

func MerkleSimple() {
	secret := "secret"

	// Data for the tree
	data := [][]byte{
		[]byte("bob"),
		[]byte("1-1-1970"),
		[]byte("berlin/germany"),
		[]byte("1234567"),
		[]byte("3531234567"),
	}

	// Create the tree
	tree, err := merkletree.NewUsing(data, New(secret), []byte{})

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
	proof, err := tree.GenerateProof(data[2])
	if err != nil {
		panic(err)
	}
	webproof := []string{hex.EncodeToString(root), fmt.Sprint(proof.Index)}
	for _, p := range proof.Hashes {
		webproof = append(webproof, hex.EncodeToString(p))
	}
	fmt.Printf(">>>> proof:\n%s\n", strings.Join(webproof, ":"))

	verified, err := merkletree.VerifyProofUsing(
		data[2],
		proof,
		root,
		New(secret),
		[]byte{},
	)
	if !verified {
		fmt.Println("failed to verify proof for data in tree")
	}

	if !verified {
		fmt.Println("failed to verify proof for data in tree")
	}

	fmt.Println("verified proof for data in tree")
}
