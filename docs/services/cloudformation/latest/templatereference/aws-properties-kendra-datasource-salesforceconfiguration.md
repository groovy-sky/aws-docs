This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceConfiguration

Provides the configuration information to connect to Salesforce as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChatterFeedConfiguration" : SalesforceChatterFeedConfiguration,
  "CrawlAttachments" : Boolean,
  "ExcludeAttachmentFilePatterns" : [ String, ... ],
  "IncludeAttachmentFilePatterns" : [ String, ... ],
  "KnowledgeArticleConfiguration" : SalesforceKnowledgeArticleConfiguration,
  "SecretArn" : String,
  "ServerUrl" : String,
  "StandardObjectAttachmentConfiguration" : SalesforceStandardObjectAttachmentConfiguration,
  "StandardObjectConfigurations" : [ SalesforceStandardObjectConfiguration, ... ]
}

```

### YAML

```yaml

  ChatterFeedConfiguration:
    SalesforceChatterFeedConfiguration
  CrawlAttachments: Boolean
  ExcludeAttachmentFilePatterns:
    - String
  IncludeAttachmentFilePatterns:
    - String
  KnowledgeArticleConfiguration:
    SalesforceKnowledgeArticleConfiguration
  SecretArn: String
  ServerUrl: String
  StandardObjectAttachmentConfiguration:
    SalesforceStandardObjectAttachmentConfiguration
  StandardObjectConfigurations:
    - SalesforceStandardObjectConfiguration

```

## Properties

`ChatterFeedConfiguration`

Configuration information for Salesforce chatter feeds.

_Required_: No

_Type_: [SalesforceChatterFeedConfiguration](aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlAttachments`

Indicates whether Amazon Kendra should index attachments to Salesforce
objects.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeAttachmentFilePatterns`

A list of regular expression patterns to exclude certain documents in your Salesforce.
Documents that match the patterns are excluded from the index. Documents that don't
match the patterns are included in the index. If a document matches both an inclusion
and exclusion pattern, the exclusion pattern takes precedence and the document isn't
included in the index.

The pattern is applied to the name of the attached file.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeAttachmentFilePatterns`

A list of regular expression patterns to include certain documents in your Salesforce.
Documents that match the patterns are included in the index. Documents that don't match
the patterns are excluded from the index. If a document matches both an inclusion and
exclusion pattern, the exclusion pattern takes precedence and the document isn't
included in the index.

The pattern is applied to the name of the attached file.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeArticleConfiguration`

Configuration information for the knowledge article types that Amazon Kendra
indexes. Amazon Kendra indexes standard knowledge articles and the standard fields
of knowledge articles, or the custom fields of custom knowledge articles, but not
both.

_Required_: No

_Type_: [SalesforceKnowledgeArticleConfiguration](aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of an AWS Secrets Managersecret that contains the
key/value pairs required to connect to your Salesforce instance. The secret must contain
a JSON structure with the following keys:

- authenticationUrl - The OAUTH endpoint that Amazon Kendra connects to get
an OAUTH token.

- consumerKey - The application public key generated when you created your
Salesforce application.

- consumerSecret - The application private key generated when you created your
Salesforce application.

- password - The password associated with the user logging in to the Salesforce
instance.

- securityToken - The token associated with the user logging in to the
Salesforce instance.

- username - The user name of the user logging in to the Salesforce
instance.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerUrl`

The instance URL for the Salesforce site that you want to index.

_Required_: Yes

_Type_: String

_Pattern_: `^(https?|ftp|file)://([^\s]*)`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardObjectAttachmentConfiguration`

Configuration information for processing attachments to Salesforce standard objects.

_Required_: No

_Type_: [SalesforceStandardObjectAttachmentConfiguration](aws-properties-kendra-datasource-salesforcestandardobjectattachmentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardObjectConfigurations`

Configuration of the Salesforce standard objects that Amazon Kendra
indexes.

_Required_: No

_Type_: Array of [SalesforceStandardObjectConfiguration](aws-properties-kendra-datasource-salesforcestandardobjectconfiguration.md)

_Minimum_: `1`

_Maximum_: `17`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceChatterFeedConfiguration

SalesforceCustomKnowledgeArticleTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
