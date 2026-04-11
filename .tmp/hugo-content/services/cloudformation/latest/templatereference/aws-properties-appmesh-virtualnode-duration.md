This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode Duration

An object that represents a duration of time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Unit: String
  Value: Integer

```

## Properties

`Unit`

A unit of time.

_Required_: Yes

_Type_: String

_Allowed values_: `s | ms`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A number of time units.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DnsServiceDiscovery

FileAccessLog

All content copied from https://docs.aws.amazon.com/.
