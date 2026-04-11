This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction EnableIoTLoggingParams

Parameters used when defining a mitigation action that enable AWS IoT Core logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogLevel" : String,
  "RoleArnForLogging" : String
}

```

### YAML

```yaml

  LogLevel: String
  RoleArnForLogging: String

```

## Properties

`LogLevel`

Specifies the type of information to be logged.

_Required_: Yes

_Type_: String

_Allowed values_: `DEBUG | INFO | ERROR | WARN | UNSET_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArnForLogging`

The Amazon Resource Name (ARN) of the IAM role used for logging.

_Required_: Yes

_Type_: String

_Minimum_: `11`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddThingsToThingGroupParams

PublishFindingToSnsParams

All content copied from https://docs.aws.amazon.com/.
