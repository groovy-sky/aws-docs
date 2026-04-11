This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application PropertyGroup

Property key-value pairs passed into an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PropertyGroupId" : String,
  "PropertyMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  PropertyGroupId: String
  PropertyMap:
    Key: Value

```

## Properties

`PropertyGroupId`

Describes the key of an application execution property key-value pair.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyMap`

Describes the value of an application execution property key-value pair.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PropertyGroup](../../../managed-flink/latest/apiv2/api-propertygroup.md) in the _Amazon Kinesis Data Analytics_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParallelismConfiguration

RecordColumn

All content copied from https://docs.aws.amazon.com/.
