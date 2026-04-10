This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup SocketAddress

The address of the endpoint, such as
`192.168.1.1`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Port" : Integer
}

```

### YAML

```yaml

  Name: String
  Port: Integer

```

## Properties

`Name`

The name of the endpoint, such as
`Endpoint 1`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port of the endpoint, such as
`55888`.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Create a SocketAddress

The following example creates a Ground Station `SocketAddress`

#### JSON

```json

{
  "Address": {
    "Name": "172.10.0.2",
    "Port": 44720
  }
}
```

#### YAML

```yaml

Address:
  Name: 172.10.0.2
  Port: 44720
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityDetails

Tag

All content copied from https://docs.aws.amazon.com/.
