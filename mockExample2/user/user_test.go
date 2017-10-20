package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/Duxxie/mockExample/mock_doer"
	"fmt"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mock_doer.NewMockDoer(mockCtrl)
	testUser := &User{
		Doer: mockDoer,
	}

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(fmt.Errorf("xyz")).Times(1)

	err := testUser.Use()
	if err.Error() != "xyz" {
		t.Fatalf("expected: %s, got: %s", "xyz", err.Error())
	}
}
