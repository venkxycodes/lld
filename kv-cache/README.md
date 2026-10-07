# Problem: In-Memory Key-Value Store

Build an in-memory key-value database in Go, similar to a very small subset of Redis.

The store supports string keys and string values.

## Core API

```text
SET key value
GET key
DELETE key
EXISTS key
```

Example:

```text
SET name venkat
> OK

GET name
> venkat

EXISTS name
> true

DELETE name
> true

GET name
> nil
```

## TTL Support

Keys can optionally have an expiry.

```text
SET key value TTL 10
```

The key should become unavailable after 10 seconds.

```text
TTL key
> remaining seconds

TTL key_without_expiry
> -1

TTL missing_key
> -2
```

## Concurrency

Multiple goroutines may access the store concurrently.

For example:

```text
100 goroutines
      |
      v
SET / GET / DELETE
      |
      v
   KV Store
```

Operations must be thread-safe. There should be no corrupted state or data races.

## Constraints

- Up to 1 million keys
- Keys <= 256 bytes
- Values <= 1 MB
- High read/write concurrency
- Everything lives in memory
- Single process
- Persistence is not required

## Scope

Expose the key-value store as a Go library.

Do NOT build:

- HTTP APIs
- CLI
- Persistence
- Distributed replication

## Tests should cover

- SET and GET
- Updating an existing key
- DELETE
- EXISTS
- TTL
- Expired keys
- Keys without TTL
- Concurrent reads
- Concurrent writes
- Concurrent reads/writes
- Relevant edge cases

## Design decision

One important design decision is intentionally unspecified:

**How should expired keys actually be removed from memory?**

Choose an approach yourself and be prepared to explain:

- Why you chose it
- Its time/space complexity
- What happens when there are millions of keys
- What happens when many keys expire simultaneously
- What alternatives you considered

## Goal

Produce a correct, thread-safe implementation and be prepared to explain the design and trade-offs as if this were a machine-coding/system-design interview.
