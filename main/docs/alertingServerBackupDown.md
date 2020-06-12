# Alerting Server Backup Down

Alerting API from kibana to Autoticket of other system

**URL** : `/api/myalert/outofcredit`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**: None

## Success Response

**Condition** : None

**Code** : `200 OK`

**Content example**

```json
{
  "Code": 200,
  "Status": "Success: ..."
}
```

## Error Responses

**Condition** : None

**Code** : `200 OK`

**Headers** : None

**Content** : 

```json
{
  "Code": 500,
  "Status": "Failled: ..."
}
```