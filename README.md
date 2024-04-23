# blog_aggregator

## Description
API provides way for users to aggregate blog posts from followed RSS feeds.

## Prereqs
- Postgres database connection
- API testing application such as postman
  
## Quick Start
- Clone the repository
- Add a .env file to the project directory with the contents
```
  PORT="<port number for server to listen to>"
  CONN=<postgres connection string>
```
- Run the precompiled binary
```
./out
```

## Endpoints

### /v1/users
#### POST - Create a new user entity. Provide a user name in the body. Return a response with an api key for further functionality. 
```
//REQUEST BODY
{
  "name": "Test"
}

//RESPONSE BODY
{
    "id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
    "created_at": "2024-04-22T19:16:09.365623Z",
    "updated_at": "2024-04-22T19:16:09.365623Z",
    "name": "Test",
    "api_key": "415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"
}
```

#### GET - Return info about the user entity. Need to provide the previously given api key for that user entity as a key value pairing in the header.
```
//REQUEST HEADER
"Authorization: ApiKey 415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"

//RESPONSE BODY
{
    "id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
    "created_at": "2024-04-22T19:16:09.365623Z",
    "updated_at": "2024-04-22T19:16:09.365623Z",
    "name": "Test",
    "api_key": "415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"
}
```

### /v1/feeds
#### POST - Create a new feed. Provide a link to an RSS blog feed's xml resource in the request body, and a api key in the request header. A response of the created feed entity and a feed follow entity for that user/feed pairing.
```
//REQUEST BODY
{
  "name": "ESPN Poker Blog",
  "url": "https://www.espn.com/espn/rss/poker/master"
}
//REQUEST HEADER
"Authorization: ApiKey 415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"


//RESPONSE BODY
{
    "feed": {
        "id": "1d82fca6-91dd-496c-8f3e-fb0b865d976b",
        "created_at": "2024-04-22T19:49:14.283358Z",
        "updated_at": "2024-04-22T19:49:14.283359Z",
        "name": "ESPN Poker Blog",
        "url": "https://www.espn.com/espn/rss/poker/master",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "last_fetched_at": null
    },
    "feed_follow": {
        "id": "b9d91e64-f0b3-40cf-9200-9b3b414e3a70",
        "created_at": "2024-04-22T19:49:14.284365Z",
        "updated_at": "2024-04-22T19:49:14.284365Z",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "feed_id": "1d82fca6-91dd-496c-8f3e-fb0b865d976b"
    }
}
```

#### GET - Returns a list of all feed entities
```
//RESPONSE BODY
[
    {
        "id": "03c3402b-f18b-4a8c-8c1c-8e94abec669c",
        "created_at": "2024-04-05T20:38:38.833911Z",
        "updated_at": "2024-04-22T15:30:38.749411Z",
        "name": "The espn.com Blog",
        "url": "https://www.espn.com/espn/rss/nfl/news/index.xml",
        "user_id": "9be3eeec-b520-4ed0-a292-2a673f38b8e2",
        "last_fetched_at": "2024-04-22T15:30:38.749411Z"
    },
    {
        "id": "1d82fca6-91dd-496c-8f3e-fb0b865d976b",
        "created_at": "2024-04-22T19:49:14.283358Z",
        "updated_at": "2024-04-22T15:30:38.755953Z",
        "name": "ESPN Poker Blog",
        "url": "https://www.espn.com/espn/rss/poker/master",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "last_fetched_at": "2024-04-22T15:30:38.755953Z"
    },
    {
        "id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228",
        "created_at": "2024-04-06T01:25:38.194026Z",
        "updated_at": "2024-04-22T15:30:38.756496Z",
        "name": "NBA Blog",
        "url": "https://www.espn.com/espn/rss/nba/news/index.xml",
        "user_id": "38d6de49-e496-44f8-98c8-209ccc4ea899",
        "last_fetched_at": "2024-04-22T15:30:38.756496Z"
    },
    {
        "id": "a3216665-206c-401e-874e-136a6ebc18f4",
        "created_at": "2024-04-06T01:14:10.311108Z",
        "updated_at": "2024-04-22T15:30:38.757347Z",
        "name": "SOCCER Blog",
        "url": "https://www.espn.com/espn/rss/soccer/news/index.xml",
        "user_id": "38d6de49-e496-44f8-98c8-209ccc4ea899",
        "last_fetched_at": "2024-04-22T15:30:38.757347Z"
    },
]
```

### /v1/feed_follows
#### POST - Create a new feed_follow entity. Which is essentially a user following a feed. Requires ApiKey.
```
//REQUEST BODY
{
  "feed_id": "a3216665-206c-401e-874e-136a6ebc18f4"
}
//REQUEST HEADER
"Authorization: ApiKey 415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"

//RESPONSE BODY
{
    "id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
    "created_at": "2024-04-22T19:16:09.365623Z",
    "updated_at": "2024-04-22T19:16:09.365623Z",
    "name": "Test",
    "api_key": "415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"
}
```

#### GET - Return a list of feed_follow entities for a user. Requires ApiKey.
```
//REQUEST HEADER
"Authorization: ApiKey 415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"

//RESPONSE BODY
[
    {
        "id": "0ac48f10-9ed4-4742-a208-5b67d66a402c",
        "created_at": "2024-04-22T19:44:32.136531Z",
        "updated_at": "2024-04-22T19:44:32.136531Z",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "feed_id": "936a617e-4c33-42c7-9da6-39ce3db8ea00"
    },
    {
        "id": "b9d91e64-f0b3-40cf-9200-9b3b414e3a70",
        "created_at": "2024-04-22T19:49:14.284365Z",
        "updated_at": "2024-04-22T19:49:14.284365Z",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "feed_id": "1d82fca6-91dd-496c-8f3e-fb0b865d976b"
    },
    {
        "id": "910cf597-5a57-4045-b438-402f5948deba",
        "created_at": "2024-04-22T20:33:40.351188Z",
        "updated_at": "2024-04-22T20:33:40.351188Z",
        "user_id": "977916ff-be87-4bed-bcaf-5000b11c5f57",
        "feed_id": "a3216665-206c-401e-874e-136a6ebc18f4"
    }
]
```

#### DELETE - Deletes the feed_follow entity for a user. Essentially unfollowing the feed. Requires ApiKey and the feed_follow id.
```
//URL: /v1/feed_follows/{feed_follow_id} -> http://localhost:8080/v1/feed_follows/910cf597-5a57-4045-b438-402f5948deba
//REQUEST HEADER
"Authorization: ApiKey 415d2f56b49ffd8433ee815edaa0989e3fd27f4c2af7da2025bfebff90d041c6"

//RESPONSE BODY
{}
```

### /v1/posts
#### GET - Return a list of posts aggregated across users various followed feeds.
```
//REQUEST HEADER
"Authorization: ApiKey 289e0b03f6141281725d3d7d30038f6ab0511089077c546b7365bb4c05c43a62"

//RESPONSE BODY
[
    {
        "id": "c2ffd992-3bc1-48fa-b4df-72ed7fbbc6fd",
        "created_at": "2024-04-23T00:24:25.341143Z",
        "updated_at": "2024-04-23T00:24:25.341143Z",
        "title": "NBA breaks 6 ties among 15 teams for June draft",
        "url": "https://www.espn.com/nba/story/_/id/40001243/nba-breaks-six-ties-15-teams-june-draft",
        "description": "The NBA broke six ties among 15 teams Monday to determine the order of selection for the 2024 NBA Draft, which will be held June 26 and 27, and the Hornets won a tiebreaker with the Trail Blazers, tentatively slotting them third and fourth respectively.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "6ab921c7-d47d-437b-a738-3a40edbde944",
        "created_at": "2024-04-23T00:24:25.341663Z",
        "updated_at": "2024-04-23T00:24:25.341663Z",
        "title": "Buffs' Williams, likely lottery pick, to enter draft",
        "url": "https://www.espn.com/nba/story/_/id/39998166/colorado-cody-williams-expected-lottery-pick-enter-nba-draft",
        "description": "Colorado freshman guard Cody Williams -- an expected lottery pick -- is entering the 2024 NBA draft, he told ESPN on Monday.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "1d1ca0d8-2573-4929-8302-383a7132fb85",
        "created_at": "2024-04-23T00:24:25.342574Z",
        "updated_at": "2024-04-23T00:24:25.342574Z",
        "title": "Projected No. 1 pick Risacher to enter NBA draft",
        "url": "https://www.espn.com/nba/story/_/id/39995286/projected-no-1-pick-zaccharie-risacher-enter-nba-draft",
        "description": "Projected No. 1 pick and current JL Bourg France player Zaccharie Risacher, 19, has submitted paperwork to make him eligible for the 2024 NBA draft.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "1577533e-9023-4d7b-9bb3-bf23164c305c",
        "created_at": "2024-04-23T00:24:25.343242Z",
        "updated_at": "2024-04-23T00:24:25.343242Z",
        "title": "Nets officially hire Fernandez as head coach",
        "url": "https://www.espn.com/nba/story/_/id/39995598/brooklyn-nets-officially-hire-jordi-fernandez-head-coach",
        "description": "Jordi Fernandez, who served as an associate head coach for the Kings for the past two seasons, was officially hired as head coach of the Brooklyn Nets.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "f8a6d934-085a-4a0f-bb5c-34d200a88486",
        "created_at": "2024-04-23T00:24:25.34398Z",
        "updated_at": "2024-04-23T00:24:25.34398Z",
        "title": "'Setting the tone': Dame's 35-pt. half lifts Bucks",
        "url": "https://www.espn.com/nba/story/_/id/39992727/damian-lillard-35-point-first-half-bucks-game-1-win-pacers",
        "description": "Damian Lillard's first-half masterpiece carried the Bucks, who were without Giannis Antetokounmpo, to a win over the Pacers on Sunday in Game 1 of their series.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "d0d8e4ea-163b-4b1e-b195-37c6a3cd7747",
        "created_at": "2024-04-23T00:24:25.344614Z",
        "updated_at": "2024-04-23T00:24:25.344614Z",
        "title": "Harden shows he can 'still score with best of 'em'",
        "url": "https://www.espn.com/nba/story/_/id/39991975/clippers-score-big-game-1-win-james-harden-28",
        "description": "James Harden netted 20 of his 28 points in the opening half to help the Clippers build a 29-point lead en route to a 109-97 win over the Mavericks in Game 1. \"Still can score with the best of 'em,\" he said.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    },
    {
        "id": "8a699987-1e21-4536-a615-806b32f8de98",
        "created_at": "2024-04-23T00:24:25.253556Z",
        "updated_at": "2024-04-23T00:24:25.253556Z",
        "title": "Matt Ryan retires: 'Honored' to do so as a Falcon",
        "url": "https://www.espn.com/nfl/story/_/id/39995362/matt-ryan-former-qb-falcons-colts-officially-retires",
        "description": "Former Falcons and Colts quarterback Matt Ryan announced his retirement Monday, officially ending his NFL career.",
        "published_at": null,
        "feed_id": "03c3402b-f18b-4a8c-8c1c-8e94abec669c"
    },
    {
        "id": "6a4ac58f-3d71-41c1-994c-a818db824a38",
        "created_at": "2024-04-23T00:24:25.254144Z",
        "updated_at": "2024-04-23T00:24:25.254144Z",
        "title": "Irsay won't be in Colts draft room amid recovery",
        "url": "https://www.espn.com/nfl/story/_/id/39997462/colts-owner-jim-irsay-draft-room-amid-recovery",
        "description": "Colts owner Jim Irsay will not be in the team's draft room this week as he recovers from lower back surgery, adding that team affairs are running per usual.",
        "published_at": null,
        "feed_id": "03c3402b-f18b-4a8c-8c1c-8e94abec669c"
    },
    {
        "id": "c633abf6-6840-4c64-8762-712947932a76",
        "created_at": "2024-04-23T00:24:25.254842Z",
        "updated_at": "2024-04-23T00:24:25.254842Z",
        "title": "Reinstated DE Toney released by Commanders",
        "url": "https://www.espn.com/nfl/story/_/id/39995818/shaka-toney-reinstated-gambling-ban-cut-commanders",
        "description": "The Commanders have released defensive end Shaka Toney, who was recently reinstated from a season-long gambling suspension.",
        "published_at": null,
        "feed_id": "03c3402b-f18b-4a8c-8c1c-8e94abec669c"
    },
    {
        "id": "672804cd-7ce3-41da-b6c6-947577ee940d",
        "created_at": "2024-04-23T00:24:25.34036Z",
        "updated_at": "2024-04-23T00:24:25.34036Z",
        "title": "Maxey a go as 76ers look to even series vs. Knicks",
        "url": "https://www.espn.com/nba/story/_/id/39995661/76ers-tyrese-maxey-questionable-game-2-due-illness",
        "description": "76ers All-Star guard Tyrese Maxey was in the starting lineup for Monday night's Game 2 against the Knicks after initially being listed as questionable due to an illness.",
        "published_at": null,
        "feed_id": "d31a85dd-9d40-4f9d-8e7b-76b9d241e228"
    }
]
```
