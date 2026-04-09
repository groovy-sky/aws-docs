# Setting up DynamoDB local (downloadable version)

With the downloadable version of Amazon DynamoDB, you can develop and test applications
without accessing the DynamoDB web service. Instead, the database is self-contained on your
computer. When you're ready to deploy your application in production, you remove the
local endpoint in the code, and then it points to the DynamoDB web service.

Having this local version helps you save on throughput, data storage, and data
transfer fees. In addition, you don't need an internet connection while you develop your
application.

DynamoDB local is available as a [download](dynamodblocal-downloadingandrunning.md#DynamoDBLocal.DownloadingAndRunning.title) (requires JRE),
as an [Apache Maven dependency](dynamodblocal-downloadingandrunning.md#apache-maven), or as a [Docker image](dynamodblocal-downloadingandrunning.md#docker).

If you prefer to use the Amazon DynamoDB web service instead, see [Setting up DynamoDB (web service)](settingup-dynamowebservice.md).

###### Topics

- [Deploying DynamoDB locally on your computer](dynamodblocal-downloadingandrunning.md)

- [DynamoDB local usage notes](dynamodblocal-usagenotes.md)

- [Release history for DynamoDB local](dynamodblocalhistory.md)

- [Telemetry in DynamoDB local](dynamodblocaltelemetry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting up DynamoDB (web service)

Deploying

All content copied from https://docs.aws.amazon.com/.
