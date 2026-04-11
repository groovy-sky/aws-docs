# Unsubscribe from participant disconnected events in Amazon Connect Agent Workspace

Unsubscribes from participant disconnected events.

**Signature**

```typescript

offParticipantDisconnected(handler: ParticipantDisconnectedHandler, contactId?: string): void

```

**Usage**

```typescript

contactClient.offParticipantDisconnected(handleParticipantDisconnected);

```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantDisconnectedHandlerEvent handler function to removecontactIdstringOptional contact ID to unsubscribe from specific contact events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onParticipantDisconnected()

onParticipantStateChanged()

All content copied from https://docs.aws.amazon.com/.
