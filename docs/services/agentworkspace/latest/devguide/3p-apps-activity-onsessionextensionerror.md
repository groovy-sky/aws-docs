# Subscribe to session extension errors in Amazon Connect Agent Workspace

Subscribes a callback function to be invoked when an attempt to extend the agent's
session fails.

**Signature**

```typescript

onSessionExtensionError(handler: SessionExtensionErrorHandler);
```

**Usage**

```typescript

const handler: SessionExtensionErrorHandler = (details: SessionExtensionErrorData) => {
    console.log("Failed to extend my session!", details);
}

sessionExpirationWarningClient.onSessionExtensionError(handler);

// SessionExtensionErrorData Structure
{
    isWarningActive: boolean;
    errorDetails: Record<string, unknown>;
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onExpirationWarningCleared()

sendActivity()

All content copied from https://docs.aws.amazon.com/.
