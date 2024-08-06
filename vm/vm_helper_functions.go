package vm

import (
	"fmt"
	"monkey/object"
)

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.Null:
		return false
	default:
		return true
	}
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) callFunction(numArgs int) error {
	fn, ok := vm.stack[vm.sp-1-numArgs].(*object.CompiledFunction)
	if !ok {
		return fmt.Errorf("calling non-function")
	}
	
	if numArgs != fn.NumParameters {
		return fmt.Errorf("wrong number of arguments: want=%d, got=%d", fn.NumParameters, numArgs)
	}

	frame := NewFrame(fn, vm.sp-numArgs)
	vm.pushFrame(frame)

	vm.sp = frame.basePointer + fn.NumLocals

	return nil
}
