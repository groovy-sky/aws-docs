This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluenceConfiguration

Provides the configuration information to connect to Confluence as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentConfiguration" : ConfluenceAttachmentConfiguration,
  "BlogConfiguration" : ConfluenceBlogConfiguration,
  "ExclusionPatterns" : [ String, ... ],
  "InclusionPatterns" : [ String, ... ],
  "PageConfiguration" : ConfluencePageConfiguration,
  "SecretArn" : String,
  "ServerUrl" : String,
  "SpaceConfiguration" : ConfluenceSpaceConfiguration,
  "Version" : String,
  "VpcConfiguration" : DataSourceVpcConfiguration
}

```

### YAML

```yaml

  AttachmentConfiguration:
    ConfluenceAttachmentConfiguration
  BlogConfiguration:
    ConfluenceBlogConfiguration
  ExclusionPatterns:
    - String
  InclusionPatterns:
    - String
  PageConfiguration:
    ConfluencePageConfiguration
  SecretArn: String
  ServerUrl: String
  SpaceConfiguration:
    ConfluenceSpaceConfiguration
  Version: String
  VpcConfiguration:
    DataSourceVpcConfiguration

```

## Properties

`AttachmentConfiguration`

Configuration information for indexing attachments to Confluence blogs and
pages.

_Required_: No

_Type_: [ConfluenceAttachmentConfiguration](aws-properties-kendra-datasource-confluenceattachmentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlogConfiguration`

Configuration information for indexing Confluence blogs.

_Required_: No

_Type_: [ConfluenceBlogConfiguration](aws-properties-kendra-datasource-confluenceblogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of regular expression patterns to exclude certain blog posts, pages, spaces, or
attachments in your Confluence. Content that matches the patterns are excluded from the
index. Content that doesn't match the patterns is included in the index. If content
matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence
and the content isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of regular expression patterns to include certain blog posts, pages, spaces, or
attachments in your Confluence. Content that matches the patterns are included in the
index. Content that doesn't match the patterns is excluded from the index. If content
matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence
and the content isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageConfiguration`

Configuration information for indexing Confluence pages.

_Required_: No

_Type_: [ConfluencePageConfiguration](aws-properties-kendra-datasource-confluencepageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the
user name and password required to connect to the Confluence instance. If you use
Confluence Cloud, you use a generated API token as the password.

You can also provide authentication credentials in the form of a personal access
token. For more information, see [Using a Confluence data\
source](../../../kendra/latest/dg/data-source-confluence.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerUrl`

The URL of your Confluence instance. Use the full URL of the server. For example,
_https://server.example.com:port/_. You can also use an IP
address, for example, _https://192.168.1.113/_.

_Required_: Yes

_Type_: String

_Pattern_: `^(https?|ftp|file)://([^\s]*)`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceConfiguration`

Configuration information for indexing Confluence spaces.

_Required_: No

_Type_: [ConfluenceSpaceConfiguration](aws-properties-kendra-datasource-confluencespaceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version or the type of Confluence installation to connect to.

_Required_: Yes

_Type_: String

_Allowed values_: `CLOUD | SERVER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

Configuration information for an Amazon Virtual Private Cloud to connect to your Confluence.
For more information, see [Configuring a VPC](../../../kendra/latest/dg/vpc-configuration.md).

_Required_: No

_Type_: [DataSourceVpcConfiguration](aws-properties-kendra-datasource-datasourcevpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceBlogToIndexFieldMapping

ConfluencePageConfiguration

All content copied from https://docs.aws.amazon.com/.
