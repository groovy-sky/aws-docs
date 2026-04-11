This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template ApplicationPolicies

Application policies describe what the certificate can be used for.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Critical" : Boolean,
  "Policies" : [ ApplicationPolicy, ... ]
}

```

### YAML

```yaml

  Critical: Boolean
  Policies:
    - ApplicationPolicy

```

## Properties

`Critical`

Marks the application policy extension as critical.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

Application policies describe what the certificate can be used for.

_Required_: Yes

_Type_: Array of [ApplicationPolicy](aws-properties-pcaconnectorad-template-applicationpolicy.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCAConnectorAD::Template

ApplicationPolicy

All content copied from https://docs.aws.amazon.com/.
