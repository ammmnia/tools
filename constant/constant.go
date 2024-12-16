// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constant

const (

	// OnlineStatus  = "online"
	// OfflineStatus = "offline"
	// Registered    = "registered"
	// UnRegistered  = "unregistered"

	Online  = 1
	Offline = 0

	Registered   = 1
	UnRegistered = 0
)

const (
	WriteDiffusion = 0
	ReadDiffusion  = 1
)

const (
	UnreliableNotification    = 1
	ReliableNotificationNoMsg = 2
	ReliableNotificationMsg   = 3
)

const (
	FieldRecvMsgOpt    = 1
	FieldIsPinned      = 2
	FieldAttachedInfo  = 3
	FieldIsPrivateChat = 4
	FieldGroupAtType   = 5
	FieldEx            = 7
	FieldUnread        = 8
	FieldBurnDuration  = 9
	FieldHasReadSeq    = 10
)

const (
	IMOrdinaryUser       = 0
	AppOrdinaryUsers     = 1
	AppAdmin             = 2
	AppNotificationAdmin = 3

	GroupOwner         = 100
	GroupAdmin         = 60
	GroupOrdinaryUsers = 20

	Male   = 1
	Female = 2
)

const (
	GroupRPCRecvSize = 30
	GroupRPCSendSize = 30
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"

const LogFileName = "OpenIM.log"

const LocalHost = "0.0.0.0"

const OpenIMCommonConfigKey = "OpenIMServerConfig"

const BatchNum = 100

// Subscribe to user constants
const (
	SubscriberUser = 1
	Unsubscribe    = 2
)

const (
	GroupSearchPositionHead = 1
	GroupSearchPositionAny  = 2
)

const (
	FirstPageNumber   = 1
	MaxSyncPullNumber = 500
)
