# Unsubscribe from participant added events in Amazon Connect Agent Workspace

Unsubscribes from participant added events.

**Signature**

```typescript

offParticipantAdded(handler: ParticipantAddedHandler, contactId?: string): void

```

**Usage**

```typescript

contactClient.offParticipantAdded(handleParticipantAdded);

```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantAddedHandlerEvent handler function to removecontactIdstringOptional contact ID to unsubscribe from specific contact events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onParticipantAdded()

onParticipantDisconnected()

All content copied from https://docs.aws.amazon.com/.
