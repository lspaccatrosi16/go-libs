package fs

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Decompress(src io.Reader, outDir string, t CompressionType) error {
	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	switch t {
	case TarGz:
		return tarGzDecompress(src, outDir)
	case Zip:
		b := bytes.NewBuffer(nil)
		size, _ := io.Copy(b, src)
		srcReader := bytes.NewReader(b.Bytes())

		return zipDecompress(srcReader, size, outDir)
	default:
		return fmt.Errorf("unknown compression type, %s", t)
	}
}

func tarGzDecompress(src io.Reader, out string) error {
	gr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		target := filepath.Join(out, header.Name)
		if header.FileInfo().IsDir() {
			err := os.MkdirAll(target, 0755)
			if err != nil {
				return err
			}
		} else {
			f, err := os.Create(target)
			if err != nil {
				return err
			}

			io.Copy(f, tr)
			f.Close()
		}
	}

	return gr.Close()
}

func zipDecompress(src io.ReaderAt, size int64, out string) error {
	r, err := zip.NewReader(src, size)
	if err != nil {
		return err
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(filepath.Join(out, f.Name), 0755)
			if err != nil {
				return err
			}
		} else {
			dst, err := os.Create(filepath.Join(out, f.Name))
			if err != nil {
				return err
			}
			io.Copy(dst, rc)
			dst.Close()
		}

		rc.Close()
	}

	return nil
}
