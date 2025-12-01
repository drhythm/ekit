package list

import (
	"reflect"
	"testing"
)

// TestArrayList_Get verifies the logic of the Get method.
func TestArrayList_Get(t *testing.T) {
	// Define the test table structure.
	tests := []struct {
		name	string
		list	*ArrayList[int]
		index	int
		wantVal	int
		wantErr	error
	}{
		// Case 1: Normal retrieval
		{
			name:	"get valid index",
			list:	NewArrayListOf([]int{10, 20, 30}),
			index:	1,
			wantVal: 20,
			wantErr: nil,
		},
		// Case 2: Index too large
		{
			name:	"index out of bounds(too large)",
			list:	NewArrayListOf([]int{10, 20, 30}),
			index: 	5,
			wantVal: 0,
			wantErr: ErrIndexOutOfRange,
		},
		// Case 3: Index negative
		{
			name:	"index out of bounds(negative)",
			list:	NewArrayListOf([]int{10, 20, 30}),
			index:	-1,
			wantVal: 0,
			wantErr: ErrIndexOutOfRange,
		},
	}
	// Iterate over the table.
	for _, tt := range tests {
		// t.Run creates a sub-test for each case.
		t.Run(tt.name, func(t *testing.T) {
			// Act: Call the method being tested.
			gotVal, err := tt.list.Get(tt.index)
			// Assert Error: Check if the error matches expectation.
			if err != tt.wantErr {
				// Log the error mismatch.
				t.Errorf("Get() err = %v, wantErr = %v\n", err, tt.wantErr)
				return
			}
			// Assert Value: Check if the returned value matches expectation.
			if gotVal != tt.wantVal {
				// Log the value mismatch.
				t.Errorf("Get() gotVal = %v, wantVal = %v\n", gotVal, tt.wantVal)
			}
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	makeList := func(length, capacity int) *ArrayList[int] {
		vals := make([]int, length, capacity)
		for i := 0; i < length; i++ {
			vals[i] = i + 1
		}
		return NewArrayListOf(vals)
	}
	tests := []struct{
		name	string
		list	*ArrayList[int]
		index	int
		wantVal int
		wantErr error
		wantSlice []int
		wantCap	int
	}{
		// Case 1: Normal deletion
		{
			name:	"delete valid index",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	2,
			wantVal: 3,
			wantErr: nil,
			wantSlice: []int{1, 2, 4, 5},
			wantCap: 5,
		},
		// Case 2: Index too large
		{
			name:	"index out of bounds(too large)",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	5,
			wantVal: 0,
			wantErr: ErrIndexOutOfRange,
			wantSlice: []int{1, 2, 3, 4, 5},
			wantCap: 5,
		},
		// Case 3: Index negative
		{
			name:	"index out of bounds(negative)",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	-1,
			wantVal: 0,
			wantErr: ErrIndexOutOfRange,
			wantSlice: []int{1, 2, 3, 4, 5},
			wantCap: 5,
		},
		// Case 4: Shrink Capacity
		{
			name:	"trigger shrink",
			list:	makeList(25, 100),
			index:	24,
			wantVal: 25,
			wantErr: nil,
			wantSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			wantCap: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotVal, err := tt.list.Delete(tt.index)
			if err != tt.wantErr {
				t.Errorf("Delete() err %v, wantErr %v\n", err, tt.wantErr)
				return
			}
			if gotVal != tt.wantVal {
				t.Errorf("Delete() gotArr %v, wantArr %v\n", gotVal, tt.wantVal)
				return
			}
			if !reflect.DeepEqual(tt.list.vals, tt.wantSlice) {
				t.Errorf("Delete() vals %v, wantSlice %v\n", tt.list.vals, tt.wantSlice)
				return
			}
			if cap(tt.list.vals) != tt.wantCap {
				t.Errorf("Delete() cap %v, wantCap %v\n", cap(tt.list.vals), tt.wantCap)
			}
		})
	}
}


func TestArray_Insert(t *testing.T){
	tests := []struct {
		name	string
		list	*ArrayList[int]
		index	int
		value	int
		wantSlice []int
		wantErr	error
	}{
		// Normal insert
		{
			name:	"insert valid index",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	2,
			value:	5,
			wantSlice: []int{1, 2, 5, 3, 4, 5},
			wantErr: nil,
		},
		// Index equals length
		{
			name:	"append element",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	5,
			value:	6,
			wantSlice: []int{1, 2, 3, 4, 5, 6},
			wantErr: nil,
		},
		// Index too large
		{
			name:	"index out of bounds(too large)",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	6,
			value:	0,
			wantSlice: []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
		// Index negative
		{
			name:	"index out of bounds(negative)",
			list:	NewArrayListOf([]int{1, 2, 3, 4, 5}),
			index:	-1,
			value:	0,
			wantSlice: []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if err := tt.list.Insert(tt.index, tt.value); err != tt.wantErr {
				t.Errorf("Insert() err %v, wantErr %v\n", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.list.vals, tt.wantSlice) {
				t.Errorf("Insert() vals %v, wantSlice %v\n", tt.list.vals, tt.wantSlice)
			}
		})
	}
}