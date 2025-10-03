package main

import (
	"sort"
)

type StackCMSComponent struct {
	Position    int
	RowNumber   int
	Alignment   string
	RowSpan     int
	ColumnSpan  int
}

func placeCmsComponentInGrid(rowNumber int, alignment string, col int, columnSpan int, offset int) int {
	position := 0
	if alignment == "Left" {
		position = (rowNumber-1)*col - offset
	} else {
		position = rowNumber*col - columnSpan - offset
	}
	return position + 1
}

func fetchPositionForCMSComponent(stackCMSComponents []StackCMSComponent, col int) []StackCMSComponent {
	offset := 0
	oldEndRowNumber := 0
	var result []StackCMSComponent

	containsRowNumberAlignment := false
	for _, stackCMScomponent := range stackCMSComponents {
		if stackCMScomponent.Position == 0 && stackCMScomponent.RowNumber != 0 && stackCMScomponent.Alignment != "" {
			containsRowNumberAlignment = true
			break
		}
	}

	if containsRowNumberAlignment {
		filteredComponents := []StackCMSComponent{}
		for _, stackCMSComponent := range stackCMSComponents {
			if stackCMSComponent.Position == 0 && stackCMSComponent.RowNumber != 0 && stackCMSComponent.Alignment != "" {
				filteredComponents = append(filteredComponents, stackCMSComponent)
			}
		}

		sort.Slice(filteredComponents, func(i, j int) bool {
			if filteredComponents[i].RowNumber == filteredComponents[j].RowNumber {
				if filteredComponents[i].Alignment == filteredComponents[j].Alignment {
					return false
				}
				return filteredComponents[i].Alignment == "Left"
			}
			return filteredComponents[i].RowNumber < filteredComponents[j].RowNumber
		})

		for _, component := range filteredComponents {
			newComponent := component
			if newComponent.ColumnSpan > col {
				newComponent.ColumnSpan = col
			}
			if newComponent.RowNumber > oldEndRowNumber {
				newComponent.Position  = placeCmsComponentInGrid(newComponent.RowNumber, newComponent.Alignment, col, newComponent.ColumnSpan, offset)
				result = append(result, newComponent)
				offset += newComponent.ColumnSpan * newComponent.RowSpan
				oldEndRowNumber = newComponent.RowNumber + newComponent.RowSpan - 1
			}
		}
	} else {
		for _, component := range stackCMSComponents {
			if component.Position != 0 {
				result = append(result, component)
			}
		}
	}
	return result
}

