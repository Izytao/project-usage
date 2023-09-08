package tools

import "sort"

// SortMap key首字母排序，结果输出
func SortMap(m map[string]interface{}) []interface{} {
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	result := make([]interface{}, 0)
	sort.Strings(keys)
	for _, k := range keys {
		result = append(result, m[k])
	}
	return result
}

// 快排分区操作
// 接受一个指向整数切片的指针 arr，以及两个整数 left 和 right，它们表示要对切片进行分区的范围。
func partition(arr *[]int, left int, right int) int {
	//基准元素
	privot := (*arr)[right]
	//初始时，没有元素小于基准，所以 i 设为 left - 1。
	i := left - 1
	//从 left 到 right-1。对于每个元素，
	//如果它小于基准元素 privot，则将 i 的值增加 1，然后交换 arr[i] 和 arr[j]，
	//这样将小于基准的元素移到了 i 的左侧。
	for j := left; j < right; j++ {
		if (*arr)[j] < privot {
			i++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = temp
		}
	}
	//最后，将基准元素 privot 与 arr[i+1] 交换，将基准元素放在了正确的位置。
	//此时，基准元素左边的元素都小于它，右边的元素都大于它。
	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[right]
	(*arr)[right] = temp
	//返回 i+1，这是基准元素的最终位置，它将在后续的快速排序递归中被使用，将数组分成两部分并进行排序。
	return i + 1
}

// QuickSort 快速排序算法
func QuickSort(arr *[]int, left int, right int) {
	//表示数组范围内只有一个元素或没有元素，无需进行排序，直接返回。
	if left >= right {
		return
	}
	//将数组分成两个子数组，并返回基准元素（pivot）的索引。
	privot := partition(arr, left, right)
	//递归
	QuickSort(arr, left, privot-1)
	QuickSort(arr, privot+1, right)
}

// QuickSort2 快排2
//找到一个基准，左边是所有比它小的，右边是比它大的，分别递归左右
func QuickSort2(arr *[]int, left int, right int) {
	if left >= right {
		return
	}
	privot := (*arr)[left]
	i := left
	j := right
	for i < j {
		for i < j && (*arr)[j] > privot {
			j--
		}
		for i < j && (*arr)[i] <= privot {
			i++
		}
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = temp
	}
	(*arr)[left] = (*arr)[i]
	(*arr)[i] = privot

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

// BubbleSort 冒泡排序
//比较相邻元素，较大的往右移
func BubbleSort(arr *[]int) {
	flag := true
	lastSwapIndex := 0
	for i := 0; i < len(*arr)-1; i++ {
		sortBorder := len(*arr) - 1 - i
		for j := 0; j < sortBorder; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				flag = false
				lastSwapIndex = j
			}
		}
		sortBorder = lastSwapIndex
		if flag {
			break
		}
	}
}

// InsertionSort 插入排序
//将未排序部分插入到已排序部分的适当位置
func InsertionSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		curKey := (*arr)[i]
		j := i - 1
		for curKey < (*arr)[j] {
			(*arr)[j+1] = (*arr)[j]
			j--
			if j < 0 {
				break
			}
		}
		(*arr)[j+1] = curKey
	}
}

// SelectionSort 选择排序
//选择一个最小值，再寻找比它还小的进行交换
func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minIndex := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = temp
	}
}

// MergeSort 归并排序
//利用临时数组合并两个有序数组
func MergeSort(arr *[]int, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	i := left
	j := mid + 1
	p := 0
	temp := make([]int, right-left+1)
	for i <= mid && j <= right {
		if (*arr)[i] <= (*arr)[j] {
			temp[p] = (*arr)[i]
			i++
		} else {
			temp[p] = (*arr)[j]
			j++
		}
		p++
	}

	for i <= mid {
		temp[p] = (*arr)[i]
		i++
		p++
	}
	for j <= right {
		temp[p] = (*arr)[j]
		j++
		p++
	}
	for i = 0; i < len(temp); i++ {
		(*arr)[left+i] = temp[i]
	}
}
