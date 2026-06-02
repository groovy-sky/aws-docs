---
title: "Unsubscribe a callback function when an Connect Customer agent workspace agent state changes - Deprecated"
---

# Unsubscribe a callback function when an Connect Customer agent workspace agent state changes - Deprecated

###### Note

This API is deprecated, use [offAvailabilityStateChanged()](3p-apps-agent-events-availabilitystatechanged-unsub.md) instead.

Unsubscribes the callback function from the agent stated change event in the Connect Customer
agent workspace.

**Signature**

```typescript

offStateChanged(handler: AgentStateChangedHandler)

```

**Usage**

```typescript

agentClient.offStateChanged(handler);

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onStateChanged() - Deprecated

getDefaultOutboundQueue()

All content copied from https://docs.aws.amazon.com/.
