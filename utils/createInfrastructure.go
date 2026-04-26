package utils

import (
	"fmt"
)

func CreateInfrastructure(path string) error {

	var infrastructureLayer = path + "/infrastructure"
	var layerServices = "/services"
	var layerInterfaces = "/interfaces"
	var layerDTOs = "/DTOs"
	var layerMappers = "/mappers"

	serviceError := CreateDir((infrastructureLayer + layerServices))
	if serviceError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerServices)
		return serviceError
	}
	interfaceError := CreateDir((infrastructureLayer + layerInterfaces))
	if interfaceError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerInterfaces)
		return interfaceError
	}
	dtoError := CreateDir((infrastructureLayer + layerDTOs))
	if dtoError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerDTOs)
		return dtoError
	}
	mapperError := CreateDir((infrastructureLayer + layerMappers))
	if mapperError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerMappers)
		return mapperError
	}

	return nil
}
