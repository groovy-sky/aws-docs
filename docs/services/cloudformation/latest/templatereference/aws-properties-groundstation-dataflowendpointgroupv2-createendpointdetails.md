This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2 CreateEndpointDetails

Endpoint definition used for creating a dataflow endpoint

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DownlinkAwsGroundStationAgentEndpoint" : DownlinkAwsGroundStationAgentEndpoint,
  "UplinkAwsGroundStationAgentEndpoint" : UplinkAwsGroundStationAgentEndpoint
}

```

### YAML

```yaml

  DownlinkAwsGroundStationAgentEndpoint:
    DownlinkAwsGroundStationAgentEndpoint
  UplinkAwsGroundStationAgentEndpoint:
    UplinkAwsGroundStationAgentEndpoint

```

## Properties

`DownlinkAwsGroundStationAgentEndpoint`

Definition for a downlink agent endpoint

_Required_: No

_Type_: [DownlinkAwsGroundStationAgentEndpoint](aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UplinkAwsGroundStationAgentEndpoint`

Definition for an uplink agent endpoint

_Required_: No

_Type_: [UplinkAwsGroundStationAgentEndpoint](aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionDetails

DownlinkAwsGroundStationAgentEndpoint

All content copied from https://docs.aws.amazon.com/.
