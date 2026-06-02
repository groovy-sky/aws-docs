---
title: "Get the queues on the agent's routing profile in Connect Customer agent workspace"
---

# Get the queues on the agent's routing profile in Connect Customer agent workspace

Returns the list of queues on the routing profile of the agent currently logged
in to the Connect Customer agent workspace. Each `Queue` contains
`name`, `queueARN`, and `queueId`.

```typescript

async getRoutingProfileQueues(): Promise<Queue[]>

```

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getDefaultOutboundQueue()

getAvailabilityState()

All content copied from https://docs.aws.amazon.com/.
