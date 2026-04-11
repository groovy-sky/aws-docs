This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Domain

Describes a domain where you are storing recordsets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Domain",
  "Properties" : {
      "DomainEntries" : [ DomainEntry, ... ],
      "DomainName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Domain
Properties:
  DomainEntries:
    - DomainEntry
  DomainName: String
  Tags:
    - Tag

```

## Properties

`DomainEntries`

An array of key-value pairs containing information about the domain entries.

_Required_: No

_Type_: Array of [DomainEntry](aws-properties-lightsail-domain-domainentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The fully qualified domain name in the certificate request.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tag keys and optional values for the resource. For more information about tags in
Lightsail, see the [Amazon Lightsail Developer Guide](../../../lightsail/latest/userguide/amazon-lightsail-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-domain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the domain recordset
( `arn:aws:lightsail:global:123456789101:Domain/824cede0-abc7-4f84-8dbc-12345EXAMPLE`).

`CreatedAt`

The date when the domain recordset was created.

`ResourceType`

The resource type.

`SupportCode`

The support code. Include this code in your email to support when you have questions about
an instance or another resource in Lightsail. This code enables our support team to look up
your Lightsail information more easily.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DomainEntry

All content copied from https://docs.aws.amazon.com/.
