---
title: "Subscribe a callback function when an Connect Customer agent workspace contact turns to Error state"
---

# Subscribe a callback function when an Connect Customer agent workspace contact turns to Error state

Subscribes a callback function to-be-invoked whenever a contact turns to Error
state in the Connect Customer agent workspace. If no contact ID is provided, then
it uses the context of the current contact that the 3P app was opened on.

**Signature**

```typescript

onError(handler: ContactErrorHandler, contactId?: string)

```

**Usage**

```typescript

const handler: ContactErrorHandler = async (data: ContactError) => {
    console.log("Contact Error occurred! " + data.contactId);
};

contactClient.onError(handler);

// ContactError Structure
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

offConnecting()

offError()

All content copied from https://docs.aws.amazon.com/.
