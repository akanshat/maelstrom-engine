# Distributed Systems Engine
A Go-based distributed systems node built to solve the Fly.io Maelstrom Challenge Series, a hands-on curriculum for implementing core distributed algorithms benchmarked against [Jepsen Maelstrom's](https://github.com/jepsen-io/maelstrom) real-time fault-injection harness.

This repository tracks my implementation of key concepts from Designing Data-Intensive Applications (DDIA)—including thread-safe RPC routing, uncoordinated identifier generation, state replication over lossy networks, and distributed consensus.

## Architecture Overview

               ┌──────────────────────────────────────────────┐
               │           Maelstrom CLI Test Runner          │
               └──────────────────────┬───────────────────────┘
                                      │
                         JSON RPC over STDIN / STDOUT
                                      │
               ┌──────────────────────┴───────────────────────┐
               │              cmd/node/main.go                │
               │   - Asynchronous Message Handler Loop        │
               │   - Thread-safe Lazy Initialization (Once)   │
               └──────────────────────┬───────────────────────┘
                                      │
                                      ▼
                      ┌──────────────────────────────┐
                      │    pkg/idgen (Snowflake)    │
                      │  - Bit-packed Timestamps    │
                      │  - Worker Node Parsing      │
                      │  - Atomic Sequence Counter  │
                      └──────────────────────────────┘

## Repository Structure
```
maelstrom-engine/
├── cmd/
│   └── node/
│       └── main.go       # Entry point; wires Maelstrom RPC routes & thread-safe handlers
├── pkg/
│   └── idgen/
│       └── idgen.go      # Snowflake generator wrapper & worker ID string parser
├── go.mod                # Module dependencies (bwmarrin/snowflake, maelstrom SDK)
└── README.md
```

## How to Build & Test
### Prerequisites: 
- Go 1.21+
- Java Runtime Environment (JRE) (Required by Jepsen Maelstrom CLI harness)
- Maelstrom CLI Tool installed in your executable path 
`go build -o ~/go/bin/maelstrom-engine ./cmd/node/main.go`

### Run Maelstrom Verification Suites

#### 1. Echo Challenge Test
`maelstrom test -w echo --bin ~/go/bin/maelstrom-engine --node-count 1 --time-limit 10`
#### 2. Unique ID Generator Test (With Partition Nemesis)
`maelstrom test -w unique-ids --bin ~/go/bin/maelstrom-engine --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition`

## Challenge Roadmap
- [x] Challenge 1: Echo Protocol
- [x] Challenge 2: Distributed Unique ID Generator
- [ ] Challenge 3: Broadcast & Gossip Protocols
- [ ] Challenge 4: Grow-Only Counter (CRDTs)
- [ ] Challenge 5: Kafka-Style Multi-Node Replication Log
- [ ] Challenge 6: Linearizable Key-Value Store