This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::TemplateGroupAccessControlEntry AccessRights

Allow or deny permissions for an Active Directory group to enroll or autoenroll certificates for a
template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoEnroll" : String,
  "Enroll" : String
}

```

### YAML

```yaml

  AutoEnroll: String
  Enroll: String

```

## Properties

`AutoEnroll`

Allow or deny an Active Directory group from autoenrolling certificates issued against a template.
The Active Directory group must be allowed to enroll to allow autoenrollment

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enroll`

Allow or deny an Active Directory group from enrolling certificates issued against a
template.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCAConnectorAD::TemplateGroupAccessControlEntry

Next

All content copied from https://docs.aws.amazon.com/.
