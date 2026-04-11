This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input MediaConnectFlowRequest

Settings that apply only if the input is a MediaConnect input.

The parent of this entity is Input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowArn" : String
}

```

### YAML

```yaml

  FlowArn: String

```

## Properties

`FlowArn`

The ARN of one or two MediaConnect flows that are the sources for this MediaConnect input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputVpcRequest

MulticastSettingsCreateRequest

All content copied from https://docs.aws.amazon.com/.
