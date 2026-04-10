This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode SubjectAlternativeNames

An object that represents the subject alternative names secured by the
certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Match" : SubjectAlternativeNameMatchers
}

```

### YAML

```yaml

  Match:
    SubjectAlternativeNameMatchers

```

## Properties

`Match`

An object that represents the criteria for determining a SANs match.

_Required_: Yes

_Type_: [SubjectAlternativeNameMatchers](aws-properties-appmesh-virtualnode-subjectalternativenamematchers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubjectAlternativeNameMatchers

Tag

All content copied from https://docs.aws.amazon.com/.
