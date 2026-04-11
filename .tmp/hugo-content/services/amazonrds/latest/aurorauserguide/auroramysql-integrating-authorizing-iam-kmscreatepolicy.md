# Creating an IAM policy to access AWS KMS resources

Aurora can access the AWS KMS keys
used for encrypting their database backups.
However, you must first create an IAM policy that provides the
permissions that allow Aurora to access KMS keys.

The following policy adds the permissions required by Aurora to access KMS keys on
your behalf.

JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Sid": "AllowAuroraToAccessKey",
      "Effect": "Allow",
      "Action": [
        "kms:Decrypt"
      ],
      "Resource": "arn:aws:kms:us-east-1:123456789012:key/key-ID"
    }
  ]
}

```

You can use the following steps to create an IAM policy that provides the
minimum required permissions for Aurora to access KMS keys on your behalf.

###### To create an IAM policy to grant access to your KMS keys

01. Open the [IAM\
     console](https://console.aws.amazon.com/iam/home?).

02. In the navigation pane, choose **Policies**.

03. Choose **Create policy**.

04. On the **Visual editor** tab, choose **Choose**
    **a service**, and then choose **KMS**.

05. For **Actions**, choose **Write**, and then choose
     **Decrypt**.

06. Choose **Resources**, and choose **Add ARN**.

07. In the **Add ARN(s)** dialog box, enter the following values:

- **Region** – Type the AWS Region, such as `us-west-2`.

- **Account** – Type the user account number.

- **Log Stream Name** – Type the KMS key identifier.

08. In the **Add ARN(s)** dialog box, choose **Add**.

09. Choose **Review policy**.

10. Set **Name** to a name for your IAM policy, for
     example `AmazonRDSKMSKey`. You use this name when you
     create an IAM role to associate with your Aurora DB cluster. You can also add
     an optional **Description** value.

11. Choose **Create policy**.

12. Complete the steps in [Creating an IAM role to allow Amazon Aurora to access AWS services](auroramysql-integrating-authorizing-iam-createrole.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating an IAM policy to access CloudWatch Logs

Creating an IAM role to access AWS services

All content copied from https://docs.aws.amazon.com/.
