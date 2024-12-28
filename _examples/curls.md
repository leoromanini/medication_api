#### /medications
* `GET`     : Get all medications

```bash
curl -X GET 'http://localhost:4000/v1/medications'
```
```json
[
   {
      "ID":1,
      "Name":"Paracetamol",
      "Dosage":"500 mg",
      "Form":"Tablet",
      "Created":"2024-12-28T20:35:35Z",
      "LastUpdate":"2024-12-28T20:35:35Z"
   },
   {
      "ID":2,
      "Name":"Lexapro",
      "Dosage":"10 mg",
      "Form":"Tablet",
      "Created":"2024-12-28T20:35:35Z",
      "LastUpdate":"2024-12-28T20:35:35Z"
   },
   {
      "ID":3,
      "Name":"Melatonin",
      "Dosage":"5 mg",
      "Form":"Capsule",
      "Created":"2024-12-28T20:35:35Z",
      "LastUpdate":"2024-12-28T20:35:35Z"
   }
]
```

* `POST`    : Create a new medication
```bash
curl --location 'http://localhost:4000/v1/medications' \
--header 'Content-Type: application/json' \
--data '{"name": "Aspirin","dosage": "300 mg", "form": "Capsule"}'
```
``` json
{
   "ID":4,
   "Name":"Aspirin",
   "Dosage":"300 mg",
   "Form":"Capsule",
   "Created":"2024-12-28T20:48:57Z",
   "LastUpdate":"2024-12-28T20:48:57Z"
}
```

#### /medications/:id
* `GET`     : Get a medications
```bash
curl --location 'http://localhost:4000/v1/medications/2'
```
``` json
{
   "ID":2,
   "Name":"Lexapro",
   "Dosage":"10 mg",
   "Form":"Tablet",
   "Created":"2024-12-28T20:35:35Z",
   "LastUpdate":"2024-12-28T20:35:35Z"
}
````

* `PATCH`   : Update a medications

```bash
curl --location --request PATCH 'http://localhost:4000/v1/medications/3' \
--header 'Content-Type: application/json' \
--data '{
    "name" : "Lipitor"
}'
```
``` json
{
   "ID":3,
   "Name":"Lipitor",
   "Dosage":"5 mg",
   "Form":"Capsule",
   "Created":"2024-12-28T20:35:35Z",
   "LastUpdate":"2024-12-28T20:53:57Z"
}
````

* `DELETE`  : Delete a medications
```bash
curl --location --request DELETE 'http://localhost:4000/v1/medications/1'
```
``` json
{
   "ID":1,
   "Name":"Paracetamol",
   "Dosage":"500 mg",
   "Form":"Tablet",
   "Created":"2024-12-28T20:35:35Z",
   "LastUpdate":"2024-12-28T20:35:35Z"
}
```

#### /health
* `GET` : Web app healthcheck

```bash
curl --location 'http://localhost:4000/health'
```
``` json
{"status": "ok"}
```

