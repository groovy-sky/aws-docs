This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DomainName MutualTlsAuthentication

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TruststoreUri" : String,
  "TruststoreVersion" : String
}

```

### YAML

```yaml

  TruststoreUri: String
  TruststoreVersion: String

```

## Properties

`TruststoreUri`

An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example
`s3://bucket-name/key-name`. The truststore can contain certificates from public or private
certificate authorities. To update the truststore, upload a new version to S3, and then update
your custom domain name to use the new version. To update the truststore, you must have
permissions to access the S3 object.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TruststoreVersion`

The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EndpointConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
