This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain DomainEndpointOptions

Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomEndpoint" : String,
  "CustomEndpointCertificateArn" : String,
  "CustomEndpointEnabled" : Boolean,
  "EnforceHTTPS" : Boolean,
  "TLSSecurityPolicy" : String
}

```

### YAML

```yaml

  CustomEndpoint: String
  CustomEndpointCertificateArn: String
  CustomEndpointEnabled: Boolean
  EnforceHTTPS: Boolean
  TLSSecurityPolicy: String

```

## Properties

`CustomEndpoint`

The fully qualified URL for your custom endpoint. Required if you enabled a custom endpoint
for the domain.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomEndpointCertificateArn`

The AWS Certificate Manager ARN for your domain's SSL/TLS certificate. Required if you
enabled a custom endpoint for the domain.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomEndpointEnabled`

True to enable a custom endpoint for the domain. If enabled, you must also provide values for `CustomEndpoint` and `CustomEndpointCertificateArn`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforceHTTPS`

True to require that all traffic to the domain arrive over HTTPS.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLSSecurityPolicy`

The minimum TLS version required for traffic to the domain. Valid values are TLS 1.3 (recommended) or 1.2:

- `Policy-Min-TLS-1-0-2019-07`

- `Policy-Min-TLS-1-2-2019-07`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ColdStorageOptions

EBSOptions
