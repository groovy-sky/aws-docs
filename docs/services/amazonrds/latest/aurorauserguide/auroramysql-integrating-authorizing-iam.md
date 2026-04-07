# Setting up IAM roles to access AWS services

To permit your Aurora DB cluster to access another AWS service, do the
following:

1. Create an IAM policy that grants permission to the AWS service. For more
    information, see the following topics.

- [Creating an IAM policy to access Amazon S3 resources](auroramysql-integrating-authorizing-iam-s3createpolicy.md)

- [Creating an IAM policy to access AWS Lambda resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.LambdaCreatePolicy.html)

- [Creating an IAM policy to access CloudWatch Logs resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.CWCreatePolicy.html)

- [Creating an IAM policy to access AWS KMS resources](auroramysql-integrating-authorizing-iam-kmscreatepolicy.md)

2. Create an IAM role and attach the policy that you created. For more
    information, see [Creating an IAM role to allow Amazon Aurora to access AWS services](auroramysql-integrating-authorizing-iam-createrole.md).

3. Associate that IAM role with your Aurora DB cluster. For more information, see
    [Associating an IAM role with an Amazon Aurora MySQL DB cluster](auroramysql-integrating-authorizing-iam-addroletodbcluster.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Authorizing Aurora MySQL to access AWS services

Creating an IAM policy to access Amazon S3
