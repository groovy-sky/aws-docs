# Amazon MQ for RabbitMQ maximum resource limit

You can configure resource limits up to the maximum values shown in the following tables. To learn how to update resource limits for your broker, see [Resource Limit Configuration](configure-resource-limits.md).

## Sizing guidelines for m7g with quorum queues for single instance deployment

The following table shows the **maximum** limit values for each instance type for single instance brokers.

Instance TypeConnectionsChannelsConsumers per channelQueuesVhostsShovelsExchangesMessage Size in Bytesmq.m7g.medium3009001,0002,5001015012500134217728mq.m7g.large5,00015,0001,00020,0001500250100,000134217728mq.m7g.xlarge10,00030,0001,00030,0001,500500150,000134217728mq.m7g.2xlarge20,00060,0001,00040,0001,5001,000200,000134217728mq.m7g.4xlarge40,000120,0001,00060,0001,5002000300,000134217728mq.m7g.8xlarge80,000240,0001,00080,0001,5004000400,000134217728mq.m7g.12xlarge120,000360,0001,000100,0001,5006,000500,000134217728mq.m7g.16xlarge160,000480,0001,000120,0001,5008,000600,000134217728

## Sizing guidelines for m7g with quorum queues for cluster deployment

The following table shows the **maximum** limit values for each instance type for cluster brokers.

Instance TypeConnections per NodeChannels per NodeConsumers per channelQueuesVhostsShovelsExchangesMessage Size in Bytesmq.m7g.medium3009001,0005001050500134217728mq.m7g.large5,00015,0001,00010,0001,50015050,000134217728mq.m7g.xlarge10,00030,0001,00015,0001,50030075,000134217728mq.m7g.2xlarge20,00060,0001,00020,0001,500600100,000134217728mq.m7g.4xlarge40,000120,0001,00030,0001,5001200150,000134217728mq.m7g.8xlarge80,000240,0001,00040,0001,5002,400200,000134217728mq.m7g.12xlarge120,000360,0001,00050,0001,5003,600250,000134217728mq.m7g.16xlarge160,000480,0001,00060,0001,5004,800300,000134217728

The following table shows the **maximum** limit values for each instance type for single instance brokers.

Instance TypeConnectionsChannelsConsumers per channelQueuesVhostsShovelsm5.large5,00015,0001,00030,0001500250m5.xlarge10,00030,0001,00060,0001500500m5.2xlarge20,00060,0001,000120,00015001,000m5.4xlarge40,000120,0001000240,0001,0002,000

The following table shows the **maximum** limit values for each instance type for cluster brokers.

Instance TypeQueuesConsumers per channelShovelsm5.large10,0001,000150m5.xlarge15,0001,000300m5.2xlarge20,0001,000600m5.4xlarge30,0001,0001200

The following connection and channel limits are applied per node:

Instance TypeConnectionsChannelsm5.large500015,000m5.xlarge10,00030,000m5.2xlarge20,00060,000m5.4xlarge40,000120,000

The exact limit values for a cluster broker
may be lower than the indicated value depending on the number of available nodes
and how RabbitMQ distributes resources among the available nodes.
If you exceed the limit values, you can create a new connection to a different node and try again,
or you can upgrade the instance size to increase the maximum limits

## Error messages

The following error messages are returned when limits are exceeded.
All values are based on the `m7.large` single instance limits.

###### Note

The error codes for the following messages may change based on the client library you are using.

**Connection**

`ConnectionClosedByBroker 500 "NOT_ALLOWED - connection refused: node connection limit (5000) is reached"`

**Channel**

`ConnectionClosedByBroker 1500 "NOT_ALLOWED - number of channels opened on node
             'rabbit@ip-10-0-23-173.us-west-2.compute.internal' has reached the maximum allowed limit of (15,000)"`

**Consumer**

`ConnectionClosedByBroker: (530, 'NOT_ALLOWED - reached maximum (1,000) of consumers per channel')`

**Maximum message size**

`
                    (406, 'PRECONDITION_FAILED - message size 524289 is larger than configured max size 524288')
                `

**Exchange**

`
                    (406, "PRECONDITION_FAILED - cannot declare exchange 'limit_test_3' in vhost '/': exchange limit of 10 is reached")
                `

###### Note

The following error messages use the HTTP Management API format.

**Queue**

`{"error":"bad_request","reason":"cannot declare queue 'my_queue': queue limit in cluster (10,000) is reached"}]`

**Shovel**

`{"error":"bad_request","reason":"Validation failed\n\ncomponent shovel is limited to 150 per node\n"}`

**Vhost**

`{"error":"bad_request","reason":"cannot create vhost 'my_vhost': vhost limit of 1500 is reached"}`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Default resource limits

Broker defaults

All content copied from https://docs.aws.amazon.com/.
