This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::DevicePool Rule

Represents a condition for a device pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : String,
  "Operator" : String,
  "Value" : String
}

```

### YAML

```yaml

  Attribute: String
  Operator: String
  Value: String

```

## Properties

`Attribute`

The rule's stringified attribute. For example, specify the value as
`"\"abc\""`.

The supported operators for each attribute are provided in the following
list.

APPIUM\_VERSION

The Appium version for the test.

Supported operators: `CONTAINS`

ARN

The Amazon Resource Name (ARN) of the device (for example,
`arn:aws:devicefarm:us-west-2::device:12345Example`.

Supported operators: `EQUALS`,
`IN`, `NOT_IN`

AVAILABILITY

The current availability of the device. Valid values are AVAILABLE,
HIGHLY\_AVAILABLE, BUSY, or TEMPORARY\_NOT\_AVAILABLE.

Supported operators: `EQUALS`

FLEET\_TYPE

The fleet type. Valid values are PUBLIC or PRIVATE.

Supported operators: `EQUALS`

FORM\_FACTOR

The device form factor. Valid values are PHONE or TABLET.

Supported operators: `EQUALS`,
`IN`, `NOT_IN`

INSTANCE\_ARN

The Amazon Resource Name (ARN) of the device instance.

Supported operators: `IN`,
`NOT_IN`

INSTANCE\_LABELS

The label of the device instance.

Supported operators: `CONTAINS`

MANUFACTURER

The device manufacturer (for example, Apple).

Supported operators: `EQUALS`,
`IN`, `NOT_IN`

MODEL

The device model, such as Apple iPad Air 2 or Google Pixel.

Supported operators: `CONTAINS`,
`EQUALS`, `IN`, `NOT_IN`

OS\_VERSION

The operating system version (for example, 10.3.2).

Supported operators: `EQUALS`,
`GREATER_THAN`, `GREATER_THAN_OR_EQUALS`,
`IN`, `LESS_THAN`,
`LESS_THAN_OR_EQUALS`, `NOT_IN`

PLATFORM

The device platform. Valid values are ANDROID or IOS.

Supported operators: `EQUALS`,
`IN`, `NOT_IN`

REMOTE\_ACCESS\_ENABLED

Whether the device is enabled for remote access. Valid values are TRUE
or FALSE.

Supported operators: `EQUALS`

REMOTE\_DEBUG\_ENABLED

Whether the device is enabled for remote debugging. Valid values are
TRUE or FALSE.

Supported operators: `EQUALS`

Because remote debugging is [no\
longer supported](../../../devicefarm/latest/developerguide/history.md), this filter is ignored.

_Required_: No

_Type_: String

_Allowed values_: `ARN | PLATFORM | FORM_FACTOR | MANUFACTURER | REMOTE_ACCESS_ENABLED | REMOTE_DEBUG_ENABLED | APPIUM_VERSION | INSTANCE_ARN | INSTANCE_LABELS | FLEET_TYPE | OS_VERSION | MODEL | AVAILABILITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

Specifies how Device Farm compares the rule's attribute to the value. For the
operators that are supported by each attribute, see the attribute
descriptions.

_Required_: No

_Type_: String

_Allowed values_: `EQUALS | LESS_THAN | LESS_THAN_OR_EQUALS | GREATER_THAN | GREATER_THAN_OR_EQUALS | IN | NOT_IN | CONTAINS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The rule's value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DeviceFarm::DevicePool

Tag

All content copied from https://docs.aws.amazon.com/.
