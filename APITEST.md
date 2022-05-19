## API 測試（利用postman完成）

### [GET] `/heroes`測試
1. Header為空

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -X GET 127.0.0.1:8080/heroes
```

**Response 200**

```jsonc
{
    "heroes": [
        {
            "id": "1",
            "name": "Daredevil",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg"
        },
        {
            "id": "2",
            "name": "Thor",
            "image": "http://x.annihil.us/u/prod/marvel/i/mg/5/a0/537bc7036ab02/standard_xlarge.jpg"
        },
        {
            "id": "3",
            "name": "Iron Man",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/a0/55b6a25e654e6/standard_xlarge.jpg"
        },
        {
            "id": "4",
            "name": "Hulk",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0/standard_xlarge.jpg"
        }
    ]
}
```
2. 正確的Name和Password

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: hahow" -H "Password: rocks" -X GET 127.0.0.1:8080/heroes
```

**Response 200**

```jsonc
{
    "heroes": [
        {
            "id": "1",
            "name": "Daredevil",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
            "profile": {
                "str": 2,
                "int": 7,
                "agi": 9,
                "luk": 7
            }
        },
        {
            "id": "2",
            "name": "Thor",
            "image": "http://x.annihil.us/u/prod/marvel/i/mg/5/a0/537bc7036ab02/standard_xlarge.jpg",
            "profile": {
                "str": 8,
                "int": 2,
                "agi": 5,
                "luk": 9
            }
        },
        {
            "id": "3",
            "name": "Iron Man",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/a0/55b6a25e654e6/standard_xlarge.jpg",
            "profile": {
                "str": 6,
                "int": 9,
                "agi": 6,
                "luk": 9
            }
        },
        {
            "id": "4",
            "name": "Hulk",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0/standard_xlarge.jpg",
            "profile": {
                "str": 10,
                "int": 1,
                "agi": 4,
                "luk": 2
            }
        }
    ]
}
```
3. 錯的Name及Password

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: whatever" -H "Password: it is" -X GET 127.0.0.1:8080/heroes
```

**Response 200**

```jsonc
{
    "heroes": [
        {
            "id": "1",
            "name": "Daredevil",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg"
        },
        {
            "id": "2",
            "name": "Thor",
            "image": "http://x.annihil.us/u/prod/marvel/i/mg/5/a0/537bc7036ab02/standard_xlarge.jpg"
        },
        {
            "id": "3",
            "name": "Iron Man",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/a0/55b6a25e654e6/standard_xlarge.jpg"
        },
        {
            "id": "4",
            "name": "Hulk",
            "image": "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0/standard_xlarge.jpg"
        }
    ]
}
```

### GET] `/heroes/:heroId`測試

1. 帶存在的heroId 

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -X GET 127.0.0.1:8080/heroes/1
```

**Response 200**

```jsonc
{
  "id": "1",
  "name": "Daredevil",
  "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg"
}
```

2. 帶不存在的heroId

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -X GET 127.0.0.1:8080/heroes/3939889
```

**Response 404**

```
"NotFound"
```

3. 正確的Name及Password及存在的heroId

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: hahow" -H "Password: rocks" -X GET 127.0.0.1:8080/heroes/1
```

**Response 200**

```jsonc
{
    "id": "1",
    "name": "Daredevil",
    "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
    "profile": {
        "str": 2,
        "int": 7,
        "agi": 9,
        "luk": 7
    }
}
```
4. 錯誤的Name及Password及存在的heroId

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: whatever" -H "Password: it is" -X GET 127.0.0.1:8080/heroes/1
```
**Response 200**

```jsonc
{
    "id": "1",
    "name": "Daredevil",
    "image": "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg"
}
```

5. 正確的Name及Password及不存在的heroId

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: hahow" -H "Password: rocks" -X GET 127.0.0.1:8080/heroes/3939889
```

**Response 404**

```
"NotFound"
```

6. 錯誤的Name及Password及不存在的heroId

**Request**

```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -H "Name: whatever" -H "Password: it is" -X GET 127.0.0.1:8080/heroes/3939889
```

**Response 404**

```
"NotFound"
```