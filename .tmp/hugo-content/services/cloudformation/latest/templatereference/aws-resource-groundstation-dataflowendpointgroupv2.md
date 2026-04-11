This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroupV2

Creates a `DataflowEndpoint` group containing the specified list of Ground Station Agent based endpoints.

The `name` field in each endpoint is used in your mission profile `DataflowEndpointConfig`
to specify which endpoints to use during a contact.

When a contact uses multiple `DataflowEndpointConfig` objects, each `Config`
must match a `DataflowEndpoint` in the same group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GroundStation::DataflowEndpointGroupV2",
  "Properties" : {
      "ContactPostPassDurationSeconds" : Integer,
      "ContactPrePassDurationSeconds" : Integer,
      "Endpoints" : [ CreateEndpointDetails, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GroundStation::DataflowEndpointGroupV2
Properties:
  ContactPostPassDurationSeconds: Integer
  ContactPrePassDurationSeconds: Integer
  Endpoints:
    - CreateEndpointDetails
  Tags:
    - Tag

```

## Properties

`ContactPostPassDurationSeconds`

Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `480`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContactPrePassDurationSeconds`

Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `480`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Endpoints`

List of endpoints for the dataflow endpoint group.

_Required_: No

_Type_: Array of [CreateEndpointDetails](aws-properties-groundstation-dataflowendpointgroupv2-createendpointdetails.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags assigned to a resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-groundstation-dataflowendpointgroupv2-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

`EndpointDetails`

Information about the endpoint details.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ConnectionDetails

All content copied from https://docs.aws.amazon.com/.
