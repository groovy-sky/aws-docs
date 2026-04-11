# Subscribe to agent enabled channel list changes in Amazon Connect Agent Workspace

Creates a subscription for EnabledChannelListChanged event. This gets triggered when
an Agent's enabled channels get updated.

**Signature**

```

const handler: EnabledChannelListChangedHandler = async (data: EnabledChannelListChanged) => {
    console.log("Agent channel list change occurred! " + data);
};

agentClient.onEnabledChannelListChanged(handler);

// EnabledChannelListChanged Structure
{
    enabledChannels: AgentRoutingProfileChannelTypes[];
    previous?: {
        enabledChannels: AgentRoutingProfileChannelTypes[];
    };
}

```

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offRoutingProfileChanged()

onRoutingProfileChanged()

All content copied from https://docs.aws.amazon.com/.
