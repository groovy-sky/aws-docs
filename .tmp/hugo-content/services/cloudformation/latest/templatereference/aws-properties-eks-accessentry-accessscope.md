This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::AccessEntry AccessScope

The scope of an `AccessPolicy` that's associated to an
`AccessEntry`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespaces" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Namespaces:
    - String
  Type: String

```

## Properties

`Namespaces`

A Kubernetes `namespace` that an access policy is scoped to. A value is required
if you specified `namespace` for `Type`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The scope type of an access policy.

_Required_: Yes

_Type_: String

_Allowed values_: `namespace | cluster`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessPolicy

Tag

All content copied from https://docs.aws.amazon.com/.
