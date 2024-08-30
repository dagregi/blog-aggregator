# Feeds

## Create a new feed

**POST /v1/feeds**

-   Create a new feed with the provided information.
-   Authentication Required
-   Request Body

    -   name: string
    -   url: string

-   Example Request

```json
{
    "name": "My Feed",
    "url": "https://myfeed.com/index.xml"
}
```

-   Example Responses:
-   201 Created:

```json
{
    "feed": {
        "id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf",
        "created_at": "2024-08-30T07:16:18.959106Z",
        "updated_at": "2024-08-30T07:16:18.959107Z",
        "name": "My Feed",
        "url": "https://myfeed.com/index.xml",
        "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
        "last_fetched_at": null
    },
    "feed_follow": {
        "id": "20385721-63fc-4395-8825-18391d69979f",
        "feed_id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf",
        "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
        "created_at": "2024-08-30T07:16:18.99673Z",
        "updated_at": "2024-08-30T07:16:18.99673Z"
    }
}
```

-   400 Bad Request:

```json
{
    "error": "Invalid Request"
}
```

-   500 Internal Server Error:

```json
{
    "error": "Internal Server Error"
}
```

## Get all feeds

**GET /v1/feeds**

-   Get a list of all feeds.
-   Example Responses:
-   200 OK:

```json
[
    {
        "id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b",
        "created_at": "2024-08-28T10:22:39.867156Z",
        "updated_at": "2024-08-30T07:21:25.329496Z",
        "name": "Boot.Dev Blog",
        "url": "https://blog.boot.dev/index.xml",
        "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
        "last_fetched_at": "2024-08-30T07:21:25.329496Z"
    },
    {
        "id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf",
        "created_at": "2024-08-30T07:16:18.959106Z",
        "updated_at": "2024-08-30T07:16:18.959107Z",
        "name": "My Feed",
        "url": "https://myfeed.com/index.xml",
        "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
        "last_fetched_at": null
    }
]
```

-   404 Not Found:

```json
{
    "error": "Couldn't get feeds"
}
```

-   500 Internal Server Error:

```json
{
    "error": "Internal Server Error"
}
```
