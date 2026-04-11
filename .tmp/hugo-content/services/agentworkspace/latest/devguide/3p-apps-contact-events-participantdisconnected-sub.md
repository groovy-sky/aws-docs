# Subscribe to participant disconnected events in Amazon Connect Agent Workspace

Subscribes to participant disconnected events. This event fires when a participant
leaves
or is removed from a contact.

**Signature**

```typescript

onParticipantDisconnected(handler: ParticipantDisconnectedHandler, contactId?: string): void

```

**Usage**

```typescript

const handleParticipantDisconnected = (event) => {
    console.log(`Participant disconnected: ${event.participant.participantId}`);
};

contactClient.onParticipantDisconnected(
    handleParticipantDisconnected,
    "contact-123"
);

```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantDisconnectedHandlerEvent handler function to call when participants disconnectcontactIdstringOptional contact ID to filter events for a specific contact

**Event Structure - ParticipantDisconnected**

The handler receives a ParticipantDisconnected event with:

- `participant`: ParticipantData - The participant that was
disconnected

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offParticipantAdded()

offParticipantDisconnected()

All content copied from https://docs.aws.amazon.com/.
