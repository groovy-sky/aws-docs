# Get all the availability states configured for the current agent in Amazon Connect Agent Workspace

Get all the availability states configured for the current agent.

**Signature**

```

listAvailabilityStates(): Promise<AgentState[]>

```

**Usage**

```

const availabilityStates: AgentState[] = await agentClient.listAvailabilityStates();

```

**Output - AgentState**

**Parameter****Type****Description** agentStateARN  string  Amazon Reference Number of agent state  type  string  It could be "routable" \| "not\_routable" \|
"after\_call\_work" \| "system" \| "error"
\| "offline"  name  string  Name of the agent state like `Available` or `Offline` startTimestamp  Date  A `Date` object that indicates when the state was set.

**Permissions required:**

```

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getState()

listQuickConnects()

All content copied from https://docs.aws.amazon.com/.
