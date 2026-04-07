This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppIntegrations::DataIntegration

Creates and persists a DataIntegration resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppIntegrations::DataIntegration",
  "Properties" : {
      "Description" : String,
      "FileConfiguration" : FileConfiguration,
      "KmsKey" : String,
      "Name" : String,
      "ObjectConfiguration" : {Key: Value, ...},
      "ScheduleConfig" : ScheduleConfig,
      "SourceURI" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppIntegrations::DataIntegration
Properties:
  Description: String
  FileConfiguration:
    FileConfiguration
  KmsKey: String
  Name: String
  ObjectConfiguration:
    Key: Value
  ScheduleConfig:
    ScheduleConfig
  SourceURI: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the DataIntegration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileConfiguration`

The configuration for what files should be pulled from the source.

_Required_: No

_Type_: [FileConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appintegrations-dataintegration-fileconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

The KMS key for the DataIntegration.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the DataIntegration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9/\._\-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectConfiguration`

The configuration for what data should be pulled from the source.

_Required_: No

_Type_: Object of Object

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleConfig`

The name of the data and how often it should be pulled from the source.

_Required_: No

_Type_: [ScheduleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appintegrations-dataintegration-scheduleconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceURI`

The URI of the data source.

_Required_: Yes

_Type_: String

_Pattern_: `^(\w+\:\/\/[\w.-]+[\w/!@#+=.-]+$)|(\w+\:\/\/[\w.-]+[\w/!@#+=.-]+[\w/!@#+=.-]+[\w/!@#+=.,-]+$)`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appintegrations-dataintegration-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DataIntegration name. For example:

`{ "Ref": "myDataIntegrationName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DataIntegrationArn`

The Amazon Resource Name (ARN) for the DataIntegration.

`Id`

A unique identifier.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

FileConfiguration
