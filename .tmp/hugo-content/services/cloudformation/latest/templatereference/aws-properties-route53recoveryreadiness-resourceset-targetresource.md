This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ResourceSet TargetResource

The target resource that the Route 53 record points to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NLBResource" : NLBResource,
  "R53Resource" : R53ResourceRecord
}

```

### YAML

```yaml

  NLBResource:
    NLBResource
  R53Resource:
    R53ResourceRecord

```

## Properties

`NLBResource`

The Network Load Balancer resource that a DNS target resource points to.

_Required_: No

_Type_: [NLBResource](aws-properties-route53recoveryreadiness-resourceset-nlbresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`R53Resource`

The Route 53 resource that a DNS target resource record points to.

_Required_: No

_Type_: [R53ResourceRecord](aws-properties-route53recoveryreadiness-resourceset-r53resourcerecord.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
