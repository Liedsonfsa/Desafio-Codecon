# Desafio Codecon

The objective of this challenge is to develop an API that receives a JSON file with 100,000 users and offers high-performance and well-structured endpoints for data analysis.

To learn more about the challenge, click on the link <a href="https://github.com/codecon-dev/desafio-1-1s-vs-3j">codecon-dev</a>

The developed solution was made using only Go without any framework

## Endpoints

### `POST /users`
Receives and stores users in memory.

### `GET /superusers`
Returns all users with a score greater than 900 and who are active.

### `GET /top-countries`
Returns the 5 countries with the highest number of superusers.

### `GET /team-insights`
Returns: total members, leaders, completed projects and % of active members.

### `GET /active-users-per-day`
Counts how many logins happened by date.