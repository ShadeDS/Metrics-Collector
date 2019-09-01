# Metrics Collector

* [Build](#build) 
* [Rest API](#rest-api)
    * [Submit metric](#submit-metric)

Service is used to store metrics about happen events

## Build
To build this application execute `go build -buildmode=default -o <path_to_target_file>`

## Rest API
### Submit metric

This API allows to save metric in storage   
* **URI:**  `{collector_host}/api/v1/metric`  
* **Method:** `POST` 
* **Headers:**  
`Content-Type: application/json`  
* **Request body:**
```json
{
  "eventType": "<name of event type>"
}
```  
* **Success Response:**  
*Http code*: `201` - if a metric was submitted successfully  
* **Error Response:**  
*Http code*: `400` - if body content is invalid   
*Response body:* 
```json
{
  "error": "<error message>"
}
```
* **Sample call**  

Request:
```bash
curl -X POST \
  http://localhost:8080/api/v1/metric \
  -H 'Content-Type: application/json' \
  -d '{
  "eventType": "redirect"
}'
```
Response:  
*Http code*: `201`
