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

package errs

var (
	ErrArgs             = NewCodeError(ArgsError, "ArgsError")
	ErrNoPermission     = NewCodeError(NoPermissionError, "NoPermissionError")
	ErrInternalServer   = NewCodeError(ServerInternalError, "ServerInternalError")
	ErrRecordNotFound   = NewCodeError(RecordNotFoundError, "RecordNotFoundError")
	ErrDuplicateKey     = NewCodeError(DuplicateKeyError, "DuplicateKeyError")
	ErrTokenExpired     = NewCodeError(TokenExpiredError, "TokenExpiredError")
	ErrTokenInvalid     = NewCodeError(TokenInvalidError, "TokenInvalidError")
	ErrTokenMalformed   = NewCodeError(TokenMalformedError, "TokenMalformedError")
	ErrTokenNotValidYet = NewCodeError(TokenNotValidYetError, "TokenNotValidYetError")
	ErrTokenUnknown     = NewCodeError(TokenUnknownError, "TokenUnknownError")
	ErrTokenKicked      = NewCodeError(TokenKickedError, "TokenKickedError")
	ErrTokenNotExist    = NewCodeError(TokenNotExistError, "TokenNotExistError")
	ErrVerionError      = NewCodeError(VersionError, "VersionError")

	ErrPassword                 = NewCodeError(20001, "PasswordError")
	ErrAccountNotFound          = NewCodeError(20002, "AccountNotFound")
	ErrPhoneAlreadyRegister     = NewCodeError(20003, "PhoneAlreadyRegister")
	ErrAccountAlreadyRegister   = NewCodeError(20004, "AccountAlreadyRegister")
	ErrVerifyCodeSendFrequently = NewCodeError(20005, "VerifyCodeSendFrequently")
	ErrVerifyCodeNotMatch       = NewCodeError(20006, "VerifyCodeNotMatch")
	ErrVerifyCodeExpired        = NewCodeError(20007, "VerifyCodeExpired")
	ErrVerifyCodeMaxCount       = NewCodeError(20008, "VerifyCodeMaxCount")
	ErrVerifyCodeUsed           = NewCodeError(20009, "VerifyCodeUsed")
	ErrInvitationCodeUsed       = NewCodeError(20010, "InvitationCodeUsed")
	ErrInvitationNotFound       = NewCodeError(20011, "InvitationNotFound")
	ErrForbidden                = NewCodeError(20012, "Forbidden")
	ErrRefuseFriend             = NewCodeError(20013, "RefuseFriend")
	ErrEmailAlreadyRegister     = NewCodeError(20014, "EmailAlreadyRegister")

	ErrChatTokenNotExist = NewCodeError(20101, "ChatTokenNotExist")
)
