# Get specific participant information in Amazon Connect Agent Workspace

Retrieves information for a specific participant.

**Signature**

```typescript

getParticipant(participantId: string): Promise<ParticipantData>

```

**Usage**

```typescript

const participant = await contactClient.getParticipant("participant-456");
console.log(`Participant type: ${participant.type.value}`);
console.log(`Is initial: ${participant.isInitial}`);

```

**Input**

**Parameter****Type****Description**participantId _Required_stringThe unique identifier for the participant

**Output - ParticipantData**

Returns a ParticipantData object with participant details.

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getInitialContactId()

getParticipantState()

All content copied from https://docs.aws.amazon.com/.
