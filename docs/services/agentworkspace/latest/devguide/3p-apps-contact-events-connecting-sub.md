---
title: "Subscribe a callback function when an Connect Customer agent workspace contact turns to Connecting state"
---

# Subscribe a callback function when an Connect Customer agent workspace contact turns to Connecting state

Subscribes a callback function to-be-invoked whenever a contact turns to
Connecting state in the Connect Customer agent workspace. The Connecting state
means the contact is being routed to the agent and has not yet been fully
assigned. If no contact ID is provided, then it uses the context of the current
contact that the 3P app was opened on.

**Signature**

```typescript

onConnecting(handler: ContactConnectingHandler, contactId?: string)

```

**Usage**

```typescript

const handler: ContactConnectingHandler = async (data: ContactConnecting) => {
    console.log("Contact Connecting occurred! " + data.contactId);
};

contactClient.onConnecting(handler);

// ContactConnecting Structure
{
    contactId: string;
    initialContactId: string | undefined;
    type: string;
    subtype: string;
}

```

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

isAutoAcceptEnabled()

offConnecting()

All content copied from https://docs.aws.amazon.com/.
