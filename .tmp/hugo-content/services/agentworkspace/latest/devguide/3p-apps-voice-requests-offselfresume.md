# Unsubscribe from self resume events in Amazon Connect Agent Workspace

Unsubscribes from self resume events.

**Signature**

```

offSelfResume(
  handler: ParticipantResumeHandler,
  contactId?: string
): void
```

**Usage**

```

voiceClient.offSelfResume(handleSelfResume);
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantResumeHandlerEvent handler function to removecontactIdstringOptional contact ID to unsubscribe from specific contact events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offSelfHold()

offVoiceEnhancementModeChanged()

All content copied from https://docs.aws.amazon.com/.
