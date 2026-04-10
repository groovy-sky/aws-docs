This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingGroup LogConfiguration

Sets a custom Amazon CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.

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

Sets a custom Amazon CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.

_Required_: No

_Type_: String

_Pattern_: `\A\/aws\/MediaPackage\/[0-9a-zA-Z-_\/\.#]+\Z`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Authorization

Tag

All content copied from https://docs.aws.amazon.com/.
