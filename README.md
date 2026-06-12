# Go URL Shortener

A high-performance, ultra-low latency URL shortener built in Go. It converts long URLs into compressed, shareable links, leveraging an in-memory database to handle instant HTTP redirections under heavy loads.

## ✨ Features
* **Sub-Millisecond Redirects:** Delivers near-instantaneous routing by serving lookups entirely from memory.
* **Custom Aliases & TTL:** Supports user-defined short keys and automatically expires links using Redis TTL.
* **API Rate Limiting:** Includes built-in middleware to track client IPs and prevent endpoint abuse.

## 🛠️ Tech Stack
* **Backend:** Go (Golang)
* **Framework:** Go-Fiber (`fasthttp` engine)
* **Database:** Redis (In-memory key-value store)
