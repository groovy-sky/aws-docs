This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MultiRegionEndpoint

Creates a multi-region endpoint (global-endpoint).

The primary region is going to be the AWS-Region where the operation is executed.
The secondary region has to be provided in request's parameters.
From the data flow standpoint there is no difference between primary
and secondary regions - sending traffic will be split equally between the two.
The primary region is the region where the resource has been created and where it can be managed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MultiRegionEndpoint",
  "Properties" : {
      "Details" : Details,
      "EndpointName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MultiRegionEndpoint
Properties:
  Details:
    Details
  EndpointName: String
  Tags:
    - Tag

```

## Properties

`Details`

Contains details of a multi-region endpoint (global-endpoint) being created.

_Required_: Yes

_Type_: [Details](aws-properties-ses-multiregionendpoint-details.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointName`

The name of the multi-region endpoint (global-endpoint).

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of objects that define the tags (keys and values) to associate with the multi-region endpoint (global-endpoint).

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-multiregionendpoint-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Details

All content copied from https://docs.aws.amazon.com/.
