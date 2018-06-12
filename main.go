package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/gonum/stat"
	convtree "github.com/visheratin/conv-tree"
)

func test(filePath string) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var points []convtree.Point
	err = json.Unmarshal(rawData, &points)
	if err != nil {
		fmt.Println(err)
		return
	}
	minX := math.MaxFloat64
	maxX := -math.MaxFloat64
	minY := math.MaxFloat64
	maxY := -math.MaxFloat64
	for _, point := range points {
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}
	topLeft := convtree.Point{
		X: minX,
		Y: minY,
	}
	bottomRight := convtree.Point{
		X: maxX,
		Y: maxY,
	}
	fmt.Println("Total number of points -", len(points))
	minXLen, minYLen := 0.001, 0.001
	maxPoints, maxDepth := 20, 50
	convNum, gridSize := 4, 25
	tree, err := convtree.NewConvTree(topLeft, bottomRight, minXLen, minYLen, maxPoints, maxDepth, convNum, gridSize, nil, points)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ConvTree stats")
	fmt.Println("Tree depth:")
	depths, lengths := checkConvTree(tree)
	min := math.MaxFloat64
	max := -math.MaxFloat64
	for _, e := range depths {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	fmt.Println("Min, max -", min, max)
	fmt.Println("Mean -", stat.Mean(depths, nil))
	fmt.Println("Standard deviaion -", stat.StdDev(depths, nil))
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, e := range lengths {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	fmt.Println("Number of points in leaf:")
	fmt.Println("Min, max -", min, max)
	fmt.Println("Mean -", stat.Mean(lengths, nil))
	fmt.Println("Standard deviaion -", stat.StdDev(lengths, nil))
	distances, leafsNum := analyzeConvDepth(tree, tree, map[string]bool{})
	fmt.Println("Number of leafs -", leafsNum)
	min = math.MaxFloat64
	max = -math.MaxFloat64
	total := 0.0
	for _, e := range distances {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
		total += e
	}
	fmt.Println("Number of transitions to neighbour leafs:")
	fmt.Println("Min, max -", min, max)
	fmt.Println("Total -", total)
	fmt.Println("Mean -", stat.Mean(distances, nil))
	fmt.Println("Standard deviation -", stat.StdDev(distances, nil))

	fmt.Println("-----")
	quad, err := convtree.NewQuadTree(topLeft, bottomRight, minXLen, minYLen, maxPoints, maxDepth, points)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("QuadTree stats")
	fmt.Println("Tree depth:")
	depths, lengths = checkQuadTree(quad)
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, e := range depths {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	fmt.Println("Min, max -", min, max)
	fmt.Println("Mean -", stat.Mean(depths, nil))
	fmt.Println("Standard deviation -", stat.StdDev(depths, nil))
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, e := range lengths {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	fmt.Println("Number of points in leaf:")
	fmt.Println("Min, max -", min, max)
	fmt.Println("Mean -", stat.Mean(lengths, nil))
	fmt.Println("Standard deviation -", stat.StdDev(lengths, nil))
	distances, leafsNum = analyzeQuadDepth(quad, quad, map[string]bool{})
	fmt.Println("Number of leafs -", leafsNum)
	min = math.MaxFloat64
	max = -math.MaxFloat64
	total = 0.0
	for _, e := range distances {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
		total += e
	}
	fmt.Println("Number of transitions to neighbour leafs:")
	fmt.Println("Min, max -", min, max)
	fmt.Println("Total -", total)
	fmt.Println("Mean -", stat.Mean(distances, nil))
	fmt.Println("Standard deviation -", stat.StdDev(distances, nil))
}

func main() {
	smallDataPath := "./data/small.json"
	fmt.Println("Checking small file.")
	test(smallDataPath)
	fmt.Println()
	fmt.Println("=====================")
	fmt.Println()
	mediumDataPath := "./data/medium.json"
	fmt.Println("Checking medium file.")
	test(mediumDataPath)
	fmt.Println()
	fmt.Println("=====================")
	fmt.Println()
	largeDataPath := "./data/large.json"
	fmt.Println("Checking large file.")
	test(largeDataPath)
}

func checkConvTree(tree convtree.ConvTree) ([]float64, []float64) {
	if tree.IsLeaf {
		total := 0
		for _, point := range tree.Points {
			total += point.Weight
		}
		return []float64{float64(tree.Depth)}, []float64{float64(total)}
	} else {
		depth, length := []float64{}, []float64{}
		d, l := checkConvTree(*tree.ChildBottomLeft)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkConvTree(*tree.ChildBottomRight)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkConvTree(*tree.ChildTopLeft)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkConvTree(*tree.ChildTopRight)
		depth = append(depth, d...)
		length = append(length, l...)
		return depth, length
	}
}

func checkQuadTree(tree convtree.QuadTree) ([]float64, []float64) {
	if tree.IsLeaf {
		total := 0
		for _, point := range tree.Points {
			total += point.Weight
		}
		return []float64{float64(tree.Depth)}, []float64{float64(total)}
	} else {
		depth, length := []float64{}, []float64{}
		d, l := checkQuadTree(*tree.ChildBottomLeft)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkQuadTree(*tree.ChildBottomRight)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkQuadTree(*tree.ChildTopLeft)
		depth = append(depth, d...)
		length = append(length, l...)
		d, l = checkQuadTree(*tree.ChildTopRight)
		depth = append(depth, d...)
		length = append(length, l...)
		return depth, length
	}
}

func analyzeConvDepth(inTree convtree.ConvTree, outTree convtree.ConvTree, distance map[string]bool) ([]float64, int) {
	if inTree.IsLeaf {
		if outTree.IsLeaf {
			if inTree.ID != outTree.ID && (inTree.TopLeft.X == outTree.TopLeft.X || inTree.TopLeft.X == outTree.BottomRight.X ||
				inTree.BottomRight.X == outTree.TopLeft.X || inTree.BottomRight.X == outTree.BottomRight.X) {
				if inTree.TopLeft.Y <= outTree.BottomRight.Y && inTree.BottomRight.Y >= outTree.TopLeft.Y {
					return []float64{float64(len(distance) + 2)}, 0
				}
			}
			if inTree.ID != outTree.ID && (inTree.TopLeft.Y == outTree.TopLeft.Y || inTree.TopLeft.Y == outTree.BottomRight.Y ||
				inTree.BottomRight.Y == outTree.TopLeft.Y || inTree.BottomRight.Y == outTree.BottomRight.Y) {
				if inTree.TopLeft.X <= outTree.BottomRight.X && inTree.BottomRight.X >= outTree.TopLeft.X {
					return []float64{float64(len(distance) + 2)}, 0
				}
			}
		} else {
			d := copyMap(distance)
			transitions, _ := analyzeConvDepth(outTree, inTree, d)
			return transitions, 1
		}

	} else {
		result := []float64{}
		leafs := 0
		if _, ok := distance[inTree.ID]; ok {
			delete(distance, inTree.ID)
		} else {
			distance[inTree.ID] = true
		}
		d := copyMap(distance)
		childTransitions, leafsNum := analyzeConvDepth(*inTree.ChildBottomLeft, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeConvDepth(*inTree.ChildBottomRight, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeConvDepth(*inTree.ChildTopLeft, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeConvDepth(*inTree.ChildTopRight, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		return result, leafs
	}
	return nil, 0
}

func copyMap(input map[string]bool) map[string]bool {
	result := map[string]bool{}
	for k, v := range input {
		result[k] = v
	}
	return result
}

func analyzeQuadDepth(inTree convtree.QuadTree, outTree convtree.QuadTree, distance map[string]bool) ([]float64, int) {
	if inTree.IsLeaf {
		if outTree.IsLeaf {
			if inTree.ID != outTree.ID && (inTree.TopLeft.X == outTree.TopLeft.X || inTree.TopLeft.X == outTree.BottomRight.X ||
				inTree.BottomRight.X == outTree.TopLeft.X || inTree.BottomRight.X == outTree.BottomRight.X) {
				if inTree.TopLeft.Y <= outTree.BottomRight.Y && inTree.BottomRight.Y >= outTree.TopLeft.Y {
					return []float64{float64(len(distance) + 2)}, 0
				}
			}
			if inTree.ID != outTree.ID && (inTree.TopLeft.Y == outTree.TopLeft.Y || inTree.TopLeft.Y == outTree.BottomRight.Y ||
				inTree.BottomRight.Y == outTree.TopLeft.Y || inTree.BottomRight.Y == outTree.BottomRight.Y) {
				if inTree.TopLeft.X <= outTree.BottomRight.X && inTree.BottomRight.X >= outTree.TopLeft.X {
					return []float64{float64(len(distance) + 2)}, 0
				}
			}
		} else {
			d := copyMap(distance)
			transitions, _ := analyzeQuadDepth(outTree, inTree, d)
			return transitions, 1
		}

	} else {
		result := []float64{}
		leafs := 0
		if _, ok := distance[inTree.ID]; ok {
			delete(distance, inTree.ID)
		} else {
			distance[inTree.ID] = true
		}
		d := copyMap(distance)
		childTransitions, leafsNum := analyzeQuadDepth(*inTree.ChildBottomLeft, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeQuadDepth(*inTree.ChildBottomRight, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeQuadDepth(*inTree.ChildTopLeft, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		d = copyMap(distance)
		childTransitions, leafsNum = analyzeQuadDepth(*inTree.ChildTopRight, outTree, d)
		result = append(result, childTransitions...)
		leafs += leafsNum
		return result, leafs
	}
	return nil, 0
}
