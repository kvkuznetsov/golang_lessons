package triangle

import (
	"math"
)

func IsExist(x1_, y1_, x2_, y2_, x3_, y3_ int) bool {
	var x1, y1, x2, y2, x3, y3 float64 = convertPointsToFloat64(x1_, y1_, x2_, y2_, x3_, y3_)
	var edge12Len = getEdgeLen(x1, y1, x2, y2)
	var edge13Len = getEdgeLen(x1, y1, x3, y3)
	var edge23Len = getEdgeLen(x2, y2, x3, y3)
	if edge12Len + edge13Len > edge23Len && edge12Len + edge23Len > edge13Len && edge13Len + edge23Len > edge12Len {
		return true
	} else {
		return false
	}
}

func GetArea(x1_, y1_, x2_, y2_, x3_, y3_ int) float64 {
	var x1, y1, x2, y2, x3, y3 float64 = convertPointsToFloat64(x1_, y1_, x2_, y2_, x3_, y3_)
	var area float64 = 0.5  * math.Abs((x2 - x1) * (y3 - y1) - (x3 - x1) * (y2 - y1))
	return area
}

func IsRectangular(x1_, y1_, x2_, y2_, x3_, y3_ int) bool {
	var x1, y1, x2, y2, x3, y3 float64 = convertPointsToFloat64(x1_, y1_, x2_, y2_, x3_, y3_)
	var edge12Len = getEdgeLen(x1, y1, x2, y2)
	var edge13Len = getEdgeLen(x1, y1, x3, y3)
	var edge23Len = getEdgeLen(x2, y2, x3, y3)
	var hypotenuseLen float64 = math.Max(math.Max(edge12Len, edge13Len), edge23Len)
	var edgesLens [3]float64 = [3]float64{edge12Len, edge13Len, edge23Len}
	var cathetsPowLen float64 = 0

	for _, edgeLen := range edgesLens {
		if edgeLen != hypotenuseLen {
			cathetsPowLen += math.Pow(edgeLen, 2)
		}
	}
	return roundFloat(math.Pow(hypotenuseLen, 2), 2) == roundFloat(cathetsPowLen, 2)
}

func getEdgeLen(x1, y1, x2, y2 float64) float64 {
	var edgeLen float64 = math.Sqrt(math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2))
	return edgeLen
}

func convertPointsToFloat64(x1_, y1_, x2_, y2_, x3_, y3_ int) (float64, float64, float64, float64, float64, float64){
	return float64(x1_), float64(y1_), float64(x2_), float64(y2_), float64(x3_), float64(y3_)
}

func roundFloat(val float64, prec int) float64 {
	var ratio float64 = math.Pow(10, float64(prec))
	return math.Round(val * ratio) / ratio
}