This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2 UplinkConnectionDetails

Connection details for customer to Agent and Agent to Ground Station

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentIpAndPortAddress" : RangedConnectionDetails,
  "IngressAddressAndPort" : ConnectionDetails
}

```

### YAML

```yaml

  AgentIpAndPortAddress:
    RangedConnectionDetails
  IngressAddressAndPort:
    ConnectionDetails

```

## Properties

`AgentIpAndPortAddress`

Agent IP and port address for the uplink connection.

_Required_: Yes

_Type_: [RangedConnectionDetails](aws-properties-groundstation-dataflowendpointgroupv2-rangedconnectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IngressAddressAndPort`

Ingress address and port for the uplink connection.

_Required_: Yes

_Type_: [ConnectionDetails](aws-properties-groundstation-dataflowendpointgroupv2-connectiondetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UplinkAwsGroundStationAgentEndpointDetails

UplinkDataflowDetails

All content copied from https://docs.aws.amazon.com/.
