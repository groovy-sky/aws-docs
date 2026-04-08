# Subscribe a callback function when an Amazon Connect Agent Workspace contact is connected

Subscribes a callback function to-be-invoked whenever a contact Connected event
occurs in the Amazon Connect agent workspace. If no contact ID is provided, then it uses the context of
the current contact that the 3P app was opened on.

**Signature**

```typescript

onConnected(handler: ContactConnectedHandler, contactId?: string)

```

**Usage**

```typescript

const handler: ContactConnectedHandler = async (data: ContactConnected) => {
    console.log("Contact Connected occurred! " + data);
};

contactClient.onConnected(handler);

// ContactConnected Structure
{
    contactId: string;
}

```

**Permissions required:**

```typescript

Contact.Details.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offCleared()

offConnected()

All content copied from https://docs.aws.amazon.com/.
