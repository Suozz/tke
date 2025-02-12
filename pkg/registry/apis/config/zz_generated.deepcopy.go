// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemStorage) DeepCopyInto(out *FileSystemStorage) {
	*out = *in
	if in.MaxThreads != nil {
		in, out := &in.MaxThreads, &out.MaxThreads
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemStorage.
func (in *FileSystemStorage) DeepCopy() *FileSystemStorage {
	if in == nil {
		return nil
	}
	out := new(FileSystemStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryStorage) DeepCopyInto(out *InMemoryStorage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryStorage.
func (in *InMemoryStorage) DeepCopy() *InMemoryStorage {
	if in == nil {
		return nil
	}
	out := new(InMemoryStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfiguration) DeepCopyInto(out *RegistryConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Storage.DeepCopyInto(&out.Storage)
	in.Security.DeepCopyInto(&out.Security)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfiguration.
func (in *RegistryConfiguration) DeepCopy() *RegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Storage) DeepCopyInto(out *S3Storage) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(string)
		**out = **in
	}
	if in.RegionEndpoint != nil {
		in, out := &in.RegionEndpoint, &out.RegionEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.Secure != nil {
		in, out := &in.Secure, &out.Secure
		*out = new(bool)
		**out = **in
	}
	if in.SkipVerify != nil {
		in, out := &in.SkipVerify, &out.SkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.V4Auth != nil {
		in, out := &in.V4Auth, &out.V4Auth
		*out = new(bool)
		**out = **in
	}
	if in.ChunkSize != nil {
		in, out := &in.ChunkSize, &out.ChunkSize
		*out = new(int64)
		**out = **in
	}
	if in.MultipartCopyChunkSize != nil {
		in, out := &in.MultipartCopyChunkSize, &out.MultipartCopyChunkSize
		*out = new(int64)
		**out = **in
	}
	if in.MultipartCopyMaxConcurrency != nil {
		in, out := &in.MultipartCopyMaxConcurrency, &out.MultipartCopyMaxConcurrency
		*out = new(int64)
		**out = **in
	}
	if in.MultipartCopyThresholdSize != nil {
		in, out := &in.MultipartCopyThresholdSize, &out.MultipartCopyThresholdSize
		*out = new(int64)
		**out = **in
	}
	if in.RootDirectory != nil {
		in, out := &in.RootDirectory, &out.RootDirectory
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(S3StorageClass)
		**out = **in
	}
	if in.UserAgent != nil {
		in, out := &in.UserAgent, &out.UserAgent
		*out = new(string)
		**out = **in
	}
	if in.ObjectACL != nil {
		in, out := &in.ObjectACL, &out.ObjectACL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Storage.
func (in *S3Storage) DeepCopy() *S3Storage {
	if in == nil {
		return nil
	}
	out := new(S3Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Security) DeepCopyInto(out *Security) {
	*out = *in
	if in.TokenExpiredHours != nil {
		in, out := &in.TokenExpiredHours, &out.TokenExpiredHours
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Security.
func (in *Security) DeepCopy() *Security {
	if in == nil {
		return nil
	}
	out := new(Security)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.FileSystem != nil {
		in, out := &in.FileSystem, &out.FileSystem
		*out = new(FileSystemStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.InMemory != nil {
		in, out := &in.InMemory, &out.InMemory
		*out = new(InMemoryStorage)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Storage)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}
