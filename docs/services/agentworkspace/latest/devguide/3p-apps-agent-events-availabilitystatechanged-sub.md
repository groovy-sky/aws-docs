---
title: "Subscribe a callback function when an Connect Customer agent workspace agent's availability state changes"
---

# Subscribe a callback function when an Connect Customer agent workspace agent's availability state changes

Subscribes a callback function to be invoked whenever the agent's availability
state changes in the Connect Customer agent workspace.

This API supersedes [onStateChanged()](3p-apps-agent-events-statechanged-sub.md), which is now deprecated.

**Signature**

```typescript

onAvailabilityStateChanged(handler: AvailabilityStateChangedHandler)

```

**Usage**

```typescript

const handler: AvailabilityStateChangedHandler = async (data: AgentAvailabilityStateChanged) => {
    console.log("Agent availability state changed! " + data.state.name);
};

agentClient.onAvailabilityStateChanged(handler);

// AgentAvailabilityStateChanged Structure
{
  state: AgentState;
  previous?: {
    state: AgentState;
  };
}

```

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

listSecurityProfilePermissions()

offAvailabilityStateChanged()

All content copied from https://docs.aws.amazon.com/.
