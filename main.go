/*
 *   Copyright (c) 2024 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/arvind5/gsc-sample-app/go-gramine"
	"github.com/intel/trustauthority-client/go-connector"
	"github.com/pkg/errors"
)

const (
	envTrustAuthorityUrl    = "TRUSTAUTHORITY_URL"
	envTrustAuthorityApiUrl = "TRUSTAUTHORITY_API_URL"
	envTrustAuthorityApiKey = "TRUSTAUTHORITY_API_KEY"
)

func main() {
	trustAuthorityUrl := os.Getenv(envTrustAuthorityUrl)
	trustAuthorityApiUrl := os.Getenv(envTrustAuthorityApiUrl)
	trustAuthorityApiKey := os.Getenv(envTrustAuthorityApiKey)

	if trustAuthorityUrl == "" || trustAuthorityApiUrl == "" || trustAuthorityApiKey == "" {
		fmt.Println("Either Trust Authority URL, API URL or API Key is missing in config")
		os.Exit(1)
	}

	cfg := &connector.Config{
		TlsCfg: &tls.Config{
			InsecureSkipVerify: true,
		},
		BaseUrl: trustAuthorityUrl,
		ApiUrl:  trustAuthorityApiUrl,
		ApiKey:  trustAuthorityApiKey,
	}

	trustAuthorityConnector, err := connector.New(cfg)
	if err != nil {
		panic(err)
	}

	_, pubBytes, err := generateKeyPair()
	if err != nil {
		panic(err)
	}

	adapter, err := gramine.NewEvidenceAdapter(pubBytes)
	if err != nil {
		panic(err)
	}

	req := connector.AttestArgs{
		Adapter: adapter,
	}
	resp, err := trustAuthorityConnector.Attest(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nTOKEN: %s\n", string(resp.Token))

	token, err := trustAuthorityConnector.VerifyToken(string(resp.Token))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nCLAIMS: %+v\n", token.Claims)
}

func generateKeyPair() (*rsa.PrivateKey, []byte, error) {
	keyPair, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while generating RSA key pair")
	}

	// Public key format : <exponent:E_SIZE_IN_BYTES><modulus:N_SIZE_IN_BYTES>
	pub := keyPair.PublicKey
	pubBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(pubBytes, uint32(pub.E))
	pubBytes = append(pubBytes, pub.N.Bytes()...)
	return keyPair, pubBytes, nil
}
