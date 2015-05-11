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
		fmt.Printf("Usage: test_read [filename]\n")
		return
	}
	fmt.Printf("Filename: %s\n", filename)
	
	fmt.Printf("GDAL v%d.%d.%d Build %d Date %d Name %s\n", 
		gdal.VERSION_MAJOR, gdal.VERSION_MINOR, gdal.VERSION_REV, gdal.VERSION_BUILD, 
		gdal.RELEASE_DATE, gdal.RELEASE_NAME)
		
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

