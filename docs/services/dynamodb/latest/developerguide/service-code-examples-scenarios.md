# Scenarios for DynamoDB using AWS SDKs

The following code examples show you how to implement common scenarios in DynamoDB
with AWS SDKs. These scenarios show you how to accomplish specific tasks by calling multiple functions
within DynamoDB or combined with other AWS services.
Each scenario includes a link to the complete source code, where you can find instructions on how to set up and run the code.

Scenarios target an intermediate level of experience to help you understand service actions in context.

###### Examples

- [Accelerate reads with DAX](example-dynamodb-usage-daxdemo-section.md)

- [Advanced Global Secondary Index scenarios](example-dynamodb-scenario-gsiadvanced-section.md)

- [Build an app to submit data to a DynamoDB table](example-cross-submitdataapp-section.md)

- [Compare multiple values with a single attribute](example-dynamodb-scenario-comparemultiplevalues-section.md)

- [Conditionally update an item's TTL](example-dynamodb-updateitemconditionalttl-section.md)

- [Connect to a local instance](example-dynamodb-local-section.md)

- [Count expression operators](example-dynamodb-scenario-expressionoperatorcounting-section.md)

- [Create a REST API to track COVID-19 data](example-cross-apigatewaydatatracker-section.md)

- [Create a messenger application](example-cross-stepfunctionsmessenger-section.md)

- [Create a serverless application to manage photos](example-cross-pam-section.md)

- [Create a table with global secondary index](example-dynamodb-createtablewithglobalsecondaryindex-section.md)

- [Create a table with warm throughput enabled](example-dynamodb-createtablewarmthroughput-section.md)

- [Create a web application to track DynamoDB data](example-cross-dynamodbdatatracker-section.md)

- [Create a websocket chat application](example-cross-apigatewaywebsocketchat-section.md)

- [Create an item with a TTL](example-dynamodb-putitemttl-section.md)

- [Create and manage MRSC global tables](example-dynamodb-scenario-mrscglobaltables-section.md)

- [Create and manage global tables demonstrating MREC](example-dynamodb-scenario-globaltableoperations-section.md)

- [Delete data using PartiQL DELETE](example-dynamodb-partiqldelete-section.md)

- [Detect PPE in images](example-cross-rekognitionphotoanalyzerppe-section.md)

- [Insert data using PartiQL INSERT](example-dynamodb-partiqlinsert-section.md)

- [Invoke a Lambda function from a browser](example-cross-lambdaforbrowser-section.md)

- [Manage Global Secondary Indexes](example-dynamodb-scenario-gsilifecycle-section.md)

- [Manage resource-based policies](example-dynamodb-scenario-resourcepolicylifecycle-section.md)

- [Monitor DynamoDB performance](example-cross-monitordynamodb-section.md)

- [Perform advanced query operations](example-dynamodb-scenario-advancedquerytechniques-section.md)

- [Perform list operations](example-dynamodb-scenario-listoperations-section.md)

- [Perform map operations](example-dynamodb-scenario-mapoperations-section.md)

- [Perform set operations](example-dynamodb-scenario-setoperations-section.md)

- [Query a table by using batches of PartiQL statements](example-dynamodb-scenario-partiqlbatch-section.md)

- [Query a table using PartiQL](example-dynamodb-scenario-partiqlsingle-section.md)

- [Query a table using a Global Secondary Index](example-dynamodb-scenarios-querywithglobalsecondaryindex-section.md)

- [Query a table using a begins\_with condition](example-dynamodb-scenarios-querywithbeginswithcondition-section.md)

- [Query a table using a date range](example-dynamodb-scenarios-querywithdaterange-section.md)

- [Query a table with a complex filter expression](example-dynamodb-scenarios-querywithcomplexfilter-section.md)

- [Query a table with a dynamic filter expression](example-dynamodb-scenarios-querywithdynamicfilter-section.md)

- [Query a table with a filter expression and limit](example-dynamodb-scenarios-querywithfilterandlimit-section.md)

- [Query a table with nested attributes](example-dynamodb-scenarios-querywithnestedattributes-section.md)

- [Query a table with pagination](example-dynamodb-scenarios-querywithpagination-section.md)

- [Query a table with strongly consistent reads](example-dynamodb-scenarios-querywithstronglyconsistentreads-section.md)

- [Query data using PartiQL SELECT](example-dynamodb-partiqlselect-section.md)

- [Query for TTL items](example-dynamodb-queryfilteredttl-section.md)

- [Query tables using date and time patterns](example-dynamodb-scenario-datetimequeries-section.md)

- [Save EXIF and other image information](example-cross-detectlabels-section.md)

- [Set up Attribute-Based Access Control](example-dynamodb-scenario-abacsetup-section.md)

- [Understand update expression order](example-dynamodb-scenario-updateexpressionorder-section.md)

- [Update a table's warm throughput setting](example-dynamodb-updatetablewarmthroughput-section.md)

- [Update an item's TTL](example-dynamodb-updateitemttl-section.md)

- [Update data using PartiQL UPDATE](example-dynamodb-partiqlupdate-section.md)

- [Use API Gateway to invoke a Lambda function](example-cross-lambdaapigateway-section.md)

- [Use Step Functions to invoke Lambda functions](example-cross-serverlessworkflows-section.md)

- [Use a document model](example-dynamodb-midlevelinterface-section.md)

- [Use a high-level object persistence model](example-dynamodb-highlevelinterface-section.md)

- [Use atomic counter operations](example-dynamodb-scenario-atomiccounteroperations-section.md)

- [Use conditional operations](example-dynamodb-scenario-conditionaloperations-section.md)

- [Use expression attribute names](example-dynamodb-scenario-expressionattributenames-section.md)

- [Use scheduled events to invoke a Lambda function](example-cross-lambdascheduledevents-section.md)

- [Work with Local Secondary Indexes](example-dynamodb-scenario-lsiexamples-section.md)

- [Work with Streams and Time-to-Live](example-dynamodb-scenario-streamsandttl-section.md)

- [Work with global tables and multi-Region replication eventual consistency (MREC)](example-dynamodb-scenario-multiregionreplication-section.md)

- [Work with resource tagging](example-dynamodb-scenario-taggingexamples-section.md)

- [Work with table encryption](example-dynamodb-scenario-encryptionexamples-section.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateTimeToLive

Accelerate reads with DAX

All content copied from https://docs.aws.amazon.com/.
