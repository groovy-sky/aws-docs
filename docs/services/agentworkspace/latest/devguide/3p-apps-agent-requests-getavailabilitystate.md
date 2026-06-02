---
title: "Get the current availability state of the agent in Connect Customer agent workspace"
---

# Get the current availability state of the agent in Connect Customer agent workspace

Returns the current availability state of the agent currently logged in to the
Connect Customer agent workspace, along with the next pending state if one is queued
because the agent is handling an active contact.

This API supersedes [getState()](3p-apps-agent-requests-getstate.md), which is now deprecated.

```typescript

async getAvailabilityState(): Promise<GetAvailabilityStateResult>

```

**Output - GetAvailabilityStateResult**

**Parameter****Type****Description**agentStateARNstring (optional)The ARN of the agent's current availability state.namestringThe name of the agent's current availability state.typeAgentStateTypeThe agent's current availability state type. One of
`routable`, `not_routable`,
`after_call_work`, `system`,
`error`, or `offline`.startTimestampDate (optional)The time at which the agent entered this state.nextStateAgentState (optional)The next state the agent will transition to once all active
contacts are cleared, when one is queued.

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getRoutingProfileQueues()

listSecurityProfilePermissions()

All content copied from https://docs.aws.amazon.com/.
