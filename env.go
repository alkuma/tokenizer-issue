package main

import (
	"github.com/anush008/fastembed-go"
)

type env struct {
	EmbeddingBatchSize int
	// EmbeddingModel must be one of
	// AllMiniLML6V2 EmbeddingModel = "fast-all-MiniLM-L6-v2"
	// BGEBaseEN     EmbeddingModel = "fast-bge-base-en"
	// BGEBaseENV15  EmbeddingModel = "fast-bge-base-en-v1.5"
	// BGESmallEN    EmbeddingModel = "fast-bge-small-en"
	// BGESmallENV15 EmbeddingModel = "fast-bge-small-en-v1.5" <-- this is default
	// BGESmallZH    EmbeddingModel = "fast-bge-small-zh-v1.5"
	EmbeddingModel fastembed.EmbeddingModel
	// InputSeqMaxLen is the max length for the embedding api
	InputSeqMaxLen int
}
