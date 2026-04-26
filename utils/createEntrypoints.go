package utils

import (
	"fmt"
)

func CreateEntrypoints(path string) error {

	var infrastructureLayer = path + "/entrypoints"
	var layerDTOs = "/DTOs"
	var layerMappers = "/mappers"
	var layerRouters = "/routers"
	var layerAPIs = "/APIs"

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
	routerError := CreateDir((infrastructureLayer + layerRouters))
	if routerError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerRouters)
		return routerError
	}
	apiError := CreateDir((infrastructureLayer + layerAPIs))
	if apiError != nil {
		fmt.Println("Error creating layer: " + infrastructureLayer + layerAPIs)
		return apiError
	}

	return nil
}
