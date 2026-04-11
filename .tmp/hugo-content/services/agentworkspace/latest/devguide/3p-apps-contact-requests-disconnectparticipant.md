# Disconnect a participant from a contact in Amazon Connect Agent Workspace

Disconnects a specific participant from the contact.

**Signature**

```typescript

disconnectParticipant(participantId: string): Promise<void>

```

**Usage**

```typescript

await contactClient.disconnectParticipant("participant-456");
console.log("Participant disconnected");

```

**Input**

**Parameter****Type****Description**participantId _Required_stringThe unique identifier for the participant to disconnect

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offConnected()

engagePreviewContact()

All content copied from https://docs.aws.amazon.com/.
