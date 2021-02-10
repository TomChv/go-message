# go-message

A simple Go API that manage users and message

## Entities

### User
- `userID` : **int**
- `username` : **string**
- `profilePicture` : **string**
- `password` : **string**

### Message
- `messageID` : **int**
- `message` : **string**
- `sender` : **User**

### Room
- `roomID` : **int**
- `messages` : **Message[]**


## HTTP Routes

Quick description of routes exposed by Go-message API

### Login

Route: `xxx/login`<br> 
Method: **POST**<br>
Body:
```
{
  username: "xxxx" // hashé
  password: "xxxx" // hashé
}
```
Return:
 - **404** : User doesn't exist
 - **200** : User connected
 - **500** : Internal server error
 
### Register

Route: `xxx/register`<br>
Method: **POST**<br>
Body:
```
{
  username: "xxxx" // hashé
  password: "xxxx" // hashé
  profilePicture: "xxxx" // non hashé
}
```
Return:
 - **201** : User created
 - **400** : Bad request
 - **403** : User already exist
 - **500** : Internal server error
 
### joinRoom

Route: `xxx/joinRoom?roomID=xxx`<br> 
Method: **POST**<br>
Return:
 - **200** : Room joined
 - **400** : Bad request
 - **401** : User not connected
 - **404** : Room not found
 - **500** : Internal server error
 
### createRoom

Route: `xxx/createRoom`<br>
Method: **POST**<br>
Return:
 - **201** : Room created (with roomID)
 - **401** : User not connected
 - **500** : Internal server error
 
### roomMessage (GET)

Route: `xxx/roomMessage?roomID=xxxx`<br>
Method: **GET**<br>
Return:
 - **200** : Room
 - **400** : Bad request
 - **401** : User not connected
 - **404** : Room not found
 - **500** : Internal server error
 
### roomMessage (POST)

Route: `xxx/roomMessage?roomID=xxxx`<br>
Method: **POST**<br>
Body:
```
{
  message: "xxx"
}
```
Return:
 - **201** : Message created
 - **400** : Bad request
 - **401** : User not connected
 - **404** : Room not found
 - **500** : Internal server error
 
### roomMessage (PUT)

Route: `xxx/roomMessage?roomID=xxxx&messageID=xxx`<br>
Method: **PUT**<br>
Body:
```
{
  message: "xxx"
}
```
Return:
 - **201** : Message created
 - **400** : Bad request
 - **401** : User not connected
 - **404** : Room not found
 - **500** : Internal server error 

### roomMessage (DELETE)

Route: `xxx/roomMessage?roomID=xxxx&messageID=xxxx`<br>
Method: **DELETE**<br>
Return:
 - **200** : Message deleted
 - **400** : Bad request
 - **401** : User not connected
 - **404** : Room not found
 - **500** : Internal server error