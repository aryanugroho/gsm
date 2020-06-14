package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// MockClient is the mock client
type MockClient struct {
	GetSecretFunc            func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
	ListSecretVersionsFunc   func(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator
	AccessSecretVersionFunc  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(req *secretmanagerpb.DeleteSecretRequest) error
	ListSecretsFunc          func(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator
	GetSecretVersionFunc     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
}

func (m *MockClient) ListSecretVersions(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator {
	return ListSecretVersionsFunc(req)
}

func (m *MockClient) AccessSecretVersion(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return AccessSecretVersionFunc(req)
}

func (m *MockClient) DestroySecretVersion(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DestroySecretVersionFunc(req)
}

func (m *MockClient) CreateSecret(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return CreateSecretFunc(req)
}

func (m *MockClient) AddSecretVersion(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return AddSecretVersionFunc(req)
}

func (m *MockClient) DeleteSecret(req *secretmanagerpb.DeleteSecretRequest) error {
	return DeleteSecretFunc(req)
}

func (m *MockClient) ListSecrets(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator {
	return m.ListSecretsFunc(req)
}

func (m *MockClient) GetSecretVersion(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return GetSecretVersionFunc(req)
}

func (m *MockClient) DisableSecretVersion(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DisableSecretVersionFunc(req)
}

func (m *MockClient) EnableSecretVersion(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return EnableSecretVersionFunc(req)
}

func (m *MockClient) Close() error {
	return nil
}

var (
	// GetSecretFunc fetches the mock client's `GetSecret` func
	GetSecretFunc            func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
	ListSecretVersionsFunc   func(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator
	AccessSecretVersionFunc  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(req *secretmanagerpb.DeleteSecretRequest) error
	//	ListSecrets          func(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator
	GetSecretVersionFunc     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
)

func (m *MockClient) GetSecret(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return GetSecretFunc(req)
}
