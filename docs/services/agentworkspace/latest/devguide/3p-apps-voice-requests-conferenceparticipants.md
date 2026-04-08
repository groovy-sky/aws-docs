# Conference all participants on a contact in Amazon Connect Agent Workspace

Conferences all participants on a contact together, removing any hold states and
enabling all participants to communicate with each other.

**Signature**

```

conferenceParticipants(contactId: string): Promise<void>
```

**Usage**

```

await voiceClient.conferenceParticipants("contact-123");
console.log("All participants are now conferenced");
```

**Input**

**Parameter****Type****Description**contactId _Required_stringThe unique identifier for the contact

[Document Conventions](../../../../general/latest/gr/docconventions.md)

canResumeSelf()

createOutboundCall()

All content copied from https://docs.aws.amazon.com/.
