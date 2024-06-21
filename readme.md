# Create a New Post

Create a new post.

**Endpoint**: `posts/create-post`

**Method**: `POST`

## Headers

- `Content-Type: application/json`

## Payload

```json
{
    "title": "This is title 5",
    "body": "This is the body of 5!"
}
```
## Response

```json
{
    "data": {
        "body": "This is title 5",
        "id": "73761e7c-2216-4f59-88be-32c8983ccad7",
        "title": "This is the body of 5!"
    },
    "status": "Post created!"
}
```
---

# Fetch All Posts

Fetch all the posts available.

**Endpoint**: `posts/fetch-all-post`

**Method**: `GET`

## Headers

- `Content-Type: application/json`

## Response

```json
{
  "data": [
    {
      "body": "This is title 1",
      "id": "7ab0d44e-f6c1-4ef5-9b40-8ecabfea5541",
      "title": "This is the body of 1!"
    },
    {
      "body": "This is a title 2!",
      "id": "f206646c-b0f6-4daa-a8ea-bfec1a580d09",
      "title": "This is the body 2!"
    },
    {
      "body": "This is title 3",
      "id": "e349b35d-bb33-4868-9529-2015b19c90fb",
      "title": "This is the body of 3!"
    },
    {
      "body": "This is title 4",
      "id": "9b8174fb-c78f-4fe9-93ff-845ca029a9d6",
      "title": "This is the body of 4!"
    },
    {
      "body": "This is a title 5!",
      "id": "73761e7c-2216-4f59-88be-32c8983ccad7",
      "title": "This is the body 5!"
    }
  ],
  "status": "All posts fetched"
}
```
---

# Fetch a Post by ID

Fetch a post by its id.

**Endpoint**: `posts/fetch-post/:id`

**Method**: `GET`

## Headers

- `Content-Type: application/json`

## Response

```json
{
  "data": {
    "body": "This is title 2",
    "id": "f206646c-b0f6-4daa-a8ea-bfec1a580d09",
    "title": "This is the body of 2!"
  },
  "status": "Post fetched"
}
```
---

# Update Post

Update a post.

**Endpoint**: `posts/update-post/:id`

**Method**: `PATCH`

## Headers

- `Content-Type: application/json`

## Payload

```json
{
  "title": "This is a title 55!",
  "body": "This is the body 55!"
}
```
## Response

```json
{
    "data": {
        "body": "This is title 55",
        "id": "73761e7c-2216-4f59-88be-32c8983ccad7",
        "title": "This is the body of 55!"
    },
    "status": "Post updated!"
}
```
---

# Delete a Post 

Delete a post .

**Endpoint**: `posts/delete-post/:id`

**Method**: `DELETE`

## Headers

- `Content-Type: application/json`

## Response

```json
{
  "status": "Post deleted"
}
```
---