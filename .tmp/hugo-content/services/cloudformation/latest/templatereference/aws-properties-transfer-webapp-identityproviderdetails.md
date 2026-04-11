This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp IdentityProviderDetails

A structure that describes the values to use for the IAM Identity Center settings when
you create or update a web app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationArn" : String,
  "InstanceArn" : String,
  "Role" : String
}

```

### YAML

```yaml

  ApplicationArn: String
  InstanceArn: String
  Role: String

```

## Properties

`ApplicationArn`

The Amazon Resource Name (ARN) for the IAM Identity Center application: this value is
set automatically when you create your web app.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w-]+:sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) for the IAM Identity Center used for the web app.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w-]+:sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Role`

The IAM role in IAM Identity Center used for the web app.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z-]+:iam::[0-9]{12}:role[:/]\S+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EndpointDetails

Tag

All content copied from https://docs.aws.amazon.com/.
