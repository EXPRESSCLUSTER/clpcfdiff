package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <xml_file1> <xml_file2>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Compare two XML files and output differences in CSV format.\n")
		os.Exit(1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	// Extract paths from both XML files
	pathValues1, err := ExtractXMLPathsAndValues(file1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file %s: %v\n", file1, err)
		os.Exit(1)
	}

	pathValues2, err := ExtractXMLPathsAndValues(file2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file %s: %v\n", file2, err)
		os.Exit(1)
	}

	// Compare paths and output differences
	differences := ComparePathValues(pathValues1, pathValues2)
	
	// Output CSV
	err = OutputCSV(differences)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CSV output: %v\n", err)
		os.Exit(1)
	}
}

// PathValueDifference represents a difference between two XML files including values
type PathValueDifference struct {
	Path1  string // Path from first file (empty if not present)
	Value1 string // Value from first file
	Path2  string // Path from second file (empty if not present)
	Value2 string // Value from second file
}

// ComparePathValues compares two sets of XML path-value pairs and returns differences
func ComparePathValues(pathValues1, pathValues2 []XMLPathValue) []PathValueDifference {
	var differences []PathValueDifference
	
	// Create maps for fast lookup
	map1 := make(map[string]XMLPathValue)
	map2 := make(map[string]XMLPathValue)
	
	for _, pathValue := range pathValues1 {
		map1[pathValue.Path] = pathValue
	}
	
	for _, pathValue := range pathValues2 {
		map2[pathValue.Path] = pathValue
	}
	
	// Find all unique paths from both files
	allPaths := make(map[string]bool)
	for _, pathValue := range pathValues1 {
		allPaths[pathValue.Path] = true
	}
	for _, pathValue := range pathValues2 {
		allPaths[pathValue.Path] = true
	}
	
	// Check each path and create differences
	for path := range allPaths {
		pathValue1, exists1 := map1[path]
		pathValue2, exists2 := map2[path]
		
		path1Value := ""
		value1 := ""
		path2Value := ""
		value2 := ""
		
		if exists1 {
			path1Value = pathValue1.Path
			value1 = pathValue1.Value
		}
		if exists2 {
			path2Value = pathValue2.Path
			value2 = pathValue2.Value
		}
		
		// Only add if there's a actual difference (path exists in only one file or values differ)
		if !exists1 || !exists2 || value1 != value2 {
			differences = append(differences, PathValueDifference{
				Path1:  path1Value,
				Value1: value1,
				Path2:  path2Value,
				Value2: value2,
			})
		}
	}
	
	// Sort by File2_Path (3rd column) in ascending order
	sort.Slice(differences, func(i, j int) bool {
		// Empty paths should come last
		if differences[i].Path2 == "" && differences[j].Path2 != "" {
			return false
		}
		if differences[i].Path2 != "" && differences[j].Path2 == "" {
			return true
		}
		return differences[i].Path2 < differences[j].Path2
	})
	
	return differences
}

// OutputCSV outputs the differences in CSV format to stdout
func OutputCSV(differences []PathValueDifference) error {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()
	
	// Write CSV header
	err := writer.Write([]string{"File1_Path", "File1_Value", "File2_Path", "File2_Value"})
	if err != nil {
		return err
	}
	
	// Write differences
	for _, diff := range differences {
		err := writer.Write([]string{diff.Path1, diff.Value1, diff.Path2, diff.Value2})
		if err != nil {
			return err
		}
	}
	
	return nil
}