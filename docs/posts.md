# Posts

## Get all posts

**GET /v1/posts**

-   Get a list of all posts for the user.
-   Authentication Required
-   Parameters:
    -   limit: int(optional)
-   Example Responses:

-   200 OK:

```json
[
    {
        "id": "626910d6-e78e-4a44-88d6-de91532de709",
        "created_at": "2024-08-28T17:54:25.142119Z",
        "updated_at": "2024-08-28T17:54:25.14212Z",
        "title": "The Boot.dev Beat. August 2024",
        "url": "https://blog.boot.dev/news/bootdev-beat-2024-08/",
        "description": "I&rsquo;m writing this one a bit early because I&rsquo;m going out of town next week&hellip; better early than late?",
        "published_at": "2024-07-26T00:00:00Z",
        "feed_id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b"
    },
    {
        "id": "78109e37-f454-4f7b-bead-23441bc461d7",
        "created_at": "2024-08-28T17:54:25.12697Z",
        "updated_at": "2024-08-28T17:54:25.126971Z",
        "title": "Is there a course on <insert technology>?",
        "url": "https://blog.boot.dev/education/is-there-a-course-on/",
        "description": "In the Boot.dev Discord server, this question appears now and then. Will there be a Rust course?",
        "published_at": "2024-07-26T00:00:00Z",
        "feed_id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b"
    },
    {
        "id": "7e273b9e-5811-4285-8ee5-4433c8e92b83",
        "created_at": "2024-08-28T17:54:25.146Z",
        "updated_at": "2024-08-28T17:54:25.146001Z",
        "title": "Are Pointers in Go Faster Than Values?",
        "url": "https://blog.boot.dev/golang/pointers-faster-than-values/",
        "description": "I was recently working on a lesson about pointer performance for Boot.dev&rsquo;s Golang course when I found myself repeating some advice I&rsquo;ve given many times before.",
        "published_at": "2024-07-19T00:00:00Z",
        "feed_id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b"
    },
    {
        "id": "8a463f7b-8960-43dd-a182-5ebcfc722af4",
        "created_at": "2024-08-28T17:54:25.148876Z",
        "updated_at": "2024-08-28T17:54:25.148876Z",
        "title": "The Boot.dev Beat. July 2024",
        "url": "https://blog.boot.dev/news/bootdev-beat-2024-07/",
        "description": "One million lessons. Well, to be precise, you have all completed 1,122,050 lessons just in June.",
        "published_at": "2024-07-10T00:00:00Z",
        "feed_id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b"
    },
    {
        "id": "540b6f46-92d8-4c53-8462-90a9004adb8e",
        "created_at": "2024-08-28T17:54:25.151877Z",
        "updated_at": "2024-08-28T17:54:25.151877Z",
        "title": "The Boot.dev Beat. June 2024",
        "url": "https://blog.boot.dev/news/bootdev-beat-2024-06/",
        "description": "ThePrimeagen&rsquo;s new Git course is live. A new boss battle is on the horizon, and we&rsquo;ve made massive speed improvements to the site.",
        "published_at": "2024-06-05T00:00:00Z",
        "feed_id": "1b7a0f51-1f68-4fd3-a061-93936c210f7b"
    }
]
```

-   401 Unauthorized:

```json
{
    "error": "Couldn't find API key"
}
```

-   404 Not Found:

```json
{
    "error": "Couldn't get posts"
}
```
