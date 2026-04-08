# Subscribe to participant resume capability change events in Amazon Connect Agent Workspace

Subscribes to events when a participant's capability to be resumed from hold changes.

**Signature**

```

onCanResumeParticipantChanged(
  handler: CanResumeParticipantChangedHandler,
  participantId?: string
): void
```

**Usage**

```

const handleCanResumeChanged = (event) => {
  console.log(`Participant ${event.participantId} resume capability: ${event.canResumeConnection}`);
  if (event.canResumeConnection) {
    // Enable resume button for this participant
  } else {
    // Disable resume button for this participant
  }
};
voiceClient.onCanResumeParticipantChanged(handleCanResumeChanged, "participant-456");
```

**Input**

**Parameter****Type****Description**handler _Required_CanResumeParticipantChangedHandlerEvent handler function to call when the capability changesparticipantIdstringOptional participant ID to filter events for a specific participant

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offVoiceEnhancementModeChanged()

onCanResumeSelfChanged()

All content copied from https://docs.aws.amazon.com/.
