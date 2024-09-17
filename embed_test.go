package main

import (
	"fmt"
	"testing"
)

func TestEmbedding(t *testing.T) {
	var testEnv = &env{
		EmbeddingBatchSize: 25,
		EmbeddingModel:     "fast-bge-base-en",
		InputSeqMaxLen:     200,
	}
	// chunk673 works
	chunk673 := "AS YOU LIKE IT\nDRAMATIS PERSONAE\nDUKE SENIOR     living in banishment.\nDUKE FREDERICK  his brother, an usurper of his dominions.\nAMIENS  | |  lords attending on the banished duke. JAQUES       |\nLE BEAU a courtier attending upon Frederick.\nCHARLES wrestler to Frederick.\nOLIVER          | | JAQUES (JAQUES DE BOYS:)    |  sons of Sir Rowland de Boys. | ORLANDO               |\nADAM    | |  servants to Oliver. DENNIS |\nTOUCHSTONE      a clown.\nSIR OLIVER MARTEXT      a vicar.\nCORIN   |\n|  shepherds.\nSILVIUS |\nWILLIAM a country fellow in love with Audrey.\nA person representing HYMEN. (HYMEN:)\nROSALIND        daughter to the banished duke.\nCELIA   daughter to Frederick.\n"

	fmt.Print(chunk673)
	fmt.Println(len(chunk673))
	_, err := generateChunkEmbedding(chunk673, testEnv.EmbeddingModel, testEnv.InputSeqMaxLen, testEnv.EmbeddingBatchSize)
	if err != nil {
		t.Errorf("unable to generate chunk embeddings text length 673: %v", err)
	}
	// chunk835 fails
	chunk835 := "AS YOU LIKE IT\nDRAMATIS PERSONAE\nDUKE SENIOR     living in banishment.\nDUKE FREDERICK  his brother, an usurper of his dominions.\nAMIENS  | |  lords attending on the banished duke. JAQUES       |\nLE BEAU a courtier attending upon Frederick.\nCHARLES wrestler to Frederick.\nOLIVER          | | JAQUES (JAQUES DE BOYS:)    |  sons of Sir Rowland de Boys. | ORLANDO               |\nADAM    | |  servants to Oliver. DENNIS |\nTOUCHSTONE      a clown.\nSIR OLIVER MARTEXT      a vicar.\nCORIN   |\n|  shepherds.\nSILVIUS |\nWILLIAM a country fellow in love with Audrey.\nA person representing HYMEN. (HYMEN:)\nROSALIND        daughter to the banished duke.\nCELIA   daughter to Frederick.\nPHEBE   a shepherdess.\nAUDREY  a country wench.\nLords, pages, and attendants, &c. (Forester:) (A Lord:) (First Lord:) (Second Lord:) (First Page:) (Second Page:)\n"
	fmt.Print(chunk835)
	fmt.Println(len(chunk835))
	_, err = generateChunkEmbedding(chunk835, testEnv.EmbeddingModel, testEnv.InputSeqMaxLen, testEnv.EmbeddingBatchSize)
	if err != nil {
		t.Errorf("unable to generate chunk embeddings text length 673: %v", err)
	}

}
