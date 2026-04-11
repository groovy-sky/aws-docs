# Subscribe a callback function when an Amazon Connect Agent Workspace contact is missed

Subscribes a callback function to-be-invoked whenever a contact missed event
occurs in the Amazon Connect agent workspace. If no contact ID is provided, then it uses the context of
the current contact that the 3P app was opened on.

**Signature**

```typescript

onMissed(handler: ContactMissedHandler, contactId?: string)

```

**Usage**

```typescript

const handler: ContactMissedHandler = async (data: ContactMissed) => {
    console.log("Contact missed occurred! " + data);
};

contactClient.onMissed(handler);

// ContactMissed Structure
{
  contactId: string;
}

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

listParticipants()

offMissed()

All content copied from https://docs.aws.amazon.com/.
