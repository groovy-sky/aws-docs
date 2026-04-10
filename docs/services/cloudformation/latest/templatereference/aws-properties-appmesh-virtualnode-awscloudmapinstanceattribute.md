This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode AwsCloudMapInstanceAttribute

An object that represents the AWS Cloud Map attribute information for your
virtual node.

###### Note

AWS Cloud Map is not available in the eu-south-1 Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The name of an AWS Cloud Map service instance attribute key. Any AWS Cloud Map service instance that contains the specified key and value is
returned.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9!-~]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of an AWS Cloud Map service instance attribute key. Any AWS Cloud Map service instance that contains the specified key and value is
returned.

_Required_: Yes

_Type_: String

_Pattern_: `([a-zA-Z0-9!-~][  a-zA-Z0-9!-~]*){0,1}[a-zA-Z0-9!-~]{0,1}`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessLog

AwsCloudMapServiceDiscovery

All content copied from https://docs.aws.amazon.com/.
