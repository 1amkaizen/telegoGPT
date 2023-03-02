# telegoGPT
[![forthebadge](http://forthebadge.com/images/badges/made-with-go.svg)](http://forthebadge.com)
[![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com)


![GitHub last commit](https://img.shields.io/github/last-commit/1amkaizen/telegoGPT) ![GitHub issues](https://img.shields.io/github/issues/1amkaizen/telegoGPT) ![GitHub pull requests](https://img.shields.io/github/issues-pr/1amkaizen/telegoGPT) ![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/1amkaizen/telegoGPT) ![GitHub contributors](https://img.shields.io/github/contributors/1amkaizen/telegoGPT)

telegoGPT adalah bot telegram dan integrasi openai dengan menggunakan bahasa golang.

telegoGPT is a telegram bot and openai integration using golang language.


![GitHub forks](https://img.shields.io/github/forks/1amkaizen/telegoGPT?style=social) ![GitHub Repo stars](https://img.shields.io/github/stars/1amkaizen/telegoGPT?style=social)

catatan:bot ini masih tahap awal dan sedang dalam pengembangan

note: this bot is still in its early stages and is under development

# 🚀 cara menggunakan :

- Buat bot Telegram baru :  Buka Telegram dan cari BotFather kemudian buat bot baru dengan mengirimkan perintah `/newbot` . Ikuti instruksi untuk memberi nama dan nama pengguna pada bot Anda. Anda akan menerima token API untuk bot Anda yang akan Anda gunakan untuk mengautentikasi dengan API Telegram.
Gunakan kunci API Telegram untuk membuat instance bot baru.
`bot, err := tgbotapi.NewBotAPI("TELEGRAMBOT_TOKEN")`.

- Buat akun di https://beta.openai.com/ dan dapatkan api nya,kemudian massukan ke sini:
`c := gogpt.NewClient("OPENAI_API")`

## Demo

<a href="https://t.me/TelegoGPTbot"><img src="https://img.shields.io/badge/telegoGPT-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" />

## follow
<a href="https://twitter.com/1amkaizen"><img src="https://img.shields.io/badge/Twitter-2CA5E0?style=for-the-badge&logo=twitter&logoColor=white" />
<a href="https://replit.com/@1amkaizen?tab=status"><img src="https://img.shields.io/badge/Replit-ff5722?style=for-the-badge&logo=replit&logoColor=white" />
<a href="https://railway.app?referralCode=v-jhtw"><img src="https://img.shields.io/badge/Railway-ff5722?style=for-the-badge&logo=railway&logoColor=white" />
