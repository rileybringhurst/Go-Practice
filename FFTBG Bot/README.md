FFTBG is essentially a computer playing against itself in an old Playstation game, with viewers betting fake money on the results.
Article- https://kotaku.com/final-fantasy-tactics-battleground-lets-you-gamble-fake-1840998880
Official site- https://www.twitch.tv/fftbattleground

I built the API calls to get data on the upcoming match and repurposed a Twitch chat bot library to track results and place bets.

I stopped building it at the point when all that was left was the betting logic.
I wanted to use machine learning to optimize the bets, but decided it wouldn't be the best use of my time.

I chose not to upload the chat bot part, because so much of it was copied from another Github user.
I did have fun running an infinite loop to handle chats, while also making sure to avoid spam restrictions.
