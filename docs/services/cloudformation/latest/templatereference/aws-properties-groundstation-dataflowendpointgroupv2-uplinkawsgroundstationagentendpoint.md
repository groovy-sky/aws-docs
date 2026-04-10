This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2 UplinkAwsGroundStationAgentEndpoint

Definition for an uplink agent endpoint

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataflowDetails" : UplinkDataflowDetails,
  "Name" : String
}

```

### YAML

```yaml

  DataflowDetails:
    UplinkDataflowDetails
  Name: String

```

## Properties

`DataflowDetails`

Dataflow details for the uplink endpoint

_Required_: Yes

_Type_: [UplinkDataflowDetails](aws-properties-groundstation-dataflowendpointgroupv2-uplinkdataflowdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Uplink dataflow endpoint name

_Required_: Yes

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UplinkAwsGroundStationAgentEndpointDetails

All content copied from https://docs.aws.amazon.com/.
