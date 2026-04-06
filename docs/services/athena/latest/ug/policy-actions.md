# Control access through JDBC and ODBC connections

To gain access to AWS services and resources, such as Athena and the Amazon S3 buckets,
provide the JDBC or ODBC driver credentials to your application. If you are using the
JDBC or ODBC driver, ensure that the IAM permissions policy includes all of the
actions listed in [AWS managed policy: AWSQuicksightAthenaAccess](security-iam-awsmanpol.md#awsquicksightathenaaccess-managed-policy).

Whenever you use IAM policies, make sure that you follow IAM best practices. For more information, see [Security best practices in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) in the _IAM User Guide_.

## Authentication methods

The Athena JDBC and ODBC drivers support SAML 2.0-based authentication, including
the following identity providers:

- Active Directory Federation Services (AD FS)

- Azure Active Directory (AD)

- Okta

- PingFederate

For more information, see the installation and configuration guides for the
respective drivers, downloadable in PDF format from the [JDBC](connect-with-jdbc.md) and
[ODBC](connect-with-odbc.md)
driver pages. For additional related information, see the following:

- [Enable federated access to the Athena API](https://docs.aws.amazon.com/athena/latest/ug/access-federation-saml.html)

- [Use Lake Formation and JDBC or ODBC drivers for federated access to Athena](https://docs.aws.amazon.com/athena/latest/ug/security-athena-lake-formation-jdbc.html)

- [Configure single sign-on using ODBC, SAML 2.0, and the Okta Identity Provider](https://docs.aws.amazon.com/athena/latest/ug/okta-saml-sso.html)

For information about the latest versions of the JDBC and ODBC drivers and their
documentation, see [Connect to Amazon Athena with JDBC](connect-with-jdbc.md) and [Connect to Amazon Athena with ODBC](connect-with-odbc.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Data perimeters

Control access to Amazon S3 from Athena
