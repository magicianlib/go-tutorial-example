package main

import (
	"fmt"
	"math"
)

//
// Radian: https://en.wikipedia.org/wiki/Radian
//

// ToRadians 角度转弧度
//
//	angle in radians = angle in degrees * π / 180°
func ToRadians(degrees float64) (radians float64) {
	radians = degrees * math.Pi / 180.0
	return
}

// ToDegrees 弧度转角度
//
//	angle in degrees = angle in radians * 180° / π
func ToDegrees(radians float64) (degrees float64) {
	degrees = radians * 180.0 / math.Pi
	return
}

// EarthRadius 地球平均半径(单位: 米)
const EarthRadius float64 = 6378137.0

// CoordsDistance 计算两个经纬度之间距离(单位: 米)
func CoordsDistance(lng1, lat1, lng2, lat2 float64) (dist float64) {

	lng1 = ToRadians(lng1)
	lat1 = ToRadians(lat1)
	lng2 = ToRadians(lng2)
	lat2 = ToRadians(lat2)

	theta := lng2 - lng1

	dist = math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	dist = dist * EarthRadius

	return
}

func main() {
	dist := CoordsDistance(121.455244, 31.234076, 121.488301, 31.237534)
	fmt.Println(dist, "米")
}
