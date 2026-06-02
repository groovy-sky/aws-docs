---
title: "Get the extension of the agent in Connect Customer agent workspace"
---

# Get the extension of the agent in Connect Customer agent workspace

Returns phone number of the agent currently logged in to the Connect Customer agent workspace. This is
the phone number that is dialed by the Connect Customer to connect calls to the agent for incoming
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
