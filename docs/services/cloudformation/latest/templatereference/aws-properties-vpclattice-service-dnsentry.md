This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Service DnsEntry

Describes the DNS information of a service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "HostedZoneId" : String
}

```

### YAML

```yaml

  DomainName: String
  HostedZoneId: String

```

## Properties

`DomainName`

The domain name of the service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostedZoneId`

The ID of the hosted zone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::Service

Tag

All content copied from https://docs.aws.amazon.com/.
