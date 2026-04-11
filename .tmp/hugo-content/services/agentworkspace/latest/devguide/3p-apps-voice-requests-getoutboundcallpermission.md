# Gets the outbound call permission configured for the agent in Amazon Connect Agent Workspace

Gets true if the agent has the security profile permission for making outbound
calls, false otherwise.

**Signature**

```

getOutboundCallPermission(): Promise<boolean>

```

**Usage**

```

const outboundCallPermission: boolean = await voiceClient.getOutboundCallPermission();

```

**Permissions required:**

```

User.Configuration.View

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getInitialCustomerPhoneNumber()

holdParticipant()

All content copied from https://docs.aws.amazon.com/.
