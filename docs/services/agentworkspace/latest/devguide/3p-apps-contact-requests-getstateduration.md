---
title: "Get the duration of the contact state in Connect Customer agent workspace"
---

# Get the duration of the contact state in Connect Customer agent workspace

Returns the duration of the contact state in milliseconds relative to local time,
in the Connect Customer agent workspace. This takes into account time skew between the JS client and the
Connect Customer backend servers.

```typescript

async getStateDuration(contactId: string): Promise<number>

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getQueueTimestamp()

isPreviewMode()

All content copied from https://docs.aws.amazon.com/.
