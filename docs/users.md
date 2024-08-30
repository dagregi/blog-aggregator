# Users

## Create a new user

**POST /v1/users**

-   Create a new user with the provided information.
-   Request Body

    -   name: string

-   Example Request:

```json
{
    "name": "John Doe"
}
```

-   Example Responses:
-   200 OK:

```json
{
    "id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
    "created_at": "2024-08-30T06:44:02.734438Z",
    "updated_at": "2024-08-30T06:44:02.734438Z",
    "name": "John Doe",
    "api_key": "fe4713950b883851a44e918668e82c87fe8407a4bed2fc839d086587956ea9f9"
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
    "error": "Couldn't create user"
}
```

## Get user

**GET /v1/users**

-   Get the information of a user.
-   Authentication Required

-   Example Responses:
    -   200 OK:

```json
{
    "id": "e2acf4f3-277c-4a4b-ad92-b45c4094d57a",
    "created_at": "2024-08-30T06:44:02.734438Z",
    "updated_at": "2024-08-30T06:44:02.734438Z",
    "name": "John Doe",
    "api_key": "fe4713950b883851a44e918668e82c87fe8407a4bed2fc839d086587956ea9f9"
}
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
    "error": "Couldn't get user"
}
```
