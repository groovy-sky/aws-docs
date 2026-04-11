# Subscribe to agent routing profile changes in Amazon Connect Agent Workspace

Creates a subscription for RoutingProfileChanged event. This gets triggered when an
Agent's routing profile gets updated.

**Signature**

```

const handler: RoutingProfileChangedHandler = async (data: AgentRoutingProfileChanged) => {
    console.log("Agent routing profile change occurred! " + data);
};

agentClient.onRoutingProfileChanged(handler);

// AgentRoutingProfileChanged Structure
{
    routingProfile: AgentRoutingProfile;
    previous?: {
        routingProfile: AgentRoutingProfile;
    };
}

```

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onEnabledChannelListChanged()

setAvailabilityState()

All content copied from https://docs.aws.amazon.com/.
