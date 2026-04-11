# Unsubscribe from draft email creation notifications in Amazon Connect Agent Workspace

Unsubscribes a callback function from the event that is fired when a draft email
contact is created.

**Signature**

```

offDraftEmailCreated(handler: SubscriptionHandler<EmailContactId>, contactId?: string): void
```

**Usage**

```typescript

emailClient.offDraftEmailCreated(handler);
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onDraftEmailCreated()

getEmailData()

All content copied from https://docs.aws.amazon.com/.
