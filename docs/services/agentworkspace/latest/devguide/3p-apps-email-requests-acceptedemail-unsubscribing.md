# Unsubscribe from accepted email notifications in Amazon Connect Agent Workspace

Unsubscribes a callback function from the event that is fired when an inbound
email contact is accepted.

**Signature**

```

offAcceptedEmail(handler: SubscriptionHandler<EmailContactId>, contactId?: string): void
```

**Usage**

```typescript

emailClient.offAcceptedEmail(handler);
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onAcceptedEmail()

createDraftEmail()

All content copied from https://docs.aws.amazon.com/.
