This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ResourceSet Resource

The resource element of a resource set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentId" : String,
  "DnsTargetResource" : DNSTargetResource,
  "ReadinessScopes" : [ String, ... ],
  "ResourceArn" : String
}

```

### YAML

```yaml

  ComponentId: String
  DnsTargetResource:
    DNSTargetResource
  ReadinessScopes:
    - String
  ResourceArn: String

```

## Properties

`ComponentId`

The component identifier of the resource, generated when DNS target resource is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsTargetResource`

A component for DNS/routing control readiness checks. This is a required setting when `ResourceSet` `ResourceSetType` is
set to `AWS::Route53RecoveryReadiness::DNSTargetResource`. Do not set it for any other `ResourceSetType` setting.

_Required_: Conditional

_Type_: [DNSTargetResource](aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadinessScopes`

The recovery group Amazon Resource Name (ARN) or the cell ARN that the readiness checks for this resource set are scoped to.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the AWS resource. This is a required setting for all `ResourceSet` `ResourceSetType`
settings except `AWS::Route53RecoveryReadiness::DNSTargetResource`. Do not set this when `ResourceSetType` is
set to `AWS::Route53RecoveryReadiness::DNSTargetResource`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

R53ResourceRecord

Tag

All content copied from https://docs.aws.amazon.com/.
