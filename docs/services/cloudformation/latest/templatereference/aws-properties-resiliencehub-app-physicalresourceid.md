This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::App PhysicalResourceId

Defines a physical resource identifier.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccountId" : String,
  "AwsRegion" : String,
  "Identifier" : String,
  "Type" : String
}

```

### YAML

```yaml

  AwsAccountId: String
  AwsRegion: String
  Identifier: String
  Type: String

```

## Properties

`AwsAccountId`

The AWS account that owns the physical resource.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsRegion`

The AWS Region that the physical resource is located in.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Identifier`

Identifier of the physical resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of physical resource identifier.

Arn

The resource identifier is an Amazon Resource Name (ARN) and it can identify the
following list of resources:

- `AWS::ECS::Service`

- `AWS::EFS::FileSystem`

- `AWS::ElasticLoadBalancingV2::LoadBalancer`

- `AWS::Lambda::Function`

- `AWS::SNS::Topic`

Native

The resource identifier is an AWS Resilience Hub-native identifier and it can
identify the following list of resources:

- `AWS::ApiGateway::RestApi`

- `AWS::ApiGatewayV2::Api`

- `AWS::AutoScaling::AutoScalingGroup`

- `AWS::DocDB::DBCluster`

- `AWS::DocDB::DBGlobalCluster`

- `AWS::DocDB::DBInstance`

- `AWS::DynamoDB::GlobalTable`

- `AWS::DynamoDB::Table`

- `AWS::EC2::EC2Fleet`

- `AWS::EC2::Instance`

- `AWS::EC2::NatGateway`

- `AWS::EC2::Volume`

- `AWS::ElasticLoadBalancing::LoadBalancer`

- `AWS::RDS::DBCluster`

- `AWS::RDS::DBInstance`

- `AWS::RDS::GlobalCluster`

- `AWS::Route53::RecordSet`

- `AWS::S3::Bucket`

- `AWS::SQS::Queue`

_Required_: Yes

_Type_: String

_Pattern_: `Arn|Native`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PermissionModel

ResourceMapping

All content copied from https://docs.aws.amazon.com/.
