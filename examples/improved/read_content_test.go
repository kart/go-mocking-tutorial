package improved

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kart/go-mocking-tutorial/base"
	"github.com/kart/go-mocking-tutorial/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_empty_name(t *testing.T) {
	b, err := ReadContent("")

	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "no name", err.Error())
}

func Test_Stat_returns_error(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(nil, errors.New("stat error"))

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "stat error", err.Error())
}

func Test_Stat_returns_size_less_than_10(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	mockFileInfo := mocks.NewMockFileInfo(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockFileInfo.
		EXPECT().
		Size().
		Return(int64(8))

	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(mockFileInfo, nil)

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "too small", err.Error())
}

func Test_Open_fails(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	mockFileInfo := mocks.NewMockFileInfo(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockFileInfo.
		EXPECT().
		Size().
		Return(int64(11))

	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(mockFileInfo, nil)

	mockOs.
		EXPECT().
		Open("test.txt").
		Return(nil, errors.New("open failed"))

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "open failed", err.Error())
}

func Test_Read_fails(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	mockFileInfo := mocks.NewMockFileInfo(ctrl)
	mockFile := mocks.NewMockFile(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockFileInfo.
		EXPECT().
		Size().
		Return(int64(11)).
		Times(2)

	mockFile.
		EXPECT().
		Read(gomock.AssignableToTypeOf([]byte{})).
		Return(0, errors.New("read failed"))

	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(mockFileInfo, nil)

	mockOs.
		EXPECT().
		Open("test.txt").
		Return(mockFile, nil)

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())
}

func Test_partial_Read(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	mockFileInfo := mocks.NewMockFileInfo(ctrl)
	mockFile := mocks.NewMockFile(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockFileInfo.
		EXPECT().
		Size().
		Return(int64(20)).
		Times(3)

	mockFile.
		EXPECT().
		Read(gomock.AssignableToTypeOf([]byte{})).
		Do(func(b []byte) {
			assert.NotNil(t, b)
			assert.Equal(t, 20, len(b))
			// Copy a 5 byte array to trigger partial read.
			copy(b, "hello")
		}).
		Return(5, nil)

	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(mockFileInfo, nil)

	mockOs.
		EXPECT().
		Open("test.txt").
		Return(mockFile, nil)

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Equal(t, "partial read", err.Error())
}

func Test_Read_completely_successful(t *testing.T) {
	// Global setup
	ctrl := gomock.NewController(t)
	defer func() {
		base.AppOs = &base.DefaultOs{}
		ctrl.Finish()
	}()
	mockOs := mocks.NewMockOs(ctrl)
	mockFileInfo := mocks.NewMockFileInfo(ctrl)
	mockFile := mocks.NewMockFile(ctrl)
	base.AppOs = mockOs

	// Expectations
	mockFileInfo.
		EXPECT().
		Size().
		Return(int64(20)).
		Times(3)

	mockFile.
		EXPECT().
		Read(gomock.AssignableToTypeOf([]byte{})).
		Do(func(b []byte) {
			assert.NotNil(t, b)
			assert.Equal(t, 20, len(b))
			// Copy a 20 byte array
			copy(b, "helloworldworldhello")
		}).
		Return(20, nil)

	mockOs.
		EXPECT().
		Stat("test.txt").
		Return(mockFileInfo, nil)

	mockOs.
		EXPECT().
		Open("test.txt").
		Return(mockFile, nil)

	// SUT
	b, err := ReadContent("test.txt")

	// Assertions
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.Equal(t, []byte("helloworldworldhello"), b)
}
