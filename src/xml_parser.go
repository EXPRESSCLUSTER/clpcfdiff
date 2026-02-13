package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding"
)

// CharsetReader supports various encodings for XML parsing
func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	var decoder *encoding.Decoder
	
	switch strings.ToLower(charset) {
	case "euc-jp":
		decoder = japanese.EUCJP.NewDecoder()
	case "shift_jis", "shift-jis", "sjis":
		decoder = japanese.ShiftJIS.NewDecoder()
	case "iso-2022-jp":
		decoder = japanese.ISO2022JP.NewDecoder()
	case "windows-1252":
		decoder = charmap.Windows1252.NewDecoder()
	case "iso-8859-1":
		decoder = charmap.ISO8859_1.NewDecoder()
	default:
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}
	
	return decoder.Reader(input), nil
}

// XMLNode represents an XML element with its path
type XMLNode struct {
	Name       xml.Name
	Attributes map[string]string
	Path       string
}

// XMLPathValue represents an XML path with its value
type XMLPathValue struct {
	Path  string
	Value string
}

// ExtractXMLPathsAndValues extracts all unique XML paths and their values from the given XML file
func ExtractXMLPathsAndValues(filename string) ([]XMLPathValue, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	// Set charset reader to support various encodings
	decoder.CharsetReader = CharsetReader
	var pathValues []XMLPathValue
	var currentPath []string
	var currentText strings.Builder

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse XML: %w", err)
		}

		switch elem := token.(type) {
		case xml.StartElement:
			// Reset text builder for new element
			currentText.Reset()
			
			// Add element name with attributes to path
			elementName := elem.Name.Local
			if elem.Name.Space != "" {
				// Use namespace prefix if available
				elementName = fmt.Sprintf("ns:%s", elem.Name.Local)
			}
			
			// Create path with attributes for uniqueness
			element := elementName
			if len(elem.Attr) > 0 {
				var attrs []string
				for _, attr := range elem.Attr {
					attrName := attr.Name.Local
					if attr.Name.Space != "" {
						attrName = fmt.Sprintf("%s:%s", attr.Name.Space, attr.Name.Local)
					}
					attrs = append(attrs, fmt.Sprintf("@%s='%s'", attrName, attr.Value))
				}
				// Sort attributes for consistent output
				sort.Strings(attrs)
				element = fmt.Sprintf("%s[%s]", elementName, strings.Join(attrs, ","))
			}
			
			currentPath = append(currentPath, element)

		case xml.CharData:
			// Accumulate text data
			currentText.Write(elem)

		case xml.EndElement:
			if len(currentPath) > 0 {
				// Get current path and text value
				path := "/" + strings.Join(currentPath, "/")
				textValue := strings.TrimSpace(currentText.String())
				
				// Add path-value pair
				pathValues = append(pathValues, XMLPathValue{
					Path:  path,
					Value: textValue,
				})
				
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
	}

	// Remove duplicates and sort
	uniquePathValues := removeDuplicatePathValues(pathValues)
	sort.Slice(uniquePathValues, func(i, j int) bool {
		return uniquePathValues[i].Path < uniquePathValues[j].Path
	})
	
	return uniquePathValues, nil
}

// removeDuplicates removes duplicate strings from a slice
func removeDuplicates(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	
	for _, item := range strSlice {
		if !keys[item] {
			keys[item] = true
			list = append(list, item)
		}
	}
	
	return list
}

// removeDuplicatePathValues removes duplicate XMLPathValue entries
func removeDuplicatePathValues(pathValues []XMLPathValue) []XMLPathValue {
	keys := make(map[string]XMLPathValue)
	
	for _, item := range pathValues {
		// Use path as key, keeping the last value if duplicates exist
		keys[item.Path] = item
	}
	
	var result []XMLPathValue
	for _, value := range keys {
		result = append(result, value)
	}
	
	return result
}