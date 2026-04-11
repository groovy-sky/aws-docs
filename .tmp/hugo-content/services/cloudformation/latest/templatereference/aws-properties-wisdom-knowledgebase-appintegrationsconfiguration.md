This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase AppIntegrationsConfiguration

Configuration information for Amazon AppIntegrations to automatically ingest content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppIntegrationArn" : String,
  "ObjectFields" : [ String, ... ]
}

```

### YAML

```yaml

  AppIntegrationArn: String
  ObjectFields:
    - String

```

## Properties

`AppIntegrationArn`

The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.

- For [Salesforce](https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm), your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not
provided, including at least `Id`, `ArticleNumber`, `VersionNumber`,
`Title`, `PublishStatus`, and `IsDeleted` as source fields.

- For [ServiceNow](https://developer.servicenow.com/dev.do), your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not
provided, including at least `number`, `short_description`, `sys_mod_count`,
`workflow_state`, and `active` as source fields.

- For [Zendesk](https://developer.zendesk.com/api-reference/help_center/help-center-api/articles), your AppIntegrations DataIntegration must have an ObjectConfiguration if `objectFields` is
not provided, including at least `id`, `title`, `updated_at`, and
`draft` as source fields.

- For [SharePoint](https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index), your AppIntegrations DataIntegration must have a FileConfiguration, including only file extensions
that are among `docx`, `pdf`, `html`, `htm`, and `txt`.

- For [Amazon S3](https://aws.amazon.com/s3), the ObjectConfiguration and
FileConfiguration of your AppIntegrations DataIntegration must be null. The `SourceURI` of your DataIntegration
must use the following format: `s3://your_s3_bucket_name`.

###### Important

The bucket policy of the corresponding S3 bucket must allow the AWS principal
`app-integrations.amazonaws.com` to perform `s3:ListBucket`, `s3:GetObject`, and
`s3:GetBucketLocation` against the bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]+?:[a-z-]+?:[a-z0-9-]*?:([0-9]{12})?:[a-zA-Z0-9-:/]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectFields`

The fields from the source that are made available to your agents in Amazon Q in Connect. Optional if ObjectConfiguration
is included in the provided DataIntegration.

- For [Salesforce](https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm), you must include at least `Id`, `ArticleNumber`,
`VersionNumber`, `Title`, `PublishStatus`, and `IsDeleted`.

- For [ServiceNow](https://developer.servicenow.com/dev.do), you must include at least `number`, `short_description`,
`sys_mod_count`, `workflow_state`, and `active`.

- For [Zendesk](https://developer.zendesk.com/api-reference/help_center/help-center-api/articles), you must include at least `id`, `title`, `updated_at`, and
`draft`.

Make sure to include additional fields. These fields are indexed and used to source recommendations.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `4096 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Wisdom::KnowledgeBase

BedrockFoundationModelConfiguration

All content copied from https://docs.aws.amazon.com/.
