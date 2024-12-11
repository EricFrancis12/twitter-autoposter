# Twitter Auto-Poster

## Overview

This automation tool is designed to:

- Periodically fetch posts from defined sources.
- Publish new content as tweets via Twitter's API.
- Avoid duplicating posts by storing already published tweets in a database.

This program is useful for automating Twitter accounts to share updates, fetch RSS feeds, or repost curated content.

---

## Features

- **Configuration Management**: Easily define multiple Twitter accounts and their sources in a JSON configuration file.
- **Database Integration**: Tracks already published tweets to ensure no duplicates are posted.
- **Error Handling with Retry Logic**: Handles transient failures gracefully to ensure continuous operation.
- **Delay Mechanisms**: Avoids rate-limiting and bot detection by introducing sleep intervals between retries and operations.

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/EricFrancis12/twitter-autoposter.git
cd autoposter
```

### 2. Install Dependencies

Ensure you have Go installed, then run the following command to build the executable:

```bash
make build
```

### 3. Configuration

Run the following command to create a blank `config.json` file at the root of the project:

```bash
make create_config
```

Add your desired twitter account details and sources. Take a look at `config.example.json` for the correct format.

### 4. Run The Application

```bash
make run
```

## ⚙️ How It Works
The AutoPoster runs in a continuous loop, performing these steps:

1. Read Configuration Settings.
2. Iterate Over All Defined Twitter Accounts and Their Sources.
3. Fetch Content from RSS Feeds or other sources.
4. Compare fetched content against the database to check for previously published posts.
5. If a new post is found:
    - Format and publish the tweet via Twitter's API.
    - Save the newly published tweet in the database to avoid duplication.
6. Wait (Sleep()) between retries or iterations to avoid API rate limits.

## 💻 API Credentials

You’ll need these credentials for authentication with the Twitter API:

- API Key
- API Secret Key
- Access Token
- Access Token Secret

> <b>NOTE:</b> The Access Token and Access Token Secret need to have <u>read and write</u> permissions to be able to post Tweets (see below screenshot). API permissions are configurable in the [X Developer Portal](https://developer.x.com/en/portal).

<img src="https://github.com/user-attachments/assets/7d8bf440-c108-44ba-a36e-343ef64ec645" />

## License
[MIT](https://mit-license.org/)
