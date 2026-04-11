This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServer

Specifies a route server to manage dynamic routing in a VPC.

Amazon VPC Route Server simplifies routing for traffic between workloads that are deployed within a VPC and its internet gateways. With this feature, VPC Route Server dynamically updates VPC and internet gateway route tables with your preferred routes to achieve routing fault tolerance for those workloads. This enables you to automatically reroute traffic within a VPC, which increases the manageability of VPC routing and interoperability with third-party workloads.

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteServer",
  "Properties" : {
      "AmazonSideAsn" : Integer,
      "PersistRoutes" : String,
      "PersistRoutesDuration" : Integer,
      "SnsNotificationsEnabled" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteServer
Properties:
  AmazonSideAsn: Integer
  PersistRoutes: String
  PersistRoutesDuration: Integer
  SnsNotificationsEnabled: Boolean
  Tags:
    - Tag

```

## Properties

`AmazonSideAsn`

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) for the appliance. Valid values are from 1 to 4294967295. We recommend using a private ASN in the 64512–65534 (16-bit ASN) or 4200000000–4294967294 (32-bit ASN) range.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4294967294`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PersistRoutes`

Indicates whether routes should be persisted after all BGP sessions are terminated.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PersistRoutesDuration`

The number of minutes a route server will wait after BGP is re-established to unpersist the routes in the FIB and RIB. Value must be in the range of 1-5. The default value is 1. Only valid if `persistRoutesState` is 'enabled'.

If you set the duration to 1 minute, then when your network appliance re-establishes BGP with route server, it has 1 minute to relearn it's adjacent network and advertise those routes to route server before route server resumes normal functionality. In most cases, 1 minute is probably sufficient. If, however, you have concerns that your BGP network may not be capable of fully re-establishing and re-learning everything in 1 minute, you can increase the duration up to 5 minutes.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsNotificationsEnabled`

Indicates whether SNS notifications are enabled for the route server. Enabling SNS notifications persists BGP status changes to an SNS topic provisioned by AWS.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Any tags assigned to the route server.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-routeserver-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the route server ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the route server.

`Id`

The ID of the route server.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::Route

Tag

All content copied from https://docs.aws.amazon.com/.
