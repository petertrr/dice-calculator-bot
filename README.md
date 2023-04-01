# dice-calculator-bot
A Discord bot for TTRPGs that provides a calculator-like interface to input and evaluate dice notation (expressions like `2d6+d4+2`).

# Running with Docker
1. Create a Discord or Telegram bot token
2. Run Discord bot: `docker run [--entrypoint dice-calculator-bot-discord] -d ghcr.io/petertrr/dice-calculator-bot:latest -t <token>`
2. Run Telegram bot: `docker run --entrypoint dice-calculator-bot-telegram -d ghcr.io/petertrr/dice-calculator-bot:latest -t <token>`

# Building and running a binary
Run: `make build && ./bin/dice-calc-bot-<discord|telegram> -t <token>`