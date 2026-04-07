# Default credentials

You can use the default credentials that you configure on your client system to
connect to Amazon Athena by setting the following connection parameters. For information
about using default credentials, see [Using the\
Default Credential Provider Chain](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default) in the
_AWS SDK for Java Developer Guide_.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `DefaultChain`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProvider_AWSCredentialsProviderClass_
_(deprecated)_Requirednone`DefaultChain`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM

AWS
configuration profile
