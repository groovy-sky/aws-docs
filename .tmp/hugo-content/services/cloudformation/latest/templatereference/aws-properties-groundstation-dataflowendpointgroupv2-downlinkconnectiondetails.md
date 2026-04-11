This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2 DownlinkConnectionDetails

Connection details for Ground Station to Agent and Agent to customer

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentIpAndPortAddress" : RangedConnectionDetails,
  "EgressAddressAndPort" : ConnectionDetails
}

```

### YAML

```yaml

  AgentIpAndPortAddress:
    RangedConnectionDetails
  EgressAddressAndPort:
    ConnectionDetails

```

## Properties

`AgentIpAndPortAddress`

Agent IP and port address for the downlink connection.

_Required_: Yes

_Type_: [RangedConnectionDetails](aws-properties-groundstation-dataflowendpointgroupv2-rangedconnectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EgressAddressAndPort`

Egress address and port for the downlink connection.

_Required_: Yes

_Type_: [ConnectionDetails](aws-properties-groundstation-dataflowendpointgroupv2-connectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DownlinkAwsGroundStationAgentEndpointDetails

DownlinkDataflowDetails

All content copied from https://docs.aws.amazon.com/.
