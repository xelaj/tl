package main

import (
	"context"
	tl "github.com/xelaj/tl"
)

type IAccountDaysTTL interface {
	tl.Object
	_IAccountDaysTTL()
}

var (
	_ IAccountDaysTTL = (*AccountDaysTTL)(nil)
)

type AccountDaysTTL struct {
	Days int32
}

func (*AccountDaysTTL) CRC() uint32 {
	return 0xb8d0afdf
}
func (*AccountDaysTTL) _IAccountDaysTTL() {}

type IAppWebViewResult interface {
	tl.Object
	_IAppWebViewResult()
}

var (
	_ IAppWebViewResult = (*AppWebViewResultURL)(nil)
)

type AppWebViewResultURL struct {
	URL string
}

func (*AppWebViewResultURL) CRC() uint32 {
	return 0x3c1b4f0d
}
func (*AppWebViewResultURL) _IAppWebViewResult() {}

type IAttachMenuBot interface {
	tl.Object
	_IAttachMenuBot()
}

var (
	_ IAttachMenuBot = (*AttachMenuBot)(nil)
)

type AttachMenuBot struct {
	_                        struct{} `tl:"flags,bitflag"`
	Inactive                 bool     `tl:",omitempty:flags:0,implicit"`
	HasSettings              bool     `tl:",omitempty:flags:1,implicit"`
	RequestWriteAccess       bool     `tl:",omitempty:flags:2,implicit"`
	ShowInAttachMenu         bool     `tl:",omitempty:flags:3,implicit"`
	ShowInSideMenu           bool     `tl:",omitempty:flags:4,implicit"`
	SideMenuDisclaimerNeeded bool     `tl:",omitempty:flags:5,implicit"`
	BotID                    int64
	ShortName                string
	PeerTypes                []EnumAttachMenuPeerType `tl:",omitempty:flags:3"`
	Icons                    []IAttachMenuBotIcon
}

func (*AttachMenuBot) CRC() uint32 {
	return 0xd90d8dfe
}
func (*AttachMenuBot) _IAttachMenuBot() {}

type IAttachMenuBotIcon interface {
	tl.Object
	_IAttachMenuBotIcon()
}

var (
	_ IAttachMenuBotIcon = (*AttachMenuBotIcon)(nil)
)

type AttachMenuBotIcon struct {
	_      struct{} `tl:"flags,bitflag"`
	Name   string
	Icon   IDocument
	Colors []IAttachMenuBotIconColor `tl:",omitempty:flags:0"`
}

func (*AttachMenuBotIcon) CRC() uint32 {
	return 0xb2a7386b
}
func (*AttachMenuBotIcon) _IAttachMenuBotIcon() {}

type IAttachMenuBotIconColor interface {
	tl.Object
	_IAttachMenuBotIconColor()
}

var (
	_ IAttachMenuBotIconColor = (*AttachMenuBotIconColor)(nil)
)

type AttachMenuBotIconColor struct {
	Name  string
	Color int32
}

func (*AttachMenuBotIconColor) CRC() uint32 {
	return 0x4576f3f0
}
func (*AttachMenuBotIconColor) _IAttachMenuBotIconColor() {}

type IAttachMenuBots interface {
	tl.Object
	_IAttachMenuBots()
}

var (
	_ IAttachMenuBots = (*AttachMenuBotsNotModified)(nil)
	_ IAttachMenuBots = (*AttachMenuBots)(nil)
)

type AttachMenuBotsNotModified struct{}

func (*AttachMenuBotsNotModified) CRC() uint32 {
	return 0xf1d88a5c
}
func (*AttachMenuBotsNotModified) _IAttachMenuBots() {}

type AttachMenuBots struct {
	Hash  int64
	Bots  []IAttachMenuBot
	Users []IUser
}

func (*AttachMenuBots) CRC() uint32 {
	return 0x3c4301c0
}
func (*AttachMenuBots) _IAttachMenuBots() {}

type IAttachMenuBotsBot interface {
	tl.Object
	_IAttachMenuBotsBot()
}

var (
	_ IAttachMenuBotsBot = (*AttachMenuBotsBot)(nil)
)

type AttachMenuBotsBot struct {
	Bot   IAttachMenuBot
	Users []IUser
}

func (*AttachMenuBotsBot) CRC() uint32 {
	return 0x93bf667f
}
func (*AttachMenuBotsBot) _IAttachMenuBotsBot() {}

type IAuthorization interface {
	tl.Object
	_IAuthorization()
}

var (
	_ IAuthorization = (*Authorization)(nil)
)

type Authorization struct {
	_                         struct{} `tl:"flags,bitflag"`
	Current                   bool     `tl:",omitempty:flags:0,implicit"`
	OfficialApp               bool     `tl:",omitempty:flags:1,implicit"`
	PasswordPending           bool     `tl:",omitempty:flags:2,implicit"`
	EncryptedRequestsDisabled bool     `tl:",omitempty:flags:3,implicit"`
	CallRequestsDisabled      bool     `tl:",omitempty:flags:4,implicit"`
	Unconfirmed               bool     `tl:",omitempty:flags:5,implicit"`
	Hash                      int64
	DeviceModel               string
	Platform                  string
	SystemVersion             string
	APIID                     int32
	AppName                   string
	AppVersion                string
	DateCreated               int32
	DateActive                int32
	Ip                        string
	Country                   string
	Region                    string
}

func (*Authorization) CRC() uint32 {
	return 0xad01d61d
}
func (*Authorization) _IAuthorization() {}

type IAutoDownloadSettings interface {
	tl.Object
	_IAutoDownloadSettings()
}

var (
	_ IAutoDownloadSettings = (*AutoDownloadSettings)(nil)
)

type AutoDownloadSettings struct {
	_                             struct{} `tl:"flags,bitflag"`
	Disabled                      bool     `tl:",omitempty:flags:0,implicit"`
	VideoPreloadLarge             bool     `tl:",omitempty:flags:1,implicit"`
	AudioPreloadNext              bool     `tl:",omitempty:flags:2,implicit"`
	PhonecallsLessData            bool     `tl:",omitempty:flags:3,implicit"`
	StoriesPreload                bool     `tl:",omitempty:flags:4,implicit"`
	PhotoSizeMax                  int32
	VideoSizeMax                  int64
	FileSizeMax                   int64
	VideoUploadMaxbitrate         int32
	SmallQueueActiveOperationsMax int32
	LargeQueueActiveOperationsMax int32
}

func (*AutoDownloadSettings) CRC() uint32 {
	return 0xbaa57628
}
func (*AutoDownloadSettings) _IAutoDownloadSettings() {}

type IAutoSaveException interface {
	tl.Object
	_IAutoSaveException()
}

var (
	_ IAutoSaveException = (*AutoSaveException)(nil)
)

type AutoSaveException struct {
	Peer     IPeer
	Settings IAutoSaveSettings
}

func (*AutoSaveException) CRC() uint32 {
	return 0x81602d47
}
func (*AutoSaveException) _IAutoSaveException() {}

type IAutoSaveSettings interface {
	tl.Object
	_IAutoSaveSettings()
}

var (
	_ IAutoSaveSettings = (*AutoSaveSettings)(nil)
)

type AutoSaveSettings struct {
	_            struct{} `tl:"flags,bitflag"`
	Photos       bool     `tl:",omitempty:flags:0,implicit"`
	Videos       bool     `tl:",omitempty:flags:1,implicit"`
	VideoMaxSize *int64   `tl:",omitempty:flags:2"`
}

func (*AutoSaveSettings) CRC() uint32 {
	return 0xc84834ce
}
func (*AutoSaveSettings) _IAutoSaveSettings() {}

type IAvailableReaction interface {
	tl.Object
	_IAvailableReaction()
}

var (
	_ IAvailableReaction = (*AvailableReaction)(nil)
)

type AvailableReaction struct {
	_                 struct{} `tl:"flags,bitflag"`
	Inactive          bool     `tl:",omitempty:flags:0,implicit"`
	Premium           bool     `tl:",omitempty:flags:2,implicit"`
	Reaction          string
	Title             string
	StaticIcon        IDocument
	AppearAnimation   IDocument
	SelectAnimation   IDocument
	ActivateAnimation IDocument
	EffectAnimation   IDocument
	AroundAnimation   IDocument `tl:",omitempty:flags:1"`
	CenterIcon        IDocument `tl:",omitempty:flags:1"`
}

func (*AvailableReaction) CRC() uint32 {
	return 0xc077ec01
}
func (*AvailableReaction) _IAvailableReaction() {}

type IBankCardOpenURL interface {
	tl.Object
	_IBankCardOpenURL()
}

var (
	_ IBankCardOpenURL = (*BankCardOpenURL)(nil)
)

type BankCardOpenURL struct {
	URL  string
	Name string
}

func (*BankCardOpenURL) CRC() uint32 {
	return 0xf568028a
}
func (*BankCardOpenURL) _IBankCardOpenURL() {}

type IBoost interface {
	tl.Object
	_IBoost()
}

var (
	_ IBoost = (*Boost)(nil)
)

type Boost struct {
	_             struct{} `tl:"flags,bitflag"`
	Gift          bool     `tl:",omitempty:flags:1,implicit"`
	Giveaway      bool     `tl:",omitempty:flags:2,implicit"`
	Unclaimed     bool     `tl:",omitempty:flags:3,implicit"`
	ID            string
	UserID        *int64 `tl:",omitempty:flags:0"`
	GiveawayMsgID *int32 `tl:",omitempty:flags:2"`
	Date          int32
	Expires       int32
	UsedGiftSlug  *string `tl:",omitempty:flags:4"`
	Multiplier    *int32  `tl:",omitempty:flags:5"`
}

func (*Boost) CRC() uint32 {
	return 0x2a1c8c71
}
func (*Boost) _IBoost() {}

type IBotApp interface {
	tl.Object
	_IBotApp()
}

var (
	_ IBotApp = (*BotAppNotModified)(nil)
	_ IBotApp = (*BotApp)(nil)
)

type BotAppNotModified struct{}

func (*BotAppNotModified) CRC() uint32 {
	return 0x5da674b7
}
func (*BotAppNotModified) _IBotApp() {}

type BotApp struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       IPhoto
	Document    IDocument `tl:",omitempty:flags:0"`
	Hash        int64
}

func (*BotApp) CRC() uint32 {
	return 0x95fcd1d6
}
func (*BotApp) _IBotApp() {}

type IBotCommand interface {
	tl.Object
	_IBotCommand()
}

var (
	_ IBotCommand = (*BotCommand)(nil)
)

type BotCommand struct {
	Command     string
	Description string
}

func (*BotCommand) CRC() uint32 {
	return 0xc27ac8c7
}
func (*BotCommand) _IBotCommand() {}

type IBotCommandScope interface {
	tl.Object
	_IBotCommandScope()
}

var (
	_ IBotCommandScope = (*BotCommandScopeDefault)(nil)
	_ IBotCommandScope = (*BotCommandScopeUsers)(nil)
	_ IBotCommandScope = (*BotCommandScopeChats)(nil)
	_ IBotCommandScope = (*BotCommandScopeChatAdmins)(nil)
	_ IBotCommandScope = (*BotCommandScopePeer)(nil)
	_ IBotCommandScope = (*BotCommandScopePeerAdmins)(nil)
	_ IBotCommandScope = (*BotCommandScopePeerUser)(nil)
)

type BotCommandScopeDefault struct{}

func (*BotCommandScopeDefault) CRC() uint32 {
	return 0x2f6cb2ab
}
func (*BotCommandScopeDefault) _IBotCommandScope() {}

type BotCommandScopeUsers struct{}

func (*BotCommandScopeUsers) CRC() uint32 {
	return 0x3c4f04d8
}
func (*BotCommandScopeUsers) _IBotCommandScope() {}

type BotCommandScopeChats struct{}

func (*BotCommandScopeChats) CRC() uint32 {
	return 0x6fe1a881
}
func (*BotCommandScopeChats) _IBotCommandScope() {}

type BotCommandScopeChatAdmins struct{}

func (*BotCommandScopeChatAdmins) CRC() uint32 {
	return 0xb9aa606a
}
func (*BotCommandScopeChatAdmins) _IBotCommandScope() {}

type BotCommandScopePeer struct {
	Peer IInputPeer
}

func (*BotCommandScopePeer) CRC() uint32 {
	return 0xdb9d897d
}
func (*BotCommandScopePeer) _IBotCommandScope() {}

type BotCommandScopePeerAdmins struct {
	Peer IInputPeer
}

func (*BotCommandScopePeerAdmins) CRC() uint32 {
	return 0x3fd863d1
}
func (*BotCommandScopePeerAdmins) _IBotCommandScope() {}

type BotCommandScopePeerUser struct {
	Peer   IInputPeer
	UserID IInputUser
}

func (*BotCommandScopePeerUser) CRC() uint32 {
	return 0xa1321f3
}
func (*BotCommandScopePeerUser) _IBotCommandScope() {}

type IBotInfo interface {
	tl.Object
	_IBotInfo()
}

var (
	_ IBotInfo = (*BotInfo)(nil)
)

type BotInfo struct {
	_                   struct{}       `tl:"flags,bitflag"`
	UserID              *int64         `tl:",omitempty:flags:0"`
	Description         *string        `tl:",omitempty:flags:1"`
	DescriptionPhoto    IPhoto         `tl:",omitempty:flags:4"`
	DescriptionDocument IDocument      `tl:",omitempty:flags:5"`
	Commands            []IBotCommand  `tl:",omitempty:flags:2"`
	MenuButton          IBotMenuButton `tl:",omitempty:flags:3"`
}

func (*BotInfo) CRC() uint32 {
	return 0x8f300b57
}
func (*BotInfo) _IBotInfo() {}

type IBotInlineMessage interface {
	tl.Object
	_IBotInlineMessage()
}

var (
	_ IBotInlineMessage = (*BotInlineMessageMediaAuto)(nil)
	_ IBotInlineMessage = (*BotInlineMessageText)(nil)
	_ IBotInlineMessage = (*BotInlineMessageMediaGeo)(nil)
	_ IBotInlineMessage = (*BotInlineMessageMediaVenue)(nil)
	_ IBotInlineMessage = (*BotInlineMessageMediaContact)(nil)
	_ IBotInlineMessage = (*BotInlineMessageMediaInvoice)(nil)
	_ IBotInlineMessage = (*BotInlineMessageMediaWebPage)(nil)
)

type BotInlineMessageMediaAuto struct {
	_           struct{} `tl:"flags,bitflag"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup IReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaAuto) CRC() uint32 {
	return 0x764cf810
}
func (*BotInlineMessageMediaAuto) _IBotInlineMessage() {}

type BotInlineMessageText struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup IReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageText) CRC() uint32 {
	return 0x8c7f65e2
}
func (*BotInlineMessageText) _IBotInlineMessage() {}

type BotInlineMessageMediaGeo struct {
	_                           struct{} `tl:"flags,bitflag"`
	Geo                         IGeoPoint
	Heading                     *int32       `tl:",omitempty:flags:0"`
	Period                      *int32       `tl:",omitempty:flags:1"`
	ProximityNotificationRadius *int32       `tl:",omitempty:flags:3"`
	ReplyMarkup                 IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaGeo) CRC() uint32 {
	return 0x51846fd
}
func (*BotInlineMessageMediaGeo) _IBotInlineMessage() {}

type BotInlineMessageMediaVenue struct {
	_           struct{} `tl:"flags,bitflag"`
	Geo         IGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaVenue) CRC() uint32 {
	return 0x8a86659c
}
func (*BotInlineMessageMediaVenue) _IBotInlineMessage() {}

type BotInlineMessageMediaContact struct {
	_           struct{} `tl:"flags,bitflag"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaContact) CRC() uint32 {
	return 0x18d1cdc2
}
func (*BotInlineMessageMediaContact) _IBotInlineMessage() {}

type BotInlineMessageMediaInvoice struct {
	_                        struct{} `tl:"flags,bitflag"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:1,implicit"`
	Test                     bool     `tl:",omitempty:flags:3,implicit"`
	Title                    string
	Description              string
	Photo                    IWebDocument `tl:",omitempty:flags:0"`
	Currency                 string
	TotalAmount              int64
	ReplyMarkup              IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaInvoice) CRC() uint32 {
	return 0x354a9b09
}
func (*BotInlineMessageMediaInvoice) _IBotInlineMessage() {}

type BotInlineMessageMediaWebPage struct {
	_               struct{} `tl:"flags,bitflag"`
	InvertMedia     bool     `tl:",omitempty:flags:3,implicit"`
	ForceLargeMedia bool     `tl:",omitempty:flags:4,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:5,implicit"`
	Manual          bool     `tl:",omitempty:flags:7,implicit"`
	Safe            bool     `tl:",omitempty:flags:8,implicit"`
	Message         string
	Entities        []IMessageEntity `tl:",omitempty:flags:1"`
	URL             string
	ReplyMarkup     IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaWebPage) CRC() uint32 {
	return 0x809ad9a6
}
func (*BotInlineMessageMediaWebPage) _IBotInlineMessage() {}

type IBotInlineResult interface {
	tl.Object
	_IBotInlineResult()
}

var (
	_ IBotInlineResult = (*BotInlineResult)(nil)
	_ IBotInlineResult = (*BotInlineMediaResult)(nil)
)

type BotInlineResult struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string      `tl:",omitempty:flags:1"`
	Description *string      `tl:",omitempty:flags:2"`
	URL         *string      `tl:",omitempty:flags:3"`
	Thumb       IWebDocument `tl:",omitempty:flags:4"`
	Content     IWebDocument `tl:",omitempty:flags:5"`
	SendMessage IBotInlineMessage
}

func (*BotInlineResult) CRC() uint32 {
	return 0x11965f3a
}
func (*BotInlineResult) _IBotInlineResult() {}

type BotInlineMediaResult struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Photo       IPhoto    `tl:",omitempty:flags:0"`
	Document    IDocument `tl:",omitempty:flags:1"`
	Title       *string   `tl:",omitempty:flags:2"`
	Description *string   `tl:",omitempty:flags:3"`
	SendMessage IBotInlineMessage
}

func (*BotInlineMediaResult) CRC() uint32 {
	return 0x17db940b
}
func (*BotInlineMediaResult) _IBotInlineResult() {}

type IBotMenuButton interface {
	tl.Object
	_IBotMenuButton()
}

var (
	_ IBotMenuButton = (*BotMenuButtonDefault)(nil)
	_ IBotMenuButton = (*BotMenuButtonCommands)(nil)
	_ IBotMenuButton = (*BotMenuButton)(nil)
)

type BotMenuButtonDefault struct{}

func (*BotMenuButtonDefault) CRC() uint32 {
	return 0x7533a588
}
func (*BotMenuButtonDefault) _IBotMenuButton() {}

type BotMenuButtonCommands struct{}

func (*BotMenuButtonCommands) CRC() uint32 {
	return 0x4258c205
}
func (*BotMenuButtonCommands) _IBotMenuButton() {}

type BotMenuButton struct {
	Text string
	URL  string
}

func (*BotMenuButton) CRC() uint32 {
	return 0xc7b57ce6
}
func (*BotMenuButton) _IBotMenuButton() {}

type ICdnConfig interface {
	tl.Object
	_ICdnConfig()
}

var (
	_ ICdnConfig = (*CdnConfig)(nil)
)

type CdnConfig struct {
	PublicKeys []ICdnPublicKey
}

func (*CdnConfig) CRC() uint32 {
	return 0x5725e40a
}
func (*CdnConfig) _ICdnConfig() {}

type ICdnPublicKey interface {
	tl.Object
	_ICdnPublicKey()
}

var (
	_ ICdnPublicKey = (*CdnPublicKey)(nil)
)

type CdnPublicKey struct {
	DcID      int32
	PublicKey string
}

func (*CdnPublicKey) CRC() uint32 {
	return 0xc982eaba
}
func (*CdnPublicKey) _ICdnPublicKey() {}

type IChannelAdminLogEvent interface {
	tl.Object
	_IChannelAdminLogEvent()
}

var (
	_ IChannelAdminLogEvent = (*ChannelAdminLogEvent)(nil)
)

type ChannelAdminLogEvent struct {
	ID     int64
	Date   int32
	UserID int64
	Action IChannelAdminLogEventAction
}

func (*ChannelAdminLogEvent) CRC() uint32 {
	return 0x1fad68cd
}
func (*ChannelAdminLogEvent) _IChannelAdminLogEvent() {}

type IChannelAdminLogEventAction interface {
	tl.Object
	_IChannelAdminLogEventAction()
}

var (
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeTitle)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeAbout)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeUsername)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangePhoto)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleInvites)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleSignatures)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionUpdatePinned)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionEditMessage)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionDeleteMessage)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoin)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantLeave)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantInvite)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantToggleBan)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantToggleAdmin)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeStickerSet)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionTogglePreHistoryHidden)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionDefaultBannedRights)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionStopPoll)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeLinkedChat)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeLocation)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleSlowMode)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionStartGroupCall)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionDiscardGroupCall)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantMute)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantUnmute)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleGroupCallSetting)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoinByInvite)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteDelete)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteRevoke)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteEdit)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantVolume)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeHistoryTTL)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoinByRequest)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleNoForwards)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionSendMessage)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeAvailableReactions)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeUsernames)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleForum)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionCreateTopic)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionEditTopic)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionDeleteTopic)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionPinTopic)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleAntiSpam)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangePeerColor)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeProfilePeerColor)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeWallpaper)(nil)
	_ IChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeEmojiStatus)(nil)
)

type ChannelAdminLogEventActionChangeTitle struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeTitle) CRC() uint32 {
	return 0xe6dfb825
}
func (*ChannelAdminLogEventActionChangeTitle) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAbout struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeAbout) CRC() uint32 {
	return 0x55188a2e
}
func (*ChannelAdminLogEventActionChangeAbout) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsername struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeUsername) CRC() uint32 {
	return 0x6a4afc38
}
func (*ChannelAdminLogEventActionChangeUsername) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePhoto struct {
	PrevPhoto IPhoto
	NewPhoto  IPhoto
}

func (*ChannelAdminLogEventActionChangePhoto) CRC() uint32 {
	return 0x434bd2af
}
func (*ChannelAdminLogEventActionChangePhoto) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleInvites struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleInvites) CRC() uint32 {
	return 0x1b7907ae
}
func (*ChannelAdminLogEventActionToggleInvites) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSignatures struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleSignatures) CRC() uint32 {
	return 0x26ae0971
}
func (*ChannelAdminLogEventActionToggleSignatures) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionUpdatePinned struct {
	Message IMessage
}

func (*ChannelAdminLogEventActionUpdatePinned) CRC() uint32 {
	return 0xe9e82c18
}
func (*ChannelAdminLogEventActionUpdatePinned) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditMessage struct {
	PrevMessage IMessage
	NewMessage  IMessage
}

func (*ChannelAdminLogEventActionEditMessage) CRC() uint32 {
	return 0x709b2405
}
func (*ChannelAdminLogEventActionEditMessage) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteMessage struct {
	Message IMessage
}

func (*ChannelAdminLogEventActionDeleteMessage) CRC() uint32 {
	return 0x42e047bb
}
func (*ChannelAdminLogEventActionDeleteMessage) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoin struct{}

func (*ChannelAdminLogEventActionParticipantJoin) CRC() uint32 {
	return 0x183040d3
}
func (*ChannelAdminLogEventActionParticipantJoin) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantLeave struct{}

func (*ChannelAdminLogEventActionParticipantLeave) CRC() uint32 {
	return 0xf89777f2
}
func (*ChannelAdminLogEventActionParticipantLeave) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantInvite struct {
	Participant IChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantInvite) CRC() uint32 {
	return 0xe31c34d8
}
func (*ChannelAdminLogEventActionParticipantInvite) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleBan struct {
	PrevParticipant IChannelParticipant
	NewParticipant  IChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantToggleBan) CRC() uint32 {
	return 0xe6d83d7e
}
func (*ChannelAdminLogEventActionParticipantToggleBan) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleAdmin struct {
	PrevParticipant IChannelParticipant
	NewParticipant  IChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantToggleAdmin) CRC() uint32 {
	return 0xd5676710
}
func (*ChannelAdminLogEventActionParticipantToggleAdmin) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeStickerSet struct {
	PrevStickerset IInputStickerSet
	NewStickerset  IInputStickerSet
}

func (*ChannelAdminLogEventActionChangeStickerSet) CRC() uint32 {
	return 0xb1c3caa7
}
func (*ChannelAdminLogEventActionChangeStickerSet) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionTogglePreHistoryHidden struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionTogglePreHistoryHidden) CRC() uint32 {
	return 0x5f5c95f1
}
func (*ChannelAdminLogEventActionTogglePreHistoryHidden) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDefaultBannedRights struct {
	PrevBannedRights IChatBannedRights
	NewBannedRights  IChatBannedRights
}

func (*ChannelAdminLogEventActionDefaultBannedRights) CRC() uint32 {
	return 0x2df5fc0a
}
func (*ChannelAdminLogEventActionDefaultBannedRights) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStopPoll struct {
	Message IMessage
}

func (*ChannelAdminLogEventActionStopPoll) CRC() uint32 {
	return 0x8f079643
}
func (*ChannelAdminLogEventActionStopPoll) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLinkedChat struct {
	PrevValue int64
	NewValue  int64
}

func (*ChannelAdminLogEventActionChangeLinkedChat) CRC() uint32 {
	return 0x50c7ac8
}
func (*ChannelAdminLogEventActionChangeLinkedChat) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLocation struct {
	PrevValue IChannelLocation
	NewValue  IChannelLocation
}

func (*ChannelAdminLogEventActionChangeLocation) CRC() uint32 {
	return 0xe6b76ae
}
func (*ChannelAdminLogEventActionChangeLocation) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSlowMode struct {
	PrevValue int32
	NewValue  int32
}

func (*ChannelAdminLogEventActionToggleSlowMode) CRC() uint32 {
	return 0x53909779
}
func (*ChannelAdminLogEventActionToggleSlowMode) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStartGroupCall struct {
	Call IInputGroupCall
}

func (*ChannelAdminLogEventActionStartGroupCall) CRC() uint32 {
	return 0x23209745
}
func (*ChannelAdminLogEventActionStartGroupCall) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDiscardGroupCall struct {
	Call IInputGroupCall
}

func (*ChannelAdminLogEventActionDiscardGroupCall) CRC() uint32 {
	return 0xdb9f9140
}
func (*ChannelAdminLogEventActionDiscardGroupCall) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantMute struct {
	Participant IGroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantMute) CRC() uint32 {
	return 0xf92424d2
}
func (*ChannelAdminLogEventActionParticipantMute) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantUnmute struct {
	Participant IGroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantUnmute) CRC() uint32 {
	return 0xe64429c0
}
func (*ChannelAdminLogEventActionParticipantUnmute) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleGroupCallSetting struct {
	JoinMuted bool
}

func (*ChannelAdminLogEventActionToggleGroupCallSetting) CRC() uint32 {
	return 0x56d6a247
}
func (*ChannelAdminLogEventActionToggleGroupCallSetting) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByInvite struct {
	_           struct{} `tl:"flags,bitflag"`
	ViaChatlist bool     `tl:",omitempty:flags:0,implicit"`
	Invite      IExportedChatInvite
}

func (*ChannelAdminLogEventActionParticipantJoinByInvite) CRC() uint32 {
	return 0xfe9fc158
}
func (*ChannelAdminLogEventActionParticipantJoinByInvite) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteDelete struct {
	Invite IExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteDelete) CRC() uint32 {
	return 0x5a50fca4
}
func (*ChannelAdminLogEventActionExportedInviteDelete) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteRevoke struct {
	Invite IExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteRevoke) CRC() uint32 {
	return 0x410a134e
}
func (*ChannelAdminLogEventActionExportedInviteRevoke) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteEdit struct {
	PrevInvite IExportedChatInvite
	NewInvite  IExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteEdit) CRC() uint32 {
	return 0xe90ebb59
}
func (*ChannelAdminLogEventActionExportedInviteEdit) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantVolume struct {
	Participant IGroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantVolume) CRC() uint32 {
	return 0x3e7f6847
}
func (*ChannelAdminLogEventActionParticipantVolume) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeHistoryTTL struct {
	PrevValue int32
	NewValue  int32
}

func (*ChannelAdminLogEventActionChangeHistoryTTL) CRC() uint32 {
	return 0x6e941a38
}
func (*ChannelAdminLogEventActionChangeHistoryTTL) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByRequest struct {
	Invite     IExportedChatInvite
	ApprovedBy int64
}

func (*ChannelAdminLogEventActionParticipantJoinByRequest) CRC() uint32 {
	return 0xafb6144a
}
func (*ChannelAdminLogEventActionParticipantJoinByRequest) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleNoForwards struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleNoForwards) CRC() uint32 {
	return 0xcb2ac766
}
func (*ChannelAdminLogEventActionToggleNoForwards) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionSendMessage struct {
	Message IMessage
}

func (*ChannelAdminLogEventActionSendMessage) CRC() uint32 {
	return 0x278f2868
}
func (*ChannelAdminLogEventActionSendMessage) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAvailableReactions struct {
	PrevValue IChatReactions
	NewValue  IChatReactions
}

func (*ChannelAdminLogEventActionChangeAvailableReactions) CRC() uint32 {
	return 0xbe4e0ef8
}
func (*ChannelAdminLogEventActionChangeAvailableReactions) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsernames struct {
	PrevValue []string
	NewValue  []string
}

func (*ChannelAdminLogEventActionChangeUsernames) CRC() uint32 {
	return 0xf04fb3a9
}
func (*ChannelAdminLogEventActionChangeUsernames) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleForum struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleForum) CRC() uint32 {
	return 0x2cc6383
}
func (*ChannelAdminLogEventActionToggleForum) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionCreateTopic struct {
	Topic IForumTopic
}

func (*ChannelAdminLogEventActionCreateTopic) CRC() uint32 {
	return 0x58707d28
}
func (*ChannelAdminLogEventActionCreateTopic) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditTopic struct {
	PrevTopic IForumTopic
	NewTopic  IForumTopic
}

func (*ChannelAdminLogEventActionEditTopic) CRC() uint32 {
	return 0xf06fe208
}
func (*ChannelAdminLogEventActionEditTopic) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteTopic struct {
	Topic IForumTopic
}

func (*ChannelAdminLogEventActionDeleteTopic) CRC() uint32 {
	return 0xae168909
}
func (*ChannelAdminLogEventActionDeleteTopic) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionPinTopic struct {
	_         struct{}    `tl:"flags,bitflag"`
	PrevTopic IForumTopic `tl:",omitempty:flags:0"`
	NewTopic  IForumTopic `tl:",omitempty:flags:1"`
}

func (*ChannelAdminLogEventActionPinTopic) CRC() uint32 {
	return 0x5d8d353b
}
func (*ChannelAdminLogEventActionPinTopic) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleAntiSpam struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleAntiSpam) CRC() uint32 {
	return 0x64f36dfc
}
func (*ChannelAdminLogEventActionToggleAntiSpam) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePeerColor struct {
	PrevValue IPeerColor
	NewValue  IPeerColor
}

func (*ChannelAdminLogEventActionChangePeerColor) CRC() uint32 {
	return 0x5796e780
}
func (*ChannelAdminLogEventActionChangePeerColor) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeProfilePeerColor struct {
	PrevValue IPeerColor
	NewValue  IPeerColor
}

func (*ChannelAdminLogEventActionChangeProfilePeerColor) CRC() uint32 {
	return 0x5e477b25
}
func (*ChannelAdminLogEventActionChangeProfilePeerColor) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeWallpaper struct {
	PrevValue IWallPaper
	NewValue  IWallPaper
}

func (*ChannelAdminLogEventActionChangeWallpaper) CRC() uint32 {
	return 0x31bb5d52
}
func (*ChannelAdminLogEventActionChangeWallpaper) _IChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeEmojiStatus struct {
	PrevValue IEmojiStatus
	NewValue  IEmojiStatus
}

func (*ChannelAdminLogEventActionChangeEmojiStatus) CRC() uint32 {
	return 0x3ea9feb1
}
func (*ChannelAdminLogEventActionChangeEmojiStatus) _IChannelAdminLogEventAction() {}

type IChannelAdminLogEventsFilter interface {
	tl.Object
	_IChannelAdminLogEventsFilter()
}

var (
	_ IChannelAdminLogEventsFilter = (*ChannelAdminLogEventsFilter)(nil)
)

type ChannelAdminLogEventsFilter struct {
	_         struct{} `tl:"flags,bitflag"`
	Join      bool     `tl:",omitempty:flags:0,implicit"`
	Leave     bool     `tl:",omitempty:flags:1,implicit"`
	Invite    bool     `tl:",omitempty:flags:2,implicit"`
	Ban       bool     `tl:",omitempty:flags:3,implicit"`
	Unban     bool     `tl:",omitempty:flags:4,implicit"`
	Kick      bool     `tl:",omitempty:flags:5,implicit"`
	Unkick    bool     `tl:",omitempty:flags:6,implicit"`
	Promote   bool     `tl:",omitempty:flags:7,implicit"`
	Demote    bool     `tl:",omitempty:flags:8,implicit"`
	Info      bool     `tl:",omitempty:flags:9,implicit"`
	Settings  bool     `tl:",omitempty:flags:10,implicit"`
	Pinned    bool     `tl:",omitempty:flags:11,implicit"`
	Edit      bool     `tl:",omitempty:flags:12,implicit"`
	Delete    bool     `tl:",omitempty:flags:13,implicit"`
	GroupCall bool     `tl:",omitempty:flags:14,implicit"`
	Invites   bool     `tl:",omitempty:flags:15,implicit"`
	Send      bool     `tl:",omitempty:flags:16,implicit"`
	Forums    bool     `tl:",omitempty:flags:17,implicit"`
}

func (*ChannelAdminLogEventsFilter) CRC() uint32 {
	return 0xea107ae4
}
func (*ChannelAdminLogEventsFilter) _IChannelAdminLogEventsFilter() {}

type IChannelLocation interface {
	tl.Object
	_IChannelLocation()
}

var (
	_ IChannelLocation = (*ChannelLocationEmpty)(nil)
	_ IChannelLocation = (*ChannelLocation)(nil)
)

type ChannelLocationEmpty struct{}

func (*ChannelLocationEmpty) CRC() uint32 {
	return 0xbfb5ad8b
}
func (*ChannelLocationEmpty) _IChannelLocation() {}

type ChannelLocation struct {
	GeoPoint IGeoPoint
	Address  string
}

func (*ChannelLocation) CRC() uint32 {
	return 0x209b82db
}
func (*ChannelLocation) _IChannelLocation() {}

type IChannelMessagesFilter interface {
	tl.Object
	_IChannelMessagesFilter()
}

var (
	_ IChannelMessagesFilter = (*ChannelMessagesFilterEmpty)(nil)
	_ IChannelMessagesFilter = (*ChannelMessagesFilter)(nil)
)

type ChannelMessagesFilterEmpty struct{}

func (*ChannelMessagesFilterEmpty) CRC() uint32 {
	return 0x94d42ee7
}
func (*ChannelMessagesFilterEmpty) _IChannelMessagesFilter() {}

type ChannelMessagesFilter struct {
	_                  struct{} `tl:"flags,bitflag"`
	ExcludeNewMessages bool     `tl:",omitempty:flags:1,implicit"`
	Ranges             []IMessageRange
}

func (*ChannelMessagesFilter) CRC() uint32 {
	return 0xcd77d957
}
func (*ChannelMessagesFilter) _IChannelMessagesFilter() {}

type IChannelParticipant interface {
	tl.Object
	_IChannelParticipant()
}

var (
	_ IChannelParticipant = (*ChannelParticipant)(nil)
	_ IChannelParticipant = (*ChannelParticipantSelf)(nil)
	_ IChannelParticipant = (*ChannelParticipantCreator)(nil)
	_ IChannelParticipant = (*ChannelParticipantAdmin)(nil)
	_ IChannelParticipant = (*ChannelParticipantBanned)(nil)
	_ IChannelParticipant = (*ChannelParticipantLeft)(nil)
)

type ChannelParticipant struct {
	UserID int64
	Date   int32
}

func (*ChannelParticipant) CRC() uint32 {
	return 0xc00c07c0
}
func (*ChannelParticipant) _IChannelParticipant() {}

type ChannelParticipantSelf struct {
	_          struct{} `tl:"flags,bitflag"`
	ViaRequest bool     `tl:",omitempty:flags:0,implicit"`
	UserID     int64
	InviterID  int64
	Date       int32
}

func (*ChannelParticipantSelf) CRC() uint32 {
	return 0x35a8bfa7
}
func (*ChannelParticipantSelf) _IChannelParticipant() {}

type ChannelParticipantCreator struct {
	_           struct{} `tl:"flags,bitflag"`
	UserID      int64
	AdminRights IChatAdminRights
	Rank        *string `tl:",omitempty:flags:0"`
}

func (*ChannelParticipantCreator) CRC() uint32 {
	return 0x2fe601d3
}
func (*ChannelParticipantCreator) _IChannelParticipant() {}

type ChannelParticipantAdmin struct {
	_           struct{} `tl:"flags,bitflag"`
	CanEdit     bool     `tl:",omitempty:flags:0,implicit"`
	Self        bool     `tl:",omitempty:flags:1,implicit"`
	UserID      int64
	InviterID   *int64 `tl:",omitempty:flags:1"`
	PromotedBy  int64
	Date        int32
	AdminRights IChatAdminRights
	Rank        *string `tl:",omitempty:flags:2"`
}

func (*ChannelParticipantAdmin) CRC() uint32 {
	return 0x34c3bb53
}
func (*ChannelParticipantAdmin) _IChannelParticipant() {}

type ChannelParticipantBanned struct {
	_            struct{} `tl:"flags,bitflag"`
	Left         bool     `tl:",omitempty:flags:0,implicit"`
	Peer         IPeer
	KickedBy     int64
	Date         int32
	BannedRights IChatBannedRights
}

func (*ChannelParticipantBanned) CRC() uint32 {
	return 0x6df8014e
}
func (*ChannelParticipantBanned) _IChannelParticipant() {}

type ChannelParticipantLeft struct {
	Peer IPeer
}

func (*ChannelParticipantLeft) CRC() uint32 {
	return 0x1b03f006
}
func (*ChannelParticipantLeft) _IChannelParticipant() {}

type IChannelParticipantsFilter interface {
	tl.Object
	_IChannelParticipantsFilter()
}

var (
	_ IChannelParticipantsFilter = (*ChannelParticipantsRecent)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsAdmins)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsKicked)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsBots)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsBanned)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsSearch)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsContacts)(nil)
	_ IChannelParticipantsFilter = (*ChannelParticipantsMentions)(nil)
)

type ChannelParticipantsRecent struct{}

func (*ChannelParticipantsRecent) CRC() uint32 {
	return 0xde3f3c79
}
func (*ChannelParticipantsRecent) _IChannelParticipantsFilter() {}

type ChannelParticipantsAdmins struct{}

func (*ChannelParticipantsAdmins) CRC() uint32 {
	return 0xb4608969
}
func (*ChannelParticipantsAdmins) _IChannelParticipantsFilter() {}

type ChannelParticipantsKicked struct {
	Q string
}

func (*ChannelParticipantsKicked) CRC() uint32 {
	return 0xa3b54985
}
func (*ChannelParticipantsKicked) _IChannelParticipantsFilter() {}

type ChannelParticipantsBots struct{}

func (*ChannelParticipantsBots) CRC() uint32 {
	return 0xb0d1865b
}
func (*ChannelParticipantsBots) _IChannelParticipantsFilter() {}

type ChannelParticipantsBanned struct {
	Q string
}

func (*ChannelParticipantsBanned) CRC() uint32 {
	return 0x1427a5e1
}
func (*ChannelParticipantsBanned) _IChannelParticipantsFilter() {}

type ChannelParticipantsSearch struct {
	Q string
}

func (*ChannelParticipantsSearch) CRC() uint32 {
	return 0x656ac4b
}
func (*ChannelParticipantsSearch) _IChannelParticipantsFilter() {}

type ChannelParticipantsContacts struct {
	Q string
}

func (*ChannelParticipantsContacts) CRC() uint32 {
	return 0xbb6ae88d
}
func (*ChannelParticipantsContacts) _IChannelParticipantsFilter() {}

type ChannelParticipantsMentions struct {
	_        struct{} `tl:"flags,bitflag"`
	Q        *string  `tl:",omitempty:flags:0"`
	TopMsgID *int32   `tl:",omitempty:flags:1"`
}

func (*ChannelParticipantsMentions) CRC() uint32 {
	return 0xe04b5ceb
}
func (*ChannelParticipantsMentions) _IChannelParticipantsFilter() {}

// Object defines a group.
type IChat interface {
	tl.Object
	_IChat()
}

var (
	_ IChat = (*ChatEmpty)(nil)
	_ IChat = (*Chat)(nil)
	_ IChat = (*ChatForbidden)(nil)
	_ IChat = (*Channel)(nil)
	_ IChat = (*ChannelForbidden)(nil)
)

type ChatEmpty struct {
	ID int64 // Group identifier
}

func (*ChatEmpty) CRC() uint32 {
	return 0x29562865
}
func (*ChatEmpty) _IChat() {}

// Info about a group
type Chat struct {
	_                   struct{}          `tl:"flags,bitflag"`
	Creator             bool              `tl:",omitempty:flags:0,implicit"` // Whether the current user is the creator of the group
	Left                bool              `tl:",omitempty:flags:2,implicit"` // Whether the current user has left the group
	Deactivated         bool              `tl:",omitempty:flags:5,implicit"` // Whether the group was migrated
	CallActive          bool              `tl:",omitempty:flags:23,implicit"`
	CallNotEmpty        bool              `tl:",omitempty:flags:24,implicit"`
	Noforwards          bool              `tl:",omitempty:flags:25,implicit"`
	ID                  int64             // ID of the group
	Title               string            // Title
	Photo               IChatPhoto        // Chat photo
	ParticipantsCount   int32             // Participant count
	Date                int32             // Date of creation of the group
	Version             int32             // Used in basic groups to reorder updates and make sure that all of them were received.
	MigratedTo          IInputChannel     `tl:",omitempty:flags:6"`  // Means this chat was upgraded to a supergroup
	AdminRights         IChatAdminRights  `tl:",omitempty:flags:14"` // Admin rights of the user in the group
	DefaultBannedRights IChatBannedRights `tl:",omitempty:flags:18"` // Default banned rights of all users in the group
}

func (*Chat) CRC() uint32 {
	return 0x41cbf256
}
func (*Chat) _IChat() {}

// A group to which the user has no access. E.g., because the user was kicked from the group.
type ChatForbidden struct {
	ID    int64  // User identifier
	Title string // Group name
}

func (*ChatForbidden) CRC() uint32 {
	return 0x6592a1a7
}
func (*ChatForbidden) _IChat() {}

// Channel/supergroup info
type Channel struct {
	_                   struct{}             `tl:"flags,bitflag"`
	Creator             bool                 `tl:",omitempty:flags:0,implicit"`  // Whether the current user is the creator of this channel
	Left                bool                 `tl:",omitempty:flags:2,implicit"`  // Whether the current user has left this channel
	Broadcast           bool                 `tl:",omitempty:flags:5,implicit"`  // Is this a channel?
	Verified            bool                 `tl:",omitempty:flags:7,implicit"`  // Is this channel verified by telegram?
	Megagroup           bool                 `tl:",omitempty:flags:8,implicit"`  // Is this a supergroup?
	Restricted          bool                 `tl:",omitempty:flags:9,implicit"`  // Whether viewing/writing in this channel for a reason (see restriction_reason
	Signatures          bool                 `tl:",omitempty:flags:11,implicit"` // Whether signatures are enabled (channels)
	Min                 bool                 `tl:",omitempty:flags:12,implicit"` // See min
	Scam                bool                 `tl:",omitempty:flags:19,implicit"` // This channel/supergroup is probably a scam
	HasLink             bool                 `tl:",omitempty:flags:20,implicit"` // Whether this channel has a private join link
	HasGeo              bool                 `tl:",omitempty:flags:21,implicit"` // Whether this chanel has a geoposition
	SlowmodeEnabled     bool                 `tl:",omitempty:flags:22,implicit"` // Whether slow mode is enabled for groups to prevent flood in chat
	CallActive          bool                 `tl:",omitempty:flags:23,implicit"`
	CallNotEmpty        bool                 `tl:",omitempty:flags:24,implicit"`
	Fake                bool                 `tl:",omitempty:flags:25,implicit"`
	Gigagroup           bool                 `tl:",omitempty:flags:26,implicit"`
	Noforwards          bool                 `tl:",omitempty:flags:27,implicit"`
	JoinToSend          bool                 `tl:",omitempty:flags:28,implicit"`
	JoinRequest         bool                 `tl:",omitempty:flags:29,implicit"`
	Forum               bool                 `tl:",omitempty:flags:30,implicit"`
	_                   struct{}             `tl:"flags2,bitflag"`
	StoriesHidden       bool                 `tl:",omitempty:flags2:1,implicit"`
	StoriesHiddenMin    bool                 `tl:",omitempty:flags2:2,implicit"`
	StoriesUnavailable  bool                 `tl:",omitempty:flags2:3,implicit"`
	ID                  int64                // ID of the channel
	AccessHash          *int64               `tl:",omitempty:flags:13"` // Access hash
	Title               string               // Title
	Username            *string              `tl:",omitempty:flags:6"` // Username
	Photo               IChatPhoto           // Profile photo
	Date                int32                // Date when the user joined the supergroup/channel, or if the user isn't a member, its creation date
	RestrictionReason   []IRestrictionReason `tl:",omitempty:flags:9"`  // Contains the reason why access to this channel must be restricted.
	AdminRights         IChatAdminRights     `tl:",omitempty:flags:14"` // Admin rights of the user in this channel (see rights)
	BannedRights        IChatBannedRights    `tl:",omitempty:flags:15"` // Banned rights of the user in this channel (see rights)
	DefaultBannedRights IChatBannedRights    `tl:",omitempty:flags:18"` // Default chat rights (see rights)
	ParticipantsCount   *int32               `tl:",omitempty:flags:17"` // Participant count
	Usernames           []IUsername          `tl:",omitempty:flags2:0"`
	StoriesMaxID        *int32               `tl:",omitempty:flags2:4"`
	Color               IPeerColor           `tl:",omitempty:flags2:7"`
	ProfileColor        IPeerColor           `tl:",omitempty:flags2:8"`
	EmojiStatus         IEmojiStatus         `tl:",omitempty:flags2:9"`
	Level               *int32               `tl:",omitempty:flags2:10"`
}

func (*Channel) CRC() uint32 {
	return 0xaadfc8f
}
func (*Channel) _IChat() {}

// Indicates a channel/supergroup we can't access because we were banned, or for some other reason.
type ChannelForbidden struct {
	_          struct{} `tl:"flags,bitflag"`
	Broadcast  bool     `tl:",omitempty:flags:5,implicit"` // Is this a channel
	Megagroup  bool     `tl:",omitempty:flags:8,implicit"` // Is this a supergroup
	ID         int64    // Channel ID
	AccessHash int64    // Access hash
	Title      string   // Title
	UntilDate  *int32   `tl:",omitempty:flags:16"` // The ban is valid until the specified date
}

func (*ChannelForbidden) CRC() uint32 {
	return 0x17d493d5
}
func (*ChannelForbidden) _IChat() {}

type IChatAdminRights interface {
	tl.Object
	_IChatAdminRights()
}

var (
	_ IChatAdminRights = (*ChatAdminRights)(nil)
)

type ChatAdminRights struct {
	_              struct{} `tl:"flags,bitflag"`
	ChangeInfo     bool     `tl:",omitempty:flags:0,implicit"`
	PostMessages   bool     `tl:",omitempty:flags:1,implicit"`
	EditMessages   bool     `tl:",omitempty:flags:2,implicit"`
	DeleteMessages bool     `tl:",omitempty:flags:3,implicit"`
	BanUsers       bool     `tl:",omitempty:flags:4,implicit"`
	InviteUsers    bool     `tl:",omitempty:flags:5,implicit"`
	PinMessages    bool     `tl:",omitempty:flags:7,implicit"`
	AddAdmins      bool     `tl:",omitempty:flags:9,implicit"`
	Anonymous      bool     `tl:",omitempty:flags:10,implicit"`
	ManageCall     bool     `tl:",omitempty:flags:11,implicit"`
	Other          bool     `tl:",omitempty:flags:12,implicit"`
	ManageTopics   bool     `tl:",omitempty:flags:13,implicit"`
	PostStories    bool     `tl:",omitempty:flags:14,implicit"`
	EditStories    bool     `tl:",omitempty:flags:15,implicit"`
	DeleteStories  bool     `tl:",omitempty:flags:16,implicit"`
}

func (*ChatAdminRights) CRC() uint32 {
	return 0x5fb224d5
}
func (*ChatAdminRights) _IChatAdminRights() {}

type IChatAdminWithInvites interface {
	tl.Object
	_IChatAdminWithInvites()
}

var (
	_ IChatAdminWithInvites = (*ChatAdminWithInvites)(nil)
)

type ChatAdminWithInvites struct {
	AdminID             int64
	InvitesCount        int32
	RevokedInvitesCount int32
}

func (*ChatAdminWithInvites) CRC() uint32 {
	return 0xf2ecef23
}
func (*ChatAdminWithInvites) _IChatAdminWithInvites() {}

type IChatBannedRights interface {
	tl.Object
	_IChatBannedRights()
}

var (
	_ IChatBannedRights = (*ChatBannedRights)(nil)
)

type ChatBannedRights struct {
	_               struct{} `tl:"flags,bitflag"`
	ViewMessages    bool     `tl:",omitempty:flags:0,implicit"`
	SendMessages    bool     `tl:",omitempty:flags:1,implicit"`
	SendMedia       bool     `tl:",omitempty:flags:2,implicit"`
	SendStickers    bool     `tl:",omitempty:flags:3,implicit"`
	SendGifs        bool     `tl:",omitempty:flags:4,implicit"`
	SendGames       bool     `tl:",omitempty:flags:5,implicit"`
	SendInline      bool     `tl:",omitempty:flags:6,implicit"`
	EmbedLinks      bool     `tl:",omitempty:flags:7,implicit"`
	SendPolls       bool     `tl:",omitempty:flags:8,implicit"`
	ChangeInfo      bool     `tl:",omitempty:flags:10,implicit"`
	InviteUsers     bool     `tl:",omitempty:flags:15,implicit"`
	PinMessages     bool     `tl:",omitempty:flags:17,implicit"`
	ManageTopics    bool     `tl:",omitempty:flags:18,implicit"`
	SendPhotos      bool     `tl:",omitempty:flags:19,implicit"`
	SendVideos      bool     `tl:",omitempty:flags:20,implicit"`
	SendRoundvideos bool     `tl:",omitempty:flags:21,implicit"`
	SendAudios      bool     `tl:",omitempty:flags:22,implicit"`
	SendVoices      bool     `tl:",omitempty:flags:23,implicit"`
	SendDocs        bool     `tl:",omitempty:flags:24,implicit"`
	SendPlain       bool     `tl:",omitempty:flags:25,implicit"`
	UntilDate       int32
}

func (*ChatBannedRights) CRC() uint32 {
	return 0x9f120418
}
func (*ChatBannedRights) _IChatBannedRights() {}

// Object containing detailed group info
type IChatFull interface {
	tl.Object
	_IChatFull()
}

var (
	_ IChatFull = (*ChatFull)(nil)
	_ IChatFull = (*ChannelFull)(nil)
)

// Detailed chat info
type ChatFull struct {
	_                      struct{}            `tl:"flags,bitflag"`
	CanSetUsername         bool                `tl:",omitempty:flags:7,implicit"` // Can we change the username of this chat
	HasScheduled           bool                `tl:",omitempty:flags:8,implicit"` // Whether scheduled messages are available
	TranslationsDisabled   bool                `tl:",omitempty:flags:19,implicit"`
	ID                     int64               // ID of the chat
	About                  string              // About string for this chat
	Participants           IChatParticipants   // Participant list
	ChatPhoto              IPhoto              `tl:",omitempty:flags:2"` // Chat photo
	NotifySettings         IPeerNotifySettings // Notification settings
	ExportedInvite         IExportedChatInvite `tl:",omitempty:flags:13"` // Chat invite
	BotInfo                []IBotInfo          `tl:",omitempty:flags:3"`  // Info about bots that are in this chat
	PinnedMsgID            *int32              `tl:",omitempty:flags:6"`  // Message ID of the last pinned message
	FolderID               *int32              `tl:",omitempty:flags:11"` // Peer folder ID, for more info click here
	Call                   IInputGroupCall     `tl:",omitempty:flags:12"`
	TTLPeriod              *int32              `tl:",omitempty:flags:14"`
	GroupcallDefaultJoinAs IPeer               `tl:",omitempty:flags:15"`
	ThemeEmoticon          *string             `tl:",omitempty:flags:16"`
	RequestsPending        *int32              `tl:",omitempty:flags:17"`
	RecentRequesters       []int64             `tl:",omitempty:flags:17"`
	AvailableReactions     IChatReactions      `tl:",omitempty:flags:18"`
}

func (*ChatFull) CRC() uint32 {
	return 0xc9d31138
}
func (*ChatFull) _IChatFull() {}

// Full info about a channel/supergroup
type ChannelFull struct {
	_                      struct{}            `tl:"flags,bitflag"`
	CanViewParticipants    bool                `tl:",omitempty:flags:3,implicit"`  // Can we vew the participant list?
	CanSetUsername         bool                `tl:",omitempty:flags:6,implicit"`  // Can we set the channel's username?
	CanSetStickers         bool                `tl:",omitempty:flags:7,implicit"`  // Can we associate a stickerpack to the supergroup?
	HiddenPrehistory       bool                `tl:",omitempty:flags:10,implicit"` // Is the history before we joined hidden to us?
	CanSetLocation         bool                `tl:",omitempty:flags:16,implicit"` // Can we set the geolocation of this group (for geogroups)
	HasScheduled           bool                `tl:",omitempty:flags:19,implicit"` // Whether scheduled messages are available
	CanViewStats           bool                `tl:",omitempty:flags:20,implicit"` // Can the user view channel/supergroup statistics
	Blocked                bool                `tl:",omitempty:flags:22,implicit"` // Whether any anonymous admin of this supergroup was blocked: if set, you won't receive messages from anonymous group admins in discussion replies via @replies
	_                      struct{}            `tl:"flags2,bitflag"`
	CanDeleteChannel       bool                `tl:",omitempty:flags2:0,implicit"`
	Antispam               bool                `tl:",omitempty:flags2:1,implicit"`
	ParticipantsHidden     bool                `tl:",omitempty:flags2:2,implicit"`
	TranslationsDisabled   bool                `tl:",omitempty:flags2:3,implicit"`
	StoriesPinnedAvailable bool                `tl:",omitempty:flags2:5,implicit"`
	ViewForumAsMessages    bool                `tl:",omitempty:flags2:6,implicit"`
	ID                     int64               // ID of the channel
	About                  string              // Info about the channel
	ParticipantsCount      *int32              `tl:",omitempty:flags:0"`  // Number of participants of the channel
	AdminsCount            *int32              `tl:",omitempty:flags:1"`  // Number of channel admins
	KickedCount            *int32              `tl:",omitempty:flags:2"`  // Number of users kicked from the channel
	BannedCount            *int32              `tl:",omitempty:flags:2"`  // Number of users banned from the channel
	OnlineCount            *int32              `tl:",omitempty:flags:13"` // Number of users currently online
	ReadInboxMaxID         int32               // Position up to which all incoming messages are read.
	ReadOutboxMaxID        int32               // Position up to which all outgoing messages are read.
	UnreadCount            int32               // Count of unread messages
	ChatPhoto              IPhoto              // Channel picture
	NotifySettings         IPeerNotifySettings // Notification settings
	ExportedInvite         IExportedChatInvite `tl:",omitempty:flags:23"` // Invite link
	BotInfo                []IBotInfo          // Info about bots in the channel/supergrup
	MigratedFromChatID     *int64              `tl:",omitempty:flags:4"`  // The chat ID from which this group was migrated
	MigratedFromMaxID      *int32              `tl:",omitempty:flags:4"`  // The message ID in the original chat at which this group was migrated
	PinnedMsgID            *int32              `tl:",omitempty:flags:5"`  // Message ID of the last pinned message
	Stickerset             IStickerSet         `tl:",omitempty:flags:8"`  // Associated stickerset
	AvailableMinID         *int32              `tl:",omitempty:flags:9"`  // Identifier of a maximum unavailable message in a channel due to hidden history.
	FolderID               *int32              `tl:",omitempty:flags:11"` // Peer folder ID, for more info click here
	LinkedChatID           *int64              `tl:",omitempty:flags:14"` // ID of the linked discussion chat for channels
	Location               IChannelLocation    `tl:",omitempty:flags:15"` // Location of the geogroup
	SlowmodeSeconds        *int32              `tl:",omitempty:flags:17"` // If specified, users in supergroups will only be able to send one message every slowmode_seconds seconds
	SlowmodeNextSendDate   *int32              `tl:",omitempty:flags:18"` // Indicates when the user will be allowed to send another message in the supergroup (unixdate)
	StatsDc                *int32              `tl:",omitempty:flags:12"` // If set, specifies the DC to use for fetching channel statistics
	Pts                    int32               // Latest PTS for this channel
	Call                   IInputGroupCall     `tl:",omitempty:flags:21"`
	TTLPeriod              *int32              `tl:",omitempty:flags:24"`
	PendingSuggestions     []string            `tl:",omitempty:flags:25"`
	GroupcallDefaultJoinAs IPeer               `tl:",omitempty:flags:26"`
	ThemeEmoticon          *string             `tl:",omitempty:flags:27"`
	RequestsPending        *int32              `tl:",omitempty:flags:28"`
	RecentRequesters       []int64             `tl:",omitempty:flags:28"`
	DefaultSendAs          IPeer               `tl:",omitempty:flags:29"`
	AvailableReactions     IChatReactions      `tl:",omitempty:flags:30"`
	Stories                IPeerStories        `tl:",omitempty:flags2:4"`
	Wallpaper              IWallPaper          `tl:",omitempty:flags2:7"`
}

func (*ChannelFull) CRC() uint32 {
	return 0xf2bcb6f
}
func (*ChannelFull) _IChatFull() {}

type IChatInvite interface {
	tl.Object
	_IChatInvite()
}

var (
	_ IChatInvite = (*ChatInviteAlready)(nil)
	_ IChatInvite = (*ChatInvite)(nil)
	_ IChatInvite = (*ChatInvitePeek)(nil)
)

type ChatInviteAlready struct {
	Chat IChat
}

func (*ChatInviteAlready) CRC() uint32 {
	return 0x5a686d7c
}
func (*ChatInviteAlready) _IChatInvite() {}

type ChatInvite struct {
	_                 struct{} `tl:"flags,bitflag"`
	Channel           bool     `tl:",omitempty:flags:0,implicit"`
	Broadcast         bool     `tl:",omitempty:flags:1,implicit"`
	Public            bool     `tl:",omitempty:flags:2,implicit"`
	Megagroup         bool     `tl:",omitempty:flags:3,implicit"`
	RequestNeeded     bool     `tl:",omitempty:flags:6,implicit"`
	Verified          bool     `tl:",omitempty:flags:7,implicit"`
	Scam              bool     `tl:",omitempty:flags:8,implicit"`
	Fake              bool     `tl:",omitempty:flags:9,implicit"`
	Title             string
	About             *string `tl:",omitempty:flags:5"`
	Photo             IPhoto
	ParticipantsCount int32
	Participants      []IUser `tl:",omitempty:flags:4"`
	Color             int32
}

func (*ChatInvite) CRC() uint32 {
	return 0xcde0ec40
}
func (*ChatInvite) _IChatInvite() {}

type ChatInvitePeek struct {
	Chat    IChat
	Expires int32
}

func (*ChatInvitePeek) CRC() uint32 {
	return 0x61695cb0
}
func (*ChatInvitePeek) _IChatInvite() {}

type IChatInviteImporter interface {
	tl.Object
	_IChatInviteImporter()
}

var (
	_ IChatInviteImporter = (*ChatInviteImporter)(nil)
)

type ChatInviteImporter struct {
	_           struct{} `tl:"flags,bitflag"`
	Requested   bool     `tl:",omitempty:flags:0,implicit"`
	ViaChatlist bool     `tl:",omitempty:flags:3,implicit"`
	UserID      int64
	Date        int32
	About       *string `tl:",omitempty:flags:2"`
	ApprovedBy  *int64  `tl:",omitempty:flags:1"`
}

func (*ChatInviteImporter) CRC() uint32 {
	return 0x8c5adfd9
}
func (*ChatInviteImporter) _IChatInviteImporter() {}

type IChatOnlines interface {
	tl.Object
	_IChatOnlines()
}

var (
	_ IChatOnlines = (*ChatOnlines)(nil)
)

type ChatOnlines struct {
	Onlines int32
}

func (*ChatOnlines) CRC() uint32 {
	return 0xf041e250
}
func (*ChatOnlines) _IChatOnlines() {}

type IChatParticipant interface {
	tl.Object
	_IChatParticipant()
}

var (
	_ IChatParticipant = (*ChatParticipant)(nil)
	_ IChatParticipant = (*ChatParticipantCreator)(nil)
	_ IChatParticipant = (*ChatParticipantAdmin)(nil)
)

type ChatParticipant struct {
	UserID    int64
	InviterID int64
	Date      int32
}

func (*ChatParticipant) CRC() uint32 {
	return 0xc02d4007
}
func (*ChatParticipant) _IChatParticipant() {}

type ChatParticipantCreator struct {
	UserID int64
}

func (*ChatParticipantCreator) CRC() uint32 {
	return 0xe46bcee4
}
func (*ChatParticipantCreator) _IChatParticipant() {}

type ChatParticipantAdmin struct {
	UserID    int64
	InviterID int64
	Date      int32
}

func (*ChatParticipantAdmin) CRC() uint32 {
	return 0xa0933f5b
}
func (*ChatParticipantAdmin) _IChatParticipant() {}

type IChatParticipants interface {
	tl.Object
	_IChatParticipants()
}

var (
	_ IChatParticipants = (*ChatParticipantsForbidden)(nil)
	_ IChatParticipants = (*ChatParticipants)(nil)
)

type ChatParticipantsForbidden struct {
	_               struct{} `tl:"flags,bitflag"`
	ChatID          int64
	SelfParticipant IChatParticipant `tl:",omitempty:flags:0"`
}

func (*ChatParticipantsForbidden) CRC() uint32 {
	return 0x8763d3e1
}
func (*ChatParticipantsForbidden) _IChatParticipants() {}

type ChatParticipants struct {
	ChatID       int64
	Participants []IChatParticipant
	Version      int32
}

func (*ChatParticipants) CRC() uint32 {
	return 0x3cbc93f8
}
func (*ChatParticipants) _IChatParticipants() {}

type IChatPhoto interface {
	tl.Object
	_IChatPhoto()
}

var (
	_ IChatPhoto = (*ChatPhotoEmpty)(nil)
	_ IChatPhoto = (*ChatPhoto)(nil)
)

type ChatPhotoEmpty struct{}

func (*ChatPhotoEmpty) CRC() uint32 {
	return 0x37c1011c
}
func (*ChatPhotoEmpty) _IChatPhoto() {}

type ChatPhoto struct {
	_             struct{} `tl:"flags,bitflag"`
	HasVideo      bool     `tl:",omitempty:flags:0,implicit"`
	PhotoID       int64
	StrippedThumb *[]byte `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*ChatPhoto) CRC() uint32 {
	return 0x1c6e1c11
}
func (*ChatPhoto) _IChatPhoto() {}

type IChatReactions interface {
	tl.Object
	_IChatReactions()
}

var (
	_ IChatReactions = (*ChatReactionsNone)(nil)
	_ IChatReactions = (*ChatReactionsAll)(nil)
	_ IChatReactions = (*ChatReactionsSome)(nil)
)

type ChatReactionsNone struct{}

func (*ChatReactionsNone) CRC() uint32 {
	return 0xeafc32bc
}
func (*ChatReactionsNone) _IChatReactions() {}

type ChatReactionsAll struct {
	_           struct{} `tl:"flags,bitflag"`
	AllowCustom bool     `tl:",omitempty:flags:0,implicit"`
}

func (*ChatReactionsAll) CRC() uint32 {
	return 0x52928bca
}
func (*ChatReactionsAll) _IChatReactions() {}

type ChatReactionsSome struct {
	Reactions []IReaction
}

func (*ChatReactionsSome) CRC() uint32 {
	return 0x661d4037
}
func (*ChatReactionsSome) _IChatReactions() {}

// Settings for the code type to send
type ICodeSettings interface {
	tl.Object
	_ICodeSettings()
}

var (
	_ ICodeSettings = (*CodeSettings)(nil)
)

// Settings used by telegram servers for sending the confirm code. Example implementations: [telegram for android](https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/ui/LoginActivity.java), [tdlib](https://github.com/tdlib/td/tree/master/td/telegram/SendCodeHelper.cpp).
type CodeSettings struct {
	_               struct{} `tl:"flags,bitflag"`
	AllowFlashcall  bool     `tl:",omitempty:flags:0,implicit"` // Whether to allow phone verification via [phone calls](https://core.telegram.org/api/auth).
	CurrentNumber   bool     `tl:",omitempty:flags:1,implicit"` // Pass true if the phone number is used on the current device.Ignored if allow_flashcall is not set.
	AllowAppHash    bool     `tl:",omitempty:flags:4,implicit"` // If a token that will be included in eventually sent SMSs is required: required in newer versions of android, to use the [android SMS receiver APIs](https://developers.google.com/identity/sms-retriever/overview)
	AllowMissedCall bool     `tl:",omitempty:flags:5,implicit"` // Whether this device supports receiving the code using the [auth.codeTypeMissedCall](https://core.telegram.org/constructor/auth.codeTypeMissedCall) method
	AllowFirebase   bool     `tl:",omitempty:flags:7,implicit"` // Whether Firebase auth is supported
	LogoutTokens    [][]byte `tl:",omitempty:flags:6"`          // Previously stored future auth tokens, [see the documentation for more info](https://core.telegram.org/api/auth#future-auth-tokens)
	Token           *string  `tl:",omitempty:flags:8"`          // Used only by official iOS apps for Firebase auth: device token for apple push.
	AppSandbox      *bool    `tl:",omitempty:flags:8"`          // Used only by official iOS apps for firebase auth: whether a sandbox - certificate will be used during transmission of the push notification.
}

func (*CodeSettings) CRC() uint32 {
	return 0xad253d78
}
func (*CodeSettings) _ICodeSettings() {}

// Object contains info on API configuring parameters.
type IConfig interface {
	tl.Object
	_IConfig()
}

var (
	_ IConfig = (*Config)(nil)
)

// Current configuration
type Config struct {
	_                       struct{}    `tl:"flags,bitflag"`
	DefaultP2PContacts      bool        `tl:",omitempty:flags:3,implicit"`  // Whether the client should use P2P by default for phone calls with contacts
	PreloadFeaturedStickers bool        `tl:",omitempty:flags:4,implicit"`  // Whether the client should preload featured stickers
	RevokePmInbox           bool        `tl:",omitempty:flags:6,implicit"`  // Whether incoming private messages can be deleted for both participants
	BlockedMode             bool        `tl:",omitempty:flags:8,implicit"`  // Indicates that telegram is probably censored by governments/ISPs in the current region
	ForceTryIpv6            bool        `tl:",omitempty:flags:14,implicit"` // Whether to forcefully connect using IPv6 [dcOptions](https://core.telegram.org/type/DcOption), even if the client knows that IPv4 is available.
	Date                    int32       // Current date at the server
	Expires                 int32       // Expiration date of this config: when it expires it'll have to be refetched using [help.getConfig](https://core.telegram.org/method/help.getConfig)
	TestMode                bool        // Whether we're connected to the test DCs
	ThisDc                  int32       // ID of the DC that returned the reply
	DcOptions               []IDcOption // DC IP list
	DcTxtDomainName         string      // Domain name for fetching encrypted DC list from DNS TXT record
	ChatSizeMax             int32       // Maximum member count for normal [groups](https://core.telegram.org/api/channel)
	MegagroupSizeMax        int32       // Maximum member count for [supergroups](https://core.telegram.org/api/channel)
	ForwardedCountMax       int32       // Maximum number of messages that can be forwarded at once using [messages.forwardMessages](https://core.telegram.org/method/messages.forwardMessages).
	OnlineUpdatePeriodMs    int32       // The client should [update its online status](https://core.telegram.org/method/account.updateStatus) every N milliseconds
	OfflineBlurTimeoutMs    int32       // Delay before offline status needs to be sent to the server
	OfflineIdleTimeoutMs    int32       // Time without any user activity after which it should be treated offline
	OnlineCloudTimeoutMs    int32       // If we are offline, but were online from some other client in last `online_cloud_timeout_ms` milliseconds after we had gone offline, then delay offline notification for `notify_cloud_delay_ms` milliseconds.
	NotifyCloudDelayMs      int32       // If we are offline, but online from some other client then delay sending the offline notification for `notify_cloud_delay_ms` milliseconds.
	NotifyDefaultDelayMs    int32       // If some other client is online, then delay notification for `notification_default_delay_ms` milliseconds
	PushChatPeriodMs        int32       // Not for client use
	PushChatLimit           int32       // Not for client use
	EditTimeLimit           int32       // Only messages with age smaller than the one specified can be edited
	RevokeTimeLimit         int32       // Only channel/supergroup messages with age smaller than the specified can be deleted
	RevokePmTimeLimit       int32       // Only private messages with age smaller than the specified can be deleted
	RatingEDecay            int32       // Exponential decay rate for computing [top peer rating](https://core.telegram.org/api/top-rating)
	StickersRecentLimit     int32       // Maximum number of recent stickers
	ChannelsReadMediaPeriod int32       // Indicates that round videos (video notes) and voice messages sent in channels and older than the specified period must be marked as read
	TmpSessions             *int32      `tl:",omitempty:flags:0"` // Temporary [passport](https://core.telegram.org/passport) sessions
	CallReceiveTimeoutMs    int32       // Maximum allowed outgoing ring time in VoIP calls: if the user we're calling doesn't reply within the specified time (in milliseconds), we should hang up the call
	CallRingTimeoutMs       int32       // Maximum allowed incoming ring time in VoIP calls: if the current user doesn't reply within the specified time (in milliseconds), the call will be automatically refused
	CallConnectTimeoutMs    int32       // VoIP connection timeout: if the instance of libtgvoip on the other side of the call doesn't connect to our instance of libtgvoip within the specified time (in milliseconds), the call must be aborted
	CallPacketTimeoutMs     int32       // If during a VoIP call a packet isn't received for the specified period of time, the call must be aborted
	MeURLPrefix             string      // The domain to use to parse [deep links](https://core.telegram.org/api/links).
	AutoupdateURLPrefix     *string     `tl:",omitempty:flags:7"`  // URL to use to auto-update the current app
	GifSearchUsername       *string     `tl:",omitempty:flags:9"`  // Username of the bot to use to search for GIFs
	VenueSearchUsername     *string     `tl:",omitempty:flags:10"` // Username of the bot to use to search for venues
	ImgSearchUsername       *string     `tl:",omitempty:flags:11"` // Username of the bot to use for image search
	StaticMapsProvider      *string     `tl:",omitempty:flags:12"` // ID of the map provider to use for venues
	CaptionLengthMax        int32       // Maximum length of caption (length in utf8 codepoints)
	MessageLengthMax        int32       // Maximum length of messages (length in utf8 codepoints)
	WebfileDcID             int32       // DC ID to use to download [webfiles](https://core.telegram.org/api/files#downloading-webfiles)
	SuggestedLangCode       *string     `tl:",omitempty:flags:2"`  // Suggested language code
	LangPackVersion         *int32      `tl:",omitempty:flags:2"`  // Language pack version
	BaseLangPackVersion     *int32      `tl:",omitempty:flags:2"`  // Basic language pack version
	ReactionsDefault        IReaction   `tl:",omitempty:flags:15"` // Default [message reaction](https://core.telegram.org/api/reactions)
	AutologinToken          *string     `tl:",omitempty:flags:16"` // Autologin token, [more info on URL authorization](https://core.telegram.org/api/url-authorization#link-url-authorization).
}

func (*Config) CRC() uint32 {
	return 0xcc1a241e
}
func (*Config) _IConfig() {}

type IContact interface {
	tl.Object
	_IContact()
}

var (
	_ IContact = (*Contact)(nil)
)

type Contact struct {
	UserID int64
	Mutual bool
}

func (*Contact) CRC() uint32 {
	return 0x145ade0b
}
func (*Contact) _IContact() {}

type IContactStatus interface {
	tl.Object
	_IContactStatus()
}

var (
	_ IContactStatus = (*ContactStatus)(nil)
)

type ContactStatus struct {
	UserID int64
	Status IUserStatus
}

func (*ContactStatus) CRC() uint32 {
	return 0x16d9703b
}
func (*ContactStatus) _IContactStatus() {}

type IDataJSON interface {
	tl.Object
	_IDataJSON()
}

var (
	_ IDataJSON = (*DataJSON)(nil)
)

type DataJSON struct {
	Data string
}

func (*DataJSON) CRC() uint32 {
	return 0x7d748d04
}
func (*DataJSON) _IDataJSON() {}

type IDcOption interface {
	tl.Object
	_IDcOption()
}

var (
	_ IDcOption = (*DcOption)(nil)
)

type DcOption struct {
	_            struct{} `tl:"flags,bitflag"`
	Ipv6         bool     `tl:",omitempty:flags:0,implicit"`
	MediaOnly    bool     `tl:",omitempty:flags:1,implicit"`
	TcpoOnly     bool     `tl:",omitempty:flags:2,implicit"`
	Cdn          bool     `tl:",omitempty:flags:3,implicit"`
	Static       bool     `tl:",omitempty:flags:4,implicit"`
	ThisPortOnly bool     `tl:",omitempty:flags:5,implicit"`
	ID           int32
	IpAddress    string
	Port         int32
	Secret       *[]byte `tl:",omitempty:flags:10"`
}

func (*DcOption) CRC() uint32 {
	return 0x18b7a10d
}
func (*DcOption) _IDcOption() {}

type IDefaultHistoryTTL interface {
	tl.Object
	_IDefaultHistoryTTL()
}

var (
	_ IDefaultHistoryTTL = (*DefaultHistoryTTL)(nil)
)

type DefaultHistoryTTL struct {
	Period int32
}

func (*DefaultHistoryTTL) CRC() uint32 {
	return 0x43b46b20
}
func (*DefaultHistoryTTL) _IDefaultHistoryTTL() {}

type IDialog interface {
	tl.Object
	_IDialog()
}

var (
	_ IDialog = (*Dialog)(nil)
	_ IDialog = (*DialogFolder)(nil)
)

type Dialog struct {
	_                    struct{} `tl:"flags,bitflag"`
	Pinned               bool     `tl:",omitempty:flags:2,implicit"`
	UnreadMark           bool     `tl:",omitempty:flags:3,implicit"`
	ViewForumAsMessages  bool     `tl:",omitempty:flags:6,implicit"`
	Peer                 IPeer
	TopMessage           int32
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	UnreadMentionsCount  int32
	UnreadReactionsCount int32
	NotifySettings       IPeerNotifySettings
	Pts                  *int32        `tl:",omitempty:flags:0"`
	Draft                IDraftMessage `tl:",omitempty:flags:1"`
	FolderID             *int32        `tl:",omitempty:flags:4"`
	TTLPeriod            *int32        `tl:",omitempty:flags:5"`
}

func (*Dialog) CRC() uint32 {
	return 0xd58a08c6
}
func (*Dialog) _IDialog() {}

type DialogFolder struct {
	_                          struct{} `tl:"flags,bitflag"`
	Pinned                     bool     `tl:",omitempty:flags:2,implicit"`
	Folder                     IFolder
	Peer                       IPeer
	TopMessage                 int32
	UnreadMutedPeersCount      int32
	UnreadUnmutedPeersCount    int32
	UnreadMutedMessagesCount   int32
	UnreadUnmutedMessagesCount int32
}

func (*DialogFolder) CRC() uint32 {
	return 0x71bd134c
}
func (*DialogFolder) _IDialog() {}

type IDialogFilter interface {
	tl.Object
	_IDialogFilter()
}

var (
	_ IDialogFilter = (*DialogFilter)(nil)
	_ IDialogFilter = (*DialogFilterDefault)(nil)
	_ IDialogFilter = (*DialogFilterChatlist)(nil)
)

type DialogFilter struct {
	_               struct{} `tl:"flags,bitflag"`
	Contacts        bool     `tl:",omitempty:flags:0,implicit"`
	NonContacts     bool     `tl:",omitempty:flags:1,implicit"`
	Groups          bool     `tl:",omitempty:flags:2,implicit"`
	Broadcasts      bool     `tl:",omitempty:flags:3,implicit"`
	Bots            bool     `tl:",omitempty:flags:4,implicit"`
	ExcludeMuted    bool     `tl:",omitempty:flags:11,implicit"`
	ExcludeRead     bool     `tl:",omitempty:flags:12,implicit"`
	ExcludeArchived bool     `tl:",omitempty:flags:13,implicit"`
	ID              int32
	Title           string
	Emoticon        *string `tl:",omitempty:flags:25"`
	PinnedPeers     []IInputPeer
	IncludePeers    []IInputPeer
	ExcludePeers    []IInputPeer
}

func (*DialogFilter) CRC() uint32 {
	return 0x7438f7e8
}
func (*DialogFilter) _IDialogFilter() {}

type DialogFilterDefault struct{}

func (*DialogFilterDefault) CRC() uint32 {
	return 0x363293ae
}
func (*DialogFilterDefault) _IDialogFilter() {}

type DialogFilterChatlist struct {
	_            struct{} `tl:"flags,bitflag"`
	HasMyInvites bool     `tl:",omitempty:flags:26,implicit"`
	ID           int32
	Title        string
	Emoticon     *string `tl:",omitempty:flags:25"`
	PinnedPeers  []IInputPeer
	IncludePeers []IInputPeer
}

func (*DialogFilterChatlist) CRC() uint32 {
	return 0xd64a04a8
}
func (*DialogFilterChatlist) _IDialogFilter() {}

type IDialogFilterSuggested interface {
	tl.Object
	_IDialogFilterSuggested()
}

var (
	_ IDialogFilterSuggested = (*DialogFilterSuggested)(nil)
)

type DialogFilterSuggested struct {
	Filter      IDialogFilter
	Description string
}

func (*DialogFilterSuggested) CRC() uint32 {
	return 0x77744d4a
}
func (*DialogFilterSuggested) _IDialogFilterSuggested() {}

type IDialogPeer interface {
	tl.Object
	_IDialogPeer()
}

var (
	_ IDialogPeer = (*DialogPeer)(nil)
	_ IDialogPeer = (*DialogPeerFolder)(nil)
)

type DialogPeer struct {
	Peer IPeer
}

func (*DialogPeer) CRC() uint32 {
	return 0xe56dbf05
}
func (*DialogPeer) _IDialogPeer() {}

type DialogPeerFolder struct {
	FolderID int32
}

func (*DialogPeerFolder) CRC() uint32 {
	return 0x514519e2
}
func (*DialogPeerFolder) _IDialogPeer() {}

type IDocument interface {
	tl.Object
	_IDocument()
}

var (
	_ IDocument = (*DocumentEmpty)(nil)
	_ IDocument = (*Document)(nil)
)

type DocumentEmpty struct {
	ID int64
}

func (*DocumentEmpty) CRC() uint32 {
	return 0x36f8c871
}
func (*DocumentEmpty) _IDocument() {}

type Document struct {
	_             struct{} `tl:"flags,bitflag"`
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	MimeType      string
	Size          int64
	Thumbs        []IPhotoSize `tl:",omitempty:flags:0"`
	VideoThumbs   []IVideoSize `tl:",omitempty:flags:1"`
	DcID          int32
	Attributes    []IDocumentAttribute
}

func (*Document) CRC() uint32 {
	return 0x8fd4c4d8
}
func (*Document) _IDocument() {}

type IDocumentAttribute interface {
	tl.Object
	_IDocumentAttribute()
}

var (
	_ IDocumentAttribute = (*DocumentAttributeImageSize)(nil)
	_ IDocumentAttribute = (*DocumentAttributeAnimated)(nil)
	_ IDocumentAttribute = (*DocumentAttributeSticker)(nil)
	_ IDocumentAttribute = (*DocumentAttributeVideo)(nil)
	_ IDocumentAttribute = (*DocumentAttributeAudio)(nil)
	_ IDocumentAttribute = (*DocumentAttributeFilename)(nil)
	_ IDocumentAttribute = (*DocumentAttributeHasStickers)(nil)
	_ IDocumentAttribute = (*DocumentAttributeCustomEmoji)(nil)
)

type DocumentAttributeImageSize struct {
	W int32
	H int32
}

func (*DocumentAttributeImageSize) CRC() uint32 {
	return 0x6c37c15c
}
func (*DocumentAttributeImageSize) _IDocumentAttribute() {}

type DocumentAttributeAnimated struct{}

func (*DocumentAttributeAnimated) CRC() uint32 {
	return 0x11b58939
}
func (*DocumentAttributeAnimated) _IDocumentAttribute() {}

type DocumentAttributeSticker struct {
	_          struct{} `tl:"flags,bitflag"`
	Mask       bool     `tl:",omitempty:flags:1,implicit"`
	Alt        string
	Stickerset IInputStickerSet
	MaskCoords IMaskCoords `tl:",omitempty:flags:0"`
}

func (*DocumentAttributeSticker) CRC() uint32 {
	return 0x6319d612
}
func (*DocumentAttributeSticker) _IDocumentAttribute() {}

type DocumentAttributeVideo struct {
	_                 struct{} `tl:"flags,bitflag"`
	RoundMessage      bool     `tl:",omitempty:flags:0,implicit"`
	SupportsStreaming bool     `tl:",omitempty:flags:1,implicit"`
	Nosound           bool     `tl:",omitempty:flags:3,implicit"`
	Duration          float64
	W                 int32
	H                 int32
	PreloadPrefixSize *int32 `tl:",omitempty:flags:2"`
}

func (*DocumentAttributeVideo) CRC() uint32 {
	return 0xd38ff1c2
}
func (*DocumentAttributeVideo) _IDocumentAttribute() {}

type DocumentAttributeAudio struct {
	_         struct{} `tl:"flags,bitflag"`
	Voice     bool     `tl:",omitempty:flags:10,implicit"`
	Duration  int32
	Title     *string `tl:",omitempty:flags:0"`
	Performer *string `tl:",omitempty:flags:1"`
	Waveform  *[]byte `tl:",omitempty:flags:2"`
}

func (*DocumentAttributeAudio) CRC() uint32 {
	return 0x9852f9c6
}
func (*DocumentAttributeAudio) _IDocumentAttribute() {}

type DocumentAttributeFilename struct {
	FileName string
}

func (*DocumentAttributeFilename) CRC() uint32 {
	return 0x15590068
}
func (*DocumentAttributeFilename) _IDocumentAttribute() {}

type DocumentAttributeHasStickers struct{}

func (*DocumentAttributeHasStickers) CRC() uint32 {
	return 0x9801d2f7
}
func (*DocumentAttributeHasStickers) _IDocumentAttribute() {}

type DocumentAttributeCustomEmoji struct {
	_          struct{} `tl:"flags,bitflag"`
	Free       bool     `tl:",omitempty:flags:0,implicit"`
	TextColor  bool     `tl:",omitempty:flags:1,implicit"`
	Alt        string
	Stickerset IInputStickerSet
}

func (*DocumentAttributeCustomEmoji) CRC() uint32 {
	return 0xfd149899
}
func (*DocumentAttributeCustomEmoji) _IDocumentAttribute() {}

type IDraftMessage interface {
	tl.Object
	_IDraftMessage()
}

var (
	_ IDraftMessage = (*DraftMessageEmpty)(nil)
	_ IDraftMessage = (*DraftMessage)(nil)
)

type DraftMessageEmpty struct {
	_    struct{} `tl:"flags,bitflag"`
	Date *int32   `tl:",omitempty:flags:0"`
}

func (*DraftMessageEmpty) CRC() uint32 {
	return 0x1b0c841a
}
func (*DraftMessageEmpty) _IDraftMessage() {}

type DraftMessage struct {
	_           struct{}      `tl:"flags,bitflag"`
	NoWebpage   bool          `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool          `tl:",omitempty:flags:6,implicit"`
	ReplyTo     IInputReplyTo `tl:",omitempty:flags:4"`
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:3"`
	Media       IInputMedia      `tl:",omitempty:flags:5"`
	Date        int32
}

func (*DraftMessage) CRC() uint32 {
	return 0x3fccf7ef
}
func (*DraftMessage) _IDraftMessage() {}

type IEmailVerification interface {
	tl.Object
	_IEmailVerification()
}

var (
	_ IEmailVerification = (*EmailVerificationCode)(nil)
	_ IEmailVerification = (*EmailVerificationGoogle)(nil)
	_ IEmailVerification = (*EmailVerificationApple)(nil)
)

type EmailVerificationCode struct {
	Code string
}

func (*EmailVerificationCode) CRC() uint32 {
	return 0x922e55a9
}
func (*EmailVerificationCode) _IEmailVerification() {}

type EmailVerificationGoogle struct {
	Token string
}

func (*EmailVerificationGoogle) CRC() uint32 {
	return 0xdb909ec2
}
func (*EmailVerificationGoogle) _IEmailVerification() {}

type EmailVerificationApple struct {
	Token string
}

func (*EmailVerificationApple) CRC() uint32 {
	return 0x96d074fd
}
func (*EmailVerificationApple) _IEmailVerification() {}

type IEmailVerifyPurpose interface {
	tl.Object
	_IEmailVerifyPurpose()
}

var (
	_ IEmailVerifyPurpose = (*EmailVerifyPurposeLoginSetup)(nil)
	_ IEmailVerifyPurpose = (*EmailVerifyPurposeLoginChange)(nil)
	_ IEmailVerifyPurpose = (*EmailVerifyPurposePassport)(nil)
)

type EmailVerifyPurposeLoginSetup struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*EmailVerifyPurposeLoginSetup) CRC() uint32 {
	return 0x4345be73
}
func (*EmailVerifyPurposeLoginSetup) _IEmailVerifyPurpose() {}

type EmailVerifyPurposeLoginChange struct{}

func (*EmailVerifyPurposeLoginChange) CRC() uint32 {
	return 0x527d22eb
}
func (*EmailVerifyPurposeLoginChange) _IEmailVerifyPurpose() {}

type EmailVerifyPurposePassport struct{}

func (*EmailVerifyPurposePassport) CRC() uint32 {
	return 0xbbf51685
}
func (*EmailVerifyPurposePassport) _IEmailVerifyPurpose() {}

type IEmojiGroup interface {
	tl.Object
	_IEmojiGroup()
}

var (
	_ IEmojiGroup = (*EmojiGroup)(nil)
)

type EmojiGroup struct {
	Title       string
	IconEmojiID int64
	Emoticons   []string
}

func (*EmojiGroup) CRC() uint32 {
	return 0x7a9abda9
}
func (*EmojiGroup) _IEmojiGroup() {}

type IEmojiKeyword interface {
	tl.Object
	_IEmojiKeyword()
}

var (
	_ IEmojiKeyword = (*EmojiKeyword)(nil)
	_ IEmojiKeyword = (*EmojiKeywordDeleted)(nil)
)

type EmojiKeyword struct {
	Keyword   string
	Emoticons []string
}

func (*EmojiKeyword) CRC() uint32 {
	return 0xd5b3b9f9
}
func (*EmojiKeyword) _IEmojiKeyword() {}

type EmojiKeywordDeleted struct {
	Keyword   string
	Emoticons []string
}

func (*EmojiKeywordDeleted) CRC() uint32 {
	return 0x236df622
}
func (*EmojiKeywordDeleted) _IEmojiKeyword() {}

type IEmojiKeywordsDifference interface {
	tl.Object
	_IEmojiKeywordsDifference()
}

var (
	_ IEmojiKeywordsDifference = (*EmojiKeywordsDifference)(nil)
)

type EmojiKeywordsDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Keywords    []IEmojiKeyword
}

func (*EmojiKeywordsDifference) CRC() uint32 {
	return 0x5cc761bd
}
func (*EmojiKeywordsDifference) _IEmojiKeywordsDifference() {}

type IEmojiLanguage interface {
	tl.Object
	_IEmojiLanguage()
}

var (
	_ IEmojiLanguage = (*EmojiLanguage)(nil)
)

type EmojiLanguage struct {
	LangCode string
}

func (*EmojiLanguage) CRC() uint32 {
	return 0xb3fb5361
}
func (*EmojiLanguage) _IEmojiLanguage() {}

type IEmojiList interface {
	tl.Object
	_IEmojiList()
}

var (
	_ IEmojiList = (*EmojiListNotModified)(nil)
	_ IEmojiList = (*EmojiList)(nil)
)

type EmojiListNotModified struct{}

func (*EmojiListNotModified) CRC() uint32 {
	return 0x481eadfa
}
func (*EmojiListNotModified) _IEmojiList() {}

type EmojiList struct {
	Hash       int64
	DocumentID []int64
}

func (*EmojiList) CRC() uint32 {
	return 0x7a1e11d1
}
func (*EmojiList) _IEmojiList() {}

type IEmojiStatus interface {
	tl.Object
	_IEmojiStatus()
}

var (
	_ IEmojiStatus = (*EmojiStatusEmpty)(nil)
	_ IEmojiStatus = (*EmojiStatus)(nil)
	_ IEmojiStatus = (*EmojiStatusUntil)(nil)
)

type EmojiStatusEmpty struct{}

func (*EmojiStatusEmpty) CRC() uint32 {
	return 0x2de11aae
}
func (*EmojiStatusEmpty) _IEmojiStatus() {}

type EmojiStatus struct {
	DocumentID int64
}

func (*EmojiStatus) CRC() uint32 {
	return 0x929b619d
}
func (*EmojiStatus) _IEmojiStatus() {}

type EmojiStatusUntil struct {
	DocumentID int64
	Until      int32
}

func (*EmojiStatusUntil) CRC() uint32 {
	return 0xfa30a8c7
}
func (*EmojiStatusUntil) _IEmojiStatus() {}

type IEmojiURL interface {
	tl.Object
	_IEmojiURL()
}

var (
	_ IEmojiURL = (*EmojiURL)(nil)
)

type EmojiURL struct {
	URL string
}

func (*EmojiURL) CRC() uint32 {
	return 0xa575739d
}
func (*EmojiURL) _IEmojiURL() {}

type IEncryptedChat interface {
	tl.Object
	_IEncryptedChat()
}

var (
	_ IEncryptedChat = (*EncryptedChatEmpty)(nil)
	_ IEncryptedChat = (*EncryptedChatWaiting)(nil)
	_ IEncryptedChat = (*EncryptedChatRequested)(nil)
	_ IEncryptedChat = (*EncryptedChat)(nil)
	_ IEncryptedChat = (*EncryptedChatDiscarded)(nil)
)

type EncryptedChatEmpty struct {
	ID int32
}

func (*EncryptedChatEmpty) CRC() uint32 {
	return 0xab7ec0a0
}
func (*EncryptedChatEmpty) _IEncryptedChat() {}

type EncryptedChatWaiting struct {
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
}

func (*EncryptedChatWaiting) CRC() uint32 {
	return 0x66b25953
}
func (*EncryptedChatWaiting) _IEncryptedChat() {}

type EncryptedChatRequested struct {
	_             struct{} `tl:"flags,bitflag"`
	FolderID      *int32   `tl:",omitempty:flags:0"`
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GA            []byte
}

func (*EncryptedChatRequested) CRC() uint32 {
	return 0x48f1d94c
}
func (*EncryptedChatRequested) _IEncryptedChat() {}

type EncryptedChat struct {
	ID             int32
	AccessHash     int64
	Date           int32
	AdminID        int64
	ParticipantID  int64
	GAOrB          []byte
	KeyFingerprint int64
}

func (*EncryptedChat) CRC() uint32 {
	return 0x61f0d4c7
}
func (*EncryptedChat) _IEncryptedChat() {}

type EncryptedChatDiscarded struct {
	_              struct{} `tl:"flags,bitflag"`
	HistoryDeleted bool     `tl:",omitempty:flags:0,implicit"`
	ID             int32
}

func (*EncryptedChatDiscarded) CRC() uint32 {
	return 0x1e1c7c45
}
func (*EncryptedChatDiscarded) _IEncryptedChat() {}

type IEncryptedFile interface {
	tl.Object
	_IEncryptedFile()
}

var (
	_ IEncryptedFile = (*EncryptedFileEmpty)(nil)
	_ IEncryptedFile = (*EncryptedFile)(nil)
)

type EncryptedFileEmpty struct{}

func (*EncryptedFileEmpty) CRC() uint32 {
	return 0xc21f497e
}
func (*EncryptedFileEmpty) _IEncryptedFile() {}

type EncryptedFile struct {
	ID             int64
	AccessHash     int64
	Size           int64
	DcID           int32
	KeyFingerprint int32
}

func (*EncryptedFile) CRC() uint32 {
	return 0xa8008cd8
}
func (*EncryptedFile) _IEncryptedFile() {}

type IEncryptedMessage interface {
	tl.Object
	_IEncryptedMessage()
}

var (
	_ IEncryptedMessage = (*EncryptedMessage)(nil)
	_ IEncryptedMessage = (*EncryptedMessageService)(nil)
)

type EncryptedMessage struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
	File     IEncryptedFile
}

func (*EncryptedMessage) CRC() uint32 {
	return 0xed18c118
}
func (*EncryptedMessage) _IEncryptedMessage() {}

type EncryptedMessageService struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
}

func (*EncryptedMessageService) CRC() uint32 {
	return 0x23734b06
}
func (*EncryptedMessageService) _IEncryptedMessage() {}

type IExportedChatInvite interface {
	tl.Object
	_IExportedChatInvite()
}

var (
	_ IExportedChatInvite = (*ChatInviteExported)(nil)
	_ IExportedChatInvite = (*ChatInvitePublicJoinRequests)(nil)
)

type ChatInviteExported struct {
	_             struct{} `tl:"flags,bitflag"`
	Revoked       bool     `tl:",omitempty:flags:0,implicit"`
	Permanent     bool     `tl:",omitempty:flags:5,implicit"`
	RequestNeeded bool     `tl:",omitempty:flags:6,implicit"`
	Link          string
	AdminID       int64
	Date          int32
	StartDate     *int32  `tl:",omitempty:flags:4"`
	ExpireDate    *int32  `tl:",omitempty:flags:1"`
	UsageLimit    *int32  `tl:",omitempty:flags:2"`
	Usage         *int32  `tl:",omitempty:flags:3"`
	Requested     *int32  `tl:",omitempty:flags:7"`
	Title         *string `tl:",omitempty:flags:8"`
}

func (*ChatInviteExported) CRC() uint32 {
	return 0xab4a819
}
func (*ChatInviteExported) _IExportedChatInvite() {}

type ChatInvitePublicJoinRequests struct{}

func (*ChatInvitePublicJoinRequests) CRC() uint32 {
	return 0xed107ab7
}
func (*ChatInvitePublicJoinRequests) _IExportedChatInvite() {}

type IExportedChatlistInvite interface {
	tl.Object
	_IExportedChatlistInvite()
}

var (
	_ IExportedChatlistInvite = (*ExportedChatlistInvite)(nil)
)

type ExportedChatlistInvite struct {
	_     struct{} `tl:"flags,bitflag"`
	Title string
	URL   string
	Peers []IPeer
}

func (*ExportedChatlistInvite) CRC() uint32 {
	return 0xc5181ac
}
func (*ExportedChatlistInvite) _IExportedChatlistInvite() {}

type IExportedContactToken interface {
	tl.Object
	_IExportedContactToken()
}

var (
	_ IExportedContactToken = (*ExportedContactToken)(nil)
)

type ExportedContactToken struct {
	URL     string
	Expires int32
}

func (*ExportedContactToken) CRC() uint32 {
	return 0x41bf109b
}
func (*ExportedContactToken) _IExportedContactToken() {}

type IExportedMessageLink interface {
	tl.Object
	_IExportedMessageLink()
}

var (
	_ IExportedMessageLink = (*ExportedMessageLink)(nil)
)

type ExportedMessageLink struct {
	Link string
	Html string
}

func (*ExportedMessageLink) CRC() uint32 {
	return 0x5dab1af4
}
func (*ExportedMessageLink) _IExportedMessageLink() {}

type IExportedStoryLink interface {
	tl.Object
	_IExportedStoryLink()
}

var (
	_ IExportedStoryLink = (*ExportedStoryLink)(nil)
)

type ExportedStoryLink struct {
	Link string
}

func (*ExportedStoryLink) CRC() uint32 {
	return 0x3fc9053b
}
func (*ExportedStoryLink) _IExportedStoryLink() {}

type IFileHash interface {
	tl.Object
	_IFileHash()
}

var (
	_ IFileHash = (*FileHash)(nil)
)

type FileHash struct {
	Offset int64
	Limit  int32
	Hash   []byte
}

func (*FileHash) CRC() uint32 {
	return 0xf39b035c
}
func (*FileHash) _IFileHash() {}

type IFolder interface {
	tl.Object
	_IFolder()
}

var (
	_ IFolder = (*Folder)(nil)
)

type Folder struct {
	_                         struct{} `tl:"flags,bitflag"`
	AutofillNewBroadcasts     bool     `tl:",omitempty:flags:0,implicit"`
	AutofillPublicGroups      bool     `tl:",omitempty:flags:1,implicit"`
	AutofillNewCorrespondents bool     `tl:",omitempty:flags:2,implicit"`
	ID                        int32
	Title                     string
	Photo                     IChatPhoto `tl:",omitempty:flags:3"`
}

func (*Folder) CRC() uint32 {
	return 0xff544e65
}
func (*Folder) _IFolder() {}

type IFolderPeer interface {
	tl.Object
	_IFolderPeer()
}

var (
	_ IFolderPeer = (*FolderPeer)(nil)
)

type FolderPeer struct {
	Peer     IPeer
	FolderID int32
}

func (*FolderPeer) CRC() uint32 {
	return 0xe9baa668
}
func (*FolderPeer) _IFolderPeer() {}

type IForumTopic interface {
	tl.Object
	_IForumTopic()
}

var (
	_ IForumTopic = (*ForumTopicDeleted)(nil)
	_ IForumTopic = (*ForumTopic)(nil)
)

type ForumTopicDeleted struct {
	ID int32
}

func (*ForumTopicDeleted) CRC() uint32 {
	return 0x23f109b
}
func (*ForumTopicDeleted) _IForumTopic() {}

type ForumTopic struct {
	_                    struct{} `tl:"flags,bitflag"`
	My                   bool     `tl:",omitempty:flags:1,implicit"`
	Closed               bool     `tl:",omitempty:flags:2,implicit"`
	Pinned               bool     `tl:",omitempty:flags:3,implicit"`
	Short                bool     `tl:",omitempty:flags:5,implicit"`
	Hidden               bool     `tl:",omitempty:flags:6,implicit"`
	ID                   int32
	Date                 int32
	Title                string
	IconColor            int32
	IconEmojiID          *int64 `tl:",omitempty:flags:0"`
	TopMessage           int32
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	UnreadMentionsCount  int32
	UnreadReactionsCount int32
	FromID               IPeer
	NotifySettings       IPeerNotifySettings
	Draft                IDraftMessage `tl:",omitempty:flags:4"`
}

func (*ForumTopic) CRC() uint32 {
	return 0x71701da9
}
func (*ForumTopic) _IForumTopic() {}

type IGame interface {
	tl.Object
	_IGame()
}

var (
	_ IGame = (*Game)(nil)
)

type Game struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       IPhoto
	Document    IDocument `tl:",omitempty:flags:0"`
}

func (*Game) CRC() uint32 {
	return 0xbdf9653b
}
func (*Game) _IGame() {}

type IGeoPoint interface {
	tl.Object
	_IGeoPoint()
}

var (
	_ IGeoPoint = (*GeoPointEmpty)(nil)
	_ IGeoPoint = (*GeoPoint)(nil)
)

type GeoPointEmpty struct{}

func (*GeoPointEmpty) CRC() uint32 {
	return 0x1117dd5f
}
func (*GeoPointEmpty) _IGeoPoint() {}

type GeoPoint struct {
	_              struct{} `tl:"flags,bitflag"`
	Long           float64
	Lat            float64
	AccessHash     int64
	AccuracyRadius *int32 `tl:",omitempty:flags:0"`
}

func (*GeoPoint) CRC() uint32 {
	return 0xb2a2f663
}
func (*GeoPoint) _IGeoPoint() {}

type IGlobalPrivacySettings interface {
	tl.Object
	_IGlobalPrivacySettings()
}

var (
	_ IGlobalPrivacySettings = (*GlobalPrivacySettings)(nil)
)

type GlobalPrivacySettings struct {
	_                                struct{} `tl:"flags,bitflag"`
	ArchiveAndMuteNewNoncontactPeers bool     `tl:",omitempty:flags:0,implicit"`
	KeepArchivedUnmuted              bool     `tl:",omitempty:flags:1,implicit"`
	KeepArchivedFolders              bool     `tl:",omitempty:flags:2,implicit"`
}

func (*GlobalPrivacySettings) CRC() uint32 {
	return 0x734c4ccb
}
func (*GlobalPrivacySettings) _IGlobalPrivacySettings() {}

type IGroupCall interface {
	tl.Object
	_IGroupCall()
}

var (
	_ IGroupCall = (*GroupCallDiscarded)(nil)
	_ IGroupCall = (*GroupCall)(nil)
)

type GroupCallDiscarded struct {
	ID         int64
	AccessHash int64
	Duration   int32
}

func (*GroupCallDiscarded) CRC() uint32 {
	return 0x7780bcb4
}
func (*GroupCallDiscarded) _IGroupCall() {}

type GroupCall struct {
	_                       struct{} `tl:"flags,bitflag"`
	JoinMuted               bool     `tl:",omitempty:flags:1,implicit"`
	CanChangeJoinMuted      bool     `tl:",omitempty:flags:2,implicit"`
	JoinDateAsc             bool     `tl:",omitempty:flags:6,implicit"`
	ScheduleStartSubscribed bool     `tl:",omitempty:flags:8,implicit"`
	CanStartVideo           bool     `tl:",omitempty:flags:9,implicit"`
	RecordVideoActive       bool     `tl:",omitempty:flags:11,implicit"`
	RtmpStream              bool     `tl:",omitempty:flags:12,implicit"`
	ListenersHidden         bool     `tl:",omitempty:flags:13,implicit"`
	ID                      int64
	AccessHash              int64
	ParticipantsCount       int32
	Title                   *string `tl:",omitempty:flags:3"`
	StreamDcID              *int32  `tl:",omitempty:flags:4"`
	RecordStartDate         *int32  `tl:",omitempty:flags:5"`
	ScheduleDate            *int32  `tl:",omitempty:flags:7"`
	UnmutedVideoCount       *int32  `tl:",omitempty:flags:10"`
	UnmutedVideoLimit       int32
	Version                 int32
}

func (*GroupCall) CRC() uint32 {
	return 0xd597650c
}
func (*GroupCall) _IGroupCall() {}

type IGroupCallParticipant interface {
	tl.Object
	_IGroupCallParticipant()
}

var (
	_ IGroupCallParticipant = (*GroupCallParticipant)(nil)
)

type GroupCallParticipant struct {
	_               struct{} `tl:"flags,bitflag"`
	Muted           bool     `tl:",omitempty:flags:0,implicit"`
	Left            bool     `tl:",omitempty:flags:1,implicit"`
	CanSelfUnmute   bool     `tl:",omitempty:flags:2,implicit"`
	JustJoined      bool     `tl:",omitempty:flags:4,implicit"`
	Versioned       bool     `tl:",omitempty:flags:5,implicit"`
	Min             bool     `tl:",omitempty:flags:8,implicit"`
	MutedByYou      bool     `tl:",omitempty:flags:9,implicit"`
	VolumeByAdmin   bool     `tl:",omitempty:flags:10,implicit"`
	Self            bool     `tl:",omitempty:flags:12,implicit"`
	VideoJoined     bool     `tl:",omitempty:flags:15,implicit"`
	Peer            IPeer
	Date            int32
	ActiveDate      *int32 `tl:",omitempty:flags:3"`
	Source          int32
	Volume          *int32                     `tl:",omitempty:flags:7"`
	About           *string                    `tl:",omitempty:flags:11"`
	RaiseHandRating *int64                     `tl:",omitempty:flags:13"`
	Video           IGroupCallParticipantVideo `tl:",omitempty:flags:6"`
	Presentation    IGroupCallParticipantVideo `tl:",omitempty:flags:14"`
}

func (*GroupCallParticipant) CRC() uint32 {
	return 0xeba636fe
}
func (*GroupCallParticipant) _IGroupCallParticipant() {}

type IGroupCallParticipantVideo interface {
	tl.Object
	_IGroupCallParticipantVideo()
}

var (
	_ IGroupCallParticipantVideo = (*GroupCallParticipantVideo)(nil)
)

type GroupCallParticipantVideo struct {
	_            struct{} `tl:"flags,bitflag"`
	Paused       bool     `tl:",omitempty:flags:0,implicit"`
	Endpoint     string
	SourceGroups []IGroupCallParticipantVideoSourceGroup
	AudioSource  *int32 `tl:",omitempty:flags:1"`
}

func (*GroupCallParticipantVideo) CRC() uint32 {
	return 0x67753ac8
}
func (*GroupCallParticipantVideo) _IGroupCallParticipantVideo() {}

type IGroupCallParticipantVideoSourceGroup interface {
	tl.Object
	_IGroupCallParticipantVideoSourceGroup()
}

var (
	_ IGroupCallParticipantVideoSourceGroup = (*GroupCallParticipantVideoSourceGroup)(nil)
)

type GroupCallParticipantVideoSourceGroup struct {
	Semantics string
	Sources   []int32
}

func (*GroupCallParticipantVideoSourceGroup) CRC() uint32 {
	return 0xdcb118b7
}
func (*GroupCallParticipantVideoSourceGroup) _IGroupCallParticipantVideoSourceGroup() {}

type IGroupCallStreamChannel interface {
	tl.Object
	_IGroupCallStreamChannel()
}

var (
	_ IGroupCallStreamChannel = (*GroupCallStreamChannel)(nil)
)

type GroupCallStreamChannel struct {
	Channel         int32
	Scale           int32
	LastTimestampMs int64
}

func (*GroupCallStreamChannel) CRC() uint32 {
	return 0x80eb48af
}
func (*GroupCallStreamChannel) _IGroupCallStreamChannel() {}

type IHighScore interface {
	tl.Object
	_IHighScore()
}

var (
	_ IHighScore = (*HighScore)(nil)
)

type HighScore struct {
	Pos    int32
	UserID int64
	Score  int32
}

func (*HighScore) CRC() uint32 {
	return 0x73a379eb
}
func (*HighScore) _IHighScore() {}

type IImportedContact interface {
	tl.Object
	_IImportedContact()
}

var (
	_ IImportedContact = (*ImportedContact)(nil)
)

type ImportedContact struct {
	UserID   int64
	ClientID int64
}

func (*ImportedContact) CRC() uint32 {
	return 0xc13e3c50
}
func (*ImportedContact) _IImportedContact() {}

type IInlineBotSwitchPm interface {
	tl.Object
	_IInlineBotSwitchPm()
}

var (
	_ IInlineBotSwitchPm = (*InlineBotSwitchPm)(nil)
)

type InlineBotSwitchPm struct {
	Text       string
	StartParam string
}

func (*InlineBotSwitchPm) CRC() uint32 {
	return 0x3c20629f
}
func (*InlineBotSwitchPm) _IInlineBotSwitchPm() {}

type IInlineBotWebView interface {
	tl.Object
	_IInlineBotWebView()
}

var (
	_ IInlineBotWebView = (*InlineBotWebView)(nil)
)

type InlineBotWebView struct {
	Text string
	URL  string
}

func (*InlineBotWebView) CRC() uint32 {
	return 0xb57295d5
}
func (*InlineBotWebView) _IInlineBotWebView() {}

type IInputAppEvent interface {
	tl.Object
	_IInputAppEvent()
}

var (
	_ IInputAppEvent = (*InputAppEvent)(nil)
)

type InputAppEvent struct {
	Time float64
	Type string
	Peer int64
	Data IJSONValue
}

func (*InputAppEvent) CRC() uint32 {
	return 0x1d1b1245
}
func (*InputAppEvent) _IInputAppEvent() {}

type IInputBotApp interface {
	tl.Object
	_IInputBotApp()
}

var (
	_ IInputBotApp = (*InputBotAppID)(nil)
	_ IInputBotApp = (*InputBotAppShortName)(nil)
)

type InputBotAppID struct {
	ID         int64
	AccessHash int64
}

func (*InputBotAppID) CRC() uint32 {
	return 0xa920bd7a
}
func (*InputBotAppID) _IInputBotApp() {}

type InputBotAppShortName struct {
	BotID     IInputUser
	ShortName string
}

func (*InputBotAppShortName) CRC() uint32 {
	return 0x908c0407
}
func (*InputBotAppShortName) _IInputBotApp() {}

type IInputBotInlineMessage interface {
	tl.Object
	_IInputBotInlineMessage()
}

var (
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaAuto)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageText)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaGeo)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaVenue)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaContact)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageGame)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaInvoice)(nil)
	_ IInputBotInlineMessage = (*InputBotInlineMessageMediaWebPage)(nil)
)

type InputBotInlineMessageMediaAuto struct {
	_           struct{} `tl:"flags,bitflag"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup IReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaAuto) CRC() uint32 {
	return 0x3380c786
}
func (*InputBotInlineMessageMediaAuto) _IInputBotInlineMessage() {}

type InputBotInlineMessageText struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup IReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageText) CRC() uint32 {
	return 0x3dcd7a87
}
func (*InputBotInlineMessageText) _IInputBotInlineMessage() {}

type InputBotInlineMessageMediaGeo struct {
	_                           struct{} `tl:"flags,bitflag"`
	GeoPoint                    IInputGeoPoint
	Heading                     *int32       `tl:",omitempty:flags:0"`
	Period                      *int32       `tl:",omitempty:flags:1"`
	ProximityNotificationRadius *int32       `tl:",omitempty:flags:3"`
	ReplyMarkup                 IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaGeo) CRC() uint32 {
	return 0x96929a85
}
func (*InputBotInlineMessageMediaGeo) _IInputBotInlineMessage() {}

type InputBotInlineMessageMediaVenue struct {
	_           struct{} `tl:"flags,bitflag"`
	GeoPoint    IInputGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaVenue) CRC() uint32 {
	return 0x417bbf11
}
func (*InputBotInlineMessageMediaVenue) _IInputBotInlineMessage() {}

type InputBotInlineMessageMediaContact struct {
	_           struct{} `tl:"flags,bitflag"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaContact) CRC() uint32 {
	return 0xa6edbffd
}
func (*InputBotInlineMessageMediaContact) _IInputBotInlineMessage() {}

type InputBotInlineMessageGame struct {
	_           struct{}     `tl:"flags,bitflag"`
	ReplyMarkup IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageGame) CRC() uint32 {
	return 0x4b425864
}
func (*InputBotInlineMessageGame) _IInputBotInlineMessage() {}

type InputBotInlineMessageMediaInvoice struct {
	_            struct{} `tl:"flags,bitflag"`
	Title        string
	Description  string
	Photo        IInputWebDocument `tl:",omitempty:flags:0"`
	Invoice      IInvoice
	Payload      []byte
	Provider     string
	ProviderData IDataJSON
	ReplyMarkup  IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaInvoice) CRC() uint32 {
	return 0xd7e78225
}
func (*InputBotInlineMessageMediaInvoice) _IInputBotInlineMessage() {}

type InputBotInlineMessageMediaWebPage struct {
	_               struct{} `tl:"flags,bitflag"`
	InvertMedia     bool     `tl:",omitempty:flags:3,implicit"`
	ForceLargeMedia bool     `tl:",omitempty:flags:4,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:5,implicit"`
	Optional        bool     `tl:",omitempty:flags:6,implicit"`
	Message         string
	Entities        []IMessageEntity `tl:",omitempty:flags:1"`
	URL             string
	ReplyMarkup     IReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaWebPage) CRC() uint32 {
	return 0xbddcc510
}
func (*InputBotInlineMessageMediaWebPage) _IInputBotInlineMessage() {}

type IInputBotInlineMessageID interface {
	tl.Object
	_IInputBotInlineMessageID()
}

var (
	_ IInputBotInlineMessageID = (*InputBotInlineMessageID)(nil)
	_ IInputBotInlineMessageID = (*InputBotInlineMessageID64)(nil)
)

type InputBotInlineMessageID struct {
	DcID       int32
	ID         int64
	AccessHash int64
}

func (*InputBotInlineMessageID) CRC() uint32 {
	return 0x890c3d89
}
func (*InputBotInlineMessageID) _IInputBotInlineMessageID() {}

type InputBotInlineMessageID64 struct {
	DcID       int32
	OwnerID    int64
	ID         int32
	AccessHash int64
}

func (*InputBotInlineMessageID64) CRC() uint32 {
	return 0xb6d915d7
}
func (*InputBotInlineMessageID64) _IInputBotInlineMessageID() {}

type IInputBotInlineResult interface {
	tl.Object
	_IInputBotInlineResult()
}

var (
	_ IInputBotInlineResult = (*InputBotInlineResult)(nil)
	_ IInputBotInlineResult = (*InputBotInlineResultPhoto)(nil)
	_ IInputBotInlineResult = (*InputBotInlineResultDocument)(nil)
	_ IInputBotInlineResult = (*InputBotInlineResultGame)(nil)
)

type InputBotInlineResult struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string           `tl:",omitempty:flags:1"`
	Description *string           `tl:",omitempty:flags:2"`
	URL         *string           `tl:",omitempty:flags:3"`
	Thumb       IInputWebDocument `tl:",omitempty:flags:4"`
	Content     IInputWebDocument `tl:",omitempty:flags:5"`
	SendMessage IInputBotInlineMessage
}

func (*InputBotInlineResult) CRC() uint32 {
	return 0x88bf9319
}
func (*InputBotInlineResult) _IInputBotInlineResult() {}

type InputBotInlineResultPhoto struct {
	ID          string
	Type        string
	Photo       IInputPhoto
	SendMessage IInputBotInlineMessage
}

func (*InputBotInlineResultPhoto) CRC() uint32 {
	return 0xa8d864a7
}
func (*InputBotInlineResultPhoto) _IInputBotInlineResult() {}

type InputBotInlineResultDocument struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string `tl:",omitempty:flags:1"`
	Description *string `tl:",omitempty:flags:2"`
	Document    IInputDocument
	SendMessage IInputBotInlineMessage
}

func (*InputBotInlineResultDocument) CRC() uint32 {
	return 0xfff8fdc4
}
func (*InputBotInlineResultDocument) _IInputBotInlineResult() {}

type InputBotInlineResultGame struct {
	ID          string
	ShortName   string
	SendMessage IInputBotInlineMessage
}

func (*InputBotInlineResultGame) CRC() uint32 {
	return 0x4fa417f2
}
func (*InputBotInlineResultGame) _IInputBotInlineResult() {}

type IInputChannel interface {
	tl.Object
	_IInputChannel()
}

var (
	_ IInputChannel = (*InputChannelEmpty)(nil)
	_ IInputChannel = (*InputChannel)(nil)
	_ IInputChannel = (*InputChannelFromMessage)(nil)
)

type InputChannelEmpty struct{}

func (*InputChannelEmpty) CRC() uint32 {
	return 0xee8c1e86
}
func (*InputChannelEmpty) _IInputChannel() {}

type InputChannel struct {
	ChannelID  int64
	AccessHash int64
}

func (*InputChannel) CRC() uint32 {
	return 0xf35aec28
}
func (*InputChannel) _IInputChannel() {}

type InputChannelFromMessage struct {
	Peer      IInputPeer
	MsgID     int32
	ChannelID int64
}

func (*InputChannelFromMessage) CRC() uint32 {
	return 0x5b934f9d
}
func (*InputChannelFromMessage) _IInputChannel() {}

// Defines a new group profile photo.
type IInputChatPhoto interface {
	tl.Object
	_IInputChatPhoto()
}

var (
	_ IInputChatPhoto = (*InputChatPhotoEmpty)(nil)
	_ IInputChatPhoto = (*InputChatUploadedPhoto)(nil)
	_ IInputChatPhoto = (*InputChatPhoto)(nil)
)

// Empty constructor, remove group photo.
type InputChatPhotoEmpty struct{}

func (*InputChatPhotoEmpty) CRC() uint32 {
	return 0x1ca48f57
}
func (*InputChatPhotoEmpty) _IInputChatPhoto() {}

// New photo to be set as group profile photo.
type InputChatUploadedPhoto struct {
	_                struct{}   `tl:"flags,bitflag"`
	File             IInputFile `tl:",omitempty:flags:0"` // File saved in parts using the method upload.saveFilePart
	Video            IInputFile `tl:",omitempty:flags:1"` // Square video for animated profile picture"
	VideoStartTs     *float64   `tl:",omitempty:flags:2"` // Timestamp that should be shown as static preview to the user (seconds)
	VideoEmojiMarkup IVideoSize `tl:",omitempty:flags:3"`
}

func (*InputChatUploadedPhoto) CRC() uint32 {
	return 0xbdcdaec0
}
func (*InputChatUploadedPhoto) _IInputChatPhoto() {}

// Existing photo to be set as a chat profile photo.
type InputChatPhoto struct {
	ID IInputPhoto // Existing photo
}

func (*InputChatPhoto) CRC() uint32 {
	return 0x8953ad37
}
func (*InputChatPhoto) _IInputChatPhoto() {}

type IInputChatlist interface {
	tl.Object
	_IInputChatlist()
}

var (
	_ IInputChatlist = (*InputChatlistDialogFilter)(nil)
)

type InputChatlistDialogFilter struct {
	FilterID int32
}

func (*InputChatlistDialogFilter) CRC() uint32 {
	return 0xf3e0da33
}
func (*InputChatlistDialogFilter) _IInputChatlist() {}

type IInputCheckPasswordSRP interface {
	tl.Object
	_IInputCheckPasswordSRP()
}

var (
	_ IInputCheckPasswordSRP = (*InputCheckPasswordEmpty)(nil)
	_ IInputCheckPasswordSRP = (*InputCheckPasswordSRP)(nil)
)

type InputCheckPasswordEmpty struct{}

func (*InputCheckPasswordEmpty) CRC() uint32 {
	return 0x9880f658
}
func (*InputCheckPasswordEmpty) _IInputCheckPasswordSRP() {}

type InputCheckPasswordSRP struct {
	SRPID int64
	A     []byte
	M1    []byte
}

func (*InputCheckPasswordSRP) CRC() uint32 {
	return 0xd27ff082
}
func (*InputCheckPasswordSRP) _IInputCheckPasswordSRP() {}

type IInputClientProxy interface {
	tl.Object
	_IInputClientProxy()
}

var (
	_ IInputClientProxy = (*InputClientProxy)(nil)
)

type InputClientProxy struct {
	Address string
	Port    int32
}

func (*InputClientProxy) CRC() uint32 {
	return 0x75588b3f
}
func (*InputClientProxy) _IInputClientProxy() {}

// Object defines a contact from the user's phonebook.
type IInputContact interface {
	tl.Object
	_IInputContact()
}

var (
	_ IInputContact = (*InputPhoneContact)(nil)
)

// Phone contact. The client_id is just an arbitrary contact ID: it should be set, for example, to an incremental number when using contacts.importContacts, in order to retry importing only the contacts that weren't imported successfully.
type InputPhoneContact struct {
	ClientID  int64  // User identifier on the client
	Phone     string // Phone number
	FirstName string // Contact's first name
	LastName  string // Contact's last name
}

func (*InputPhoneContact) CRC() uint32 {
	return 0xf392b7f4
}
func (*InputPhoneContact) _IInputContact() {}

type IInputDialogPeer interface {
	tl.Object
	_IInputDialogPeer()
}

var (
	_ IInputDialogPeer = (*InputDialogPeer)(nil)
	_ IInputDialogPeer = (*InputDialogPeerFolder)(nil)
)

type InputDialogPeer struct {
	Peer IInputPeer
}

func (*InputDialogPeer) CRC() uint32 {
	return 0xfcaafeb7
}
func (*InputDialogPeer) _IInputDialogPeer() {}

type InputDialogPeerFolder struct {
	FolderID int32
}

func (*InputDialogPeerFolder) CRC() uint32 {
	return 0x64600527
}
func (*InputDialogPeerFolder) _IInputDialogPeer() {}

type IInputDocument interface {
	tl.Object
	_IInputDocument()
}

var (
	_ IInputDocument = (*InputDocumentEmpty)(nil)
	_ IInputDocument = (*InputDocument)(nil)
)

type InputDocumentEmpty struct{}

func (*InputDocumentEmpty) CRC() uint32 {
	return 0x72f0eaae
}
func (*InputDocumentEmpty) _IInputDocument() {}

type InputDocument struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

func (*InputDocument) CRC() uint32 {
	return 0x1abfb575
}
func (*InputDocument) _IInputDocument() {}

type IInputEncryptedChat interface {
	tl.Object
	_IInputEncryptedChat()
}

var (
	_ IInputEncryptedChat = (*InputEncryptedChat)(nil)
)

type InputEncryptedChat struct {
	ChatID     int32
	AccessHash int64
}

func (*InputEncryptedChat) CRC() uint32 {
	return 0xf141b5e1
}
func (*InputEncryptedChat) _IInputEncryptedChat() {}

type IInputEncryptedFile interface {
	tl.Object
	_IInputEncryptedFile()
}

var (
	_ IInputEncryptedFile = (*InputEncryptedFileEmpty)(nil)
	_ IInputEncryptedFile = (*InputEncryptedFileUploaded)(nil)
	_ IInputEncryptedFile = (*InputEncryptedFile)(nil)
	_ IInputEncryptedFile = (*InputEncryptedFileBigUploaded)(nil)
)

type InputEncryptedFileEmpty struct{}

func (*InputEncryptedFileEmpty) CRC() uint32 {
	return 0x1837c364
}
func (*InputEncryptedFileEmpty) _IInputEncryptedFile() {}

type InputEncryptedFileUploaded struct {
	ID             int64
	Parts          int32
	Md5Checksum    string
	KeyFingerprint int32
}

func (*InputEncryptedFileUploaded) CRC() uint32 {
	return 0x64bd0306
}
func (*InputEncryptedFileUploaded) _IInputEncryptedFile() {}

type InputEncryptedFile struct {
	ID         int64
	AccessHash int64
}

func (*InputEncryptedFile) CRC() uint32 {
	return 0x5a17b5e5
}
func (*InputEncryptedFile) _IInputEncryptedFile() {}

type InputEncryptedFileBigUploaded struct {
	ID             int64
	Parts          int32
	KeyFingerprint int32
}

func (*InputEncryptedFileBigUploaded) CRC() uint32 {
	return 0x2dc173c8
}
func (*InputEncryptedFileBigUploaded) _IInputEncryptedFile() {}

// Defines a file uploaded by the client.
type IInputFile interface {
	tl.Object
	_IInputFile()
}

var (
	_ IInputFile = (*InputFile)(nil)
	_ IInputFile = (*InputFileBig)(nil)
)

// Defines a file saved in parts using the method upload.saveFilePart.
type InputFile struct {
	ID          int64  // Random file identifier created by the client
	Parts       int32  // Number of parts saved
	Name        string // Full name of the file
	Md5Checksum string // In case the file's md5-hash was passed, contents of the file will be checked prior to use
}

func (*InputFile) CRC() uint32 {
	return 0xf52ff27f
}
func (*InputFile) _IInputFile() {}

// Assigns a big file (over 10Mb in size), saved in part using the method upload.saveBigFilePart.
type InputFileBig struct {
	ID    int64  // Random file id, created by the client
	Parts int32  // Number of parts saved
	Name  string // Full file name
}

func (*InputFileBig) CRC() uint32 {
	return 0xfa4f0bb5
}
func (*InputFileBig) _IInputFile() {}

type IInputFileLocation interface {
	tl.Object
	_IInputFileLocation()
}

var (
	_ IInputFileLocation = (*InputFileLocation)(nil)
	_ IInputFileLocation = (*InputEncryptedFileLocation)(nil)
	_ IInputFileLocation = (*InputDocumentFileLocation)(nil)
	_ IInputFileLocation = (*InputSecureFileLocation)(nil)
	_ IInputFileLocation = (*InputTakeoutFileLocation)(nil)
	_ IInputFileLocation = (*InputPhotoFileLocation)(nil)
	_ IInputFileLocation = (*InputPhotoLegacyFileLocation)(nil)
	_ IInputFileLocation = (*InputPeerPhotoFileLocation)(nil)
	_ IInputFileLocation = (*InputStickerSetThumb)(nil)
	_ IInputFileLocation = (*InputGroupCallStream)(nil)
)

type InputFileLocation struct {
	VolumeID      int64
	LocalID       int32
	Secret        int64
	FileReference []byte
}

func (*InputFileLocation) CRC() uint32 {
	return 0xdfdaabe1
}
func (*InputFileLocation) _IInputFileLocation() {}

type InputEncryptedFileLocation struct {
	ID         int64
	AccessHash int64
}

func (*InputEncryptedFileLocation) CRC() uint32 {
	return 0xf5235d55
}
func (*InputEncryptedFileLocation) _IInputFileLocation() {}

type InputDocumentFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

func (*InputDocumentFileLocation) CRC() uint32 {
	return 0xbad07584
}
func (*InputDocumentFileLocation) _IInputFileLocation() {}

type InputSecureFileLocation struct {
	ID         int64
	AccessHash int64
}

func (*InputSecureFileLocation) CRC() uint32 {
	return 0xcbc7ee28
}
func (*InputSecureFileLocation) _IInputFileLocation() {}

type InputTakeoutFileLocation struct{}

func (*InputTakeoutFileLocation) CRC() uint32 {
	return 0x29be5899
}
func (*InputTakeoutFileLocation) _IInputFileLocation() {}

type InputPhotoFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

func (*InputPhotoFileLocation) CRC() uint32 {
	return 0x40181ffe
}
func (*InputPhotoFileLocation) _IInputFileLocation() {}

type InputPhotoLegacyFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	VolumeID      int64
	LocalID       int32
	Secret        int64
}

func (*InputPhotoLegacyFileLocation) CRC() uint32 {
	return 0xd83466f3
}
func (*InputPhotoLegacyFileLocation) _IInputFileLocation() {}

type InputPeerPhotoFileLocation struct {
	_       struct{} `tl:"flags,bitflag"`
	Big     bool     `tl:",omitempty:flags:0,implicit"`
	Peer    IInputPeer
	PhotoID int64
}

func (*InputPeerPhotoFileLocation) CRC() uint32 {
	return 0x37257e99
}
func (*InputPeerPhotoFileLocation) _IInputFileLocation() {}

type InputStickerSetThumb struct {
	Stickerset   IInputStickerSet
	ThumbVersion int32
}

func (*InputStickerSetThumb) CRC() uint32 {
	return 0x9d84f3db
}
func (*InputStickerSetThumb) _IInputFileLocation() {}

type InputGroupCallStream struct {
	_            struct{} `tl:"flags,bitflag"`
	Call         IInputGroupCall
	TimeMs       int64
	Scale        int32
	VideoChannel *int32 `tl:",omitempty:flags:0"`
	VideoQuality *int32 `tl:",omitempty:flags:0"`
}

func (*InputGroupCallStream) CRC() uint32 {
	return 0x598a92a
}
func (*InputGroupCallStream) _IInputFileLocation() {}

type IInputFolderPeer interface {
	tl.Object
	_IInputFolderPeer()
}

var (
	_ IInputFolderPeer = (*InputFolderPeer)(nil)
)

type InputFolderPeer struct {
	Peer     IInputPeer
	FolderID int32
}

func (*InputFolderPeer) CRC() uint32 {
	return 0xfbd2c296
}
func (*InputFolderPeer) _IInputFolderPeer() {}

type IInputGame interface {
	tl.Object
	_IInputGame()
}

var (
	_ IInputGame = (*InputGameID)(nil)
	_ IInputGame = (*InputGameShortName)(nil)
)

type InputGameID struct {
	ID         int64
	AccessHash int64
}

func (*InputGameID) CRC() uint32 {
	return 0x32c3e77
}
func (*InputGameID) _IInputGame() {}

type InputGameShortName struct {
	BotID     IInputUser
	ShortName string
}

func (*InputGameShortName) CRC() uint32 {
	return 0xc331e80a
}
func (*InputGameShortName) _IInputGame() {}

// Defines a GeoPoint.
type IInputGeoPoint interface {
	tl.Object
	_IInputGeoPoint()
}

var (
	_ IInputGeoPoint = (*InputGeoPointEmpty)(nil)
	_ IInputGeoPoint = (*InputGeoPoint)(nil)
)

// Empty GeoPoint constructor.
type InputGeoPointEmpty struct{}

func (*InputGeoPointEmpty) CRC() uint32 {
	return 0xe4c123d6
}
func (*InputGeoPointEmpty) _IInputGeoPoint() {}

// Defines a GeoPoint by its coordinates.
type InputGeoPoint struct {
	_              struct{} `tl:"flags,bitflag"`
	Lat            float64  // Latitide
	Long           float64  // Longtitude
	AccuracyRadius *int32   `tl:",omitempty:flags:0"` // The estimated horizontal accuracy of the location, in meters; as defined by the sender.
}

func (*InputGeoPoint) CRC() uint32 {
	return 0x48222faf
}
func (*InputGeoPoint) _IInputGeoPoint() {}

type IInputGroupCall interface {
	tl.Object
	_IInputGroupCall()
}

var (
	_ IInputGroupCall = (*InputGroupCall)(nil)
)

type InputGroupCall struct {
	ID         int64
	AccessHash int64
}

func (*InputGroupCall) CRC() uint32 {
	return 0xd8aa840f
}
func (*InputGroupCall) _IInputGroupCall() {}

type IInputInvoice interface {
	tl.Object
	_IInputInvoice()
}

var (
	_ IInputInvoice = (*InputInvoiceMessage)(nil)
	_ IInputInvoice = (*InputInvoiceSlug)(nil)
	_ IInputInvoice = (*InputInvoicePremiumGiftCode)(nil)
)

type InputInvoiceMessage struct {
	Peer  IInputPeer
	MsgID int32
}

func (*InputInvoiceMessage) CRC() uint32 {
	return 0xc5b56859
}
func (*InputInvoiceMessage) _IInputInvoice() {}

type InputInvoiceSlug struct {
	Slug string
}

func (*InputInvoiceSlug) CRC() uint32 {
	return 0xc326caef
}
func (*InputInvoiceSlug) _IInputInvoice() {}

type InputInvoicePremiumGiftCode struct {
	Purpose IInputStorePaymentPurpose
	Option  IPremiumGiftCodeOption
}

func (*InputInvoicePremiumGiftCode) CRC() uint32 {
	return 0x98986c0d
}
func (*InputInvoicePremiumGiftCode) _IInputInvoice() {}

// Defines media content of a message.
type IInputMedia interface {
	tl.Object
	_IInputMedia()
}

var (
	_ IInputMedia = (*InputMediaEmpty)(nil)
	_ IInputMedia = (*InputMediaUploadedPhoto)(nil)
	_ IInputMedia = (*InputMediaPhoto)(nil)
	_ IInputMedia = (*InputMediaGeoPoint)(nil)
	_ IInputMedia = (*InputMediaContact)(nil)
	_ IInputMedia = (*InputMediaUploadedDocument)(nil)
	_ IInputMedia = (*InputMediaDocument)(nil)
	_ IInputMedia = (*InputMediaVenue)(nil)
	_ IInputMedia = (*InputMediaPhotoExternal)(nil)
	_ IInputMedia = (*InputMediaDocumentExternal)(nil)
	_ IInputMedia = (*InputMediaGame)(nil)
	_ IInputMedia = (*InputMediaInvoice)(nil)
	_ IInputMedia = (*InputMediaGeoLive)(nil)
	_ IInputMedia = (*InputMediaPoll)(nil)
	_ IInputMedia = (*InputMediaDice)(nil)
	_ IInputMedia = (*InputMediaStory)(nil)
	_ IInputMedia = (*InputMediaWebPage)(nil)
)

// Empty media content of a message.
type InputMediaEmpty struct{}

func (*InputMediaEmpty) CRC() uint32 {
	return 0x9664f57f
}
func (*InputMediaEmpty) _IInputMedia() {}

// Photo
type InputMediaUploadedPhoto struct {
	_          struct{}         `tl:"flags,bitflag"`
	Spoiler    bool             `tl:",omitempty:flags:2,implicit"`
	File       IInputFile       // The uploaded file
	Stickers   []IInputDocument `tl:",omitempty:flags:0"` // Attached mask stickers
	TTLSeconds *int32           `tl:",omitempty:flags:1"` // Time to live in seconds of self-destructing photo
}

func (*InputMediaUploadedPhoto) CRC() uint32 {
	return 0x1e287d04
}
func (*InputMediaUploadedPhoto) _IInputMedia() {}

type InputMediaPhoto struct {
	_          struct{}    `tl:"flags,bitflag"`
	Spoiler    bool        `tl:",omitempty:flags:1,implicit"`
	ID         IInputPhoto // Photo to be forwarded
	TTLSeconds *int32      `tl:",omitempty:flags:0"` // Time to live in seconds of self-destructing photo
}

func (*InputMediaPhoto) CRC() uint32 {
	return 0xb3ba0635
}
func (*InputMediaPhoto) _IInputMedia() {}

// Map.
type InputMediaGeoPoint struct {
	GeoPoint IInputGeoPoint // GeoPoint
}

func (*InputMediaGeoPoint) CRC() uint32 {
	return 0xf9c44144
}
func (*InputMediaGeoPoint) _IInputMedia() {}

// Phonebook contact
type InputMediaContact struct {
	PhoneNumber string // Number of parts saved
	FirstName   string // Contact's first name
	LastName    string // Contact's last name
	Vcard       string // Contact vcard
}

func (*InputMediaContact) CRC() uint32 {
	return 0xf8ab7dfb
}
func (*InputMediaContact) _IInputMedia() {}

// New document
type InputMediaUploadedDocument struct {
	_            struct{}             `tl:"flags,bitflag"`
	NosoundVideo bool                 `tl:",omitempty:flags:3,implicit"` // Whether the specified document is a video file with no audio tracks (a GIF animation (even as MPEG4), for example)
	ForceFile    bool                 `tl:",omitempty:flags:4,implicit"` // Force the media file to be uploaded as document
	Spoiler      bool                 `tl:",omitempty:flags:5,implicit"`
	File         IInputFile           // The uploaded file
	Thumb        IInputFile           `tl:",omitempty:flags:2"` // Thumbnail of the document, uploaded as for the file
	MimeType     string               // MIME type of document
	Attributes   []IDocumentAttribute // Attributes that specify the type of the document (video, audio, voice, sticker, etc.)
	Stickers     []IInputDocument     `tl:",omitempty:flags:0"` // Attached stickers
	TTLSeconds   *int32               `tl:",omitempty:flags:1"` // Time to live in seconds of self-destructing document
}

func (*InputMediaUploadedDocument) CRC() uint32 {
	return 0x5b38c6c1
}
func (*InputMediaUploadedDocument) _IInputMedia() {}

// Forwarded document
type InputMediaDocument struct {
	_          struct{}       `tl:"flags,bitflag"`
	Spoiler    bool           `tl:",omitempty:flags:2,implicit"`
	ID         IInputDocument // The document to be forwarded.
	TTLSeconds *int32         `tl:",omitempty:flags:0"` // Time to live of self-destructing document
	Query      *string        `tl:",omitempty:flags:1"`
}

func (*InputMediaDocument) CRC() uint32 {
	return 0x33473058
}
func (*InputMediaDocument) _IInputMedia() {}

// Can be used to send a venue geolocation.
type InputMediaVenue struct {
	GeoPoint  IInputGeoPoint // Geolocation
	Title     string         // Venue name
	Address   string         // Physical address of the venue
	Provider  string         // Venue provider: currently only "foursquare" needs to be supported
	VenueID   string         // Venue ID in the provider's database
	VenueType string         // Venue type in the provider's database
}

func (*InputMediaVenue) CRC() uint32 {
	return 0xc13d1c11
}
func (*InputMediaVenue) _IInputMedia() {}

// New photo that will be uploaded by the server using the specified URL
type InputMediaPhotoExternal struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:1,implicit"`
	URL        string   // URL of the photo"
	TTLSeconds *int32   `tl:",omitempty:flags:0"` // Self-destruct time to live of photo"
}

func (*InputMediaPhotoExternal) CRC() uint32 {
	return 0xe5bbfe1a
}
func (*InputMediaPhotoExternal) _IInputMedia() {}

// Document that will be downloaded by the telegram servers
type InputMediaDocumentExternal struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:1,implicit"`
	URL        string   // URL of the document
	TTLSeconds *int32   `tl:",omitempty:flags:0"` // Self-destruct time to live of document
}

func (*InputMediaDocumentExternal) CRC() uint32 {
	return 0xfb52dc99
}
func (*InputMediaDocumentExternal) _IInputMedia() {}

// A game
type InputMediaGame struct {
	ID IInputGame // The game to forward
}

func (*InputMediaGame) CRC() uint32 {
	return 0xd33f43f3
}
func (*InputMediaGame) _IInputMedia() {}

// Generated invoice of a bot payment
type InputMediaInvoice struct {
	_             struct{}          `tl:"flags,bitflag"`
	Title         string            // roduct name, 1-32 characters
	Description   string            // Product description, 1-255 characters
	Photo         IInputWebDocument `tl:",omitempty:flags:0"` // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	Invoice       IInvoice          // The actual invoice
	Payload       []byte            // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	Provider      string            // Payments provider token, obtained via Botfather
	ProviderData  IDataJSON         // JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	StartParam    *string           `tl:",omitempty:flags:1"` // Start parameter
	ExtendedMedia IInputMedia       `tl:",omitempty:flags:2"`
}

func (*InputMediaInvoice) CRC() uint32 {
	return 0x8eb5a6d5
}
func (*InputMediaInvoice) _IInputMedia() {}

// Live geolocation
type InputMediaGeoLive struct {
	_                           struct{}       `tl:"flags,bitflag"`
	Stopped                     bool           `tl:",omitempty:flags:0,implicit"` // Whether sending of the geolocation was stopped
	GeoPoint                    IInputGeoPoint // Current geolocation
	Heading                     *int32         `tl:",omitempty:flags:2"` // For live locations, a direction in which the location moves, in degrees; 1-360.
	Period                      *int32         `tl:",omitempty:flags:1"` // Validity period of the current location
	ProximityNotificationRadius *int32         `tl:",omitempty:flags:3"` // For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000)
}

func (*InputMediaGeoLive) CRC() uint32 {
	return 0x971fa843
}
func (*InputMediaGeoLive) _IInputMedia() {}

// A poll
type InputMediaPoll struct {
	_                struct{}         `tl:"flags,bitflag"`
	Poll             IPoll            // The poll to send
	CorrectAnswers   [][]byte         `tl:",omitempty:flags:0"` // Correct answer IDs (for quiz polls)
	Solution         *string          `tl:",omitempty:flags:1"` // Explanation of quiz solution
	SolutionEntities []IMessageEntity `tl:",omitempty:flags:1"` // Message entities for styled text
}

func (*InputMediaPoll) CRC() uint32 {
	return 0xf94e5f1
}
func (*InputMediaPoll) _IInputMedia() {}

// Assigns a big file (over 10Mb in size), saved in part using the method upload.saveBigFilePart.
type InputMediaDice struct {
	Emoticon string // The emoji, for now 🏀, 🎲 and 🎯 are supported
}

func (*InputMediaDice) CRC() uint32 {
	return 0xe66fbf7b
}
func (*InputMediaDice) _IInputMedia() {}

type InputMediaStory struct {
	Peer IInputPeer
	ID   int32
}

func (*InputMediaStory) CRC() uint32 {
	return 0x89fdd778
}
func (*InputMediaStory) _IInputMedia() {}

type InputMediaWebPage struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:",omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:1,implicit"`
	Optional        bool     `tl:",omitempty:flags:2,implicit"`
	URL             string
}

func (*InputMediaWebPage) CRC() uint32 {
	return 0xc21b8849
}
func (*InputMediaWebPage) _IInputMedia() {}

type IInputMessage interface {
	tl.Object
	_IInputMessage()
}

var (
	_ IInputMessage = (*InputMessageID)(nil)
	_ IInputMessage = (*InputMessageReplyTo)(nil)
	_ IInputMessage = (*InputMessagePinned)(nil)
	_ IInputMessage = (*InputMessageCallbackQuery)(nil)
)

type InputMessageID struct {
	ID int32
}

func (*InputMessageID) CRC() uint32 {
	return 0xa676a322
}
func (*InputMessageID) _IInputMessage() {}

type InputMessageReplyTo struct {
	ID int32
}

func (*InputMessageReplyTo) CRC() uint32 {
	return 0xbad88395
}
func (*InputMessageReplyTo) _IInputMessage() {}

type InputMessagePinned struct{}

func (*InputMessagePinned) CRC() uint32 {
	return 0x86872538
}
func (*InputMessagePinned) _IInputMessage() {}

type InputMessageCallbackQuery struct {
	ID      int32
	QueryID int64
}

func (*InputMessageCallbackQuery) CRC() uint32 {
	return 0xacfa1a7e
}
func (*InputMessageCallbackQuery) _IInputMessage() {}

type IInputNotifyPeer interface {
	tl.Object
	_IInputNotifyPeer()
}

var (
	_ IInputNotifyPeer = (*InputNotifyPeer)(nil)
	_ IInputNotifyPeer = (*InputNotifyUsers)(nil)
	_ IInputNotifyPeer = (*InputNotifyChats)(nil)
	_ IInputNotifyPeer = (*InputNotifyBroadcasts)(nil)
	_ IInputNotifyPeer = (*InputNotifyForumTopic)(nil)
)

type InputNotifyPeer struct {
	Peer IInputPeer
}

func (*InputNotifyPeer) CRC() uint32 {
	return 0xb8bc5b0c
}
func (*InputNotifyPeer) _IInputNotifyPeer() {}

type InputNotifyUsers struct{}

func (*InputNotifyUsers) CRC() uint32 {
	return 0x193b4417
}
func (*InputNotifyUsers) _IInputNotifyPeer() {}

type InputNotifyChats struct{}

func (*InputNotifyChats) CRC() uint32 {
	return 0x4a95e84e
}
func (*InputNotifyChats) _IInputNotifyPeer() {}

type InputNotifyBroadcasts struct{}

func (*InputNotifyBroadcasts) CRC() uint32 {
	return 0xb1db7c7e
}
func (*InputNotifyBroadcasts) _IInputNotifyPeer() {}

type InputNotifyForumTopic struct {
	Peer     IInputPeer
	TopMsgID int32
}

func (*InputNotifyForumTopic) CRC() uint32 {
	return 0x5c467992
}
func (*InputNotifyForumTopic) _IInputNotifyPeer() {}

type IInputPaymentCredentials interface {
	tl.Object
	_IInputPaymentCredentials()
}

var (
	_ IInputPaymentCredentials = (*InputPaymentCredentialsSaved)(nil)
	_ IInputPaymentCredentials = (*InputPaymentCredentials)(nil)
	_ IInputPaymentCredentials = (*InputPaymentCredentialsApplePay)(nil)
	_ IInputPaymentCredentials = (*InputPaymentCredentialsGooglePay)(nil)
)

type InputPaymentCredentialsSaved struct {
	ID          string
	TmpPassword []byte
}

func (*InputPaymentCredentialsSaved) CRC() uint32 {
	return 0xc10eb2cf
}
func (*InputPaymentCredentialsSaved) _IInputPaymentCredentials() {}

type InputPaymentCredentials struct {
	_    struct{} `tl:"flags,bitflag"`
	Save bool     `tl:",omitempty:flags:0,implicit"`
	Data IDataJSON
}

func (*InputPaymentCredentials) CRC() uint32 {
	return 0x3417d728
}
func (*InputPaymentCredentials) _IInputPaymentCredentials() {}

type InputPaymentCredentialsApplePay struct {
	PaymentData IDataJSON
}

func (*InputPaymentCredentialsApplePay) CRC() uint32 {
	return 0xaa1c39f
}
func (*InputPaymentCredentialsApplePay) _IInputPaymentCredentials() {}

type InputPaymentCredentialsGooglePay struct {
	PaymentToken IDataJSON
}

func (*InputPaymentCredentialsGooglePay) CRC() uint32 {
	return 0x8ac32801
}
func (*InputPaymentCredentialsGooglePay) _IInputPaymentCredentials() {}

// Peer
type IInputPeer interface {
	tl.Object
	_IInputPeer()
}

var (
	_ IInputPeer = (*InputPeerEmpty)(nil)
	_ IInputPeer = (*InputPeerSelf)(nil)
	_ IInputPeer = (*InputPeerChat)(nil)
	_ IInputPeer = (*InputPeerUser)(nil)
	_ IInputPeer = (*InputPeerChannel)(nil)
	_ IInputPeer = (*InputPeerUserFromMessage)(nil)
	_ IInputPeer = (*InputPeerChannelFromMessage)(nil)
)

// An empty constructor, no user or chat is defined.
type InputPeerEmpty struct{}

func (*InputPeerEmpty) CRC() uint32 {
	return 0x7f3b18ea
}
func (*InputPeerEmpty) _IInputPeer() {}

// Defines the current user.
type InputPeerSelf struct{}

func (*InputPeerSelf) CRC() uint32 {
	return 0x7da07ec9
}
func (*InputPeerSelf) _IInputPeer() {}

// Defines a chat for further interaction.
type InputPeerChat struct {
	ChatID int64 // Chat idientifier
}

func (*InputPeerChat) CRC() uint32 {
	return 0x35a95cb9
}
func (*InputPeerChat) _IInputPeer() {}

// Defines a user for further interaction.
type InputPeerUser struct {
	UserID     int64 // User identifier
	AccessHash int64 // **access_hash** value from the [user](https://core.telegram.org/constructor/user) constructor
}

func (*InputPeerUser) CRC() uint32 {
	return 0xdde8a54c
}
func (*InputPeerUser) _IInputPeer() {}

// Defines a channel for further interaction.
type InputPeerChannel struct {
	ChannelID  int64 // Channel identifier
	AccessHash int64 // **access_hash** value from the [channel](https://core.telegram.org/constructor/channel) constructor
}

func (*InputPeerChannel) CRC() uint32 {
	return 0x27bcbbfc
}
func (*InputPeerChannel) _IInputPeer() {}

// Defines a [min](https://core.telegram.org/api/min) user that was seen in a certain message of a certain chat.
type InputPeerUserFromMessage struct {
	Peer   IInputPeer // The chat where the user was seen
	MsgID  int32      // The message ID
	UserID int64      // The identifier of the user that was seen
}

func (*InputPeerUserFromMessage) CRC() uint32 {
	return 0xa87b0a1c
}
func (*InputPeerUserFromMessage) _IInputPeer() {}

// Defines a [min](https://core.telegram.org/api/min) channel that was seen in a certain message of a certain chat.
type InputPeerChannelFromMessage struct {
	Peer      IInputPeer // The chat where the channel's message was seen
	MsgID     int32      // The message ID
	ChannelID int64      // The identifier of the channel that was seen
}

func (*InputPeerChannelFromMessage) CRC() uint32 {
	return 0xbd2a0840
}
func (*InputPeerChannelFromMessage) _IInputPeer() {}

type IInputPeerNotifySettings interface {
	tl.Object
	_IInputPeerNotifySettings()
}

var (
	_ IInputPeerNotifySettings = (*InputPeerNotifySettings)(nil)
)

type InputPeerNotifySettings struct {
	_                 struct{}           `tl:"flags,bitflag"`
	ShowPreviews      *bool              `tl:",omitempty:flags:0"`
	Silent            *bool              `tl:",omitempty:flags:1"`
	MuteUntil         *int32             `tl:",omitempty:flags:2"`
	Sound             INotificationSound `tl:",omitempty:flags:3"`
	StoriesMuted      *bool              `tl:",omitempty:flags:6"`
	StoriesHideSender *bool              `tl:",omitempty:flags:7"`
	StoriesSound      INotificationSound `tl:",omitempty:flags:8"`
}

func (*InputPeerNotifySettings) CRC() uint32 {
	return 0xcacb6ae2
}
func (*InputPeerNotifySettings) _IInputPeerNotifySettings() {}

type IInputPhoneCall interface {
	tl.Object
	_IInputPhoneCall()
}

var (
	_ IInputPhoneCall = (*InputPhoneCall)(nil)
)

type InputPhoneCall struct {
	ID         int64
	AccessHash int64
}

func (*InputPhoneCall) CRC() uint32 {
	return 0x1e36fded
}
func (*InputPhoneCall) _IInputPhoneCall() {}

type IInputPhoto interface {
	tl.Object
	_IInputPhoto()
}

var (
	_ IInputPhoto = (*InputPhotoEmpty)(nil)
	_ IInputPhoto = (*InputPhoto)(nil)
)

type InputPhotoEmpty struct{}

func (*InputPhotoEmpty) CRC() uint32 {
	return 0x1cd7bf0d
}
func (*InputPhotoEmpty) _IInputPhoto() {}

type InputPhoto struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

func (*InputPhoto) CRC() uint32 {
	return 0x3bb3b94a
}
func (*InputPhoto) _IInputPhoto() {}

type IInputPrivacyRule interface {
	tl.Object
	_IInputPrivacyRule()
}

var (
	_ IInputPrivacyRule = (*InputPrivacyValueAllowContacts)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueAllowAll)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueAllowUsers)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueDisallowContacts)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueDisallowAll)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueDisallowUsers)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueAllowChatParticipants)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueDisallowChatParticipants)(nil)
	_ IInputPrivacyRule = (*InputPrivacyValueAllowCloseFriends)(nil)
)

type InputPrivacyValueAllowContacts struct{}

func (*InputPrivacyValueAllowContacts) CRC() uint32 {
	return 0xd09e07b
}
func (*InputPrivacyValueAllowContacts) _IInputPrivacyRule() {}

type InputPrivacyValueAllowAll struct{}

func (*InputPrivacyValueAllowAll) CRC() uint32 {
	return 0x184b35ce
}
func (*InputPrivacyValueAllowAll) _IInputPrivacyRule() {}

type InputPrivacyValueAllowUsers struct {
	Users []IInputUser
}

func (*InputPrivacyValueAllowUsers) CRC() uint32 {
	return 0x131cc67f
}
func (*InputPrivacyValueAllowUsers) _IInputPrivacyRule() {}

type InputPrivacyValueDisallowContacts struct{}

func (*InputPrivacyValueDisallowContacts) CRC() uint32 {
	return 0xba52007
}
func (*InputPrivacyValueDisallowContacts) _IInputPrivacyRule() {}

type InputPrivacyValueDisallowAll struct{}

func (*InputPrivacyValueDisallowAll) CRC() uint32 {
	return 0xd66b66c9
}
func (*InputPrivacyValueDisallowAll) _IInputPrivacyRule() {}

type InputPrivacyValueDisallowUsers struct {
	Users []IInputUser
}

func (*InputPrivacyValueDisallowUsers) CRC() uint32 {
	return 0x90110467
}
func (*InputPrivacyValueDisallowUsers) _IInputPrivacyRule() {}

type InputPrivacyValueAllowChatParticipants struct {
	Chats []int64
}

func (*InputPrivacyValueAllowChatParticipants) CRC() uint32 {
	return 0x840649cf
}
func (*InputPrivacyValueAllowChatParticipants) _IInputPrivacyRule() {}

type InputPrivacyValueDisallowChatParticipants struct {
	Chats []int64
}

func (*InputPrivacyValueDisallowChatParticipants) CRC() uint32 {
	return 0xe94f0f86
}
func (*InputPrivacyValueDisallowChatParticipants) _IInputPrivacyRule() {}

type InputPrivacyValueAllowCloseFriends struct{}

func (*InputPrivacyValueAllowCloseFriends) CRC() uint32 {
	return 0x2f453e49
}
func (*InputPrivacyValueAllowCloseFriends) _IInputPrivacyRule() {}

type IInputReplyTo interface {
	tl.Object
	_IInputReplyTo()
}

var (
	_ IInputReplyTo = (*InputReplyToMessage)(nil)
	_ IInputReplyTo = (*InputReplyToStory)(nil)
)

type InputReplyToMessage struct {
	_             struct{} `tl:"flags,bitflag"`
	ReplyToMsgID  int32
	TopMsgID      *int32           `tl:",omitempty:flags:0"`
	ReplyToPeerID IInputPeer       `tl:",omitempty:flags:1"`
	QuoteText     *string          `tl:",omitempty:flags:2"`
	QuoteEntities []IMessageEntity `tl:",omitempty:flags:3"`
	QuoteOffset   *int32           `tl:",omitempty:flags:4"`
}

func (*InputReplyToMessage) CRC() uint32 {
	return 0x22c0f6d5
}
func (*InputReplyToMessage) _IInputReplyTo() {}

type InputReplyToStory struct {
	UserID  IInputUser
	StoryID int32
}

func (*InputReplyToStory) CRC() uint32 {
	return 0x15b0f283
}
func (*InputReplyToStory) _IInputReplyTo() {}

type IInputSecureFile interface {
	tl.Object
	_IInputSecureFile()
}

var (
	_ IInputSecureFile = (*InputSecureFileUploaded)(nil)
	_ IInputSecureFile = (*InputSecureFile)(nil)
)

type InputSecureFileUploaded struct {
	ID          int64
	Parts       int32
	Md5Checksum string
	FileHash    []byte
	Secret      []byte
}

func (*InputSecureFileUploaded) CRC() uint32 {
	return 0x3334b0f0
}
func (*InputSecureFileUploaded) _IInputSecureFile() {}

type InputSecureFile struct {
	ID         int64
	AccessHash int64
}

func (*InputSecureFile) CRC() uint32 {
	return 0x5367e5be
}
func (*InputSecureFile) _IInputSecureFile() {}

type IInputSecureValue interface {
	tl.Object
	_IInputSecureValue()
}

var (
	_ IInputSecureValue = (*InputSecureValue)(nil)
)

type InputSecureValue struct {
	_           struct{} `tl:"flags,bitflag"`
	Type        EnumSecureValueType
	Data        ISecureData        `tl:",omitempty:flags:0"`
	FrontSide   IInputSecureFile   `tl:",omitempty:flags:1"`
	ReverseSide IInputSecureFile   `tl:",omitempty:flags:2"`
	Selfie      IInputSecureFile   `tl:",omitempty:flags:3"`
	Translation []IInputSecureFile `tl:",omitempty:flags:6"`
	Files       []IInputSecureFile `tl:",omitempty:flags:4"`
	PlainData   ISecurePlainData   `tl:",omitempty:flags:5"`
}

func (*InputSecureValue) CRC() uint32 {
	return 0xdb21d0a7
}
func (*InputSecureValue) _IInputSecureValue() {}

type IInputSingleMedia interface {
	tl.Object
	_IInputSingleMedia()
}

var (
	_ IInputSingleMedia = (*InputSingleMedia)(nil)
)

type InputSingleMedia struct {
	_        struct{} `tl:"flags,bitflag"`
	Media    IInputMedia
	RandomID int64
	Message  string
	Entities []IMessageEntity `tl:",omitempty:flags:0"`
}

func (*InputSingleMedia) CRC() uint32 {
	return 0x1cc6e91f
}
func (*InputSingleMedia) _IInputSingleMedia() {}

type IInputStickerSet interface {
	tl.Object
	_IInputStickerSet()
}

var (
	_ IInputStickerSet = (*InputStickerSetEmpty)(nil)
	_ IInputStickerSet = (*InputStickerSetID)(nil)
	_ IInputStickerSet = (*InputStickerSetShortName)(nil)
	_ IInputStickerSet = (*InputStickerSetAnimatedEmoji)(nil)
	_ IInputStickerSet = (*InputStickerSetDice)(nil)
	_ IInputStickerSet = (*InputStickerSetAnimatedEmojiAnimations)(nil)
	_ IInputStickerSet = (*InputStickerSetPremiumGifts)(nil)
	_ IInputStickerSet = (*InputStickerSetEmojiGenericAnimations)(nil)
	_ IInputStickerSet = (*InputStickerSetEmojiDefaultStatuses)(nil)
	_ IInputStickerSet = (*InputStickerSetEmojiDefaultTopicIcons)(nil)
	_ IInputStickerSet = (*InputStickerSetEmojiChannelDefaultStatuses)(nil)
)

type InputStickerSetEmpty struct{}

func (*InputStickerSetEmpty) CRC() uint32 {
	return 0xffb62b95
}
func (*InputStickerSetEmpty) _IInputStickerSet() {}

type InputStickerSetID struct {
	ID         int64
	AccessHash int64
}

func (*InputStickerSetID) CRC() uint32 {
	return 0x9de7a269
}
func (*InputStickerSetID) _IInputStickerSet() {}

type InputStickerSetShortName struct {
	ShortName string
}

func (*InputStickerSetShortName) CRC() uint32 {
	return 0x861cc8a0
}
func (*InputStickerSetShortName) _IInputStickerSet() {}

type InputStickerSetAnimatedEmoji struct{}

func (*InputStickerSetAnimatedEmoji) CRC() uint32 {
	return 0x28703c8
}
func (*InputStickerSetAnimatedEmoji) _IInputStickerSet() {}

type InputStickerSetDice struct {
	Emoticon string
}

func (*InputStickerSetDice) CRC() uint32 {
	return 0xe67f520e
}
func (*InputStickerSetDice) _IInputStickerSet() {}

type InputStickerSetAnimatedEmojiAnimations struct{}

func (*InputStickerSetAnimatedEmojiAnimations) CRC() uint32 {
	return 0xcde3739
}
func (*InputStickerSetAnimatedEmojiAnimations) _IInputStickerSet() {}

type InputStickerSetPremiumGifts struct{}

func (*InputStickerSetPremiumGifts) CRC() uint32 {
	return 0xc88b3b02
}
func (*InputStickerSetPremiumGifts) _IInputStickerSet() {}

type InputStickerSetEmojiGenericAnimations struct{}

func (*InputStickerSetEmojiGenericAnimations) CRC() uint32 {
	return 0x4c4d4ce
}
func (*InputStickerSetEmojiGenericAnimations) _IInputStickerSet() {}

type InputStickerSetEmojiDefaultStatuses struct{}

func (*InputStickerSetEmojiDefaultStatuses) CRC() uint32 {
	return 0x29d0f5ee
}
func (*InputStickerSetEmojiDefaultStatuses) _IInputStickerSet() {}

type InputStickerSetEmojiDefaultTopicIcons struct{}

func (*InputStickerSetEmojiDefaultTopicIcons) CRC() uint32 {
	return 0x44c1f8e9
}
func (*InputStickerSetEmojiDefaultTopicIcons) _IInputStickerSet() {}

type InputStickerSetEmojiChannelDefaultStatuses struct{}

func (*InputStickerSetEmojiChannelDefaultStatuses) CRC() uint32 {
	return 0x49748553
}
func (*InputStickerSetEmojiChannelDefaultStatuses) _IInputStickerSet() {}

type IInputStickerSetItem interface {
	tl.Object
	_IInputStickerSetItem()
}

var (
	_ IInputStickerSetItem = (*InputStickerSetItem)(nil)
)

type InputStickerSetItem struct {
	_          struct{} `tl:"flags,bitflag"`
	Document   IInputDocument
	Emoji      string
	MaskCoords IMaskCoords `tl:",omitempty:flags:0"`
	Keywords   *string     `tl:",omitempty:flags:1"`
}

func (*InputStickerSetItem) CRC() uint32 {
	return 0x32da9e9c
}
func (*InputStickerSetItem) _IInputStickerSetItem() {}

type IInputStickeredMedia interface {
	tl.Object
	_IInputStickeredMedia()
}

var (
	_ IInputStickeredMedia = (*InputStickeredMediaPhoto)(nil)
	_ IInputStickeredMedia = (*InputStickeredMediaDocument)(nil)
)

type InputStickeredMediaPhoto struct {
	ID IInputPhoto
}

func (*InputStickeredMediaPhoto) CRC() uint32 {
	return 0x4a992157
}
func (*InputStickeredMediaPhoto) _IInputStickeredMedia() {}

type InputStickeredMediaDocument struct {
	ID IInputDocument
}

func (*InputStickeredMediaDocument) CRC() uint32 {
	return 0x438865b
}
func (*InputStickeredMediaDocument) _IInputStickeredMedia() {}

type IInputStorePaymentPurpose interface {
	tl.Object
	_IInputStorePaymentPurpose()
}

var (
	_ IInputStorePaymentPurpose = (*InputStorePaymentPremiumSubscription)(nil)
	_ IInputStorePaymentPurpose = (*InputStorePaymentGiftPremium)(nil)
	_ IInputStorePaymentPurpose = (*InputStorePaymentPremiumGiftCode)(nil)
	_ IInputStorePaymentPurpose = (*InputStorePaymentPremiumGiveaway)(nil)
)

type InputStorePaymentPremiumSubscription struct {
	_       struct{} `tl:"flags,bitflag"`
	Restore bool     `tl:",omitempty:flags:0,implicit"`
	Upgrade bool     `tl:",omitempty:flags:1,implicit"`
}

func (*InputStorePaymentPremiumSubscription) CRC() uint32 {
	return 0xa6751e66
}
func (*InputStorePaymentPremiumSubscription) _IInputStorePaymentPurpose() {}

type InputStorePaymentGiftPremium struct {
	UserID   IInputUser
	Currency string
	Amount   int64
}

func (*InputStorePaymentGiftPremium) CRC() uint32 {
	return 0x616f7fe8
}
func (*InputStorePaymentGiftPremium) _IInputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiftCode struct {
	_         struct{} `tl:"flags,bitflag"`
	Users     []IInputUser
	BoostPeer IInputPeer `tl:",omitempty:flags:0"`
	Currency  string
	Amount    int64
}

func (*InputStorePaymentPremiumGiftCode) CRC() uint32 {
	return 0xa3805f3f
}
func (*InputStorePaymentPremiumGiftCode) _IInputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiveaway struct {
	_                  struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers bool     `tl:",omitempty:flags:0,implicit"`
	WinnersAreVisible  bool     `tl:",omitempty:flags:3,implicit"`
	BoostPeer          IInputPeer
	AdditionalPeers    []IInputPeer `tl:",omitempty:flags:1"`
	CountriesIso2      []string     `tl:",omitempty:flags:2"`
	PrizeDescription   *string      `tl:",omitempty:flags:4"`
	RandomID           int64
	UntilDate          int32
	Currency           string
	Amount             int64
}

func (*InputStorePaymentPremiumGiveaway) CRC() uint32 {
	return 0x160544ca
}
func (*InputStorePaymentPremiumGiveaway) _IInputStorePaymentPurpose() {}

type IInputTheme interface {
	tl.Object
	_IInputTheme()
}

var (
	_ IInputTheme = (*InputTheme)(nil)
	_ IInputTheme = (*InputThemeSlug)(nil)
)

type InputTheme struct {
	ID         int64
	AccessHash int64
}

func (*InputTheme) CRC() uint32 {
	return 0x3c5693e9
}
func (*InputTheme) _IInputTheme() {}

type InputThemeSlug struct {
	Slug string
}

func (*InputThemeSlug) CRC() uint32 {
	return 0xf5890df1
}
func (*InputThemeSlug) _IInputTheme() {}

type IInputThemeSettings interface {
	tl.Object
	_IInputThemeSettings()
}

var (
	_ IInputThemeSettings = (*InputThemeSettings)(nil)
)

type InputThemeSettings struct {
	_                     struct{} `tl:"flags,bitflag"`
	MessageColorsAnimated bool     `tl:",omitempty:flags:2,implicit"`
	BaseTheme             EnumBaseTheme
	AccentColor           int32
	OutboxAccentColor     *int32             `tl:",omitempty:flags:3"`
	MessageColors         []int32            `tl:",omitempty:flags:0"`
	Wallpaper             IInputWallPaper    `tl:",omitempty:flags:1"`
	WallpaperSettings     IWallPaperSettings `tl:",omitempty:flags:1"`
}

func (*InputThemeSettings) CRC() uint32 {
	return 0x8fde504f
}
func (*InputThemeSettings) _IInputThemeSettings() {}

// Defines a user for subsequent interaction.
type IInputUser interface {
	tl.Object
	_IInputUser()
}

var (
	_ IInputUser = (*InputUserEmpty)(nil)
	_ IInputUser = (*InputUserSelf)(nil)
	_ IInputUser = (*InputUser)(nil)
	_ IInputUser = (*InputUserFromMessage)(nil)
)

// Empty constructor, does not define a user.
type InputUserEmpty struct{}

func (*InputUserEmpty) CRC() uint32 {
	return 0xb98886cf
}
func (*InputUserEmpty) _IInputUser() {}

// Defines the current user.
type InputUserSelf struct{}

func (*InputUserSelf) CRC() uint32 {
	return 0xf7c1b13f
}
func (*InputUserSelf) _IInputUser() {}

// Defines a user for further interaction.
type InputUser struct {
	UserID     int64 // User identifier
	AccessHash int64 // access_hash value from the user constructor
}

func (*InputUser) CRC() uint32 {
	return 0xf21158c6
}
func (*InputUser) _IInputUser() {}

// Defines a min user that was seen in a certain message of a certain chat.
type InputUserFromMessage struct {
	Peer   IInputPeer // The chat where the user was seen
	MsgID  int32      // The message ID
	UserID int64      // The identifier of the user that was seen
}

func (*InputUserFromMessage) CRC() uint32 {
	return 0x1da448e2
}
func (*InputUserFromMessage) _IInputUser() {}

type IInputWallPaper interface {
	tl.Object
	_IInputWallPaper()
}

var (
	_ IInputWallPaper = (*InputWallPaper)(nil)
	_ IInputWallPaper = (*InputWallPaperSlug)(nil)
	_ IInputWallPaper = (*InputWallPaperNoFile)(nil)
)

type InputWallPaper struct {
	ID         int64
	AccessHash int64
}

func (*InputWallPaper) CRC() uint32 {
	return 0xe630b979
}
func (*InputWallPaper) _IInputWallPaper() {}

type InputWallPaperSlug struct {
	Slug string
}

func (*InputWallPaperSlug) CRC() uint32 {
	return 0x72091c80
}
func (*InputWallPaperSlug) _IInputWallPaper() {}

type InputWallPaperNoFile struct {
	ID int64
}

func (*InputWallPaperNoFile) CRC() uint32 {
	return 0x967a462e
}
func (*InputWallPaperNoFile) _IInputWallPaper() {}

type IInputWebDocument interface {
	tl.Object
	_IInputWebDocument()
}

var (
	_ IInputWebDocument = (*InputWebDocument)(nil)
)

type InputWebDocument struct {
	URL        string
	Size       int32
	MimeType   string
	Attributes []IDocumentAttribute
}

func (*InputWebDocument) CRC() uint32 {
	return 0x9bed434d
}
func (*InputWebDocument) _IInputWebDocument() {}

type IInputWebFileLocation interface {
	tl.Object
	_IInputWebFileLocation()
}

var (
	_ IInputWebFileLocation = (*InputWebFileLocation)(nil)
	_ IInputWebFileLocation = (*InputWebFileGeoPointLocation)(nil)
	_ IInputWebFileLocation = (*InputWebFileAudioAlbumThumbLocation)(nil)
)

type InputWebFileLocation struct {
	URL        string
	AccessHash int64
}

func (*InputWebFileLocation) CRC() uint32 {
	return 0xc239d686
}
func (*InputWebFileLocation) _IInputWebFileLocation() {}

type InputWebFileGeoPointLocation struct {
	GeoPoint   IInputGeoPoint
	AccessHash int64
	W          int32
	H          int32
	Zoom       int32
	Scale      int32
}

func (*InputWebFileGeoPointLocation) CRC() uint32 {
	return 0x9f2221c9
}
func (*InputWebFileGeoPointLocation) _IInputWebFileLocation() {}

type InputWebFileAudioAlbumThumbLocation struct {
	_         struct{}       `tl:"flags,bitflag"`
	Small     bool           `tl:",omitempty:flags:2,implicit"`
	Document  IInputDocument `tl:",omitempty:flags:0"`
	Title     *string        `tl:",omitempty:flags:1"`
	Performer *string        `tl:",omitempty:flags:1"`
}

func (*InputWebFileAudioAlbumThumbLocation) CRC() uint32 {
	return 0xf46fe924
}
func (*InputWebFileAudioAlbumThumbLocation) _IInputWebFileLocation() {}

type IInvoice interface {
	tl.Object
	_IInvoice()
}

var (
	_ IInvoice = (*Invoice)(nil)
)

type Invoice struct {
	_                        struct{} `tl:"flags,bitflag"`
	Test                     bool     `tl:",omitempty:flags:0,implicit"`
	NameRequested            bool     `tl:",omitempty:flags:1,implicit"`
	PhoneRequested           bool     `tl:",omitempty:flags:2,implicit"`
	EmailRequested           bool     `tl:",omitempty:flags:3,implicit"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:4,implicit"`
	Flexible                 bool     `tl:",omitempty:flags:5,implicit"`
	PhoneToProvider          bool     `tl:",omitempty:flags:6,implicit"`
	EmailToProvider          bool     `tl:",omitempty:flags:7,implicit"`
	Recurring                bool     `tl:",omitempty:flags:9,implicit"`
	Currency                 string
	Prices                   []ILabeledPrice
	MaxTipAmount             *int64  `tl:",omitempty:flags:8"`
	SuggestedTipAmounts      []int64 `tl:",omitempty:flags:8"`
	TermsURL                 *string `tl:",omitempty:flags:10"`
}

func (*Invoice) CRC() uint32 {
	return 0x5db95a15
}
func (*Invoice) _IInvoice() {}

type IJSONObjectValue interface {
	tl.Object
	_IJSONObjectValue()
}

var (
	_ IJSONObjectValue = (*JSONObjectValue)(nil)
)

type JSONObjectValue struct {
	Key   string
	Value IJSONValue
}

func (*JSONObjectValue) CRC() uint32 {
	return 0xc0de1bd9
}
func (*JSONObjectValue) _IJSONObjectValue() {}

type IJSONValue interface {
	tl.Object
	_IJSONValue()
}

var (
	_ IJSONValue = (*JSONNull)(nil)
	_ IJSONValue = (*JSONBool)(nil)
	_ IJSONValue = (*JSONNumber)(nil)
	_ IJSONValue = (*JSONString)(nil)
	_ IJSONValue = (*JSONArray)(nil)
	_ IJSONValue = (*JSONObject)(nil)
)

type JSONNull struct{}

func (*JSONNull) CRC() uint32 {
	return 0x3f6d7b68
}
func (*JSONNull) _IJSONValue() {}

type JSONBool struct {
	Value bool
}

func (*JSONBool) CRC() uint32 {
	return 0xc7345e6a
}
func (*JSONBool) _IJSONValue() {}

type JSONNumber struct {
	Value float64
}

func (*JSONNumber) CRC() uint32 {
	return 0x2be0dfa4
}
func (*JSONNumber) _IJSONValue() {}

type JSONString struct {
	Value string
}

func (*JSONString) CRC() uint32 {
	return 0xb71e767a
}
func (*JSONString) _IJSONValue() {}

type JSONArray struct {
	Value []IJSONValue
}

func (*JSONArray) CRC() uint32 {
	return 0xf7444763
}
func (*JSONArray) _IJSONValue() {}

type JSONObject struct {
	Value []IJSONObjectValue
}

func (*JSONObject) CRC() uint32 {
	return 0x99c1d49d
}
func (*JSONObject) _IJSONValue() {}

type IKeyboardButton interface {
	tl.Object
	_IKeyboardButton()
}

var (
	_ IKeyboardButton = (*KeyboardButton)(nil)
	_ IKeyboardButton = (*KeyboardButtonURL)(nil)
	_ IKeyboardButton = (*KeyboardButtonCallback)(nil)
	_ IKeyboardButton = (*KeyboardButtonRequestPhone)(nil)
	_ IKeyboardButton = (*KeyboardButtonRequestGeoLocation)(nil)
	_ IKeyboardButton = (*KeyboardButtonSwitchInline)(nil)
	_ IKeyboardButton = (*KeyboardButtonGame)(nil)
	_ IKeyboardButton = (*KeyboardButtonBuy)(nil)
	_ IKeyboardButton = (*KeyboardButtonURLAuth)(nil)
	_ IKeyboardButton = (*InputKeyboardButtonURLAuth)(nil)
	_ IKeyboardButton = (*KeyboardButtonRequestPoll)(nil)
	_ IKeyboardButton = (*InputKeyboardButtonUserProfile)(nil)
	_ IKeyboardButton = (*KeyboardButtonUserProfile)(nil)
	_ IKeyboardButton = (*KeyboardButtonWebView)(nil)
	_ IKeyboardButton = (*KeyboardButtonSimpleWebView)(nil)
	_ IKeyboardButton = (*KeyboardButtonRequestPeer)(nil)
)

type KeyboardButton struct {
	Text string
}

func (*KeyboardButton) CRC() uint32 {
	return 0xa2fa4880
}
func (*KeyboardButton) _IKeyboardButton() {}

type KeyboardButtonURL struct {
	Text string
	URL  string
}

func (*KeyboardButtonURL) CRC() uint32 {
	return 0x258aff05
}
func (*KeyboardButtonURL) _IKeyboardButton() {}

type KeyboardButtonCallback struct {
	_                struct{} `tl:"flags,bitflag"`
	RequiresPassword bool     `tl:",omitempty:flags:0,implicit"`
	Text             string
	Data             []byte
}

func (*KeyboardButtonCallback) CRC() uint32 {
	return 0x35bbdb6b
}
func (*KeyboardButtonCallback) _IKeyboardButton() {}

type KeyboardButtonRequestPhone struct {
	Text string
}

func (*KeyboardButtonRequestPhone) CRC() uint32 {
	return 0xb16a6c29
}
func (*KeyboardButtonRequestPhone) _IKeyboardButton() {}

type KeyboardButtonRequestGeoLocation struct {
	Text string
}

func (*KeyboardButtonRequestGeoLocation) CRC() uint32 {
	return 0xfc796b3f
}
func (*KeyboardButtonRequestGeoLocation) _IKeyboardButton() {}

type KeyboardButtonSwitchInline struct {
	_         struct{} `tl:"flags,bitflag"`
	SamePeer  bool     `tl:",omitempty:flags:0,implicit"`
	Text      string
	Query     string
	PeerTypes []EnumInlineQueryPeerType `tl:",omitempty:flags:1"`
}

func (*KeyboardButtonSwitchInline) CRC() uint32 {
	return 0x93b9fbb5
}
func (*KeyboardButtonSwitchInline) _IKeyboardButton() {}

type KeyboardButtonGame struct {
	Text string
}

func (*KeyboardButtonGame) CRC() uint32 {
	return 0x50f41ccf
}
func (*KeyboardButtonGame) _IKeyboardButton() {}

type KeyboardButtonBuy struct {
	Text string
}

func (*KeyboardButtonBuy) CRC() uint32 {
	return 0xafd93fbb
}
func (*KeyboardButtonBuy) _IKeyboardButton() {}

type KeyboardButtonURLAuth struct {
	_        struct{} `tl:"flags,bitflag"`
	Text     string
	FwdText  *string `tl:",omitempty:flags:0"`
	URL      string
	ButtonID int32
}

func (*KeyboardButtonURLAuth) CRC() uint32 {
	return 0x10b78d29
}
func (*KeyboardButtonURLAuth) _IKeyboardButton() {}

type InputKeyboardButtonURLAuth struct {
	_                  struct{} `tl:"flags,bitflag"`
	RequestWriteAccess bool     `tl:",omitempty:flags:0,implicit"`
	Text               string
	FwdText            *string `tl:",omitempty:flags:1"`
	URL                string
	Bot                IInputUser
}

func (*InputKeyboardButtonURLAuth) CRC() uint32 {
	return 0xd02e7fd4
}
func (*InputKeyboardButtonURLAuth) _IKeyboardButton() {}

type KeyboardButtonRequestPoll struct {
	_    struct{} `tl:"flags,bitflag"`
	Quiz *bool    `tl:",omitempty:flags:0"`
	Text string
}

func (*KeyboardButtonRequestPoll) CRC() uint32 {
	return 0xbbc7515d
}
func (*KeyboardButtonRequestPoll) _IKeyboardButton() {}

type InputKeyboardButtonUserProfile struct {
	Text   string
	UserID IInputUser
}

func (*InputKeyboardButtonUserProfile) CRC() uint32 {
	return 0xe988037b
}
func (*InputKeyboardButtonUserProfile) _IKeyboardButton() {}

type KeyboardButtonUserProfile struct {
	Text   string
	UserID int64
}

func (*KeyboardButtonUserProfile) CRC() uint32 {
	return 0x308660c1
}
func (*KeyboardButtonUserProfile) _IKeyboardButton() {}

type KeyboardButtonWebView struct {
	Text string
	URL  string
}

func (*KeyboardButtonWebView) CRC() uint32 {
	return 0x13767230
}
func (*KeyboardButtonWebView) _IKeyboardButton() {}

type KeyboardButtonSimpleWebView struct {
	Text string
	URL  string
}

func (*KeyboardButtonSimpleWebView) CRC() uint32 {
	return 0xa0c0505c
}
func (*KeyboardButtonSimpleWebView) _IKeyboardButton() {}

type KeyboardButtonRequestPeer struct {
	Text        string
	ButtonID    int32
	PeerType    IRequestPeerType
	MaxQuantity int32
}

func (*KeyboardButtonRequestPeer) CRC() uint32 {
	return 0x53d7bfd8
}
func (*KeyboardButtonRequestPeer) _IKeyboardButton() {}

type IKeyboardButtonRow interface {
	tl.Object
	_IKeyboardButtonRow()
}

var (
	_ IKeyboardButtonRow = (*KeyboardButtonRow)(nil)
)

type KeyboardButtonRow struct {
	Buttons []IKeyboardButton
}

func (*KeyboardButtonRow) CRC() uint32 {
	return 0x77608b83
}
func (*KeyboardButtonRow) _IKeyboardButtonRow() {}

type ILabeledPrice interface {
	tl.Object
	_ILabeledPrice()
}

var (
	_ ILabeledPrice = (*LabeledPrice)(nil)
)

type LabeledPrice struct {
	Label  string
	Amount int64
}

func (*LabeledPrice) CRC() uint32 {
	return 0xcb296bf8
}
func (*LabeledPrice) _ILabeledPrice() {}

type ILangPackDifference interface {
	tl.Object
	_ILangPackDifference()
}

var (
	_ ILangPackDifference = (*LangPackDifference)(nil)
)

type LangPackDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Strings     []ILangPackString
}

func (*LangPackDifference) CRC() uint32 {
	return 0xf385c1f6
}
func (*LangPackDifference) _ILangPackDifference() {}

type ILangPackLanguage interface {
	tl.Object
	_ILangPackLanguage()
}

var (
	_ ILangPackLanguage = (*LangPackLanguage)(nil)
)

type LangPackLanguage struct {
	_               struct{} `tl:"flags,bitflag"`
	Official        bool     `tl:",omitempty:flags:0,implicit"`
	Rtl             bool     `tl:",omitempty:flags:2,implicit"`
	Beta            bool     `tl:",omitempty:flags:3,implicit"`
	Name            string
	NativeName      string
	LangCode        string
	BaseLangCode    *string `tl:",omitempty:flags:1"`
	PluralCode      string
	StringsCount    int32
	TranslatedCount int32
	TranslationsURL string
}

func (*LangPackLanguage) CRC() uint32 {
	return 0xeeca5ce3
}
func (*LangPackLanguage) _ILangPackLanguage() {}

type ILangPackString interface {
	tl.Object
	_ILangPackString()
}

var (
	_ ILangPackString = (*LangPackString)(nil)
	_ ILangPackString = (*LangPackStringPluralized)(nil)
	_ ILangPackString = (*LangPackStringDeleted)(nil)
)

type LangPackString struct {
	Key   string
	Value string
}

func (*LangPackString) CRC() uint32 {
	return 0xcad181f6
}
func (*LangPackString) _ILangPackString() {}

type LangPackStringPluralized struct {
	_          struct{} `tl:"flags,bitflag"`
	Key        string
	ZeroValue  *string `tl:",omitempty:flags:0"`
	OneValue   *string `tl:",omitempty:flags:1"`
	TwoValue   *string `tl:",omitempty:flags:2"`
	FewValue   *string `tl:",omitempty:flags:3"`
	ManyValue  *string `tl:",omitempty:flags:4"`
	OtherValue string
}

func (*LangPackStringPluralized) CRC() uint32 {
	return 0x6c47ac9f
}
func (*LangPackStringPluralized) _ILangPackString() {}

type LangPackStringDeleted struct {
	Key string
}

func (*LangPackStringDeleted) CRC() uint32 {
	return 0x2979eeb2
}
func (*LangPackStringDeleted) _ILangPackString() {}

type IMaskCoords interface {
	tl.Object
	_IMaskCoords()
}

var (
	_ IMaskCoords = (*MaskCoords)(nil)
)

type MaskCoords struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

func (*MaskCoords) CRC() uint32 {
	return 0xaed6dbb2
}
func (*MaskCoords) _IMaskCoords() {}

type IMediaArea interface {
	tl.Object
	_IMediaArea()
}

var (
	_ IMediaArea = (*MediaAreaVenue)(nil)
	_ IMediaArea = (*InputMediaAreaVenue)(nil)
	_ IMediaArea = (*MediaAreaGeoPoint)(nil)
	_ IMediaArea = (*MediaAreaSuggestedReaction)(nil)
	_ IMediaArea = (*MediaAreaChannelPost)(nil)
	_ IMediaArea = (*InputMediaAreaChannelPost)(nil)
)

type MediaAreaVenue struct {
	Coordinates IMediaAreaCoordinates
	Geo         IGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
}

func (*MediaAreaVenue) CRC() uint32 {
	return 0xbe82db9c
}
func (*MediaAreaVenue) _IMediaArea() {}

type InputMediaAreaVenue struct {
	Coordinates IMediaAreaCoordinates
	QueryID     int64
	ResultID    string
}

func (*InputMediaAreaVenue) CRC() uint32 {
	return 0xb282217f
}
func (*InputMediaAreaVenue) _IMediaArea() {}

type MediaAreaGeoPoint struct {
	Coordinates IMediaAreaCoordinates
	Geo         IGeoPoint
}

func (*MediaAreaGeoPoint) CRC() uint32 {
	return 0xdf8b3b22
}
func (*MediaAreaGeoPoint) _IMediaArea() {}

type MediaAreaSuggestedReaction struct {
	_           struct{} `tl:"flags,bitflag"`
	Dark        bool     `tl:",omitempty:flags:0,implicit"`
	Flipped     bool     `tl:",omitempty:flags:1,implicit"`
	Coordinates IMediaAreaCoordinates
	Reaction    IReaction
}

func (*MediaAreaSuggestedReaction) CRC() uint32 {
	return 0x14455871
}
func (*MediaAreaSuggestedReaction) _IMediaArea() {}

type MediaAreaChannelPost struct {
	Coordinates IMediaAreaCoordinates
	ChannelID   int64
	MsgID       int32
}

func (*MediaAreaChannelPost) CRC() uint32 {
	return 0x770416af
}
func (*MediaAreaChannelPost) _IMediaArea() {}

type InputMediaAreaChannelPost struct {
	Coordinates IMediaAreaCoordinates
	Channel     IInputChannel
	MsgID       int32
}

func (*InputMediaAreaChannelPost) CRC() uint32 {
	return 0x2271f2bf
}
func (*InputMediaAreaChannelPost) _IMediaArea() {}

type IMediaAreaCoordinates interface {
	tl.Object
	_IMediaAreaCoordinates()
}

var (
	_ IMediaAreaCoordinates = (*MediaAreaCoordinates)(nil)
)

type MediaAreaCoordinates struct {
	X        float64
	Y        float64
	W        float64
	H        float64
	Rotation float64
}

func (*MediaAreaCoordinates) CRC() uint32 {
	return 0x3d1ea4e
}
func (*MediaAreaCoordinates) _IMediaAreaCoordinates() {}

// Object describing a message.
type IMessage interface {
	tl.Object
	_IMessage()
}

var (
	_ IMessage = (*MessageEmpty)(nil)
	_ IMessage = (*Message)(nil)
	_ IMessage = (*MessageService)(nil)
)

// Empty constructor, non-existent message.
type MessageEmpty struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32    // Message identifier
	PeerID IPeer    `tl:",omitempty:flags:0"`
}

func (*MessageEmpty) CRC() uint32 {
	return 0x90a6ca84
}
func (*MessageEmpty) _IMessage() {}

// A message
type Message struct {
	_                 struct{}             `tl:"flags,bitflag"`
	Out               bool                 `tl:",omitempty:flags:1,implicit"`  // Is this an outgoing message
	Mentioned         bool                 `tl:",omitempty:flags:4,implicit"`  // Whether we were mentioned in this message
	MediaUnread       bool                 `tl:",omitempty:flags:5,implicit"`  // Whether there are unread media attachments in this message
	Silent            bool                 `tl:",omitempty:flags:13,implicit"` // Whether this is a silent message (no notification triggered)
	Post              bool                 `tl:",omitempty:flags:14,implicit"` // Whether this is a channel post
	FromScheduled     bool                 `tl:",omitempty:flags:18,implicit"` // Whether this is a scheduled message
	Legacy            bool                 `tl:",omitempty:flags:19,implicit"` // This is a legacy message: it has to be refetched with the new layer
	EditHide          bool                 `tl:",omitempty:flags:21,implicit"` // Whether the message should be shown as not modified to the user, even if an edit date is present
	Pinned            bool                 `tl:",omitempty:flags:24,implicit"` // Whether this message is pinned
	Noforwards        bool                 `tl:",omitempty:flags:26,implicit"`
	InvertMedia       bool                 `tl:",omitempty:flags:27,implicit"`
	ID                int32                // ID of the message
	FromID            IPeer                `tl:",omitempty:flags:8"` // ID of the sender of the message
	PeerID            IPeer                // Peer ID, the chat where this message was sent
	FwdFrom           IMessageFwdHeader    `tl:",omitempty:flags:2"`  // Info about forwarded messages
	ViaBotID          *int64               `tl:",omitempty:flags:11"` // ID of the inline bot that generated the message
	ReplyTo           IMessageReplyHeader  `tl:",omitempty:flags:3"`  // Reply information
	Date              int32                // Date of the message
	Message           string               // The message
	Media             IMessageMedia        `tl:",omitempty:flags:9"`  // Media attachment
	ReplyMarkup       IReplyMarkup         `tl:",omitempty:flags:6"`  // Reply markup (bot/inline keyboards)
	Entities          []IMessageEntity     `tl:",omitempty:flags:7"`  // Message entities for styled text
	Views             *int32               `tl:",omitempty:flags:10"` // View count for channel posts
	Forwards          *int32               `tl:",omitempty:flags:10"` // Forward counter
	Replies           IMessageReplies      `tl:",omitempty:flags:23"` // Info about post comments (for channels) or message replies (for groups)
	EditDate          *int32               `tl:",omitempty:flags:15"` // Last edit date of this message
	PostAuthor        *string              `tl:",omitempty:flags:16"` // Name of the author of this message for channel posts (with signatures enabled)
	GroupedID         *int64               `tl:",omitempty:flags:17"` // Multiple media messages sent using messages.sendMultiMedia with the same grouped ID indicate an album or media group
	Reactions         IMessageReactions    `tl:",omitempty:flags:20"`
	RestrictionReason []IRestrictionReason `tl:",omitempty:flags:22"` // Contains the reason why access to this message must be restricted.
	TTLPeriod         *int32               `tl:",omitempty:flags:25"`
}

func (*Message) CRC() uint32 {
	return 0x38116ee0
}
func (*Message) _IMessage() {}

// Indicates a service message
type MessageService struct {
	_           struct{}            `tl:"flags,bitflag"`
	Out         bool                `tl:",omitempty:flags:1,implicit"`  // Whether the message is outgoing
	Mentioned   bool                `tl:",omitempty:flags:4,implicit"`  // Whether we were mentioned in the message
	MediaUnread bool                `tl:",omitempty:flags:5,implicit"`  // Whether the message contains unread media
	Silent      bool                `tl:",omitempty:flags:13,implicit"` // Whether the message is silent
	Post        bool                `tl:",omitempty:flags:14,implicit"` // Whether it's a channel post
	Legacy      bool                `tl:",omitempty:flags:19,implicit"` // This is a legacy message: it has to be refetched with the new layer
	ID          int32               // Message ID
	FromID      IPeer               `tl:",omitempty:flags:8"` // ID of the sender of this message
	PeerID      IPeer               // Sender of service message
	ReplyTo     IMessageReplyHeader `tl:",omitempty:flags:3"` // Reply (thread) information
	Date        int32               // Message date
	Action      IMessageAction      // Event connected with the service message
	TTLPeriod   *int32              `tl:",omitempty:flags:25"`
}

func (*MessageService) CRC() uint32 {
	return 0x2b085862
}
func (*MessageService) _IMessage() {}

type IMessageAction interface {
	tl.Object
	_IMessageAction()
}

var (
	_ IMessageAction = (*MessageActionEmpty)(nil)
	_ IMessageAction = (*MessageActionChatCreate)(nil)
	_ IMessageAction = (*MessageActionChatEditTitle)(nil)
	_ IMessageAction = (*MessageActionChatEditPhoto)(nil)
	_ IMessageAction = (*MessageActionChatDeletePhoto)(nil)
	_ IMessageAction = (*MessageActionChatAddUser)(nil)
	_ IMessageAction = (*MessageActionChatDeleteUser)(nil)
	_ IMessageAction = (*MessageActionChatJoinedByLink)(nil)
	_ IMessageAction = (*MessageActionChannelCreate)(nil)
	_ IMessageAction = (*MessageActionChatMigrateTo)(nil)
	_ IMessageAction = (*MessageActionChannelMigrateFrom)(nil)
	_ IMessageAction = (*MessageActionPinMessage)(nil)
	_ IMessageAction = (*MessageActionHistoryClear)(nil)
	_ IMessageAction = (*MessageActionGameScore)(nil)
	_ IMessageAction = (*MessageActionPaymentSentMe)(nil)
	_ IMessageAction = (*MessageActionPaymentSent)(nil)
	_ IMessageAction = (*MessageActionPhoneCall)(nil)
	_ IMessageAction = (*MessageActionScreenshotTaken)(nil)
	_ IMessageAction = (*MessageActionCustomAction)(nil)
	_ IMessageAction = (*MessageActionBotAllowed)(nil)
	_ IMessageAction = (*MessageActionSecureValuesSentMe)(nil)
	_ IMessageAction = (*MessageActionSecureValuesSent)(nil)
	_ IMessageAction = (*MessageActionContactSignUp)(nil)
	_ IMessageAction = (*MessageActionGeoProximityReached)(nil)
	_ IMessageAction = (*MessageActionGroupCall)(nil)
	_ IMessageAction = (*MessageActionInviteToGroupCall)(nil)
	_ IMessageAction = (*MessageActionSetMessagesTTL)(nil)
	_ IMessageAction = (*MessageActionGroupCallScheduled)(nil)
	_ IMessageAction = (*MessageActionSetChatTheme)(nil)
	_ IMessageAction = (*MessageActionChatJoinedByRequest)(nil)
	_ IMessageAction = (*MessageActionWebViewDataSentMe)(nil)
	_ IMessageAction = (*MessageActionWebViewDataSent)(nil)
	_ IMessageAction = (*MessageActionGiftPremium)(nil)
	_ IMessageAction = (*MessageActionTopicCreate)(nil)
	_ IMessageAction = (*MessageActionTopicEdit)(nil)
	_ IMessageAction = (*MessageActionSuggestProfilePhoto)(nil)
	_ IMessageAction = (*MessageActionRequestedPeer)(nil)
	_ IMessageAction = (*MessageActionSetChatWallPaper)(nil)
	_ IMessageAction = (*MessageActionGiftCode)(nil)
	_ IMessageAction = (*MessageActionGiveawayLaunch)(nil)
	_ IMessageAction = (*MessageActionGiveawayResults)(nil)
)

type MessageActionEmpty struct{}

func (*MessageActionEmpty) CRC() uint32 {
	return 0xb6aef7b0
}
func (*MessageActionEmpty) _IMessageAction() {}

type MessageActionChatCreate struct {
	Title string
	Users []int64
}

func (*MessageActionChatCreate) CRC() uint32 {
	return 0xbd47cbad
}
func (*MessageActionChatCreate) _IMessageAction() {}

type MessageActionChatEditTitle struct {
	Title string
}

func (*MessageActionChatEditTitle) CRC() uint32 {
	return 0xb5a1ce5a
}
func (*MessageActionChatEditTitle) _IMessageAction() {}

type MessageActionChatEditPhoto struct {
	Photo IPhoto
}

func (*MessageActionChatEditPhoto) CRC() uint32 {
	return 0x7fcb13a8
}
func (*MessageActionChatEditPhoto) _IMessageAction() {}

type MessageActionChatDeletePhoto struct{}

func (*MessageActionChatDeletePhoto) CRC() uint32 {
	return 0x95e3fbef
}
func (*MessageActionChatDeletePhoto) _IMessageAction() {}

type MessageActionChatAddUser struct {
	Users []int64
}

func (*MessageActionChatAddUser) CRC() uint32 {
	return 0x15cefd00
}
func (*MessageActionChatAddUser) _IMessageAction() {}

type MessageActionChatDeleteUser struct {
	UserID int64
}

func (*MessageActionChatDeleteUser) CRC() uint32 {
	return 0xa43f30cc
}
func (*MessageActionChatDeleteUser) _IMessageAction() {}

type MessageActionChatJoinedByLink struct {
	InviterID int64
}

func (*MessageActionChatJoinedByLink) CRC() uint32 {
	return 0x31224c3
}
func (*MessageActionChatJoinedByLink) _IMessageAction() {}

type MessageActionChannelCreate struct {
	Title string
}

func (*MessageActionChannelCreate) CRC() uint32 {
	return 0x95d2ac92
}
func (*MessageActionChannelCreate) _IMessageAction() {}

type MessageActionChatMigrateTo struct {
	ChannelID int64
}

func (*MessageActionChatMigrateTo) CRC() uint32 {
	return 0xe1037f92
}
func (*MessageActionChatMigrateTo) _IMessageAction() {}

type MessageActionChannelMigrateFrom struct {
	Title  string
	ChatID int64
}

func (*MessageActionChannelMigrateFrom) CRC() uint32 {
	return 0xea3948e9
}
func (*MessageActionChannelMigrateFrom) _IMessageAction() {}

type MessageActionPinMessage struct{}

func (*MessageActionPinMessage) CRC() uint32 {
	return 0x94bd38ed
}
func (*MessageActionPinMessage) _IMessageAction() {}

type MessageActionHistoryClear struct{}

func (*MessageActionHistoryClear) CRC() uint32 {
	return 0x9fbab604
}
func (*MessageActionHistoryClear) _IMessageAction() {}

type MessageActionGameScore struct {
	GameID int64
	Score  int32
}

func (*MessageActionGameScore) CRC() uint32 {
	return 0x92a72876
}
func (*MessageActionGameScore) _IMessageAction() {}

type MessageActionPaymentSentMe struct {
	_                struct{} `tl:"flags,bitflag"`
	RecurringInit    bool     `tl:",omitempty:flags:2,implicit"`
	RecurringUsed    bool     `tl:",omitempty:flags:3,implicit"`
	Currency         string
	TotalAmount      int64
	Payload          []byte
	Info             IPaymentRequestedInfo `tl:",omitempty:flags:0"`
	ShippingOptionID *string               `tl:",omitempty:flags:1"`
	Charge           IPaymentCharge
}

func (*MessageActionPaymentSentMe) CRC() uint32 {
	return 0x8f31b327
}
func (*MessageActionPaymentSentMe) _IMessageAction() {}

type MessageActionPaymentSent struct {
	_             struct{} `tl:"flags,bitflag"`
	RecurringInit bool     `tl:",omitempty:flags:2,implicit"`
	RecurringUsed bool     `tl:",omitempty:flags:3,implicit"`
	Currency      string
	TotalAmount   int64
	InvoiceSlug   *string `tl:",omitempty:flags:0"`
}

func (*MessageActionPaymentSent) CRC() uint32 {
	return 0x96163f56
}
func (*MessageActionPaymentSent) _IMessageAction() {}

type MessageActionPhoneCall struct {
	_        struct{} `tl:"flags,bitflag"`
	Video    bool     `tl:",omitempty:flags:2,implicit"`
	CallID   int64
	Reason   EnumPhoneCallDiscardReason `tl:",omitempty:flags:0"`
	Duration *int32                     `tl:",omitempty:flags:1"`
}

func (*MessageActionPhoneCall) CRC() uint32 {
	return 0x80e11a7f
}
func (*MessageActionPhoneCall) _IMessageAction() {}

type MessageActionScreenshotTaken struct{}

func (*MessageActionScreenshotTaken) CRC() uint32 {
	return 0x4792929b
}
func (*MessageActionScreenshotTaken) _IMessageAction() {}

type MessageActionCustomAction struct {
	Message string
}

func (*MessageActionCustomAction) CRC() uint32 {
	return 0xfae69f56
}
func (*MessageActionCustomAction) _IMessageAction() {}

type MessageActionBotAllowed struct {
	_           struct{} `tl:"flags,bitflag"`
	AttachMenu  bool     `tl:",omitempty:flags:1,implicit"`
	FromRequest bool     `tl:",omitempty:flags:3,implicit"`
	Domain      *string  `tl:",omitempty:flags:0"`
	App         IBotApp  `tl:",omitempty:flags:2"`
}

func (*MessageActionBotAllowed) CRC() uint32 {
	return 0xc516d679
}
func (*MessageActionBotAllowed) _IMessageAction() {}

type MessageActionSecureValuesSentMe struct {
	Values      []ISecureValue
	Credentials ISecureCredentialsEncrypted
}

func (*MessageActionSecureValuesSentMe) CRC() uint32 {
	return 0x1b287353
}
func (*MessageActionSecureValuesSentMe) _IMessageAction() {}

type MessageActionSecureValuesSent struct {
	Types []EnumSecureValueType
}

func (*MessageActionSecureValuesSent) CRC() uint32 {
	return 0xd95c6154
}
func (*MessageActionSecureValuesSent) _IMessageAction() {}

type MessageActionContactSignUp struct{}

func (*MessageActionContactSignUp) CRC() uint32 {
	return 0xf3f25f76
}
func (*MessageActionContactSignUp) _IMessageAction() {}

type MessageActionGeoProximityReached struct {
	FromID   IPeer
	ToID     IPeer
	Distance int32
}

func (*MessageActionGeoProximityReached) CRC() uint32 {
	return 0x98e0d697
}
func (*MessageActionGeoProximityReached) _IMessageAction() {}

type MessageActionGroupCall struct {
	_        struct{} `tl:"flags,bitflag"`
	Call     IInputGroupCall
	Duration *int32 `tl:",omitempty:flags:0"`
}

func (*MessageActionGroupCall) CRC() uint32 {
	return 0x7a0d7f42
}
func (*MessageActionGroupCall) _IMessageAction() {}

type MessageActionInviteToGroupCall struct {
	Call  IInputGroupCall
	Users []int64
}

func (*MessageActionInviteToGroupCall) CRC() uint32 {
	return 0x502f92f7
}
func (*MessageActionInviteToGroupCall) _IMessageAction() {}

type MessageActionSetMessagesTTL struct {
	_               struct{} `tl:"flags,bitflag"`
	Period          int32
	AutoSettingFrom *int64 `tl:",omitempty:flags:0"`
}

func (*MessageActionSetMessagesTTL) CRC() uint32 {
	return 0x3c134d7b
}
func (*MessageActionSetMessagesTTL) _IMessageAction() {}

type MessageActionGroupCallScheduled struct {
	Call         IInputGroupCall
	ScheduleDate int32
}

func (*MessageActionGroupCallScheduled) CRC() uint32 {
	return 0xb3a07661
}
func (*MessageActionGroupCallScheduled) _IMessageAction() {}

type MessageActionSetChatTheme struct {
	Emoticon string
}

func (*MessageActionSetChatTheme) CRC() uint32 {
	return 0xaa786345
}
func (*MessageActionSetChatTheme) _IMessageAction() {}

type MessageActionChatJoinedByRequest struct{}

func (*MessageActionChatJoinedByRequest) CRC() uint32 {
	return 0xebbca3cb
}
func (*MessageActionChatJoinedByRequest) _IMessageAction() {}

type MessageActionWebViewDataSentMe struct {
	Text string
	Data string
}

func (*MessageActionWebViewDataSentMe) CRC() uint32 {
	return 0x47dd8079
}
func (*MessageActionWebViewDataSentMe) _IMessageAction() {}

type MessageActionWebViewDataSent struct {
	Text string
}

func (*MessageActionWebViewDataSent) CRC() uint32 {
	return 0xb4c38cb5
}
func (*MessageActionWebViewDataSent) _IMessageAction() {}

type MessageActionGiftPremium struct {
	_              struct{} `tl:"flags,bitflag"`
	Currency       string
	Amount         int64
	Months         int32
	CryptoCurrency *string `tl:",omitempty:flags:0"`
	CryptoAmount   *int64  `tl:",omitempty:flags:0"`
}

func (*MessageActionGiftPremium) CRC() uint32 {
	return 0xc83d6aec
}
func (*MessageActionGiftPremium) _IMessageAction() {}

type MessageActionTopicCreate struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string
	IconColor   int32
	IconEmojiID *int64 `tl:",omitempty:flags:0"`
}

func (*MessageActionTopicCreate) CRC() uint32 {
	return 0xd999256
}
func (*MessageActionTopicCreate) _IMessageAction() {}

type MessageActionTopicEdit struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       *string  `tl:",omitempty:flags:0"`
	IconEmojiID *int64   `tl:",omitempty:flags:1"`
	Closed      *bool    `tl:",omitempty:flags:2"`
	Hidden      *bool    `tl:",omitempty:flags:3"`
}

func (*MessageActionTopicEdit) CRC() uint32 {
	return 0xc0944820
}
func (*MessageActionTopicEdit) _IMessageAction() {}

type MessageActionSuggestProfilePhoto struct {
	Photo IPhoto
}

func (*MessageActionSuggestProfilePhoto) CRC() uint32 {
	return 0x57de635e
}
func (*MessageActionSuggestProfilePhoto) _IMessageAction() {}

type MessageActionRequestedPeer struct {
	ButtonID int32
	Peers    []IPeer
}

func (*MessageActionRequestedPeer) CRC() uint32 {
	return 0x31518e9b
}
func (*MessageActionRequestedPeer) _IMessageAction() {}

type MessageActionSetChatWallPaper struct {
	_         struct{} `tl:"flags,bitflag"`
	Same      bool     `tl:",omitempty:flags:0,implicit"`
	ForBoth   bool     `tl:",omitempty:flags:1,implicit"`
	Wallpaper IWallPaper
}

func (*MessageActionSetChatWallPaper) CRC() uint32 {
	return 0x5060a3f4
}
func (*MessageActionSetChatWallPaper) _IMessageAction() {}

type MessageActionGiftCode struct {
	_              struct{} `tl:"flags,bitflag"`
	ViaGiveaway    bool     `tl:",omitempty:flags:0,implicit"`
	Unclaimed      bool     `tl:",omitempty:flags:2,implicit"`
	BoostPeer      IPeer    `tl:",omitempty:flags:1"`
	Months         int32
	Slug           string
	Currency       *string `tl:",omitempty:flags:2"`
	Amount         *int64  `tl:",omitempty:flags:2"`
	CryptoCurrency *string `tl:",omitempty:flags:3"`
	CryptoAmount   *int64  `tl:",omitempty:flags:3"`
}

func (*MessageActionGiftCode) CRC() uint32 {
	return 0x678c2e09
}
func (*MessageActionGiftCode) _IMessageAction() {}

type MessageActionGiveawayLaunch struct{}

func (*MessageActionGiveawayLaunch) CRC() uint32 {
	return 0x332ba9ed
}
func (*MessageActionGiveawayLaunch) _IMessageAction() {}

type MessageActionGiveawayResults struct {
	WinnersCount   int32
	UnclaimedCount int32
}

func (*MessageActionGiveawayResults) CRC() uint32 {
	return 0x2a9fadc5
}
func (*MessageActionGiveawayResults) _IMessageAction() {}

type IMessageEntity interface {
	tl.Object
	_IMessageEntity()
}

var (
	_ IMessageEntity = (*MessageEntityUnknown)(nil)
	_ IMessageEntity = (*MessageEntityMention)(nil)
	_ IMessageEntity = (*MessageEntityHashtag)(nil)
	_ IMessageEntity = (*MessageEntityBotCommand)(nil)
	_ IMessageEntity = (*MessageEntityURL)(nil)
	_ IMessageEntity = (*MessageEntityEmail)(nil)
	_ IMessageEntity = (*MessageEntityBold)(nil)
	_ IMessageEntity = (*MessageEntityItalic)(nil)
	_ IMessageEntity = (*MessageEntityCode)(nil)
	_ IMessageEntity = (*MessageEntityPre)(nil)
	_ IMessageEntity = (*MessageEntityTextURL)(nil)
	_ IMessageEntity = (*MessageEntityMentionName)(nil)
	_ IMessageEntity = (*InputMessageEntityMentionName)(nil)
	_ IMessageEntity = (*MessageEntityPhone)(nil)
	_ IMessageEntity = (*MessageEntityCashtag)(nil)
	_ IMessageEntity = (*MessageEntityUnderline)(nil)
	_ IMessageEntity = (*MessageEntityStrike)(nil)
	_ IMessageEntity = (*MessageEntityBankCard)(nil)
	_ IMessageEntity = (*MessageEntitySpoiler)(nil)
	_ IMessageEntity = (*MessageEntityCustomEmoji)(nil)
	_ IMessageEntity = (*MessageEntityBlockquote)(nil)
)

type MessageEntityUnknown struct {
	Offset int32
	Length int32
}

func (*MessageEntityUnknown) CRC() uint32 {
	return 0xbb92ba95
}
func (*MessageEntityUnknown) _IMessageEntity() {}

type MessageEntityMention struct {
	Offset int32
	Length int32
}

func (*MessageEntityMention) CRC() uint32 {
	return 0xfa04579d
}
func (*MessageEntityMention) _IMessageEntity() {}

type MessageEntityHashtag struct {
	Offset int32
	Length int32
}

func (*MessageEntityHashtag) CRC() uint32 {
	return 0x6f635b0d
}
func (*MessageEntityHashtag) _IMessageEntity() {}

type MessageEntityBotCommand struct {
	Offset int32
	Length int32
}

func (*MessageEntityBotCommand) CRC() uint32 {
	return 0x6cef8ac7
}
func (*MessageEntityBotCommand) _IMessageEntity() {}

type MessageEntityURL struct {
	Offset int32
	Length int32
}

func (*MessageEntityURL) CRC() uint32 {
	return 0x6ed02538
}
func (*MessageEntityURL) _IMessageEntity() {}

type MessageEntityEmail struct {
	Offset int32
	Length int32
}

func (*MessageEntityEmail) CRC() uint32 {
	return 0x64e475c2
}
func (*MessageEntityEmail) _IMessageEntity() {}

type MessageEntityBold struct {
	Offset int32
	Length int32
}

func (*MessageEntityBold) CRC() uint32 {
	return 0xbd610bc9
}
func (*MessageEntityBold) _IMessageEntity() {}

type MessageEntityItalic struct {
	Offset int32
	Length int32
}

func (*MessageEntityItalic) CRC() uint32 {
	return 0x826f8b60
}
func (*MessageEntityItalic) _IMessageEntity() {}

type MessageEntityCode struct {
	Offset int32
	Length int32
}

func (*MessageEntityCode) CRC() uint32 {
	return 0x28a20571
}
func (*MessageEntityCode) _IMessageEntity() {}

type MessageEntityPre struct {
	Offset   int32
	Length   int32
	Language string
}

func (*MessageEntityPre) CRC() uint32 {
	return 0x73924be0
}
func (*MessageEntityPre) _IMessageEntity() {}

type MessageEntityTextURL struct {
	Offset int32
	Length int32
	URL    string
}

func (*MessageEntityTextURL) CRC() uint32 {
	return 0x76a6d327
}
func (*MessageEntityTextURL) _IMessageEntity() {}

type MessageEntityMentionName struct {
	Offset int32
	Length int32
	UserID int64
}

func (*MessageEntityMentionName) CRC() uint32 {
	return 0xdc7b1140
}
func (*MessageEntityMentionName) _IMessageEntity() {}

type InputMessageEntityMentionName struct {
	Offset int32
	Length int32
	UserID IInputUser
}

func (*InputMessageEntityMentionName) CRC() uint32 {
	return 0x208e68c9
}
func (*InputMessageEntityMentionName) _IMessageEntity() {}

type MessageEntityPhone struct {
	Offset int32
	Length int32
}

func (*MessageEntityPhone) CRC() uint32 {
	return 0x9b69e34b
}
func (*MessageEntityPhone) _IMessageEntity() {}

type MessageEntityCashtag struct {
	Offset int32
	Length int32
}

func (*MessageEntityCashtag) CRC() uint32 {
	return 0x4c4e743f
}
func (*MessageEntityCashtag) _IMessageEntity() {}

type MessageEntityUnderline struct {
	Offset int32
	Length int32
}

func (*MessageEntityUnderline) CRC() uint32 {
	return 0x9c4e7e8b
}
func (*MessageEntityUnderline) _IMessageEntity() {}

type MessageEntityStrike struct {
	Offset int32
	Length int32
}

func (*MessageEntityStrike) CRC() uint32 {
	return 0xbf0693d4
}
func (*MessageEntityStrike) _IMessageEntity() {}

type MessageEntityBankCard struct {
	Offset int32
	Length int32
}

func (*MessageEntityBankCard) CRC() uint32 {
	return 0x761e6af4
}
func (*MessageEntityBankCard) _IMessageEntity() {}

type MessageEntitySpoiler struct {
	Offset int32
	Length int32
}

func (*MessageEntitySpoiler) CRC() uint32 {
	return 0x32ca960f
}
func (*MessageEntitySpoiler) _IMessageEntity() {}

type MessageEntityCustomEmoji struct {
	Offset     int32
	Length     int32
	DocumentID int64
}

func (*MessageEntityCustomEmoji) CRC() uint32 {
	return 0xc8cf05f8
}
func (*MessageEntityCustomEmoji) _IMessageEntity() {}

type MessageEntityBlockquote struct {
	Offset int32
	Length int32
}

func (*MessageEntityBlockquote) CRC() uint32 {
	return 0x20df5d0
}
func (*MessageEntityBlockquote) _IMessageEntity() {}

type IMessageExtendedMedia interface {
	tl.Object
	_IMessageExtendedMedia()
}

var (
	_ IMessageExtendedMedia = (*MessageExtendedMediaPreview)(nil)
	_ IMessageExtendedMedia = (*MessageExtendedMedia)(nil)
)

type MessageExtendedMediaPreview struct {
	_             struct{}   `tl:"flags,bitflag"`
	W             *int32     `tl:",omitempty:flags:0"`
	H             *int32     `tl:",omitempty:flags:0"`
	Thumb         IPhotoSize `tl:",omitempty:flags:1"`
	VideoDuration *int32     `tl:",omitempty:flags:2"`
}

func (*MessageExtendedMediaPreview) CRC() uint32 {
	return 0xad628cc8
}
func (*MessageExtendedMediaPreview) _IMessageExtendedMedia() {}

type MessageExtendedMedia struct {
	Media IMessageMedia
}

func (*MessageExtendedMedia) CRC() uint32 {
	return 0xee479c64
}
func (*MessageExtendedMedia) _IMessageExtendedMedia() {}

type IMessageFwdHeader interface {
	tl.Object
	_IMessageFwdHeader()
}

var (
	_ IMessageFwdHeader = (*MessageFwdHeader)(nil)
)

type MessageFwdHeader struct {
	_              struct{} `tl:"flags,bitflag"`
	Imported       bool     `tl:",omitempty:flags:7,implicit"`
	FromID         IPeer    `tl:",omitempty:flags:0"`
	FromName       *string  `tl:",omitempty:flags:5"`
	Date           int32
	ChannelPost    *int32  `tl:",omitempty:flags:2"`
	PostAuthor     *string `tl:",omitempty:flags:3"`
	SavedFromPeer  IPeer   `tl:",omitempty:flags:4"`
	SavedFromMsgID *int32  `tl:",omitempty:flags:4"`
	PsaType        *string `tl:",omitempty:flags:6"`
}

func (*MessageFwdHeader) CRC() uint32 {
	return 0x5f777dce
}
func (*MessageFwdHeader) _IMessageFwdHeader() {}

type IMessageMedia interface {
	tl.Object
	_IMessageMedia()
}

var (
	_ IMessageMedia = (*MessageMediaEmpty)(nil)
	_ IMessageMedia = (*MessageMediaPhoto)(nil)
	_ IMessageMedia = (*MessageMediaGeo)(nil)
	_ IMessageMedia = (*MessageMediaContact)(nil)
	_ IMessageMedia = (*MessageMediaUnsupported)(nil)
	_ IMessageMedia = (*MessageMediaDocument)(nil)
	_ IMessageMedia = (*MessageMediaWebPage)(nil)
	_ IMessageMedia = (*MessageMediaVenue)(nil)
	_ IMessageMedia = (*MessageMediaGame)(nil)
	_ IMessageMedia = (*MessageMediaInvoice)(nil)
	_ IMessageMedia = (*MessageMediaGeoLive)(nil)
	_ IMessageMedia = (*MessageMediaPoll)(nil)
	_ IMessageMedia = (*MessageMediaDice)(nil)
	_ IMessageMedia = (*MessageMediaStory)(nil)
	_ IMessageMedia = (*MessageMediaGiveaway)(nil)
	_ IMessageMedia = (*MessageMediaGiveawayResults)(nil)
)

type MessageMediaEmpty struct{}

func (*MessageMediaEmpty) CRC() uint32 {
	return 0x3ded6320
}
func (*MessageMediaEmpty) _IMessageMedia() {}

type MessageMediaPhoto struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:3,implicit"`
	Photo      IPhoto   `tl:",omitempty:flags:0"`
	TTLSeconds *int32   `tl:",omitempty:flags:2"`
}

func (*MessageMediaPhoto) CRC() uint32 {
	return 0x695150d7
}
func (*MessageMediaPhoto) _IMessageMedia() {}

type MessageMediaGeo struct {
	Geo IGeoPoint
}

func (*MessageMediaGeo) CRC() uint32 {
	return 0x56e0d474
}
func (*MessageMediaGeo) _IMessageMedia() {}

type MessageMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	UserID      int64
}

func (*MessageMediaContact) CRC() uint32 {
	return 0x70322949
}
func (*MessageMediaContact) _IMessageMedia() {}

type MessageMediaUnsupported struct{}

func (*MessageMediaUnsupported) CRC() uint32 {
	return 0x9f84f49e
}
func (*MessageMediaUnsupported) _IMessageMedia() {}

type MessageMediaDocument struct {
	_           struct{}  `tl:"flags,bitflag"`
	Nopremium   bool      `tl:",omitempty:flags:3,implicit"`
	Spoiler     bool      `tl:",omitempty:flags:4,implicit"`
	Document    IDocument `tl:",omitempty:flags:0"`
	AltDocument IDocument `tl:",omitempty:flags:5"`
	TTLSeconds  *int32    `tl:",omitempty:flags:2"`
}

func (*MessageMediaDocument) CRC() uint32 {
	return 0x4cf4d72d
}
func (*MessageMediaDocument) _IMessageMedia() {}

type MessageMediaWebPage struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:",omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:1,implicit"`
	Manual          bool     `tl:",omitempty:flags:3,implicit"`
	Safe            bool     `tl:",omitempty:flags:4,implicit"`
	Webpage         IWebPage
}

func (*MessageMediaWebPage) CRC() uint32 {
	return 0xddf10c3b
}
func (*MessageMediaWebPage) _IMessageMedia() {}

type MessageMediaVenue struct {
	Geo       IGeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

func (*MessageMediaVenue) CRC() uint32 {
	return 0x2ec0533f
}
func (*MessageMediaVenue) _IMessageMedia() {}

type MessageMediaGame struct {
	Game IGame
}

func (*MessageMediaGame) CRC() uint32 {
	return 0xfdb19008
}
func (*MessageMediaGame) _IMessageMedia() {}

type MessageMediaInvoice struct {
	_                        struct{} `tl:"flags,bitflag"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:1,implicit"`
	Test                     bool     `tl:",omitempty:flags:3,implicit"`
	Title                    string
	Description              string
	Photo                    IWebDocument `tl:",omitempty:flags:0"`
	ReceiptMsgID             *int32       `tl:",omitempty:flags:2"`
	Currency                 string
	TotalAmount              int64
	StartParam               string
	ExtendedMedia            IMessageExtendedMedia `tl:",omitempty:flags:4"`
}

func (*MessageMediaInvoice) CRC() uint32 {
	return 0xf6a548d3
}
func (*MessageMediaInvoice) _IMessageMedia() {}

type MessageMediaGeoLive struct {
	_                           struct{} `tl:"flags,bitflag"`
	Geo                         IGeoPoint
	Heading                     *int32 `tl:",omitempty:flags:0"`
	Period                      int32
	ProximityNotificationRadius *int32 `tl:",omitempty:flags:1"`
}

func (*MessageMediaGeoLive) CRC() uint32 {
	return 0xb940c666
}
func (*MessageMediaGeoLive) _IMessageMedia() {}

type MessageMediaPoll struct {
	Poll    IPoll
	Results IPollResults
}

func (*MessageMediaPoll) CRC() uint32 {
	return 0x4bd6e798
}
func (*MessageMediaPoll) _IMessageMedia() {}

type MessageMediaDice struct {
	Value    int32
	Emoticon string
}

func (*MessageMediaDice) CRC() uint32 {
	return 0x3f7ee58b
}
func (*MessageMediaDice) _IMessageMedia() {}

type MessageMediaStory struct {
	_          struct{} `tl:"flags,bitflag"`
	ViaMention bool     `tl:",omitempty:flags:1,implicit"`
	Peer       IPeer
	ID         int32
	Story      IStoryItem `tl:",omitempty:flags:0"`
}

func (*MessageMediaStory) CRC() uint32 {
	return 0x68cb6283
}
func (*MessageMediaStory) _IMessageMedia() {}

type MessageMediaGiveaway struct {
	_                  struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers bool     `tl:",omitempty:flags:0,implicit"`
	WinnersAreVisible  bool     `tl:",omitempty:flags:2,implicit"`
	Channels           []int64
	CountriesIso2      []string `tl:",omitempty:flags:1"`
	PrizeDescription   *string  `tl:",omitempty:flags:3"`
	Quantity           int32
	Months             int32
	UntilDate          int32
}

func (*MessageMediaGiveaway) CRC() uint32 {
	return 0xdaad85b0
}
func (*MessageMediaGiveaway) _IMessageMedia() {}

type MessageMediaGiveawayResults struct {
	_                    struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers   bool     `tl:",omitempty:flags:0,implicit"`
	Refunded             bool     `tl:",omitempty:flags:2,implicit"`
	ChannelID            int64
	AdditionalPeersCount *int32 `tl:",omitempty:flags:3"`
	LaunchMsgID          int32
	WinnersCount         int32
	UnclaimedCount       int32
	Winners              []int64
	Months               int32
	PrizeDescription     *string `tl:",omitempty:flags:1"`
	UntilDate            int32
}

func (*MessageMediaGiveawayResults) CRC() uint32 {
	return 0xc6991068
}
func (*MessageMediaGiveawayResults) _IMessageMedia() {}

type IMessagePeerReaction interface {
	tl.Object
	_IMessagePeerReaction()
}

var (
	_ IMessagePeerReaction = (*MessagePeerReaction)(nil)
)

type MessagePeerReaction struct {
	_        struct{} `tl:"flags,bitflag"`
	Big      bool     `tl:",omitempty:flags:0,implicit"`
	Unread   bool     `tl:",omitempty:flags:1,implicit"`
	My       bool     `tl:",omitempty:flags:2,implicit"`
	PeerID   IPeer
	Date     int32
	Reaction IReaction
}

func (*MessagePeerReaction) CRC() uint32 {
	return 0x8c79b63c
}
func (*MessagePeerReaction) _IMessagePeerReaction() {}

type IMessagePeerVote interface {
	tl.Object
	_IMessagePeerVote()
}

var (
	_ IMessagePeerVote = (*MessagePeerVote)(nil)
	_ IMessagePeerVote = (*MessagePeerVoteInputOption)(nil)
	_ IMessagePeerVote = (*MessagePeerVoteMultiple)(nil)
)

type MessagePeerVote struct {
	Peer   IPeer
	Option []byte
	Date   int32
}

func (*MessagePeerVote) CRC() uint32 {
	return 0xb6cc2d5c
}
func (*MessagePeerVote) _IMessagePeerVote() {}

type MessagePeerVoteInputOption struct {
	Peer IPeer
	Date int32
}

func (*MessagePeerVoteInputOption) CRC() uint32 {
	return 0x74cda504
}
func (*MessagePeerVoteInputOption) _IMessagePeerVote() {}

type MessagePeerVoteMultiple struct {
	Peer    IPeer
	Options [][]byte
	Date    int32
}

func (*MessagePeerVoteMultiple) CRC() uint32 {
	return 0x4628f6e6
}
func (*MessagePeerVoteMultiple) _IMessagePeerVote() {}

type IMessageRange interface {
	tl.Object
	_IMessageRange()
}

var (
	_ IMessageRange = (*MessageRange)(nil)
)

type MessageRange struct {
	MinID int32
	MaxID int32
}

func (*MessageRange) CRC() uint32 {
	return 0xae30253
}
func (*MessageRange) _IMessageRange() {}

type IMessageReactions interface {
	tl.Object
	_IMessageReactions()
}

var (
	_ IMessageReactions = (*MessageReactions)(nil)
)

type MessageReactions struct {
	_               struct{} `tl:"flags,bitflag"`
	Min             bool     `tl:",omitempty:flags:0,implicit"`
	CanSeeList      bool     `tl:",omitempty:flags:2,implicit"`
	Results         []IReactionCount
	RecentReactions []IMessagePeerReaction `tl:",omitempty:flags:1"`
}

func (*MessageReactions) CRC() uint32 {
	return 0x4f2b9479
}
func (*MessageReactions) _IMessageReactions() {}

type IMessageReplies interface {
	tl.Object
	_IMessageReplies()
}

var (
	_ IMessageReplies = (*MessageReplies)(nil)
)

type MessageReplies struct {
	_              struct{} `tl:"flags,bitflag"`
	Comments       bool     `tl:",omitempty:flags:0,implicit"`
	Replies        int32
	RepliesPts     int32
	RecentRepliers []IPeer `tl:",omitempty:flags:1"`
	ChannelID      *int64  `tl:",omitempty:flags:0"`
	MaxID          *int32  `tl:",omitempty:flags:2"`
	ReadMaxID      *int32  `tl:",omitempty:flags:3"`
}

func (*MessageReplies) CRC() uint32 {
	return 0x83d60fc2
}
func (*MessageReplies) _IMessageReplies() {}

type IMessageReplyHeader interface {
	tl.Object
	_IMessageReplyHeader()
}

var (
	_ IMessageReplyHeader = (*MessageReplyHeader)(nil)
	_ IMessageReplyHeader = (*MessageReplyStoryHeader)(nil)
)

type MessageReplyHeader struct {
	_                struct{}          `tl:"flags,bitflag"`
	ReplyToScheduled bool              `tl:",omitempty:flags:2,implicit"`
	ForumTopic       bool              `tl:",omitempty:flags:3,implicit"`
	Quote            bool              `tl:",omitempty:flags:9,implicit"`
	ReplyToMsgID     *int32            `tl:",omitempty:flags:4"`
	ReplyToPeerID    IPeer             `tl:",omitempty:flags:0"`
	ReplyFrom        IMessageFwdHeader `tl:",omitempty:flags:5"`
	ReplyMedia       IMessageMedia     `tl:",omitempty:flags:8"`
	ReplyToTopID     *int32            `tl:",omitempty:flags:1"`
	QuoteText        *string           `tl:",omitempty:flags:6"`
	QuoteEntities    []IMessageEntity  `tl:",omitempty:flags:7"`
	QuoteOffset      *int32            `tl:",omitempty:flags:10"`
}

func (*MessageReplyHeader) CRC() uint32 {
	return 0xafbc09db
}
func (*MessageReplyHeader) _IMessageReplyHeader() {}

type MessageReplyStoryHeader struct {
	UserID  int64
	StoryID int32
}

func (*MessageReplyStoryHeader) CRC() uint32 {
	return 0x9c98bfc1
}
func (*MessageReplyStoryHeader) _IMessageReplyHeader() {}

type IMessageViews interface {
	tl.Object
	_IMessageViews()
}

var (
	_ IMessageViews = (*MessageViews)(nil)
)

type MessageViews struct {
	_        struct{}        `tl:"flags,bitflag"`
	Views    *int32          `tl:",omitempty:flags:0"`
	Forwards *int32          `tl:",omitempty:flags:1"`
	Replies  IMessageReplies `tl:",omitempty:flags:2"`
}

func (*MessageViews) CRC() uint32 {
	return 0x455b853d
}
func (*MessageViews) _IMessageViews() {}

// Object describes message filter.
type IMessagesFilter interface {
	tl.Object
	_IMessagesFilter()
}

var (
	_ IMessagesFilter = (*InputMessagesFilterEmpty)(nil)
	_ IMessagesFilter = (*InputMessagesFilterPhotos)(nil)
	_ IMessagesFilter = (*InputMessagesFilterVideo)(nil)
	_ IMessagesFilter = (*InputMessagesFilterPhotoVideo)(nil)
	_ IMessagesFilter = (*InputMessagesFilterDocument)(nil)
	_ IMessagesFilter = (*InputMessagesFilterURL)(nil)
	_ IMessagesFilter = (*InputMessagesFilterGif)(nil)
	_ IMessagesFilter = (*InputMessagesFilterVoice)(nil)
	_ IMessagesFilter = (*InputMessagesFilterMusic)(nil)
	_ IMessagesFilter = (*InputMessagesFilterChatPhotos)(nil)
	_ IMessagesFilter = (*InputMessagesFilterPhoneCalls)(nil)
	_ IMessagesFilter = (*InputMessagesFilterRoundVoice)(nil)
	_ IMessagesFilter = (*InputMessagesFilterRoundVideo)(nil)
	_ IMessagesFilter = (*InputMessagesFilterMyMentions)(nil)
	_ IMessagesFilter = (*InputMessagesFilterGeo)(nil)
	_ IMessagesFilter = (*InputMessagesFilterContacts)(nil)
	_ IMessagesFilter = (*InputMessagesFilterPinned)(nil)
)

// Filter is absent.
type InputMessagesFilterEmpty struct{}

func (*InputMessagesFilterEmpty) CRC() uint32 {
	return 0x57e2f66c
}
func (*InputMessagesFilterEmpty) _IMessagesFilter() {}

// Filter for messages containing photos.
type InputMessagesFilterPhotos struct{}

func (*InputMessagesFilterPhotos) CRC() uint32 {
	return 0x9609a51c
}
func (*InputMessagesFilterPhotos) _IMessagesFilter() {}

// Filter for messages containing videos.
type InputMessagesFilterVideo struct{}

func (*InputMessagesFilterVideo) CRC() uint32 {
	return 0x9fc00e65
}
func (*InputMessagesFilterVideo) _IMessagesFilter() {}

// Filter for messages containing photos or videos.
type InputMessagesFilterPhotoVideo struct{}

func (*InputMessagesFilterPhotoVideo) CRC() uint32 {
	return 0x56e9f0e4
}
func (*InputMessagesFilterPhotoVideo) _IMessagesFilter() {}

// Filter for messages containing documents.
type InputMessagesFilterDocument struct{}

func (*InputMessagesFilterDocument) CRC() uint32 {
	return 0x9eddf188
}
func (*InputMessagesFilterDocument) _IMessagesFilter() {}

// Return only messages containing URLs
type InputMessagesFilterURL struct{}

func (*InputMessagesFilterURL) CRC() uint32 {
	return 0x7ef0dd87
}
func (*InputMessagesFilterURL) _IMessagesFilter() {}

// Return only messages containing gifs
type InputMessagesFilterGif struct{}

func (*InputMessagesFilterGif) CRC() uint32 {
	return 0xffc86587
}
func (*InputMessagesFilterGif) _IMessagesFilter() {}

// Return only messages containing voice notes
type InputMessagesFilterVoice struct{}

func (*InputMessagesFilterVoice) CRC() uint32 {
	return 0x50f5c392
}
func (*InputMessagesFilterVoice) _IMessagesFilter() {}

// Return only messages containing audio files
type InputMessagesFilterMusic struct{}

func (*InputMessagesFilterMusic) CRC() uint32 {
	return 0x3751b49e
}
func (*InputMessagesFilterMusic) _IMessagesFilter() {}

// Return only chat photo changes
type InputMessagesFilterChatPhotos struct{}

func (*InputMessagesFilterChatPhotos) CRC() uint32 {
	return 0x3a20ecb8
}
func (*InputMessagesFilterChatPhotos) _IMessagesFilter() {}

// Return only phone calls
type InputMessagesFilterPhoneCalls struct {
	_      struct{} `tl:"flags,bitflag"`
	Missed bool     `tl:",omitempty:flags:0,implicit"` // Return only missed phone calls
}

func (*InputMessagesFilterPhoneCalls) CRC() uint32 {
	return 0x80c99768
}
func (*InputMessagesFilterPhoneCalls) _IMessagesFilter() {}

// Return only round videos and voice notes
type InputMessagesFilterRoundVoice struct{}

func (*InputMessagesFilterRoundVoice) CRC() uint32 {
	return 0x7a7c17a4
}
func (*InputMessagesFilterRoundVoice) _IMessagesFilter() {}

// Return only round videos
type InputMessagesFilterRoundVideo struct{}

func (*InputMessagesFilterRoundVideo) CRC() uint32 {
	return 0xb549da53
}
func (*InputMessagesFilterRoundVideo) _IMessagesFilter() {}

// Return only messages where the current user was mentioned
type InputMessagesFilterMyMentions struct{}

func (*InputMessagesFilterMyMentions) CRC() uint32 {
	return 0xc1f8e69a
}
func (*InputMessagesFilterMyMentions) _IMessagesFilter() {}

// Return only messages containing geolocations
type InputMessagesFilterGeo struct{}

func (*InputMessagesFilterGeo) CRC() uint32 {
	return 0xe7026d0d
}
func (*InputMessagesFilterGeo) _IMessagesFilter() {}

// Return only messages containing contacts
type InputMessagesFilterContacts struct{}

func (*InputMessagesFilterContacts) CRC() uint32 {
	return 0xe062db83
}
func (*InputMessagesFilterContacts) _IMessagesFilter() {}

// Fetch only pinned messages
type InputMessagesFilterPinned struct{}

func (*InputMessagesFilterPinned) CRC() uint32 {
	return 0x1bb00451
}
func (*InputMessagesFilterPinned) _IMessagesFilter() {}

type IMyBoost interface {
	tl.Object
	_IMyBoost()
}

var (
	_ IMyBoost = (*MyBoost)(nil)
)

type MyBoost struct {
	_                 struct{} `tl:"flags,bitflag"`
	Slot              int32
	Peer              IPeer `tl:",omitempty:flags:0"`
	Date              int32
	Expires           int32
	CooldownUntilDate *int32 `tl:",omitempty:flags:1"`
}

func (*MyBoost) CRC() uint32 {
	return 0xc448415c
}
func (*MyBoost) _IMyBoost() {}

type INearestDc interface {
	tl.Object
	_INearestDc()
}

var (
	_ INearestDc = (*NearestDc)(nil)
)

type NearestDc struct {
	Country   string
	ThisDc    int32
	NearestDc int32
}

func (*NearestDc) CRC() uint32 {
	return 0x8e1a1775
}
func (*NearestDc) _INearestDc() {}

type INotificationSound interface {
	tl.Object
	_INotificationSound()
}

var (
	_ INotificationSound = (*NotificationSoundDefault)(nil)
	_ INotificationSound = (*NotificationSoundNone)(nil)
	_ INotificationSound = (*NotificationSoundLocal)(nil)
	_ INotificationSound = (*NotificationSoundRingtone)(nil)
)

type NotificationSoundDefault struct{}

func (*NotificationSoundDefault) CRC() uint32 {
	return 0x97e8bebe
}
func (*NotificationSoundDefault) _INotificationSound() {}

type NotificationSoundNone struct{}

func (*NotificationSoundNone) CRC() uint32 {
	return 0x6f0c34df
}
func (*NotificationSoundNone) _INotificationSound() {}

type NotificationSoundLocal struct {
	Title string
	Data  string
}

func (*NotificationSoundLocal) CRC() uint32 {
	return 0x830b9ae4
}
func (*NotificationSoundLocal) _INotificationSound() {}

type NotificationSoundRingtone struct {
	ID int64
}

func (*NotificationSoundRingtone) CRC() uint32 {
	return 0xff6c8049
}
func (*NotificationSoundRingtone) _INotificationSound() {}

type INotifyPeer interface {
	tl.Object
	_INotifyPeer()
}

var (
	_ INotifyPeer = (*NotifyPeer)(nil)
	_ INotifyPeer = (*NotifyUsers)(nil)
	_ INotifyPeer = (*NotifyChats)(nil)
	_ INotifyPeer = (*NotifyBroadcasts)(nil)
	_ INotifyPeer = (*NotifyForumTopic)(nil)
)

type NotifyPeer struct {
	Peer IPeer
}

func (*NotifyPeer) CRC() uint32 {
	return 0x9fd40bd8
}
func (*NotifyPeer) _INotifyPeer() {}

type NotifyUsers struct{}

func (*NotifyUsers) CRC() uint32 {
	return 0xb4c83b4c
}
func (*NotifyUsers) _INotifyPeer() {}

type NotifyChats struct{}

func (*NotifyChats) CRC() uint32 {
	return 0xc007cec3
}
func (*NotifyChats) _INotifyPeer() {}

type NotifyBroadcasts struct{}

func (*NotifyBroadcasts) CRC() uint32 {
	return 0xd612e8ef
}
func (*NotifyBroadcasts) _INotifyPeer() {}

type NotifyForumTopic struct {
	Peer     IPeer
	TopMsgID int32
}

func (*NotifyForumTopic) CRC() uint32 {
	return 0x226e6308
}
func (*NotifyForumTopic) _INotifyPeer() {}

type IPage interface {
	tl.Object
	_IPage()
}

var (
	_ IPage = (*Page)(nil)
)

type Page struct {
	_         struct{} `tl:"flags,bitflag"`
	Part      bool     `tl:",omitempty:flags:0,implicit"`
	Rtl       bool     `tl:",omitempty:flags:1,implicit"`
	V2        bool     `tl:",omitempty:flags:2,implicit"`
	URL       string
	Blocks    []IPageBlock
	Photos    []IPhoto
	Documents []IDocument
	Views     *int32 `tl:",omitempty:flags:3"`
}

func (*Page) CRC() uint32 {
	return 0x98657f0d
}
func (*Page) _IPage() {}

type IPageBlock interface {
	tl.Object
	_IPageBlock()
}

var (
	_ IPageBlock = (*PageBlockUnsupported)(nil)
	_ IPageBlock = (*PageBlockTitle)(nil)
	_ IPageBlock = (*PageBlockSubtitle)(nil)
	_ IPageBlock = (*PageBlockAuthorDate)(nil)
	_ IPageBlock = (*PageBlockHeader)(nil)
	_ IPageBlock = (*PageBlockSubheader)(nil)
	_ IPageBlock = (*PageBlockParagraph)(nil)
	_ IPageBlock = (*PageBlockPreformatted)(nil)
	_ IPageBlock = (*PageBlockFooter)(nil)
	_ IPageBlock = (*PageBlockDivider)(nil)
	_ IPageBlock = (*PageBlockAnchor)(nil)
	_ IPageBlock = (*PageBlockList)(nil)
	_ IPageBlock = (*PageBlockBlockquote)(nil)
	_ IPageBlock = (*PageBlockPullquote)(nil)
	_ IPageBlock = (*PageBlockPhoto)(nil)
	_ IPageBlock = (*PageBlockVideo)(nil)
	_ IPageBlock = (*PageBlockCover)(nil)
	_ IPageBlock = (*PageBlockEmbed)(nil)
	_ IPageBlock = (*PageBlockEmbedPost)(nil)
	_ IPageBlock = (*PageBlockCollage)(nil)
	_ IPageBlock = (*PageBlockSlideshow)(nil)
	_ IPageBlock = (*PageBlockChannel)(nil)
	_ IPageBlock = (*PageBlockAudio)(nil)
	_ IPageBlock = (*PageBlockKicker)(nil)
	_ IPageBlock = (*PageBlockTable)(nil)
	_ IPageBlock = (*PageBlockOrderedList)(nil)
	_ IPageBlock = (*PageBlockDetails)(nil)
	_ IPageBlock = (*PageBlockRelatedArticles)(nil)
	_ IPageBlock = (*PageBlockMap)(nil)
)

type PageBlockUnsupported struct{}

func (*PageBlockUnsupported) CRC() uint32 {
	return 0x13567e8a
}
func (*PageBlockUnsupported) _IPageBlock() {}

type PageBlockTitle struct {
	Text IRichText
}

func (*PageBlockTitle) CRC() uint32 {
	return 0x70abc3fd
}
func (*PageBlockTitle) _IPageBlock() {}

type PageBlockSubtitle struct {
	Text IRichText
}

func (*PageBlockSubtitle) CRC() uint32 {
	return 0x8ffa9a1f
}
func (*PageBlockSubtitle) _IPageBlock() {}

type PageBlockAuthorDate struct {
	Author        IRichText
	PublishedDate int32
}

func (*PageBlockAuthorDate) CRC() uint32 {
	return 0xbaafe5e0
}
func (*PageBlockAuthorDate) _IPageBlock() {}

type PageBlockHeader struct {
	Text IRichText
}

func (*PageBlockHeader) CRC() uint32 {
	return 0xbfd064ec
}
func (*PageBlockHeader) _IPageBlock() {}

type PageBlockSubheader struct {
	Text IRichText
}

func (*PageBlockSubheader) CRC() uint32 {
	return 0xf12bb6e1
}
func (*PageBlockSubheader) _IPageBlock() {}

type PageBlockParagraph struct {
	Text IRichText
}

func (*PageBlockParagraph) CRC() uint32 {
	return 0x467a0766
}
func (*PageBlockParagraph) _IPageBlock() {}

type PageBlockPreformatted struct {
	Text     IRichText
	Language string
}

func (*PageBlockPreformatted) CRC() uint32 {
	return 0xc070d93e
}
func (*PageBlockPreformatted) _IPageBlock() {}

type PageBlockFooter struct {
	Text IRichText
}

func (*PageBlockFooter) CRC() uint32 {
	return 0x48870999
}
func (*PageBlockFooter) _IPageBlock() {}

type PageBlockDivider struct{}

func (*PageBlockDivider) CRC() uint32 {
	return 0xdb20b188
}
func (*PageBlockDivider) _IPageBlock() {}

type PageBlockAnchor struct {
	Name string
}

func (*PageBlockAnchor) CRC() uint32 {
	return 0xce0d37b0
}
func (*PageBlockAnchor) _IPageBlock() {}

type PageBlockList struct {
	Items []IPageListItem
}

func (*PageBlockList) CRC() uint32 {
	return 0xe4e88011
}
func (*PageBlockList) _IPageBlock() {}

type PageBlockBlockquote struct {
	Text    IRichText
	Caption IRichText
}

func (*PageBlockBlockquote) CRC() uint32 {
	return 0x263d7c26
}
func (*PageBlockBlockquote) _IPageBlock() {}

type PageBlockPullquote struct {
	Text    IRichText
	Caption IRichText
}

func (*PageBlockPullquote) CRC() uint32 {
	return 0x4f4456d3
}
func (*PageBlockPullquote) _IPageBlock() {}

type PageBlockPhoto struct {
	_         struct{} `tl:"flags,bitflag"`
	PhotoID   int64
	Caption   IPageCaption
	URL       *string `tl:",omitempty:flags:0"`
	WebpageID *int64  `tl:",omitempty:flags:0"`
}

func (*PageBlockPhoto) CRC() uint32 {
	return 0x1759c560
}
func (*PageBlockPhoto) _IPageBlock() {}

type PageBlockVideo struct {
	_        struct{} `tl:"flags,bitflag"`
	Autoplay bool     `tl:",omitempty:flags:0,implicit"`
	Loop     bool     `tl:",omitempty:flags:1,implicit"`
	VideoID  int64
	Caption  IPageCaption
}

func (*PageBlockVideo) CRC() uint32 {
	return 0x7c8fe7b6
}
func (*PageBlockVideo) _IPageBlock() {}

type PageBlockCover struct {
	Cover IPageBlock
}

func (*PageBlockCover) CRC() uint32 {
	return 0x39f23300
}
func (*PageBlockCover) _IPageBlock() {}

type PageBlockEmbed struct {
	_              struct{} `tl:"flags,bitflag"`
	FullWidth      bool     `tl:",omitempty:flags:0,implicit"`
	AllowScrolling bool     `tl:",omitempty:flags:3,implicit"`
	URL            *string  `tl:",omitempty:flags:1"`
	Html           *string  `tl:",omitempty:flags:2"`
	PosterPhotoID  *int64   `tl:",omitempty:flags:4"`
	W              *int32   `tl:",omitempty:flags:5"`
	H              *int32   `tl:",omitempty:flags:5"`
	Caption        IPageCaption
}

func (*PageBlockEmbed) CRC() uint32 {
	return 0xa8718dc5
}
func (*PageBlockEmbed) _IPageBlock() {}

type PageBlockEmbedPost struct {
	URL           string
	WebpageID     int64
	AuthorPhotoID int64
	Author        string
	Date          int32
	Blocks        []IPageBlock
	Caption       IPageCaption
}

func (*PageBlockEmbedPost) CRC() uint32 {
	return 0xf259a80b
}
func (*PageBlockEmbedPost) _IPageBlock() {}

type PageBlockCollage struct {
	Items   []IPageBlock
	Caption IPageCaption
}

func (*PageBlockCollage) CRC() uint32 {
	return 0x65a0fa4d
}
func (*PageBlockCollage) _IPageBlock() {}

type PageBlockSlideshow struct {
	Items   []IPageBlock
	Caption IPageCaption
}

func (*PageBlockSlideshow) CRC() uint32 {
	return 0x31f9590
}
func (*PageBlockSlideshow) _IPageBlock() {}

type PageBlockChannel struct {
	Channel IChat
}

func (*PageBlockChannel) CRC() uint32 {
	return 0xef1751b5
}
func (*PageBlockChannel) _IPageBlock() {}

type PageBlockAudio struct {
	AudioID int64
	Caption IPageCaption
}

func (*PageBlockAudio) CRC() uint32 {
	return 0x804361ea
}
func (*PageBlockAudio) _IPageBlock() {}

type PageBlockKicker struct {
	Text IRichText
}

func (*PageBlockKicker) CRC() uint32 {
	return 0x1e148390
}
func (*PageBlockKicker) _IPageBlock() {}

type PageBlockTable struct {
	_        struct{} `tl:"flags,bitflag"`
	Bordered bool     `tl:",omitempty:flags:0,implicit"`
	Striped  bool     `tl:",omitempty:flags:1,implicit"`
	Title    IRichText
	Rows     []IPageTableRow
}

func (*PageBlockTable) CRC() uint32 {
	return 0xbf4dea82
}
func (*PageBlockTable) _IPageBlock() {}

type PageBlockOrderedList struct {
	Items []IPageListOrderedItem
}

func (*PageBlockOrderedList) CRC() uint32 {
	return 0x9a8ae1e1
}
func (*PageBlockOrderedList) _IPageBlock() {}

type PageBlockDetails struct {
	_      struct{} `tl:"flags,bitflag"`
	Open   bool     `tl:",omitempty:flags:0,implicit"`
	Blocks []IPageBlock
	Title  IRichText
}

func (*PageBlockDetails) CRC() uint32 {
	return 0x76768bed
}
func (*PageBlockDetails) _IPageBlock() {}

type PageBlockRelatedArticles struct {
	Title    IRichText
	Articles []IPageRelatedArticle
}

func (*PageBlockRelatedArticles) CRC() uint32 {
	return 0x16115a96
}
func (*PageBlockRelatedArticles) _IPageBlock() {}

type PageBlockMap struct {
	Geo     IGeoPoint
	Zoom    int32
	W       int32
	H       int32
	Caption IPageCaption
}

func (*PageBlockMap) CRC() uint32 {
	return 0xa44f3ef6
}
func (*PageBlockMap) _IPageBlock() {}

type IPageCaption interface {
	tl.Object
	_IPageCaption()
}

var (
	_ IPageCaption = (*PageCaption)(nil)
)

type PageCaption struct {
	Text   IRichText
	Credit IRichText
}

func (*PageCaption) CRC() uint32 {
	return 0x6f747657
}
func (*PageCaption) _IPageCaption() {}

type IPageListItem interface {
	tl.Object
	_IPageListItem()
}

var (
	_ IPageListItem = (*PageListItemText)(nil)
	_ IPageListItem = (*PageListItemBlocks)(nil)
)

type PageListItemText struct {
	Text IRichText
}

func (*PageListItemText) CRC() uint32 {
	return 0xb92fb6cd
}
func (*PageListItemText) _IPageListItem() {}

type PageListItemBlocks struct {
	Blocks []IPageBlock
}

func (*PageListItemBlocks) CRC() uint32 {
	return 0x25e073fc
}
func (*PageListItemBlocks) _IPageListItem() {}

type IPageListOrderedItem interface {
	tl.Object
	_IPageListOrderedItem()
}

var (
	_ IPageListOrderedItem = (*PageListOrderedItemText)(nil)
	_ IPageListOrderedItem = (*PageListOrderedItemBlocks)(nil)
)

type PageListOrderedItemText struct {
	Num  string
	Text IRichText
}

func (*PageListOrderedItemText) CRC() uint32 {
	return 0x5e068047
}
func (*PageListOrderedItemText) _IPageListOrderedItem() {}

type PageListOrderedItemBlocks struct {
	Num    string
	Blocks []IPageBlock
}

func (*PageListOrderedItemBlocks) CRC() uint32 {
	return 0x98dd8936
}
func (*PageListOrderedItemBlocks) _IPageListOrderedItem() {}

type IPageRelatedArticle interface {
	tl.Object
	_IPageRelatedArticle()
}

var (
	_ IPageRelatedArticle = (*PageRelatedArticle)(nil)
)

type PageRelatedArticle struct {
	_             struct{} `tl:"flags,bitflag"`
	URL           string
	WebpageID     int64
	Title         *string `tl:",omitempty:flags:0"`
	Description   *string `tl:",omitempty:flags:1"`
	PhotoID       *int64  `tl:",omitempty:flags:2"`
	Author        *string `tl:",omitempty:flags:3"`
	PublishedDate *int32  `tl:",omitempty:flags:4"`
}

func (*PageRelatedArticle) CRC() uint32 {
	return 0xb390dc08
}
func (*PageRelatedArticle) _IPageRelatedArticle() {}

type IPageTableCell interface {
	tl.Object
	_IPageTableCell()
}

var (
	_ IPageTableCell = (*PageTableCell)(nil)
)

type PageTableCell struct {
	_            struct{}  `tl:"flags,bitflag"`
	Header       bool      `tl:",omitempty:flags:0,implicit"`
	AlignCenter  bool      `tl:",omitempty:flags:3,implicit"`
	AlignRight   bool      `tl:",omitempty:flags:4,implicit"`
	ValignMiddle bool      `tl:",omitempty:flags:5,implicit"`
	ValignBottom bool      `tl:",omitempty:flags:6,implicit"`
	Text         IRichText `tl:",omitempty:flags:7"`
	Colspan      *int32    `tl:",omitempty:flags:1"`
	Rowspan      *int32    `tl:",omitempty:flags:2"`
}

func (*PageTableCell) CRC() uint32 {
	return 0x34566b6a
}
func (*PageTableCell) _IPageTableCell() {}

type IPageTableRow interface {
	tl.Object
	_IPageTableRow()
}

var (
	_ IPageTableRow = (*PageTableRow)(nil)
)

type PageTableRow struct {
	Cells []IPageTableCell
}

func (*PageTableRow) CRC() uint32 {
	return 0xe0c0c5e5
}
func (*PageTableRow) _IPageTableRow() {}

type IPasswordKdfAlgo interface {
	tl.Object
	_IPasswordKdfAlgo()
}

var (
	_ IPasswordKdfAlgo = (*PasswordKdfAlgoUnknown)(nil)
	_ IPasswordKdfAlgo = (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPow)(nil)
)

type PasswordKdfAlgoUnknown struct{}

func (*PasswordKdfAlgoUnknown) CRC() uint32 {
	return 0xd45ab096
}
func (*PasswordKdfAlgoUnknown) _IPasswordKdfAlgo() {}

type PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPow) CRC() uint32 {
	return 0x3a912d4a
}
func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPow) _IPasswordKdfAlgo() {}

type IPaymentCharge interface {
	tl.Object
	_IPaymentCharge()
}

var (
	_ IPaymentCharge = (*PaymentCharge)(nil)
)

type PaymentCharge struct {
	ID               string
	ProviderChargeID string
}

func (*PaymentCharge) CRC() uint32 {
	return 0xea02c27e
}
func (*PaymentCharge) _IPaymentCharge() {}

type IPaymentFormMethod interface {
	tl.Object
	_IPaymentFormMethod()
}

var (
	_ IPaymentFormMethod = (*PaymentFormMethod)(nil)
)

type PaymentFormMethod struct {
	URL   string
	Title string
}

func (*PaymentFormMethod) CRC() uint32 {
	return 0x88f8f21b
}
func (*PaymentFormMethod) _IPaymentFormMethod() {}

type IPaymentRequestedInfo interface {
	tl.Object
	_IPaymentRequestedInfo()
}

var (
	_ IPaymentRequestedInfo = (*PaymentRequestedInfo)(nil)
)

type PaymentRequestedInfo struct {
	_               struct{}     `tl:"flags,bitflag"`
	Name            *string      `tl:",omitempty:flags:0"`
	Phone           *string      `tl:",omitempty:flags:1"`
	Email           *string      `tl:",omitempty:flags:2"`
	ShippingAddress IPostAddress `tl:",omitempty:flags:3"`
}

func (*PaymentRequestedInfo) CRC() uint32 {
	return 0x909c3f94
}
func (*PaymentRequestedInfo) _IPaymentRequestedInfo() {}

type IPaymentSavedCredentials interface {
	tl.Object
	_IPaymentSavedCredentials()
}

var (
	_ IPaymentSavedCredentials = (*PaymentSavedCredentialsCard)(nil)
)

type PaymentSavedCredentialsCard struct {
	ID    string
	Title string
}

func (*PaymentSavedCredentialsCard) CRC() uint32 {
	return 0xcdc27a1f
}
func (*PaymentSavedCredentialsCard) _IPaymentSavedCredentials() {}

// Chat partner or group.
type IPeer interface {
	tl.Object
	_IPeer()
}

var (
	_ IPeer = (*PeerUser)(nil)
	_ IPeer = (*PeerChat)(nil)
	_ IPeer = (*PeerChannel)(nil)
)

// Chat partner
type PeerUser struct {
	UserID int64 // User identifier
}

func (*PeerUser) CRC() uint32 {
	return 0x59511722
}
func (*PeerUser) _IPeer() {}

// Group.
type PeerChat struct {
	ChatID int64 // Group identifier
}

func (*PeerChat) CRC() uint32 {
	return 0x36c6019a
}
func (*PeerChat) _IPeer() {}

// Channel/supergroup
type PeerChannel struct {
	ChannelID int64 // Channel ID
}

func (*PeerChannel) CRC() uint32 {
	return 0xa2a5371e
}
func (*PeerChannel) _IPeer() {}

type IPeerBlocked interface {
	tl.Object
	_IPeerBlocked()
}

var (
	_ IPeerBlocked = (*PeerBlocked)(nil)
)

type PeerBlocked struct {
	PeerID IPeer
	Date   int32
}

func (*PeerBlocked) CRC() uint32 {
	return 0xe8fd8014
}
func (*PeerBlocked) _IPeerBlocked() {}

type IPeerColor interface {
	tl.Object
	_IPeerColor()
}

var (
	_ IPeerColor = (*PeerColor)(nil)
)

type PeerColor struct {
	_                 struct{} `tl:"flags,bitflag"`
	Color             *int32   `tl:",omitempty:flags:0"`
	BackgroundEmojiID *int64   `tl:",omitempty:flags:1"`
}

func (*PeerColor) CRC() uint32 {
	return 0xb54b5acf
}
func (*PeerColor) _IPeerColor() {}

type IPeerLocated interface {
	tl.Object
	_IPeerLocated()
}

var (
	_ IPeerLocated = (*PeerLocated)(nil)
	_ IPeerLocated = (*PeerSelfLocated)(nil)
)

type PeerLocated struct {
	Peer     IPeer
	Expires  int32
	Distance int32
}

func (*PeerLocated) CRC() uint32 {
	return 0xca461b5d
}
func (*PeerLocated) _IPeerLocated() {}

type PeerSelfLocated struct {
	Expires int32
}

func (*PeerSelfLocated) CRC() uint32 {
	return 0xf8ec284b
}
func (*PeerSelfLocated) _IPeerLocated() {}

type IPeerNotifySettings interface {
	tl.Object
	_IPeerNotifySettings()
}

var (
	_ IPeerNotifySettings = (*PeerNotifySettings)(nil)
)

type PeerNotifySettings struct {
	_                   struct{}           `tl:"flags,bitflag"`
	ShowPreviews        *bool              `tl:",omitempty:flags:0"`
	Silent              *bool              `tl:",omitempty:flags:1"`
	MuteUntil           *int32             `tl:",omitempty:flags:2"`
	IosSound            INotificationSound `tl:",omitempty:flags:3"`
	AndroidSound        INotificationSound `tl:",omitempty:flags:4"`
	OtherSound          INotificationSound `tl:",omitempty:flags:5"`
	StoriesMuted        *bool              `tl:",omitempty:flags:6"`
	StoriesHideSender   *bool              `tl:",omitempty:flags:7"`
	StoriesIosSound     INotificationSound `tl:",omitempty:flags:8"`
	StoriesAndroidSound INotificationSound `tl:",omitempty:flags:9"`
	StoriesOtherSound   INotificationSound `tl:",omitempty:flags:10"`
}

func (*PeerNotifySettings) CRC() uint32 {
	return 0x99622c0c
}
func (*PeerNotifySettings) _IPeerNotifySettings() {}

type IPeerSettings interface {
	tl.Object
	_IPeerSettings()
}

var (
	_ IPeerSettings = (*PeerSettings)(nil)
)

type PeerSettings struct {
	_                     struct{} `tl:"flags,bitflag"`
	ReportSpam            bool     `tl:",omitempty:flags:0,implicit"`
	AddContact            bool     `tl:",omitempty:flags:1,implicit"`
	BlockContact          bool     `tl:",omitempty:flags:2,implicit"`
	ShareContact          bool     `tl:",omitempty:flags:3,implicit"`
	NeedContactsException bool     `tl:",omitempty:flags:4,implicit"`
	ReportGeo             bool     `tl:",omitempty:flags:5,implicit"`
	Autoarchived          bool     `tl:",omitempty:flags:7,implicit"`
	InviteMembers         bool     `tl:",omitempty:flags:8,implicit"`
	RequestChatBroadcast  bool     `tl:",omitempty:flags:10,implicit"`
	GeoDistance           *int32   `tl:",omitempty:flags:6"`
	RequestChatTitle      *string  `tl:",omitempty:flags:9"`
	RequestChatDate       *int32   `tl:",omitempty:flags:9"`
}

func (*PeerSettings) CRC() uint32 {
	return 0xa518110d
}
func (*PeerSettings) _IPeerSettings() {}

type IPeerStories interface {
	tl.Object
	_IPeerStories()
}

var (
	_ IPeerStories = (*PeerStories)(nil)
)

type PeerStories struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IPeer
	MaxReadID *int32 `tl:",omitempty:flags:0"`
	Stories   []IStoryItem
}

func (*PeerStories) CRC() uint32 {
	return 0x9a35e999
}
func (*PeerStories) _IPeerStories() {}

type IPhoneCall interface {
	tl.Object
	_IPhoneCall()
}

var (
	_ IPhoneCall = (*PhoneCallEmpty)(nil)
	_ IPhoneCall = (*PhoneCallWaiting)(nil)
	_ IPhoneCall = (*PhoneCallRequested)(nil)
	_ IPhoneCall = (*PhoneCallAccepted)(nil)
	_ IPhoneCall = (*PhoneCall)(nil)
	_ IPhoneCall = (*PhoneCallDiscarded)(nil)
)

type PhoneCallEmpty struct {
	ID int64
}

func (*PhoneCallEmpty) CRC() uint32 {
	return 0x5366c915
}
func (*PhoneCallEmpty) _IPhoneCall() {}

type PhoneCallWaiting struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	Protocol      IPhoneCallProtocol
	ReceiveDate   *int32 `tl:",omitempty:flags:0"`
}

func (*PhoneCallWaiting) CRC() uint32 {
	return 0xc5226f17
}
func (*PhoneCallWaiting) _IPhoneCall() {}

type PhoneCallRequested struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GAHash        []byte
	Protocol      IPhoneCallProtocol
}

func (*PhoneCallRequested) CRC() uint32 {
	return 0x14b0ed0c
}
func (*PhoneCallRequested) _IPhoneCall() {}

type PhoneCallAccepted struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GB            []byte
	Protocol      IPhoneCallProtocol
}

func (*PhoneCallAccepted) CRC() uint32 {
	return 0x3660c311
}
func (*PhoneCallAccepted) _IPhoneCall() {}

type PhoneCall struct {
	_              struct{} `tl:"flags,bitflag"`
	P2PAllowed     bool     `tl:",omitempty:flags:5,implicit"`
	Video          bool     `tl:",omitempty:flags:6,implicit"`
	ID             int64
	AccessHash     int64
	Date           int32
	AdminID        int64
	ParticipantID  int64
	GAOrB          []byte
	KeyFingerprint int64
	Protocol       IPhoneCallProtocol
	Connections    []IPhoneConnection
	StartDate      int32
}

func (*PhoneCall) CRC() uint32 {
	return 0x967f7c67
}
func (*PhoneCall) _IPhoneCall() {}

type PhoneCallDiscarded struct {
	_          struct{} `tl:"flags,bitflag"`
	NeedRating bool     `tl:",omitempty:flags:2,implicit"`
	NeedDebug  bool     `tl:",omitempty:flags:3,implicit"`
	Video      bool     `tl:",omitempty:flags:6,implicit"`
	ID         int64
	Reason     EnumPhoneCallDiscardReason `tl:",omitempty:flags:0"`
	Duration   *int32                     `tl:",omitempty:flags:1"`
}

func (*PhoneCallDiscarded) CRC() uint32 {
	return 0x50ca4de1
}
func (*PhoneCallDiscarded) _IPhoneCall() {}

type IPhoneCallProtocol interface {
	tl.Object
	_IPhoneCallProtocol()
}

var (
	_ IPhoneCallProtocol = (*PhoneCallProtocol)(nil)
)

type PhoneCallProtocol struct {
	_               struct{} `tl:"flags,bitflag"`
	UdpP2P          bool     `tl:",omitempty:flags:0,implicit"`
	UdpReflector    bool     `tl:",omitempty:flags:1,implicit"`
	MinLayer        int32
	MaxLayer        int32
	LibraryVersions []string
}

func (*PhoneCallProtocol) CRC() uint32 {
	return 0xfc878fc8
}
func (*PhoneCallProtocol) _IPhoneCallProtocol() {}

type IPhoneConnection interface {
	tl.Object
	_IPhoneConnection()
}

var (
	_ IPhoneConnection = (*PhoneConnection)(nil)
	_ IPhoneConnection = (*PhoneConnectionWebrtc)(nil)
)

type PhoneConnection struct {
	_       struct{} `tl:"flags,bitflag"`
	Tcp     bool     `tl:",omitempty:flags:0,implicit"`
	ID      int64
	Ip      string
	Ipv6    string
	Port    int32
	PeerTag []byte
}

func (*PhoneConnection) CRC() uint32 {
	return 0x9cc123c7
}
func (*PhoneConnection) _IPhoneConnection() {}

type PhoneConnectionWebrtc struct {
	_        struct{} `tl:"flags,bitflag"`
	Turn     bool     `tl:",omitempty:flags:0,implicit"`
	Stun     bool     `tl:",omitempty:flags:1,implicit"`
	ID       int64
	Ip       string
	Ipv6     string
	Port     int32
	Username string
	Password string
}

func (*PhoneConnectionWebrtc) CRC() uint32 {
	return 0x635fe375
}
func (*PhoneConnectionWebrtc) _IPhoneConnection() {}

type IPhoto interface {
	tl.Object
	_IPhoto()
}

var (
	_ IPhoto = (*PhotoEmpty)(nil)
	_ IPhoto = (*Photo)(nil)
)

type PhotoEmpty struct {
	ID int64
}

func (*PhotoEmpty) CRC() uint32 {
	return 0x2331b22d
}
func (*PhotoEmpty) _IPhoto() {}

type Photo struct {
	_             struct{} `tl:"flags,bitflag"`
	HasStickers   bool     `tl:",omitempty:flags:0,implicit"`
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	Sizes         []IPhotoSize
	VideoSizes    []IVideoSize `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*Photo) CRC() uint32 {
	return 0xfb197a65
}
func (*Photo) _IPhoto() {}

type IPhotoSize interface {
	tl.Object
	_IPhotoSize()
}

var (
	_ IPhotoSize = (*PhotoSizeEmpty)(nil)
	_ IPhotoSize = (*PhotoSize)(nil)
	_ IPhotoSize = (*PhotoCachedSize)(nil)
	_ IPhotoSize = (*PhotoStrippedSize)(nil)
	_ IPhotoSize = (*PhotoSizeProgressive)(nil)
	_ IPhotoSize = (*PhotoPathSize)(nil)
)

type PhotoSizeEmpty struct {
	Type string
}

func (*PhotoSizeEmpty) CRC() uint32 {
	return 0xe17e23c
}
func (*PhotoSizeEmpty) _IPhotoSize() {}

type PhotoSize struct {
	Type string
	W    int32
	H    int32
	Size int32
}

func (*PhotoSize) CRC() uint32 {
	return 0x75c78e60
}
func (*PhotoSize) _IPhotoSize() {}

type PhotoCachedSize struct {
	Type  string
	W     int32
	H     int32
	Bytes []byte
}

func (*PhotoCachedSize) CRC() uint32 {
	return 0x21e1ad6
}
func (*PhotoCachedSize) _IPhotoSize() {}

type PhotoStrippedSize struct {
	Type  string
	Bytes []byte
}

func (*PhotoStrippedSize) CRC() uint32 {
	return 0xe0b0bc2e
}
func (*PhotoStrippedSize) _IPhotoSize() {}

type PhotoSizeProgressive struct {
	Type  string
	W     int32
	H     int32
	Sizes []int32
}

func (*PhotoSizeProgressive) CRC() uint32 {
	return 0xfa3efb95
}
func (*PhotoSizeProgressive) _IPhotoSize() {}

type PhotoPathSize struct {
	Type  string
	Bytes []byte
}

func (*PhotoPathSize) CRC() uint32 {
	return 0xd8214d41
}
func (*PhotoPathSize) _IPhotoSize() {}

type IPoll interface {
	tl.Object
	_IPoll()
}

var (
	_ IPoll = (*Poll)(nil)
)

type Poll struct {
	ID             int64
	_              struct{} `tl:"flags,bitflag"`
	Closed         bool     `tl:",omitempty:flags:0,implicit"`
	PublicVoters   bool     `tl:",omitempty:flags:1,implicit"`
	MultipleChoice bool     `tl:",omitempty:flags:2,implicit"`
	Quiz           bool     `tl:",omitempty:flags:3,implicit"`
	Question       string
	Answers        []IPollAnswer
	ClosePeriod    *int32 `tl:",omitempty:flags:4"`
	CloseDate      *int32 `tl:",omitempty:flags:5"`
}

func (*Poll) CRC() uint32 {
	return 0x86e18161
}
func (*Poll) _IPoll() {}

type IPollAnswer interface {
	tl.Object
	_IPollAnswer()
}

var (
	_ IPollAnswer = (*PollAnswer)(nil)
)

type PollAnswer struct {
	Text   string
	Option []byte
}

func (*PollAnswer) CRC() uint32 {
	return 0x6ca9c2e9
}
func (*PollAnswer) _IPollAnswer() {}

type IPollAnswerVoters interface {
	tl.Object
	_IPollAnswerVoters()
}

var (
	_ IPollAnswerVoters = (*PollAnswerVoters)(nil)
)

type PollAnswerVoters struct {
	_       struct{} `tl:"flags,bitflag"`
	Chosen  bool     `tl:",omitempty:flags:0,implicit"`
	Correct bool     `tl:",omitempty:flags:1,implicit"`
	Option  []byte
	Voters  int32
}

func (*PollAnswerVoters) CRC() uint32 {
	return 0x3b6ddad2
}
func (*PollAnswerVoters) _IPollAnswerVoters() {}

type IPollResults interface {
	tl.Object
	_IPollResults()
}

var (
	_ IPollResults = (*PollResults)(nil)
)

type PollResults struct {
	_                struct{}            `tl:"flags,bitflag"`
	Min              bool                `tl:",omitempty:flags:0,implicit"`
	Results          []IPollAnswerVoters `tl:",omitempty:flags:1"`
	TotalVoters      *int32              `tl:",omitempty:flags:2"`
	RecentVoters     []IPeer             `tl:",omitempty:flags:3"`
	Solution         *string             `tl:",omitempty:flags:4"`
	SolutionEntities []IMessageEntity    `tl:",omitempty:flags:4"`
}

func (*PollResults) CRC() uint32 {
	return 0x7adf2420
}
func (*PollResults) _IPollResults() {}

type IPopularContact interface {
	tl.Object
	_IPopularContact()
}

var (
	_ IPopularContact = (*PopularContact)(nil)
)

type PopularContact struct {
	ClientID  int64
	Importers int32
}

func (*PopularContact) CRC() uint32 {
	return 0x5ce14175
}
func (*PopularContact) _IPopularContact() {}

type IPostAddress interface {
	tl.Object
	_IPostAddress()
}

var (
	_ IPostAddress = (*PostAddress)(nil)
)

type PostAddress struct {
	StreetLine1 string
	StreetLine2 string
	City        string
	State       string
	CountryIso2 string
	PostCode    string
}

func (*PostAddress) CRC() uint32 {
	return 0x1e8caaeb
}
func (*PostAddress) _IPostAddress() {}

type IPostInteractionCounters interface {
	tl.Object
	_IPostInteractionCounters()
}

var (
	_ IPostInteractionCounters = (*PostInteractionCountersMessage)(nil)
	_ IPostInteractionCounters = (*PostInteractionCountersStory)(nil)
)

type PostInteractionCountersMessage struct {
	MsgID     int32
	Views     int32
	Forwards  int32
	Reactions int32
}

func (*PostInteractionCountersMessage) CRC() uint32 {
	return 0xe7058e7f
}
func (*PostInteractionCountersMessage) _IPostInteractionCounters() {}

type PostInteractionCountersStory struct {
	StoryID   int32
	Views     int32
	Forwards  int32
	Reactions int32
}

func (*PostInteractionCountersStory) CRC() uint32 {
	return 0x8a480e27
}
func (*PostInteractionCountersStory) _IPostInteractionCounters() {}

type IPremiumGiftCodeOption interface {
	tl.Object
	_IPremiumGiftCodeOption()
}

var (
	_ IPremiumGiftCodeOption = (*PremiumGiftCodeOption)(nil)
)

type PremiumGiftCodeOption struct {
	_             struct{} `tl:"flags,bitflag"`
	Users         int32
	Months        int32
	StoreProduct  *string `tl:",omitempty:flags:0"`
	StoreQuantity *int32  `tl:",omitempty:flags:1"`
	Currency      string
	Amount        int64
}

func (*PremiumGiftCodeOption) CRC() uint32 {
	return 0x257e962b
}
func (*PremiumGiftCodeOption) _IPremiumGiftCodeOption() {}

type IPremiumGiftOption interface {
	tl.Object
	_IPremiumGiftOption()
}

var (
	_ IPremiumGiftOption = (*PremiumGiftOption)(nil)
)

type PremiumGiftOption struct {
	_            struct{} `tl:"flags,bitflag"`
	Months       int32
	Currency     string
	Amount       int64
	BotURL       string
	StoreProduct *string `tl:",omitempty:flags:0"`
}

func (*PremiumGiftOption) CRC() uint32 {
	return 0x74c34319
}
func (*PremiumGiftOption) _IPremiumGiftOption() {}

type IPremiumSubscriptionOption interface {
	tl.Object
	_IPremiumSubscriptionOption()
}

var (
	_ IPremiumSubscriptionOption = (*PremiumSubscriptionOption)(nil)
)

type PremiumSubscriptionOption struct {
	_                  struct{} `tl:"flags,bitflag"`
	Current            bool     `tl:",omitempty:flags:1,implicit"`
	CanPurchaseUpgrade bool     `tl:",omitempty:flags:2,implicit"`
	Transaction        *string  `tl:",omitempty:flags:3"`
	Months             int32
	Currency           string
	Amount             int64
	BotURL             string
	StoreProduct       *string `tl:",omitempty:flags:0"`
}

func (*PremiumSubscriptionOption) CRC() uint32 {
	return 0x5f2d1df2
}
func (*PremiumSubscriptionOption) _IPremiumSubscriptionOption() {}

type IPrepaidGiveaway interface {
	tl.Object
	_IPrepaidGiveaway()
}

var (
	_ IPrepaidGiveaway = (*PrepaidGiveaway)(nil)
)

type PrepaidGiveaway struct {
	ID       int64
	Months   int32
	Quantity int32
	Date     int32
}

func (*PrepaidGiveaway) CRC() uint32 {
	return 0xb2539d54
}
func (*PrepaidGiveaway) _IPrepaidGiveaway() {}

type IPrivacyRule interface {
	tl.Object
	_IPrivacyRule()
}

var (
	_ IPrivacyRule = (*PrivacyValueAllowContacts)(nil)
	_ IPrivacyRule = (*PrivacyValueAllowAll)(nil)
	_ IPrivacyRule = (*PrivacyValueAllowUsers)(nil)
	_ IPrivacyRule = (*PrivacyValueDisallowContacts)(nil)
	_ IPrivacyRule = (*PrivacyValueDisallowAll)(nil)
	_ IPrivacyRule = (*PrivacyValueDisallowUsers)(nil)
	_ IPrivacyRule = (*PrivacyValueAllowChatParticipants)(nil)
	_ IPrivacyRule = (*PrivacyValueDisallowChatParticipants)(nil)
	_ IPrivacyRule = (*PrivacyValueAllowCloseFriends)(nil)
)

type PrivacyValueAllowContacts struct{}

func (*PrivacyValueAllowContacts) CRC() uint32 {
	return 0xfffe1bac
}
func (*PrivacyValueAllowContacts) _IPrivacyRule() {}

type PrivacyValueAllowAll struct{}

func (*PrivacyValueAllowAll) CRC() uint32 {
	return 0x65427b82
}
func (*PrivacyValueAllowAll) _IPrivacyRule() {}

type PrivacyValueAllowUsers struct {
	Users []int64
}

func (*PrivacyValueAllowUsers) CRC() uint32 {
	return 0xb8905fb2
}
func (*PrivacyValueAllowUsers) _IPrivacyRule() {}

type PrivacyValueDisallowContacts struct{}

func (*PrivacyValueDisallowContacts) CRC() uint32 {
	return 0xf888fa1a
}
func (*PrivacyValueDisallowContacts) _IPrivacyRule() {}

type PrivacyValueDisallowAll struct{}

func (*PrivacyValueDisallowAll) CRC() uint32 {
	return 0x8b73e763
}
func (*PrivacyValueDisallowAll) _IPrivacyRule() {}

type PrivacyValueDisallowUsers struct {
	Users []int64
}

func (*PrivacyValueDisallowUsers) CRC() uint32 {
	return 0xe4621141
}
func (*PrivacyValueDisallowUsers) _IPrivacyRule() {}

type PrivacyValueAllowChatParticipants struct {
	Chats []int64
}

func (*PrivacyValueAllowChatParticipants) CRC() uint32 {
	return 0x6b134e8e
}
func (*PrivacyValueAllowChatParticipants) _IPrivacyRule() {}

type PrivacyValueDisallowChatParticipants struct {
	Chats []int64
}

func (*PrivacyValueDisallowChatParticipants) CRC() uint32 {
	return 0x41c87565
}
func (*PrivacyValueDisallowChatParticipants) _IPrivacyRule() {}

type PrivacyValueAllowCloseFriends struct{}

func (*PrivacyValueAllowCloseFriends) CRC() uint32 {
	return 0xf7e8d89b
}
func (*PrivacyValueAllowCloseFriends) _IPrivacyRule() {}

type IPublicForward interface {
	tl.Object
	_IPublicForward()
}

var (
	_ IPublicForward = (*PublicForwardMessage)(nil)
	_ IPublicForward = (*PublicForwardStory)(nil)
)

type PublicForwardMessage struct {
	Message IMessage
}

func (*PublicForwardMessage) CRC() uint32 {
	return 0x1f2bf4a
}
func (*PublicForwardMessage) _IPublicForward() {}

type PublicForwardStory struct {
	Peer  IPeer
	Story IStoryItem
}

func (*PublicForwardStory) CRC() uint32 {
	return 0xedf3add0
}
func (*PublicForwardStory) _IPublicForward() {}

type IReaction interface {
	tl.Object
	_IReaction()
}

var (
	_ IReaction = (*ReactionEmpty)(nil)
	_ IReaction = (*ReactionEmoji)(nil)
	_ IReaction = (*ReactionCustomEmoji)(nil)
)

type ReactionEmpty struct{}

func (*ReactionEmpty) CRC() uint32 {
	return 0x79f5d419
}
func (*ReactionEmpty) _IReaction() {}

type ReactionEmoji struct {
	Emoticon string
}

func (*ReactionEmoji) CRC() uint32 {
	return 0x1b2286b8
}
func (*ReactionEmoji) _IReaction() {}

type ReactionCustomEmoji struct {
	DocumentID int64
}

func (*ReactionCustomEmoji) CRC() uint32 {
	return 0x8935fc73
}
func (*ReactionCustomEmoji) _IReaction() {}

type IReactionCount interface {
	tl.Object
	_IReactionCount()
}

var (
	_ IReactionCount = (*ReactionCount)(nil)
)

type ReactionCount struct {
	_           struct{} `tl:"flags,bitflag"`
	ChosenOrder *int32   `tl:",omitempty:flags:0"`
	Reaction    IReaction
	Count       int32
}

func (*ReactionCount) CRC() uint32 {
	return 0xa3d1cb80
}
func (*ReactionCount) _IReactionCount() {}

type IReadParticipantDate interface {
	tl.Object
	_IReadParticipantDate()
}

var (
	_ IReadParticipantDate = (*ReadParticipantDate)(nil)
)

type ReadParticipantDate struct {
	UserID int64
	Date   int32
}

func (*ReadParticipantDate) CRC() uint32 {
	return 0x4a4ff172
}
func (*ReadParticipantDate) _IReadParticipantDate() {}

type IReceivedNotifyMessage interface {
	tl.Object
	_IReceivedNotifyMessage()
}

var (
	_ IReceivedNotifyMessage = (*ReceivedNotifyMessage)(nil)
)

type ReceivedNotifyMessage struct {
	ID    int32
	Flags int32
}

func (*ReceivedNotifyMessage) CRC() uint32 {
	return 0xa384b779
}
func (*ReceivedNotifyMessage) _IReceivedNotifyMessage() {}

type IRecentMeURL interface {
	tl.Object
	_IRecentMeURL()
}

var (
	_ IRecentMeURL = (*RecentMeURLUnknown)(nil)
	_ IRecentMeURL = (*RecentMeURLUser)(nil)
	_ IRecentMeURL = (*RecentMeURLChat)(nil)
	_ IRecentMeURL = (*RecentMeURLChatInvite)(nil)
	_ IRecentMeURL = (*RecentMeURLStickerSet)(nil)
)

type RecentMeURLUnknown struct {
	URL string
}

func (*RecentMeURLUnknown) CRC() uint32 {
	return 0x46e1d13d
}
func (*RecentMeURLUnknown) _IRecentMeURL() {}

type RecentMeURLUser struct {
	URL    string
	UserID int64
}

func (*RecentMeURLUser) CRC() uint32 {
	return 0xb92c09e2
}
func (*RecentMeURLUser) _IRecentMeURL() {}

type RecentMeURLChat struct {
	URL    string
	ChatID int64
}

func (*RecentMeURLChat) CRC() uint32 {
	return 0xb2da71d2
}
func (*RecentMeURLChat) _IRecentMeURL() {}

type RecentMeURLChatInvite struct {
	URL        string
	ChatInvite IChatInvite
}

func (*RecentMeURLChatInvite) CRC() uint32 {
	return 0xeb49081d
}
func (*RecentMeURLChatInvite) _IRecentMeURL() {}

type RecentMeURLStickerSet struct {
	URL string
	Set IStickerSetCovered
}

func (*RecentMeURLStickerSet) CRC() uint32 {
	return 0xbc0a57dc
}
func (*RecentMeURLStickerSet) _IRecentMeURL() {}

type IReplyMarkup interface {
	tl.Object
	_IReplyMarkup()
}

var (
	_ IReplyMarkup = (*ReplyKeyboardHide)(nil)
	_ IReplyMarkup = (*ReplyKeyboardForceReply)(nil)
	_ IReplyMarkup = (*ReplyKeyboardMarkup)(nil)
	_ IReplyMarkup = (*ReplyInlineMarkup)(nil)
)

type ReplyKeyboardHide struct {
	_         struct{} `tl:"flags,bitflag"`
	Selective bool     `tl:",omitempty:flags:2,implicit"`
}

func (*ReplyKeyboardHide) CRC() uint32 {
	return 0xa03e5b85
}
func (*ReplyKeyboardHide) _IReplyMarkup() {}

type ReplyKeyboardForceReply struct {
	_           struct{} `tl:"flags,bitflag"`
	SingleUse   bool     `tl:",omitempty:flags:1,implicit"`
	Selective   bool     `tl:",omitempty:flags:2,implicit"`
	Placeholder *string  `tl:",omitempty:flags:3"`
}

func (*ReplyKeyboardForceReply) CRC() uint32 {
	return 0x86b40b08
}
func (*ReplyKeyboardForceReply) _IReplyMarkup() {}

type ReplyKeyboardMarkup struct {
	_           struct{} `tl:"flags,bitflag"`
	Resize      bool     `tl:",omitempty:flags:0,implicit"`
	SingleUse   bool     `tl:",omitempty:flags:1,implicit"`
	Selective   bool     `tl:",omitempty:flags:2,implicit"`
	Persistent  bool     `tl:",omitempty:flags:4,implicit"`
	Rows        []IKeyboardButtonRow
	Placeholder *string `tl:",omitempty:flags:3"`
}

func (*ReplyKeyboardMarkup) CRC() uint32 {
	return 0x85dd99d1
}
func (*ReplyKeyboardMarkup) _IReplyMarkup() {}

type ReplyInlineMarkup struct {
	Rows []IKeyboardButtonRow
}

func (*ReplyInlineMarkup) CRC() uint32 {
	return 0x48a30254
}
func (*ReplyInlineMarkup) _IReplyMarkup() {}

type IRequestPeerType interface {
	tl.Object
	_IRequestPeerType()
}

var (
	_ IRequestPeerType = (*RequestPeerTypeUser)(nil)
	_ IRequestPeerType = (*RequestPeerTypeChat)(nil)
	_ IRequestPeerType = (*RequestPeerTypeBroadcast)(nil)
)

type RequestPeerTypeUser struct {
	_       struct{} `tl:"flags,bitflag"`
	Bot     *bool    `tl:",omitempty:flags:0"`
	Premium *bool    `tl:",omitempty:flags:1"`
}

func (*RequestPeerTypeUser) CRC() uint32 {
	return 0x5f3b8a00
}
func (*RequestPeerTypeUser) _IRequestPeerType() {}

type RequestPeerTypeChat struct {
	_               struct{}         `tl:"flags,bitflag"`
	Creator         bool             `tl:",omitempty:flags:0,implicit"`
	BotParticipant  bool             `tl:",omitempty:flags:5,implicit"`
	HasUsername     *bool            `tl:",omitempty:flags:3"`
	Forum           *bool            `tl:",omitempty:flags:4"`
	UserAdminRights IChatAdminRights `tl:",omitempty:flags:1"`
	BotAdminRights  IChatAdminRights `tl:",omitempty:flags:2"`
}

func (*RequestPeerTypeChat) CRC() uint32 {
	return 0xc9f06e1b
}
func (*RequestPeerTypeChat) _IRequestPeerType() {}

type RequestPeerTypeBroadcast struct {
	_               struct{}         `tl:"flags,bitflag"`
	Creator         bool             `tl:",omitempty:flags:0,implicit"`
	HasUsername     *bool            `tl:",omitempty:flags:3"`
	UserAdminRights IChatAdminRights `tl:",omitempty:flags:1"`
	BotAdminRights  IChatAdminRights `tl:",omitempty:flags:2"`
}

func (*RequestPeerTypeBroadcast) CRC() uint32 {
	return 0x339bef6c
}
func (*RequestPeerTypeBroadcast) _IRequestPeerType() {}

type IRestrictionReason interface {
	tl.Object
	_IRestrictionReason()
}

var (
	_ IRestrictionReason = (*RestrictionReason)(nil)
)

type RestrictionReason struct {
	Platform string
	Reason   string
	Text     string
}

func (*RestrictionReason) CRC() uint32 {
	return 0xd072acb4
}
func (*RestrictionReason) _IRestrictionReason() {}

type IRichText interface {
	tl.Object
	_IRichText()
}

var (
	_ IRichText = (*TextEmpty)(nil)
	_ IRichText = (*TextPlain)(nil)
	_ IRichText = (*TextBold)(nil)
	_ IRichText = (*TextItalic)(nil)
	_ IRichText = (*TextUnderline)(nil)
	_ IRichText = (*TextStrike)(nil)
	_ IRichText = (*TextFixed)(nil)
	_ IRichText = (*TextURL)(nil)
	_ IRichText = (*TextEmail)(nil)
	_ IRichText = (*TextConcat)(nil)
	_ IRichText = (*TextSubscript)(nil)
	_ IRichText = (*TextSuperscript)(nil)
	_ IRichText = (*TextMarked)(nil)
	_ IRichText = (*TextPhone)(nil)
	_ IRichText = (*TextImage)(nil)
	_ IRichText = (*TextAnchor)(nil)
)

type TextEmpty struct{}

func (*TextEmpty) CRC() uint32 {
	return 0xdc3d824f
}
func (*TextEmpty) _IRichText() {}

type TextPlain struct {
	Text string
}

func (*TextPlain) CRC() uint32 {
	return 0x744694e0
}
func (*TextPlain) _IRichText() {}

type TextBold struct {
	Text IRichText
}

func (*TextBold) CRC() uint32 {
	return 0x6724abc4
}
func (*TextBold) _IRichText() {}

type TextItalic struct {
	Text IRichText
}

func (*TextItalic) CRC() uint32 {
	return 0xd912a59c
}
func (*TextItalic) _IRichText() {}

type TextUnderline struct {
	Text IRichText
}

func (*TextUnderline) CRC() uint32 {
	return 0xc12622c4
}
func (*TextUnderline) _IRichText() {}

type TextStrike struct {
	Text IRichText
}

func (*TextStrike) CRC() uint32 {
	return 0x9bf8bb95
}
func (*TextStrike) _IRichText() {}

type TextFixed struct {
	Text IRichText
}

func (*TextFixed) CRC() uint32 {
	return 0x6c3f19b9
}
func (*TextFixed) _IRichText() {}

type TextURL struct {
	Text      IRichText
	URL       string
	WebpageID int64
}

func (*TextURL) CRC() uint32 {
	return 0x3c2884c1
}
func (*TextURL) _IRichText() {}

type TextEmail struct {
	Text  IRichText
	Email string
}

func (*TextEmail) CRC() uint32 {
	return 0xde5a0dd6
}
func (*TextEmail) _IRichText() {}

type TextConcat struct {
	Texts []IRichText
}

func (*TextConcat) CRC() uint32 {
	return 0x7e6260d7
}
func (*TextConcat) _IRichText() {}

type TextSubscript struct {
	Text IRichText
}

func (*TextSubscript) CRC() uint32 {
	return 0xed6a8504
}
func (*TextSubscript) _IRichText() {}

type TextSuperscript struct {
	Text IRichText
}

func (*TextSuperscript) CRC() uint32 {
	return 0xc7fb5e01
}
func (*TextSuperscript) _IRichText() {}

type TextMarked struct {
	Text IRichText
}

func (*TextMarked) CRC() uint32 {
	return 0x34b8621
}
func (*TextMarked) _IRichText() {}

type TextPhone struct {
	Text  IRichText
	Phone string
}

func (*TextPhone) CRC() uint32 {
	return 0x1ccb966a
}
func (*TextPhone) _IRichText() {}

type TextImage struct {
	DocumentID int64
	W          int32
	H          int32
}

func (*TextImage) CRC() uint32 {
	return 0x81ccf4f
}
func (*TextImage) _IRichText() {}

type TextAnchor struct {
	Text IRichText
	Name string
}

func (*TextAnchor) CRC() uint32 {
	return 0x35553762
}
func (*TextAnchor) _IRichText() {}

type ISavedContact interface {
	tl.Object
	_ISavedContact()
}

var (
	_ ISavedContact = (*SavedPhoneContact)(nil)
)

type SavedPhoneContact struct {
	Phone     string
	FirstName string
	LastName  string
	Date      int32
}

func (*SavedPhoneContact) CRC() uint32 {
	return 0x1142bd56
}
func (*SavedPhoneContact) _ISavedContact() {}

type ISearchResultsCalendarPeriod interface {
	tl.Object
	_ISearchResultsCalendarPeriod()
}

var (
	_ ISearchResultsCalendarPeriod = (*SearchResultsCalendarPeriod)(nil)
)

type SearchResultsCalendarPeriod struct {
	Date     int32
	MinMsgID int32
	MaxMsgID int32
	Count    int32
}

func (*SearchResultsCalendarPeriod) CRC() uint32 {
	return 0xc9b0539f
}
func (*SearchResultsCalendarPeriod) _ISearchResultsCalendarPeriod() {}

type ISearchResultsPosition interface {
	tl.Object
	_ISearchResultsPosition()
}

var (
	_ ISearchResultsPosition = (*SearchResultPosition)(nil)
)

type SearchResultPosition struct {
	MsgID  int32
	Date   int32
	Offset int32
}

func (*SearchResultPosition) CRC() uint32 {
	return 0x7f648b67
}
func (*SearchResultPosition) _ISearchResultsPosition() {}

// Encrypted secure credentials
type ISecureCredentialsEncrypted interface {
	tl.Object
	_ISecureCredentialsEncrypted()
}

var (
	_ ISecureCredentialsEncrypted = (*SecureCredentialsEncrypted)(nil)
)

// Encrypted credentials required to decrypt telegram passport data.
type SecureCredentialsEncrypted struct {
	Data   []byte
	Hash   []byte
	Secret []byte
}

func (*SecureCredentialsEncrypted) CRC() uint32 {
	return 0x33f0ea47
}
func (*SecureCredentialsEncrypted) _ISecureCredentialsEncrypted() {}

type ISecureData interface {
	tl.Object
	_ISecureData()
}

var (
	_ ISecureData = (*SecureData)(nil)
)

type SecureData struct {
	Data     []byte
	DataHash []byte
	Secret   []byte
}

func (*SecureData) CRC() uint32 {
	return 0x8aeabec3
}
func (*SecureData) _ISecureData() {}

type ISecureFile interface {
	tl.Object
	_ISecureFile()
}

var (
	_ ISecureFile = (*SecureFileEmpty)(nil)
	_ ISecureFile = (*SecureFile)(nil)
)

type SecureFileEmpty struct{}

func (*SecureFileEmpty) CRC() uint32 {
	return 0x64199744
}
func (*SecureFileEmpty) _ISecureFile() {}

type SecureFile struct {
	ID         int64
	AccessHash int64
	Size       int64
	DcID       int32
	Date       int32
	FileHash   []byte
	Secret     []byte
}

func (*SecureFile) CRC() uint32 {
	return 0x7d09c27e
}
func (*SecureFile) _ISecureFile() {}

type ISecurePasswordKdfAlgo interface {
	tl.Object
	_ISecurePasswordKdfAlgo()
}

var (
	_ ISecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoUnknown)(nil)
	_ ISecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000)(nil)
	_ ISecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoSHA512)(nil)
)

type SecurePasswordKdfAlgoUnknown struct{}

func (*SecurePasswordKdfAlgoUnknown) CRC() uint32 {
	return 0x4a8537
}
func (*SecurePasswordKdfAlgoUnknown) _ISecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000 struct {
	Salt []byte
}

func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000) CRC() uint32 {
	return 0xbbf2dda0
}
func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000) _ISecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoSHA512 struct {
	Salt []byte
}

func (*SecurePasswordKdfAlgoSHA512) CRC() uint32 {
	return 0x86471d92
}
func (*SecurePasswordKdfAlgoSHA512) _ISecurePasswordKdfAlgo() {}

type ISecurePlainData interface {
	tl.Object
	_ISecurePlainData()
}

var (
	_ ISecurePlainData = (*SecurePlainPhone)(nil)
	_ ISecurePlainData = (*SecurePlainEmail)(nil)
)

type SecurePlainPhone struct {
	Phone string
}

func (*SecurePlainPhone) CRC() uint32 {
	return 0x7d6099dd
}
func (*SecurePlainPhone) _ISecurePlainData() {}

type SecurePlainEmail struct {
	Email string
}

func (*SecurePlainEmail) CRC() uint32 {
	return 0x21ec5a5f
}
func (*SecurePlainEmail) _ISecurePlainData() {}

type ISecureRequiredType interface {
	tl.Object
	_ISecureRequiredType()
}

var (
	_ ISecureRequiredType = (*SecureRequiredType)(nil)
	_ ISecureRequiredType = (*SecureRequiredTypeOneOf)(nil)
)

type SecureRequiredType struct {
	_                   struct{} `tl:"flags,bitflag"`
	NativeNames         bool     `tl:",omitempty:flags:0,implicit"`
	SelfieRequired      bool     `tl:",omitempty:flags:1,implicit"`
	TranslationRequired bool     `tl:",omitempty:flags:2,implicit"`
	Type                EnumSecureValueType
}

func (*SecureRequiredType) CRC() uint32 {
	return 0x829d99da
}
func (*SecureRequiredType) _ISecureRequiredType() {}

type SecureRequiredTypeOneOf struct {
	Types []ISecureRequiredType
}

func (*SecureRequiredTypeOneOf) CRC() uint32 {
	return 0x27477b4
}
func (*SecureRequiredTypeOneOf) _ISecureRequiredType() {}

type ISecureSecretSettings interface {
	tl.Object
	_ISecureSecretSettings()
}

var (
	_ ISecureSecretSettings = (*SecureSecretSettings)(nil)
)

type SecureSecretSettings struct {
	SecureAlgo     ISecurePasswordKdfAlgo
	SecureSecret   []byte
	SecureSecretID int64
}

func (*SecureSecretSettings) CRC() uint32 {
	return 0x1527bcac
}
func (*SecureSecretSettings) _ISecureSecretSettings() {}

type ISecureValue interface {
	tl.Object
	_ISecureValue()
}

var (
	_ ISecureValue = (*SecureValue)(nil)
)

type SecureValue struct {
	_           struct{} `tl:"flags,bitflag"`
	Type        EnumSecureValueType
	Data        ISecureData      `tl:",omitempty:flags:0"`
	FrontSide   ISecureFile      `tl:",omitempty:flags:1"`
	ReverseSide ISecureFile      `tl:",omitempty:flags:2"`
	Selfie      ISecureFile      `tl:",omitempty:flags:3"`
	Translation []ISecureFile    `tl:",omitempty:flags:6"`
	Files       []ISecureFile    `tl:",omitempty:flags:4"`
	PlainData   ISecurePlainData `tl:",omitempty:flags:5"`
	Hash        []byte
}

func (*SecureValue) CRC() uint32 {
	return 0x187fa0ca
}
func (*SecureValue) _ISecureValue() {}

type ISecureValueError interface {
	tl.Object
	_ISecureValueError()
}

var (
	_ ISecureValueError = (*SecureValueErrorData)(nil)
	_ ISecureValueError = (*SecureValueErrorFrontSide)(nil)
	_ ISecureValueError = (*SecureValueErrorReverseSide)(nil)
	_ ISecureValueError = (*SecureValueErrorSelfie)(nil)
	_ ISecureValueError = (*SecureValueErrorFile)(nil)
	_ ISecureValueError = (*SecureValueErrorFiles)(nil)
	_ ISecureValueError = (*SecureValueError)(nil)
	_ ISecureValueError = (*SecureValueErrorTranslationFile)(nil)
	_ ISecureValueError = (*SecureValueErrorTranslationFiles)(nil)
)

type SecureValueErrorData struct {
	Type     EnumSecureValueType
	DataHash []byte
	Field    string
	Text     string
}

func (*SecureValueErrorData) CRC() uint32 {
	return 0xe8a40bd9
}
func (*SecureValueErrorData) _ISecureValueError() {}

type SecureValueErrorFrontSide struct {
	Type     EnumSecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorFrontSide) CRC() uint32 {
	return 0xbe3dfa
}
func (*SecureValueErrorFrontSide) _ISecureValueError() {}

type SecureValueErrorReverseSide struct {
	Type     EnumSecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorReverseSide) CRC() uint32 {
	return 0x868a2aa5
}
func (*SecureValueErrorReverseSide) _ISecureValueError() {}

type SecureValueErrorSelfie struct {
	Type     EnumSecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorSelfie) CRC() uint32 {
	return 0xe537ced6
}
func (*SecureValueErrorSelfie) _ISecureValueError() {}

type SecureValueErrorFile struct {
	Type     EnumSecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorFile) CRC() uint32 {
	return 0x7a700873
}
func (*SecureValueErrorFile) _ISecureValueError() {}

type SecureValueErrorFiles struct {
	Type     EnumSecureValueType
	FileHash [][]byte
	Text     string
}

func (*SecureValueErrorFiles) CRC() uint32 {
	return 0x666220e9
}
func (*SecureValueErrorFiles) _ISecureValueError() {}

type SecureValueError struct {
	Type EnumSecureValueType
	Hash []byte
	Text string
}

func (*SecureValueError) CRC() uint32 {
	return 0x869d758f
}
func (*SecureValueError) _ISecureValueError() {}

type SecureValueErrorTranslationFile struct {
	Type     EnumSecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorTranslationFile) CRC() uint32 {
	return 0xa1144770
}
func (*SecureValueErrorTranslationFile) _ISecureValueError() {}

type SecureValueErrorTranslationFiles struct {
	Type     EnumSecureValueType
	FileHash [][]byte
	Text     string
}

func (*SecureValueErrorTranslationFiles) CRC() uint32 {
	return 0x34636dd8
}
func (*SecureValueErrorTranslationFiles) _ISecureValueError() {}

type ISecureValueHash interface {
	tl.Object
	_ISecureValueHash()
}

var (
	_ ISecureValueHash = (*SecureValueHash)(nil)
)

type SecureValueHash struct {
	Type EnumSecureValueType
	Hash []byte
}

func (*SecureValueHash) CRC() uint32 {
	return 0xed1ecdb0
}
func (*SecureValueHash) _ISecureValueHash() {}

type ISendAsPeer interface {
	tl.Object
	_ISendAsPeer()
}

var (
	_ ISendAsPeer = (*SendAsPeer)(nil)
)

type SendAsPeer struct {
	_               struct{} `tl:"flags,bitflag"`
	PremiumRequired bool     `tl:",omitempty:flags:0,implicit"`
	Peer            IPeer
}

func (*SendAsPeer) CRC() uint32 {
	return 0xb81c7034
}
func (*SendAsPeer) _ISendAsPeer() {}

type ISendMessageAction interface {
	tl.Object
	_ISendMessageAction()
}

var (
	_ ISendMessageAction = (*SendMessageTypingAction)(nil)
	_ ISendMessageAction = (*SendMessageCancelAction)(nil)
	_ ISendMessageAction = (*SendMessageRecordVideoAction)(nil)
	_ ISendMessageAction = (*SendMessageUploadVideoAction)(nil)
	_ ISendMessageAction = (*SendMessageRecordAudioAction)(nil)
	_ ISendMessageAction = (*SendMessageUploadAudioAction)(nil)
	_ ISendMessageAction = (*SendMessageUploadPhotoAction)(nil)
	_ ISendMessageAction = (*SendMessageUploadDocumentAction)(nil)
	_ ISendMessageAction = (*SendMessageGeoLocationAction)(nil)
	_ ISendMessageAction = (*SendMessageChooseContactAction)(nil)
	_ ISendMessageAction = (*SendMessageGamePlayAction)(nil)
	_ ISendMessageAction = (*SendMessageRecordRoundAction)(nil)
	_ ISendMessageAction = (*SendMessageUploadRoundAction)(nil)
	_ ISendMessageAction = (*SpeakingInGroupCallAction)(nil)
	_ ISendMessageAction = (*SendMessageHistoryImportAction)(nil)
	_ ISendMessageAction = (*SendMessageChooseStickerAction)(nil)
	_ ISendMessageAction = (*SendMessageEmojiInteraction)(nil)
	_ ISendMessageAction = (*SendMessageEmojiInteractionSeen)(nil)
)

type SendMessageTypingAction struct{}

func (*SendMessageTypingAction) CRC() uint32 {
	return 0x16bf744e
}
func (*SendMessageTypingAction) _ISendMessageAction() {}

type SendMessageCancelAction struct{}

func (*SendMessageCancelAction) CRC() uint32 {
	return 0xfd5ec8f5
}
func (*SendMessageCancelAction) _ISendMessageAction() {}

type SendMessageRecordVideoAction struct{}

func (*SendMessageRecordVideoAction) CRC() uint32 {
	return 0xa187d66f
}
func (*SendMessageRecordVideoAction) _ISendMessageAction() {}

type SendMessageUploadVideoAction struct {
	Progress int32
}

func (*SendMessageUploadVideoAction) CRC() uint32 {
	return 0xe9763aec
}
func (*SendMessageUploadVideoAction) _ISendMessageAction() {}

type SendMessageRecordAudioAction struct{}

func (*SendMessageRecordAudioAction) CRC() uint32 {
	return 0xd52f73f7
}
func (*SendMessageRecordAudioAction) _ISendMessageAction() {}

type SendMessageUploadAudioAction struct {
	Progress int32
}

func (*SendMessageUploadAudioAction) CRC() uint32 {
	return 0xf351d7ab
}
func (*SendMessageUploadAudioAction) _ISendMessageAction() {}

type SendMessageUploadPhotoAction struct {
	Progress int32
}

func (*SendMessageUploadPhotoAction) CRC() uint32 {
	return 0xd1d34a26
}
func (*SendMessageUploadPhotoAction) _ISendMessageAction() {}

type SendMessageUploadDocumentAction struct {
	Progress int32
}

func (*SendMessageUploadDocumentAction) CRC() uint32 {
	return 0xaa0cd9e4
}
func (*SendMessageUploadDocumentAction) _ISendMessageAction() {}

type SendMessageGeoLocationAction struct{}

func (*SendMessageGeoLocationAction) CRC() uint32 {
	return 0x176f8ba1
}
func (*SendMessageGeoLocationAction) _ISendMessageAction() {}

type SendMessageChooseContactAction struct{}

func (*SendMessageChooseContactAction) CRC() uint32 {
	return 0x628cbc6f
}
func (*SendMessageChooseContactAction) _ISendMessageAction() {}

type SendMessageGamePlayAction struct{}

func (*SendMessageGamePlayAction) CRC() uint32 {
	return 0xdd6a8f48
}
func (*SendMessageGamePlayAction) _ISendMessageAction() {}

type SendMessageRecordRoundAction struct{}

func (*SendMessageRecordRoundAction) CRC() uint32 {
	return 0x88f27fbc
}
func (*SendMessageRecordRoundAction) _ISendMessageAction() {}

type SendMessageUploadRoundAction struct {
	Progress int32
}

func (*SendMessageUploadRoundAction) CRC() uint32 {
	return 0x243e1c66
}
func (*SendMessageUploadRoundAction) _ISendMessageAction() {}

type SpeakingInGroupCallAction struct{}

func (*SpeakingInGroupCallAction) CRC() uint32 {
	return 0xd92c2285
}
func (*SpeakingInGroupCallAction) _ISendMessageAction() {}

type SendMessageHistoryImportAction struct {
	Progress int32
}

func (*SendMessageHistoryImportAction) CRC() uint32 {
	return 0xdbda9246
}
func (*SendMessageHistoryImportAction) _ISendMessageAction() {}

type SendMessageChooseStickerAction struct{}

func (*SendMessageChooseStickerAction) CRC() uint32 {
	return 0xb05ac6b1
}
func (*SendMessageChooseStickerAction) _ISendMessageAction() {}

type SendMessageEmojiInteraction struct {
	Emoticon    string
	MsgID       int32
	Interaction IDataJSON
}

func (*SendMessageEmojiInteraction) CRC() uint32 {
	return 0x25972bcb
}
func (*SendMessageEmojiInteraction) _ISendMessageAction() {}

type SendMessageEmojiInteractionSeen struct {
	Emoticon string
}

func (*SendMessageEmojiInteractionSeen) CRC() uint32 {
	return 0xb665902e
}
func (*SendMessageEmojiInteractionSeen) _ISendMessageAction() {}

type IShippingOption interface {
	tl.Object
	_IShippingOption()
}

var (
	_ IShippingOption = (*ShippingOption)(nil)
)

type ShippingOption struct {
	ID     string
	Title  string
	Prices []ILabeledPrice
}

func (*ShippingOption) CRC() uint32 {
	return 0xb6213cdf
}
func (*ShippingOption) _IShippingOption() {}

type ISimpleWebViewResult interface {
	tl.Object
	_ISimpleWebViewResult()
}

var (
	_ ISimpleWebViewResult = (*SimpleWebViewResultURL)(nil)
)

type SimpleWebViewResultURL struct {
	URL string
}

func (*SimpleWebViewResultURL) CRC() uint32 {
	return 0x882f76bb
}
func (*SimpleWebViewResultURL) _ISimpleWebViewResult() {}

type ISponsoredMessage interface {
	tl.Object
	_ISponsoredMessage()
}

var (
	_ ISponsoredMessage = (*SponsoredMessage)(nil)
)

type SponsoredMessage struct {
	_              struct{} `tl:"flags,bitflag"`
	Recommended    bool     `tl:",omitempty:flags:5,implicit"`
	ShowPeerPhoto  bool     `tl:",omitempty:flags:6,implicit"`
	RandomID       []byte
	FromID         IPeer             `tl:",omitempty:flags:3"`
	ChatInvite     IChatInvite       `tl:",omitempty:flags:4"`
	ChatInviteHash *string           `tl:",omitempty:flags:4"`
	ChannelPost    *int32            `tl:",omitempty:flags:2"`
	StartParam     *string           `tl:",omitempty:flags:0"`
	Webpage        ISponsoredWebPage `tl:",omitempty:flags:9"`
	App            IBotApp           `tl:",omitempty:flags:10"`
	Message        string
	Entities       []IMessageEntity `tl:",omitempty:flags:1"`
	ButtonText     *string          `tl:",omitempty:flags:11"`
	SponsorInfo    *string          `tl:",omitempty:flags:7"`
	AdditionalInfo *string          `tl:",omitempty:flags:8"`
}

func (*SponsoredMessage) CRC() uint32 {
	return 0xed5383f7
}
func (*SponsoredMessage) _ISponsoredMessage() {}

type ISponsoredWebPage interface {
	tl.Object
	_ISponsoredWebPage()
}

var (
	_ ISponsoredWebPage = (*SponsoredWebPage)(nil)
)

type SponsoredWebPage struct {
	_        struct{} `tl:"flags,bitflag"`
	URL      string
	SiteName string
	Photo    IPhoto `tl:",omitempty:flags:0"`
}

func (*SponsoredWebPage) CRC() uint32 {
	return 0x3db8ec63
}
func (*SponsoredWebPage) _ISponsoredWebPage() {}

type IStatsAbsValueAndPrev interface {
	tl.Object
	_IStatsAbsValueAndPrev()
}

var (
	_ IStatsAbsValueAndPrev = (*StatsAbsValueAndPrev)(nil)
)

type StatsAbsValueAndPrev struct {
	Current  float64
	Previous float64
}

func (*StatsAbsValueAndPrev) CRC() uint32 {
	return 0xcb43acde
}
func (*StatsAbsValueAndPrev) _IStatsAbsValueAndPrev() {}

type IStatsDateRangeDays interface {
	tl.Object
	_IStatsDateRangeDays()
}

var (
	_ IStatsDateRangeDays = (*StatsDateRangeDays)(nil)
)

type StatsDateRangeDays struct {
	MinDate int32
	MaxDate int32
}

func (*StatsDateRangeDays) CRC() uint32 {
	return 0xb637edaf
}
func (*StatsDateRangeDays) _IStatsDateRangeDays() {}

// Channel statistics graph
type IStatsGraph interface {
	tl.Object
	_IStatsGraph()
}

var (
	_ IStatsGraph = (*StatsGraphAsync)(nil)
	_ IStatsGraph = (*StatsGraphError)(nil)
	_ IStatsGraph = (*StatsGraph)(nil)
)

// This [channel statistics graph](https://core.telegram.org/api/stats) must be generated asynchronously using StatsLoadAsyncGraph to reduce server load
type StatsGraphAsync struct {
	Token string // Token to use for fetching the async graph
}

func (*StatsGraphAsync) CRC() uint32 {
	return 0x4a27eb2d
}
func (*StatsGraphAsync) _IStatsGraph() {}

// An error occurred while generating the statistics graph
type StatsGraphError struct {
	Error string // The error
}

func (*StatsGraphError) CRC() uint32 {
	return 0xbedc9822
}
func (*StatsGraphError) _IStatsGraph() {}

type StatsGraph struct {
	_         struct{}  `tl:"flags,bitflag"`
	JSON      IDataJSON // Statistics data
	ZoomToken *string   `tl:",omitempty:flags:0"` // Zoom token
}

func (*StatsGraph) CRC() uint32 {
	return 0x8ea464b6
}
func (*StatsGraph) _IStatsGraph() {}

type IStatsGroupTopAdmin interface {
	tl.Object
	_IStatsGroupTopAdmin()
}

var (
	_ IStatsGroupTopAdmin = (*StatsGroupTopAdmin)(nil)
)

type StatsGroupTopAdmin struct {
	UserID  int64
	Deleted int32
	Kicked  int32
	Banned  int32
}

func (*StatsGroupTopAdmin) CRC() uint32 {
	return 0xd7584c87
}
func (*StatsGroupTopAdmin) _IStatsGroupTopAdmin() {}

type IStatsGroupTopInviter interface {
	tl.Object
	_IStatsGroupTopInviter()
}

var (
	_ IStatsGroupTopInviter = (*StatsGroupTopInviter)(nil)
)

type StatsGroupTopInviter struct {
	UserID      int64
	Invitations int32
}

func (*StatsGroupTopInviter) CRC() uint32 {
	return 0x535f779d
}
func (*StatsGroupTopInviter) _IStatsGroupTopInviter() {}

type IStatsGroupTopPoster interface {
	tl.Object
	_IStatsGroupTopPoster()
}

var (
	_ IStatsGroupTopPoster = (*StatsGroupTopPoster)(nil)
)

type StatsGroupTopPoster struct {
	UserID   int64
	Messages int32
	AvgChars int32
}

func (*StatsGroupTopPoster) CRC() uint32 {
	return 0x9d04af9b
}
func (*StatsGroupTopPoster) _IStatsGroupTopPoster() {}

type IStatsPercentValue interface {
	tl.Object
	_IStatsPercentValue()
}

var (
	_ IStatsPercentValue = (*StatsPercentValue)(nil)
)

type StatsPercentValue struct {
	Part  float64
	Total float64
}

func (*StatsPercentValue) CRC() uint32 {
	return 0xcbce2fe0
}
func (*StatsPercentValue) _IStatsPercentValue() {}

type IStatsURL interface {
	tl.Object
	_IStatsURL()
}

var (
	_ IStatsURL = (*StatsURL)(nil)
)

type StatsURL struct {
	URL string
}

func (*StatsURL) CRC() uint32 {
	return 0x47a971e0
}
func (*StatsURL) _IStatsURL() {}

type IStickerKeyword interface {
	tl.Object
	_IStickerKeyword()
}

var (
	_ IStickerKeyword = (*StickerKeyword)(nil)
)

type StickerKeyword struct {
	DocumentID int64
	Keyword    []string
}

func (*StickerKeyword) CRC() uint32 {
	return 0xfcfeb29c
}
func (*StickerKeyword) _IStickerKeyword() {}

type IStickerPack interface {
	tl.Object
	_IStickerPack()
}

var (
	_ IStickerPack = (*StickerPack)(nil)
)

type StickerPack struct {
	Emoticon  string
	Documents []int64
}

func (*StickerPack) CRC() uint32 {
	return 0x12b299d4
}
func (*StickerPack) _IStickerPack() {}

type IStickerSet interface {
	tl.Object
	_IStickerSet()
}

var (
	_ IStickerSet = (*StickerSet)(nil)
)

type StickerSet struct {
	_                  struct{} `tl:"flags,bitflag"`
	Archived           bool     `tl:",omitempty:flags:1,implicit"`
	Official           bool     `tl:",omitempty:flags:2,implicit"`
	Masks              bool     `tl:",omitempty:flags:3,implicit"`
	Animated           bool     `tl:",omitempty:flags:5,implicit"`
	Videos             bool     `tl:",omitempty:flags:6,implicit"`
	Emojis             bool     `tl:",omitempty:flags:7,implicit"`
	TextColor          bool     `tl:",omitempty:flags:9,implicit"`
	ChannelEmojiStatus bool     `tl:",omitempty:flags:10,implicit"`
	InstalledDate      *int32   `tl:",omitempty:flags:0"`
	ID                 int64
	AccessHash         int64
	Title              string
	ShortName          string
	Thumbs             []IPhotoSize `tl:",omitempty:flags:4"`
	ThumbDcID          *int32       `tl:",omitempty:flags:4"`
	ThumbVersion       *int32       `tl:",omitempty:flags:4"`
	ThumbDocumentID    *int64       `tl:",omitempty:flags:8"`
	Count              int32
	Hash               int32
}

func (*StickerSet) CRC() uint32 {
	return 0x2dd14edc
}
func (*StickerSet) _IStickerSet() {}

type IStickerSetCovered interface {
	tl.Object
	_IStickerSetCovered()
}

var (
	_ IStickerSetCovered = (*StickerSetCovered)(nil)
	_ IStickerSetCovered = (*StickerSetMultiCovered)(nil)
	_ IStickerSetCovered = (*StickerSetFullCovered)(nil)
	_ IStickerSetCovered = (*StickerSetNoCovered)(nil)
)

type StickerSetCovered struct {
	Set   IStickerSet
	Cover IDocument
}

func (*StickerSetCovered) CRC() uint32 {
	return 0x6410a5d2
}
func (*StickerSetCovered) _IStickerSetCovered() {}

type StickerSetMultiCovered struct {
	Set    IStickerSet
	Covers []IDocument
}

func (*StickerSetMultiCovered) CRC() uint32 {
	return 0x3407e51b
}
func (*StickerSetMultiCovered) _IStickerSetCovered() {}

type StickerSetFullCovered struct {
	Set       IStickerSet
	Packs     []IStickerPack
	Keywords  []IStickerKeyword
	Documents []IDocument
}

func (*StickerSetFullCovered) CRC() uint32 {
	return 0x40d13c0e
}
func (*StickerSetFullCovered) _IStickerSetCovered() {}

type StickerSetNoCovered struct {
	Set IStickerSet
}

func (*StickerSetNoCovered) CRC() uint32 {
	return 0x77b15d1c
}
func (*StickerSetNoCovered) _IStickerSetCovered() {}

type IStoriesStealthMode interface {
	tl.Object
	_IStoriesStealthMode()
}

var (
	_ IStoriesStealthMode = (*StoriesStealthMode)(nil)
)

type StoriesStealthMode struct {
	_                 struct{} `tl:"flags,bitflag"`
	ActiveUntilDate   *int32   `tl:",omitempty:flags:0"`
	CooldownUntilDate *int32   `tl:",omitempty:flags:1"`
}

func (*StoriesStealthMode) CRC() uint32 {
	return 0x712e27fd
}
func (*StoriesStealthMode) _IStoriesStealthMode() {}

type IStoryFwdHeader interface {
	tl.Object
	_IStoryFwdHeader()
}

var (
	_ IStoryFwdHeader = (*StoryFwdHeader)(nil)
)

type StoryFwdHeader struct {
	_        struct{} `tl:"flags,bitflag"`
	Modified bool     `tl:",omitempty:flags:3,implicit"`
	From     IPeer    `tl:",omitempty:flags:0"`
	FromName *string  `tl:",omitempty:flags:1"`
	StoryID  *int32   `tl:",omitempty:flags:2"`
}

func (*StoryFwdHeader) CRC() uint32 {
	return 0xb826e150
}
func (*StoryFwdHeader) _IStoryFwdHeader() {}

type IStoryItem interface {
	tl.Object
	_IStoryItem()
}

var (
	_ IStoryItem = (*StoryItemDeleted)(nil)
	_ IStoryItem = (*StoryItemSkipped)(nil)
	_ IStoryItem = (*StoryItem)(nil)
)

type StoryItemDeleted struct {
	ID int32
}

func (*StoryItemDeleted) CRC() uint32 {
	return 0x51e6ee4f
}
func (*StoryItemDeleted) _IStoryItem() {}

type StoryItemSkipped struct {
	_            struct{} `tl:"flags,bitflag"`
	CloseFriends bool     `tl:",omitempty:flags:8,implicit"`
	ID           int32
	Date         int32
	ExpireDate   int32
}

func (*StoryItemSkipped) CRC() uint32 {
	return 0xffadc913
}
func (*StoryItemSkipped) _IStoryItem() {}

type StoryItem struct {
	_                struct{} `tl:"flags,bitflag"`
	Pinned           bool     `tl:",omitempty:flags:5,implicit"`
	Public           bool     `tl:",omitempty:flags:7,implicit"`
	CloseFriends     bool     `tl:",omitempty:flags:8,implicit"`
	Min              bool     `tl:",omitempty:flags:9,implicit"`
	Noforwards       bool     `tl:",omitempty:flags:10,implicit"`
	Edited           bool     `tl:",omitempty:flags:11,implicit"`
	Contacts         bool     `tl:",omitempty:flags:12,implicit"`
	SelectedContacts bool     `tl:",omitempty:flags:13,implicit"`
	Out              bool     `tl:",omitempty:flags:16,implicit"`
	ID               int32
	Date             int32
	FwdFrom          IStoryFwdHeader `tl:",omitempty:flags:17"`
	ExpireDate       int32
	Caption          *string          `tl:",omitempty:flags:0"`
	Entities         []IMessageEntity `tl:",omitempty:flags:1"`
	Media            IMessageMedia
	MediaAreas       []IMediaArea   `tl:",omitempty:flags:14"`
	Privacy          []IPrivacyRule `tl:",omitempty:flags:2"`
	Views            IStoryViews    `tl:",omitempty:flags:3"`
	SentReaction     IReaction      `tl:",omitempty:flags:15"`
}

func (*StoryItem) CRC() uint32 {
	return 0xaf6365a1
}
func (*StoryItem) _IStoryItem() {}

type IStoryReaction interface {
	tl.Object
	_IStoryReaction()
}

var (
	_ IStoryReaction = (*StoryReaction)(nil)
	_ IStoryReaction = (*StoryReactionPublicForward)(nil)
	_ IStoryReaction = (*StoryReactionPublicRepost)(nil)
)

type StoryReaction struct {
	PeerID   IPeer
	Date     int32
	Reaction IReaction
}

func (*StoryReaction) CRC() uint32 {
	return 0x6090d6d5
}
func (*StoryReaction) _IStoryReaction() {}

type StoryReactionPublicForward struct {
	Message IMessage
}

func (*StoryReactionPublicForward) CRC() uint32 {
	return 0xbbab2643
}
func (*StoryReactionPublicForward) _IStoryReaction() {}

type StoryReactionPublicRepost struct {
	PeerID IPeer
	Story  IStoryItem
}

func (*StoryReactionPublicRepost) CRC() uint32 {
	return 0xcfcd0f13
}
func (*StoryReactionPublicRepost) _IStoryReaction() {}

type IStoryView interface {
	tl.Object
	_IStoryView()
}

var (
	_ IStoryView = (*StoryView)(nil)
	_ IStoryView = (*StoryViewPublicForward)(nil)
	_ IStoryView = (*StoryViewPublicRepost)(nil)
)

type StoryView struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	UserID               int64
	Date                 int32
	Reaction             IReaction `tl:",omitempty:flags:2"`
}

func (*StoryView) CRC() uint32 {
	return 0xb0bdeac5
}
func (*StoryView) _IStoryView() {}

type StoryViewPublicForward struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	Message              IMessage
}

func (*StoryViewPublicForward) CRC() uint32 {
	return 0x9083670b
}
func (*StoryViewPublicForward) _IStoryView() {}

type StoryViewPublicRepost struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	PeerID               IPeer
	Story                IStoryItem
}

func (*StoryViewPublicRepost) CRC() uint32 {
	return 0xbd74cf49
}
func (*StoryViewPublicRepost) _IStoryView() {}

type IStoryViews interface {
	tl.Object
	_IStoryViews()
}

var (
	_ IStoryViews = (*StoryViews)(nil)
)

type StoryViews struct {
	_              struct{} `tl:"flags,bitflag"`
	HasViewers     bool     `tl:",omitempty:flags:1,implicit"`
	ViewsCount     int32
	ForwardsCount  *int32           `tl:",omitempty:flags:2"`
	Reactions      []IReactionCount `tl:",omitempty:flags:3"`
	ReactionsCount *int32           `tl:",omitempty:flags:4"`
	RecentViewers  []int64          `tl:",omitempty:flags:0"`
}

func (*StoryViews) CRC() uint32 {
	return 0x8d595cd6
}
func (*StoryViews) _IStoryViews() {}

type ITextWithEntities interface {
	tl.Object
	_ITextWithEntities()
}

var (
	_ ITextWithEntities = (*TextWithEntities)(nil)
)

type TextWithEntities struct {
	Text     string
	Entities []IMessageEntity
}

func (*TextWithEntities) CRC() uint32 {
	return 0x751f3146
}
func (*TextWithEntities) _ITextWithEntities() {}

type ITheme interface {
	tl.Object
	_ITheme()
}

var (
	_ ITheme = (*Theme)(nil)
)

type Theme struct {
	_             struct{} `tl:"flags,bitflag"`
	Creator       bool     `tl:",omitempty:flags:0,implicit"`
	Default       bool     `tl:",omitempty:flags:1,implicit"`
	ForChat       bool     `tl:",omitempty:flags:5,implicit"`
	ID            int64
	AccessHash    int64
	Slug          string
	Title         string
	Document      IDocument        `tl:",omitempty:flags:2"`
	Settings      []IThemeSettings `tl:",omitempty:flags:3"`
	Emoticon      *string          `tl:",omitempty:flags:6"`
	InstallsCount *int32           `tl:",omitempty:flags:4"`
}

func (*Theme) CRC() uint32 {
	return 0xa00e67d6
}
func (*Theme) _ITheme() {}

type IThemeSettings interface {
	tl.Object
	_IThemeSettings()
}

var (
	_ IThemeSettings = (*ThemeSettings)(nil)
)

type ThemeSettings struct {
	_                     struct{} `tl:"flags,bitflag"`
	MessageColorsAnimated bool     `tl:",omitempty:flags:2,implicit"`
	BaseTheme             EnumBaseTheme
	AccentColor           int32
	OutboxAccentColor     *int32     `tl:",omitempty:flags:3"`
	MessageColors         []int32    `tl:",omitempty:flags:0"`
	Wallpaper             IWallPaper `tl:",omitempty:flags:1"`
}

func (*ThemeSettings) CRC() uint32 {
	return 0xfa58b6d4
}
func (*ThemeSettings) _IThemeSettings() {}

type ITopPeer interface {
	tl.Object
	_ITopPeer()
}

var (
	_ ITopPeer = (*TopPeer)(nil)
)

type TopPeer struct {
	Peer   IPeer
	Rating float64
}

func (*TopPeer) CRC() uint32 {
	return 0xedcdc05b
}
func (*TopPeer) _ITopPeer() {}

type ITopPeerCategoryPeers interface {
	tl.Object
	_ITopPeerCategoryPeers()
}

var (
	_ ITopPeerCategoryPeers = (*TopPeerCategoryPeers)(nil)
)

type TopPeerCategoryPeers struct {
	Category EnumTopPeerCategory
	Count    int32
	Peers    []ITopPeer
}

func (*TopPeerCategoryPeers) CRC() uint32 {
	return 0xfb834291
}
func (*TopPeerCategoryPeers) _ITopPeerCategoryPeers() {}

// Object contains info on events occured.
type IUpdate interface {
	tl.Object
	_IUpdate()
}

var (
	_ IUpdate = (*UpdateNewMessage)(nil)
	_ IUpdate = (*UpdateMessageID)(nil)
	_ IUpdate = (*UpdateDeleteMessages)(nil)
	_ IUpdate = (*UpdateUserTyping)(nil)
	_ IUpdate = (*UpdateChatUserTyping)(nil)
	_ IUpdate = (*UpdateChatParticipants)(nil)
	_ IUpdate = (*UpdateUserStatus)(nil)
	_ IUpdate = (*UpdateUserName)(nil)
	_ IUpdate = (*UpdateNewAuthorization)(nil)
	_ IUpdate = (*UpdateNewEncryptedMessage)(nil)
	_ IUpdate = (*UpdateEncryptedChatTyping)(nil)
	_ IUpdate = (*UpdateEncryption)(nil)
	_ IUpdate = (*UpdateEncryptedMessagesRead)(nil)
	_ IUpdate = (*UpdateChatParticipantAdd)(nil)
	_ IUpdate = (*UpdateChatParticipantDelete)(nil)
	_ IUpdate = (*UpdateDcOptions)(nil)
	_ IUpdate = (*UpdateNotifySettings)(nil)
	_ IUpdate = (*UpdateServiceNotification)(nil)
	_ IUpdate = (*UpdatePrivacy)(nil)
	_ IUpdate = (*UpdateUserPhone)(nil)
	_ IUpdate = (*UpdateReadHistoryInbox)(nil)
	_ IUpdate = (*UpdateReadHistoryOutbox)(nil)
	_ IUpdate = (*UpdateWebPage)(nil)
	_ IUpdate = (*UpdateReadMessagesContents)(nil)
	_ IUpdate = (*UpdateChannelTooLong)(nil)
	_ IUpdate = (*UpdateChannel)(nil)
	_ IUpdate = (*UpdateNewChannelMessage)(nil)
	_ IUpdate = (*UpdateReadChannelInbox)(nil)
	_ IUpdate = (*UpdateDeleteChannelMessages)(nil)
	_ IUpdate = (*UpdateChannelMessageViews)(nil)
	_ IUpdate = (*UpdateChatParticipantAdmin)(nil)
	_ IUpdate = (*UpdateNewStickerSet)(nil)
	_ IUpdate = (*UpdateStickerSetsOrder)(nil)
	_ IUpdate = (*UpdateStickerSets)(nil)
	_ IUpdate = (*UpdateSavedGifs)(nil)
	_ IUpdate = (*UpdateBotInlineQuery)(nil)
	_ IUpdate = (*UpdateBotInlineSend)(nil)
	_ IUpdate = (*UpdateEditChannelMessage)(nil)
	_ IUpdate = (*UpdateBotCallbackQuery)(nil)
	_ IUpdate = (*UpdateEditMessage)(nil)
	_ IUpdate = (*UpdateInlineBotCallbackQuery)(nil)
	_ IUpdate = (*UpdateReadChannelOutbox)(nil)
	_ IUpdate = (*UpdateDraftMessage)(nil)
	_ IUpdate = (*UpdateReadFeaturedStickers)(nil)
	_ IUpdate = (*UpdateRecentStickers)(nil)
	_ IUpdate = (*UpdateConfig)(nil)
	_ IUpdate = (*UpdatePtsChanged)(nil)
	_ IUpdate = (*UpdateChannelWebPage)(nil)
	_ IUpdate = (*UpdateDialogPinned)(nil)
	_ IUpdate = (*UpdatePinnedDialogs)(nil)
	_ IUpdate = (*UpdateBotWebhookJSON)(nil)
	_ IUpdate = (*UpdateBotWebhookJSONQuery)(nil)
	_ IUpdate = (*UpdateBotShippingQuery)(nil)
	_ IUpdate = (*UpdateBotPrecheckoutQuery)(nil)
	_ IUpdate = (*UpdatePhoneCall)(nil)
	_ IUpdate = (*UpdateLangPackTooLong)(nil)
	_ IUpdate = (*UpdateLangPack)(nil)
	_ IUpdate = (*UpdateFavedStickers)(nil)
	_ IUpdate = (*UpdateChannelReadMessagesContents)(nil)
	_ IUpdate = (*UpdateContactsReset)(nil)
	_ IUpdate = (*UpdateChannelAvailableMessages)(nil)
	_ IUpdate = (*UpdateDialogUnreadMark)(nil)
	_ IUpdate = (*UpdateMessagePoll)(nil)
	_ IUpdate = (*UpdateChatDefaultBannedRights)(nil)
	_ IUpdate = (*UpdateFolderPeers)(nil)
	_ IUpdate = (*UpdatePeerSettings)(nil)
	_ IUpdate = (*UpdatePeerLocated)(nil)
	_ IUpdate = (*UpdateNewScheduledMessage)(nil)
	_ IUpdate = (*UpdateDeleteScheduledMessages)(nil)
	_ IUpdate = (*UpdateTheme)(nil)
	_ IUpdate = (*UpdateGeoLiveViewed)(nil)
	_ IUpdate = (*UpdateLoginToken)(nil)
	_ IUpdate = (*UpdateMessagePollVote)(nil)
	_ IUpdate = (*UpdateDialogFilter)(nil)
	_ IUpdate = (*UpdateDialogFilterOrder)(nil)
	_ IUpdate = (*UpdateDialogFilters)(nil)
	_ IUpdate = (*UpdatePhoneCallSignalingData)(nil)
	_ IUpdate = (*UpdateChannelMessageForwards)(nil)
	_ IUpdate = (*UpdateReadChannelDiscussionInbox)(nil)
	_ IUpdate = (*UpdateReadChannelDiscussionOutbox)(nil)
	_ IUpdate = (*UpdatePeerBlocked)(nil)
	_ IUpdate = (*UpdateChannelUserTyping)(nil)
	_ IUpdate = (*UpdatePinnedMessages)(nil)
	_ IUpdate = (*UpdatePinnedChannelMessages)(nil)
	_ IUpdate = (*UpdateChat)(nil)
	_ IUpdate = (*UpdateGroupCallParticipants)(nil)
	_ IUpdate = (*UpdateGroupCall)(nil)
	_ IUpdate = (*UpdatePeerHistoryTTL)(nil)
	_ IUpdate = (*UpdateChatParticipant)(nil)
	_ IUpdate = (*UpdateChannelParticipant)(nil)
	_ IUpdate = (*UpdateBotStopped)(nil)
	_ IUpdate = (*UpdateGroupCallConnection)(nil)
	_ IUpdate = (*UpdateBotCommands)(nil)
	_ IUpdate = (*UpdatePendingJoinRequests)(nil)
	_ IUpdate = (*UpdateBotChatInviteRequester)(nil)
	_ IUpdate = (*UpdateMessageReactions)(nil)
	_ IUpdate = (*UpdateAttachMenuBots)(nil)
	_ IUpdate = (*UpdateWebViewResultSent)(nil)
	_ IUpdate = (*UpdateBotMenuButton)(nil)
	_ IUpdate = (*UpdateSavedRingtones)(nil)
	_ IUpdate = (*UpdateTranscribedAudio)(nil)
	_ IUpdate = (*UpdateReadFeaturedEmojiStickers)(nil)
	_ IUpdate = (*UpdateUserEmojiStatus)(nil)
	_ IUpdate = (*UpdateRecentEmojiStatuses)(nil)
	_ IUpdate = (*UpdateRecentReactions)(nil)
	_ IUpdate = (*UpdateMoveStickerSetToTop)(nil)
	_ IUpdate = (*UpdateMessageExtendedMedia)(nil)
	_ IUpdate = (*UpdateChannelPinnedTopic)(nil)
	_ IUpdate = (*UpdateChannelPinnedTopics)(nil)
	_ IUpdate = (*UpdateUser)(nil)
	_ IUpdate = (*UpdateAutoSaveSettings)(nil)
	_ IUpdate = (*UpdateGroupInvitePrivacyForbidden)(nil)
	_ IUpdate = (*UpdateStory)(nil)
	_ IUpdate = (*UpdateReadStories)(nil)
	_ IUpdate = (*UpdateStoryID)(nil)
	_ IUpdate = (*UpdateStoriesStealthMode)(nil)
	_ IUpdate = (*UpdateSentStoryReaction)(nil)
	_ IUpdate = (*UpdateBotChatBoost)(nil)
	_ IUpdate = (*UpdateChannelViewForumAsMessages)(nil)
	_ IUpdate = (*UpdatePeerWallpaper)(nil)
	_ IUpdate = (*UpdateBotMessageReaction)(nil)
	_ IUpdate = (*UpdateBotMessageReactions)(nil)
)

// New message.
type UpdateNewMessage struct {
	Message  IMessage
	Pts      int32
	PtsCount int32
}

func (*UpdateNewMessage) CRC() uint32 {
	return 0x1f2b0afd
}
func (*UpdateNewMessage) _IUpdate() {}

// Sent message with random_id client identifier was assigned an identifier.
type UpdateMessageID struct {
	ID       int32
	RandomID int64
}

func (*UpdateMessageID) CRC() uint32 {
	return 0x4e90bfd6
}
func (*UpdateMessageID) _IUpdate() {}

// Messages were deleted.
type UpdateDeleteMessages struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

func (*UpdateDeleteMessages) CRC() uint32 {
	return 0xa20db0e5
}
func (*UpdateDeleteMessages) _IUpdate() {}

// The user is preparing a message; typing, recording, uploading, etc. This update is valid for 6 seconds. If no repeated update received after 6 seconds, it should be considered that the user stopped doing whatever he's been doing.
type UpdateUserTyping struct {
	UserID int64
	Action ISendMessageAction
}

func (*UpdateUserTyping) CRC() uint32 {
	return 0xc01e857f
}
func (*UpdateUserTyping) _IUpdate() {}

// The user is preparing a message in a group; typing, recording, uploading, etc. This update is valid for 6 seconds. If no repeated update received after 6 seconds, it should be considered that the user stopped doing whatever he's been doing.
type UpdateChatUserTyping struct {
	ChatID int64
	FromID IPeer
	Action ISendMessageAction
}

func (*UpdateChatUserTyping) CRC() uint32 {
	return 0x83487af0
}
func (*UpdateChatUserTyping) _IUpdate() {}

// Composition of chat participants changed.
type UpdateChatParticipants struct {
	Participants IChatParticipants
}

func (*UpdateChatParticipants) CRC() uint32 {
	return 0x7761198
}
func (*UpdateChatParticipants) _IUpdate() {}

// Contact status update.
type UpdateUserStatus struct {
	UserID int64
	Status IUserStatus
}

func (*UpdateUserStatus) CRC() uint32 {
	return 0xe5bdf8de
}
func (*UpdateUserStatus) _IUpdate() {}

// Changes the user's first name, last name and username.
type UpdateUserName struct {
	UserID    int64
	FirstName string
	LastName  string
	Usernames []IUsername
}

func (*UpdateUserName) CRC() uint32 {
	return 0xa7848924
}
func (*UpdateUserName) _IUpdate() {}

type UpdateNewAuthorization struct {
	_           struct{} `tl:"flags,bitflag"`
	Unconfirmed bool     `tl:",omitempty:flags:0,implicit"`
	Hash        int64
	Date        *int32  `tl:",omitempty:flags:0"`
	Device      *string `tl:",omitempty:flags:0"`
	Location    *string `tl:",omitempty:flags:0"`
}

func (*UpdateNewAuthorization) CRC() uint32 {
	return 0x8951abef
}
func (*UpdateNewAuthorization) _IUpdate() {}

// New encrypted message.
type UpdateNewEncryptedMessage struct {
	Message IEncryptedMessage
	Qts     int32
}

func (*UpdateNewEncryptedMessage) CRC() uint32 {
	return 0x12bcbd9a
}
func (*UpdateNewEncryptedMessage) _IUpdate() {}

// Interlocutor is typing a message in an encrypted chat. Update period is 6 second. If upon this time there is no repeated update, it shall be considered that the interlocutor stopped typing.
type UpdateEncryptedChatTyping struct {
	ChatID int32
}

func (*UpdateEncryptedChatTyping) CRC() uint32 {
	return 0x1710f156
}
func (*UpdateEncryptedChatTyping) _IUpdate() {}

// Change of state in an encrypted chat.
type UpdateEncryption struct {
	Chat IEncryptedChat
	Date int32
}

func (*UpdateEncryption) CRC() uint32 {
	return 0xb4a2e88d
}
func (*UpdateEncryption) _IUpdate() {}

// Communication history in an encrypted chat was marked as read.
type UpdateEncryptedMessagesRead struct {
	ChatID  int32
	MaxDate int32
	Date    int32
}

func (*UpdateEncryptedMessagesRead) CRC() uint32 {
	return 0x38fe25b7
}
func (*UpdateEncryptedMessagesRead) _IUpdate() {}

// New group member.
type UpdateChatParticipantAdd struct {
	ChatID    int64
	UserID    int64
	InviterID int64
	Date      int32
	Version   int32
}

func (*UpdateChatParticipantAdd) CRC() uint32 {
	return 0x3dda5451
}
func (*UpdateChatParticipantAdd) _IUpdate() {}

// A member has left the group.
type UpdateChatParticipantDelete struct {
	ChatID  int64
	UserID  int64
	Version int32
}

func (*UpdateChatParticipantDelete) CRC() uint32 {
	return 0xe32f3d77
}
func (*UpdateChatParticipantDelete) _IUpdate() {}

// Changes in the data center configuration options.
type UpdateDcOptions struct {
	DcOptions []IDcOption
}

func (*UpdateDcOptions) CRC() uint32 {
	return 0x8e5e9873
}
func (*UpdateDcOptions) _IUpdate() {}

// Changes in notification settings.
type UpdateNotifySettings struct {
	Peer           INotifyPeer
	NotifySettings IPeerNotifySettings
}

func (*UpdateNotifySettings) CRC() uint32 {
	return 0xbec268ef
}
func (*UpdateNotifySettings) _IUpdate() {}

// A service message for the user. The app must show the message to the user upon receiving this update. In case the popup parameter was passed, the text message must be displayed in a popup alert immediately upon receipt. It is recommended to handle the text as you would an ordinary message in terms of highlighting links, etc. The message must also be stored locally as part of the message history with the user id 777000 (Telegram Notifications).
type UpdateServiceNotification struct {
	_           struct{} `tl:"flags,bitflag"`
	Popup       bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:2,implicit"`
	InboxDate   *int32   `tl:",omitempty:flags:1"`
	Type        string
	Message     string
	Media       IMessageMedia
	Entities    []IMessageEntity
}

func (*UpdateServiceNotification) CRC() uint32 {
	return 0xebe46819
}
func (*UpdateServiceNotification) _IUpdate() {}

// Privacy rules were changed
type UpdatePrivacy struct {
	Key   EnumPrivacyKey
	Rules []IPrivacyRule
}

func (*UpdatePrivacy) CRC() uint32 {
	return 0xee3b272a
}
func (*UpdatePrivacy) _IUpdate() {}

// A user's phone number was changed
type UpdateUserPhone struct {
	UserID int64
	Phone  string
}

func (*UpdateUserPhone) CRC() uint32 {
	return 0x5492a13
}
func (*UpdateUserPhone) _IUpdate() {}

// Incoming messages were read
type UpdateReadHistoryInbox struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:",omitempty:flags:0"`
	Peer             IPeer
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
	PtsCount         int32
}

func (*UpdateReadHistoryInbox) CRC() uint32 {
	return 0x9c974fdf
}
func (*UpdateReadHistoryInbox) _IUpdate() {}

// Outgoing messages were read
type UpdateReadHistoryOutbox struct {
	Peer     IPeer
	MaxID    int32
	Pts      int32
	PtsCount int32
}

func (*UpdateReadHistoryOutbox) CRC() uint32 {
	return 0x2f2f21bf
}
func (*UpdateReadHistoryOutbox) _IUpdate() {}

// An (IV) webpage preview was generated
type UpdateWebPage struct {
	Webpage  IWebPage
	Pts      int32
	PtsCount int32
}

func (*UpdateWebPage) CRC() uint32 {
	return 0x7f891213
}
func (*UpdateWebPage) _IUpdate() {}

// Contents of messages in the common message box were read
type UpdateReadMessagesContents struct {
	_        struct{} `tl:"flags,bitflag"`
	Messages []int32
	Pts      int32
	PtsCount int32
	Date     *int32 `tl:",omitempty:flags:0"`
}

func (*UpdateReadMessagesContents) CRC() uint32 {
	return 0xf8227181
}
func (*UpdateReadMessagesContents) _IUpdate() {}

// There are new updates in the specified channel, the client must fetch them, eventually starting the specified pts if the difference is too long or if the channel isn't currently in the states.
type UpdateChannelTooLong struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	Pts       *int32 `tl:",omitempty:flags:0"`
}

func (*UpdateChannelTooLong) CRC() uint32 {
	return 0x108d941f
}
func (*UpdateChannelTooLong) _IUpdate() {}

// A new channel is available
type UpdateChannel struct {
	ChannelID int64
}

func (*UpdateChannel) CRC() uint32 {
	return 0x635b4c09
}
func (*UpdateChannel) _IUpdate() {}

// A new message was sent in a channel/supergroup
type UpdateNewChannelMessage struct {
	Message  IMessage
	Pts      int32
	PtsCount int32
}

func (*UpdateNewChannelMessage) CRC() uint32 {
	return 0x62ba04d9
}
func (*UpdateNewChannelMessage) _IUpdate() {}

// Incoming messages in a channel/supergroup were read
type UpdateReadChannelInbox struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:",omitempty:flags:0"`
	ChannelID        int64
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
}

func (*UpdateReadChannelInbox) CRC() uint32 {
	return 0x922e6e10
}
func (*UpdateReadChannelInbox) _IUpdate() {}

// Some messages in a supergroup/channel were deleted
type UpdateDeleteChannelMessages struct {
	ChannelID int64
	Messages  []int32
	Pts       int32
	PtsCount  int32
}

func (*UpdateDeleteChannelMessages) CRC() uint32 {
	return 0xc32d5b12
}
func (*UpdateDeleteChannelMessages) _IUpdate() {}

// The view counter of a message in a channel has changed
type UpdateChannelMessageViews struct {
	ChannelID int64
	ID        int32
	Views     int32
}

func (*UpdateChannelMessageViews) CRC() uint32 {
	return 0xf226ac08
}
func (*UpdateChannelMessageViews) _IUpdate() {}

// Admin permissions of a user in a legacy group were changed
type UpdateChatParticipantAdmin struct {
	ChatID  int64
	UserID  int64
	IsAdmin bool
	Version int32
}

func (*UpdateChatParticipantAdmin) CRC() uint32 {
	return 0xd7ca61a2
}
func (*UpdateChatParticipantAdmin) _IUpdate() {}

// A new stickerset was installed
type UpdateNewStickerSet struct {
	Stickerset IMessagesStickerSet
}

func (*UpdateNewStickerSet) CRC() uint32 {
	return 0x688a30aa
}
func (*UpdateNewStickerSet) _IUpdate() {}

// The order of stickersets was changed
type UpdateStickerSetsOrder struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
	Order  []int64
}

func (*UpdateStickerSetsOrder) CRC() uint32 {
	return 0xbb2d201
}
func (*UpdateStickerSetsOrder) _IUpdate() {}

// Installed stickersets have changed, the client should refetch them using messages.getAllStickers
type UpdateStickerSets struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
}

func (*UpdateStickerSets) CRC() uint32 {
	return 0x31c24808
}
func (*UpdateStickerSets) _IUpdate() {}

// The saved gif list has changed, the client should refetch it using messages.getSavedGifs
type UpdateSavedGifs struct{}

func (*UpdateSavedGifs) CRC() uint32 {
	return 0x9375341e
}
func (*UpdateSavedGifs) _IUpdate() {}

// An incoming inline query
type UpdateBotInlineQuery struct {
	_        struct{} `tl:"flags,bitflag"`
	QueryID  int64
	UserID   int64
	Query    string
	Geo      IGeoPoint               `tl:",omitempty:flags:0"`
	PeerType EnumInlineQueryPeerType `tl:",omitempty:flags:1"`
	Offset   string
}

func (*UpdateBotInlineQuery) CRC() uint32 {
	return 0x496f379c
}
func (*UpdateBotInlineQuery) _IUpdate() {}

// The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
type UpdateBotInlineSend struct {
	_      struct{} `tl:"flags,bitflag"`
	UserID int64
	Query  string
	Geo    IGeoPoint `tl:",omitempty:flags:0"`
	ID     string
	MsgID  IInputBotInlineMessageID `tl:",omitempty:flags:1"`
}

func (*UpdateBotInlineSend) CRC() uint32 {
	return 0x12f12a07
}
func (*UpdateBotInlineSend) _IUpdate() {}

// A message was edited in a channel/supergroup
type UpdateEditChannelMessage struct {
	Message  IMessage
	Pts      int32
	PtsCount int32
}

func (*UpdateEditChannelMessage) CRC() uint32 {
	return 0x1b3f4df7
}
func (*UpdateEditChannelMessage) _IUpdate() {}

// A callback button was pressed, and the button data was sent to the bot that created the button
type UpdateBotCallbackQuery struct {
	_             struct{} `tl:"flags,bitflag"`
	QueryID       int64
	UserID        int64
	Peer          IPeer
	MsgID         int32
	ChatInstance  int64
	Data          *[]byte `tl:",omitempty:flags:0"`
	GameShortName *string `tl:",omitempty:flags:1"`
}

func (*UpdateBotCallbackQuery) CRC() uint32 {
	return 0xb9cfc48d
}
func (*UpdateBotCallbackQuery) _IUpdate() {}

// A message was edited
type UpdateEditMessage struct {
	Message  IMessage
	Pts      int32
	PtsCount int32
}

func (*UpdateEditMessage) CRC() uint32 {
	return 0xe40370a3
}
func (*UpdateEditMessage) _IUpdate() {}

// This notification is received by bots when a button is pressed
type UpdateInlineBotCallbackQuery struct {
	_             struct{} `tl:"flags,bitflag"`
	QueryID       int64
	UserID        int64
	MsgID         IInputBotInlineMessageID
	ChatInstance  int64
	Data          *[]byte `tl:",omitempty:flags:0"`
	GameShortName *string `tl:",omitempty:flags:1"`
}

func (*UpdateInlineBotCallbackQuery) CRC() uint32 {
	return 0x691e9052
}
func (*UpdateInlineBotCallbackQuery) _IUpdate() {}

// Outgoing messages in a channel/supergroup were read
type UpdateReadChannelOutbox struct {
	ChannelID int64
	MaxID     int32
}

func (*UpdateReadChannelOutbox) CRC() uint32 {
	return 0xb75f99a9
}
func (*UpdateReadChannelOutbox) _IUpdate() {}

// Notifies a change of a message draft.
type UpdateDraftMessage struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
	Draft    IDraftMessage
}

func (*UpdateDraftMessage) CRC() uint32 {
	return 0x1b49ec6d
}
func (*UpdateDraftMessage) _IUpdate() {}

// Some featured stickers were marked as read
type UpdateReadFeaturedStickers struct{}

func (*UpdateReadFeaturedStickers) CRC() uint32 {
	return 0x571d2742
}
func (*UpdateReadFeaturedStickers) _IUpdate() {}

// The recent sticker list was updated
type UpdateRecentStickers struct{}

func (*UpdateRecentStickers) CRC() uint32 {
	return 0x9a422c20
}
func (*UpdateRecentStickers) _IUpdate() {}

// The server-side configuration has changed; the client should re-fetch the config using help.getConfig
type UpdateConfig struct{}

func (*UpdateConfig) CRC() uint32 {
	return 0xa229dd06
}
func (*UpdateConfig) _IUpdate() {}

// Common message box sequence PTS has changed, state has to be refetched using updates.getState
type UpdatePtsChanged struct{}

func (*UpdatePtsChanged) CRC() uint32 {
	return 0x3354678f
}
func (*UpdatePtsChanged) _IUpdate() {}

// A webpage preview of a link in a channel/supergroup message was generated
type UpdateChannelWebPage struct {
	ChannelID int64
	Webpage   IWebPage
	Pts       int32
	PtsCount  int32
}

func (*UpdateChannelWebPage) CRC() uint32 {
	return 0x2f2ba99f
}
func (*UpdateChannelWebPage) _IUpdate() {}

// A dialog was pinned/unpinned
type UpdateDialogPinned struct {
	_        struct{} `tl:"flags,bitflag"`
	Pinned   bool     `tl:",omitempty:flags:0,implicit"`
	FolderID *int32   `tl:",omitempty:flags:1"`
	Peer     IDialogPeer
}

func (*UpdateDialogPinned) CRC() uint32 {
	return 0x6e6fe51c
}
func (*UpdateDialogPinned) _IUpdate() {}

// Pinned dialogs were updated
type UpdatePinnedDialogs struct {
	_        struct{}      `tl:"flags,bitflag"`
	FolderID *int32        `tl:",omitempty:flags:1"`
	Order    []IDialogPeer `tl:",omitempty:flags:0"`
}

func (*UpdatePinnedDialogs) CRC() uint32 {
	return 0xfa0f3ca2
}
func (*UpdatePinnedDialogs) _IUpdate() {}

// A new incoming event; for bots only
type UpdateBotWebhookJSON struct {
	Data IDataJSON
}

func (*UpdateBotWebhookJSON) CRC() uint32 {
	return 0x8317c0c3
}
func (*UpdateBotWebhookJSON) _IUpdate() {}

// A new incoming query; for bots only
type UpdateBotWebhookJSONQuery struct {
	QueryID int64
	Data    IDataJSON
	Timeout int32
}

func (*UpdateBotWebhookJSONQuery) CRC() uint32 {
	return 0x9b9240a6
}
func (*UpdateBotWebhookJSONQuery) _IUpdate() {}

// This object contains information about an incoming shipping query.
type UpdateBotShippingQuery struct {
	QueryID         int64
	UserID          int64
	Payload         []byte
	ShippingAddress IPostAddress
}

func (*UpdateBotShippingQuery) CRC() uint32 {
	return 0xb5aefd7d
}
func (*UpdateBotShippingQuery) _IUpdate() {}

// This object contains information about an incoming pre-checkout query.
type UpdateBotPrecheckoutQuery struct {
	_                struct{} `tl:"flags,bitflag"`
	QueryID          int64
	UserID           int64
	Payload          []byte
	Info             IPaymentRequestedInfo `tl:",omitempty:flags:0"`
	ShippingOptionID *string               `tl:",omitempty:flags:1"`
	Currency         string
	TotalAmount      int64
}

func (*UpdateBotPrecheckoutQuery) CRC() uint32 {
	return 0x8caa9a96
}
func (*UpdateBotPrecheckoutQuery) _IUpdate() {}

// An incoming phone call
type UpdatePhoneCall struct {
	PhoneCall IPhoneCall
}

func (*UpdatePhoneCall) CRC() uint32 {
	return 0xab0f6b1e
}
func (*UpdatePhoneCall) _IUpdate() {}

// A language pack has changed, the client should manually fetch the changed strings using langpack.getDifference
type UpdateLangPackTooLong struct {
	LangCode string
}

func (*UpdateLangPackTooLong) CRC() uint32 {
	return 0x46560264
}
func (*UpdateLangPackTooLong) _IUpdate() {}

// Language pack updated
type UpdateLangPack struct {
	Difference ILangPackDifference
}

func (*UpdateLangPack) CRC() uint32 {
	return 0x56022f4d
}
func (*UpdateLangPack) _IUpdate() {}

// The list of favorited stickers was changed, the client should call messages.getFavedStickers to refetch the new list
type UpdateFavedStickers struct{}

func (*UpdateFavedStickers) CRC() uint32 {
	return 0xe511996d
}
func (*UpdateFavedStickers) _IUpdate() {}

// The specified channel/supergroup messages were read
type UpdateChannelReadMessagesContents struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	Messages  []int32
}

func (*UpdateChannelReadMessagesContents) CRC() uint32 {
	return 0xea29055d
}
func (*UpdateChannelReadMessagesContents) _IUpdate() {}

// All contacts were deleted
type UpdateContactsReset struct{}

func (*UpdateContactsReset) CRC() uint32 {
	return 0x7084a7be
}
func (*UpdateContactsReset) _IUpdate() {}

// The history of a channel/supergroup was hidden.
type UpdateChannelAvailableMessages struct {
	ChannelID      int64
	AvailableMinID int32
}

func (*UpdateChannelAvailableMessages) CRC() uint32 {
	return 0xb23fc698
}
func (*UpdateChannelAvailableMessages) _IUpdate() {}

// The manual unread mark of a chat was changed
type UpdateDialogUnreadMark struct {
	_      struct{} `tl:"flags,bitflag"`
	Unread bool     `tl:",omitempty:flags:0,implicit"`
	Peer   IDialogPeer
}

func (*UpdateDialogUnreadMark) CRC() uint32 {
	return 0xe16459c3
}
func (*UpdateDialogUnreadMark) _IUpdate() {}

// The results of a poll have changed
type UpdateMessagePoll struct {
	_       struct{} `tl:"flags,bitflag"`
	PollID  int64
	Poll    IPoll `tl:",omitempty:flags:0"`
	Results IPollResults
}

func (*UpdateMessagePoll) CRC() uint32 {
	return 0xaca1657b
}
func (*UpdateMessagePoll) _IUpdate() {}

// Default banned rights in a normal chat were updated
type UpdateChatDefaultBannedRights struct {
	Peer                IPeer
	DefaultBannedRights IChatBannedRights
	Version             int32
}

func (*UpdateChatDefaultBannedRights) CRC() uint32 {
	return 0x54c01850
}
func (*UpdateChatDefaultBannedRights) _IUpdate() {}

// The dialog list of a folder was changed
type UpdateFolderPeers struct {
	FolderPeers []IFolderPeer
	Pts         int32
	PtsCount    int32
}

func (*UpdateFolderPeers) CRC() uint32 {
	return 0x19360dc0
}
func (*UpdateFolderPeers) _IUpdate() {}

// Settings of a certain peer have changed
type UpdatePeerSettings struct {
	Peer     IPeer
	Settings IPeerSettings
}

func (*UpdatePeerSettings) CRC() uint32 {
	return 0x6a7e7366
}
func (*UpdatePeerSettings) _IUpdate() {}

// List of peers near you was updated
type UpdatePeerLocated struct {
	Peers []IPeerLocated
}

func (*UpdatePeerLocated) CRC() uint32 {
	return 0xb4afcfb0
}
func (*UpdatePeerLocated) _IUpdate() {}

// New incoming scheduled message
type UpdateNewScheduledMessage struct {
	Message IMessage
}

func (*UpdateNewScheduledMessage) CRC() uint32 {
	return 0x39a51dfb
}
func (*UpdateNewScheduledMessage) _IUpdate() {}

// Some scheduled messages were deleted
type UpdateDeleteScheduledMessages struct {
	Peer     IPeer
	Messages []int32
}

func (*UpdateDeleteScheduledMessages) CRC() uint32 {
	return 0x90866cee
}
func (*UpdateDeleteScheduledMessages) _IUpdate() {}

// A new cloud theme was installed
type UpdateTheme struct {
	Theme ITheme
}

func (*UpdateTheme) CRC() uint32 {
	return 0x8216fba3
}
func (*UpdateTheme) _IUpdate() {}

// Live geoposition message was viewed
type UpdateGeoLiveViewed struct {
	Peer  IPeer // The user that viewed the live geoposition
	MsgID int32 // Message ID of geoposition message
}

func (*UpdateGeoLiveViewed) CRC() uint32 {
	return 0x871fb939
}
func (*UpdateGeoLiveViewed) _IUpdate() {}

// A login token (for login via QR code) was generated
type UpdateLoginToken struct{}

func (*UpdateLoginToken) CRC() uint32 {
	return 0x564fe691
}
func (*UpdateLoginToken) _IUpdate() {}

// A specific user has voted in a poll
type UpdateMessagePollVote struct {
	PollID  int64 // Poll ID
	Peer    IPeer
	Options [][]byte // Chosen option(s)
	Qts     int32
}

func (*UpdateMessagePollVote) CRC() uint32 {
	return 0x24f40e77
}
func (*UpdateMessagePollVote) _IUpdate() {}

// A new folder was added
type UpdateDialogFilter struct {
	_      struct{}      `tl:"flags,bitflag"`
	ID     int32         // Folder ID
	Filter IDialogFilter `tl:",omitempty:flags:0"` // Folder info
}

func (*UpdateDialogFilter) CRC() uint32 {
	return 0x26ffde7d
}
func (*UpdateDialogFilter) _IUpdate() {}

// New chat folders order
type UpdateDialogFilterOrder struct {
	Order []int32 // Ordered folder IDs
}

func (*UpdateDialogFilterOrder) CRC() uint32 {
	return 0xa5d72105
}
func (*UpdateDialogFilterOrder) _IUpdate() {}

// Update folder list
type UpdateDialogFilters struct{}

func (*UpdateDialogFilters) CRC() uint32 {
	return 0x3504914f
}
func (*UpdateDialogFilters) _IUpdate() {}

// Incoming phone call signaling payload
type UpdatePhoneCallSignalingData struct {
	PhoneCallID int64  // Phone call ID
	Data        []byte // Signaling payload
}

func (*UpdatePhoneCallSignalingData) CRC() uint32 {
	return 0x2661bf09
}
func (*UpdatePhoneCallSignalingData) _IUpdate() {}

// The forward counter of a message in a channel has changed
type UpdateChannelMessageForwards struct {
	ChannelID int64
	ID        int32 // ID of the message
	Forwards  int32 // New forward counter
}

func (*UpdateChannelMessageForwards) CRC() uint32 {
	return 0xd29a27f4
}
func (*UpdateChannelMessageForwards) _IUpdate() {}

// Incoming messages were marked as read
type UpdateReadChannelDiscussionInbox struct {
	_             struct{} `tl:"flags,bitflag"`
	ChannelID     int64    // Discussion group ID
	TopMsgID      int32    // ID of the group message that started the thread (message in linked discussion group)
	ReadMaxID     int32    // Message ID of latest read incoming message for this thread
	BroadcastID   *int64   `tl:",omitempty:flags:0"` // If set, contains the ID of the channel that contains the post that started the comment thread in the discussion group (channel_id)
	BroadcastPost *int32   `tl:",omitempty:flags:0"` // If set, contains the ID of the channel post that started the the comment thread
}

func (*UpdateReadChannelDiscussionInbox) CRC() uint32 {
	return 0xd6b19546
}
func (*UpdateReadChannelDiscussionInbox) _IUpdate() {}

// Outgoing comments in a discussion thread were marked as read
type UpdateReadChannelDiscussionOutbox struct {
	ChannelID int64
	TopMsgID  int32 // ID of the group message that started the thread
	ReadMaxID int32 // Message ID of latest read outgoing message for this thread
}

func (*UpdateReadChannelDiscussionOutbox) CRC() uint32 {
	return 0x695c9e7c
}
func (*UpdateReadChannelDiscussionOutbox) _IUpdate() {}

// A peer was blocked
type UpdatePeerBlocked struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"` // Whether the peer was blocked or unblocked
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	PeerID               IPeer
}

func (*UpdatePeerBlocked) CRC() uint32 {
	return 0xebe07752
}
func (*UpdatePeerBlocked) _IUpdate() {}

// user is typing in a supergroup, channel or message thread
type UpdateChannelUserTyping struct {
	_         struct{}           `tl:"flags,bitflag"`
	ChannelID int64              // Channel ID
	TopMsgID  *int32             `tl:",omitempty:flags:0"` // Thread ID
	FromID    IPeer              // The peer that is typing
	Action    ISendMessageAction // Whether the user is typing, sending a media or doing something else
}

func (*UpdateChannelUserTyping) CRC() uint32 {
	return 0x8c88c923
}
func (*UpdateChannelUserTyping) _IUpdate() {}

// Some messages were pinned in a chat
type UpdatePinnedMessages struct {
	_        struct{} `tl:"flags,bitflag"`
	Pinned   bool     `tl:",omitempty:flags:0,implicit"` // Whether the messages were pinned or unpinned
	Peer     IPeer    // Peer
	Messages []int32  // Message IDs
	Pts      int32    // Event count after generation
	PtsCount int32    // Number of events that were generated
}

func (*UpdatePinnedMessages) CRC() uint32 {
	return 0xed85eab5
}
func (*UpdatePinnedMessages) _IUpdate() {}

// Messages were pinned/unpinned in a channel/supergroup
type UpdatePinnedChannelMessages struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:",omitempty:flags:0,implicit"` // Whether the messages were pinned or unpinned
	ChannelID int64    // Channel ID
	Messages  []int32  // Messages
	Pts       int32    // Event count after generation
	PtsCount  int32    // Number of events that were generated
}

func (*UpdatePinnedChannelMessages) CRC() uint32 {
	return 0x5bb98608
}
func (*UpdatePinnedChannelMessages) _IUpdate() {}

type UpdateChat struct {
	ChatID int64
}

func (*UpdateChat) CRC() uint32 {
	return 0xf89a6a4e
}
func (*UpdateChat) _IUpdate() {}

type UpdateGroupCallParticipants struct {
	Call         IInputGroupCall
	Participants []IGroupCallParticipant
	Version      int32
}

func (*UpdateGroupCallParticipants) CRC() uint32 {
	return 0xf2ebdb4e
}
func (*UpdateGroupCallParticipants) _IUpdate() {}

type UpdateGroupCall struct {
	ChatID int64
	Call   IGroupCall
}

func (*UpdateGroupCall) CRC() uint32 {
	return 0x14b24500
}
func (*UpdateGroupCall) _IUpdate() {}

type UpdatePeerHistoryTTL struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IPeer
	TTLPeriod *int32 `tl:",omitempty:flags:0"`
}

func (*UpdatePeerHistoryTTL) CRC() uint32 {
	return 0xbb9bb9a5
}
func (*UpdatePeerHistoryTTL) _IUpdate() {}

type UpdateChatParticipant struct {
	_               struct{} `tl:"flags,bitflag"`
	ChatID          int64
	Date            int32
	ActorID         int64
	UserID          int64
	PrevParticipant IChatParticipant    `tl:",omitempty:flags:0"`
	NewParticipant  IChatParticipant    `tl:",omitempty:flags:1"`
	Invite          IExportedChatInvite `tl:",omitempty:flags:2"`
	Qts             int32
}

func (*UpdateChatParticipant) CRC() uint32 {
	return 0xd087663a
}
func (*UpdateChatParticipant) _IUpdate() {}

// A participant has left/joined a channel or supergroup.
type UpdateChannelParticipant struct {
	_               struct{} `tl:"flags,bitflag"`
	ViaChatlist     bool     `tl:",omitempty:flags:3,implicit"`
	ChannelID       int64    // Channel ID
	Date            int32    // Date of the event
	ActorID         int64
	UserID          int64               // User in question
	PrevParticipant IChannelParticipant `tl:",omitempty:flags:0"` // Previous participant status
	NewParticipant  IChannelParticipant `tl:",omitempty:flags:1"` // New participant status
	Invite          IExportedChatInvite `tl:",omitempty:flags:2"`
	Qts             int32               // [PTS](https://core.telegram.org/api/updates)
}

func (*UpdateChannelParticipant) CRC() uint32 {
	return 0x985d3abb
}
func (*UpdateChannelParticipant) _IUpdate() {}

type UpdateBotStopped struct {
	UserID  int64
	Date    int32
	Stopped bool
	Qts     int32
}

func (*UpdateBotStopped) CRC() uint32 {
	return 0xc4870a49
}
func (*UpdateBotStopped) _IUpdate() {}

type UpdateGroupCallConnection struct {
	_            struct{} `tl:"flags,bitflag"`
	Presentation bool     `tl:",omitempty:flags:0,implicit"`
	Params       IDataJSON
}

func (*UpdateGroupCallConnection) CRC() uint32 {
	return 0xb783982
}
func (*UpdateGroupCallConnection) _IUpdate() {}

type UpdateBotCommands struct {
	Peer     IPeer
	BotID    int64
	Commands []IBotCommand
}

func (*UpdateBotCommands) CRC() uint32 {
	return 0x4d712f2e
}
func (*UpdateBotCommands) _IUpdate() {}

type UpdatePendingJoinRequests struct {
	Peer             IPeer
	RequestsPending  int32
	RecentRequesters []int64
}

func (*UpdatePendingJoinRequests) CRC() uint32 {
	return 0x7063c3db
}
func (*UpdatePendingJoinRequests) _IUpdate() {}

type UpdateBotChatInviteRequester struct {
	Peer   IPeer
	Date   int32
	UserID int64
	About  string
	Invite IExportedChatInvite
	Qts    int32
}

func (*UpdateBotChatInviteRequester) CRC() uint32 {
	return 0x11dfa986
}
func (*UpdateBotChatInviteRequester) _IUpdate() {}

type UpdateMessageReactions struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IPeer
	MsgID     int32
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	Reactions IMessageReactions
}

func (*UpdateMessageReactions) CRC() uint32 {
	return 0x5e1b3cb8
}
func (*UpdateMessageReactions) _IUpdate() {}

type UpdateAttachMenuBots struct{}

func (*UpdateAttachMenuBots) CRC() uint32 {
	return 0x17b7a20b
}
func (*UpdateAttachMenuBots) _IUpdate() {}

type UpdateWebViewResultSent struct {
	QueryID int64
}

func (*UpdateWebViewResultSent) CRC() uint32 {
	return 0x1592b79d
}
func (*UpdateWebViewResultSent) _IUpdate() {}

type UpdateBotMenuButton struct {
	BotID  int64
	Button IBotMenuButton
}

func (*UpdateBotMenuButton) CRC() uint32 {
	return 0x14b85813
}
func (*UpdateBotMenuButton) _IUpdate() {}

type UpdateSavedRingtones struct{}

func (*UpdateSavedRingtones) CRC() uint32 {
	return 0x74d8be99
}
func (*UpdateSavedRingtones) _IUpdate() {}

type UpdateTranscribedAudio struct {
	_               struct{} `tl:"flags,bitflag"`
	Pending         bool     `tl:",omitempty:flags:0,implicit"`
	Peer            IPeer
	MsgID           int32
	TranscriptionID int64
	Text            string
}

func (*UpdateTranscribedAudio) CRC() uint32 {
	return 0x84cd5a
}
func (*UpdateTranscribedAudio) _IUpdate() {}

type UpdateReadFeaturedEmojiStickers struct{}

func (*UpdateReadFeaturedEmojiStickers) CRC() uint32 {
	return 0xfb4c496c
}
func (*UpdateReadFeaturedEmojiStickers) _IUpdate() {}

type UpdateUserEmojiStatus struct {
	UserID      int64
	EmojiStatus IEmojiStatus
}

func (*UpdateUserEmojiStatus) CRC() uint32 {
	return 0x28373599
}
func (*UpdateUserEmojiStatus) _IUpdate() {}

type UpdateRecentEmojiStatuses struct{}

func (*UpdateRecentEmojiStatuses) CRC() uint32 {
	return 0x30f443db
}
func (*UpdateRecentEmojiStatuses) _IUpdate() {}

type UpdateRecentReactions struct{}

func (*UpdateRecentReactions) CRC() uint32 {
	return 0x6f7863f4
}
func (*UpdateRecentReactions) _IUpdate() {}

type UpdateMoveStickerSetToTop struct {
	_          struct{} `tl:"flags,bitflag"`
	Masks      bool     `tl:",omitempty:flags:0,implicit"`
	Emojis     bool     `tl:",omitempty:flags:1,implicit"`
	Stickerset int64
}

func (*UpdateMoveStickerSetToTop) CRC() uint32 {
	return 0x86fccf85
}
func (*UpdateMoveStickerSetToTop) _IUpdate() {}

type UpdateMessageExtendedMedia struct {
	Peer          IPeer
	MsgID         int32
	ExtendedMedia IMessageExtendedMedia
}

func (*UpdateMessageExtendedMedia) CRC() uint32 {
	return 0x5a73a98c
}
func (*UpdateMessageExtendedMedia) _IUpdate() {}

type UpdateChannelPinnedTopic struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:",omitempty:flags:0,implicit"`
	ChannelID int64
	TopicID   int32
}

func (*UpdateChannelPinnedTopic) CRC() uint32 {
	return 0x192efbe3
}
func (*UpdateChannelPinnedTopic) _IUpdate() {}

type UpdateChannelPinnedTopics struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	Order     []int32 `tl:",omitempty:flags:0"`
}

func (*UpdateChannelPinnedTopics) CRC() uint32 {
	return 0xfe198602
}
func (*UpdateChannelPinnedTopics) _IUpdate() {}

type UpdateUser struct {
	UserID int64
}

func (*UpdateUser) CRC() uint32 {
	return 0x20529438
}
func (*UpdateUser) _IUpdate() {}

type UpdateAutoSaveSettings struct{}

func (*UpdateAutoSaveSettings) CRC() uint32 {
	return 0xec05b097
}
func (*UpdateAutoSaveSettings) _IUpdate() {}

type UpdateGroupInvitePrivacyForbidden struct {
	UserID int64
}

func (*UpdateGroupInvitePrivacyForbidden) CRC() uint32 {
	return 0xccf08ad6
}
func (*UpdateGroupInvitePrivacyForbidden) _IUpdate() {}

type UpdateStory struct {
	Peer  IPeer
	Story IStoryItem
}

func (*UpdateStory) CRC() uint32 {
	return 0x75b3b798
}
func (*UpdateStory) _IUpdate() {}

type UpdateReadStories struct {
	Peer  IPeer
	MaxID int32
}

func (*UpdateReadStories) CRC() uint32 {
	return 0xf74e932b
}
func (*UpdateReadStories) _IUpdate() {}

type UpdateStoryID struct {
	ID       int32
	RandomID int64
}

func (*UpdateStoryID) CRC() uint32 {
	return 0x1bf335b9
}
func (*UpdateStoryID) _IUpdate() {}

type UpdateStoriesStealthMode struct {
	StealthMode IStoriesStealthMode
}

func (*UpdateStoriesStealthMode) CRC() uint32 {
	return 0x2c084dc1
}
func (*UpdateStoriesStealthMode) _IUpdate() {}

type UpdateSentStoryReaction struct {
	Peer     IPeer
	StoryID  int32
	Reaction IReaction
}

func (*UpdateSentStoryReaction) CRC() uint32 {
	return 0x7d627683
}
func (*UpdateSentStoryReaction) _IUpdate() {}

type UpdateBotChatBoost struct {
	Peer  IPeer
	Boost IBoost
	Qts   int32
}

func (*UpdateBotChatBoost) CRC() uint32 {
	return 0x904dd49c
}
func (*UpdateBotChatBoost) _IUpdate() {}

type UpdateChannelViewForumAsMessages struct {
	ChannelID int64
	Enabled   bool
}

func (*UpdateChannelViewForumAsMessages) CRC() uint32 {
	return 0x7b68920
}
func (*UpdateChannelViewForumAsMessages) _IUpdate() {}

type UpdatePeerWallpaper struct {
	_                   struct{} `tl:"flags,bitflag"`
	WallpaperOverridden bool     `tl:",omitempty:flags:1,implicit"`
	Peer                IPeer
	Wallpaper           IWallPaper `tl:",omitempty:flags:0"`
}

func (*UpdatePeerWallpaper) CRC() uint32 {
	return 0xae3f101d
}
func (*UpdatePeerWallpaper) _IUpdate() {}

type UpdateBotMessageReaction struct {
	Peer         IPeer
	MsgID        int32
	Date         int32
	Actor        IPeer
	OldReactions []IReaction
	NewReactions []IReaction
	Qts          int32
}

func (*UpdateBotMessageReaction) CRC() uint32 {
	return 0xac21d3ce
}
func (*UpdateBotMessageReaction) _IUpdate() {}

type UpdateBotMessageReactions struct {
	Peer      IPeer
	MsgID     int32
	Date      int32
	Reactions []IReactionCount
	Qts       int32
}

func (*UpdateBotMessageReactions) CRC() uint32 {
	return 0x9cb7759
}
func (*UpdateBotMessageReactions) _IUpdate() {}

type IUpdates interface {
	tl.Object
	_IUpdates()
}

var (
	_ IUpdates = (*UpdatesTooLong)(nil)
	_ IUpdates = (*UpdateShortMessage)(nil)
	_ IUpdates = (*UpdateShortChatMessage)(nil)
	_ IUpdates = (*UpdateShort)(nil)
	_ IUpdates = (*UpdatesCombined)(nil)
	_ IUpdates = (*Updates)(nil)
	_ IUpdates = (*UpdateShortSentMessage)(nil)
)

type UpdatesTooLong struct{}

func (*UpdatesTooLong) CRC() uint32 {
	return 0xe317af7e
}
func (*UpdatesTooLong) _IUpdates() {}

type UpdateShortMessage struct {
	_           struct{} `tl:"flags,bitflag"`
	Out         bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned   bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread bool     `tl:",omitempty:flags:5,implicit"`
	Silent      bool     `tl:",omitempty:flags:13,implicit"`
	ID          int32
	UserID      int64
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
	FwdFrom     IMessageFwdHeader   `tl:",omitempty:flags:2"`
	ViaBotID    *int64              `tl:",omitempty:flags:11"`
	ReplyTo     IMessageReplyHeader `tl:",omitempty:flags:3"`
	Entities    []IMessageEntity    `tl:",omitempty:flags:7"`
	TTLPeriod   *int32              `tl:",omitempty:flags:25"`
}

func (*UpdateShortMessage) CRC() uint32 {
	return 0x313bc7f8
}
func (*UpdateShortMessage) _IUpdates() {}

type UpdateShortChatMessage struct {
	_           struct{} `tl:"flags,bitflag"`
	Out         bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned   bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread bool     `tl:",omitempty:flags:5,implicit"`
	Silent      bool     `tl:",omitempty:flags:13,implicit"`
	ID          int32
	FromID      int64
	ChatID      int64
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
	FwdFrom     IMessageFwdHeader   `tl:",omitempty:flags:2"`
	ViaBotID    *int64              `tl:",omitempty:flags:11"`
	ReplyTo     IMessageReplyHeader `tl:",omitempty:flags:3"`
	Entities    []IMessageEntity    `tl:",omitempty:flags:7"`
	TTLPeriod   *int32              `tl:",omitempty:flags:25"`
}

func (*UpdateShortChatMessage) CRC() uint32 {
	return 0x4d6deea5
}
func (*UpdateShortChatMessage) _IUpdates() {}

type UpdateShort struct {
	Update IUpdate
	Date   int32
}

func (*UpdateShort) CRC() uint32 {
	return 0x78d4dec1
}
func (*UpdateShort) _IUpdates() {}

type UpdatesCombined struct {
	Updates  []IUpdate
	Users    []IUser
	Chats    []IChat
	Date     int32
	SeqStart int32
	Seq      int32
}

func (*UpdatesCombined) CRC() uint32 {
	return 0x725b04c3
}
func (*UpdatesCombined) _IUpdates() {}

type Updates struct {
	Updates []IUpdate
	Users   []IUser
	Chats   []IChat
	Date    int32
	Seq     int32
}

func (*Updates) CRC() uint32 {
	return 0x74ae4240
}
func (*Updates) _IUpdates() {}

type UpdateShortSentMessage struct {
	_         struct{} `tl:"flags,bitflag"`
	Out       bool     `tl:",omitempty:flags:1,implicit"`
	ID        int32
	Pts       int32
	PtsCount  int32
	Date      int32
	Media     IMessageMedia    `tl:",omitempty:flags:9"`
	Entities  []IMessageEntity `tl:",omitempty:flags:7"`
	TTLPeriod *int32           `tl:",omitempty:flags:25"`
}

func (*UpdateShortSentMessage) CRC() uint32 {
	return 0x9015e101
}
func (*UpdateShortSentMessage) _IUpdates() {}

type IURLAuthResult interface {
	tl.Object
	_IURLAuthResult()
}

var (
	_ IURLAuthResult = (*URLAuthResultRequest)(nil)
	_ IURLAuthResult = (*URLAuthResultAccepted)(nil)
	_ IURLAuthResult = (*URLAuthResultDefault)(nil)
)

type URLAuthResultRequest struct {
	_                  struct{} `tl:"flags,bitflag"`
	RequestWriteAccess bool     `tl:",omitempty:flags:0,implicit"`
	Bot                IUser
	Domain             string
}

func (*URLAuthResultRequest) CRC() uint32 {
	return 0x92d33a0e
}
func (*URLAuthResultRequest) _IURLAuthResult() {}

type URLAuthResultAccepted struct {
	URL string
}

func (*URLAuthResultAccepted) CRC() uint32 {
	return 0x8f8c0e4e
}
func (*URLAuthResultAccepted) _IURLAuthResult() {}

type URLAuthResultDefault struct{}

func (*URLAuthResultDefault) CRC() uint32 {
	return 0xa9d6db1f
}
func (*URLAuthResultDefault) _IURLAuthResult() {}

// Object defines a user.
type IUser interface {
	tl.Object
	_IUser()
}

var (
	_ IUser = (*UserEmpty)(nil)
	_ IUser = (*User)(nil)
)

// Empty constructor, non-existent user.
type UserEmpty struct {
	ID int64 // User identifier or `0`
}

func (*UserEmpty) CRC() uint32 {
	return 0xd3bc4b7a
}
func (*UserEmpty) _IUser() {}

// Indicates info about a certain user
type User struct {
	_                    struct{}             `tl:"flags,bitflag"`
	Self                 bool                 `tl:",omitempty:flags:10,implicit"` // Whether this user indicates the currently logged in user
	Contact              bool                 `tl:",omitempty:flags:11,implicit"` // Whether this user is a contact
	MutualContact        bool                 `tl:",omitempty:flags:12,implicit"` // Whether this user is a mutual contact
	Deleted              bool                 `tl:",omitempty:flags:13,implicit"` // Whether the account of this user was deleted
	Bot                  bool                 `tl:",omitempty:flags:14,implicit"` // Is this user a bot?
	BotChatHistory       bool                 `tl:",omitempty:flags:15,implicit"` // Can the bot see all messages in groups?
	BotNochats           bool                 `tl:",omitempty:flags:16,implicit"` // Can the bot be added to groups?
	Verified             bool                 `tl:",omitempty:flags:17,implicit"` // Whether this user is verified
	Restricted           bool                 `tl:",omitempty:flags:18,implicit"` // Access to this user must be restricted for the reason specified in restriction_reason
	Min                  bool                 `tl:",omitempty:flags:20,implicit"` // see [min](https://core.telegram.org/api/min)
	BotInlineGeo         bool                 `tl:",omitempty:flags:21,implicit"` // Whether the bot can request our geolocation in inline mode
	Support              bool                 `tl:",omitempty:flags:23,implicit"` // Whether this is an official support user
	Scam                 bool                 `tl:",omitempty:flags:24,implicit"` // This may be a scam user
	ApplyMinPhoto        bool                 `tl:",omitempty:flags:25,implicit"` // If set, the profile picture for this user should be refetched
	Fake                 bool                 `tl:",omitempty:flags:26,implicit"`
	BotAttachMenu        bool                 `tl:",omitempty:flags:27,implicit"`
	Premium              bool                 `tl:",omitempty:flags:28,implicit"`
	AttachMenuEnabled    bool                 `tl:",omitempty:flags:29,implicit"`
	_                    struct{}             `tl:"flags2,bitflag"`
	BotCanEdit           bool                 `tl:",omitempty:flags2:1,implicit"`
	CloseFriend          bool                 `tl:",omitempty:flags2:2,implicit"`
	StoriesHidden        bool                 `tl:",omitempty:flags2:3,implicit"`
	StoriesUnavailable   bool                 `tl:",omitempty:flags2:4,implicit"`
	ID                   int64                // ID of the user
	AccessHash           *int64               `tl:",omitempty:flags:0"`  // Access hash of the user
	FirstName            *string              `tl:",omitempty:flags:1"`  // First name
	LastName             *string              `tl:",omitempty:flags:2"`  // Last name
	Username             *string              `tl:",omitempty:flags:3"`  // Username
	Phone                *string              `tl:",omitempty:flags:4"`  // Phone number
	Photo                IUserProfilePhoto    `tl:",omitempty:flags:5"`  // Profile picture of user
	Status               IUserStatus          `tl:",omitempty:flags:6"`  // Online status of user
	BotInfoVersion       *int32               `tl:",omitempty:flags:14"` // Version of the bot_info field in userFull, incremented every time it changes
	RestrictionReason    []IRestrictionReason `tl:",omitempty:flags:18"` // Contains the reason why access to this user must be restricted.
	BotInlinePlaceholder *string              `tl:",omitempty:flags:19"` // Inline placeholder for this inline bot
	LangCode             *string              `tl:",omitempty:flags:22"` // Language code of the user
	EmojiStatus          IEmojiStatus         `tl:",omitempty:flags:30"`
	Usernames            []IUsername          `tl:",omitempty:flags2:0"`
	StoriesMaxID         *int32               `tl:",omitempty:flags2:5"`
	Color                IPeerColor           `tl:",omitempty:flags2:8"`
	ProfileColor         IPeerColor           `tl:",omitempty:flags2:9"`
}

func (*User) CRC() uint32 {
	return 0x215c4438
}
func (*User) _IUser() {}

type IUserFull interface {
	tl.Object
	_IUserFull()
}

var (
	_ IUserFull = (*UserFull)(nil)
)

type UserFull struct {
	_                       struct{} `tl:"flags,bitflag"`
	Blocked                 bool     `tl:",omitempty:flags:0,implicit"`
	PhoneCallsAvailable     bool     `tl:",omitempty:flags:4,implicit"`
	PhoneCallsPrivate       bool     `tl:",omitempty:flags:5,implicit"`
	CanPinMessage           bool     `tl:",omitempty:flags:7,implicit"`
	HasScheduled            bool     `tl:",omitempty:flags:12,implicit"`
	VideoCallsAvailable     bool     `tl:",omitempty:flags:13,implicit"`
	VoiceMessagesForbidden  bool     `tl:",omitempty:flags:20,implicit"`
	TranslationsDisabled    bool     `tl:",omitempty:flags:23,implicit"`
	StoriesPinnedAvailable  bool     `tl:",omitempty:flags:26,implicit"`
	BlockedMyStoriesFrom    bool     `tl:",omitempty:flags:27,implicit"`
	WallpaperOverridden     bool     `tl:",omitempty:flags:28,implicit"`
	ID                      int64
	About                   *string `tl:",omitempty:flags:1"`
	Settings                IPeerSettings
	PersonalPhoto           IPhoto `tl:",omitempty:flags:21"`
	ProfilePhoto            IPhoto `tl:",omitempty:flags:2"`
	FallbackPhoto           IPhoto `tl:",omitempty:flags:22"`
	NotifySettings          IPeerNotifySettings
	BotInfo                 IBotInfo `tl:",omitempty:flags:3"`
	PinnedMsgID             *int32   `tl:",omitempty:flags:6"`
	CommonChatsCount        int32
	FolderID                *int32               `tl:",omitempty:flags:11"`
	TTLPeriod               *int32               `tl:",omitempty:flags:14"`
	ThemeEmoticon           *string              `tl:",omitempty:flags:15"`
	PrivateForwardName      *string              `tl:",omitempty:flags:16"`
	BotGroupAdminRights     IChatAdminRights     `tl:",omitempty:flags:17"`
	BotBroadcastAdminRights IChatAdminRights     `tl:",omitempty:flags:18"`
	PremiumGifts            []IPremiumGiftOption `tl:",omitempty:flags:19"`
	Wallpaper               IWallPaper           `tl:",omitempty:flags:24"`
	Stories                 IPeerStories         `tl:",omitempty:flags:25"`
}

func (*UserFull) CRC() uint32 {
	return 0xb9b12c6c
}
func (*UserFull) _IUserFull() {}

type IUserProfilePhoto interface {
	tl.Object
	_IUserProfilePhoto()
}

var (
	_ IUserProfilePhoto = (*UserProfilePhotoEmpty)(nil)
	_ IUserProfilePhoto = (*UserProfilePhoto)(nil)
)

type UserProfilePhotoEmpty struct{}

func (*UserProfilePhotoEmpty) CRC() uint32 {
	return 0x4f11bae1
}
func (*UserProfilePhotoEmpty) _IUserProfilePhoto() {}

type UserProfilePhoto struct {
	_             struct{} `tl:"flags,bitflag"`
	HasVideo      bool     `tl:",omitempty:flags:0,implicit"`
	Personal      bool     `tl:",omitempty:flags:2,implicit"`
	PhotoID       int64
	StrippedThumb *[]byte `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*UserProfilePhoto) CRC() uint32 {
	return 0x82d1f706
}
func (*UserProfilePhoto) _IUserProfilePhoto() {}

type IUserStatus interface {
	tl.Object
	_IUserStatus()
}

var (
	_ IUserStatus = (*UserStatusEmpty)(nil)
	_ IUserStatus = (*UserStatusOnline)(nil)
	_ IUserStatus = (*UserStatusOffline)(nil)
	_ IUserStatus = (*UserStatusRecently)(nil)
	_ IUserStatus = (*UserStatusLastWeek)(nil)
	_ IUserStatus = (*UserStatusLastMonth)(nil)
)

type UserStatusEmpty struct{}

func (*UserStatusEmpty) CRC() uint32 {
	return 0x9d05049
}
func (*UserStatusEmpty) _IUserStatus() {}

type UserStatusOnline struct {
	Expires int32
}

func (*UserStatusOnline) CRC() uint32 {
	return 0xedb93949
}
func (*UserStatusOnline) _IUserStatus() {}

type UserStatusOffline struct {
	WasOnline int32
}

func (*UserStatusOffline) CRC() uint32 {
	return 0x8c703f
}
func (*UserStatusOffline) _IUserStatus() {}

type UserStatusRecently struct{}

func (*UserStatusRecently) CRC() uint32 {
	return 0xe26f42f1
}
func (*UserStatusRecently) _IUserStatus() {}

type UserStatusLastWeek struct{}

func (*UserStatusLastWeek) CRC() uint32 {
	return 0x7bf09fc
}
func (*UserStatusLastWeek) _IUserStatus() {}

type UserStatusLastMonth struct{}

func (*UserStatusLastMonth) CRC() uint32 {
	return 0x77ebc742
}
func (*UserStatusLastMonth) _IUserStatus() {}

type IUsername interface {
	tl.Object
	_IUsername()
}

var (
	_ IUsername = (*Username)(nil)
)

type Username struct {
	_        struct{} `tl:"flags,bitflag"`
	Editable bool     `tl:",omitempty:flags:0,implicit"`
	Active   bool     `tl:",omitempty:flags:1,implicit"`
	Username string
}

func (*Username) CRC() uint32 {
	return 0xb4073647
}
func (*Username) _IUsername() {}

type IVideoSize interface {
	tl.Object
	_IVideoSize()
}

var (
	_ IVideoSize = (*VideoSize)(nil)
	_ IVideoSize = (*VideoSizeEmojiMarkup)(nil)
	_ IVideoSize = (*VideoSizeStickerMarkup)(nil)
)

type VideoSize struct {
	_            struct{} `tl:"flags,bitflag"`
	Type         string
	W            int32
	H            int32
	Size         int32
	VideoStartTs *float64 `tl:",omitempty:flags:0"`
}

func (*VideoSize) CRC() uint32 {
	return 0xde33b094
}
func (*VideoSize) _IVideoSize() {}

type VideoSizeEmojiMarkup struct {
	EmojiID          int64
	BackgroundColors []int32
}

func (*VideoSizeEmojiMarkup) CRC() uint32 {
	return 0xf85c413c
}
func (*VideoSizeEmojiMarkup) _IVideoSize() {}

type VideoSizeStickerMarkup struct {
	Stickerset       IInputStickerSet
	StickerID        int64
	BackgroundColors []int32
}

func (*VideoSizeStickerMarkup) CRC() uint32 {
	return 0xda082fe
}
func (*VideoSizeStickerMarkup) _IVideoSize() {}

type IWallPaper interface {
	tl.Object
	_IWallPaper()
}

var (
	_ IWallPaper = (*WallPaper)(nil)
	_ IWallPaper = (*WallPaperNoFile)(nil)
)

type WallPaper struct {
	ID         int64
	_          struct{} `tl:"flags,bitflag"`
	Creator    bool     `tl:",omitempty:flags:0,implicit"`
	Default    bool     `tl:",omitempty:flags:1,implicit"`
	Pattern    bool     `tl:",omitempty:flags:3,implicit"`
	Dark       bool     `tl:",omitempty:flags:4,implicit"`
	AccessHash int64
	Slug       string
	Document   IDocument
	Settings   IWallPaperSettings `tl:",omitempty:flags:2"`
}

func (*WallPaper) CRC() uint32 {
	return 0xa437c3ed
}
func (*WallPaper) _IWallPaper() {}

type WallPaperNoFile struct {
	ID       int64
	_        struct{}           `tl:"flags,bitflag"`
	Default  bool               `tl:",omitempty:flags:1,implicit"`
	Dark     bool               `tl:",omitempty:flags:4,implicit"`
	Settings IWallPaperSettings `tl:",omitempty:flags:2"`
}

func (*WallPaperNoFile) CRC() uint32 {
	return 0xe0804116
}
func (*WallPaperNoFile) _IWallPaper() {}

type IWallPaperSettings interface {
	tl.Object
	_IWallPaperSettings()
}

var (
	_ IWallPaperSettings = (*WallPaperSettings)(nil)
)

type WallPaperSettings struct {
	_                     struct{} `tl:"flags,bitflag"`
	Blur                  bool     `tl:",omitempty:flags:1,implicit"`
	Motion                bool     `tl:",omitempty:flags:2,implicit"`
	BackgroundColor       *int32   `tl:",omitempty:flags:0"`
	SecondBackgroundColor *int32   `tl:",omitempty:flags:4"`
	ThirdBackgroundColor  *int32   `tl:",omitempty:flags:5"`
	FourthBackgroundColor *int32   `tl:",omitempty:flags:6"`
	Intensity             *int32   `tl:",omitempty:flags:3"`
	Rotation              *int32   `tl:",omitempty:flags:4"`
	Emoticon              *string  `tl:",omitempty:flags:7"`
}

func (*WallPaperSettings) CRC() uint32 {
	return 0x372efcd0
}
func (*WallPaperSettings) _IWallPaperSettings() {}

type IWebAuthorization interface {
	tl.Object
	_IWebAuthorization()
}

var (
	_ IWebAuthorization = (*WebAuthorization)(nil)
)

type WebAuthorization struct {
	Hash        int64
	BotID       int64
	Domain      string
	Browser     string
	Platform    string
	DateCreated int32
	DateActive  int32
	Ip          string
	Region      string
}

func (*WebAuthorization) CRC() uint32 {
	return 0xa6f8f452
}
func (*WebAuthorization) _IWebAuthorization() {}

type IWebDocument interface {
	tl.Object
	_IWebDocument()
}

var (
	_ IWebDocument = (*WebDocument)(nil)
	_ IWebDocument = (*WebDocumentNoProxy)(nil)
)

type WebDocument struct {
	URL        string
	AccessHash int64
	Size       int32
	MimeType   string
	Attributes []IDocumentAttribute
}

func (*WebDocument) CRC() uint32 {
	return 0x1c570ed1
}
func (*WebDocument) _IWebDocument() {}

type WebDocumentNoProxy struct {
	URL        string
	Size       int32
	MimeType   string
	Attributes []IDocumentAttribute
}

func (*WebDocumentNoProxy) CRC() uint32 {
	return 0xf9c8bcc6
}
func (*WebDocumentNoProxy) _IWebDocument() {}

// Instant View webpage preview
type IWebPage interface {
	tl.Object
	_IWebPage()
}

var (
	_ IWebPage = (*WebPageEmpty)(nil)
	_ IWebPage = (*WebPagePending)(nil)
	_ IWebPage = (*WebPage)(nil)
	_ IWebPage = (*WebPageNotModified)(nil)
)

// No preview is available for the webpage
type WebPageEmpty struct {
	_   struct{} `tl:"flags,bitflag"`
	ID  int64    // Preview ID
	URL *string  `tl:",omitempty:flags:0"`
}

func (*WebPageEmpty) CRC() uint32 {
	return 0x211a1788
}
func (*WebPageEmpty) _IWebPage() {}

// A preview of the webpage is currently being generated
type WebPagePending struct {
	_    struct{} `tl:"flags,bitflag"`
	ID   int64    // ID of preview
	URL  *string  `tl:",omitempty:flags:0"`
	Date int32    // When was the processing started
}

func (*WebPagePending) CRC() uint32 {
	return 0xb0d13e47
}
func (*WebPagePending) _IWebPage() {}

// Webpage preview
type WebPage struct {
	_             struct{}            `tl:"flags,bitflag"` // Flags, see TL conditional fields
	HasLargeMedia bool                `tl:",omitempty:flags:13,implicit"`
	ID            int64               // Preview ID
	URL           string              // URL of previewed webpage
	DisplayURL    string              // Webpage URL to be displayed to the user
	Hash          int32               // Hash for pagination, for more info click here
	Type          *string             `tl:",omitempty:flags:0"`  // Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
	SiteName      *string             `tl:",omitempty:flags:1"`  // Short name of the site (e.g., Google Docs, App Store)
	Title         *string             `tl:",omitempty:flags:2"`  // Title of the content
	Description   *string             `tl:",omitempty:flags:3"`  // Content description
	Photo         IPhoto              `tl:",omitempty:flags:4"`  // Image representing the content
	EmbedURL      *string             `tl:",omitempty:flags:5"`  // URL to show in the embedded preview
	EmbedType     *string             `tl:",omitempty:flags:5"`  // MIME type of the embedded preview, (e.g., text/html or video/mp4)
	EmbedWidth    *int32              `tl:",omitempty:flags:6"`  // Width of the embedded preview
	EmbedHeight   *int32              `tl:",omitempty:flags:6"`  // Height of the embedded preview
	Duration      *int32              `tl:",omitempty:flags:7"`  // Duration of the content, in seconds
	Author        *string             `tl:",omitempty:flags:8"`  // Author of the content
	Document      IDocument           `tl:",omitempty:flags:9"`  // Preview of the content as a media file
	CachedPage    IPage               `tl:",omitempty:flags:10"` // Page contents in instant view format
	Attributes    []IWebPageAttribute `tl:",omitempty:flags:12"` // Webpage attributes
}

func (*WebPage) CRC() uint32 {
	return 0xe89c45b2
}
func (*WebPage) _IWebPage() {}

// The preview of the webpage hasn't changed
type WebPageNotModified struct {
	_               struct{} `tl:"flags,bitflag"`      // Flags, see TL conditional fields
	CachedPageViews *int32   `tl:",omitempty:flags:0"` // Page view count
}

func (*WebPageNotModified) CRC() uint32 {
	return 0x7311ca11
}
func (*WebPageNotModified) _IWebPage() {}

type IWebPageAttribute interface {
	tl.Object
	_IWebPageAttribute()
}

var (
	_ IWebPageAttribute = (*WebPageAttributeTheme)(nil)
	_ IWebPageAttribute = (*WebPageAttributeStory)(nil)
)

type WebPageAttributeTheme struct {
	_         struct{}       `tl:"flags,bitflag"`
	Documents []IDocument    `tl:",omitempty:flags:0"`
	Settings  IThemeSettings `tl:",omitempty:flags:1"`
}

func (*WebPageAttributeTheme) CRC() uint32 {
	return 0x54b56617
}
func (*WebPageAttributeTheme) _IWebPageAttribute() {}

type WebPageAttributeStory struct {
	_     struct{} `tl:"flags,bitflag"`
	Peer  IPeer
	ID    int32
	Story IStoryItem `tl:",omitempty:flags:0"`
}

func (*WebPageAttributeStory) CRC() uint32 {
	return 0x2e94c3e7
}
func (*WebPageAttributeStory) _IWebPageAttribute() {}

type IWebViewMessageSent interface {
	tl.Object
	_IWebViewMessageSent()
}

var (
	_ IWebViewMessageSent = (*WebViewMessageSent)(nil)
)

type WebViewMessageSent struct {
	_     struct{}                 `tl:"flags,bitflag"`
	MsgID IInputBotInlineMessageID `tl:",omitempty:flags:0"`
}

func (*WebViewMessageSent) CRC() uint32 {
	return 0xc94511c
}
func (*WebViewMessageSent) _IWebViewMessageSent() {}

type IWebViewResult interface {
	tl.Object
	_IWebViewResult()
}

var (
	_ IWebViewResult = (*WebViewResultURL)(nil)
)

type WebViewResultURL struct {
	QueryID int64
	URL     string
}

func (*WebViewResultURL) CRC() uint32 {
	return 0xc14557c
}
func (*WebViewResultURL) _IWebViewResult() {}

// Authorization form
type IAccountAuthorizationForm interface {
	tl.Object
	_IAccountAuthorizationForm()
}

var (
	_ IAccountAuthorizationForm = (*AccountAuthorizationForm)(nil)
)

// Telegram Passport authorization form
type AccountAuthorizationForm struct {
	_                struct{}              `tl:"flags,bitflag"`
	RequiredTypes    []ISecureRequiredType // Required Telegram Passport¹ documents
	Values           []ISecureValue        // Already submitted Telegram Passport¹ documents
	Errors           []ISecureValueError
	Users            []IUser // Info about the bot to which the form will be submitted
	PrivacyPolicyURL *string `tl:",omitempty:flags:0"` // URL of the service's privacy policy
}

func (*AccountAuthorizationForm) CRC() uint32 {
	return 0xad2e1cd8
}
func (*AccountAuthorizationForm) _IAccountAuthorizationForm() {}

type IAccountAuthorizations interface {
	tl.Object
	_IAccountAuthorizations()
}

var (
	_ IAccountAuthorizations = (*AccountAuthorizations)(nil)
)

type AccountAuthorizations struct {
	AuthorizationTTLDays int32
	Authorizations       []IAuthorization
}

func (*AccountAuthorizations) CRC() uint32 {
	return 0x4bff8ea0
}
func (*AccountAuthorizations) _IAccountAuthorizations() {}

type IAccountAutoDownloadSettings interface {
	tl.Object
	_IAccountAutoDownloadSettings()
}

var (
	_ IAccountAutoDownloadSettings = (*AccountAutoDownloadSettings)(nil)
)

type AccountAutoDownloadSettings struct {
	Low    IAutoDownloadSettings
	Medium IAutoDownloadSettings
	High   IAutoDownloadSettings
}

func (*AccountAutoDownloadSettings) CRC() uint32 {
	return 0x63cacf26
}
func (*AccountAutoDownloadSettings) _IAccountAutoDownloadSettings() {}

type IAccountAutoSaveSettings interface {
	tl.Object
	_IAccountAutoSaveSettings()
}

var (
	_ IAccountAutoSaveSettings = (*AccountAutoSaveSettings)(nil)
)

type AccountAutoSaveSettings struct {
	UsersSettings      IAutoSaveSettings
	ChatsSettings      IAutoSaveSettings
	BroadcastsSettings IAutoSaveSettings
	Exceptions         []IAutoSaveException
	Chats              []IChat
	Users              []IUser
}

func (*AccountAutoSaveSettings) CRC() uint32 {
	return 0x4c3e069d
}
func (*AccountAutoSaveSettings) _IAccountAutoSaveSettings() {}

type IAccountContentSettings interface {
	tl.Object
	_IAccountContentSettings()
}

var (
	_ IAccountContentSettings = (*AccountContentSettings)(nil)
)

type AccountContentSettings struct {
	_                  struct{} `tl:"flags,bitflag"`
	SensitiveEnabled   bool     `tl:",omitempty:flags:0,implicit"`
	SensitiveCanChange bool     `tl:",omitempty:flags:1,implicit"`
}

func (*AccountContentSettings) CRC() uint32 {
	return 0x57e28221
}
func (*AccountContentSettings) _IAccountContentSettings() {}

type IAccountEmailVerified interface {
	tl.Object
	_IAccountEmailVerified()
}

var (
	_ IAccountEmailVerified = (*AccountEmailVerified)(nil)
	_ IAccountEmailVerified = (*AccountEmailVerifiedLogin)(nil)
)

type AccountEmailVerified struct {
	Email string
}

func (*AccountEmailVerified) CRC() uint32 {
	return 0x2b96cd1b
}
func (*AccountEmailVerified) _IAccountEmailVerified() {}

type AccountEmailVerifiedLogin struct {
	Email    string
	SentCode IAuthSentCode
}

func (*AccountEmailVerifiedLogin) CRC() uint32 {
	return 0xe1bb0d61
}
func (*AccountEmailVerifiedLogin) _IAccountEmailVerified() {}

type IAccountEmojiStatuses interface {
	tl.Object
	_IAccountEmojiStatuses()
}

var (
	_ IAccountEmojiStatuses = (*AccountEmojiStatusesNotModified)(nil)
	_ IAccountEmojiStatuses = (*AccountEmojiStatuses)(nil)
)

type AccountEmojiStatusesNotModified struct{}

func (*AccountEmojiStatusesNotModified) CRC() uint32 {
	return 0xd08ce645
}
func (*AccountEmojiStatusesNotModified) _IAccountEmojiStatuses() {}

type AccountEmojiStatuses struct {
	Hash     int64
	Statuses []IEmojiStatus
}

func (*AccountEmojiStatuses) CRC() uint32 {
	return 0x90c467d1
}
func (*AccountEmojiStatuses) _IAccountEmojiStatuses() {}

type IAccountPassword interface {
	tl.Object
	_IAccountPassword()
}

var (
	_ IAccountPassword = (*AccountPassword)(nil)
)

type AccountPassword struct {
	_                       struct{}         `tl:"flags,bitflag"`
	HasRecovery             bool             `tl:",omitempty:flags:0,implicit"`
	HasSecureValues         bool             `tl:",omitempty:flags:1,implicit"`
	HasPassword             bool             `tl:",omitempty:flags:2,implicit"`
	CurrentAlgo             IPasswordKdfAlgo `tl:",omitempty:flags:2"`
	SRPB                    *[]byte          `tl:",omitempty:flags:2"`
	SRPID                   *int64           `tl:",omitempty:flags:2"`
	Hint                    *string          `tl:",omitempty:flags:3"`
	EmailUnconfirmedPattern *string          `tl:",omitempty:flags:4"`
	NewAlgo                 IPasswordKdfAlgo
	NewSecureAlgo           ISecurePasswordKdfAlgo
	SecureRandom            []byte
	PendingResetDate        *int32  `tl:",omitempty:flags:5"`
	LoginEmailPattern       *string `tl:",omitempty:flags:6"`
}

func (*AccountPassword) CRC() uint32 {
	return 0x957b50fb
}
func (*AccountPassword) _IAccountPassword() {}

type IAccountPasswordInputSettings interface {
	tl.Object
	_IAccountPasswordInputSettings()
}

var (
	_ IAccountPasswordInputSettings = (*AccountPasswordInputSettings)(nil)
)

type AccountPasswordInputSettings struct {
	_                 struct{}              `tl:"flags,bitflag"`
	NewAlgo           IPasswordKdfAlgo      `tl:",omitempty:flags:0"`
	NewPasswordHash   *[]byte               `tl:",omitempty:flags:0"`
	Hint              *string               `tl:",omitempty:flags:0"`
	Email             *string               `tl:",omitempty:flags:1"`
	NewSecureSettings ISecureSecretSettings `tl:",omitempty:flags:2"`
}

func (*AccountPasswordInputSettings) CRC() uint32 {
	return 0xc23727c9
}
func (*AccountPasswordInputSettings) _IAccountPasswordInputSettings() {}

type IAccountPasswordSettings interface {
	tl.Object
	_IAccountPasswordSettings()
}

var (
	_ IAccountPasswordSettings = (*AccountPasswordSettings)(nil)
)

type AccountPasswordSettings struct {
	_              struct{}              `tl:"flags,bitflag"`
	Email          *string               `tl:",omitempty:flags:0"`
	SecureSettings ISecureSecretSettings `tl:",omitempty:flags:1"`
}

func (*AccountPasswordSettings) CRC() uint32 {
	return 0x9a5c33e5
}
func (*AccountPasswordSettings) _IAccountPasswordSettings() {}

type IAccountPrivacyRules interface {
	tl.Object
	_IAccountPrivacyRules()
}

var (
	_ IAccountPrivacyRules = (*AccountPrivacyRules)(nil)
)

type AccountPrivacyRules struct {
	Rules []IPrivacyRule
	Chats []IChat
	Users []IUser
}

func (*AccountPrivacyRules) CRC() uint32 {
	return 0x50a04e45
}
func (*AccountPrivacyRules) _IAccountPrivacyRules() {}

type IAccountResetPasswordResult interface {
	tl.Object
	_IAccountResetPasswordResult()
}

var (
	_ IAccountResetPasswordResult = (*AccountResetPasswordFailedWait)(nil)
	_ IAccountResetPasswordResult = (*AccountResetPasswordRequestedWait)(nil)
	_ IAccountResetPasswordResult = (*AccountResetPasswordOk)(nil)
)

type AccountResetPasswordFailedWait struct {
	RetryDate int32
}

func (*AccountResetPasswordFailedWait) CRC() uint32 {
	return 0xe3779861
}
func (*AccountResetPasswordFailedWait) _IAccountResetPasswordResult() {}

type AccountResetPasswordRequestedWait struct {
	UntilDate int32
}

func (*AccountResetPasswordRequestedWait) CRC() uint32 {
	return 0xe9effc7d
}
func (*AccountResetPasswordRequestedWait) _IAccountResetPasswordResult() {}

type AccountResetPasswordOk struct{}

func (*AccountResetPasswordOk) CRC() uint32 {
	return 0xe926d63e
}
func (*AccountResetPasswordOk) _IAccountResetPasswordResult() {}

type IAccountSavedRingtone interface {
	tl.Object
	_IAccountSavedRingtone()
}

var (
	_ IAccountSavedRingtone = (*AccountSavedRingtone)(nil)
	_ IAccountSavedRingtone = (*AccountSavedRingtoneConverted)(nil)
)

type AccountSavedRingtone struct{}

func (*AccountSavedRingtone) CRC() uint32 {
	return 0xb7263f6d
}
func (*AccountSavedRingtone) _IAccountSavedRingtone() {}

type AccountSavedRingtoneConverted struct {
	Document IDocument
}

func (*AccountSavedRingtoneConverted) CRC() uint32 {
	return 0x1f307eb7
}
func (*AccountSavedRingtoneConverted) _IAccountSavedRingtone() {}

type IAccountSavedRingtones interface {
	tl.Object
	_IAccountSavedRingtones()
}

var (
	_ IAccountSavedRingtones = (*AccountSavedRingtonesNotModified)(nil)
	_ IAccountSavedRingtones = (*AccountSavedRingtones)(nil)
)

type AccountSavedRingtonesNotModified struct{}

func (*AccountSavedRingtonesNotModified) CRC() uint32 {
	return 0xfbf6e8b1
}
func (*AccountSavedRingtonesNotModified) _IAccountSavedRingtones() {}

type AccountSavedRingtones struct {
	Hash      int64
	Ringtones []IDocument
}

func (*AccountSavedRingtones) CRC() uint32 {
	return 0xc1e92cc5
}
func (*AccountSavedRingtones) _IAccountSavedRingtones() {}

type IAccountSentEmailCode interface {
	tl.Object
	_IAccountSentEmailCode()
}

var (
	_ IAccountSentEmailCode = (*AccountSentEmailCode)(nil)
)

type AccountSentEmailCode struct {
	EmailPattern string
	Length       int32
}

func (*AccountSentEmailCode) CRC() uint32 {
	return 0x811f854f
}
func (*AccountSentEmailCode) _IAccountSentEmailCode() {}

type IAccountTakeout interface {
	tl.Object
	_IAccountTakeout()
}

var (
	_ IAccountTakeout = (*AccountTakeout)(nil)
)

type AccountTakeout struct {
	ID int64
}

func (*AccountTakeout) CRC() uint32 {
	return 0x4dba4501
}
func (*AccountTakeout) _IAccountTakeout() {}

type IAccountThemes interface {
	tl.Object
	_IAccountThemes()
}

var (
	_ IAccountThemes = (*AccountThemesNotModified)(nil)
	_ IAccountThemes = (*AccountThemes)(nil)
)

type AccountThemesNotModified struct{}

func (*AccountThemesNotModified) CRC() uint32 {
	return 0xf41eb622
}
func (*AccountThemesNotModified) _IAccountThemes() {}

type AccountThemes struct {
	Hash   int64
	Themes []ITheme
}

func (*AccountThemes) CRC() uint32 {
	return 0x9a3d8c6d
}
func (*AccountThemes) _IAccountThemes() {}

type IAccountTmpPassword interface {
	tl.Object
	_IAccountTmpPassword()
}

var (
	_ IAccountTmpPassword = (*AccountTmpPassword)(nil)
)

type AccountTmpPassword struct {
	TmpPassword []byte
	ValidUntil  int32
}

func (*AccountTmpPassword) CRC() uint32 {
	return 0xdb64fd34
}
func (*AccountTmpPassword) _IAccountTmpPassword() {}

type IAccountWallPapers interface {
	tl.Object
	_IAccountWallPapers()
}

var (
	_ IAccountWallPapers = (*AccountWallPapersNotModified)(nil)
	_ IAccountWallPapers = (*AccountWallPapers)(nil)
)

type AccountWallPapersNotModified struct{}

func (*AccountWallPapersNotModified) CRC() uint32 {
	return 0x1c199183
}
func (*AccountWallPapersNotModified) _IAccountWallPapers() {}

type AccountWallPapers struct {
	Hash       int64
	Wallpapers []IWallPaper
}

func (*AccountWallPapers) CRC() uint32 {
	return 0xcdc3858c
}
func (*AccountWallPapers) _IAccountWallPapers() {}

type IAccountWebAuthorizations interface {
	tl.Object
	_IAccountWebAuthorizations()
}

var (
	_ IAccountWebAuthorizations = (*AccountWebAuthorizations)(nil)
)

type AccountWebAuthorizations struct {
	Authorizations []IWebAuthorization
	Users          []IUser
}

func (*AccountWebAuthorizations) CRC() uint32 {
	return 0xed56c9fc
}
func (*AccountWebAuthorizations) _IAccountWebAuthorizations() {}

type IAuthAuthorization interface {
	tl.Object
	_IAuthAuthorization()
}

var (
	_ IAuthAuthorization = (*AuthAuthorization)(nil)
	_ IAuthAuthorization = (*AuthAuthorizationSignUpRequired)(nil)
)

type AuthAuthorization struct {
	_                     struct{} `tl:"flags,bitflag"`
	SetupPasswordRequired bool     `tl:",omitempty:flags:1,implicit"`
	OtherwiseReloginDays  *int32   `tl:",omitempty:flags:1"`
	TmpSessions           *int32   `tl:",omitempty:flags:0"`
	FutureAuthToken       *[]byte  `tl:",omitempty:flags:2"`
	User                  IUser
}

func (*AuthAuthorization) CRC() uint32 {
	return 0x2ea2c0d4
}
func (*AuthAuthorization) _IAuthAuthorization() {}

type AuthAuthorizationSignUpRequired struct {
	_              struct{}            `tl:"flags,bitflag"`
	TermsOfService IHelpTermsOfService `tl:",omitempty:flags:0"`
}

func (*AuthAuthorizationSignUpRequired) CRC() uint32 {
	return 0x44747e9a
}
func (*AuthAuthorizationSignUpRequired) _IAuthAuthorization() {}

type IAuthExportedAuthorization interface {
	tl.Object
	_IAuthExportedAuthorization()
}

var (
	_ IAuthExportedAuthorization = (*AuthExportedAuthorization)(nil)
)

type AuthExportedAuthorization struct {
	ID    int64
	Bytes []byte
}

func (*AuthExportedAuthorization) CRC() uint32 {
	return 0xb434e2b8
}
func (*AuthExportedAuthorization) _IAuthExportedAuthorization() {}

type IAuthLoggedOut interface {
	tl.Object
	_IAuthLoggedOut()
}

var (
	_ IAuthLoggedOut = (*AuthLoggedOut)(nil)
)

type AuthLoggedOut struct {
	_               struct{} `tl:"flags,bitflag"`
	FutureAuthToken *[]byte  `tl:",omitempty:flags:0"`
}

func (*AuthLoggedOut) CRC() uint32 {
	return 0xc3a2835f
}
func (*AuthLoggedOut) _IAuthLoggedOut() {}

type IAuthLoginToken interface {
	tl.Object
	_IAuthLoginToken()
}

var (
	_ IAuthLoginToken = (*AuthLoginToken)(nil)
	_ IAuthLoginToken = (*AuthLoginTokenMigrateTo)(nil)
	_ IAuthLoginToken = (*AuthLoginTokenSuccess)(nil)
)

type AuthLoginToken struct {
	Expires int32
	Token   []byte
}

func (*AuthLoginToken) CRC() uint32 {
	return 0x629f1980
}
func (*AuthLoginToken) _IAuthLoginToken() {}

type AuthLoginTokenMigrateTo struct {
	DcID  int32
	Token []byte
}

func (*AuthLoginTokenMigrateTo) CRC() uint32 {
	return 0x68e9916
}
func (*AuthLoginTokenMigrateTo) _IAuthLoginToken() {}

type AuthLoginTokenSuccess struct {
	Authorization IAuthAuthorization
}

func (*AuthLoginTokenSuccess) CRC() uint32 {
	return 0x390d5c5e
}
func (*AuthLoginTokenSuccess) _IAuthLoginToken() {}

type IAuthPasswordRecovery interface {
	tl.Object
	_IAuthPasswordRecovery()
}

var (
	_ IAuthPasswordRecovery = (*AuthPasswordRecovery)(nil)
)

type AuthPasswordRecovery struct {
	EmailPattern string
}

func (*AuthPasswordRecovery) CRC() uint32 {
	return 0x137948a5
}
func (*AuthPasswordRecovery) _IAuthPasswordRecovery() {}

type IAuthSentCode interface {
	tl.Object
	_IAuthSentCode()
}

var (
	_ IAuthSentCode = (*AuthSentCode)(nil)
	_ IAuthSentCode = (*AuthSentCodeSuccess)(nil)
)

type AuthSentCode struct {
	_             struct{} `tl:"flags,bitflag"`
	Type          IAuthSentCodeType
	PhoneCodeHash string
	NextType      EnumAuthCodeType `tl:",omitempty:flags:1"`
	Timeout       *int32           `tl:",omitempty:flags:2"`
}

func (*AuthSentCode) CRC() uint32 {
	return 0x5e002502
}
func (*AuthSentCode) _IAuthSentCode() {}

type AuthSentCodeSuccess struct {
	Authorization IAuthAuthorization
}

func (*AuthSentCodeSuccess) CRC() uint32 {
	return 0x2390fe44
}
func (*AuthSentCodeSuccess) _IAuthSentCode() {}

type IAuthSentCodeType interface {
	tl.Object
	_IAuthSentCodeType()
}

var (
	_ IAuthSentCodeType = (*AuthSentCodeTypeApp)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeSms)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeCall)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeFlashCall)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeMissedCall)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeEmailCode)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeSetUpEmailRequired)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeFragmentSms)(nil)
	_ IAuthSentCodeType = (*AuthSentCodeTypeFirebaseSms)(nil)
)

type AuthSentCodeTypeApp struct {
	Length int32
}

func (*AuthSentCodeTypeApp) CRC() uint32 {
	return 0x3dbb5986
}
func (*AuthSentCodeTypeApp) _IAuthSentCodeType() {}

type AuthSentCodeTypeSms struct {
	Length int32
}

func (*AuthSentCodeTypeSms) CRC() uint32 {
	return 0xc000bba2
}
func (*AuthSentCodeTypeSms) _IAuthSentCodeType() {}

type AuthSentCodeTypeCall struct {
	Length int32
}

func (*AuthSentCodeTypeCall) CRC() uint32 {
	return 0x5353e5a7
}
func (*AuthSentCodeTypeCall) _IAuthSentCodeType() {}

type AuthSentCodeTypeFlashCall struct {
	Pattern string
}

func (*AuthSentCodeTypeFlashCall) CRC() uint32 {
	return 0xab03c6d9
}
func (*AuthSentCodeTypeFlashCall) _IAuthSentCodeType() {}

type AuthSentCodeTypeMissedCall struct {
	Prefix string
	Length int32
}

func (*AuthSentCodeTypeMissedCall) CRC() uint32 {
	return 0x82006484
}
func (*AuthSentCodeTypeMissedCall) _IAuthSentCodeType() {}

type AuthSentCodeTypeEmailCode struct {
	_                    struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed   bool     `tl:",omitempty:flags:0,implicit"`
	GoogleSigninAllowed  bool     `tl:",omitempty:flags:1,implicit"`
	EmailPattern         string
	Length               int32
	ResetAvailablePeriod *int32 `tl:",omitempty:flags:3"`
	ResetPendingDate     *int32 `tl:",omitempty:flags:4"`
}

func (*AuthSentCodeTypeEmailCode) CRC() uint32 {
	return 0xf450f59b
}
func (*AuthSentCodeTypeEmailCode) _IAuthSentCodeType() {}

type AuthSentCodeTypeSetUpEmailRequired struct {
	_                   struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed  bool     `tl:",omitempty:flags:0,implicit"`
	GoogleSigninAllowed bool     `tl:",omitempty:flags:1,implicit"`
}

func (*AuthSentCodeTypeSetUpEmailRequired) CRC() uint32 {
	return 0xa5491dea
}
func (*AuthSentCodeTypeSetUpEmailRequired) _IAuthSentCodeType() {}

type AuthSentCodeTypeFragmentSms struct {
	URL    string
	Length int32
}

func (*AuthSentCodeTypeFragmentSms) CRC() uint32 {
	return 0xd9565c39
}
func (*AuthSentCodeTypeFragmentSms) _IAuthSentCodeType() {}

type AuthSentCodeTypeFirebaseSms struct {
	_           struct{} `tl:"flags,bitflag"`
	Nonce       *[]byte  `tl:",omitempty:flags:0"`
	Receipt     *string  `tl:",omitempty:flags:1"`
	PushTimeout *int32   `tl:",omitempty:flags:1"`
	Length      int32
}

func (*AuthSentCodeTypeFirebaseSms) CRC() uint32 {
	return 0xe57b1432
}
func (*AuthSentCodeTypeFirebaseSms) _IAuthSentCodeType() {}

type IBotsBotInfo interface {
	tl.Object
	_IBotsBotInfo()
}

var (
	_ IBotsBotInfo = (*BotsBotInfo)(nil)
)

type BotsBotInfo struct {
	Name        string
	About       string
	Description string
}

func (*BotsBotInfo) CRC() uint32 {
	return 0xe8a775b0
}
func (*BotsBotInfo) _IBotsBotInfo() {}

type IChannelsAdminLogResults interface {
	tl.Object
	_IChannelsAdminLogResults()
}

var (
	_ IChannelsAdminLogResults = (*ChannelsAdminLogResults)(nil)
)

type ChannelsAdminLogResults struct {
	Events []IChannelAdminLogEvent
	Chats  []IChat
	Users  []IUser
}

func (*ChannelsAdminLogResults) CRC() uint32 {
	return 0xed8af74d
}
func (*ChannelsAdminLogResults) _IChannelsAdminLogResults() {}

type IChannelsChannelParticipant interface {
	tl.Object
	_IChannelsChannelParticipant()
}

var (
	_ IChannelsChannelParticipant = (*ChannelsChannelParticipant)(nil)
)

type ChannelsChannelParticipant struct {
	Participant IChannelParticipant
	Chats       []IChat
	Users       []IUser
}

func (*ChannelsChannelParticipant) CRC() uint32 {
	return 0xdfb80317
}
func (*ChannelsChannelParticipant) _IChannelsChannelParticipant() {}

type IChannelsChannelParticipants interface {
	tl.Object
	_IChannelsChannelParticipants()
}

var (
	_ IChannelsChannelParticipants = (*ChannelsChannelParticipants)(nil)
	_ IChannelsChannelParticipants = (*ChannelsChannelParticipantsNotModified)(nil)
)

type ChannelsChannelParticipants struct {
	Count        int32
	Participants []IChannelParticipant
	Chats        []IChat
	Users        []IUser
}

func (*ChannelsChannelParticipants) CRC() uint32 {
	return 0x9ab0feaf
}
func (*ChannelsChannelParticipants) _IChannelsChannelParticipants() {}

type ChannelsChannelParticipantsNotModified struct{}

func (*ChannelsChannelParticipantsNotModified) CRC() uint32 {
	return 0xf0173fe9
}
func (*ChannelsChannelParticipantsNotModified) _IChannelsChannelParticipants() {}

type IChannelsSendAsPeers interface {
	tl.Object
	_IChannelsSendAsPeers()
}

var (
	_ IChannelsSendAsPeers = (*ChannelsSendAsPeers)(nil)
)

type ChannelsSendAsPeers struct {
	Peers []ISendAsPeer
	Chats []IChat
	Users []IUser
}

func (*ChannelsSendAsPeers) CRC() uint32 {
	return 0xf496b0c6
}
func (*ChannelsSendAsPeers) _IChannelsSendAsPeers() {}

type IChatlistsChatlistInvite interface {
	tl.Object
	_IChatlistsChatlistInvite()
}

var (
	_ IChatlistsChatlistInvite = (*ChatlistsChatlistInviteAlready)(nil)
	_ IChatlistsChatlistInvite = (*ChatlistsChatlistInvite)(nil)
)

type ChatlistsChatlistInviteAlready struct {
	FilterID     int32
	MissingPeers []IPeer
	AlreadyPeers []IPeer
	Chats        []IChat
	Users        []IUser
}

func (*ChatlistsChatlistInviteAlready) CRC() uint32 {
	return 0xfa87f659
}
func (*ChatlistsChatlistInviteAlready) _IChatlistsChatlistInvite() {}

type ChatlistsChatlistInvite struct {
	_        struct{} `tl:"flags,bitflag"`
	Title    string
	Emoticon *string `tl:",omitempty:flags:0"`
	Peers    []IPeer
	Chats    []IChat
	Users    []IUser
}

func (*ChatlistsChatlistInvite) CRC() uint32 {
	return 0x1dcd839d
}
func (*ChatlistsChatlistInvite) _IChatlistsChatlistInvite() {}

type IChatlistsChatlistUpdates interface {
	tl.Object
	_IChatlistsChatlistUpdates()
}

var (
	_ IChatlistsChatlistUpdates = (*ChatlistsChatlistUpdates)(nil)
)

type ChatlistsChatlistUpdates struct {
	MissingPeers []IPeer
	Chats        []IChat
	Users        []IUser
}

func (*ChatlistsChatlistUpdates) CRC() uint32 {
	return 0x93bd878d
}
func (*ChatlistsChatlistUpdates) _IChatlistsChatlistUpdates() {}

type IChatlistsExportedChatlistInvite interface {
	tl.Object
	_IChatlistsExportedChatlistInvite()
}

var (
	_ IChatlistsExportedChatlistInvite = (*ChatlistsExportedChatlistInvite)(nil)
)

type ChatlistsExportedChatlistInvite struct {
	Filter IDialogFilter
	Invite IExportedChatlistInvite
}

func (*ChatlistsExportedChatlistInvite) CRC() uint32 {
	return 0x10e6e3a6
}
func (*ChatlistsExportedChatlistInvite) _IChatlistsExportedChatlistInvite() {}

type IChatlistsExportedInvites interface {
	tl.Object
	_IChatlistsExportedInvites()
}

var (
	_ IChatlistsExportedInvites = (*ChatlistsExportedInvites)(nil)
)

type ChatlistsExportedInvites struct {
	Invites []IExportedChatlistInvite
	Chats   []IChat
	Users   []IUser
}

func (*ChatlistsExportedInvites) CRC() uint32 {
	return 0x10ab6dc7
}
func (*ChatlistsExportedInvites) _IChatlistsExportedInvites() {}

type IContactsBlocked interface {
	tl.Object
	_IContactsBlocked()
}

var (
	_ IContactsBlocked = (*ContactsBlocked)(nil)
	_ IContactsBlocked = (*ContactsBlockedSlice)(nil)
)

type ContactsBlocked struct {
	Blocked []IPeerBlocked
	Chats   []IChat
	Users   []IUser
}

func (*ContactsBlocked) CRC() uint32 {
	return 0xade1591
}
func (*ContactsBlocked) _IContactsBlocked() {}

type ContactsBlockedSlice struct {
	Count   int32
	Blocked []IPeerBlocked
	Chats   []IChat
	Users   []IUser
}

func (*ContactsBlockedSlice) CRC() uint32 {
	return 0xe1664194
}
func (*ContactsBlockedSlice) _IContactsBlocked() {}

// Info on the current user's contact list.
type IContactsContacts interface {
	tl.Object
	_IContactsContacts()
}

var (
	_ IContactsContacts = (*ContactsContactsNotModified)(nil)
	_ IContactsContacts = (*ContactsContacts)(nil)
)

// Contact list on the server is the same as the list on the client.
type ContactsContactsNotModified struct{}

func (*ContactsContactsNotModified) CRC() uint32 {
	return 0xb74ba9d2
}
func (*ContactsContactsNotModified) _IContactsContacts() {}

// The current user's contact list and info on users.
type ContactsContacts struct {
	Contacts   []IContact // Contact list
	SavedCount int32      // Number of contacts that were saved successfully
	Users      []IUser    // User list
}

func (*ContactsContacts) CRC() uint32 {
	return 0xeae87e42
}
func (*ContactsContacts) _IContactsContacts() {}

type IContactsFound interface {
	tl.Object
	_IContactsFound()
}

var (
	_ IContactsFound = (*ContactsFound)(nil)
)

type ContactsFound struct {
	MyResults []IPeer
	Results   []IPeer
	Chats     []IChat
	Users     []IUser
}

func (*ContactsFound) CRC() uint32 {
	return 0xb3134d9d
}
func (*ContactsFound) _IContactsFound() {}

type IContactsImportedContacts interface {
	tl.Object
	_IContactsImportedContacts()
}

var (
	_ IContactsImportedContacts = (*ContactsImportedContacts)(nil)
)

type ContactsImportedContacts struct {
	Imported       []IImportedContact
	PopularInvites []IPopularContact
	RetryContacts  []int64
	Users          []IUser
}

func (*ContactsImportedContacts) CRC() uint32 {
	return 0x77d01c3b
}
func (*ContactsImportedContacts) _IContactsImportedContacts() {}

type IContactsResolvedPeer interface {
	tl.Object
	_IContactsResolvedPeer()
}

var (
	_ IContactsResolvedPeer = (*ContactsResolvedPeer)(nil)
)

type ContactsResolvedPeer struct {
	Peer  IPeer
	Chats []IChat
	Users []IUser
}

func (*ContactsResolvedPeer) CRC() uint32 {
	return 0x7f077ad9
}
func (*ContactsResolvedPeer) _IContactsResolvedPeer() {}

type IContactsTopPeers interface {
	tl.Object
	_IContactsTopPeers()
}

var (
	_ IContactsTopPeers = (*ContactsTopPeersNotModified)(nil)
	_ IContactsTopPeers = (*ContactsTopPeers)(nil)
	_ IContactsTopPeers = (*ContactsTopPeersDisabled)(nil)
)

type ContactsTopPeersNotModified struct{}

func (*ContactsTopPeersNotModified) CRC() uint32 {
	return 0xde266ef5
}
func (*ContactsTopPeersNotModified) _IContactsTopPeers() {}

type ContactsTopPeers struct {
	Categories []ITopPeerCategoryPeers
	Chats      []IChat
	Users      []IUser
}

func (*ContactsTopPeers) CRC() uint32 {
	return 0x70b772a8
}
func (*ContactsTopPeers) _IContactsTopPeers() {}

type ContactsTopPeersDisabled struct{}

func (*ContactsTopPeersDisabled) CRC() uint32 {
	return 0xb52c939d
}
func (*ContactsTopPeersDisabled) _IContactsTopPeers() {}

type IHelpAppConfig interface {
	tl.Object
	_IHelpAppConfig()
}

var (
	_ IHelpAppConfig = (*HelpAppConfigNotModified)(nil)
	_ IHelpAppConfig = (*HelpAppConfig)(nil)
)

type HelpAppConfigNotModified struct{}

func (*HelpAppConfigNotModified) CRC() uint32 {
	return 0x7cde641d
}
func (*HelpAppConfigNotModified) _IHelpAppConfig() {}

type HelpAppConfig struct {
	Hash   int32
	Config IJSONValue
}

func (*HelpAppConfig) CRC() uint32 {
	return 0xdd18782e
}
func (*HelpAppConfig) _IHelpAppConfig() {}

type IHelpAppUpdate interface {
	tl.Object
	_IHelpAppUpdate()
}

var (
	_ IHelpAppUpdate = (*HelpAppUpdate)(nil)
	_ IHelpAppUpdate = (*HelpNoAppUpdate)(nil)
)

type HelpAppUpdate struct {
	_          struct{} `tl:"flags,bitflag"`
	CanNotSkip bool     `tl:",omitempty:flags:0,implicit"`
	ID         int32
	Version    string
	Text       string
	Entities   []IMessageEntity
	Document   IDocument `tl:",omitempty:flags:1"`
	URL        *string   `tl:",omitempty:flags:2"`
	Sticker    IDocument `tl:",omitempty:flags:3"`
}

func (*HelpAppUpdate) CRC() uint32 {
	return 0xccbbce30
}
func (*HelpAppUpdate) _IHelpAppUpdate() {}

type HelpNoAppUpdate struct{}

func (*HelpNoAppUpdate) CRC() uint32 {
	return 0xc45a6536
}
func (*HelpNoAppUpdate) _IHelpAppUpdate() {}

type IHelpCountriesList interface {
	tl.Object
	_IHelpCountriesList()
}

var (
	_ IHelpCountriesList = (*HelpCountriesListNotModified)(nil)
	_ IHelpCountriesList = (*HelpCountriesList)(nil)
)

type HelpCountriesListNotModified struct{}

func (*HelpCountriesListNotModified) CRC() uint32 {
	return 0x93cc1f32
}
func (*HelpCountriesListNotModified) _IHelpCountriesList() {}

type HelpCountriesList struct {
	Countries []IHelpCountry
	Hash      int32
}

func (*HelpCountriesList) CRC() uint32 {
	return 0x87d0759e
}
func (*HelpCountriesList) _IHelpCountriesList() {}

type IHelpCountry interface {
	tl.Object
	_IHelpCountry()
}

var (
	_ IHelpCountry = (*HelpCountry)(nil)
)

type HelpCountry struct {
	_            struct{} `tl:"flags,bitflag"`
	Hidden       bool     `tl:",omitempty:flags:0,implicit"`
	Iso2         string
	DefaultName  string
	Name         *string `tl:",omitempty:flags:1"`
	CountryCodes []IHelpCountryCode
}

func (*HelpCountry) CRC() uint32 {
	return 0xc3878e23
}
func (*HelpCountry) _IHelpCountry() {}

type IHelpCountryCode interface {
	tl.Object
	_IHelpCountryCode()
}

var (
	_ IHelpCountryCode = (*HelpCountryCode)(nil)
)

type HelpCountryCode struct {
	_           struct{} `tl:"flags,bitflag"`
	CountryCode string
	Prefixes    []string `tl:",omitempty:flags:0"`
	Patterns    []string `tl:",omitempty:flags:1"`
}

func (*HelpCountryCode) CRC() uint32 {
	return 0x4203c5ef
}
func (*HelpCountryCode) _IHelpCountryCode() {}

type IHelpDeepLinkInfo interface {
	tl.Object
	_IHelpDeepLinkInfo()
}

var (
	_ IHelpDeepLinkInfo = (*HelpDeepLinkInfoEmpty)(nil)
	_ IHelpDeepLinkInfo = (*HelpDeepLinkInfo)(nil)
)

type HelpDeepLinkInfoEmpty struct{}

func (*HelpDeepLinkInfoEmpty) CRC() uint32 {
	return 0x66afa166
}
func (*HelpDeepLinkInfoEmpty) _IHelpDeepLinkInfo() {}

type HelpDeepLinkInfo struct {
	_         struct{} `tl:"flags,bitflag"`
	UpdateApp bool     `tl:",omitempty:flags:0,implicit"`
	Message   string
	Entities  []IMessageEntity `tl:",omitempty:flags:1"`
}

func (*HelpDeepLinkInfo) CRC() uint32 {
	return 0x6a4ee832
}
func (*HelpDeepLinkInfo) _IHelpDeepLinkInfo() {}

type IHelpInviteText interface {
	tl.Object
	_IHelpInviteText()
}

var (
	_ IHelpInviteText = (*HelpInviteText)(nil)
)

type HelpInviteText struct {
	Message string
}

func (*HelpInviteText) CRC() uint32 {
	return 0x18cb9f78
}
func (*HelpInviteText) _IHelpInviteText() {}

type IHelpPassportConfig interface {
	tl.Object
	_IHelpPassportConfig()
}

var (
	_ IHelpPassportConfig = (*HelpPassportConfigNotModified)(nil)
	_ IHelpPassportConfig = (*HelpPassportConfig)(nil)
)

type HelpPassportConfigNotModified struct{}

func (*HelpPassportConfigNotModified) CRC() uint32 {
	return 0xbfb9f457
}
func (*HelpPassportConfigNotModified) _IHelpPassportConfig() {}

type HelpPassportConfig struct {
	Hash           int32
	CountriesLangs IDataJSON
}

func (*HelpPassportConfig) CRC() uint32 {
	return 0xa098d6af
}
func (*HelpPassportConfig) _IHelpPassportConfig() {}

type IHelpPeerColorOption interface {
	tl.Object
	_IHelpPeerColorOption()
}

var (
	_ IHelpPeerColorOption = (*HelpPeerColorOption)(nil)
)

type HelpPeerColorOption struct {
	_               struct{} `tl:"flags,bitflag"`
	Hidden          bool     `tl:",omitempty:flags:0,implicit"`
	ColorID         int32
	Colors          IHelpPeerColorSet `tl:",omitempty:flags:1"`
	DarkColors      IHelpPeerColorSet `tl:",omitempty:flags:2"`
	ChannelMinLevel *int32            `tl:",omitempty:flags:3"`
}

func (*HelpPeerColorOption) CRC() uint32 {
	return 0xef8430ab
}
func (*HelpPeerColorOption) _IHelpPeerColorOption() {}

type IHelpPeerColorSet interface {
	tl.Object
	_IHelpPeerColorSet()
}

var (
	_ IHelpPeerColorSet = (*HelpPeerColorSet)(nil)
	_ IHelpPeerColorSet = (*HelpPeerColorProfileSet)(nil)
)

type HelpPeerColorSet struct {
	Colors []int32
}

func (*HelpPeerColorSet) CRC() uint32 {
	return 0x26219a58
}
func (*HelpPeerColorSet) _IHelpPeerColorSet() {}

type HelpPeerColorProfileSet struct {
	PaletteColors []int32
	BgColors      []int32
	StoryColors   []int32
}

func (*HelpPeerColorProfileSet) CRC() uint32 {
	return 0x767d61eb
}
func (*HelpPeerColorProfileSet) _IHelpPeerColorSet() {}

type IHelpPeerColors interface {
	tl.Object
	_IHelpPeerColors()
}

var (
	_ IHelpPeerColors = (*HelpPeerColorsNotModified)(nil)
	_ IHelpPeerColors = (*HelpPeerColors)(nil)
)

type HelpPeerColorsNotModified struct{}

func (*HelpPeerColorsNotModified) CRC() uint32 {
	return 0x2ba1f5ce
}
func (*HelpPeerColorsNotModified) _IHelpPeerColors() {}

type HelpPeerColors struct {
	Hash   int32
	Colors []IHelpPeerColorOption
}

func (*HelpPeerColors) CRC() uint32 {
	return 0xf8ed08
}
func (*HelpPeerColors) _IHelpPeerColors() {}

type IHelpPremiumPromo interface {
	tl.Object
	_IHelpPremiumPromo()
}

var (
	_ IHelpPremiumPromo = (*HelpPremiumPromo)(nil)
)

type HelpPremiumPromo struct {
	StatusText     string
	StatusEntities []IMessageEntity
	VideoSections  []string
	Videos         []IDocument
	PeriodOptions  []IPremiumSubscriptionOption
	Users          []IUser
}

func (*HelpPremiumPromo) CRC() uint32 {
	return 0x5334759c
}
func (*HelpPremiumPromo) _IHelpPremiumPromo() {}

type IHelpPromoData interface {
	tl.Object
	_IHelpPromoData()
}

var (
	_ IHelpPromoData = (*HelpPromoDataEmpty)(nil)
	_ IHelpPromoData = (*HelpPromoData)(nil)
)

type HelpPromoDataEmpty struct {
	Expires int32
}

func (*HelpPromoDataEmpty) CRC() uint32 {
	return 0x98f6ac75
}
func (*HelpPromoDataEmpty) _IHelpPromoData() {}

type HelpPromoData struct {
	_          struct{} `tl:"flags,bitflag"`
	Proxy      bool     `tl:",omitempty:flags:0,implicit"`
	Expires    int32
	Peer       IPeer
	Chats      []IChat
	Users      []IUser
	PsaType    *string `tl:",omitempty:flags:1"`
	PsaMessage *string `tl:",omitempty:flags:2"`
}

func (*HelpPromoData) CRC() uint32 {
	return 0x8c39793f
}
func (*HelpPromoData) _IHelpPromoData() {}

type IHelpRecentMeUrls interface {
	tl.Object
	_IHelpRecentMeUrls()
}

var (
	_ IHelpRecentMeUrls = (*HelpRecentMeUrls)(nil)
)

type HelpRecentMeUrls struct {
	Urls  []IRecentMeURL
	Chats []IChat
	Users []IUser
}

func (*HelpRecentMeUrls) CRC() uint32 {
	return 0xe0310d7
}
func (*HelpRecentMeUrls) _IHelpRecentMeUrls() {}

type IHelpSupport interface {
	tl.Object
	_IHelpSupport()
}

var (
	_ IHelpSupport = (*HelpSupport)(nil)
)

type HelpSupport struct {
	PhoneNumber string
	User        IUser
}

func (*HelpSupport) CRC() uint32 {
	return 0x17c6b5f6
}
func (*HelpSupport) _IHelpSupport() {}

type IHelpSupportName interface {
	tl.Object
	_IHelpSupportName()
}

var (
	_ IHelpSupportName = (*HelpSupportName)(nil)
)

type HelpSupportName struct {
	Name string
}

func (*HelpSupportName) CRC() uint32 {
	return 0x8c05f1c9
}
func (*HelpSupportName) _IHelpSupportName() {}

type IHelpTermsOfService interface {
	tl.Object
	_IHelpTermsOfService()
}

var (
	_ IHelpTermsOfService = (*HelpTermsOfService)(nil)
)

type HelpTermsOfService struct {
	_             struct{} `tl:"flags,bitflag"`
	Popup         bool     `tl:",omitempty:flags:0,implicit"`
	ID            IDataJSON
	Text          string
	Entities      []IMessageEntity
	MinAgeConfirm *int32 `tl:",omitempty:flags:1"`
}

func (*HelpTermsOfService) CRC() uint32 {
	return 0x780a0310
}
func (*HelpTermsOfService) _IHelpTermsOfService() {}

type IHelpTermsOfServiceUpdate interface {
	tl.Object
	_IHelpTermsOfServiceUpdate()
}

var (
	_ IHelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdateEmpty)(nil)
	_ IHelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdate)(nil)
)

type HelpTermsOfServiceUpdateEmpty struct {
	Expires int32
}

func (*HelpTermsOfServiceUpdateEmpty) CRC() uint32 {
	return 0xe3309f7f
}
func (*HelpTermsOfServiceUpdateEmpty) _IHelpTermsOfServiceUpdate() {}

type HelpTermsOfServiceUpdate struct {
	Expires        int32
	TermsOfService IHelpTermsOfService
}

func (*HelpTermsOfServiceUpdate) CRC() uint32 {
	return 0x28ecf961
}
func (*HelpTermsOfServiceUpdate) _IHelpTermsOfServiceUpdate() {}

type IHelpUserInfo interface {
	tl.Object
	_IHelpUserInfo()
}

var (
	_ IHelpUserInfo = (*HelpUserInfoEmpty)(nil)
	_ IHelpUserInfo = (*HelpUserInfo)(nil)
)

type HelpUserInfoEmpty struct{}

func (*HelpUserInfoEmpty) CRC() uint32 {
	return 0xf3ae2eed
}
func (*HelpUserInfoEmpty) _IHelpUserInfo() {}

type HelpUserInfo struct {
	Message  string
	Entities []IMessageEntity
	Author   string
	Date     int32
}

func (*HelpUserInfo) CRC() uint32 {
	return 0x1eb3758
}
func (*HelpUserInfo) _IHelpUserInfo() {}

type IMessagesAffectedFoundMessages interface {
	tl.Object
	_IMessagesAffectedFoundMessages()
}

var (
	_ IMessagesAffectedFoundMessages = (*MessagesAffectedFoundMessages)(nil)
)

type MessagesAffectedFoundMessages struct {
	Pts      int32
	PtsCount int32
	Offset   int32
	Messages []int32
}

func (*MessagesAffectedFoundMessages) CRC() uint32 {
	return 0xef8d3e6c
}
func (*MessagesAffectedFoundMessages) _IMessagesAffectedFoundMessages() {}

type IMessagesAffectedHistory interface {
	tl.Object
	_IMessagesAffectedHistory()
}

var (
	_ IMessagesAffectedHistory = (*MessagesAffectedHistory)(nil)
)

type MessagesAffectedHistory struct {
	Pts      int32
	PtsCount int32
	Offset   int32
}

func (*MessagesAffectedHistory) CRC() uint32 {
	return 0xb45c69d1
}
func (*MessagesAffectedHistory) _IMessagesAffectedHistory() {}

type IMessagesAffectedMessages interface {
	tl.Object
	_IMessagesAffectedMessages()
}

var (
	_ IMessagesAffectedMessages = (*MessagesAffectedMessages)(nil)
)

type MessagesAffectedMessages struct {
	Pts      int32
	PtsCount int32
}

func (*MessagesAffectedMessages) CRC() uint32 {
	return 0x84d19185
}
func (*MessagesAffectedMessages) _IMessagesAffectedMessages() {}

type IMessagesAllStickers interface {
	tl.Object
	_IMessagesAllStickers()
}

var (
	_ IMessagesAllStickers = (*MessagesAllStickersNotModified)(nil)
	_ IMessagesAllStickers = (*MessagesAllStickers)(nil)
)

type MessagesAllStickersNotModified struct{}

func (*MessagesAllStickersNotModified) CRC() uint32 {
	return 0xe86602c3
}
func (*MessagesAllStickersNotModified) _IMessagesAllStickers() {}

type MessagesAllStickers struct {
	Hash int64
	Sets []IStickerSet
}

func (*MessagesAllStickers) CRC() uint32 {
	return 0xcdbbcebb
}
func (*MessagesAllStickers) _IMessagesAllStickers() {}

type IMessagesArchivedStickers interface {
	tl.Object
	_IMessagesArchivedStickers()
}

var (
	_ IMessagesArchivedStickers = (*MessagesArchivedStickers)(nil)
)

type MessagesArchivedStickers struct {
	Count int32
	Sets  []IStickerSetCovered
}

func (*MessagesArchivedStickers) CRC() uint32 {
	return 0x4fcba9c8
}
func (*MessagesArchivedStickers) _IMessagesArchivedStickers() {}

type IMessagesAvailableReactions interface {
	tl.Object
	_IMessagesAvailableReactions()
}

var (
	_ IMessagesAvailableReactions = (*MessagesAvailableReactionsNotModified)(nil)
	_ IMessagesAvailableReactions = (*MessagesAvailableReactions)(nil)
)

type MessagesAvailableReactionsNotModified struct{}

func (*MessagesAvailableReactionsNotModified) CRC() uint32 {
	return 0x9f071957
}
func (*MessagesAvailableReactionsNotModified) _IMessagesAvailableReactions() {}

type MessagesAvailableReactions struct {
	Hash      int32
	Reactions []IAvailableReaction
}

func (*MessagesAvailableReactions) CRC() uint32 {
	return 0x768e3aad
}
func (*MessagesAvailableReactions) _IMessagesAvailableReactions() {}

type IMessagesBotApp interface {
	tl.Object
	_IMessagesBotApp()
}

var (
	_ IMessagesBotApp = (*MessagesBotApp)(nil)
)

type MessagesBotApp struct {
	_                  struct{} `tl:"flags,bitflag"`
	Inactive           bool     `tl:",omitempty:flags:0,implicit"`
	RequestWriteAccess bool     `tl:",omitempty:flags:1,implicit"`
	HasSettings        bool     `tl:",omitempty:flags:2,implicit"`
	App                IBotApp
}

func (*MessagesBotApp) CRC() uint32 {
	return 0xeb50adf5
}
func (*MessagesBotApp) _IMessagesBotApp() {}

type IMessagesBotCallbackAnswer interface {
	tl.Object
	_IMessagesBotCallbackAnswer()
}

var (
	_ IMessagesBotCallbackAnswer = (*MessagesBotCallbackAnswer)(nil)
)

type MessagesBotCallbackAnswer struct {
	_         struct{} `tl:"flags,bitflag"`
	Alert     bool     `tl:",omitempty:flags:1,implicit"`
	HasURL    bool     `tl:",omitempty:flags:3,implicit"`
	NativeUi  bool     `tl:",omitempty:flags:4,implicit"`
	Message   *string  `tl:",omitempty:flags:0"`
	URL       *string  `tl:",omitempty:flags:2"`
	CacheTime int32
}

func (*MessagesBotCallbackAnswer) CRC() uint32 {
	return 0x36585ea4
}
func (*MessagesBotCallbackAnswer) _IMessagesBotCallbackAnswer() {}

type IMessagesBotResults interface {
	tl.Object
	_IMessagesBotResults()
}

var (
	_ IMessagesBotResults = (*MessagesBotResults)(nil)
)

type MessagesBotResults struct {
	_             struct{} `tl:"flags,bitflag"`
	Gallery       bool     `tl:",omitempty:flags:0,implicit"`
	QueryID       int64
	NextOffset    *string            `tl:",omitempty:flags:1"`
	SwitchPm      IInlineBotSwitchPm `tl:",omitempty:flags:2"`
	SwitchWebview IInlineBotWebView  `tl:",omitempty:flags:3"`
	Results       []IBotInlineResult
	CacheTime     int32
	Users         []IUser
}

func (*MessagesBotResults) CRC() uint32 {
	return 0xe021f2f6
}
func (*MessagesBotResults) _IMessagesBotResults() {}

type IMessagesChatAdminsWithInvites interface {
	tl.Object
	_IMessagesChatAdminsWithInvites()
}

var (
	_ IMessagesChatAdminsWithInvites = (*MessagesChatAdminsWithInvites)(nil)
)

type MessagesChatAdminsWithInvites struct {
	Admins []IChatAdminWithInvites
	Users  []IUser
}

func (*MessagesChatAdminsWithInvites) CRC() uint32 {
	return 0xb69b72d7
}
func (*MessagesChatAdminsWithInvites) _IMessagesChatAdminsWithInvites() {}

type IMessagesChatFull interface {
	tl.Object
	_IMessagesChatFull()
}

var (
	_ IMessagesChatFull = (*MessagesChatFull)(nil)
)

type MessagesChatFull struct {
	FullChat IChatFull
	Chats    []IChat
	Users    []IUser
}

func (*MessagesChatFull) CRC() uint32 {
	return 0xe5d7d19c
}
func (*MessagesChatFull) _IMessagesChatFull() {}

type IMessagesChatInviteImporters interface {
	tl.Object
	_IMessagesChatInviteImporters()
}

var (
	_ IMessagesChatInviteImporters = (*MessagesChatInviteImporters)(nil)
)

type MessagesChatInviteImporters struct {
	Count     int32
	Importers []IChatInviteImporter
	Users     []IUser
}

func (*MessagesChatInviteImporters) CRC() uint32 {
	return 0x81b6b00a
}
func (*MessagesChatInviteImporters) _IMessagesChatInviteImporters() {}

type IMessagesChats interface {
	tl.Object
	_IMessagesChats()
}

var (
	_ IMessagesChats = (*MessagesChats)(nil)
	_ IMessagesChats = (*MessagesChatsSlice)(nil)
)

type MessagesChats struct {
	Chats []IChat
}

func (*MessagesChats) CRC() uint32 {
	return 0x64ff9fd5
}
func (*MessagesChats) _IMessagesChats() {}

type MessagesChatsSlice struct {
	Count int32
	Chats []IChat
}

func (*MessagesChatsSlice) CRC() uint32 {
	return 0x9cd81144
}
func (*MessagesChatsSlice) _IMessagesChats() {}

type IMessagesCheckedHistoryImportPeer interface {
	tl.Object
	_IMessagesCheckedHistoryImportPeer()
}

var (
	_ IMessagesCheckedHistoryImportPeer = (*MessagesCheckedHistoryImportPeer)(nil)
)

type MessagesCheckedHistoryImportPeer struct {
	ConfirmText string
}

func (*MessagesCheckedHistoryImportPeer) CRC() uint32 {
	return 0xa24de717
}
func (*MessagesCheckedHistoryImportPeer) _IMessagesCheckedHistoryImportPeer() {}

type IMessagesDhConfig interface {
	tl.Object
	_IMessagesDhConfig()
}

var (
	_ IMessagesDhConfig = (*MessagesDhConfigNotModified)(nil)
	_ IMessagesDhConfig = (*MessagesDhConfig)(nil)
)

type MessagesDhConfigNotModified struct {
	Random []byte
}

func (*MessagesDhConfigNotModified) CRC() uint32 {
	return 0xc0e24635
}
func (*MessagesDhConfigNotModified) _IMessagesDhConfig() {}

type MessagesDhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

func (*MessagesDhConfig) CRC() uint32 {
	return 0x2c221edd
}
func (*MessagesDhConfig) _IMessagesDhConfig() {}

type IMessagesDialogs interface {
	tl.Object
	_IMessagesDialogs()
}

var (
	_ IMessagesDialogs = (*MessagesDialogs)(nil)
	_ IMessagesDialogs = (*MessagesDialogsSlice)(nil)
	_ IMessagesDialogs = (*MessagesDialogsNotModified)(nil)
)

type MessagesDialogs struct {
	Dialogs  []IDialog
	Messages []IMessage
	Chats    []IChat
	Users    []IUser
}

func (*MessagesDialogs) CRC() uint32 {
	return 0x15ba6c40
}
func (*MessagesDialogs) _IMessagesDialogs() {}

type MessagesDialogsSlice struct {
	Count    int32
	Dialogs  []IDialog
	Messages []IMessage
	Chats    []IChat
	Users    []IUser
}

func (*MessagesDialogsSlice) CRC() uint32 {
	return 0x71e094f3
}
func (*MessagesDialogsSlice) _IMessagesDialogs() {}

type MessagesDialogsNotModified struct {
	Count int32
}

func (*MessagesDialogsNotModified) CRC() uint32 {
	return 0xf0e3e596
}
func (*MessagesDialogsNotModified) _IMessagesDialogs() {}

type IMessagesDiscussionMessage interface {
	tl.Object
	_IMessagesDiscussionMessage()
}

var (
	_ IMessagesDiscussionMessage = (*MessagesDiscussionMessage)(nil)
)

type MessagesDiscussionMessage struct {
	_               struct{} `tl:"flags,bitflag"`
	Messages        []IMessage
	MaxID           *int32 `tl:",omitempty:flags:0"`
	ReadInboxMaxID  *int32 `tl:",omitempty:flags:1"`
	ReadOutboxMaxID *int32 `tl:",omitempty:flags:2"`
	UnreadCount     int32
	Chats           []IChat
	Users           []IUser
}

func (*MessagesDiscussionMessage) CRC() uint32 {
	return 0xa6341782
}
func (*MessagesDiscussionMessage) _IMessagesDiscussionMessage() {}

type IMessagesEmojiGroups interface {
	tl.Object
	_IMessagesEmojiGroups()
}

var (
	_ IMessagesEmojiGroups = (*MessagesEmojiGroupsNotModified)(nil)
	_ IMessagesEmojiGroups = (*MessagesEmojiGroups)(nil)
)

type MessagesEmojiGroupsNotModified struct{}

func (*MessagesEmojiGroupsNotModified) CRC() uint32 {
	return 0x6fb4ad87
}
func (*MessagesEmojiGroupsNotModified) _IMessagesEmojiGroups() {}

type MessagesEmojiGroups struct {
	Hash   int32
	Groups []IEmojiGroup
}

func (*MessagesEmojiGroups) CRC() uint32 {
	return 0x881fb94b
}
func (*MessagesEmojiGroups) _IMessagesEmojiGroups() {}

type IMessagesExportedChatInvite interface {
	tl.Object
	_IMessagesExportedChatInvite()
}

var (
	_ IMessagesExportedChatInvite = (*MessagesExportedChatInvite)(nil)
	_ IMessagesExportedChatInvite = (*MessagesExportedChatInviteReplaced)(nil)
)

type MessagesExportedChatInvite struct {
	Invite IExportedChatInvite
	Users  []IUser
}

func (*MessagesExportedChatInvite) CRC() uint32 {
	return 0x1871be50
}
func (*MessagesExportedChatInvite) _IMessagesExportedChatInvite() {}

type MessagesExportedChatInviteReplaced struct {
	Invite    IExportedChatInvite
	NewInvite IExportedChatInvite
	Users     []IUser
}

func (*MessagesExportedChatInviteReplaced) CRC() uint32 {
	return 0x222600ef
}
func (*MessagesExportedChatInviteReplaced) _IMessagesExportedChatInvite() {}

type IMessagesExportedChatInvites interface {
	tl.Object
	_IMessagesExportedChatInvites()
}

var (
	_ IMessagesExportedChatInvites = (*MessagesExportedChatInvites)(nil)
)

type MessagesExportedChatInvites struct {
	Count   int32
	Invites []IExportedChatInvite
	Users   []IUser
}

func (*MessagesExportedChatInvites) CRC() uint32 {
	return 0xbdc62dcc
}
func (*MessagesExportedChatInvites) _IMessagesExportedChatInvites() {}

type IMessagesFavedStickers interface {
	tl.Object
	_IMessagesFavedStickers()
}

var (
	_ IMessagesFavedStickers = (*MessagesFavedStickersNotModified)(nil)
	_ IMessagesFavedStickers = (*MessagesFavedStickers)(nil)
)

type MessagesFavedStickersNotModified struct{}

func (*MessagesFavedStickersNotModified) CRC() uint32 {
	return 0x9e8fa6d3
}
func (*MessagesFavedStickersNotModified) _IMessagesFavedStickers() {}

type MessagesFavedStickers struct {
	Hash     int64
	Packs    []IStickerPack
	Stickers []IDocument
}

func (*MessagesFavedStickers) CRC() uint32 {
	return 0x2cb51097
}
func (*MessagesFavedStickers) _IMessagesFavedStickers() {}

type IMessagesFeaturedStickers interface {
	tl.Object
	_IMessagesFeaturedStickers()
}

var (
	_ IMessagesFeaturedStickers = (*MessagesFeaturedStickersNotModified)(nil)
	_ IMessagesFeaturedStickers = (*MessagesFeaturedStickers)(nil)
)

type MessagesFeaturedStickersNotModified struct {
	Count int32
}

func (*MessagesFeaturedStickersNotModified) CRC() uint32 {
	return 0xc6dc0c66
}
func (*MessagesFeaturedStickersNotModified) _IMessagesFeaturedStickers() {}

type MessagesFeaturedStickers struct {
	_       struct{} `tl:"flags,bitflag"`
	Premium bool     `tl:",omitempty:flags:0,implicit"`
	Hash    int64
	Count   int32
	Sets    []IStickerSetCovered
	Unread  []int64
}

func (*MessagesFeaturedStickers) CRC() uint32 {
	return 0xbe382906
}
func (*MessagesFeaturedStickers) _IMessagesFeaturedStickers() {}

type IMessagesForumTopics interface {
	tl.Object
	_IMessagesForumTopics()
}

var (
	_ IMessagesForumTopics = (*MessagesForumTopics)(nil)
)

type MessagesForumTopics struct {
	_                 struct{} `tl:"flags,bitflag"`
	OrderByCreateDate bool     `tl:",omitempty:flags:0,implicit"`
	Count             int32
	Topics            []IForumTopic
	Messages          []IMessage
	Chats             []IChat
	Users             []IUser
	Pts               int32
}

func (*MessagesForumTopics) CRC() uint32 {
	return 0x367617d3
}
func (*MessagesForumTopics) _IMessagesForumTopics() {}

type IMessagesFoundStickerSets interface {
	tl.Object
	_IMessagesFoundStickerSets()
}

var (
	_ IMessagesFoundStickerSets = (*MessagesFoundStickerSetsNotModified)(nil)
	_ IMessagesFoundStickerSets = (*MessagesFoundStickerSets)(nil)
)

type MessagesFoundStickerSetsNotModified struct{}

func (*MessagesFoundStickerSetsNotModified) CRC() uint32 {
	return 0xd54b65d
}
func (*MessagesFoundStickerSetsNotModified) _IMessagesFoundStickerSets() {}

type MessagesFoundStickerSets struct {
	Hash int64
	Sets []IStickerSetCovered
}

func (*MessagesFoundStickerSets) CRC() uint32 {
	return 0x8af09dd2
}
func (*MessagesFoundStickerSets) _IMessagesFoundStickerSets() {}

type IMessagesHighScores interface {
	tl.Object
	_IMessagesHighScores()
}

var (
	_ IMessagesHighScores = (*MessagesHighScores)(nil)
)

type MessagesHighScores struct {
	Scores []IHighScore
	Users  []IUser
}

func (*MessagesHighScores) CRC() uint32 {
	return 0x9a3bfd99
}
func (*MessagesHighScores) _IMessagesHighScores() {}

type IMessagesHistoryImport interface {
	tl.Object
	_IMessagesHistoryImport()
}

var (
	_ IMessagesHistoryImport = (*MessagesHistoryImport)(nil)
)

type MessagesHistoryImport struct {
	ID int64
}

func (*MessagesHistoryImport) CRC() uint32 {
	return 0x1662af0b
}
func (*MessagesHistoryImport) _IMessagesHistoryImport() {}

type IMessagesHistoryImportParsed interface {
	tl.Object
	_IMessagesHistoryImportParsed()
}

var (
	_ IMessagesHistoryImportParsed = (*MessagesHistoryImportParsed)(nil)
)

type MessagesHistoryImportParsed struct {
	_     struct{} `tl:"flags,bitflag"`
	Pm    bool     `tl:",omitempty:flags:0,implicit"`
	Group bool     `tl:",omitempty:flags:1,implicit"`
	Title *string  `tl:",omitempty:flags:2"`
}

func (*MessagesHistoryImportParsed) CRC() uint32 {
	return 0x5e0fb7b9
}
func (*MessagesHistoryImportParsed) _IMessagesHistoryImportParsed() {}

type IMessagesInactiveChats interface {
	tl.Object
	_IMessagesInactiveChats()
}

var (
	_ IMessagesInactiveChats = (*MessagesInactiveChats)(nil)
)

type MessagesInactiveChats struct {
	Dates []int32
	Chats []IChat
	Users []IUser
}

func (*MessagesInactiveChats) CRC() uint32 {
	return 0xa927fec5
}
func (*MessagesInactiveChats) _IMessagesInactiveChats() {}

type IMessagesMessageEditData interface {
	tl.Object
	_IMessagesMessageEditData()
}

var (
	_ IMessagesMessageEditData = (*MessagesMessageEditData)(nil)
)

type MessagesMessageEditData struct {
	_       struct{} `tl:"flags,bitflag"`
	Caption bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesMessageEditData) CRC() uint32 {
	return 0x26b5dde6
}
func (*MessagesMessageEditData) _IMessagesMessageEditData() {}

type IMessagesMessageReactionsList interface {
	tl.Object
	_IMessagesMessageReactionsList()
}

var (
	_ IMessagesMessageReactionsList = (*MessagesMessageReactionsList)(nil)
)

type MessagesMessageReactionsList struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Reactions  []IMessagePeerReaction
	Chats      []IChat
	Users      []IUser
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*MessagesMessageReactionsList) CRC() uint32 {
	return 0x31bd492d
}
func (*MessagesMessageReactionsList) _IMessagesMessageReactionsList() {}

type IMessagesMessageViews interface {
	tl.Object
	_IMessagesMessageViews()
}

var (
	_ IMessagesMessageViews = (*MessagesMessageViews)(nil)
)

type MessagesMessageViews struct {
	Views []IMessageViews
	Chats []IChat
	Users []IUser
}

func (*MessagesMessageViews) CRC() uint32 {
	return 0xb6c4f543
}
func (*MessagesMessageViews) _IMessagesMessageViews() {}

type IMessagesMessages interface {
	tl.Object
	_IMessagesMessages()
}

var (
	_ IMessagesMessages = (*MessagesMessages)(nil)
	_ IMessagesMessages = (*MessagesMessagesSlice)(nil)
	_ IMessagesMessages = (*MessagesChannelMessages)(nil)
	_ IMessagesMessages = (*MessagesMessagesNotModified)(nil)
)

type MessagesMessages struct {
	Messages []IMessage
	Chats    []IChat
	Users    []IUser
}

func (*MessagesMessages) CRC() uint32 {
	return 0x8c718e87
}
func (*MessagesMessages) _IMessagesMessages() {}

type MessagesMessagesSlice struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:1,implicit"`
	Count          int32
	NextRate       *int32 `tl:",omitempty:flags:0"`
	OffsetIDOffset *int32 `tl:",omitempty:flags:2"`
	Messages       []IMessage
	Chats          []IChat
	Users          []IUser
}

func (*MessagesMessagesSlice) CRC() uint32 {
	return 0x3a54685e
}
func (*MessagesMessagesSlice) _IMessagesMessages() {}

type MessagesChannelMessages struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:1,implicit"`
	Pts            int32
	Count          int32
	OffsetIDOffset *int32 `tl:",omitempty:flags:2"`
	Messages       []IMessage
	Topics         []IForumTopic
	Chats          []IChat
	Users          []IUser
}

func (*MessagesChannelMessages) CRC() uint32 {
	return 0xc776ba4e
}
func (*MessagesChannelMessages) _IMessagesMessages() {}

type MessagesMessagesNotModified struct {
	Count int32
}

func (*MessagesMessagesNotModified) CRC() uint32 {
	return 0x74535f21
}
func (*MessagesMessagesNotModified) _IMessagesMessages() {}

type IMessagesPeerDialogs interface {
	tl.Object
	_IMessagesPeerDialogs()
}

var (
	_ IMessagesPeerDialogs = (*MessagesPeerDialogs)(nil)
)

type MessagesPeerDialogs struct {
	Dialogs  []IDialog
	Messages []IMessage
	Chats    []IChat
	Users    []IUser
	State    IUpdatesState
}

func (*MessagesPeerDialogs) CRC() uint32 {
	return 0x3371c354
}
func (*MessagesPeerDialogs) _IMessagesPeerDialogs() {}

type IMessagesPeerSettings interface {
	tl.Object
	_IMessagesPeerSettings()
}

var (
	_ IMessagesPeerSettings = (*MessagesPeerSettings)(nil)
)

type MessagesPeerSettings struct {
	Settings IPeerSettings
	Chats    []IChat
	Users    []IUser
}

func (*MessagesPeerSettings) CRC() uint32 {
	return 0x6880b94d
}
func (*MessagesPeerSettings) _IMessagesPeerSettings() {}

type IMessagesReactions interface {
	tl.Object
	_IMessagesReactions()
}

var (
	_ IMessagesReactions = (*MessagesReactionsNotModified)(nil)
	_ IMessagesReactions = (*MessagesReactions)(nil)
)

type MessagesReactionsNotModified struct{}

func (*MessagesReactionsNotModified) CRC() uint32 {
	return 0xb06fdbdf
}
func (*MessagesReactionsNotModified) _IMessagesReactions() {}

type MessagesReactions struct {
	Hash      int64
	Reactions []IReaction
}

func (*MessagesReactions) CRC() uint32 {
	return 0xeafdf716
}
func (*MessagesReactions) _IMessagesReactions() {}

type IMessagesRecentStickers interface {
	tl.Object
	_IMessagesRecentStickers()
}

var (
	_ IMessagesRecentStickers = (*MessagesRecentStickersNotModified)(nil)
	_ IMessagesRecentStickers = (*MessagesRecentStickers)(nil)
)

type MessagesRecentStickersNotModified struct{}

func (*MessagesRecentStickersNotModified) CRC() uint32 {
	return 0xb17f890
}
func (*MessagesRecentStickersNotModified) _IMessagesRecentStickers() {}

type MessagesRecentStickers struct {
	Hash     int64
	Packs    []IStickerPack
	Stickers []IDocument
	Dates    []int32
}

func (*MessagesRecentStickers) CRC() uint32 {
	return 0x88d37c56
}
func (*MessagesRecentStickers) _IMessagesRecentStickers() {}

type IMessagesSavedGifs interface {
	tl.Object
	_IMessagesSavedGifs()
}

var (
	_ IMessagesSavedGifs = (*MessagesSavedGifsNotModified)(nil)
	_ IMessagesSavedGifs = (*MessagesSavedGifs)(nil)
)

type MessagesSavedGifsNotModified struct{}

func (*MessagesSavedGifsNotModified) CRC() uint32 {
	return 0xe8025ca2
}
func (*MessagesSavedGifsNotModified) _IMessagesSavedGifs() {}

type MessagesSavedGifs struct {
	Hash int64
	Gifs []IDocument
}

func (*MessagesSavedGifs) CRC() uint32 {
	return 0x84a02a0d
}
func (*MessagesSavedGifs) _IMessagesSavedGifs() {}

type IMessagesSearchCounter interface {
	tl.Object
	_IMessagesSearchCounter()
}

var (
	_ IMessagesSearchCounter = (*MessagesSearchCounter)(nil)
)

type MessagesSearchCounter struct {
	_       struct{} `tl:"flags,bitflag"`
	Inexact bool     `tl:",omitempty:flags:1,implicit"`
	Filter  IMessagesFilter
	Count   int32
}

func (*MessagesSearchCounter) CRC() uint32 {
	return 0xe844ebff
}
func (*MessagesSearchCounter) _IMessagesSearchCounter() {}

type IMessagesSearchResultsCalendar interface {
	tl.Object
	_IMessagesSearchResultsCalendar()
}

var (
	_ IMessagesSearchResultsCalendar = (*MessagesSearchResultsCalendar)(nil)
)

type MessagesSearchResultsCalendar struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:0,implicit"`
	Count          int32
	MinDate        int32
	MinMsgID       int32
	OffsetIDOffset *int32 `tl:",omitempty:flags:1"`
	Periods        []ISearchResultsCalendarPeriod
	Messages       []IMessage
	Chats          []IChat
	Users          []IUser
}

func (*MessagesSearchResultsCalendar) CRC() uint32 {
	return 0x147ee23c
}
func (*MessagesSearchResultsCalendar) _IMessagesSearchResultsCalendar() {}

type IMessagesSearchResultsPositions interface {
	tl.Object
	_IMessagesSearchResultsPositions()
}

var (
	_ IMessagesSearchResultsPositions = (*MessagesSearchResultsPositions)(nil)
)

type MessagesSearchResultsPositions struct {
	Count     int32
	Positions []ISearchResultsPosition
}

func (*MessagesSearchResultsPositions) CRC() uint32 {
	return 0x53b22baf
}
func (*MessagesSearchResultsPositions) _IMessagesSearchResultsPositions() {}

type IMessagesSentEncryptedMessage interface {
	tl.Object
	_IMessagesSentEncryptedMessage()
}

var (
	_ IMessagesSentEncryptedMessage = (*MessagesSentEncryptedMessage)(nil)
	_ IMessagesSentEncryptedMessage = (*MessagesSentEncryptedFile)(nil)
)

type MessagesSentEncryptedMessage struct {
	Date int32
}

func (*MessagesSentEncryptedMessage) CRC() uint32 {
	return 0x560f8935
}
func (*MessagesSentEncryptedMessage) _IMessagesSentEncryptedMessage() {}

type MessagesSentEncryptedFile struct {
	Date int32
	File IEncryptedFile
}

func (*MessagesSentEncryptedFile) CRC() uint32 {
	return 0x9493ff32
}
func (*MessagesSentEncryptedFile) _IMessagesSentEncryptedMessage() {}

type IMessagesSponsoredMessages interface {
	tl.Object
	_IMessagesSponsoredMessages()
}

var (
	_ IMessagesSponsoredMessages = (*MessagesSponsoredMessages)(nil)
	_ IMessagesSponsoredMessages = (*MessagesSponsoredMessagesEmpty)(nil)
)

type MessagesSponsoredMessages struct {
	_            struct{} `tl:"flags,bitflag"`
	PostsBetween *int32   `tl:",omitempty:flags:0"`
	Messages     []ISponsoredMessage
	Chats        []IChat
	Users        []IUser
}

func (*MessagesSponsoredMessages) CRC() uint32 {
	return 0xc9ee1d87
}
func (*MessagesSponsoredMessages) _IMessagesSponsoredMessages() {}

type MessagesSponsoredMessagesEmpty struct{}

func (*MessagesSponsoredMessagesEmpty) CRC() uint32 {
	return 0x1839490f
}
func (*MessagesSponsoredMessagesEmpty) _IMessagesSponsoredMessages() {}

type IMessagesStickerSet interface {
	tl.Object
	_IMessagesStickerSet()
}

var (
	_ IMessagesStickerSet = (*MessagesStickerSet)(nil)
	_ IMessagesStickerSet = (*MessagesStickerSetNotModified)(nil)
)

type MessagesStickerSet struct {
	Set       IStickerSet
	Packs     []IStickerPack
	Keywords  []IStickerKeyword
	Documents []IDocument
}

func (*MessagesStickerSet) CRC() uint32 {
	return 0x6e153f16
}
func (*MessagesStickerSet) _IMessagesStickerSet() {}

type MessagesStickerSetNotModified struct{}

func (*MessagesStickerSetNotModified) CRC() uint32 {
	return 0xd3f924eb
}
func (*MessagesStickerSetNotModified) _IMessagesStickerSet() {}

type IMessagesStickerSetInstallResult interface {
	tl.Object
	_IMessagesStickerSetInstallResult()
}

var (
	_ IMessagesStickerSetInstallResult = (*MessagesStickerSetInstallResultSuccess)(nil)
	_ IMessagesStickerSetInstallResult = (*MessagesStickerSetInstallResultArchive)(nil)
)

type MessagesStickerSetInstallResultSuccess struct{}

func (*MessagesStickerSetInstallResultSuccess) CRC() uint32 {
	return 0x38641628
}
func (*MessagesStickerSetInstallResultSuccess) _IMessagesStickerSetInstallResult() {}

type MessagesStickerSetInstallResultArchive struct {
	Sets []IStickerSetCovered
}

func (*MessagesStickerSetInstallResultArchive) CRC() uint32 {
	return 0x35e410a8
}
func (*MessagesStickerSetInstallResultArchive) _IMessagesStickerSetInstallResult() {}

type IMessagesStickers interface {
	tl.Object
	_IMessagesStickers()
}

var (
	_ IMessagesStickers = (*MessagesStickersNotModified)(nil)
	_ IMessagesStickers = (*MessagesStickers)(nil)
)

type MessagesStickersNotModified struct{}

func (*MessagesStickersNotModified) CRC() uint32 {
	return 0xf1749a22
}
func (*MessagesStickersNotModified) _IMessagesStickers() {}

type MessagesStickers struct {
	Hash     int64
	Stickers []IDocument
}

func (*MessagesStickers) CRC() uint32 {
	return 0x30a6ec7e
}
func (*MessagesStickers) _IMessagesStickers() {}

type IMessagesTranscribedAudio interface {
	tl.Object
	_IMessagesTranscribedAudio()
}

var (
	_ IMessagesTranscribedAudio = (*MessagesTranscribedAudio)(nil)
)

type MessagesTranscribedAudio struct {
	_                     struct{} `tl:"flags,bitflag"`
	Pending               bool     `tl:",omitempty:flags:0,implicit"`
	TranscriptionID       int64
	Text                  string
	TrialRemainsNum       *int32 `tl:",omitempty:flags:1"`
	TrialRemainsUntilDate *int32 `tl:",omitempty:flags:1"`
}

func (*MessagesTranscribedAudio) CRC() uint32 {
	return 0xcfb9d957
}
func (*MessagesTranscribedAudio) _IMessagesTranscribedAudio() {}

type IMessagesTranslatedText interface {
	tl.Object
	_IMessagesTranslatedText()
}

var (
	_ IMessagesTranslatedText = (*MessagesTranslateResult)(nil)
)

type MessagesTranslateResult struct {
	Result []ITextWithEntities
}

func (*MessagesTranslateResult) CRC() uint32 {
	return 0x33db32f8
}
func (*MessagesTranslateResult) _IMessagesTranslatedText() {}

type IMessagesVotesList interface {
	tl.Object
	_IMessagesVotesList()
}

var (
	_ IMessagesVotesList = (*MessagesVotesList)(nil)
)

type MessagesVotesList struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Votes      []IMessagePeerVote
	Chats      []IChat
	Users      []IUser
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*MessagesVotesList) CRC() uint32 {
	return 0x4899484e
}
func (*MessagesVotesList) _IMessagesVotesList() {}

type IMessagesWebPage interface {
	tl.Object
	_IMessagesWebPage()
}

var (
	_ IMessagesWebPage = (*MessagesWebPage)(nil)
)

type MessagesWebPage struct {
	Webpage IWebPage
	Chats   []IChat
	Users   []IUser
}

func (*MessagesWebPage) CRC() uint32 {
	return 0xfd5e12bd
}
func (*MessagesWebPage) _IMessagesWebPage() {}

type IPaymentsBankCardData interface {
	tl.Object
	_IPaymentsBankCardData()
}

var (
	_ IPaymentsBankCardData = (*PaymentsBankCardData)(nil)
)

type PaymentsBankCardData struct {
	Title    string
	OpenUrls []IBankCardOpenURL
}

func (*PaymentsBankCardData) CRC() uint32 {
	return 0x3e24e573
}
func (*PaymentsBankCardData) _IPaymentsBankCardData() {}

type IPaymentsCheckedGiftCode interface {
	tl.Object
	_IPaymentsCheckedGiftCode()
}

var (
	_ IPaymentsCheckedGiftCode = (*PaymentsCheckedGiftCode)(nil)
)

type PaymentsCheckedGiftCode struct {
	_             struct{} `tl:"flags,bitflag"`
	ViaGiveaway   bool     `tl:",omitempty:flags:2,implicit"`
	FromID        IPeer    `tl:",omitempty:flags:4"`
	GiveawayMsgID *int32   `tl:",omitempty:flags:3"`
	ToID          *int64   `tl:",omitempty:flags:0"`
	Date          int32
	Months        int32
	UsedDate      *int32 `tl:",omitempty:flags:1"`
	Chats         []IChat
	Users         []IUser
}

func (*PaymentsCheckedGiftCode) CRC() uint32 {
	return 0x284a1096
}
func (*PaymentsCheckedGiftCode) _IPaymentsCheckedGiftCode() {}

type IPaymentsExportedInvoice interface {
	tl.Object
	_IPaymentsExportedInvoice()
}

var (
	_ IPaymentsExportedInvoice = (*PaymentsExportedInvoice)(nil)
)

type PaymentsExportedInvoice struct {
	URL string
}

func (*PaymentsExportedInvoice) CRC() uint32 {
	return 0xaed0cbd9
}
func (*PaymentsExportedInvoice) _IPaymentsExportedInvoice() {}

type IPaymentsGiveawayInfo interface {
	tl.Object
	_IPaymentsGiveawayInfo()
}

var (
	_ IPaymentsGiveawayInfo = (*PaymentsGiveawayInfo)(nil)
	_ IPaymentsGiveawayInfo = (*PaymentsGiveawayInfoResults)(nil)
)

type PaymentsGiveawayInfo struct {
	_                     struct{} `tl:"flags,bitflag"`
	Participating         bool     `tl:",omitempty:flags:0,implicit"`
	PreparingResults      bool     `tl:",omitempty:flags:3,implicit"`
	StartDate             int32
	JoinedTooEarlyDate    *int32  `tl:",omitempty:flags:1"`
	AdminDisallowedChatID *int64  `tl:",omitempty:flags:2"`
	DisallowedCountry     *string `tl:",omitempty:flags:4"`
}

func (*PaymentsGiveawayInfo) CRC() uint32 {
	return 0x4367daa0
}
func (*PaymentsGiveawayInfo) _IPaymentsGiveawayInfo() {}

type PaymentsGiveawayInfoResults struct {
	_              struct{} `tl:"flags,bitflag"`
	Winner         bool     `tl:",omitempty:flags:0,implicit"`
	Refunded       bool     `tl:",omitempty:flags:1,implicit"`
	StartDate      int32
	GiftCodeSlug   *string `tl:",omitempty:flags:0"`
	FinishDate     int32
	WinnersCount   int32
	ActivatedCount int32
}

func (*PaymentsGiveawayInfoResults) CRC() uint32 {
	return 0xcd5570
}
func (*PaymentsGiveawayInfoResults) _IPaymentsGiveawayInfo() {}

type IPaymentsPaymentForm interface {
	tl.Object
	_IPaymentsPaymentForm()
}

var (
	_ IPaymentsPaymentForm = (*PaymentsPaymentForm)(nil)
)

// Payment form
type PaymentsPaymentForm struct {
	_                  struct{} `tl:"flags,bitflag"`               // Flags, see TL conditional fields
	CanSaveCredentials bool     `tl:",omitempty:flags:2,implicit"` // Whether the user can choose to save credentials.
	PasswordMissing    bool     `tl:",omitempty:flags:3,implicit"` // Indicates that the user can save payment credentials, but only after setting up a 2FA password (currently the account doesn't have a 2FA password)
	FormID             int64
	BotID              int64 // Bot ID
	Title              string
	Description        string
	Photo              IWebDocument               `tl:",omitempty:flags:5"`
	Invoice            IInvoice                   // Invoice
	ProviderID         int64                      // Payment provider ID.
	URL                string                     // Payment form URL
	NativeProvider     *string                    `tl:",omitempty:flags:4"` // Payment provider name.One of the following:- stripe
	NativeParams       IDataJSON                  `tl:",omitempty:flags:4"` // Contains information about the payment provider, if available, to support it natively without the need for opening the URL.A JSON object that can contain the following fields:- publishable_key: Stripe API publishable key- apple_pay_merchant_id: Apple Pay merchant ID- android_pay_public_key: Android Pay public key- android_pay_bgcolor: Android Pay form background color- android_pay_inverse: Whether to use the dark theme in the Android Pay form- need_country: True, if the user country must be provided,- need_zip: True, if the user ZIP/postal code must be provided,- need_cardholder_name: True, if the cardholder name must be provided
	AdditionalMethods  []IPaymentFormMethod       `tl:",omitempty:flags:6"`
	SavedInfo          IPaymentRequestedInfo      `tl:",omitempty:flags:0"` // Saved server-side order information
	SavedCredentials   []IPaymentSavedCredentials `tl:",omitempty:flags:1"` // Contains information about saved card credentials
	Users              []IUser                    // Users
}

func (*PaymentsPaymentForm) CRC() uint32 {
	return 0xa0058751
}
func (*PaymentsPaymentForm) _IPaymentsPaymentForm() {}

type IPaymentsPaymentReceipt interface {
	tl.Object
	_IPaymentsPaymentReceipt()
}

var (
	_ IPaymentsPaymentReceipt = (*PaymentsPaymentReceipt)(nil)
)

type PaymentsPaymentReceipt struct {
	_                struct{} `tl:"flags,bitflag"`
	Date             int32
	BotID            int64
	ProviderID       int64
	Title            string
	Description      string
	Photo            IWebDocument `tl:",omitempty:flags:2"`
	Invoice          IInvoice
	Info             IPaymentRequestedInfo `tl:",omitempty:flags:0"`
	Shipping         IShippingOption       `tl:",omitempty:flags:1"`
	TipAmount        *int64                `tl:",omitempty:flags:3"`
	Currency         string
	TotalAmount      int64
	CredentialsTitle string
	Users            []IUser
}

func (*PaymentsPaymentReceipt) CRC() uint32 {
	return 0x70c4fe03
}
func (*PaymentsPaymentReceipt) _IPaymentsPaymentReceipt() {}

type IPaymentsPaymentResult interface {
	tl.Object
	_IPaymentsPaymentResult()
}

var (
	_ IPaymentsPaymentResult = (*PaymentsPaymentResult)(nil)
	_ IPaymentsPaymentResult = (*PaymentsPaymentVerificationNeeded)(nil)
)

type PaymentsPaymentResult struct {
	Updates IUpdates
}

func (*PaymentsPaymentResult) CRC() uint32 {
	return 0x4e5f810d
}
func (*PaymentsPaymentResult) _IPaymentsPaymentResult() {}

type PaymentsPaymentVerificationNeeded struct {
	URL string
}

func (*PaymentsPaymentVerificationNeeded) CRC() uint32 {
	return 0xd8411139
}
func (*PaymentsPaymentVerificationNeeded) _IPaymentsPaymentResult() {}

type IPaymentsSavedInfo interface {
	tl.Object
	_IPaymentsSavedInfo()
}

var (
	_ IPaymentsSavedInfo = (*PaymentsSavedInfo)(nil)
)

type PaymentsSavedInfo struct {
	_                   struct{}              `tl:"flags,bitflag"`
	HasSavedCredentials bool                  `tl:",omitempty:flags:1,implicit"`
	SavedInfo           IPaymentRequestedInfo `tl:",omitempty:flags:0"`
}

func (*PaymentsSavedInfo) CRC() uint32 {
	return 0xfb8fe43c
}
func (*PaymentsSavedInfo) _IPaymentsSavedInfo() {}

type IPaymentsValidatedRequestedInfo interface {
	tl.Object
	_IPaymentsValidatedRequestedInfo()
}

var (
	_ IPaymentsValidatedRequestedInfo = (*PaymentsValidatedRequestedInfo)(nil)
)

type PaymentsValidatedRequestedInfo struct {
	_               struct{}          `tl:"flags,bitflag"`
	ID              *string           `tl:",omitempty:flags:0"`
	ShippingOptions []IShippingOption `tl:",omitempty:flags:1"`
}

func (*PaymentsValidatedRequestedInfo) CRC() uint32 {
	return 0xd1451883
}
func (*PaymentsValidatedRequestedInfo) _IPaymentsValidatedRequestedInfo() {}

type IPhoneExportedGroupCallInvite interface {
	tl.Object
	_IPhoneExportedGroupCallInvite()
}

var (
	_ IPhoneExportedGroupCallInvite = (*PhoneExportedGroupCallInvite)(nil)
)

type PhoneExportedGroupCallInvite struct {
	Link string
}

func (*PhoneExportedGroupCallInvite) CRC() uint32 {
	return 0x204bd158
}
func (*PhoneExportedGroupCallInvite) _IPhoneExportedGroupCallInvite() {}

type IPhoneGroupCall interface {
	tl.Object
	_IPhoneGroupCall()
}

var (
	_ IPhoneGroupCall = (*PhoneGroupCall)(nil)
)

type PhoneGroupCall struct {
	Call                   IGroupCall
	Participants           []IGroupCallParticipant
	ParticipantsNextOffset string
	Chats                  []IChat
	Users                  []IUser
}

func (*PhoneGroupCall) CRC() uint32 {
	return 0x9e727aad
}
func (*PhoneGroupCall) _IPhoneGroupCall() {}

type IPhoneGroupCallStreamChannels interface {
	tl.Object
	_IPhoneGroupCallStreamChannels()
}

var (
	_ IPhoneGroupCallStreamChannels = (*PhoneGroupCallStreamChannels)(nil)
)

type PhoneGroupCallStreamChannels struct {
	Channels []IGroupCallStreamChannel
}

func (*PhoneGroupCallStreamChannels) CRC() uint32 {
	return 0xd0e482b2
}
func (*PhoneGroupCallStreamChannels) _IPhoneGroupCallStreamChannels() {}

type IPhoneGroupCallStreamRtmpURL interface {
	tl.Object
	_IPhoneGroupCallStreamRtmpURL()
}

var (
	_ IPhoneGroupCallStreamRtmpURL = (*PhoneGroupCallStreamRtmpURL)(nil)
)

type PhoneGroupCallStreamRtmpURL struct {
	URL string
	Key string
}

func (*PhoneGroupCallStreamRtmpURL) CRC() uint32 {
	return 0x2dbf3432
}
func (*PhoneGroupCallStreamRtmpURL) _IPhoneGroupCallStreamRtmpURL() {}

type IPhoneGroupParticipants interface {
	tl.Object
	_IPhoneGroupParticipants()
}

var (
	_ IPhoneGroupParticipants = (*PhoneGroupParticipants)(nil)
)

type PhoneGroupParticipants struct {
	Count        int32
	Participants []IGroupCallParticipant
	NextOffset   string
	Chats        []IChat
	Users        []IUser
	Version      int32
}

func (*PhoneGroupParticipants) CRC() uint32 {
	return 0xf47751b6
}
func (*PhoneGroupParticipants) _IPhoneGroupParticipants() {}

type IPhoneJoinAsPeers interface {
	tl.Object
	_IPhoneJoinAsPeers()
}

var (
	_ IPhoneJoinAsPeers = (*PhoneJoinAsPeers)(nil)
)

type PhoneJoinAsPeers struct {
	Peers []IPeer
	Chats []IChat
	Users []IUser
}

func (*PhoneJoinAsPeers) CRC() uint32 {
	return 0xafe5623f
}
func (*PhoneJoinAsPeers) _IPhoneJoinAsPeers() {}

type IPhonePhoneCall interface {
	tl.Object
	_IPhonePhoneCall()
}

var (
	_ IPhonePhoneCall = (*PhonePhoneCall)(nil)
)

type PhonePhoneCall struct {
	PhoneCall IPhoneCall
	Users     []IUser
}

func (*PhonePhoneCall) CRC() uint32 {
	return 0xec82e140
}
func (*PhonePhoneCall) _IPhonePhoneCall() {}

type IPhotosPhoto interface {
	tl.Object
	_IPhotosPhoto()
}

var (
	_ IPhotosPhoto = (*PhotosPhoto)(nil)
)

type PhotosPhoto struct {
	Photo IPhoto
	Users []IUser
}

func (*PhotosPhoto) CRC() uint32 {
	return 0x20212ca8
}
func (*PhotosPhoto) _IPhotosPhoto() {}

type IPhotosPhotos interface {
	tl.Object
	_IPhotosPhotos()
}

var (
	_ IPhotosPhotos = (*PhotosPhotos)(nil)
	_ IPhotosPhotos = (*PhotosPhotosSlice)(nil)
)

type PhotosPhotos struct {
	Photos []IPhoto
	Users  []IUser
}

func (*PhotosPhotos) CRC() uint32 {
	return 0x8dca6aa5
}
func (*PhotosPhotos) _IPhotosPhotos() {}

type PhotosPhotosSlice struct {
	Count  int32
	Photos []IPhoto
	Users  []IUser
}

func (*PhotosPhotosSlice) CRC() uint32 {
	return 0x15051f54
}
func (*PhotosPhotosSlice) _IPhotosPhotos() {}

type IPremiumBoostsList interface {
	tl.Object
	_IPremiumBoostsList()
}

var (
	_ IPremiumBoostsList = (*PremiumBoostsList)(nil)
)

type PremiumBoostsList struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Boosts     []IBoost
	NextOffset *string `tl:",omitempty:flags:0"`
	Users      []IUser
}

func (*PremiumBoostsList) CRC() uint32 {
	return 0x86f8613c
}
func (*PremiumBoostsList) _IPremiumBoostsList() {}

type IPremiumBoostsStatus interface {
	tl.Object
	_IPremiumBoostsStatus()
}

var (
	_ IPremiumBoostsStatus = (*PremiumBoostsStatus)(nil)
)

type PremiumBoostsStatus struct {
	_                  struct{} `tl:"flags,bitflag"`
	MyBoost            bool     `tl:",omitempty:flags:2,implicit"`
	Level              int32
	CurrentLevelBoosts int32
	Boosts             int32
	GiftBoosts         *int32             `tl:",omitempty:flags:4"`
	NextLevelBoosts    *int32             `tl:",omitempty:flags:0"`
	PremiumAudience    IStatsPercentValue `tl:",omitempty:flags:1"`
	BoostURL           string
	PrepaidGiveaways   []IPrepaidGiveaway `tl:",omitempty:flags:3"`
	MyBoostSlots       []int32            `tl:",omitempty:flags:2"`
}

func (*PremiumBoostsStatus) CRC() uint32 {
	return 0x4959427a
}
func (*PremiumBoostsStatus) _IPremiumBoostsStatus() {}

type IPremiumMyBoosts interface {
	tl.Object
	_IPremiumMyBoosts()
}

var (
	_ IPremiumMyBoosts = (*PremiumMyBoosts)(nil)
)

type PremiumMyBoosts struct {
	MyBoosts []IMyBoost
	Chats    []IChat
	Users    []IUser
}

func (*PremiumMyBoosts) CRC() uint32 {
	return 0x9ae228e2
}
func (*PremiumMyBoosts) _IPremiumMyBoosts() {}

// Channel statistics
type IStatsBroadcastStats interface {
	tl.Object
	_IStatsBroadcastStats()
}

var (
	_ IStatsBroadcastStats = (*StatsBroadcastStats)(nil)
)

// [Channel statistics](https://core.telegram.org/api/stats)
type StatsBroadcastStats struct {
	Period                       IStatsDateRangeDays   // Period in consideration
	Followers                    IStatsAbsValueAndPrev // Follower count change for period in consideration
	ViewsPerPost                 IStatsAbsValueAndPrev // `total_viewcount/postcount`, for posts posted during the period in consideration (`views_per_post`). Note that in this case, `current` refers to the `period` in consideration (`min_date` till `max_date`), and `prev` refers to the previous period (`(min_date - (max_date - min_date))` till `min_date`)
	SharesPerPost                IStatsAbsValueAndPrev // `total_viewcount/postcount`, for posts posted during the period in consideration (`views_per_post`). Note that in this case, `current` refers to the `period` in consideration (`min_date` till `max_date`), and `prev` refers to the previous period (`(min_date - (max_date - min_date))` till `min_date`)
	ReactionsPerPost             IStatsAbsValueAndPrev
	ViewsPerStory                IStatsAbsValueAndPrev
	SharesPerStory               IStatsAbsValueAndPrev
	ReactionsPerStory            IStatsAbsValueAndPrev
	EnabledNotifications         IStatsPercentValue // Percentage of subscribers with enabled notifications
	GrowthGraph                  IStatsGraph        // Channel growth graph (absolute subscriber count)
	FollowersGraph               IStatsGraph        // Followers growth graph (relative subscriber count)
	MuteGraph                    IStatsGraph        // Muted users graph (relative)
	TopHoursGraph                IStatsGraph        // Views per hour graph (absolute)
	InteractionsGraph            IStatsGraph        // Interactions graph (absolute)
	IvInteractionsGraph          IStatsGraph        // IV interactions graph (absolute)
	ViewsBySourceGraph           IStatsGraph        // Views by source graph (absolute)
	NewFollowersBySourceGraph    IStatsGraph        // New followers by source graph (absolute)
	LanguagesGraph               IStatsGraph        // Subscriber language graph (piechart)
	ReactionsByEmotionGraph      IStatsGraph
	StoryInteractionsGraph       IStatsGraph
	StoryReactionsByEmotionGraph IStatsGraph
	RecentPostsInteractions      []IPostInteractionCounters // Recent post interactions
}

func (*StatsBroadcastStats) CRC() uint32 {
	return 0x396ca5fc
}
func (*StatsBroadcastStats) _IStatsBroadcastStats() {}

type IStatsMegagroupStats interface {
	tl.Object
	_IStatsMegagroupStats()
}

var (
	_ IStatsMegagroupStats = (*StatsMegagroupStats)(nil)
)

type StatsMegagroupStats struct {
	Period                  IStatsDateRangeDays
	Members                 IStatsAbsValueAndPrev
	Messages                IStatsAbsValueAndPrev
	Viewers                 IStatsAbsValueAndPrev
	Posters                 IStatsAbsValueAndPrev
	GrowthGraph             IStatsGraph
	MembersGraph            IStatsGraph
	NewMembersBySourceGraph IStatsGraph
	LanguagesGraph          IStatsGraph
	MessagesGraph           IStatsGraph
	ActionsGraph            IStatsGraph
	TopHoursGraph           IStatsGraph
	WeekdaysGraph           IStatsGraph
	TopPosters              []IStatsGroupTopPoster
	TopAdmins               []IStatsGroupTopAdmin
	TopInviters             []IStatsGroupTopInviter
	Users                   []IUser
}

func (*StatsMegagroupStats) CRC() uint32 {
	return 0xef7ff916
}
func (*StatsMegagroupStats) _IStatsMegagroupStats() {}

type IStatsMessageStats interface {
	tl.Object
	_IStatsMessageStats()
}

var (
	_ IStatsMessageStats = (*StatsMessageStats)(nil)
)

type StatsMessageStats struct {
	ViewsGraph              IStatsGraph
	ReactionsByEmotionGraph IStatsGraph
}

func (*StatsMessageStats) CRC() uint32 {
	return 0x7fe91c14
}
func (*StatsMessageStats) _IStatsMessageStats() {}

type IStatsPublicForwards interface {
	tl.Object
	_IStatsPublicForwards()
}

var (
	_ IStatsPublicForwards = (*StatsPublicForwards)(nil)
)

type StatsPublicForwards struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Forwards   []IPublicForward
	NextOffset *string `tl:",omitempty:flags:0"`
	Chats      []IChat
	Users      []IUser
}

func (*StatsPublicForwards) CRC() uint32 {
	return 0x93037e20
}
func (*StatsPublicForwards) _IStatsPublicForwards() {}

type IStatsStoryStats interface {
	tl.Object
	_IStatsStoryStats()
}

var (
	_ IStatsStoryStats = (*StatsStoryStats)(nil)
)

type StatsStoryStats struct {
	ViewsGraph              IStatsGraph
	ReactionsByEmotionGraph IStatsGraph
}

func (*StatsStoryStats) CRC() uint32 {
	return 0x50cd067c
}
func (*StatsStoryStats) _IStatsStoryStats() {}

type IStickersSuggestedShortName interface {
	tl.Object
	_IStickersSuggestedShortName()
}

var (
	_ IStickersSuggestedShortName = (*StickersSuggestedShortName)(nil)
)

type StickersSuggestedShortName struct {
	ShortName string
}

func (*StickersSuggestedShortName) CRC() uint32 {
	return 0x85fea03f
}
func (*StickersSuggestedShortName) _IStickersSuggestedShortName() {}

type IStoriesAllStories interface {
	tl.Object
	_IStoriesAllStories()
}

var (
	_ IStoriesAllStories = (*StoriesAllStoriesNotModified)(nil)
	_ IStoriesAllStories = (*StoriesAllStories)(nil)
)

type StoriesAllStoriesNotModified struct {
	_           struct{} `tl:"flags,bitflag"`
	State       string
	StealthMode IStoriesStealthMode
}

func (*StoriesAllStoriesNotModified) CRC() uint32 {
	return 0x1158fe3e
}
func (*StoriesAllStoriesNotModified) _IStoriesAllStories() {}

type StoriesAllStories struct {
	_           struct{} `tl:"flags,bitflag"`
	HasMore     bool     `tl:",omitempty:flags:0,implicit"`
	Count       int32
	State       string
	PeerStories []IPeerStories
	Chats       []IChat
	Users       []IUser
	StealthMode IStoriesStealthMode
}

func (*StoriesAllStories) CRC() uint32 {
	return 0x6efc5e81
}
func (*StoriesAllStories) _IStoriesAllStories() {}

type IStoriesPeerStories interface {
	tl.Object
	_IStoriesPeerStories()
}

var (
	_ IStoriesPeerStories = (*StoriesPeerStories)(nil)
)

type StoriesPeerStories struct {
	Stories IPeerStories
	Chats   []IChat
	Users   []IUser
}

func (*StoriesPeerStories) CRC() uint32 {
	return 0xcae68768
}
func (*StoriesPeerStories) _IStoriesPeerStories() {}

type IStoriesStories interface {
	tl.Object
	_IStoriesStories()
}

var (
	_ IStoriesStories = (*StoriesStories)(nil)
)

type StoriesStories struct {
	Count   int32
	Stories []IStoryItem
	Chats   []IChat
	Users   []IUser
}

func (*StoriesStories) CRC() uint32 {
	return 0x5dd8c3c8
}
func (*StoriesStories) _IStoriesStories() {}

type IStoriesStoryReactionsList interface {
	tl.Object
	_IStoriesStoryReactionsList()
}

var (
	_ IStoriesStoryReactionsList = (*StoriesStoryReactionsList)(nil)
)

type StoriesStoryReactionsList struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Reactions  []IStoryReaction
	Chats      []IChat
	Users      []IUser
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*StoriesStoryReactionsList) CRC() uint32 {
	return 0xaa5f789c
}
func (*StoriesStoryReactionsList) _IStoriesStoryReactionsList() {}

type IStoriesStoryViews interface {
	tl.Object
	_IStoriesStoryViews()
}

var (
	_ IStoriesStoryViews = (*StoriesStoryViews)(nil)
)

type StoriesStoryViews struct {
	Views []IStoryViews
	Users []IUser
}

func (*StoriesStoryViews) CRC() uint32 {
	return 0xde9eed1d
}
func (*StoriesStoryViews) _IStoriesStoryViews() {}

type IStoriesStoryViewsList interface {
	tl.Object
	_IStoriesStoryViewsList()
}

var (
	_ IStoriesStoryViewsList = (*StoriesStoryViewsList)(nil)
)

type StoriesStoryViewsList struct {
	_              struct{} `tl:"flags,bitflag"`
	Count          int32
	ViewsCount     int32
	ForwardsCount  int32
	ReactionsCount int32
	Views          []IStoryView
	Chats          []IChat
	Users          []IUser
	NextOffset     *string `tl:",omitempty:flags:0"`
}

func (*StoriesStoryViewsList) CRC() uint32 {
	return 0x59d78fc5
}
func (*StoriesStoryViewsList) _IStoriesStoryViewsList() {}

type IUpdatesChannelDifference interface {
	tl.Object
	_IUpdatesChannelDifference()
}

var (
	_ IUpdatesChannelDifference = (*UpdatesChannelDifferenceEmpty)(nil)
	_ IUpdatesChannelDifference = (*UpdatesChannelDifferenceTooLong)(nil)
	_ IUpdatesChannelDifference = (*UpdatesChannelDifference)(nil)
)

type UpdatesChannelDifferenceEmpty struct {
	_       struct{} `tl:"flags,bitflag"`
	Final   bool     `tl:",omitempty:flags:0,implicit"`
	Pts     int32
	Timeout *int32 `tl:",omitempty:flags:1"`
}

func (*UpdatesChannelDifferenceEmpty) CRC() uint32 {
	return 0x3e11affb
}
func (*UpdatesChannelDifferenceEmpty) _IUpdatesChannelDifference() {}

type UpdatesChannelDifferenceTooLong struct {
	_        struct{} `tl:"flags,bitflag"`
	Final    bool     `tl:",omitempty:flags:0,implicit"`
	Timeout  *int32   `tl:",omitempty:flags:1"`
	Dialog   IDialog
	Messages []IMessage
	Chats    []IChat
	Users    []IUser
}

func (*UpdatesChannelDifferenceTooLong) CRC() uint32 {
	return 0xa4bcc6fe
}
func (*UpdatesChannelDifferenceTooLong) _IUpdatesChannelDifference() {}

type UpdatesChannelDifference struct {
	_            struct{} `tl:"flags,bitflag"`
	Final        bool     `tl:",omitempty:flags:0,implicit"`
	Pts          int32
	Timeout      *int32 `tl:",omitempty:flags:1"`
	NewMessages  []IMessage
	OtherUpdates []IUpdate
	Chats        []IChat
	Users        []IUser
}

func (*UpdatesChannelDifference) CRC() uint32 {
	return 0x2064674e
}
func (*UpdatesChannelDifference) _IUpdatesChannelDifference() {}

type IUpdatesDifference interface {
	tl.Object
	_IUpdatesDifference()
}

var (
	_ IUpdatesDifference = (*UpdatesDifferenceEmpty)(nil)
	_ IUpdatesDifference = (*UpdatesDifference)(nil)
	_ IUpdatesDifference = (*UpdatesDifferenceSlice)(nil)
	_ IUpdatesDifference = (*UpdatesDifferenceTooLong)(nil)
)

type UpdatesDifferenceEmpty struct {
	Date int32
	Seq  int32
}

func (*UpdatesDifferenceEmpty) CRC() uint32 {
	return 0x5d75a138
}
func (*UpdatesDifferenceEmpty) _IUpdatesDifference() {}

type UpdatesDifference struct {
	NewMessages          []IMessage
	NewEncryptedMessages []IEncryptedMessage
	OtherUpdates         []IUpdate
	Chats                []IChat
	Users                []IUser
	State                IUpdatesState
}

func (*UpdatesDifference) CRC() uint32 {
	return 0xf49ca0
}
func (*UpdatesDifference) _IUpdatesDifference() {}

type UpdatesDifferenceSlice struct {
	NewMessages          []IMessage
	NewEncryptedMessages []IEncryptedMessage
	OtherUpdates         []IUpdate
	Chats                []IChat
	Users                []IUser
	IntermediateState    IUpdatesState
}

func (*UpdatesDifferenceSlice) CRC() uint32 {
	return 0xa8fb1981
}
func (*UpdatesDifferenceSlice) _IUpdatesDifference() {}

type UpdatesDifferenceTooLong struct {
	Pts int32
}

func (*UpdatesDifferenceTooLong) CRC() uint32 {
	return 0x4afe8f6d
}
func (*UpdatesDifferenceTooLong) _IUpdatesDifference() {}

type IUpdatesState interface {
	tl.Object
	_IUpdatesState()
}

var (
	_ IUpdatesState = (*UpdatesState)(nil)
)

type UpdatesState struct {
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

func (*UpdatesState) CRC() uint32 {
	return 0xa56c2a3e
}
func (*UpdatesState) _IUpdatesState() {}

type IUploadCdnFile interface {
	tl.Object
	_IUploadCdnFile()
}

var (
	_ IUploadCdnFile = (*UploadCdnFileReuploadNeeded)(nil)
	_ IUploadCdnFile = (*UploadCdnFile)(nil)
)

type UploadCdnFileReuploadNeeded struct {
	RequestToken []byte
}

func (*UploadCdnFileReuploadNeeded) CRC() uint32 {
	return 0xeea8e46e
}
func (*UploadCdnFileReuploadNeeded) _IUploadCdnFile() {}

type UploadCdnFile struct {
	Bytes []byte
}

func (*UploadCdnFile) CRC() uint32 {
	return 0xa99fca4f
}
func (*UploadCdnFile) _IUploadCdnFile() {}

type IUploadFile interface {
	tl.Object
	_IUploadFile()
}

var (
	_ IUploadFile = (*UploadFile)(nil)
	_ IUploadFile = (*UploadFileCdnRedirect)(nil)
)

type UploadFile struct {
	Type  EnumStorageFileType
	Mtime int32
	Bytes []byte
}

func (*UploadFile) CRC() uint32 {
	return 0x96a18d5
}
func (*UploadFile) _IUploadFile() {}

type UploadFileCdnRedirect struct {
	DcID          int32
	FileToken     []byte
	EncryptionKey []byte
	EncryptionIv  []byte
	FileHashes    []IFileHash
}

func (*UploadFileCdnRedirect) CRC() uint32 {
	return 0xf18cda44
}
func (*UploadFileCdnRedirect) _IUploadFile() {}

type IUploadWebFile interface {
	tl.Object
	_IUploadWebFile()
}

var (
	_ IUploadWebFile = (*UploadWebFile)(nil)
)

type UploadWebFile struct {
	Size     int32
	MimeType string
	FileType EnumStorageFileType
	Mtime    int32
	Bytes    []byte
}

func (*UploadWebFile) CRC() uint32 {
	return 0x21e753bc
}
func (*UploadWebFile) _IUploadWebFile() {}

type IUsersUserFull interface {
	tl.Object
	_IUsersUserFull()
}

var (
	_ IUsersUserFull = (*UsersUserFull)(nil)
)

type UsersUserFull struct {
	FullUser IUserFull
	Chats    []IChat
	Users    []IUser
}

func (*UsersUserFull) CRC() uint32 {
	return 0x3b6d152e
}
func (*UsersUserFull) _IUsersUserFull() {}

type EnumAttachMenuPeerType uint32

const (
	AttachMenuPeerTypeSameBotPm EnumAttachMenuPeerType = 0x7d6be90e
	AttachMenuPeerTypeBotPm     EnumAttachMenuPeerType = 0xc32bfa1a
	AttachMenuPeerTypePm        EnumAttachMenuPeerType = 0xf146d31f
	AttachMenuPeerTypeChat      EnumAttachMenuPeerType = 0x509113f
	AttachMenuPeerTypeBroadcast EnumAttachMenuPeerType = 0x7bfbdefc
)

type EnumBaseTheme uint32

const (
	BaseThemeClassic EnumBaseTheme = 0xc3a12462 // Classic theme
	BaseThemeDay     EnumBaseTheme = 0xfbd81688 // Day theme
	BaseThemeNight   EnumBaseTheme = 0xb7b31ea8 // Night theme
	BaseThemeTinted  EnumBaseTheme = 0x6d5f77ee // Tinted theme
	BaseThemeArctic  EnumBaseTheme = 0x5b11125a // Arctic theme
)

type EnumInlineQueryPeerType uint32

const (
	InlineQueryPeerTypeSameBotPm EnumInlineQueryPeerType = 0x3081ed9d
	InlineQueryPeerTypePm        EnumInlineQueryPeerType = 0x833c0fac
	InlineQueryPeerTypeChat      EnumInlineQueryPeerType = 0xd766c50a
	InlineQueryPeerTypeMegagroup EnumInlineQueryPeerType = 0x5ec4be43
	InlineQueryPeerTypeBroadcast EnumInlineQueryPeerType = 0x6334ee9a
	InlineQueryPeerTypeBotPm     EnumInlineQueryPeerType = 0xe3b2d0c
)

type EnumInputPrivacyKey uint32

const (
	InputPrivacyKeyStatusTimestamp EnumInputPrivacyKey = 0x4f96cb18 // Whether we can see the exact last online timestamp of the user
	InputPrivacyKeyChatInvite      EnumInputPrivacyKey = 0xbdfb0426 // Whether the user can be invited to chats
	InputPrivacyKeyPhoneCall       EnumInputPrivacyKey = 0xfabadc5f // Whether the user will accept phone calls
	InputPrivacyKeyPhoneP2P        EnumInputPrivacyKey = 0xdb9e70d2 // Whether the user allows P2P communication during VoIP calls
	InputPrivacyKeyForwards        EnumInputPrivacyKey = 0xa4dd4c08 // Whether messages forwarded from this user will be [anonymous](https://telegram.org/blog/unsend-privacy-emoji#anonymous-forwarding)
	InputPrivacyKeyProfilePhoto    EnumInputPrivacyKey = 0x5719bacc // Whether people will be able to see the user's profile picture
	InputPrivacyKeyPhoneNumber     EnumInputPrivacyKey = 0x352dafa  // Whether people will be able to see the user's phone number
	InputPrivacyKeyAddedByPhone    EnumInputPrivacyKey = 0xd1219bdd // Whether people can add you to their contact list by your phone number
	InputPrivacyKeyVoiceMessages   EnumInputPrivacyKey = 0xaee69d68
	InputPrivacyKeyAbout           EnumInputPrivacyKey = 0x3823cc40
)

type EnumPhoneCallDiscardReason uint32

const (
	PhoneCallDiscardReasonMissed     EnumPhoneCallDiscardReason = 0x85e42301 // The phone call was missed
	PhoneCallDiscardReasonDisconnect EnumPhoneCallDiscardReason = 0xe095c1a0 // The phone call was disconnected
	PhoneCallDiscardReasonHangup     EnumPhoneCallDiscardReason = 0x57adc690 // The phone call was ended normally
	PhoneCallDiscardReasonBusy       EnumPhoneCallDiscardReason = 0xfaf7e8c9 // The phone call was discared because the user is busy in another call
)

type EnumPrivacyKey uint32

const (
	PrivacyKeyStatusTimestamp EnumPrivacyKey = 0xbc2eab30 // Whether we can see the last online timestamp
	PrivacyKeyChatInvite      EnumPrivacyKey = 0x500e6dfa // Whether the user can be invited to chats
	PrivacyKeyPhoneCall       EnumPrivacyKey = 0x3d662b7b // Whether the user accepts phone calls
	PrivacyKeyPhoneP2P        EnumPrivacyKey = 0x39491cc8 // Whether P2P connections in phone calls are allowed
	PrivacyKeyForwards        EnumPrivacyKey = 0x69ec56a3 // Whether messages forwarded from the user will be anonymously forwarded
	PrivacyKeyProfilePhoto    EnumPrivacyKey = 0x96151fed // Whether the profile picture of the user is visible
	PrivacyKeyPhoneNumber     EnumPrivacyKey = 0xd19ae46d // Whether the user allows us to see his phone number
	PrivacyKeyAddedByPhone    EnumPrivacyKey = 0x42ffd42b // Whether people can add you to their contact list by your phone number
	PrivacyKeyVoiceMessages   EnumPrivacyKey = 0x697f414
	PrivacyKeyAbout           EnumPrivacyKey = 0xa486b761
)

type EnumReportReason uint32

const (
	InputReportReasonSpam            EnumReportReason = 0x58dbcab8
	InputReportReasonViolence        EnumReportReason = 0x1e22c78d
	InputReportReasonPornography     EnumReportReason = 0x2e59d922
	InputReportReasonChildAbuse      EnumReportReason = 0xadf44ee3
	InputReportReasonOther           EnumReportReason = 0xc1e4a2b1
	InputReportReasonCopyright       EnumReportReason = 0x9b89f93a
	InputReportReasonGeoIrrelevant   EnumReportReason = 0xdbd4feed
	InputReportReasonFake            EnumReportReason = 0xf5ddd6e7
	InputReportReasonIllegalDrugs    EnumReportReason = 0xa8eb2be
	InputReportReasonPersonalDetails EnumReportReason = 0x9ec7863d
)

type EnumSecureValueType uint32

const (
	SecureValueTypePersonalDetails       EnumSecureValueType = 0x9d2a81e3 // Personal details
	SecureValueTypePassport              EnumSecureValueType = 0x3dac6a00 // Passport
	SecureValueTypeDriverLicense         EnumSecureValueType = 0x6e425c4  // Driver's license
	SecureValueTypeIdentityCard          EnumSecureValueType = 0xa0d0744b // Identity card
	SecureValueTypeInternalPassport      EnumSecureValueType = 0x99a48f23 // Internal [passport](https://core.telegram.org/passport)
	SecureValueTypeAddress               EnumSecureValueType = 0xcbe31e26 // Address
	SecureValueTypeUtilityBill           EnumSecureValueType = 0xfc36954e // Utility bill
	SecureValueTypeBankStatement         EnumSecureValueType = 0x89137c0d // Bank statement
	SecureValueTypeRentalAgreement       EnumSecureValueType = 0x8b883488 // Rental agreement
	SecureValueTypePassportRegistration  EnumSecureValueType = 0x99e3806a // Internal registration [passport](https://core.telegram.org/passport)
	SecureValueTypeTemporaryRegistration EnumSecureValueType = 0xea02ec33 // Temporary registration
	SecureValueTypePhone                 EnumSecureValueType = 0xb320aadb // Phone
	SecureValueTypeEmail                 EnumSecureValueType = 0x8e3ca7ee // Email
)

type EnumTopPeerCategory uint32

const (
	TopPeerCategoryBotsPm         EnumTopPeerCategory = 0xab661b5b // Most used bots
	TopPeerCategoryBotsInline     EnumTopPeerCategory = 0x148677e2 // Most used inline bots
	TopPeerCategoryCorrespondents EnumTopPeerCategory = 0x637b7ed  // Users we've chatted most frequently with
	TopPeerCategoryGroups         EnumTopPeerCategory = 0xbd17a14a // Often-opened groups and supergroups
	TopPeerCategoryChannels       EnumTopPeerCategory = 0x161d9628 // Most frequently visited channels
	TopPeerCategoryPhoneCalls     EnumTopPeerCategory = 0x1e76a78c // Most frequently called users
	TopPeerCategoryForwardUsers   EnumTopPeerCategory = 0xa8406ca9 // Users to which the users often forwards messages to
	TopPeerCategoryForwardChats   EnumTopPeerCategory = 0xfbeec0f0 // Chats to which the users often forwards messages to
)

type EnumAuthCodeType uint32

const (
	AuthCodeTypeSms         EnumAuthCodeType = 0x72a3158c // Type of verification code that will be sent next if you call the resendCode method: SMS code
	AuthCodeTypeCall        EnumAuthCodeType = 0x741cd3e3 // Type of verification code that will be sent next if you call the resendCode method: Phone call
	AuthCodeTypeFlashCall   EnumAuthCodeType = 0x226ccefb // Type of verification code that will be sent next if you call the resendCode method: Flash call
	AuthCodeTypeMissedCall  EnumAuthCodeType = 0xd61ad6ee
	AuthCodeTypeFragmentSms EnumAuthCodeType = 0x6ed998c
)

type EnumStorageFileType uint32

const (
	StorageFileUnknown EnumStorageFileType = 0xaa963b05 // Unknown type.
	StorageFilePartial EnumStorageFileType = 0x40bc6f52 // Part of a bigger file.
	StorageFileJpeg    EnumStorageFileType = 0x7efe0e   // JPEG image. MIME type: image/jpeg.
	StorageFileGif     EnumStorageFileType = 0xcae1aadf // GIF image. MIME type: image/gif.
	StorageFilePng     EnumStorageFileType = 0xa4f63c0  // PNG image. MIME type: image/png.
	StorageFilePdf     EnumStorageFileType = 0xae1e508d // PDF document image. MIME type: application/pdf.
	StorageFileMp3     EnumStorageFileType = 0x528a0677 // Mp3 audio. MIME type: audio/mpeg.
	StorageFileMov     EnumStorageFileType = 0x4b09ebbc // Quicktime video. MIME type: video/quicktime.
	StorageFileMp4     EnumStorageFileType = 0xb3cea0e4 // MPEG-4 video. MIME type: video/mp4.
	StorageFileWebp    EnumStorageFileType = 0x1081464c // WEBP image. MIME type: image/webp.
)

type Requester interface {
	MakeRequest(ctx context.Context, msg []byte) ([]byte, error)
}

func request[IN, OUT any](ctx context.Context, m Requester, in *IN, out *OUT) error {
	if msg, err := tl.Marshal(in); err != nil {
		return fmt.Errorf("marshaling: %w", err)
	} else if respRaw, err := m.MakeRequest(ctx, msg); err != nil {
		return fmt.Errorf("sending: %w", err)
	} else if err := Unmarshal(respRaw, out); err != nil {
		return fmt.Errorf("got invalid response type: %w", err)
	}

	return nil
}

type AccountRegisterDeviceRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	NoMuted    bool     `tl:",omitempty:flags:0,implicit"`
	TokenType  int32
	Token      string
	AppSandbox bool
	Secret     []byte
	OtherUids  []int64
}

func (*AccountRegisterDeviceRequest) CRC() uint32 {
	return 0xec86017a
}

func AccountRegisterDevice(ctx context.Context, m Requester, i AccountRegisterDeviceRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUnregisterDeviceRequest struct {
	TokenType int32
	Token     string
	OtherUids []int64
}

func (*AccountUnregisterDeviceRequest) CRC() uint32 {
	return 0x6a0d3206
}

func AccountUnregisterDevice(ctx context.Context, m Requester, i AccountUnregisterDeviceRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateNotifySettingsRequest struct {
	Peer     IInputNotifyPeer
	Settings IInputPeerNotifySettings
}

func (*AccountUpdateNotifySettingsRequest) CRC() uint32 {
	return 0x84be5b93
}

func AccountUpdateNotifySettings(ctx context.Context, m Requester, i AccountUpdateNotifySettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetNotifySettingsRequest struct {
	Peer IInputNotifyPeer
}

func (*AccountGetNotifySettingsRequest) CRC() uint32 {
	return 0x12b3ad31
}

func AccountGetNotifySettings(ctx context.Context, m Requester, i AccountGetNotifySettingsRequest) (IPeerNotifySettings, error) {
	var res IPeerNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AccountResetNotifySettingsRequest struct{}

func (*AccountResetNotifySettingsRequest) CRC() uint32 {
	return 0xdb7e1747
}

func AccountResetNotifySettings(ctx context.Context, m Requester, i AccountResetNotifySettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateProfileRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	FirstName *string  `tl:",omitempty:flags:0"`
	LastName  *string  `tl:",omitempty:flags:1"`
	About     *string  `tl:",omitempty:flags:2"`
}

func (*AccountUpdateProfileRequest) CRC() uint32 {
	return 0x78515775
}

func AccountUpdateProfile(ctx context.Context, m Requester, i AccountUpdateProfileRequest) (IUser, error) {
	var res IUser
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateStatusRequest struct {
	Offline bool
}

func (*AccountUpdateStatusRequest) CRC() uint32 {
	return 0x6628562c
}

func AccountUpdateStatus(ctx context.Context, m Requester, i AccountUpdateStatusRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPapersRequest struct {
	Hash int64
}

func (*AccountGetWallPapersRequest) CRC() uint32 {
	return 0x7967d36
}

func AccountGetWallPapers(ctx context.Context, m Requester, i AccountGetWallPapersRequest) (IAccountWallPapers, error) {
	var res IAccountWallPapers
	return res, request(ctx, m, &i, &res)
}

type AccountReportPeerRequest struct {
	Peer    IInputPeer
	Reason  EnumReportReason
	Message string
}

func (*AccountReportPeerRequest) CRC() uint32 {
	return 0xc5ba3d86
}

func AccountReportPeer(ctx context.Context, m Requester, i AccountReportPeerRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCheckUsernameRequest struct {
	Username string
}

func (*AccountCheckUsernameRequest) CRC() uint32 {
	return 0x2714d86c
}

func AccountCheckUsername(ctx context.Context, m Requester, i AccountCheckUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateUsernameRequest struct {
	Username string
}

func (*AccountUpdateUsernameRequest) CRC() uint32 {
	return 0x3e0bdd7c
}

func AccountUpdateUsername(ctx context.Context, m Requester, i AccountUpdateUsernameRequest) (IUser, error) {
	var res IUser
	return res, request(ctx, m, &i, &res)
}

type AccountGetPrivacyRequest struct {
	Key EnumInputPrivacyKey
}

func (*AccountGetPrivacyRequest) CRC() uint32 {
	return 0xdadbc950
}

func AccountGetPrivacy(ctx context.Context, m Requester, i AccountGetPrivacyRequest) (IAccountPrivacyRules, error) {
	var res IAccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountSetPrivacyRequest struct {
	Key   EnumInputPrivacyKey
	Rules []IInputPrivacyRule
}

func (*AccountSetPrivacyRequest) CRC() uint32 {
	return 0xc9f81ce8
}

func AccountSetPrivacy(ctx context.Context, m Requester, i AccountSetPrivacyRequest) (IAccountPrivacyRules, error) {
	var res IAccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteAccountRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Reason   string
	Password IInputCheckPasswordSRP `tl:",omitempty:flags:0"`
}

func (*AccountDeleteAccountRequest) CRC() uint32 {
	return 0xa2c0cf74
}

func AccountDeleteAccount(ctx context.Context, m Requester, i AccountDeleteAccountRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAccountTTLRequest struct{}

func (*AccountGetAccountTTLRequest) CRC() uint32 {
	return 0x8fc711d
}

func AccountGetAccountTTL(ctx context.Context, m Requester, i AccountGetAccountTTLRequest) (IAccountDaysTTL, error) {
	var res IAccountDaysTTL
	return res, request(ctx, m, &i, &res)
}

type AccountSetAccountTTLRequest struct {
	TTL IAccountDaysTTL
}

func (*AccountSetAccountTTLRequest) CRC() uint32 {
	return 0x2442485e
}

func AccountSetAccountTTL(ctx context.Context, m Requester, i AccountSetAccountTTLRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendChangePhoneCodeRequest struct {
	PhoneNumber string
	Settings    ICodeSettings
}

func (*AccountSendChangePhoneCodeRequest) CRC() uint32 {
	return 0x82574ae5
}

func AccountSendChangePhoneCode(ctx context.Context, m Requester, i AccountSendChangePhoneCodeRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountChangePhoneRequest struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountChangePhoneRequest) CRC() uint32 {
	return 0x70c32edb
}

func AccountChangePhone(ctx context.Context, m Requester, i AccountChangePhoneRequest) (IUser, error) {
	var res IUser
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateDeviceLockedRequest struct {
	Period int32
}

func (*AccountUpdateDeviceLockedRequest) CRC() uint32 {
	return 0x38df3532
}

func AccountUpdateDeviceLocked(ctx context.Context, m Requester, i AccountUpdateDeviceLockedRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAuthorizationsRequest struct{}

func (*AccountGetAuthorizationsRequest) CRC() uint32 {
	return 0xe320c158
}

func AccountGetAuthorizations(ctx context.Context, m Requester, i AccountGetAuthorizationsRequest) (IAccountAuthorizations, error) {
	var res IAccountAuthorizations
	return res, request(ctx, m, &i, &res)
}

type AccountResetAuthorizationRequest struct {
	Hash int64
}

func (*AccountResetAuthorizationRequest) CRC() uint32 {
	return 0xdf77f3bc
}

func AccountResetAuthorization(ctx context.Context, m Requester, i AccountResetAuthorizationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetPasswordRequest struct{}

func (*AccountGetPasswordRequest) CRC() uint32 {
	return 0x548a30f5
}

func AccountGetPassword(ctx context.Context, m Requester, i AccountGetPasswordRequest) (IAccountPassword, error) {
	var res IAccountPassword
	return res, request(ctx, m, &i, &res)
}

type AccountGetPasswordSettingsRequest struct {
	Password IInputCheckPasswordSRP
}

func (*AccountGetPasswordSettingsRequest) CRC() uint32 {
	return 0x9cd4eaf9
}

func AccountGetPasswordSettings(ctx context.Context, m Requester, i AccountGetPasswordSettingsRequest) (IAccountPasswordSettings, error) {
	var res IAccountPasswordSettings
	return res, request(ctx, m, &i, &res)
}

type AccountUpdatePasswordSettingsRequest struct {
	Password    IInputCheckPasswordSRP
	NewSettings IAccountPasswordInputSettings
}

func (*AccountUpdatePasswordSettingsRequest) CRC() uint32 {
	return 0xa59b102f
}

func AccountUpdatePasswordSettings(ctx context.Context, m Requester, i AccountUpdatePasswordSettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendConfirmPhoneCodeRequest struct {
	Hash     string
	Settings ICodeSettings
}

func (*AccountSendConfirmPhoneCodeRequest) CRC() uint32 {
	return 0x1b3faa88
}

func AccountSendConfirmPhoneCode(ctx context.Context, m Requester, i AccountSendConfirmPhoneCodeRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPhoneRequest struct {
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountConfirmPhoneRequest) CRC() uint32 {
	return 0x5f2178c3
}

func AccountConfirmPhone(ctx context.Context, m Requester, i AccountConfirmPhoneRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetTmpPasswordRequest struct {
	Password IInputCheckPasswordSRP
	Period   int32
}

func (*AccountGetTmpPasswordRequest) CRC() uint32 {
	return 0x449e0b51
}

func AccountGetTmpPassword(ctx context.Context, m Requester, i AccountGetTmpPasswordRequest) (IAccountTmpPassword, error) {
	var res IAccountTmpPassword
	return res, request(ctx, m, &i, &res)
}

type AccountGetWebAuthorizationsRequest struct{}

func (*AccountGetWebAuthorizationsRequest) CRC() uint32 {
	return 0x182e6d6f
}

func AccountGetWebAuthorizations(ctx context.Context, m Requester, i AccountGetWebAuthorizationsRequest) (IAccountWebAuthorizations, error) {
	var res IAccountWebAuthorizations
	return res, request(ctx, m, &i, &res)
}

type AccountResetWebAuthorizationRequest struct {
	Hash int64
}

func (*AccountResetWebAuthorizationRequest) CRC() uint32 {
	return 0x2d01b9ef
}

func AccountResetWebAuthorization(ctx context.Context, m Requester, i AccountResetWebAuthorizationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetWebAuthorizationsRequest struct{}

func (*AccountResetWebAuthorizationsRequest) CRC() uint32 {
	return 0x682d2594
}

func AccountResetWebAuthorizations(ctx context.Context, m Requester, i AccountResetWebAuthorizationsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAllSecureValuesRequest struct{}

func (*AccountGetAllSecureValuesRequest) CRC() uint32 {
	return 0xb288bc7d
}

func AccountGetAllSecureValues(ctx context.Context, m Requester, i AccountGetAllSecureValuesRequest) ([]ISecureValue, error) {
	var res []ISecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountGetSecureValueRequest struct {
	Types []EnumSecureValueType
}

func (*AccountGetSecureValueRequest) CRC() uint32 {
	return 0x73665bc2
}

func AccountGetSecureValue(ctx context.Context, m Requester, i AccountGetSecureValueRequest) ([]ISecureValue, error) {
	var res []ISecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountSaveSecureValueRequest struct {
	Value          IInputSecureValue
	SecureSecretID int64
}

func (*AccountSaveSecureValueRequest) CRC() uint32 {
	return 0x899fe31d
}

func AccountSaveSecureValue(ctx context.Context, m Requester, i AccountSaveSecureValueRequest) (ISecureValue, error) {
	var res ISecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteSecureValueRequest struct {
	Types []EnumSecureValueType
}

func (*AccountDeleteSecureValueRequest) CRC() uint32 {
	return 0xb880bc4b
}

func AccountDeleteSecureValue(ctx context.Context, m Requester, i AccountDeleteSecureValueRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAuthorizationFormRequest struct {
	BotID     int64
	Scope     string
	PublicKey string
}

func (*AccountGetAuthorizationFormRequest) CRC() uint32 {
	return 0xa929597a
}

func AccountGetAuthorizationForm(ctx context.Context, m Requester, i AccountGetAuthorizationFormRequest) (IAccountAuthorizationForm, error) {
	var res IAccountAuthorizationForm
	return res, request(ctx, m, &i, &res)
}

type AccountAcceptAuthorizationRequest struct {
	BotID       int64
	Scope       string
	PublicKey   string
	ValueHashes []ISecureValueHash
	Credentials ISecureCredentialsEncrypted
}

func (*AccountAcceptAuthorizationRequest) CRC() uint32 {
	return 0xf3ed4c73
}

func AccountAcceptAuthorization(ctx context.Context, m Requester, i AccountAcceptAuthorizationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyPhoneCodeRequest struct {
	PhoneNumber string
	Settings    ICodeSettings
}

func (*AccountSendVerifyPhoneCodeRequest) CRC() uint32 {
	return 0xa5a356f9
}

func AccountSendVerifyPhoneCode(ctx context.Context, m Requester, i AccountSendVerifyPhoneCodeRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyPhoneRequest struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountVerifyPhoneRequest) CRC() uint32 {
	return 0x4dd3a7f6
}

func AccountVerifyPhone(ctx context.Context, m Requester, i AccountVerifyPhoneRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyEmailCodeRequest struct {
	Purpose IEmailVerifyPurpose
	Email   string
}

func (*AccountSendVerifyEmailCodeRequest) CRC() uint32 {
	return 0x98e037bb
}

func AccountSendVerifyEmailCode(ctx context.Context, m Requester, i AccountSendVerifyEmailCodeRequest) (IAccountSentEmailCode, error) {
	var res IAccountSentEmailCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyEmailRequest struct {
	Purpose      IEmailVerifyPurpose
	Verification IEmailVerification
}

func (*AccountVerifyEmailRequest) CRC() uint32 {
	return 0x32da4cf
}

func AccountVerifyEmail(ctx context.Context, m Requester, i AccountVerifyEmailRequest) (IAccountEmailVerified, error) {
	var res IAccountEmailVerified
	return res, request(ctx, m, &i, &res)
}

type AccountInitTakeoutSessionRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	Contacts          bool     `tl:",omitempty:flags:0,implicit"`
	MessageUsers      bool     `tl:",omitempty:flags:1,implicit"`
	MessageChats      bool     `tl:",omitempty:flags:2,implicit"`
	MessageMegagroups bool     `tl:",omitempty:flags:3,implicit"`
	MessageChannels   bool     `tl:",omitempty:flags:4,implicit"`
	Files             bool     `tl:",omitempty:flags:5,implicit"`
	FileMaxSize       *int64   `tl:",omitempty:flags:5"`
}

func (*AccountInitTakeoutSessionRequest) CRC() uint32 {
	return 0x8ef3eab0
}

func AccountInitTakeoutSession(ctx context.Context, m Requester, i AccountInitTakeoutSessionRequest) (IAccountTakeout, error) {
	var res IAccountTakeout
	return res, request(ctx, m, &i, &res)
}

type AccountFinishTakeoutSessionRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Success bool     `tl:",omitempty:flags:0,implicit"`
}

func (*AccountFinishTakeoutSessionRequest) CRC() uint32 {
	return 0x1d2652ee
}

func AccountFinishTakeoutSession(ctx context.Context, m Requester, i AccountFinishTakeoutSessionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPasswordEmailRequest struct {
	Code string
}

func (*AccountConfirmPasswordEmailRequest) CRC() uint32 {
	return 0x8fdf1920
}

func AccountConfirmPasswordEmail(ctx context.Context, m Requester, i AccountConfirmPasswordEmailRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResendPasswordEmailRequest struct{}

func (*AccountResendPasswordEmailRequest) CRC() uint32 {
	return 0x7a7f2a15
}

func AccountResendPasswordEmail(ctx context.Context, m Requester, i AccountResendPasswordEmailRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCancelPasswordEmailRequest struct{}

func (*AccountCancelPasswordEmailRequest) CRC() uint32 {
	return 0xc1cbd5b6
}

func AccountCancelPasswordEmail(ctx context.Context, m Requester, i AccountCancelPasswordEmailRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetContactSignUpNotificationRequest struct{}

func (*AccountGetContactSignUpNotificationRequest) CRC() uint32 {
	return 0x9f07c728
}

func AccountGetContactSignUpNotification(ctx context.Context, m Requester, i AccountGetContactSignUpNotificationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSetContactSignUpNotificationRequest struct {
	Silent bool
}

func (*AccountSetContactSignUpNotificationRequest) CRC() uint32 {
	return 0xcff43f61
}

func AccountSetContactSignUpNotification(ctx context.Context, m Requester, i AccountSetContactSignUpNotificationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetNotifyExceptionsRequest struct {
	_              struct{}         `tl:"flags,bitflag"`
	CompareSound   bool             `tl:",omitempty:flags:1,implicit"`
	CompareStories bool             `tl:",omitempty:flags:2,implicit"`
	Peer           IInputNotifyPeer `tl:",omitempty:flags:0"`
}

func (*AccountGetNotifyExceptionsRequest) CRC() uint32 {
	return 0x53577479
}

func AccountGetNotifyExceptions(ctx context.Context, m Requester, i AccountGetNotifyExceptionsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPaperRequest struct {
	Wallpaper IInputWallPaper
}

func (*AccountGetWallPaperRequest) CRC() uint32 {
	return 0xfc8ddbea
}

func AccountGetWallPaper(ctx context.Context, m Requester, i AccountGetWallPaperRequest) (IWallPaper, error) {
	var res IWallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountUploadWallPaperRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	ForChat  bool     `tl:",omitempty:flags:0,implicit"`
	File     IInputFile
	MimeType string
	Settings IWallPaperSettings
}

func (*AccountUploadWallPaperRequest) CRC() uint32 {
	return 0xe39a8f03
}

func AccountUploadWallPaper(ctx context.Context, m Requester, i AccountUploadWallPaperRequest) (IWallPaper, error) {
	var res IWallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountSaveWallPaperRequest struct {
	Wallpaper IInputWallPaper
	Unsave    bool
	Settings  IWallPaperSettings
}

func (*AccountSaveWallPaperRequest) CRC() uint32 {
	return 0x6c5a5b37
}

func AccountSaveWallPaper(ctx context.Context, m Requester, i AccountSaveWallPaperRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInstallWallPaperRequest struct {
	Wallpaper IInputWallPaper
	Settings  IWallPaperSettings
}

func (*AccountInstallWallPaperRequest) CRC() uint32 {
	return 0xfeed5769
}

func AccountInstallWallPaper(ctx context.Context, m Requester, i AccountInstallWallPaperRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetWallPapersRequest struct{}

func (*AccountResetWallPapersRequest) CRC() uint32 {
	return 0xbb3b9804
}

func AccountResetWallPapers(ctx context.Context, m Requester, i AccountResetWallPapersRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAutoDownloadSettingsRequest struct{}

func (*AccountGetAutoDownloadSettingsRequest) CRC() uint32 {
	return 0x56da0b3f
}

func AccountGetAutoDownloadSettings(ctx context.Context, m Requester, i AccountGetAutoDownloadSettingsRequest) (IAccountAutoDownloadSettings, error) {
	var res IAccountAutoDownloadSettings
	return res, request(ctx, m, &i, &res)
}

type AccountSaveAutoDownloadSettingsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Low      bool     `tl:",omitempty:flags:0,implicit"`
	High     bool     `tl:",omitempty:flags:1,implicit"`
	Settings IAutoDownloadSettings
}

func (*AccountSaveAutoDownloadSettingsRequest) CRC() uint32 {
	return 0x76f36233
}

func AccountSaveAutoDownloadSettings(ctx context.Context, m Requester, i AccountSaveAutoDownloadSettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUploadThemeRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	File     IInputFile
	Thumb    IInputFile `tl:",omitempty:flags:0"`
	FileName string
	MimeType string
}

func (*AccountUploadThemeRequest) CRC() uint32 {
	return 0x1c3db333
}

func AccountUploadTheme(ctx context.Context, m Requester, i AccountUploadThemeRequest) (IDocument, error) {
	var res IDocument
	return res, request(ctx, m, &i, &res)
}

type AccountCreateThemeRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Slug     string
	Title    string
	Document IInputDocument        `tl:",omitempty:flags:2"`
	Settings []IInputThemeSettings `tl:",omitempty:flags:3"`
}

func (*AccountCreateThemeRequest) CRC() uint32 {
	return 0x652e4400
}

func AccountCreateTheme(ctx context.Context, m Requester, i AccountCreateThemeRequest) (ITheme, error) {
	var res ITheme
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateThemeRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Format   string
	Theme    IInputTheme
	Slug     *string               `tl:",omitempty:flags:0"`
	Title    *string               `tl:",omitempty:flags:1"`
	Document IInputDocument        `tl:",omitempty:flags:2"`
	Settings []IInputThemeSettings `tl:",omitempty:flags:3"`
}

func (*AccountUpdateThemeRequest) CRC() uint32 {
	return 0x2bf40ccc
}

func AccountUpdateTheme(ctx context.Context, m Requester, i AccountUpdateThemeRequest) (ITheme, error) {
	var res ITheme
	return res, request(ctx, m, &i, &res)
}

type AccountSaveThemeRequest struct {
	Theme  IInputTheme
	Unsave bool
}

func (*AccountSaveThemeRequest) CRC() uint32 {
	return 0xf257106c
}

func AccountSaveTheme(ctx context.Context, m Requester, i AccountSaveThemeRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInstallThemeRequest struct {
	_         struct{}      `tl:"flags,bitflag"`
	Dark      bool          `tl:",omitempty:flags:0,implicit"`
	Theme     IInputTheme   `tl:",omitempty:flags:1"`
	Format    *string       `tl:",omitempty:flags:2"`
	BaseTheme EnumBaseTheme `tl:",omitempty:flags:3"`
}

func (*AccountInstallThemeRequest) CRC() uint32 {
	return 0xc727bb3b
}

func AccountInstallTheme(ctx context.Context, m Requester, i AccountInstallThemeRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemeRequest struct {
	Format string
	Theme  IInputTheme
}

func (*AccountGetThemeRequest) CRC() uint32 {
	return 0x3a5869ec
}

func AccountGetTheme(ctx context.Context, m Requester, i AccountGetThemeRequest) (ITheme, error) {
	var res ITheme
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemesRequest struct {
	Format string
	Hash   int64
}

func (*AccountGetThemesRequest) CRC() uint32 {
	return 0x7206e458
}

func AccountGetThemes(ctx context.Context, m Requester, i AccountGetThemesRequest) (IAccountThemes, error) {
	var res IAccountThemes
	return res, request(ctx, m, &i, &res)
}

type AccountSetContentSettingsRequest struct {
	_                struct{} `tl:"flags,bitflag"`
	SensitiveEnabled bool     `tl:",omitempty:flags:0,implicit"`
}

func (*AccountSetContentSettingsRequest) CRC() uint32 {
	return 0xb574b16b
}

func AccountSetContentSettings(ctx context.Context, m Requester, i AccountSetContentSettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetContentSettingsRequest struct{}

func (*AccountGetContentSettingsRequest) CRC() uint32 {
	return 0x8b9b4dae
}

func AccountGetContentSettings(ctx context.Context, m Requester, i AccountGetContentSettingsRequest) (IAccountContentSettings, error) {
	var res IAccountContentSettings
	return res, request(ctx, m, &i, &res)
}

type AccountGetMultiWallPapersRequest struct {
	Wallpapers []IInputWallPaper
}

func (*AccountGetMultiWallPapersRequest) CRC() uint32 {
	return 0x65ad71dc
}

func AccountGetMultiWallPapers(ctx context.Context, m Requester, i AccountGetMultiWallPapersRequest) ([]IWallPaper, error) {
	var res []IWallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountGetGlobalPrivacySettingsRequest struct{}

func (*AccountGetGlobalPrivacySettingsRequest) CRC() uint32 {
	return 0xeb2b4cf6
}

func AccountGetGlobalPrivacySettings(ctx context.Context, m Requester, i AccountGetGlobalPrivacySettingsRequest) (IGlobalPrivacySettings, error) {
	var res IGlobalPrivacySettings
	return res, request(ctx, m, &i, &res)
}

type AccountSetGlobalPrivacySettingsRequest struct {
	Settings IGlobalPrivacySettings
}

func (*AccountSetGlobalPrivacySettingsRequest) CRC() uint32 {
	return 0x1edaaac2
}

func AccountSetGlobalPrivacySettings(ctx context.Context, m Requester, i AccountSetGlobalPrivacySettingsRequest) (IGlobalPrivacySettings, error) {
	var res IGlobalPrivacySettings
	return res, request(ctx, m, &i, &res)
}

type AccountReportProfilePhotoRequest struct {
	Peer    IInputPeer
	PhotoID IInputPhoto
	Reason  EnumReportReason
	Message string
}

func (*AccountReportProfilePhotoRequest) CRC() uint32 {
	return 0xfa8cc6f5
}

func AccountReportProfilePhoto(ctx context.Context, m Requester, i AccountReportProfilePhotoRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetPasswordRequest struct{}

func (*AccountResetPasswordRequest) CRC() uint32 {
	return 0x9308ce1b
}

func AccountResetPassword(ctx context.Context, m Requester, i AccountResetPasswordRequest) (IAccountResetPasswordResult, error) {
	var res IAccountResetPasswordResult
	return res, request(ctx, m, &i, &res)
}

type AccountDeclinePasswordResetRequest struct{}

func (*AccountDeclinePasswordResetRequest) CRC() uint32 {
	return 0x4c9409f6
}

func AccountDeclinePasswordReset(ctx context.Context, m Requester, i AccountDeclinePasswordResetRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetChatThemesRequest struct {
	Hash int64
}

func (*AccountGetChatThemesRequest) CRC() uint32 {
	return 0xd638de89
}

func AccountGetChatThemes(ctx context.Context, m Requester, i AccountGetChatThemesRequest) (IAccountThemes, error) {
	var res IAccountThemes
	return res, request(ctx, m, &i, &res)
}

type AccountSetAuthorizationTTLRequest struct {
	AuthorizationTTLDays int32
}

func (*AccountSetAuthorizationTTLRequest) CRC() uint32 {
	return 0xbf899aa0
}

func AccountSetAuthorizationTTL(ctx context.Context, m Requester, i AccountSetAuthorizationTTLRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountChangeAuthorizationSettingsRequest struct {
	_                         struct{} `tl:"flags,bitflag"`
	Confirmed                 bool     `tl:",omitempty:flags:3,implicit"`
	Hash                      int64
	EncryptedRequestsDisabled *bool `tl:",omitempty:flags:0"`
	CallRequestsDisabled      *bool `tl:",omitempty:flags:1"`
}

func (*AccountChangeAuthorizationSettingsRequest) CRC() uint32 {
	return 0x40f48462
}

func AccountChangeAuthorizationSettings(ctx context.Context, m Requester, i AccountChangeAuthorizationSettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetSavedRingtonesRequest struct {
	Hash int64
}

func (*AccountGetSavedRingtonesRequest) CRC() uint32 {
	return 0xe1902288
}

func AccountGetSavedRingtones(ctx context.Context, m Requester, i AccountGetSavedRingtonesRequest) (IAccountSavedRingtones, error) {
	var res IAccountSavedRingtones
	return res, request(ctx, m, &i, &res)
}

type AccountSaveRingtoneRequest struct {
	ID     IInputDocument
	Unsave bool
}

func (*AccountSaveRingtoneRequest) CRC() uint32 {
	return 0x3dea5b03
}

func AccountSaveRingtone(ctx context.Context, m Requester, i AccountSaveRingtoneRequest) (IAccountSavedRingtone, error) {
	var res IAccountSavedRingtone
	return res, request(ctx, m, &i, &res)
}

type AccountUploadRingtoneRequest struct {
	File     IInputFile
	FileName string
	MimeType string
}

func (*AccountUploadRingtoneRequest) CRC() uint32 {
	return 0x831a83a2
}

func AccountUploadRingtone(ctx context.Context, m Requester, i AccountUploadRingtoneRequest) (IDocument, error) {
	var res IDocument
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateEmojiStatusRequest struct {
	EmojiStatus IEmojiStatus
}

func (*AccountUpdateEmojiStatusRequest) CRC() uint32 {
	return 0xfbd3de6b
}

func AccountUpdateEmojiStatus(ctx context.Context, m Requester, i AccountUpdateEmojiStatusRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultEmojiStatusesRequest struct {
	Hash int64
}

func (*AccountGetDefaultEmojiStatusesRequest) CRC() uint32 {
	return 0xd6753386
}

func AccountGetDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetDefaultEmojiStatusesRequest) (IAccountEmojiStatuses, error) {
	var res IAccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetRecentEmojiStatusesRequest struct {
	Hash int64
}

func (*AccountGetRecentEmojiStatusesRequest) CRC() uint32 {
	return 0xf578105
}

func AccountGetRecentEmojiStatuses(ctx context.Context, m Requester, i AccountGetRecentEmojiStatusesRequest) (IAccountEmojiStatuses, error) {
	var res IAccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountClearRecentEmojiStatusesRequest struct{}

func (*AccountClearRecentEmojiStatusesRequest) CRC() uint32 {
	return 0x18201aae
}

func AccountClearRecentEmojiStatuses(ctx context.Context, m Requester, i AccountClearRecentEmojiStatusesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountReorderUsernamesRequest struct {
	Order []string
}

func (*AccountReorderUsernamesRequest) CRC() uint32 {
	return 0xef500eab
}

func AccountReorderUsernames(ctx context.Context, m Requester, i AccountReorderUsernamesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleUsernameRequest struct {
	Username string
	Active   bool
}

func (*AccountToggleUsernameRequest) CRC() uint32 {
	return 0x58d6b376
}

func AccountToggleUsername(ctx context.Context, m Requester, i AccountToggleUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultProfilePhotoEmojisRequest struct {
	Hash int64
}

func (*AccountGetDefaultProfilePhotoEmojisRequest) CRC() uint32 {
	return 0xe2750328
}

func AccountGetDefaultProfilePhotoEmojis(ctx context.Context, m Requester, i AccountGetDefaultProfilePhotoEmojisRequest) (IEmojiList, error) {
	var res IEmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultGroupPhotoEmojisRequest struct {
	Hash int64
}

func (*AccountGetDefaultGroupPhotoEmojisRequest) CRC() uint32 {
	return 0x915860ae
}

func AccountGetDefaultGroupPhotoEmojis(ctx context.Context, m Requester, i AccountGetDefaultGroupPhotoEmojisRequest) (IEmojiList, error) {
	var res IEmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetAutoSaveSettingsRequest struct{}

func (*AccountGetAutoSaveSettingsRequest) CRC() uint32 {
	return 0xadcbbcda
}

func AccountGetAutoSaveSettings(ctx context.Context, m Requester, i AccountGetAutoSaveSettingsRequest) (IAccountAutoSaveSettings, error) {
	var res IAccountAutoSaveSettings
	return res, request(ctx, m, &i, &res)
}

type AccountSaveAutoSaveSettingsRequest struct {
	_          struct{}   `tl:"flags,bitflag"`
	Users      bool       `tl:",omitempty:flags:0,implicit"`
	Chats      bool       `tl:",omitempty:flags:1,implicit"`
	Broadcasts bool       `tl:",omitempty:flags:2,implicit"`
	Peer       IInputPeer `tl:",omitempty:flags:3"`
	Settings   IAutoSaveSettings
}

func (*AccountSaveAutoSaveSettingsRequest) CRC() uint32 {
	return 0xd69b8361
}

func AccountSaveAutoSaveSettings(ctx context.Context, m Requester, i AccountSaveAutoSaveSettingsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteAutoSaveExceptionsRequest struct{}

func (*AccountDeleteAutoSaveExceptionsRequest) CRC() uint32 {
	return 0x53bc0020
}

func AccountDeleteAutoSaveExceptions(ctx context.Context, m Requester, i AccountDeleteAutoSaveExceptionsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInvalidateSignInCodesRequest struct {
	Codes []string
}

func (*AccountInvalidateSignInCodesRequest) CRC() uint32 {
	return 0xca8ae8ba
}

func AccountInvalidateSignInCodes(ctx context.Context, m Requester, i AccountInvalidateSignInCodesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateColorRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	ForProfile        bool     `tl:",omitempty:flags:1,implicit"`
	Color             *int32   `tl:",omitempty:flags:2"`
	BackgroundEmojiID *int64   `tl:",omitempty:flags:0"`
}

func (*AccountUpdateColorRequest) CRC() uint32 {
	return 0x7cefa15d
}

func AccountUpdateColor(ctx context.Context, m Requester, i AccountUpdateColorRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultBackgroundEmojisRequest struct {
	Hash int64
}

func (*AccountGetDefaultBackgroundEmojisRequest) CRC() uint32 {
	return 0xa60ab9ce
}

func AccountGetDefaultBackgroundEmojis(ctx context.Context, m Requester, i AccountGetDefaultBackgroundEmojisRequest) (IEmojiList, error) {
	var res IEmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelDefaultEmojiStatusesRequest struct {
	Hash int64
}

func (*AccountGetChannelDefaultEmojiStatusesRequest) CRC() uint32 {
	return 0x7727a7d5
}

func AccountGetChannelDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetChannelDefaultEmojiStatusesRequest) (IAccountEmojiStatuses, error) {
	var res IAccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelRestrictedStatusEmojisRequest struct {
	Hash int64
}

func (*AccountGetChannelRestrictedStatusEmojisRequest) CRC() uint32 {
	return 0x35a9e0d5
}

func AccountGetChannelRestrictedStatusEmojis(ctx context.Context, m Requester, i AccountGetChannelRestrictedStatusEmojisRequest) (IEmojiList, error) {
	var res IEmojiList
	return res, request(ctx, m, &i, &res)
}

// Send the verification code for login
type AuthSendCodeRequest struct {
	PhoneNumber string        // Phone number in international format
	APIID       int32         // Application identifier (see App configuration)
	APIHash     string        // Application secret hash (see App configuration)
	Settings    ICodeSettings // Settings for the code type to send
}

func (*AuthSendCodeRequest) CRC() uint32 {
	return 0xa677244f
}

func AuthSendCode(ctx context.Context, m Requester, i AuthSendCodeRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

// Registers a validated phone number in the system.
type AuthSignUpRequest struct {
	PhoneNumber   string // Phone number in international format
	PhoneCodeHash string // SMS-message ID
	FirstName     string // New user first name
	LastName      string // New user last name
}

func (*AuthSignUpRequest) CRC() uint32 {
	return 0x80eee427
}

func AuthSignUp(ctx context.Context, m Requester, i AuthSignUpRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthSignInRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	PhoneNumber       string
	PhoneCodeHash     string
	PhoneCode         *string            `tl:",omitempty:flags:0"`
	EmailVerification IEmailVerification `tl:",omitempty:flags:1"`
}

func (*AuthSignInRequest) CRC() uint32 {
	return 0x8d52a951
}

func AuthSignIn(ctx context.Context, m Requester, i AuthSignInRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthLogOutRequest struct{}

func (*AuthLogOutRequest) CRC() uint32 {
	return 0x3e72ba19
}

func AuthLogOut(ctx context.Context, m Requester, i AuthLogOutRequest) (IAuthLoggedOut, error) {
	var res IAuthLoggedOut
	return res, request(ctx, m, &i, &res)
}

type AuthResetAuthorizationsRequest struct{}

func (*AuthResetAuthorizationsRequest) CRC() uint32 {
	return 0x9fab0d1a
}

func AuthResetAuthorizations(ctx context.Context, m Requester, i AuthResetAuthorizationsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthExportAuthorizationRequest struct {
	DcID int32
}

func (*AuthExportAuthorizationRequest) CRC() uint32 {
	return 0xe5bfffcd
}

func AuthExportAuthorization(ctx context.Context, m Requester, i AuthExportAuthorizationRequest) (IAuthExportedAuthorization, error) {
	var res IAuthExportedAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthImportAuthorizationRequest struct {
	ID    int64
	Bytes []byte
}

func (*AuthImportAuthorizationRequest) CRC() uint32 {
	return 0xa57a7dad
}

func AuthImportAuthorization(ctx context.Context, m Requester, i AuthImportAuthorizationRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthBindTempAuthKeyRequest struct {
	PermAuthKeyID    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

func (*AuthBindTempAuthKeyRequest) CRC() uint32 {
	return 0xcdd42a05
}

func AuthBindTempAuthKey(ctx context.Context, m Requester, i AuthBindTempAuthKeyRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportBotAuthorizationRequest struct {
	Flags        int32
	APIID        int32
	APIHash      string
	BotAuthToken string
}

func (*AuthImportBotAuthorizationRequest) CRC() uint32 {
	return 0x67a3ff2c
}

func AuthImportBotAuthorization(ctx context.Context, m Requester, i AuthImportBotAuthorizationRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckPasswordRequest struct {
	Password IInputCheckPasswordSRP
}

func (*AuthCheckPasswordRequest) CRC() uint32 {
	return 0xd18b4d16
}

func AuthCheckPassword(ctx context.Context, m Requester, i AuthCheckPasswordRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthRequestPasswordRecoveryRequest struct{}

func (*AuthRequestPasswordRecoveryRequest) CRC() uint32 {
	return 0xd897bc66
}

func AuthRequestPasswordRecovery(ctx context.Context, m Requester, i AuthRequestPasswordRecoveryRequest) (IAuthPasswordRecovery, error) {
	var res IAuthPasswordRecovery
	return res, request(ctx, m, &i, &res)
}

type AuthRecoverPasswordRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Code        string
	NewSettings IAccountPasswordInputSettings `tl:",omitempty:flags:0"`
}

func (*AuthRecoverPasswordRequest) CRC() uint32 {
	return 0x37096c70
}

func AuthRecoverPassword(ctx context.Context, m Requester, i AuthRecoverPasswordRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthResendCodeRequest struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*AuthResendCodeRequest) CRC() uint32 {
	return 0x3ef1a9bf
}

func AuthResendCode(ctx context.Context, m Requester, i AuthResendCodeRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthCancelCodeRequest struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*AuthCancelCodeRequest) CRC() uint32 {
	return 0x1f040578
}

func AuthCancelCode(ctx context.Context, m Requester, i AuthCancelCodeRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthDropTempAuthKeysRequest struct {
	ExceptAuthKeys []int64
}

func (*AuthDropTempAuthKeysRequest) CRC() uint32 {
	return 0x8e48a188
}

func AuthDropTempAuthKeys(ctx context.Context, m Requester, i AuthDropTempAuthKeysRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthExportLoginTokenRequest struct {
	APIID     int32
	APIHash   string
	ExceptIds []int64
}

func (*AuthExportLoginTokenRequest) CRC() uint32 {
	return 0xb7e085fe
}

func AuthExportLoginToken(ctx context.Context, m Requester, i AuthExportLoginTokenRequest) (IAuthLoginToken, error) {
	var res IAuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthImportLoginTokenRequest struct {
	Token []byte
}

func (*AuthImportLoginTokenRequest) CRC() uint32 {
	return 0x95ac5ce4
}

func AuthImportLoginToken(ctx context.Context, m Requester, i AuthImportLoginTokenRequest) (IAuthLoginToken, error) {
	var res IAuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthAcceptLoginTokenRequest struct {
	Token []byte
}

func (*AuthAcceptLoginTokenRequest) CRC() uint32 {
	return 0xe894ad4d
}

func AuthAcceptLoginToken(ctx context.Context, m Requester, i AuthAcceptLoginTokenRequest) (IAuthorization, error) {
	var res IAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckRecoveryPasswordRequest struct {
	Code string
}

func (*AuthCheckRecoveryPasswordRequest) CRC() uint32 {
	return 0xd36bf79
}

func AuthCheckRecoveryPassword(ctx context.Context, m Requester, i AuthCheckRecoveryPasswordRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportWebTokenAuthorizationRequest struct {
	APIID        int32
	APIHash      string
	WebAuthToken string
}

func (*AuthImportWebTokenAuthorizationRequest) CRC() uint32 {
	return 0x2db873a9
}

func AuthImportWebTokenAuthorization(ctx context.Context, m Requester, i AuthImportWebTokenAuthorizationRequest) (IAuthAuthorization, error) {
	var res IAuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthRequestFirebaseSmsRequest struct {
	_              struct{} `tl:"flags,bitflag"`
	PhoneNumber    string
	PhoneCodeHash  string
	SafetyNetToken *string `tl:",omitempty:flags:0"`
	IosPushSecret  *string `tl:",omitempty:flags:1"`
}

func (*AuthRequestFirebaseSmsRequest) CRC() uint32 {
	return 0x89464b50
}

func AuthRequestFirebaseSms(ctx context.Context, m Requester, i AuthRequestFirebaseSmsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthResetLoginEmailRequest struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*AuthResetLoginEmailRequest) CRC() uint32 {
	return 0x7e960193
}

func AuthResetLoginEmail(ctx context.Context, m Requester, i AuthResetLoginEmailRequest) (IAuthSentCode, error) {
	var res IAuthSentCode
	return res, request(ctx, m, &i, &res)
}

type BotsSendCustomRequestRequest struct {
	CustomMethod string
	Params       IDataJSON
}

func (*BotsSendCustomRequestRequest) CRC() uint32 {
	return 0xaa2769ed
}

func BotsSendCustomRequest(ctx context.Context, m Requester, i BotsSendCustomRequestRequest) (IDataJSON, error) {
	var res IDataJSON
	return res, request(ctx, m, &i, &res)
}

type BotsAnswerWebhookJSONQueryRequest struct {
	QueryID int64
	Data    IDataJSON
}

func (*BotsAnswerWebhookJSONQueryRequest) CRC() uint32 {
	return 0xe6213f4d
}

func BotsAnswerWebhookJSONQuery(ctx context.Context, m Requester, i BotsAnswerWebhookJSONQueryRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotCommandsRequest struct {
	Scope    IBotCommandScope
	LangCode string
	Commands []IBotCommand
}

func (*BotsSetBotCommandsRequest) CRC() uint32 {
	return 0x517165a
}

func BotsSetBotCommands(ctx context.Context, m Requester, i BotsSetBotCommandsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsResetBotCommandsRequest struct {
	Scope    IBotCommandScope
	LangCode string
}

func (*BotsResetBotCommandsRequest) CRC() uint32 {
	return 0x3d8de0f9
}

func BotsResetBotCommands(ctx context.Context, m Requester, i BotsResetBotCommandsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotCommandsRequest struct {
	Scope    IBotCommandScope
	LangCode string
}

func (*BotsGetBotCommandsRequest) CRC() uint32 {
	return 0xe34c0dd6
}

func BotsGetBotCommands(ctx context.Context, m Requester, i BotsGetBotCommandsRequest) ([]IBotCommand, error) {
	var res []IBotCommand
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotMenuButtonRequest struct {
	UserID IInputUser
	Button IBotMenuButton
}

func (*BotsSetBotMenuButtonRequest) CRC() uint32 {
	return 0x4504d54f
}

func BotsSetBotMenuButton(ctx context.Context, m Requester, i BotsSetBotMenuButtonRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotMenuButtonRequest struct {
	UserID IInputUser
}

func (*BotsGetBotMenuButtonRequest) CRC() uint32 {
	return 0x9c60eb28
}

func BotsGetBotMenuButton(ctx context.Context, m Requester, i BotsGetBotMenuButtonRequest) (IBotMenuButton, error) {
	var res IBotMenuButton
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotBroadcastDefaultAdminRightsRequest struct {
	AdminRights IChatAdminRights
}

func (*BotsSetBotBroadcastDefaultAdminRightsRequest) CRC() uint32 {
	return 0x788464e1
}

func BotsSetBotBroadcastDefaultAdminRights(ctx context.Context, m Requester, i BotsSetBotBroadcastDefaultAdminRightsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotGroupDefaultAdminRightsRequest struct {
	AdminRights IChatAdminRights
}

func (*BotsSetBotGroupDefaultAdminRightsRequest) CRC() uint32 {
	return 0x925ec9ea
}

func BotsSetBotGroupDefaultAdminRights(ctx context.Context, m Requester, i BotsSetBotGroupDefaultAdminRightsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotInfoRequest struct {
	_           struct{}   `tl:"flags,bitflag"`
	Bot         IInputUser `tl:",omitempty:flags:2"`
	LangCode    string
	Name        *string `tl:",omitempty:flags:3"`
	About       *string `tl:",omitempty:flags:0"`
	Description *string `tl:",omitempty:flags:1"`
}

func (*BotsSetBotInfoRequest) CRC() uint32 {
	return 0x10cf3123
}

func BotsSetBotInfo(ctx context.Context, m Requester, i BotsSetBotInfoRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotInfoRequest struct {
	_        struct{}   `tl:"flags,bitflag"`
	Bot      IInputUser `tl:",omitempty:flags:0"`
	LangCode string
}

func (*BotsGetBotInfoRequest) CRC() uint32 {
	return 0xdcd914fd
}

func BotsGetBotInfo(ctx context.Context, m Requester, i BotsGetBotInfoRequest) (IBotsBotInfo, error) {
	var res IBotsBotInfo
	return res, request(ctx, m, &i, &res)
}

type BotsReorderUsernamesRequest struct {
	Bot   IInputUser
	Order []string
}

func (*BotsReorderUsernamesRequest) CRC() uint32 {
	return 0x9709b1c2
}

func BotsReorderUsernames(ctx context.Context, m Requester, i BotsReorderUsernamesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsToggleUsernameRequest struct {
	Bot      IInputUser
	Username string
	Active   bool
}

func (*BotsToggleUsernameRequest) CRC() uint32 {
	return 0x53ca973
}

func BotsToggleUsername(ctx context.Context, m Requester, i BotsToggleUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsCanSendMessageRequest struct {
	Bot IInputUser
}

func (*BotsCanSendMessageRequest) CRC() uint32 {
	return 0x1359f4e6
}

func BotsCanSendMessage(ctx context.Context, m Requester, i BotsCanSendMessageRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsAllowSendMessageRequest struct {
	Bot IInputUser
}

func (*BotsAllowSendMessageRequest) CRC() uint32 {
	return 0xf132e3ef
}

func BotsAllowSendMessage(ctx context.Context, m Requester, i BotsAllowSendMessageRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type BotsInvokeWebViewCustomMethodRequest struct {
	Bot          IInputUser
	CustomMethod string
	Params       IDataJSON
}

func (*BotsInvokeWebViewCustomMethodRequest) CRC() uint32 {
	return 0x87fc5e7
}

func BotsInvokeWebViewCustomMethod(ctx context.Context, m Requester, i BotsInvokeWebViewCustomMethodRequest) (IDataJSON, error) {
	var res IDataJSON
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadHistoryRequest struct {
	Channel IInputChannel
	MaxID   int32
}

func (*ChannelsReadHistoryRequest) CRC() uint32 {
	return 0xcc104937
}

func ChannelsReadHistory(ctx context.Context, m Requester, i ChannelsReadHistoryRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteMessagesRequest struct {
	Channel IInputChannel
	ID      []int32
}

func (*ChannelsDeleteMessagesRequest) CRC() uint32 {
	return 0x84c1fd4e
}

func ChannelsDeleteMessages(ctx context.Context, m Requester, i ChannelsDeleteMessagesRequest) (IMessagesAffectedMessages, error) {
	var res IMessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportSpamRequest struct {
	Channel     IInputChannel
	Participant IInputPeer
	ID          []int32
}

func (*ChannelsReportSpamRequest) CRC() uint32 {
	return 0xf44a8315
}

func ChannelsReportSpam(ctx context.Context, m Requester, i ChannelsReportSpamRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetMessagesRequest struct {
	Channel IInputChannel
	ID      []IInputMessage
}

func (*ChannelsGetMessagesRequest) CRC() uint32 {
	return 0xad8c9a23
}

func ChannelsGetMessages(ctx context.Context, m Requester, i ChannelsGetMessagesRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

// Get the participants of a channel
type ChannelsGetParticipantsRequest struct {
	Channel IInputChannel              // Channel
	Filter  IChannelParticipantsFilter // Which participant types to fetch
	Offset  int32
	Limit   int32
	Hash    int64
}

func (*ChannelsGetParticipantsRequest) CRC() uint32 {
	return 0x77ced9d0
}

func ChannelsGetParticipants(ctx context.Context, m Requester, i ChannelsGetParticipantsRequest) (IChannelsChannelParticipants, error) {
	var res IChannelsChannelParticipants
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetParticipantRequest struct {
	Channel     IInputChannel
	Participant IInputPeer
}

func (*ChannelsGetParticipantRequest) CRC() uint32 {
	return 0xa0ab6cc6
}

func ChannelsGetParticipant(ctx context.Context, m Requester, i ChannelsGetParticipantRequest) (IChannelsChannelParticipant, error) {
	var res IChannelsChannelParticipant
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetChannelsRequest struct {
	ID []IInputChannel
}

func (*ChannelsGetChannelsRequest) CRC() uint32 {
	return 0xa7f6bbb
}

func ChannelsGetChannels(ctx context.Context, m Requester, i ChannelsGetChannelsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetFullChannelRequest struct {
	Channel IInputChannel
}

func (*ChannelsGetFullChannelRequest) CRC() uint32 {
	return 0x8736a09
}

func ChannelsGetFullChannel(ctx context.Context, m Requester, i ChannelsGetFullChannelRequest) (IMessagesChatFull, error) {
	var res IMessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateChannelRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Broadcast bool     `tl:",omitempty:flags:0,implicit"`
	Megagroup bool     `tl:",omitempty:flags:1,implicit"`
	ForImport bool     `tl:",omitempty:flags:3,implicit"`
	Forum     bool     `tl:",omitempty:flags:5,implicit"`
	Title     string
	About     string
	GeoPoint  IInputGeoPoint `tl:",omitempty:flags:2"`
	Address   *string        `tl:",omitempty:flags:2"`
	TTLPeriod *int32         `tl:",omitempty:flags:4"`
}

func (*ChannelsCreateChannelRequest) CRC() uint32 {
	return 0x91006707
}

func ChannelsCreateChannel(ctx context.Context, m Requester, i ChannelsCreateChannelRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditAdminRequest struct {
	Channel     IInputChannel
	UserID      IInputUser
	AdminRights IChatAdminRights
	Rank        string
}

func (*ChannelsEditAdminRequest) CRC() uint32 {
	return 0xd33c8902
}

func ChannelsEditAdmin(ctx context.Context, m Requester, i ChannelsEditAdminRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditTitleRequest struct {
	Channel IInputChannel
	Title   string
}

func (*ChannelsEditTitleRequest) CRC() uint32 {
	return 0x566decd0
}

func ChannelsEditTitle(ctx context.Context, m Requester, i ChannelsEditTitleRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditPhotoRequest struct {
	Channel IInputChannel
	Photo   IInputChatPhoto
}

func (*ChannelsEditPhotoRequest) CRC() uint32 {
	return 0xf12e57c9
}

func ChannelsEditPhoto(ctx context.Context, m Requester, i ChannelsEditPhotoRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCheckUsernameRequest struct {
	Channel  IInputChannel
	Username string
}

func (*ChannelsCheckUsernameRequest) CRC() uint32 {
	return 0x10e6bd2c
}

func ChannelsCheckUsername(ctx context.Context, m Requester, i ChannelsCheckUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateUsernameRequest struct {
	Channel  IInputChannel
	Username string
}

func (*ChannelsUpdateUsernameRequest) CRC() uint32 {
	return 0x3514b3de
}

func ChannelsUpdateUsername(ctx context.Context, m Requester, i ChannelsUpdateUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsJoinChannelRequest struct {
	Channel IInputChannel
}

func (*ChannelsJoinChannelRequest) CRC() uint32 {
	return 0x24b524c5
}

func ChannelsJoinChannel(ctx context.Context, m Requester, i ChannelsJoinChannelRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsLeaveChannelRequest struct {
	Channel IInputChannel
}

func (*ChannelsLeaveChannelRequest) CRC() uint32 {
	return 0xf836aa95
}

func ChannelsLeaveChannel(ctx context.Context, m Requester, i ChannelsLeaveChannelRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsInviteToChannelRequest struct {
	Channel IInputChannel
	Users   []IInputUser
}

func (*ChannelsInviteToChannelRequest) CRC() uint32 {
	return 0x199f3a6c
}

func ChannelsInviteToChannel(ctx context.Context, m Requester, i ChannelsInviteToChannelRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteChannelRequest struct {
	Channel IInputChannel
}

func (*ChannelsDeleteChannelRequest) CRC() uint32 {
	return 0xc0111fe3
}

func ChannelsDeleteChannel(ctx context.Context, m Requester, i ChannelsDeleteChannelRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsExportMessageLinkRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Grouped bool     `tl:",omitempty:flags:0,implicit"`
	Thread  bool     `tl:",omitempty:flags:1,implicit"`
	Channel IInputChannel
	ID      int32
}

func (*ChannelsExportMessageLinkRequest) CRC() uint32 {
	return 0xe63fadeb
}

func ChannelsExportMessageLink(ctx context.Context, m Requester, i ChannelsExportMessageLinkRequest) (IExportedMessageLink, error) {
	var res IExportedMessageLink
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSignaturesRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleSignaturesRequest) CRC() uint32 {
	return 0x1f69b606
}

func ChannelsToggleSignatures(ctx context.Context, m Requester, i ChannelsToggleSignaturesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetAdminedPublicChannelsRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	ByLocation bool     `tl:",omitempty:flags:0,implicit"`
	CheckLimit bool     `tl:",omitempty:flags:1,implicit"`
}

func (*ChannelsGetAdminedPublicChannelsRequest) CRC() uint32 {
	return 0xf8b036af
}

func ChannelsGetAdminedPublicChannels(ctx context.Context, m Requester, i ChannelsGetAdminedPublicChannelsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditBannedRequest struct {
	Channel      IInputChannel
	Participant  IInputPeer
	BannedRights IChatBannedRights
}

func (*ChannelsEditBannedRequest) CRC() uint32 {
	return 0x96e6cd81
}

func ChannelsEditBanned(ctx context.Context, m Requester, i ChannelsEditBannedRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetAdminLogRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Channel      IInputChannel
	Q            string
	EventsFilter IChannelAdminLogEventsFilter `tl:",omitempty:flags:0"`
	Admins       []IInputUser                 `tl:",omitempty:flags:1"`
	MaxID        int64
	MinID        int64
	Limit        int32
}

func (*ChannelsGetAdminLogRequest) CRC() uint32 {
	return 0x33ddf480
}

func ChannelsGetAdminLog(ctx context.Context, m Requester, i ChannelsGetAdminLogRequest) (IChannelsAdminLogResults, error) {
	var res IChannelsAdminLogResults
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetStickersRequest struct {
	Channel    IInputChannel
	Stickerset IInputStickerSet
}

func (*ChannelsSetStickersRequest) CRC() uint32 {
	return 0xea8ca4f9
}

func ChannelsSetStickers(ctx context.Context, m Requester, i ChannelsSetStickersRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadMessageContentsRequest struct {
	Channel IInputChannel
	ID      []int32
}

func (*ChannelsReadMessageContentsRequest) CRC() uint32 {
	return 0xeab5dc38
}

func ChannelsReadMessageContents(ctx context.Context, m Requester, i ChannelsReadMessageContentsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteHistoryRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	ForEveryone bool     `tl:",omitempty:flags:0,implicit"`
	Channel     IInputChannel
	MaxID       int32
}

func (*ChannelsDeleteHistoryRequest) CRC() uint32 {
	return 0x9baa9647
}

func ChannelsDeleteHistory(ctx context.Context, m Requester, i ChannelsDeleteHistoryRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsTogglePreHistoryHiddenRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsTogglePreHistoryHiddenRequest) CRC() uint32 {
	return 0xeabbb94c
}

func ChannelsTogglePreHistoryHidden(ctx context.Context, m Requester, i ChannelsTogglePreHistoryHiddenRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetLeftChannelsRequest struct {
	Offset int32
}

func (*ChannelsGetLeftChannelsRequest) CRC() uint32 {
	return 0x8341ecc0
}

func ChannelsGetLeftChannels(ctx context.Context, m Requester, i ChannelsGetLeftChannelsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetGroupsForDiscussionRequest struct{}

func (*ChannelsGetGroupsForDiscussionRequest) CRC() uint32 {
	return 0xf5dad378
}

func ChannelsGetGroupsForDiscussion(ctx context.Context, m Requester, i ChannelsGetGroupsForDiscussionRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetDiscussionGroupRequest struct {
	Broadcast IInputChannel
	Group     IInputChannel
}

func (*ChannelsSetDiscussionGroupRequest) CRC() uint32 {
	return 0x40582bb2
}

func ChannelsSetDiscussionGroup(ctx context.Context, m Requester, i ChannelsSetDiscussionGroupRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditCreatorRequest struct {
	Channel  IInputChannel
	UserID   IInputUser
	Password IInputCheckPasswordSRP
}

func (*ChannelsEditCreatorRequest) CRC() uint32 {
	return 0x8f38cd1f
}

func ChannelsEditCreator(ctx context.Context, m Requester, i ChannelsEditCreatorRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditLocationRequest struct {
	Channel  IInputChannel
	GeoPoint IInputGeoPoint
	Address  string
}

func (*ChannelsEditLocationRequest) CRC() uint32 {
	return 0x58e63f6d
}

func ChannelsEditLocation(ctx context.Context, m Requester, i ChannelsEditLocationRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSlowModeRequest struct {
	Channel IInputChannel
	Seconds int32
}

func (*ChannelsToggleSlowModeRequest) CRC() uint32 {
	return 0xedd49ef0
}

func ChannelsToggleSlowMode(ctx context.Context, m Requester, i ChannelsToggleSlowModeRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetInactiveChannelsRequest struct{}

func (*ChannelsGetInactiveChannelsRequest) CRC() uint32 {
	return 0x11e831ee
}

func ChannelsGetInactiveChannels(ctx context.Context, m Requester, i ChannelsGetInactiveChannelsRequest) (IMessagesInactiveChats, error) {
	var res IMessagesInactiveChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsConvertToGigagroupRequest struct {
	Channel IInputChannel
}

func (*ChannelsConvertToGigagroupRequest) CRC() uint32 {
	return 0xb290c69
}

func ChannelsConvertToGigagroup(ctx context.Context, m Requester, i ChannelsConvertToGigagroupRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsViewSponsoredMessageRequest struct {
	Channel  IInputChannel
	RandomID []byte
}

func (*ChannelsViewSponsoredMessageRequest) CRC() uint32 {
	return 0xbeaedb94
}

func ChannelsViewSponsoredMessage(ctx context.Context, m Requester, i ChannelsViewSponsoredMessageRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSponsoredMessagesRequest struct {
	Channel IInputChannel
}

func (*ChannelsGetSponsoredMessagesRequest) CRC() uint32 {
	return 0xec210fbf
}

func ChannelsGetSponsoredMessages(ctx context.Context, m Requester, i ChannelsGetSponsoredMessagesRequest) (IMessagesSponsoredMessages, error) {
	var res IMessagesSponsoredMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSendAsRequest struct {
	Peer IInputPeer
}

func (*ChannelsGetSendAsRequest) CRC() uint32 {
	return 0xdc770ee
}

func ChannelsGetSendAs(ctx context.Context, m Requester, i ChannelsGetSendAsRequest) (IChannelsSendAsPeers, error) {
	var res IChannelsSendAsPeers
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteParticipantHistoryRequest struct {
	Channel     IInputChannel
	Participant IInputPeer
}

func (*ChannelsDeleteParticipantHistoryRequest) CRC() uint32 {
	return 0x367544db
}

func ChannelsDeleteParticipantHistory(ctx context.Context, m Requester, i ChannelsDeleteParticipantHistoryRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinToSendRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleJoinToSendRequest) CRC() uint32 {
	return 0xe4cb9580
}

func ChannelsToggleJoinToSend(ctx context.Context, m Requester, i ChannelsToggleJoinToSendRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinRequestRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleJoinRequestRequest) CRC() uint32 {
	return 0x4c2985b6
}

func ChannelsToggleJoinRequest(ctx context.Context, m Requester, i ChannelsToggleJoinRequestRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderUsernamesRequest struct {
	Channel IInputChannel
	Order   []string
}

func (*ChannelsReorderUsernamesRequest) CRC() uint32 {
	return 0xb45ced1d
}

func ChannelsReorderUsernames(ctx context.Context, m Requester, i ChannelsReorderUsernamesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleUsernameRequest struct {
	Channel  IInputChannel
	Username string
	Active   bool
}

func (*ChannelsToggleUsernameRequest) CRC() uint32 {
	return 0x50f24105
}

func ChannelsToggleUsername(ctx context.Context, m Requester, i ChannelsToggleUsernameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeactivateAllUsernamesRequest struct {
	Channel IInputChannel
}

func (*ChannelsDeactivateAllUsernamesRequest) CRC() uint32 {
	return 0xa245dd3
}

func ChannelsDeactivateAllUsernames(ctx context.Context, m Requester, i ChannelsDeactivateAllUsernamesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleForumRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleForumRequest) CRC() uint32 {
	return 0xa4298b29
}

func ChannelsToggleForum(ctx context.Context, m Requester, i ChannelsToggleForumRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateForumTopicRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     IInputChannel
	Title       string
	IconColor   *int32 `tl:",omitempty:flags:0"`
	IconEmojiID *int64 `tl:",omitempty:flags:3"`
	RandomID    int64
	SendAs      IInputPeer `tl:",omitempty:flags:2"`
}

func (*ChannelsCreateForumTopicRequest) CRC() uint32 {
	return 0xf40c0224
}

func ChannelsCreateForumTopic(ctx context.Context, m Requester, i ChannelsCreateForumTopicRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     IInputChannel
	Q           *string `tl:",omitempty:flags:0"`
	OffsetDate  int32
	OffsetID    int32
	OffsetTopic int32
	Limit       int32
}

func (*ChannelsGetForumTopicsRequest) CRC() uint32 {
	return 0xde560d1
}

func ChannelsGetForumTopics(ctx context.Context, m Requester, i ChannelsGetForumTopicsRequest) (IMessagesForumTopics, error) {
	var res IMessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsByIDRequest struct {
	Channel IInputChannel
	Topics  []int32
}

func (*ChannelsGetForumTopicsByIDRequest) CRC() uint32 {
	return 0xb0831eb9
}

func ChannelsGetForumTopicsByID(ctx context.Context, m Requester, i ChannelsGetForumTopicsByIDRequest) (IMessagesForumTopics, error) {
	var res IMessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditForumTopicRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     IInputChannel
	TopicID     int32
	Title       *string `tl:",omitempty:flags:0"`
	IconEmojiID *int64  `tl:",omitempty:flags:1"`
	Closed      *bool   `tl:",omitempty:flags:2"`
	Hidden      *bool   `tl:",omitempty:flags:3"`
}

func (*ChannelsEditForumTopicRequest) CRC() uint32 {
	return 0xf4dfa185
}

func ChannelsEditForumTopic(ctx context.Context, m Requester, i ChannelsEditForumTopicRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdatePinnedForumTopicRequest struct {
	Channel IInputChannel
	TopicID int32
	Pinned  bool
}

func (*ChannelsUpdatePinnedForumTopicRequest) CRC() uint32 {
	return 0x6c2d9026
}

func ChannelsUpdatePinnedForumTopic(ctx context.Context, m Requester, i ChannelsUpdatePinnedForumTopicRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteTopicHistoryRequest struct {
	Channel  IInputChannel
	TopMsgID int32
}

func (*ChannelsDeleteTopicHistoryRequest) CRC() uint32 {
	return 0x34435f2d
}

func ChannelsDeleteTopicHistory(ctx context.Context, m Requester, i ChannelsDeleteTopicHistoryRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderPinnedForumTopicsRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Force   bool     `tl:",omitempty:flags:0,implicit"`
	Channel IInputChannel
	Order   []int32
}

func (*ChannelsReorderPinnedForumTopicsRequest) CRC() uint32 {
	return 0x2950a18f
}

func ChannelsReorderPinnedForumTopics(ctx context.Context, m Requester, i ChannelsReorderPinnedForumTopicsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleAntiSpamRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleAntiSpamRequest) CRC() uint32 {
	return 0x68f3e4eb
}

func ChannelsToggleAntiSpam(ctx context.Context, m Requester, i ChannelsToggleAntiSpamRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportAntiSpamFalsePositiveRequest struct {
	Channel IInputChannel
	MsgID   int32
}

func (*ChannelsReportAntiSpamFalsePositiveRequest) CRC() uint32 {
	return 0xa850a693
}

func ChannelsReportAntiSpamFalsePositive(ctx context.Context, m Requester, i ChannelsReportAntiSpamFalsePositiveRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleParticipantsHiddenRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleParticipantsHiddenRequest) CRC() uint32 {
	return 0x6a6e7854
}

func ChannelsToggleParticipantsHidden(ctx context.Context, m Requester, i ChannelsToggleParticipantsHiddenRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsClickSponsoredMessageRequest struct {
	Channel  IInputChannel
	RandomID []byte
}

func (*ChannelsClickSponsoredMessageRequest) CRC() uint32 {
	return 0x18afbc93
}

func ChannelsClickSponsoredMessage(ctx context.Context, m Requester, i ChannelsClickSponsoredMessageRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateColorRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	ForProfile        bool     `tl:",omitempty:flags:1,implicit"`
	Channel           IInputChannel
	Color             *int32 `tl:",omitempty:flags:2"`
	BackgroundEmojiID *int64 `tl:",omitempty:flags:0"`
}

func (*ChannelsUpdateColorRequest) CRC() uint32 {
	return 0xd8aa3671
}

func ChannelsUpdateColor(ctx context.Context, m Requester, i ChannelsUpdateColorRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleViewForumAsMessagesRequest struct {
	Channel IInputChannel
	Enabled bool
}

func (*ChannelsToggleViewForumAsMessagesRequest) CRC() uint32 {
	return 0x9738bb15
}

func ChannelsToggleViewForumAsMessages(ctx context.Context, m Requester, i ChannelsToggleViewForumAsMessagesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetChannelRecommendationsRequest struct {
	Channel IInputChannel
}

func (*ChannelsGetChannelRecommendationsRequest) CRC() uint32 {
	return 0x83b70d97
}

func ChannelsGetChannelRecommendations(ctx context.Context, m Requester, i ChannelsGetChannelRecommendationsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateEmojiStatusRequest struct {
	Channel     IInputChannel
	EmojiStatus IEmojiStatus
}

func (*ChannelsUpdateEmojiStatusRequest) CRC() uint32 {
	return 0xf0d3e6a8
}

func ChannelsUpdateEmojiStatus(ctx context.Context, m Requester, i ChannelsUpdateEmojiStatusRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsExportChatlistInviteRequest struct {
	Chatlist IInputChatlist
	Title    string
	Peers    []IInputPeer
}

func (*ChatlistsExportChatlistInviteRequest) CRC() uint32 {
	return 0x8472478e
}

func ChatlistsExportChatlistInvite(ctx context.Context, m Requester, i ChatlistsExportChatlistInviteRequest) (IChatlistsExportedChatlistInvite, error) {
	var res IChatlistsExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsDeleteExportedInviteRequest struct {
	Chatlist IInputChatlist
	Slug     string
}

func (*ChatlistsDeleteExportedInviteRequest) CRC() uint32 {
	return 0x719c5c5e
}

func ChatlistsDeleteExportedInvite(ctx context.Context, m Requester, i ChatlistsDeleteExportedInviteRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsEditExportedInviteRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Chatlist IInputChatlist
	Slug     string
	Title    *string      `tl:",omitempty:flags:1"`
	Peers    []IInputPeer `tl:",omitempty:flags:2"`
}

func (*ChatlistsEditExportedInviteRequest) CRC() uint32 {
	return 0x653db63d
}

func ChatlistsEditExportedInvite(ctx context.Context, m Requester, i ChatlistsEditExportedInviteRequest) (IExportedChatlistInvite, error) {
	var res IExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetExportedInvitesRequest struct {
	Chatlist IInputChatlist
}

func (*ChatlistsGetExportedInvitesRequest) CRC() uint32 {
	return 0xce03da83
}

func ChatlistsGetExportedInvites(ctx context.Context, m Requester, i ChatlistsGetExportedInvitesRequest) (IChatlistsExportedInvites, error) {
	var res IChatlistsExportedInvites
	return res, request(ctx, m, &i, &res)
}

type ChatlistsCheckChatlistInviteRequest struct {
	Slug string
}

func (*ChatlistsCheckChatlistInviteRequest) CRC() uint32 {
	return 0x41c10fff
}

func ChatlistsCheckChatlistInvite(ctx context.Context, m Requester, i ChatlistsCheckChatlistInviteRequest) (IChatlistsChatlistInvite, error) {
	var res IChatlistsChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistInviteRequest struct {
	Slug  string
	Peers []IInputPeer
}

func (*ChatlistsJoinChatlistInviteRequest) CRC() uint32 {
	return 0xa6b1e39a
}

func ChatlistsJoinChatlistInvite(ctx context.Context, m Requester, i ChatlistsJoinChatlistInviteRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetChatlistUpdatesRequest struct {
	Chatlist IInputChatlist
}

func (*ChatlistsGetChatlistUpdatesRequest) CRC() uint32 {
	return 0x89419521
}

func ChatlistsGetChatlistUpdates(ctx context.Context, m Requester, i ChatlistsGetChatlistUpdatesRequest) (IChatlistsChatlistUpdates, error) {
	var res IChatlistsChatlistUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistUpdatesRequest struct {
	Chatlist IInputChatlist
	Peers    []IInputPeer
}

func (*ChatlistsJoinChatlistUpdatesRequest) CRC() uint32 {
	return 0xe089f8f5
}

func ChatlistsJoinChatlistUpdates(ctx context.Context, m Requester, i ChatlistsJoinChatlistUpdatesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsHideChatlistUpdatesRequest struct {
	Chatlist IInputChatlist
}

func (*ChatlistsHideChatlistUpdatesRequest) CRC() uint32 {
	return 0x66e486fb
}

func ChatlistsHideChatlistUpdates(ctx context.Context, m Requester, i ChatlistsHideChatlistUpdatesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetLeaveChatlistSuggestionsRequest struct {
	Chatlist IInputChatlist
}

func (*ChatlistsGetLeaveChatlistSuggestionsRequest) CRC() uint32 {
	return 0xfdbcd714
}

func ChatlistsGetLeaveChatlistSuggestions(ctx context.Context, m Requester, i ChatlistsGetLeaveChatlistSuggestionsRequest) ([]IPeer, error) {
	var res []IPeer
	return res, request(ctx, m, &i, &res)
}

type ChatlistsLeaveChatlistRequest struct {
	Chatlist IInputChatlist
	Peers    []IInputPeer
}

func (*ChatlistsLeaveChatlistRequest) CRC() uint32 {
	return 0x74fae13a
}

func ChatlistsLeaveChatlist(ctx context.Context, m Requester, i ChatlistsLeaveChatlistRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetContactIDsRequest struct {
	Hash int64
}

func (*ContactsGetContactIDsRequest) CRC() uint32 {
	return 0x7adc669d
}

func ContactsGetContactIDs(ctx context.Context, m Requester, i ContactsGetContactIDsRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type ContactsGetStatusesRequest struct{}

func (*ContactsGetStatusesRequest) CRC() uint32 {
	return 0xc4a353ee
}

func ContactsGetStatuses(ctx context.Context, m Requester, i ContactsGetStatusesRequest) ([]IContactStatus, error) {
	var res []IContactStatus
	return res, request(ctx, m, &i, &res)
}

type ContactsGetContactsRequest struct {
	Hash int64
}

func (*ContactsGetContactsRequest) CRC() uint32 {
	return 0x5dd69e12
}

func ContactsGetContacts(ctx context.Context, m Requester, i ContactsGetContactsRequest) (IContactsContacts, error) {
	var res IContactsContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsImportContactsRequest struct {
	Contacts []IInputContact
}

func (*ContactsImportContactsRequest) CRC() uint32 {
	return 0x2c800be5
}

func ContactsImportContacts(ctx context.Context, m Requester, i ContactsImportContactsRequest) (IContactsImportedContacts, error) {
	var res IContactsImportedContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteContactsRequest struct {
	ID []IInputUser
}

func (*ContactsDeleteContactsRequest) CRC() uint32 {
	return 0x96a0e00
}

func ContactsDeleteContacts(ctx context.Context, m Requester, i ContactsDeleteContactsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteByPhonesRequest struct {
	Phones []string
}

func (*ContactsDeleteByPhonesRequest) CRC() uint32 {
	return 0x1013fd9e
}

func ContactsDeleteByPhones(ctx context.Context, m Requester, i ContactsDeleteByPhonesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsBlockRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            IInputPeer
}

func (*ContactsBlockRequest) CRC() uint32 {
	return 0x2e2e8734
}

func ContactsBlock(ctx context.Context, m Requester, i ContactsBlockRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsUnblockRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            IInputPeer
}

func (*ContactsUnblockRequest) CRC() uint32 {
	return 0xb550d328
}

func ContactsUnblock(ctx context.Context, m Requester, i ContactsUnblockRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsGetBlockedRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	Offset        int32
	Limit         int32
}

func (*ContactsGetBlockedRequest) CRC() uint32 {
	return 0x9a868f80
}

func ContactsGetBlocked(ctx context.Context, m Requester, i ContactsGetBlockedRequest) (IContactsBlocked, error) {
	var res IContactsBlocked
	return res, request(ctx, m, &i, &res)
}

type ContactsSearchRequest struct {
	Q     string
	Limit int32
}

func (*ContactsSearchRequest) CRC() uint32 {
	return 0x11f812d8
}

func ContactsSearch(ctx context.Context, m Requester, i ContactsSearchRequest) (IContactsFound, error) {
	var res IContactsFound
	return res, request(ctx, m, &i, &res)
}

type ContactsResolveUsernameRequest struct {
	Username string
}

func (*ContactsResolveUsernameRequest) CRC() uint32 {
	return 0xf93ccba3
}

func ContactsResolveUsername(ctx context.Context, m Requester, i ContactsResolveUsernameRequest) (IContactsResolvedPeer, error) {
	var res IContactsResolvedPeer
	return res, request(ctx, m, &i, &res)
}

type ContactsGetTopPeersRequest struct {
	_              struct{} `tl:"flags,bitflag"`
	Correspondents bool     `tl:",omitempty:flags:0,implicit"`
	BotsPm         bool     `tl:",omitempty:flags:1,implicit"`
	BotsInline     bool     `tl:",omitempty:flags:2,implicit"`
	PhoneCalls     bool     `tl:",omitempty:flags:3,implicit"`
	ForwardUsers   bool     `tl:",omitempty:flags:4,implicit"`
	ForwardChats   bool     `tl:",omitempty:flags:5,implicit"`
	Groups         bool     `tl:",omitempty:flags:10,implicit"`
	Channels       bool     `tl:",omitempty:flags:15,implicit"`
	Offset         int32
	Limit          int32
	Hash           int64
}

func (*ContactsGetTopPeersRequest) CRC() uint32 {
	return 0x973478b6
}

func ContactsGetTopPeers(ctx context.Context, m Requester, i ContactsGetTopPeersRequest) (IContactsTopPeers, error) {
	var res IContactsTopPeers
	return res, request(ctx, m, &i, &res)
}

type ContactsResetTopPeerRatingRequest struct {
	Category EnumTopPeerCategory
	Peer     IInputPeer
}

func (*ContactsResetTopPeerRatingRequest) CRC() uint32 {
	return 0x1ae373ac
}

func ContactsResetTopPeerRating(ctx context.Context, m Requester, i ContactsResetTopPeerRatingRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsResetSavedRequest struct{}

func (*ContactsResetSavedRequest) CRC() uint32 {
	return 0x879537f1
}

func ContactsResetSaved(ctx context.Context, m Requester, i ContactsResetSavedRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsGetSavedRequest struct{}

func (*ContactsGetSavedRequest) CRC() uint32 {
	return 0x82f1e39f
}

func ContactsGetSaved(ctx context.Context, m Requester, i ContactsGetSavedRequest) ([]ISavedContact, error) {
	var res []ISavedContact
	return res, request(ctx, m, &i, &res)
}

type ContactsToggleTopPeersRequest struct {
	Enabled bool
}

func (*ContactsToggleTopPeersRequest) CRC() uint32 {
	return 0x8514bdda
}

func ContactsToggleTopPeers(ctx context.Context, m Requester, i ContactsToggleTopPeersRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsAddContactRequest struct {
	_                        struct{} `tl:"flags,bitflag"`
	AddPhonePrivacyException bool     `tl:",omitempty:flags:0,implicit"`
	ID                       IInputUser
	FirstName                string
	LastName                 string
	Phone                    string
}

func (*ContactsAddContactRequest) CRC() uint32 {
	return 0xe8f463d0
}

func ContactsAddContact(ctx context.Context, m Requester, i ContactsAddContactRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsAcceptContactRequest struct {
	ID IInputUser
}

func (*ContactsAcceptContactRequest) CRC() uint32 {
	return 0xf831a20f
}

func ContactsAcceptContact(ctx context.Context, m Requester, i ContactsAcceptContactRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetLocatedRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Background  bool     `tl:",omitempty:flags:1,implicit"`
	GeoPoint    IInputGeoPoint
	SelfExpires *int32 `tl:",omitempty:flags:0"`
}

func (*ContactsGetLocatedRequest) CRC() uint32 {
	return 0xd348bc44
}

func ContactsGetLocated(ctx context.Context, m Requester, i ContactsGetLocatedRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsBlockFromRepliesRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	DeleteMessage bool     `tl:",omitempty:flags:0,implicit"`
	DeleteHistory bool     `tl:",omitempty:flags:1,implicit"`
	ReportSpam    bool     `tl:",omitempty:flags:2,implicit"`
	MsgID         int32
}

func (*ContactsBlockFromRepliesRequest) CRC() uint32 {
	return 0x29a8962c
}

func ContactsBlockFromReplies(ctx context.Context, m Requester, i ContactsBlockFromRepliesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type ContactsResolvePhoneRequest struct {
	Phone string
}

func (*ContactsResolvePhoneRequest) CRC() uint32 {
	return 0x8af94344
}

func ContactsResolvePhone(ctx context.Context, m Requester, i ContactsResolvePhoneRequest) (IContactsResolvedPeer, error) {
	var res IContactsResolvedPeer
	return res, request(ctx, m, &i, &res)
}

type ContactsExportContactTokenRequest struct{}

func (*ContactsExportContactTokenRequest) CRC() uint32 {
	return 0xf8654027
}

func ContactsExportContactToken(ctx context.Context, m Requester, i ContactsExportContactTokenRequest) (IExportedContactToken, error) {
	var res IExportedContactToken
	return res, request(ctx, m, &i, &res)
}

type ContactsImportContactTokenRequest struct {
	Token string
}

func (*ContactsImportContactTokenRequest) CRC() uint32 {
	return 0x13005788
}

func ContactsImportContactToken(ctx context.Context, m Requester, i ContactsImportContactTokenRequest) (IUser, error) {
	var res IUser
	return res, request(ctx, m, &i, &res)
}

type ContactsEditCloseFriendsRequest struct {
	ID []int64
}

func (*ContactsEditCloseFriendsRequest) CRC() uint32 {
	return 0xba6705f0
}

func ContactsEditCloseFriends(ctx context.Context, m Requester, i ContactsEditCloseFriendsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsSetBlockedRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            []IInputPeer
	Limit         int32
}

func (*ContactsSetBlockedRequest) CRC() uint32 {
	return 0x94c65c76
}

func ContactsSetBlocked(ctx context.Context, m Requester, i ContactsSetBlockedRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type FoldersEditPeerFoldersRequest struct {
	FolderPeers []IInputFolderPeer
}

func (*FoldersEditPeerFoldersRequest) CRC() uint32 {
	return 0x6847d0ab
}

func FoldersEditPeerFolders(ctx context.Context, m Requester, i FoldersEditPeerFoldersRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type HelpGetConfigRequest struct{}

func (*HelpGetConfigRequest) CRC() uint32 {
	return 0xc4f9186b
}

func HelpGetConfig(ctx context.Context, m Requester, i HelpGetConfigRequest) (IConfig, error) {
	var res IConfig
	return res, request(ctx, m, &i, &res)
}

type HelpGetNearestDcRequest struct{}

func (*HelpGetNearestDcRequest) CRC() uint32 {
	return 0x1fb33026
}

func HelpGetNearestDc(ctx context.Context, m Requester, i HelpGetNearestDcRequest) (INearestDc, error) {
	var res INearestDc
	return res, request(ctx, m, &i, &res)
}

type HelpGetAppUpdateRequest struct {
	Source string
}

func (*HelpGetAppUpdateRequest) CRC() uint32 {
	return 0x522d5a7d
}

func HelpGetAppUpdate(ctx context.Context, m Requester, i HelpGetAppUpdateRequest) (IHelpAppUpdate, error) {
	var res IHelpAppUpdate
	return res, request(ctx, m, &i, &res)
}

type HelpGetInviteTextRequest struct{}

func (*HelpGetInviteTextRequest) CRC() uint32 {
	return 0x4d392343
}

func HelpGetInviteText(ctx context.Context, m Requester, i HelpGetInviteTextRequest) (IHelpInviteText, error) {
	var res IHelpInviteText
	return res, request(ctx, m, &i, &res)
}

type HelpGetSupportRequest struct{}

func (*HelpGetSupportRequest) CRC() uint32 {
	return 0x9cdf08cd
}

func HelpGetSupport(ctx context.Context, m Requester, i HelpGetSupportRequest) (IHelpSupport, error) {
	var res IHelpSupport
	return res, request(ctx, m, &i, &res)
}

type HelpSetBotUpdatesStatusRequest struct {
	PendingUpdatesCount int32
	Message             string
}

func (*HelpSetBotUpdatesStatusRequest) CRC() uint32 {
	return 0xec22cfcd
}

func HelpSetBotUpdatesStatus(ctx context.Context, m Requester, i HelpSetBotUpdatesStatusRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetCdnConfigRequest struct{}

func (*HelpGetCdnConfigRequest) CRC() uint32 {
	return 0x52029342
}

func HelpGetCdnConfig(ctx context.Context, m Requester, i HelpGetCdnConfigRequest) (ICdnConfig, error) {
	var res ICdnConfig
	return res, request(ctx, m, &i, &res)
}

type HelpGetRecentMeUrlsRequest struct {
	Referer string
}

func (*HelpGetRecentMeUrlsRequest) CRC() uint32 {
	return 0x3dc0f114
}

func HelpGetRecentMeUrls(ctx context.Context, m Requester, i HelpGetRecentMeUrlsRequest) (IHelpRecentMeUrls, error) {
	var res IHelpRecentMeUrls
	return res, request(ctx, m, &i, &res)
}

type HelpGetTermsOfServiceUpdateRequest struct{}

func (*HelpGetTermsOfServiceUpdateRequest) CRC() uint32 {
	return 0x2ca51fd1
}

func HelpGetTermsOfServiceUpdate(ctx context.Context, m Requester, i HelpGetTermsOfServiceUpdateRequest) (IHelpTermsOfServiceUpdate, error) {
	var res IHelpTermsOfServiceUpdate
	return res, request(ctx, m, &i, &res)
}

type HelpAcceptTermsOfServiceRequest struct {
	ID IDataJSON
}

func (*HelpAcceptTermsOfServiceRequest) CRC() uint32 {
	return 0xee72f79a
}

func HelpAcceptTermsOfService(ctx context.Context, m Requester, i HelpAcceptTermsOfServiceRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetDeepLinkInfoRequest struct {
	Path string
}

func (*HelpGetDeepLinkInfoRequest) CRC() uint32 {
	return 0x3fedc75f
}

func HelpGetDeepLinkInfo(ctx context.Context, m Requester, i HelpGetDeepLinkInfoRequest) (IHelpDeepLinkInfo, error) {
	var res IHelpDeepLinkInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetAppConfigRequest struct {
	Hash int32
}

func (*HelpGetAppConfigRequest) CRC() uint32 {
	return 0x61e3f854
}

func HelpGetAppConfig(ctx context.Context, m Requester, i HelpGetAppConfigRequest) (IHelpAppConfig, error) {
	var res IHelpAppConfig
	return res, request(ctx, m, &i, &res)
}

type HelpSaveAppLogRequest struct {
	Events []IInputAppEvent
}

func (*HelpSaveAppLogRequest) CRC() uint32 {
	return 0x6f02f748
}

func HelpSaveAppLog(ctx context.Context, m Requester, i HelpSaveAppLogRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetPassportConfigRequest struct {
	Hash int32
}

func (*HelpGetPassportConfigRequest) CRC() uint32 {
	return 0xc661ad08
}

func HelpGetPassportConfig(ctx context.Context, m Requester, i HelpGetPassportConfigRequest) (IHelpPassportConfig, error) {
	var res IHelpPassportConfig
	return res, request(ctx, m, &i, &res)
}

type HelpGetSupportNameRequest struct{}

func (*HelpGetSupportNameRequest) CRC() uint32 {
	return 0xd360e72c
}

func HelpGetSupportName(ctx context.Context, m Requester, i HelpGetSupportNameRequest) (IHelpSupportName, error) {
	var res IHelpSupportName
	return res, request(ctx, m, &i, &res)
}

type HelpGetUserInfoRequest struct {
	UserID IInputUser
}

func (*HelpGetUserInfoRequest) CRC() uint32 {
	return 0x38a08d3
}

func HelpGetUserInfo(ctx context.Context, m Requester, i HelpGetUserInfoRequest) (IHelpUserInfo, error) {
	var res IHelpUserInfo
	return res, request(ctx, m, &i, &res)
}

type HelpEditUserInfoRequest struct {
	UserID   IInputUser
	Message  string
	Entities []IMessageEntity
}

func (*HelpEditUserInfoRequest) CRC() uint32 {
	return 0x66b91b70
}

func HelpEditUserInfo(ctx context.Context, m Requester, i HelpEditUserInfoRequest) (IHelpUserInfo, error) {
	var res IHelpUserInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetPromoDataRequest struct{}

func (*HelpGetPromoDataRequest) CRC() uint32 {
	return 0xc0977421
}

func HelpGetPromoData(ctx context.Context, m Requester, i HelpGetPromoDataRequest) (IHelpPromoData, error) {
	var res IHelpPromoData
	return res, request(ctx, m, &i, &res)
}

type HelpHidePromoDataRequest struct {
	Peer IInputPeer
}

func (*HelpHidePromoDataRequest) CRC() uint32 {
	return 0x1e251c95
}

func HelpHidePromoData(ctx context.Context, m Requester, i HelpHidePromoDataRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpDismissSuggestionRequest struct {
	Peer       IInputPeer
	Suggestion string
}

func (*HelpDismissSuggestionRequest) CRC() uint32 {
	return 0xf50dbaa1
}

func HelpDismissSuggestion(ctx context.Context, m Requester, i HelpDismissSuggestionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetCountriesListRequest struct {
	LangCode string
	Hash     int32
}

func (*HelpGetCountriesListRequest) CRC() uint32 {
	return 0x735787a8
}

func HelpGetCountriesList(ctx context.Context, m Requester, i HelpGetCountriesListRequest) (IHelpCountriesList, error) {
	var res IHelpCountriesList
	return res, request(ctx, m, &i, &res)
}

type HelpGetPremiumPromoRequest struct{}

func (*HelpGetPremiumPromoRequest) CRC() uint32 {
	return 0xb81b93d4
}

func HelpGetPremiumPromo(ctx context.Context, m Requester, i HelpGetPremiumPromoRequest) (IHelpPremiumPromo, error) {
	var res IHelpPremiumPromo
	return res, request(ctx, m, &i, &res)
}

type HelpGetPeerColorsRequest struct {
	Hash int32
}

func (*HelpGetPeerColorsRequest) CRC() uint32 {
	return 0xda80f42f
}

func HelpGetPeerColors(ctx context.Context, m Requester, i HelpGetPeerColorsRequest) (IHelpPeerColors, error) {
	var res IHelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type HelpGetPeerProfileColorsRequest struct {
	Hash int32
}

func (*HelpGetPeerProfileColorsRequest) CRC() uint32 {
	return 0xabcfa9fd
}

func HelpGetPeerProfileColors(ctx context.Context, m Requester, i HelpGetPeerProfileColorsRequest) (IHelpPeerColors, error) {
	var res IHelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLangPackRequest struct {
	LangPack string
	LangCode string
}

func (*LangpackGetLangPackRequest) CRC() uint32 {
	return 0xf2f2330a
}

func LangpackGetLangPack(ctx context.Context, m Requester, i LangpackGetLangPackRequest) (ILangPackDifference, error) {
	var res ILangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetStringsRequest struct {
	LangPack string
	LangCode string
	Keys     []string
}

func (*LangpackGetStringsRequest) CRC() uint32 {
	return 0xefea3803
}

func LangpackGetStrings(ctx context.Context, m Requester, i LangpackGetStringsRequest) ([]ILangPackString, error) {
	var res []ILangPackString
	return res, request(ctx, m, &i, &res)
}

type LangpackGetDifferenceRequest struct {
	LangPack    string
	LangCode    string
	FromVersion int32
}

func (*LangpackGetDifferenceRequest) CRC() uint32 {
	return 0xcd984aa5
}

func LangpackGetDifference(ctx context.Context, m Requester, i LangpackGetDifferenceRequest) (ILangPackDifference, error) {
	var res ILangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguagesRequest struct {
	LangPack string
}

func (*LangpackGetLanguagesRequest) CRC() uint32 {
	return 0x42c6978f
}

func LangpackGetLanguages(ctx context.Context, m Requester, i LangpackGetLanguagesRequest) ([]ILangPackLanguage, error) {
	var res []ILangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguageRequest struct {
	LangPack string
	LangCode string
}

func (*LangpackGetLanguageRequest) CRC() uint32 {
	return 0x6a596502
}

func LangpackGetLanguage(ctx context.Context, m Requester, i LangpackGetLanguageRequest) (ILangPackLanguage, error) {
	var res ILangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesRequest struct {
	ID []IInputMessage
}

func (*MessagesGetMessagesRequest) CRC() uint32 {
	return 0x63c66506
}

func MessagesGetMessages(ctx context.Context, m Requester, i MessagesGetMessagesRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogsRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	ExcludePinned bool     `tl:",omitempty:flags:0,implicit"`
	FolderID      *int32   `tl:",omitempty:flags:1"`
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    IInputPeer
	Limit         int32
	Hash          int64
}

func (*MessagesGetDialogsRequest) CRC() uint32 {
	return 0xa0f4cb4f
}

func MessagesGetDialogs(ctx context.Context, m Requester, i MessagesGetDialogsRequest) (IMessagesDialogs, error) {
	var res IMessagesDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesGetHistoryRequest struct {
	Peer       IInputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int64
}

func (*MessagesGetHistoryRequest) CRC() uint32 {
	return 0x4423e6c5
}

func MessagesGetHistory(ctx context.Context, m Requester, i MessagesGetHistoryRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IInputPeer
	Q         string
	FromID    IInputPeer `tl:",omitempty:flags:0"`
	TopMsgID  *int32     `tl:",omitempty:flags:1"`
	Filter    IMessagesFilter
	MinDate   int32
	MaxDate   int32
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
	Hash      int64
}

func (*MessagesSearchRequest) CRC() uint32 {
	return 0xa0fda762
}

func MessagesSearch(ctx context.Context, m Requester, i MessagesSearchRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadHistoryRequest struct {
	Peer  IInputPeer
	MaxID int32
}

func (*MessagesReadHistoryRequest) CRC() uint32 {
	return 0xe306d3a
}

func MessagesReadHistory(ctx context.Context, m Requester, i MessagesReadHistoryRequest) (IMessagesAffectedMessages, error) {
	var res IMessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteHistoryRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	JustClear bool     `tl:",omitempty:flags:0,implicit"`
	Revoke    bool     `tl:",omitempty:flags:1,implicit"`
	Peer      IInputPeer
	MaxID     int32
	MinDate   *int32 `tl:",omitempty:flags:2"`
	MaxDate   *int32 `tl:",omitempty:flags:3"`
}

func (*MessagesDeleteHistoryRequest) CRC() uint32 {
	return 0xb08f922a
}

func MessagesDeleteHistory(ctx context.Context, m Requester, i MessagesDeleteHistoryRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteMessagesRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Revoke bool     `tl:",omitempty:flags:0,implicit"`
	ID     []int32
}

func (*MessagesDeleteMessagesRequest) CRC() uint32 {
	return 0xe58e95d2
}

func MessagesDeleteMessages(ctx context.Context, m Requester, i MessagesDeleteMessagesRequest) (IMessagesAffectedMessages, error) {
	var res IMessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedMessagesRequest struct {
	MaxID int32
}

func (*MessagesReceivedMessagesRequest) CRC() uint32 {
	return 0x5a954c0
}

func MessagesReceivedMessages(ctx context.Context, m Requester, i MessagesReceivedMessagesRequest) ([]IReceivedNotifyMessage, error) {
	var res []IReceivedNotifyMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSetTypingRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
	Action   ISendMessageAction
}

func (*MessagesSetTypingRequest) CRC() uint32 {
	return 0x58943ee2
}

func MessagesSetTyping(ctx context.Context, m Requester, i MessagesSetTypingRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMessageRequest struct {
	_                      struct{} `tl:"flags,bitflag"`
	NoWebpage              bool     `tl:",omitempty:flags:1,implicit"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   IInputPeer
	ReplyTo                IInputReplyTo `tl:",omitempty:flags:0"`
	Message                string
	RandomID               int64
	ReplyMarkup            IReplyMarkup     `tl:",omitempty:flags:2"`
	Entities               []IMessageEntity `tl:",omitempty:flags:3"`
	ScheduleDate           *int32           `tl:",omitempty:flags:10"`
	SendAs                 IInputPeer       `tl:",omitempty:flags:13"`
}

func (*MessagesSendMessageRequest) CRC() uint32 {
	return 0x280d096f
}

func MessagesSendMessage(ctx context.Context, m Requester, i MessagesSendMessageRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMediaRequest struct {
	_                      struct{} `tl:"flags,bitflag"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   IInputPeer
	ReplyTo                IInputReplyTo `tl:",omitempty:flags:0"`
	Media                  IInputMedia
	Message                string
	RandomID               int64
	ReplyMarkup            IReplyMarkup     `tl:",omitempty:flags:2"`
	Entities               []IMessageEntity `tl:",omitempty:flags:3"`
	ScheduleDate           *int32           `tl:",omitempty:flags:10"`
	SendAs                 IInputPeer       `tl:",omitempty:flags:13"`
}

func (*MessagesSendMediaRequest) CRC() uint32 {
	return 0x72ccc23d
}

func MessagesSendMedia(ctx context.Context, m Requester, i MessagesSendMediaRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesForwardMessagesRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	Silent            bool     `tl:",omitempty:flags:5,implicit"`
	Background        bool     `tl:",omitempty:flags:6,implicit"`
	WithMyScore       bool     `tl:",omitempty:flags:8,implicit"`
	DropAuthor        bool     `tl:",omitempty:flags:11,implicit"`
	DropMediaCaptions bool     `tl:",omitempty:flags:12,implicit"`
	Noforwards        bool     `tl:",omitempty:flags:14,implicit"`
	FromPeer          IInputPeer
	ID                []int32
	RandomID          []int64
	ToPeer            IInputPeer
	TopMsgID          *int32     `tl:",omitempty:flags:9"`
	ScheduleDate      *int32     `tl:",omitempty:flags:10"`
	SendAs            IInputPeer `tl:",omitempty:flags:13"`
}

func (*MessagesForwardMessagesRequest) CRC() uint32 {
	return 0xc661bbc4
}

func MessagesForwardMessages(ctx context.Context, m Requester, i MessagesForwardMessagesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesReportSpamRequest struct {
	Peer IInputPeer
}

func (*MessagesReportSpamRequest) CRC() uint32 {
	return 0xcf1592db
}

func MessagesReportSpam(ctx context.Context, m Requester, i MessagesReportSpamRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerSettingsRequest struct {
	Peer IInputPeer
}

func (*MessagesGetPeerSettingsRequest) CRC() uint32 {
	return 0xefd9a6a2
}

func MessagesGetPeerSettings(ctx context.Context, m Requester, i MessagesGetPeerSettingsRequest) (IMessagesPeerSettings, error) {
	var res IMessagesPeerSettings
	return res, request(ctx, m, &i, &res)
}

type MessagesReportRequest struct {
	Peer    IInputPeer
	ID      []int32
	Reason  EnumReportReason
	Message string
}

func (*MessagesReportRequest) CRC() uint32 {
	return 0x8953ab4e
}

func MessagesReport(ctx context.Context, m Requester, i MessagesReportRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatsRequest struct {
	ID []int64
}

func (*MessagesGetChatsRequest) CRC() uint32 {
	return 0x49e9528f
}

func MessagesGetChats(ctx context.Context, m Requester, i MessagesGetChatsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFullChatRequest struct {
	ChatID int64
}

func (*MessagesGetFullChatRequest) CRC() uint32 {
	return 0xaeb00b34
}

func MessagesGetFullChat(ctx context.Context, m Requester, i MessagesGetFullChatRequest) (IMessagesChatFull, error) {
	var res IMessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatTitleRequest struct {
	ChatID int64
	Title  string
}

func (*MessagesEditChatTitleRequest) CRC() uint32 {
	return 0x73783ffd
}

func MessagesEditChatTitle(ctx context.Context, m Requester, i MessagesEditChatTitleRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatPhotoRequest struct {
	ChatID int64
	Photo  IInputChatPhoto
}

func (*MessagesEditChatPhotoRequest) CRC() uint32 {
	return 0x35ddd674
}

func MessagesEditChatPhoto(ctx context.Context, m Requester, i MessagesEditChatPhotoRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesAddChatUserRequest struct {
	ChatID   int64
	UserID   IInputUser
	FwdLimit int32
}

func (*MessagesAddChatUserRequest) CRC() uint32 {
	return 0xf24753e3
}

func MessagesAddChatUser(ctx context.Context, m Requester, i MessagesAddChatUserRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatUserRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	RevokeHistory bool     `tl:",omitempty:flags:0,implicit"`
	ChatID        int64
	UserID        IInputUser
}

func (*MessagesDeleteChatUserRequest) CRC() uint32 {
	return 0xa2185cab
}

func MessagesDeleteChatUser(ctx context.Context, m Requester, i MessagesDeleteChatUserRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesCreateChatRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Users     []IInputUser
	Title     string
	TTLPeriod *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesCreateChatRequest) CRC() uint32 {
	return 0x34a818
}

func MessagesCreateChat(ctx context.Context, m Requester, i MessagesCreateChatRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDhConfigRequest struct {
	Version      int32
	RandomLength int32
}

func (*MessagesGetDhConfigRequest) CRC() uint32 {
	return 0x26cf8950
}

func MessagesGetDhConfig(ctx context.Context, m Requester, i MessagesGetDhConfigRequest) (IMessagesDhConfig, error) {
	var res IMessagesDhConfig
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestEncryptionRequest struct {
	UserID   IInputUser
	RandomID int32
	GA       []byte
}

func (*MessagesRequestEncryptionRequest) CRC() uint32 {
	return 0xf64daf43
}

func MessagesRequestEncryption(ctx context.Context, m Requester, i MessagesRequestEncryptionRequest) (IEncryptedChat, error) {
	var res IEncryptedChat
	return res, request(ctx, m, &i, &res)
}

type MessagesAcceptEncryptionRequest struct {
	Peer           IInputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

func (*MessagesAcceptEncryptionRequest) CRC() uint32 {
	return 0x3dbc0415
}

func MessagesAcceptEncryption(ctx context.Context, m Requester, i MessagesAcceptEncryptionRequest) (IEncryptedChat, error) {
	var res IEncryptedChat
	return res, request(ctx, m, &i, &res)
}

type MessagesDiscardEncryptionRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	DeleteHistory bool     `tl:",omitempty:flags:0,implicit"`
	ChatID        int32
}

func (*MessagesDiscardEncryptionRequest) CRC() uint32 {
	return 0xf393aea0
}

func MessagesDiscardEncryption(ctx context.Context, m Requester, i MessagesDiscardEncryptionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSetEncryptedTypingRequest struct {
	Peer   IInputEncryptedChat
	Typing bool
}

func (*MessagesSetEncryptedTypingRequest) CRC() uint32 {
	return 0x791451ed
}

func MessagesSetEncryptedTyping(ctx context.Context, m Requester, i MessagesSetEncryptedTypingRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadEncryptedHistoryRequest struct {
	Peer    IInputEncryptedChat
	MaxDate int32
}

func (*MessagesReadEncryptedHistoryRequest) CRC() uint32 {
	return 0x7f4b690a
}

func MessagesReadEncryptedHistory(ctx context.Context, m Requester, i MessagesReadEncryptedHistoryRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Silent   bool     `tl:",omitempty:flags:0,implicit"`
	Peer     IInputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedRequest) CRC() uint32 {
	return 0x44fa7a15
}

func MessagesSendEncrypted(ctx context.Context, m Requester, i MessagesSendEncryptedRequest) (IMessagesSentEncryptedMessage, error) {
	var res IMessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedFileRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Silent   bool     `tl:",omitempty:flags:0,implicit"`
	Peer     IInputEncryptedChat
	RandomID int64
	Data     []byte
	File     IInputEncryptedFile
}

func (*MessagesSendEncryptedFileRequest) CRC() uint32 {
	return 0x5559481d
}

func MessagesSendEncryptedFile(ctx context.Context, m Requester, i MessagesSendEncryptedFileRequest) (IMessagesSentEncryptedMessage, error) {
	var res IMessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedServiceRequest struct {
	Peer     IInputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedServiceRequest) CRC() uint32 {
	return 0x32d439a4
}

func MessagesSendEncryptedService(ctx context.Context, m Requester, i MessagesSendEncryptedServiceRequest) (IMessagesSentEncryptedMessage, error) {
	var res IMessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedQueueRequest struct {
	MaxQts int32
}

func (*MessagesReceivedQueueRequest) CRC() uint32 {
	return 0x55a5bb66
}

func MessagesReceivedQueue(ctx context.Context, m Requester, i MessagesReceivedQueueRequest) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type MessagesReportEncryptedSpamRequest struct {
	Peer IInputEncryptedChat
}

func (*MessagesReportEncryptedSpamRequest) CRC() uint32 {
	return 0x4b0c8c0f
}

func MessagesReportEncryptedSpam(ctx context.Context, m Requester, i MessagesReportEncryptedSpamRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMessageContentsRequest struct {
	ID []int32
}

func (*MessagesReadMessageContentsRequest) CRC() uint32 {
	return 0x36a73f77
}

func MessagesReadMessageContents(ctx context.Context, m Requester, i MessagesReadMessageContentsRequest) (IMessagesAffectedMessages, error) {
	var res IMessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickersRequest struct {
	Emoticon string
	Hash     int64
}

func (*MessagesGetStickersRequest) CRC() uint32 {
	return 0xd5a5d3a1
}

func MessagesGetStickers(ctx context.Context, m Requester, i MessagesGetStickersRequest) (IMessagesStickers, error) {
	var res IMessagesStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAllStickersRequest struct {
	Hash int64
}

func (*MessagesGetAllStickersRequest) CRC() uint32 {
	return 0xb8a0a1a8
}

func MessagesGetAllStickers(ctx context.Context, m Requester, i MessagesGetAllStickersRequest) (IMessagesAllStickers, error) {
	var res IMessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPagePreviewRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Message  string
	Entities []IMessageEntity `tl:",omitempty:flags:3"`
}

func (*MessagesGetWebPagePreviewRequest) CRC() uint32 {
	return 0x8b68b0cc
}

func MessagesGetWebPagePreview(ctx context.Context, m Requester, i MessagesGetWebPagePreviewRequest) (IMessageMedia, error) {
	var res IMessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesExportChatInviteRequest struct {
	_                     struct{} `tl:"flags,bitflag"`
	LegacyRevokePermanent bool     `tl:",omitempty:flags:2,implicit"`
	RequestNeeded         bool     `tl:",omitempty:flags:3,implicit"`
	Peer                  IInputPeer
	ExpireDate            *int32  `tl:",omitempty:flags:0"`
	UsageLimit            *int32  `tl:",omitempty:flags:1"`
	Title                 *string `tl:",omitempty:flags:4"`
}

func (*MessagesExportChatInviteRequest) CRC() uint32 {
	return 0xa02ce5d5
}

func MessagesExportChatInvite(ctx context.Context, m Requester, i MessagesExportChatInviteRequest) (IExportedChatInvite, error) {
	var res IExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckChatInviteRequest struct {
	Hash string
}

func (*MessagesCheckChatInviteRequest) CRC() uint32 {
	return 0x3eadb1bb
}

func MessagesCheckChatInvite(ctx context.Context, m Requester, i MessagesCheckChatInviteRequest) (IChatInvite, error) {
	var res IChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesImportChatInviteRequest struct {
	Hash string
}

func (*MessagesImportChatInviteRequest) CRC() uint32 {
	return 0x6c50051c
}

func MessagesImportChatInvite(ctx context.Context, m Requester, i MessagesImportChatInviteRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickerSetRequest struct {
	Stickerset IInputStickerSet
	Hash       int32
}

func (*MessagesGetStickerSetRequest) CRC() uint32 {
	return 0xc8a0ec74
}

func MessagesGetStickerSet(ctx context.Context, m Requester, i MessagesGetStickerSetRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type MessagesInstallStickerSetRequest struct {
	Stickerset IInputStickerSet
	Archived   bool
}

func (*MessagesInstallStickerSetRequest) CRC() uint32 {
	return 0xc78fe460
}

func MessagesInstallStickerSet(ctx context.Context, m Requester, i MessagesInstallStickerSetRequest) (IMessagesStickerSetInstallResult, error) {
	var res IMessagesStickerSetInstallResult
	return res, request(ctx, m, &i, &res)
}

type MessagesUninstallStickerSetRequest struct {
	Stickerset IInputStickerSet
}

func (*MessagesUninstallStickerSetRequest) CRC() uint32 {
	return 0xf96e55de
}

func MessagesUninstallStickerSet(ctx context.Context, m Requester, i MessagesUninstallStickerSetRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesStartBotRequest struct {
	Bot        IInputUser
	Peer       IInputPeer
	RandomID   int64
	StartParam string
}

func (*MessagesStartBotRequest) CRC() uint32 {
	return 0xe6df7378
}

func MessagesStartBot(ctx context.Context, m Requester, i MessagesStartBotRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesViewsRequest struct {
	Peer      IInputPeer
	ID        []int32
	Increment bool
}

func (*MessagesGetMessagesViewsRequest) CRC() uint32 {
	return 0x5784d3e1
}

func MessagesGetMessagesViews(ctx context.Context, m Requester, i MessagesGetMessagesViewsRequest) (IMessagesMessageViews, error) {
	var res IMessagesMessageViews
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAdminRequest struct {
	ChatID  int64
	UserID  IInputUser
	IsAdmin bool
}

func (*MessagesEditChatAdminRequest) CRC() uint32 {
	return 0xa85bd1c2
}

func MessagesEditChatAdmin(ctx context.Context, m Requester, i MessagesEditChatAdminRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesMigrateChatRequest struct {
	ChatID int64
}

func (*MessagesMigrateChatRequest) CRC() uint32 {
	return 0xa2875319
}

func MessagesMigrateChat(ctx context.Context, m Requester, i MessagesMigrateChatRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchGlobalRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	FolderID   *int32   `tl:",omitempty:flags:0"`
	Q          string
	Filter     IMessagesFilter
	MinDate    int32
	MaxDate    int32
	OffsetRate int32
	OffsetPeer IInputPeer
	OffsetID   int32
	Limit      int32
}

func (*MessagesSearchGlobalRequest) CRC() uint32 {
	return 0x4bc6589a
}

func MessagesSearchGlobal(ctx context.Context, m Requester, i MessagesSearchGlobalRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderStickerSetsRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
	Order  []int64
}

func (*MessagesReorderStickerSetsRequest) CRC() uint32 {
	return 0x78337739
}

func MessagesReorderStickerSets(ctx context.Context, m Requester, i MessagesReorderStickerSetsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDocumentByHashRequest struct {
	SHA256   []byte
	Size     int64
	MimeType string
}

func (*MessagesGetDocumentByHashRequest) CRC() uint32 {
	return 0xb1f2061f
}

func MessagesGetDocumentByHash(ctx context.Context, m Requester, i MessagesGetDocumentByHashRequest) (IDocument, error) {
	var res IDocument
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedGifsRequest struct {
	Hash int64
}

func (*MessagesGetSavedGifsRequest) CRC() uint32 {
	return 0x5cf09635
}

func MessagesGetSavedGifs(ctx context.Context, m Requester, i MessagesGetSavedGifsRequest) (IMessagesSavedGifs, error) {
	var res IMessagesSavedGifs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveGifRequest struct {
	ID     IInputDocument
	Unsave bool
}

func (*MessagesSaveGifRequest) CRC() uint32 {
	return 0x327a30cb
}

func MessagesSaveGif(ctx context.Context, m Requester, i MessagesSaveGifRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineBotResultsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Bot      IInputUser
	Peer     IInputPeer
	GeoPoint IInputGeoPoint `tl:",omitempty:flags:0"`
	Query    string
	Offset   string
}

func (*MessagesGetInlineBotResultsRequest) CRC() uint32 {
	return 0x514e999d
}

func MessagesGetInlineBotResults(ctx context.Context, m Requester, i MessagesGetInlineBotResultsRequest) (IMessagesBotResults, error) {
	var res IMessagesBotResults
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineBotResultsRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	Gallery       bool     `tl:",omitempty:flags:0,implicit"`
	Private       bool     `tl:",omitempty:flags:1,implicit"`
	QueryID       int64
	Results       []IInputBotInlineResult
	CacheTime     int32
	NextOffset    *string            `tl:",omitempty:flags:2"`
	SwitchPm      IInlineBotSwitchPm `tl:",omitempty:flags:3"`
	SwitchWebview IInlineBotWebView  `tl:",omitempty:flags:4"`
}

func (*MessagesSetInlineBotResultsRequest) CRC() uint32 {
	return 0xbb12a419
}

func MessagesSetInlineBotResults(ctx context.Context, m Requester, i MessagesSetInlineBotResultsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendInlineBotResultRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Silent       bool     `tl:",omitempty:flags:5,implicit"`
	Background   bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft   bool     `tl:",omitempty:flags:7,implicit"`
	HideVia      bool     `tl:",omitempty:flags:11,implicit"`
	Peer         IInputPeer
	ReplyTo      IInputReplyTo `tl:",omitempty:flags:0"`
	RandomID     int64
	QueryID      int64
	ID           string
	ScheduleDate *int32     `tl:",omitempty:flags:10"`
	SendAs       IInputPeer `tl:",omitempty:flags:13"`
}

func (*MessagesSendInlineBotResultRequest) CRC() uint32 {
	return 0xf7bc68ba
}

func MessagesSendInlineBotResult(ctx context.Context, m Requester, i MessagesSendInlineBotResultRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageEditDataRequest struct {
	Peer IInputPeer
	ID   int32
}

func (*MessagesGetMessageEditDataRequest) CRC() uint32 {
	return 0xfda68d36
}

func MessagesGetMessageEditData(ctx context.Context, m Requester, i MessagesGetMessageEditDataRequest) (IMessagesMessageEditData, error) {
	var res IMessagesMessageEditData
	return res, request(ctx, m, &i, &res)
}

type MessagesEditMessageRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	NoWebpage    bool     `tl:",omitempty:flags:1,implicit"`
	InvertMedia  bool     `tl:",omitempty:flags:16,implicit"`
	Peer         IInputPeer
	ID           int32
	Message      *string          `tl:",omitempty:flags:11"`
	Media        IInputMedia      `tl:",omitempty:flags:14"`
	ReplyMarkup  IReplyMarkup     `tl:",omitempty:flags:2"`
	Entities     []IMessageEntity `tl:",omitempty:flags:3"`
	ScheduleDate *int32           `tl:",omitempty:flags:15"`
}

func (*MessagesEditMessageRequest) CRC() uint32 {
	return 0x48f71778
}

func MessagesEditMessage(ctx context.Context, m Requester, i MessagesEditMessageRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditInlineBotMessageRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:16,implicit"`
	ID          IInputBotInlineMessageID
	Message     *string          `tl:",omitempty:flags:11"`
	Media       IInputMedia      `tl:",omitempty:flags:14"`
	ReplyMarkup IReplyMarkup     `tl:",omitempty:flags:2"`
	Entities    []IMessageEntity `tl:",omitempty:flags:3"`
}

func (*MessagesEditInlineBotMessageRequest) CRC() uint32 {
	return 0x83557dba
}

func MessagesEditInlineBotMessage(ctx context.Context, m Requester, i MessagesEditInlineBotMessageRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotCallbackAnswerRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Game     bool     `tl:",omitempty:flags:1,implicit"`
	Peer     IInputPeer
	MsgID    int32
	Data     *[]byte                `tl:",omitempty:flags:0"`
	Password IInputCheckPasswordSRP `tl:",omitempty:flags:2"`
}

func (*MessagesGetBotCallbackAnswerRequest) CRC() uint32 {
	return 0x9342ca07
}

func MessagesGetBotCallbackAnswer(ctx context.Context, m Requester, i MessagesGetBotCallbackAnswerRequest) (IMessagesBotCallbackAnswer, error) {
	var res IMessagesBotCallbackAnswer
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotCallbackAnswerRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Alert     bool     `tl:",omitempty:flags:1,implicit"`
	QueryID   int64
	Message   *string `tl:",omitempty:flags:0"`
	URL       *string `tl:",omitempty:flags:2"`
	CacheTime int32
}

func (*MessagesSetBotCallbackAnswerRequest) CRC() uint32 {
	return 0xd58f130a
}

func MessagesSetBotCallbackAnswer(ctx context.Context, m Requester, i MessagesSetBotCallbackAnswerRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerDialogsRequest struct {
	Peers []IInputDialogPeer
}

func (*MessagesGetPeerDialogsRequest) CRC() uint32 {
	return 0xe470bcfd
}

func MessagesGetPeerDialogs(ctx context.Context, m Requester, i MessagesGetPeerDialogsRequest) (IMessagesPeerDialogs, error) {
	var res IMessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDraftRequest struct {
	_           struct{}      `tl:"flags,bitflag"`
	NoWebpage   bool          `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool          `tl:",omitempty:flags:6,implicit"`
	ReplyTo     IInputReplyTo `tl:",omitempty:flags:4"`
	Peer        IInputPeer
	Message     string
	Entities    []IMessageEntity `tl:",omitempty:flags:3"`
	Media       IInputMedia      `tl:",omitempty:flags:5"`
}

func (*MessagesSaveDraftRequest) CRC() uint32 {
	return 0x7ff3b806
}

func MessagesSaveDraft(ctx context.Context, m Requester, i MessagesSaveDraftRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAllDraftsRequest struct{}

func (*MessagesGetAllDraftsRequest) CRC() uint32 {
	return 0x6a3f8d65
}

func MessagesGetAllDrafts(ctx context.Context, m Requester, i MessagesGetAllDraftsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFeaturedStickersRequest struct {
	Hash int64
}

func (*MessagesGetFeaturedStickersRequest) CRC() uint32 {
	return 0x64780b14
}

func MessagesGetFeaturedStickers(ctx context.Context, m Requester, i MessagesGetFeaturedStickersRequest) (IMessagesFeaturedStickers, error) {
	var res IMessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReadFeaturedStickersRequest struct {
	ID []int64
}

func (*MessagesReadFeaturedStickersRequest) CRC() uint32 {
	return 0x5b118126
}

func MessagesReadFeaturedStickers(ctx context.Context, m Requester, i MessagesReadFeaturedStickersRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentStickersRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
	Hash     int64
}

func (*MessagesGetRecentStickersRequest) CRC() uint32 {
	return 0x9da9403b
}

func MessagesGetRecentStickers(ctx context.Context, m Requester, i MessagesGetRecentStickersRequest) (IMessagesRecentStickers, error) {
	var res IMessagesRecentStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveRecentStickerRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
	ID       IInputDocument
	Unsave   bool
}

func (*MessagesSaveRecentStickerRequest) CRC() uint32 {
	return 0x392718f8
}

func MessagesSaveRecentSticker(ctx context.Context, m Requester, i MessagesSaveRecentStickerRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesClearRecentStickersRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesClearRecentStickersRequest) CRC() uint32 {
	return 0x8999602d
}

func MessagesClearRecentStickers(ctx context.Context, m Requester, i MessagesClearRecentStickersRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetArchivedStickersRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Masks    bool     `tl:",omitempty:flags:0,implicit"`
	Emojis   bool     `tl:",omitempty:flags:1,implicit"`
	OffsetID int64
	Limit    int32
}

func (*MessagesGetArchivedStickersRequest) CRC() uint32 {
	return 0x57f17692
}

func MessagesGetArchivedStickers(ctx context.Context, m Requester, i MessagesGetArchivedStickersRequest) (IMessagesArchivedStickers, error) {
	var res IMessagesArchivedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMaskStickersRequest struct {
	Hash int64
}

func (*MessagesGetMaskStickersRequest) CRC() uint32 {
	return 0x640f82b8
}

func MessagesGetMaskStickers(ctx context.Context, m Requester, i MessagesGetMaskStickersRequest) (IMessagesAllStickers, error) {
	var res IMessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachedStickersRequest struct {
	Media IInputStickeredMedia
}

func (*MessagesGetAttachedStickersRequest) CRC() uint32 {
	return 0xcc5b67cc
}

func MessagesGetAttachedStickers(ctx context.Context, m Requester, i MessagesGetAttachedStickersRequest) ([]IStickerSetCovered, error) {
	var res []IStickerSetCovered
	return res, request(ctx, m, &i, &res)
}

type MessagesSetGameScoreRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	EditMessage bool     `tl:",omitempty:flags:0,implicit"`
	Force       bool     `tl:",omitempty:flags:1,implicit"`
	Peer        IInputPeer
	ID          int32
	UserID      IInputUser
	Score       int32
}

func (*MessagesSetGameScoreRequest) CRC() uint32 {
	return 0x8ef8ecc0
}

func MessagesSetGameScore(ctx context.Context, m Requester, i MessagesSetGameScoreRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineGameScoreRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	EditMessage bool     `tl:",omitempty:flags:0,implicit"`
	Force       bool     `tl:",omitempty:flags:1,implicit"`
	ID          IInputBotInlineMessageID
	UserID      IInputUser
	Score       int32
}

func (*MessagesSetInlineGameScoreRequest) CRC() uint32 {
	return 0x15ad9f64
}

func MessagesSetInlineGameScore(ctx context.Context, m Requester, i MessagesSetInlineGameScoreRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetGameHighScoresRequest struct {
	Peer   IInputPeer
	ID     int32
	UserID IInputUser
}

func (*MessagesGetGameHighScoresRequest) CRC() uint32 {
	return 0xe822649d
}

func MessagesGetGameHighScores(ctx context.Context, m Requester, i MessagesGetGameHighScoresRequest) (IMessagesHighScores, error) {
	var res IMessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineGameHighScoresRequest struct {
	ID     IInputBotInlineMessageID
	UserID IInputUser
}

func (*MessagesGetInlineGameHighScoresRequest) CRC() uint32 {
	return 0xf635e1b
}

func MessagesGetInlineGameHighScores(ctx context.Context, m Requester, i MessagesGetInlineGameHighScoresRequest) (IMessagesHighScores, error) {
	var res IMessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCommonChatsRequest struct {
	UserID IInputUser
	MaxID  int64
	Limit  int32
}

func (*MessagesGetCommonChatsRequest) CRC() uint32 {
	return 0xe40ca104
}

func MessagesGetCommonChats(ctx context.Context, m Requester, i MessagesGetCommonChatsRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPageRequest struct {
	URL  string
	Hash int32
}

func (*MessagesGetWebPageRequest) CRC() uint32 {
	return 0x8d9692a3
}

func MessagesGetWebPage(ctx context.Context, m Requester, i MessagesGetWebPageRequest) (IMessagesWebPage, error) {
	var res IMessagesWebPage
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleDialogPinRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Pinned bool     `tl:",omitempty:flags:0,implicit"`
	Peer   IInputDialogPeer
}

func (*MessagesToggleDialogPinRequest) CRC() uint32 {
	return 0xa731e257
}

func MessagesToggleDialogPin(ctx context.Context, m Requester, i MessagesToggleDialogPinRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderPinnedDialogsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Force    bool     `tl:",omitempty:flags:0,implicit"`
	FolderID int32
	Order    []IInputDialogPeer
}

func (*MessagesReorderPinnedDialogsRequest) CRC() uint32 {
	return 0x3b1adf37
}

func MessagesReorderPinnedDialogs(ctx context.Context, m Requester, i MessagesReorderPinnedDialogsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPinnedDialogsRequest struct {
	FolderID int32
}

func (*MessagesGetPinnedDialogsRequest) CRC() uint32 {
	return 0xd6b94df2
}

func MessagesGetPinnedDialogs(ctx context.Context, m Requester, i MessagesGetPinnedDialogsRequest) (IMessagesPeerDialogs, error) {
	var res IMessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotShippingResultsRequest struct {
	_               struct{} `tl:"flags,bitflag"`
	QueryID         int64
	Error           *string           `tl:",omitempty:flags:0"`
	ShippingOptions []IShippingOption `tl:",omitempty:flags:1"`
}

func (*MessagesSetBotShippingResultsRequest) CRC() uint32 {
	return 0xe5f672fa
}

func MessagesSetBotShippingResults(ctx context.Context, m Requester, i MessagesSetBotShippingResultsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotPrecheckoutResultsRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Success bool     `tl:",omitempty:flags:1,implicit"`
	QueryID int64
	Error   *string `tl:",omitempty:flags:0"`
}

func (*MessagesSetBotPrecheckoutResultsRequest) CRC() uint32 {
	return 0x9c2dd95
}

func MessagesSetBotPrecheckoutResults(ctx context.Context, m Requester, i MessagesSetBotPrecheckoutResultsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadMediaRequest struct {
	Peer  IInputPeer
	Media IInputMedia
}

func (*MessagesUploadMediaRequest) CRC() uint32 {
	return 0x519bc2b1
}

func MessagesUploadMedia(ctx context.Context, m Requester, i MessagesUploadMediaRequest) (IMessageMedia, error) {
	var res IMessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScreenshotNotificationRequest struct {
	Peer     IInputPeer
	ReplyTo  IInputReplyTo
	RandomID int64
}

func (*MessagesSendScreenshotNotificationRequest) CRC() uint32 {
	return 0xa1405817
}

func MessagesSendScreenshotNotification(ctx context.Context, m Requester, i MessagesSendScreenshotNotificationRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFavedStickersRequest struct {
	Hash int64
}

func (*MessagesGetFavedStickersRequest) CRC() uint32 {
	return 0x4f1aaa9
}

func MessagesGetFavedStickers(ctx context.Context, m Requester, i MessagesGetFavedStickersRequest) (IMessagesFavedStickers, error) {
	var res IMessagesFavedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesFaveStickerRequest struct {
	ID     IInputDocument
	Unfave bool
}

func (*MessagesFaveStickerRequest) CRC() uint32 {
	return 0xb9ffc55b
}

func MessagesFaveSticker(ctx context.Context, m Requester, i MessagesFaveStickerRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadMentionsRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IInputPeer
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

func (*MessagesGetUnreadMentionsRequest) CRC() uint32 {
	return 0xf107e790
}

func MessagesGetUnreadMentions(ctx context.Context, m Requester, i MessagesGetUnreadMentionsRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMentionsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesReadMentionsRequest) CRC() uint32 {
	return 0x36e5bf4d
}

func MessagesReadMentions(ctx context.Context, m Requester, i MessagesReadMentionsRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentLocationsRequest struct {
	Peer  IInputPeer
	Limit int32
	Hash  int64
}

func (*MessagesGetRecentLocationsRequest) CRC() uint32 {
	return 0x702a40e0
}

func MessagesGetRecentLocations(ctx context.Context, m Requester, i MessagesGetRecentLocationsRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMultiMediaRequest struct {
	_                      struct{} `tl:"flags,bitflag"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   IInputPeer
	ReplyTo                IInputReplyTo `tl:",omitempty:flags:0"`
	MultiMedia             []IInputSingleMedia
	ScheduleDate           *int32     `tl:",omitempty:flags:10"`
	SendAs                 IInputPeer `tl:",omitempty:flags:13"`
}

func (*MessagesSendMultiMediaRequest) CRC() uint32 {
	return 0x456e8987
}

func MessagesSendMultiMedia(ctx context.Context, m Requester, i MessagesSendMultiMediaRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadEncryptedFileRequest struct {
	Peer IInputEncryptedChat
	File IInputEncryptedFile
}

func (*MessagesUploadEncryptedFileRequest) CRC() uint32 {
	return 0x5057c497
}

func MessagesUploadEncryptedFile(ctx context.Context, m Requester, i MessagesUploadEncryptedFileRequest) (IEncryptedFile, error) {
	var res IEncryptedFile
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchStickerSetsRequest struct {
	_               struct{} `tl:"flags,bitflag"`
	ExcludeFeatured bool     `tl:",omitempty:flags:0,implicit"`
	Q               string
	Hash            int64
}

func (*MessagesSearchStickerSetsRequest) CRC() uint32 {
	return 0x35705b8a
}

func MessagesSearchStickerSets(ctx context.Context, m Requester, i MessagesSearchStickerSetsRequest) (IMessagesFoundStickerSets, error) {
	var res IMessagesFoundStickerSets
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSplitRangesRequest struct{}

func (*MessagesGetSplitRangesRequest) CRC() uint32 {
	return 0x1cff7e08
}

func MessagesGetSplitRanges(ctx context.Context, m Requester, i MessagesGetSplitRangesRequest) ([]IMessageRange, error) {
	var res []IMessageRange
	return res, request(ctx, m, &i, &res)
}

type MessagesMarkDialogUnreadRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Unread bool     `tl:",omitempty:flags:0,implicit"`
	Peer   IInputDialogPeer
}

func (*MessagesMarkDialogUnreadRequest) CRC() uint32 {
	return 0xc286d98f
}

func MessagesMarkDialogUnread(ctx context.Context, m Requester, i MessagesMarkDialogUnreadRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogUnreadMarksRequest struct{}

func (*MessagesGetDialogUnreadMarksRequest) CRC() uint32 {
	return 0x22e24e22
}

func MessagesGetDialogUnreadMarks(ctx context.Context, m Requester, i MessagesGetDialogUnreadMarksRequest) ([]IDialogPeer, error) {
	var res []IDialogPeer
	return res, request(ctx, m, &i, &res)
}

type MessagesClearAllDraftsRequest struct{}

func (*MessagesClearAllDraftsRequest) CRC() uint32 {
	return 0x7e58ee9c
}

func MessagesClearAllDrafts(ctx context.Context, m Requester, i MessagesClearAllDraftsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdatePinnedMessageRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Silent    bool     `tl:",omitempty:flags:0,implicit"`
	Unpin     bool     `tl:",omitempty:flags:1,implicit"`
	PmOneside bool     `tl:",omitempty:flags:2,implicit"`
	Peer      IInputPeer
	ID        int32
}

func (*MessagesUpdatePinnedMessageRequest) CRC() uint32 {
	return 0xd2aaf7ec
}

func MessagesUpdatePinnedMessage(ctx context.Context, m Requester, i MessagesUpdatePinnedMessageRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendVoteRequest struct {
	Peer    IInputPeer
	MsgID   int32
	Options [][]byte
}

func (*MessagesSendVoteRequest) CRC() uint32 {
	return 0x10ea6184
}

func MessagesSendVote(ctx context.Context, m Requester, i MessagesSendVoteRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollResultsRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*MessagesGetPollResultsRequest) CRC() uint32 {
	return 0x73bb643b
}

func MessagesGetPollResults(ctx context.Context, m Requester, i MessagesGetPollResultsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOnlinesRequest struct {
	Peer IInputPeer
}

func (*MessagesGetOnlinesRequest) CRC() uint32 {
	return 0x6e2be050
}

func MessagesGetOnlines(ctx context.Context, m Requester, i MessagesGetOnlinesRequest) (IChatOnlines, error) {
	var res IChatOnlines
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAboutRequest struct {
	Peer  IInputPeer
	About string
}

func (*MessagesEditChatAboutRequest) CRC() uint32 {
	return 0xdef60797
}

func MessagesEditChatAbout(ctx context.Context, m Requester, i MessagesEditChatAboutRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatDefaultBannedRightsRequest struct {
	Peer         IInputPeer
	BannedRights IChatBannedRights
}

func (*MessagesEditChatDefaultBannedRightsRequest) CRC() uint32 {
	return 0xa5866b41
}

func MessagesEditChatDefaultBannedRights(ctx context.Context, m Requester, i MessagesEditChatDefaultBannedRightsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsRequest struct {
	LangCode string
}

func (*MessagesGetEmojiKeywordsRequest) CRC() uint32 {
	return 0x35a0e062
}

func MessagesGetEmojiKeywords(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsRequest) (IEmojiKeywordsDifference, error) {
	var res IEmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsDifferenceRequest struct {
	LangCode    string
	FromVersion int32
}

func (*MessagesGetEmojiKeywordsDifferenceRequest) CRC() uint32 {
	return 0x1508b6af
}

func MessagesGetEmojiKeywordsDifference(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsDifferenceRequest) (IEmojiKeywordsDifference, error) {
	var res IEmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsLanguagesRequest struct {
	LangCodes []string
}

func (*MessagesGetEmojiKeywordsLanguagesRequest) CRC() uint32 {
	return 0x4e9963b2
}

func MessagesGetEmojiKeywordsLanguages(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsLanguagesRequest) ([]IEmojiLanguage, error) {
	var res []IEmojiLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiURLRequest struct {
	LangCode string
}

func (*MessagesGetEmojiURLRequest) CRC() uint32 {
	return 0xd5b10c26
}

func MessagesGetEmojiURL(ctx context.Context, m Requester, i MessagesGetEmojiURLRequest) (IEmojiURL, error) {
	var res IEmojiURL
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchCountersRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
	Filters  []IMessagesFilter
}

func (*MessagesGetSearchCountersRequest) CRC() uint32 {
	return 0xae7cc1
}

func MessagesGetSearchCounters(ctx context.Context, m Requester, i MessagesGetSearchCountersRequest) ([]IMessagesSearchCounter, error) {
	var res []IMessagesSearchCounter
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestURLAuthRequest struct {
	_        struct{}   `tl:"flags,bitflag"`
	Peer     IInputPeer `tl:",omitempty:flags:1"`
	MsgID    *int32     `tl:",omitempty:flags:1"`
	ButtonID *int32     `tl:",omitempty:flags:1"`
	URL      *string    `tl:",omitempty:flags:2"`
}

func (*MessagesRequestURLAuthRequest) CRC() uint32 {
	return 0x198fb446
}

func MessagesRequestURLAuth(ctx context.Context, m Requester, i MessagesRequestURLAuthRequest) (IURLAuthResult, error) {
	var res IURLAuthResult
	return res, request(ctx, m, &i, &res)
}

type MessagesAcceptURLAuthRequest struct {
	_            struct{}   `tl:"flags,bitflag"`
	WriteAllowed bool       `tl:",omitempty:flags:0,implicit"`
	Peer         IInputPeer `tl:",omitempty:flags:1"`
	MsgID        *int32     `tl:",omitempty:flags:1"`
	ButtonID     *int32     `tl:",omitempty:flags:1"`
	URL          *string    `tl:",omitempty:flags:2"`
}

func (*MessagesAcceptURLAuthRequest) CRC() uint32 {
	return 0xb12c7125
}

func MessagesAcceptURLAuth(ctx context.Context, m Requester, i MessagesAcceptURLAuthRequest) (IURLAuthResult, error) {
	var res IURLAuthResult
	return res, request(ctx, m, &i, &res)
}

type MessagesHidePeerSettingsBarRequest struct {
	Peer IInputPeer
}

func (*MessagesHidePeerSettingsBarRequest) CRC() uint32 {
	return 0x4facb138
}

func MessagesHidePeerSettingsBar(ctx context.Context, m Requester, i MessagesHidePeerSettingsBarRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledHistoryRequest struct {
	Peer IInputPeer
	Hash int64
}

func (*MessagesGetScheduledHistoryRequest) CRC() uint32 {
	return 0xf516760b
}

func MessagesGetScheduledHistory(ctx context.Context, m Requester, i MessagesGetScheduledHistoryRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledMessagesRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*MessagesGetScheduledMessagesRequest) CRC() uint32 {
	return 0xbdbb0464
}

func MessagesGetScheduledMessages(ctx context.Context, m Requester, i MessagesGetScheduledMessagesRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScheduledMessagesRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*MessagesSendScheduledMessagesRequest) CRC() uint32 {
	return 0xbd38850a
}

func MessagesSendScheduledMessages(ctx context.Context, m Requester, i MessagesSendScheduledMessagesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteScheduledMessagesRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*MessagesDeleteScheduledMessagesRequest) CRC() uint32 {
	return 0x59ae2b16
}

func MessagesDeleteScheduledMessages(ctx context.Context, m Requester, i MessagesDeleteScheduledMessagesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollVotesRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Peer   IInputPeer
	ID     int32
	Option *[]byte `tl:",omitempty:flags:0"`
	Offset *string `tl:",omitempty:flags:1"`
	Limit  int32
}

func (*MessagesGetPollVotesRequest) CRC() uint32 {
	return 0xb86e380e
}

func MessagesGetPollVotes(ctx context.Context, m Requester, i MessagesGetPollVotesRequest) (IMessagesVotesList, error) {
	var res IMessagesVotesList
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleStickerSetsRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Uninstall   bool     `tl:",omitempty:flags:0,implicit"`
	Archive     bool     `tl:",omitempty:flags:1,implicit"`
	Unarchive   bool     `tl:",omitempty:flags:2,implicit"`
	Stickersets []IInputStickerSet
}

func (*MessagesToggleStickerSetsRequest) CRC() uint32 {
	return 0xb5052fea
}

func MessagesToggleStickerSets(ctx context.Context, m Requester, i MessagesToggleStickerSetsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogFiltersRequest struct{}

func (*MessagesGetDialogFiltersRequest) CRC() uint32 {
	return 0xf19ed96d
}

func MessagesGetDialogFilters(ctx context.Context, m Requester, i MessagesGetDialogFiltersRequest) ([]IDialogFilter, error) {
	var res []IDialogFilter
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSuggestedDialogFiltersRequest struct{}

func (*MessagesGetSuggestedDialogFiltersRequest) CRC() uint32 {
	return 0xa29cd42c
}

func MessagesGetSuggestedDialogFilters(ctx context.Context, m Requester, i MessagesGetSuggestedDialogFiltersRequest) ([]IDialogFilterSuggested, error) {
	var res []IDialogFilterSuggested
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateDialogFilterRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32
	Filter IDialogFilter `tl:",omitempty:flags:0"`
}

func (*MessagesUpdateDialogFilterRequest) CRC() uint32 {
	return 0x1ad4a04a
}

func MessagesUpdateDialogFilter(ctx context.Context, m Requester, i MessagesUpdateDialogFilterRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateDialogFiltersOrderRequest struct {
	Order []int32
}

func (*MessagesUpdateDialogFiltersOrderRequest) CRC() uint32 {
	return 0xc563c1e4
}

func MessagesUpdateDialogFiltersOrder(ctx context.Context, m Requester, i MessagesUpdateDialogFiltersOrderRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOldFeaturedStickersRequest struct {
	Offset int32
	Limit  int32
	Hash   int64
}

func (*MessagesGetOldFeaturedStickersRequest) CRC() uint32 {
	return 0x7ed094a1
}

func MessagesGetOldFeaturedStickers(ctx context.Context, m Requester, i MessagesGetOldFeaturedStickersRequest) (IMessagesFeaturedStickers, error) {
	var res IMessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRepliesRequest struct {
	Peer       IInputPeer
	MsgID      int32
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int64
}

func (*MessagesGetRepliesRequest) CRC() uint32 {
	return 0x22ddd30c
}

func MessagesGetReplies(ctx context.Context, m Requester, i MessagesGetRepliesRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDiscussionMessageRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*MessagesGetDiscussionMessageRequest) CRC() uint32 {
	return 0x446972fd
}

func MessagesGetDiscussionMessage(ctx context.Context, m Requester, i MessagesGetDiscussionMessageRequest) (IMessagesDiscussionMessage, error) {
	var res IMessagesDiscussionMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReadDiscussionRequest struct {
	Peer      IInputPeer
	MsgID     int32
	ReadMaxID int32
}

func (*MessagesReadDiscussionRequest) CRC() uint32 {
	return 0xf731a9f4
}

func MessagesReadDiscussion(ctx context.Context, m Requester, i MessagesReadDiscussionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUnpinAllMessagesRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesUnpinAllMessagesRequest) CRC() uint32 {
	return 0xee22b9a8
}

func MessagesUnpinAllMessages(ctx context.Context, m Requester, i MessagesUnpinAllMessagesRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatRequest struct {
	ChatID int64
}

func (*MessagesDeleteChatRequest) CRC() uint32 {
	return 0x5bd0ee50
}

func MessagesDeleteChat(ctx context.Context, m Requester, i MessagesDeleteChatRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeletePhoneCallHistoryRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Revoke bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesDeletePhoneCallHistoryRequest) CRC() uint32 {
	return 0xf9cbe409
}

func MessagesDeletePhoneCallHistory(ctx context.Context, m Requester, i MessagesDeletePhoneCallHistoryRequest) (IMessagesAffectedFoundMessages, error) {
	var res IMessagesAffectedFoundMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportRequest struct {
	ImportHead string
}

func (*MessagesCheckHistoryImportRequest) CRC() uint32 {
	return 0x43fe19f3
}

func MessagesCheckHistoryImport(ctx context.Context, m Requester, i MessagesCheckHistoryImportRequest) (IMessagesHistoryImportParsed, error) {
	var res IMessagesHistoryImportParsed
	return res, request(ctx, m, &i, &res)
}

type MessagesInitHistoryImportRequest struct {
	Peer       IInputPeer
	File       IInputFile
	MediaCount int32
}

func (*MessagesInitHistoryImportRequest) CRC() uint32 {
	return 0x34090c3b
}

func MessagesInitHistoryImport(ctx context.Context, m Requester, i MessagesInitHistoryImportRequest) (IMessagesHistoryImport, error) {
	var res IMessagesHistoryImport
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadImportedMediaRequest struct {
	Peer     IInputPeer
	ImportID int64
	FileName string
	Media    IInputMedia
}

func (*MessagesUploadImportedMediaRequest) CRC() uint32 {
	return 0x2a862092
}

func MessagesUploadImportedMedia(ctx context.Context, m Requester, i MessagesUploadImportedMediaRequest) (IMessageMedia, error) {
	var res IMessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesStartHistoryImportRequest struct {
	Peer     IInputPeer
	ImportID int64
}

func (*MessagesStartHistoryImportRequest) CRC() uint32 {
	return 0xb43df344
}

func MessagesStartHistoryImport(ctx context.Context, m Requester, i MessagesStartHistoryImportRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInvitesRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	Revoked    bool     `tl:",omitempty:flags:3,implicit"`
	Peer       IInputPeer
	AdminID    IInputUser
	OffsetDate *int32  `tl:",omitempty:flags:2"`
	OffsetLink *string `tl:",omitempty:flags:2"`
	Limit      int32
}

func (*MessagesGetExportedChatInvitesRequest) CRC() uint32 {
	return 0xa2b5a3f6
}

func MessagesGetExportedChatInvites(ctx context.Context, m Requester, i MessagesGetExportedChatInvitesRequest) (IMessagesExportedChatInvites, error) {
	var res IMessagesExportedChatInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInviteRequest struct {
	Peer IInputPeer
	Link string
}

func (*MessagesGetExportedChatInviteRequest) CRC() uint32 {
	return 0x73746f5c
}

func MessagesGetExportedChatInvite(ctx context.Context, m Requester, i MessagesGetExportedChatInviteRequest) (IMessagesExportedChatInvite, error) {
	var res IMessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesEditExportedChatInviteRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	Revoked       bool     `tl:",omitempty:flags:2,implicit"`
	Peer          IInputPeer
	Link          string
	ExpireDate    *int32  `tl:",omitempty:flags:0"`
	UsageLimit    *int32  `tl:",omitempty:flags:1"`
	RequestNeeded *bool   `tl:",omitempty:flags:3"`
	Title         *string `tl:",omitempty:flags:4"`
}

func (*MessagesEditExportedChatInviteRequest) CRC() uint32 {
	return 0xbdca2f75
}

func MessagesEditExportedChatInvite(ctx context.Context, m Requester, i MessagesEditExportedChatInviteRequest) (IMessagesExportedChatInvite, error) {
	var res IMessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteRevokedExportedChatInvitesRequest struct {
	Peer    IInputPeer
	AdminID IInputUser
}

func (*MessagesDeleteRevokedExportedChatInvitesRequest) CRC() uint32 {
	return 0x56987bd5
}

func MessagesDeleteRevokedExportedChatInvites(ctx context.Context, m Requester, i MessagesDeleteRevokedExportedChatInvitesRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteExportedChatInviteRequest struct {
	Peer IInputPeer
	Link string
}

func (*MessagesDeleteExportedChatInviteRequest) CRC() uint32 {
	return 0xd464a42b
}

func MessagesDeleteExportedChatInvite(ctx context.Context, m Requester, i MessagesDeleteExportedChatInviteRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAdminsWithInvitesRequest struct {
	Peer IInputPeer
}

func (*MessagesGetAdminsWithInvitesRequest) CRC() uint32 {
	return 0x3920e6ef
}

func MessagesGetAdminsWithInvites(ctx context.Context, m Requester, i MessagesGetAdminsWithInvitesRequest) (IMessagesChatAdminsWithInvites, error) {
	var res IMessagesChatAdminsWithInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatInviteImportersRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	Requested  bool     `tl:",omitempty:flags:0,implicit"`
	Peer       IInputPeer
	Link       *string `tl:",omitempty:flags:1"`
	Q          *string `tl:",omitempty:flags:2"`
	OffsetDate int32
	OffsetUser IInputUser
	Limit      int32
}

func (*MessagesGetChatInviteImportersRequest) CRC() uint32 {
	return 0xdf04dd4e
}

func MessagesGetChatInviteImporters(ctx context.Context, m Requester, i MessagesGetChatInviteImportersRequest) (IMessagesChatInviteImporters, error) {
	var res IMessagesChatInviteImporters
	return res, request(ctx, m, &i, &res)
}

type MessagesSetHistoryTTLRequest struct {
	Peer   IInputPeer
	Period int32
}

func (*MessagesSetHistoryTTLRequest) CRC() uint32 {
	return 0xb80e5fe4
}

func MessagesSetHistoryTTL(ctx context.Context, m Requester, i MessagesSetHistoryTTLRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportPeerRequest struct {
	Peer IInputPeer
}

func (*MessagesCheckHistoryImportPeerRequest) CRC() uint32 {
	return 0x5dc60f03
}

func MessagesCheckHistoryImportPeer(ctx context.Context, m Requester, i MessagesCheckHistoryImportPeerRequest) (IMessagesCheckedHistoryImportPeer, error) {
	var res IMessagesCheckedHistoryImportPeer
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatThemeRequest struct {
	Peer     IInputPeer
	Emoticon string
}

func (*MessagesSetChatThemeRequest) CRC() uint32 {
	return 0xe63be13f
}

func MessagesSetChatTheme(ctx context.Context, m Requester, i MessagesSetChatThemeRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReadParticipantsRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*MessagesGetMessageReadParticipantsRequest) CRC() uint32 {
	return 0x31c1c44f
}

func MessagesGetMessageReadParticipants(ctx context.Context, m Requester, i MessagesGetMessageReadParticipantsRequest) ([]IReadParticipantDate, error) {
	var res []IReadParticipantDate
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsCalendarRequest struct {
	Peer       IInputPeer
	Filter     IMessagesFilter
	OffsetID   int32
	OffsetDate int32
}

func (*MessagesGetSearchResultsCalendarRequest) CRC() uint32 {
	return 0x49f0bde9
}

func MessagesGetSearchResultsCalendar(ctx context.Context, m Requester, i MessagesGetSearchResultsCalendarRequest) (IMessagesSearchResultsCalendar, error) {
	var res IMessagesSearchResultsCalendar
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsPositionsRequest struct {
	Peer     IInputPeer
	Filter   IMessagesFilter
	OffsetID int32
	Limit    int32
}

func (*MessagesGetSearchResultsPositionsRequest) CRC() uint32 {
	return 0x6e9583a3
}

func MessagesGetSearchResultsPositions(ctx context.Context, m Requester, i MessagesGetSearchResultsPositionsRequest) (IMessagesSearchResultsPositions, error) {
	var res IMessagesSearchResultsPositions
	return res, request(ctx, m, &i, &res)
}

type MessagesHideChatJoinRequestRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Approved bool     `tl:",omitempty:flags:0,implicit"`
	Peer     IInputPeer
	UserID   IInputUser
}

func (*MessagesHideChatJoinRequestRequest) CRC() uint32 {
	return 0x7fe7e815
}

func MessagesHideChatJoinRequest(ctx context.Context, m Requester, i MessagesHideChatJoinRequestRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesHideAllChatJoinRequestsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Approved bool     `tl:",omitempty:flags:0,implicit"`
	Peer     IInputPeer
	Link     *string `tl:",omitempty:flags:1"`
}

func (*MessagesHideAllChatJoinRequestsRequest) CRC() uint32 {
	return 0xe085f4ea
}

func MessagesHideAllChatJoinRequests(ctx context.Context, m Requester, i MessagesHideAllChatJoinRequestsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleNoForwardsRequest struct {
	Peer    IInputPeer
	Enabled bool
}

func (*MessagesToggleNoForwardsRequest) CRC() uint32 {
	return 0xb11eafa2
}

func MessagesToggleNoForwards(ctx context.Context, m Requester, i MessagesToggleNoForwardsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDefaultSendAsRequest struct {
	Peer   IInputPeer
	SendAs IInputPeer
}

func (*MessagesSaveDefaultSendAsRequest) CRC() uint32 {
	return 0xccfddf96
}

func MessagesSaveDefaultSendAs(ctx context.Context, m Requester, i MessagesSaveDefaultSendAsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendReactionRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Big         bool     `tl:",omitempty:flags:1,implicit"`
	AddToRecent bool     `tl:",omitempty:flags:2,implicit"`
	Peer        IInputPeer
	MsgID       int32
	Reaction    []IReaction `tl:",omitempty:flags:0"`
}

func (*MessagesSendReactionRequest) CRC() uint32 {
	return 0xd30d78d4
}

func MessagesSendReaction(ctx context.Context, m Requester, i MessagesSendReactionRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesReactionsRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*MessagesGetMessagesReactionsRequest) CRC() uint32 {
	return 0x8bba90e6
}

func MessagesGetMessagesReactions(ctx context.Context, m Requester, i MessagesGetMessagesReactionsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReactionsListRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	ID       int32
	Reaction IReaction `tl:",omitempty:flags:0"`
	Offset   *string   `tl:",omitempty:flags:1"`
	Limit    int32
}

func (*MessagesGetMessageReactionsListRequest) CRC() uint32 {
	return 0x461b3f48
}

func MessagesGetMessageReactionsList(ctx context.Context, m Requester, i MessagesGetMessageReactionsListRequest) (IMessagesMessageReactionsList, error) {
	var res IMessagesMessageReactionsList
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatAvailableReactionsRequest struct {
	Peer               IInputPeer
	AvailableReactions IChatReactions
}

func (*MessagesSetChatAvailableReactionsRequest) CRC() uint32 {
	return 0xfeb16771
}

func MessagesSetChatAvailableReactions(ctx context.Context, m Requester, i MessagesSetChatAvailableReactionsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAvailableReactionsRequest struct {
	Hash int32
}

func (*MessagesGetAvailableReactionsRequest) CRC() uint32 {
	return 0x18dea0ac
}

func MessagesGetAvailableReactions(ctx context.Context, m Requester, i MessagesGetAvailableReactionsRequest) (IMessagesAvailableReactions, error) {
	var res IMessagesAvailableReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultReactionRequest struct {
	Reaction IReaction
}

func (*MessagesSetDefaultReactionRequest) CRC() uint32 {
	return 0x4f47a016
}

func MessagesSetDefaultReaction(ctx context.Context, m Requester, i MessagesSetDefaultReactionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesTranslateTextRequest struct {
	_      struct{}            `tl:"flags,bitflag"`
	Peer   IInputPeer          `tl:",omitempty:flags:0"`
	ID     []int32             `tl:",omitempty:flags:0"`
	Text   []ITextWithEntities `tl:",omitempty:flags:1"`
	ToLang string
}

func (*MessagesTranslateTextRequest) CRC() uint32 {
	return 0x63183030
}

func MessagesTranslateText(ctx context.Context, m Requester, i MessagesTranslateTextRequest) (IMessagesTranslatedText, error) {
	var res IMessagesTranslatedText
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadReactionsRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      IInputPeer
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

func (*MessagesGetUnreadReactionsRequest) CRC() uint32 {
	return 0x3223495b
}

func MessagesGetUnreadReactions(ctx context.Context, m Requester, i MessagesGetUnreadReactionsRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadReactionsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     IInputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesReadReactionsRequest) CRC() uint32 {
	return 0x54aa7f8e
}

func MessagesReadReactions(ctx context.Context, m Requester, i MessagesReadReactionsRequest) (IMessagesAffectedHistory, error) {
	var res IMessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchSentMediaRequest struct {
	Q      string
	Filter IMessagesFilter
	Limit  int32
}

func (*MessagesSearchSentMediaRequest) CRC() uint32 {
	return 0x107e31a0
}

func MessagesSearchSentMedia(ctx context.Context, m Requester, i MessagesSearchSentMediaRequest) (IMessagesMessages, error) {
	var res IMessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotsRequest struct {
	Hash int64
}

func (*MessagesGetAttachMenuBotsRequest) CRC() uint32 {
	return 0x16fcc2cb
}

func MessagesGetAttachMenuBots(ctx context.Context, m Requester, i MessagesGetAttachMenuBotsRequest) (IAttachMenuBots, error) {
	var res IAttachMenuBots
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotRequest struct {
	Bot IInputUser
}

func (*MessagesGetAttachMenuBotRequest) CRC() uint32 {
	return 0x77216192
}

func MessagesGetAttachMenuBot(ctx context.Context, m Requester, i MessagesGetAttachMenuBotRequest) (IAttachMenuBotsBot, error) {
	var res IAttachMenuBotsBot
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleBotInAttachMenuRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	WriteAllowed bool     `tl:",omitempty:flags:0,implicit"`
	Bot          IInputUser
	Enabled      bool
}

func (*MessagesToggleBotInAttachMenuRequest) CRC() uint32 {
	return 0x69f59d69
}

func MessagesToggleBotInAttachMenu(ctx context.Context, m Requester, i MessagesToggleBotInAttachMenuRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestWebViewRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	FromBotMenu bool     `tl:",omitempty:flags:4,implicit"`
	Silent      bool     `tl:",omitempty:flags:5,implicit"`
	Peer        IInputPeer
	Bot         IInputUser
	URL         *string   `tl:",omitempty:flags:1"`
	StartParam  *string   `tl:",omitempty:flags:3"`
	ThemeParams IDataJSON `tl:",omitempty:flags:2"`
	Platform    string
	ReplyTo     IInputReplyTo `tl:",omitempty:flags:0"`
	SendAs      IInputPeer    `tl:",omitempty:flags:13"`
}

func (*MessagesRequestWebViewRequest) CRC() uint32 {
	return 0x269dc2c1
}

func MessagesRequestWebView(ctx context.Context, m Requester, i MessagesRequestWebViewRequest) (IWebViewResult, error) {
	var res IWebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesProlongWebViewRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Silent  bool     `tl:",omitempty:flags:5,implicit"`
	Peer    IInputPeer
	Bot     IInputUser
	QueryID int64
	ReplyTo IInputReplyTo `tl:",omitempty:flags:0"`
	SendAs  IInputPeer    `tl:",omitempty:flags:13"`
}

func (*MessagesProlongWebViewRequest) CRC() uint32 {
	return 0xb0d81a83
}

func MessagesProlongWebView(ctx context.Context, m Requester, i MessagesProlongWebViewRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestSimpleWebViewRequest struct {
	_                 struct{} `tl:"flags,bitflag"`
	FromSwitchWebview bool     `tl:",omitempty:flags:1,implicit"`
	FromSideMenu      bool     `tl:",omitempty:flags:2,implicit"`
	Bot               IInputUser
	URL               *string   `tl:",omitempty:flags:3"`
	StartParam        *string   `tl:",omitempty:flags:4"`
	ThemeParams       IDataJSON `tl:",omitempty:flags:0"`
	Platform          string
}

func (*MessagesRequestSimpleWebViewRequest) CRC() uint32 {
	return 0x1a46500a
}

func MessagesRequestSimpleWebView(ctx context.Context, m Requester, i MessagesRequestSimpleWebViewRequest) (ISimpleWebViewResult, error) {
	var res ISimpleWebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewResultMessageRequest struct {
	BotQueryID string
	Result     IInputBotInlineResult
}

func (*MessagesSendWebViewResultMessageRequest) CRC() uint32 {
	return 0xa4314f5
}

func MessagesSendWebViewResultMessage(ctx context.Context, m Requester, i MessagesSendWebViewResultMessageRequest) (IWebViewMessageSent, error) {
	var res IWebViewMessageSent
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewDataRequest struct {
	Bot        IInputUser
	RandomID   int64
	ButtonText string
	Data       string
}

func (*MessagesSendWebViewDataRequest) CRC() uint32 {
	return 0xdc0242c8
}

func MessagesSendWebViewData(ctx context.Context, m Requester, i MessagesSendWebViewDataRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesTranscribeAudioRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*MessagesTranscribeAudioRequest) CRC() uint32 {
	return 0x269e9a49
}

func MessagesTranscribeAudio(ctx context.Context, m Requester, i MessagesTranscribeAudioRequest) (IMessagesTranscribedAudio, error) {
	var res IMessagesTranscribedAudio
	return res, request(ctx, m, &i, &res)
}

type MessagesRateTranscribedAudioRequest struct {
	Peer            IInputPeer
	MsgID           int32
	TranscriptionID int64
	Good            bool
}

func (*MessagesRateTranscribedAudioRequest) CRC() uint32 {
	return 0x7f1d072f
}

func MessagesRateTranscribedAudio(ctx context.Context, m Requester, i MessagesRateTranscribedAudioRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCustomEmojiDocumentsRequest struct {
	DocumentID []int64
}

func (*MessagesGetCustomEmojiDocumentsRequest) CRC() uint32 {
	return 0xd9ab0f54
}

func MessagesGetCustomEmojiDocuments(ctx context.Context, m Requester, i MessagesGetCustomEmojiDocumentsRequest) ([]IDocument, error) {
	var res []IDocument
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStickersRequest struct {
	Hash int64
}

func (*MessagesGetEmojiStickersRequest) CRC() uint32 {
	return 0xfbfca18f
}

func MessagesGetEmojiStickers(ctx context.Context, m Requester, i MessagesGetEmojiStickersRequest) (IMessagesAllStickers, error) {
	var res IMessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFeaturedEmojiStickersRequest struct {
	Hash int64
}

func (*MessagesGetFeaturedEmojiStickersRequest) CRC() uint32 {
	return 0xecf6736
}

func MessagesGetFeaturedEmojiStickers(ctx context.Context, m Requester, i MessagesGetFeaturedEmojiStickersRequest) (IMessagesFeaturedStickers, error) {
	var res IMessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReportReactionRequest struct {
	Peer         IInputPeer
	ID           int32
	ReactionPeer IInputPeer
}

func (*MessagesReportReactionRequest) CRC() uint32 {
	return 0x3f64c076
}

func MessagesReportReaction(ctx context.Context, m Requester, i MessagesReportReactionRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetTopReactionsRequest struct {
	Limit int32
	Hash  int64
}

func (*MessagesGetTopReactionsRequest) CRC() uint32 {
	return 0xbb8125ba
}

func MessagesGetTopReactions(ctx context.Context, m Requester, i MessagesGetTopReactionsRequest) (IMessagesReactions, error) {
	var res IMessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentReactionsRequest struct {
	Limit int32
	Hash  int64
}

func (*MessagesGetRecentReactionsRequest) CRC() uint32 {
	return 0x39461db2
}

func MessagesGetRecentReactions(ctx context.Context, m Requester, i MessagesGetRecentReactionsRequest) (IMessagesReactions, error) {
	var res IMessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesClearRecentReactionsRequest struct{}

func (*MessagesClearRecentReactionsRequest) CRC() uint32 {
	return 0x9dfeefb4
}

func MessagesClearRecentReactions(ctx context.Context, m Requester, i MessagesClearRecentReactionsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExtendedMediaRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*MessagesGetExtendedMediaRequest) CRC() uint32 {
	return 0x84f80814
}

func MessagesGetExtendedMedia(ctx context.Context, m Requester, i MessagesGetExtendedMediaRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultHistoryTTLRequest struct {
	Period int32
}

func (*MessagesSetDefaultHistoryTTLRequest) CRC() uint32 {
	return 0x9eb51445
}

func MessagesSetDefaultHistoryTTL(ctx context.Context, m Requester, i MessagesSetDefaultHistoryTTLRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDefaultHistoryTTLRequest struct{}

func (*MessagesGetDefaultHistoryTTLRequest) CRC() uint32 {
	return 0x658b7188
}

func MessagesGetDefaultHistoryTTL(ctx context.Context, m Requester, i MessagesGetDefaultHistoryTTLRequest) (IDefaultHistoryTTL, error) {
	var res IDefaultHistoryTTL
	return res, request(ctx, m, &i, &res)
}

type MessagesSendBotRequestedPeerRequest struct {
	Peer           IInputPeer
	MsgID          int32
	ButtonID       int32
	RequestedPeers []IInputPeer
}

func (*MessagesSendBotRequestedPeerRequest) CRC() uint32 {
	return 0x91b2d060
}

func MessagesSendBotRequestedPeer(ctx context.Context, m Requester, i MessagesSendBotRequestedPeerRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiGroupsRequest struct {
	Hash int32
}

func (*MessagesGetEmojiGroupsRequest) CRC() uint32 {
	return 0x7488ce5b
}

func MessagesGetEmojiGroups(ctx context.Context, m Requester, i MessagesGetEmojiGroupsRequest) (IMessagesEmojiGroups, error) {
	var res IMessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStatusGroupsRequest struct {
	Hash int32
}

func (*MessagesGetEmojiStatusGroupsRequest) CRC() uint32 {
	return 0x2ecd56cd
}

func MessagesGetEmojiStatusGroups(ctx context.Context, m Requester, i MessagesGetEmojiStatusGroupsRequest) (IMessagesEmojiGroups, error) {
	var res IMessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiProfilePhotoGroupsRequest struct {
	Hash int32
}

func (*MessagesGetEmojiProfilePhotoGroupsRequest) CRC() uint32 {
	return 0x21a548f3
}

func MessagesGetEmojiProfilePhotoGroups(ctx context.Context, m Requester, i MessagesGetEmojiProfilePhotoGroupsRequest) (IMessagesEmojiGroups, error) {
	var res IMessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchCustomEmojiRequest struct {
	Emoticon string
	Hash     int64
}

func (*MessagesSearchCustomEmojiRequest) CRC() uint32 {
	return 0x2c11c0d7
}

func MessagesSearchCustomEmoji(ctx context.Context, m Requester, i MessagesSearchCustomEmojiRequest) (IEmojiList, error) {
	var res IEmojiList
	return res, request(ctx, m, &i, &res)
}

type MessagesTogglePeerTranslationsRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Disabled bool     `tl:",omitempty:flags:0,implicit"`
	Peer     IInputPeer
}

func (*MessagesTogglePeerTranslationsRequest) CRC() uint32 {
	return 0xe47cb579
}

func MessagesTogglePeerTranslations(ctx context.Context, m Requester, i MessagesTogglePeerTranslationsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotAppRequest struct {
	App  IInputBotApp
	Hash int64
}

func (*MessagesGetBotAppRequest) CRC() uint32 {
	return 0x34fdc5c3
}

func MessagesGetBotApp(ctx context.Context, m Requester, i MessagesGetBotAppRequest) (IMessagesBotApp, error) {
	var res IMessagesBotApp
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestAppWebViewRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	WriteAllowed bool     `tl:",omitempty:flags:0,implicit"`
	Peer         IInputPeer
	App          IInputBotApp
	StartParam   *string   `tl:",omitempty:flags:1"`
	ThemeParams  IDataJSON `tl:",omitempty:flags:2"`
	Platform     string
}

func (*MessagesRequestAppWebViewRequest) CRC() uint32 {
	return 0x8c5a3b3c
}

func MessagesRequestAppWebView(ctx context.Context, m Requester, i MessagesRequestAppWebViewRequest) (IAppWebViewResult, error) {
	var res IAppWebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatWallPaperRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	ForBoth   bool     `tl:",omitempty:flags:3,implicit"`
	Revert    bool     `tl:",omitempty:flags:4,implicit"`
	Peer      IInputPeer
	Wallpaper IInputWallPaper    `tl:",omitempty:flags:0"`
	Settings  IWallPaperSettings `tl:",omitempty:flags:2"`
	ID        *int32             `tl:",omitempty:flags:1"`
}

func (*MessagesSetChatWallPaperRequest) CRC() uint32 {
	return 0x8ffacae1
}

func MessagesSetChatWallPaper(ctx context.Context, m Requester, i MessagesSetChatWallPaperRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchEmojiStickerSetsRequest struct {
	_               struct{} `tl:"flags,bitflag"`
	ExcludeFeatured bool     `tl:",omitempty:flags:0,implicit"`
	Q               string
	Hash            int64
}

func (*MessagesSearchEmojiStickerSetsRequest) CRC() uint32 {
	return 0x92b4494c
}

func MessagesSearchEmojiStickerSets(ctx context.Context, m Requester, i MessagesSearchEmojiStickerSetsRequest) (IMessagesFoundStickerSets, error) {
	var res IMessagesFoundStickerSets
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentFormRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Invoice     IInputInvoice
	ThemeParams IDataJSON `tl:",omitempty:flags:0"`
}

func (*PaymentsGetPaymentFormRequest) CRC() uint32 {
	return 0x37148dbb
}

func PaymentsGetPaymentForm(ctx context.Context, m Requester, i PaymentsGetPaymentFormRequest) (IPaymentsPaymentForm, error) {
	var res IPaymentsPaymentForm
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentReceiptRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*PaymentsGetPaymentReceiptRequest) CRC() uint32 {
	return 0x2478d1cc
}

func PaymentsGetPaymentReceipt(ctx context.Context, m Requester, i PaymentsGetPaymentReceiptRequest) (IPaymentsPaymentReceipt, error) {
	var res IPaymentsPaymentReceipt
	return res, request(ctx, m, &i, &res)
}

type PaymentsValidateRequestedInfoRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Save    bool     `tl:",omitempty:flags:0,implicit"`
	Invoice IInputInvoice
	Info    IPaymentRequestedInfo
}

func (*PaymentsValidateRequestedInfoRequest) CRC() uint32 {
	return 0xb6c8f12b
}

func PaymentsValidateRequestedInfo(ctx context.Context, m Requester, i PaymentsValidateRequestedInfoRequest) (IPaymentsValidatedRequestedInfo, error) {
	var res IPaymentsValidatedRequestedInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsSendPaymentFormRequest struct {
	_                struct{} `tl:"flags,bitflag"`
	FormID           int64
	Invoice          IInputInvoice
	RequestedInfoID  *string `tl:",omitempty:flags:0"`
	ShippingOptionID *string `tl:",omitempty:flags:1"`
	Credentials      IInputPaymentCredentials
	TipAmount        *int64 `tl:",omitempty:flags:2"`
}

func (*PaymentsSendPaymentFormRequest) CRC() uint32 {
	return 0x2d03522f
}

func PaymentsSendPaymentForm(ctx context.Context, m Requester, i PaymentsSendPaymentFormRequest) (IPaymentsPaymentResult, error) {
	var res IPaymentsPaymentResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetSavedInfoRequest struct{}

func (*PaymentsGetSavedInfoRequest) CRC() uint32 {
	return 0x227d824b
}

func PaymentsGetSavedInfo(ctx context.Context, m Requester, i PaymentsGetSavedInfoRequest) (IPaymentsSavedInfo, error) {
	var res IPaymentsSavedInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsClearSavedInfoRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	Credentials bool     `tl:",omitempty:flags:0,implicit"`
	Info        bool     `tl:",omitempty:flags:1,implicit"`
}

func (*PaymentsClearSavedInfoRequest) CRC() uint32 {
	return 0xd83d70c1
}

func PaymentsClearSavedInfo(ctx context.Context, m Requester, i PaymentsClearSavedInfoRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetBankCardDataRequest struct {
	Number string
}

func (*PaymentsGetBankCardDataRequest) CRC() uint32 {
	return 0x2e79d779
}

func PaymentsGetBankCardData(ctx context.Context, m Requester, i PaymentsGetBankCardDataRequest) (IPaymentsBankCardData, error) {
	var res IPaymentsBankCardData
	return res, request(ctx, m, &i, &res)
}

type PaymentsExportInvoiceRequest struct {
	InvoiceMedia IInputMedia
}

func (*PaymentsExportInvoiceRequest) CRC() uint32 {
	return 0xf91b065
}

func PaymentsExportInvoice(ctx context.Context, m Requester, i PaymentsExportInvoiceRequest) (IPaymentsExportedInvoice, error) {
	var res IPaymentsExportedInvoice
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignAppStoreTransactionRequest struct {
	Receipt []byte
	Purpose IInputStorePaymentPurpose
}

func (*PaymentsAssignAppStoreTransactionRequest) CRC() uint32 {
	return 0x80ed747d
}

func PaymentsAssignAppStoreTransaction(ctx context.Context, m Requester, i PaymentsAssignAppStoreTransactionRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignPlayMarketTransactionRequest struct {
	Receipt IDataJSON
	Purpose IInputStorePaymentPurpose
}

func (*PaymentsAssignPlayMarketTransactionRequest) CRC() uint32 {
	return 0xdffd50d3
}

func PaymentsAssignPlayMarketTransaction(ctx context.Context, m Requester, i PaymentsAssignPlayMarketTransactionRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PaymentsCanPurchasePremiumRequest struct {
	Purpose IInputStorePaymentPurpose
}

func (*PaymentsCanPurchasePremiumRequest) CRC() uint32 {
	return 0x9fc19eb6
}

func PaymentsCanPurchasePremium(ctx context.Context, m Requester, i PaymentsCanPurchasePremiumRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPremiumGiftCodeOptionsRequest struct {
	_         struct{}   `tl:"flags,bitflag"`
	BoostPeer IInputPeer `tl:",omitempty:flags:0"`
}

func (*PaymentsGetPremiumGiftCodeOptionsRequest) CRC() uint32 {
	return 0x2757ba54
}

func PaymentsGetPremiumGiftCodeOptions(ctx context.Context, m Requester, i PaymentsGetPremiumGiftCodeOptionsRequest) ([]IPremiumGiftCodeOption, error) {
	var res []IPremiumGiftCodeOption
	return res, request(ctx, m, &i, &res)
}

type PaymentsCheckGiftCodeRequest struct {
	Slug string
}

func (*PaymentsCheckGiftCodeRequest) CRC() uint32 {
	return 0x8e51b4c1
}

func PaymentsCheckGiftCode(ctx context.Context, m Requester, i PaymentsCheckGiftCodeRequest) (IPaymentsCheckedGiftCode, error) {
	var res IPaymentsCheckedGiftCode
	return res, request(ctx, m, &i, &res)
}

type PaymentsApplyGiftCodeRequest struct {
	Slug string
}

func (*PaymentsApplyGiftCodeRequest) CRC() uint32 {
	return 0xf6e26854
}

func PaymentsApplyGiftCode(ctx context.Context, m Requester, i PaymentsApplyGiftCodeRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetGiveawayInfoRequest struct {
	Peer  IInputPeer
	MsgID int32
}

func (*PaymentsGetGiveawayInfoRequest) CRC() uint32 {
	return 0xf4239425
}

func PaymentsGetGiveawayInfo(ctx context.Context, m Requester, i PaymentsGetGiveawayInfoRequest) (IPaymentsGiveawayInfo, error) {
	var res IPaymentsGiveawayInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsLaunchPrepaidGiveawayRequest struct {
	Peer       IInputPeer
	GiveawayID int64
	Purpose    IInputStorePaymentPurpose
}

func (*PaymentsLaunchPrepaidGiveawayRequest) CRC() uint32 {
	return 0x5ff58f20
}

func PaymentsLaunchPrepaidGiveaway(ctx context.Context, m Requester, i PaymentsLaunchPrepaidGiveawayRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetCallConfigRequest struct{}

func (*PhoneGetCallConfigRequest) CRC() uint32 {
	return 0x55451fa9
}

func PhoneGetCallConfig(ctx context.Context, m Requester, i PhoneGetCallConfigRequest) (IDataJSON, error) {
	var res IDataJSON
	return res, request(ctx, m, &i, &res)
}

type PhoneRequestCallRequest struct {
	_        struct{} `tl:"flags,bitflag"`
	Video    bool     `tl:",omitempty:flags:0,implicit"`
	UserID   IInputUser
	RandomID int32
	GAHash   []byte
	Protocol IPhoneCallProtocol
}

func (*PhoneRequestCallRequest) CRC() uint32 {
	return 0x42ff96ed
}

func PhoneRequestCall(ctx context.Context, m Requester, i PhoneRequestCallRequest) (IPhonePhoneCall, error) {
	var res IPhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneAcceptCallRequest struct {
	Peer     IInputPhoneCall
	GB       []byte
	Protocol IPhoneCallProtocol
}

func (*PhoneAcceptCallRequest) CRC() uint32 {
	return 0x3bd2b4a0
}

func PhoneAcceptCall(ctx context.Context, m Requester, i PhoneAcceptCallRequest) (IPhonePhoneCall, error) {
	var res IPhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneConfirmCallRequest struct {
	Peer           IInputPhoneCall
	GA             []byte
	KeyFingerprint int64
	Protocol       IPhoneCallProtocol
}

func (*PhoneConfirmCallRequest) CRC() uint32 {
	return 0x2efe1722
}

func PhoneConfirmCall(ctx context.Context, m Requester, i PhoneConfirmCallRequest) (IPhonePhoneCall, error) {
	var res IPhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneReceivedCallRequest struct {
	Peer IInputPhoneCall
}

func (*PhoneReceivedCallRequest) CRC() uint32 {
	return 0x17d54f61
}

func PhoneReceivedCall(ctx context.Context, m Requester, i PhoneReceivedCallRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardCallRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Video        bool     `tl:",omitempty:flags:0,implicit"`
	Peer         IInputPhoneCall
	Duration     int32
	Reason       EnumPhoneCallDiscardReason
	ConnectionID int64
}

func (*PhoneDiscardCallRequest) CRC() uint32 {
	return 0xb2cbc1c0
}

func PhoneDiscardCall(ctx context.Context, m Requester, i PhoneDiscardCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneSetCallRatingRequest struct {
	_              struct{} `tl:"flags,bitflag"`
	UserInitiative bool     `tl:",omitempty:flags:0,implicit"`
	Peer           IInputPhoneCall
	Rating         int32
	Comment        string
}

func (*PhoneSetCallRatingRequest) CRC() uint32 {
	return 0x59ead627
}

func PhoneSetCallRating(ctx context.Context, m Requester, i PhoneSetCallRatingRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallDebugRequest struct {
	Peer  IInputPhoneCall
	Debug IDataJSON
}

func (*PhoneSaveCallDebugRequest) CRC() uint32 {
	return 0x277add7e
}

func PhoneSaveCallDebug(ctx context.Context, m Requester, i PhoneSaveCallDebugRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneSendSignalingDataRequest struct {
	Peer IInputPhoneCall
	Data []byte
}

func (*PhoneSendSignalingDataRequest) CRC() uint32 {
	return 0xff7a9383
}

func PhoneSendSignalingData(ctx context.Context, m Requester, i PhoneSendSignalingDataRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneCreateGroupCallRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	RtmpStream   bool     `tl:",omitempty:flags:2,implicit"`
	Peer         IInputPeer
	RandomID     int32
	Title        *string `tl:",omitempty:flags:0"`
	ScheduleDate *int32  `tl:",omitempty:flags:1"`
}

func (*PhoneCreateGroupCallRequest) CRC() uint32 {
	return 0x48cdc6d8
}

func PhoneCreateGroupCall(ctx context.Context, m Requester, i PhoneCreateGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Muted        bool     `tl:",omitempty:flags:0,implicit"`
	VideoStopped bool     `tl:",omitempty:flags:2,implicit"`
	Call         IInputGroupCall
	JoinAs       IInputPeer
	InviteHash   *string `tl:",omitempty:flags:1"`
	Params       IDataJSON
}

func (*PhoneJoinGroupCallRequest) CRC() uint32 {
	return 0xb132ff7b
}

func PhoneJoinGroupCall(ctx context.Context, m Requester, i PhoneJoinGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallRequest struct {
	Call   IInputGroupCall
	Source int32
}

func (*PhoneLeaveGroupCallRequest) CRC() uint32 {
	return 0x500377f9
}

func PhoneLeaveGroupCall(ctx context.Context, m Requester, i PhoneLeaveGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneInviteToGroupCallRequest struct {
	Call  IInputGroupCall
	Users []IInputUser
}

func (*PhoneInviteToGroupCallRequest) CRC() uint32 {
	return 0x7b393160
}

func PhoneInviteToGroupCall(ctx context.Context, m Requester, i PhoneInviteToGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardGroupCallRequest struct {
	Call IInputGroupCall
}

func (*PhoneDiscardGroupCallRequest) CRC() uint32 {
	return 0x7a777135
}

func PhoneDiscardGroupCall(ctx context.Context, m Requester, i PhoneDiscardGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallSettingsRequest struct {
	_               struct{} `tl:"flags,bitflag"`
	ResetInviteHash bool     `tl:",omitempty:flags:1,implicit"`
	Call            IInputGroupCall
	JoinMuted       *bool `tl:",omitempty:flags:0"`
}

func (*PhoneToggleGroupCallSettingsRequest) CRC() uint32 {
	return 0x74bbb43d
}

func PhoneToggleGroupCallSettings(ctx context.Context, m Requester, i PhoneToggleGroupCallSettingsRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallRequest struct {
	Call  IInputGroupCall
	Limit int32
}

func (*PhoneGetGroupCallRequest) CRC() uint32 {
	return 0x41845db
}

func PhoneGetGroupCall(ctx context.Context, m Requester, i PhoneGetGroupCallRequest) (IPhoneGroupCall, error) {
	var res IPhoneGroupCall
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupParticipantsRequest struct {
	Call    IInputGroupCall
	Ids     []IInputPeer
	Sources []int32
	Offset  string
	Limit   int32
}

func (*PhoneGetGroupParticipantsRequest) CRC() uint32 {
	return 0xc558d8ab
}

func PhoneGetGroupParticipants(ctx context.Context, m Requester, i PhoneGetGroupParticipantsRequest) (IPhoneGroupParticipants, error) {
	var res IPhoneGroupParticipants
	return res, request(ctx, m, &i, &res)
}

type PhoneCheckGroupCallRequest struct {
	Call    IInputGroupCall
	Sources []int32
}

func (*PhoneCheckGroupCallRequest) CRC() uint32 {
	return 0xb59cf977
}

func PhoneCheckGroupCall(ctx context.Context, m Requester, i PhoneCheckGroupCallRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallRecordRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	Start         bool     `tl:",omitempty:flags:0,implicit"`
	Video         bool     `tl:",omitempty:flags:2,implicit"`
	Call          IInputGroupCall
	Title         *string `tl:",omitempty:flags:1"`
	VideoPortrait *bool   `tl:",omitempty:flags:2"`
}

func (*PhoneToggleGroupCallRecordRequest) CRC() uint32 {
	return 0xf128c708
}

func PhoneToggleGroupCallRecord(ctx context.Context, m Requester, i PhoneToggleGroupCallRecordRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallParticipantRequest struct {
	_                  struct{} `tl:"flags,bitflag"`
	Call               IInputGroupCall
	Participant        IInputPeer
	Muted              *bool  `tl:",omitempty:flags:0"`
	Volume             *int32 `tl:",omitempty:flags:1"`
	RaiseHand          *bool  `tl:",omitempty:flags:2"`
	VideoStopped       *bool  `tl:",omitempty:flags:3"`
	VideoPaused        *bool  `tl:",omitempty:flags:4"`
	PresentationPaused *bool  `tl:",omitempty:flags:5"`
}

func (*PhoneEditGroupCallParticipantRequest) CRC() uint32 {
	return 0xa5273abf
}

func PhoneEditGroupCallParticipant(ctx context.Context, m Requester, i PhoneEditGroupCallParticipantRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallTitleRequest struct {
	Call  IInputGroupCall
	Title string
}

func (*PhoneEditGroupCallTitleRequest) CRC() uint32 {
	return 0x1ca6ac0a
}

func PhoneEditGroupCallTitle(ctx context.Context, m Requester, i PhoneEditGroupCallTitleRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallJoinAsRequest struct {
	Peer IInputPeer
}

func (*PhoneGetGroupCallJoinAsRequest) CRC() uint32 {
	return 0xef7c213a
}

func PhoneGetGroupCallJoinAs(ctx context.Context, m Requester, i PhoneGetGroupCallJoinAsRequest) (IPhoneJoinAsPeers, error) {
	var res IPhoneJoinAsPeers
	return res, request(ctx, m, &i, &res)
}

type PhoneExportGroupCallInviteRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	CanSelfUnmute bool     `tl:",omitempty:flags:0,implicit"`
	Call          IInputGroupCall
}

func (*PhoneExportGroupCallInviteRequest) CRC() uint32 {
	return 0xe6aa647f
}

func PhoneExportGroupCallInvite(ctx context.Context, m Requester, i PhoneExportGroupCallInviteRequest) (IPhoneExportedGroupCallInvite, error) {
	var res IPhoneExportedGroupCallInvite
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallStartSubscriptionRequest struct {
	Call       IInputGroupCall
	Subscribed bool
}

func (*PhoneToggleGroupCallStartSubscriptionRequest) CRC() uint32 {
	return 0x219c34e6
}

func PhoneToggleGroupCallStartSubscription(ctx context.Context, m Requester, i PhoneToggleGroupCallStartSubscriptionRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneStartScheduledGroupCallRequest struct {
	Call IInputGroupCall
}

func (*PhoneStartScheduledGroupCallRequest) CRC() uint32 {
	return 0x5680e342
}

func PhoneStartScheduledGroupCall(ctx context.Context, m Requester, i PhoneStartScheduledGroupCallRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveDefaultGroupCallJoinAsRequest struct {
	Peer   IInputPeer
	JoinAs IInputPeer
}

func (*PhoneSaveDefaultGroupCallJoinAsRequest) CRC() uint32 {
	return 0x575e1f8c
}

func PhoneSaveDefaultGroupCallJoinAs(ctx context.Context, m Requester, i PhoneSaveDefaultGroupCallJoinAsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallPresentationRequest struct {
	Call   IInputGroupCall
	Params IDataJSON
}

func (*PhoneJoinGroupCallPresentationRequest) CRC() uint32 {
	return 0xcbea6bc4
}

func PhoneJoinGroupCallPresentation(ctx context.Context, m Requester, i PhoneJoinGroupCallPresentationRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallPresentationRequest struct {
	Call IInputGroupCall
}

func (*PhoneLeaveGroupCallPresentationRequest) CRC() uint32 {
	return 0x1c50d144
}

func PhoneLeaveGroupCallPresentation(ctx context.Context, m Requester, i PhoneLeaveGroupCallPresentationRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamChannelsRequest struct {
	Call IInputGroupCall
}

func (*PhoneGetGroupCallStreamChannelsRequest) CRC() uint32 {
	return 0x1ab21940
}

func PhoneGetGroupCallStreamChannels(ctx context.Context, m Requester, i PhoneGetGroupCallStreamChannelsRequest) (IPhoneGroupCallStreamChannels, error) {
	var res IPhoneGroupCallStreamChannels
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamRtmpURLRequest struct {
	Peer   IInputPeer
	Revoke bool
}

func (*PhoneGetGroupCallStreamRtmpURLRequest) CRC() uint32 {
	return 0xdeb3abbf
}

func PhoneGetGroupCallStreamRtmpURL(ctx context.Context, m Requester, i PhoneGetGroupCallStreamRtmpURLRequest) (IPhoneGroupCallStreamRtmpURL, error) {
	var res IPhoneGroupCallStreamRtmpURL
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallLogRequest struct {
	Peer IInputPhoneCall
	File IInputFile
}

func (*PhoneSaveCallLogRequest) CRC() uint32 {
	return 0x41248786
}

func PhoneSaveCallLog(ctx context.Context, m Requester, i PhoneSaveCallLogRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhotosUpdateProfilePhotoRequest struct {
	_        struct{}   `tl:"flags,bitflag"`
	Fallback bool       `tl:",omitempty:flags:0,implicit"`
	Bot      IInputUser `tl:",omitempty:flags:1"`
	ID       IInputPhoto
}

func (*PhotosUpdateProfilePhotoRequest) CRC() uint32 {
	return 0x9e82039
}

func PhotosUpdateProfilePhoto(ctx context.Context, m Requester, i PhotosUpdateProfilePhotoRequest) (IPhotosPhoto, error) {
	var res IPhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadProfilePhotoRequest struct {
	_                struct{}   `tl:"flags,bitflag"`
	Fallback         bool       `tl:",omitempty:flags:3,implicit"`
	Bot              IInputUser `tl:",omitempty:flags:5"`
	File             IInputFile `tl:",omitempty:flags:0"`
	Video            IInputFile `tl:",omitempty:flags:1"`
	VideoStartTs     *float64   `tl:",omitempty:flags:2"`
	VideoEmojiMarkup IVideoSize `tl:",omitempty:flags:4"`
}

func (*PhotosUploadProfilePhotoRequest) CRC() uint32 {
	return 0x388a3b5
}

func PhotosUploadProfilePhoto(ctx context.Context, m Requester, i PhotosUploadProfilePhotoRequest) (IPhotosPhoto, error) {
	var res IPhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosDeletePhotosRequest struct {
	ID []IInputPhoto
}

func (*PhotosDeletePhotosRequest) CRC() uint32 {
	return 0x87cf7f2f
}

func PhotosDeletePhotos(ctx context.Context, m Requester, i PhotosDeletePhotosRequest) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type PhotosGetUserPhotosRequest struct {
	UserID IInputUser
	Offset int32
	MaxID  int64
	Limit  int32
}

func (*PhotosGetUserPhotosRequest) CRC() uint32 {
	return 0x91cd32a8
}

func PhotosGetUserPhotos(ctx context.Context, m Requester, i PhotosGetUserPhotosRequest) (IPhotosPhotos, error) {
	var res IPhotosPhotos
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadContactProfilePhotoRequest struct {
	_                struct{} `tl:"flags,bitflag"`
	Suggest          bool     `tl:",omitempty:flags:3,implicit"`
	Save             bool     `tl:",omitempty:flags:4,implicit"`
	UserID           IInputUser
	File             IInputFile `tl:",omitempty:flags:0"`
	Video            IInputFile `tl:",omitempty:flags:1"`
	VideoStartTs     *float64   `tl:",omitempty:flags:2"`
	VideoEmojiMarkup IVideoSize `tl:",omitempty:flags:5"`
}

func (*PhotosUploadContactProfilePhotoRequest) CRC() uint32 {
	return 0xe14c4a71
}

func PhotosUploadContactProfilePhoto(ctx context.Context, m Requester, i PhotosUploadContactProfilePhotoRequest) (IPhotosPhoto, error) {
	var res IPhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsListRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Gifts  bool     `tl:",omitempty:flags:0,implicit"`
	Peer   IInputPeer
	Offset string
	Limit  int32
}

func (*PremiumGetBoostsListRequest) CRC() uint32 {
	return 0x60f67660
}

func PremiumGetBoostsList(ctx context.Context, m Requester, i PremiumGetBoostsListRequest) (IPremiumBoostsList, error) {
	var res IPremiumBoostsList
	return res, request(ctx, m, &i, &res)
}

type PremiumGetMyBoostsRequest struct{}

func (*PremiumGetMyBoostsRequest) CRC() uint32 {
	return 0xbe77b4a
}

func PremiumGetMyBoosts(ctx context.Context, m Requester, i PremiumGetMyBoostsRequest) (IPremiumMyBoosts, error) {
	var res IPremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumApplyBoostRequest struct {
	_     struct{} `tl:"flags,bitflag"`
	Slots []int32  `tl:",omitempty:flags:0"`
	Peer  IInputPeer
}

func (*PremiumApplyBoostRequest) CRC() uint32 {
	return 0x6b7da746
}

func PremiumApplyBoost(ctx context.Context, m Requester, i PremiumApplyBoostRequest) (IPremiumMyBoosts, error) {
	var res IPremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsStatusRequest struct {
	Peer IInputPeer
}

func (*PremiumGetBoostsStatusRequest) CRC() uint32 {
	return 0x42f1f61
}

func PremiumGetBoostsStatus(ctx context.Context, m Requester, i PremiumGetBoostsStatusRequest) (IPremiumBoostsStatus, error) {
	var res IPremiumBoostsStatus
	return res, request(ctx, m, &i, &res)
}

type PremiumGetUserBoostsRequest struct {
	Peer   IInputPeer
	UserID IInputUser
}

func (*PremiumGetUserBoostsRequest) CRC() uint32 {
	return 0x39854d1f
}

func PremiumGetUserBoosts(ctx context.Context, m Requester, i PremiumGetUserBoostsRequest) (IPremiumBoostsList, error) {
	var res IPremiumBoostsList
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastStatsRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel IInputChannel
}

func (*StatsGetBroadcastStatsRequest) CRC() uint32 {
	return 0xab42441a
}

func StatsGetBroadcastStats(ctx context.Context, m Requester, i StatsGetBroadcastStatsRequest) (IStatsBroadcastStats, error) {
	var res IStatsBroadcastStats
	return res, request(ctx, m, &i, &res)
}

type StatsLoadAsyncGraphRequest struct {
	_     struct{} `tl:"flags,bitflag"`
	Token string
	X     *int64 `tl:",omitempty:flags:0"`
}

func (*StatsLoadAsyncGraphRequest) CRC() uint32 {
	return 0x621d5fa0
}

func StatsLoadAsyncGraph(ctx context.Context, m Requester, i StatsLoadAsyncGraphRequest) (IStatsGraph, error) {
	var res IStatsGraph
	return res, request(ctx, m, &i, &res)
}

type StatsGetMegagroupStatsRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel IInputChannel
}

func (*StatsGetMegagroupStatsRequest) CRC() uint32 {
	return 0xdcdf8607
}

func StatsGetMegagroupStats(ctx context.Context, m Requester, i StatsGetMegagroupStatsRequest) (IStatsMegagroupStats, error) {
	var res IStatsMegagroupStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessagePublicForwardsRequest struct {
	Channel IInputChannel
	MsgID   int32
	Offset  string
	Limit   int32
}

func (*StatsGetMessagePublicForwardsRequest) CRC() uint32 {
	return 0x5f150144
}

func StatsGetMessagePublicForwards(ctx context.Context, m Requester, i StatsGetMessagePublicForwardsRequest) (IStatsPublicForwards, error) {
	var res IStatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessageStatsRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel IInputChannel
	MsgID   int32
}

func (*StatsGetMessageStatsRequest) CRC() uint32 {
	return 0xb6e0a3f5
}

func StatsGetMessageStats(ctx context.Context, m Requester, i StatsGetMessageStatsRequest) (IStatsMessageStats, error) {
	var res IStatsMessageStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryStatsRequest struct {
	_    struct{} `tl:"flags,bitflag"`
	Dark bool     `tl:",omitempty:flags:0,implicit"`
	Peer IInputPeer
	ID   int32
}

func (*StatsGetStoryStatsRequest) CRC() uint32 {
	return 0x374fef40
}

func StatsGetStoryStats(ctx context.Context, m Requester, i StatsGetStoryStatsRequest) (IStatsStoryStats, error) {
	var res IStatsStoryStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryPublicForwardsRequest struct {
	Peer   IInputPeer
	ID     int32
	Offset string
	Limit  int32
}

func (*StatsGetStoryPublicForwardsRequest) CRC() uint32 {
	return 0xa6437ef6
}

func StatsGetStoryPublicForwards(ctx context.Context, m Requester, i StatsGetStoryPublicForwardsRequest) (IStatsPublicForwards, error) {
	var res IStatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StickersCreateStickerSetRequest struct {
	_         struct{} `tl:"flags,bitflag"`
	Masks     bool     `tl:",omitempty:flags:0,implicit"`
	Animated  bool     `tl:",omitempty:flags:1,implicit"`
	Videos    bool     `tl:",omitempty:flags:4,implicit"`
	Emojis    bool     `tl:",omitempty:flags:5,implicit"`
	TextColor bool     `tl:",omitempty:flags:6,implicit"`
	UserID    IInputUser
	Title     string
	ShortName string
	Thumb     IInputDocument `tl:",omitempty:flags:2"`
	Stickers  []IInputStickerSetItem
	Software  *string `tl:",omitempty:flags:3"`
}

func (*StickersCreateStickerSetRequest) CRC() uint32 {
	return 0x9021ab67
}

func StickersCreateStickerSet(ctx context.Context, m Requester, i StickersCreateStickerSetRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRemoveStickerFromSetRequest struct {
	Sticker IInputDocument
}

func (*StickersRemoveStickerFromSetRequest) CRC() uint32 {
	return 0xf7760f51
}

func StickersRemoveStickerFromSet(ctx context.Context, m Requester, i StickersRemoveStickerFromSetRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerPositionRequest struct {
	Sticker  IInputDocument
	Position int32
}

func (*StickersChangeStickerPositionRequest) CRC() uint32 {
	return 0xffb6d4ca
}

func StickersChangeStickerPosition(ctx context.Context, m Requester, i StickersChangeStickerPositionRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersAddStickerToSetRequest struct {
	Stickerset IInputStickerSet
	Sticker    IInputStickerSetItem
}

func (*StickersAddStickerToSetRequest) CRC() uint32 {
	return 0x8653febe
}

func StickersAddStickerToSet(ctx context.Context, m Requester, i StickersAddStickerToSetRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersSetStickerSetThumbRequest struct {
	_               struct{} `tl:"flags,bitflag"`
	Stickerset      IInputStickerSet
	Thumb           IInputDocument `tl:",omitempty:flags:0"`
	ThumbDocumentID *int64         `tl:",omitempty:flags:1"`
}

func (*StickersSetStickerSetThumbRequest) CRC() uint32 {
	return 0xa76a5392
}

func StickersSetStickerSetThumb(ctx context.Context, m Requester, i StickersSetStickerSetThumbRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersCheckShortNameRequest struct {
	ShortName string
}

func (*StickersCheckShortNameRequest) CRC() uint32 {
	return 0x284b3639
}

func StickersCheckShortName(ctx context.Context, m Requester, i StickersCheckShortNameRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StickersSuggestShortNameRequest struct {
	Title string
}

func (*StickersSuggestShortNameRequest) CRC() uint32 {
	return 0x4dafc503
}

func StickersSuggestShortName(ctx context.Context, m Requester, i StickersSuggestShortNameRequest) (IStickersSuggestedShortName, error) {
	var res IStickersSuggestedShortName
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerRequest struct {
	_          struct{} `tl:"flags,bitflag"`
	Sticker    IInputDocument
	Emoji      *string     `tl:",omitempty:flags:0"`
	MaskCoords IMaskCoords `tl:",omitempty:flags:1"`
	Keywords   *string     `tl:",omitempty:flags:2"`
}

func (*StickersChangeStickerRequest) CRC() uint32 {
	return 0xf5537ebc
}

func StickersChangeSticker(ctx context.Context, m Requester, i StickersChangeStickerRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRenameStickerSetRequest struct {
	Stickerset IInputStickerSet
	Title      string
}

func (*StickersRenameStickerSetRequest) CRC() uint32 {
	return 0x124b1c00
}

func StickersRenameStickerSet(ctx context.Context, m Requester, i StickersRenameStickerSetRequest) (IMessagesStickerSet, error) {
	var res IMessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersDeleteStickerSetRequest struct {
	Stickerset IInputStickerSet
}

func (*StickersDeleteStickerSetRequest) CRC() uint32 {
	return 0x87704394
}

func StickersDeleteStickerSet(ctx context.Context, m Requester, i StickersDeleteStickerSetRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesCanSendStoryRequest struct {
	Peer IInputPeer
}

func (*StoriesCanSendStoryRequest) CRC() uint32 {
	return 0xc7dfdfdd
}

func StoriesCanSendStory(ctx context.Context, m Requester, i StoriesCanSendStoryRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesSendStoryRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Pinned       bool     `tl:",omitempty:flags:2,implicit"`
	Noforwards   bool     `tl:",omitempty:flags:4,implicit"`
	FwdModified  bool     `tl:",omitempty:flags:7,implicit"`
	Peer         IInputPeer
	Media        IInputMedia
	MediaAreas   []IMediaArea     `tl:",omitempty:flags:5"`
	Caption      *string          `tl:",omitempty:flags:0"`
	Entities     []IMessageEntity `tl:",omitempty:flags:1"`
	PrivacyRules []IInputPrivacyRule
	RandomID     int64
	Period       *int32     `tl:",omitempty:flags:3"`
	FwdFromID    IInputPeer `tl:",omitempty:flags:6"`
	FwdFromStory *int32     `tl:",omitempty:flags:6"`
}

func (*StoriesSendStoryRequest) CRC() uint32 {
	return 0xe4e6694b
}

func StoriesSendStory(ctx context.Context, m Requester, i StoriesSendStoryRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type StoriesEditStoryRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Peer         IInputPeer
	ID           int32
	Media        IInputMedia         `tl:",omitempty:flags:0"`
	MediaAreas   []IMediaArea        `tl:",omitempty:flags:3"`
	Caption      *string             `tl:",omitempty:flags:1"`
	Entities     []IMessageEntity    `tl:",omitempty:flags:1"`
	PrivacyRules []IInputPrivacyRule `tl:",omitempty:flags:2"`
}

func (*StoriesEditStoryRequest) CRC() uint32 {
	return 0xb583ba46
}

func StoriesEditStory(ctx context.Context, m Requester, i StoriesEditStoryRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type StoriesDeleteStoriesRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*StoriesDeleteStoriesRequest) CRC() uint32 {
	return 0xae59db5f
}

func StoriesDeleteStories(ctx context.Context, m Requester, i StoriesDeleteStoriesRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePinnedRequest struct {
	Peer   IInputPeer
	ID     []int32
	Pinned bool
}

func (*StoriesTogglePinnedRequest) CRC() uint32 {
	return 0x9a75a1ef
}

func StoriesTogglePinned(ctx context.Context, m Requester, i StoriesTogglePinnedRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesGetAllStoriesRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Next   bool     `tl:",omitempty:flags:1,implicit"`
	Hidden bool     `tl:",omitempty:flags:2,implicit"`
	State  *string  `tl:",omitempty:flags:0"`
}

func (*StoriesGetAllStoriesRequest) CRC() uint32 {
	return 0xeeb0d625
}

func StoriesGetAllStories(ctx context.Context, m Requester, i StoriesGetAllStoriesRequest) (IStoriesAllStories, error) {
	var res IStoriesAllStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPinnedStoriesRequest struct {
	Peer     IInputPeer
	OffsetID int32
	Limit    int32
}

func (*StoriesGetPinnedStoriesRequest) CRC() uint32 {
	return 0x5821a5dc
}

func StoriesGetPinnedStories(ctx context.Context, m Requester, i StoriesGetPinnedStoriesRequest) (IStoriesStories, error) {
	var res IStoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesArchiveRequest struct {
	Peer     IInputPeer
	OffsetID int32
	Limit    int32
}

func (*StoriesGetStoriesArchiveRequest) CRC() uint32 {
	return 0xb4352016
}

func StoriesGetStoriesArchive(ctx context.Context, m Requester, i StoriesGetStoriesArchiveRequest) (IStoriesStories, error) {
	var res IStoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesByIDRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*StoriesGetStoriesByIDRequest) CRC() uint32 {
	return 0x5774ca74
}

func StoriesGetStoriesByID(ctx context.Context, m Requester, i StoriesGetStoriesByIDRequest) (IStoriesStories, error) {
	var res IStoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesToggleAllStoriesHiddenRequest struct {
	Hidden bool
}

func (*StoriesToggleAllStoriesHiddenRequest) CRC() uint32 {
	return 0x7c2557c4
}

func StoriesToggleAllStoriesHidden(ctx context.Context, m Requester, i StoriesToggleAllStoriesHiddenRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesReadStoriesRequest struct {
	Peer  IInputPeer
	MaxID int32
}

func (*StoriesReadStoriesRequest) CRC() uint32 {
	return 0xa556dac8
}

func StoriesReadStories(ctx context.Context, m Requester, i StoriesReadStoriesRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesIncrementStoryViewsRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*StoriesIncrementStoryViewsRequest) CRC() uint32 {
	return 0xb2028afb
}

func StoriesIncrementStoryViews(ctx context.Context, m Requester, i StoriesIncrementStoryViewsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryViewsListRequest struct {
	_              struct{} `tl:"flags,bitflag"`
	JustContacts   bool     `tl:",omitempty:flags:0,implicit"`
	ReactionsFirst bool     `tl:",omitempty:flags:2,implicit"`
	ForwardsFirst  bool     `tl:",omitempty:flags:3,implicit"`
	Peer           IInputPeer
	Q              *string `tl:",omitempty:flags:1"`
	ID             int32
	Offset         string
	Limit          int32
}

func (*StoriesGetStoryViewsListRequest) CRC() uint32 {
	return 0x7ed23c57
}

func StoriesGetStoryViewsList(ctx context.Context, m Requester, i StoriesGetStoryViewsListRequest) (IStoriesStoryViewsList, error) {
	var res IStoriesStoryViewsList
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesViewsRequest struct {
	Peer IInputPeer
	ID   []int32
}

func (*StoriesGetStoriesViewsRequest) CRC() uint32 {
	return 0x28e16cc8
}

func StoriesGetStoriesViews(ctx context.Context, m Requester, i StoriesGetStoriesViewsRequest) (IStoriesStoryViews, error) {
	var res IStoriesStoryViews
	return res, request(ctx, m, &i, &res)
}

type StoriesExportStoryLinkRequest struct {
	Peer IInputPeer
	ID   int32
}

func (*StoriesExportStoryLinkRequest) CRC() uint32 {
	return 0x7b8def20
}

func StoriesExportStoryLink(ctx context.Context, m Requester, i StoriesExportStoryLinkRequest) (IExportedStoryLink, error) {
	var res IExportedStoryLink
	return res, request(ctx, m, &i, &res)
}

type StoriesReportRequest struct {
	Peer    IInputPeer
	ID      []int32
	Reason  EnumReportReason
	Message string
}

func (*StoriesReportRequest) CRC() uint32 {
	return 0x1923fa8c
}

func StoriesReport(ctx context.Context, m Requester, i StoriesReportRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesActivateStealthModeRequest struct {
	_      struct{} `tl:"flags,bitflag"`
	Past   bool     `tl:",omitempty:flags:0,implicit"`
	Future bool     `tl:",omitempty:flags:1,implicit"`
}

func (*StoriesActivateStealthModeRequest) CRC() uint32 {
	return 0x57bbd166
}

func StoriesActivateStealthMode(ctx context.Context, m Requester, i StoriesActivateStealthModeRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type StoriesSendReactionRequest struct {
	_           struct{} `tl:"flags,bitflag"`
	AddToRecent bool     `tl:",omitempty:flags:0,implicit"`
	Peer        IInputPeer
	StoryID     int32
	Reaction    IReaction
}

func (*StoriesSendReactionRequest) CRC() uint32 {
	return 0x7fd736b2
}

func StoriesSendReaction(ctx context.Context, m Requester, i StoriesSendReactionRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPeerStoriesRequest struct {
	Peer IInputPeer
}

func (*StoriesGetPeerStoriesRequest) CRC() uint32 {
	return 0x2c4ada50
}

func StoriesGetPeerStories(ctx context.Context, m Requester, i StoriesGetPeerStoriesRequest) (IStoriesPeerStories, error) {
	var res IStoriesPeerStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetAllReadPeerStoriesRequest struct{}

func (*StoriesGetAllReadPeerStoriesRequest) CRC() uint32 {
	return 0x9b5ae7f9
}

func StoriesGetAllReadPeerStories(ctx context.Context, m Requester, i StoriesGetAllReadPeerStoriesRequest) (IUpdates, error) {
	var res IUpdates
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPeerMaxIDsRequest struct {
	ID []IInputPeer
}

func (*StoriesGetPeerMaxIDsRequest) CRC() uint32 {
	return 0x535983c3
}

func StoriesGetPeerMaxIDs(ctx context.Context, m Requester, i StoriesGetPeerMaxIDsRequest) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesGetChatsToSendRequest struct{}

func (*StoriesGetChatsToSendRequest) CRC() uint32 {
	return 0xa56a8b60
}

func StoriesGetChatsToSend(ctx context.Context, m Requester, i StoriesGetChatsToSendRequest) (IMessagesChats, error) {
	var res IMessagesChats
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePeerStoriesHiddenRequest struct {
	Peer   IInputPeer
	Hidden bool
}

func (*StoriesTogglePeerStoriesHiddenRequest) CRC() uint32 {
	return 0xbd0415c4
}

func StoriesTogglePeerStoriesHidden(ctx context.Context, m Requester, i StoriesTogglePeerStoriesHiddenRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryReactionsListRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	ForwardsFirst bool     `tl:",omitempty:flags:2,implicit"`
	Peer          IInputPeer
	ID            int32
	Reaction      IReaction `tl:",omitempty:flags:0"`
	Offset        *string   `tl:",omitempty:flags:1"`
	Limit         int32
}

func (*StoriesGetStoryReactionsListRequest) CRC() uint32 {
	return 0xb9b2881f
}

func StoriesGetStoryReactionsList(ctx context.Context, m Requester, i StoriesGetStoryReactionsListRequest) (IStoriesStoryReactionsList, error) {
	var res IStoriesStoryReactionsList
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetStateRequest struct{}

func (*UpdatesGetStateRequest) CRC() uint32 {
	return 0xedd4882a
}

func UpdatesGetState(ctx context.Context, m Requester, i UpdatesGetStateRequest) (IUpdatesState, error) {
	var res IUpdatesState
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetDifferenceRequest struct {
	_             struct{} `tl:"flags,bitflag"`
	Pts           int32
	PtsLimit      *int32 `tl:",omitempty:flags:1"`
	PtsTotalLimit *int32 `tl:",omitempty:flags:0"`
	Date          int32
	Qts           int32
	QtsLimit      *int32 `tl:",omitempty:flags:2"`
}

func (*UpdatesGetDifferenceRequest) CRC() uint32 {
	return 0x19c2f763
}

func UpdatesGetDifference(ctx context.Context, m Requester, i UpdatesGetDifferenceRequest) (IUpdatesDifference, error) {
	var res IUpdatesDifference
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetChannelDifferenceRequest struct {
	_       struct{} `tl:"flags,bitflag"`
	Force   bool     `tl:",omitempty:flags:0,implicit"`
	Channel IInputChannel
	Filter  IChannelMessagesFilter
	Pts     int32
	Limit   int32
}

func (*UpdatesGetChannelDifferenceRequest) CRC() uint32 {
	return 0x3173d78
}

func UpdatesGetChannelDifference(ctx context.Context, m Requester, i UpdatesGetChannelDifferenceRequest) (IUpdatesChannelDifference, error) {
	var res IUpdatesChannelDifference
	return res, request(ctx, m, &i, &res)
}

type UploadSaveFilePartRequest struct {
	FileID   int64
	FilePart int32
	Bytes    []byte
}

func (*UploadSaveFilePartRequest) CRC() uint32 {
	return 0xb304a621
}

func UploadSaveFilePart(ctx context.Context, m Requester, i UploadSaveFilePartRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileRequest struct {
	_            struct{} `tl:"flags,bitflag"`
	Precise      bool     `tl:",omitempty:flags:0,implicit"`
	CdnSupported bool     `tl:",omitempty:flags:1,implicit"`
	Location     IInputFileLocation
	Offset       int64
	Limit        int32
}

func (*UploadGetFileRequest) CRC() uint32 {
	return 0xbe5335be
}

func UploadGetFile(ctx context.Context, m Requester, i UploadGetFileRequest) (IUploadFile, error) {
	var res IUploadFile
	return res, request(ctx, m, &i, &res)
}

type UploadSaveBigFilePartRequest struct {
	FileID         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

func (*UploadSaveBigFilePartRequest) CRC() uint32 {
	return 0xde7b673d
}

func UploadSaveBigFilePart(ctx context.Context, m Requester, i UploadSaveBigFilePartRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetWebFileRequest struct {
	Location IInputWebFileLocation
	Offset   int32
	Limit    int32
}

func (*UploadGetWebFileRequest) CRC() uint32 {
	return 0x24e6818d
}

func UploadGetWebFile(ctx context.Context, m Requester, i UploadGetWebFileRequest) (IUploadWebFile, error) {
	var res IUploadWebFile
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileRequest struct {
	FileToken []byte
	Offset    int64
	Limit     int32
}

func (*UploadGetCdnFileRequest) CRC() uint32 {
	return 0x395f69da
}

func UploadGetCdnFile(ctx context.Context, m Requester, i UploadGetCdnFileRequest) (IUploadCdnFile, error) {
	var res IUploadCdnFile
	return res, request(ctx, m, &i, &res)
}

type UploadReuploadCdnFileRequest struct {
	FileToken    []byte
	RequestToken []byte
}

func (*UploadReuploadCdnFileRequest) CRC() uint32 {
	return 0x9b2754a8
}

func UploadReuploadCdnFile(ctx context.Context, m Requester, i UploadReuploadCdnFileRequest) ([]IFileHash, error) {
	var res []IFileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileHashesRequest struct {
	FileToken []byte
	Offset    int64
}

func (*UploadGetCdnFileHashesRequest) CRC() uint32 {
	return 0x91dc3f31
}

func UploadGetCdnFileHashes(ctx context.Context, m Requester, i UploadGetCdnFileHashesRequest) ([]IFileHash, error) {
	var res []IFileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileHashesRequest struct {
	Location IInputFileLocation
	Offset   int64
}

func (*UploadGetFileHashesRequest) CRC() uint32 {
	return 0x9156982a
}

func UploadGetFileHashes(ctx context.Context, m Requester, i UploadGetFileHashesRequest) ([]IFileHash, error) {
	var res []IFileHash
	return res, request(ctx, m, &i, &res)
}

type UsersGetUsersRequest struct {
	ID []IInputUser
}

func (*UsersGetUsersRequest) CRC() uint32 {
	return 0xd91a548
}

func UsersGetUsers(ctx context.Context, m Requester, i UsersGetUsersRequest) ([]IUser, error) {
	var res []IUser
	return res, request(ctx, m, &i, &res)
}

type UsersGetFullUserRequest struct {
	ID IInputUser
}

func (*UsersGetFullUserRequest) CRC() uint32 {
	return 0xb60f5918
}

func UsersGetFullUser(ctx context.Context, m Requester, i UsersGetFullUserRequest) (IUsersUserFull, error) {
	var res IUsersUserFull
	return res, request(ctx, m, &i, &res)
}

type UsersSetSecureValueErrorsRequest struct {
	ID     IInputUser
	Errors []ISecureValueError
}

func (*UsersSetSecureValueErrorsRequest) CRC() uint32 {
	return 0x90c894b5
}

func UsersSetSecureValueErrors(ctx context.Context, m Requester, i UsersSetSecureValueErrorsRequest) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}
