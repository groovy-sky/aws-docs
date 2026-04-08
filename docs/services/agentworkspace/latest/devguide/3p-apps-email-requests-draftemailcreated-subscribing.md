# Subscribe to draft email creation notifications in Amazon Connect Agent Workspace

Subscribes a callback function to-be-invoked whenever a draft email contact has
been created.

**Signature**

```

onDraftEmailCreated(handler: SubscriptionHandler<EmailContactId>, contactId?: string): void
```

**Usage**

```typescript

const handler: SubscriptionHandler<EmailContactId> = async (emailContact: EmailContactId) => {
   const { contactId } = emailContact;
   console.log(`Draft Email Contact Created with Id: ${contactId}`);
}

emailClient.onDraftEmailCreated(handler);

// EmailContactId Structure
{
   contactId: string;
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

createDraftEmail()

offDraftEmailCreated()

All content copied from https://docs.aws.amazon.com/.
