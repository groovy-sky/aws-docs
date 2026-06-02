---
title: "Get the routing profile of the agent in Connect Customer agent workspace"
---

# Get the routing profile of the agent in Connect Customer agent workspace

Returns the routing profile of the agent currently logged in to the Connect Customer agent workspace.
The routing profile contains the following fields:

- `channelConcurrencyMap`: See agent. [Get the limit of contacts for the agent in Connect Customer agent workspace](3p-apps-agent-requests-getchannelconcurrency.md) for more
info.

- `defaultOutboundQueue`: The default queue which should be
associated with outbound contacts. See queues for details on properties.

- `name`: The name of the routing profile.

- `queues`: The queues contained in the routing profile. Each queue
object has the following properties:

- `name`: The name of the queue.

- `queueARN`: The ARN of the queue.

- `queueId`: Alias for queueARN.

- `routingProfileARN`: The routing profile ARN.

- `routingProfileId`: Alias for `routingProfileARN`.

```typescript

async getRoutingProfile(): Promise<AgentRoutingProfile>

```

**Permissions required:**

```typescript

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getName()

getState() - Deprecated

All content copied from https://docs.aws.amazon.com/.
