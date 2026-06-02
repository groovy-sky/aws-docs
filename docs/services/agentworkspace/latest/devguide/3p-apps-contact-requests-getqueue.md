---
title: "Get the queue of the contact in Connect Customer agent workspace"
---

# Get the queue of the contact in Connect Customer agent workspace

Returns the queue associated with the contact in the Connect Customer agent workspace. The `Queue`
object has the following fields:

- `name`: The name of the queue.

- `queueARN`: The ARN of the queue.

- `queueId`: Alias for `queueARN`.

```typescript

async getQueue(contactId: string): Promise<Queue>

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getPreviewConfiguration()

getQueueTimestamp()

All content copied from https://docs.aws.amazon.com/.
