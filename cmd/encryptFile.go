package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt the file passed",
	Long: `Encrypts the file passed to the file specified
  For Example: 
  cryptic encrypt -f file.txt -k "sample key here" -o encryptedfile.txt`,

	Run: encrypt,
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().StringP("file", "f", "", "file to be encrypted")
	encryptCmd.Flags().StringP("outfile", "o", "encrypted.enc", "file to export encrypted text")
	encryptCmd.Flags().StringP("key", "k", "Ja1URn%rp|F3=2n]VMgELG*(J&bY8aHY", "key to encrypt the text with")
}

func encrypt(cmd *cobra.Command, args []string) {
	plainfile, _ := cmd.Flags().GetString("file")
	outputfile, _ := cmd.Flags().GetString("outfile")
	key, _ := cmd.Flags().GetString("key")

	plaintext, err := os.ReadFile(plainfile)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Println(string(ciphertext))
	os.WriteFile(outputfile, ciphertext, 0644)

	fmt.Println("Your encrypted file saved as ", outputfile)
}
