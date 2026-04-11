This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::HookVersion LoggingConfig

The `LoggingConfig` property type specifies logging configuration
information for an extension.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String,
  "LogRoleArn" : String
}

```

### YAML

```yaml

  LogGroupName: String
  LogRoleArn: String

```

## Properties

`LogGroupName`

The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the
extension's handlers.

_Required_: No

_Type_: String

_Pattern_: `^[\.\-_/#A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogRoleArn`

The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log
entries to CloudWatch Logs.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::HookVersion

AWS::CloudFormation::LambdaHook

All content copied from https://docs.aws.amazon.com/.
