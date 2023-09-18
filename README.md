# telegoGPT
![](https://img.shields.io/badge/Code-Go-informational?style=flat&logo=go&color=61DAFB) ![](https://img.shields.io/badge/Code-JavaScript-informational?style=flat&logo=JavaScript&color=F7DF1E) ![](https://img.shields.io/badge/Code-HTML5-informational?style=flat&logo=HTML5&color=E34F26) ![](https://img.shields.io/badge/Code-MySQL-informational?style=flat&logo=mysql&color=336791)
 

telegoGPT adalah bot telegram dan integrasi openai dengan menggunakan bahasa golang.


![GitHub forks](https://img.shields.io/github/forks/1amkaizen/telegoGPT?style=social) ![GitHub Repo stars](https://img.shields.io/github/stars/1amkaizen/telegoGPT?style=social)

catatan:bot ini masih tahap awal dan sedang dalam pengembangan


# 🚀 cara menggunakan :

- Buat bot Telegram baru :  Buka Telegram dan cari BotFather kemudian buat bot baru dengan mengirimkan perintah `/newbot` . Ikuti instruksi untuk memberi nama dan nama pengguna pada bot Anda. Anda akan menerima token API untuk bot Anda yang akan Anda gunakan untuk mengautentikasi dengan API Telegram.
Gunakan kunci API Telegram untuk membuat instance bot baru.
```
bot, err := tgbotapi.NewBotAPI("TELEGRAMBOT_TOKEN")
```

- Buat akun di https://beta.openai.com/ dan dapatkan api nya,kemudian massukan ke sini:
```c := gogpt.NewClient("OPENAI_API")```

- Hosting ke railway dan konfigurasi variabel environment.
- Buat database di railway
- Konfigurasi variabel environment database telegoGPT
- Contoh konfigurasi variabel environment

```
 {
  "DATABASES": "root:5YUTSCHvx0yXwcsJUQFW@tcp(containers-us-west-139.railway.app:6522)/railway",
  "MYSQLDATABASE": "${{MySQL.MYSQLDATABASE}}",
  "MYSQLHOST": "${{MySQL.MYSQLHOST}}",
  "MYSQLPASSWORD": "${{MySQL.MYSQLPASSWORD}}",
  "MYSQLPORT": "${{MySQL.MYSQLPORT}}",
  "MYSQLUSER": "${{MySQL.MYSQLUSER}}",
  "MYSQL_URL": "${{MySQL.MYSQL_URL}}",
  "OPENAI_API": "sk-526u5wZVplXk518QhaJaT3BlbkJZTk4k1PtcaS5yBbVsnB",
  "TELEGRAM_BOT_TOKEN": "5898177748:AAE-4HBM9Ma9Wm7bICqhmMIQFJK0GnlVbf8"
}
```
**Catatan:** Yang harus di ganti hanya **DATABASES**,**TELEGRAM_BOT_TOKEN** & **OPENAI_API**.





## Demo

<a href="https://t.me/TelegoGPTbot"><img src="https://img.shields.io/badge/telegoGPT-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" />

## follow
<a href="https://twitter.com/1amkaizen"><img src="https://img.shields.io/badge/Twitter-2CA5E0?style=for-the-badge&logo=twitter&logoColor=white" />
<a href="https://replit.com/@1amkaizen?tab=status"><img src="https://img.shields.io/badge/Replit-ff5722?style=for-the-badge&logo=replit&logoColor=white" />
<a href="https://railway.app?referralCode=v-jhtw"><img src="https://img.shields.io/badge/Railway-ff5722?style=for-the-badge&logo=railway&logoColor=white" />
---
<p align="center"><img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" /></p>
<p align="center">Copyright &copy; 2023-present <a href="https://github.com/1amkaizen" target="_blank">1amkaizen</a>
<p align="center"><a href="https://github.com/1amkaizen/telegoGPT/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=for-the-badge&label=License&message=MIT&logoColor=d9e0ee&colorA=363a4f&colorB=b7bdf8"/></a></p>
