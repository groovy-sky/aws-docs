# Get the timestamp of the contact in Amazon Connect Agent Workspace

Returns a `Date` object with the timestamp associated with when the contact
was placed in the queue in the Amazon Connect agent workspace.

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
