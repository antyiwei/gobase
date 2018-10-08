package M18_templatemethod

import "testing"

func ExampleNewHTTPDownloader() {
	var downloader Downloader = NewHTTPDownloader()

	downloader.Download("http://example.com/abc.zip")
}

func ExampleNewFTPDownloader() {
	var downloader Downloader = NewFTPDownloader()

	downloader.Download("ftp://example.com/abc.zip")
}

func TestNewFTPDownloader(t *testing.T) {
	ExampleNewFTPDownloader()
}

func TestNewHTTPDownloader(t *testing.T) {
	ExampleNewHTTPDownloader()
}
