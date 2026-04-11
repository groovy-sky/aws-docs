This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Application PersonalizationConfiguration

Configuration information about chat response personalization. For more information,
see [Personalizing chat responses](../../../amazonq/latest/qbusiness-ug/personalizing-chat-responses.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PersonalizationControlMode" : String
}

```

### YAML

```yaml

  PersonalizationControlMode: String

```

## Properties

`PersonalizationControlMode`

An option to allow Amazon Q Business to customize chat responses using user
specific metadata—specifically, location and job information—in your
IAM Identity Center instance.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

QAppsConfiguration

All content copied from https://docs.aws.amazon.com/.
