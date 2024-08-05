package vm

import (
	"fmt"
	"monkey/object"
)

func (vm *VM) buildArray(startIndex, endIndex int) object.Object {
	elements := make([]object.Object, endIndex-startIndex)
	for i := startIndex; i < endIndex; i++ {
		elements[i-startIndex] = vm.stack[i]
	}
	return &object.Array{Elements: elements}
}

func (vm *VM) buildHash(startIndex, endIndex int) (object.Object, error) {
	hashedPairs := make(map[object.HashKey]object.HashPair)
	for i := startIndex; i < endIndex; i += 2 {
		key := vm.stack[i]
		value := vm.stack[i+1]
		pair := object.HashPair{Key: key, Value: value}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", key.Type())
		}
		hashedPairs[hashKey.HashKey()] = pair
	}
	return &object.Hash{Pairs: hashedPairs}, nil
}
