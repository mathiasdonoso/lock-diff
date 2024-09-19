package internal

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
)

type ReportAdapter struct {
	writer io.Writer
}

func NewReportAdapter(writer io.Writer) *ReportAdapter {
	return &ReportAdapter{
		writer: writer,
	}
}

func (r *ReportAdapter) Print(diff []Diff) error {
	w := tabwriter.NewWriter(r.writer, 1, 1, 1, ' ', 0)

	fmt.Fprintf(w, "%s\t%s\t%s\t\n", "Dependency", "file-1", "file-2")
	for _, v := range diff {
		statusColor := checkDiff(v.VersionLockFile1, v.VersionLockFile2)
		fmt.Fprintf(w, "%s\t%s\t\t\t   %s\t\n", v.Name, printColorText(statusColor, v.VersionLockFile1), printColorText(statusColor, v.VersionLockFile2))
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func checkDiff(v1, v2 string) Color {
	format1 := strings.Split(v1, "-")
	format2 := strings.Split(v2, "-")
	version1 := strings.Split(format1[0], ".")
	version2 := strings.Split(format2[0], ".")

	var v1i []int
	var v2i []int

	for _, v := range version1 {
		v, _ := strconv.Atoi(v)

		v1i = append(v1i, v)
	}

	for _, v := range version2 {
		v, _ := strconv.Atoi(v)

		v2i = append(v2i, v)
	}

	if v1i[0] != v2i[0] {
		return RED
	}

	if v1i[1] != v2i[1] {
		return YELLOW
	}

	return GREEN
}

type Color int

const (
	RED Color = iota
	YELLOW
	GREEN
)

func printColorText(color Color, text string) string {
	if color == RED {
		return fmt.Sprintf("\033[0;31m%s\033[0m", text)
	}

	if color == YELLOW {
		return fmt.Sprintf("\033[0;33m%s\033[0m", text)
	}

	if color == GREEN {
		return fmt.Sprintf("\033[0;32m%s\033[0m", text)
	}

	return text
}
