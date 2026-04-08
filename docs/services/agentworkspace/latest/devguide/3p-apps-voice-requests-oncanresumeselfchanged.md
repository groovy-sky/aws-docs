# Subscribe to self resume capability change events in Amazon Connect Agent Workspace

Subscribes to events when the current user's capability to be resumed from hold
changes.

**Signature**

```

onCanResumeSelfChanged(
  handler: CanResumeParticipantChangedHandler,
  contactId?: string
): void
```

**Usage**

```

const handleCanResumeSelfChanged = (event) => {
  if (event.canResumeConnection) {
    console.log("You can now be resumed from hold");
    // Enable resume button in UI
  } else {
    console.log("You cannot be resumed at this time");
    // Disable resume button in UI
  }
};
voiceClient.onCanResumeSelfChanged(handleCanResumeSelfChanged, "contact-123");
```

**Input**

**Parameter****Type****Description**handler _Required_CanResumeParticipantChangedHandlerEvent handler function to call when the capability changescontactIdstringOptional contact ID to filter events for a specific contact

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onCanResumeParticipantChanged()

onParticipantHold()

All content copied from https://docs.aws.amazon.com/.
