This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::Channel OutputHeaderConfiguration

The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PublishMQCS" : Boolean
}

```

### YAML

```yaml

  PublishMQCS: Boolean

```

## Properties

`PublishMQCS`

When true, AWS Elemental MediaPackage includes the MQCS in responses to the CDN. This setting is valid only when `InputType` is `CMAF`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputSwitchConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
