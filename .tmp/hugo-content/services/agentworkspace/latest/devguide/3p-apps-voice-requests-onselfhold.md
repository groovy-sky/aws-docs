# Subscribe to self hold events in Amazon Connect Agent Workspace

Subscribes to events when the current user's participant is put on hold.

**Signature**

```

onSelfHold(
  handler: ParticipantHoldHandler,
  contactId?: string
): void
```

**Usage**

```

const handleSelfHold = (event) => {
  console.log("You have been put on hold");
  console.log(`Contact: ${event.contactId}`);
};
// Subscribe to all contacts
voiceClient.onSelfHold(handleSelfHold);
// Or subscribe to a specific contact
voiceClient.onSelfHold(handleSelfHold, "contact-123");
```

**Input**

**Parameter****Type****Description**handler _Required_ParticipantHoldHandlerEvent handler function to call when the current user's participant is
put on holdcontactIdstringOptional contact ID to filter events for a specific contact

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onParticipantResume()

onSelfResume()

All content copied from https://docs.aws.amazon.com/.
