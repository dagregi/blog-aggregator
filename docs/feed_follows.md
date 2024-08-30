# Feed Follows

## Create a new feed follow

**POST /v1/feed_follows**

-   Create a new feed follow with the provided information.
-   Authentication Required
-   Request Body
    -   feed_id: string
-   Example Request:

```json
{
    "feed_id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf"
}
```

-   Example Responses:

-   201 Created:

```json
{
    "id": "20385721-63fc-4395-8825-18391d69979f",
    "feed_id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf",
    "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
    "created_at": "2024-08-30T07:16:18.99673Z",
    "updated_at": "2024-08-30T07:16:18.99673Z"
}
```

-   400 Bad Request:

```json
{
    "error": "Invalid Request"
}
```

-   401 Unauthorized:

```json
{
    "error": "Couldn't find API key"
}
```

-   500 Internal Server Error:

```json
{
    "error": "Couldn't follow feed"
}
```

## Get all feed follows

**GET /v1/feed_follows**

-   Get a list of all user feed follows.
-   Authentication Required
-   Example Responses:

-   200 OK:

```json
[
    {
        "id": "20385721-63fc-4395-8825-18391d69979f",
        "feed_id": "4ffa35b2-6d17-45d3-a071-680a25ce96bf",
        "user_id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
        "created_at": "2024-08-30T07:16:18.99673Z",
        "updated_at": "2024-08-30T07:16:18.99673Z"
    }
]
```

-   401 Unauthorized:

```json
{
    "error": "Couldn't find API key"
}
```

-   500 Internal Server Error:

```json
{
    "error": "Couldn't get follows"
}
```

## Delete a feed follow

**DELETE /v1/feed_follows/{feedFollowID}**

-   Delete a feed follow by ID.
-   Parameters

    -   feedFollowID: ID of the feed follow to delete

-   Example Responses:

-   204 No Content:

```json
{}
```

-   400 Bad Request:

```json
{
    "error": "Invalid Request"
}
```

-   401 Unauthorized:

```json
{
    "error": "Couldn't find API key"
}
```

-   500 Internal Server Error:

```json
{
    "error": "Couldn't delete follow"
}
```
