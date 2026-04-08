# Get the duration of the contact state in Amazon Connect Agent Workspace

Returns the duration of the contact state in milliseconds relative to local time,
in the Amazon Connect agent workspace. This takes into account time skew between the JS client and the
Amazon Connect backend servers.

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
