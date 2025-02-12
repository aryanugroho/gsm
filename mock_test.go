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
	"errors"
	"testing"

	opt "github.com/googleapis/gax-go/v2"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func TestMockClient_Close(t *testing.T) {
	AddSecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	GetSecretFunc = func(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	GetSecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	CreateSecretFunc = func(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...opt.CallOption) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	DestroySecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	DisableSecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	EnableSecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	AccessSecretVersionFunc = func(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...opt.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
		return &secretmanagerpb.AccessSecretVersionResponse{
			Name:    "projects/myProjects/secrets/mySecrets/versions/latest",
			Payload: &secretmanagerpb.SecretPayload{Data: []byte("mySecret")},
		}, nil
	}
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				GetSecretFunc:            nil,
				AccessSecretVersionFunc:  nil,
				DestroySecretVersionFunc: nil,
				CreateSecretFunc:         nil,
				AddSecretVersionFunc:     nil,
				DeleteSecretFunc:         nil,
				GetSecretVersionFunc:     nil,
				DisableSecretVersionFunc: nil,
				EnableSecretVersionFunc:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockClient{
				GetSecretFunc:            tt.fields.GetSecretFunc,
				AccessSecretVersionFunc:  tt.fields.AccessSecretVersionFunc,
				DestroySecretVersionFunc: tt.fields.DestroySecretVersionFunc,
				CreateSecretFunc:         tt.fields.CreateSecretFunc,
				AddSecretVersionFunc:     tt.fields.AddSecretVersionFunc,
				DeleteSecretFunc:         tt.fields.DeleteSecretFunc,
				GetSecretVersionFunc:     tt.fields.GetSecretVersionFunc,
				DisableSecretVersionFunc: tt.fields.DisableSecretVersionFunc,
				EnableSecretVersionFunc:  tt.fields.EnableSecretVersionFunc,
			}
			if err := m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
