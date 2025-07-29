# very simple echo bot

You must have a token from the bot.
And write the token to .env

## this bot works using method /getUpdates

https://core.telegram.org/bots/api#getupdates

```
https://api.telegram.org/bot{BOT_TOKEN}/getUpdates?offset={int}
```

## the first time it starts /deleteWebhook

for work method /getUpdates

https://core.telegram.org/bots/api#deletewebhook

```
https://api.telegram.org/bot{BOT_TOKEN}/deleteWebhook
```
