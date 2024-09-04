package utils

import (
  "io"
  "mime/multipart"
  "os"
  "path/filepath"
)

func SaveFile(file *multipart.FileHeader) (string, error) {
  dst := filepath.Join("uploads", file.Filename)

  // Create the directory if it doesn't exist
  if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
    return "", err
  }

  // Open the uploaded file
  src, err := file.Open()
  if err != nil {
    return "", err
  }
  defer src.Close()

  // Create the destination file
  out, err := os.Create(dst)
  if err != nil {
    return "", err
  }
  defer out.Close()

  // Copy the uploaded file data to the destination file
  if _, err = io.Copy(out, src); err != nil {
    return "", err
  }

  return dst, nil
}
