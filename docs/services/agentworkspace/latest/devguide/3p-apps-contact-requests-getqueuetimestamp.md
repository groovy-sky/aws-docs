---
title: "Get the timestamp of the contact in Connect Customer agent workspace"
---

# Get the timestamp of the contact in Connect Customer agent workspace

Returns a `Date` object with the timestamp associated with when the contact
was placed in the queue in the Connect Customer agent workspace.

```typescript

async getQueueTimestamp(contactId: string): Promise<Date | undefined>

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getQueue()

getStateDuration()

All content copied from https://docs.aws.amazon.com/.
