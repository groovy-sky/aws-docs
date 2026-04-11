This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ResourceSet R53ResourceRecord

The Amazon Route 53 resource that a DNS target resource record points to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "RecordSetId" : String
}

```

### YAML

```yaml

  DomainName: String
  RecordSetId: String

```

## Properties

`DomainName`

The DNS target domain name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordSetId`

The Amazon Route 53 Resource Record Set ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NLBResource

Resource

All content copied from https://docs.aws.amazon.com/.
