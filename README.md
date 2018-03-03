# go-twitter-bot
Many users follow you back if you follow them, and this is what this bot aims to exploit.

The twitter bot will search for users tweeting about configurable topics and give them a follow. Then after 6 hours (by default), it will unfollow them, and in most of the cases that user will forget or miss to unfollow you.

It is quite a toxic behavior, but it is an efficient way to get new followers if you don't mind that.

**For more details, please see https://thecuriousdev.org/golang-twitter-bot/**

**DISCLAIMER!** If you use the bot it's at your own risk. I am not responsible if Twitter would decide to take some action against you. With that in mind, the bots default values prevents spam and most likely you will be fine.

## Getting started
There are many ways to get started, you could clone the project (and install Golang on your computer if you don't have it already) and then simply compile the code by executing the command `go build` in the directory which will give you an executable file.

I have also uploaded a couple of different executables compiled for the most popular platforms which can be found here: https://snieking-owncloud.cloud.seedboxes.cc/index.php/s/JGeOtiFntelfi5I

### Configuring the bot
Modify/Create `config.json` file. An example file can be found in this repository. Fill in your details, and make sure that you pick up Twitter API keys from https://apps.twitter.com and enter those into the config file. DO NOT SHARE your API keys with anyone.

### Executing it
You start the bot by simply executing the binary. For example: `./go-twitter-bot.exe` or `./go-twitter-bot`

## Extra functionality
The twitter bot supports three different running modes and you can also configure some default values. This is done by providing some extra program arguments. Execute `./go-twitter-bot.exe -h` for details. 

`-clean` will clear all previous followed users.
`-unfollowAll` will unfollow all users that your account follows **(USE WITH CAUTION)**.

