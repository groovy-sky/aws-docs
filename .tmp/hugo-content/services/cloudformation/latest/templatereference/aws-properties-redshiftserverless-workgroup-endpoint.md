This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup Endpoint

The VPC endpoint object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String,
  "Port" : Integer,
  "VpcEndpoints" : [ VpcEndpoint, ... ]
}

```

### YAML

```yaml

  Address: String
  Port: Integer
  VpcEndpoints:
    - VpcEndpoint

```

## Properties

`Address`

The DNS address of the VPC endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port that Amazon Redshift Serverless listens on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpoints`

An array of `VpcEndpoint` objects.

_Required_: No

_Type_: Array of [VpcEndpoint](aws-properties-redshiftserverless-workgroup-vpcendpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigParameter

NetworkInterface

All content copied from https://docs.aws.amazon.com/.
