package main

import (
	"fmt"
	"github.com/anush008/fastembed-go"
)

const (
	// TODO move to env
	cacheDir = "model_cache"
)

func generateChunkEmbedding(chunk string, embeddingModel fastembed.EmbeddingModel, inputSeqMaxLen int, embeddingBatchSize int) ([][]float32, error) {
	// With custom options
	options := fastembed.InitOptions{
		// TODO parameterization and correct defaults
		Model:     embeddingModel,
		CacheDir:  cacheDir,
		MaxLength: inputSeqMaxLen,
	}

	model, err := fastembed.NewFlagEmbedding(&options)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func(model *fastembed.FlagEmbedding) {
		err := model.Destroy()
		if err != nil {
			fmt.Println(err)
		}
	}(model)

	documents := []string{
		chunk,
	}

	fmt.Println(chunk)
	fmt.Println(len(chunk))
	embeddings, err := model.PassageEmbed(documents, embeddingBatchSize) //  -> Embeddings length: 4
	if err != nil {
		panic(err)
	}
	fmt.Println(embeddings)
	return embeddings, nil
}
