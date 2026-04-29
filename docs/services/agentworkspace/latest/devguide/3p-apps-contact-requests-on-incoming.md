---
title: "Subscribe to incoming contact events in Amazon Connect Customer AI agent workspace"
---

# Subscribe to incoming contact events in Amazon Connect Customer AI agent workspace

Creates a subscription whenever a new incoming event occurs in the Amazon Connect
Customer AI agent workspace.

**Signature**

```

onIncoming(handler: ContactIncomingHandler, contactId?: string): void

```

**Usage**

```

const handler: ContactIncomingHandler = async (data: ContactIncoming) => {
    console.log("Contact incoming occurred! " + data);
};

contactClient.onIncoming(handler);

// ContactIncoming Structure
{
    contactId: string;
    initialContactId: string | undefined;
    type: ContactChannelType["type"];
    subtype: ContactChannelType["subtype"];
}

```

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offIncoming()

onParticipantAdded()

All content copied from https://docs.aws.amazon.com/.
