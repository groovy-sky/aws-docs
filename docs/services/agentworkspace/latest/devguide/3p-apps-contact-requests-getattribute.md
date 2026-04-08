# Get specific attributes for a contact in Amazon Connect Agent Workspace

Returns the requested attribute associated with the contact in the Amazon Connect
agent workspace.

```typescript

async getAttribute(
  contactId: string,
  attribute: string,
): Promise<string | undefined>

```

**Permissions required:**

```typescript

Contact.Attributes.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

engagePreviewContact()

getAttributes()

All content copied from https://docs.aws.amazon.com/.
