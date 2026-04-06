# Connecting Gmail to Amazon Q Business

With Amazon Q Business, you can connect your Gmail enterprise email
system to unlock valuable organizational knowledge stored in email communications. When you
connect Gmail to Amazon Q Business, your users can search and get
answers from email content and conversations directly through the Amazon Q web
experience.

You can connect your Gmail instance to Amazon Q Business using either the AWS Management Console or the [CreateDataSource](../api-reference/api-createdatasource.md) API. This connection enables your organization to leverage email-based knowledge for improved decision-making and faster information discovery.

###### Topics

- [Gmail connector versions](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-versions.html)

- [Gmail connector overview](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-overview.html)

- [Prerequisites for connecting Amazon Q Business to Gmail](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-prereqs.html)

- [Connecting Amazon Q Business to Gmail using the latest connector (Console)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-console-new.html)

- [Connecting Amazon Q Business to Gmail using the legacy connector (Console)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-console-original.html)

- [Connecting Amazon Q Business to Gmail using the new connector (API)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-new-api.html)

- [Connecting Amazon Q Business to Gmail using the original connector (API)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-original-api.html)

- [How Amazon Q Business connector crawls Gmail ACLs](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-user-management.html)

- [Gmail data source connector field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-field-mappings.html)

- [IAM role for Amazon Q Business Gmail connector](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-iam-role.html)

- [Understand error codes in the Amazon Q Business Gmail connector](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/gmail-error-codes.html)

**Learn more**

- For an overview of the Amazon Q web experience creation process using IAM Identity Center, see [Configuring an application using IAM Identity Center](create-application.md).

- For an overview of the Amazon Q web experience creation process using AWS Identity and Access Management, see [Configuring an application using IAM](create-application-iam.md).

- For an overview of connector features, see [Data source connector concepts](connector-concepts.md).

- For information about connector configuration best practices, see [Connector configuration best practices](connector-best-practices.md).

###### Note

**ACL behavior by connector version:**

- **New connector:** ACL and identity crawling is automatically enabled and cannot be disabled. No manual configuration is required.

- **Original connector:** ACL and identity crawling can be manually configured during setup.

###### Note

**Original connector only:** Field mappings are only available when using the original Gmail connector. The new connector uses optimized default field mappings that cannot be customized.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM role

Connector versions
