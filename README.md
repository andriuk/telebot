# Telegram Bot for Writing Messages to CSV on local host

This simple Go-based Telegram bot allows you to capture incoming messages and log them into a CSV file. 
Useful for logging or archiving conversations and messages received on Telegram.

Example: 
Used for write data from Apple Shortcuts 

## Installation Guide

### Prerequisites

- Go (Golang) installed on your system
- Git for cloning the repository
- Telegram account for creating a bot and obtaining API token

### Installation Steps

1. **Clone this repository** by running the following command in your terminal:

`$ git clone git@github.com:andriuk/kbot.git`

2. **Create your own Telegram bot:**

- Go to [BotFather](https://t.me/BotFather) on Telegram.
- Use the `/newbot` command to create a new bot and follow the instructions to get the API token.

3. **Store your Telegram bot API token in an environment variable:**

- Use the command below to set the API token as an environment variable. Replace `<YOUR_API_TOKEN>` with the obtained token.


### Paste your API token
For more secure use command below:

   ` $ read -s TELE_TOKEN`

Paste your API token 

   `$ export TELE_TOKEN=<YOUR_API_TOKEN>`

4. **Run the bot** using the following command:

`$ ./kbot start`

5. **Interact with your Telegram bot** by sending messages to log them in the CSV file.


#### 6. Well done 

---

If you have any questions or need assistance, feel free to create an issue in this repository. 
We're here to help!