# telegoGPT
![](https://img.shields.io/badge/Code-Go-informational?style=flat&logo=go&color=61DAFB) ![](https://img.shields.io/badge/Code-JavaScript-informational?style=flat&logo=JavaScript&color=F7DF1E) ![](https://img.shields.io/badge/Code-HTML5-informational?style=flat&logo=HTML5&color=E34F26) ![](https://img.shields.io/badge/Code-MySQL-informational?style=flat&logo=mysql&color=336791)
 

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
