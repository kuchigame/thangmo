package bot

var (
	newStateMessages map[string]string = map[string]string{
		"hackernews": "Do you want to receive posts from HackerNews?",
		"rss":        "Do you want to add RSS channels?",
		"finished":   "Done. You will here from me soon (or not).\nHave a nice day!",
	}

	commandHelpMessages map[string]string = map[string]string{
		"hackernews": `
		Use this command to enable or disable HackerNews.

		To receive news items with a minimum score:
		/hackernews enable <minimum score>
		Example: /hackernews enable 200

		To disable hackernews:
		/hackernews disable
		`,
		"rss": `
		Use this command to add, remove or display your current RSS channels.

		To add a channel:
		/rss add <channel name> <channel URL>
		Example: /rss add BBC Vietnamese https://www.bbc.co.uk/vietnamese/index.xml

		To remove a channel:
		/rss remove <channel name>
		Example: /rss remove BBC Vietnamese

		To display your current channels:
		/rss list
		`,
	}
)