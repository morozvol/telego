package objects

import (
	"bytes"
	"encoding/json"
	"io"
	mp "mime/multipart"
	"strconv"
	"strings"
)

type MethodArguments interface {
	ToJson() []byte
	ToMultiPart(wr *mp.Writer)
}

type GetUpdatesArgs struct {
	/*Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.*/
	Offset int `json:"offset,omitempty"`
	/*Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.*/
	Limit int `json:"limit,omitempty"`
	/*Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.*/
	Timeout int `json:"timeout,omitempty"`
	/*A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.
	Please note that this parameter doesnt affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.*/
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (args *GetUpdatesArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *GetUpdatesArgs) ToMultiPart(wr *mp.Writer) {
	//The arguments of this method are never passed as multipart.
}

type DefaultSendMethodsArguments struct {
	/*Unique identifier for the target chat or Username of the target channel (in the format @channelusername).*/
	ChatId json.RawMessage `json:"chat_id"`
	/*Sends the message silently. Users will receive a notification with no sound.*/
	DisableNotification bool `json:"disable_notification,omitempty"`
	/*If the message is a reply, ID of the original message*/
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	/*Pass True, if the message should be sent even if the specified replied-to message is not found*/
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	/*Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.*/
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (df *DefaultSendMethodsArguments) toMultiPart(wr *mp.Writer) {
	fw, _ := wr.CreateFormField("chat_id")
	_, _ = io.Copy(fw, strings.NewReader(string(df.ChatId)))
	fw, _ = wr.CreateFormField("disable_notification")
	_, _ = io.Copy(fw, strings.NewReader(strconv.FormatBool(df.DisableNotification)))
	if df.ReplyToMessageId != 0 {
		fw, _ = wr.CreateFormField("reply_to_message_id")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(df.ReplyToMessageId)))
	}
	fw, _ = wr.CreateFormField("allow_sending_without_reply")
	_, _ = io.Copy(fw, strings.NewReader(strconv.FormatBool(df.AllowSendingWithoutReply)))
	if df.ReplyMarkup != nil {
		fw, _ = wr.CreateFormField("reply_markup")
		bt, _ := json.Marshal(df.ReplyMarkup)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
}

type SendMessageArgs struct {
	DefaultSendMethodsArguments
	/*Text of the message to be sent, 1-4096 characters after entities parsing*/
	Text string `json:"text"`
	/*Mode for parsing entities in the message text. */
	ParseMode string `json:"parse_mode,omitempty"`
	/*A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode*/
	Entities []MessageEntity `json:"entities,omitempty"`
	/*Disables link previews for links in this message*/
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

func (args *SendMessageArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendMessageArgs) ToMultiPart(wr *mp.Writer) {
	//The arguments of this method are never passed as multipart.
}

type ForwardMessageArgs struct {
	/*Unique identifier for the target chat or Username of the target channel (in the format @channelusername).*/
	ChatId json.RawMessage `json:"chat_id"`
	/*Unique identifier for the chat where the original message was sent or Channel username in the format @channelusername*/
	FromChatId json.RawMessage `json:"from_chat_id"`
	/*Sends the message silently. Users will receive a notification with no sound.*/
	DisableNotification bool `json:"disable_notification,omitempty"`
	/*Message identifier in the chat specified in from_chat_id*/
	MessageId int `json:"message_id"`
}

func (args *ForwardMessageArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *ForwardMessageArgs) ToMultiPart(wr *mp.Writer) {
	//The arguments of this method are never passed as multipart.
}

type CopyMessageArgs struct {
	ForwardMessageArgs
	/*New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept*/
	Caption string `json:"caption,omitempty"`
	/*Mode for parsing entities in the message text. */
	ParseMode string `json:"parse_mode,omitempty"`
	/*A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode*/
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	/*If the message is a reply, ID of the original message*/
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	/*Pass True, if the message should be sent even if the specified replied-to message is not found*/
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	/*Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.*/
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (args *CopyMessageArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *CopyMessageArgs) ToMultiPart(wr *mp.Writer) {
	//The arguments of this method are never passed as multipart.
}

type SendPhotoArgs struct {
	DefaultSendMethodsArguments
	/*Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20.*/
	Photo string `json:"photo"`
	/*Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing*/
	Caption string `json:"caption,omitempty"`
	/*Mode for parsing entities in the photo caption. */
	ParseMode string `json:"parse_mode,omitempty"`
	/*A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode*/
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

func (args *SendPhotoArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendPhotoArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("photo")
	_, _ = io.Copy(fw, strings.NewReader(args.Photo))
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
}

type SendAudioArgs struct {
	DefaultSendMethodsArguments
	/*Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data.*/
	Audio           string          `json:"audio"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	/*Duration of the audio in secconds*/
	Duration  int    `json:"duration,omitempty"`
	Performer string `json:"performer,omitempty"`
	/*Track name*/
	Title string `json:"title,omitempty"`
	/*Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.*/
	Thumb string `json:"thumb,omitempty"`
}

func (args *SendAudioArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendAudioArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("audio")
	_, _ = io.Copy(fw, strings.NewReader(args.Audio))
	if args.Duration != 0 {
		fw, _ = wr.CreateFormField("duration")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Duration)))
	}
	if args.Thumb != "" {
		fw, _ = wr.CreateFormField("thumb")
		_, _ = io.Copy(fw, strings.NewReader(args.Thumb))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
	if args.Performer != "" {
		fw, _ = wr.CreateFormField("performer")
		_, _ = io.Copy(fw, strings.NewReader(args.Performer))
	}
	if args.Title != "" {
		fw, _ = wr.CreateFormField("title")
		_, _ = io.Copy(fw, strings.NewReader(args.Title))
	}
}

type SendDocumentArgs struct {
	DefaultSendMethodsArguments
	Document string `json:"document"`
	/*Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.*/
	Thumb           string          `json:"thumb,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	/*Disables automatic server-side content type detection for files uploaded using multipart/form-data*/
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

func (args *SendDocumentArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendDocumentArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("document")
	_, _ = io.Copy(fw, strings.NewReader(args.Document))
	if args.Thumb != "" {
		fw, _ = wr.CreateFormField("thumb")
		_, _ = io.Copy(fw, strings.NewReader(args.Thumb))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
	fw, _ = wr.CreateFormField("disable_content_type_detection")
	_, _ = io.Copy(fw, strings.NewReader(strconv.FormatBool(args.DisableContentTypeDetection)))
}

type SendVideoArgs struct {
	DefaultSendMethodsArguments
	Video string `json:"video"`
	/*Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.*/
	Thumb           string          `json:"thumb,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	/*Duration of sent video in seconds.*/
	Duration int `json:"duration,omitempty"`
	/*Pass True, if the uploaded video is suitable for streaming*/
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

func (args *SendVideoArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendVideoArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("video")
	_, _ = io.Copy(fw, strings.NewReader(args.Video))
	fw, _ = wr.CreateFormField("supports_streaming")
	_, _ = io.Copy(fw, strings.NewReader(strconv.FormatBool(args.SupportsStreaming)))
	if args.Duration != 0 {
		fw, _ = wr.CreateFormField("duration")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Duration)))
	}
	if args.Thumb != "" {
		fw, _ = wr.CreateFormField("thumb")
		_, _ = io.Copy(fw, strings.NewReader(args.Thumb))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
}

type SendAnimationArgs struct {
	DefaultSendMethodsArguments
	Animation string `json:"animation"`
	/*Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.*/
	Thumb           string          `json:"thumb,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Width           int             `json:"width,omitempty"`
	Height          int             `json:"height,omitempty"`
	/*Duration of sent video in seconds.*/
	Duration int `json:"duration,omitempty"`
}

func (args *SendAnimationArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendAnimationArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("animation")
	_, _ = io.Copy(fw, strings.NewReader(args.Animation))
	if args.Duration != 0 {
		fw, _ = wr.CreateFormField("duration")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Duration)))
	}
	if args.Thumb != "" {
		fw, _ = wr.CreateFormField("thumb")
		_, _ = io.Copy(fw, strings.NewReader(args.Thumb))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
	if args.Width != 0 {
		fw, _ = wr.CreateFormField("width")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Width)))
	}
	if args.Height != 0 {
		fw, _ = wr.CreateFormField("height")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Height)))
	}
}

type SendVoiceArgs struct {
	DefaultSendMethodsArguments
	Voice           string          `json:"voice"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	/*Duration of sent video in seconds.*/
	Duration int `json:"duration,omitempty"`
}

func (args *SendVoiceArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendVoiceArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("voice")
	_, _ = io.Copy(fw, strings.NewReader(args.Voice))
	if args.Duration != 0 {
		fw, _ = wr.CreateFormField("duration")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Duration)))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
}

type SendVideoNoteArgs struct {
	DefaultSendMethodsArguments
	VideoNote string `json:"video_note"`
	/*Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.*/
	Thumb           string          `json:"thumb,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Length          int             `json:"length,omitempty"`
	/*Duration of sent video in seconds.*/
	Duration int `json:"duration,omitempty"`
}

func (args *SendVideoNoteArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendVideoNoteArgs) ToMultiPart(wr *mp.Writer) {
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("video_note")
	_, _ = io.Copy(fw, strings.NewReader(args.VideoNote))
	if args.Duration != 0 {
		fw, _ = wr.CreateFormField("duration")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Duration)))
	}
	if args.Thumb != "" {
		fw, _ = wr.CreateFormField("thumb")
		_, _ = io.Copy(fw, strings.NewReader(args.Thumb))
	}
	if args.Caption != "" {
		fw, _ = wr.CreateFormField("caption")
		_, _ = io.Copy(fw, strings.NewReader(args.Caption))
	}
	if args.ParseMode != "" {
		fw, _ = wr.CreateFormField("parse_mode")
		_, _ = io.Copy(fw, strings.NewReader(args.ParseMode))
	}
	if args.CaptionEntities != nil {
		fw, _ = wr.CreateFormField("caption_entities")
		bt, _ := json.Marshal(args.CaptionEntities)
		_, _ = io.Copy(fw, bytes.NewReader(bt))
	}
	if args.Length != 0 {
		fw, _ = wr.CreateFormField("lengt")
		_, _ = io.Copy(fw, strings.NewReader(strconv.Itoa(args.Length)))
	}
}

type SendMediaGroupArgs struct {
	DefaultSendMethodsArguments
	Media []InputMedia `json:"media"`
}

func (args *SendMediaGroupArgs) ToJson() []byte {
	args.ReplyMarkup = nil
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendMediaGroupArgs) ToMultiPart(wr *mp.Writer) {
	args.ReplyMarkup = nil
	args.toMultiPart(wr)
	fw, _ := wr.CreateFormField("media")
	bt, _ := json.Marshal(args.Media)
	_, _ = io.Copy(fw, bytes.NewReader(bt))
}

type SendLocationArgs struct {
	DefaultSendMethodsArguments
	/*Latitude of the location*/
	Latitude float32 `json:"latitude"`
	/*Longitude of the location*/
	Longitude float32 `json:"longitude"`
	/*The radius of uncertainty for the location, measured in meters; 0-1500*/
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	/*Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.*/
	LivePeriod int `json:"live_period,omitempty"`
	/*	For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.*/
	Heading int `json:"heading,omitempty"`
	/*For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.*/
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

func (args *SendLocationArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendLocationArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type EditMessageLiveLocationArgs struct {
	/*Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)*/
	ChatId json.RawMessage `json:"chat_id,omitempty"`
	/*Required if inline_message_id is not specified. Identifier of the message to edit*/
	MessageId int `json:"message_id,omitempty"`
	/*Required if chat_id and message_id are not specified. Identifier of the inline message*/
	InlineMessageId string `json:"inline_message_id,omitempty"`
	/*Latitude of the location*/
	Latitude float32 `json:"latitude"`
	/*Longitude of the location*/
	Longitude float32 `json:"longitude"`
	/*The radius of uncertainty for the location, measured in meters; 0-1500*/
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	/*	For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.*/
	Heading int `json:"heading,omitempty"`
	/*For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.*/
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	/*Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.*/
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (args *EditMessageLiveLocationArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *EditMessageLiveLocationArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type StopMessageLiveLocationArgs struct {
	/*Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)*/
	ChatId json.RawMessage `json:"chat_id,omitempty"`
	/*Required if inline_message_id is not specified. Identifier of the message to edit*/
	MessageId int `json:"message_id,omitempty"`
	/*Required if chat_id and message_id are not specified. Identifier of the inline message*/
	InlineMessageId string `json:"inline_message_id,omitempty"`
	/*Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.*/
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (args *StopMessageLiveLocationArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *StopMessageLiveLocationArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SendVenueArgs struct {
	DefaultSendMethodsArguments
	/*Latitude of the location*/
	Latitude float32 `json:"latitude"`
	/*Longitude of the location*/
	Longitude float32 `json:"longitude"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	/*Foursquare identifier of the venue*/
	FoursquareId string `json:"foursquare_id,omitempty"`
	/*Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)*/
	FoursquareType  string `json:"foursquare_type,omitempty"`
	GooglePlaceId   string `json:"google_place_id,omitempty"`
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

func (args *SendVenueArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendVenueArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SendContactArgs struct {
	DefaultSendMethodsArguments
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

func (args *SendContactArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendContactArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SendPollArgs struct {
	DefaultSendMethodsArguments
	/*Poll question, 1-300 characters*/
	Question string `json:"question"`
	/*A JSON-serialized list of answer options, 2-10 strings 1-100 characters each*/
	Options []string `json:"options"`
	/*Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.*/
	IsClosed bool `json:"is_closed,omitempty"`
	/*True, if the poll needs to be anonymous, defaults to True*/
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	/*Poll type, “quiz” or “regular”, defaults to “regular”*/
	Type string `json:"type,omitempty"`
	/*	True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False*/
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
	/*0-based identifier of the correct answer option, required for polls in quiz mode*/
	CorrectOptionId int `json:"correct_option_id"`
	/*Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing*/
	Explanation string `json:"explanation,omitempty"`
	/*Mode for parsing entities in the explanation. See formatting options for more details.*/
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`
	/*	A JSON-serialized list of special entities that appear in the poll explanation, which can be specified instead of parse_mode*/
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	/*	Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.*/
	OpenPeriod int `json:"open_period,omitempty"`
	/*the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.*/
	CloseDate int `json:"close_date,omitempty"`
}

func (args *SendPollArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendPollArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SendDiceArgs struct {
	DefaultSendMethodsArguments
	/*Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults to “🎲”*/
	Emoji string `json:"emoji,omitempty"`
}

func (args *SendDiceArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendDiceArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SendChatActionArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	/*Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.*/
	Action string `json:"action"`
}

func (args *SendChatActionArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SendChatActionArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type GetUserProfilePhototsArgs struct {
	/*Unique identifier of the target user*/
	UserId int `json:"user_id"`
	/*Sequential number of the first photo to be returned. By default, all photos are returned.*/
	Offset int `json:"offset,omitempty"`
	/*Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.*/
	Limit int `json:"limit,omitempty"`
}

func (args *GetUserProfilePhototsArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *GetUserProfilePhototsArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type GetFileArgs struct {
	FileId string `json:"file_id"`
}

func (args *GetFileArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *GetFileArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type DefaultChatArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
}

func (args *DefaultChatArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *DefaultChatArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type BanChatMemberArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	UserId int             `json:"user_id"`
	/*Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.*/
	UntilDate int `json:"until_date,omitempty"`
	/*Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.*/
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

func (args *BanChatMemberArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *BanChatMemberArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type UnbanChatMemberArgsArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	UserId int             `json:"user_id"`
	/*Do nothing if the user is not banned*/
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

func (args *UnbanChatMemberArgsArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *UnbanChatMemberArgsArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type RestrictChatMemberArgs struct {
	ChatId     json.RawMessage `json:"chat_id"`
	UserId     int             `json:"user_id"`
	Permission ChatPermissions `json:"permissions"`
	UntilDate  int             `json:"until_date,omitempty"`
}

func (args *RestrictChatMemberArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *RestrictChatMemberArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type PromoteChatMemberArgs struct {
	ChatId              json.RawMessage `json:"chat_id"`
	UserId              int             `json:"user_id"`
	IsAnonymous         bool            `json:"is_anonymous"`
	CanManageChat       bool            `json:"can_manage_chat"`
	CanPostMessages     bool            `json:"can_post_messages"`
	CanEditMessages     bool            `json:"can_edit_messages"`
	CanDeleteMessages   bool            `json:"can_delete_messages"`
	CanManageVoiceChats bool            `json:"can_manage_voice_chats"`
	CanRestrictMembers  bool            `json:"can_restrict_members"`
	CanPromoteMembers   bool            `json:"can_promote_members"`
	CanChangeInfo       bool            `json:"can_change_info"`
	CanInviteUsers      bool            `json:"can_invite_users"`
	CanPinMessages      bool            `json:"can_pin_messages"`
}

func (args *PromoteChatMemberArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *PromoteChatMemberArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SetChatAdministratorCustomTitleArgs struct {
	ChatId      json.RawMessage `json:"chat_id"`
	UserId      int             `json:"user_id"`
	CustomTitle string          `json:"custom_title"`
}

func (args *SetChatAdministratorCustomTitleArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SetChatAdministratorCustomTitleArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type BanChatSenderChatArgs struct {
	ChatId       json.RawMessage `json:"chat_id"`
	SenderChatId int             `json:"sender_chat_id"`
}

func (args *BanChatSenderChatArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *BanChatSenderChatArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type UnbanChatSenderChatArgs struct {
	ChatId       json.RawMessage `json:"chat_id"`
	SenderChatId int             `json:"sender_chat_id"`
}

func (args *UnbanChatSenderChatArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *UnbanChatSenderChatArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SetChatPermissionsArgs struct {
	ChatId      json.RawMessage `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (args *SetChatPermissionsArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SetChatPermissionsArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type CreateChatInviteLinkArgs struct {
	ChatId             json.RawMessage `json:"chat_id"`
	Name               string          `json:"name,omitempty"`
	ExpireDate         int             `json:"expire_date,omitempty"`
	MemberLimit        int             `json:"member_limit,omitempty"`
	CreatesjoinRequest bool            `json:"creates_join_request,omitempty"`
}

func (args *CreateChatInviteLinkArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *CreateChatInviteLinkArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type EditChatInviteLinkArgs struct {
	ChatId             json.RawMessage `json:"chat_id"`
	InviteLink         string          `json:"invite_link"`
	Name               string          `json:"name,omitempty"`
	ExpireDate         int             `json:"expire_date,omitempty"`
	MemberLimit        int             `json:"member_limit,omitempty"`
	CreatesjoinRequest bool            `json:"creates_join_request,omitempty"`
}

func (args *EditChatInviteLinkArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *EditChatInviteLinkArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type RevokeChatInviteLinkArgs struct {
	ChatId     json.RawMessage `json:"chat_id"`
	InviteLink string          `json:"invite_link"`
}

func (args *RevokeChatInviteLinkArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *RevokeChatInviteLinkArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type ApproveChatJoinRequestArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	UserId int             `json:"user_id"`
}

func (args *ApproveChatJoinRequestArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *ApproveChatJoinRequestArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type DeclineChatJoinRequestArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	UserId int             `json:"user_id"`
}

func (args *DeclineChatJoinRequestArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *DeclineChatJoinRequestArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SetChatPhotoArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	Photo  string          `json:"photo"`
}

func (args *SetChatPhotoArgs) ToJson() []byte {
	//The arguments of this method are never passed as json.
	return nil
}

func (args *SetChatPhotoArgs) ToMultiPart(wr *mp.Writer) {
	fw, _ := wr.CreateFormField("chat_id")
	_, _ = io.Copy(fw, bytes.NewReader(args.ChatId))
	fw, _ = wr.CreateFormField("photo")
	_, _ = io.Copy(fw, strings.NewReader(args.Photo))
}

type SetChatTitleArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	Title  string          `json:"title"`
}

func (args *SetChatTitleArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SetChatTitleArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SetChatDescriptionArgs struct {
	ChatId      json.RawMessage `json:"chat_id"`
	Description string          `json:"description"`
}

func (args *SetChatDescriptionArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SetChatDescriptionArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type PinChatMessageArgs struct {
	ChatId              json.RawMessage `json:"chat_id"`
	MessageId           int             `json:"message_id"`
	DisableNotification bool            `json:"disable_notification"`
}

func (args *PinChatMessageArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *PinChatMessageArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type UnpinChatMessageArgs struct {
	ChatId    json.RawMessage `json:"chat_id"`
	MessageId int             `json:"message_id"`
}

func (args *UnpinChatMessageArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *UnpinChatMessageArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type GetChatMemberArgs struct {
	ChatId json.RawMessage `json:"chat_id"`
	UserId int             `json:"user_id"`
}

func (args *GetChatMemberArgs) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *GetChatMemberArgs) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}

type SetChatStcikerSet struct {
	ChatId         json.RawMessage `json:"chat_id"`
	StickerSetName string          `json:"sticker_set_name"`
}

func (args *SetChatStcikerSet) ToJson() []byte {
	bt, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return bt
}

func (args *SetChatStcikerSet) ToMultiPart(wr *mp.Writer) {
	//This method arguments are never passed as multipart
}
