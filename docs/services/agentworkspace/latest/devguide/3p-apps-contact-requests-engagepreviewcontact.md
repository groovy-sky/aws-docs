# Engage the preview contact for the given contactId in Amazon Connect Agent Workspace

When an agent is previewing a preview contact, this API will actually initiate the
outbound dial to the end customer, ending the preview experience.

**Signature**

```

engagePreviewContact(contactId: string): Promise<AddParticipantResult>

```

**Usage**

```

const addParticipantResult: AddParticipantResult = await contactClient.engagePreviewContact(contactId);

```

**Input**

**Parameter****Type****Description** contactId _Required_ string  The id of the contact which is being previewed by the agent to
which a participant needs to be added.

**Output - AddParticipantResult**

**Parameter****Type****Description** participantId  string  The id of the newly added participant

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

disconnectParticipant()

getAttribute()

All content copied from https://docs.aws.amazon.com/.
