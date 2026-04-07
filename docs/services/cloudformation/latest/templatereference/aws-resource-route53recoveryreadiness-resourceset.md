This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ResourceSet

Creates a resource set in Amazon Route 53 Application Recovery Controller. A resource set is a set of resources of one type, such as Network
Load Balancers, that span multiple cells. You can associate a resource set with a readiness check to have Route 53 ARC continually monitor the
resources in the set for failover readiness.

You typically create a resource set and a readiness check for each supported type of AWS resource in your application.

For more information, see
[Readiness checks, resource sets, \
and readiness scopes](https://docs.aws.amazon.com/r53recovery/latest/dg/recovery-readiness.recovery-groups.readiness-scope.html)
in the Amazon Route 53 Application Recovery Controller Developer Guide.

Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryReadiness::ResourceSet",
  "Properties" : {
      "Resources" : [ Resource, ... ],
      "ResourceSetName" : String,
      "ResourceSetType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryReadiness::ResourceSet
Properties:
  Resources:
    - Resource
  ResourceSetName: String
  ResourceSetType: String
  Tags:
    - Tag

```

## Properties

`Resources`

A list of resource objects in the resource set.

_Required_: Yes

_Type_: Array of [Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53recoveryreadiness-resourceset-resource.html)

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceSetName`

The name of the resource set to create.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceSetType`

The resource type of the resources in the resource set. Enter one of the following values for resource type:

AWS::ApiGateway::Stage,
AWS::ApiGatewayV2::Stage,
AWS::AutoScaling::AutoScalingGroup,
AWS::CloudWatch::Alarm,
AWS::EC2::CustomerGateway,
AWS::DynamoDB::Table,
AWS::EC2::Volume,
AWS::ElasticLoadBalancing::LoadBalancer,
AWS::ElasticLoadBalancingV2::LoadBalancer,
AWS::Lambda::Function,
AWS::MSK::Cluster,
AWS::RDS::DBCluster,
AWS::Route53::HealthCheck,
AWS::SQS::Queue,
AWS::SNS::Topic,
AWS::SNS::Subscription,
AWS::EC2::VPC,
AWS::EC2::VPNConnection,
AWS::EC2::VPNGateway,
AWS::Route53RecoveryReadiness::DNSTargetResource.

Note that AWS::Route53RecoveryReadiness::DNSTargetResource is only used for this setting. It isn't an actual AWS CloudFormation resource type.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A tag to associate with the parameters for a resource set.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53recoveryreadiness-resourceset-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResourceSetName` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceSetArn`

The Amazon Resource Name (ARN) of the resource set.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DNSTargetResource
