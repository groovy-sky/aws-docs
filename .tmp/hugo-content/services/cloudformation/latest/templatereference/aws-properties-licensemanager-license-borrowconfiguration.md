This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LicenseManager::License BorrowConfiguration

Details about a borrow configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowEarlyCheckIn" : Boolean,
  "MaxTimeToLiveInMinutes" : Integer
}

```

### YAML

```yaml

  AllowEarlyCheckIn: Boolean
  MaxTimeToLiveInMinutes: Integer

```

## Properties

`AllowEarlyCheckIn`

Indicates whether early check-ins are allowed.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTimeToLiveInMinutes`

Maximum time for the borrow configuration, in minutes.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::LicenseManager::License

ConsumptionConfiguration

All content copied from https://docs.aws.amazon.com/.
