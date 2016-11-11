# Linebot Golang
	使用 golang 製作的 linebot

# 開發進度
	11/12: echo

# 環境設置
	在這個專案內，我們設定 CHANNEL_TOKEN, CHANNEL_SECRET 的方式是利用環境變數
	對應到的 code 會是:
	`os.Getenv("CHANNEL_TOKEN"), os.Getenv("CHANNEL_SECRET")`

# 測試方式
	利用 ngrok 架設一個暫時的 https server
	`$./ngrok http <your_port>`
	如果我們在設定 PORT 的時候，也將它設置為環境變數的話，就可以直接讀取

