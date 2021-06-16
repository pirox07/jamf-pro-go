# jamf-pro-go
Jamf Pro API Client

## Usage

```
export JAMF_BASE_URL=https://<your tenant>.jamfcloud.com
export JAMF_USER=<Jamf Pro Username>
export JAMF_USER_PASSWORD=<Jamf Pro User Password>
```

[exapmple code](./examples/main.go)

```
package main

import (
	"fmt"
	"os"

	"https://github.com/pirox07/jamf-pro-go"
)


func main() {
	url := os.Getenv("JAMF_BASE_URL")
	userName := os.Getenv("JAMF_USER")
	password := os.Getenv("JAMF_USER_PASSWORD")
	conf, err := jamf.NewConfig(url, userName, password)
	if err !=nil{
		fmt.Println(err.Error())
	}
	client := jamf.NewClient(conf)
	
	...
}
```

## References

- [Jamf Pro API](https://www.jamf.com/developers/apis/jamf-pro/reference/)
- [Jamf Classic API](https://www.jamf.com/developers/apis/classic/reference/)
- [Jamf Pro の API Client をつくった](https://note.com/pirox/n/n6b08712720a2)

## APIs

- [Policies](https://www.jamf.com/developers/apis/classic/reference/#/policies)
  - `GET /policies`: Finds all policies
  - `GET /policies/id/{id}`: Finds policies by ID
  - `POST /policies/id/{id}`: Creates a new policy by ID
  - `PUT /policies/id/{id}`: Updates an existing policy by ID
  - `DELETE /policies/id/{id}`: Deletes a policy by ID

- [Scrips](https://www.jamf.com/developers/apis/jamf-pro/reference/#/scripts)
  - `GET /v1/scripts`: Search for sorted and paged Scripts
  - `GET /v1/scripts/{id}`: Retrieve a full script object
  - `POST /v1/scripts/{id}`: Creates a script
  - `PUT /v1/scripts/{id}`: Replace the script at the id with the supplied information
  - `DELETE /v1/scripts/{id}`: Delete a Script at the specified id
