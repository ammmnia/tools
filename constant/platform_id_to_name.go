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

// fixme 1<--->IOS 2<--->Android  3<--->Windows
// fixme  4<--->OSX  5<--->Web  6<--->MiniWeb 7<--->Linux.

type PlatformID int
type TerminalType string

const (
	// Platform ID.
	IOSPlatformID        PlatformID = 1
	AndroidPlatformID    PlatformID = 2
	WindowsPlatformID    PlatformID = 3
	OSXPlatformID        PlatformID = 4
	WebPlatformID        PlatformID = 5
	MiniWebPlatformID    PlatformID = 6
	LinuxPlatformID      PlatformID = 7
	AndroidPadPlatformID PlatformID = 8
	IPadPlatformID       PlatformID = 9
	AdminPlatformID      PlatformID = 10

	// Platform string match to Platform ID.
	IOSPlatformStr        = "IOS"
	AndroidPlatformStr    = "Android"
	WindowsPlatformStr    = "Windows"
	OSXPlatformStr        = "OSX"
	WebPlatformStr        = "Web"
	MiniWebPlatformStr    = "MiniWeb"
	LinuxPlatformStr      = "Linux"
	AndroidPadPlatformStr = "APad"
	IPadPlatformStr       = "IPad"
	AdminPlatformStr      = "Admin"

	// terminal types.
	TerminalPC     TerminalType = "PC"
	TerminalMobile TerminalType = "Mobile"
	TerminalWeb    TerminalType = "Web"
	TerminalAdmin  TerminalType = "Admin"
)

var PlatformID2Name = map[PlatformID]string{
	IOSPlatformID:        IOSPlatformStr,
	AndroidPlatformID:    AndroidPlatformStr,
	WindowsPlatformID:    WindowsPlatformStr,
	OSXPlatformID:        OSXPlatformStr,
	WebPlatformID:        WebPlatformStr,
	MiniWebPlatformID:    MiniWebPlatformStr,
	LinuxPlatformID:      LinuxPlatformStr,
	AndroidPadPlatformID: AndroidPadPlatformStr,
	IPadPlatformID:       IPadPlatformStr,
	AdminPlatformID:      AdminPlatformStr,
}

var PlatformName2ID = map[string]PlatformID{
	IOSPlatformStr:        IOSPlatformID,
	AndroidPlatformStr:    AndroidPlatformID,
	WindowsPlatformStr:    WindowsPlatformID,
	OSXPlatformStr:        OSXPlatformID,
	WebPlatformStr:        WebPlatformID,
	MiniWebPlatformStr:    MiniWebPlatformID,
	LinuxPlatformStr:      LinuxPlatformID,
	AndroidPadPlatformStr: AndroidPadPlatformID,
	IPadPlatformStr:       IPadPlatformID,
	AdminPlatformStr:      AdminPlatformID,
}

var PlatformName2class = map[string]TerminalType{
	IOSPlatformStr:        TerminalMobile,
	AndroidPlatformStr:    TerminalMobile,
	MiniWebPlatformStr:    TerminalWeb,
	WebPlatformStr:        TerminalWeb,
	WindowsPlatformStr:    TerminalPC,
	OSXPlatformStr:        TerminalPC,
	LinuxPlatformStr:      TerminalPC,
	AndroidPadPlatformStr: TerminalMobile,
	IPadPlatformStr:       TerminalMobile,
	AdminPlatformStr:      TerminalAdmin,
}

var PlatformID2class = map[PlatformID]TerminalType{
	IOSPlatformID:        TerminalMobile,
	AndroidPlatformID:    TerminalMobile,
	MiniWebPlatformID:    TerminalWeb,
	WebPlatformID:        TerminalWeb,
	WindowsPlatformID:    TerminalPC,
	OSXPlatformID:        TerminalPC,
	LinuxPlatformID:      TerminalPC,
	AndroidPadPlatformID: TerminalMobile,
	IPadPlatformID:       TerminalMobile,
	AdminPlatformID:      TerminalAdmin,
}

func PlatformIDToName(num PlatformID) string {
	return PlatformID2Name[num]
}

func PlatformNameToID(name string) PlatformID {
	return PlatformName2ID[name]
}

func PlatformNameToClass(name string) TerminalType {
	return PlatformName2class[name]
}

func PlatformIDToClass(num PlatformID) TerminalType {
	return PlatformID2class[num]
}
