This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LicenseManager::License ConsumptionConfiguration

Details about a consumption configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BorrowConfiguration" : BorrowConfiguration,
  "ProvisionalConfiguration" : ProvisionalConfiguration,
  "RenewType" : String
}

```

### YAML

```yaml

  BorrowConfiguration:
    BorrowConfiguration
  ProvisionalConfiguration:
    ProvisionalConfiguration
  RenewType: String

```

## Properties

`BorrowConfiguration`

Details about a borrow configuration.

_Required_: No

_Type_: [BorrowConfiguration](aws-properties-licensemanager-license-borrowconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionalConfiguration`

Details about a provisional configuration.

_Required_: No

_Type_: [ProvisionalConfiguration](aws-properties-licensemanager-license-provisionalconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenewType`

Renewal frequency.

_Required_: No

_Type_: String

_Allowed values_: `None | Weekly | Monthly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BorrowConfiguration

Entitlement

All content copied from https://docs.aws.amazon.com/.
