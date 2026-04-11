This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ResourceSet DNSTargetResource

A component for DNS/routing control readiness checks and architecture checks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "HostedZoneArn" : String,
  "RecordSetId" : String,
  "RecordType" : String,
  "TargetResource" : TargetResource
}

```

### YAML

```yaml

  DomainName: String
  HostedZoneArn: String
  RecordSetId: String
  RecordType: String
  TargetResource:
    TargetResource

```

## Properties

`DomainName`

The domain name that acts as an ingress point to a portion of the customer application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostedZoneArn`

The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordSetId`

The Amazon Route 53 record set ID that uniquely identifies a DNS record, given a name and a type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordType`

The type of DNS record of the target resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetResource`

The target resource that the Route 53 record points to.

_Required_: No

_Type_: [TargetResource](aws-properties-route53recoveryreadiness-resourceset-targetresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53RecoveryReadiness::ResourceSet

NLBResource

All content copied from https://docs.aws.amazon.com/.
