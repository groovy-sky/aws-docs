# Unsubscribe a callback function when an Amazon Connect Agent Workspace contact starts ACW

Unsubscribes the callback function from the contact StartingAcw event in the Amazon Connect
agent workspace.

**Signature**

```typescript

offStartingAcw(handler: ContactStartingAcwHandler, contactId?: string)

```

**Usage**

```typescript

contactClient.offStartingAcw(handler);

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onStartingAcw()

transfer()

All content copied from https://docs.aws.amazon.com/.
