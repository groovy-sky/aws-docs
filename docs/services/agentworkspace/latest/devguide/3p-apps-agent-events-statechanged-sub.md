# Subscribe a callback function when an Amazon Connect Agent Workspace agent state changes

Subscribes a callback function to-be-invoked whenever an agent state changed event
occurs in the Amazon Connect agent workspace.

**Signature**

```typescript

onStateChanged(handler: AgentStateChangedHandler)

```

**Usage**

```typescript

const handler: AgentStateChangedHandler = async (data: AgentStateChangedEventData) => {
    console.log("Agent state change occurred! " + data);
};

agentClient.onStateChanged(handler);

// AgentStateChangedEventData Structure
{
  state: string;
  previous: {
    state: string;
  };
}

```

**Permissions required:**

```typescript

User.Status.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

setOffline()

offStateChanged()

All content copied from https://docs.aws.amazon.com/.
