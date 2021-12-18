package telebot

import (
	"os"

	objs "github.com/SakoDroid/telebot/objects"
)

/*An advanced type of bot which will give you alot more customization for the bot.
Methods which are uniquely for advanced bot start with 'A' .*/
type AdvancedBot struct {
	*Bot
}

/*Send a text message to a chat (not channel, use SendMessageToChannel method for sending messages to channles) and returns the sent message on success
If you want to ignore "parseMode" pass empty string. To ignore replyTo pass 0.*/
func (bot *AdvancedBot) ASendMessage(chatId int, text, parseMode string, replyTo int, silent bool, entites []objs.MessageEntity, disabelWebPagePreview, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendMessage(chatId, "", text, parseMode, entites, disabelWebPagePreview, silent, allowSendingWithoutReply, replyTo, replyMarkup)
}

/*Send a text message to a channel and returns the sent message on success
If you want to ignore "parseMode" pass empty string. To ignore replyTo pass 0.*/
func (bot *AdvancedBot) ASendMesssageToChannel(chatId, text, parseMode string, replyTo int, silent bool, entites []objs.MessageEntity, disabelWebPagePreview, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendMessage(0, chatId, text, parseMode, entites, disabelWebPagePreview, silent, allowSendingWithoutReply, replyTo, replyMarkup)
}

/*Returns a MessageCopier which has several methods for copying a message*/
func (bot *AdvancedBot) ACopyMessage(messageId int, disableNotif bool, replyTo int, caption, parseMode string, captionEntites []objs.MessageEntity, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MessageCopier {
	return &MessageCopier{bot: bot.Bot, messageId: messageId, disableNotif: disableNotif, caption: caption, parseMode: parseMode, captionEntities: captionEntites, allowSendingWihtouReply: allowSendingWithoutReply, replyTo: replyTo, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a photo. This method is only used for sending a photo to all types of chat except channels. To send a photo to a channel use "SendPhotoToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")*/
func (bot *AdvancedBot) ASendPhoto(chatId, replyTo int, caption, parseMode string, captionEntites []objs.MessageEntity, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: PHOTO, bot: bot.Bot, chatIdInt: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntites, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a photo. This method is only used for sending a photo to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")
*/
func (bot *AdvancedBot) ASendPhotoToChannel(chatId string, replyTo int, caption, parseMode string, captionEntites []objs.MessageEntity, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: PHOTO, bot: bot.Bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntites, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a video. This method is only used for sending a video to all types of chat except channels. To send a video to a channel use "SendVideoToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendVideo(chatId int, replyTo int, caption, parseMode string, captionEntites []objs.MessageEntity, duration int, supportsStreaming, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VIDEO, bot: bot.Bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntites, duration: duration, supportsStreaming: supportsStreaming, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a video. This method is only used for sending a video to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendVideoToChannel(chatId string, replyTo int, caption, parseMode string, captionEntites []objs.MessageEntity, duration int, supportsStreaming, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VIDEO, bot: bot.Bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntites, duration: duration, supportsStreaming: supportsStreaming, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a audio. This method is only used for sending a audio to all types of chat except channels. To send a audio to a channel use "SendAudioToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.*/
func (bot *AdvancedBot) ASendAudio(chatId, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, duration int, performer, title string, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: AUDIO, bot: bot.Bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntities, performer: performer, title: title, duration: duration, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a audio. This method is only used for sending a audio to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.*/
func (bot *AdvancedBot) ASendAudioToChannel(chatId string, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, duration int, performer, title string, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: AUDIO, bot: bot.Bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntities, performer: performer, title: title, duration: duration, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a document. This method is only used for sending a document to all types of chat except channels. To send a audio to a channel use "SendDocumentToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendDocument(chatId, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, disableContentTypeDetection, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: DOCUMENT, bot: bot.Bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntities, disableContentTypeDetection: disableContentTypeDetection, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a document. This method is only used for sending a document to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendDocumentToChannel(chatId string, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, disableContentTypeDetection, allowSendingWithoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: DOCUMENT, bot: bot.Bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode, captionEntities: captionEntities, disableContentTypeDetection: disableContentTypeDetection, allowSendingWihoutReply: allowSendingWithoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending an animation. This method is only used for sending an animation to all types of chat except channels. To send a audio to a channel use "SendAnimationToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendAnimation(chatId int, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, width, height, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: ANIMATION, chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, duration: duration, width: width, height: height, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending an animation. This method is only used for sending an animation to channels
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendAnimationToChannel(chatId string, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, width, height, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: ANIMATION, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, duration: duration, width: width, height: height, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a voice. This method is only used for sending a voice to all types of chat except channels. To send a voice to a channel use "SendVoiceToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendVoice(chatId int, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VOICE, chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, duration: duration, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a voice. This method is only used for sending a voice to channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *AdvancedBot) ASendVoiceToChannel(chatId string, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VOICE, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, duration: duration, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup}
}

/*Returns a MediaSender which has several methods for sending a video note. This method is only used for sending a video note to all types of chat except channels. To send a video note to a channel use "SendVideoNoteToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendVideoNote(chatId int, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, length, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VIDEONOTE, chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup, length: length, duration: duration}
}

/*Returns an MediaSender which has several methods for sending a video note. This method is only used for sending a video note to channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendVideoNoteToChannel(chatId string, replyTo int, caption, parseMode string, captionEntities []objs.MessageEntity, length, duration int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaSender {
	return &MediaSender{mediaType: VIDEONOTE, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot.Bot, caption: caption, parseMode: parseMode, captionEntities: captionEntities, allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup, length: length, duration: duration}
}

/*To ignore replyTo argument, pass 0.*/
func (bot *AdvancedBot) ACreateAlbum(replyTo int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *MediaGroup {
	return &MediaGroup{replyTo: replyTo, bot: bot.Bot, media: make([]objs.InputMedia, 0), files: make([]*os.File, 0), allowSendingWihoutReply: allowSendingWihtoutReply, replyMarkup: replyMarkup}
}

/*Sends a venue to all types of chat but channels. To send it to channels use "SendVenueToChannel" method.

---------------------------------

Official telegram doc :

Use this method to send information about a venue. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendVenue(chatId, replyTo int, latitude, longitude float32, title, address, foursquareId, foursquareType, googlePlaceId, googlePlaceType string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendVenue(
		chatId, "", latitude, longitude, title, address, foursquareId, foursquareType, googlePlaceId, googlePlaceType, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a venue to a channel.

---------------------------------

Official telegram doc :

Use this method to send information about a venue. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendVenueTOChannel(chatId string, replyTo int, latitude, longitude float32, title, address, foursquareId, foursquareType, googlePlaceId, googlePlaceType string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendVenue(
		0, chatId, latitude, longitude, title, address, foursquareId, foursquareType, googlePlaceId, googlePlaceType, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a contact to all types of chat but channels. To send it to channels use "SendContactToChannel" method.

---------------------------------

Official telegram doc :

Use this method to send phone contacts. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendContact(chatId, replyTo int, phoneNumber, firstName, lastName, vCard string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendContact(
		chatId, "", phoneNumber, firstName, lastName, vCard, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a contact to a channel.

---------------------------------

Official telegram doc :

Use this method to send phone contacts. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendContactToChannel(chatId string, replyTo int, phoneNumber, firstName, lastName, vCard string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendContact(
		0, chatId, phoneNumber, firstName, lastName, vCard, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a dice message to all types of chat but channels. To send it to channels use "SendDiceToChannel" method.

Available emojies : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.

---------------------------------

Official telegram doc :

Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned*/
func (bot *AdvancedBot) ASendDice(chatId, replyTo int, emoji string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendDice(
		chatId, "", emoji, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a dice message to a channel.

Available emojies : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.

---------------------------------

Official telegram doc :

Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned*/
func (bot *AdvancedBot) ASendDiceToChannel(chatId string, replyTo int, emoji string, silent bool, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendDice(
		0, chatId, emoji, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Creates a live location which has several methods for managing it.*/
func (bot *AdvancedBot) ACreateLiveLocation(latitude, longitude, accuracy float32, livePeriod, heading, proximtyAlertRadius, replyTo int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) *LiveLocation {
	return &LiveLocation{bot: bot.Bot, replyTo: replyTo, allowSendingWihoutReply: allowSendingWihtoutReply, latitude: latitude, longitude: longitude, livePeriod: livePeriod, horizontalAccuracy: accuracy, heading: heading, proximityAlertRadius: proximtyAlertRadius, replyMarkUp: replyMarkup}
}

/*Sends a location (not live) to all types of chats but channels. To send it to channel use "SendLocationToChannel" method.

You can not use this methods to send a live location. To send a live location use "ACreateLiveLocation" method.

---------------------------------

Official telegram doc :

Use this method to send point on the map. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendLocation(chatId int, silent bool, latitude, longitude, accuracy float32, replyTo int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendLocation(
		chatId, "", latitude, longitude, accuracy, 0, 0, 0, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Sends a location (not live) to a channel.

You can not use this methods to send a live location. To send a live location use "ACreateLiveLocation" method.

---------------------------------

Official telegram doc :

Use this method to send point on the map. On success, the sent Message is returned.*/
func (bot *AdvancedBot) ASendLocationToChannel(chatId string, silent bool, latitude, longitude, accuracy float32, replyTo int, allowSendingWihtoutReply bool, replyMarkup objs.ReplyMarkup) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendLocation(
		0, chatId, latitude, longitude, accuracy, 0, 0, 0, replyTo, silent, allowSendingWihtoutReply, replyMarkup,
	)
}

/*Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.

Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @Botfather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.*/
func (bot *Bot) AAnswerCallbackQuery(callbackQueryId, text string, showAlert bool, url string, cacheTime int) (*objs.LogicalResult, error) {
	return bot.apiInterface.AnswerCallbackQuery(callbackQueryId, text, url, showAlert, cacheTime)
}
