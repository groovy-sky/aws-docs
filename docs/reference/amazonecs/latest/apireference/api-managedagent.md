# ManagedAgent

Details about the managed agent status for the container.

## Contents

**lastStartedAt**

The Unix timestamp for the time when the managed agent was last started.

Type: Timestamp

Required: No

**lastStatus**

The last known status of the managed agent.

Type: String

Required: No

**name**

The name of the managed agent. When the execute command feature is turned on, the
managed agent name is `ExecuteCommandAgent`.

Type: String

Valid Values: `ExecuteCommandAgent`

Required: No

**reason**

The reason for why the managed agent is in the state it is in.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/managedagent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/managedagent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/managedagent.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LogConfiguration

ManagedAgentStateChange
