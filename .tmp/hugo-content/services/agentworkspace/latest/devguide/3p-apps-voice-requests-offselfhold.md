# Unsubscribe from self hold events in Amazon Connect Agent Workspace

Unsubscribes from self hold events.

**Signature**

```

offSelfHold(
  handler: ParticipantHoldHandler,
  contactId?: string
): void
```

**Usage**

```

voiceClient.offSelfHold(handleSelfHold);
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantHoldHandlerEvent handler function to removecontactIdstringOptional contact ID to unsubscribe from specific contact events

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offParticipantResume()

offSelfResume()

All content copied from https://docs.aws.amazon.com/.
