# Sending Message To Telegram

Sending message to telegram from kibana alert

**URL** : `/api/telegrams/send`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**: 
```json
{
  "text":"Hi Webhooker!"
}
```
or
```json
{
  "text": "Hi Webhooker!",
  "token":  "..."
}
```
## Success Response

**Condition** : None

**Code** : `200 OK`

**Content example**

```json
{
  "Code": 200,
  "Status": "Success"
}
```

## Error Responses

**Condition** : None

**Code** : `200 OK`

**Headers** : None

**Content** : 

```json
{
  "Status": "Error for posting server",
  "Code": 500
}
```