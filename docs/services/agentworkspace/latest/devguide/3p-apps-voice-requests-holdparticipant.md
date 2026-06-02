---
title: "Place a participant on hold in Connect Customer agent workspace"
---

# Place a participant on hold in Connect Customer agent workspace

Places a specific participant on hold.

**Signature**

```

holdParticipant(participantId: string): Promise<void>
```

**Usage**

```

await voiceClient.holdParticipant("participant-456");
console.log("Participant is now on hold");
```

**Input**

**Parameter****Type****Description**participantId _Required_stringThe unique identifier for the participant to place on hold

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getOutboundCallPermission()

getVoiceEnhancementMode()

All content copied from https://docs.aws.amazon.com/.
