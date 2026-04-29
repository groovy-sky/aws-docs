---
title: "Get the extension of the agent in Amazon Connect Customer AI agent workspace"
---

# Get the extension of the agent in Amazon Connect Customer AI agent workspace

Returns phone number of the agent currently logged in to the Amazon Connect Customer AI agent workspace. This is
the phone number that is dialed by the Amazon Connect to connect calls to the agent for incoming
and outgoing calls if soft phone is not enabled.

```typescript

async getExtension(): Promise<string | null>

```

**Permissions required:**

```typescript

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getChannelConcurrency()

getName()

All content copied from https://docs.aws.amazon.com/.
