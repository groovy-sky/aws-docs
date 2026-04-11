# Get the initial ID of the contact in Amazon Connect Agent Workspace

Returns the original (initial) contact id from which this contact was transferred
in the Amazon Connect agent workspace, or none if this is not an internal Connect transfer. This is
typically a contact owned by another agent, thus this agent will not be able to
manipulate it. It is for reference and association purposes only, and can be used to
share data between transferred contacts externally if it is linked by
originalContactId.

```typescript

async getInitialContactId(contactId: string): Promise<string | undefined>

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getContact()

getParticipant()

All content copied from https://docs.aws.amazon.com/.
