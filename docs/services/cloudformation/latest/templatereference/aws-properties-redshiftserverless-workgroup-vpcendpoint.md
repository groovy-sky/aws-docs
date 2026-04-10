This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup VpcEndpoint

The connection endpoint for connecting to Amazon Redshift Serverless through the proxy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkInterfaces" : [ NetworkInterface, ... ],
  "VpcEndpointId" : String,
  "VpcId" : String
}

```

### YAML

```yaml

  NetworkInterfaces:
    - NetworkInterface
  VpcEndpointId: String
  VpcId: String

```

## Properties

`NetworkInterfaces`

One or more network interfaces of the endpoint. Also known as an interface endpoint.

_Required_: No

_Type_: Array of [NetworkInterface](aws-properties-redshiftserverless-workgroup-networkinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpointId`

The connection endpoint ID for connecting to Amazon Redshift Serverless.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The VPC identifier that the endpoint is associated with.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Workgroup

All content copied from https://docs.aws.amazon.com/.
