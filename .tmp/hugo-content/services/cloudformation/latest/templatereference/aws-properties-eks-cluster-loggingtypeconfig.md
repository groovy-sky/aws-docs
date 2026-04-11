This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster LoggingTypeConfig

The enabled logging type. For a list of the valid logging types, see the [`types` property of `LogSetup`](../../../../reference/eks/latest/apireference/api-logsetup.md#AmazonEKS-Type-LogSetup-types) in the
_Amazon EKS API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The name of the log type.

_Required_: No

_Type_: String

_Allowed values_: `api | audit | authenticator | controllerManager | scheduler`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging

OutpostConfig

All content copied from https://docs.aws.amazon.com/.
