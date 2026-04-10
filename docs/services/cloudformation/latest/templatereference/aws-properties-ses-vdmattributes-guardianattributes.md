This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::VdmAttributes GuardianAttributes

An object containing additional settings for your VDM configuration as applicable to
the Guardian.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OptimizedSharedDelivery" : String
}

```

### YAML

```yaml

  OptimizedSharedDelivery: String

```

## Properties

`OptimizedSharedDelivery`

Specifies the status of your VDM optimized shared delivery. Can be one of the
following:

- `ENABLED` – Amazon SES enables optimized shared delivery for your
account.

- `DISABLED` – Amazon SES disables optimized shared delivery for
your account.

_Required_: No

_Type_: String

_Pattern_: `ENABLED|DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashboardAttributes

Next

All content copied from https://docs.aws.amazon.com/.
