This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Command

The `AWS::IoT::Command` resource Property description not available. for IoT.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Command",
  "Properties" : {
      "CommandId" : String,
      "CreatedAt" : String,
      "Deprecated" : Boolean,
      "Description" : String,
      "DisplayName" : String,
      "LastUpdatedAt" : String,
      "MandatoryParameters" : [ CommandParameter, ... ],
      "Namespace" : String,
      "Payload" : CommandPayload,
      "PayloadTemplate" : String,
      "PendingDeletion" : Boolean,
      "Preprocessor" : CommandPreprocessor,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Command
Properties:
  CommandId: String
  CreatedAt: String
  Deprecated: Boolean
  Description: String
  DisplayName: String
  LastUpdatedAt: String
  MandatoryParameters:
    - CommandParameter
  Namespace: String
  Payload:
    CommandPayload
  PayloadTemplate: String
  PendingDeletion: Boolean
  Preprocessor:
    CommandPreprocessor
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`CommandId`

The unique identifier of the command.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreatedAt`

The timestamp, when the command was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Deprecated`

Indicates whether the command has been deprecated.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the command parameter.

_Required_: No

_Type_: String

_Maximum_: `2028`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the command.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUpdatedAt`

The timestamp, when the command was last updated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MandatoryParameters`

Property description not available.

_Required_: No

_Type_: Array of [CommandParameter](aws-properties-iot-command-commandparameter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `AWS-IoT | AWS-IoT-FleetWise`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

Property description not available.

_Required_: No

_Type_: [CommandPayload](aws-properties-iot-command-commandpayload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadTemplate`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `32768`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PendingDeletion`

Indicates whether the command is pending deletion.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Preprocessor`

Property description not available.

_Required_: No

_Type_: [CommandPreprocessor](aws-properties-iot-command-commandpreprocessor.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2028`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-command-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CommandArn`

The Amazon Resource Name (ARN) of the command.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AwsJsonSubstitutionCommandPreprocessorConfig

All content copied from https://docs.aws.amazon.com/.
