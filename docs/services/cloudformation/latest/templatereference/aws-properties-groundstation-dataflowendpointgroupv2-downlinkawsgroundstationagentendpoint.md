This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2 DownlinkAwsGroundStationAgentEndpoint

Definition for a downlink agent endpoint

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataflowDetails" : DownlinkDataflowDetails,
  "Name" : String
}

```

### YAML

```yaml

  DataflowDetails:
    DownlinkDataflowDetails
  Name: String

```

## Properties

`DataflowDetails`

Dataflow details for the downlink endpoint

_Required_: Yes

_Type_: [DownlinkDataflowDetails](aws-properties-groundstation-dataflowendpointgroupv2-downlinkdataflowdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Downlink dataflow endpoint name

_Required_: Yes

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateEndpointDetails

DownlinkAwsGroundStationAgentEndpointDetails

All content copied from https://docs.aws.amazon.com/.
