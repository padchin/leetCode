package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func uploadFile(filePath string, targetURL string, ctx context.Context) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, targetURL, file)
	if err != nil {
		return err
	}

	request.ContentLength = fileInfo.Size()

	client := &http.Client{}
	resp, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	// Specify the path to the file you want to upload
	filePath := "go.mod"

	// Specify the URL of your virtual server
	serverURL := "http://warming-spb.online:8080/upload"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		// Ожидайте нажатия клавиши "Enter" для остановки передачи
		fmt.Println("Press Enter to stop the upload...")
		_, _ = fmt.Scanln()
		cancel() // Отмена загрузки по нажатию клавиши "Enter"
	}()

	err := uploadFile(filePath, serverURL, ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("File uploaded successfully!")
}
