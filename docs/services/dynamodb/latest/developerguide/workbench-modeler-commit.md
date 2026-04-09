# Committing a data model to DynamoDB

When you are satisfied with your data model, you can commit the model to
Amazon DynamoDB.

###### Note

- This action creates server-side resources in AWS for the tables and global secondary indexes represented in the data model.

- NoSQL Workbench creates tables and indexes with on-demand capacity by default.

###### To commit the data model to DynamoDB

1. Open NoSQL Workbench, and on the main screen, click on the name of the model that you want to commit.

2. In the top bar, click **Commit**.

3. Choose an existing connection, or create a new connection by clicking the **Add new connection** button.

- To add a new connection, specify the following information:

- **Account Alias**

- **AWS Region**

- **Access key ID**

- **Secret access key**

For more information about how to obtain the access keys, see [Getting an AWS access key](settingup-dynamowebservice.md#SettingUp.DynamoWebService.GetCredentials).

- You can optionally specify the following:

- [**Session token**](../../../iam/latest/userguide/id-credentials-temp-use-resources.md)

- [**IAM role ARN**](../../../iam/latest/userguide/reference-identifiers.md#identifiers-arns)

4. If you prefer to use [DynamoDB local](dynamodblocal.md):

1. Choose the **Local connection** tab.

2. Click the **Add new connection** button.

3. Specify the **Connection name** and **Port**.

###### Note

To use DynamoDB local, you will need to turn it on by using the **DynamoDB local** toggle at the bottom left of the NoSQL Workbench screen.

5. Click **Commit**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Exporting a model

Operation builder

All content copied from https://docs.aws.amazon.com/.
