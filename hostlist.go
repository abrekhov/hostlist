package hostlist

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// ExpandNodeList function expands SLURM's and other schedulers
// lists of nodes to array of nodes
// e.g. n02p[001-003] -> ["n02p001", "n02p002", "n02p003"]
func ExpandNodeList(nodeListString string) []string {
	//https://github.com/LLNL/py-hostlist/blob/master/hostlist/hostlist.py
	var resultHostlist []string
	nodeSlice := strings.Split(nodeListString, ", ")
	re := regexp.MustCompile(`(\w+-?)\[((,?[0-9]+-?,?-?){0,})\](.*)?`)
	for _, node := range nodeSlice {
		match := re.FindStringSubmatch(node)
		if len(match) != 0 {
			oldstr := match[2]
			leftBr := strings.Replace(oldstr, "[", "", -1)
			rightBr := strings.Replace(leftBr, "]", "", -1)
			numList := strings.Split(rightBr, ",")
			var finalList []int
			leadZeros := 0
			leadZerosStr := ""
			for _, elem := range numList {
				if strings.Count(elem, "-") > 0 {
					tmpReplaced := strings.Replace(elem, "-", ",", -1)
					tmpList := strings.Split(tmpReplaced, ",")
					for _, digit := range tmpList[0] {
						if string(digit) == "0" {
							leadZeros++
							leadZerosStr = leadZerosStr + "0"
						}
					}
					tmpListFirst, err := strconv.ParseInt(tmpList[0], 10, 0)
					if err != nil {
						return []string{}
					}
					tmpListLast, err := strconv.ParseInt(tmpList[1], 10, 0)
					if err != nil {
						return []string{}
					}
					rngList := makeRange(int(tmpListFirst), int(tmpListLast))
					finalList = append(finalList, rngList...)
				} else {
					integ, err := strconv.ParseInt(elem, 10, 0)
					if err != nil {
						return []string{}
					}
					finalList = append(finalList, int(integ))
				}

			}
			sort.Ints(finalList)
			var hostlistTmp []string
			for _, elem := range finalList {
				if (leadZeros > 0) && (len(fmt.Sprint(elem)) <= len(leadZerosStr)) {
					hostlistTmp = append(hostlistTmp, fmt.Sprintf("%0*d", leadZeros+1, int(elem)))
				} else {
					hostlistTmp = append(hostlistTmp, strconv.Itoa(elem))
				}
			}
			var hostlistnoSuffix []string
			for _, elem := range hostlistTmp {
				hostlistnoSuffix = append(hostlistnoSuffix, match[1]+elem)
			}
			var finalHostlist []string
			for _, elem := range hostlistnoSuffix {
				finalHostlist = append(finalHostlist, elem+match[4])
			}
			resultHostlist = append(resultHostlist, finalHostlist...)
		} else {
			resultHostlist = append(resultHostlist, node)
		}
	}
	return resultHostlist
}

// ExpandCPUList function expands SLURM's and other schedulers
// lists of used logical CPUs to array of CPU
// e.g. 36-38,41 -> ["36", "37", "38", "41"]
func ExpandCPUList(cpuListString string) []string {
	var resultHostlist []string
	cpuSlice := strings.Split(cpuListString, ", ")
	re := regexp.MustCompile(`((,?[0-9]+-?,?-?){0,})`)
	for _, node := range cpuSlice {
		match := re.FindStringSubmatch(node)
		if len(match) != 0 {
			oldstr := match[1]
			leftBr := strings.Replace(oldstr, "[", "", -1)
			rightBr := strings.Replace(leftBr, "]", "", -1)
			numList := strings.Split(rightBr, ",")
			var finalList []int
			leadZeros := 0
			leadZerosStr := ""
			for _, elem := range numList {
				if strings.Count(elem, "-") > 0 {
					tmpReplaced := strings.Replace(elem, "-", ",", -1)
					tmpList := strings.Split(tmpReplaced, ",")
					for _, digit := range tmpList[0] {
						if string(digit) == "0" {
							leadZeros++
							leadZerosStr = leadZerosStr + "0"
						}
					}
					tmpListFirst, err := strconv.ParseInt(tmpList[0], 10, 0)
					if err != nil {
						return []string{}
					}
					tmpListLast, err := strconv.ParseInt(tmpList[1], 10, 0)
					if err != nil {
						return []string{}
					}
					rngList := makeRange(int(tmpListFirst), int(tmpListLast))
					finalList = append(finalList, rngList...)
				} else {
					integ, err := strconv.ParseInt(elem, 10, 0)
					if err != nil {
						return []string{}
					}
					finalList = append(finalList, int(integ))
				}

			}
			sort.Ints(finalList)
			var hostlistTmp []string
			for _, elem := range finalList {
				if (leadZeros > 0) && (len(fmt.Sprint(elem)) <= len(leadZerosStr)) {
					hostlistTmp = append(hostlistTmp, fmt.Sprintf("%d", int(elem)))
				} else {
					hostlistTmp = append(hostlistTmp, strconv.Itoa(elem))
				}
			}
			resultHostlist = append(resultHostlist, hostlistTmp...)
		} else {
			resultHostlist = append(resultHostlist, node)
		}
	}
	return resultHostlist
}
