This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ServiceNowConfiguration

Provides the configuration information to connect to ServiceNow as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String,
  "HostUrl" : String,
  "KnowledgeArticleConfiguration" : ServiceNowKnowledgeArticleConfiguration,
  "SecretArn" : String,
  "ServiceCatalogConfiguration" : ServiceNowServiceCatalogConfiguration,
  "ServiceNowBuildVersion" : String
}

```

### YAML

```yaml

  AuthenticationType: String
  HostUrl: String
  KnowledgeArticleConfiguration:
    ServiceNowKnowledgeArticleConfiguration
  SecretArn: String
  ServiceCatalogConfiguration:
    ServiceNowServiceCatalogConfiguration
  ServiceNowBuildVersion: String

```

## Properties

`AuthenticationType`

The type of authentication used to connect to the ServiceNow instance. If you choose
`HTTP_BASIC`, Amazon Kendra is authenticated using the user name and
password provided in the AWS Secrets Manager secret in the `SecretArn`
field. If you choose `OAUTH2`, Amazon Kendra is authenticated using the
credentials of client ID, client secret, user name and password.

When you use `OAUTH2` authentication, you must generate a token and a
client secret using the ServiceNow console. For more information, see [Using a\
ServiceNow data source](../../../kendra/latest/dg/data-source-servicenow.md).

_Required_: No

_Type_: String

_Allowed values_: `HTTP_BASIC | OAUTH2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostUrl`

The ServiceNow instance that the data source connects to. The host endpoint should
look like the following: _{instance}.service-now.com._

_Required_: Yes

_Type_: String

_Pattern_: `^(?!(^(https?|ftp|file):\/\/))[a-z0-9-]+(\.service-now\.com)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeArticleConfiguration`

Configuration information for crawling knowledge articles in the ServiceNow
site.

_Required_: No

_Type_: [ServiceNowKnowledgeArticleConfiguration](aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the
user name and password required to connect to the ServiceNow instance. You can also
provide OAuth authentication credentials of user name, password, client ID, and client
secret. For more information, see [Using a ServiceNow data\
source](../../../kendra/latest/dg/data-source-servicenow.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceCatalogConfiguration`

Configuration information for crawling service catalogs in the ServiceNow site.

_Required_: No

_Type_: [ServiceNowServiceCatalogConfiguration](aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNowBuildVersion`

The identifier of the release that the ServiceNow host is running. If the host is not
running the `LONDON` release, use `OTHERS`.

_Required_: Yes

_Type_: String

_Allowed values_: `LONDON | OTHERS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceStandardObjectConfiguration

ServiceNowKnowledgeArticleConfiguration

All content copied from https://docs.aws.amazon.com/.
