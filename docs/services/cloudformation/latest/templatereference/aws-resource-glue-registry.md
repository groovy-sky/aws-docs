This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Registry

The AWS::Glue::Registry is an AWS Glue resource type that manages registries of schemas in the AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Registry",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Registry
Properties:
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the registry.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the registry.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

AWS tags that contain a key value pair and may be searched by console, command line, or API.

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-registry-tag.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a combination of "VersionId\|Key\|Value" as a string.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageDescriptor

Tag

All content copied from https://docs.aws.amazon.com/.
