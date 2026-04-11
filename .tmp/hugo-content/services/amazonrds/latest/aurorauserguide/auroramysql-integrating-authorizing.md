# Authorizing Amazon Aurora MySQL to access other AWS services on your behalf

For your Aurora MySQL DB cluster to access other services on your behalf, create and
configure an AWS Identity and Access Management (IAM) role. This role authorizes database users in your DB
cluster to access other AWS services. For more information, see [Setting up IAM roles to access AWS services](auroramysql-integrating-authorizing-iam.md).

You must also configure your Aurora DB cluster to allow outbound connections to the
target AWS service. For more information, see [Enabling network communication from Amazon Aurora to other AWS services](auroramysql-integrating-authorizing-network.md).

If you do so, your database users can perform these actions using other AWS
services:

- Synchronously or asynchronously invoke an AWS Lambda function using the native functions
`lambda_sync` or `lambda_async`. Or, asynchronously invoke an AWS Lambda
function using the `mysql.lambda_async` procedure. For more information, see [Invoking a Lambda function with an Aurora MySQL native function](auroramysql-integrating-nativelambda.md).

- Load data from text or XML files stored in an Amazon S3 bucket into your DB cluster
by using the `LOAD DATA FROM S3` or `LOAD XML FROM S3`
statement. For more information, see [Loading data into an Amazon Aurora MySQL DB cluster from text files in an Amazon S3 bucket](auroramysql-integrating-loadfroms3.md).

- Save data from your DB cluster into text files stored in an Amazon S3 bucket by
using the `SELECT INTO OUTFILE S3` statement. For more information,
see [Saving data from an Amazon Aurora MySQL DB cluster into text files in an Amazon S3 bucket](auroramysql-integrating-saveintos3.md).

- Export log data to Amazon CloudWatch Logs MySQL. For more information, see [Publishing Amazon Aurora MySQL logs to Amazon CloudWatch Logs](auroramysql-integrating-cloudwatch.md).

- Automatically add or remove Aurora Replicas with Application Auto Scaling. For more information,
see [Amazon Aurora Auto Scaling with Aurora Replicas](aurora-integrating-autoscaling.md).

## Related topics

- [Integrating Aurora with other AWS services](aurora-integrating.md)

- [Managing an Amazon Aurora DB cluster](chap-aurora.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Integrating Aurora MySQL with AWS services

Setting up IAM roles to access AWS services

All content copied from https://docs.aws.amazon.com/.
