This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Capability

Instantiates a capability based on the specified parameters. A trading capability contains the information required to transform incoming EDI documents into JSON or XML outputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::B2BI::Capability",
  "Properties" : {
      "Configuration" : CapabilityConfiguration,
      "InstructionsDocuments" : [ S3Location, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::B2BI::Capability
Properties:
  Configuration:
    CapabilityConfiguration
  InstructionsDocuments:
    - S3Location
  Name: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Configuration`

Specifies a structure that contains the details for a capability.

_Required_: Yes

_Type_: [CapabilityConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-capability-capabilityconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstructionsDocuments`

Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability. Each item contains the name of the bucket and the key, to identify the document's location.

_Required_: No

_Type_: Array of [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-capability-s3location.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The display name of the capability.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type. You can attach this metadata to resources (capabilities, partnerships, and so on) for any purpose.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-capability-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Returns the type of the capability. Currently, only `edi` is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `edi`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CapabilityArn`

Returns an Amazon Resource Name (ARN) for a specific AWS resource, such as a capability, partnership, profile, or transformer.

`CapabilityId`

Returns a system-assigned unique identifier for the capability.

`CreatedAt`

Returns a timestamp for creation date and time of the capability.

`ModifiedAt`

Returns a timestamp that identifies the most recent date and time that the capability
was modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS B2B Data Interchange

CapabilityConfiguration
