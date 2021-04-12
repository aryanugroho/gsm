/*
 * // Licensed to the Apache Software Foundation (ASF) under one
 * // or more contributor license agreements.  See the NOTICE file
 * // distributed with this work for additional information
 * // regarding copyright ownership.  The ASF licenses this file
 * // to you under the Apache License, Version 2.0 (the
 * // "License"); you may not use this file except in compliance
 * // with the License.  You may obtain a copy of the License at
 * //
 * //   http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing,
 * // software distributed under the License is distributed on an
 * // "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * // KIND, either express or implied.  See the License for the
 * // specific language governing permissions and limitations
 * // under the License.
 *
 *
 *
 *
 * author: Eddy Kioi
 * project: gcp-secret-manager
 * date: 15/06/2020, 14:17
 */

package gsm

import (
	"context"

	opt "github.com/googleapis/gax-go/v2"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Declare Mock funcs
var (
	GetSecretFunc            func(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error)
	AccessSecretVersionFunc  func(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...opt.CallOption) error
	GetSecretVersionFunc     func(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
)

// MockClient is the mock client
type MockClient struct {
	GetSecretFunc            func(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error)
	AccessSecretVersionFunc  func(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...opt.CallOption) error
	GetSecretVersionFunc     func(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error)
}

// GetSecret Mock Get Secret
func (m *MockClient) GetSecret(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error) {
	return GetSecretFunc(ctx, req)
}

// AccessSecretVersion Mock Access SecretVersion
func (m *MockClient) AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return AccessSecretVersionFunc(ctx, req)
}

// DestroySecretVersion Mock Destroy Secret Version
func (m *MockClient) DestroySecretVersion(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return DestroySecretVersionFunc(ctx, req)
}

// CreateSecret Mock Create Secret Version
func (m *MockClient) CreateSecret(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error) {
	return CreateSecretFunc(ctx, req)
}

// AddSecretVersion Mock Add Secret Version
func (m *MockClient) AddSecretVersion(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return AddSecretVersionFunc(ctx, req)
}

// DeleteSecret Mock Delete Secret
func (m *MockClient) DeleteSecret(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...opt.CallOption) error {
	return DeleteSecretFunc(ctx, req)
}

// GetSecretVersion Mock Get Secret Version
func (m *MockClient) GetSecretVersion(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return GetSecretVersionFunc(ctx, req)
}

// DisableSecretVersion Mock Disable Secret Version
func (m *MockClient) DisableSecretVersion(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return DisableSecretVersionFunc(ctx, req)
}

// EnableSecretVersion Mock Enable Secret Version
func (m *MockClient) EnableSecretVersion(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return EnableSecretVersionFunc(ctx, req)
}

// Close Mock Close Client
func (m *MockClient) Close() error {
	return nil
}
