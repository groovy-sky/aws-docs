This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Application

Creates an OpenSearch UI application. For more information, see [Using the OpenSearch user interface in Amazon OpenSearch Service](../../../opensearch-service/latest/developerguide/application.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchService::Application",
  "Properties" : {
      "AppConfigs" : [ AppConfig, ... ],
      "DataSources" : [ DataSource, ... ],
      "Endpoint" : String,
      "IamIdentityCenterOptions" : IamIdentityCenterOptions,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchService::Application
Properties:
  AppConfigs:
    - AppConfig
  DataSources:
    - DataSource
  Endpoint: String
  IamIdentityCenterOptions:
    IamIdentityCenterOptions
  Name: String
  Tags:
    - Tag

```

## Properties

`AppConfigs`

Property description not available.

_Required_: No

_Type_: Array of [AppConfig](aws-properties-opensearchservice-application-appconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSources`

Property description not available.

_Required_: No

_Type_: Array of [DataSource](aws-properties-opensearchservice-application-datasource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The endpoint URL of an OpenSearch application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamIdentityCenterOptions`

Settings container for integrating IAM Identity Center with OpenSearch UI applications,
which enables enabling secure user authentication and access control across multiple data
sources. This setup supports single sign-on (SSO) through IAM Identity Center, allowing
centralized user management.

_Required_: No

_Type_: [IamIdentityCenterOptions](aws-properties-opensearchservice-application-iamidentitycenteroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of an OpenSearch application.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]+`

_Minimum_: `3`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-opensearchservice-application-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the domain. See [Identifiers for IAM Entities](../../../iam/latest/userguide/index.md) in
_Using AWS Identity and Access Management_ for
more information.

`Id`

The unique identifier of an OpenSearch application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon OpenSearch Service

AppConfig

All content copied from https://docs.aws.amazon.com/.
