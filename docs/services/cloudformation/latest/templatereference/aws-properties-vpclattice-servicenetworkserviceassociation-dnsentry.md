This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetworkServiceAssociation DnsEntry

The DNS information.

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::VpcLattice::ServiceNetworkServiceAssociation

Tag
