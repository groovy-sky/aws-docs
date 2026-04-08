# Get the limit of contacts for the agent in Amazon Connect Agent Workspace

Returns a map of `ChannelType`-to-number indicating how many concurrent
contacts can an Amazon Connect agent workspace agent have on a given channel. 0 represents a disabled
channel.

```typescript

async getChannelConcurrency(): Promise<AgentChannelConcurrencyMap>

```

**Permissions required:**

```typescript

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getARN()

getExtension()

All content copied from https://docs.aws.amazon.com/.
