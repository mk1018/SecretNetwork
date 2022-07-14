package remote_attestation

import (
	"fmt"
	"strconv"
)

const rootIntelPEM = `-----BEGIN CERTIFICATE-----
MIIFSzCCA7OgAwIBAgIJANEHdl0yo7CUMA0GCSqGSIb3DQEBCwUAMH4xCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTEUMBIGA1UEBwwLU2FudGEgQ2xhcmExGjAYBgNV
BAoMEUludGVsIENvcnBvcmF0aW9uMTAwLgYDVQQDDCdJbnRlbCBTR1ggQXR0ZXN0
YXRpb24gUmVwb3J0IFNpZ25pbmcgQ0EwIBcNMTYxMTE0MTUzNzMxWhgPMjA0OTEy
MzEyMzU5NTlaMH4xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEUMBIGA1UEBwwL
U2FudGEgQ2xhcmExGjAYBgNVBAoMEUludGVsIENvcnBvcmF0aW9uMTAwLgYDVQQD
DCdJbnRlbCBTR1ggQXR0ZXN0YXRpb24gUmVwb3J0IFNpZ25pbmcgQ0EwggGiMA0G
CSqGSIb3DQEBAQUAA4IBjwAwggGKAoIBgQCfPGR+tXc8u1EtJzLA10Feu1Wg+p7e
LmSRmeaCHbkQ1TF3Nwl3RmpqXkeGzNLd69QUnWovYyVSndEMyYc3sHecGgfinEeh
rgBJSEdsSJ9FpaFdesjsxqzGRa20PYdnnfWcCTvFoulpbFR4VBuXnnVLVzkUvlXT
L/TAnd8nIZk0zZkFJ7P5LtePvykkar7LcSQO85wtcQe0R1Raf/sQ6wYKaKmFgCGe
NpEJUmg4ktal4qgIAxk+QHUxQE42sxViN5mqglB0QJdUot/o9a/V/mMeH8KvOAiQ
byinkNndn+Bgk5sSV5DFgF0DffVqmVMblt5p3jPtImzBIH0QQrXJq39AT8cRwP5H
afuVeLHcDsRp6hol4P+ZFIhu8mmbI1u0hH3W/0C2BuYXB5PC+5izFFh/nP0lc2Lf
6rELO9LZdnOhpL1ExFOq9H/B8tPQ84T3Sgb4nAifDabNt/zu6MmCGo5U8lwEFtGM
RoOaX4AS+909x00lYnmtwsDVWv9vBiJCXRsCAwEAAaOByTCBxjBgBgNVHR8EWTBX
MFWgU6BRhk9odHRwOi8vdHJ1c3RlZHNlcnZpY2VzLmludGVsLmNvbS9jb250ZW50
L0NSTC9TR1gvQXR0ZXN0YXRpb25SZXBvcnRTaWduaW5nQ0EuY3JsMB0GA1UdDgQW
BBR4Q3t2pn680K9+QjfrNXw7hwFRPDAfBgNVHSMEGDAWgBR4Q3t2pn680K9+Qjfr
NXw7hwFRPDAOBgNVHQ8BAf8EBAMCAQYwEgYDVR0TAQH/BAgwBgEB/wIBADANBgkq
hkiG9w0BAQsFAAOCAYEAeF8tYMXICvQqeXYQITkV2oLJsp6J4JAqJabHWxYJHGir
IEqucRiJSSx+HjIJEUVaj8E0QjEud6Y5lNmXlcjqRXaCPOqK0eGRz6hi+ripMtPZ
sFNaBwLQVV905SDjAzDzNIDnrcnXyB4gcDFCvwDFKKgLRjOB/WAqgscDUoGq5ZVi
zLUzTqiQPmULAQaB9c6Oti6snEFJiCQ67JLyW/E83/frzCmO5Ru6WjU4tmsmy8Ra
Ud4APK0wZTGtfPXU7w+IBdG5Ez0kE1qzxGQaL4gINJ1zMyleDnbuS8UicjJijvqA
152Sq049ESDz+1rRGc2NVEqh1KaGXmtXvqxXcTB+Ljy5Bw2ke0v8iGngFBPqCTVB
3op5KBG3RjbF6RRSzwzuWfL7QErNC8WEy5yDVARzTA5+xmBc388v9Dm21HGfcC8O
DD+gT9sSpssq0ascmvH49MOgjt1yoysLtdCtJW/9FZpoOypaHx0R+mJTLwPXVMrv
DaVzWh5aiEx+idkSGMnX
-----END CERTIFICATE-----`

type QuoteReport struct {
	ID                    string   `json:"id"`
	Timestamp             string   `json:"timestamp"`
	Version               int      `json:"version"`
	IsvEnclaveQuoteStatus string   `json:"isvEnclaveQuoteStatus"`
	PlatformInfoBlob      string   `json:"platformInfoBlob"`
	IsvEnclaveQuoteBody   string   `json:"isvEnclaveQuoteBody"`
	AdvisoryIDs           []string `json:"advisoryIDs"`
}

type Certificate []byte

//TODO: add more origin field if needed
type QuoteReportData struct {
	version    int
	signType   int
	reportBody QuoteReportBody
}

//TODO: add more origin filed if needed
type QuoteReportBody struct {
	mrEnclave  string
	mrSigner   string
	reportData string
}

type EndorsedAttestationReport struct {
	Report      []byte `json:"report"`
	Signature   []byte `json:"signature"`
	SigningCert []byte `json:"signing_cert"`
}

type PlatformInfoBlob struct {
	SgxEpidGroupFlags       uint8             `json:"sgx_epid_group_flags"`
	SgxTcbEvaluationFlags   uint32            `json:"sgx_tcb_evaluation_flags"`
	PseEvaluationFlags      uint32            `json:"pse_evaluation_flags"`
	LatestEquivalentTcbPsvn string            `json:"latest_equivalent_tcb_psvn"`
	LatestPseIsvsvn         string            `json:"latest_pse_isvsvn"`
	LatestPsdaSvn           string            `json:"latest_psda_svn"`
	Xeid                    uint32            `json:"xeid"`
	Gid                     uint32            `json:"gid"`
	SgxEc256SignatureT      SGXEC256Signature `json:"sgx_ec256_signature_t"`
}

type SGXEC256Signature struct {
	Gx string `json:"gx"`
	Gy string `json:"gy"`
}

// directly read from []byte
func parseReport(quoteBytes []byte, quoteHex string) *QuoteReportData {
	qrData := &QuoteReportData{reportBody: QuoteReportBody{}}
	qrData.version = int(quoteBytes[0])
	qrData.signType = int(quoteBytes[2])
	qrData.reportBody.mrEnclave = quoteHex[224:288]
	qrData.reportBody.mrSigner = quoteHex[352:416]
	qrData.reportBody.reportData = quoteHex[736:864]
	return qrData
}

// directly read from []byte
func parsePlatform(piBlobByte []byte) *PlatformInfoBlob {
	piBlob := &PlatformInfoBlob{SgxEc256SignatureT: SGXEC256Signature{}}
	piBlob.SgxEpidGroupFlags = uint8(piBlobByte[0])
	piBlob.SgxTcbEvaluationFlags = computeDec(piBlobByte[1:3])
	piBlob.PseEvaluationFlags = computeDec(piBlobByte[3:5])
	piBlob.LatestEquivalentTcbPsvn = bytesToString(piBlobByte[5:23])
	piBlob.LatestPseIsvsvn = bytesToString(piBlobByte[23:25])
	piBlob.LatestPsdaSvn = bytesToString(piBlobByte[25:29])
	piBlob.Xeid = computeDec(piBlobByte[29:33])
	piBlob.Gid = computeDec(piBlobByte[33:37])
	piBlob.SgxEc256SignatureT.Gx = bytesToString(piBlobByte[37:69])
	piBlob.SgxEc256SignatureT.Gy = bytesToString(piBlobByte[69:])

	return piBlob
}

func computeDec(piBlobSlice []byte) uint32 {
	var hexString string
	for i := len(piBlobSlice) - 1; i >= 0; i-- {
		hexString += fmt.Sprintf("%02x", piBlobSlice[i])
	}
	s, _ := strconv.ParseInt(hexString, 16, 32)

	return uint32(s)
}

func bytesToString(byteSlice []byte) string {
	var byteString string
	for i := 0; i < len(byteSlice); i++ {
		byteString += strconv.Itoa(int(byteSlice[i])) + ", "
	}
	byteString = "[" + byteString[:len(byteString)-2] + "]"
	return byteString
}
