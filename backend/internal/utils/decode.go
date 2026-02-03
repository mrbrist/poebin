package utils

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func RawToGo(raw string) (*PathOfBuilding, error) {
	b, err := decompress(raw)
	if err != nil {
		return nil, err
	}

	var data PathOfBuilding
	if err := xml.Unmarshal([]byte(b), &data); err != nil {
		return nil, fmt.Errorf("XML unmarshal error: %w", err)
	}

	return &data, nil
}

func decompress(data string) (string, error) {
	decoded, err := decode(data)
	if err != nil {
		return "", err
	}

	return deflate(decoded)
}

func decode(data string) ([]byte, error) {
	data = strings.TrimSpace(data)
	decoded, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("Base64 decode error: %w", err)
	}
	return decoded, nil
}

func deflate(inp []byte) (string, error) {
	b := bytes.NewReader(inp)
	r, err := zlib.NewReader(b)
	if err != nil {
		return "", fmt.Errorf("zlib decompression error: %w", err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read decompressed data error: %w", err)
	}

	// Try UTF-8 first
	str := string(buf)
	if !isValidUTF8(buf) {
		// Fallback to Windows-1252
		decoder := charmap.Windows1252.NewDecoder()
		strDecoded, err := decoder.String(str)
		if err != nil {
			return "", fmt.Errorf("string decode error: %w", err)
		}
		return strDecoded, nil
	}
	return str, nil
}

func isValidUTF8(data []byte) bool {
	for len(data) > 0 {
		r, size := decodeRune(data)
		if r == '\uFFFD' && size == 1 {
			return false
		}
		data = data[size:]
	}
	return true
}

func decodeRune(p []byte) (r rune, size int) {
	if len(p) == 0 {
		return 0, 0
	}
	r, size = rune(p[0]), 1
	if r < 0x80 {
		return r, size
	}
	return '\uFFFD', 1
}
