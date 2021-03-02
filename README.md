# Go Scrum Poker Bot

[![build](https://github.com/isayakhov/duty-schedule-bot/workflows/Linters%20And%20Tests/badge.svg)](https://github.com/isayakhov/duty-schedule-bot/blob/main/.github/workflows/lint-and-tests.yml)
[![codecov](https://codecov.io/gh/isayakhov/duty-schedule-bot/branch/main/graph/badge.svg?token=7DVLEWCKQR)](https://codecov.io/gh/isayakhov/duty-schedule-bot)
[![code_quality](https://www.code-inspector.com/project/18495/status/svg)](https://frontend.code-inspector.com/public/project/18495/duty-schedule-bot/dashboard)
[![code_score](https://www.code-inspector.com/project/18495/score/svg)](https://frontend.code-inspector.com/public/project/18495/duty-schedule-bot/dashboard)

## What is this?

This bot can help you to play **Scrum Poker** with your team. You need to create new application in your apps in Slack and add it rights.

## Why?

Because it's fun and better to have your own Scrum Poker instrument instead of using free or payable soft.

## How to run it?

1. First you need to create your own bot in Slack and add it to your group
2. Next you need to pass to environment variables your bot's token into `SLACK_TOKEN` variable
   - You also need to configure a public domain with SSL
   - You also need to configure slash commands on your Slack Applications page
3. Build and run `main.go`
4. That's all :)

## Environment variables

### Common

|Name     | Required | Default | Description|
|:--------|:-------- |:------- |:-----------|
| REDIS_HOST                    |   | redis    | Redis host address                |
| REDIS_PORT                    |   | 6379     | Redis port number                 |
| REDIS_DB                      |   | 0        | Redis database number             |
| SLACK_WEB_SERVER_ADDRESS      |   | :8000    | Web server address                |

### Slack

|Name     | Required | Default | Description|
|:--------|:-------- |:------- |:-----------|
| SLACK_TOKEN                   |   | ""       | Slack bot token        |
| SLACK_WEB_SERVER_ADDRESS      |   | :8000    | Web server address     |