This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::Integration Tag

Metadata assigned to an Amazon RDS resource consisting of a key-value pair.

For more information, see
[Tagging Amazon RDS resources](../../../amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with `aws:` or `rds:`. The string can only contain only the set of Unicode letters, digits, white-space, '\_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-@\]\*)$").

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with `aws:` or `rds:`. The string can only contain only the set of Unicode letters, digits, white-space, '\_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-@\]\*)$").

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example specifies a tag for a zero-ETL integration.

#### JSON

```json

"Tags" : [
   {
      "Key" : "keyname1",
      "Value" : "value1"
   },
   {
      "Key" : "keyname2",
      "Value" : "value2"
   }
]
```

#### YAML

```yaml

Tags:
  - Key: keyname1
    Value: value1
  - Key: keyname2
    Value: value2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RDS::Integration

AWS::RDS::OptionGroup

All content copied from https://docs.aws.amazon.com/.
