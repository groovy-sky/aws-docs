This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway SubjectAlternativeNameMatchers

An object that represents the methods by which a subject alternative name on a peer
Transport Layer Security (TLS) certificate can be matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : [ String, ... ]
}

```

### YAML

```yaml

  Exact:
    - String

```

## Properties

`Exact`

The values sent must match the specified values exactly.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingFormat

SubjectAlternativeNames

All content copied from https://docs.aws.amazon.com/.
