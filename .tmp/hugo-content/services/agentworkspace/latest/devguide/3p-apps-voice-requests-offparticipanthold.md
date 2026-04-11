# Unsubscribe from participant hold events in Amazon Connect Agent Workspace

Unsubscribes from participant hold events.

**Signature**

```

offParticipantHold(
  handler: ParticipantHoldHandler,
  participantId?: string
): void
```

**Usage**

```

voiceClient.offParticipantHold(handleParticipantHold);
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantHoldHandlerEvent handler function to removeparticipantIdstringOptional participant ID to unsubscribe from specific participant
events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offCanResumeSelfChanged()

offParticipantResume()

All content copied from https://docs.aws.amazon.com/.
