# Get participant state in Amazon Connect Agent Workspace

Retrieves the current state of a specific participant.

**Signature**

```typescript

getParticipantState(participantId: string): Promise<ParticipantState>

```

**Usage**

```typescript

const state = await contactClient.getParticipantState("participant-456");
if (state.value === "connected") {
    console.log("Participant is connected");
} else if (state.value === "hold") {
    console.log("Participant is on hold");
}

```

**Input**

**Parameter****Type****Description**participantId _Required_stringThe unique identifier for the participant

**Output - ParticipantState**

The ParticipantState type can be:

- { value: ParticipantStateType } where ParticipantStateType includes:
connecting,
connected, hold, disconnected, rejected, silent\_monitor, barge

- { value: "other"; actual: string } for unknown states

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getParticipant()

getPreviewConfiguration()

All content copied from https://docs.aws.amazon.com/.
