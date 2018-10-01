package sublist

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(arr1, arr2 []int) Relation {
	if isEqual(arr1, arr2) {
		return RelationEqual
	} else if contains(arr1, arr2) {
		return RelationSuperlist
	} else if contains(arr2, arr1) {
		return RelationSublist
	}
	return RelationUnequal
}

func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2){
		return false
	}
	for ind, i := range arr1 {
		if i != arr2[ind] {
			return false
		}
	}
	return true
}

func contains(arr1, arr2 []int) bool {
	if len(arr2) == 0 && len(arr1) >= 0 {
		return true
	}
	for ind1, val1 := range arr1 {
		//the first element of array 2 should be present in array one to make it a sublist
		if arr2[0] == val1 {
			//Variable to hold the number of common elements
			a := 1
			//check if there is enough elements left in array 1 starting from the found element to hold all elements of array 2
			for ind2, _ := range arr2 {
				if len(arr1) < len(arr2) + ind1{
					return false
				}
				//loop through array 2 to check if all its elements are present in array 1 starting from the found element
				if len(arr1) > ind1+ind2 {
					if arr2[ind2] == arr1[ind1+ind2]{
						//increment number of elements found
						a++
					}
				}
			}
			//check if number of found elements corresponds to the length of array 2, if it's not the case continue looking for the next match of arr2[0] in arr1 (line 46)
			if a > len(arr2){
				return true
			}
		}
	}

	return false
}