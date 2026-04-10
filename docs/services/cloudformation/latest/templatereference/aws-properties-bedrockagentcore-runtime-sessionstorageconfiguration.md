This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime SessionStorageConfiguration

The `SessionStorageConfiguration` property type specifies Property description not available. for an [AWS::BedrockAgentCore::Runtime](aws-resource-bedrockagentcore-runtime.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MountPath" : String
}

```

### YAML

```yaml

  MountPath: String

```

## Properties

`MountPath`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^/mnt/[a-zA-Z0-9._-]+/?$`

_Minimum_: `6`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Location

VpcConfig

All content copied from https://docs.aws.amazon.com/.
