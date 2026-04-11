# Connecting to your DB cluster using IAM authentication

With IAM database authentication, you use an authentication token when you connect
to your DB cluster. An _authentication token_ is a
string of characters that you use instead of a password. After you generate an
authentication token, it's valid for 15 minutes before it expires. If you try to
connect using an expired token, the connection request is denied.

Every authentication token must be accompanied by a valid signature, using AWS
signature version 4. (For more information, see [Signature Version 4 signing\
process](../../../../general/latest/gr/signature-version-4.md) in the _AWS General Reference._) The AWS CLI
and an AWS SDK, such as the AWS SDK for Java or AWS SDK for Python (Boto3), can automatically sign each token you create.

You can use an authentication token when you connect to
Amazon Aurora from another AWS service,
such as AWS Lambda. By using a token, you can avoid placing a password in your code.
Alternatively, you can use an AWS SDK to programmatically create and programmatically sign an
authentication token.

After you have a signed IAM authentication token, you can connect to an Aurora DB
cluster. Following, you can find out how to do this using either a command
line tool or an AWS SDK, such as the AWS SDK for Java or AWS SDK for Python (Boto3).

For more information, see the following blog posts:

- [Use IAM authentication to connect with SQL Workbench/J to Aurora MySQL or Amazon RDS for MySQL](https://aws.amazon.com/blogs/database/use-iam-authentication-to-connect-with-sql-workbenchj-to-amazon-aurora-mysql-or-amazon-rds-for-mysql)

- [Using IAM authentication to connect with pgAdmin Amazon Aurora PostgreSQL or Amazon RDS for PostgreSQL](https://aws.amazon.com/blogs/database/using-iam-authentication-to-connect-with-pgadmin-amazon-aurora-postgresql-or-amazon-rds-for-postgresql)

###### Prerequisites

The following are prerequisites for connecting to your DB cluster using IAM authentication:

- [Enabling and disabling IAM database authentication](usingwithrds-iamdbauth-enabling.md)

- [Creating and using an IAM policy for IAM database access](usingwithrds-iamdbauth-iampolicy.md)

- [Creating a database account using IAM authentication](usingwithrds-iamdbauth-dbaccounts.md)

###### Topics

- [Connecting to your DB cluster using IAM authentication with the AWS drivers](iamdbauth-connecting-drivers.md)

- [Connecting to your DB cluster using IAM authentication from the command line: AWS CLI and mysql client](usingwithrds-iamdbauth-connecting-awscli.md)

- [Connecting to your DB cluster using IAM authentication from the command line: AWS CLI and psql client](usingwithrds-iamdbauth-connecting-awscli-postgresql.md)

- [Connecting to your DB cluster using IAM authentication and the AWS SDK for .NET](usingwithrds-iamdbauth-connecting-net.md)

- [Connecting to your DB cluster using IAM authentication and the AWS SDK for Go](usingwithrds-iamdbauth-connecting-go.md)

- [Connecting to your DB cluster using IAM authentication and the AWS SDK for Java](usingwithrds-iamdbauth-connecting-java.md)

- [Connecting to your DB cluster using IAM authentication and the AWS SDK for Python (Boto3)](usingwithrds-iamdbauth-connecting-python.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating a database account using IAM authentication

Connecting to your DB cluster using IAM authentication with the AWS drivers

All content copied from https://docs.aws.amazon.com/.
