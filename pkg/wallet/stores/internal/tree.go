package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sonrhq/core/pkg/crypto"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type AccountTree struct {
	Name     string
	Children []*AccountTree
}

func buildAccountTree(accounts []*vaultv1.AccountConfig) *AccountTree {
	tree := &AccountTree{Name: "Blockchains"}
	for _, acc := range accounts {
		current := tree
		parts := []string{"Blockchains", fmt.Sprintf("%d", acc.CoinTypeIndex)}
		for _, part := range parts {
			child := current.findChild(part)
			if child == nil {
				child = &AccountTree{Name: part}
				current.Children = append(current.Children, child)
			}
			current = child
		}
		current.Children = append(current.Children, &AccountTree{Name: acc.Name})
	}
	return tree
}

func (node *AccountTree) findChild(name string) *AccountTree {
	for _, child := range node.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func printAccountTree(node *AccountTree, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Name)
	for _, child := range node.Children {
		printAccountTree(child, indent+1)
	}
}

func getAccountPath(basePath string, coinType crypto.CoinType) (string, error) {
	// Create the parent directories for the file path
	err := os.MkdirAll(filepath.Join(basePath, bip44Prefix, strconv.Itoa(int(coinType))), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("could not create directory for coin type %d: %w", coinType, err)
	}

	// Check if the directory for the coin type exists, create it if it doesn't
	coinTypeDir := filepath.Join(basePath, bip44Prefix, strconv.Itoa(int(coinType.BipPath())))
	if _, err := os.Stat(coinTypeDir); os.IsNotExist(err) {
		if err := os.Mkdir(coinTypeDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("could not create directory for coin type %d: %w", coinType, err)
		}
	}
	// Find the next available account number for the coin type
	nextAccountNum := 1
	err = filepath.Walk(coinTypeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".json" {
			return nil
		}
		// Extract the account number from the filename and update nextAccountNum if necessary
		var accountNum int
		_, err = fmt.Sscanf(info.Name(), "account%d.json", &accountNum)
		if err == nil && accountNum >= nextAccountNum {
			nextAccountNum = accountNum + 1
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("could not find next available account number: %w", err)
	}
	// Return the path for the new account
	accountPath := filepath.Join(coinTypeDir, fmt.Sprintf("account%d.json", nextAccountNum))
	return accountPath, nil
}
