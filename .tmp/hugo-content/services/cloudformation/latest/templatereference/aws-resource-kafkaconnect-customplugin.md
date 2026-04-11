This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::CustomPlugin

Creates a custom plugin using the specified properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KafkaConnect::CustomPlugin",
  "Properties" : {
      "ContentType" : String,
      "Description" : String,
      "Location" : CustomPluginLocation,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KafkaConnect::CustomPlugin
Properties:
  ContentType: String
  Description: String
  Location:
    CustomPluginLocation
  Name: String
  Tags:
    - Tag

```

## Properties

`ContentType`

The format of the plugin file.

_Required_: Yes

_Type_: String

_Allowed values_: `JAR | ZIP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the custom plugin.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

Information about the location of the custom plugin.

_Required_: Yes

_Type_: [CustomPluginLocation](aws-properties-kafkaconnect-customplugin-custompluginlocation.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the custom plugin.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-kafkaconnect-customplugin-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CustomPluginArn`

The Amazon Resource Name (ARN) of the custom plugin.

`Revision`

The revision of the custom plugin.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkerLogDelivery

CustomPluginFileDescription

All content copied from https://docs.aws.amazon.com/.
