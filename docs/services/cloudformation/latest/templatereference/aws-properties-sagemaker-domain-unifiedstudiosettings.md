This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain UnifiedStudioSettings

The settings that apply to an Amazon SageMaker AI domain when you use it in
Amazon SageMaker Unified Studio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainAccountId" : String,
  "DomainId" : String,
  "DomainRegion" : String,
  "EnvironmentId" : String,
  "ProjectId" : String,
  "ProjectS3Path" : String,
  "StudioWebPortalAccess" : String
}

```

### YAML

```yaml

  DomainAccountId: String
  DomainId: String
  DomainRegion: String
  EnvironmentId: String
  ProjectId: String
  ProjectS3Path: String
  StudioWebPortalAccess: String

```

## Properties

`DomainAccountId`

The ID of the AWS account that has the Amazon SageMaker Unified Studio domain. The default value, if
you don't specify an ID, is the ID of the account that has the Amazon SageMaker AI
domain.

_Required_: No

_Type_: String

_Pattern_: `^\d+$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainId`

The ID of the Amazon SageMaker Unified Studio domain associated with this domain.

_Required_: No

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainRegion`

The AWS Region where the domain is located in Amazon SageMaker Unified Studio. The default value, if
you don't specify a Region, is the Region where the Amazon SageMaker AI domain is
located.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z]{2}-[a-zA-Z\-]+-\d+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentId`

The ID of the environment that Amazon SageMaker Unified Studio associates with the domain.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectId`

The ID of the Amazon SageMaker Unified Studio project that corresponds to the domain.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectS3Path`

The location where Amazon S3 stores temporary execution data and other artifacts
for the project that corresponds to the domain.

_Required_: No

_Type_: String

_Pattern_: `[\w\.-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StudioWebPortalAccess`

Sets whether you can access the domain in Amazon SageMaker Studio:

ENABLED

You can access the domain in Amazon SageMaker Studio. If you migrate the domain to
Amazon SageMaker Unified Studio, you can access it in both studio interfaces.

DISABLED

You can't access the domain in Amazon SageMaker Studio. If you migrate the domain to
Amazon SageMaker Unified Studio, you can access it only in that studio interface.

To migrate a domain to Amazon SageMaker Unified Studio, you specify the UnifiedStudioSettings data type when you
use the UpdateDomain action.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UserSettings

All content copied from https://docs.aws.amazon.com/.
