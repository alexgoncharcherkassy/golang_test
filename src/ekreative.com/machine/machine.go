package machine

import (
	"awesomeProject/src/ekreative.com/product"
	"errors"
)

type Machine struct {
	buckets         map[int]Bucket
	numberOfBuckets int
	bucketSize      int
}

func (m Machine) Buckets() map[int]Bucket {
	return m.buckets
}

func MakeMachine(numberOfBuckets int, numberOfProductsInBucket int) Machine {
	buckets := make(map[int]Bucket)

	for i := 1; i <= numberOfBuckets; i++ {
		buckets[len(buckets)+1] = MakeBucket(numberOfProductsInBucket)
	}

	return Machine{buckets, numberOfBuckets, numberOfProductsInBucket}
}

func InArray(needle int, source []int) bool {
	isPresent := false

	for _, item := range source {
		if item == needle {
			isPresent = true
			break
		}
	}

	return isPresent
}

func (m *Machine) GetProducts(products []int) ([]int, error) {
	currentPositions := map[int]int{}
	result := map[int][]int{}

	for key, bucket := range m.buckets {
		currentPositions[key] = 0

		var productTmp []int
		var productPosition []int
		for _, elem := range bucket.elements {
			numberFoundProductsBefore := len(productTmp)

			for pos, productName := range products {
				if elem.Name() == productName && InArray(pos, productPosition) == false {
					productTmp = append(productTmp, productName)
					productPosition = append(productPosition, pos)
					break
				}
			}

			if len(productTmp) == numberFoundProductsBefore {
				break
			}
		}

		if len(productTmp) > 0 {
			result[key] = productTmp
		}
	}

	var foundProducts []int

	if len(result) > 0 {
		processedBuckets := map[int]int{}

		result = CompareItems(result, products, processedBuckets, -1, -1)

		for bucketNumber, item := range result {
			for _, val := range item {
				pr, err := m.GetProductFromBucket(val, bucketNumber)

				if err != nil {
					return []int{}, err
				}

				foundProducts = append(foundProducts, pr.Name())
			}
		}
	}

	return foundProducts, nil
}

func DiffSlice(slice1, slice2 []int) []int {
	var diff []int
	tmp1 := MakeCopy(slice1)
	tmp2 := MakeCopy(slice2)

	for _, sourceItem := range tmp1 {
		isPresent := false

		for i := 0; i < len(tmp2); i++ {
			if tmp2[i] == sourceItem {
				isPresent = true
				tmp2 = append(tmp2[:i], tmp2[i+1:]...)
				break
			}
		}

		if isPresent == false {
			diff = append(diff, sourceItem)
		}
	}

	return diff
}

func GetMaxLenBucket(source map[int][]int, processedBuckets map[int]int) int {
	var maxLen int
	var bucketNumber int

	for key, val := range source {
		_, ok := processedBuckets[key]

		if len(val) > maxLen && ok == false {
			maxLen = len(val)
			bucketNumber = key
		}
	}

	return bucketNumber
}

func MakeCopy(source []int) []int {
	copySlice := make([]int, len(source))
	copy(copySlice, source)

	return copySlice
}

func CompareItems(result map[int][]int, products []int, processedBuckets map[int]int, baseBucket int, baseBucketPosition int) map[int][]int {
	var baseBucketItem []int
	var baseBucketNumber int
	tmpResult := map[int][]int{}

	if baseBucket >= 0 {
		baseBucketItem = MakeCopy(result[baseBucket])
		baseBucketNumber = baseBucket
	} else {
		baseBucketNumber = GetMaxLenBucket(result, processedBuckets)
		baseBucketItem = MakeCopy(result[baseBucketNumber])
	}

	tmpResult[baseBucketNumber] = baseBucketItem

	if baseBucketPosition >= 0 {
		baseBucketItem = append(baseBucketItem[:baseBucketPosition], baseBucketItem[baseBucketPosition+(len(baseBucketItem)-baseBucketPosition):]...)
		tmpResult[baseBucketNumber] = baseBucketItem
	}

	notFoundInBucket := DiffSlice(products, baseBucketItem)

	for bucketNumber, item := range result {
		if bucketNumber != baseBucketNumber {
			if len(notFoundInBucket) == 0 {
				break
			}

			var numberProcessedElem int

			beforeLength := len(notFoundInBucket)
			bItem := MakeCopy(item)
			for _, bucketItem := range bItem {
				for _, val := range notFoundInBucket {
					if bucketItem == val {
						notFoundInBucket = notFoundInBucket[1:]
						numberProcessedElem++
						break
					}
				}

				if beforeLength == len(notFoundInBucket) || len(notFoundInBucket) == 0 {
					break
				}
			}

			for i := numberProcessedElem; i < len(bItem); i++ {
				bItem = append(bItem[:i], bItem[i+1:]...)
				i--
			}

			if len(bItem) > 0 {
				tmpResult[bucketNumber] = bItem
			}
		}
	}

	if len(notFoundInBucket) > 0 {
		baseBucketItem = tmpResult[baseBucketNumber]
		var position int
		if len(baseBucketItem) > 0 {
			baseBucketItem = baseBucketItem[:len(baseBucketItem)-1]
			position = len(baseBucketItem)
		} else {
			position = -1
		}

		if position < 0 {

			if len(processedBuckets) == len(result) {
				return map[int][]int{}
			}

			processedBuckets[baseBucketNumber] = baseBucketNumber
			baseBucketNumber = GetMaxLenBucket(result, processedBuckets)
		}

		tmpResult = CompareItems(result, products, processedBuckets, baseBucketNumber, position)
	}

	return tmpResult
}

func (m *Machine) AddProduct(productName int, numberOfBucket int) error {
	if numberOfBucket > len(m.buckets) || numberOfBucket < 1 {
		return errors.New("incorrect bucket")
	}

	bucket := m.buckets[numberOfBucket]
	err := bucket.Push(product.MakeProduct(productName))

	if err != nil {
		return err
	}

	m.buckets[numberOfBucket] = bucket

	return nil
}

func (m *Machine) GetProductFromBucket(productName int, numberOfBucket int) (*product.Product, error) {
	if numberOfBucket > len(m.buckets) || numberOfBucket < 1 {
		return nil, errors.New("incorrect bucket")
	}

	bucket := m.buckets[numberOfBucket]
	element, err := bucket.Pop()

	if err != nil {
		return nil, err
	}

	m.buckets[numberOfBucket] = bucket

	return element, nil
}
