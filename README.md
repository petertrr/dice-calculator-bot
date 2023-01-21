# dice-calculator-bot
A Discord bot for TTRPGs that provides a calculator-like interface to input and evaluate dice notation (expressions like `2d6+d4+2`).

# Running with Docker
1. Create a Discord bot token
2. `docker run -d ghcr.io/petertrr/dice-calculator-bot:latest -t <token>`

# Building and running a binary
Run: `make build && ./bin/dice-calc-bot -t <token>`