package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testFormsControlRef    = "testRef"
	testFormsControlName   = "Test Name"
	testFormsControlString = "Test Control"
)

func TestFormsControl_ValdiateSuccess(t *testing.T) {
	control := FormControl{Ref: testFormsControlRef, Name: testFormsControlName, Control: testFormsControlString}

	err := control.Validate()

	assert.Nil(t, err)
}

func TestFormsControl_ValdiateError(t *testing.T) {
	controlMissingRef := FormControl{Name: testFormsControlName, Control: testFormsControlString}
	controlMissingName := FormControl{Ref: testFormsControlRef, Control: testFormsControlString}
	controlMissingControl := FormControl{Ref: testFormsControlRef, Name: testFormsControlName}

	err := controlMissingRef.Validate()
	assert.NotNil(t, err)

	err = controlMissingName.Validate()
	assert.NotNil(t, err)

	err = controlMissingControl.Validate()
	assert.NotNil(t, err)
}