# Data sources

In the previous section, we learned that a schema defines the shape of your data. However, we never
explained where that data came from. In real projects, your schema is like a gateway that handles all
requests made to the server. When a request is made, the schema acts as the single endpoint that interfaces
with the client. The schema will access, process, and relay data from the data source back to the client.
See the infographic below:

![GraphQL schema integrating multiple AWS services for a single endpoint API architecture.](https://docs.aws.amazon.com/images/appsync/latest/devguide/images/aws-flow-infographic.png)

AWS AppSync and GraphQL superbly implement Backend For Frontend (BFF) solutions. They work in tandem to reduce
complexity at scale by abstracting the backend. If your service uses different data sources and/or
microservices, you can essentially abstract some of the complexity away by defining the shape of the data of
each source (subgraph) in a single schema (supergraph). This means your GraphQL API is not limited to using
one data source. You can associate any number of data sources with your GraphQL API and specify in your code
how they will interact with the service.

As you can see in the infographic, the GraphQL schema contains all of the information clients need to
request data. This means everything can be processed in a single request rather than multiple requests as is
the case with REST. These requests go through the schema, which is the sole endpoint of the service. When
requests are processed, a resolver (explained in the next section) executes its code to process the data
from the relevant data source. When the response is returned, the subgraph tied to the data source will be
populated with the data in the schema.

AWS AppSync supports many different data source types. In the table below, we'll describe each type, list some
of the benefits of each, and provide useful links for additional context.

Data sourceDescriptionBenefitsSupplemental informationAmazon DynamoDB"Amazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable
performance with seamless scalability. DynamoDB lets you offload the administrative burdens of
operating and scaling a distributed database so that you don't have to worry about hardware
provisioning, setup and configuration, replication, software patching, or cluster scaling. DynamoDB
also offers encryption at rest, which eliminates the operational burden and complexity involved
in protecting sensitive data."

- **Performance at scale**: DynamoDB is designed around
consistent performance at any scale. This is possible through the use of partitions.
DynamoDB will automatically partition your tables into several allocations that will be
stored in multiple SSDs across several nodes. This will generally increase network
throughput and reduce latency.

- **Capacity at scale**: DynamoDB monitors your traffic and
allows you to automatically scale your throughput if the network remains overloaded for
extended periods of time.

- **Availability and fault tolerance**: DynamoDB is
supported by several physically isolated Regions, each containing several physically
isolated Availability Zones. DynamoDB will automatically switch to a backup zone in the
event of a service disruption. You can also back up and replicate your data manually
for data assurance.

- **Logging and monitoring**: DynamoDB provides several
analytical tools for your tables. You can monitor your table's performance and create
alarms to notify you of drastic changes to the service.

- **Security**: DynamoDB follows strict protocols to ensure
your data complies with your organization's security requirements.

- **Integration with AWS AppSync**: DynamoDB is seamlessly
integrated with our service. You can create new DynamoDB tables and automatically generate
a schema from them to streamline your development process. We also provide an entire
collection of operations to easily request data from existing DynamoDB tables in your
account in your resolver.

- [DynamoDB official documentation](../../../dynamodb/latest/developerguide/introduction.md)

- [Partitions](../../../dynamodb/latest/developerguide/howitworks-partitions.md)

- [Auto scaling](../../../dynamodb/latest/developerguide/autoscaling.md)

- [Fault tolerance](../../../dynamodb/latest/developerguide/disaster-recovery-resiliency.md)

- [Monitoring](../../../dynamodb/latest/developerguide/monitoring.md)

- [Security](../../../dynamodb/latest/developerguide/security.md)

- [GraphQL\
and DynamoDB](https://aws.amazon.com/graphql/graphql-dynamodb-data-modeling)

- [Resolver\
operations for DynamoDB](js-resolver-reference-dynamodb.md)

- [Pricing model](https://aws.amazon.com/dynamodb/pricing)

AWS Lambda"AWS Lambda is a compute service that lets you run code without provisioning or managing
servers.

Lambda runs your code on a high-availability compute infrastructure and performs
all of the administration of the compute resources, including server and operating system
maintenance, capacity provisioning and automatic scaling, and logging. With Lambda, all you
need to do is supply your code in one of the language runtimes that Lambda
supports."

- **Pay-as-you-use model**: Lambda only charges you when
you use their resources. They also allow you to scale the amount of resources used with
your application needs.

- **Automatic scaling**: Sometimes your application may
require extra computing power for a particular process. Lambda allows you to
automatically scale computing resources to fit the needs of your application.

- **Faster deployment times**: You can streamline your
development process via a deployment package. Use a package to upload your function
code to the Lambda service. You can then use their runtime environments to test and
execute your functions.

- **Versatility**: Lambda can be used in a multitude of
use cases. You can seamlessly integrate Lambda with third-party services and AWS
services alike. Some examples include [CI/CD pipelines](../../../codepipeline/latest/userguide/actions-invoke-lambda-function.md) and [mass mailing services](../../../ses/latest/dg/receiving-email-action-lambda-example-functions.md).

- **Integration with AWS AppSync**: You can easily
invoke your Lambda functions in your resolver to handle requests. Our service provides a
streamlined request operation to perform Lambda calls. We allow both single and batched
calls.

- [Official\
documentation](../../../lambda/latest/dg/welcome.md)

- [Scaling](../../../lambda/latest/dg/lambda-concurrency.md)

- [deployment](../../../lambda/latest/dg/gettingstarted-package.md)

- [runtimes](../../../lambda/latest/dg/lambda-runtimes.md)

- [Lambda resolver\
tutorial](tutorial-lambda-resolvers-js.md)

- [Pricing model](https://aws.amazon.com/lambda/pricing)

OpenSearch"Amazon OpenSearch Service is a managed service that makes it easy to deploy, operate, and scale
OpenSearch clusters in the AWS Cloud. Amazon OpenSearch Service supports OpenSearch and legacy Elasticsearch
OSS (up to 7.10, the final open-source version of the software). When you create a cluster, you
have the option of which search engine to use.

**OpenSearch** is a fully open-source search and analytics engine for use cases
such as log analytics, real-time application monitoring, and clickstream analysis. For more
information, see the [OpenSearch\
documentation](https://opensearch.org/docs).

**Amazon OpenSearch Service** provisions
all the resources for your OpenSearch cluster and launches it. It also automatically
detects and replaces failed OpenSearch Service nodes, reducing the overhead associated with
self-managed infrastructures. You can scale your cluster with a single API call or a few
clicks in the console."

- **Scaling**: You can easily scale the service to fit
your service requirements through OpenSearch Serverless.

- **Data ingestion**: You can use OpenSearch Ingestion
to import, process, and analyze data. There are many applications for data ingestion,
which you can find [here](../../../opensearch-service/latest/developerguide/use-cases-overview.md).

- **Security**: OpenSearch can manage your AWS security
configuration including IAM, CloudTrail, VPCs, authentication, etc.

- **Availability**: OpenSearch also supports different
Regions and Availability Zones in its service.

- **Integration with AWS AppSync**: In AWS AppSync,
you can use GraphQL APIs to store and retrieve data from existing OpenSearch Service
domains in your account.

- [Official documentation](../../../opensearch-service/latest/developerguide/what-is.md)

- [Serverless](../../../opensearch-service/latest/developerguide/serverless.md)

- [Pricing\
model](https://aws.amazon.com/opensearch-service/pricing)

HTTP endpointsYou can use HTTP endpoints as data sources. AWS AppSync can send requests to the endpoints
with the relevant information like params and payload. The HTTP response will be exposed to the
resolver, which will return the final response after it finishes its operation(s).

- Useful for simple applications that aren't as integrated with services like
Lambda.

- [Resolver\
reference](resolver-reference-http-js.md)

Amazon EventBridge"EventBridge is a serverless service that uses events to connect application components together,
making it easier for you to build scalable event-driven applications. Use it to route events
from sources such as home-grown applications, AWS services, and third-party software to
consumer applications across your organization. EventBridge provides a simple and consistent way to
ingest, filter, transform, and deliver events so you can build new applications
quickly."

- **Event-driven architecture**: You can take advantage
of [event-driven\
architecture](https://aws.amazon.com/event-driven-architecture).

- **Scheduling**: You can use the EventBridge Scheduler to
automate your tasks and rules using cron expressions or set time intervals as an
alternative to event patterns.

- **Pipes**: Using EventBridge Pipes, you can replace the event
bus with a pipe that includes additional filtering event patterns and enrichment
through data transforms before sending the event to the target.

- **Integration with AWS AppSync**: AWS AppSync
allows you to send events to event buses using your resolver.

- [Official documentation](../../../eventbridge/latest/userguide/eb-what-is.md)

- [Pipes](../../../eventbridge/latest/userguide/eb-pipes.md)

- [Scheduler](../../../eventbridge/latest/userguide/scheduler.md)

- [Resolver\
reference](resolver-reference-eventbridge-js.md)

- [Pricing\
model](https://aws.amazon.com/eventbridge/pricing)

Relational databases"Amazon Relational Database Service (Amazon RDS) is a web service that makes it easier to set up,
operate, and scale a relational database in the AWS Cloud. It provides cost-efficient,
resizable capacity for an industry-standard relational database and manages common database
administration tasks."

- **Managing made easy**: Periodically, RDS performs
maintenance on its resources. Maintenance most often involves updates to the DB
instance's underlying hardware, underlying operating system (OS), or database engine
version. Under normal circumstances, you can decide when to perform updates (exceptions
include security patches).

- **Recommendations**: RDS' recommendation feature
provides automated suggestions for fixing potential issues in your instance.

- **Availability**: RDS is available in different
physical Regions across the world. You can easily distribute your database needs across
different nodes to provide better service to your customers.

- **Customisation**: RDS is tailored to meet the
requirements of large corporations. RDS provides various options for computing, quick
deployment, scalability, and storage.

- **Security**: RDS is integrated with several tools and
services to maintain database security on the user, database, and network
levels.

- **Integration with AWS AppSync**: If you're looking
for a mature backend solution, AWS AppSync allows you to send, process, store, and
return data using your instance as the data source.

- [Official documentation](../../../amazonrds/latest/userguide/welcome.md)

- [Features](../../../amazonrds/latest/userguide/concepts-rdsfeaturesregionsdbengines-grids.md)

- [Maintenance](../../../amazonrds/latest/userguide/user-upgradedbinstance-maintenance.md)

- [Recommendations](../../../amazonrds/latest/userguide/accessing-monitoring.md#USER_Recommendations)

- [Storage options](../../../amazonrds/latest/userguide/chap-storage.md)

- [Availability](../../../amazonrds/latest/userguide/concepts-regionsandavailabilityzones.md)

- [Security](../../../amazonrds/latest/userguide/usingwithrds.md)

- [Pricing model](https://aws.amazon.com/rds/pricing)

None data sourceIf you aren't planning on using a data source service, you can set it to `none`.
A `none` data source, while still explicitly categorized as a data source, isn't a
storage medium. Despite that, it's still useful in certain instances for data manipulation and
pass-throughs.

- Potentially useful for things like data conversion

- Useful when resolving something locally

- [Resolver reference](resolver-reference-none-js.md)

###### Tip

For more information about how data sources interact with AWS AppSync, see [Attaching a data source](attaching-a-data-source.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GraphQL fields

Resolvers

All content copied from https://docs.aws.amazon.com/.
