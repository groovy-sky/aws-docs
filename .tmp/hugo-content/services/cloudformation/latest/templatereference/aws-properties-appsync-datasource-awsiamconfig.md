This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource AwsIamConfig

Use the `AwsIamConfig` property type to specify `AwsIamConfig`
for a AWS AppSync authorizaton.

`AwsIamConfig` is a property of the [AWS AppSync DataSource AuthorizationConfig](../userguide/aws-properties-appsync-datasource-httpconfig-authorizationconfig.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SigningRegion" : String,
  "SigningServiceName" : String
}

```

### YAML

```yaml

  SigningRegion: String
  SigningServiceName: String

```

## Properties

`SigningRegion`

The signing Region for AWS Identity and Access Management authorization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningServiceName`

The signing service name for AWS Identity and Access Management authorization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationConfig

DeltaSyncConfig

All content copied from https://docs.aws.amazon.com/.
