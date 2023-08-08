package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testFormControlName = "test name"
	testIssueRef        = "i.abcd.1234"
	testIssueValue      = "test value"
	testFormControlRef  = "j.1111.2222"
	testFormControlDesc = "test description"
)

func TestNewFormControlRef(t *testing.T) {
	n := NewFormControlRef()
	assert.NotNil(t, n)
	assert.Equal(t, FormControlRefTable, n.Ref[0:1])
}

func TestFormsControlRef_ValidateSuccess(t *testing.T) {
	var controlRef FormControlRef

	err := controlRef.ValidateStringParams(testFormControlName, "testError")
	assert.Nil(t, err)

	err = controlRef.ValidateStringParams(testIssueRef, "testError")
	assert.Nil(t, err)

	err = controlRef.ValidateStringParams(testIssueValue, "testError")
	assert.Nil(t, err)
}

func TestFormsControlRef_ValidateError(t *testing.T) {
	var controlRef FormControlRef

	err := controlRef.ValidateStringParams(testFormControlName, "testError")
	assert.NotNil(t, err)

	err = controlRef.ValidateStringParams(testIssueRef, "testError")
	assert.NotNil(t, err)

	err = controlRef.ValidateStringParams(testIssueValue, "testError")
	assert.NotNil(t, err)
}

func TestFormsControlRef_BuildFormControlRefSuccess(t *testing.T) {
	var controlRef FormControlRef
	err := controlRef.BuildFormControlRefForCreate(getValidFormControl(), testIssueRef, testIssueValue)

	assert.Nil(t, err)
}

func TestFormsControlRef_BuildFormControlRefInvalidForm(t *testing.T) {
	var controlRef FormControlRef
	err := controlRef.BuildFormControlRefForCreate(getInvalidFormControl(), testIssueRef, testIssueValue)

	assert.NotNil(t, err)
}

func TestFormsControlRef_BuildFormControlRefInvalidJSON(t *testing.T) {
	var controlRef FormControlRef
	err := controlRef.BuildFormControlRefForCreate(getInvalidFormControlWithBadJSON(), testIssueRef, testIssueValue)

	assert.NotNil(t, err)
}

func TestFormsControlRef_BuildFormControlRefMissingRef(t *testing.T) {
	var controlRef FormControlRef
	err := controlRef.BuildFormControlRefForCreate(getValidFormControl(), "", testIssueValue)

	assert.NotNil(t, err)
}

func TestFormsControlRef_BuildFormControlRefMissingValue(t *testing.T) {
	var controlRef FormControlRef
	err := controlRef.BuildFormControlRefForCreate(getValidFormControl(), testIssueRef, "")

	assert.NotNil(t, err)
}

func getValidFormControl() FormControl {
	return FormControl{
		Id:          1,
		Name:        testFormsControlName,
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "{\"key\": \"text\",  \"type\": \"text\",  \"templateOptions\": {    \"label\": \"Text\", \"placeholder\": \"Name, email or phone number of Area Manager\", \"required\": true  }}",
	}
}

func getInvalidFormControl() FormControl {
	return FormControl{
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "{\"key\": \"text\",  \"type\": \"text\",  \"templateOptions\": {    \"label\": \"Text\", \"placeholder\": \"Name, email or phone number of Area Manager\", \"required\": true  }}",
	}
}

func getInvalidFormControlWithBadJSON() FormControl {
	return FormControl{
		Id:          1,
		Name:        testFormsControlName,
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "invalid json",
	}
}
