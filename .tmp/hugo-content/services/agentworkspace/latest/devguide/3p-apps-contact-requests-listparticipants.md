# List all participants for a contact in Amazon Connect Agent Workspace

Retrieves all participants associated with a specific contact.

**Signature**

```typescript

listParticipants(contactId: string): Promise<ParticipantData[]>

```

**Usage**

```typescript

const participants = await contactClient.listParticipants("contact-123");
participants.forEach((p) => {
    console.log(`Participant ${p.participantId}: ${p.type.value}`);
    if (p.isSelf) {
        console.log("This is the current user");
    }
});

```

**Input**

**Parameter****Type****Description**contactId _Required_stringThe unique identifier for the contact

**Output - ParticipantData\[\]**

The ParticipantData interface includes:

- `participantId`: string - Unique identifier for the participant

- `contactId`: string - Contact this participant belongs to

- `type`: ParticipantType - Type of participant (agent, outbound,
inbound, monitoring, other)

- `isInitial`: boolean - Whether this is the initial participant

- `isSelf`: boolean - Whether this participant is associated with the
current user

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

listContacts()

onMissed()

All content copied from https://docs.aws.amazon.com/.
