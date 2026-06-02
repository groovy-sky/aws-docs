---
title: "Subscribe to participant resume events in Connect Customer agent workspace"
---

# Subscribe to participant resume events in Connect Customer agent workspace

Subscribes to events when any participant is taken off hold.

**Signature**

```

onParticipantResume(
  handler: ParticipantResumeHandler,
  participantId?: string
): void
```

**Usage**

```

const handleParticipantResume = (event) => {
  console.log(`Participant ${event.participantId} has been resumed`);
};
voiceClient.onParticipantResume(handleParticipantResume, "participant-456");
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantResumeHandlerEvent handler function to call when participants are taken off holdparticipantIdstringOptional participant ID to filter events for a specific participant

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onParticipantHold()

onSelfHold()

All content copied from https://docs.aws.amazon.com/.
