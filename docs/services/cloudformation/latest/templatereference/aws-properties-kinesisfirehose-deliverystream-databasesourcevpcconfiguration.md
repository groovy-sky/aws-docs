This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DatabaseSourceVPCConfiguration

The structure for details of the VPC Endpoint Service which Firehose uses to create a PrivateLink to the database.

Amazon Data Firehose is in preview release and is subject to change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcEndpointServiceName" : String
}

```

### YAML

```yaml

  VpcEndpointServiceName: String

```

## Properties

`VpcEndpointServiceName`

The VPC endpoint service name which Firehose uses to create a PrivateLink to the database. The endpoint service must have the Firehose service principle `firehose.amazonaws.com` as an allowed principal on the VPC endpoint service. The VPC endpoint service name is a string that looks like `com.amazonaws.vpce.<region>.<vpc-endpoint-service-id>`.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: String

_Pattern_: `([a-zA-Z0-9\-\_]+\.){2,3}vpce\.[a-zA-Z0-9\-]*\.vpce-svc\-[a-zA-Z0-9\-]{17}$`

_Minimum_: `47`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseSourceConfiguration

DatabaseTables

All content copied from https://docs.aws.amazon.com/.
