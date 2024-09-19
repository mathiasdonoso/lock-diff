package internal

import (
	"fmt"
	"os"
)

type ReportCommandHandler struct {
	service ReporterService
}

func NewHandler(service ReporterService) *ReportCommandHandler {
	return &ReportCommandHandler{
		service: service,
	}
}

func (h *ReportCommandHandler) HandleReport(file1Path, file2Path string) error {
	file1, err := os.Open(file1Path)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return err
	}

	file2, err := os.Open(file2Path)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return err
	}

	err = h.service.GetReport(file1, file2)
	if err != nil {
		return err
	}

	return nil
}
