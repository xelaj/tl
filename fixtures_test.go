// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl_test

// это набор разнообразных фикстур, которые хранятся исключительно для тестов
// фикстуры являются полноценно сгенерированными объектами из пакета telegram

import (
	. "github.com/xelaj/tl"
)

type MultipleChats struct {
	Chats []any
}

func (*MultipleChats) CRC() uint32 {
	return uint32(0xff1144cc)
}

type Chat struct {
	_                 struct{} `tl:"flag,bitflag"`
	Creator           bool     `tl:",omitempty:flag:0,implicit"`
	Kicked            bool     `tl:",omitempty:flag:1,implicit"`
	Left              bool     `tl:",omitempty:flag:2,implicit"`
	Deactivated       bool     `tl:",omitempty:flag:5,implicit"`
	ID                int32
	Title             string
	Photo             string
	ParticipantsCount int32
	Date              int32
	Version           int32
	AdminRights       *Rights `tl:",omitempty:flag:14"`
	BannedRights      *Rights `tl:",omitempty:flag:18"`
}

func (*Chat) CRC() uint32 {
	return uint32(0x3bda1bde)
}

func (*Chat) FlagIndex() int {
	return 0
}

type AuthSentCode struct {
	_             struct{} `tl:"flag,bitflag"`
	Type          AuthSentCodeType
	PhoneCodeHash string
	NextType      AuthCodeType `tl:",omitempty:flag:1"`
	Timeout       *int32       `tl:",omitempty:flag:2"`
}

func (*AuthSentCode) CRC() uint32 {
	return uint32(0x5e002502)
}

func (*AuthSentCode) FlagIndex() int {
	return 0
}

type SomeNullStruct struct{}

func (*SomeNullStruct) CRC() uint32 {
	return uint32(0xc4f9186b)
}

type AuthSentCodeType interface {
	Object
	ImplementsAuthSentCodeType()
}

type AuthSentCodeTypeApp struct {
	Length int32
}

func (*AuthSentCodeTypeApp) CRC() uint32 {
	return uint32(0x3dbb5986)
}

func (*AuthSentCodeTypeApp) ImplementsAuthSentCodeType() {}

type Rights struct {
	_              struct{} `tl:"flag,bitflag"`
	DeleteMessages bool     `tl:",omitempty:flag:3,implicit"`
	BanUsers       bool     `tl:",omitempty:flag:4,implicit"`
}

func (*Rights) CRC() uint32 {
	return uint32(0x5fb224d5)
}

func (*Rights) FlagIndex() int {
	return 0
}

type AuthCodeType uint32

const (
	AuthCodeTypeSms       AuthCodeType = 1923290508
	AuthCodeTypeCall      AuthCodeType = 1948046307
	AuthCodeTypeFlashCall AuthCodeType = 577556219
)

func (e AuthCodeType) String() string {
	switch e {
	case AuthCodeTypeSms:
		return "auth.codeTypeSms"
	case AuthCodeTypeCall:
		return "auth.codeTypeCall"
	case AuthCodeTypeFlashCall:
		return "auth.codeTypeFlashCall"
	default:
		return "<UNKNOWN auth.CodeType>"
	}
}
func (e AuthCodeType) CRC() uint32 {
	return uint32(e)
}

type PollResults struct { //nolint:maligned required ordering
	_                struct{}            `tl:"flag,bitflag"`
	Min              bool                `tl:",omitempty:flag:0,implicit"`
	Results          []*PollAnswerVoters `tl:",omitempty:flag:1"`
	TotalVoters      *int32              `tl:",omitempty:flag:2"`
	RecentVoters     []int32             `tl:",omitempty:flag:3"`
	Solution         *string             `tl:",omitempty:flag:4"`
	SolutionEntities []MessageEntity     `tl:",omitempty:flag:4"`
}

func (*PollResults) CRC() uint32 {
	return uint32(0xbadcc1a3)
}
func (*PollResults) FlagIndex() int {
	return 0
}

type PollAnswerVoters struct { //nolint:maligned required ordering
	_       struct{} `tl:"flag,bitflag"`
	Chosen  bool     `tl:",omitempty:flag:0,implicit"`
	Correct bool     `tl:",omitempty:flag:1,implicit"`
	Option  []byte
	Voters  int32
}

func (*PollAnswerVoters) CRC() uint32 {
	return uint32(0x3b6ddad2)
}
func (*PollAnswerVoters) FlagIndex() int {
	return 0
}

type MessageEntity interface {
	Object
	ImplementsMessageEntity()
}

// type MessageEntityUnknown struct {
// 	Offset int32
// 	Length int32
// }
//
// func (*MessageEntityUnknown) CRC() uint32 {
// 	return uint32(0xbb92ba95)
// }
//
// func (*MessageEntityUnknown) ImplementsMessageEntity() {}

type AccountInstallThemeParams struct {
	_      struct{}   `tl:"flag,bitflag"`
	Dark   bool       `tl:",omitempty:flag:0,implicit"`
	Format *string    `tl:",omitempty:flag:1"`
	Theme  InputTheme `tl:",omitempty:flag:1"`
}

func (e *AccountInstallThemeParams) CRC() uint32 {
	return uint32(0x7ae43737)
}

func (*AccountInstallThemeParams) FlagIndex() int {
	return 0
}

type InputTheme interface {
	Object
	ImplementsInputTheme()
}

type InputThemeObj struct {
	ID         int64
	AccessHash int64
}

func (*InputThemeObj) CRC() uint32 {
	return uint32(0x3c5693e9)
}

func (*InputThemeObj) ImplementsInputTheme() {}

type AccountUnregisterDeviceParams struct {
	TokenType int32
	Token     string
	OtherUids []int32
}

func (e *AccountUnregisterDeviceParams) CRC() uint32 {
	return uint32(0x3076c4bf)
}

type AnyStructWithAnyType struct {
	SomeInt int32
	Data    any
}

func (*AnyStructWithAnyType) CRC() uint32 {
	return uint32(0xfdfd4646)
}

type AnyStructWithAnyObject struct {
	SomeInt int32
	Data    Object
}

func (*AnyStructWithAnyObject) CRC() uint32 {
	return uint32(0xfd46fd46)
}

type InvokeWithLayerParams struct {
	Layer int32
	Query any
}

func (*InvokeWithLayerParams) CRC() uint32 {
	return 0xda9b0d0d
}

type InitConnectionParams struct {
	_              struct{} `tl:"flag,bitflag"`
	APIID          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          any `tl:",omitempty:flag:0"`
	Params         any `tl:",omitempty:flag:1"`

	Query any
}

func (*InitConnectionParams) CRC() uint32 {
	return 0xc1cd5ea9
}

func (*InitConnectionParams) FlagIndex() int {
	return 0
}

type ResPQ struct {
	Nonce        *Int128
	ServerNonce  *Int128
	Pq           []byte
	Fingerprints []int64
}

func (*ResPQ) CRC() uint32 {
	return 0x05162463
}

// issue #59 fixtures

type Poll struct {
	ID             int64
	_              struct{} `tl:"flag,bitflag"`
	Closed         bool     `tl:",omitempty:flag:0,implicit"`
	PublicVoters   bool     `tl:",omitempty:flag:1,implicit"`
	MultipleChoice bool     `tl:",omitempty:flag:2,implicit"`
	Quiz           bool     `tl:",omitempty:flag:3,implicit"`
	Question       string
	Answers        []*PollAnswer
	ClosePeriod    *int32 `tl:",omitempty:flag:4"`
	CloseDate      *int32 `tl:",omitempty:flag:5"`
}

func (*Poll) CRC() uint32 {
	return 0x86e18161
}

func (*Poll) FlagIndex() int {
	return 1
}

type PollAnswer struct {
	Text   string
	Option []byte
}

func (*PollAnswer) CRC() uint32 {
	return 0x6ca9c2e9
}

func ptr[T any](value T) *T { return &value }
