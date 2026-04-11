This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service IngressConfiguration

Network configuration settings for inbound network traffic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsPubliclyAccessible" : Boolean
}

```

### YAML

```yaml

  IsPubliclyAccessible: Boolean

```

## Properties

`IsPubliclyAccessible`

Specifies whether your App Runner service is publicly accessible. To make the service publicly accessible set it to `True`. To make the service
privately accessible, from only within an Amazon VPC set it to `False`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageRepository

InstanceConfiguration

All content copied from https://docs.aws.amazon.com/.
