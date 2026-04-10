This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource OneDriveConfiguration

Provides the configuration information to connect to OneDrive as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisableLocalGroups" : Boolean,
  "ExclusionPatterns" : [ String, ... ],
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "InclusionPatterns" : [ String, ... ],
  "OneDriveUsers" : OneDriveUsers,
  "SecretArn" : String,
  "TenantDomain" : String
}

```

### YAML

```yaml

  DisableLocalGroups: Boolean
  ExclusionPatterns:
    - String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  InclusionPatterns:
    - String
  OneDriveUsers:
    OneDriveUsers
  SecretArn: String
  TenantDomain: String

```

## Properties

`DisableLocalGroups`

`TRUE` to disable local groups information.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of regular expression patterns to exclude certain documents in your OneDrive.
Documents that match the patterns are excluded from the index. Documents that don't
match the patterns are included in the index. If a document matches both an inclusion
and exclusion pattern, the exclusion pattern takes precedence and the document isn't
included in the index.

The pattern is applied to the file name.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

A list of `DataSourceToIndexFieldMapping` objects that map OneDrive data
source attributes or field names to Amazon Kendra index field names. To create
custom fields, use the `UpdateIndex` API before you map to OneDrive fields.
For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
OneDrive data source field names must exist in your OneDrive custom metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of regular expression patterns to include certain documents in your OneDrive.
Documents that match the patterns are included in the index. Documents that don't match
the patterns are excluded from the index. If a document matches both an inclusion and
exclusion pattern, the exclusion pattern takes precedence and the document isn't
included in the index.

The pattern is applied to the file name.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneDriveUsers`

A list of user accounts whose documents should be indexed.

_Required_: Yes

_Type_: [OneDriveUsers](aws-properties-kendra-datasource-onedriveusers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that
contains the user name and password to connect to OneDrive. The user name should be the
application ID for the OneDrive application, and the password is the application key for
the OneDrive application.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TenantDomain`

The Azure Active Directory domain of the organization.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-z]{2,}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InlineCustomDocumentEnrichmentConfiguration

OneDriveUsers

All content copied from https://docs.aws.amazon.com/.
