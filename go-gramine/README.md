# Intel® Trust Authority Go Gramine Adapter

The **go-gramine** adapter enables a confidential computing client running in a Gramine Shielded Container (GSC) to collect a quote for attestation by Intel Trust Authority. The go-gramine adapter is used with the [**go-connector**](https://github.com/intel/trustauthority-client-for-go/tree/main/go-connector) to request an attestation token. 

## Requirements

- Use **Go 1.19 or newer**. See [https://go.dev/doc/install](https://go.dev/doc/install) for installation of Go.

## Usage

Create a new Go Gramine adapter, then use the adapter to collect quote from SGX enabled platform. 

```go
import "github.com/arvind5/gsc-sample-app/go-gramine"

adapter, err := gramine.NewEvidenceAdapter(enclaveHeldData)
if err != nil {
    return err
}

evidence, err := adapter.CollectEvidence(nonce)
if err != nil {
    return err
}
```
