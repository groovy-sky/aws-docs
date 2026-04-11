This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectTriggers

`ProjectTriggers` is a property of the [AWS CodeBuild Project](../userguide/aws-resource-codebuild-project.md)
resource that specifies webhooks that trigger an AWS CodeBuild build.

###### Note

The Webhook feature isn't available in AWS CloudFormation for GitHub Enterprise projects. Use the AWS CLI or AWS CodeBuild console to create the webhook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BuildType" : String,
  "FilterGroups" : [ [ , ... ], ... ],
  "ScopeConfiguration" : ScopeConfiguration,
  "Webhook" : Boolean
}

```

### YAML

```yaml

  BuildType: String
  FilterGroups:
    -
    -
  ScopeConfiguration:
    ScopeConfiguration
  Webhook: Boolean

```

## Properties

`BuildType`

Specifies the type of build this webhook will trigger. Allowed values are:

BUILD

A single build

BUILD\_BATCH

A batch build

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroups`

A list of lists of `WebhookFilter` objects used to determine which webhook
events are triggered. At least one `WebhookFilter` in the array must specify `EVENT` as its type.

_Required_: No

_Type_: Array of Array

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScopeConfiguration`

Contains configuration information about the scope for a webhook.

_Required_: No

_Type_: [ScopeConfiguration](aws-properties-codebuild-project-scopeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Webhook`

Specifies whether or not to begin automatically rebuilding the source code every time a code change is pushed to the repository.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectSourceVersion

RegistryCredential

All content copied from https://docs.aws.amazon.com/.
