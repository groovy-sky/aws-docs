# List all contacts for the current agent in Amazon Connect Agent Workspace

Lists all contacts for the current agent.

**Signature**

```typescript

listContacts(): Promise<ListContactsResult>

```

**Usage**

```typescript

const contacts = await contactClient.listContacts();
console.log(`Active contacts: ${contacts.length}`);
contacts.forEach((contact) => {
    console.log(`Contact ${contact.contactId}: ${contact.type}`);
});

```

**Output - ListContactsResult**

Returns an array of contact data objects (currently typed as CoreContactData\[\]).

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

isPreviewMode()

listParticipants()

All content copied from https://docs.aws.amazon.com/.
