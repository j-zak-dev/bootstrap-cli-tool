package utils

import (
	"fmt"
)

func CreateDomain(path string) error {

	var domainLayer = path + "/domain"
	var layerServices = "/services"
	var layerInterfaces = "/interfaces"
	var layerAggregates = "/aggregates"
	var layerVObjects = "/value-objects"

	serviceError := CreateDir((domainLayer + layerServices))
	if serviceError != nil {
		fmt.Println("Error creating layer: " + domainLayer + layerServices)
		return serviceError
	}
	vObjectError := CreateDir((domainLayer + layerVObjects))
	if vObjectError != nil {
		fmt.Println("Error creating layer: " + domainLayer + layerVObjects)
		return vObjectError
	}
	aggregateError := CreateDir((domainLayer + layerAggregates))
	if aggregateError != nil {
		fmt.Println("Error creating layer: " + domainLayer + layerAggregates)
		return aggregateError
	}
	interfaceError := CreateDir((domainLayer + layerInterfaces))
	if interfaceError != nil {
		fmt.Println("Error creating layer: " + domainLayer + layerInterfaces)
		return interfaceError
	}

	return nil
}
