# Creating an IAM role to allow Amazon Aurora to access AWS services

After creating an IAM policy to allow Aurora to access AWS resources, you must
create an IAM role and attach the IAM policy to the new IAM role.

To create an IAM role to permit your Amazon RDS cluster to communicate with other AWS
services on your behalf, take the following steps.

###### To create an IAM role to allow Amazon RDS to access AWS services

01. Open the [IAM\
     console](https://console.aws.amazon.com/iam/home?).

02. In the navigation pane, choose **Roles**.

03. Choose **Create role**.

04. Under **AWS service**, choose **RDS**.

05. Under **Select your use case**, choose **RDS – Add Role to Database**.

06. Choose **Next**.

07. On the **Permissions policies** page, enter the name of your policy
     in the **Search** field.

08. When it appears in the list, select the policy that you defined earlier using the instructions
     in one of the following sections:

- [Creating an IAM policy to access Amazon S3 resources](auroramysql-integrating-authorizing-iam-s3createpolicy.md)

- [Creating an IAM policy to access AWS Lambda resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.LambdaCreatePolicy.html)

- [Creating an IAM policy to access CloudWatch Logs resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.CWCreatePolicy.html)

- [Creating an IAM policy to access AWS KMS resources](auroramysql-integrating-authorizing-iam-kmscreatepolicy.md)

09. Choose **Next**.

10. In **Role name**, enter a name for your IAM role, for
     example `RDSLoadFromS3`. You can also add
     an optional **Description** value.

11. Choose **Create Role**.

12. Complete the steps in [Associating an IAM role with an Amazon Aurora MySQL DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.AddRoleToDBCluster.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating an IAM policy to access AWS KMS

Associating an IAM role with a DB cluster
