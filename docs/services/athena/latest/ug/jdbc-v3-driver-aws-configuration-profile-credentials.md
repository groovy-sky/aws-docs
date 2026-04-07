# AWS configuration profile credentials

You can use credentials stored in an AWS configuration profile by setting the
following connection parameters. AWS configuration profiles are typically stored in
files in the `~/.aws` directory). For information about AWS
configuration profiles, see [Use\
profiles](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/credentials-profiles.html) in the _AWS SDK for Java Developer Guide_.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `ProfileCredentials`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProvider_AWSCredentialsProviderClass_
_(deprecated)_Requirednone`ProfileCredentials`

## Profile name

The name of the AWS configuration profile whose credentials should be used to
authenticate the request to Athena.

Parameter nameAliasParameter typeDefault valueProfileNamenoneRequirednone

###### Note

The profile name can also be specified as the value of the
`CredentialsProviderArguments` parameter, although this use is
deprecated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Default

Instance
profile
