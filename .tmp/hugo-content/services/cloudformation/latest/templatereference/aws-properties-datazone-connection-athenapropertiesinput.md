This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection AthenaPropertiesInput

The Amazon Athena properties of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkgroupName" : String
}

```

### YAML

```yaml

  WorkgroupName: String

```

## Properties

`WorkgroupName`

The Amazon Athena workgroup name of a connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]+$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonQPropertiesInput

AuthenticationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
