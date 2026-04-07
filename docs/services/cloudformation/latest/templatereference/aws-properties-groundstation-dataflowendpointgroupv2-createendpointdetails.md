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

_Type_: [DownlinkAwsGroundStationAgentEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UplinkAwsGroundStationAgentEndpoint`

Definition for an uplink agent endpoint

_Required_: No

_Type_: [UplinkAwsGroundStationAgentEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionDetails

DownlinkAwsGroundStationAgentEndpoint
