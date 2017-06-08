package qrcode

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	imagePath = "/qrcode/out.jpg"
)

const (
	createApi     = "https://api.qrserver.com/v1/create-qr-code/?size=350x350&data="
	ipLocalPrefix = "192.168.1."
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(dir, "src") {
		imagePath = "." + imagePath
	} else {
		imagePath = "./src" + imagePath
	}
}

func QRCode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "image/jpeg")

	file, err := getImage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(file)
}

func getImage() ([]byte, error) {
	err := downloadQRCode()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(imagePath)
}

func downloadQRCode() error {
	ip := getLocalIpAddress()
	err := createQRCodeImage(ip)
	if err != nil {
		return err
	}

	return nil
}

func getLocalIpAddress() string {
	addresses, _ := net.InterfaceAddrs()
	for _, address := range addresses {
		if strings.Contains(address.String(), ipLocalPrefix) {
			ip := removeMask(address.String())
			return ip
		}
	}

	return ""
}

func removeMask(ip string) string {
	value := strings.Split(ip, "/")
	return value[0]
}

func createQRCodeImage(ip string) error {
	url := createApi + ip
	response, err := http.Get(url)
	if err != nil {
		return errors.New("Error while creating QR Code image")
	}

	defer response.Body.Close()

	return createImage(response.Body)
}

func createImage(body io.ReadCloser) error {
	file, err := os.Create(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
