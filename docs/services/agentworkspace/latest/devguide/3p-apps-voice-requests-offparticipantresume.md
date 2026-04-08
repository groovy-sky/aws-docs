# Unsubscribe from participant resume events in Amazon Connect Agent Workspace

Unsubscribes from participant resume events.

**Signature**

```

offParticipantResume(
  handler: ParticipantResumeHandler,
  participantId?: string
): void
```

**Usage**

```

voiceClient.offParticipantResume(handleParticipantResume);
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantResumeHandlerEvent handler function to removeparticipantIdstringOptional participant ID to unsubscribe from specific participant
events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offParticipantHold()

offSelfHold()

All content copied from https://docs.aws.amazon.com/.
