package types

import (
	fmt "fmt"
	"testing"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/minio/sha256-simd"

	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/stretchr/testify/assert"
)

func TestBlockHash(t *testing.T) {
	blockHash := func(block *Block) []byte {
		header := block.Header
		digest := sha256.New()
		writeBlockHeaderOld(digest, header)
		return digest.Sum(nil)
	}

	txIn := make([]*Tx, 0)
	block := NewBlock(nil, nil, nil, txIn, nil, 0)

	h1 := blockHash(block)
	h2 := block.calculateBlockHash()

	assert.Equal(t, h1, h2)
}

func genKeyPair(assert *assert.Assertions) (crypto.PrivKey, crypto.PubKey) {
	privKey, pubKey, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	assert.Nil(err)

	return privKey, pubKey
}

func genSign(assert *assert.Assertions, block *Block, privKey crypto.PrivKey) []byte {
	msg, err := block.Header.bytesForDigest()
	assert.Nil(err)

	sig, err := privKey.Sign(msg)
	assert.Nil(err)

	return sig
}

func TestBlockSignBasic(t *testing.T) {
	signAssert := assert.New(t)

	message := func(block *Block) []byte {
		msg, err := block.Header.bytesForDigest()
		signAssert.Nil(err)

		return msg
	}

	sign := func(block *Block, privKey crypto.PrivKey) []byte {
		sig, err := privKey.Sign(message(block))
		signAssert.Nil(err)

		return sig
	}

	verify := func(block *Block, sig []byte, pubKey crypto.PubKey) bool {
		valid, err := pubKey.Verify(message(block), sig)
		signAssert.Nil(err)

		return valid
	}

	block := NewBlock(nil, nil, nil, make([]*Tx, 0), nil, 0)

	privKey, pubKey := genKeyPair(signAssert)
	sig := sign(block, privKey)
	valid := verify(block, sig, pubKey)
	signAssert.True(valid)
}

func TestBlockSign(t *testing.T) {
	signAssert := assert.New(t)
	block := NewBlock(nil, nil, nil, make([]*Tx, 0), nil, 0)

	privKey, _ := genKeyPair(signAssert)
	signAssert.Nil(block.Sign(privKey))

	sig := genSign(signAssert, block, privKey)
	signAssert.NotNil(sig)

	signAssert.Equal(sig, block.Header.Sign)

	valid, err := block.VerifySign()
	signAssert.Nil(err)
	signAssert.True(valid)
}

func TestMovingAverage(t *testing.T) {
	size := int64(10)
	mv := NewMovingAverage(int(size))

	assert.Equal(t, int64(0), mv.calculateAvg())

	var sum = int64(0)
	var expAvg int64
	for i := int64(1); i <= size; i++ {
		avg := mv.Add(int64(i))
		assert.Equal(t, i, int64(mv.count))

		sum += i
		expAvg = sum / i
		assert.Equal(t, expAvg, avg)
	}

	for i := int64(size + 1); i <= (size + 10); i++ {
		avg := mv.Add(int64(i))
		assert.Equal(t, size, int64(mv.count))

		sum = sum + i - (i - size)
		expAvg = sum / size
		assert.Equal(t, expAvg, avg)
	}
}

func TestUpdateAvgVerifyTime(t *testing.T) {
	var size = int64(10)
	avgTime := NewAvgTime(int(size))

	val := avgTime.Get()

	assert.Equal(t, int64(0), int64(val))

	var sum = int64(0)
	var expAvg int64

	for i := int64(1); i <= size; i++ {
		avgTime.UpdateAverage(time.Duration(i))
		newAvg := avgTime.Get()

		sum += i
		expAvg = sum / i
		assert.Equal(t, expAvg, int64(newAvg))
	}

	for i := int64(size + 1); i <= (size + 10); i++ {
		avgTime.UpdateAverage(time.Duration(i))
		newAvg := avgTime.Get()

		sum = sum + i - (i - size)
		expAvg = sum / size
		assert.Equal(t, expAvg, int64(newAvg))
	}
}

func TestBlockHeader(t *testing.T) {
	a := assert.New(t)

	chainID := "01234567890123456789"
	dig := sha256.New()
	dig.Write([]byte(chainID))
	h := dig.Sum(nil)

	addr, err := DecodeAddress("AmLquXjQSiDdR8FTDK78LJ16Ycrq3uNL6NQuiwqXRCGw9Riq2DE4")
	a.Nil(err, "invalid coinbase account")

	priv, pub := genKeyPair(a)

	block := &Block{
		Header: &BlockHeader{
			ChainID:          []byte(chainID),
			PrevBlockHash:    h,
			BlockNo:          10,
			Timestamp:        time.Now().UnixNano(),
			BlocksRootHash:   h,
			TxsRootHash:      h,
			ReceiptsRootHash: h,
			Confirms:         23,
			CoinbaseAccount:  addr,
		},
	}
	err = block.setPubKey(pub)
	a.Nil(err, "PubKey set failed")

	err = block.Sign(priv)
	a.Nil(err, "block sign failed")

	fmt.Println("block id:", block.ID())
	fmt.Println("block size:", proto.Size(block))
	fmt.Println("signature len:", len(block.Header.Sign))
}
