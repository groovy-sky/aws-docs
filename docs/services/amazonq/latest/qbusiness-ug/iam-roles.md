# IAM roles for Amazon Q Business

When you create an application or a web experience with Amazon Q Business, or
connect a data source to it, Amazon Q Business needs access to the required AWS resources.

If you use the AWS CLI or an AWS SDK, you must create an AWS Identity and Access Management (IAM)
policy before you create the Amazon Q Business resource. When you call an API
operation, you provide the Amazon Resource Name (ARN) role with the policy attached.

If you use the AWS Management Console, you can create a new IAM role in the Amazon Q console or use an existing IAM role. The console displays
roles that have the string **qbusiness** or **QBusiness**
in the role name.

To learn more about IAM roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the _AWS Identity and Access Management User_
_Guide_.

The following topics provide details for the required policies. If you create IAM roles using the Amazon Q Business console, these policies are created
on your behalf, unless otherwise noted.

###### Topics

- [IAM role for an Amazon Q Business application](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-application-iam-role.html)

- [IAM role for an Amazon Q Business web experience](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/deploy-experience-iam-role.html)

- [IAM role for Amazon Q Business data source connectors](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/iam-roles-ds.html)

- [IAM role for Amazon Q Business plugins](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/plugin-iam-role.html)

- [IAM roles for custom document enrichment in Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cde-iam-roles.html)

- [IAM role for an Amazon Kendra retriever](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/kendra-retriever-iam-role.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting up

Amazon Q Business application
