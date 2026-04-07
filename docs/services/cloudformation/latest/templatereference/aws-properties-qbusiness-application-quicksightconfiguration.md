This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Application QuickSightConfiguration

The Amazon Quick configuration for an Amazon Q Business application that uses Quick as the identity provider.
For more information, see [Creating an\
Amazon Quick integrated application](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-quicksight-integrated-application.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientNamespace" : String
}

```

### YAML

```yaml

  ClientNamespace: String

```

## Properties

`ClientNamespace`

The Amazon Quick namespace that is used as the identity provider. For more information about Quick namespaces, see
[Namespace operations](https://docs.aws.amazon.com/quicksight/latest/developerguide/namespace-operations.html).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QAppsConfiguration

Tag
