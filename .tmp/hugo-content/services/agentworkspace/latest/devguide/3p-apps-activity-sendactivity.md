# Inform Amazon Connect that the agent is active

Sends a signal to the Amazon Connect indicating that the agent is active and should
not be logged out. It takes a provider as a parameter.

**Signature**

```typescript

sendActivity(provider): void
```

**Usage**

```typescript

import { sendActivity } from '@amazon-connect/activity';

const handleActivity = () => {
   sendActivity(sampleProvider);
};

window.addEventListener("click", handleActivity);
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onSessionExtensionError()

Agent

All content copied from https://docs.aws.amazon.com/.
