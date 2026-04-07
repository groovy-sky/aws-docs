This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain DomainEndpointOptions

Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.

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

_Pattern_: `^(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])$`

_Minimum_: `1`

_Maximum_: `255`

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

True to require that all traffic to the domain arrive over HTTPS. Required if you enable
fine-grained access control in [AdvancedSecurityOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html).

_Required_: Conditional

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`TLSSecurityPolicy`

The minimum TLS version required for traffic to the domain. The policy can be one of the
following values:

- **Policy-Min-TLS-1-0-2019-07:** TLS security policy that
supports TLS version 1.0 to TLS version 1.2

- **Policy-Min-TLS-1-2-2019-07:** TLS security policy that
supports only TLS version 1.2

- **Policy-Min-TLS-1-2-PFS-2023-10:** TLS security policy
that supports TLS version 1.2 to TLS version 1.3 with perfect forward secrecy cipher
suites

_Required_: No

_Type_: String

_Allowed values_: `Policy-Min-TLS-1-0-2019-07 | Policy-Min-TLS-1-2-2019-07 | Policy-Min-TLS-1-2-PFS-2023-10 | Policy-Min-TLS-1-2-RFC9151-FIPS-2024-08`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeploymentStrategyOptions

EBSOptions
