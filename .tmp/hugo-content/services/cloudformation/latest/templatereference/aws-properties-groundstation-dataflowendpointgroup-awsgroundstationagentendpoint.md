This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup AwsGroundStationAgentEndpoint

Information about AwsGroundStationAgentEndpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentStatus" : String,
  "AuditResults" : String,
  "EgressAddress" : ConnectionDetails,
  "IngressAddress" : RangedConnectionDetails,
  "Name" : String
}

```

### YAML

```yaml

  AgentStatus: String
  AuditResults: String
  EgressAddress:
    ConnectionDetails
  IngressAddress:
    RangedConnectionDetails
  Name: String

```

## Properties

`AgentStatus`

The status of AgentEndpoint.

_Required_: No

_Type_: String

_Allowed values_: `SUCCESS | FAILED | ACTIVE | INACTIVE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuditResults`

The results of the audit.

_Required_: No

_Type_: String

_Allowed values_: `HEALTHY | UNHEALTHY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EgressAddress`

The egress address of AgentEndpoint.

_Required_: No

_Type_: [ConnectionDetails](aws-properties-groundstation-dataflowendpointgroup-connectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IngressAddress`

The ingress address of AgentEndpoint.

_Required_: No

_Type_: [RangedConnectionDetails](aws-properties-groundstation-dataflowendpointgroup-rangedconnectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Name string associated with AgentEndpoint. Used as a human-readable identifier for AgentEndpoint.

_Required_: No

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GroundStation::DataflowEndpointGroup

ConnectionDetails

All content copied from https://docs.aws.amazon.com/.
