# Subscribe a callback function when an Amazon Connect Agent Workspace contact starts ACW

Subscribes a callback function to-be-invoked whenever a contact StartingAcw event
occurs in the Amazon Connect agent workspace. If no contact ID is provided, then it uses the context of
the current contact that the 3P app was opened on.

**Signature**

```typescript

onStartingAcw(handler: ContactStartingAcwHandler, contactId?: string)

```

**Usage**

```typescript

const handler: ContactStartingAcwHandler = async (data: ContactStartingAcw) => {
    console.log("Contact StartingAcw occurred! " + data);
};

contactClient.onStartingAcw(handler);

// ContactStartingAcw Structure
{
  contactId: string;
}

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onParticipantStateChanged()

offStartingAcw()

All content copied from https://docs.aws.amazon.com/.
