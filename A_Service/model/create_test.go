package model

import (
	"errors"
	"my_task/A_Service/views"
	"testing"
)

type MockDAOInt struct {
	FakeConnect func()
	FakeClose func()
	FakeCreate func(string, string) error
	FakeDelete func(string) error
	FakePing func() bool
}

func (m *MockDAOInt) Connect() {
	panic("implement me")
}

func (m *MockDAOInt) Close() {
	panic("implement me")
}

func (m *MockDAOInt) Read(title string, c *Cache) ([]views.Movie, error) {
	panic("implement me")
}

func (m *MockDAOInt) Delete(title string) error {
	if m.FakeDelete != nil {
		return m.FakeDelete(title)
	}
	return nil
}

func (m *MockDAOInt) Ping() bool {
	if m.FakePing != nil {
		return m.FakePing()
	}
	return true
}

func (m *MockDAOInt) Create(title, year string) error {
	if m.FakeCreate != nil {
		return m.FakeCreate(title, year)
	}
	return nil
}

func TestCreate(t *testing.T){
	t.Run("Create return nil", func(t *testing.T) {
		DB = &MockDAOInt{
			FakeCreate: func(string, string) error {
				return nil
			},
		}
		if DB.Create("", "") != nil {
			t.Errorf("Test did not return expected results")
		}
	})
	t.Run("Create return error", func(t *testing.T) {
		DB = &MockDAOInt{
			FakeCreate: func(string, string) error {
				return errors.New("expected error")
			},
		}
		if DB.Create("", "") == nil {
			t.Errorf("Test did not return expected results")
		}
	})
}



