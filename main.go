package main

import (
	"fmt"
	"time"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func main() {
	fmt.Println("Starting program!!")

	polygon := orb.MultiPolygon{NewPolygon()}

	featureCollection := geojson.FeatureCollection{
		Features: []*geojson.Feature{},
	}

	points := []orb.Point{
		{105.5, 13.7},
		{100.51444172859192, 13.752068124818038},
		{100.50998926162718, 13.751818014017923},
		{100.5129289627075, 13.750963466768232},
		{100.51170587539673, 13.749942174500532},
	}

	isInPolygon := isPointInsidePolygon(&featureCollection, points[0])
	fmt.Println("isInPolygon: ", isInPolygon)

	start := time.Now()
	for _, point := range points {
		inPolygon := isPointInsidePolygon2(polygon, point)
		fmt.Printf("%f, %f is in polygon: %v\n", point[1], point[0], inPolygon)
	}
	fmt.Println("Use ", time.Now().Sub(start))
}

// isPointInsidePolygon runs through the MultiPolygon and Polygons within a
// feature collection and checks if a point (long/lat) lies within it.
func isPointInsidePolygon(fc *geojson.FeatureCollection, point orb.Point) bool {
	for _, feature := range fc.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			if planar.MultiPolygonContains(multiPoly, point) {
				return true
			}
		} else {
			// Fallback to Polygon
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if planar.PolygonContains(polygon, point) {
					return true
				}
			}
		}
	}
	return false
}

func isPointInsidePolygon2(polygon orb.MultiPolygon, point orb.Point) bool {
	return planar.MultiPolygonContains(polygon, point)
}

// NewPolygon from http://geojson.io/#data=data:application/json,%7B%22type%22%3A%22Polygon%22%2C%22coordinates%22%3A%5B%5B%5B100.50924897193909%2C13.754621323948488%5D%2C%5B100.50923824310303%2C13.753547941584374%5D%2C%5B100.50910949707031%2C13.752985195465651%5D%2C%5B100.50885200500488%2C13.752537081848029%5D%2C%5B100.50863742828368%2C13.751922226883776%5D%2C%5B100.51023602485657%2C13.751390740782975%5D%2C%5B100.51098704338074%2C13.751119786620173%5D%2C%5B100.51671624183653%2C13.749077198993833%5D%2C%5B100.51697373390198%2C13.751422004704647%5D%2C%5B100.51707029342651%2C13.753047722878971%5D%2C%5B100.50924897193909%2C13.754621323948488%5D%5D%5D%7D
func NewPolygon() orb.Polygon {
	geometry := [][]float64{{
		100.50924897193909,
		13.754621323948488,
	},
		{
			100.50923824310303,
			13.753547941584374,
		},
		{
			100.50910949707031,
			13.752985195465651,
		},
		{
			100.50885200500488,
			13.752537081848029,
		},
		{
			100.50863742828368,
			13.751922226883776,
		},
		{
			100.51023602485657,
			13.751390740782975,
		},
		{
			100.51098704338074,
			13.751119786620173,
		},
		{
			100.51671624183653,
			13.749077198993833,
		},
		{
			100.51697373390198,
			13.751422004704647,
		},
		{
			100.51707029342651,
			13.753047722878971,
		},
		{
			100.50924897193909,
			13.754621323948488,
		}}

	points := []orb.Point{}
	for _, geo := range geometry {
		points = append(points, orb.Point{geo[0], geo[1]})
	}

	polygon := orb.Polygon{points}

	return polygon
}
