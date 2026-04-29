---
title: "Unsubscribe from participant resume capability change events in Amazon Connect Customer AI agent workspace"
---

# Unsubscribe from participant resume capability change events in Amazon Connect Customer AI agent workspace

Unsubscribes from participant capability change events.

**Signature**

```

offCanResumeParticipantChanged(
  handler: CanResumeParticipantChangedHandler,
  participantId?: string
): void
```

**Usage**

```

voiceClient.offCanResumeParticipantChanged(handleCanResumeChanged);
```

**Input**

**Parameter****Type****Description**handler _Required_CanResumeParticipantChangedHandlerEvent handler function to removeparticipantIdstringOptional participant ID to unsubscribe from specific participant
events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

listDialableCountries()

offCanResumeSelfChanged()

All content copied from https://docs.aws.amazon.com/.
