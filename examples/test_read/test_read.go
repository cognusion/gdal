package main

import (
	"fmt"
	"flag"
	"github.com/lukeroth/gdal"

)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Usage: test_tiff [filename]\n")
		return
	}
	fmt.Printf("Filename: %s\n", filename)
	
	/*
	fmt.Printf("Loading driver\n")
	driver, err := gdal.GetDriverByName("GTiff")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	*/
	
	dataset, err := gdal.Open(filename,gdal.ReadOnly)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dataset.Close()
	
	fmt.Printf("Projection: %s\n", dataset.Projection())
	fmt.Printf("X: %d Y: %d\n", dataset.RasterXSize(), dataset.RasterYSize())
	fmt.Printf("Bands: %d\n", dataset.RasterCount())
	
	if dataset.RasterCount() > 0 {
		fmt.Printf("Loading first band...\n")
		rasterband := dataset.RasterBand(1)
		fmt.Printf("\tData type: %s\n", rasterband.RasterDataType())
		
		nodata, ndvalid := rasterband.NoDataValue()
		if ndvalid {
			fmt.Printf("\tNoData: %e, %f\n", nodata, nodata)
		}
		bx,by := rasterband.BlockSize()
		fmt.Printf("\tBlocksize X: %d, Y: %d\n", bx, by)
	}
	

}
