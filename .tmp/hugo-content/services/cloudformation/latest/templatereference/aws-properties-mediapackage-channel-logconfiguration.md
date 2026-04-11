This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::Channel LogConfiguration

The access log configuration parameters for your channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String
}

```

### YAML

```yaml

  LogGroupName: String

```

## Properties

`LogGroupName`

Sets a custom Amazon CloudWatch log group name.

_Required_: No

_Type_: String

_Pattern_: `\A^(\/aws\/MediaPackage\/)[a-zA-Z0-9_-]+\Z`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngestEndpoint

Tag

All content copied from https://docs.aws.amazon.com/.
