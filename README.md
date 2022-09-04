# Cards Bot

## Usage

This bot is operated using slash commands, type "/" in any channel the bot has access to, and you will be shown a list of all slash commands (for all bots), click on the cards bot logo to the left to see all commands for this bot.

## Commands

### 1. Heartbeat

Command name: `/heartbeat`

This command is used to check if the bot is working, and reading messages as it should be. If the bot is online and working, it should reply "I'm alive!"

### 2. Find Pokemon card by id

Command name: `/card-id` 

Command arguments: `id <id>`

Example:
Response to command `/card-id xy1-2`:


<img src="docs/pokemon/imgs/pkm-by-id.png" style="width:400px;"/>

### 3. Find Pokemon card/s by parameters

Command name: `/card-where`

The following parameters can be used:
- name
- hp <hp, to have a hp range, type [x TO y]>
- artist
- rarity
- 
Example:
Response to command `/card-where name: charizard hp: 330`:
  
<img src="docs/pokemon/imgs/pkm-by-params.png" style="width:400px;"/>
  
## Running this bot

To run this bot, run:

1. `docker build --tag docker-bot .` -> Builds binary
2. `docker run --env-file .env docker-bot` -> Executes binary

The env file, should contain the following:
```
TOKEN=<discord token>
```
