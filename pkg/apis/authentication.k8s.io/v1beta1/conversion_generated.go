// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	api "k8s.io/kubernetes/pkg/api"
	authentication_k8s_io "k8s.io/kubernetes/pkg/apis/authentication.k8s.io"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_TokenReview_To_authenticationk8sio_TokenReview,
		Convert_authenticationk8sio_TokenReview_To_v1beta1_TokenReview,
		Convert_v1beta1_TokenReviewSpec_To_authenticationk8sio_TokenReviewSpec,
		Convert_authenticationk8sio_TokenReviewSpec_To_v1beta1_TokenReviewSpec,
		Convert_v1beta1_TokenReviewStatus_To_authenticationk8sio_TokenReviewStatus,
		Convert_authenticationk8sio_TokenReviewStatus_To_v1beta1_TokenReviewStatus,
		Convert_v1beta1_UserInfo_To_authenticationk8sio_UserInfo,
		Convert_authenticationk8sio_UserInfo_To_v1beta1_UserInfo,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1beta1_TokenReview_To_authenticationk8sio_TokenReview(in *TokenReview, out *authentication_k8s_io.TokenReview, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewSpec_To_authenticationk8sio_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewStatus_To_authenticationk8sio_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_TokenReview_To_authenticationk8sio_TokenReview(in *TokenReview, out *authentication_k8s_io.TokenReview, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReview_To_authenticationk8sio_TokenReview(in, out, s)
}

func autoConvert_authenticationk8sio_TokenReview_To_v1beta1_TokenReview(in *authentication_k8s_io.TokenReview, out *TokenReview, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_authenticationk8sio_TokenReviewSpec_To_v1beta1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authenticationk8sio_TokenReviewStatus_To_v1beta1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authenticationk8sio_TokenReview_To_v1beta1_TokenReview(in *authentication_k8s_io.TokenReview, out *TokenReview, s conversion.Scope) error {
	return autoConvert_authenticationk8sio_TokenReview_To_v1beta1_TokenReview(in, out, s)
}

func autoConvert_v1beta1_TokenReviewSpec_To_authenticationk8sio_TokenReviewSpec(in *TokenReviewSpec, out *authentication_k8s_io.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_v1beta1_TokenReviewSpec_To_authenticationk8sio_TokenReviewSpec(in *TokenReviewSpec, out *authentication_k8s_io.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewSpec_To_authenticationk8sio_TokenReviewSpec(in, out, s)
}

func autoConvert_authenticationk8sio_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication_k8s_io.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_authenticationk8sio_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication_k8s_io.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_authenticationk8sio_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in, out, s)
}

func autoConvert_v1beta1_TokenReviewStatus_To_authenticationk8sio_TokenReviewStatus(in *TokenReviewStatus, out *authentication_k8s_io.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_v1beta1_UserInfo_To_authenticationk8sio_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_TokenReviewStatus_To_authenticationk8sio_TokenReviewStatus(in *TokenReviewStatus, out *authentication_k8s_io.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewStatus_To_authenticationk8sio_TokenReviewStatus(in, out, s)
}

func autoConvert_authenticationk8sio_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication_k8s_io.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_authenticationk8sio_UserInfo_To_v1beta1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func Convert_authenticationk8sio_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication_k8s_io.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_authenticationk8sio_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in, out, s)
}

func autoConvert_v1beta1_UserInfo_To_authenticationk8sio_UserInfo(in *UserInfo, out *authentication_k8s_io.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = in.Groups
	out.Extra = in.Extra
	return nil
}

func Convert_v1beta1_UserInfo_To_authenticationk8sio_UserInfo(in *UserInfo, out *authentication_k8s_io.UserInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_UserInfo_To_authenticationk8sio_UserInfo(in, out, s)
}

func autoConvert_authenticationk8sio_UserInfo_To_v1beta1_UserInfo(in *authentication_k8s_io.UserInfo, out *UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = in.Groups
	out.Extra = in.Extra
	return nil
}

func Convert_authenticationk8sio_UserInfo_To_v1beta1_UserInfo(in *authentication_k8s_io.UserInfo, out *UserInfo, s conversion.Scope) error {
	return autoConvert_authenticationk8sio_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
