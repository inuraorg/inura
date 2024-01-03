package reassemble

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sort"

	"github.com/inuraorg/inura/inura-node/chaincfg"
	"github.com/inuraorg/inura/inura-node/rollup"

	"github.com/inuraorg/inura/inura-node/cmd/batch_decoder/fetch"
	"github.com/inuraorg/inura/inura-node/rollup/derive"
	"github.com/inuraorg/inura/inura-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

type ChannelWithMetadata struct {
	ID             derive.ChannelID    `json:"id"`
	IsReady        bool                `json:"is_ready"`
	InvalidFrames  bool                `json:"invalid_frames"`
	InvalidBatches bool                `json:"invalid_batches"`
	Frames         []FrameWithMetadata `json:"frames"`
	Batches        []derive.BatchData  `json:"batches"`
}

type FrameWithMetadata struct {
	TxHash         common.Hash  `json:"transaction_hash"`
	InclusionBlock uint64       `json:"inclusion_block"`
	Timestamp      uint64       `json:"timestamp"`
	BlockHash      common.Hash  `json:"block_hash"`
	Frame          derive.Frame `json:"frame"`
}

type Config struct {
	BatchInbox   common.Address
	InDirectory  string
	OutDirectory string
}

func LoadFrames(directory string, inbox common.Address) []FrameWithMetadata {
	txns := loadTransactions(directory, inbox)
	// Sort first by block number then by transaction index inside the block number range.
	// This is to match the order they are processed in derivation.
	sort.Slice(txns, func(i, j int) bool {
		if txns[i].BlockNumber == txns[j].BlockNumber {
			return txns[i].TxIndex < txns[j].TxIndex
		} else {
			return txns[i].BlockNumber < txns[j].BlockNumber
		}

	})
	return transactionsToFrames(txns)
}

// Channels loads all transactions from the given input directory that are submitted to the
// specified batch inbox and then re-assembles all channels & writes the re-assembled channels
// to the out directory.
func Channels(config Config) {
	if err := os.MkdirAll(config.OutDirectory, 0750); err != nil {
		log.Fatal(err)
	}
	frames := LoadFrames(config.InDirectory, config.BatchInbox)
	framesByChannel := make(map[derive.ChannelID][]FrameWithMetadata)
	for _, frame := range frames {
		framesByChannel[frame.Frame.ID] = append(framesByChannel[frame.Frame.ID], frame)
	}
	cfg := chaincfg.Mainnet
	for id, frames := range framesByChannel {
		ch := processFrames(cfg, id, frames)
		filename := path.Join(config.OutDirectory, fmt.Sprintf("%s.json", id.String()))
		if err := writeChannel(ch, filename); err != nil {
			log.Fatal(err)
		}
	}
}

func writeChannel(ch ChannelWithMetadata, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	return enc.Encode(ch)
}

func processFrames(cfg *rollup.Config, id derive.ChannelID, frames []FrameWithMetadata) ChannelWithMetadata {
	ch := derive.NewChannel(id, eth.L1BlockRef{Number: frames[0].InclusionBlock})
	invalidFrame := false

	for _, frame := range frames {
		if ch.IsReady() {
			fmt.Printf("Channel %v is ready despite having more frames\n", id.String())
			invalidFrame = true
			break
		}
		if err := ch.AddFrame(frame.Frame, eth.L1BlockRef{Number: frame.InclusionBlock}); err != nil {
			fmt.Printf("Error adding to channel %v. Err: %v\n", id.String(), err)
			invalidFrame = true
		}
	}

	var batches []derive.BatchData
	invalidBatches := false
	if ch.IsReady() {
		br, err := derive.BatchReader(ch.Reader())
		if err == nil {
			for batch, err := br(); err != io.EOF; batch, err = br() {
				if err != nil {
					fmt.Printf("Error reading batch for channel %v. Err: %v\n", id.String(), err)
					invalidBatches = true
				} else {
					batches = append(batches, *batch)
				}
			}
		} else {
			fmt.Printf("Error creating batch reader for channel %v. Err: %v\n", id.String(), err)
		}
	} else {
		fmt.Printf("Channel %v is not ready\n", id.String())
	}

	return ChannelWithMetadata{
		ID:             id,
		Frames:         frames,
		IsReady:        ch.IsReady(),
		InvalidFrames:  invalidFrame,
		InvalidBatches: invalidBatches,
		Batches:        batches,
	}
}

func transactionsToFrames(txns []fetch.TransactionWithMetadata) []FrameWithMetadata {
	var out []FrameWithMetadata
	for _, tx := range txns {
		for _, frame := range tx.Frames {
			fm := FrameWithMetadata{
				TxHash:         tx.Tx.Hash(),
				InclusionBlock: tx.BlockNumber,
				BlockHash:      tx.BlockHash,
				Timestamp:      tx.BlockTime,
				Frame:          frame,
			}
			out = append(out, fm)
		}
	}
	return out
}

// if inbox is the zero address, it will load all frames
func loadTransactions(dir string, inbox common.Address) []fetch.TransactionWithMetadata {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var out []fetch.TransactionWithMetadata
	for _, file := range files {
		f := path.Join(dir, file.Name())
		txm := loadTransactionsFile(f)
		if (inbox == common.Address{} || txm.InboxAddr == inbox) && txm.ValidSender {
			out = append(out, txm)
		}
	}
	return out
}

func loadTransactionsFile(file string) fetch.TransactionWithMetadata {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var txm fetch.TransactionWithMetadata
	if err := dec.Decode(&txm); err != nil {
		log.Fatalf("Failed to decode %v. Err: %v\n", file, err)
	}
	return txm
}
