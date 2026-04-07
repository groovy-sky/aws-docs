This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudExadataInfrastructure CustomerContact

A contact to receive notification from Oracle about maintenance updates for a specific Exadata infrastructure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Email" : String
}

```

### YAML

```yaml

  Email: String

```

## Properties

`Email`

The email address of the contact.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `320`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ODB::CloudExadataInfrastructure

MaintenanceWindow
