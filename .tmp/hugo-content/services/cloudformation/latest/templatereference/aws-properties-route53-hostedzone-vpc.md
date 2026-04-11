This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HostedZone VPC

_Private hosted zones only:_ A complex type that contains information about an Amazon VPC. Route 53 Resolver
uses the records in the private hosted zone to route traffic in that VPC.

###### Note

For public hosted zones, omit `VPCs`, `VPCId`, and `VPCRegion`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VPCId" : String,
  "VPCRegion" : String
}

```

### YAML

```yaml

  VPCId: String
  VPCRegion: String

```

## Properties

`VPCId`

_Private hosted zones only:_ The ID of an Amazon VPC.

###### Note

For public hosted zones, omit `VPCs`, `VPCId`, and `VPCRegion`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCRegion`

_Private hosted zones only:_ The region that an Amazon VPC was created in.

###### Note

For public hosted zones, omit `VPCs`, `VPCId`, and `VPCRegion`.

_Required_: Yes

_Type_: String

_Allowed values_: `us-east-1 | us-east-2 | us-west-1 | us-west-2 | eu-west-1 | eu-west-2 | eu-west-3 | eu-central-1 | eu-central-2 | ap-east-1 | me-south-1 | us-gov-west-1 | us-gov-east-1 | us-iso-east-1 | us-iso-west-1 | us-isob-east-1 | me-central-1 | ap-southeast-1 | ap-southeast-2 | ap-southeast-3 | ap-south-1 | ap-south-2 | ap-northeast-1 | ap-northeast-2 | ap-northeast-3 | eu-north-1 | sa-east-1 | ca-central-1 | cn-north-1 | cn-northwest-1 | af-south-1 | eu-south-1 | eu-south-2 | ap-southeast-4 | il-central-1 | ca-west-1 | ap-southeast-5 | mx-central-1 | us-isof-south-1 | us-isof-east-1 | ap-southeast-7 | ap-east-2 | eu-isoe-west-1 | ap-southeast-6 | us-isob-west-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](../userguide/aws-resource-route53-hostedzone.md#aws-resource-route53-hostedzone-return-values)
in the topic
[AWS::Route53::HostedZone](../userguide/aws-resource-route53-hostedzone.md)

- [VPC](../../../../reference/route53/latest/apireference/api-vpc.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryLoggingConfig

AWS::Route53::KeySigningKey

All content copied from https://docs.aws.amazon.com/.
