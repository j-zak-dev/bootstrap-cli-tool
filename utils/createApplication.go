package utils

import (
	"fmt"
)

func CreateApplication(path string) error {

	var applicationLayer = path + "/application"
	var layerServices = "/services"
	var layerInterfaces = "/interfaces"
	var layerDTOs = "/DTOs"
	var layerMappers = "/mappers"

	serviceError := CreateDir((applicationLayer + layerServices))
	if serviceError != nil {
		fmt.Println("Error creating layer: " + applicationLayer + layerServices)
		return serviceError
	}
	interfaceError := CreateDir((applicationLayer + layerInterfaces))
	if interfaceError != nil {
		fmt.Println("Error creating layer: " + applicationLayer + layerInterfaces)
		return interfaceError
	}
	dtoError := CreateDir((applicationLayer + layerDTOs))
	if dtoError != nil {
		fmt.Println("Error creating layer: " + applicationLayer + layerDTOs)
		return dtoError
	}
	mapperError := CreateDir((applicationLayer + layerMappers))
	if mapperError != nil {
		fmt.Println("Error creating layer: " + applicationLayer + layerMappers)
		return mapperError
	}

	return nil
}
