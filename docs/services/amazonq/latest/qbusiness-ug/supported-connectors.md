# Connecting Amazon Q Business data source connectors

###### Note

Before you connect an Amazon Q Business data source to your application, you [must add an index and retriever](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/select-retriever.html) to it.

To connect a data source to your Amazon Q Business application environment, you can use
the AWS Management Console or the [CreateDataSource](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_CreateDataSource.html) API operation.

By using the `CreateDataSource` API operation, you can configure tags, sync
run schedules, and configure Amazon VPC settings. Then, you can use the
`configuration` parameter to provide all other configuration information
specific to your data source connector.

###### Note

This procedure is available if you chose to [add\
a new Amazon Q Business index](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/native-retriever.html) to your application environment.

This section contains an overview of data source connector features, recommended best
practices for configuration, and configuration information specific to your data source
connector.

###### Topics

- [Data source connector concepts](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connector-concepts.html)

- [What is a document?](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connector-doc-crawl.html)

- [Best practices for data source connector configuration in Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connector-best-practices.html)

- [Supported connectors](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connectors-list.html)

- [Understanding Amazon Q Business User Store](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connector-principal-store.html)

- [Using Amazon VPC with Amazon Q Business connectors](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/connector-vpc.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Uploading files

Concepts
