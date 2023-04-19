# Newsletter Bot

The Newsletter Bot is a Telegram bot that provides news recommendations and updates in different categories. It also provides the following features:

- Receiving daily mailings with the latest news updates.
- Automatic news recommendations based on your preferences.
- Receiving news updates by category.

## Available commands

The bot supports the following commands:

- `/start`: Start the bot and receive a welcome message.
- `/help`: Get help with using the bot.
- `/addrecomendations`: Add recommended news sources to the bot's database.
- `/addcategories`: Add categories to the bot's database.
- `/news`: Get the latest news updates from all categories.
- `/hotnewsus`: Get the latest hot news updates from the US.
- `/hotnewseu`: Get the latest hot news updates from the EU.
- `/sportnews`: Get the latest sports news updates.
- `/politicsnews`: Get the latest politics news updates.
- `/economicsnews`: Get the latest economics news updates.
- `/scienceandtechnews`: Get the latest science and technology news updates.
- `/healthnews`: Get the latest health news updates.
- `/entertainmentnews`: Get the latest entertainment news updates.
- `/travelnews`: Get the latest travel news updates.
- `/lifestylenews`: Get the latest lifestyle news updates.

## Installation and usage

To use the Newsletter Bot, you will need to:

1. Clone the repository and navigate to the root directory.
2. Run `go build` to build the binary.
3. Set the following environment variables:
    - `TELEGRAM_BOT_TOKEN`: Your Telegram bot token.
    - `NEWS_API_KEY`: Your NewsAPI.org API key.
    - `DATABASE_URL`: The URL for your PostgreSQL database.
4. Run the binary.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
