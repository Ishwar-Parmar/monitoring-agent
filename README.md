# Blockchain Monitoring Agent

## Setup Instructions

### Prerequisites
- Docker
- Docker Compose
- Go

### Running the Local Blockchain

1. **Using Anvil:**
   ```bash
   anvil
Using Docker Compose:

docker-compose up -d
Running the Monitoring Agent
Clone the repository:

git clone https://github.com/yourusername/blockchain-monitoring-agent.git
cd blockchain-monitoring-agent
Build and run the Go application:

go build -o monitor cmd/monitor/main.go
./monitor
