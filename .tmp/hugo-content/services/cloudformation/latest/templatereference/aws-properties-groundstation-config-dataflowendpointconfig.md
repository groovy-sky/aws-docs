This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config DataflowEndpointConfig

Provides information to AWS Ground Station about which IP endpoints to use during a contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataflowEndpointName" : String,
  "DataflowEndpointRegion" : String
}

```

### YAML

```yaml

  DataflowEndpointName: String
  DataflowEndpointRegion: String

```

## Properties

`DataflowEndpointName`

The name of the dataflow endpoint to use during contacts.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataflowEndpointRegion`

The region of the dataflow endpoint to use during contacts. When omitted, Ground Station will use the region of the contact.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a DataflowEndpointConfig

The following example creates a Ground Station `DataflowEndpointConfig`

#### JSON

```json

{
  "DataflowEndpointConfig": {
    "DataflowEndpointName": "Downlink Demod Decode",
    "DataflowEndpointRegion": "us-east-2"
  }
}
```

#### YAML

```yaml

DataflowEndpointConfig:
  DataflowEndpointName: "Downlink Demod Decode"
  DataflowEndpointRegion: "us-east-2"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigData

DecodeConfig

All content copied from https://docs.aws.amazon.com/.
