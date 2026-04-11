This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::AccessLogSubscription

Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner
can use the access logs to audit the services in the network. The service network owner can only
see access logs from clients and services that are associated with their service network. Access
log entries represent traffic originated from VPCs associated with that network. For more
information, see [Access logs](../../../vpc-lattice/latest/ug/monitoring-access-logs.md) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::AccessLogSubscription",
  "Properties" : {
      "DestinationArn" : String,
      "ResourceIdentifier" : String,
      "ServiceNetworkLogType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::AccessLogSubscription
Properties:
  DestinationArn: String
  ResourceIdentifier: String
  ServiceNetworkLogType: String
  Tags:
    - Tag

```

## Properties

`DestinationArn`

The Amazon Resource Name (ARN) of the destination. The supported destination types are
CloudWatch Log groups, Kinesis Data Firehose delivery streams, and Amazon S3 buckets.

_Required_: Yes

_Type_: String

_Pattern_: `^arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:([^/].*)?$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

The ID or ARN of the service network or service.

_Required_: No

_Type_: String

_Pattern_: `^((((sn)|(svc)|(rcfg))-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:((servicenetwork/sn)|(resourceconfiguration/rcfg)|(service/svc))-[0-9a-z]{17}))$`

_Minimum_: `17`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceNetworkLogType`

Log type of the service network.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE | RESOURCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the access log subscription.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-accesslogsubscription-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the access log
subscription.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the access log subscription.

`Id`

The ID of the access log subscription.

`ResourceArn`

The Amazon Resource Name (ARN) of the access log subscription.

`ResourceId`

The ID of the service network or service.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon VPC Lattice

Tag

All content copied from https://docs.aws.amazon.com/.
