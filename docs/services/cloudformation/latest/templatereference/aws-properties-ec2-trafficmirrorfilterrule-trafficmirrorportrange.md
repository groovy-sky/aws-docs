This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TrafficMirrorFilterRule TrafficMirrorPortRange

Describes the Traffic Mirror port range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  ToPort: Integer

```

## Properties

`FromPort`

The start of the Traffic Mirror port range. This applies to the TCP and UDP protocols.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

The end of the Traffic Mirror port range. This applies to the TCP and UDP protocols.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::TrafficMirrorSession

All content copied from https://docs.aws.amazon.com/.
