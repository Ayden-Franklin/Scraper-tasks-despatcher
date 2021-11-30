# Description
## User Story
Currently, collecting data is in both centralized and decentralized fashion. Through the passive data collection vector, data is collected through Extension user activity on Amazon or Walmart, but only when users take action and only for the products they are actively looking at. I feel it poses little risk to collect data through the clients im a more comprehensive manner.

## Feature
This is the back-end service for an extension distributed decentralized data collection solution.

## Core functionalities
- Prepare the ASINs for dispatching.
- Sort the ASINs following the current strategy
- Construct tasks for all the ASINs to track the processing status
- Dispatch a batch of ASINs for a client
- Receive notifications if any ASINs has done (whether successful or failed)
- Track clients’ behavior, contribution and performance





# How to run
1. Clone this project and run `go run cmd/dispatcher-server/main.go`
2. (Optional) Run `go test ./...` for unit tests.
