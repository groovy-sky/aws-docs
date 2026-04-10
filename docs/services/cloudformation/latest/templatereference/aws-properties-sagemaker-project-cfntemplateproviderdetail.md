This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Project CfnTemplateProviderDetail

Details about a CloudFormation template provider configuration and associated provisioning information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : [ CfnStackParameter, ... ],
  "RoleARN" : String,
  "TemplateName" : String,
  "TemplateURL" : String
}

```

### YAML

```yaml

  Parameters:
    - CfnStackParameter
  RoleARN: String
  TemplateName: String
  TemplateURL: String

```

## Properties

`Parameters`

An array of CloudFormation stack parameters.

_Required_: No

_Type_: Array of [CfnStackParameter](aws-properties-sagemaker-project-cfnstackparameter.md)

_Minimum_: `0`

_Maximum_: `180`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleARN`

The IAM role used by CloudFormation to create the stack.

_Required_: No

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateName`

The unique identifier of the template within the project.

_Required_: Yes

_Type_: String

_Pattern_: `(?=.{1,32}$)[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateURL`

The Amazon S3 URL of the CloudFormation template.

_Required_: Yes

_Type_: String

_Pattern_: `(?=.{1,1024}$)(https)://([^/]+)/(.+)`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CfnStackParameter

ProvisioningParameter

All content copied from https://docs.aws.amazon.com/.
