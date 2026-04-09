# AgentInfo

Information about agents associated with the user’s AWS account. Information includes
agent IDs, IP addresses, media access control (MAC) addresses, agent or collector status,
hostname where the agent resides, and agent version for each agent.

## Contents

**agentId**

The agent or collector ID.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 20.

Pattern: `\S+`

Required: No

**agentNetworkInfoList**

Network details about the host where the agent or collector resides.

Type: Array of [AgentNetworkInfo](api-agentnetworkinfo.md) objects

Required: No

**agentType**

Type of agent.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**collectionStatus**

Status of the collection process for an agent.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**connectorId**

The ID of the connector.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**health**

The health of the agent.

Type: String

Valid Values: `HEALTHY | UNHEALTHY | RUNNING | UNKNOWN | BLACKLISTED | SHUTDOWN`

Required: No

**hostName**

The name of the host where the agent or collector resides. The host can be a server or
virtual machine.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**lastHealthPingTime**

Time since agent health was reported.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**registeredTime**

Agent's first registration timestamp in UTC.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

**version**

The agent or collector version.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/agentinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/agentinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/agentinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentConfigurationStatus

AgentNetworkInfo

All content copied from https://docs.aws.amazon.com/.
