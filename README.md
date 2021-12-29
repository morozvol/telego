# Telego

[![Go Reference](https://pkg.go.dev/badge/github.com/SakoDroid/telego.svg)](https://pkg.go.dev/github.com/SakoDroid/telego)
[![telegram bot api](https://img.shields.io/badge/telegram-telegram%20bot%20api-blue)](https://core.telegram.org/bots/api)
![Version](https://img.shields.io/badge/%20%20Version%20%20-1.2.3-success)
![Development status](https://img.shields.io/badge/%20%20Development%20%20-%20%20Active%20%20-blueviolet)

A Go library for creating telegram bots.

![telego logo inspired by Golang logo](https://github.com/SakoDroid/telego/blob/master/telego-logo.jpg?raw=true)

* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
    * [Quick start](#quick-start)
    * [Step by step](#step-by-step)
        * [Creating the bot](#creating-the-bot)
        * [Receiving updates](#receiving-updates)
            * [Handlers](#handlers)
            * [Special channels](#special-channels)
            * [Update receving priority](#update-receiving-priority)
        * [Methods](#methods)
            * [Text messages](#text-messages)
            * [Media messages](#media-messages)
            * [Media group messages](#media-group-messages)
            * [Polls](#polls)
            * [Files](#files)
* [License](#license)

---------------------------------

## Features
* Fast and reliable
* Highly customizable
* Full support for [telegram bot api](https://core.telegram.org/bots/api)
* Offers two different ways for managing the bot updates :
    1. [Handlers](#handlers) ( for text messages only ).
    2. [Special channels](#special-channels)
* Automatic poll management : You don't need to worry about poll updates. Telego takes care of that for you. Just create a poll, send it and sit back 
and monitor the poll update via a go channel.

---------------------------------

## Requirements
  * Go 1.17 or higher.
  * Small and basic knowledge about telegram bots.
---------------------------------

## Installation
 Install the package into your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go command](https://golang.org/cmd/go/ "go command") from terminal :
 ```
 $ go get -u github.com/SakoDroid/telego
 ```
 Git needs to be installed on your computer.

 --------------------------------

 ## Usage

 ### Quick start

 The following code creates a bot and starts receving updates. If the update is a text message that contains "hi" the bot will respond "hi to you too!".

 ```
 import (
    "fmt"
    
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
 )

 func main(){
    up := cfg.DefaultUpdateConfigs()
    
    cf := cfg.BotConfigs{BotAPI: cfg.DefaultBotAPI, APIKey: "your api key", UpdateConfigs: up, Webhook: false, LogFileAddress: cfg.DefaultLogFile}

    bot, err := bt.NewBot(&cf)

    if err == nil{

        err == bot.Run()

        if err == nil{
            go start(bot)
        }
    }
 }

 func start(bot *bt.Bot){

     //The general update channel.
     updateChannel := bot.GetUpdateChannel()

    //Adding a handler. Everytime the bot receives message "hi" in a private chat, it will respond "hi to you too".
    bot.AddHandler("hi",func(u *objs.Update) {
		_,err := bot.SendMessage(u.Message.Chat.Id,"hi to you too","",u.Message.MessageId,false)
		if err != nil{
			fmt.Println(err)
		}
	},"private")

    //Monitores any other update. (Updates that don't contain text message "hi" in a private chat)
     for {
         update := <- updateChannel

        //Some processing on the update
     }
 }
 ```
 ## Step by step

### **Creating the bot**
 First you need to import required libraries :

 ```
 import (
    bt "github.com/SakoDroid/telego"
    cfg "github.com/SakoDroid/telego/configs"
    objs "github.com/SakoDroid/telego/objects"
 )
 ```

 Then you need to create bot configs. **BotConfigs** struct is located in configs package and contains these fields :

 ```
 /*This is the bot api server. If you dont have a local bot api server, use "configs.DefaultBotAPI" for this field.*/

 BotAPI string

 /*The API key for your bot. You can get the api key (token) from botfather*/

 APIKey string

 /*The settings related to getting updates from the api server. This field shoud only be populated when Webhook field is false, otherwise it is ignored.*/

 UpdateConfigs *UpdateConfigs

 /*This field idicates if webhook should be used for receiving updates or not.
 Recommend : false*/

 Webhook bool

 /*All the logs related to bot will be written in this file. You can use configs.DefaultLogFile for default value*/

 LogFileAddress string
```

 * **Note** : telego library currently does not support webhooks so Webhook field should always be *false*.

To create bot configs you need an UpdateConfigs to populate related field in BotConfigs. **UpdateConfigs** struct contains following fields :

```
/*Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.*/

 Limit int

 /*Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.*/

 Timeout int

 /*List of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.
 Please note that this parameter doesnt affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.*/

 AllowedUpdates []string

 /*This field indicates the frequency to call getUpdates method. Default is one second*/

 UpdateFrequency time.Duration
 ```
 You can use **`configs.DefaultUpdateConfigs()`** to create default update configs. Otherwise you can create your own custom update configs.

 After you have created BotConfigs you can create the bot by passing the `BotConfigs` struct you've created to **NewBot** method located in **telego** package. After bot is created call **Run()** method and your bot will start working and will receive updates from the api server: 
 ```
 import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
 )

 func main(){
    up := cfg.DefaultUpdateConfigs()
    
    cf := cfg.BotConfigs{BotAPI: cfg.DefaultBotAPI, APIKey: "your api key", UpdateConfigs: up, Webhook: false, LogFileAddress: cfg.DefaultLogFile}

    bot, err := bt.NewBot(&cf)
    if err == nil{
        err == bot.Run()
        if err == nil{
            //Do anything you want with the bot.
        }
    }
 }
```

Now that the bot is running it will receive updates from api server and passes them into UpdateeChannel. So you can use this channel to know if an update is received from api server. You can get the channel via **GetUpdateChannel()** method of the bot :

 ```
 import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
 )

 func main(){
    up := cfg.DefaultUpdateConfigs()
    
    cf := cfg.BotConfigs{BotAPI: cfg.DefaultBotAPI, APIKey: "your api key", UpdateConfigs: up, Webhook: false, LogFileAddress: cfg.DefaultLogFile}

    bot, err := bt.NewBot(&cf)
    if err == nil{
        err == bot.Run()
        if err == nil{
            go start(bot)
        }
    }
 }

 func start(bot *bt.Bot){
     updateChannel := bot.GetUpdateChannel()
     for {
         update := <- updateChannel
         //Do your own processing.
     }
 }
```

### Receiving updates
#### **Handlers**

You can use handlers for routing text messages. You specify a function and everytime a text message is recevied which the handlers regex matches with the text the specified function will be called. Function format should be like this `exampleFunction(*objs.Update)`. To add a handler you sipmly call `AddHandler(pattern string, handler func(*objs.Update), chatTypes ...string)`. Arguments :
1. "Pattern" is the regex pattern which will be matched against the received text message.
2. "chatType" : is the a string array containing chat types which the handler will act on. It can be "private","group","supergroup","channel" and "all".
3. "handler" : is the function that will be called.

Handlers are super easy to use and you can see an example in [Quick start](#quick-start) section.

#### **Special channels**

In telego you can register special channels. Special channels are channels for a specified update. Meaning this channels will be uptaded when the the specified update is received from api server, giving the developers a lot more felxibility. To use special channels you need to call `RegisterChannel(chatId string, mediaType string)` method of the **advanced bot** (so for using this method, first you should call `AdvancedMode()` method of the bot). This method is fully documented in the source code but we will describe it here too. This method takes two arguments : 
1. chatId : This is a string representing a certain chat which this channel will be dedicated to. This argument can be chat identificator of a chat or username of a channel or supergroup.
2. mediaType : This argument specifies an update type which the channel will be dedicated to. For example if you pass "message", the returned channel will only be updated when an update containing message field [for a specified chat] is received.

**Note :** Both arguments can be used together to create channels that will be updated only when a certain field (mediaType) is present in the received update for a specified chat (chatId).

Examples :

This method can be used in four ways :
1. RegisterChannel("123456","message") : The returned channel will be updated when a message (text,photo,video ...) is received from a chat with "123456" as it's chat id.

2. RegiterChannel("","message") : The returned channel will be updated everytime a message is received from any chat.

3. RegisterChannel("123456","") : The returned channel will be updated everytime an update of anykind is received for the specified chat.

4. RegisterChannel("","") : The returned is the global update channel which will be updated everytime an update is received. You can get this channel by calling `getUpdateChannel()` method too.

**Note :** When a channel is registered it is not editable. Meaning that calling the `RegisterChannel` method with the same arugments won't create a new channel and the previously created channel will be returned.

Once a channel is created it cannot be edited, But it can be deleted. To delete a channel (unregister it) call `UnRegisterChannel(chatId string,mediaType string)` method of the **AdvancedBot**. **If** a channel has been registered for the given arguments it will be cleared.

#### **Update receving priority :**

Since different types of channels and handlers may get involved it's important to know the priority of them. Meaning when an update is received which methods have higher priority to be executed and in case of channels which channels will be first considered to have the update passed into them. Basically this is how handlers and channels are prioritized :

1. Hanlders
2. Chat channels :
    1. Update types
    2. General
2. Global channels :
    1. Updates types
    2. General channel

When an update is received, first it is compared against all the handlers. If a handler's regex matching is successfull the handler will be executed. If not handler is successfull then channels are checked. 

After none of the handlers are executed then the update is checked to see if it has chat information and if it does, channels registered for that chat are checked. If a channel is registered for the field that the update contains it will be passed into the channel. If no channel is registered for the field then it will be passed into the general channel for the chat.( For example lets assume you haved called `RegisterChannel("123456","message")` method, in this case if an update for a chat that it's chat id is "123456" is received that contains `message` field, it will be passed into this channel. ) If this step fails (does not have chat information or no channel is registered for the chat) then the *update type channels* are checked and if the update contains a field that does have a channel registered for it the related field will be passed into the channel.(For example if the update contains message field and you have called `RegisterChannel("","message")` method, the update will be passed into the channel). If this step fails too then the update will be passed into general update channel. 

To summarize :

```
Update is received -> Handlers
                          |
                          |
If no hanlder is executed |
                          |
                          |                                                / Specified update type channel
                     Chat channels (if update is relevant to a chat) ----- 
                             |                                             \ General chat channel
                             |
if chat channel check fails  |
                             |
                             |----------> General update type channels
                                                   |
                                                   |
                              if this check fails  |
                                                   |
                                                   |----------> General update channel
                              
```

**Note :** 

Handlers and special channels cnan be used together. For example the below code add a hander for text message "hi". Everytime the bot receives "hi" in a private chat it responds "hi to you too, send a location". Then it rgisters a channel for receiving messages in that chat and waits for the user to send a message. After message is received it sends the exact same location the user has sent back to the user : 

```
import (
    "fmt"
    
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
 )

 func main(){
    up := cfg.DefaultUpdateConfigs()
    
    cf := cfg.BotConfigs{BotAPI: cfg.DefaultBotAPI, APIKey: "your api key", UpdateConfigs: up, Webhook: false, LogFileAddress: cfg.DefaultLogFile}

    bot, err := bt.NewBot(&cf)

    if err == nil{

        err == bot.Run()

        if err == nil{
            go start(bot)
        }
    }
 }

 func start(bot *bt.Bot){

     //The general update channel.
     updateChannel := bot.GetUpdateChannel()

    //Adding a handler. Everytime the bot receives message "hi" in a private chat, it will respond "hi to you too".

    bot.AddHandler("hi",func(u *objs.Update) {

        //Register channel for receiving messages from this chat.
		cc, _ := bot.AdvancedMode().RegisterChannel(strconv.Itoa(u.Message.Chat.Id), "message")

        //Sends back a message
		_, err := bot.SendMessage(u.Message.Chat.Id, "hi to you too, send me a location", "", u.Message.MessageId, false)
		if err != nil {
			fmt.Println(err)
		}

        //Waits for an update from this chat
		up := <-*cc

        //Sends back the received location
		_, err = bot.SendLocation(up.Message.Chat.Id, false, up.Message.Location.Latitude, up.Message.Location.Longitude, up.Message.Location.HorizontalAccuracy, up.Message.MessageId)


		if err != nil {
			fmt.Println(err)
		}
	},"private")

    //Monitores any other update. (Updates that don't contain text message "hi" in a private chat)
     for {
         update := <- updateChannel

        //Some processing on the update
     }
 }
```

### **Methods**

 To send back text or media (such as photo, video, gif, ...) you can use Send methods. There are several send methods such as **SendMessage** and **SendPhoto**. There is two ways to send back data to the client. First way is using unique chat ids (which are integers that are unique for each chat) to send data to private chats, groups and supergroups. Second way is using chat username which can be used to send back data to supergroups (with username) and channels. Methods that use username as chat identificator end with `UN`.
 
 We will cover some of the methods below. All these methods are fully documented in the source code and will be described here briefly. In all methods you can ignore `number` arguments (int or float) by passing 0 and ignore `string` arguments by passing empty string ("").
  * **Note** : All bot methods are simplified to avoid unnecessary arguments. To access more options for each method you can call `AdvancedMode()` method of the bot that will return an advanced version of bot which will give you full access.

 #### **Text messages**

 To send back text you can use **SendMessage** (chat id) or **SendMessageUN** (username). 

 #### **Media messages**

 To send media types such as photo,video,gif,audio,voice,videonote,mpeg4 gif,sticker and document you can use their specified method. In general there are three ways to send media :
 
 1. **By file id** : File id is a unique id for a file that already exists in telegram servers. [Telegram bot api documentation](https://core.telegram.org/bots/api) recommends using file id.
 2. **By URL** : You can pass an HTTP url to send. The file will be downloaded in telegram servers and then it will be sent to the specified chat.
 3. **By file** : You can send a file on your computer. The file will be uploaded to telegram servers and then it will be sent to the specified chat.

 Calling each media sending related method returnes a MediaSender. MediaSender has all methods that are needed to send a media. For example lets send photo in our computer :

 ```
 photoFile,err := os.Open("photo.jpg")

 if err == nil{

    ms := bot.SendPhoto(chatId, messageId, "custom caption", "")

    _,err = ms.SendByFile(photoFile,false)

    if err != nil{
        fmt.Println(err)
    }

 }
 ```
 
 #### **Media group messages**

 To send a group of medias (aka albums) first you need to create a *`MediaGroup`* by calling `CreateAlbum(replyto int)` method of the bot. MediaGroup has several methods for adding photo,video,audio and other media types to the album. Keep in mind that according to [Telegram bot api documentation about media groups](https://core.telegram.org/bots/api#sendmediagroup), documents and audio files can be only grouped in an album with messages of the same type. Also the media group must include 2-10 items. The code below shows how to create a media group, add some photo to it and send it :

 ```
 mg := bot.CreateAlbum(messageId)

//Add a file on the computer.
fl,_ := os.Open("file.jpg")
 pa1,_ := mg.AddPhoto("", "", nil)
 err := pa1.AddByFile(fl)
 if err != nil{
     fmt.Println(err)
 }

//Add a photo by file id or url.
 pa2,_ ;= mg.AddPhoto("","",nil)
 err = pa2.AddByFileIdOrURL("fileId or HTTP url")
 if err != nil{
     fmt.Println(err)
 }

//Send the media group
_, err = mg.Send(chatId, false)
if err != nil {
    fmt.Println(err)
}
```

#### **Polls**

telego library offers automatic poll management. When you create a poll and send the poll bot will receive updates about the poll. Whene you create a poll by **`CreatePoll`** method, it will return a Poll which has methods for managing the poll. You should keep the returned pointer (to Poll) somewhere because everytime an update about a poll is received the bot will process the update and update the related poll and notifies user through a [bool]channel (which you can get by calling `GetUpdateChannel` method of the poll). 

* **Note** : If an update is received that contains update about a poll and the poll is not registered with the Polls map, the given update is passed into *UpdateChannel* of the bot. Otherwise as described above, the related poll will be updated.

Let's see an example :

```

//A custom function that creates and sends a poll and listens to its updates.
func pollTest(chatId int) {

    //Creates the poll
	poll, _ := bot.CreatePoll(chatId, "How are you?", "regular")

    //Adds some options
	poll.AddOption("good")
	poll.AddOption("not bad")
	poll.AddOption("alright")
	poll.AddOption("bad")

    //Adds an explanation for the poll.
	poll.SetExplanation("This is just a test for telego framework", "", nil)

    //Sends the poll
	err := poll.Send(false, 0)

	if err != nil {
		fmt.Println(err)
	} else {

        //Starts waiting for updates and when poll is updated, the updated result of the bot is printed.
		ch := poll.GetUpdateChannel()
		for {
			<-*ch
			fmt.Println("poll updated.")
			for _, val := range poll.GetResult() {
				fmt.Println(val.Text, ":", val.VoterCount)
			}
		}
	}
}
```

#### **Files**

You can get informations of a file that is stored in telegram servers and download it into your computer by calling **`GetFile`** method. If you want to download the file, pass true for *download* argument of the method. The below example downloads a received sticker from the user and saves it into the given file (read full documentation of the method for more information) :

```
//Receives upadate
update := <- updateChannel

//Get sticker file id
fi := update.Message.Sticker.FileId

//Open a file in the computer.
fl, _ := os.OpenFile("sticker.webp", os.O_CREATE|os.O_WRONLY, 0666)

//Gets the file info and downloads it.
_, err := bot.GetFile(fi, true, fl)
if err != nil {
    fmt.Println(err)
}
fl.Close()

```
---------------------------

## License

telego is licensed under [MIT lisence](https://en.wikipedia.org/wiki/MIT_License). Which means it can be used for commerical and private apps and can be modified.

---------------------------

![telego logo inspired by Golang logo](https://github.com/SakoDroid/telego/blob/master/telego-logo.jpg?raw=true)
