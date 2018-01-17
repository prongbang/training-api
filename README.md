# RESTful API (CRUD post)

## CREATE
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts
Request Method: POST
```
> Header
```
Content-Type: application/json
```
> Body
```
{
	"subject": "Hello",
	"content": "Test Content"
}
```
> Response
```
{
    "id": 1,
    "created_at": "2018-01-17T14:12:51.326479934+07:00",
    "updated_at": "2018-01-17T14:12:51.327361469+07:00",
    "subject": "Hello",
    "content": "Test Content"
}
```

## GET ALL
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts
Request Method: GET
```
> Response
```
[
    {
        "id": 1,
        "created_at": "2018-01-17T14:12:51.326479934+07:00",
        "updated_at": "2018-01-17T14:12:51.327361469+07:00",
        "subject": "Hello",
        "content": "Test Content"
    }
]
```

## GET BY ID
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts/1
Request Method: GET
```
> Response
```
{
    "id": 1,
    "created_at": "2018-01-17T14:12:51+07:00",
    "updated_at": "2018-01-17T14:12:51+07:00",
    "subject": "Hello",
    "content": "Test Content"
}
```

## PUT
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts/6
Request Method: PUT
```
> Header
```
Content-Type: application/json
```
> Body
```
{
	"subject":"Put subject create 6",
	"content":"Put content create 6"
}
```
> Response
```
{
    "id": 6,
    "created_at": "2018-01-17T14:28:10.700395654+07:00",
    "updated_at": "2018-01-17T14:28:10.70064638+07:00",
    "subject": "Put create 6",
    "content": "content put create 6"
}
```

## PATCH
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts/6
Request Method: PATCH
```
> Header
```
Content-Type: application/json
```
> Body
```
{
	"content":"Patch content create 6"
}
```
> Response
```
{
    "id": 6,
    "created_at": "2018-01-17T14:28:10+07:00",
    "updated_at": "2018-01-17T14:30:24.127595332+07:00",
    "subject": "Put create 6",
    "content": "Patch content create 6"
}
```

## DELETE
> Request
```
Request URL: http://192.168.7.76:1323/api/v1/posts/6
Request Method: DELETE
```
> Response
```
{
    "id": 6,
    "created_at": "2018-01-17T14:28:10+07:00",
    "updated_at": "2018-01-17T14:30:24+07:00",
    "subject": "Put create 6",
    "content": "Patch content create 6"
}
```

# MySQL
> Database Name: ```trainingdb```
```sql
CREATE TABLE `post` (
  `id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  `subject` varchar(200) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

```