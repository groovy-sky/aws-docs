This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin BasicAuthConfiguration

Information about the basic authentication credentials used to configure a
plugin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  RoleArn: String
  SecretArn: String

```

## Properties

`RoleArn`

The ARN of an IAM role used by Amazon Q Business to access the basic
authentication credentials stored in a Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The ARN of the Secrets Manager secret that stores the basic authentication
credentials used for plugin configuration..

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

APISchema

CustomPluginConfiguration

All content copied from https://docs.aws.amazon.com/.
