package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"os"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt the file passed",
	Long: `Decrypts the file and store the decrypted text passed to the file specified
  For Example: 
  cryptic decrypt -f file.txt -k "sample key here" -o decryptedfile.enc`,

	Run: decrypt,
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringP("file", "f", "", "file to be decrypted")
	decryptCmd.Flags().StringP("outfile", "o", "decrypted.dec", "file to export decrypted text")
	decryptCmd.Flags().StringP("key", "k", "Ja1URn%rp|F3=2n]VMgELG*(J&bY8aHY", "key to decrypt the text with")
}

func decrypt(cmd *cobra.Command, args []string) {
	encryptedfile, _ := cmd.Flags().GetString("file")
	outputfile, _ := cmd.Flags().GetString("outfile")
	key, _ := cmd.Flags().GetString("key")

	ciphertext, err := os.ReadFile(encryptedfile)
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

	NonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:NonceSize], ciphertext[NonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), ciphertext, nil)
	if err != nil {
		panic(err)
	}

	os.WriteFile(outputfile, plaintext, 0644)
}
