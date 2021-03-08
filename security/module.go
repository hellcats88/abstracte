package security

import (
	"crypto"
	"crypto/cipher"
	"crypto/elliptic"

	"github.com/hellcats88/abstracte/runtime"
)

// GenerateRSAKeyPairReq contains all parameters used to generate an RSA key pair
type GenerateRSAKeyPairReq struct {
	Alias string
	Bits  int
}

// GenerateECDSAKeyPairReq contains all parameters used to generate an ECDSA key pair
type GenerateECDSAKeyPairReq struct {
	Alias string
	Bits  elliptic.Curve
}

// GenerateAESKeyReq contains all parameters used to generate a symmetric AES key
type GenerateAESKeyReq struct {
	Alias string
	Bits  int
}

type CapabilityResp struct {
	Name string
	Len  []int
}

type CapabilitiesResp struct {
	Asymmetric []CapabilityResp
	Symmetric  []CapabilityResp
}

// SecureModule abstract the backend implementation of the Encryption managemend
type SecureModule interface {
	GenerateRSAKeyPair(ctx runtime.Context, req GenerateRSAKeyPairReq) (crypto.PublicKey, error)
	GenerateECDSAKeyPair(ctx runtime.Context, req GenerateECDSAKeyPairReq) (crypto.PublicKey, error)
	Signer(ctx runtime.Context, alias string) (crypto.Signer, error)
	Block(ctx runtime.Context, alias string) (cipher.Block, error)
	GenerateAESKey(ctx runtime.Context, req GenerateAESKeyReq) (cipher.Block, error)
	Capabilities() CapabilitiesResp
}
