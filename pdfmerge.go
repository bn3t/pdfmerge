package main

import (
	"fmt"
	"os"

	unicommon "github.com/unidoc/unidoc/common"
	//unilicense "github.com/unidoc/unidoc/license"
	unipdf "github.com/unidoc/unidoc/pdf"
)

func initUnidoc() error {
	// err := unilicense.SetLicenseKey("")
	// if err != nil {
	// 	return err
	// }
	unicommon.SetLogger(unicommon.DummyLogger{})
	return nil
}

func mergePdf(output string, input1 string, input2 string) error {
	pdfWriter := unipdf.NewPdfWriter()
	finput1, err := os.Open(input1)
	if err != nil {
		return err
	}
	defer finput1.Close()
	finput2, err := os.Open(input2)
	if err != nil {
		return err
	}
	defer finput2.Close()
	pdfReader1, err := unipdf.NewPdfReader(finput1)
	if err != nil {
		return err
	}
	pdfReader2, err := unipdf.NewPdfReader(finput2)
	if err != nil {
		return err
	}
	numPages1, err := pdfReader1.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Println("Number of pages in input1: ", numPages1)
	numPages2, err := pdfReader2.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Println("Number of pages in input2: ", numPages2)

	for i := 0; i < numPages1; i++ {

		page1, err := pdfReader1.GetPage(i + 1)
		if err != nil {
			return err
		}
		err = pdfWriter.AddPage(page1)
		if err != nil {
			return err
		}

		if numPages1-i-1 < numPages2 {
			page2, err := pdfReader2.GetPage(numPages1 - i)
			err = pdfWriter.AddPage(page2)
			if err != nil {
				return err
			}
		}
	}

	fWrite, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Wrong number of arguments")
		return
	}
	err := initUnidoc()
	if err != nil {
		fmt.Println("Error initializing unidoc", err)
		return
	}

	err = mergePdf(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println("Error merging: ", err)
		return
	}
}
