# Go Scrum Poker Bot

[![build](https://github.com/isayakhov/go-scrum-poker-bot/workflows/Lint,%20Tests,%20Build/badge.svg)](https://github.com/isayakhov/go-scrum-poker-bot/blob/master/.github/workflows/lint-tests-build.yml)
[![codecov](https://codecov.io/gh/isayakhov/go-scrum-poker-bot/branch/master/graph/badge.svg?token=GGJ61UJ6T3)](https://codecov.io/gh/isayakhov/go-scrum-poker-bot)
[![code_quality](https://www.code-inspector.com/project/19706/status/svg#1)](https://frontend.code-inspector.com/project/19706/dashboard)
[![code_score](https://www.code-inspector.com/project/19706/score/svg#1)](https://frontend.code-inspector.com/project/19706/dashboard)

## What is this?

This bot can help you to play **Scrum Poker** with your team with your own type of ranks (by default it's basic Fibonacci numbers from 0 to 100).

## How it looks like?

<details>
  <summary>See animation</summary>

  ![Demonstration](demo.gif)
</details>

## How to run it?

1. First you need to create your own bot in Slack and add it to your group
2. Next you need to pass to environment variables your bot's token into `SLACK_TOKEN` variable
   - You also need to configure a public domain with SSL (for development and testing purposes I recommend to use **ngrok**)
   - You also need to configure slash commands on your Slack Applications page (see below)
3. Build and run `go build .` or just type `docker-compose up`
4. Application will open `:8000` port and you can use it:
   - Healthcheck path for your metrics: **GET:** `/healthcheck`
   - Init poker session through Slack slash command: **POST**: `/play-poker`
   - Interact with UI through Slack: **POST**: `/interactivity`
4. That's all :)

## How to set up it?

1. You need to create new app into your Slack workplace
2. Give this app these rights: `chat:write`, `commands` and reinstall it to your workplace
3. Copy your token and paste it as env variable
4. Create new slash command e.g. `/poker`. Request URL: `https://your-domain.com/play-poker`
5. Go to **Interactivity & Shortcuts** and fill request URL: `https://your-domain.com/interactivity`
6. Go to your channel and try: `/poker LINK_TO_YOUR_TASK e.g. https://jira.example.com/ABC-124`

## Environment variables

### App

|Name     | Required | Default | Description|
|:--------|:-------- |:------- |:-----------|
| WEB_SERVER_ADDRESS   |   | :8000                          | Web server address     |
| POKER_RANKS          |   | ?,0,0.5,1,2,3,5,8,13,20,40,100 | Poker cards ranks      |

### Redis

|Name     | Required | Default | Description|
|:--------|:-------- |:------- |:-----------|
| REDIS_HOST |   | redis | Redis host address     |
| REDIS_PORT |   | 6379  | Redis port number      |
| REDIS_DB   |   | 0     | Redis database number  |

### Slack

|Name     | Required | Default | Description|
|:--------|:-------- |:------- |:-----------|
| SLACK_TOKEN  |   | ""  | Slack bot token  |