package fs

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type CompressionType int

const (
	TarGz CompressionType = iota
	Zip
)

func (c CompressionType) String() string {
	switch c {
	case TarGz:
		return "tar.gz"
	case Zip:
		return "zip"
	default:
		return "unknown"
	}
}

func Compress(srcDir string, out io.Writer, t CompressionType) error {
	switch t {
	case TarGz:
		return tarGzCompress(srcDir, out)
	case Zip:
		return zipCompress(srcDir, out)
	default:
		return fmt.Errorf("unknown compression type, %s", t)
	}
}

func tarGzCompress(src string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	tw := tar.NewWriter(gw)

	err := filepath.Walk(src, func(file string, fi fs.FileInfo, _ error) error {
		header, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, file)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		header.Name = filepath.ToSlash(relPath)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			io.Copy(tw, data)
			data.Close()
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err = tw.Close(); err != nil {
		return err
	}

	if err = gw.Close(); err != nil {
		return err
	}

	return nil
}

func zipCompress(src string, buf io.Writer) error {
	w := zip.NewWriter(buf)

	err := filepath.Walk(src, func(file string, fi fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, file)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}
		header.Name = filepath.ToSlash(relPath)

		f, err := w.CreateHeader(header)
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			io.Copy(f, data)
			data.Close()
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
