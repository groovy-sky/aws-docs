# Unsubscribe from self resume capability change events in Amazon Connect Agent Workspace

Unsubscribes from capability change events for the current user.

**Signature**

```

offCanResumeSelfChanged(
  handler: CanResumeParticipantChangedHandler,
  contactId?: string
): void
```

**Usage**

```

voiceClient.offCanResumeSelfChanged(handleCanResumeSelfChanged);
```

**Input**

**Parameter****Type****Description**handler _Required_CanResumeParticipantChangedHandlerEvent handler function to removecontactIdstringOptional contact ID to unsubscribe from specific contact events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offCanResumeParticipantChanged()

offParticipantHold()

All content copied from https://docs.aws.amazon.com/.
