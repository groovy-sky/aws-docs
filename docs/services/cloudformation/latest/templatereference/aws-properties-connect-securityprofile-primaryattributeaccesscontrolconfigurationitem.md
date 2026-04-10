This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityProfile PrimaryAttributeAccessControlConfigurationItem

A primary attribute access control configuration item.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrimaryAttributeValues" : [ PrimaryAttributeValue, ... ]
}

```

### YAML

```yaml

  PrimaryAttributeValues:
    - PrimaryAttributeValue

```

## Properties

`PrimaryAttributeValues`

The item's primary attribute values.

_Required_: Yes

_Type_: Array of [PrimaryAttributeValue](aws-properties-connect-securityprofile-primaryattributevalue.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GranularAccessControlConfiguration

PrimaryAttributeValue

All content copied from https://docs.aws.amazon.com/.
