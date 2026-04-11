This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup DataflowEndpoint

Contains information such as socket address and name that defines an endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : SocketAddress,
  "Mtu" : Integer,
  "Name" : String
}

```

### YAML

```yaml

  Address:
    SocketAddress
  Mtu: Integer
  Name: String

```

## Properties

`Address`

The address and port of an endpoint.

_Required_: No

_Type_: [SocketAddress](aws-properties-groundstation-dataflowendpointgroup-socketaddress.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mtu`

Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
Valid values are between 1400 and 1500. A default value of 1500 is used if not set.

_Required_: No

_Type_: Integer

_Minimum_: `1400`

_Maximum_: `1500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The endpoint name.

When listing available contacts for a satellite, Ground Station searches for a dataflow endpoint whose name matches the value specified by the dataflow endpoint config of the selected mission profile. If no matching dataflow endpoints are found then Ground Station will not display any available contacts for the satellite.

_Required_: No

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Create a DataflowEndpoint

The following example creates a Ground Station `DataflowEndpoint`

#### JSON

```json

{
  "Endpoint": {
    "Name": "myEndpoint",
    "Address": {
      "Name": "172.10.0.2",
      "Port": 44720
    },
    "Mtu": 1500
}
```

#### YAML

```yaml

Endpoint:
  Name: myEndpoint
  Address:
    Name: 172.10.0.2
    Port: 44720
  Mtu: 1500
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionDetails

EndpointDetails

All content copied from https://docs.aws.amazon.com/.
