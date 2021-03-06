/*

SPDX-Copyright: Copyright (c) Brad Rydzewski, project contributors, Capital One Services, LLC
SPDX-License-Identifier: Apache-2.0
Copyright 2017 Brad Rydzewski, project contributors, Capital One Services, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

*/
package model

type User struct {
	ID     int64  `json:"id"      meddler:"user_id,pk"`
	Login  string `json:"login"   meddler:"user_login"`
	Token  string `json:"-"       meddler:"user_token"`
	Avatar string `json:"avatar"  meddler:"user_avatar"`
	Secret string `json:"-"       meddler:"user_secret"`
	Scopes string `json:"-"       meddler:"user_scopes"`
}
