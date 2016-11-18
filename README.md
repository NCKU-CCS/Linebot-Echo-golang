Linebot Golang Echo
===================
使用 golang 製作的 linebot

這份資料是我在系上分享的：[如何架設 LINE BOT](https://hackmd.io/s/SyJUciYWg)

Development Note
----------------
- 11/12: echo
<br>

Environment Settings
--------------------
在這個專案內，我們需要設定的環境變數：
- LineBot Setting:
	- CHANNEL_TOKEN
	- CHANNEL_SECRET

	對應到的 code 會是:
	`os.Getenv("CHANNEL_TOKEN"), os.Getenv("CHANNEL_SECRET")`


Build a Testing Environment
---------------------------
- using [ngrok](https://ngrok.com/download)
	- step1 `$./ngrok http <your_port>`
	- step2 `Set your local IP to Linebot whitelists`


